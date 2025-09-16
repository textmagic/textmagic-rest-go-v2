# CreateEmailCampaignResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique campaign ID. | [default to null]
**Status** | **string** | Current campaign status. | [default to null]
**EmailSenderId** | **int32** | Email sender ID used for this campaign. | [optional] [default to null]
**StartAt** | **string** | Campaign start timestamp. | [default to null]
**CreatedBy** | [***UserPersonalInfo**](UserPersonalInfo.md) |  | [default to null]
**FromName** | **string** | Sender name displayed in recipient&#39;s inbox. | [optional] [default to null]
**FromEmail** | **string** | Sender email address. | [default to null]
**ReplyToEmail** | **string** | Reply-to email address. | [default to null]
**Subject** | **string** | Email subject line. | [default to null]
**Html** | **string** | HTML email content. | [default to null]
**Cost** | **float32** | Total campaign cost. | [default to null]
**Totals** | [***EmailCampaignStatisticTotals**](EmailCampaignStatisticTotals.md) |  | [default to null]
**OutboundEmail** | [***OutboundEmailResponse**](OutboundEmailResponse.md) |  | [optional] [default to null]
**FailedReason** | **string** | Reason for campaign failure if applicable. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


