# Chat

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Chat ID. | [default to null]
**OriginalId** | **int32** |  | [default to null]
**Phone** | **string** | Chat partner&#39;s phone number. | [default to null]
**Contact** | [***Contact**](Contact.md) |  | [default to null]
**UnsubscribedContactId** | **int32** | If this field has a value, it means that the chat phone number has been unsubscribed from you and this value is an ID of an Unsubscribed contact entity. See [Get all unsubscribed contacts](https://docs.textmagic.com/#operation/getUnsubscribers). | [default to null]
**Unread** | **int32** | Total unread incoming messages. | [default to null]
**UpdatedAt** | **string** | Time when the last incoming message arrived at this chat. | [default to null]
**Status** | **string** | Chat status:   * **a** - Active;   * **c** - Closed;   * **d** - Deleted.  | [default to null]
**Mute** | **int32** | Indicates when the chat is muted. | [default to null]
**LastMessage** | **string** | The last message content of a chat. | [default to null]
**Direction** | **string** | Last message type: * **ci** - incoming call; * **co** - outgoing call; * **i** - incoming message; * **o** - outgoing message.  | [default to null]
**ReplyOptionsType** | **string** | Used for chats prices. | [default to null]
**From** | **string** | If filled, the value will be used as a sender number for all outgoing messages of a chat. | [default to null]
**MutedUntil** | **string** | Date and time until the chat will be muted. | [default to null]
**TimeLeftMute** | **int32** | Time left untill the chat will be unmuted (seconds). | [default to null]
**Country** | [***Country**](Country.md) |  | [default to null]
**Pinned** | **bool** | Indicates when the chat is pinned. | [default to null]
**Type_** | **string** | Chat type. | [default to null]
**SmsPrice** | **float32** |  | [default to null]
**MmsPrice** | **float32** |  | [default to null]
**Tags** | [**[]Tag**](Tag.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


