# MessagingStatItem

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReplyRate** | **float32** | The number of incoming messages divided by the number of total messages. | [default to null]
**Date** | [**time.Time**](time.Time.md) | Time interval start, empty if the **by** parameter was set to **off**.  | [default to null]
**DeliveryRate** | **float32** | Message delivery rate:the number of delivered messages divided by the number of total messages. | [default to null]
**Costs** | **float32** | Cost for sent messages during this period. The costs are in the [Account](http://docs.textmagictesting.com/#tag/User) currency.  | [default to null]
**MessagesReceived** | **int32** | Total received messages count. | [default to null]
**MessagesSentDelivered** | **int32** | Delivered messages count. As messages are retried for up to 48 hours, this value could change. | [default to null]
**MessagesSentAccepted** | **int32** | Messages accepted for delivery (in queue), but not yet delivered. | [default to null]
**MessagesSentBuffered** | **int32** | Messages buffered by endpoint cell phone operators. | [default to null]
**MessagesSentFailed** | **int32** | Messages that have failed for whatever reason, e.g. the destination phone was switched off for 48 hours or the recipient phone account is out of service. | [default to null]
**MessagesSentRejected** | **int32** | Messages that were rejected: invalid Sender ID used (e.g. you cannot use the Sender ID or your own mobile number when sending to the United States and Canada.)  | [default to null]
**MessagesSentParts** | **int32** | Total sent messages **parts** count. Note that this is not equal to the sent messages count, because one message could consist of 1 to 6 parts and users are charged per part, not per message. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


