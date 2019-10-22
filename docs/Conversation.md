# Conversation

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [default to null]
**Direction** | **string** | Message type: inbound or outbound.  | [default to null]
**Sender** | **string** | Sender phone number. | [default to null]
**MessageTime** | [**time.Time**](time.Time.md) | Time when message arrived at TextMagic. | [default to null]
**Text** | **string** | Message text. | [default to null]
**Receiver** | **string** | Receiver phone number. | [default to null]
**Status** | **string** | Message status (for chats outbound only). See [message delivery statuses](http://docs.textmagictesting.com/#section/Delivery-status-codes) for details. | [default to null]
**FirstName** | **string** | Contact first name. | [default to null]
**LastName** | **string** | Contact last name. | [default to null]
**SessionId** | **int32** | Session ID of a message. See [message sessions](http://docs.textmagictesting.com/#tag/Outbound-Message-Sessions) for details. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


