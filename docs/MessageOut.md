# MessageOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Message ID. | 
**Sender** | Pointer to **NullableString** | Message sender (phone number or alphanumeric Sender ID). | [optional] 
**Receiver** | Pointer to **string** | Recipient&#x60;s phone number. | [optional] 
**Text** | **string** |  | 
**Status** | **string** | Delivery status of the message. See [message delivery statuses](https://docs.textmagic.com/#section/Delivery-status-codes) for details.  | 
**RejectReason** | Pointer to **string** | Rejection reason. | [optional] 
**ContactId** | **NullableInt32** | Recipient contact ID. | 
**SessionId** | **NullableInt32** | Message Session ID of a message. | 
**MessageTime** | **time.Time** | Sending time. | 
**Avatar** | **NullableString** |  | 
**Deleted** | Pointer to **bool** | Indicates that the message has been deleted. | [optional] 
**Charset** | **NullableString** | Message charset. Could be: *   **ISO-8859-1** for plaintext SMS; *   **UTF-16BE** for Unicode SMS.  | 
**CharsetLabel** | **NullableString** | Human-readable message charset label. Could be: *   **ISO-8859-1** for plaintext SMS; *   **UTF-16BE** for Unicode SMS; *   **Voice** for voice services (Text-to-Speech or Voice Broadcast) messages.  | 
**FirstName** | **NullableString** | Contact first name. Could be substituted from your [Contacts](https://docs.textmagic.com/#tag/Contacts) (even if you submitted the phone number instead of the contact ID).  | 
**LastName** | **NullableString** | Contact last name. | 
**Country** | **NullableString** | The 2-letter ISO country code of the recipient&#39;s phone number.  | 
**Phone** | Pointer to **NullableString** | Receipent&#x60;s phone number. | [optional] 
**Price** | Pointer to **NullableFloat32** | Message price. | [optional] 
**PartsCount** | **NullableInt32** | Message parts (multiples of 160 characters) count. | 
**FromEmail** | Pointer to **NullableString** | The user email which this message came from. For Email2SMS and Distribution Lists the messages, it is an original email address - in other cases, it is an account email address. | [optional] 
**FromNumber** | Pointer to **NullableString** | The Phone number used to send the SMS. | [optional] 
**SenderSource** | Pointer to [**MessageOutSenderSource**](MessageOutSenderSource.md) |  | [optional] 
**Session** | Pointer to [**MessageOutSession**](MessageOutSession.md) |  | [optional] 

## Methods

### NewMessageOut

`func NewMessageOut(id int32, text string, status string, contactId NullableInt32, sessionId NullableInt32, messageTime time.Time, avatar NullableString, charset NullableString, charsetLabel NullableString, firstName NullableString, lastName NullableString, country NullableString, partsCount NullableInt32, ) *MessageOut`

NewMessageOut instantiates a new MessageOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageOutWithDefaults

`func NewMessageOutWithDefaults() *MessageOut`

NewMessageOutWithDefaults instantiates a new MessageOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MessageOut) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MessageOut) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MessageOut) SetId(v int32)`

SetId sets Id field to given value.


### GetSender

`func (o *MessageOut) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *MessageOut) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *MessageOut) SetSender(v string)`

SetSender sets Sender field to given value.

### HasSender

`func (o *MessageOut) HasSender() bool`

HasSender returns a boolean if a field has been set.

### SetSenderNil

`func (o *MessageOut) SetSenderNil(b bool)`

 SetSenderNil sets the value for Sender to be an explicit nil

### UnsetSender
`func (o *MessageOut) UnsetSender()`

UnsetSender ensures that no value is present for Sender, not even an explicit nil
### GetReceiver

`func (o *MessageOut) GetReceiver() string`

GetReceiver returns the Receiver field if non-nil, zero value otherwise.

### GetReceiverOk

`func (o *MessageOut) GetReceiverOk() (*string, bool)`

GetReceiverOk returns a tuple with the Receiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiver

`func (o *MessageOut) SetReceiver(v string)`

SetReceiver sets Receiver field to given value.

### HasReceiver

`func (o *MessageOut) HasReceiver() bool`

HasReceiver returns a boolean if a field has been set.

### GetText

`func (o *MessageOut) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *MessageOut) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *MessageOut) SetText(v string)`

SetText sets Text field to given value.


### GetStatus

`func (o *MessageOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MessageOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MessageOut) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetRejectReason

