# UpdateBalanceNotificationSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LowBalanceNotification** | Pointer to **bool** | Should user receive low balance notification. | [optional] 
**AlertBalance** | Pointer to **string** | If balance is below this value, user receive low balance notification. | [optional] 
**AlertPhone** | Pointer to **string** | Low balance notification phone number. | [optional] 
**AlertEmail1** | Pointer to **string** | Low balance notification email 1. | [optional] 
**AlertEmail2** | Pointer to **string** | Low balance notification email 2. | [optional] 
**AlertEmail3** | Pointer to **string** | Low balance notification email 3. | [optional] 

## Methods

### NewUpdateBalanceNotificationSettingsRequest

`func NewUpdateBalanceNotificationSettingsRequest() *UpdateBalanceNotificationSettingsRequest`

NewUpdateBalanceNotificationSettingsRequest instantiates a new UpdateBalanceNotificationSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateBalanceNotificationSettingsRequestWithDefaults

`func NewUpdateBalanceNotificationSettingsRequestWithDefaults() *UpdateBalanceNotificationSettingsRequest`

NewUpdateBalanceNotificationSettingsRequestWithDefaults instantiates a new UpdateBalanceNotificationSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLowBalanceNotification

`func (o *UpdateBalanceNotificationSettingsRequest) GetLowBalanceNotification() bool`

GetLowBalanceNotification returns the LowBalanceNotification field if non-nil, zero value otherwise.

### GetLowBalanceNotificationOk

`func (o *UpdateBalanceNotificationSettingsRequest) GetLowBalanceNotificationOk() (*bool, bool)`

GetLowBalanceNotificationOk returns a tuple with the LowBalanceNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowBalanceNotification

`func (o *UpdateBalanceNotificationSettingsRequest) SetLowBalanceNotification(v bool)`

SetLowBalanceNotification sets LowBalanceNotification field to given value.

### HasLowBalanceNotification

`func (o *UpdateBalanceNotificationSettingsRequest) HasLowBalanceNotification() bool`

HasLowBalanceNotification returns a boolean if a field has been set.

### GetAlertBalance

`func (o *UpdateBalanceNotificationSettingsRequest) GetAlertBalance() string`

GetAlertBalance returns the AlertBalance field if non-nil, zero value otherwise.

### GetAlertBalanceOk

`func (o *UpdateBalanceNotificationSettingsRequest) GetAlertBalanceOk() (*string, bool)`

GetAlertBalanceOk returns a tuple with the AlertBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertBalance

`func (o *UpdateBalanceNotificationSettingsRequest) SetAlertBalance(v string)`

SetAlertBalance sets AlertBalance field to given value.

### HasAlertBalance

`func (o *UpdateBalanceNotificationSettingsRequest) HasAlertBalance() bool`

HasAlertBalance returns a boolean if a field has been set.

### GetAlertPhone

`func (o *UpdateBalanceNotificationSettingsRequest) GetAlertPhone() string`

GetAlertPhone returns the AlertPhone field if non-nil, zero value otherwise.

### GetAlertPhoneOk

`func (o *UpdateBalanceNotificationSettingsRequest) GetAlertPhoneOk() (*string, bool)`

GetAlertPhoneOk returns a tuple with the AlertPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertPhone

`func (o *UpdateBalanceNotificationSettingsRequest) SetAlertPhone(v string)`

SetAlertPhone sets AlertPhone field to given value.

### HasAlertPhone

`func (o *UpdateBalanceNotificationSettingsRequest) HasAlertPhone() bool`

HasAlertPhone returns a boolean if a field has been set.

### GetAlertEmail1

`func (o *UpdateBalanceNotificationSettingsRequest) GetAlertEmail1() string`

GetAlertEmail1 returns the AlertEmail1 field if non-nil, zero value otherwise.

### GetAlertEmail1Ok

`func (o *UpdateBalanceNotificationSettingsRequest) GetAlertEmail1Ok() (*string, bool)`

GetAlertEmail1Ok returns a tuple with the AlertEmail1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertEmail1

`func (o *UpdateBalanceNotificationSettingsRequest) SetAlertEmail1(v string)`

SetAlertEmail1 sets AlertEmail1 field to given value.

### HasAlertEmail1

`func (o *UpdateBalanceNotificationSettingsRequest) HasAlertEmail1() bool`

HasAlertEmail1 returns a boolean if a field has been set.

### GetAlertEmail2

`func (o *UpdateBalanceNotificationSettingsRequest) GetAlertEmail2() string`

GetAlertEmail2 returns the AlertEmail2 field if non-nil, zero value otherwise.

### GetAlertEmail2Ok

`func (o *UpdateBalanceNotificationSettingsRequest) GetAlertEmail2Ok() (*string, bool)`

GetAlertEmail2Ok returns a tuple with the AlertEmail2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertEmail2

`func (o *UpdateBalanceNotificationSettingsRequest) SetAlertEmail2(v string)`

SetAlertEmail2 sets AlertEmail2 field to given value.

### HasAlertEmail2

`func (o *UpdateBalanceNotificationSettingsRequest) HasAlertEmail2() bool`

HasAlertEmail2 returns a boolean if a field has been set.

### GetAlertEmail3

`func (o *UpdateBalanceNotificationSettingsRequest) GetAlertEmail3() string`

GetAlertEmail3 returns the AlertEmail3 field if non-nil, zero value otherwise.

### GetAlertEmail3Ok

`func (o *UpdateBalanceNotificationSettingsRequest) GetAlertEmail3Ok() (*string, bool)`

GetAlertEmail3Ok returns a tuple with the AlertEmail3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertEmail3

`func (o *UpdateBalanceNotificationSettingsRequest) SetAlertEmail3(v string)`

SetAlertEmail3 sets AlertEmail3 field to given value.

### HasAlertEmail3

`func (o *UpdateBalanceNotificationSettingsRequest) HasAlertEmail3() bool`

HasAlertEmail3 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


