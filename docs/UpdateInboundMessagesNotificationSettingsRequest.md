# UpdateInboundMessagesNotificationSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InboundMessageNotification** | Pointer to **bool** | Should user receive notification about new incoming messages. | [optional] 
**IncludeSmsHistory** | Pointer to **bool** | Include SMS history into notification Email. | [optional] 
**SendInHtmlFormat** | Pointer to **bool** | Send Email notification in HTML format. | [optional] 
**AlertEmail1** | Pointer to **string** | New message notification email 2. | [optional] 
**AlertEmail2** | Pointer to **string** | New message notification email 2. | [optional] 
**AlertEmail3** | Pointer to **string** | New message notification email 3. | [optional] 

## Methods

### NewUpdateInboundMessagesNotificationSettingsRequest

`func NewUpdateInboundMessagesNotificationSettingsRequest() *UpdateInboundMessagesNotificationSettingsRequest`

NewUpdateInboundMessagesNotificationSettingsRequest instantiates a new UpdateInboundMessagesNotificationSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateInboundMessagesNotificationSettingsRequestWithDefaults

`func NewUpdateInboundMessagesNotificationSettingsRequestWithDefaults() *UpdateInboundMessagesNotificationSettingsRequest`

NewUpdateInboundMessagesNotificationSettingsRequestWithDefaults instantiates a new UpdateInboundMessagesNotificationSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInboundMessageNotification

`func (o *UpdateInboundMessagesNotificationSettingsRequest) GetInboundMessageNotification() bool`

GetInboundMessageNotification returns the InboundMessageNotification field if non-nil, zero value otherwise.

### GetInboundMessageNotificationOk

`func (o *UpdateInboundMessagesNotificationSettingsRequest) GetInboundMessageNotificationOk() (*bool, bool)`

GetInboundMessageNotificationOk returns a tuple with the InboundMessageNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundMessageNotification

`func (o *UpdateInboundMessagesNotificationSettingsRequest) SetInboundMessageNotification(v bool)`

SetInboundMessageNotification sets InboundMessageNotification field to given value.

### HasInboundMessageNotification

`func (o *UpdateInboundMessagesNotificationSettingsRequest) HasInboundMessageNotification() bool`

HasInboundMessageNotification returns a boolean if a field has been set.

### GetIncludeSmsHistory

`func (o *UpdateInboundMessagesNotificationSettingsRequest) GetIncludeSmsHistory() bool`

GetIncludeSmsHistory returns the IncludeSmsHistory field if non-nil, zero value otherwise.

### GetIncludeSmsHistoryOk

`func (o *UpdateInboundMessagesNotificationSettingsRequest) GetIncludeSmsHistoryOk() (*bool, bool)`

GetIncludeSmsHistoryOk returns a tuple with the IncludeSmsHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSmsHistory

`func (o *UpdateInboundMessagesNotificationSettingsRequest) SetIncludeSmsHistory(v bool)`

SetIncludeSmsHistory sets IncludeSmsHistory field to given value.

### HasIncludeSmsHistory

`func (o *UpdateInboundMessagesNotificationSettingsRequest) HasIncludeSmsHistory() bool`

HasIncludeSmsHistory returns a boolean if a field has been set.

### GetSendInHtmlFormat

`func (o *UpdateInboundMessagesNotificationSettingsRequest) GetSendInHtmlFormat() bool`

GetSendInHtmlFormat returns the SendInHtmlFormat field if non-nil, zero value otherwise.

### GetSendInHtmlFormatOk

`func (o *UpdateInboundMessagesNotificationSettingsRequest) GetSendInHtmlFormatOk() (*bool, bool)`

GetSendInHtmlFormatOk returns a tuple with the SendInHtmlFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendInHtmlFormat

`func (o *UpdateInboundMessagesNotificationSettingsRequest) SetSendInHtmlFormat(v bool)`

SetSendInHtmlFormat sets SendInHtmlFormat field to given value.

### HasSendInHtmlFormat

`func (o *UpdateInboundMessagesNotificationSettingsRequest) HasSendInHtmlFormat() bool`

HasSendInHtmlFormat returns a boolean if a field has been set.

### GetAlertEmail1

`func (o *UpdateInboundMessagesNotificationSettingsRequest) GetAlertEmail1() string`

GetAlertEmail1 returns the AlertEmail1 field if non-nil, zero value otherwise.

### GetAlertEmail1Ok

`func (o *UpdateInboundMessagesNotificationSettingsRequest) GetAlertEmail1Ok() (*string, bool)`

GetAlertEmail1Ok returns a tuple with the AlertEmail1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertEmail1

`func (o *UpdateInboundMessagesNotificationSettingsRequest) SetAlertEmail1(v string)`

SetAlertEmail1 sets AlertEmail1 field to given value.

### HasAlertEmail1

`func (o *UpdateInboundMessagesNotificationSettingsRequest) HasAlertEmail1() bool`

HasAlertEmail1 returns a boolean if a field has been set.

### GetAlertEmail2

`func (o *UpdateInboundMessagesNotificationSettingsRequest) GetAlertEmail2() string`

GetAlertEmail2 returns the AlertEmail2 field if non-nil, zero value otherwise.

### GetAlertEmail2Ok

`func (o *UpdateInboundMessagesNotificationSettingsRequest) GetAlertEmail2Ok() (*string, bool)`

GetAlertEmail2Ok returns a tuple with the AlertEmail2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertEmail2

`func (o *UpdateInboundMessagesNotificationSettingsRequest) SetAlertEmail2(v string)`

SetAlertEmail2 sets AlertEmail2 field to given value.

### HasAlertEmail2

`func (o *UpdateInboundMessagesNotificationSettingsRequest) HasAlertEmail2() bool`

HasAlertEmail2 returns a boolean if a field has been set.

### GetAlertEmail3

`func (o *UpdateInboundMessagesNotificationSettingsRequest) GetAlertEmail3() string`

GetAlertEmail3 returns the AlertEmail3 field if non-nil, zero value otherwise.

### GetAlertEmail3Ok

`func (o *UpdateInboundMessagesNotificationSettingsRequest) GetAlertEmail3Ok() (*string, bool)`

GetAlertEmail3Ok returns a tuple with the AlertEmail3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertEmail3

`func (o *UpdateInboundMessagesNotificationSettingsRequest) SetAlertEmail3(v string)`

SetAlertEmail3 sets AlertEmail3 field to given value.

### HasAlertEmail3

`func (o *UpdateInboundMessagesNotificationSettingsRequest) HasAlertEmail3() bool`

HasAlertEmail3 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


