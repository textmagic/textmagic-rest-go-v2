# MessagingStatItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReplyRate** | **NullableFloat32** | The number of incoming messages divided by the number of total messages. | 
**Date** | **NullableTime** | Time interval start: empty if the **by** parameter was set to **off**.  | 
**DeliveryRate** | **NullableFloat32** | Message delivery rate:the number of delivered messages divided by the number of total messages. | 
**Costs** | **NullableFloat32** | Cost for sent messages during this period. The costs are in the [Account](https://docs.textmagic.com/#tag/User) currency.  | 
**MessagesReceived** | **NullableInt32** | Total received messages count. | 
**MessagesSentDelivered** | **NullableInt32** | Delivered messages count. As messages are retried for up to 48 hours, this value could change. | 
**MessagesSentAccepted** | **NullableInt32** | Messages accepted for delivery (in queue) but not yet delivered. | 
**MessagesSentBuffered** | **NullableInt32** | Messages buffered by endpoint cell phone operators. | 
**MessagesSentFailed** | **NullableInt32** | Messages that have failed for whatever reason, e.g. the destination phone was switched off for 48 hours or the recipient&#39;s phone account is out of service. | 
**MessagesSentRejected** | **NullableInt32** | Messages that were rejected: invalid Sender ID used (e.g. you cannot use the Sender ID or your own mobile number when sending to the United States and Canada.)  | 
**MessagesSentParts** | **NullableInt32** | Total sent messages **parts** count. Note that this is not equal to the sent messages count, because one message could consist of 1 to 6 parts and users are charged per part, not per message. | 

## Methods

### NewMessagingStatItem

`func NewMessagingStatItem(replyRate NullableFloat32, date NullableTime, deliveryRate NullableFloat32, costs NullableFloat32, messagesReceived NullableInt32, messagesSentDelivered NullableInt32, messagesSentAccepted NullableInt32, messagesSentBuffered NullableInt32, messagesSentFailed NullableInt32, messagesSentRejected NullableInt32, messagesSentParts NullableInt32, ) *MessagingStatItem`

NewMessagingStatItem instantiates a new MessagingStatItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessagingStatItemWithDefaults

`func NewMessagingStatItemWithDefaults() *MessagingStatItem`

NewMessagingStatItemWithDefaults instantiates a new MessagingStatItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReplyRate

`func (o *MessagingStatItem) GetReplyRate() float32`

GetReplyRate returns the ReplyRate field if non-nil, zero value otherwise.

### GetReplyRateOk

`func (o *MessagingStatItem) GetReplyRateOk() (*float32, bool)`

GetReplyRateOk returns a tuple with the ReplyRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyRate

`func (o *MessagingStatItem) SetReplyRate(v float32)`

SetReplyRate sets ReplyRate field to given value.


### SetReplyRateNil

`func (o *MessagingStatItem) SetReplyRateNil(b bool)`

 SetReplyRateNil sets the value for ReplyRate to be an explicit nil

### UnsetReplyRate
`func (o *MessagingStatItem) UnsetReplyRate()`

UnsetReplyRate ensures that no value is present for ReplyRate, not even an explicit nil
### GetDate

`func (o *MessagingStatItem) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *MessagingStatItem) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *MessagingStatItem) SetDate(v time.Time)`

SetDate sets Date field to given value.


### SetDateNil

`func (o *MessagingStatItem) SetDateNil(b bool)`

 SetDateNil sets the value for Date to be an explicit nil

### UnsetDate
`func (o *MessagingStatItem) UnsetDate()`

UnsetDate ensures that no value is present for Date, not even an explicit nil
### GetDeliveryRate

`func (o *MessagingStatItem) GetDeliveryRate() float32`

GetDeliveryRate returns the DeliveryRate field if non-nil, zero value otherwise.

### GetDeliveryRateOk

`func (o *MessagingStatItem) GetDeliveryRateOk() (*float32, bool)`

GetDeliveryRateOk returns a tuple with the DeliveryRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryRate

`func (o *MessagingStatItem) SetDeliveryRate(v float32)`

SetDeliveryRate sets DeliveryRate field to given value.


### SetDeliveryRateNil

`func (o *MessagingStatItem) SetDeliveryRateNil(b bool)`

 SetDeliveryRateNil sets the value for DeliveryRate to be an explicit nil

### UnsetDeliveryRate
`func (o *MessagingStatItem) UnsetDeliveryRate()`

UnsetDeliveryRate ensures that no value is present for DeliveryRate, not even an explicit nil
### GetCosts

`func (o *MessagingStatItem) GetCosts() float32`

GetCosts returns the Costs field if non-nil, zero value otherwise.

### GetCostsOk

`func (o *MessagingStatItem) GetCostsOk() (*float32, bool)`

GetCostsOk returns a tuple with the Costs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCosts

`func (o *MessagingStatItem) SetCosts(v float32)`

SetCosts sets Costs field to given value.


### SetCostsNil

`func (o *MessagingStatItem) SetCostsNil(b bool)`

 SetCostsNil sets the value for Costs to be an explicit nil

### UnsetCosts
`func (o *MessagingStatItem) UnsetCosts()`

UnsetCosts ensures that no value is present for Costs, not even an explicit nil
### GetMessagesReceived

`func (o *MessagingStatItem) GetMessagesReceived() int32`

GetMessagesReceived returns the MessagesReceived field if non-nil, zero value otherwise.

### GetMessagesReceivedOk

`func (o *MessagingStatItem) GetMessagesReceivedOk() (*int32, bool)`

GetMessagesReceivedOk returns a tuple with the MessagesReceived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagesReceived

`func (o *MessagingStatItem) SetMessagesReceived(v int32)`

SetMessagesReceived sets MessagesReceived field to given value.


### SetMessagesReceivedNil

`func (o *MessagingStatItem) SetMessagesReceivedNil(b bool)`

 SetMessagesReceivedNil sets the value for MessagesReceived to be an explicit nil

### UnsetMessagesReceived
`func (o *MessagingStatItem) UnsetMessagesReceived()`

UnsetMessagesReceived ensures that no value is present for MessagesReceived, not even an explicit nil
### GetMessagesSentDelivered

`func (o *MessagingStatItem) GetMessagesSentDelivered() int32`

GetMessagesSentDelivered returns the MessagesSentDelivered field if non-nil, zero value otherwise.

### GetMessagesSentDeliveredOk

`func (o *MessagingStatItem) GetMessagesSentDeliveredOk() (*int32, bool)`

GetMessagesSentDeliveredOk returns a tuple with the MessagesSentDelivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagesSentDelivered

`func (o *MessagingStatItem) SetMessagesSentDelivered(v int32)`

SetMessagesSentDelivered sets MessagesSentDelivered field to given value.


### SetMessagesSentDeliveredNil

`func (o *MessagingStatItem) SetMessagesSentDeliveredNil(b bool)`

 SetMessagesSentDeliveredNil sets the value for MessagesSentDelivered to be an explicit nil

### UnsetMessagesSentDelivered
`func (o *MessagingStatItem) UnsetMessagesSentDelivered()`

UnsetMessagesSentDelivered ensures that no value is present for MessagesSentDelivered, not even an explicit nil
### GetMessagesSentAccepted

`func (o *MessagingStatItem) GetMessagesSentAccepted() int32`

GetMessagesSentAccepted returns the MessagesSentAccepted field if non-nil, zero value otherwise.

### GetMessagesSentAcceptedOk

`func (o *MessagingStatItem) GetMessagesSentAcceptedOk() (*int32, bool)`

GetMessagesSentAcceptedOk returns a tuple with the MessagesSentAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagesSentAccepted

`func (o *MessagingStatItem) SetMessagesSentAccepted(v int32)`

SetMessagesSentAccepted sets MessagesSentAccepted field to given value.


### SetMessagesSentAcceptedNil

`func (o *MessagingStatItem) SetMessagesSentAcceptedNil(b bool)`

 SetMessagesSentAcceptedNil sets the value for MessagesSentAccepted to be an explicit nil

### UnsetMessagesSentAccepted
`func (o *MessagingStatItem) UnsetMessagesSentAccepted()`

UnsetMessagesSentAccepted ensures that no value is present for MessagesSentAccepted, not even an explicit nil
### GetMessagesSentBuffered

`func (o *MessagingStatItem) GetMessagesSentBuffered() int32`

GetMessagesSentBuffered returns the MessagesSentBuffered field if non-nil, zero value otherwise.

### GetMessagesSentBufferedOk

`func (o *MessagingStatItem) GetMessagesSentBufferedOk() (*int32, bool)`

GetMessagesSentBufferedOk returns a tuple with the MessagesSentBuffered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagesSentBuffered

`func (o *MessagingStatItem) SetMessagesSentBuffered(v int32)`

SetMessagesSentBuffered sets MessagesSentBuffered field to given value.


### SetMessagesSentBufferedNil

`func (o *MessagingStatItem) SetMessagesSentBufferedNil(b bool)`

 SetMessagesSentBufferedNil sets the value for MessagesSentBuffered to be an explicit nil

### UnsetMessagesSentBuffered
`func (o *MessagingStatItem) UnsetMessagesSentBuffered()`

UnsetMessagesSentBuffered ensures that no value is present for MessagesSentBuffered, not even an explicit nil
### GetMessagesSentFailed

`func (o *MessagingStatItem) GetMessagesSentFailed() int32`

GetMessagesSentFailed returns the MessagesSentFailed field if non-nil, zero value otherwise.

### GetMessagesSentFailedOk

`func (o *MessagingStatItem) GetMessagesSentFailedOk() (*int32, bool)`

GetMessagesSentFailedOk returns a tuple with the MessagesSentFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagesSentFailed

`func (o *MessagingStatItem) SetMessagesSentFailed(v int32)`

SetMessagesSentFailed sets MessagesSentFailed field to given value.


### SetMessagesSentFailedNil

`func (o *MessagingStatItem) SetMessagesSentFailedNil(b bool)`

 SetMessagesSentFailedNil sets the value for MessagesSentFailed to be an explicit nil

### UnsetMessagesSentFailed
`func (o *MessagingStatItem) UnsetMessagesSentFailed()`

UnsetMessagesSentFailed ensures that no value is present for MessagesSentFailed, not even an explicit nil
### GetMessagesSentRejected

`func (o *MessagingStatItem) GetMessagesSentRejected() int32`

GetMessagesSentRejected returns the MessagesSentRejected field if non-nil, zero value otherwise.

### GetMessagesSentRejectedOk

`func (o *MessagingStatItem) GetMessagesSentRejectedOk() (*int32, bool)`

GetMessagesSentRejectedOk returns a tuple with the MessagesSentRejected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagesSentRejected

`func (o *MessagingStatItem) SetMessagesSentRejected(v int32)`

SetMessagesSentRejected sets MessagesSentRejected field to given value.


### SetMessagesSentRejectedNil

`func (o *MessagingStatItem) SetMessagesSentRejectedNil(b bool)`

 SetMessagesSentRejectedNil sets the value for MessagesSentRejected to be an explicit nil

### UnsetMessagesSentRejected
`func (o *MessagingStatItem) UnsetMessagesSentRejected()`

UnsetMessagesSentRejected ensures that no value is present for MessagesSentRejected, not even an explicit nil
### GetMessagesSentParts

`func (o *MessagingStatItem) GetMessagesSentParts() int32`

GetMessagesSentParts returns the MessagesSentParts field if non-nil, zero value otherwise.

### GetMessagesSentPartsOk

`func (o *MessagingStatItem) GetMessagesSentPartsOk() (*int32, bool)`

GetMessagesSentPartsOk returns a tuple with the MessagesSentParts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagesSentParts

`func (o *MessagingStatItem) SetMessagesSentParts(v int32)`

SetMessagesSentParts sets MessagesSentParts field to given value.


### SetMessagesSentPartsNil

`func (o *MessagingStatItem) SetMessagesSentPartsNil(b bool)`

 SetMessagesSentPartsNil sets the value for MessagesSentParts to be an explicit nil

### UnsetMessagesSentParts
`func (o *MessagingStatItem) UnsetMessagesSentParts()`

UnsetMessagesSentParts ensures that no value is present for MessagesSentParts, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


