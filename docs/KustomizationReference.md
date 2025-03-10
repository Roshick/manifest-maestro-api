# KustomizationReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RepositoryType** | **string** |  | 
**RepositoryURL** | **string** |  | 
**GitReference** | **string** |  | 
**Path** | Pointer to **string** |  | [optional] 

## Methods

### NewKustomizationReference

`func NewKustomizationReference(repositoryType string, repositoryURL string, gitReference string, ) *KustomizationReference`

NewKustomizationReference instantiates a new KustomizationReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKustomizationReferenceWithDefaults

`func NewKustomizationReferenceWithDefaults() *KustomizationReference`

NewKustomizationReferenceWithDefaults instantiates a new KustomizationReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepositoryType

`func (o *KustomizationReference) GetRepositoryType() string`

GetRepositoryType returns the RepositoryType field if non-nil, zero value otherwise.

### GetRepositoryTypeOk

`func (o *KustomizationReference) GetRepositoryTypeOk() (*string, bool)`

GetRepositoryTypeOk returns a tuple with the RepositoryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryType

`func (o *KustomizationReference) SetRepositoryType(v string)`

SetRepositoryType sets RepositoryType field to given value.


### GetRepositoryURL

`func (o *KustomizationReference) GetRepositoryURL() string`

GetRepositoryURL returns the RepositoryURL field if non-nil, zero value otherwise.

### GetRepositoryURLOk

`func (o *KustomizationReference) GetRepositoryURLOk() (*string, bool)`

GetRepositoryURLOk returns a tuple with the RepositoryURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryURL

`func (o *KustomizationReference) SetRepositoryURL(v string)`

SetRepositoryURL sets RepositoryURL field to given value.


### GetGitReference

`func (o *KustomizationReference) GetGitReference() string`

GetGitReference returns the GitReference field if non-nil, zero value otherwise.

### GetGitReferenceOk

`func (o *KustomizationReference) GetGitReferenceOk() (*string, bool)`

GetGitReferenceOk returns a tuple with the GitReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitReference

`func (o *KustomizationReference) SetGitReference(v string)`

SetGitReference sets GitReference field to given value.


### GetPath

`func (o *KustomizationReference) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *KustomizationReference) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *KustomizationReference) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *KustomizationReference) HasPath() bool`

HasPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


