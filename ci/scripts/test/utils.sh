mount_storage() {
  # Configure cgroup
  mount -t tmpfs cgroup_root /sys/fs/cgroup
  mkdir -p /sys/fs/cgroup/devices
  mkdir -p /sys/fs/cgroup/memory

  mount -tcgroup -odevices cgroup:devices /sys/fs/cgroup/devices
  devices_mount_info=$(cat /proc/self/cgroup | grep devices)
  devices_subdir=$(echo $devices_mount_info | cut -d: -f3)
  echo 'b 7:* rwm' > /sys/fs/cgroup/devices/devices.allow
  echo 'b 7:* rwm' > /sys/fs/cgroup/devices${devices_subdir}/devices.allow

  mount -tcgroup -omemory cgroup:memory /sys/fs/cgroup/memory

  # Setup loop devices
  for i in {0..256}
  do
    mknod -m777 /dev/loop$i b 7 $i
  done

  # Make and Mount EXT4 Volume
  mkdir /mnt/ext4
  truncate -s 256M /ext4_volume
  mkfs.ext4 /ext4_volume
  mount /ext4_volume /mnt/ext4
  chmod 777 /mnt/ext4

  # Make BTRFS volume
  truncate -s 1G /btrfs_volume
  mkfs.btrfs --nodesize 4k -s 4k /btrfs_volume

  # Mount BTRFS
  mkdir /mnt/btrfs
  mount -t btrfs -o user_subvol_rm_allowed,rw /btrfs_volume /mnt/btrfs
  chmod 777 -R /mnt/btrfs
  btrfs quota enable /mnt/btrfs

  for i in {1..9}
  do
    # Make XFS Volume
    truncate -s 1G /xfs_volume_${i}
    mkfs.xfs -b size=4096 /xfs_volume_${i}

    # Mount XFS
    mkdir /mnt/xfs-${i}
    mount -t xfs -o pquota /xfs_volume_${i} /mnt/xfs-${i}
    chmod 777 -R /mnt/xfs-${i}
  done

  # Make and mount another xfs store for init-store tests
  truncate -s 1G /init_store_xfs_volume
  mkfs.xfs -b size=4096 /init_store_xfs_volume
  mkdir /opt/init_store_overlay-xfs
  mount -t xfs -o pquota /init_store_xfs_volume /opt/init_store_overlay-xfs
  chmod 777 -R /opt/init_store_overlay-xfs

  # Make and mount another btrfs store for init-store tests
  truncate -s 1G /init_store_btrfs_volume
  mkfs.btrfs --nodesize 4k -s 4k /init_store_btrfs_volume
  mkdir /opt/init_store_btrfs
  mount -t btrfs -o user_subvol_rm_allowed,rw /init_store_btrfs_volume /opt/init_store_btrfs
  chmod 777 -R /opt/init_store_btrfs
  btrfs quota enable /opt/init_store_btrfs
}

sudo_mount_storage() {
  local MOUNT_STORAGE_FUNC=$(declare -f mount_storage)
  sudo bash -c "$MOUNT_STORAGE_FUNC; mount_storage"
}

move_to_gopath() {
  thing_i_want_moved=$1
  dest_path=/go/src/code.cloudfoundry.org/${thing_i_want_moved}

  # remove the original grootfs package path
  [ -d $dest_path ] && rmdir $dest_path

  # link the uploaded source (from build) to the GOPATH
  ln -s $PWD/src/code.cloudfoundry.org/${thing_i_want_moved} $dest_path

  # because the uploaded source is owned by the user that runs fly, we need
  # to chown
  pushd $dest_path
    sudo chown -R groot:groot .
  popd

  echo $dest_path
}

install_dependencies() {
  if ! [ -d vendor ]; then
    glide install
  fi
}

setup_drax() {
  drax_path=$1
  cp $drax_path /usr/local/bin/drax
  chown root:root /usr/local/bin/drax
  chmod u+s /usr/local/bin/drax
}

sudo_setup_drax() {
  drax_path=$1

  local SETUP_DRAX_FUNC=$(declare -f setup_drax)
  sudo bash -c "$SETUP_DRAX_FUNC; setup_drax $drax_path"
}
