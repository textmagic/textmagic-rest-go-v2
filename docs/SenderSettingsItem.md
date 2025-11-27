# SenderSettingsItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Country** | **string** | The 2-letter ISO country code of the recipient&#39;s phone number.  | 
**Phone** | **string** | Phone enabled for sending to a specified country. | 

## Methods

### NewSenderSettingsItem

`func NewSenderSettingsItem(country string, phone string, ) *SenderSettingsItem`

NewSenderSettingsItem instantiates a new SenderSettingsItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSenderSettingsItemWithDefaults

`func NewSenderSettingsItemWithDefaults() *SenderSettingsItem`

NewSenderSettingsItemWithDefaults instantiates a new SenderSettingsItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountry

`func (o *SenderSettingsItem) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *SenderSettingsItem) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *SenderSettingsItem) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetPhone

`func (o *SenderSettingsItem) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *SenderSettingsItem) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *SenderSettingsItem) SetPhone(v string)`

SetPhone sets Phone field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


