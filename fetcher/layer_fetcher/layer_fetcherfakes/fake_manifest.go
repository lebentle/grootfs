// Code generated by counterfeiter. DO NOT EDIT.
package layer_fetcherfakes

import (
	"context"
	"sync"

	"code.cloudfoundry.org/grootfs/fetcher/layer_fetcher"
	"github.com/containers/image/docker/reference"
	"github.com/containers/image/types"
	"github.com/opencontainers/image-spec/specs-go/v1"
)

type FakeManifest struct {
	ReferenceStub        func() types.ImageReference
	referenceMutex       sync.RWMutex
	referenceArgsForCall []struct{}
	referenceReturns     struct {
		result1 types.ImageReference
	}
	referenceReturnsOnCall map[int]struct {
		result1 types.ImageReference
	}
	ManifestStub        func(ctx context.Context) ([]byte, string, error)
	manifestMutex       sync.RWMutex
	manifestArgsForCall []struct {
		ctx context.Context
	}
	manifestReturns struct {
		result1 []byte
		result2 string
		result3 error
	}
	manifestReturnsOnCall map[int]struct {
		result1 []byte
		result2 string
		result3 error
	}
	SignaturesStub        func(ctx context.Context) ([][]byte, error)
	signaturesMutex       sync.RWMutex
	signaturesArgsForCall []struct {
		ctx context.Context
	}
	signaturesReturns struct {
		result1 [][]byte
		result2 error
	}
	signaturesReturnsOnCall map[int]struct {
		result1 [][]byte
		result2 error
	}
	ConfigInfoStub        func() types.BlobInfo
	configInfoMutex       sync.RWMutex
	configInfoArgsForCall []struct{}
	configInfoReturns     struct {
		result1 types.BlobInfo
	}
	configInfoReturnsOnCall map[int]struct {
		result1 types.BlobInfo
	}
	ConfigBlobStub        func(context.Context) ([]byte, error)
	configBlobMutex       sync.RWMutex
	configBlobArgsForCall []struct {
		arg1 context.Context
	}
	configBlobReturns struct {
		result1 []byte
		result2 error
	}
	configBlobReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	OCIConfigStub        func(context.Context) (*v1.Image, error)
	oCIConfigMutex       sync.RWMutex
	oCIConfigArgsForCall []struct {
		arg1 context.Context
	}
	oCIConfigReturns struct {
		result1 *v1.Image
		result2 error
	}
	oCIConfigReturnsOnCall map[int]struct {
		result1 *v1.Image
		result2 error
	}
	LayerInfosStub        func() []types.BlobInfo
	layerInfosMutex       sync.RWMutex
	layerInfosArgsForCall []struct{}
	layerInfosReturns     struct {
		result1 []types.BlobInfo
	}
	layerInfosReturnsOnCall map[int]struct {
		result1 []types.BlobInfo
	}
	LayerInfosForCopyStub        func(context.Context) ([]types.BlobInfo, error)
	layerInfosForCopyMutex       sync.RWMutex
	layerInfosForCopyArgsForCall []struct {
		arg1 context.Context
	}
	layerInfosForCopyReturns struct {
		result1 []types.BlobInfo
		result2 error
	}
	layerInfosForCopyReturnsOnCall map[int]struct {
		result1 []types.BlobInfo
		result2 error
	}
	EmbeddedDockerReferenceConflictsStub        func(ref reference.Named) bool
	embeddedDockerReferenceConflictsMutex       sync.RWMutex
	embeddedDockerReferenceConflictsArgsForCall []struct {
		ref reference.Named
	}
	embeddedDockerReferenceConflictsReturns struct {
		result1 bool
	}
	embeddedDockerReferenceConflictsReturnsOnCall map[int]struct {
		result1 bool
	}
	InspectStub        func(context.Context) (*types.ImageInspectInfo, error)
	inspectMutex       sync.RWMutex
	inspectArgsForCall []struct {
		arg1 context.Context
	}
	inspectReturns struct {
		result1 *types.ImageInspectInfo
		result2 error
	}
	inspectReturnsOnCall map[int]struct {
		result1 *types.ImageInspectInfo
		result2 error
	}
	UpdatedImageNeedsLayerDiffIDsStub        func(options types.ManifestUpdateOptions) bool
	updatedImageNeedsLayerDiffIDsMutex       sync.RWMutex
	updatedImageNeedsLayerDiffIDsArgsForCall []struct {
		options types.ManifestUpdateOptions
	}
	updatedImageNeedsLayerDiffIDsReturns struct {
		result1 bool
	}
	updatedImageNeedsLayerDiffIDsReturnsOnCall map[int]struct {
		result1 bool
	}
	UpdatedImageStub        func(ctx context.Context, options types.ManifestUpdateOptions) (types.Image, error)
	updatedImageMutex       sync.RWMutex
	updatedImageArgsForCall []struct {
		ctx     context.Context
		options types.ManifestUpdateOptions
	}
	updatedImageReturns struct {
		result1 types.Image
		result2 error
	}
	updatedImageReturnsOnCall map[int]struct {
		result1 types.Image
		result2 error
	}
	SizeStub        func() (int64, error)
	sizeMutex       sync.RWMutex
	sizeArgsForCall []struct{}
	sizeReturns     struct {
		result1 int64
		result2 error
	}
	sizeReturnsOnCall map[int]struct {
		result1 int64
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeManifest) Reference() types.ImageReference {
	fake.referenceMutex.Lock()
	ret, specificReturn := fake.referenceReturnsOnCall[len(fake.referenceArgsForCall)]
	fake.referenceArgsForCall = append(fake.referenceArgsForCall, struct{}{})
	fake.recordInvocation("Reference", []interface{}{})
	fake.referenceMutex.Unlock()
	if fake.ReferenceStub != nil {
		return fake.ReferenceStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.referenceReturns.result1
}

func (fake *FakeManifest) ReferenceCallCount() int {
	fake.referenceMutex.RLock()
	defer fake.referenceMutex.RUnlock()
	return len(fake.referenceArgsForCall)
}

func (fake *FakeManifest) ReferenceReturns(result1 types.ImageReference) {
	fake.ReferenceStub = nil
	fake.referenceReturns = struct {
		result1 types.ImageReference
	}{result1}
}

func (fake *FakeManifest) ReferenceReturnsOnCall(i int, result1 types.ImageReference) {
	fake.ReferenceStub = nil
	if fake.referenceReturnsOnCall == nil {
		fake.referenceReturnsOnCall = make(map[int]struct {
			result1 types.ImageReference
		})
	}
	fake.referenceReturnsOnCall[i] = struct {
		result1 types.ImageReference
	}{result1}
}

func (fake *FakeManifest) Manifest(ctx context.Context) ([]byte, string, error) {
	fake.manifestMutex.Lock()
	ret, specificReturn := fake.manifestReturnsOnCall[len(fake.manifestArgsForCall)]
	fake.manifestArgsForCall = append(fake.manifestArgsForCall, struct {
		ctx context.Context
	}{ctx})
	fake.recordInvocation("Manifest", []interface{}{ctx})
	fake.manifestMutex.Unlock()
	if fake.ManifestStub != nil {
		return fake.ManifestStub(ctx)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.manifestReturns.result1, fake.manifestReturns.result2, fake.manifestReturns.result3
}

func (fake *FakeManifest) ManifestCallCount() int {
	fake.manifestMutex.RLock()
	defer fake.manifestMutex.RUnlock()
	return len(fake.manifestArgsForCall)
}

func (fake *FakeManifest) ManifestArgsForCall(i int) context.Context {
	fake.manifestMutex.RLock()
	defer fake.manifestMutex.RUnlock()
	return fake.manifestArgsForCall[i].ctx
}

func (fake *FakeManifest) ManifestReturns(result1 []byte, result2 string, result3 error) {
	fake.ManifestStub = nil
	fake.manifestReturns = struct {
		result1 []byte
		result2 string
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeManifest) ManifestReturnsOnCall(i int, result1 []byte, result2 string, result3 error) {
	fake.ManifestStub = nil
	if fake.manifestReturnsOnCall == nil {
		fake.manifestReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 string
			result3 error
		})
	}
	fake.manifestReturnsOnCall[i] = struct {
		result1 []byte
		result2 string
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeManifest) Signatures(ctx context.Context) ([][]byte, error) {
	fake.signaturesMutex.Lock()
	ret, specificReturn := fake.signaturesReturnsOnCall[len(fake.signaturesArgsForCall)]
	fake.signaturesArgsForCall = append(fake.signaturesArgsForCall, struct {
		ctx context.Context
	}{ctx})
	fake.recordInvocation("Signatures", []interface{}{ctx})
	fake.signaturesMutex.Unlock()
	if fake.SignaturesStub != nil {
		return fake.SignaturesStub(ctx)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.signaturesReturns.result1, fake.signaturesReturns.result2
}

func (fake *FakeManifest) SignaturesCallCount() int {
	fake.signaturesMutex.RLock()
	defer fake.signaturesMutex.RUnlock()
	return len(fake.signaturesArgsForCall)
}

func (fake *FakeManifest) SignaturesArgsForCall(i int) context.Context {
	fake.signaturesMutex.RLock()
	defer fake.signaturesMutex.RUnlock()
	return fake.signaturesArgsForCall[i].ctx
}

func (fake *FakeManifest) SignaturesReturns(result1 [][]byte, result2 error) {
	fake.SignaturesStub = nil
	fake.signaturesReturns = struct {
		result1 [][]byte
		result2 error
	}{result1, result2}
}

func (fake *FakeManifest) SignaturesReturnsOnCall(i int, result1 [][]byte, result2 error) {
	fake.SignaturesStub = nil
	if fake.signaturesReturnsOnCall == nil {
		fake.signaturesReturnsOnCall = make(map[int]struct {
			result1 [][]byte
			result2 error
		})
	}
	fake.signaturesReturnsOnCall[i] = struct {
		result1 [][]byte
		result2 error
	}{result1, result2}
}

func (fake *FakeManifest) ConfigInfo() types.BlobInfo {
	fake.configInfoMutex.Lock()
	ret, specificReturn := fake.configInfoReturnsOnCall[len(fake.configInfoArgsForCall)]
	fake.configInfoArgsForCall = append(fake.configInfoArgsForCall, struct{}{})
	fake.recordInvocation("ConfigInfo", []interface{}{})
	fake.configInfoMutex.Unlock()
	if fake.ConfigInfoStub != nil {
		return fake.ConfigInfoStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.configInfoReturns.result1
}

func (fake *FakeManifest) ConfigInfoCallCount() int {
	fake.configInfoMutex.RLock()
	defer fake.configInfoMutex.RUnlock()
	return len(fake.configInfoArgsForCall)
}

func (fake *FakeManifest) ConfigInfoReturns(result1 types.BlobInfo) {
	fake.ConfigInfoStub = nil
	fake.configInfoReturns = struct {
		result1 types.BlobInfo
	}{result1}
}

func (fake *FakeManifest) ConfigInfoReturnsOnCall(i int, result1 types.BlobInfo) {
	fake.ConfigInfoStub = nil
	if fake.configInfoReturnsOnCall == nil {
		fake.configInfoReturnsOnCall = make(map[int]struct {
			result1 types.BlobInfo
		})
	}
	fake.configInfoReturnsOnCall[i] = struct {
		result1 types.BlobInfo
	}{result1}
}

func (fake *FakeManifest) ConfigBlob(arg1 context.Context) ([]byte, error) {
	fake.configBlobMutex.Lock()
	ret, specificReturn := fake.configBlobReturnsOnCall[len(fake.configBlobArgsForCall)]
	fake.configBlobArgsForCall = append(fake.configBlobArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	fake.recordInvocation("ConfigBlob", []interface{}{arg1})
	fake.configBlobMutex.Unlock()
	if fake.ConfigBlobStub != nil {
		return fake.ConfigBlobStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.configBlobReturns.result1, fake.configBlobReturns.result2
}

func (fake *FakeManifest) ConfigBlobCallCount() int {
	fake.configBlobMutex.RLock()
	defer fake.configBlobMutex.RUnlock()
	return len(fake.configBlobArgsForCall)
}

func (fake *FakeManifest) ConfigBlobArgsForCall(i int) context.Context {
	fake.configBlobMutex.RLock()
	defer fake.configBlobMutex.RUnlock()
	return fake.configBlobArgsForCall[i].arg1
}

func (fake *FakeManifest) ConfigBlobReturns(result1 []byte, result2 error) {
	fake.ConfigBlobStub = nil
	fake.configBlobReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeManifest) ConfigBlobReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.ConfigBlobStub = nil
	if fake.configBlobReturnsOnCall == nil {
		fake.configBlobReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.configBlobReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeManifest) OCIConfig(arg1 context.Context) (*v1.Image, error) {
	fake.oCIConfigMutex.Lock()
	ret, specificReturn := fake.oCIConfigReturnsOnCall[len(fake.oCIConfigArgsForCall)]
	fake.oCIConfigArgsForCall = append(fake.oCIConfigArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	fake.recordInvocation("OCIConfig", []interface{}{arg1})
	fake.oCIConfigMutex.Unlock()
	if fake.OCIConfigStub != nil {
		return fake.OCIConfigStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.oCIConfigReturns.result1, fake.oCIConfigReturns.result2
}

func (fake *FakeManifest) OCIConfigCallCount() int {
	fake.oCIConfigMutex.RLock()
	defer fake.oCIConfigMutex.RUnlock()
	return len(fake.oCIConfigArgsForCall)
}

func (fake *FakeManifest) OCIConfigArgsForCall(i int) context.Context {
	fake.oCIConfigMutex.RLock()
	defer fake.oCIConfigMutex.RUnlock()
	return fake.oCIConfigArgsForCall[i].arg1
}

func (fake *FakeManifest) OCIConfigReturns(result1 *v1.Image, result2 error) {
	fake.OCIConfigStub = nil
	fake.oCIConfigReturns = struct {
		result1 *v1.Image
		result2 error
	}{result1, result2}
}

func (fake *FakeManifest) OCIConfigReturnsOnCall(i int, result1 *v1.Image, result2 error) {
	fake.OCIConfigStub = nil
	if fake.oCIConfigReturnsOnCall == nil {
		fake.oCIConfigReturnsOnCall = make(map[int]struct {
			result1 *v1.Image
			result2 error
		})
	}
	fake.oCIConfigReturnsOnCall[i] = struct {
		result1 *v1.Image
		result2 error
	}{result1, result2}
}

func (fake *FakeManifest) LayerInfos() []types.BlobInfo {
	fake.layerInfosMutex.Lock()
	ret, specificReturn := fake.layerInfosReturnsOnCall[len(fake.layerInfosArgsForCall)]
	fake.layerInfosArgsForCall = append(fake.layerInfosArgsForCall, struct{}{})
	fake.recordInvocation("LayerInfos", []interface{}{})
	fake.layerInfosMutex.Unlock()
	if fake.LayerInfosStub != nil {
		return fake.LayerInfosStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.layerInfosReturns.result1
}

func (fake *FakeManifest) LayerInfosCallCount() int {
	fake.layerInfosMutex.RLock()
	defer fake.layerInfosMutex.RUnlock()
	return len(fake.layerInfosArgsForCall)
}

func (fake *FakeManifest) LayerInfosReturns(result1 []types.BlobInfo) {
	fake.LayerInfosStub = nil
	fake.layerInfosReturns = struct {
		result1 []types.BlobInfo
	}{result1}
}

func (fake *FakeManifest) LayerInfosReturnsOnCall(i int, result1 []types.BlobInfo) {
	fake.LayerInfosStub = nil
	if fake.layerInfosReturnsOnCall == nil {
		fake.layerInfosReturnsOnCall = make(map[int]struct {
			result1 []types.BlobInfo
		})
	}
	fake.layerInfosReturnsOnCall[i] = struct {
		result1 []types.BlobInfo
	}{result1}
}

func (fake *FakeManifest) LayerInfosForCopy(arg1 context.Context) ([]types.BlobInfo, error) {
	fake.layerInfosForCopyMutex.Lock()
	ret, specificReturn := fake.layerInfosForCopyReturnsOnCall[len(fake.layerInfosForCopyArgsForCall)]
	fake.layerInfosForCopyArgsForCall = append(fake.layerInfosForCopyArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	fake.recordInvocation("LayerInfosForCopy", []interface{}{arg1})
	fake.layerInfosForCopyMutex.Unlock()
	if fake.LayerInfosForCopyStub != nil {
		return fake.LayerInfosForCopyStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.layerInfosForCopyReturns.result1, fake.layerInfosForCopyReturns.result2
}

func (fake *FakeManifest) LayerInfosForCopyCallCount() int {
	fake.layerInfosForCopyMutex.RLock()
	defer fake.layerInfosForCopyMutex.RUnlock()
	return len(fake.layerInfosForCopyArgsForCall)
}

func (fake *FakeManifest) LayerInfosForCopyArgsForCall(i int) context.Context {
	fake.layerInfosForCopyMutex.RLock()
	defer fake.layerInfosForCopyMutex.RUnlock()
	return fake.layerInfosForCopyArgsForCall[i].arg1
}

func (fake *FakeManifest) LayerInfosForCopyReturns(result1 []types.BlobInfo, result2 error) {
	fake.LayerInfosForCopyStub = nil
	fake.layerInfosForCopyReturns = struct {
		result1 []types.BlobInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeManifest) LayerInfosForCopyReturnsOnCall(i int, result1 []types.BlobInfo, result2 error) {
	fake.LayerInfosForCopyStub = nil
	if fake.layerInfosForCopyReturnsOnCall == nil {
		fake.layerInfosForCopyReturnsOnCall = make(map[int]struct {
			result1 []types.BlobInfo
			result2 error
		})
	}
	fake.layerInfosForCopyReturnsOnCall[i] = struct {
		result1 []types.BlobInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeManifest) EmbeddedDockerReferenceConflicts(ref reference.Named) bool {
	fake.embeddedDockerReferenceConflictsMutex.Lock()
	ret, specificReturn := fake.embeddedDockerReferenceConflictsReturnsOnCall[len(fake.embeddedDockerReferenceConflictsArgsForCall)]
	fake.embeddedDockerReferenceConflictsArgsForCall = append(fake.embeddedDockerReferenceConflictsArgsForCall, struct {
		ref reference.Named
	}{ref})
	fake.recordInvocation("EmbeddedDockerReferenceConflicts", []interface{}{ref})
	fake.embeddedDockerReferenceConflictsMutex.Unlock()
	if fake.EmbeddedDockerReferenceConflictsStub != nil {
		return fake.EmbeddedDockerReferenceConflictsStub(ref)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.embeddedDockerReferenceConflictsReturns.result1
}

func (fake *FakeManifest) EmbeddedDockerReferenceConflictsCallCount() int {
	fake.embeddedDockerReferenceConflictsMutex.RLock()
	defer fake.embeddedDockerReferenceConflictsMutex.RUnlock()
	return len(fake.embeddedDockerReferenceConflictsArgsForCall)
}

func (fake *FakeManifest) EmbeddedDockerReferenceConflictsArgsForCall(i int) reference.Named {
	fake.embeddedDockerReferenceConflictsMutex.RLock()
	defer fake.embeddedDockerReferenceConflictsMutex.RUnlock()
	return fake.embeddedDockerReferenceConflictsArgsForCall[i].ref
}

func (fake *FakeManifest) EmbeddedDockerReferenceConflictsReturns(result1 bool) {
	fake.EmbeddedDockerReferenceConflictsStub = nil
	fake.embeddedDockerReferenceConflictsReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeManifest) EmbeddedDockerReferenceConflictsReturnsOnCall(i int, result1 bool) {
	fake.EmbeddedDockerReferenceConflictsStub = nil
	if fake.embeddedDockerReferenceConflictsReturnsOnCall == nil {
		fake.embeddedDockerReferenceConflictsReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.embeddedDockerReferenceConflictsReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeManifest) Inspect(arg1 context.Context) (*types.ImageInspectInfo, error) {
	fake.inspectMutex.Lock()
	ret, specificReturn := fake.inspectReturnsOnCall[len(fake.inspectArgsForCall)]
	fake.inspectArgsForCall = append(fake.inspectArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	fake.recordInvocation("Inspect", []interface{}{arg1})
	fake.inspectMutex.Unlock()
	if fake.InspectStub != nil {
		return fake.InspectStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.inspectReturns.result1, fake.inspectReturns.result2
}

func (fake *FakeManifest) InspectCallCount() int {
	fake.inspectMutex.RLock()
	defer fake.inspectMutex.RUnlock()
	return len(fake.inspectArgsForCall)
}

func (fake *FakeManifest) InspectArgsForCall(i int) context.Context {
	fake.inspectMutex.RLock()
	defer fake.inspectMutex.RUnlock()
	return fake.inspectArgsForCall[i].arg1
}

func (fake *FakeManifest) InspectReturns(result1 *types.ImageInspectInfo, result2 error) {
	fake.InspectStub = nil
	fake.inspectReturns = struct {
		result1 *types.ImageInspectInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeManifest) InspectReturnsOnCall(i int, result1 *types.ImageInspectInfo, result2 error) {
	fake.InspectStub = nil
	if fake.inspectReturnsOnCall == nil {
		fake.inspectReturnsOnCall = make(map[int]struct {
			result1 *types.ImageInspectInfo
			result2 error
		})
	}
	fake.inspectReturnsOnCall[i] = struct {
		result1 *types.ImageInspectInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeManifest) UpdatedImageNeedsLayerDiffIDs(options types.ManifestUpdateOptions) bool {
	fake.updatedImageNeedsLayerDiffIDsMutex.Lock()
	ret, specificReturn := fake.updatedImageNeedsLayerDiffIDsReturnsOnCall[len(fake.updatedImageNeedsLayerDiffIDsArgsForCall)]
	fake.updatedImageNeedsLayerDiffIDsArgsForCall = append(fake.updatedImageNeedsLayerDiffIDsArgsForCall, struct {
		options types.ManifestUpdateOptions
	}{options})
	fake.recordInvocation("UpdatedImageNeedsLayerDiffIDs", []interface{}{options})
	fake.updatedImageNeedsLayerDiffIDsMutex.Unlock()
	if fake.UpdatedImageNeedsLayerDiffIDsStub != nil {
		return fake.UpdatedImageNeedsLayerDiffIDsStub(options)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.updatedImageNeedsLayerDiffIDsReturns.result1
}

func (fake *FakeManifest) UpdatedImageNeedsLayerDiffIDsCallCount() int {
	fake.updatedImageNeedsLayerDiffIDsMutex.RLock()
	defer fake.updatedImageNeedsLayerDiffIDsMutex.RUnlock()
	return len(fake.updatedImageNeedsLayerDiffIDsArgsForCall)
}

func (fake *FakeManifest) UpdatedImageNeedsLayerDiffIDsArgsForCall(i int) types.ManifestUpdateOptions {
	fake.updatedImageNeedsLayerDiffIDsMutex.RLock()
	defer fake.updatedImageNeedsLayerDiffIDsMutex.RUnlock()
	return fake.updatedImageNeedsLayerDiffIDsArgsForCall[i].options
}

func (fake *FakeManifest) UpdatedImageNeedsLayerDiffIDsReturns(result1 bool) {
	fake.UpdatedImageNeedsLayerDiffIDsStub = nil
	fake.updatedImageNeedsLayerDiffIDsReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeManifest) UpdatedImageNeedsLayerDiffIDsReturnsOnCall(i int, result1 bool) {
	fake.UpdatedImageNeedsLayerDiffIDsStub = nil
	if fake.updatedImageNeedsLayerDiffIDsReturnsOnCall == nil {
		fake.updatedImageNeedsLayerDiffIDsReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.updatedImageNeedsLayerDiffIDsReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeManifest) UpdatedImage(ctx context.Context, options types.ManifestUpdateOptions) (types.Image, error) {
	fake.updatedImageMutex.Lock()
	ret, specificReturn := fake.updatedImageReturnsOnCall[len(fake.updatedImageArgsForCall)]
	fake.updatedImageArgsForCall = append(fake.updatedImageArgsForCall, struct {
		ctx     context.Context
		options types.ManifestUpdateOptions
	}{ctx, options})
	fake.recordInvocation("UpdatedImage", []interface{}{ctx, options})
	fake.updatedImageMutex.Unlock()
	if fake.UpdatedImageStub != nil {
		return fake.UpdatedImageStub(ctx, options)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.updatedImageReturns.result1, fake.updatedImageReturns.result2
}

func (fake *FakeManifest) UpdatedImageCallCount() int {
	fake.updatedImageMutex.RLock()
	defer fake.updatedImageMutex.RUnlock()
	return len(fake.updatedImageArgsForCall)
}

func (fake *FakeManifest) UpdatedImageArgsForCall(i int) (context.Context, types.ManifestUpdateOptions) {
	fake.updatedImageMutex.RLock()
	defer fake.updatedImageMutex.RUnlock()
	return fake.updatedImageArgsForCall[i].ctx, fake.updatedImageArgsForCall[i].options
}

func (fake *FakeManifest) UpdatedImageReturns(result1 types.Image, result2 error) {
	fake.UpdatedImageStub = nil
	fake.updatedImageReturns = struct {
		result1 types.Image
		result2 error
	}{result1, result2}
}

func (fake *FakeManifest) UpdatedImageReturnsOnCall(i int, result1 types.Image, result2 error) {
	fake.UpdatedImageStub = nil
	if fake.updatedImageReturnsOnCall == nil {
		fake.updatedImageReturnsOnCall = make(map[int]struct {
			result1 types.Image
			result2 error
		})
	}
	fake.updatedImageReturnsOnCall[i] = struct {
		result1 types.Image
		result2 error
	}{result1, result2}
}

func (fake *FakeManifest) Size() (int64, error) {
	fake.sizeMutex.Lock()
	ret, specificReturn := fake.sizeReturnsOnCall[len(fake.sizeArgsForCall)]
	fake.sizeArgsForCall = append(fake.sizeArgsForCall, struct{}{})
	fake.recordInvocation("Size", []interface{}{})
	fake.sizeMutex.Unlock()
	if fake.SizeStub != nil {
		return fake.SizeStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.sizeReturns.result1, fake.sizeReturns.result2
}

func (fake *FakeManifest) SizeCallCount() int {
	fake.sizeMutex.RLock()
	defer fake.sizeMutex.RUnlock()
	return len(fake.sizeArgsForCall)
}

func (fake *FakeManifest) SizeReturns(result1 int64, result2 error) {
	fake.SizeStub = nil
	fake.sizeReturns = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *FakeManifest) SizeReturnsOnCall(i int, result1 int64, result2 error) {
	fake.SizeStub = nil
	if fake.sizeReturnsOnCall == nil {
		fake.sizeReturnsOnCall = make(map[int]struct {
			result1 int64
			result2 error
		})
	}
	fake.sizeReturnsOnCall[i] = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *FakeManifest) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.referenceMutex.RLock()
	defer fake.referenceMutex.RUnlock()
	fake.manifestMutex.RLock()
	defer fake.manifestMutex.RUnlock()
	fake.signaturesMutex.RLock()
	defer fake.signaturesMutex.RUnlock()
	fake.configInfoMutex.RLock()
	defer fake.configInfoMutex.RUnlock()
	fake.configBlobMutex.RLock()
	defer fake.configBlobMutex.RUnlock()
	fake.oCIConfigMutex.RLock()
	defer fake.oCIConfigMutex.RUnlock()
	fake.layerInfosMutex.RLock()
	defer fake.layerInfosMutex.RUnlock()
	fake.layerInfosForCopyMutex.RLock()
	defer fake.layerInfosForCopyMutex.RUnlock()
	fake.embeddedDockerReferenceConflictsMutex.RLock()
	defer fake.embeddedDockerReferenceConflictsMutex.RUnlock()
	fake.inspectMutex.RLock()
	defer fake.inspectMutex.RUnlock()
	fake.updatedImageNeedsLayerDiffIDsMutex.RLock()
	defer fake.updatedImageNeedsLayerDiffIDsMutex.RUnlock()
	fake.updatedImageMutex.RLock()
	defer fake.updatedImageMutex.RUnlock()
	fake.sizeMutex.RLock()
	defer fake.sizeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeManifest) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ layer_fetcher.Manifest = new(FakeManifest)
