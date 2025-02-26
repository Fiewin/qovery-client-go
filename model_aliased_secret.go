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

// AliasedSecret struct for AliasedSecret
type AliasedSecret struct {
	Id    *string `json:"id,omitempty"`
	Key   *string `json:"key,omitempty"`
	Scope *string `json:"scope,omitempty"`
}

// NewAliasedSecret instantiates a new AliasedSecret object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAliasedSecret() *AliasedSecret {
	this := AliasedSecret{}
	return &this
}

// NewAliasedSecretWithDefaults instantiates a new AliasedSecret object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAliasedSecretWithDefaults() *AliasedSecret {
	this := AliasedSecret{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AliasedSecret) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AliasedSecret) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AliasedSecret) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AliasedSecret) SetId(v string) {
	o.Id = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *AliasedSecret) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AliasedSecret) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *AliasedSecret) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *AliasedSecret) SetKey(v string) {
	o.Key = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *AliasedSecret) GetScope() string {
	if o == nil || o.Scope == nil {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AliasedSecret) GetScopeOk() (*string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *AliasedSecret) HasScope() bool {
	if o != nil && o.Scope != nil {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *AliasedSecret) SetScope(v string) {
	o.Scope = &v
}

func (o AliasedSecret) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	return json.Marshal(toSerialize)
}

type NullableAliasedSecret struct {
	value *AliasedSecret
	isSet bool
}

func (v NullableAliasedSecret) Get() *AliasedSecret {
	return v.value
}

func (v *NullableAliasedSecret) Set(val *AliasedSecret) {
	v.value = val
	v.isSet = true
}

func (v NullableAliasedSecret) IsSet() bool {
	return v.isSet
}

func (v *NullableAliasedSecret) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAliasedSecret(val *AliasedSecret) *NullableAliasedSecret {
	return &NullableAliasedSecret{value: val, isSet: true}
}

func (v NullableAliasedSecret) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAliasedSecret) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
