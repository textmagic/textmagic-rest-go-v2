# ScheduledEmailCampaignDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique scheduled campaign ID. | 
**Status** | **string** | Current scheduled campaign status. | 
**EmailSenderId** | Pointer to **NullableInt32** | Email sender ID used for this campaign. | [optional] 
**StartAt** | **time.Time** | Scheduled start timestamp (UTC). | 
**EndAt** | Pointer to **NullableTime** | End timestamp for recurring campaigns (UTC). | [optional] 
**NextSendAt** | Pointer to **NullableTime** | Next scheduled send timestamp (UTC). | [optional] 
**CreatedBy** | [**UserPersonalInfo**](UserPersonalInfo.md) |  | 
**CreatedAt** | **time.Time** | Campaign creation timestamp. | 
**UpdatedAt** | **time.Time** | Last update timestamp. | 
**Type** | **string** | Campaign recurrence type. | 
**FromName** | Pointer to **NullableString** | Sender name displayed in recipient&#39;s inbox. | [optional] 
**FromEmail** | **string** | Sender email address. | 
**ReplyToEmail** | **string** | Reply-to email address. | 
**Subject** | **string** | Email subject line. | 
**Html** | **string** | HTML email content. | 
**RecipientsCount** | **int32** | Number of recipients for this campaign. | 
**SendingTimezone** | **string** | Timezone for sending the campaign. | 
**Rrule** | Pointer to **NullableString** | RFC 5545 recurrence rule for recurring campaigns. | [optional] 
**OccurrenceSummary** | Pointer to **NullableString** | Human-readable schedule description. | [optional] 

## Methods

### NewScheduledEmailCampaignDetails

`func NewScheduledEmailCampaignDetails(id int32, status string, startAt time.Time, createdBy UserPersonalInfo, createdAt time.Time, updatedAt time.Time, type_ string, fromEmail string, replyToEmail string, subject string, html string, recipientsCount int32, sendingTimezone string, ) *ScheduledEmailCampaignDetails`

NewScheduledEmailCampaignDetails instantiates a new ScheduledEmailCampaignDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduledEmailCampaignDetailsWithDefaults

`func NewScheduledEmailCampaignDetailsWithDefaults() *ScheduledEmailCampaignDetails`

NewScheduledEmailCampaignDetailsWithDefaults instantiates a new ScheduledEmailCampaignDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ScheduledEmailCampaignDetails) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScheduledEmailCampaignDetails) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScheduledEmailCampaignDetails) SetId(v int32)`

SetId sets Id field to given value.


### GetStatus

`func (o *ScheduledEmailCampaignDetails) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ScheduledEmailCampaignDetails) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ScheduledEmailCampaignDetails) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetEmailSenderId

`func (o *ScheduledEmailCampaignDetails) GetEmailSenderId() int32`

GetEmailSenderId returns the EmailSenderId field if non-nil, zero value otherwise.

### GetEmailSenderIdOk

`func (o *ScheduledEmailCampaignDetails) GetEmailSenderIdOk() (*int32, bool)`

GetEmailSenderIdOk returns a tuple with the EmailSenderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailSenderId

`func (o *ScheduledEmailCampaignDetails) SetEmailSenderId(v int32)`

SetEmailSenderId sets EmailSenderId field to given value.

### HasEmailSenderId

`func (o *ScheduledEmailCampaignDetails) HasEmailSenderId() bool`

HasEmailSenderId returns a boolean if a field has been set.

### SetEmailSenderIdNil

`func (o *ScheduledEmailCampaignDetails) SetEmailSenderIdNil(b bool)`

 SetEmailSenderIdNil sets the value for EmailSenderId to be an explicit nil

### UnsetEmailSenderId
`func (o *ScheduledEmailCampaignDetails) UnsetEmailSenderId()`

UnsetEmailSenderId ensures that no value is present for EmailSenderId, not even an explicit nil
### GetStartAt

`func (o *ScheduledEmailCampaignDetails) GetStartAt() time.Time`

GetStartAt returns the StartAt field if non-nil, zero value otherwise.

### GetStartAtOk

`func (o *ScheduledEmailCampaignDetails) GetStartAtOk() (*time.Time, bool)`

GetStartAtOk returns a tuple with the StartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAt

`func (o *ScheduledEmailCampaignDetails) SetStartAt(v time.Time)`

SetStartAt sets StartAt field to given value.


### GetEndAt

`func (o *ScheduledEmailCampaignDetails) GetEndAt() time.Time`

