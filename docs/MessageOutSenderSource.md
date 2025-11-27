# MessageOutSenderSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** |  | [optional] 
**Phone** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**CountryId** | Pointer to **NullableString** |  | [optional] 
**Label** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**CarrierStatus** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewMessageOutSenderSource

`func NewMessageOutSenderSource() *MessageOutSenderSource`

NewMessageOutSenderSource instantiates a new MessageOutSenderSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageOutSenderSourceWithDefaults

`func NewMessageOutSenderSourceWithDefaults() *MessageOutSenderSource`

NewMessageOutSenderSourceWithDefaults instantiates a new MessageOutSenderSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MessageOutSenderSource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MessageOutSenderSource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MessageOutSenderSource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MessageOutSenderSource) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *MessageOutSenderSource) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *MessageOutSenderSource) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetPhone

`func (o *MessageOutSenderSource) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *MessageOutSenderSource) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *MessageOutSenderSource) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *MessageOutSenderSource) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *MessageOutSenderSource) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *MessageOutSenderSource) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetType

`func (o *MessageOutSenderSource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MessageOutSenderSource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MessageOutSenderSource) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MessageOutSenderSource) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *MessageOutSenderSource) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *MessageOutSenderSource) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetCountryId

`func (o *MessageOutSenderSource) GetCountryId() string`

GetCountryId returns the CountryId field if non-nil, zero value otherwise.

### GetCountryIdOk

`func (o *MessageOutSenderSource) GetCountryIdOk() (*string, bool)`

GetCountryIdOk returns a tuple with the CountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryId

`func (o *MessageOutSenderSource) SetCountryId(v string)`

SetCountryId sets CountryId field to given value.

### HasCountryId

`func (o *MessageOutSenderSource) HasCountryId() bool`

HasCountryId returns a boolean if a field has been set.

### SetCountryIdNil

`func (o *MessageOutSenderSource) SetCountryIdNil(b bool)`

 SetCountryIdNil sets the value for CountryId to be an explicit nil

### UnsetCountryId
`func (o *MessageOutSenderSource) UnsetCountryId()`

UnsetCountryId ensures that no value is present for CountryId, not even an explicit nil
### GetLabel

`func (o *MessageOutSenderSource) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *MessageOutSenderSource) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *MessageOutSenderSource) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *MessageOutSenderSource) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *MessageOutSenderSource) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *MessageOutSenderSource) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetStatus

`func (o *MessageOutSenderSource) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MessageOutSenderSource) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MessageOutSenderSource) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MessageOutSenderSource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *MessageOutSenderSource) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *MessageOutSenderSource) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetCarrierStatus

`func (o *MessageOutSenderSource) GetCarrierStatus() string`

GetCarrierStatus returns the CarrierStatus field if non-nil, zero value otherwise.

### GetCarrierStatusOk

`func (o *MessageOutSenderSource) GetCarrierStatusOk() (*string, bool)`

GetCarrierStatusOk returns a tuple with the CarrierStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierStatus

`func (o *MessageOutSenderSource) SetCarrierStatus(v string)`

SetCarrierStatus sets CarrierStatus field to given value.

### HasCarrierStatus

`func (o *MessageOutSenderSource) HasCarrierStatus() bool`

HasCarrierStatus returns a boolean if a field has been set.

### SetCarrierStatusNil

`func (o *MessageOutSenderSource) SetCarrierStatusNil(b bool)`

 SetCarrierStatusNil sets the value for CarrierStatus to be an explicit nil

### UnsetCarrierStatus
`func (o *MessageOutSenderSource) UnsetCarrierStatus()`

UnsetCarrierStatus ensures that no value is present for CarrierStatus, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


