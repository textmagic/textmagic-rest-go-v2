# MessageOutSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** |  | [optional] 
**StartTime** | Pointer to **NullableString** |  | [optional] 
**Text** | Pointer to **NullableString** |  | [optional] 
**Price** | Pointer to **NullableFloat32** |  | [optional] 
**NumbersCount** | Pointer to **NullableInt32** |  | [optional] 
**Destination** | Pointer to **NullableString** |  | [optional] 
**Source** | Pointer to **NullableString** |  | [optional] 
**ReferenceId** | Pointer to **NullableString** |  | [optional] 
**InitiatorId** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewMessageOutSession

`func NewMessageOutSession() *MessageOutSession`

NewMessageOutSession instantiates a new MessageOutSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageOutSessionWithDefaults

`func NewMessageOutSessionWithDefaults() *MessageOutSession`

NewMessageOutSessionWithDefaults instantiates a new MessageOutSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MessageOutSession) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MessageOutSession) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MessageOutSession) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MessageOutSession) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *MessageOutSession) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *MessageOutSession) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetStartTime

`func (o *MessageOutSession) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *MessageOutSession) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *MessageOutSession) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *MessageOutSession) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *MessageOutSession) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *MessageOutSession) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetText

`func (o *MessageOutSession) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *MessageOutSession) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *MessageOutSession) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *MessageOutSession) HasText() bool`

HasText returns a boolean if a field has been set.

### SetTextNil

`func (o *MessageOutSession) SetTextNil(b bool)`

 SetTextNil sets the value for Text to be an explicit nil

### UnsetText
`func (o *MessageOutSession) UnsetText()`

UnsetText ensures that no value is present for Text, not even an explicit nil
### GetPrice

`func (o *MessageOutSession) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *MessageOutSession) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *MessageOutSession) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *MessageOutSession) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *MessageOutSession) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *MessageOutSession) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetNumbersCount

`func (o *MessageOutSession) GetNumbersCount() int32`

GetNumbersCount returns the NumbersCount field if non-nil, zero value otherwise.

### GetNumbersCountOk

`func (o *MessageOutSession) GetNumbersCountOk() (*int32, bool)`

GetNumbersCountOk returns a tuple with the NumbersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumbersCount

`func (o *MessageOutSession) SetNumbersCount(v int32)`

SetNumbersCount sets NumbersCount field to given value.

### HasNumbersCount

`func (o *MessageOutSession) HasNumbersCount() bool`

HasNumbersCount returns a boolean if a field has been set.

### SetNumbersCountNil

`func (o *MessageOutSession) SetNumbersCountNil(b bool)`

 SetNumbersCountNil sets the value for NumbersCount to be an explicit nil

### UnsetNumbersCount
`func (o *MessageOutSession) UnsetNumbersCount()`

UnsetNumbersCount ensures that no value is present for NumbersCount, not even an explicit nil
### GetDestination

`func (o *MessageOutSession) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *MessageOutSession) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *MessageOutSession) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *MessageOutSession) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### SetDestinationNil

`func (o *MessageOutSession) SetDestinationNil(b bool)`

 SetDestinationNil sets the value for Destination to be an explicit nil

### UnsetDestination
`func (o *MessageOutSession) UnsetDestination()`

UnsetDestination ensures that no value is present for Destination, not even an explicit nil
### GetSource

`func (o *MessageOutSession) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *MessageOutSession) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *MessageOutSession) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *MessageOutSession) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *MessageOutSession) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *MessageOutSession) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetReferenceId

`func (o *MessageOutSession) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *MessageOutSession) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *MessageOutSession) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *MessageOutSession) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### SetReferenceIdNil

`func (o *MessageOutSession) SetReferenceIdNil(b bool)`

 SetReferenceIdNil sets the value for ReferenceId to be an explicit nil

### UnsetReferenceId
`func (o *MessageOutSession) UnsetReferenceId()`

UnsetReferenceId ensures that no value is present for ReferenceId, not even an explicit nil
### GetInitiatorId

`func (o *MessageOutSession) GetInitiatorId() int32`

GetInitiatorId returns the InitiatorId field if non-nil, zero value otherwise.

### GetInitiatorIdOk

`func (o *MessageOutSession) GetInitiatorIdOk() (*int32, bool)`

GetInitiatorIdOk returns a tuple with the InitiatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatorId

`func (o *MessageOutSession) SetInitiatorId(v int32)`

SetInitiatorId sets InitiatorId field to given value.

### HasInitiatorId

`func (o *MessageOutSession) HasInitiatorId() bool`

HasInitiatorId returns a boolean if a field has been set.

### SetInitiatorIdNil

`func (o *MessageOutSession) SetInitiatorIdNil(b bool)`

 SetInitiatorIdNil sets the value for InitiatorId to be an explicit nil

### UnsetInitiatorId
`func (o *MessageOutSession) UnsetInitiatorId()`

UnsetInitiatorId ensures that no value is present for InitiatorId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


