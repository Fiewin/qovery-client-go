# ApplicationStorageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Storage** | Pointer to [**[]ApplicationStorageRequestStorage**](ApplicationStorageRequestStorage.md) |  | [optional] 

## Methods

### NewApplicationStorageRequest

`func NewApplicationStorageRequest() *ApplicationStorageRequest`

NewApplicationStorageRequest instantiates a new ApplicationStorageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationStorageRequestWithDefaults

`func NewApplicationStorageRequestWithDefaults() *ApplicationStorageRequest`

NewApplicationStorageRequestWithDefaults instantiates a new ApplicationStorageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStorage

`func (o *ApplicationStorageRequest) GetStorage() []ApplicationStorageRequestStorage`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *ApplicationStorageRequest) GetStorageOk() (*[]ApplicationStorageRequestStorage, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *ApplicationStorageRequest) SetStorage(v []ApplicationStorageRequestStorage)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *ApplicationStorageRequest) HasStorage() bool`

HasStorage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


