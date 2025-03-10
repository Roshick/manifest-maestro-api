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
)

// checks if the HelmRenderChartActionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HelmRenderChartActionResponse{}

// HelmRenderChartActionResponse struct for HelmRenderChartActionResponse
type HelmRenderChartActionResponse struct {
	Manifests []Manifest `json:"manifests,omitempty"`
	Metadata *HelmRenderMetadata `json:"metadata,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HelmRenderChartActionResponse HelmRenderChartActionResponse

// NewHelmRenderChartActionResponse instantiates a new HelmRenderChartActionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHelmRenderChartActionResponse() *HelmRenderChartActionResponse {
	this := HelmRenderChartActionResponse{}
	return &this
}

// NewHelmRenderChartActionResponseWithDefaults instantiates a new HelmRenderChartActionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHelmRenderChartActionResponseWithDefaults() *HelmRenderChartActionResponse {
	this := HelmRenderChartActionResponse{}
	return &this
}

// GetManifests returns the Manifests field value if set, zero value otherwise.
func (o *HelmRenderChartActionResponse) GetManifests() []Manifest {
	if o == nil || IsNil(o.Manifests) {
		var ret []Manifest
		return ret
	}
	return o.Manifests
}

// GetManifestsOk returns a tuple with the Manifests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmRenderChartActionResponse) GetManifestsOk() ([]Manifest, bool) {
	if o == nil || IsNil(o.Manifests) {
		return nil, false
	}
	return o.Manifests, true
}

// HasManifests returns a boolean if a field has been set.
func (o *HelmRenderChartActionResponse) HasManifests() bool {
	if o != nil && !IsNil(o.Manifests) {
		return true
	}

	return false
}

// SetManifests gets a reference to the given []Manifest and assigns it to the Manifests field.
func (o *HelmRenderChartActionResponse) SetManifests(v []Manifest) {
	o.Manifests = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *HelmRenderChartActionResponse) GetMetadata() HelmRenderMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret HelmRenderMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmRenderChartActionResponse) GetMetadataOk() (*HelmRenderMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *HelmRenderChartActionResponse) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given HelmRenderMetadata and assigns it to the Metadata field.
func (o *HelmRenderChartActionResponse) SetMetadata(v HelmRenderMetadata) {
	o.Metadata = &v
}

func (o HelmRenderChartActionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HelmRenderChartActionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Manifests) {
		toSerialize["manifests"] = o.Manifests
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HelmRenderChartActionResponse) UnmarshalJSON(data []byte) (err error) {
	varHelmRenderChartActionResponse := _HelmRenderChartActionResponse{}

	err = json.Unmarshal(data, &varHelmRenderChartActionResponse)

	if err != nil {
		return err
	}

	*o = HelmRenderChartActionResponse(varHelmRenderChartActionResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "manifests")
		delete(additionalProperties, "metadata")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHelmRenderChartActionResponse struct {
	value *HelmRenderChartActionResponse
	isSet bool
}

func (v NullableHelmRenderChartActionResponse) Get() *HelmRenderChartActionResponse {
	return v.value
}

func (v *NullableHelmRenderChartActionResponse) Set(val *HelmRenderChartActionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableHelmRenderChartActionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableHelmRenderChartActionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHelmRenderChartActionResponse(val *HelmRenderChartActionResponse) *NullableHelmRenderChartActionResponse {
	return &NullableHelmRenderChartActionResponse{value: val, isSet: true}
}

func (v NullableHelmRenderChartActionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHelmRenderChartActionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


