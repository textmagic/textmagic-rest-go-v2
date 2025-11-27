# GetMessagePriceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **float32** | Total price of the message. | 
**Parts** | **NullableInt32** | Message parts (multiples of 160 characters) count. | 
**Countries** | [**[]GetMessagePriceResponseCountriesItem**](GetMessagePriceResponseCountriesItem.md) |  | 

## Methods

### NewGetMessagePriceResponse

`func NewGetMessagePriceResponse(total float32, parts NullableInt32, countries []GetMessagePriceResponseCountriesItem, ) *GetMessagePriceResponse`

NewGetMessagePriceResponse instantiates a new GetMessagePriceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMessagePriceResponseWithDefaults

`func NewGetMessagePriceResponseWithDefaults() *GetMessagePriceResponse`

NewGetMessagePriceResponseWithDefaults instantiates a new GetMessagePriceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *GetMessagePriceResponse) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetMessagePriceResponse) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetMessagePriceResponse) SetTotal(v float32)`

SetTotal sets Total field to given value.


### GetParts

`func (o *GetMessagePriceResponse) GetParts() int32`

GetParts returns the Parts field if non-nil, zero value otherwise.

### GetPartsOk

`func (o *GetMessagePriceResponse) GetPartsOk() (*int32, bool)`

GetPartsOk returns a tuple with the Parts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParts

`func (o *GetMessagePriceResponse) SetParts(v int32)`

SetParts sets Parts field to given value.


### SetPartsNil

`func (o *GetMessagePriceResponse) SetPartsNil(b bool)`

 SetPartsNil sets the value for Parts to be an explicit nil

### UnsetParts
`func (o *GetMessagePriceResponse) UnsetParts()`

UnsetParts ensures that no value is present for Parts, not even an explicit nil
### GetCountries

`func (o *GetMessagePriceResponse) GetCountries() []GetMessagePriceResponseCountriesItem`

GetCountries returns the Countries field if non-nil, zero value otherwise.

### GetCountriesOk

`func (o *GetMessagePriceResponse) GetCountriesOk() (*[]GetMessagePriceResponseCountriesItem, bool)`

GetCountriesOk returns a tuple with the Countries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountries

`func (o *GetMessagePriceResponse) SetCountries(v []GetMessagePriceResponseCountriesItem)`

SetCountries sets Countries field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


