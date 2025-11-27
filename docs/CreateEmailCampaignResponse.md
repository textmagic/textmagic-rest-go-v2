# CreateEmailCampaignResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique campaign ID. | 
**Status** | **string** | Current campaign status. | 
**EmailSenderId** | Pointer to **NullableInt32** | Email sender ID used for this campaign. | [optional] 
**StartAt** | **time.Time** | Campaign start timestamp. | 
**CreatedBy** | [**UserPersonalInfo**](UserPersonalInfo.md) |  | 
**FromName** | Pointer to **NullableString** | Sender name displayed in recipient&#39;s inbox. | [optional] 
**FromEmail** | **string** | Sender email address. | 
**ReplyToEmail** | **string** | Reply-to email address. | 
**Subject** | **string** | Email subject line. | 
**Html** | **string** | HTML email content. | 
**Cost** | **float32** | Total campaign cost. | 
**Totals** | [**EmailCampaignStatisticTotals**](EmailCampaignStatisticTotals.md) |  | 
**OutboundEmail** | Pointer to [**OutboundEmailResponse**](OutboundEmailResponse.md) |  | [optional] 
**FailedReason** | Pointer to **NullableString** | Reason for campaign failure if applicable. | [optional] 

## Methods

### NewCreateEmailCampaignResponse

`func NewCreateEmailCampaignResponse(id int32, status string, startAt time.Time, createdBy UserPersonalInfo, fromEmail string, replyToEmail string, subject string, html string, cost float32, totals EmailCampaignStatisticTotals, ) *CreateEmailCampaignResponse`

NewCreateEmailCampaignResponse instantiates a new CreateEmailCampaignResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEmailCampaignResponseWithDefaults

`func NewCreateEmailCampaignResponseWithDefaults() *CreateEmailCampaignResponse`

NewCreateEmailCampaignResponseWithDefaults instantiates a new CreateEmailCampaignResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateEmailCampaignResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateEmailCampaignResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateEmailCampaignResponse) SetId(v int32)`

SetId sets Id field to given value.


### GetStatus

`func (o *CreateEmailCampaignResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateEmailCampaignResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateEmailCampaignResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetEmailSenderId

`func (o *CreateEmailCampaignResponse) GetEmailSenderId() int32`

GetEmailSenderId returns the EmailSenderId field if non-nil, zero value otherwise.

### GetEmailSenderIdOk

`func (o *CreateEmailCampaignResponse) GetEmailSenderIdOk() (*int32, bool)`

GetEmailSenderIdOk returns a tuple with the EmailSenderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailSenderId

`func (o *CreateEmailCampaignResponse) SetEmailSenderId(v int32)`

SetEmailSenderId sets EmailSenderId field to given value.

### HasEmailSenderId

`func (o *CreateEmailCampaignResponse) HasEmailSenderId() bool`

HasEmailSenderId returns a boolean if a field has been set.

### SetEmailSenderIdNil

`func (o *CreateEmailCampaignResponse) SetEmailSenderIdNil(b bool)`

 SetEmailSenderIdNil sets the value for EmailSenderId to be an explicit nil

### UnsetEmailSenderId
`func (o *CreateEmailCampaignResponse) UnsetEmailSenderId()`

UnsetEmailSenderId ensures that no value is present for EmailSenderId, not even an explicit nil
### GetStartAt

`func (o *CreateEmailCampaignResponse) GetStartAt() time.Time`

GetStartAt returns the StartAt field if non-nil, zero value otherwise.

### GetStartAtOk

`func (o *CreateEmailCampaignResponse) GetStartAtOk() (*time.Time, bool)`

GetStartAtOk returns a tuple with the StartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAt

`func (o *CreateEmailCampaignResponse) SetStartAt(v time.Time)`

SetStartAt sets StartAt field to given value.


### GetCreatedBy

`func (o *CreateEmailCampaignResponse) GetCreatedBy() UserPersonalInfo`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *CreateEmailCampaignResponse) GetCreatedByOk() (*UserPersonalInfo, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *CreateEmailCampaignResponse) SetCreatedBy(v UserPersonalInfo)`

SetCreatedBy sets CreatedBy field to given value.


### GetFromName

`func (o *CreateEmailCampaignResponse) GetFromName() string`

GetFromName returns the FromName field if non-nil, zero value otherwise.

### GetFromNameOk

`func (o *CreateEmailCampaignResponse) GetFromNameOk() (*string, bool)`

GetFromNameOk returns a tuple with the FromName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromName

`func (o *CreateEmailCampaignResponse) SetFromName(v string)`

SetFromName sets FromName field to given value.

### HasFromName

`func (o *CreateEmailCampaignResponse) HasFromName() bool`

HasFromName returns a boolean if a field has been set.

### SetFromNameNil

`func (o *CreateEmailCampaignResponse) SetFromNameNil(b bool)`

 SetFromNameNil sets the value for FromName to be an explicit nil

### UnsetFromName
`func (o *CreateEmailCampaignResponse) UnsetFromName()`

UnsetFromName ensures that no value is present for FromName, not even an explicit nil
### GetFromEmail

`func (o *CreateEmailCampaignResponse) GetFromEmail() string`

GetFromEmail returns the FromEmail field if non-nil, zero value otherwise.

### GetFromEmailOk

`func (o *CreateEmailCampaignResponse) GetFromEmailOk() (*string, bool)`

GetFromEmailOk returns a tuple with the FromEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromEmail

`func (o *CreateEmailCampaignResponse) SetFromEmail(v string)`

SetFromEmail sets FromEmail field to given value.


### GetReplyToEmail

`func (o *CreateEmailCampaignResponse) GetReplyToEmail() string`

GetReplyToEmail returns the ReplyToEmail field if non-nil, zero value otherwise.

### GetReplyToEmailOk

`func (o *CreateEmailCampaignResponse) GetReplyToEmailOk() (*string, bool)`

GetReplyToEmailOk returns a tuple with the ReplyToEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyToEmail

`func (o *CreateEmailCampaignResponse) SetReplyToEmail(v string)`

SetReplyToEmail sets ReplyToEmail field to given value.


### GetSubject

`func (o *CreateEmailCampaignResponse) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CreateEmailCampaignResponse) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CreateEmailCampaignResponse) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetHtml

