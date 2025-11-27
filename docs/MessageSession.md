# MessageSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Session ID. | 
**StartTime** | **NullableString** | Session creation time. | 
**Text** | **NullableString** | Session text. If a template was used for the session text (see [Messages: Send](https://docs.textmagic.com/#tag/Outbound-Messages) for details), it may contain template tags.  | 
**Source** | **NullableString** | *   **O** – for Textmagic Online; *   **A** – for API; *   **M** – for Textmagic Messenger; *   **E** – for [Email to SMS](https://docs.textmagic.com/#tag/Send-Email-to-SMS); *   **X** – for [Distribution Lists](https://docs.textmagic.com/#tag/Distribution-Lists).  | 
**ReferenceId** | **NullableString** | Custom reference ID (see [Messages: Send](https://docs.textmagic.com/#tag/Send-Email-to-SMS) for details).  | 
**Price** | **NullableFloat32** | Session cost (in account currency). | 
**NumbersCount** | **NullableInt32** | Session recipient count. | 
**Destination** | **NullableString** | Destination type of a Message Session: * **t** – text SMS; * **s** – text-to-speech; * **v** – voice broadcast.  | 
**InitiatorId** | **NullableInt32** | Initiator ID. | 
**Title** | **NullableString** |  | 

## Methods

### NewMessageSession

`func NewMessageSession(id int32, startTime NullableString, text NullableString, source NullableString, referenceId NullableString, price NullableFloat32, numbersCount NullableInt32, destination NullableString, initiatorId NullableInt32, title NullableString, ) *MessageSession`

NewMessageSession instantiates a new MessageSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageSessionWithDefaults

`func NewMessageSessionWithDefaults() *MessageSession`

NewMessageSessionWithDefaults instantiates a new MessageSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MessageSession) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MessageSession) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MessageSession) SetId(v int32)`

SetId sets Id field to given value.


### GetStartTime

`func (o *MessageSession) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *MessageSession) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *MessageSession) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.


### SetStartTimeNil

`func (o *MessageSession) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *MessageSession) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetText

`func (o *MessageSession) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *MessageSession) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *MessageSession) SetText(v string)`

SetText sets Text field to given value.


### SetTextNil

`func (o *MessageSession) SetTextNil(b bool)`

 SetTextNil sets the value for Text to be an explicit nil

### UnsetText
`func (o *MessageSession) UnsetText()`

UnsetText ensures that no value is present for Text, not even an explicit nil
### GetSource

`func (o *MessageSession) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *MessageSession) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *MessageSession) SetSource(v string)`

SetSource sets Source field to given value.


### SetSourceNil

`func (o *MessageSession) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *MessageSession) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetReferenceId

`func (o *MessageSession) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *MessageSession) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *MessageSession) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.


### SetReferenceIdNil

`func (o *MessageSession) SetReferenceIdNil(b bool)`

 SetReferenceIdNil sets the value for ReferenceId to be an explicit nil

### UnsetReferenceId
`func (o *MessageSession) UnsetReferenceId()`

UnsetReferenceId ensures that no value is present for ReferenceId, not even an explicit nil
### GetPrice

`func (o *MessageSession) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *MessageSession) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *MessageSession) SetPrice(v float32)`

SetPrice sets Price field to given value.


### SetPriceNil

`func (o *MessageSession) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *MessageSession) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetNumbersCount

`func (o *MessageSession) GetNumbersCount() int32`

GetNumbersCount returns the NumbersCount field if non-nil, zero value otherwise.

### GetNumbersCountOk

`func (o *MessageSession) GetNumbersCountOk() (*int32, bool)`

GetNumbersCountOk returns a tuple with the NumbersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumbersCount

`func (o *MessageSession) SetNumbersCount(v int32)`

SetNumbersCount sets NumbersCount field to given value.


### SetNumbersCountNil

`func (o *MessageSession) SetNumbersCountNil(b bool)`

 SetNumbersCountNil sets the value for NumbersCount to be an explicit nil

### UnsetNumbersCount
`func (o *MessageSession) UnsetNumbersCount()`

UnsetNumbersCount ensures that no value is present for NumbersCount, not even an explicit nil
### GetDestination

`func (o *MessageSession) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *MessageSession) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *MessageSession) SetDestination(v string)`

SetDestination sets Destination field to given value.


### SetDestinationNil

`func (o *MessageSession) SetDestinationNil(b bool)`

 SetDestinationNil sets the value for Destination to be an explicit nil

### UnsetDestination
`func (o *MessageSession) UnsetDestination()`

UnsetDestination ensures that no value is present for Destination, not even an explicit nil
### GetInitiatorId

`func (o *MessageSession) GetInitiatorId() int32`

GetInitiatorId returns the InitiatorId field if non-nil, zero value otherwise.

### GetInitiatorIdOk

`func (o *MessageSession) GetInitiatorIdOk() (*int32, bool)`

GetInitiatorIdOk returns a tuple with the InitiatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatorId

`func (o *MessageSession) SetInitiatorId(v int32)`

SetInitiatorId sets InitiatorId field to given value.


### SetInitiatorIdNil

`func (o *MessageSession) SetInitiatorIdNil(b bool)`

 SetInitiatorIdNil sets the value for InitiatorId to be an explicit nil

### UnsetInitiatorId
`func (o *MessageSession) UnsetInitiatorId()`

UnsetInitiatorId ensures that no value is present for InitiatorId, not even an explicit nil
### GetTitle

`func (o *MessageSession) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MessageSession) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MessageSession) SetTitle(v string)`

SetTitle sets Title field to given value.


### SetTitleNil

`func (o *MessageSession) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *MessageSession) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


