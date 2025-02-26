/*
 * [BETA] Qovery API
 *
 * - Qovery is the fastest way to deploy your full-stack apps on any Cloud provider. - ℹ️ The API is in Beta and still in progress. Some endpoints are not available yet.
 *
 * API version: 1.0.0
 * Contact: support+api+documentation@qovery.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qovery

import (
	"encoding/json"
)

// ProjectResponseList struct for ProjectResponseList
type ProjectResponseList struct {
	Results *[]ProjectResponse `json:"results,omitempty"`
}

// NewProjectResponseList instantiates a new ProjectResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectResponseList() *ProjectResponseList {
	this := ProjectResponseList{}
	return &this
}

// NewProjectResponseListWithDefaults instantiates a new ProjectResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectResponseListWithDefaults() *ProjectResponseList {
	this := ProjectResponseList{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *ProjectResponseList) GetResults() []ProjectResponse {
	if o == nil || o.Results == nil {
		var ret []ProjectResponse
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectResponseList) GetResultsOk() (*[]ProjectResponse, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ProjectResponseList) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []ProjectResponse and assigns it to the Results field.
func (o *ProjectResponseList) SetResults(v []ProjectResponse) {
	o.Results = &v
}

func (o ProjectResponseList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableProjectResponseList struct {
	value *ProjectResponseList
	isSet bool
}

func (v NullableProjectResponseList) Get() *ProjectResponseList {
	return v.value
}

func (v *NullableProjectResponseList) Set(val *ProjectResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectResponseList(val *ProjectResponseList) *NullableProjectResponseList {
	return &NullableProjectResponseList{value: val, isSet: true}
}

func (v NullableProjectResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
