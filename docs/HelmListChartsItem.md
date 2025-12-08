# HelmListChartsItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Versions** | [**[]HelmListChartsItemVersion**](HelmListChartsItemVersion.md) |  | 

## Methods

### NewHelmListChartsItem

`func NewHelmListChartsItem(name string, versions []HelmListChartsItemVersion, ) *HelmListChartsItem`

NewHelmListChartsItem instantiates a new HelmListChartsItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelmListChartsItemWithDefaults

`func NewHelmListChartsItemWithDefaults() *HelmListChartsItem`

NewHelmListChartsItemWithDefaults instantiates a new HelmListChartsItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *HelmListChartsItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HelmListChartsItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HelmListChartsItem) SetName(v string)`

SetName sets Name field to given value.


### GetVersions

`func (o *HelmListChartsItem) GetVersions() []HelmListChartsItemVersion`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *HelmListChartsItem) GetVersionsOk() (*[]HelmListChartsItemVersion, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *HelmListChartsItem) SetVersions(v []HelmListChartsItemVersion)`

SetVersions sets Versions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


