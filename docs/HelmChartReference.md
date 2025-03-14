# HelmChartReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RepositoryURL** | **string** |  | 
**ChartName** | **string** |  | 
**ChartVersion** | Pointer to **string** |  | [optional] 
**Reference** | **string** |  | 
**Path** | Pointer to **string** |  | [optional] 

## Methods

### NewHelmChartReference

`func NewHelmChartReference(repositoryURL string, chartName string, reference string, ) *HelmChartReference`

NewHelmChartReference instantiates a new HelmChartReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelmChartReferenceWithDefaults

`func NewHelmChartReferenceWithDefaults() *HelmChartReference`

NewHelmChartReferenceWithDefaults instantiates a new HelmChartReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepositoryURL

`func (o *HelmChartReference) GetRepositoryURL() string`

GetRepositoryURL returns the RepositoryURL field if non-nil, zero value otherwise.

### GetRepositoryURLOk

`func (o *HelmChartReference) GetRepositoryURLOk() (*string, bool)`

GetRepositoryURLOk returns a tuple with the RepositoryURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryURL

`func (o *HelmChartReference) SetRepositoryURL(v string)`

SetRepositoryURL sets RepositoryURL field to given value.


### GetChartName

`func (o *HelmChartReference) GetChartName() string`

GetChartName returns the ChartName field if non-nil, zero value otherwise.

### GetChartNameOk

`func (o *HelmChartReference) GetChartNameOk() (*string, bool)`

GetChartNameOk returns a tuple with the ChartName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartName

`func (o *HelmChartReference) SetChartName(v string)`

SetChartName sets ChartName field to given value.


### GetChartVersion

`func (o *HelmChartReference) GetChartVersion() string`

GetChartVersion returns the ChartVersion field if non-nil, zero value otherwise.

### GetChartVersionOk

`func (o *HelmChartReference) GetChartVersionOk() (*string, bool)`

GetChartVersionOk returns a tuple with the ChartVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartVersion

`func (o *HelmChartReference) SetChartVersion(v string)`

SetChartVersion sets ChartVersion field to given value.

### HasChartVersion

`func (o *HelmChartReference) HasChartVersion() bool`

HasChartVersion returns a boolean if a field has been set.

### GetReference

`func (o *HelmChartReference) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *HelmChartReference) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *HelmChartReference) SetReference(v string)`

SetReference sets Reference field to given value.


### GetPath

`func (o *HelmChartReference) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *HelmChartReference) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *HelmChartReference) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *HelmChartReference) HasPath() bool`

HasPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


