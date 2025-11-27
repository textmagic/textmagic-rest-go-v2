# CallPriceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Outbound** | **float32** | Price for outbound message. | 
**Inbound** | **float32** | Price for inbound message. | 
**Forward** | **float32** | Price for forward. | 
**Country** | **string** | The 2-letter ISO country code for local phone numbers, used when local is  set to true. Default is account country. | 

## Methods

### NewCallPriceResponse

`func NewCallPriceResponse(outbound float32, inbound float32, forward float32, country string, ) *CallPriceResponse`

NewCallPriceResponse instantiates a new CallPriceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallPriceResponseWithDefaults

`func NewCallPriceResponseWithDefaults() *CallPriceResponse`

NewCallPriceResponseWithDefaults instantiates a new CallPriceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOutbound

`func (o *CallPriceResponse) GetOutbound() float32`

GetOutbound returns the Outbound field if non-nil, zero value otherwise.

### GetOutboundOk

`func (o *CallPriceResponse) GetOutboundOk() (*float32, bool)`

GetOutboundOk returns a tuple with the Outbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutbound

`func (o *CallPriceResponse) SetOutbound(v float32)`

SetOutbound sets Outbound field to given value.


### GetInbound

`func (o *CallPriceResponse) GetInbound() float32`

GetInbound returns the Inbound field if non-nil, zero value otherwise.

### GetInboundOk

`func (o *CallPriceResponse) GetInboundOk() (*float32, bool)`

GetInboundOk returns a tuple with the Inbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbound

`func (o *CallPriceResponse) SetInbound(v float32)`

SetInbound sets Inbound field to given value.


### GetForward

`func (o *CallPriceResponse) GetForward() float32`

GetForward returns the Forward field if non-nil, zero value otherwise.

### GetForwardOk

`func (o *CallPriceResponse) GetForwardOk() (*float32, bool)`

GetForwardOk returns a tuple with the Forward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForward

`func (o *CallPriceResponse) SetForward(v float32)`

SetForward sets Forward field to given value.


### GetCountry

`func (o *CallPriceResponse) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CallPriceResponse) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CallPriceResponse) SetCountry(v string)`

SetCountry sets Country field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


