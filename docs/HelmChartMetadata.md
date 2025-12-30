# HelmChartMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Version** | **string** |  | 
**AppVersion** | Pointer to **string** |  | [optional] 
**Dependencies** | [**[]HelmChartMetadata**](HelmChartMetadata.md) |  | 

## Methods

### NewHelmChartMetadata

`func NewHelmChartMetadata(name string, version string, dependencies []HelmChartMetadata, ) *HelmChartMetadata`

NewHelmChartMetadata instantiates a new HelmChartMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelmChartMetadataWithDefaults

`func NewHelmChartMetadataWithDefaults() *HelmChartMetadata`

NewHelmChartMetadataWithDefaults instantiates a new HelmChartMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *HelmChartMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HelmChartMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HelmChartMetadata) SetName(v string)`

SetName sets Name field to given value.


### GetVersion

`func (o *HelmChartMetadata) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HelmChartMetadata) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HelmChartMetadata) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetAppVersion

`func (o *HelmChartMetadata) GetAppVersion() string`

GetAppVersion returns the AppVersion field if non-nil, zero value otherwise.

### GetAppVersionOk

`func (o *HelmChartMetadata) GetAppVersionOk() (*string, bool)`

GetAppVersionOk returns a tuple with the AppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppVersion

`func (o *HelmChartMetadata) SetAppVersion(v string)`

SetAppVersion sets AppVersion field to given value.

### HasAppVersion

`func (o *HelmChartMetadata) HasAppVersion() bool`

HasAppVersion returns a boolean if a field has been set.

### GetDependencies

`func (o *HelmChartMetadata) GetDependencies() []HelmChartMetadata`

GetDependencies returns the Dependencies field if non-nil, zero value otherwise.

### GetDependenciesOk

`func (o *HelmChartMetadata) GetDependenciesOk() (*[]HelmChartMetadata, bool)`

GetDependenciesOk returns a tuple with the Dependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencies

`func (o *HelmChartMetadata) SetDependencies(v []HelmChartMetadata)`

SetDependencies sets Dependencies field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


