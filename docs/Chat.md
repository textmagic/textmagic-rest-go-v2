# Chat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Chat ID. | 
**OriginalId** | **NullableInt32** |  | 
**Phone** | **string** | Chat partner&#39;s phone number. | 
**Contact** | [**NullableContact**](Contact.md) |  | 
**UnsubscribedContactId** | **NullableInt32** | If this field has a value, it means that the chat phone number has been unsubscribed from you and this value is an ID of an Unsubscribed contact entity. See [Get all unsubscribed contacts](https://docs.textmagic.com/#operation/getUnsubscribers). | 
**Unread** | **int32** | Total unread incoming messages. | 
**UpdatedAt** | **time.Time** | Time when the last incoming message arrived at this chat. | 
**Status** | **string** | Chat status:   * **a** - Active;   * **c** - Closed;   * **d** - Deleted.  | 
**Mute** | **int32** | Indicates when the chat is muted. | 
**LastMessage** | **NullableString** | The last message content of a chat. | 
**Direction** | **NullableString** | Last message type: * **ci** - incoming call; * **co** - outgoing call; * **i** - incoming message; * **o** - outgoing message.  | 
**ReplyOptionsType** | **NullableString** | Used for chats prices. | 
**From** | **NullableString** | If filled, the value will be used as a sender number for all outgoing messages of a chat. | 
**MutedUntil** | **NullableTime** | Date and time until the chat will be muted. | 
**TimeLeftMute** | **int32** | Time left untill the chat will be unmuted (seconds). | 
**Country** | [**NullableCountry**](Country.md) |  | 
**Pinned** | **NullableBool** | Indicates when the chat is pinned. | 
**Type** | **string** | Chat type. | 
**SmsPrice** | **float32** |  | 
**MmsPrice** | **float32** |  | 
**Tags** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 

## Methods

### NewChat

`func NewChat(id int32, originalId NullableInt32, phone string, contact NullableContact, unsubscribedContactId NullableInt32, unread int32, updatedAt time.Time, status string, mute int32, lastMessage NullableString, direction NullableString, replyOptionsType NullableString, from NullableString, mutedUntil NullableTime, timeLeftMute int32, country NullableCountry, pinned NullableBool, type_ string, smsPrice float32, mmsPrice float32, ) *Chat`

NewChat instantiates a new Chat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatWithDefaults

`func NewChatWithDefaults() *Chat`

NewChatWithDefaults instantiates a new Chat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Chat) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Chat) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Chat) SetId(v int32)`

SetId sets Id field to given value.


### GetOriginalId

`func (o *Chat) GetOriginalId() int32`

GetOriginalId returns the OriginalId field if non-nil, zero value otherwise.

### GetOriginalIdOk

`func (o *Chat) GetOriginalIdOk() (*int32, bool)`

GetOriginalIdOk returns a tuple with the OriginalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalId

`func (o *Chat) SetOriginalId(v int32)`

SetOriginalId sets OriginalId field to given value.


### SetOriginalIdNil

`func (o *Chat) SetOriginalIdNil(b bool)`

 SetOriginalIdNil sets the value for OriginalId to be an explicit nil

### UnsetOriginalId
`func (o *Chat) UnsetOriginalId()`

UnsetOriginalId ensures that no value is present for OriginalId, not even an explicit nil
### GetPhone

`func (o *Chat) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *Chat) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *Chat) SetPhone(v string)`

SetPhone sets Phone field to given value.


### GetContact

`func (o *Chat) GetContact() Contact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *Chat) GetContactOk() (*Contact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *Chat) SetContact(v Contact)`

SetContact sets Contact field to given value.


### SetContactNil

`func (o *Chat) SetContactNil(b bool)`

 SetContactNil sets the value for Contact to be an explicit nil

### UnsetContact
`func (o *Chat) UnsetContact()`

UnsetContact ensures that no value is present for Contact, not even an explicit nil
### GetUnsubscribedContactId

`func (o *Chat) GetUnsubscribedContactId() int32`

GetUnsubscribedContactId returns the UnsubscribedContactId field if non-nil, zero value otherwise.

### GetUnsubscribedContactIdOk

`func (o *Chat) GetUnsubscribedContactIdOk() (*int32, bool)`

GetUnsubscribedContactIdOk returns a tuple with the UnsubscribedContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsubscribedContactId

`func (o *Chat) SetUnsubscribedContactId(v int32)`

SetUnsubscribedContactId sets UnsubscribedContactId field to given value.


### SetUnsubscribedContactIdNil

`func (o *Chat) SetUnsubscribedContactIdNil(b bool)`

 SetUnsubscribedContactIdNil sets the value for UnsubscribedContactId to be an explicit nil

### UnsetUnsubscribedContactId
`func (o *Chat) UnsetUnsubscribedContactId()`

UnsetUnsubscribedContactId ensures that no value is present for UnsubscribedContactId, not even an explicit nil
### GetUnread

`func (o *Chat) GetUnread() int32`

GetUnread returns the Unread field if non-nil, zero value otherwise.

### GetUnreadOk

`func (o *Chat) GetUnreadOk() (*int32, bool)`

GetUnreadOk returns a tuple with the Unread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnread

`func (o *Chat) SetUnread(v int32)`

SetUnread sets Unread field to given value.


### GetUpdatedAt

`func (o *Chat) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Chat) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Chat) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetStatus

`func (o *Chat) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Chat) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Chat) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetMute

`func (o *Chat) GetMute() int32`

GetMute returns the Mute field if non-nil, zero value otherwise.

### GetMuteOk

`func (o *Chat) GetMuteOk() (*int32, bool)`

GetMuteOk returns a tuple with the Mute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMute

`func (o *Chat) SetMute(v int32)`

SetMute sets Mute field to given value.


### GetLastMessage

`func (o *Chat) GetLastMessage() string`

GetLastMessage returns the LastMessage field if non-nil, zero value otherwise.

### GetLastMessageOk

`func (o *Chat) GetLastMessageOk() (*string, bool)`

GetLastMessageOk returns a tuple with the LastMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMessage

`func (o *Chat) SetLastMessage(v string)`

SetLastMessage sets LastMessage field to given value.


### SetLastMessageNil

`func (o *Chat) SetLastMessageNil(b bool)`

 SetLastMessageNil sets the value for LastMessage to be an explicit nil

### UnsetLastMessage
`func (o *Chat) UnsetLastMessage()`

UnsetLastMessage ensures that no value is present for LastMessage, not even an explicit nil
### GetDirection

`func (o *Chat) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *Chat) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *Chat) SetDirection(v string)`

SetDirection sets Direction field to given value.


### SetDirectionNil

`func (o *Chat) SetDirectionNil(b bool)`

 SetDirectionNil sets the value for Direction to be an explicit nil

### UnsetDirection
`func (o *Chat) UnsetDirection()`

UnsetDirection ensures that no value is present for Direction, not even an explicit nil
### GetReplyOptionsType

`func (o *Chat) GetReplyOptionsType() string`

GetReplyOptionsType returns the ReplyOptionsType field if non-nil, zero value otherwise.

### GetReplyOptionsTypeOk

`func (o *Chat) GetReplyOptionsTypeOk() (*string, bool)`

GetReplyOptionsTypeOk returns a tuple with the ReplyOptionsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyOptionsType

`func (o *Chat) SetReplyOptionsType(v string)`

SetReplyOptionsType sets ReplyOptionsType field to given value.


### SetReplyOptionsTypeNil

`func (o *Chat) SetReplyOptionsTypeNil(b bool)`

 SetReplyOptionsTypeNil sets the value for ReplyOptionsType to be an explicit nil

### UnsetReplyOptionsType
`func (o *Chat) UnsetReplyOptionsType()`

UnsetReplyOptionsType ensures that no value is present for ReplyOptionsType, not even an explicit nil
### GetFrom

`func (o *Chat) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *Chat) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *Chat) SetFrom(v string)`

SetFrom sets From field to given value.


### SetFromNil

`func (o *Chat) SetFromNil(b bool)`

 SetFromNil sets the value for From to be an explicit nil

### UnsetFrom
`func (o *Chat) UnsetFrom()`

UnsetFrom ensures that no value is present for From, not even an explicit nil
### GetMutedUntil

`func (o *Chat) GetMutedUntil() time.Time`

GetMutedUntil returns the MutedUntil field if non-nil, zero value otherwise.

### GetMutedUntilOk

`func (o *Chat) GetMutedUntilOk() (*time.Time, bool)`

GetMutedUntilOk returns a tuple with the MutedUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutedUntil

`func (o *Chat) SetMutedUntil(v time.Time)`

SetMutedUntil sets MutedUntil field to given value.


### SetMutedUntilNil

`func (o *Chat) SetMutedUntilNil(b bool)`

 SetMutedUntilNil sets the value for MutedUntil to be an explicit nil

### UnsetMutedUntil
`func (o *Chat) UnsetMutedUntil()`

UnsetMutedUntil ensures that no value is present for MutedUntil, not even an explicit nil
### GetTimeLeftMute

`func (o *Chat) GetTimeLeftMute() int32`

GetTimeLeftMute returns the TimeLeftMute field if non-nil, zero value otherwise.

### GetTimeLeftMuteOk

`func (o *Chat) GetTimeLeftMuteOk() (*int32, bool)`

GetTimeLeftMuteOk returns a tuple with the TimeLeftMute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeLeftMute

`func (o *Chat) SetTimeLeftMute(v int32)`

SetTimeLeftMute sets TimeLeftMute field to given value.


### GetCountry

`func (o *Chat) GetCountry() Country`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Chat) GetCountryOk() (*Country, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Chat) SetCountry(v Country)`

