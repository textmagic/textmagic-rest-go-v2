# MessagesIcsTextParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cost** | **NullableFloat32** | Cost to check that one number is constant – 0.04 in your account currency. | 
**Parts** | **NullableInt32** | Message parts (multiples of 160 characters) count. | 
**Chars** | **NullableInt32** | Characters count. | 
**Encoding** | **NullableString** | Message charset. Could be: * **ISO-8859-1** – for plaintext SMS; * **UTF-16BE** – for Unicode SMS.  | 
**Countries** | **[]string** |  | 
**CharsetLabel** | **NullableString** | Human-readable message charset label. Could be: *   **ISO-8859-1** for plaintext SMS; *   **UTF-16BE** for Unicode SMS; *   **Voice** for voice services (Text-to-Speech or Voice Broadcast) messages.  | 

## Methods

### NewMessagesIcsTextParameters

`func NewMessagesIcsTextParameters(cost NullableFloat32, parts NullableInt32, chars NullableInt32, encoding NullableString, countries []string, charsetLabel NullableString, ) *MessagesIcsTextParameters`

NewMessagesIcsTextParameters instantiates a new MessagesIcsTextParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessagesIcsTextParametersWithDefaults

`func NewMessagesIcsTextParametersWithDefaults() *MessagesIcsTextParameters`

NewMessagesIcsTextParametersWithDefaults instantiates a new MessagesIcsTextParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCost

`func (o *MessagesIcsTextParameters) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *MessagesIcsTextParameters) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *MessagesIcsTextParameters) SetCost(v float32)`

SetCost sets Cost field to given value.


### SetCostNil

`func (o *MessagesIcsTextParameters) SetCostNil(b bool)`

 SetCostNil sets the value for Cost to be an explicit nil

### UnsetCost
`func (o *MessagesIcsTextParameters) UnsetCost()`

UnsetCost ensures that no value is present for Cost, not even an explicit nil
### GetParts

`func (o *MessagesIcsTextParameters) GetParts() int32`

GetParts returns the Parts field if non-nil, zero value otherwise.

### GetPartsOk

`func (o *MessagesIcsTextParameters) GetPartsOk() (*int32, bool)`

GetPartsOk returns a tuple with the Parts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParts

`func (o *MessagesIcsTextParameters) SetParts(v int32)`

SetParts sets Parts field to given value.


### SetPartsNil

`func (o *MessagesIcsTextParameters) SetPartsNil(b bool)`

 SetPartsNil sets the value for Parts to be an explicit nil

### UnsetParts
`func (o *MessagesIcsTextParameters) UnsetParts()`

UnsetParts ensures that no value is present for Parts, not even an explicit nil
### GetChars

`func (o *MessagesIcsTextParameters) GetChars() int32`

GetChars returns the Chars field if non-nil, zero value otherwise.

### GetCharsOk

`func (o *MessagesIcsTextParameters) GetCharsOk() (*int32, bool)`

GetCharsOk returns a tuple with the Chars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChars

`func (o *MessagesIcsTextParameters) SetChars(v int32)`

SetChars sets Chars field to given value.


### SetCharsNil

`func (o *MessagesIcsTextParameters) SetCharsNil(b bool)`

 SetCharsNil sets the value for Chars to be an explicit nil

### UnsetChars
`func (o *MessagesIcsTextParameters) UnsetChars()`

UnsetChars ensures that no value is present for Chars, not even an explicit nil
### GetEncoding

`func (o *MessagesIcsTextParameters) GetEncoding() string`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *MessagesIcsTextParameters) GetEncodingOk() (*string, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *MessagesIcsTextParameters) SetEncoding(v string)`

SetEncoding sets Encoding field to given value.


### SetEncodingNil

`func (o *MessagesIcsTextParameters) SetEncodingNil(b bool)`

 SetEncodingNil sets the value for Encoding to be an explicit nil

### UnsetEncoding
`func (o *MessagesIcsTextParameters) UnsetEncoding()`

UnsetEncoding ensures that no value is present for Encoding, not even an explicit nil
### GetCountries

`func (o *MessagesIcsTextParameters) GetCountries() []string`

GetCountries returns the Countries field if non-nil, zero value otherwise.

### GetCountriesOk

`func (o *MessagesIcsTextParameters) GetCountriesOk() (*[]string, bool)`

GetCountriesOk returns a tuple with the Countries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountries

`func (o *MessagesIcsTextParameters) SetCountries(v []string)`

SetCountries sets Countries field to given value.


### SetCountriesNil

`func (o *MessagesIcsTextParameters) SetCountriesNil(b bool)`

 SetCountriesNil sets the value for Countries to be an explicit nil

### UnsetCountries
`func (o *MessagesIcsTextParameters) UnsetCountries()`

UnsetCountries ensures that no value is present for Countries, not even an explicit nil
### GetCharsetLabel

`func (o *MessagesIcsTextParameters) GetCharsetLabel() string`

GetCharsetLabel returns the CharsetLabel field if non-nil, zero value otherwise.

### GetCharsetLabelOk

`func (o *MessagesIcsTextParameters) GetCharsetLabelOk() (*string, bool)`

GetCharsetLabelOk returns a tuple with the CharsetLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharsetLabel

`func (o *MessagesIcsTextParameters) SetCharsetLabel(v string)`

SetCharsetLabel sets CharsetLabel field to given value.


### SetCharsetLabelNil

`func (o *MessagesIcsTextParameters) SetCharsetLabelNil(b bool)`

 SetCharsetLabelNil sets the value for CharsetLabel to be an explicit nil

### UnsetCharsetLabel
`func (o *MessagesIcsTextParameters) UnsetCharsetLabel()`

UnsetCharsetLabel ensures that no value is present for CharsetLabel, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


