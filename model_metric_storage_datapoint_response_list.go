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

// MetricStorageDatapointResponseList struct for MetricStorageDatapointResponseList
type MetricStorageDatapointResponseList struct {
	Results *[]MetricStorageDatapointResponse `json:"results,omitempty"`
}

// NewMetricStorageDatapointResponseList instantiates a new MetricStorageDatapointResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricStorageDatapointResponseList() *MetricStorageDatapointResponseList {
	this := MetricStorageDatapointResponseList{}
	return &this
}

// NewMetricStorageDatapointResponseListWithDefaults instantiates a new MetricStorageDatapointResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricStorageDatapointResponseListWithDefaults() *MetricStorageDatapointResponseList {
	this := MetricStorageDatapointResponseList{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *MetricStorageDatapointResponseList) GetResults() []MetricStorageDatapointResponse {
	if o == nil || o.Results == nil {
		var ret []MetricStorageDatapointResponse
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricStorageDatapointResponseList) GetResultsOk() (*[]MetricStorageDatapointResponse, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *MetricStorageDatapointResponseList) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []MetricStorageDatapointResponse and assigns it to the Results field.
func (o *MetricStorageDatapointResponseList) SetResults(v []MetricStorageDatapointResponse) {
	o.Results = &v
}

func (o MetricStorageDatapointResponseList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableMetricStorageDatapointResponseList struct {
	value *MetricStorageDatapointResponseList
	isSet bool
}

func (v NullableMetricStorageDatapointResponseList) Get() *MetricStorageDatapointResponseList {
	return v.value
}

func (v *NullableMetricStorageDatapointResponseList) Set(val *MetricStorageDatapointResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricStorageDatapointResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricStorageDatapointResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricStorageDatapointResponseList(val *MetricStorageDatapointResponseList) *NullableMetricStorageDatapointResponseList {
	return &NullableMetricStorageDatapointResponseList{value: val, isSet: true}
}

func (v NullableMetricStorageDatapointResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricStorageDatapointResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