GetEndAt returns the EndAt field if non-nil, zero value otherwise.

### GetEndAtOk

`func (o *ScheduledEmailCampaignDetails) GetEndAtOk() (*time.Time, bool)`

GetEndAtOk returns a tuple with the EndAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAt

`func (o *ScheduledEmailCampaignDetails) SetEndAt(v time.Time)`

SetEndAt sets EndAt field to given value.

### HasEndAt

`func (o *ScheduledEmailCampaignDetails) HasEndAt() bool`

HasEndAt returns a boolean if a field has been set.

### SetEndAtNil

`func (o *ScheduledEmailCampaignDetails) SetEndAtNil(b bool)`

 SetEndAtNil sets the value for EndAt to be an explicit nil

### UnsetEndAt
`func (o *ScheduledEmailCampaignDetails) UnsetEndAt()`

UnsetEndAt ensures that no value is present for EndAt, not even an explicit nil
### GetNextSendAt

`func (o *ScheduledEmailCampaignDetails) GetNextSendAt() time.Time`

GetNextSendAt returns the NextSendAt field if non-nil, zero value otherwise.

### GetNextSendAtOk

`func (o *ScheduledEmailCampaignDetails) GetNextSendAtOk() (*time.Time, bool)`

GetNextSendAtOk returns a tuple with the NextSendAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextSendAt

`func (o *ScheduledEmailCampaignDetails) SetNextSendAt(v time.Time)`

SetNextSendAt sets NextSendAt field to given value.

### HasNextSendAt

`func (o *ScheduledEmailCampaignDetails) HasNextSendAt() bool`

HasNextSendAt returns a boolean if a field has been set.

### SetNextSendAtNil

`func (o *ScheduledEmailCampaignDetails) SetNextSendAtNil(b bool)`

 SetNextSendAtNil sets the value for NextSendAt to be an explicit nil

### UnsetNextSendAt
`func (o *ScheduledEmailCampaignDetails) UnsetNextSendAt()`

UnsetNextSendAt ensures that no value is present for NextSendAt, not even an explicit nil
### GetCreatedBy

`func (o *ScheduledEmailCampaignDetails) GetCreatedBy() UserPersonalInfo`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ScheduledEmailCampaignDetails) GetCreatedByOk() (*UserPersonalInfo, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ScheduledEmailCampaignDetails) SetCreatedBy(v UserPersonalInfo)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreatedAt

`func (o *ScheduledEmailCampaignDetails) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ScheduledEmailCampaignDetails) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ScheduledEmailCampaignDetails) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ScheduledEmailCampaignDetails) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ScheduledEmailCampaignDetails) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ScheduledEmailCampaignDetails) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetType

`func (o *ScheduledEmailCampaignDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ScheduledEmailCampaignDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ScheduledEmailCampaignDetails) SetType(v string)`

SetType sets Type field to given value.


### GetFromName

`func (o *ScheduledEmailCampaignDetails) GetFromName() string`

GetFromName returns the FromName field if non-nil, zero value otherwise.

### GetFromNameOk

`func (o *ScheduledEmailCampaignDetails) GetFromNameOk() (*string, bool)`

GetFromNameOk returns a tuple with the FromName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromName

`func (o *ScheduledEmailCampaignDetails) SetFromName(v string)`

SetFromName sets FromName field to given value.

### HasFromName

`func (o *ScheduledEmailCampaignDetails) HasFromName() bool`

HasFromName returns a boolean if a field has been set.

### SetFromNameNil

`func (o *ScheduledEmailCampaignDetails) SetFromNameNil(b bool)`

 SetFromNameNil sets the value for FromName to be an explicit nil

### UnsetFromName
`func (o *ScheduledEmailCampaignDetails) UnsetFromName()`

UnsetFromName ensures that no value is present for FromName, not even an explicit nil
### GetFromEmail

`func (o *ScheduledEmailCampaignDetails) GetFromEmail() string`

GetFromEmail returns the FromEmail field if non-nil, zero value otherwise.

### GetFromEmailOk

`func (o *ScheduledEmailCampaignDetails) GetFromEmailOk() (*string, bool)`

GetFromEmailOk returns a tuple with the FromEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromEmail

`func (o *ScheduledEmailCampaignDetails) SetFromEmail(v string)`

SetFromEmail sets FromEmail field to given value.


### GetReplyToEmail

`func (o *ScheduledEmailCampaignDetails) GetReplyToEmail() string`

GetReplyToEmail returns the ReplyToEmail field if non-nil, zero value otherwise.

### GetReplyToEmailOk

`func (o *ScheduledEmailCampaignDetails) GetReplyToEmailOk() (*string, bool)`

GetReplyToEmailOk returns a tuple with the ReplyToEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyToEmail

`func (o *ScheduledEmailCampaignDetails) SetReplyToEmail(v string)`

SetReplyToEmail sets ReplyToEmail field to given value.


### GetSubject

`func (o *ScheduledEmailCampaignDetails) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *ScheduledEmailCampaignDetails) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *ScheduledEmailCampaignDetails) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetHtml

