# MessagePriceItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Country name. | 
**Price** | **string** | Price to send message to desired country. | 
**Country** | **string** | The 2-letter ISO country code of the recipient&#39;s phone number. | 

## Methods

### NewMessagePriceItem

`func NewMessagePriceItem(name string, price string, country string, ) *MessagePriceItem`

NewMessagePriceItem instantiates a new MessagePriceItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessagePriceItemWithDefaults

`func NewMessagePriceItemWithDefaults() *MessagePriceItem`

NewMessagePriceItemWithDefaults instantiates a new MessagePriceItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MessagePriceItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MessagePriceItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MessagePriceItem) SetName(v string)`

SetName sets Name field to given value.


### GetPrice

`func (o *MessagePriceItem) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *MessagePriceItem) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *MessagePriceItem) SetPrice(v string)`

SetPrice sets Price field to given value.


### GetCountry

`func (o *MessagePriceItem) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *MessagePriceItem) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *MessagePriceItem) SetCountry(v string)`

SetCountry sets Country field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


