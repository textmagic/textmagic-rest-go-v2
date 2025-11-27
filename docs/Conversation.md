# Conversation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Direction** | **string** | Message type: inbound or outbound.  | 
**Sender** | **NullableString** | Sender phone number. | 
**MessageTime** | **time.Time** | Time when  the message arrived at Textmagic. | 
**Text** | **string** | Message text. | 
**Receiver** | **string** | Receiver&#39;s phone number. | 
**Status** | **string** | Message status (for chats outbound only). See [message delivery statuses](https://docs.textmagic.com/#section/Delivery-status-codes) for details. | 
**FirstName** | **string** | Contact first name. | 
**LastName** | **string** | Contact last name. | 
**SessionId** | **NullableInt32** | Session ID of a message. See [message sessions](https://docs.textmagic.com/#tag/Outbound-Message-Sessions) for details. | 
**InitiatorId** | Pointer to **NullableInt32** | Initiator ID of a message. See [message sessions](https://docs.textmagic.com/#tag/Outbound-Message-Sessions) for details. | [optional] 
**MessageFileId** | Pointer to **NullableInt32** | Message file id. | [optional] 
**Type** | Pointer to **string** | Message type. | [optional] 
**ChatType** | Pointer to **string** | Chat type. | [optional] 
**ChatId** | Pointer to **int32** | Chat id. | [optional] 
**IsEdited** | Pointer to **bool** |  | [optional] 
**ErrorCode** | Pointer to **NullableString** | Error code. | [optional] 
**Files** | Pointer to [**[]File**](File.md) |  | [optional] 
**Payload** | Pointer to [**NullableMessagePayload**](MessagePayload.md) |  | [optional] 
**Avatar** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewConversation

`func NewConversation(id int32, direction string, sender NullableString, messageTime time.Time, text string, receiver string, status string, firstName string, lastName string, sessionId NullableInt32, ) *Conversation`

NewConversation instantiates a new Conversation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConversationWithDefaults

`func NewConversationWithDefaults() *Conversation`

NewConversationWithDefaults instantiates a new Conversation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Conversation) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Conversation) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Conversation) SetId(v int32)`

SetId sets Id field to given value.


### GetDirection

`func (o *Conversation) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *Conversation) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *Conversation) SetDirection(v string)`

SetDirection sets Direction field to given value.


### GetSender

`func (o *Conversation) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *Conversation) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *Conversation) SetSender(v string)`

SetSender sets Sender field to given value.


### SetSenderNil

`func (o *Conversation) SetSenderNil(b bool)`

 SetSenderNil sets the value for Sender to be an explicit nil

### UnsetSender
`func (o *Conversation) UnsetSender()`

UnsetSender ensures that no value is present for Sender, not even an explicit nil
### GetMessageTime

`func (o *Conversation) GetMessageTime() time.Time`

GetMessageTime returns the MessageTime field if non-nil, zero value otherwise.

### GetMessageTimeOk

`func (o *Conversation) GetMessageTimeOk() (*time.Time, bool)`

GetMessageTimeOk returns a tuple with the MessageTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageTime

`func (o *Conversation) SetMessageTime(v time.Time)`

SetMessageTime sets MessageTime field to given value.


### GetText

`func (o *Conversation) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *Conversation) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *Conversation) SetText(v string)`

SetText sets Text field to given value.


### GetReceiver

`func (o *Conversation) GetReceiver() string`

GetReceiver returns the Receiver field if non-nil, zero value otherwise.

### GetReceiverOk

`func (o *Conversation) GetReceiverOk() (*string, bool)`

GetReceiverOk returns a tuple with the Receiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiver

`func (o *Conversation) SetReceiver(v string)`

SetReceiver sets Receiver field to given value.


### GetStatus

`func (o *Conversation) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Conversation) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Conversation) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetFirstName

`func (o *Conversation) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *Conversation) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *Conversation) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *Conversation) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *Conversation) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *Conversation) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetSessionId

`func (o *Conversation) GetSessionId() int32`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *Conversation) GetSessionIdOk() (*int32, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *Conversation) SetSessionId(v int32)`

SetSessionId sets SessionId field to given value.


### SetSessionIdNil

`func (o *Conversation) SetSessionIdNil(b bool)`

 SetSessionIdNil sets the value for SessionId to be an explicit nil

### UnsetSessionId
`func (o *Conversation) UnsetSessionId()`

UnsetSessionId ensures that no value is present for SessionId, not even an explicit nil
### GetInitiatorId

`func (o *Conversation) GetInitiatorId() int32`

