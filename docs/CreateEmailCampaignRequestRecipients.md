# CreateEmailCampaignRequestRecipients

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContactIds** | **[]int32** | Array of contact IDs to send to. | 
**Emails** | **[]string** | Array of email addresses to send to. | 
**GroupIds** | **[]int32** | Array of group IDs to send to. | 

## Methods

### NewCreateEmailCampaignRequestRecipients

`func NewCreateEmailCampaignRequestRecipients(contactIds []int32, emails []string, groupIds []int32, ) *CreateEmailCampaignRequestRecipients`

NewCreateEmailCampaignRequestRecipients instantiates a new CreateEmailCampaignRequestRecipients object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEmailCampaignRequestRecipientsWithDefaults

`func NewCreateEmailCampaignRequestRecipientsWithDefaults() *CreateEmailCampaignRequestRecipients`

NewCreateEmailCampaignRequestRecipientsWithDefaults instantiates a new CreateEmailCampaignRequestRecipients object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContactIds

`func (o *CreateEmailCampaignRequestRecipients) GetContactIds() []int32`

GetContactIds returns the ContactIds field if non-nil, zero value otherwise.

### GetContactIdsOk

`func (o *CreateEmailCampaignRequestRecipients) GetContactIdsOk() (*[]int32, bool)`

GetContactIdsOk returns a tuple with the ContactIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactIds

`func (o *CreateEmailCampaignRequestRecipients) SetContactIds(v []int32)`

SetContactIds sets ContactIds field to given value.


### GetEmails

`func (o *CreateEmailCampaignRequestRecipients) GetEmails() []string`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *CreateEmailCampaignRequestRecipients) GetEmailsOk() (*[]string, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *CreateEmailCampaignRequestRecipients) SetEmails(v []string)`

SetEmails sets Emails field to given value.


### GetGroupIds

`func (o *CreateEmailCampaignRequestRecipients) GetGroupIds() []int32`

GetGroupIds returns the GroupIds field if non-nil, zero value otherwise.

### GetGroupIdsOk

`func (o *CreateEmailCampaignRequestRecipients) GetGroupIdsOk() (*[]int32, bool)`

GetGroupIdsOk returns a tuple with the GroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIds

`func (o *CreateEmailCampaignRequestRecipients) SetGroupIds(v []int32)`

SetGroupIds sets GroupIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


