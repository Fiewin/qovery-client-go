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

// MetricStorageResponse struct for MetricStorageResponse
type MetricStorageResponse struct {
	StorageId *string                          `json:"storage_id,omitempty"`
	Data      []MetricStorageDatapointResponse `json:"data"`
}

// NewMetricStorageResponse instantiates a new MetricStorageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricStorageResponse(data []MetricStorageDatapointResponse) *MetricStorageResponse {
	this := MetricStorageResponse{}
	this.Data = data
	return &this
}

// NewMetricStorageResponseWithDefaults instantiates a new MetricStorageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricStorageResponseWithDefaults() *MetricStorageResponse {
	this := MetricStorageResponse{}
	return &this
}

// GetStorageId returns the StorageId field value if set, zero value otherwise.
func (o *MetricStorageResponse) GetStorageId() string {
	if o == nil || o.StorageId == nil {
		var ret string
		return ret
	}
	return *o.StorageId
}

// GetStorageIdOk returns a tuple with the StorageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricStorageResponse) GetStorageIdOk() (*string, bool) {
	if o == nil || o.StorageId == nil {
		return nil, false
	}
	return o.StorageId, true
}

// HasStorageId returns a boolean if a field has been set.
func (o *MetricStorageResponse) HasStorageId() bool {
	if o != nil && o.StorageId != nil {
		return true
	}

	return false
}

// SetStorageId gets a reference to the given string and assigns it to the StorageId field.
func (o *MetricStorageResponse) SetStorageId(v string) {
	o.StorageId = &v
}

// GetData returns the Data field value
func (o *MetricStorageResponse) GetData() []MetricStorageDatapointResponse {
	if o == nil {
		var ret []MetricStorageDatapointResponse
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *MetricStorageResponse) GetDataOk() (*[]MetricStorageDatapointResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *MetricStorageResponse) SetData(v []MetricStorageDatapointResponse) {
	o.Data = v
}

func (o MetricStorageResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.StorageId != nil {
		toSerialize["storage_id"] = o.StorageId
	}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableMetricStorageResponse struct {
	value *MetricStorageResponse
	isSet bool
}

func (v NullableMetricStorageResponse) Get() *MetricStorageResponse {
	return v.value
}

func (v *NullableMetricStorageResponse) Set(val *MetricStorageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricStorageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricStorageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricStorageResponse(val *MetricStorageResponse) *NullableMetricStorageResponse {
	return &NullableMetricStorageResponse{value: val, isSet: true}
}

func (v NullableMetricStorageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricStorageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