GetInitiatorId returns the InitiatorId field if non-nil, zero value otherwise.

### GetInitiatorIdOk

`func (o *Conversation) GetInitiatorIdOk() (*int32, bool)`

GetInitiatorIdOk returns a tuple with the InitiatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatorId

`func (o *Conversation) SetInitiatorId(v int32)`

SetInitiatorId sets InitiatorId field to given value.

### HasInitiatorId

`func (o *Conversation) HasInitiatorId() bool`

HasInitiatorId returns a boolean if a field has been set.

### SetInitiatorIdNil

`func (o *Conversation) SetInitiatorIdNil(b bool)`

 SetInitiatorIdNil sets the value for InitiatorId to be an explicit nil

### UnsetInitiatorId
`func (o *Conversation) UnsetInitiatorId()`

UnsetInitiatorId ensures that no value is present for InitiatorId, not even an explicit nil
### GetMessageFileId

`func (o *Conversation) GetMessageFileId() int32`

GetMessageFileId returns the MessageFileId field if non-nil, zero value otherwise.

### GetMessageFileIdOk

`func (o *Conversation) GetMessageFileIdOk() (*int32, bool)`

GetMessageFileIdOk returns a tuple with the MessageFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageFileId

`func (o *Conversation) SetMessageFileId(v int32)`

SetMessageFileId sets MessageFileId field to given value.

### HasMessageFileId

`func (o *Conversation) HasMessageFileId() bool`

HasMessageFileId returns a boolean if a field has been set.

### SetMessageFileIdNil

`func (o *Conversation) SetMessageFileIdNil(b bool)`

 SetMessageFileIdNil sets the value for MessageFileId to be an explicit nil

### UnsetMessageFileId
`func (o *Conversation) UnsetMessageFileId()`

UnsetMessageFileId ensures that no value is present for MessageFileId, not even an explicit nil
### GetType

`func (o *Conversation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Conversation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Conversation) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Conversation) HasType() bool`

HasType returns a boolean if a field has been set.

### GetChatType

`func (o *Conversation) GetChatType() string`

GetChatType returns the ChatType field if non-nil, zero value otherwise.

### GetChatTypeOk

`func (o *Conversation) GetChatTypeOk() (*string, bool)`

GetChatTypeOk returns a tuple with the ChatType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatType

`func (o *Conversation) SetChatType(v string)`

SetChatType sets ChatType field to given value.

### HasChatType

`func (o *Conversation) HasChatType() bool`

HasChatType returns a boolean if a field has been set.

### GetChatId

`func (o *Conversation) GetChatId() int32`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *Conversation) GetChatIdOk() (*int32, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *Conversation) SetChatId(v int32)`

SetChatId sets ChatId field to given value.

### HasChatId

`func (o *Conversation) HasChatId() bool`

HasChatId returns a boolean if a field has been set.

### GetIsEdited

`func (o *Conversation) GetIsEdited() bool`

GetIsEdited returns the IsEdited field if non-nil, zero value otherwise.

### GetIsEditedOk

`func (o *Conversation) GetIsEditedOk() (*bool, bool)`

GetIsEditedOk returns a tuple with the IsEdited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEdited

`func (o *Conversation) SetIsEdited(v bool)`

SetIsEdited sets IsEdited field to given value.

### HasIsEdited

`func (o *Conversation) HasIsEdited() bool`

HasIsEdited returns a boolean if a field has been set.

### GetErrorCode

`func (o *Conversation) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *Conversation) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *Conversation) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *Conversation) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *Conversation) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *Conversation) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
### GetFiles

`func (o *Conversation) GetFiles() []File`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *Conversation) GetFilesOk() (*[]File, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *Conversation) SetFiles(v []File)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *Conversation) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetPayload

`func (o *Conversation) GetPayload() MessagePayload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *Conversation) GetPayloadOk() (*MessagePayload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *Conversation) SetPayload(v MessagePayload)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *Conversation) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### SetPayloadNil

`func (o *Conversation) SetPayloadNil(b bool)`

 SetPayloadNil sets the value for Payload to be an explicit nil

### UnsetPayload
`func (o *Conversation) UnsetPayload()`

UnsetPayload ensures that no value is present for Payload, not even an explicit nil
### GetAvatar

`func (o *Conversation) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *Conversation) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *Conversation) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *Conversation) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### SetAvatarNil

`func (o *Conversation) SetAvatarNil(b bool)`

 SetAvatarNil sets the value for Avatar to be an explicit nil

### UnsetAvatar
`func (o *Conversation) UnsetAvatar()`

UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


