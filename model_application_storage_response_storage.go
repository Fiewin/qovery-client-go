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

// ApplicationStorageResponseStorage struct for ApplicationStorageResponseStorage
type ApplicationStorageResponseStorage struct {
	Id   *string `json:"id,omitempty"`
	Type string  `json:"type"`
	// unit is GB
	Size       float32 `json:"size"`
	MountPoint string  `json:"mount_point"`
}

// NewApplicationStorageResponseStorage instantiates a new ApplicationStorageResponseStorage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationStorageResponseStorage(type_ string, size float32, mountPoint string) *ApplicationStorageResponseStorage {
	this := ApplicationStorageResponseStorage{}
	this.Type = type_
	this.Size = size
	this.MountPoint = mountPoint
	return &this
}

// NewApplicationStorageResponseStorageWithDefaults instantiates a new ApplicationStorageResponseStorage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationStorageResponseStorageWithDefaults() *ApplicationStorageResponseStorage {
	this := ApplicationStorageResponseStorage{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApplicationStorageResponseStorage) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationStorageResponseStorage) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApplicationStorageResponseStorage) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApplicationStorageResponseStorage) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value
func (o *ApplicationStorageResponseStorage) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ApplicationStorageResponseStorage) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ApplicationStorageResponseStorage) SetType(v string) {
	o.Type = v
}

// GetSize returns the Size field value
func (o *ApplicationStorageResponseStorage) GetSize() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *ApplicationStorageResponseStorage) GetSizeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *ApplicationStorageResponseStorage) SetSize(v float32) {
	o.Size = v
}

// GetMountPoint returns the MountPoint field value
func (o *ApplicationStorageResponseStorage) GetMountPoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MountPoint
}

// GetMountPointOk returns a tuple with the MountPoint field value
// and a boolean to check if the value has been set.
func (o *ApplicationStorageResponseStorage) GetMountPointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MountPoint, true
}

// SetMountPoint sets field value
func (o *ApplicationStorageResponseStorage) SetMountPoint(v string) {
	o.MountPoint = v
}

func (o ApplicationStorageResponseStorage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["size"] = o.Size
	}
	if true {
		toSerialize["mount_point"] = o.MountPoint
	}
	return json.Marshal(toSerialize)
}

type NullableApplicationStorageResponseStorage struct {
	value *ApplicationStorageResponseStorage
	isSet bool
}

func (v NullableApplicationStorageResponseStorage) Get() *ApplicationStorageResponseStorage {
	return v.value
}

func (v *NullableApplicationStorageResponseStorage) Set(val *ApplicationStorageResponseStorage) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationStorageResponseStorage) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationStorageResponseStorage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationStorageResponseStorage(val *ApplicationStorageResponseStorage) *NullableApplicationStorageResponseStorage {
	return &NullableApplicationStorageResponseStorage{value: val, isSet: true}
}

func (v NullableApplicationStorageResponseStorage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationStorageResponseStorage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
