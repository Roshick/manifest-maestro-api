# HelmRepositoryReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RepositoryURL** | **string** |  | 
**Reference** | **string** |  | 

## Methods

### NewHelmRepositoryReference

`func NewHelmRepositoryReference(repositoryURL string, reference string, ) *HelmRepositoryReference`

NewHelmRepositoryReference instantiates a new HelmRepositoryReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelmRepositoryReferenceWithDefaults

`func NewHelmRepositoryReferenceWithDefaults() *HelmRepositoryReference`

NewHelmRepositoryReferenceWithDefaults instantiates a new HelmRepositoryReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepositoryURL

`func (o *HelmRepositoryReference) GetRepositoryURL() string`

GetRepositoryURL returns the RepositoryURL field if non-nil, zero value otherwise.

### GetRepositoryURLOk

`func (o *HelmRepositoryReference) GetRepositoryURLOk() (*string, bool)`

GetRepositoryURLOk returns a tuple with the RepositoryURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryURL

`func (o *HelmRepositoryReference) SetRepositoryURL(v string)`

SetRepositoryURL sets RepositoryURL field to given value.


### GetReference

`func (o *HelmRepositoryReference) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *HelmRepositoryReference) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *HelmRepositoryReference) SetReference(v string)`

SetReference sets Reference field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


