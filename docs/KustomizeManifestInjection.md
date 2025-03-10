# KustomizeManifestInjection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileName** | **string** |  | 
**Manifests** | [**[]Manifest**](Manifest.md) |  | 

## Methods

### NewKustomizeManifestInjection

`func NewKustomizeManifestInjection(fileName string, manifests []Manifest, ) *KustomizeManifestInjection`

NewKustomizeManifestInjection instantiates a new KustomizeManifestInjection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKustomizeManifestInjectionWithDefaults

`func NewKustomizeManifestInjectionWithDefaults() *KustomizeManifestInjection`

NewKustomizeManifestInjectionWithDefaults instantiates a new KustomizeManifestInjection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileName

`func (o *KustomizeManifestInjection) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *KustomizeManifestInjection) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *KustomizeManifestInjection) SetFileName(v string)`

SetFileName sets FileName field to given value.


### GetManifests

`func (o *KustomizeManifestInjection) GetManifests() []Manifest`

GetManifests returns the Manifests field if non-nil, zero value otherwise.

### GetManifestsOk

`func (o *KustomizeManifestInjection) GetManifestsOk() (*[]Manifest, bool)`

GetManifestsOk returns a tuple with the Manifests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifests

`func (o *KustomizeManifestInjection) SetManifests(v []Manifest)`

SetManifests sets Manifests field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


