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

// ApplicationRequest struct for ApplicationRequest
type ApplicationRequest struct {
	// name is case insensitive
	Name string `json:"name"`
	// give a description to this application
	Description   *string                         `json:"description,omitempty"`
	GitRepository ApplicationGitRepositoryRequest `json:"git_repository"`
	// `DOCKER` requires `dockerfile_path` `BUILDPACKS` does not require any `dockerfile_path`
	BuildMode *string `json:"build_mode,omitempty"`
	// The path of the associated Dockerfile. Only if you are using build_mode = DOCKER
	DockerfilePath *string `json:"dockerfile_path,omitempty"`
	// unit is millicores (m). 1000m = 1 cpu
	Cpu *float32 `json:"cpu,omitempty"`
	// unit is MB. 1024 MB = 1GB
	Memory *float32 `json:"memory,omitempty"`
	// Minimum number of instances running. This resource auto-scale based on the CPU and Memory consumption. Note: 0 means that there is no application running.
	MinRunningInstances *int32 `json:"min_running_instances,omitempty"`
	// Maximum number of instances running. This resource auto-scale based on the CPU and Memory consumption. Note: -1 means that there is no limit.
	MaxRunningInstances *int32                              `json:"max_running_instances,omitempty"`
	Healthcheck         *Healthcheck                        `json:"healthcheck,omitempty"`
	Storage             *[]ApplicationStorageRequestStorage `json:"storage,omitempty"`
	Ports               *[]ApplicationPortRequestPorts      `json:"ports,omitempty"`
}

// NewApplicationRequest instantiates a new ApplicationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationRequest(name string, gitRepository ApplicationGitRepositoryRequest) *ApplicationRequest {
	this := ApplicationRequest{}
	return &this
}

// NewApplicationRequestWithDefaults instantiates a new ApplicationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationRequestWithDefaults() *ApplicationRequest {
	this := ApplicationRequest{}
	var buildMode string = "BUILDPACKS"
	this.BuildMode = &buildMode
	var cpu float32 = 250
	this.Cpu = &cpu
	var memory float32 = 256
	this.Memory = &memory
	var minRunningInstances int32 = 1
	this.MinRunningInstances = &minRunningInstances
	var maxRunningInstances int32 = 1
	this.MaxRunningInstances = &maxRunningInstances
	return &this
}

// GetName returns the Name field value
func (o *ApplicationRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApplicationRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApplicationRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ApplicationRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ApplicationRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ApplicationRequest) SetDescription(v string) {
	o.Description = &v
}

// GetGitRepository returns the GitRepository field value
func (o *ApplicationRequest) GetGitRepository() ApplicationGitRepositoryRequest {
	if o == nil {
		var ret ApplicationGitRepositoryRequest
		return ret
	}

	return o.GitRepository
}

// GetGitRepositoryOk returns a tuple with the GitRepository field value
// and a boolean to check if the value has been set.
func (o *ApplicationRequest) GetGitRepositoryOk() (*ApplicationGitRepositoryRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GitRepository, true
}

// SetGitRepository sets field value
func (o *ApplicationRequest) SetGitRepository(v ApplicationGitRepositoryRequest) {
	o.GitRepository = v
}

// GetBuildMode returns the BuildMode field value if set, zero value otherwise.
func (o *ApplicationRequest) GetBuildMode() string {
	if o == nil || o.BuildMode == nil {
		var ret string
		return ret
	}
	return *o.BuildMode
}

// GetBuildModeOk returns a tuple with the BuildMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationRequest) GetBuildModeOk() (*string, bool) {
	if o == nil || o.BuildMode == nil {
		return nil, false
	}
	return o.BuildMode, true
}

// HasBuildMode returns a boolean if a field has been set.
func (o *ApplicationRequest) HasBuildMode() bool {
	if o != nil && o.BuildMode != nil {
		return true
	}

	return false
}

// SetBuildMode gets a reference to the given string and assigns it to the BuildMode field.
func (o *ApplicationRequest) SetBuildMode(v string) {
	o.BuildMode = &v
}

