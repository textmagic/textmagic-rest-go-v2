# SendMessageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | Pointer to **string** | Message text. Required if the **template_id** is not set. | [optional] 
**TemplateId** | Pointer to **int32** | Template used instead of message text. Required if the **text** is not set. | [optional] 
**SendingTime** | Pointer to **int32** | DEPRECATED, consider using sendingDateTime and sendingTimezone parameters instead: Optional (required with rrule set). Message sending time in unix timestamp format. Default is now. | [optional] 
**SendingDateTime** | Pointer to **string** | Sending time in Y-m-d H:i:s format (e.g. 2022-05-27 13:05:10). This time is relative to **sendingTimezone**. Note: for correct operation, the value of seconds must not be less than 10. | [optional] 
**SendingTimezone** | Pointer to **string** | ID or ISO-name of timezone used for sending when sendingDateTime parameter is set. E.g. if you specify sendingDateTime &#x3D; \\\&quot;2016-05-27 13:02:33\\\&quot; and sendingTimezone &#x3D; \\\&quot;America/Buenos_Aires\\\&quot;, your message will be sent at May 27, 2016 13:02:33 Buenos Aires time, or 16:02:33 UTC. Default is account timezone. | [optional] 
**Contacts** | Pointer to **string** | Comma separated array of contact resources id message will be sent to. | [optional] 
**Lists** | Pointer to **string** | Comma separated array of list resources id message will be sent to. | [optional] 
**Phones** | Pointer to **string** | Comma separated array of E.164 phone numbers message will be sent to. | [optional] 
**CutExtra** | Pointer to **bool** | Should sending method cut extra characters which not fit supplied partsCount or return 400 Bad request response instead. | [optional] [default to false]
**PartsCount** | Pointer to **int32** | Maximum message parts count (Textmagic allows sending 1 to 6 message parts). | [optional] 
**ReferenceId** | Pointer to **int32** | Custom message reference id which can be used in your application infrastructure. | [optional] 
**From** | Pointer to **string** | One of allowed Sender ID (phone number or alphanumeric sender ID). If specified Sender ID is not allowed for some destinations, a fallback default Sender ID will be used to ensure delivery. See [Get timezones](https://docs.textmagic.com/#tag/Sender-IDs). | [optional] 
**Rrule** | Pointer to **string** | iCal RRULE parameter to create recurrent scheduled messages. When used, sendingTime is mandatory as start point of sending. See https://www.textmagic.com/free-tools/rrule-generator for format details. | [optional] 
**CreateChat** | Pointer to **bool** | Should sending method try to create new Chat (if not exist) with specified recipients? | [optional] [default to false]
**Tts** | Pointer to **bool** | Send a Text-to-Speech message. | [optional] [default to false]
**Local** | Pointer to **bool** | Treat phone numbers passed in the \\&#39;phones\\&#39; field as local. | [optional] [default to false]
**LocalCountry** | Pointer to **string** | The 2-letter ISO country code for local phone numbers, used when \\&#39;local\\&#39; is set to true. Default is the account country. | [optional] 
**Destination** | Pointer to **string** | Messsage destination type allowed [mms, tts]. | [optional] 
**Resources** | Pointer to **string** | File name from mms attachment response (named as resource) | [optional] 

## Methods

### NewSendMessageRequest

`func NewSendMessageRequest() *SendMessageRequest`

NewSendMessageRequest instantiates a new SendMessageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendMessageRequestWithDefaults

`func NewSendMessageRequestWithDefaults() *SendMessageRequest`

NewSendMessageRequestWithDefaults instantiates a new SendMessageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *SendMessageRequest) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *SendMessageRequest) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *SendMessageRequest) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *SendMessageRequest) HasText() bool`

HasText returns a boolean if a field has been set.

### GetTemplateId

`func (o *SendMessageRequest) GetTemplateId() int32`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *SendMessageRequest) GetTemplateIdOk() (*int32, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *SendMessageRequest) SetTemplateId(v int32)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *SendMessageRequest) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetSendingTime

`func (o *SendMessageRequest) GetSendingTime() int32`

GetSendingTime returns the SendingTime field if non-nil, zero value otherwise.

### GetSendingTimeOk

`func (o *SendMessageRequest) GetSendingTimeOk() (*int32, bool)`

GetSendingTimeOk returns a tuple with the SendingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendingTime

`func (o *SendMessageRequest) SetSendingTime(v int32)`

SetSendingTime sets SendingTime field to given value.

### HasSendingTime

`func (o *SendMessageRequest) HasSendingTime() bool`

HasSendingTime returns a boolean if a field has been set.

### GetSendingDateTime

`func (o *SendMessageRequest) GetSendingDateTime() string`

GetSendingDateTime returns the SendingDateTime field if non-nil, zero value otherwise.

### GetSendingDateTimeOk

`func (o *SendMessageRequest) GetSendingDateTimeOk() (*string, bool)`

GetSendingDateTimeOk returns a tuple with the SendingDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendingDateTime

`func (o *SendMessageRequest) SetSendingDateTime(v string)`

SetSendingDateTime sets SendingDateTime field to given value.

### HasSendingDateTime

`func (o *SendMessageRequest) HasSendingDateTime() bool`

HasSendingDateTime returns a boolean if a field has been set.

### GetSendingTimezone

`func (o *SendMessageRequest) GetSendingTimezone() string`

GetSendingTimezone returns the SendingTimezone field if non-nil, zero value otherwise.

### GetSendingTimezoneOk

`func (o *SendMessageRequest) GetSendingTimezoneOk() (*string, bool)`

GetSendingTimezoneOk returns a tuple with the SendingTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendingTimezone

`func (o *SendMessageRequest) SetSendingTimezone(v string)`

SetSendingTimezone sets SendingTimezone field to given value.

### HasSendingTimezone

`func (o *SendMessageRequest) HasSendingTimezone() bool`

HasSendingTimezone returns a boolean if a field has been set.

### GetContacts

`func (o *SendMessageRequest) GetContacts() string`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *SendMessageRequest) GetContactsOk() (*string, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *SendMessageRequest) SetContacts(v string)`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *SendMessageRequest) HasContacts() bool`

HasContacts returns a boolean if a field has been set.

### GetLists

`func (o *SendMessageRequest) GetLists() string`

GetLists returns the Lists field if non-nil, zero value otherwise.

### GetListsOk

`func (o *SendMessageRequest) GetListsOk() (*string, bool)`

GetListsOk returns a tuple with the Lists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLists

`func (o *SendMessageRequest) SetLists(v string)`

SetLists sets Lists field to given value.

### HasLists

`func (o *SendMessageRequest) HasLists() bool`

HasLists returns a boolean if a field has been set.

### GetPhones

`func (o *SendMessageRequest) GetPhones() string`

GetPhones returns the Phones field if non-nil, zero value otherwise.

### GetPhonesOk

`func (o *SendMessageRequest) GetPhonesOk() (*string, bool)`

GetPhonesOk returns a tuple with the Phones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhones

`func (o *SendMessageRequest) SetPhones(v string)`

SetPhones sets Phones field to given value.

### HasPhones

`func (o *SendMessageRequest) HasPhones() bool`

HasPhones returns a boolean if a field has been set.

### GetCutExtra

`func (o *SendMessageRequest) GetCutExtra() bool`

GetCutExtra returns the CutExtra field if non-nil, zero value otherwise.

### GetCutExtraOk

`func (o *SendMessageRequest) GetCutExtraOk() (*bool, bool)`

GetCutExtraOk returns a tuple with the CutExtra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCutExtra

`func (o *SendMessageRequest) SetCutExtra(v bool)`

SetCutExtra sets CutExtra field to given value.

### HasCutExtra

`func (o *SendMessageRequest) HasCutExtra() bool`

HasCutExtra returns a boolean if a field has been set.

### GetPartsCount

`func (o *SendMessageRequest) GetPartsCount() int32`

GetPartsCount returns the PartsCount field if non-nil, zero value otherwise.

### GetPartsCountOk

`func (o *SendMessageRequest) GetPartsCountOk() (*int32, bool)`

GetPartsCountOk returns a tuple with the PartsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartsCount

`func (o *SendMessageRequest) SetPartsCount(v int32)`

SetPartsCount sets PartsCount field to given value.

### HasPartsCount

`func (o *SendMessageRequest) HasPartsCount() bool`

HasPartsCount returns a boolean if a field has been set.

### GetReferenceId

`func (o *SendMessageRequest) GetReferenceId() int32`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *SendMessageRequest) GetReferenceIdOk() (*int32, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *SendMessageRequest) SetReferenceId(v int32)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *SendMessageRequest) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetFrom

`func (o *SendMessageRequest) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *SendMessageRequest) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *SendMessageRequest) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *SendMessageRequest) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetRrule

