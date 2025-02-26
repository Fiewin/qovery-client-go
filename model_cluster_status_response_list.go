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

// ClusterStatusResponseList struct for ClusterStatusResponseList
type ClusterStatusResponseList struct {
	Id     string  `json:"id"`
	Status *string `json:"status,omitempty"`
}

// NewClusterStatusResponseList instantiates a new ClusterStatusResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterStatusResponseList(id string) *ClusterStatusResponseList {
	this := ClusterStatusResponseList{}
	this.Id = id
	return &this
}

// NewClusterStatusResponseListWithDefaults instantiates a new ClusterStatusResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterStatusResponseListWithDefaults() *ClusterStatusResponseList {
	this := ClusterStatusResponseList{}
	return &this
}

// GetId returns the Id field value
func (o *ClusterStatusResponseList) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ClusterStatusResponseList) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ClusterStatusResponseList) SetId(v string) {
	o.Id = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ClusterStatusResponseList) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterStatusResponseList) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ClusterStatusResponseList) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ClusterStatusResponseList) SetStatus(v string) {
	o.Status = &v
}

func (o ClusterStatusResponseList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableClusterStatusResponseList struct {
	value *ClusterStatusResponseList
	isSet bool
}

func (v NullableClusterStatusResponseList) Get() *ClusterStatusResponseList {
	return v.value
}

func (v *NullableClusterStatusResponseList) Set(val *ClusterStatusResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterStatusResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterStatusResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterStatusResponseList(val *ClusterStatusResponseList) *NullableClusterStatusResponseList {
	return &NullableClusterStatusResponseList{value: val, isSet: true}
}

func (v NullableClusterStatusResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterStatusResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
