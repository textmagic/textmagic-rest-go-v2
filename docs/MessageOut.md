# MessageOut

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Message ID. | [default to null]
**Sender** | **string** | Message sender (phone number or alphanumeric Sender ID). | [optional] [default to null]
**Receiver** | **string** | Recipient phone number. | [optional] [default to null]
**Text** | **string** |  | [default to null]
**Status** | **string** | Delivery status of the message. See [message delivery statuses](http://docs.textmagictesting.com/#section/Delivery-status-codes) for details.  | [default to null]
**ContactId** | **int32** | Recipient contact ID. | [default to null]
**SessionId** | **int32** | Message Session ID of a Message. | [default to null]
**MessageTime** | [**time.Time**](time.Time.md) | Sending time. | [default to null]
**Avatar** | **string** |  | [default to null]
**Deleted** | **bool** | Indicates that message has been deleted. | [optional] [default to null]
**Charset** | **string** | Message charset. Could be: *   **ISO-8859-1** for plaintext SMS *   **UTF-16BE** for Unicode SMS  | [default to null]
**CharsetLabel** | **string** | Human-readable message charset label. Could be: *   **ISO-8859-1** for plaintext SMS *   **UTF-16BE** for Unicode SMS *   **Voice** for voice services (Text-to-Speech or Voice Broadcast) messages  | [default to null]
**FirstName** | **string** | Contact first name. Could be substituted from your [Contacts](http://docs.textmagictesting.com/#tag/Contacts) (even if you submitted phone number instead of contact ID).  | [default to null]
**LastName** | **string** | Contact last name. | [default to null]
**Country** | **string** | Two-letter ISO country code of the recipient phone number.  | [default to null]
**Phone** | **string** | Receipent phone number. | [optional] [default to null]
**Price** | **float32** | Message price. | [optional] [default to null]
**PartsCount** | **int32** | Message parts (multiples of 160 characters) count. | [default to null]
**FromEmail** | **string** | User email which this message came from. For Email2SMS and Distribution Lists messages it will be an original email address, in other cases it is an account email address. | [optional] [default to null]
**FromNumber** | **string** | Phone number which is used to send SMS. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


