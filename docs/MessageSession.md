# MessageSession

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Session ID. | [default to null]
**StartTime** | **string** | Session creation time. | [default to null]
**Text** | **string** | Session text. If a template was used for the session text (see [Messages: Send](http://docs.textmagictesting.com/#tag/Outbound-Messages) for details), it may contain template tags.  | [default to null]
**Source** | **string** | *   **O** for TextMagic Online *   **A** for API *   **M** for TextMagic Messenger *   **E** for [Email to SMS](http://docs.textmagictesting.com/#tag/Send-Email-to-SMS) *   **X** for [Distribution lists](http://docs.textmagictesting.com/#tag/Distribution-Lists)  | [default to null]
**ReferenceId** | **string** | Custom reference ID (see [Messages: Send](http://docs.textmagictesting.com/#tag/Send-Email-to-SMS) for details).  | [default to null]
**Price** | **float32** | Session cost (in account currency). | [default to null]
**NumbersCount** | **int32** | Session recipient count. | [default to null]
**Destination** | **string** | Destination type of a Message Session: * **t** - text SMS * **s** - text to speech * **v** - voice broadcast  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