`func (o *SendMessageRequest) GetRrule() string`

GetRrule returns the Rrule field if non-nil, zero value otherwise.

### GetRruleOk

`func (o *SendMessageRequest) GetRruleOk() (*string, bool)`

GetRruleOk returns a tuple with the Rrule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrule

`func (o *SendMessageRequest) SetRrule(v string)`

SetRrule sets Rrule field to given value.

### HasRrule

`func (o *SendMessageRequest) HasRrule() bool`

HasRrule returns a boolean if a field has been set.

### GetCreateChat

`func (o *SendMessageRequest) GetCreateChat() bool`

GetCreateChat returns the CreateChat field if non-nil, zero value otherwise.

### GetCreateChatOk

`func (o *SendMessageRequest) GetCreateChatOk() (*bool, bool)`

GetCreateChatOk returns a tuple with the CreateChat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateChat

`func (o *SendMessageRequest) SetCreateChat(v bool)`

SetCreateChat sets CreateChat field to given value.

### HasCreateChat

`func (o *SendMessageRequest) HasCreateChat() bool`

HasCreateChat returns a boolean if a field has been set.

### GetTts

`func (o *SendMessageRequest) GetTts() bool`

GetTts returns the Tts field if non-nil, zero value otherwise.

### GetTtsOk

