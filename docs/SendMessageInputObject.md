# SendMessageInputObject

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | **string** | Message text. Required if **template_id** is not set. | [default to null]
**TemplateId** | **int32** | Template used instead of message text. Required if **text** is not set. | [optional] [default to null]
**SendingTime** | **int32** | DEPRECATED, consider using sendingDateTime and sendingTimezone parameters instead: Optional (required with rrule set). Message sending time in unix timestamp format. Default is now. | [optional] [default to null]
**SendingDateTime** | **string** | Sending time in Y-m-d H:i:s format (e.g. 2016-05-27 13:02:33). This time is relative to **sendingTimezone**. | [optional] [default to null]
**SendingTimezone** | **string** | ID or ISO-name of timezone used for sending when sendingDateTime parameter is set. E.g. if you specify sendingDateTime &#x3D; \\\&quot;2016-05-27 13:02:33\\\&quot; and sendingTimezone &#x3D; \\\&quot;America/Buenos_Aires\\\&quot;, your message will be sent at May 27, 2016 13:02:33 Buenos Aires time, or 16:02:33 UTC. Default is account timezone. | [optional] [default to null]
**Contacts** | **string** | Comma separated array of contact resources id message will be sent to. | [optional] [default to null]
**Lists** | **string** | Comma separated array of list resources id message will be sent to. | [optional] [default to null]
**Phones** | **string** | Comma separated array of E.164 phone numbers message will be sent to. | [default to null]
**CutExtra** | **bool** | Should sending method cut extra characters which not fit supplied partsCount or return 400 Bad request response instead. | [optional] [default to null]
**PartsCount** | **int32** | Maximum message parts count (TextMagic allows sending 1 to 6 message parts). | [optional] [default to null]
**ReferenceId** | **int32** | Custom message reference id which can be used in your application infrastructure. | [optional] [default to null]
**From** | **string** | One of allowed Sender ID (phone number or alphanumeric sender ID). If specified Sender ID is not allowed for some destinations, a fallback default Sender ID will be used to ensure delivery. See [Get timezones](http://docs.textmagictesting.com/#tag/Sender-IDs). | [optional] [default to null]
**Rrule** | **string** | iCal RRULE parameter to create recurrent scheduled messages. When used, sendingTime is mandatory as start point of sending. See https://www.textmagic.com/free-tools/rrule-generator for format details. | [optional] [default to null]
**CreateChat** | **bool** | Should sending method try to create new Chat(if not exist) with specified recipients. | [optional] [default to null]
**Tts** | **bool** | Send Text to Speech message. | [optional] [default to null]
**Local** | **bool** | Treat phone numbers passed in \\&#39;phones\\&#39; field as local. | [optional] [default to null]
**LocalCountry** | **string** | 2-letter ISO country code for local phone numbers, used when \\&#39;local\\&#39; is set to true. Default is account country. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


