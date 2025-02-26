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
	"time"
)

// EnvironmentLogResponse struct for EnvironmentLogResponse
type EnvironmentLogResponse struct {
	Id        string                       `json:"id"`
	CreatedAt time.Time                    `json:"created_at"`
	Scope     *EnvironmentLogResponseScope `json:"scope,omitempty"`
	// Status is a state machine. It starts with `BUILDING` or `DEPLOYING` state (or `INITIALIZED`if auto-deploy is deactivated). Then finish with `*_ERROR` or any termination state.
	State *string `json:"state,omitempty"`
	// Log message
	Message NullableString `json:"message"`
	// Only for errors. Helps Qovery team to investigate.
	ExecutionId *string `json:"execution_id,omitempty"`
	Hint        *string `json:"hint,omitempty"`
}

// NewEnvironmentLogResponse instantiates a new EnvironmentLogResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentLogResponse(id string, createdAt time.Time, message NullableString) *EnvironmentLogResponse {
	this := EnvironmentLogResponse{}
	this.Id = id
	this.CreatedAt = createdAt
	this.Message = message
	return &this
}

// NewEnvironmentLogResponseWithDefaults instantiates a new EnvironmentLogResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentLogResponseWithDefaults() *EnvironmentLogResponse {
	this := EnvironmentLogResponse{}
	return &this
}

// GetId returns the Id field value
func (o *EnvironmentLogResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EnvironmentLogResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EnvironmentLogResponse) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *EnvironmentLogResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *EnvironmentLogResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *EnvironmentLogResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *EnvironmentLogResponse) GetScope() EnvironmentLogResponseScope {
	if o == nil || o.Scope == nil {
		var ret EnvironmentLogResponseScope
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentLogResponse) GetScopeOk() (*EnvironmentLogResponseScope, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *EnvironmentLogResponse) HasScope() bool {
	if o != nil && o.Scope != nil {
		return true
	}

	return false
}

// SetScope gets a reference to the given EnvironmentLogResponseScope and assigns it to the Scope field.
func (o *EnvironmentLogResponse) SetScope(v EnvironmentLogResponseScope) {
	o.Scope = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *EnvironmentLogResponse) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentLogResponse) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *EnvironmentLogResponse) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *EnvironmentLogResponse) SetState(v string) {
	o.State = &v
}

// GetMessage returns the Message field value
// If the value is explicit nil, the zero value for string will be returned
func (o *EnvironmentLogResponse) GetMessage() string {
	if o == nil || o.Message.Get() == nil {
		var ret string
		return ret
	}

	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentLogResponse) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// SetMessage sets field value
func (o *EnvironmentLogResponse) SetMessage(v string) {
	o.Message.Set(&v)
}

// GetExecutionId returns the ExecutionId field value if set, zero value otherwise.
func (o *EnvironmentLogResponse) GetExecutionId() string {
	if o == nil || o.ExecutionId == nil {
		var ret string
		return ret
	}
	return *o.ExecutionId
}

// GetExecutionIdOk returns a tuple with the ExecutionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentLogResponse) GetExecutionIdOk() (*string, bool) {
	if o == nil || o.ExecutionId == nil {
		return nil, false
	}
	return o.ExecutionId, true
}

// HasExecutionId returns a boolean if a field has been set.
func (o *EnvironmentLogResponse) HasExecutionId() bool {
	if o != nil && o.ExecutionId != nil {
		return true
	}

	return false
}

// SetExecutionId gets a reference to the given string and assigns it to the ExecutionId field.
func (o *EnvironmentLogResponse) SetExecutionId(v string) {
	o.ExecutionId = &v
}

// GetHint returns the Hint field value if set, zero value otherwise.
func (o *EnvironmentLogResponse) GetHint() string {
	if o == nil || o.Hint == nil {
		var ret string
		return ret
	}
	return *o.Hint
}

// GetHintOk returns a tuple with the Hint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentLogResponse) GetHintOk() (*string, bool) {
	if o == nil || o.Hint == nil {
		return nil, false
	}
	return o.Hint, true
}

// HasHint returns a boolean if a field has been set.
func (o *EnvironmentLogResponse) HasHint() bool {
	if o != nil && o.Hint != nil {
		return true
	}

	return false
}

// SetHint gets a reference to the given string and assigns it to the Hint field.
func (o *EnvironmentLogResponse) SetHint(v string) {
	o.Hint = &v
}

func (o EnvironmentLogResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if true {
		toSerialize["message"] = o.Message.Get()
	}
	if o.ExecutionId != nil {
		toSerialize["execution_id"] = o.ExecutionId
	}
	if o.Hint != nil {
		toSerialize["hint"] = o.Hint
	}
	return json.Marshal(toSerialize)
}

type NullableEnvironmentLogResponse struct {
	value *EnvironmentLogResponse
	isSet bool
}

func (v NullableEnvironmentLogResponse) Get() *EnvironmentLogResponse {
	return v.value
}

func (v *NullableEnvironmentLogResponse) Set(val *EnvironmentLogResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentLogResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentLogResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentLogResponse(val *EnvironmentLogResponse) *NullableEnvironmentLogResponse {
	return &NullableEnvironmentLogResponse{value: val, isSet: true}
}

func (v NullableEnvironmentLogResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentLogResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
