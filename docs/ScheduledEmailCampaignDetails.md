# ScheduledEmailCampaignDetails

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique scheduled campaign ID. | [default to null]
**Status** | **string** | Current scheduled campaign status. | [default to null]
**EmailSenderId** | **int32** | Email sender ID used for this campaign. | [optional] [default to null]
**StartAt** | **string** | Scheduled start timestamp (UTC). | [default to null]
**EndAt** | **string** | End timestamp for recurring campaigns (UTC). | [optional] [default to null]
**NextSendAt** | **string** | Next scheduled send timestamp (UTC). | [optional] [default to null]
**CreatedBy** | [***UserPersonalInfo**](UserPersonalInfo.md) |  | [default to null]
**CreatedAt** | **string** | Campaign creation timestamp. | [default to null]
**UpdatedAt** | **string** | Last update timestamp. | [default to null]
**Type_** | **string** | Campaign recurrence type. | [default to null]
**FromName** | **string** | Sender name displayed in recipient&#39;s inbox. | [optional] [default to null]
**FromEmail** | **string** | Sender email address. | [default to null]
**ReplyToEmail** | **string** | Reply-to email address. | [default to null]
**Subject** | **string** | Email subject line. | [default to null]
**Html** | **string** | HTML email content. | [default to null]
**RecipientsCount** | **int32** | Number of recipients for this campaign. | [default to null]
**SendingTimezone** | **string** | Timezone for sending the campaign. | [default to null]
**Rrule** | **string** | RFC 5545 recurrence rule for recurring campaigns. | [optional] [default to null]
**OccurrenceSummary** | **string** | Human-readable schedule description. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


