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

// ApplicationPortRequest struct for ApplicationPortRequest
type ApplicationPortRequest struct {
	Ports *[]ApplicationPortRequestPorts `json:"ports,omitempty"`
}

// NewApplicationPortRequest instantiates a new ApplicationPortRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationPortRequest() *ApplicationPortRequest {
	this := ApplicationPortRequest{}
	return &this
}

// NewApplicationPortRequestWithDefaults instantiates a new ApplicationPortRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationPortRequestWithDefaults() *ApplicationPortRequest {
	this := ApplicationPortRequest{}
	return &this
}

// GetPorts returns the Ports field value if set, zero value otherwise.
func (o *ApplicationPortRequest) GetPorts() []ApplicationPortRequestPorts {
	if o == nil || o.Ports == nil {
		var ret []ApplicationPortRequestPorts
		return ret
	}
	return *o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationPortRequest) GetPortsOk() (*[]ApplicationPortRequestPorts, bool) {
	if o == nil || o.Ports == nil {
		return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *ApplicationPortRequest) HasPorts() bool {
	if o != nil && o.Ports != nil {
		return true
	}

	return false
}

// SetPorts gets a reference to the given []ApplicationPortRequestPorts and assigns it to the Ports field.
func (o *ApplicationPortRequest) SetPorts(v []ApplicationPortRequestPorts) {
	o.Ports = &v
}

func (o ApplicationPortRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Ports != nil {
		toSerialize["ports"] = o.Ports
	}
	return json.Marshal(toSerialize)
}

type NullableApplicationPortRequest struct {
	value *ApplicationPortRequest
	isSet bool
}

func (v NullableApplicationPortRequest) Get() *ApplicationPortRequest {
	return v.value
}

func (v *NullableApplicationPortRequest) Set(val *ApplicationPortRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationPortRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationPortRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationPortRequest(val *ApplicationPortRequest) *NullableApplicationPortRequest {
	return &NullableApplicationPortRequest{value: val, isSet: true}
}

func (v NullableApplicationPortRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationPortRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