// GetDockerfilePath returns the DockerfilePath field value if set, zero value otherwise.
func (o *ApplicationRequest) GetDockerfilePath() string {
	if o == nil || o.DockerfilePath == nil {
		var ret string
		return ret
	}
	return *o.DockerfilePath
}

// GetDockerfilePathOk returns a tuple with the DockerfilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationRequest) GetDockerfilePathOk() (*string, bool) {
	if o == nil || o.DockerfilePath == nil {
		return nil, false
	}
	return o.DockerfilePath, true
}

// HasDockerfilePath returns a boolean if a field has been set.
func (o *ApplicationRequest) HasDockerfilePath() bool {
	if o != nil && o.DockerfilePath != nil {
		return true
	}

	return false
}

// SetDockerfilePath gets a reference to the given string and assigns it to the DockerfilePath field.
func (o *ApplicationRequest) SetDockerfilePath(v string) {
	o.DockerfilePath = &v
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *ApplicationRequest) GetCpu() float32 {
	if o == nil || o.Cpu == nil {
		var ret float32
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationRequest) GetCpuOk() (*float32, bool) {
	if o == nil || o.Cpu == nil {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *ApplicationRequest) HasCpu() bool {
	if o != nil && o.Cpu != nil {
		return true
	}

	return false
}

// SetCpu gets a reference to the given float32 and assigns it to the Cpu field.
func (o *ApplicationRequest) SetCpu(v float32) {
	o.Cpu = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *ApplicationRequest) GetMemory() float32 {
	if o == nil || o.Memory == nil {
		var ret float32
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationRequest) GetMemoryOk() (*float32, bool) {
	if o == nil || o.Memory == nil {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *ApplicationRequest) HasMemory() bool {
	if o != nil && o.Memory != nil {
		return true
	}

	return false
}

// SetMemory gets a reference to the given float32 and assigns it to the Memory field.
func (o *ApplicationRequest) SetMemory(v float32) {
	o.Memory = &v
}

// GetMinRunningInstances returns the MinRunningInstances field value if set, zero value otherwise.
func (o *ApplicationRequest) GetMinRunningInstances() int32 {
	if o == nil || o.MinRunningInstances == nil {
		var ret int32
		return ret
	}
	return *o.MinRunningInstances
}

// GetMinRunningInstancesOk returns a tuple with the MinRunningInstances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationRequest) GetMinRunningInstancesOk() (*int32, bool) {
	if o == nil || o.MinRunningInstances == nil {
		return nil, false
	}
	return o.MinRunningInstances, true
}

// HasMinRunningInstances returns a boolean if a field has been set.
func (o *ApplicationRequest) HasMinRunningInstances() bool {
	if o != nil && o.MinRunningInstances != nil {
		return true
	}

	return false
}

// SetMinRunningInstances gets a reference to the given int32 and assigns it to the MinRunningInstances field.
func (o *ApplicationRequest) SetMinRunningInstances(v int32) {
	o.MinRunningInstances = &v
}

// GetMaxRunningInstances returns the MaxRunningInstances field value if set, zero value otherwise.
func (o *ApplicationRequest) GetMaxRunningInstances() int32 {
	if o == nil || o.MaxRunningInstances == nil {
		var ret int32
		return ret
	}
	return *o.MaxRunningInstances
}

// GetMaxRunningInstancesOk returns a tuple with the MaxRunningInstances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationRequest) GetMaxRunningInstancesOk() (*int32, bool) {
	if o == nil || o.MaxRunningInstances == nil {
		return nil, false
	}
	return o.MaxRunningInstances, true
}

// HasMaxRunningInstances returns a boolean if a field has been set.
func (o *ApplicationRequest) HasMaxRunningInstances() bool {
	if o != nil && o.MaxRunningInstances != nil {
		return true
	}

	return false
}

// SetMaxRunningInstances gets a reference to the given int32 and assigns it to the MaxRunningInstances field.
func (o *ApplicationRequest) SetMaxRunningInstances(v int32) {
	o.MaxRunningInstances = &v
}

// GetHealthcheck returns the Healthcheck field value if set, zero value otherwise.
func (o *ApplicationRequest) GetHealthcheck() Healthcheck {
	if o == nil || o.Healthcheck == nil {
		var ret Healthcheck
		return ret
	}
	return *o.Healthcheck
}

// GetHealthcheckOk returns a tuple with the Healthcheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationRequest) GetHealthcheckOk() (*Healthcheck, bool) {
	if o == nil || o.Healthcheck == nil {
		return nil, false
	}
	return o.Healthcheck, true
}

// HasHealthcheck returns a boolean if a field has been set.
func (o *ApplicationRequest) HasHealthcheck() bool {
	if o != nil && o.Healthcheck != nil {
		return true
	}

	return false
}

// SetHealthcheck gets a reference to the given Healthcheck and assigns it to the Healthcheck field.
func (o *ApplicationRequest) SetHealthcheck(v Healthcheck) {
	o.Healthcheck = &v
}

// GetStorage returns the Storage field value if set, zero value otherwise.
func (o *ApplicationRequest) GetStorage() []ApplicationStorageRequestStorage {
	if o == nil || o.Storage == nil {
		var ret []ApplicationStorageRequestStorage
		return ret
	}
	return *o.Storage
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationRequest) GetStorageOk() (*[]ApplicationStorageRequestStorage, bool) {
	if o == nil || o.Storage == nil {
		return nil, false
	}
	return o.Storage, true
}

// HasStorage returns a boolean if a field has been set.
func (o *ApplicationRequest) HasStorage() bool {
	if o != nil && o.Storage != nil {
		return true
	}

	return false
}

// SetStorage gets a reference to the given []ApplicationStorageRequestStorage and assigns it to the Storage field.
func (o *ApplicationRequest) SetStorage(v []ApplicationStorageRequestStorage) {
	o.Storage = &v
}

// GetPorts returns the Ports field value if set, zero value otherwise.
func (o *ApplicationRequest) GetPorts() []ApplicationPortRequestPorts {
	if o == nil || o.Ports == nil {
		var ret []ApplicationPortRequestPorts
		return ret
	}
	return *o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationRequest) GetPortsOk() (*[]ApplicationPortRequestPorts, bool) {
	if o == nil || o.Ports == nil {
		return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *ApplicationRequest) HasPorts() bool {
	if o != nil && o.Ports != nil {
		return true
	}

	return false
}

// SetPorts gets a reference to the given []ApplicationPortRequestPorts and assigns it to the Ports field.
func (o *ApplicationRequest) SetPorts(v []ApplicationPortRequestPorts) {
	o.Ports = &v
}

func (o ApplicationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["git_repository"] = o.GitRepository
	}
	if o.BuildMode != nil {
		toSerialize["build_mode"] = o.BuildMode
	}
	if o.DockerfilePath != nil {
		toSerialize["dockerfile_path"] = o.DockerfilePath
	}
	if o.Cpu != nil {
		toSerialize["cpu"] = o.Cpu
	}
	if o.Memory != nil {
		toSerialize["memory"] = o.Memory
	}
	if o.MinRunningInstances != nil {
		toSerialize["min_running_instances"] = o.MinRunningInstances
	}
	if o.MaxRunningInstances != nil {
		toSerialize["max_running_instances"] = o.MaxRunningInstances
	}
	if o.Healthcheck != nil {
		toSerialize["healthcheck"] = o.Healthcheck
	}
	if o.Storage != nil {
		toSerialize["storage"] = o.Storage
	}
	if o.Ports != nil {
		toSerialize["ports"] = o.Ports
	}
	return json.Marshal(toSerialize)
}

type NullableApplicationRequest struct {
	value *ApplicationRequest
	isSet bool
}

func (v NullableApplicationRequest) Get() *ApplicationRequest {
	return v.value
}

func (v *NullableApplicationRequest) Set(val *ApplicationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationRequest(val *ApplicationRequest) *NullableApplicationRequest {
	return &NullableApplicationRequest{value: val, isSet: true}
}

func (v NullableApplicationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
