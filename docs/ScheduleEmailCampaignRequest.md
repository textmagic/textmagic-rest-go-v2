# ScheduleEmailCampaignRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailSenderId** | Pointer to **int32** | Email sender ID from your configured senders. | [optional] 
**Subject** | Pointer to **string** | Email subject line. | [optional] 
**Message** | Pointer to **string** | HTML email content. | [optional] 
**FromName** | Pointer to **NullableString** | Optional custom sender name. | [optional] 
**ReplyToEmail** | Pointer to **NullableString** | Optional custom reply-to email address. | [optional] 
**Recipients** | Pointer to [**ScheduleEmailCampaignRequestRecipients**](ScheduleEmailCampaignRequestRecipients.md) |  | [optional] 
**ScheduleParams** | Pointer to [**ScheduleEmailCampaignRequestScheduleParams**](ScheduleEmailCampaignRequestScheduleParams.md) |  | [optional] 

## Methods

### NewScheduleEmailCampaignRequest

`func NewScheduleEmailCampaignRequest() *ScheduleEmailCampaignRequest`

NewScheduleEmailCampaignRequest instantiates a new ScheduleEmailCampaignRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleEmailCampaignRequestWithDefaults

`func NewScheduleEmailCampaignRequestWithDefaults() *ScheduleEmailCampaignRequest`

NewScheduleEmailCampaignRequestWithDefaults instantiates a new ScheduleEmailCampaignRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailSenderId

`func (o *ScheduleEmailCampaignRequest) GetEmailSenderId() int32`

GetEmailSenderId returns the EmailSenderId field if non-nil, zero value otherwise.

### GetEmailSenderIdOk

`func (o *ScheduleEmailCampaignRequest) GetEmailSenderIdOk() (*int32, bool)`

GetEmailSenderIdOk returns a tuple with the EmailSenderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailSenderId

`func (o *ScheduleEmailCampaignRequest) SetEmailSenderId(v int32)`

SetEmailSenderId sets EmailSenderId field to given value.

### HasEmailSenderId

`func (o *ScheduleEmailCampaignRequest) HasEmailSenderId() bool`

HasEmailSenderId returns a boolean if a field has been set.

### GetSubject

`func (o *ScheduleEmailCampaignRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *ScheduleEmailCampaignRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *ScheduleEmailCampaignRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *ScheduleEmailCampaignRequest) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetMessage

`func (o *ScheduleEmailCampaignRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ScheduleEmailCampaignRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ScheduleEmailCampaignRequest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ScheduleEmailCampaignRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetFromName

`func (o *ScheduleEmailCampaignRequest) GetFromName() string`

GetFromName returns the FromName field if non-nil, zero value otherwise.

### GetFromNameOk

`func (o *ScheduleEmailCampaignRequest) GetFromNameOk() (*string, bool)`

GetFromNameOk returns a tuple with the FromName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromName

`func (o *ScheduleEmailCampaignRequest) SetFromName(v string)`

SetFromName sets FromName field to given value.

### HasFromName

`func (o *ScheduleEmailCampaignRequest) HasFromName() bool`

HasFromName returns a boolean if a field has been set.

### SetFromNameNil

`func (o *ScheduleEmailCampaignRequest) SetFromNameNil(b bool)`

 SetFromNameNil sets the value for FromName to be an explicit nil

### UnsetFromName
`func (o *ScheduleEmailCampaignRequest) UnsetFromName()`

UnsetFromName ensures that no value is present for FromName, not even an explicit nil
### GetReplyToEmail

`func (o *ScheduleEmailCampaignRequest) GetReplyToEmail() string`

GetReplyToEmail returns the ReplyToEmail field if non-nil, zero value otherwise.

### GetReplyToEmailOk

`func (o *ScheduleEmailCampaignRequest) GetReplyToEmailOk() (*string, bool)`

GetReplyToEmailOk returns a tuple with the ReplyToEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyToEmail

`func (o *ScheduleEmailCampaignRequest) SetReplyToEmail(v string)`

SetReplyToEmail sets ReplyToEmail field to given value.

### HasReplyToEmail

`func (o *ScheduleEmailCampaignRequest) HasReplyToEmail() bool`

HasReplyToEmail returns a boolean if a field has been set.

### SetReplyToEmailNil

`func (o *ScheduleEmailCampaignRequest) SetReplyToEmailNil(b bool)`

 SetReplyToEmailNil sets the value for ReplyToEmail to be an explicit nil

### UnsetReplyToEmail
`func (o *ScheduleEmailCampaignRequest) UnsetReplyToEmail()`

UnsetReplyToEmail ensures that no value is present for ReplyToEmail, not even an explicit nil
### GetRecipients

`func (o *ScheduleEmailCampaignRequest) GetRecipients() ScheduleEmailCampaignRequestRecipients`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *ScheduleEmailCampaignRequest) GetRecipientsOk() (*ScheduleEmailCampaignRequestRecipients, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *ScheduleEmailCampaignRequest) SetRecipients(v ScheduleEmailCampaignRequestRecipients)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *ScheduleEmailCampaignRequest) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.

### GetScheduleParams

`func (o *ScheduleEmailCampaignRequest) GetScheduleParams() ScheduleEmailCampaignRequestScheduleParams`

GetScheduleParams returns the ScheduleParams field if non-nil, zero value otherwise.

### GetScheduleParamsOk

`func (o *ScheduleEmailCampaignRequest) GetScheduleParamsOk() (*ScheduleEmailCampaignRequestScheduleParams, bool)`

GetScheduleParamsOk returns a tuple with the ScheduleParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleParams

`func (o *ScheduleEmailCampaignRequest) SetScheduleParams(v ScheduleEmailCampaignRequestScheduleParams)`

SetScheduleParams sets ScheduleParams field to given value.

### HasScheduleParams

`func (o *ScheduleEmailCampaignRequest) HasScheduleParams() bool`

HasScheduleParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