`func (o *SendMessageRequest) GetTtsOk() (*bool, bool)`

GetTtsOk returns a tuple with the Tts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTts

`func (o *SendMessageRequest) SetTts(v bool)`

SetTts sets Tts field to given value.

### HasTts

`func (o *SendMessageRequest) HasTts() bool`

HasTts returns a boolean if a field has been set.

### GetLocal

`func (o *SendMessageRequest) GetLocal() bool`

GetLocal returns the Local field if non-nil, zero value otherwise.

### GetLocalOk

`func (o *SendMessageRequest) GetLocalOk() (*bool, bool)`

GetLocalOk returns a tuple with the Local field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocal

`func (o *SendMessageRequest) SetLocal(v bool)`

SetLocal sets Local field to given value.

### HasLocal

`func (o *SendMessageRequest) HasLocal() bool`

HasLocal returns a boolean if a field has been set.

### GetLocalCountry

`func (o *SendMessageRequest) GetLocalCountry() string`

GetLocalCountry returns the LocalCountry field if non-nil, zero value otherwise.

### GetLocalCountryOk

`func (o *SendMessageRequest) GetLocalCountryOk() (*string, bool)`

GetLocalCountryOk returns a tuple with the LocalCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalCountry

`func (o *SendMessageRequest) SetLocalCountry(v string)`

SetLocalCountry sets LocalCountry field to given value.

### HasLocalCountry

`func (o *SendMessageRequest) HasLocalCountry() bool`

HasLocalCountry returns a boolean if a field has been set.

### GetDestination

`func (o *SendMessageRequest) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *SendMessageRequest) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *SendMessageRequest) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *SendMessageRequest) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetResources

`func (o *SendMessageRequest) GetResources() string`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *SendMessageRequest) GetResourcesOk() (*string, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *SendMessageRequest) SetResources(v string)`

SetResources sets Resources field to given value.

### HasResources

`func (o *SendMessageRequest) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


