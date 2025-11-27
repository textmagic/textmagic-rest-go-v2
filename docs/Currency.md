# Currency

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The 3-letter ISO currency ID. | 
**UnicodeSymbol** | **string** | Unicode-compatible currency symbol. | 
**HtmlSymbol** | **string** | HTML-compatible currency symbol. | 

## Methods

### NewCurrency

`func NewCurrency(id string, unicodeSymbol string, htmlSymbol string, ) *Currency`

NewCurrency instantiates a new Currency object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrencyWithDefaults

`func NewCurrencyWithDefaults() *Currency`

NewCurrencyWithDefaults instantiates a new Currency object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Currency) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Currency) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Currency) SetId(v string)`

SetId sets Id field to given value.


### GetUnicodeSymbol

`func (o *Currency) GetUnicodeSymbol() string`

GetUnicodeSymbol returns the UnicodeSymbol field if non-nil, zero value otherwise.

### GetUnicodeSymbolOk

`func (o *Currency) GetUnicodeSymbolOk() (*string, bool)`

GetUnicodeSymbolOk returns a tuple with the UnicodeSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnicodeSymbol

`func (o *Currency) SetUnicodeSymbol(v string)`

SetUnicodeSymbol sets UnicodeSymbol field to given value.


### GetHtmlSymbol

`func (o *Currency) GetHtmlSymbol() string`

GetHtmlSymbol returns the HtmlSymbol field if non-nil, zero value otherwise.

### GetHtmlSymbolOk

`func (o *Currency) GetHtmlSymbolOk() (*string, bool)`

GetHtmlSymbolOk returns a tuple with the HtmlSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlSymbol

`func (o *Currency) SetHtmlSymbol(v string)`

SetHtmlSymbol sets HtmlSymbol field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