`func (o *ScheduledEmailCampaignDetails) GetHtml() string`

GetHtml returns the Html field if non-nil, zero value otherwise.

### GetHtmlOk

`func (o *ScheduledEmailCampaignDetails) GetHtmlOk() (*string, bool)`

GetHtmlOk returns a tuple with the Html field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtml

`func (o *ScheduledEmailCampaignDetails) SetHtml(v string)`

SetHtml sets Html field to given value.


### GetRecipientsCount

`func (o *ScheduledEmailCampaignDetails) GetRecipientsCount() int32`

GetRecipientsCount returns the RecipientsCount field if non-nil, zero value otherwise.

### GetRecipientsCountOk

`func (o *ScheduledEmailCampaignDetails) GetRecipientsCountOk() (*int32, bool)`

GetRecipientsCountOk returns a tuple with the RecipientsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientsCount

`func (o *ScheduledEmailCampaignDetails) SetRecipientsCount(v int32)`

SetRecipientsCount sets RecipientsCount field to given value.


### GetSendingTimezone

`func (o *ScheduledEmailCampaignDetails) GetSendingTimezone() string`

GetSendingTimezone returns the SendingTimezone field if non-nil, zero value otherwise.

### GetSendingTimezoneOk

`func (o *ScheduledEmailCampaignDetails) GetSendingTimezoneOk() (*string, bool)`

GetSendingTimezoneOk returns a tuple with the SendingTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendingTimezone

`func (o *ScheduledEmailCampaignDetails) SetSendingTimezone(v string)`

SetSendingTimezone sets SendingTimezone field to given value.


### GetRrule

`func (o *ScheduledEmailCampaignDetails) GetRrule() string`

GetRrule returns the Rrule field if non-nil, zero value otherwise.

### GetRruleOk

`func (o *ScheduledEmailCampaignDetails) GetRruleOk() (*string, bool)`

GetRruleOk returns a tuple with the Rrule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrule

`func (o *ScheduledEmailCampaignDetails) SetRrule(v string)`

SetRrule sets Rrule field to given value.

### HasRrule

`func (o *ScheduledEmailCampaignDetails) HasRrule() bool`

HasRrule returns a boolean if a field has been set.

### SetRruleNil

`func (o *ScheduledEmailCampaignDetails) SetRruleNil(b bool)`

 SetRruleNil sets the value for Rrule to be an explicit nil

### UnsetRrule
`func (o *ScheduledEmailCampaignDetails) UnsetRrule()`

UnsetRrule ensures that no value is present for Rrule, not even an explicit nil
### GetOccurrenceSummary

`func (o *ScheduledEmailCampaignDetails) GetOccurrenceSummary() string`

GetOccurrenceSummary returns the OccurrenceSummary field if non-nil, zero value otherwise.

### GetOccurrenceSummaryOk

`func (o *ScheduledEmailCampaignDetails) GetOccurrenceSummaryOk() (*string, bool)`

GetOccurrenceSummaryOk returns a tuple with the OccurrenceSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurrenceSummary

`func (o *ScheduledEmailCampaignDetails) SetOccurrenceSummary(v string)`

SetOccurrenceSummary sets OccurrenceSummary field to given value.

### HasOccurrenceSummary

`func (o *ScheduledEmailCampaignDetails) HasOccurrenceSummary() bool`

HasOccurrenceSummary returns a boolean if a field has been set.

### SetOccurrenceSummaryNil

`func (o *ScheduledEmailCampaignDetails) SetOccurrenceSummaryNil(b bool)`

 SetOccurrenceSummaryNil sets the value for OccurrenceSummary to be an explicit nil

### UnsetOccurrenceSummary
`func (o *ScheduledEmailCampaignDetails) UnsetOccurrenceSummary()`

UnsetOccurrenceSummary ensures that no value is present for OccurrenceSummary, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


