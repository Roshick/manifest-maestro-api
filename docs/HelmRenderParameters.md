# HelmRenderParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReleaseName** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**KubeVersion** | Pointer to **string** |  | [optional] 
**ApiVersions** | Pointer to **[]string** |  | [optional] 
**ValueFiles** | Pointer to **[]string** |  | [optional] 
**Values** | Pointer to  |  | [optional] 
**ValuesFlat** | Pointer to **[]string** |  | [optional] 
**StringValues** | Pointer to  |  | [optional] 
**StringValuesFlat** | Pointer to **[]string** |  | [optional] 
**ComplexValues** | Pointer to **map[string]interface{}** |  | [optional] 
**IncludeCRDs** | Pointer to **bool** |  | [optional] [default to true]
**IncludeHooks** | Pointer to **bool** |  | [optional] [default to true]
**IgnoreMissingValueFiles** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewHelmRenderParameters

`func NewHelmRenderParameters() *HelmRenderParameters`

NewHelmRenderParameters instantiates a new HelmRenderParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelmRenderParametersWithDefaults

`func NewHelmRenderParametersWithDefaults() *HelmRenderParameters`

NewHelmRenderParametersWithDefaults instantiates a new HelmRenderParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReleaseName

`func (o *HelmRenderParameters) GetReleaseName() string`

GetReleaseName returns the ReleaseName field if non-nil, zero value otherwise.

### GetReleaseNameOk

`func (o *HelmRenderParameters) GetReleaseNameOk() (*string, bool)`

GetReleaseNameOk returns a tuple with the ReleaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseName

`func (o *HelmRenderParameters) SetReleaseName(v string)`

SetReleaseName sets ReleaseName field to given value.

### HasReleaseName

`func (o *HelmRenderParameters) HasReleaseName() bool`

HasReleaseName returns a boolean if a field has been set.

### GetNamespace

