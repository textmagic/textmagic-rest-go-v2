# GetMessagePriceResponseCountriesItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Country** | **string** | The 2-letter ISO country code. | 
**CountryName** | **string** | Country name. | 
**AllowDedicated** | **bool** | Is allowed to use a dedicated number? | 
**Count** | **float32** | Parts count to send. | 
**Max** | **float32** | Maximum parts to send. | 
**Sum** | **string** | Total price to send. | 
**Landline** | **float32** | Is this a landline number? | 

## Methods

### NewGetMessagePriceResponseCountriesItem

`func NewGetMessagePriceResponseCountriesItem(country string, countryName string, allowDedicated bool, count float32, max float32, sum string, landline float32, ) *GetMessagePriceResponseCountriesItem`

NewGetMessagePriceResponseCountriesItem instantiates a new GetMessagePriceResponseCountriesItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMessagePriceResponseCountriesItemWithDefaults

`func NewGetMessagePriceResponseCountriesItemWithDefaults() *GetMessagePriceResponseCountriesItem`

NewGetMessagePriceResponseCountriesItemWithDefaults instantiates a new GetMessagePriceResponseCountriesItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountry

`func (o *GetMessagePriceResponseCountriesItem) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *GetMessagePriceResponseCountriesItem) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *GetMessagePriceResponseCountriesItem) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetCountryName

`func (o *GetMessagePriceResponseCountriesItem) GetCountryName() string`

GetCountryName returns the CountryName field if non-nil, zero value otherwise.

### GetCountryNameOk

`func (o *GetMessagePriceResponseCountriesItem) GetCountryNameOk() (*string, bool)`

GetCountryNameOk returns a tuple with the CountryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryName

`func (o *GetMessagePriceResponseCountriesItem) SetCountryName(v string)`

SetCountryName sets CountryName field to given value.


### GetAllowDedicated

`func (o *GetMessagePriceResponseCountriesItem) GetAllowDedicated() bool`

GetAllowDedicated returns the AllowDedicated field if non-nil, zero value otherwise.

### GetAllowDedicatedOk

`func (o *GetMessagePriceResponseCountriesItem) GetAllowDedicatedOk() (*bool, bool)`

GetAllowDedicatedOk returns a tuple with the AllowDedicated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowDedicated

`func (o *GetMessagePriceResponseCountriesItem) SetAllowDedicated(v bool)`

SetAllowDedicated sets AllowDedicated field to given value.


### GetCount

`func (o *GetMessagePriceResponseCountriesItem) GetCount() float32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GetMessagePriceResponseCountriesItem) GetCountOk() (*float32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GetMessagePriceResponseCountriesItem) SetCount(v float32)`

SetCount sets Count field to given value.


### GetMax

`func (o *GetMessagePriceResponseCountriesItem) GetMax() float32`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *GetMessagePriceResponseCountriesItem) GetMaxOk() (*float32, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *GetMessagePriceResponseCountriesItem) SetMax(v float32)`

SetMax sets Max field to given value.


### GetSum

`func (o *GetMessagePriceResponseCountriesItem) GetSum() string`

GetSum returns the Sum field if non-nil, zero value otherwise.

### GetSumOk

`func (o *GetMessagePriceResponseCountriesItem) GetSumOk() (*string, bool)`

GetSumOk returns a tuple with the Sum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSum

`func (o *GetMessagePriceResponseCountriesItem) SetSum(v string)`

SetSum sets Sum field to given value.


### GetLandline

`func (o *GetMessagePriceResponseCountriesItem) GetLandline() float32`

GetLandline returns the Landline field if non-nil, zero value otherwise.

### GetLandlineOk

`func (o *GetMessagePriceResponseCountriesItem) GetLandlineOk() (*float32, bool)`

GetLandlineOk returns a tuple with the Landline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLandline

`func (o *GetMessagePriceResponseCountriesItem) SetLandline(v float32)`

SetLandline sets Landline field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


