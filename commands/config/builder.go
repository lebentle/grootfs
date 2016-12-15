package config

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	BaseStorePath             string   `yaml:"store_path"`
	CleanThresholdBytes       uint64   `yaml:"clean_threshold_bytes"`
	DiskLimitSizeBytes        int64    `yaml:"disk_limit_size_bytes"`
	DraxBin                   string   `yaml:"drax_bin"`
	ExcludeBaseImageFromQuota bool     `yaml:"exclude_base_image_from_quota"`
	GIDMappings               []string `yaml:"gid_mappings"`
	UIDMappings               []string `yaml:"uid_mappings"`
	IgnoreBaseImages          []string `yaml:"ignore_base_images"`
	InsecureRegistries        []string `yaml:"insecure_registries"`
	MetronEndpoint            string   `yaml:"metron_endpoint"`
	UserBasedStorePath        string
}

type Builder struct {
	config *Config
}

func NewBuilder(pathToYaml string) (*Builder, error) {
	config := Config{}

	if pathToYaml != "" {
		var err error
		config, err = load(pathToYaml)
		if err != nil {
			return nil, err
		}
	}

	b := &Builder{
		config: &config,
	}

	return b.WithStorePath(config.BaseStorePath, ""), nil
}

func (b *Builder) Build() (Config, error) {
	if b.config.DiskLimitSizeBytes < 0 {
		return *b.config, errors.New("invalid argument: disk limit cannot be negative")
	}

	return *b.config, nil
}

func (b *Builder) WithInsecureRegistries(insecureRegistries []string) *Builder {
	if insecureRegistries == nil || len(insecureRegistries) == 0 {
		return b
	}

	b.config.InsecureRegistries = insecureRegistries
	return b
}

func (b *Builder) WithIgnoreBaseImages(ignoreBaseImages []string) *Builder {
	if ignoreBaseImages == nil || len(ignoreBaseImages) == 0 {
		return b
	}

	b.config.IgnoreBaseImages = ignoreBaseImages
	return b
}

func (b *Builder) WithStorePath(storePath, defaultStorePath string) *Builder {
	if b.config.BaseStorePath != "" && storePath == defaultStorePath {
		b.config.UserBasedStorePath = userBasedStorePath(b.config.BaseStorePath)
		return b
	}

	b.config.BaseStorePath = storePath
	b.config.UserBasedStorePath = userBasedStorePath(storePath)
	return b
}

func (b *Builder) WithDraxBin(draxBin, defaultDraxBin string) *Builder {
	if b.config.DraxBin != "" && draxBin == defaultDraxBin {
		return b
	}

	b.config.DraxBin = draxBin
	return b
}

func (b *Builder) WithMetronEndpoint(metronEndpoint string) *Builder {
	if metronEndpoint == "" {
		return b
	}

	b.config.MetronEndpoint = metronEndpoint
	return b
}

func (b *Builder) WithUIDMappings(uidMappings []string) *Builder {
	if uidMappings == nil || len(uidMappings) == 0 {
		return b
	}

	b.config.UIDMappings = uidMappings
	return b
}

func (b *Builder) WithGIDMappings(gidMappings []string) *Builder {
	if gidMappings == nil || len(gidMappings) == 0 {
		return b
	}

	b.config.GIDMappings = gidMappings
	return b
}

func (b *Builder) WithDiskLimitSizeBytes(limit int64) *Builder {
	b.config.DiskLimitSizeBytes = limit
	return b
}

func (b *Builder) WithExcludeBaseImageFromQuota(exclude bool) *Builder {
	b.config.ExcludeBaseImageFromQuota = exclude
	return b
}

func (b *Builder) WithCleanThresholdBytes(threshold uint64) *Builder {
	b.config.CleanThresholdBytes = threshold
	return b
}

func load(configPath string) (Config, error) {
	configContent, err := ioutil.ReadFile(configPath)
	if err != nil {
		return Config{}, fmt.Errorf("invalid config path: %s", err)
	}

	var config Config
	err = yaml.Unmarshal(configContent, &config)
	if err != nil {
		return Config{}, fmt.Errorf("invalid config file: %s", err)
	}

	return config, nil
}

func userBasedStorePath(storePath string) string {
	userID := os.Getuid()
	return filepath.Join(storePath, strconv.Itoa(userID))
}