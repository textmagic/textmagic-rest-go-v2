# OutboundEmailResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Outbound email ID. | 
**SendTime** | **time.Time** | Email send timestamp. | 
**FromName** | Pointer to **NullableString** | Sender name. | [optional] 
**FromEmail** | **string** | Sender email address. | 
**ReplyToEmail** | **string** | Reply-to email address. | 
**RecipientFullName** | Pointer to **NullableString** | Recipient&#39;s full name. | [optional] 
**RecipientEmail** | Pointer to **NullableString** | Recipient&#39;s email address. | [optional] 
**EmailSubject** | **string** | Email subject line. | 
**EmailContent** | **string** | HTML email content. | 
**Source** | **string** | Source of the outbound email. | 
**Status** | **string** | Current email status. | 
**Cost** | **float32** | Cost of sending this email. | 
**StatusReason** | Pointer to **NullableString** | Detailed status reason. | [optional] 
**ContactId** | Pointer to **NullableInt32** | Associated contact ID. | [optional] 
**InitiatorId** | **int32** | ID of user who initiated the email. | 

## Methods

### NewOutboundEmailResponse

`func NewOutboundEmailResponse(id int32, sendTime time.Time, fromEmail string, replyToEmail string, emailSubject string, emailContent string, source string, status string, cost float32, initiatorId int32, ) *OutboundEmailResponse`

NewOutboundEmailResponse instantiates a new OutboundEmailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutboundEmailResponseWithDefaults

`func NewOutboundEmailResponseWithDefaults() *OutboundEmailResponse`

