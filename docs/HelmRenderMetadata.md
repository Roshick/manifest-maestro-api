# HelmRenderMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReleaseName** | **string** |  | 
**Namespace** | **string** |  | 
**ApiVersions** | **[]string** |  | 
**KubeVersion** | **string** |  | 
**HelmVersion** | **string** |  | 
**MergedValues** | **map[string]interface{}** |  | 

## Methods

### NewHelmRenderMetadata

`func NewHelmRenderMetadata(releaseName string, namespace string, apiVersions []string, kubeVersion string, helmVersion string, mergedValues map[string]interface{}, ) *HelmRenderMetadata`

NewHelmRenderMetadata instantiates a new HelmRenderMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelmRenderMetadataWithDefaults

`func NewHelmRenderMetadataWithDefaults() *HelmRenderMetadata`

NewHelmRenderMetadataWithDefaults instantiates a new HelmRenderMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReleaseName

`func (o *HelmRenderMetadata) GetReleaseName() string`

GetReleaseName returns the ReleaseName field if non-nil, zero value otherwise.

### GetReleaseNameOk

`func (o *HelmRenderMetadata) GetReleaseNameOk() (*string, bool)`

GetReleaseNameOk returns a tuple with the ReleaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseName

`func (o *HelmRenderMetadata) SetReleaseName(v string)`

SetReleaseName sets ReleaseName field to given value.


### GetNamespace

`func (o *HelmRenderMetadata) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *HelmRenderMetadata) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *HelmRenderMetadata) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### GetApiVersions

`func (o *HelmRenderMetadata) GetApiVersions() []string`

GetApiVersions returns the ApiVersions field if non-nil, zero value otherwise.

### GetApiVersionsOk

`func (o *HelmRenderMetadata) GetApiVersionsOk() (*[]string, bool)`

GetApiVersionsOk returns a tuple with the ApiVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersions

`func (o *HelmRenderMetadata) SetApiVersions(v []string)`

SetApiVersions sets ApiVersions field to given value.


### GetKubeVersion

`func (o *HelmRenderMetadata) GetKubeVersion() string`

GetKubeVersion returns the KubeVersion field if non-nil, zero value otherwise.

### GetKubeVersionOk

`func (o *HelmRenderMetadata) GetKubeVersionOk() (*string, bool)`

GetKubeVersionOk returns a tuple with the KubeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeVersion

`func (o *HelmRenderMetadata) SetKubeVersion(v string)`

SetKubeVersion sets KubeVersion field to given value.


### GetHelmVersion

`func (o *HelmRenderMetadata) GetHelmVersion() string`

GetHelmVersion returns the HelmVersion field if non-nil, zero value otherwise.

### GetHelmVersionOk

`func (o *HelmRenderMetadata) GetHelmVersionOk() (*string, bool)`

GetHelmVersionOk returns a tuple with the HelmVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelmVersion

`func (o *HelmRenderMetadata) SetHelmVersion(v string)`

SetHelmVersion sets HelmVersion field to given value.


### GetMergedValues

`func (o *HelmRenderMetadata) GetMergedValues() map[string]interface{}`

GetMergedValues returns the MergedValues field if non-nil, zero value otherwise.

### GetMergedValuesOk

`func (o *HelmRenderMetadata) GetMergedValuesOk() (*map[string]interface{}, bool)`

GetMergedValuesOk returns a tuple with the MergedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergedValues

`func (o *HelmRenderMetadata) SetMergedValues(v map[string]interface{})`

SetMergedValues sets MergedValues field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


