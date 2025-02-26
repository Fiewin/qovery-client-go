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

// BillingPeriod struct for BillingPeriod
type BillingPeriod struct {
	BillingStartedOn *time.Time `json:"billing_started_on,omitempty"`
	BillingEndedOn   *time.Time `json:"billing_ended_on,omitempty"`
}

// NewBillingPeriod instantiates a new BillingPeriod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingPeriod() *BillingPeriod {
	this := BillingPeriod{}
	return &this
}

// NewBillingPeriodWithDefaults instantiates a new BillingPeriod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingPeriodWithDefaults() *BillingPeriod {
	this := BillingPeriod{}
	return &this
}

// GetBillingStartedOn returns the BillingStartedOn field value if set, zero value otherwise.
func (o *BillingPeriod) GetBillingStartedOn() time.Time {
	if o == nil || o.BillingStartedOn == nil {
		var ret time.Time
		return ret
	}
	return *o.BillingStartedOn
}

// GetBillingStartedOnOk returns a tuple with the BillingStartedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingPeriod) GetBillingStartedOnOk() (*time.Time, bool) {
	if o == nil || o.BillingStartedOn == nil {
		return nil, false
	}
	return o.BillingStartedOn, true
}

// HasBillingStartedOn returns a boolean if a field has been set.
func (o *BillingPeriod) HasBillingStartedOn() bool {
	if o != nil && o.BillingStartedOn != nil {
		return true
	}

	return false
}

// SetBillingStartedOn gets a reference to the given time.Time and assigns it to the BillingStartedOn field.
func (o *BillingPeriod) SetBillingStartedOn(v time.Time) {
	o.BillingStartedOn = &v
}

// GetBillingEndedOn returns the BillingEndedOn field value if set, zero value otherwise.
func (o *BillingPeriod) GetBillingEndedOn() time.Time {
	if o == nil || o.BillingEndedOn == nil {
		var ret time.Time
		return ret
	}
	return *o.BillingEndedOn
}

// GetBillingEndedOnOk returns a tuple with the BillingEndedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingPeriod) GetBillingEndedOnOk() (*time.Time, bool) {
	if o == nil || o.BillingEndedOn == nil {
		return nil, false
	}
	return o.BillingEndedOn, true
}

// HasBillingEndedOn returns a boolean if a field has been set.
func (o *BillingPeriod) HasBillingEndedOn() bool {
	if o != nil && o.BillingEndedOn != nil {
		return true
	}

	return false
}

// SetBillingEndedOn gets a reference to the given time.Time and assigns it to the BillingEndedOn field.
func (o *BillingPeriod) SetBillingEndedOn(v time.Time) {
	o.BillingEndedOn = &v
}

func (o BillingPeriod) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BillingStartedOn != nil {
		toSerialize["billing_started_on"] = o.BillingStartedOn
	}
	if o.BillingEndedOn != nil {
		toSerialize["billing_ended_on"] = o.BillingEndedOn
	}
	return json.Marshal(toSerialize)
}

type NullableBillingPeriod struct {
	value *BillingPeriod
	isSet bool
}

func (v NullableBillingPeriod) Get() *BillingPeriod {
	return v.value
}

func (v *NullableBillingPeriod) Set(val *BillingPeriod) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingPeriod) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingPeriod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingPeriod(val *BillingPeriod) *NullableBillingPeriod {
	return &NullableBillingPeriod{value: val, isSet: true}
}

func (v NullableBillingPeriod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingPeriod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
