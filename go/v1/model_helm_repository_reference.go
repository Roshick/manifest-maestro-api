/*
Manifest Maestro

Renders Kubernetes manifests with the help of various tools such as Helm and Kustomize.

API version: v1
Contact: e.rieb@posteo.de
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
	"gopkg.in/validator.v2"
)

// HelmRepositoryReference - struct for HelmRepositoryReference
type HelmRepositoryReference struct {
	GitRepositoryReference *GitRepositoryReference
	HelmChartRepositoryReference *HelmChartRepositoryReference
}

// GitRepositoryReferenceAsHelmRepositoryReference is a convenience function that returns GitRepositoryReference wrapped in HelmRepositoryReference
func GitRepositoryReferenceAsHelmRepositoryReference(v *GitRepositoryReference) HelmRepositoryReference {
	return HelmRepositoryReference{
		GitRepositoryReference: v,
	}
}

// HelmChartRepositoryReferenceAsHelmRepositoryReference is a convenience function that returns HelmChartRepositoryReference wrapped in HelmRepositoryReference
func HelmChartRepositoryReferenceAsHelmRepositoryReference(v *HelmChartRepositoryReference) HelmRepositoryReference {
	return HelmRepositoryReference{
		HelmChartRepositoryReference: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *HelmRepositoryReference) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GitRepositoryReference
	err = newStrictDecoder(data).Decode(&dst.GitRepositoryReference)
	if err == nil {
		jsonGitRepositoryReference, _ := json.Marshal(dst.GitRepositoryReference)
		if string(jsonGitRepositoryReference) == "{}" { // empty struct
			dst.GitRepositoryReference = nil
		} else {
			if err = validator.Validate(dst.GitRepositoryReference); err != nil {
				dst.GitRepositoryReference = nil
			} else {
				match++
			}
		}
	} else {
		dst.GitRepositoryReference = nil
	}

	// try to unmarshal data into HelmChartRepositoryReference
	err = newStrictDecoder(data).Decode(&dst.HelmChartRepositoryReference)
	if err == nil {
		jsonHelmChartRepositoryReference, _ := json.Marshal(dst.HelmChartRepositoryReference)
		if string(jsonHelmChartRepositoryReference) == "{}" { // empty struct
			dst.HelmChartRepositoryReference = nil
		} else {
			if err = validator.Validate(dst.HelmChartRepositoryReference); err != nil {
				dst.HelmChartRepositoryReference = nil
			} else {
				match++
			}
		}
	} else {
		dst.HelmChartRepositoryReference = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.GitRepositoryReference = nil
		dst.HelmChartRepositoryReference = nil

		return fmt.Errorf("data matches more than one schema in oneOf(HelmRepositoryReference)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(HelmRepositoryReference)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src HelmRepositoryReference) MarshalJSON() ([]byte, error) {
	if src.GitRepositoryReference != nil {
		return json.Marshal(&src.GitRepositoryReference)
	}

	if src.HelmChartRepositoryReference != nil {
		return json.Marshal(&src.HelmChartRepositoryReference)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *HelmRepositoryReference) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.GitRepositoryReference != nil {
		return obj.GitRepositoryReference
	}

	if obj.HelmChartRepositoryReference != nil {
		return obj.HelmChartRepositoryReference
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj HelmRepositoryReference) GetActualInstanceValue() (interface{}) {
	if obj.GitRepositoryReference != nil {
		return *obj.GitRepositoryReference
	}

	if obj.HelmChartRepositoryReference != nil {
		return *obj.HelmChartRepositoryReference
	}

	// all schemas are nil
	return nil
}

type NullableHelmRepositoryReference struct {
	value *HelmRepositoryReference
	isSet bool
}

func (v NullableHelmRepositoryReference) Get() *HelmRepositoryReference {
	return v.value
}

func (v *NullableHelmRepositoryReference) Set(val *HelmRepositoryReference) {
	v.value = val
	v.isSet = true
}

func (v NullableHelmRepositoryReference) IsSet() bool {
	return v.isSet
}

func (v *NullableHelmRepositoryReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHelmRepositoryReference(val *HelmRepositoryReference) *NullableHelmRepositoryReference {
	return &NullableHelmRepositoryReference{value: val, isSet: true}
}

func (v NullableHelmRepositoryReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHelmRepositoryReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


