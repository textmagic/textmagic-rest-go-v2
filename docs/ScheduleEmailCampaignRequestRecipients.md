# ScheduleEmailCampaignRequestRecipients

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContactIds** | **[]int32** | Array of contact IDs to send to. | 
**Emails** | **[]string** | Array of email addresses to send to. | 
**GroupIds** | **[]int32** | Array of group IDs to send to. | 

## Methods

### NewScheduleEmailCampaignRequestRecipients

`func NewScheduleEmailCampaignRequestRecipients(contactIds []int32, emails []string, groupIds []int32, ) *ScheduleEmailCampaignRequestRecipients`

NewScheduleEmailCampaignRequestRecipients instantiates a new ScheduleEmailCampaignRequestRecipients object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleEmailCampaignRequestRecipientsWithDefaults

`func NewScheduleEmailCampaignRequestRecipientsWithDefaults() *ScheduleEmailCampaignRequestRecipients`

NewScheduleEmailCampaignRequestRecipientsWithDefaults instantiates a new ScheduleEmailCampaignRequestRecipients object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContactIds

`func (o *ScheduleEmailCampaignRequestRecipients) GetContactIds() []int32`

GetContactIds returns the ContactIds field if non-nil, zero value otherwise.

### GetContactIdsOk

`func (o *ScheduleEmailCampaignRequestRecipients) GetContactIdsOk() (*[]int32, bool)`

GetContactIdsOk returns a tuple with the ContactIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactIds

`func (o *ScheduleEmailCampaignRequestRecipients) SetContactIds(v []int32)`

SetContactIds sets ContactIds field to given value.


### GetEmails

`func (o *ScheduleEmailCampaignRequestRecipients) GetEmails() []string`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *ScheduleEmailCampaignRequestRecipients) GetEmailsOk() (*[]string, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *ScheduleEmailCampaignRequestRecipients) SetEmails(v []string)`

SetEmails sets Emails field to given value.


### GetGroupIds

`func (o *ScheduleEmailCampaignRequestRecipients) GetGroupIds() []int32`

GetGroupIds returns the GroupIds field if non-nil, zero value otherwise.

### GetGroupIdsOk

`func (o *ScheduleEmailCampaignRequestRecipients) GetGroupIdsOk() (*[]int32, bool)`

GetGroupIdsOk returns a tuple with the GroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIds

`func (o *ScheduleEmailCampaignRequestRecipients) SetGroupIds(v []int32)`

SetGroupIds sets GroupIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


