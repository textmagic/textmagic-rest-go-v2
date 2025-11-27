# UpdateSenderSettingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** | Available phone number in international E.164 format or senderid. | [optional] 
**Country** | Pointer to **string** | Country for which the setting will be set. | [optional] 
**ChatId** | Pointer to **int32** | Set sender setting for specified chat only. | [optional] 

## Methods

### NewUpdateSenderSettingRequest

`func NewUpdateSenderSettingRequest() *UpdateSenderSettingRequest`

NewUpdateSenderSettingRequest instantiates a new UpdateSenderSettingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSenderSettingRequestWithDefaults

`func NewUpdateSenderSettingRequestWithDefaults() *UpdateSenderSettingRequest`

NewUpdateSenderSettingRequestWithDefaults instantiates a new UpdateSenderSettingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *UpdateSenderSettingRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UpdateSenderSettingRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UpdateSenderSettingRequest) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *UpdateSenderSettingRequest) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetCountry

`func (o *UpdateSenderSettingRequest) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *UpdateSenderSettingRequest) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *UpdateSenderSettingRequest) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *UpdateSenderSettingRequest) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetChatId

`func (o *UpdateSenderSettingRequest) GetChatId() int32`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *UpdateSenderSettingRequest) GetChatIdOk() (*int32, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *UpdateSenderSettingRequest) SetChatId(v int32)`

SetChatId sets ChatId field to given value.

### HasChatId

`func (o *UpdateSenderSettingRequest) HasChatId() bool`

HasChatId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


