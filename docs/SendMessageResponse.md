# SendMessageResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Message ID. | [default to null]
**Href** | **string** | URI of the message session. | [default to null]
**Type_** | **string** | Message response type: * **message** – when the message is sent to a single recipient. * **session** – when the message is sent is to multiple recipients. * **schedule** - when the message is scheduled for sending. * **bulk** - when the message is sent to multiple recipients and the number of recipients requires asynchronous processing See [Sending more than 1,000 messages in one session](https://docs.textmagic.com/#section/Tutorials/Sending-more-than-1000-messages-in-one-session).  | [default to null]
**SessionId** | **int32** | Message session ID. | [default to null]
**BulkId** | **int32** | Bulk Session ID. See [Sending more than 1,000 messages in one session](https://docs.textmagic.com/#section/Tutorials/Sending-more-than-1000-messages-in-one-session). | [default to null]
**MessageId** | **int32** | Message ID. | [default to null]
**ScheduleId** | **int32** | Message Schedule ID. | [default to null]
**ChatId** | **int32** | Message Chat ID. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


