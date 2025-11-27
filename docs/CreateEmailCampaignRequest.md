# CreateEmailCampaignRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailSenderId** | Pointer to **int32** | Email sender ID from your configured senders. | [optional] 
**Subject** | Pointer to **string** | Email subject line. | [optional] 
**Message** | Pointer to **string** | HTML email content. | [optional] 
**FromName** | Pointer to **NullableString** | Optional custom sender name. | [optional] 
**ReplyToEmail** | Pointer to **NullableString** | Optional custom reply-to email address. | [optional] 
**Recipients** | Pointer to [**CreateEmailCampaignRequestRecipients**](CreateEmailCampaignRequestRecipients.md) |  | [optional] 

## Methods

### NewCreateEmailCampaignRequest

`func NewCreateEmailCampaignRequest() *CreateEmailCampaignRequest`

NewCreateEmailCampaignRequest instantiates a new CreateEmailCampaignRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEmailCampaignRequestWithDefaults

`func NewCreateEmailCampaignRequestWithDefaults() *CreateEmailCampaignRequest`

NewCreateEmailCampaignRequestWithDefaults instantiates a new CreateEmailCampaignRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailSenderId

`func (o *CreateEmailCampaignRequest) GetEmailSenderId() int32`

GetEmailSenderId returns the EmailSenderId field if non-nil, zero value otherwise.

### GetEmailSenderIdOk

`func (o *CreateEmailCampaignRequest) GetEmailSenderIdOk() (*int32, bool)`

GetEmailSenderIdOk returns a tuple with the EmailSenderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailSenderId

`func (o *CreateEmailCampaignRequest) SetEmailSenderId(v int32)`

SetEmailSenderId sets EmailSenderId field to given value.

### HasEmailSenderId

`func (o *CreateEmailCampaignRequest) HasEmailSenderId() bool`

HasEmailSenderId returns a boolean if a field has been set.

### GetSubject

`func (o *CreateEmailCampaignRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CreateEmailCampaignRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CreateEmailCampaignRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *CreateEmailCampaignRequest) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetMessage

`func (o *CreateEmailCampaignRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CreateEmailCampaignRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CreateEmailCampaignRequest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CreateEmailCampaignRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetFromName

`func (o *CreateEmailCampaignRequest) GetFromName() string`

GetFromName returns the FromName field if non-nil, zero value otherwise.

### GetFromNameOk

`func (o *CreateEmailCampaignRequest) GetFromNameOk() (*string, bool)`

GetFromNameOk returns a tuple with the FromName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromName

`func (o *CreateEmailCampaignRequest) SetFromName(v string)`

SetFromName sets FromName field to given value.

### HasFromName

`func (o *CreateEmailCampaignRequest) HasFromName() bool`

HasFromName returns a boolean if a field has been set.

### SetFromNameNil

`func (o *CreateEmailCampaignRequest) SetFromNameNil(b bool)`

 SetFromNameNil sets the value for FromName to be an explicit nil

### UnsetFromName
`func (o *CreateEmailCampaignRequest) UnsetFromName()`

UnsetFromName ensures that no value is present for FromName, not even an explicit nil
### GetReplyToEmail

`func (o *CreateEmailCampaignRequest) GetReplyToEmail() string`

GetReplyToEmail returns the ReplyToEmail field if non-nil, zero value otherwise.

### GetReplyToEmailOk

`func (o *CreateEmailCampaignRequest) GetReplyToEmailOk() (*string, bool)`

GetReplyToEmailOk returns a tuple with the ReplyToEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyToEmail

`func (o *CreateEmailCampaignRequest) SetReplyToEmail(v string)`

SetReplyToEmail sets ReplyToEmail field to given value.

### HasReplyToEmail

`func (o *CreateEmailCampaignRequest) HasReplyToEmail() bool`

HasReplyToEmail returns a boolean if a field has been set.

### SetReplyToEmailNil

`func (o *CreateEmailCampaignRequest) SetReplyToEmailNil(b bool)`

 SetReplyToEmailNil sets the value for ReplyToEmail to be an explicit nil

### UnsetReplyToEmail
`func (o *CreateEmailCampaignRequest) UnsetReplyToEmail()`

UnsetReplyToEmail ensures that no value is present for ReplyToEmail, not even an explicit nil
### GetRecipients

`func (o *CreateEmailCampaignRequest) GetRecipients() CreateEmailCampaignRequestRecipients`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *CreateEmailCampaignRequest) GetRecipientsOk() (*CreateEmailCampaignRequestRecipients, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *CreateEmailCampaignRequest) SetRecipients(v CreateEmailCampaignRequestRecipients)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *CreateEmailCampaignRequest) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


