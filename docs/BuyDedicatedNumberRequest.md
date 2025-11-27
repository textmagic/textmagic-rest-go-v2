# BuyDedicatedNumberRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phone** | Pointer to **string** | Dedicated phone number. | [optional] 
**Country** | Pointer to **string** | Country code phone number. | [optional] 
**UserId** | Pointer to **int32** | Assigned dedicated number. This number will be available for this account only. You cannot transfer numbers between sub-accounts.  | [optional] 

## Methods

### NewBuyDedicatedNumberRequest

`func NewBuyDedicatedNumberRequest() *BuyDedicatedNumberRequest`

NewBuyDedicatedNumberRequest instantiates a new BuyDedicatedNumberRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuyDedicatedNumberRequestWithDefaults

`func NewBuyDedicatedNumberRequestWithDefaults() *BuyDedicatedNumberRequest`

NewBuyDedicatedNumberRequestWithDefaults instantiates a new BuyDedicatedNumberRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhone

`func (o *BuyDedicatedNumberRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *BuyDedicatedNumberRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *BuyDedicatedNumberRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *BuyDedicatedNumberRequest) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetCountry

`func (o *BuyDedicatedNumberRequest) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *BuyDedicatedNumberRequest) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *BuyDedicatedNumberRequest) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *BuyDedicatedNumberRequest) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetUserId

`func (o *BuyDedicatedNumberRequest) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *BuyDedicatedNumberRequest) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *BuyDedicatedNumberRequest) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *BuyDedicatedNumberRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


