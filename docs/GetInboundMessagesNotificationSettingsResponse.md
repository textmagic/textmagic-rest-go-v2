# GetInboundMessagesNotificationSettingsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InboundMessageNotification** | **bool** | Should user receive notification about new incoming messages. | 
**ForwardedCallNotification** | **bool** | Should user receive notification about new forwarded calls. | 
**IncludeSmsHistory** | **bool** | Include SMS history into notification Email. | 
**SendInHtmlFormat** | **bool** | Send Email notification in HTML format. | 
**AlertEmail1** | **string** | New message notification email 1. | 
**AlertEmail2** | **string** | New message notification email 2. | 
**AlertEmail3** | **string** | New message notification email 3. | 

## Methods

### NewGetInboundMessagesNotificationSettingsResponse

`func NewGetInboundMessagesNotificationSettingsResponse(inboundMessageNotification bool, forwardedCallNotification bool, includeSmsHistory bool, sendInHtmlFormat bool, alertEmail1 string, alertEmail2 string, alertEmail3 string, ) *GetInboundMessagesNotificationSettingsResponse`

NewGetInboundMessagesNotificationSettingsResponse instantiates a new GetInboundMessagesNotificationSettingsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInboundMessagesNotificationSettingsResponseWithDefaults

`func NewGetInboundMessagesNotificationSettingsResponseWithDefaults() *GetInboundMessagesNotificationSettingsResponse`

NewGetInboundMessagesNotificationSettingsResponseWithDefaults instantiates a new GetInboundMessagesNotificationSettingsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInboundMessageNotification

`func (o *GetInboundMessagesNotificationSettingsResponse) GetInboundMessageNotification() bool`

GetInboundMessageNotification returns the InboundMessageNotification field if non-nil, zero value otherwise.

### GetInboundMessageNotificationOk

`func (o *GetInboundMessagesNotificationSettingsResponse) GetInboundMessageNotificationOk() (*bool, bool)`

GetInboundMessageNotificationOk returns a tuple with the InboundMessageNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundMessageNotification

`func (o *GetInboundMessagesNotificationSettingsResponse) SetInboundMessageNotification(v bool)`

SetInboundMessageNotification sets InboundMessageNotification field to given value.


### GetForwardedCallNotification

`func (o *GetInboundMessagesNotificationSettingsResponse) GetForwardedCallNotification() bool`

GetForwardedCallNotification returns the ForwardedCallNotification field if non-nil, zero value otherwise.

### GetForwardedCallNotificationOk

`func (o *GetInboundMessagesNotificationSettingsResponse) GetForwardedCallNotificationOk() (*bool, bool)`

GetForwardedCallNotificationOk returns a tuple with the ForwardedCallNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardedCallNotification

`func (o *GetInboundMessagesNotificationSettingsResponse) SetForwardedCallNotification(v bool)`

SetForwardedCallNotification sets ForwardedCallNotification field to given value.


### GetIncludeSmsHistory

`func (o *GetInboundMessagesNotificationSettingsResponse) GetIncludeSmsHistory() bool`

GetIncludeSmsHistory returns the IncludeSmsHistory field if non-nil, zero value otherwise.

### GetIncludeSmsHistoryOk

`func (o *GetInboundMessagesNotificationSettingsResponse) GetIncludeSmsHistoryOk() (*bool, bool)`

GetIncludeSmsHistoryOk returns a tuple with the IncludeSmsHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSmsHistory

`func (o *GetInboundMessagesNotificationSettingsResponse) SetIncludeSmsHistory(v bool)`

SetIncludeSmsHistory sets IncludeSmsHistory field to given value.


### GetSendInHtmlFormat

`func (o *GetInboundMessagesNotificationSettingsResponse) GetSendInHtmlFormat() bool`

GetSendInHtmlFormat returns the SendInHtmlFormat field if non-nil, zero value otherwise.

### GetSendInHtmlFormatOk

`func (o *GetInboundMessagesNotificationSettingsResponse) GetSendInHtmlFormatOk() (*bool, bool)`

GetSendInHtmlFormatOk returns a tuple with the SendInHtmlFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendInHtmlFormat

`func (o *GetInboundMessagesNotificationSettingsResponse) SetSendInHtmlFormat(v bool)`

SetSendInHtmlFormat sets SendInHtmlFormat field to given value.


### GetAlertEmail1

`func (o *GetInboundMessagesNotificationSettingsResponse) GetAlertEmail1() string`

GetAlertEmail1 returns the AlertEmail1 field if non-nil, zero value otherwise.

### GetAlertEmail1Ok

`func (o *GetInboundMessagesNotificationSettingsResponse) GetAlertEmail1Ok() (*string, bool)`

GetAlertEmail1Ok returns a tuple with the AlertEmail1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertEmail1

`func (o *GetInboundMessagesNotificationSettingsResponse) SetAlertEmail1(v string)`

SetAlertEmail1 sets AlertEmail1 field to given value.


### GetAlertEmail2

`func (o *GetInboundMessagesNotificationSettingsResponse) GetAlertEmail2() string`

GetAlertEmail2 returns the AlertEmail2 field if non-nil, zero value otherwise.

### GetAlertEmail2Ok

`func (o *GetInboundMessagesNotificationSettingsResponse) GetAlertEmail2Ok() (*string, bool)`

GetAlertEmail2Ok returns a tuple with the AlertEmail2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertEmail2

`func (o *GetInboundMessagesNotificationSettingsResponse) SetAlertEmail2(v string)`

SetAlertEmail2 sets AlertEmail2 field to given value.


### GetAlertEmail3

`func (o *GetInboundMessagesNotificationSettingsResponse) GetAlertEmail3() string`

GetAlertEmail3 returns the AlertEmail3 field if non-nil, zero value otherwise.

### GetAlertEmail3Ok

`func (o *GetInboundMessagesNotificationSettingsResponse) GetAlertEmail3Ok() (*string, bool)`

GetAlertEmail3Ok returns a tuple with the AlertEmail3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertEmail3

`func (o *GetInboundMessagesNotificationSettingsResponse) SetAlertEmail3(v string)`

SetAlertEmail3 sets AlertEmail3 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


