# GetBalanceNotificationSettingsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LowBalanceNotification** | **bool** | Should user receive low balance notification. | 
**AlertBalance** | **float32** | If balance is below this value, user receive low balance notification. | 
**AlertPhone** | **string** | Low balance notification phone number. | 
**AlertEmail1** | **string** | Low balance notification email 1. | 
**AlertEmail2** | **string** | Low balance notification email 2. | 
**AlertEmail3** | **string** | Low balance notification email 3. | 

## Methods

### NewGetBalanceNotificationSettingsResponse

`func NewGetBalanceNotificationSettingsResponse(lowBalanceNotification bool, alertBalance float32, alertPhone string, alertEmail1 string, alertEmail2 string, alertEmail3 string, ) *GetBalanceNotificationSettingsResponse`

NewGetBalanceNotificationSettingsResponse instantiates a new GetBalanceNotificationSettingsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBalanceNotificationSettingsResponseWithDefaults

`func NewGetBalanceNotificationSettingsResponseWithDefaults() *GetBalanceNotificationSettingsResponse`

NewGetBalanceNotificationSettingsResponseWithDefaults instantiates a new GetBalanceNotificationSettingsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLowBalanceNotification

`func (o *GetBalanceNotificationSettingsResponse) GetLowBalanceNotification() bool`

GetLowBalanceNotification returns the LowBalanceNotification field if non-nil, zero value otherwise.

### GetLowBalanceNotificationOk

`func (o *GetBalanceNotificationSettingsResponse) GetLowBalanceNotificationOk() (*bool, bool)`

GetLowBalanceNotificationOk returns a tuple with the LowBalanceNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowBalanceNotification

`func (o *GetBalanceNotificationSettingsResponse) SetLowBalanceNotification(v bool)`

SetLowBalanceNotification sets LowBalanceNotification field to given value.


### GetAlertBalance

`func (o *GetBalanceNotificationSettingsResponse) GetAlertBalance() float32`

GetAlertBalance returns the AlertBalance field if non-nil, zero value otherwise.

### GetAlertBalanceOk

`func (o *GetBalanceNotificationSettingsResponse) GetAlertBalanceOk() (*float32, bool)`

GetAlertBalanceOk returns a tuple with the AlertBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertBalance

`func (o *GetBalanceNotificationSettingsResponse) SetAlertBalance(v float32)`

SetAlertBalance sets AlertBalance field to given value.


### GetAlertPhone

`func (o *GetBalanceNotificationSettingsResponse) GetAlertPhone() string`

GetAlertPhone returns the AlertPhone field if non-nil, zero value otherwise.

### GetAlertPhoneOk

`func (o *GetBalanceNotificationSettingsResponse) GetAlertPhoneOk() (*string, bool)`

GetAlertPhoneOk returns a tuple with the AlertPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertPhone

`func (o *GetBalanceNotificationSettingsResponse) SetAlertPhone(v string)`

SetAlertPhone sets AlertPhone field to given value.


### GetAlertEmail1

`func (o *GetBalanceNotificationSettingsResponse) GetAlertEmail1() string`

GetAlertEmail1 returns the AlertEmail1 field if non-nil, zero value otherwise.

### GetAlertEmail1Ok

`func (o *GetBalanceNotificationSettingsResponse) GetAlertEmail1Ok() (*string, bool)`

GetAlertEmail1Ok returns a tuple with the AlertEmail1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertEmail1

`func (o *GetBalanceNotificationSettingsResponse) SetAlertEmail1(v string)`

SetAlertEmail1 sets AlertEmail1 field to given value.


### GetAlertEmail2

`func (o *GetBalanceNotificationSettingsResponse) GetAlertEmail2() string`

GetAlertEmail2 returns the AlertEmail2 field if non-nil, zero value otherwise.

### GetAlertEmail2Ok

`func (o *GetBalanceNotificationSettingsResponse) GetAlertEmail2Ok() (*string, bool)`

GetAlertEmail2Ok returns a tuple with the AlertEmail2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertEmail2

`func (o *GetBalanceNotificationSettingsResponse) SetAlertEmail2(v string)`

SetAlertEmail2 sets AlertEmail2 field to given value.


### GetAlertEmail3

`func (o *GetBalanceNotificationSettingsResponse) GetAlertEmail3() string`

GetAlertEmail3 returns the AlertEmail3 field if non-nil, zero value otherwise.

### GetAlertEmail3Ok

`func (o *GetBalanceNotificationSettingsResponse) GetAlertEmail3Ok() (*string, bool)`

GetAlertEmail3Ok returns a tuple with the AlertEmail3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertEmail3

`func (o *GetBalanceNotificationSettingsResponse) SetAlertEmail3(v string)`

SetAlertEmail3 sets AlertEmail3 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


