# ClusterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | name is case insensitive | 
**Description** | Pointer to **string** |  | [optional] 
**AutoUpdate** | Pointer to **bool** |  | [optional] 
**EnableHistoricMetric** | Pointer to **bool** |  | [optional] 
**EnableHistoricLog** | Pointer to **bool** |  | [optional] 
**Cpu** | Pointer to **float32** | unit is millicores (m). 1000m &#x3D; 1 cpu | [optional] [default to 250]
**Memory** | Pointer to **float32** | unit is MB. 1024 MB &#x3D; 1GB | [optional] [default to 256]
**MinRunningNodes** | Pointer to **int32** |  | [optional] [default to 1]
**MaxRunningNodes** | Pointer to **int32** |  | [optional] [default to 1]

## Methods

### NewClusterRequest

`func NewClusterRequest(name string, ) *ClusterRequest`

NewClusterRequest instantiates a new ClusterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterRequestWithDefaults

`func NewClusterRequestWithDefaults() *ClusterRequest`

NewClusterRequestWithDefaults instantiates a new ClusterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ClusterRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ClusterRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ClusterRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ClusterRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ClusterRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAutoUpdate

`func (o *ClusterRequest) GetAutoUpdate() bool`

GetAutoUpdate returns the AutoUpdate field if non-nil, zero value otherwise.

### GetAutoUpdateOk

`func (o *ClusterRequest) GetAutoUpdateOk() (*bool, bool)`

GetAutoUpdateOk returns a tuple with the AutoUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoUpdate

`func (o *ClusterRequest) SetAutoUpdate(v bool)`

SetAutoUpdate sets AutoUpdate field to given value.

### HasAutoUpdate

`func (o *ClusterRequest) HasAutoUpdate() bool`

HasAutoUpdate returns a boolean if a field has been set.

### GetEnableHistoricMetric

`func (o *ClusterRequest) GetEnableHistoricMetric() bool`

GetEnableHistoricMetric returns the EnableHistoricMetric field if non-nil, zero value otherwise.

### GetEnableHistoricMetricOk

`func (o *ClusterRequest) GetEnableHistoricMetricOk() (*bool, bool)`

GetEnableHistoricMetricOk returns a tuple with the EnableHistoricMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableHistoricMetric

`func (o *ClusterRequest) SetEnableHistoricMetric(v bool)`

SetEnableHistoricMetric sets EnableHistoricMetric field to given value.

### HasEnableHistoricMetric

`func (o *ClusterRequest) HasEnableHistoricMetric() bool`

HasEnableHistoricMetric returns a boolean if a field has been set.

### GetEnableHistoricLog

`func (o *ClusterRequest) GetEnableHistoricLog() bool`

GetEnableHistoricLog returns the EnableHistoricLog field if non-nil, zero value otherwise.

### GetEnableHistoricLogOk

`func (o *ClusterRequest) GetEnableHistoricLogOk() (*bool, bool)`

GetEnableHistoricLogOk returns a tuple with the EnableHistoricLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableHistoricLog

`func (o *ClusterRequest) SetEnableHistoricLog(v bool)`

SetEnableHistoricLog sets EnableHistoricLog field to given value.

### HasEnableHistoricLog

`func (o *ClusterRequest) HasEnableHistoricLog() bool`

HasEnableHistoricLog returns a boolean if a field has been set.

### GetCpu

`func (o *ClusterRequest) GetCpu() float32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *ClusterRequest) GetCpuOk() (*float32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *ClusterRequest) SetCpu(v float32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *ClusterRequest) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetMemory

`func (o *ClusterRequest) GetMemory() float32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *ClusterRequest) GetMemoryOk() (*float32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *ClusterRequest) SetMemory(v float32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *ClusterRequest) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetMinRunningNodes

`func (o *ClusterRequest) GetMinRunningNodes() int32`

GetMinRunningNodes returns the MinRunningNodes field if non-nil, zero value otherwise.

### GetMinRunningNodesOk

`func (o *ClusterRequest) GetMinRunningNodesOk() (*int32, bool)`

GetMinRunningNodesOk returns a tuple with the MinRunningNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRunningNodes

`func (o *ClusterRequest) SetMinRunningNodes(v int32)`

SetMinRunningNodes sets MinRunningNodes field to given value.

### HasMinRunningNodes

`func (o *ClusterRequest) HasMinRunningNodes() bool`

HasMinRunningNodes returns a boolean if a field has been set.

### GetMaxRunningNodes

`func (o *ClusterRequest) GetMaxRunningNodes() int32`

GetMaxRunningNodes returns the MaxRunningNodes field if non-nil, zero value otherwise.

### GetMaxRunningNodesOk

`func (o *ClusterRequest) GetMaxRunningNodesOk() (*int32, bool)`

GetMaxRunningNodesOk returns a tuple with the MaxRunningNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRunningNodes

`func (o *ClusterRequest) SetMaxRunningNodes(v int32)`

SetMaxRunningNodes sets MaxRunningNodes field to given value.

### HasMaxRunningNodes

`func (o *ClusterRequest) HasMaxRunningNodes() bool`

HasMaxRunningNodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


