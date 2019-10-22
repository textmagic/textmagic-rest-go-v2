# Chat

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Chat ID. | [default to null]
**OriginalId** | **int32** |  | [default to null]
**Phone** | **string** | Chat partner phone number. | [default to null]
**Contact** | [***Contact**](Contact.md) |  | [default to null]
**UnsubscribedContactId** | **int32** | If this field has a value then it means that chat phone number has been unsubscribed from you and this value is a ID of a Unsubscribed contact entity. See [Get all unsubscribed contacts](http://docs.textmagictesting.com/#operation/getUnsubscribers). | [default to null]
**Unread** | **int32** | Total unread incoming messages. | [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | Time when last incoming message arrived at this chat. | [default to null]
**Status** | **string** | Chat status:   * **a** - Active   * **c** - Closed   * **d** - Deleted  | [default to null]
**Mute** | **int32** | Indicates when chat is muted. | [default to null]
**LastMessage** | **string** | The last message content of a chat. | [default to null]
**Direction** | **string** | Last message type: * **ci** - incoming call * **co** - outgoing call * **i** - incoming message * **o** - outgoing message  | [default to null]
**From** | **string** | If filled then value will be used as a sender number for all outgoing messages of a chat. | [default to null]
**MutedUntil** | [**time.Time**](time.Time.md) | Date and time until chat will be mutted. | [default to null]
**TimeLeftMute** | **int32** | Time left till chat will be unmutted (seconds). | [default to null]
**Country** | [***Country**](Country.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