SetCountry sets Country field to given value.


### SetCountryNil

`func (o *Chat) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *Chat) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetPinned

`func (o *Chat) GetPinned() bool`

GetPinned returns the Pinned field if non-nil, zero value otherwise.

### GetPinnedOk

`func (o *Chat) GetPinnedOk() (*bool, bool)`

GetPinnedOk returns a tuple with the Pinned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinned

`func (o *Chat) SetPinned(v bool)`

SetPinned sets Pinned field to given value.


### SetPinnedNil

`func (o *Chat) SetPinnedNil(b bool)`

 SetPinnedNil sets the value for Pinned to be an explicit nil

### UnsetPinned
`func (o *Chat) UnsetPinned()`

UnsetPinned ensures that no value is present for Pinned, not even an explicit nil
### GetType

`func (o *Chat) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Chat) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Chat) SetType(v string)`

SetType sets Type field to given value.


### GetSmsPrice

`func (o *Chat) GetSmsPrice() float32`

GetSmsPrice returns the SmsPrice field if non-nil, zero value otherwise.

### GetSmsPriceOk

`func (o *Chat) GetSmsPriceOk() (*float32, bool)`

GetSmsPriceOk returns a tuple with the SmsPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsPrice

`func (o *Chat) SetSmsPrice(v float32)`

SetSmsPrice sets SmsPrice field to given value.


### GetMmsPrice

`func (o *Chat) GetMmsPrice() float32`

GetMmsPrice returns the MmsPrice field if non-nil, zero value otherwise.

### GetMmsPriceOk

`func (o *Chat) GetMmsPriceOk() (*float32, bool)`

GetMmsPriceOk returns a tuple with the MmsPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMmsPrice

`func (o *Chat) SetMmsPrice(v float32)`

SetMmsPrice sets MmsPrice field to given value.


### GetTags

`func (o *Chat) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Chat) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Chat) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Chat) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