`func (o *MessageOut) GetRejectReason() string`

GetRejectReason returns the RejectReason field if non-nil, zero value otherwise.

### GetRejectReasonOk

`func (o *MessageOut) GetRejectReasonOk() (*string, bool)`

GetRejectReasonOk returns a tuple with the RejectReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectReason

`func (o *MessageOut) SetRejectReason(v string)`

SetRejectReason sets RejectReason field to given value.

### HasRejectReason

`func (o *MessageOut) HasRejectReason() bool`

HasRejectReason returns a boolean if a field has been set.

### GetContactId

`func (o *MessageOut) GetContactId() int32`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *MessageOut) GetContactIdOk() (*int32, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *MessageOut) SetContactId(v int32)`

SetContactId sets ContactId field to given value.


### SetContactIdNil

`func (o *MessageOut) SetContactIdNil(b bool)`

 SetContactIdNil sets the value for ContactId to be an explicit nil

### UnsetContactId
`func (o *MessageOut) UnsetContactId()`

UnsetContactId ensures that no value is present for ContactId, not even an explicit nil
### GetSessionId

`func (o *MessageOut) GetSessionId() int32`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *MessageOut) GetSessionIdOk() (*int32, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *MessageOut) SetSessionId(v int32)`

SetSessionId sets SessionId field to given value.


### SetSessionIdNil

`func (o *MessageOut) SetSessionIdNil(b bool)`

 SetSessionIdNil sets the value for SessionId to be an explicit nil

### UnsetSessionId
`func (o *MessageOut) UnsetSessionId()`

UnsetSessionId ensures that no value is present for SessionId, not even an explicit nil
### GetMessageTime

`func (o *MessageOut) GetMessageTime() time.Time`

GetMessageTime returns the MessageTime field if non-nil, zero value otherwise.

### GetMessageTimeOk

`func (o *MessageOut) GetMessageTimeOk() (*time.Time, bool)`

GetMessageTimeOk returns a tuple with the MessageTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageTime

`func (o *MessageOut) SetMessageTime(v time.Time)`

SetMessageTime sets MessageTime field to given value.


### GetAvatar

`func (o *MessageOut) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *MessageOut) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *MessageOut) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.


### SetAvatarNil

`func (o *MessageOut) SetAvatarNil(b bool)`

 SetAvatarNil sets the value for Avatar to be an explicit nil

### UnsetAvatar
`func (o *MessageOut) UnsetAvatar()`

UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil
### GetDeleted

`func (o *MessageOut) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *MessageOut) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *MessageOut) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *MessageOut) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetCharset

`func (o *MessageOut) GetCharset() string`

GetCharset returns the Charset field if non-nil, zero value otherwise.

### GetCharsetOk

`func (o *MessageOut) GetCharsetOk() (*string, bool)`

GetCharsetOk returns a tuple with the Charset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharset

`func (o *MessageOut) SetCharset(v string)`

SetCharset sets Charset field to given value.


### SetCharsetNil

`func (o *MessageOut) SetCharsetNil(b bool)`

 SetCharsetNil sets the value for Charset to be an explicit nil

### UnsetCharset
`func (o *MessageOut) UnsetCharset()`

UnsetCharset ensures that no value is present for Charset, not even an explicit nil
### GetCharsetLabel

`func (o *MessageOut) GetCharsetLabel() string`

GetCharsetLabel returns the CharsetLabel field if non-nil, zero value otherwise.

### GetCharsetLabelOk

`func (o *MessageOut) GetCharsetLabelOk() (*string, bool)`

GetCharsetLabelOk returns a tuple with the CharsetLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharsetLabel

`func (o *MessageOut) SetCharsetLabel(v string)`

SetCharsetLabel sets CharsetLabel field to given value.


### SetCharsetLabelNil

`func (o *MessageOut) SetCharsetLabelNil(b bool)`

 SetCharsetLabelNil sets the value for CharsetLabel to be an explicit nil

### UnsetCharsetLabel
`func (o *MessageOut) UnsetCharsetLabel()`

UnsetCharsetLabel ensures that no value is present for CharsetLabel, not even an explicit nil
### GetFirstName

