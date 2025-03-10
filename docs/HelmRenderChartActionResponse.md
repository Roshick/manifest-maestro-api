# HelmRenderChartActionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Manifests** | Pointer to [**[]Manifest**](Manifest.md) |  | [optional] 
**Metadata** | Pointer to [**HelmRenderMetadata**](HelmRenderMetadata.md) |  | [optional] 

## Methods

### NewHelmRenderChartActionResponse

`func NewHelmRenderChartActionResponse() *HelmRenderChartActionResponse`

NewHelmRenderChartActionResponse instantiates a new HelmRenderChartActionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelmRenderChartActionResponseWithDefaults

`func NewHelmRenderChartActionResponseWithDefaults() *HelmRenderChartActionResponse`

NewHelmRenderChartActionResponseWithDefaults instantiates a new HelmRenderChartActionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetManifests

`func (o *HelmRenderChartActionResponse) GetManifests() []Manifest`

GetManifests returns the Manifests field if non-nil, zero value otherwise.

### GetManifestsOk

`func (o *HelmRenderChartActionResponse) GetManifestsOk() (*[]Manifest, bool)`

GetManifestsOk returns a tuple with the Manifests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifests

`func (o *HelmRenderChartActionResponse) SetManifests(v []Manifest)`

SetManifests sets Manifests field to given value.

### HasManifests

`func (o *HelmRenderChartActionResponse) HasManifests() bool`

HasManifests returns a boolean if a field has been set.

### GetMetadata

`func (o *HelmRenderChartActionResponse) GetMetadata() HelmRenderMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *HelmRenderChartActionResponse) GetMetadataOk() (*HelmRenderMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *HelmRenderChartActionResponse) SetMetadata(v HelmRenderMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *HelmRenderChartActionResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


