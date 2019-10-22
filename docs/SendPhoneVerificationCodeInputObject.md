# SendPhoneVerificationCodeInputObject

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phone** | **string** | Use the phone number in international E.164 format. If you need to pass phone numbers in the local format, please use them with the **country** parameter to specify the origin country of the phone number.  | [default to null]
**Brand** | **string** | Alphanumeric string with up to 18 characters you can use to personalize the verification text message body, to help users identify your company or application name. For example: “Your TextMagic PIN is …”  | [default to null]
**CodeLength** | **int32** | The length of the verification code. The value can be 4 or 6 characters.  | [default to null]
**Language** | **string** | By default, the SMS or text-to-speech (TTS) voice message is generated in the locale that matches the number. For example, the text message or TTS message for a 33\\* number is sent in French. Use this parameter to explicitly control the language, accent, and gender used for the verification request.Can be one of the following: &#x60;de-de&#x60;, &#x60;en-au&#x60;, &#x60;en-gb&#x60;, &#x60;en-us&#x60;, &#x60;en-in&#x60;, &#x60;es-es&#x60;, &#x60;es-mx&#x60;, &#x60;es-us&#x60;, &#x60;fr-ca&#x60;, &#x60;fr-fr&#x60;, &#x60;is-is&#x60;, &#x60;it-it&#x60;, &#x60;ja-jp&#x60;, &#x60;ko-kr&#x60;, &#x60;nl-nl&#x60;, &#x60;pl-pl&#x60;, &#x60;pt-pt&#x60;, &#x60;pt-br&#x60;, &#x60;ro-ro&#x60;, &#x60;ru-ru&#x60;, &#x60;sv-se&#x60;, &#x60;tr-tr&#x60;, &#x60;zh-cn&#x60; or &#x60;zh-tw&#x60;.  | [optional] [default to null]
**SenderId** | **string** | One of the available [sender settings](https://my.textmagic.com/online/reply-options/) on your TextMagic account. If specified sender setting type is not allowed for some destinations, a fallback default sender will be used to ensure the message delivery. More info about known restrictions can be found [here](https://support.textmagic.com/article/how-to-understand-sender-setting-restrictions/).  | [optional] [default to null]
**Country** | **string** | 2-letter ISO country code for local phone number. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


