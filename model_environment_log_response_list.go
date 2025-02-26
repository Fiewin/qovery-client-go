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

// EnvironmentLogResponseList struct for EnvironmentLogResponseList
type EnvironmentLogResponseList struct {
	Results *[]EnvironmentLogResponse `json:"results,omitempty"`
}

// NewEnvironmentLogResponseList instantiates a new EnvironmentLogResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentLogResponseList() *EnvironmentLogResponseList {
	this := EnvironmentLogResponseList{}
	return &this
}

// NewEnvironmentLogResponseListWithDefaults instantiates a new EnvironmentLogResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentLogResponseListWithDefaults() *EnvironmentLogResponseList {
	this := EnvironmentLogResponseList{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *EnvironmentLogResponseList) GetResults() []EnvironmentLogResponse {
	if o == nil || o.Results == nil {
		var ret []EnvironmentLogResponse
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentLogResponseList) GetResultsOk() (*[]EnvironmentLogResponse, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *EnvironmentLogResponseList) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []EnvironmentLogResponse and assigns it to the Results field.
func (o *EnvironmentLogResponseList) SetResults(v []EnvironmentLogResponse) {
	o.Results = &v
}

func (o EnvironmentLogResponseList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableEnvironmentLogResponseList struct {
	value *EnvironmentLogResponseList
	isSet bool
}

func (v NullableEnvironmentLogResponseList) Get() *EnvironmentLogResponseList {
	return v.value
}

func (v *NullableEnvironmentLogResponseList) Set(val *EnvironmentLogResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentLogResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentLogResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentLogResponseList(val *EnvironmentLogResponseList) *NullableEnvironmentLogResponseList {
	return &NullableEnvironmentLogResponseList{value: val, isSet: true}
}

func (v NullableEnvironmentLogResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentLogResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