`func (o *MessageOut) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *MessageOut) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *MessageOut) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### SetFirstNameNil

`func (o *MessageOut) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *MessageOut) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *MessageOut) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *MessageOut) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *MessageOut) SetLastName(v string)`

SetLastName sets LastName field to given value.


### SetLastNameNil

`func (o *MessageOut) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *MessageOut) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetCountry

`func (o *MessageOut) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *MessageOut) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *MessageOut) SetCountry(v string)`

SetCountry sets Country field to given value.


### SetCountryNil

`func (o *MessageOut) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *MessageOut) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetPhone

`func (o *MessageOut) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *MessageOut) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *MessageOut) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *MessageOut) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *MessageOut) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *MessageOut) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetPrice

`func (o *MessageOut) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *MessageOut) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *MessageOut) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *MessageOut) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *MessageOut) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *MessageOut) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetPartsCount

`func (o *MessageOut) GetPartsCount() int32`

GetPartsCount returns the PartsCount field if non-nil, zero value otherwise.

### GetPartsCountOk

`func (o *MessageOut) GetPartsCountOk() (*int32, bool)`

GetPartsCountOk returns a tuple with the PartsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartsCount

`func (o *MessageOut) SetPartsCount(v int32)`

SetPartsCount sets PartsCount field to given value.


### SetPartsCountNil

`func (o *MessageOut) SetPartsCountNil(b bool)`

 SetPartsCountNil sets the value for PartsCount to be an explicit nil

### UnsetPartsCount
`func (o *MessageOut) UnsetPartsCount()`

UnsetPartsCount ensures that no value is present for PartsCount, not even an explicit nil
### GetFromEmail

`func (o *MessageOut) GetFromEmail() string`

GetFromEmail returns the FromEmail field if non-nil, zero value otherwise.

### GetFromEmailOk

`func (o *MessageOut) GetFromEmailOk() (*string, bool)`

GetFromEmailOk returns a tuple with the FromEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromEmail

`func (o *MessageOut) SetFromEmail(v string)`

SetFromEmail sets FromEmail field to given value.

### HasFromEmail

`func (o *MessageOut) HasFromEmail() bool`

HasFromEmail returns a boolean if a field has been set.

### SetFromEmailNil

`func (o *MessageOut) SetFromEmailNil(b bool)`

 SetFromEmailNil sets the value for FromEmail to be an explicit nil

### UnsetFromEmail
`func (o *MessageOut) UnsetFromEmail()`

UnsetFromEmail ensures that no value is present for FromEmail, not even an explicit nil
### GetFromNumber

`func (o *MessageOut) GetFromNumber() string`

GetFromNumber returns the FromNumber field if non-nil, zero value otherwise.

### GetFromNumberOk

`func (o *MessageOut) GetFromNumberOk() (*string, bool)`

GetFromNumberOk returns a tuple with the FromNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromNumber

`func (o *MessageOut) SetFromNumber(v string)`

SetFromNumber sets FromNumber field to given value.

### HasFromNumber

`func (o *MessageOut) HasFromNumber() bool`

HasFromNumber returns a boolean if a field has been set.

### SetFromNumberNil

`func (o *MessageOut) SetFromNumberNil(b bool)`

 SetFromNumberNil sets the value for FromNumber to be an explicit nil

### UnsetFromNumber
`func (o *MessageOut) UnsetFromNumber()`

UnsetFromNumber ensures that no value is present for FromNumber, not even an explicit nil
### GetSenderSource

`func (o *MessageOut) GetSenderSource() MessageOutSenderSource`

GetSenderSource returns the SenderSource field if non-nil, zero value otherwise.

### GetSenderSourceOk

`func (o *MessageOut) GetSenderSourceOk() (*MessageOutSenderSource, bool)`

GetSenderSourceOk returns a tuple with the SenderSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderSource

`func (o *MessageOut) SetSenderSource(v MessageOutSenderSource)`

SetSenderSource sets SenderSource field to given value.

### HasSenderSource

`func (o *MessageOut) HasSenderSource() bool`

HasSenderSource returns a boolean if a field has been set.

### GetSession

`func (o *MessageOut) GetSession() MessageOutSession`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *MessageOut) GetSessionOk() (*MessageOutSession, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *MessageOut) SetSession(v MessageOutSession)`

SetSession sets Session field to given value.

### HasSession

`func (o *MessageOut) HasSession() bool`

HasSession returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


