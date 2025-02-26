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

// DatabaseRequest struct for DatabaseRequest
type DatabaseRequest struct {
	// name is case insensitive
	Name          string  `json:"name"`
	Type          string  `json:"type"`
	Version       string  `json:"version"`
	Mode          string  `json:"mode"`
	Accessibility *string `json:"accessibility,omitempty"`
	// unit is millicores (m). 1000m = 1 cpu
	Cpu *float32 `json:"cpu,omitempty"`
	// unit is MB. 1024 MB = 1GB
	Memory *float32 `json:"memory,omitempty"`
	// unit is MB
	Storage *float32 `json:"storage,omitempty"`
}

// NewDatabaseRequest instantiates a new DatabaseRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatabaseRequest(name string, type_ string, version string, mode string) *DatabaseRequest {
	this := DatabaseRequest{}
	this.Name = name
	this.Type = type_
	this.Version = version
	this.Mode = mode
	var accessibility string = "PRIVATE"
	this.Accessibility = &accessibility
	var cpu float32 = 250
	this.Cpu = &cpu
	var memory float32 = 256
	this.Memory = &memory
	var storage float32 = 10240
	this.Storage = &storage
	return &this
}

// NewDatabaseRequestWithDefaults instantiates a new DatabaseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabaseRequestWithDefaults() *DatabaseRequest {
	this := DatabaseRequest{}
	var accessibility string = "PRIVATE"
	this.Accessibility = &accessibility
	var cpu float32 = 250
	this.Cpu = &cpu
	var memory float32 = 256
	this.Memory = &memory
	var storage float32 = 10240
	this.Storage = &storage
	return &this
}

// GetName returns the Name field value
func (o *DatabaseRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DatabaseRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DatabaseRequest) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *DatabaseRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DatabaseRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DatabaseRequest) SetType(v string) {
	o.Type = v
}

// GetVersion returns the Version field value
func (o *DatabaseRequest) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *DatabaseRequest) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *DatabaseRequest) SetVersion(v string) {
	o.Version = v
}

// GetMode returns the Mode field value
func (o *DatabaseRequest) GetMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *DatabaseRequest) GetModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value
func (o *DatabaseRequest) SetMode(v string) {
	o.Mode = v
}

// GetAccessibility returns the Accessibility field value if set, zero value otherwise.
func (o *DatabaseRequest) GetAccessibility() string {
	if o == nil || o.Accessibility == nil {
		var ret string
		return ret
	}
	return *o.Accessibility
}

// GetAccessibilityOk returns a tuple with the Accessibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseRequest) GetAccessibilityOk() (*string, bool) {
	if o == nil || o.Accessibility == nil {
		return nil, false
	}
	return o.Accessibility, true
}

// HasAccessibility returns a boolean if a field has been set.
func (o *DatabaseRequest) HasAccessibility() bool {
	if o != nil && o.Accessibility != nil {
		return true
	}

	return false
}

// SetAccessibility gets a reference to the given string and assigns it to the Accessibility field.
func (o *DatabaseRequest) SetAccessibility(v string) {
	o.Accessibility = &v
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *DatabaseRequest) GetCpu() float32 {
	if o == nil || o.Cpu == nil {
		var ret float32
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseRequest) GetCpuOk() (*float32, bool) {
	if o == nil || o.Cpu == nil {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *DatabaseRequest) HasCpu() bool {
	if o != nil && o.Cpu != nil {
		return true
	}

	return false
}

// SetCpu gets a reference to the given float32 and assigns it to the Cpu field.
func (o *DatabaseRequest) SetCpu(v float32) {
	o.Cpu = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *DatabaseRequest) GetMemory() float32 {
	if o == nil || o.Memory == nil {
		var ret float32
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseRequest) GetMemoryOk() (*float32, bool) {
	if o == nil || o.Memory == nil {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *DatabaseRequest) HasMemory() bool {
	if o != nil && o.Memory != nil {
		return true
	}

	return false
}

// SetMemory gets a reference to the given float32 and assigns it to the Memory field.
func (o *DatabaseRequest) SetMemory(v float32) {
	o.Memory = &v
}

// GetStorage returns the Storage field value if set, zero value otherwise.
func (o *DatabaseRequest) GetStorage() float32 {
	if o == nil || o.Storage == nil {
		var ret float32
		return ret
	}
	return *o.Storage
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseRequest) GetStorageOk() (*float32, bool) {
	if o == nil || o.Storage == nil {
		return nil, false
	}
	return o.Storage, true
}

// HasStorage returns a boolean if a field has been set.
func (o *DatabaseRequest) HasStorage() bool {
	if o != nil && o.Storage != nil {
		return true
	}

	return false
}

// SetStorage gets a reference to the given float32 and assigns it to the Storage field.
func (o *DatabaseRequest) SetStorage(v float32) {
	o.Storage = &v
}

func (o DatabaseRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["version"] = o.Version
	}
	if true {
		toSerialize["mode"] = o.Mode
	}
	if o.Accessibility != nil {
		toSerialize["accessibility"] = o.Accessibility
	}
	if o.Cpu != nil {
		toSerialize["cpu"] = o.Cpu
	}
	if o.Memory != nil {
		toSerialize["memory"] = o.Memory
	}
	if o.Storage != nil {
		toSerialize["storage"] = o.Storage
	}
	return json.Marshal(toSerialize)
}

type NullableDatabaseRequest struct {
	value *DatabaseRequest
	isSet bool
}

func (v NullableDatabaseRequest) Get() *DatabaseRequest {
	return v.value
}

func (v *NullableDatabaseRequest) Set(val *DatabaseRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDatabaseRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDatabaseRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatabaseRequest(val *DatabaseRequest) *NullableDatabaseRequest {
	return &NullableDatabaseRequest{value: val, isSet: true}
}

func (v NullableDatabaseRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatabaseRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
