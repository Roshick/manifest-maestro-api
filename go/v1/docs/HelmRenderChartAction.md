# HelmRenderChartAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reference** | [**HelmChartReference**](HelmChartReference.md) |  | 
**Parameters** | Pointer to [**HelmRenderParameters**](HelmRenderParameters.md) |  | [optional] 

## Methods

### NewHelmRenderChartAction

`func NewHelmRenderChartAction(reference HelmChartReference, ) *HelmRenderChartAction`

NewHelmRenderChartAction instantiates a new HelmRenderChartAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelmRenderChartActionWithDefaults

`func NewHelmRenderChartActionWithDefaults() *HelmRenderChartAction`

NewHelmRenderChartActionWithDefaults instantiates a new HelmRenderChartAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReference

`func (o *HelmRenderChartAction) GetReference() HelmChartReference`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *HelmRenderChartAction) GetReferenceOk() (*HelmChartReference, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *HelmRenderChartAction) SetReference(v HelmChartReference)`

SetReference sets Reference field to given value.


### GetParameters

`func (o *HelmRenderChartAction) GetParameters() HelmRenderParameters`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *HelmRenderChartAction) GetParametersOk() (*HelmRenderParameters, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *HelmRenderChartAction) SetParameters(v HelmRenderParameters)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *HelmRenderChartAction) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


