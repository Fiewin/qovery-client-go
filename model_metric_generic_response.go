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

// MetricGenericResponse struct for MetricGenericResponse
type MetricGenericResponse struct {
	InstanceName string                           `json:"instance_name"`
	Data         []MetricGenericDatapointResponse `json:"data"`
}

// NewMetricGenericResponse instantiates a new MetricGenericResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricGenericResponse(instanceName string, data []MetricGenericDatapointResponse) *MetricGenericResponse {
	this := MetricGenericResponse{}
	this.InstanceName = instanceName
	this.Data = data
	return &this
}

// NewMetricGenericResponseWithDefaults instantiates a new MetricGenericResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricGenericResponseWithDefaults() *MetricGenericResponse {
	this := MetricGenericResponse{}
	return &this
}

// GetInstanceName returns the InstanceName field value
func (o *MetricGenericResponse) GetInstanceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstanceName
}

// GetInstanceNameOk returns a tuple with the InstanceName field value
// and a boolean to check if the value has been set.
func (o *MetricGenericResponse) GetInstanceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceName, true
}

// SetInstanceName sets field value
func (o *MetricGenericResponse) SetInstanceName(v string) {
	o.InstanceName = v
}

// GetData returns the Data field value
func (o *MetricGenericResponse) GetData() []MetricGenericDatapointResponse {
	if o == nil {
		var ret []MetricGenericDatapointResponse
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *MetricGenericResponse) GetDataOk() (*[]MetricGenericDatapointResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *MetricGenericResponse) SetData(v []MetricGenericDatapointResponse) {
	o.Data = v
}

func (o MetricGenericResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["instance_name"] = o.InstanceName
	}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableMetricGenericResponse struct {
	value *MetricGenericResponse
	isSet bool
}

func (v NullableMetricGenericResponse) Get() *MetricGenericResponse {
	return v.value
}

func (v *NullableMetricGenericResponse) Set(val *MetricGenericResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricGenericResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricGenericResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricGenericResponse(val *MetricGenericResponse) *NullableMetricGenericResponse {
	return &NullableMetricGenericResponse{value: val, isSet: true}
}

func (v NullableMetricGenericResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricGenericResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
