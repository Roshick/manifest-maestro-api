# HelmRepositoryReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RepositoryType** | **string** |  | 
**RepositoryURL** | **string** |  | 
**GitReference** | **string** |  | 

## Methods

### NewHelmRepositoryReference

`func NewHelmRepositoryReference(repositoryType string, repositoryURL string, gitReference string, ) *HelmRepositoryReference`

NewHelmRepositoryReference instantiates a new HelmRepositoryReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelmRepositoryReferenceWithDefaults

`func NewHelmRepositoryReferenceWithDefaults() *HelmRepositoryReference`

NewHelmRepositoryReferenceWithDefaults instantiates a new HelmRepositoryReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepositoryType

`func (o *HelmRepositoryReference) GetRepositoryType() string`

GetRepositoryType returns the RepositoryType field if non-nil, zero value otherwise.

### GetRepositoryTypeOk

`func (o *HelmRepositoryReference) GetRepositoryTypeOk() (*string, bool)`

GetRepositoryTypeOk returns a tuple with the RepositoryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryType

`func (o *HelmRepositoryReference) SetRepositoryType(v string)`

SetRepositoryType sets RepositoryType field to given value.


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


### GetGitReference

`func (o *HelmRepositoryReference) GetGitReference() string`

GetGitReference returns the GitReference field if non-nil, zero value otherwise.

### GetGitReferenceOk

`func (o *HelmRepositoryReference) GetGitReferenceOk() (*string, bool)`

GetGitReferenceOk returns a tuple with the GitReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitReference

`func (o *HelmRepositoryReference) SetGitReference(v string)`

SetGitReference sets GitReference field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


