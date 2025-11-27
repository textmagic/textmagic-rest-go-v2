# MessageIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The ID of the inbound message. | 
**Sender** | **string** | The sender’s phone number. | 
**Receiver** | **NullableString** | The receiver’s phone number (i.e. your dedicated or shared reply number). | 
**MessageTime** | **time.Time** | The time when the message reached the Textmagic API endpoint. | 
**Text** | **string** | The text from the received message. | 
**ContactId** | Pointer to **NullableInt32** | Sender contact ID. | [optional] 
**FirstName** | Pointer to **NullableString** | Sender contact first name. | [optional] 
**LastName** | Pointer to **NullableString** | Sender contact last name. | [optional] 
**Avatar** | **NullableString** |  | 
**Email** | Pointer to **NullableString** | Sender email. | [optional] 
**ContactUserId** | Pointer to **NullableInt32** |  | [optional] 
**UserId** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewMessageIn

`func NewMessageIn(id int32, sender string, receiver NullableString, messageTime time.Time, text string, avatar NullableString, ) *MessageIn`

NewMessageIn instantiates a new MessageIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageInWithDefaults

`func NewMessageInWithDefaults() *MessageIn`

NewMessageInWithDefaults instantiates a new MessageIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MessageIn) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MessageIn) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MessageIn) SetId(v int32)`

SetId sets Id field to given value.


### GetSender

`func (o *MessageIn) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *MessageIn) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *MessageIn) SetSender(v string)`

SetSender sets Sender field to given value.


### GetReceiver

`func (o *MessageIn) GetReceiver() string`

GetReceiver returns the Receiver field if non-nil, zero value otherwise.

### GetReceiverOk

`func (o *MessageIn) GetReceiverOk() (*string, bool)`

GetReceiverOk returns a tuple with the Receiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiver

`func (o *MessageIn) SetReceiver(v string)`

SetReceiver sets Receiver field to given value.


### SetReceiverNil

`func (o *MessageIn) SetReceiverNil(b bool)`

 SetReceiverNil sets the value for Receiver to be an explicit nil

### UnsetReceiver
`func (o *MessageIn) UnsetReceiver()`

UnsetReceiver ensures that no value is present for Receiver, not even an explicit nil
### GetMessageTime

`func (o *MessageIn) GetMessageTime() time.Time`

GetMessageTime returns the MessageTime field if non-nil, zero value otherwise.

### GetMessageTimeOk

`func (o *MessageIn) GetMessageTimeOk() (*time.Time, bool)`

GetMessageTimeOk returns a tuple with the MessageTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageTime

`func (o *MessageIn) SetMessageTime(v time.Time)`

SetMessageTime sets MessageTime field to given value.


### GetText

`func (o *MessageIn) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *MessageIn) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *MessageIn) SetText(v string)`

SetText sets Text field to given value.


### GetContactId

`func (o *MessageIn) GetContactId() int32`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *MessageIn) GetContactIdOk() (*int32, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *MessageIn) SetContactId(v int32)`

SetContactId sets ContactId field to given value.

### HasContactId

`func (o *MessageIn) HasContactId() bool`

HasContactId returns a boolean if a field has been set.

### SetContactIdNil

`func (o *MessageIn) SetContactIdNil(b bool)`

 SetContactIdNil sets the value for ContactId to be an explicit nil

### UnsetContactId
`func (o *MessageIn) UnsetContactId()`

UnsetContactId ensures that no value is present for ContactId, not even an explicit nil
### GetFirstName

`func (o *MessageIn) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *MessageIn) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *MessageIn) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *MessageIn) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *MessageIn) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *MessageIn) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *MessageIn) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *MessageIn) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *MessageIn) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *MessageIn) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *MessageIn) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *MessageIn) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetAvatar

`func (o *MessageIn) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *MessageIn) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *MessageIn) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.


### SetAvatarNil

`func (o *MessageIn) SetAvatarNil(b bool)`

 SetAvatarNil sets the value for Avatar to be an explicit nil

### UnsetAvatar
`func (o *MessageIn) UnsetAvatar()`

UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil
### GetEmail

`func (o *MessageIn) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *MessageIn) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *MessageIn) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *MessageIn) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *MessageIn) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *MessageIn) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetContactUserId

`func (o *MessageIn) GetContactUserId() int32`

GetContactUserId returns the ContactUserId field if non-nil, zero value otherwise.

### GetContactUserIdOk

`func (o *MessageIn) GetContactUserIdOk() (*int32, bool)`

GetContactUserIdOk returns a tuple with the ContactUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactUserId

`func (o *MessageIn) SetContactUserId(v int32)`

SetContactUserId sets ContactUserId field to given value.

### HasContactUserId

`func (o *MessageIn) HasContactUserId() bool`

HasContactUserId returns a boolean if a field has been set.

### SetContactUserIdNil

`func (o *MessageIn) SetContactUserIdNil(b bool)`

 SetContactUserIdNil sets the value for ContactUserId to be an explicit nil

### UnsetContactUserId
`func (o *MessageIn) UnsetContactUserId()`

UnsetContactUserId ensures that no value is present for ContactUserId, not even an explicit nil
### GetUserId

`func (o *MessageIn) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *MessageIn) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *MessageIn) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *MessageIn) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *MessageIn) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *MessageIn) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