NewOutboundEmailResponseWithDefaults instantiates a new OutboundEmailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OutboundEmailResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OutboundEmailResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OutboundEmailResponse) SetId(v int32)`

SetId sets Id field to given value.


### GetSendTime

`func (o *OutboundEmailResponse) GetSendTime() time.Time`

GetSendTime returns the SendTime field if non-nil, zero value otherwise.

### GetSendTimeOk

`func (o *OutboundEmailResponse) GetSendTimeOk() (*time.Time, bool)`

GetSendTimeOk returns a tuple with the SendTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendTime

`func (o *OutboundEmailResponse) SetSendTime(v time.Time)`

SetSendTime sets SendTime field to given value.


### GetFromName

`func (o *OutboundEmailResponse) GetFromName() string`

GetFromName returns the FromName field if non-nil, zero value otherwise.

### GetFromNameOk

`func (o *OutboundEmailResponse) GetFromNameOk() (*string, bool)`

GetFromNameOk returns a tuple with the FromName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromName

`func (o *OutboundEmailResponse) SetFromName(v string)`

SetFromName sets FromName field to given value.

### HasFromName

`func (o *OutboundEmailResponse) HasFromName() bool`

HasFromName returns a boolean if a field has been set.

### SetFromNameNil

`func (o *OutboundEmailResponse) SetFromNameNil(b bool)`

 SetFromNameNil sets the value for FromName to be an explicit nil

### UnsetFromName
`func (o *OutboundEmailResponse) UnsetFromName()`

UnsetFromName ensures that no value is present for FromName, not even an explicit nil
### GetFromEmail

`func (o *OutboundEmailResponse) GetFromEmail() string`

GetFromEmail returns the FromEmail field if non-nil, zero value otherwise.

### GetFromEmailOk

`func (o *OutboundEmailResponse) GetFromEmailOk() (*string, bool)`

GetFromEmailOk returns a tuple with the FromEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromEmail

`func (o *OutboundEmailResponse) SetFromEmail(v string)`

SetFromEmail sets FromEmail field to given value.


### GetReplyToEmail

`func (o *OutboundEmailResponse) GetReplyToEmail() string`

GetReplyToEmail returns the ReplyToEmail field if non-nil, zero value otherwise.

### GetReplyToEmailOk

`func (o *OutboundEmailResponse) GetReplyToEmailOk() (*string, bool)`

GetReplyToEmailOk returns a tuple with the ReplyToEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyToEmail

`func (o *OutboundEmailResponse) SetReplyToEmail(v string)`

SetReplyToEmail sets ReplyToEmail field to given value.


### GetRecipientFullName

`func (o *OutboundEmailResponse) GetRecipientFullName() string`

GetRecipientFullName returns the RecipientFullName field if non-nil, zero value otherwise.

### GetRecipientFullNameOk

`func (o *OutboundEmailResponse) GetRecipientFullNameOk() (*string, bool)`

GetRecipientFullNameOk returns a tuple with the RecipientFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientFullName

`func (o *OutboundEmailResponse) SetRecipientFullName(v string)`

SetRecipientFullName sets RecipientFullName field to given value.

### HasRecipientFullName

`func (o *OutboundEmailResponse) HasRecipientFullName() bool`

HasRecipientFullName returns a boolean if a field has been set.

### SetRecipientFullNameNil

`func (o *OutboundEmailResponse) SetRecipientFullNameNil(b bool)`

 SetRecipientFullNameNil sets the value for RecipientFullName to be an explicit nil

### UnsetRecipientFullName
`func (o *OutboundEmailResponse) UnsetRecipientFullName()`

UnsetRecipientFullName ensures that no value is present for RecipientFullName, not even an explicit nil
### GetRecipientEmail

`func (o *OutboundEmailResponse) GetRecipientEmail() string`

GetRecipientEmail returns the RecipientEmail field if non-nil, zero value otherwise.

### GetRecipientEmailOk

`func (o *OutboundEmailResponse) GetRecipientEmailOk() (*string, bool)`

GetRecipientEmailOk returns a tuple with the RecipientEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientEmail

`func (o *OutboundEmailResponse) SetRecipientEmail(v string)`

SetRecipientEmail sets RecipientEmail field to given value.

### HasRecipientEmail

`func (o *OutboundEmailResponse) HasRecipientEmail() bool`

HasRecipientEmail returns a boolean if a field has been set.

### SetRecipientEmailNil

`func (o *OutboundEmailResponse) SetRecipientEmailNil(b bool)`

 SetRecipientEmailNil sets the value for RecipientEmail to be an explicit nil

### UnsetRecipientEmail
`func (o *OutboundEmailResponse) UnsetRecipientEmail()`

UnsetRecipientEmail ensures that no value is present for RecipientEmail, not even an explicit nil
### GetEmailSubject

`func (o *OutboundEmailResponse) GetEmailSubject() string`

GetEmailSubject returns the EmailSubject field if non-nil, zero value otherwise.

### GetEmailSubjectOk

`func (o *OutboundEmailResponse) GetEmailSubjectOk() (*string, bool)`

GetEmailSubjectOk returns a tuple with the EmailSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailSubject

`func (o *OutboundEmailResponse) SetEmailSubject(v string)`

SetEmailSubject sets EmailSubject field to given value.


### GetEmailContent

`func (o *OutboundEmailResponse) GetEmailContent() string`

GetEmailContent returns the EmailContent field if non-nil, zero value otherwise.

### GetEmailContentOk

`func (o *OutboundEmailResponse) GetEmailContentOk() (*string, bool)`

GetEmailContentOk returns a tuple with the EmailContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailContent

`func (o *OutboundEmailResponse) SetEmailContent(v string)`

SetEmailContent sets EmailContent field to given value.


### GetSource

`func (o *OutboundEmailResponse) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *OutboundEmailResponse) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *OutboundEmailResponse) SetSource(v string)`

SetSource sets Source field to given value.


### GetStatus

`func (o *OutboundEmailResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OutboundEmailResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OutboundEmailResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCost

`func (o *OutboundEmailResponse) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *OutboundEmailResponse) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *OutboundEmailResponse) SetCost(v float32)`

SetCost sets Cost field to given value.


### GetStatusReason

`func (o *OutboundEmailResponse) GetStatusReason() string`

GetStatusReason returns the StatusReason field if non-nil, zero value otherwise.

### GetStatusReasonOk

`func (o *OutboundEmailResponse) GetStatusReasonOk() (*string, bool)`

GetStatusReasonOk returns a tuple with the StatusReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusReason

`func (o *OutboundEmailResponse) SetStatusReason(v string)`

SetStatusReason sets StatusReason field to given value.

### HasStatusReason

`func (o *OutboundEmailResponse) HasStatusReason() bool`

HasStatusReason returns a boolean if a field has been set.

### SetStatusReasonNil

`func (o *OutboundEmailResponse) SetStatusReasonNil(b bool)`

 SetStatusReasonNil sets the value for StatusReason to be an explicit nil

### UnsetStatusReason
`func (o *OutboundEmailResponse) UnsetStatusReason()`

UnsetStatusReason ensures that no value is present for StatusReason, not even an explicit nil
### GetContactId

`func (o *OutboundEmailResponse) GetContactId() int32`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *OutboundEmailResponse) GetContactIdOk() (*int32, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *OutboundEmailResponse) SetContactId(v int32)`

SetContactId sets ContactId field to given value.

### HasContactId

`func (o *OutboundEmailResponse) HasContactId() bool`

HasContactId returns a boolean if a field has been set.

### SetContactIdNil

`func (o *OutboundEmailResponse) SetContactIdNil(b bool)`

 SetContactIdNil sets the value for ContactId to be an explicit nil

### UnsetContactId
`func (o *OutboundEmailResponse) UnsetContactId()`

UnsetContactId ensures that no value is present for ContactId, not even an explicit nil
### GetInitiatorId

`func (o *OutboundEmailResponse) GetInitiatorId() int32`

GetInitiatorId returns the InitiatorId field if non-nil, zero value otherwise.

### GetInitiatorIdOk

`func (o *OutboundEmailResponse) GetInitiatorIdOk() (*int32, bool)`

GetInitiatorIdOk returns a tuple with the InitiatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatorId

`func (o *OutboundEmailResponse) SetInitiatorId(v int32)`

SetInitiatorId sets InitiatorId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


