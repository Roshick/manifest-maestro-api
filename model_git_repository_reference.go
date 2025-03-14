/*
Manifest Maestro

Renders Kubernetes manifests with the help of various tools such as Helm and Kustomize.

API version: v1.2.0
Contact: e.rieb@posteo.de
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// checks if the GitRepositoryReference type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GitRepositoryReference{}

// GitRepositoryReference struct for GitRepositoryReference
type GitRepositoryReference struct {
	RepositoryURL string `json:"repositoryURL"`
	Reference string `json:"reference"`
	AdditionalProperties map[string]interface{}
}

type _GitRepositoryReference GitRepositoryReference

// NewGitRepositoryReference instantiates a new GitRepositoryReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGitRepositoryReference(repositoryURL string, reference string) *GitRepositoryReference {
	this := GitRepositoryReference{}
	this.RepositoryURL = repositoryURL
	this.Reference = reference
	return &this
}

// NewGitRepositoryReferenceWithDefaults instantiates a new GitRepositoryReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGitRepositoryReferenceWithDefaults() *GitRepositoryReference {
	this := GitRepositoryReference{}
	return &this
}

// GetRepositoryURL returns the RepositoryURL field value
func (o *GitRepositoryReference) GetRepositoryURL() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RepositoryURL
}

// GetRepositoryURLOk returns a tuple with the RepositoryURL field value
// and a boolean to check if the value has been set.
func (o *GitRepositoryReference) GetRepositoryURLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RepositoryURL, true
}

// SetRepositoryURL sets field value
func (o *GitRepositoryReference) SetRepositoryURL(v string) {
	o.RepositoryURL = v
}

// GetReference returns the Reference field value
func (o *GitRepositoryReference) GetReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value
// and a boolean to check if the value has been set.
func (o *GitRepositoryReference) GetReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reference, true
}

// SetReference sets field value
func (o *GitRepositoryReference) SetReference(v string) {
	o.Reference = v
}

func (o GitRepositoryReference) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GitRepositoryReference) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["repositoryURL"] = o.RepositoryURL
	toSerialize["reference"] = o.Reference

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GitRepositoryReference) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"repositoryURL",
		"reference",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varGitRepositoryReference := _GitRepositoryReference{}

	err = json.Unmarshal(data, &varGitRepositoryReference)

	if err != nil {
		return err
	}

	*o = GitRepositoryReference(varGitRepositoryReference)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "repositoryURL")
		delete(additionalProperties, "reference")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGitRepositoryReference struct {
	value *GitRepositoryReference
	isSet bool
}

func (v NullableGitRepositoryReference) Get() *GitRepositoryReference {
	return v.value
}

func (v *NullableGitRepositoryReference) Set(val *GitRepositoryReference) {
	v.value = val
	v.isSet = true
}

func (v NullableGitRepositoryReference) IsSet() bool {
	return v.isSet
}

func (v *NullableGitRepositoryReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGitRepositoryReference(val *GitRepositoryReference) *NullableGitRepositoryReference {
	return &NullableGitRepositoryReference{value: val, isSet: true}
}

func (v NullableGitRepositoryReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGitRepositoryReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


