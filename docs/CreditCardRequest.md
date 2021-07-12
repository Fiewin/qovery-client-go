# CreditCardRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | **string** |  | 
**Cvv** | **string** |  | 
**ExpiryMonth** | **int32** |  | 
**ExpiryYear** | **int32** |  | 

## Methods

### NewCreditCardRequest

`func NewCreditCardRequest(number string, cvv string, expiryMonth int32, expiryYear int32, ) *CreditCardRequest`

NewCreditCardRequest instantiates a new CreditCardRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditCardRequestWithDefaults

`func NewCreditCardRequestWithDefaults() *CreditCardRequest`

NewCreditCardRequestWithDefaults instantiates a new CreditCardRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *CreditCardRequest) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *CreditCardRequest) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *CreditCardRequest) SetNumber(v string)`

SetNumber sets Number field to given value.


### GetCvv

`func (o *CreditCardRequest) GetCvv() string`

GetCvv returns the Cvv field if non-nil, zero value otherwise.

### GetCvvOk

`func (o *CreditCardRequest) GetCvvOk() (*string, bool)`

GetCvvOk returns a tuple with the Cvv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvv

`func (o *CreditCardRequest) SetCvv(v string)`

SetCvv sets Cvv field to given value.


### GetExpiryMonth

`func (o *CreditCardRequest) GetExpiryMonth() int32`

GetExpiryMonth returns the ExpiryMonth field if non-nil, zero value otherwise.

### GetExpiryMonthOk

`func (o *CreditCardRequest) GetExpiryMonthOk() (*int32, bool)`

GetExpiryMonthOk returns a tuple with the ExpiryMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryMonth

`func (o *CreditCardRequest) SetExpiryMonth(v int32)`

SetExpiryMonth sets ExpiryMonth field to given value.


### GetExpiryYear

`func (o *CreditCardRequest) GetExpiryYear() int32`

GetExpiryYear returns the ExpiryYear field if non-nil, zero value otherwise.

### GetExpiryYearOk

`func (o *CreditCardRequest) GetExpiryYearOk() (*int32, bool)`

GetExpiryYearOk returns a tuple with the ExpiryYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryYear

`func (o *CreditCardRequest) SetExpiryYear(v int32)`

SetExpiryYear sets ExpiryYear field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


