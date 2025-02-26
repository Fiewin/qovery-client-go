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

// ServiceTotalNumber struct for ServiceTotalNumber
type ServiceTotalNumber struct {
	ServiceTotalNumber *float32 `json:"service_total_number,omitempty"`
}

// NewServiceTotalNumber instantiates a new ServiceTotalNumber object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceTotalNumber() *ServiceTotalNumber {
	this := ServiceTotalNumber{}
	return &this
}

// NewServiceTotalNumberWithDefaults instantiates a new ServiceTotalNumber object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceTotalNumberWithDefaults() *ServiceTotalNumber {
	this := ServiceTotalNumber{}
	return &this
}

// GetServiceTotalNumber returns the ServiceTotalNumber field value if set, zero value otherwise.
func (o *ServiceTotalNumber) GetServiceTotalNumber() float32 {
	if o == nil || o.ServiceTotalNumber == nil {
		var ret float32
		return ret
	}
	return *o.ServiceTotalNumber
}

// GetServiceTotalNumberOk returns a tuple with the ServiceTotalNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceTotalNumber) GetServiceTotalNumberOk() (*float32, bool) {
	if o == nil || o.ServiceTotalNumber == nil {
		return nil, false
	}
	return o.ServiceTotalNumber, true
}

// HasServiceTotalNumber returns a boolean if a field has been set.
func (o *ServiceTotalNumber) HasServiceTotalNumber() bool {
	if o != nil && o.ServiceTotalNumber != nil {
		return true
	}

	return false
}

// SetServiceTotalNumber gets a reference to the given float32 and assigns it to the ServiceTotalNumber field.
func (o *ServiceTotalNumber) SetServiceTotalNumber(v float32) {
	o.ServiceTotalNumber = &v
}

func (o ServiceTotalNumber) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ServiceTotalNumber != nil {
		toSerialize["service_total_number"] = o.ServiceTotalNumber
	}
	return json.Marshal(toSerialize)
}

type NullableServiceTotalNumber struct {
	value *ServiceTotalNumber
	isSet bool
}

func (v NullableServiceTotalNumber) Get() *ServiceTotalNumber {
	return v.value
}

func (v *NullableServiceTotalNumber) Set(val *ServiceTotalNumber) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceTotalNumber) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceTotalNumber) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceTotalNumber(val *ServiceTotalNumber) *NullableServiceTotalNumber {
	return &NullableServiceTotalNumber{value: val, isSet: true}
}

func (v NullableServiceTotalNumber) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceTotalNumber) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
