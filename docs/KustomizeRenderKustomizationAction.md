# KustomizeRenderKustomizationAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reference** | [**GitRepositoryPathReference**](GitRepositoryPathReference.md) |  | 
**Parameters** | Pointer to [**KustomizeRenderParameters**](KustomizeRenderParameters.md) |  | [optional] 

## Methods

### NewKustomizeRenderKustomizationAction

`func NewKustomizeRenderKustomizationAction(reference GitRepositoryPathReference, ) *KustomizeRenderKustomizationAction`

NewKustomizeRenderKustomizationAction instantiates a new KustomizeRenderKustomizationAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKustomizeRenderKustomizationActionWithDefaults

`func NewKustomizeRenderKustomizationActionWithDefaults() *KustomizeRenderKustomizationAction`

NewKustomizeRenderKustomizationActionWithDefaults instantiates a new KustomizeRenderKustomizationAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReference

`func (o *KustomizeRenderKustomizationAction) GetReference() GitRepositoryPathReference`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *KustomizeRenderKustomizationAction) GetReferenceOk() (*GitRepositoryPathReference, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *KustomizeRenderKustomizationAction) SetReference(v GitRepositoryPathReference)`

SetReference sets Reference field to given value.


### GetParameters

`func (o *KustomizeRenderKustomizationAction) GetParameters() KustomizeRenderParameters`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *KustomizeRenderKustomizationAction) GetParametersOk() (*KustomizeRenderParameters, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *KustomizeRenderKustomizationAction) SetParameters(v KustomizeRenderParameters)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *KustomizeRenderKustomizationAction) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