`func (o *CreateEmailCampaignResponse) GetHtml() string`

GetHtml returns the Html field if non-nil, zero value otherwise.

### GetHtmlOk

`func (o *CreateEmailCampaignResponse) GetHtmlOk() (*string, bool)`

GetHtmlOk returns a tuple with the Html field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtml

`func (o *CreateEmailCampaignResponse) SetHtml(v string)`

SetHtml sets Html field to given value.


### GetCost

`func (o *CreateEmailCampaignResponse) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *CreateEmailCampaignResponse) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *CreateEmailCampaignResponse) SetCost(v float32)`

SetCost sets Cost field to given value.


### GetTotals

`func (o *CreateEmailCampaignResponse) GetTotals() EmailCampaignStatisticTotals`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *CreateEmailCampaignResponse) GetTotalsOk() (*EmailCampaignStatisticTotals, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *CreateEmailCampaignResponse) SetTotals(v EmailCampaignStatisticTotals)`

SetTotals sets Totals field to given value.


### GetOutboundEmail

`func (o *CreateEmailCampaignResponse) GetOutboundEmail() OutboundEmailResponse`

GetOutboundEmail returns the OutboundEmail field if non-nil, zero value otherwise.

### GetOutboundEmailOk

`func (o *CreateEmailCampaignResponse) GetOutboundEmailOk() (*OutboundEmailResponse, bool)`

GetOutboundEmailOk returns a tuple with the OutboundEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundEmail

`func (o *CreateEmailCampaignResponse) SetOutboundEmail(v OutboundEmailResponse)`

SetOutboundEmail sets OutboundEmail field to given value.

### HasOutboundEmail

`func (o *CreateEmailCampaignResponse) HasOutboundEmail() bool`

HasOutboundEmail returns a boolean if a field has been set.

### GetFailedReason

`func (o *CreateEmailCampaignResponse) GetFailedReason() string`

GetFailedReason returns the FailedReason field if non-nil, zero value otherwise.

### GetFailedReasonOk

`func (o *CreateEmailCampaignResponse) GetFailedReasonOk() (*string, bool)`

GetFailedReasonOk returns a tuple with the FailedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedReason

`func (o *CreateEmailCampaignResponse) SetFailedReason(v string)`

SetFailedReason sets FailedReason field to given value.

### HasFailedReason

`func (o *CreateEmailCampaignResponse) HasFailedReason() bool`

HasFailedReason returns a boolean if a field has been set.

### SetFailedReasonNil

`func (o *CreateEmailCampaignResponse) SetFailedReasonNil(b bool)`

 SetFailedReasonNil sets the value for FailedReason to be an explicit nil

### UnsetFailedReason
`func (o *CreateEmailCampaignResponse) UnsetFailedReason()`

UnsetFailedReason ensures that no value is present for FailedReason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


