# GitRepositoryPathReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RepositoryType** | **string** |  | 
**RepositoryURL** | **string** |  | 
**GitReference** | **string** |  | 
**Path** | Pointer to **string** |  | [optional] 

## Methods

### NewGitRepositoryPathReference

`func NewGitRepositoryPathReference(repositoryType string, repositoryURL string, gitReference string, ) *GitRepositoryPathReference`

NewGitRepositoryPathReference instantiates a new GitRepositoryPathReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitRepositoryPathReferenceWithDefaults

`func NewGitRepositoryPathReferenceWithDefaults() *GitRepositoryPathReference`

NewGitRepositoryPathReferenceWithDefaults instantiates a new GitRepositoryPathReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepositoryType

`func (o *GitRepositoryPathReference) GetRepositoryType() string`

GetRepositoryType returns the RepositoryType field if non-nil, zero value otherwise.

### GetRepositoryTypeOk

`func (o *GitRepositoryPathReference) GetRepositoryTypeOk() (*string, bool)`

GetRepositoryTypeOk returns a tuple with the RepositoryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryType

`func (o *GitRepositoryPathReference) SetRepositoryType(v string)`

SetRepositoryType sets RepositoryType field to given value.


### GetRepositoryURL

`func (o *GitRepositoryPathReference) GetRepositoryURL() string`

GetRepositoryURL returns the RepositoryURL field if non-nil, zero value otherwise.

### GetRepositoryURLOk

`func (o *GitRepositoryPathReference) GetRepositoryURLOk() (*string, bool)`

GetRepositoryURLOk returns a tuple with the RepositoryURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryURL

`func (o *GitRepositoryPathReference) SetRepositoryURL(v string)`

SetRepositoryURL sets RepositoryURL field to given value.


### GetGitReference

`func (o *GitRepositoryPathReference) GetGitReference() string`

GetGitReference returns the GitReference field if non-nil, zero value otherwise.

### GetGitReferenceOk

`func (o *GitRepositoryPathReference) GetGitReferenceOk() (*string, bool)`

GetGitReferenceOk returns a tuple with the GitReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitReference

`func (o *GitRepositoryPathReference) SetGitReference(v string)`

SetGitReference sets GitReference field to given value.


### GetPath

`func (o *GitRepositoryPathReference) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *GitRepositoryPathReference) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *GitRepositoryPathReference) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *GitRepositoryPathReference) HasPath() bool`

HasPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