`func (o *HelmRenderParameters) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *HelmRenderParameters) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *HelmRenderParameters) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *HelmRenderParameters) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetKubeVersion

`func (o *HelmRenderParameters) GetKubeVersion() string`

GetKubeVersion returns the KubeVersion field if non-nil, zero value otherwise.

### GetKubeVersionOk

`func (o *HelmRenderParameters) GetKubeVersionOk() (*string, bool)`

GetKubeVersionOk returns a tuple with the KubeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeVersion

`func (o *HelmRenderParameters) SetKubeVersion(v string)`

SetKubeVersion sets KubeVersion field to given value.

### HasKubeVersion

`func (o *HelmRenderParameters) HasKubeVersion() bool`

HasKubeVersion returns a boolean if a field has been set.

### GetApiVersions

`func (o *HelmRenderParameters) GetApiVersions() []string`

GetApiVersions returns the ApiVersions field if non-nil, zero value otherwise.

### GetApiVersionsOk

`func (o *HelmRenderParameters) GetApiVersionsOk() (*[]string, bool)`

GetApiVersionsOk returns a tuple with the ApiVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersions

`func (o *HelmRenderParameters) SetApiVersions(v []string)`

SetApiVersions sets ApiVersions field to given value.

### HasApiVersions

`func (o *HelmRenderParameters) HasApiVersions() bool`

HasApiVersions returns a boolean if a field has been set.

### GetValueFiles

`func (o *HelmRenderParameters) GetValueFiles() []string`

GetValueFiles returns the ValueFiles field if non-nil, zero value otherwise.

### GetValueFilesOk

`func (o *HelmRenderParameters) GetValueFilesOk() (*[]string, bool)`

GetValueFilesOk returns a tuple with the ValueFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueFiles

`func (o *HelmRenderParameters) SetValueFiles(v []string)`

SetValueFiles sets ValueFiles field to given value.

### HasValueFiles

`func (o *HelmRenderParameters) HasValueFiles() bool`

HasValueFiles returns a boolean if a field has been set.

### GetValues

`func (o *HelmRenderParameters) GetValues() map[string]string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *HelmRenderParameters) GetValuesOk() (*map[string]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *HelmRenderParameters) SetValues(v map[string]string)`

SetValues sets Values field to given value.

### HasValues

`func (o *HelmRenderParameters) HasValues() bool`

HasValues returns a boolean if a field has been set.

### SetValuesNil

`func (o *HelmRenderParameters) SetValuesNil(b bool)`

 SetValuesNil sets the value for Values to be an explicit nil

### UnsetValues
`func (o *HelmRenderParameters) UnsetValues()`

UnsetValues ensures that no value is present for Values, not even an explicit nil
### GetValuesFlat

`func (o *HelmRenderParameters) GetValuesFlat() []string`

GetValuesFlat returns the ValuesFlat field if non-nil, zero value otherwise.

### GetValuesFlatOk

`func (o *HelmRenderParameters) GetValuesFlatOk() (*[]string, bool)`

GetValuesFlatOk returns a tuple with the ValuesFlat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValuesFlat

`func (o *HelmRenderParameters) SetValuesFlat(v []string)`

SetValuesFlat sets ValuesFlat field to given value.

### HasValuesFlat

`func (o *HelmRenderParameters) HasValuesFlat() bool`

HasValuesFlat returns a boolean if a field has been set.

### GetStringValues

`func (o *HelmRenderParameters) GetStringValues() map[string]string`

GetStringValues returns the StringValues field if non-nil, zero value otherwise.

### GetStringValuesOk

`func (o *HelmRenderParameters) GetStringValuesOk() (*map[string]string, bool)`

GetStringValuesOk returns a tuple with the StringValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringValues

`func (o *HelmRenderParameters) SetStringValues(v map[string]string)`

SetStringValues sets StringValues field to given value.

### HasStringValues

`func (o *HelmRenderParameters) HasStringValues() bool`

HasStringValues returns a boolean if a field has been set.

### SetStringValuesNil

`func (o *HelmRenderParameters) SetStringValuesNil(b bool)`

 SetStringValuesNil sets the value for StringValues to be an explicit nil

### UnsetStringValues
`func (o *HelmRenderParameters) UnsetStringValues()`

UnsetStringValues ensures that no value is present for StringValues, not even an explicit nil
### GetStringValuesFlat

`func (o *HelmRenderParameters) GetStringValuesFlat() []string`

GetStringValuesFlat returns the StringValuesFlat field if non-nil, zero value otherwise.

### GetStringValuesFlatOk

`func (o *HelmRenderParameters) GetStringValuesFlatOk() (*[]string, bool)`

GetStringValuesFlatOk returns a tuple with the StringValuesFlat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringValuesFlat

`func (o *HelmRenderParameters) SetStringValuesFlat(v []string)`

SetStringValuesFlat sets StringValuesFlat field to given value.

### HasStringValuesFlat

`func (o *HelmRenderParameters) HasStringValuesFlat() bool`

HasStringValuesFlat returns a boolean if a field has been set.

### GetComplexValues

`func (o *HelmRenderParameters) GetComplexValues() map[string]interface{}`

GetComplexValues returns the ComplexValues field if non-nil, zero value otherwise.

### GetComplexValuesOk

`func (o *HelmRenderParameters) GetComplexValuesOk() (*map[string]interface{}, bool)`

GetComplexValuesOk returns a tuple with the ComplexValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplexValues

`func (o *HelmRenderParameters) SetComplexValues(v map[string]interface{})`

SetComplexValues sets ComplexValues field to given value.

### HasComplexValues

`func (o *HelmRenderParameters) HasComplexValues() bool`

HasComplexValues returns a boolean if a field has been set.

### GetIncludeCRDs

`func (o *HelmRenderParameters) GetIncludeCRDs() bool`

GetIncludeCRDs returns the IncludeCRDs field if non-nil, zero value otherwise.

### GetIncludeCRDsOk

`func (o *HelmRenderParameters) GetIncludeCRDsOk() (*bool, bool)`

GetIncludeCRDsOk returns a tuple with the IncludeCRDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeCRDs

`func (o *HelmRenderParameters) SetIncludeCRDs(v bool)`

SetIncludeCRDs sets IncludeCRDs field to given value.

### HasIncludeCRDs

`func (o *HelmRenderParameters) HasIncludeCRDs() bool`

HasIncludeCRDs returns a boolean if a field has been set.

### GetIncludeHooks

`func (o *HelmRenderParameters) GetIncludeHooks() bool`

GetIncludeHooks returns the IncludeHooks field if non-nil, zero value otherwise.

### GetIncludeHooksOk

`func (o *HelmRenderParameters) GetIncludeHooksOk() (*bool, bool)`

GetIncludeHooksOk returns a tuple with the IncludeHooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeHooks

`func (o *HelmRenderParameters) SetIncludeHooks(v bool)`

SetIncludeHooks sets IncludeHooks field to given value.

### HasIncludeHooks

`func (o *HelmRenderParameters) HasIncludeHooks() bool`

HasIncludeHooks returns a boolean if a field has been set.

### GetIgnoreMissingValueFiles

`func (o *HelmRenderParameters) GetIgnoreMissingValueFiles() bool`

GetIgnoreMissingValueFiles returns the IgnoreMissingValueFiles field if non-nil, zero value otherwise.

### GetIgnoreMissingValueFilesOk

`func (o *HelmRenderParameters) GetIgnoreMissingValueFilesOk() (*bool, bool)`

GetIgnoreMissingValueFilesOk returns a tuple with the IgnoreMissingValueFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreMissingValueFiles

`func (o *HelmRenderParameters) SetIgnoreMissingValueFiles(v bool)`

SetIgnoreMissingValueFiles sets IgnoreMissingValueFiles field to given value.

### HasIgnoreMissingValueFiles

`func (o *HelmRenderParameters) HasIgnoreMissingValueFiles() bool`

HasIgnoreMissingValueFiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


