# SendMessageResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Message ID. | [default to null]
**Href** | **string** | URI of message session. | [default to null]
**Type_** | **string** | Message response type: * **message** when message sent to a single recipient * **session** when message sent to multiple recipients * **schedule** when message has been scheduled for sending * **bulk** when message sent to multiple recipient and the number of recipients requires asynchronous processiong See [Sending more than 1,000 messages in one session](http://docs.textmagictesting.com/#section/Tutorials/Sending-more-than-1000-messages-in-one-session).  | [default to null]
**SessionId** | **int32** | Message session ID. | [default to null]
**BulkId** | **int32** | Bulk Session ID. See [Sending more than 1,000 messages in one session](http://docs.textmagictesting.com/#section/Tutorials/Sending-more-than-1000-messages-in-one-session). | [default to null]
**MessageId** | **int32** | Message ID. | [default to null]
**ScheduleId** | **int32** | Message Schedule ID. | [default to null]
**ChatId** | **int32** | Message Chat ID. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


