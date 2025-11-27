# GetAvailableDedicatedNumbersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Numbers** | **[]string** | Array of phone numbers. | 
**Price** | **float32** | Dedicated number monthly fee for this country. Returned in the current [account](https://docs.textmagic.com/#tag/User) currency. | 
**GiftType** | **string** |  | 

## Methods

### NewGetAvailableDedicatedNumbersResponse

`func NewGetAvailableDedicatedNumbersResponse(numbers []string, price float32, giftType string, ) *GetAvailableDedicatedNumbersResponse`

NewGetAvailableDedicatedNumbersResponse instantiates a new GetAvailableDedicatedNumbersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAvailableDedicatedNumbersResponseWithDefaults

`func NewGetAvailableDedicatedNumbersResponseWithDefaults() *GetAvailableDedicatedNumbersResponse`

NewGetAvailableDedicatedNumbersResponseWithDefaults instantiates a new GetAvailableDedicatedNumbersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumbers

`func (o *GetAvailableDedicatedNumbersResponse) GetNumbers() []string`

GetNumbers returns the Numbers field if non-nil, zero value otherwise.

### GetNumbersOk

`func (o *GetAvailableDedicatedNumbersResponse) GetNumbersOk() (*[]string, bool)`

GetNumbersOk returns a tuple with the Numbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumbers

`func (o *GetAvailableDedicatedNumbersResponse) SetNumbers(v []string)`

SetNumbers sets Numbers field to given value.


### GetPrice

`func (o *GetAvailableDedicatedNumbersResponse) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *GetAvailableDedicatedNumbersResponse) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *GetAvailableDedicatedNumbersResponse) SetPrice(v float32)`

SetPrice sets Price field to given value.


### GetGiftType

`func (o *GetAvailableDedicatedNumbersResponse) GetGiftType() string`

GetGiftType returns the GiftType field if non-nil, zero value otherwise.

### GetGiftTypeOk

`func (o *GetAvailableDedicatedNumbersResponse) GetGiftTypeOk() (*string, bool)`

GetGiftTypeOk returns a tuple with the GiftType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftType

`func (o *GetAvailableDedicatedNumbersResponse) SetGiftType(v string)`

SetGiftType sets GiftType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


