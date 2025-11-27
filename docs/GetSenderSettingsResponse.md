# GetSenderSettingsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | [**[]SenderSettingsItem**](SenderSettingsItem.md) |  | 
**Special** | [**[]SenderSettingsItem**](SenderSettingsItem.md) |  | 
**Other** | [**[]SenderSettingsItem**](SenderSettingsItem.md) |  | 

## Methods

### NewGetSenderSettingsResponse

`func NewGetSenderSettingsResponse(user []SenderSettingsItem, special []SenderSettingsItem, other []SenderSettingsItem, ) *GetSenderSettingsResponse`

NewGetSenderSettingsResponse instantiates a new GetSenderSettingsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSenderSettingsResponseWithDefaults

`func NewGetSenderSettingsResponseWithDefaults() *GetSenderSettingsResponse`

NewGetSenderSettingsResponseWithDefaults instantiates a new GetSenderSettingsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *GetSenderSettingsResponse) GetUser() []SenderSettingsItem`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *GetSenderSettingsResponse) GetUserOk() (*[]SenderSettingsItem, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *GetSenderSettingsResponse) SetUser(v []SenderSettingsItem)`

SetUser sets User field to given value.


### GetSpecial

`func (o *GetSenderSettingsResponse) GetSpecial() []SenderSettingsItem`

GetSpecial returns the Special field if non-nil, zero value otherwise.

### GetSpecialOk

`func (o *GetSenderSettingsResponse) GetSpecialOk() (*[]SenderSettingsItem, bool)`

GetSpecialOk returns a tuple with the Special field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecial

`func (o *GetSenderSettingsResponse) SetSpecial(v []SenderSettingsItem)`

SetSpecial sets Special field to given value.


### GetOther

`func (o *GetSenderSettingsResponse) GetOther() []SenderSettingsItem`

GetOther returns the Other field if non-nil, zero value otherwise.

### GetOtherOk

`func (o *GetSenderSettingsResponse) GetOtherOk() (*[]SenderSettingsItem, bool)`

GetOtherOk returns a tuple with the Other field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOther

`func (o *GetSenderSettingsResponse) SetOther(v []SenderSettingsItem)`

SetOther sets Other field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


