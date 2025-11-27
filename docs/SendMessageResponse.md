# SendMessageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Message ID. | 
**Href** | **string** | URI of the message session. | 
**Type** | **string** | Message response type: * **message** – when the message is sent to a single recipient. * **session** – when the message is sent is to multiple recipients. * **schedule** - when the message is scheduled for sending. * **bulk** - when the message is sent to multiple recipients and the number of recipients requires asynchronous processing See [Sending more than 1,000 messages in one session](https://docs.textmagic.com/#section/Tutorials/Sending-more-than-1000-messages-in-one-session).  | 
**SessionId** | **NullableInt32** | Message session ID. | 
**BulkId** | **NullableInt32** | Bulk Session ID. See [Sending more than 1,000 messages in one session](https://docs.textmagic.com/#section/Tutorials/Sending-more-than-1000-messages-in-one-session). | 
**MessageId** | **NullableInt32** | Message ID. | 
**ScheduleId** | **NullableInt32** | Message Schedule ID. | 
**ChatId** | **NullableInt32** | Message Chat ID. | 

## Methods

### NewSendMessageResponse

`func NewSendMessageResponse(id int32, href string, type_ string, sessionId NullableInt32, bulkId NullableInt32, messageId NullableInt32, scheduleId NullableInt32, chatId NullableInt32, ) *SendMessageResponse`

NewSendMessageResponse instantiates a new SendMessageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendMessageResponseWithDefaults

`func NewSendMessageResponseWithDefaults() *SendMessageResponse`

NewSendMessageResponseWithDefaults instantiates a new SendMessageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SendMessageResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SendMessageResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SendMessageResponse) SetId(v int32)`

SetId sets Id field to given value.


### GetHref

`func (o *SendMessageResponse) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *SendMessageResponse) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *SendMessageResponse) SetHref(v string)`

SetHref sets Href field to given value.


### GetType

`func (o *SendMessageResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SendMessageResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SendMessageResponse) SetType(v string)`

SetType sets Type field to given value.


### GetSessionId

`func (o *SendMessageResponse) GetSessionId() int32`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *SendMessageResponse) GetSessionIdOk() (*int32, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *SendMessageResponse) SetSessionId(v int32)`

SetSessionId sets SessionId field to given value.


### SetSessionIdNil

`func (o *SendMessageResponse) SetSessionIdNil(b bool)`

 SetSessionIdNil sets the value for SessionId to be an explicit nil

### UnsetSessionId
`func (o *SendMessageResponse) UnsetSessionId()`

UnsetSessionId ensures that no value is present for SessionId, not even an explicit nil
### GetBulkId

`func (o *SendMessageResponse) GetBulkId() int32`

GetBulkId returns the BulkId field if non-nil, zero value otherwise.

### GetBulkIdOk

`func (o *SendMessageResponse) GetBulkIdOk() (*int32, bool)`

GetBulkIdOk returns a tuple with the BulkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkId

`func (o *SendMessageResponse) SetBulkId(v int32)`

SetBulkId sets BulkId field to given value.


### SetBulkIdNil

`func (o *SendMessageResponse) SetBulkIdNil(b bool)`

 SetBulkIdNil sets the value for BulkId to be an explicit nil

### UnsetBulkId
`func (o *SendMessageResponse) UnsetBulkId()`

UnsetBulkId ensures that no value is present for BulkId, not even an explicit nil
### GetMessageId

`func (o *SendMessageResponse) GetMessageId() int32`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *SendMessageResponse) GetMessageIdOk() (*int32, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *SendMessageResponse) SetMessageId(v int32)`

SetMessageId sets MessageId field to given value.


### SetMessageIdNil

`func (o *SendMessageResponse) SetMessageIdNil(b bool)`

 SetMessageIdNil sets the value for MessageId to be an explicit nil

### UnsetMessageId
`func (o *SendMessageResponse) UnsetMessageId()`

UnsetMessageId ensures that no value is present for MessageId, not even an explicit nil
### GetScheduleId

`func (o *SendMessageResponse) GetScheduleId() int32`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *SendMessageResponse) GetScheduleIdOk() (*int32, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *SendMessageResponse) SetScheduleId(v int32)`

SetScheduleId sets ScheduleId field to given value.


### SetScheduleIdNil

`func (o *SendMessageResponse) SetScheduleIdNil(b bool)`

 SetScheduleIdNil sets the value for ScheduleId to be an explicit nil

### UnsetScheduleId
`func (o *SendMessageResponse) UnsetScheduleId()`

UnsetScheduleId ensures that no value is present for ScheduleId, not even an explicit nil
### GetChatId

`func (o *SendMessageResponse) GetChatId() int32`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *SendMessageResponse) GetChatIdOk() (*int32, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *SendMessageResponse) SetChatId(v int32)`

SetChatId sets ChatId field to given value.


### SetChatIdNil

`func (o *SendMessageResponse) SetChatIdNil(b bool)`

 SetChatIdNil sets the value for ChatId to be an explicit nil

### UnsetChatId
`func (o *SendMessageResponse) UnsetChatId()`

UnsetChatId ensures that no value is present for ChatId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


