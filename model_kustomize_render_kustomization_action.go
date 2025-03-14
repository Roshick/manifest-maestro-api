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

// checks if the KustomizeRenderKustomizationAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KustomizeRenderKustomizationAction{}

// KustomizeRenderKustomizationAction struct for KustomizeRenderKustomizationAction
type KustomizeRenderKustomizationAction struct {
	Reference KustomizationReference `json:"reference"`
	Parameters *KustomizeRenderParameters `json:"parameters,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KustomizeRenderKustomizationAction KustomizeRenderKustomizationAction

// NewKustomizeRenderKustomizationAction instantiates a new KustomizeRenderKustomizationAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKustomizeRenderKustomizationAction(reference KustomizationReference) *KustomizeRenderKustomizationAction {
	this := KustomizeRenderKustomizationAction{}
	this.Reference = reference
	return &this
}

// NewKustomizeRenderKustomizationActionWithDefaults instantiates a new KustomizeRenderKustomizationAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKustomizeRenderKustomizationActionWithDefaults() *KustomizeRenderKustomizationAction {
	this := KustomizeRenderKustomizationAction{}
	return &this
}

// GetReference returns the Reference field value
func (o *KustomizeRenderKustomizationAction) GetReference() KustomizationReference {
	if o == nil {
		var ret KustomizationReference
		return ret
	}

	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value
// and a boolean to check if the value has been set.
func (o *KustomizeRenderKustomizationAction) GetReferenceOk() (*KustomizationReference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reference, true
}

// SetReference sets field value
func (o *KustomizeRenderKustomizationAction) SetReference(v KustomizationReference) {
	o.Reference = v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *KustomizeRenderKustomizationAction) GetParameters() KustomizeRenderParameters {
	if o == nil || IsNil(o.Parameters) {
		var ret KustomizeRenderParameters
		return ret
	}
	return *o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KustomizeRenderKustomizationAction) GetParametersOk() (*KustomizeRenderParameters, bool) {
	if o == nil || IsNil(o.Parameters) {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *KustomizeRenderKustomizationAction) HasParameters() bool {
	if o != nil && !IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given KustomizeRenderParameters and assigns it to the Parameters field.
func (o *KustomizeRenderKustomizationAction) SetParameters(v KustomizeRenderParameters) {
	o.Parameters = &v
}

func (o KustomizeRenderKustomizationAction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KustomizeRenderKustomizationAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reference"] = o.Reference
	if !IsNil(o.Parameters) {
		toSerialize["parameters"] = o.Parameters
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *KustomizeRenderKustomizationAction) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varKustomizeRenderKustomizationAction := _KustomizeRenderKustomizationAction{}

	err = json.Unmarshal(data, &varKustomizeRenderKustomizationAction)

	if err != nil {
		return err
	}

	*o = KustomizeRenderKustomizationAction(varKustomizeRenderKustomizationAction)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "reference")
		delete(additionalProperties, "parameters")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKustomizeRenderKustomizationAction struct {
	value *KustomizeRenderKustomizationAction
	isSet bool
}

func (v NullableKustomizeRenderKustomizationAction) Get() *KustomizeRenderKustomizationAction {
	return v.value
}

func (v *NullableKustomizeRenderKustomizationAction) Set(val *KustomizeRenderKustomizationAction) {
	v.value = val
	v.isSet = true
}

func (v NullableKustomizeRenderKustomizationAction) IsSet() bool {
	return v.isSet
}

func (v *NullableKustomizeRenderKustomizationAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKustomizeRenderKustomizationAction(val *KustomizeRenderKustomizationAction) *NullableKustomizeRenderKustomizationAction {
	return &NullableKustomizeRenderKustomizationAction{value: val, isSet: true}
}

func (v NullableKustomizeRenderKustomizationAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKustomizeRenderKustomizationAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


