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

// ReferralResponse struct for ReferralResponse
type ReferralResponse struct {
	TotalInvited   *int32  `json:"total_invited,omitempty"`
	InvitationLink *string `json:"invitation_link,omitempty"`
}

// NewReferralResponse instantiates a new ReferralResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReferralResponse() *ReferralResponse {
	this := ReferralResponse{}
	return &this
}

// NewReferralResponseWithDefaults instantiates a new ReferralResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReferralResponseWithDefaults() *ReferralResponse {
	this := ReferralResponse{}
	return &this
}

// GetTotalInvited returns the TotalInvited field value if set, zero value otherwise.
func (o *ReferralResponse) GetTotalInvited() int32 {
	if o == nil || o.TotalInvited == nil {
		var ret int32
		return ret
	}
	return *o.TotalInvited
}

// GetTotalInvitedOk returns a tuple with the TotalInvited field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferralResponse) GetTotalInvitedOk() (*int32, bool) {
	if o == nil || o.TotalInvited == nil {
		return nil, false
	}
	return o.TotalInvited, true
}

// HasTotalInvited returns a boolean if a field has been set.
func (o *ReferralResponse) HasTotalInvited() bool {
	if o != nil && o.TotalInvited != nil {
		return true
	}

	return false
}

// SetTotalInvited gets a reference to the given int32 and assigns it to the TotalInvited field.
func (o *ReferralResponse) SetTotalInvited(v int32) {
	o.TotalInvited = &v
}

// GetInvitationLink returns the InvitationLink field value if set, zero value otherwise.
func (o *ReferralResponse) GetInvitationLink() string {
	if o == nil || o.InvitationLink == nil {
		var ret string
		return ret
	}
	return *o.InvitationLink
}

// GetInvitationLinkOk returns a tuple with the InvitationLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferralResponse) GetInvitationLinkOk() (*string, bool) {
	if o == nil || o.InvitationLink == nil {
		return nil, false
	}
	return o.InvitationLink, true
}

// HasInvitationLink returns a boolean if a field has been set.
func (o *ReferralResponse) HasInvitationLink() bool {
	if o != nil && o.InvitationLink != nil {
		return true
	}

	return false
}

// SetInvitationLink gets a reference to the given string and assigns it to the InvitationLink field.
func (o *ReferralResponse) SetInvitationLink(v string) {
	o.InvitationLink = &v
}

func (o ReferralResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TotalInvited != nil {
		toSerialize["total_invited"] = o.TotalInvited
	}
	if o.InvitationLink != nil {
		toSerialize["invitation_link"] = o.InvitationLink
	}
	return json.Marshal(toSerialize)
}

type NullableReferralResponse struct {
	value *ReferralResponse
	isSet bool
}

func (v NullableReferralResponse) Get() *ReferralResponse {
	return v.value
}

func (v *NullableReferralResponse) Set(val *ReferralResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReferralResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReferralResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReferralResponse(val *ReferralResponse) *NullableReferralResponse {
	return &NullableReferralResponse{value: val, isSet: true}
}

func (v NullableReferralResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReferralResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
