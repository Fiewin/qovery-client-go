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

// EnvironmentLogResponseScope struct for EnvironmentLogResponseScope
type EnvironmentLogResponseScope struct {
	Type *string `json:"type,omitempty"`
	Name *string `json:"name,omitempty"`
	Id   *string `json:"id,omitempty"`
}

// NewEnvironmentLogResponseScope instantiates a new EnvironmentLogResponseScope object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentLogResponseScope() *EnvironmentLogResponseScope {
	this := EnvironmentLogResponseScope{}
	return &this
}

// NewEnvironmentLogResponseScopeWithDefaults instantiates a new EnvironmentLogResponseScope object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentLogResponseScopeWithDefaults() *EnvironmentLogResponseScope {
	this := EnvironmentLogResponseScope{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *EnvironmentLogResponseScope) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentLogResponseScope) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *EnvironmentLogResponseScope) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *EnvironmentLogResponseScope) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EnvironmentLogResponseScope) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentLogResponseScope) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EnvironmentLogResponseScope) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EnvironmentLogResponseScope) SetName(v string) {
	o.Name = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EnvironmentLogResponseScope) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentLogResponseScope) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EnvironmentLogResponseScope) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EnvironmentLogResponseScope) SetId(v string) {
	o.Id = &v
}

func (o EnvironmentLogResponseScope) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableEnvironmentLogResponseScope struct {
	value *EnvironmentLogResponseScope
	isSet bool
}

func (v NullableEnvironmentLogResponseScope) Get() *EnvironmentLogResponseScope {
	return v.value
}

func (v *NullableEnvironmentLogResponseScope) Set(val *EnvironmentLogResponseScope) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentLogResponseScope) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentLogResponseScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentLogResponseScope(val *EnvironmentLogResponseScope) *NullableEnvironmentLogResponseScope {
	return &NullableEnvironmentLogResponseScope{value: val, isSet: true}
}

func (v NullableEnvironmentLogResponseScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentLogResponseScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
