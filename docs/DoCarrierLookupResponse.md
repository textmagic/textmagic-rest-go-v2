# DoCarrierLookupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cost** | **float32** | The cost to check that one number is constant – 0.04 in your account currency. | 
**Country** | Pointer to [**NullableCountry**](Country.md) |  | [optional] 
**Local** | **string** | Phone number in [National format](https://en.wikipedia.org/wiki/National_conventions_for_writing_telephone_numbers). | 
**Type** | **string** | Phone number type. | 
**Carrier** | **string** | Carrier name. | 
**Number164** | **string** | Phone number in [E.164 format](https://en.wikipedia.org/wiki/E.164). | 
**Valid** | **bool** | This field shows whether the entered phone number is valid or not. | 

## Methods

### NewDoCarrierLookupResponse

`func NewDoCarrierLookupResponse(cost float32, local string, type_ string, carrier string, number164 string, valid bool, ) *DoCarrierLookupResponse`

NewDoCarrierLookupResponse instantiates a new DoCarrierLookupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDoCarrierLookupResponseWithDefaults

`func NewDoCarrierLookupResponseWithDefaults() *DoCarrierLookupResponse`

NewDoCarrierLookupResponseWithDefaults instantiates a new DoCarrierLookupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCost

`func (o *DoCarrierLookupResponse) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *DoCarrierLookupResponse) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *DoCarrierLookupResponse) SetCost(v float32)`

SetCost sets Cost field to given value.


### GetCountry

`func (o *DoCarrierLookupResponse) GetCountry() Country`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *DoCarrierLookupResponse) GetCountryOk() (*Country, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *DoCarrierLookupResponse) SetCountry(v Country)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *DoCarrierLookupResponse) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *DoCarrierLookupResponse) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *DoCarrierLookupResponse) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetLocal

`func (o *DoCarrierLookupResponse) GetLocal() string`

GetLocal returns the Local field if non-nil, zero value otherwise.

### GetLocalOk

`func (o *DoCarrierLookupResponse) GetLocalOk() (*string, bool)`

GetLocalOk returns a tuple with the Local field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocal

`func (o *DoCarrierLookupResponse) SetLocal(v string)`

SetLocal sets Local field to given value.


### GetType

`func (o *DoCarrierLookupResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DoCarrierLookupResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DoCarrierLookupResponse) SetType(v string)`

SetType sets Type field to given value.


### GetCarrier

`func (o *DoCarrierLookupResponse) GetCarrier() string`

GetCarrier returns the Carrier field if non-nil, zero value otherwise.

### GetCarrierOk

`func (o *DoCarrierLookupResponse) GetCarrierOk() (*string, bool)`

GetCarrierOk returns a tuple with the Carrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrier

`func (o *DoCarrierLookupResponse) SetCarrier(v string)`

SetCarrier sets Carrier field to given value.


### GetNumber164

`func (o *DoCarrierLookupResponse) GetNumber164() string`

GetNumber164 returns the Number164 field if non-nil, zero value otherwise.

### GetNumber164Ok

`func (o *DoCarrierLookupResponse) GetNumber164Ok() (*string, bool)`

GetNumber164Ok returns a tuple with the Number164 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber164

`func (o *DoCarrierLookupResponse) SetNumber164(v string)`

SetNumber164 sets Number164 field to given value.


### GetValid

`func (o *DoCarrierLookupResponse) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *DoCarrierLookupResponse) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *DoCarrierLookupResponse) SetValid(v bool)`

SetValid sets Valid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


