# GetCallsPricesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Outbound** | **float32** | Price for outbound message. | 
**Inbound** | **float32** | Price for inbound message. | 
**Forward** | **float32** | Price for forward. | 
**Country** | **string** | 2-letter ISO country code for local phone numbers, used when local is  set to true. Default is account country. | 

## Methods

### NewGetCallsPricesResponse

`func NewGetCallsPricesResponse(outbound float32, inbound float32, forward float32, country string, ) *GetCallsPricesResponse`

NewGetCallsPricesResponse instantiates a new GetCallsPricesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCallsPricesResponseWithDefaults

`func NewGetCallsPricesResponseWithDefaults() *GetCallsPricesResponse`

NewGetCallsPricesResponseWithDefaults instantiates a new GetCallsPricesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOutbound

`func (o *GetCallsPricesResponse) GetOutbound() float32`

GetOutbound returns the Outbound field if non-nil, zero value otherwise.

### GetOutboundOk

`func (o *GetCallsPricesResponse) GetOutboundOk() (*float32, bool)`

GetOutboundOk returns a tuple with the Outbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutbound

`func (o *GetCallsPricesResponse) SetOutbound(v float32)`

SetOutbound sets Outbound field to given value.


### GetInbound

`func (o *GetCallsPricesResponse) GetInbound() float32`

GetInbound returns the Inbound field if non-nil, zero value otherwise.

### GetInboundOk

`func (o *GetCallsPricesResponse) GetInboundOk() (*float32, bool)`

GetInboundOk returns a tuple with the Inbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbound

`func (o *GetCallsPricesResponse) SetInbound(v float32)`

SetInbound sets Inbound field to given value.


### GetForward

`func (o *GetCallsPricesResponse) GetForward() float32`

GetForward returns the Forward field if non-nil, zero value otherwise.

### GetForwardOk

`func (o *GetCallsPricesResponse) GetForwardOk() (*float32, bool)`

GetForwardOk returns a tuple with the Forward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForward

`func (o *GetCallsPricesResponse) SetForward(v float32)`

SetForward sets Forward field to given value.


### GetCountry

`func (o *GetCallsPricesResponse) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *GetCallsPricesResponse) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *GetCallsPricesResponse) SetCountry(v string)`

SetCountry sets Country field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


