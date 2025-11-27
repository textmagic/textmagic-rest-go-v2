# MuteChatRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Chat ID. | [optional] 
**Mute** | Pointer to **bool** | Mute notifications sound. | [optional] 
**For** | Pointer to **int32** | Mute for N hours. | [optional] 

## Methods

### NewMuteChatRequest

`func NewMuteChatRequest() *MuteChatRequest`

NewMuteChatRequest instantiates a new MuteChatRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMuteChatRequestWithDefaults

`func NewMuteChatRequestWithDefaults() *MuteChatRequest`

NewMuteChatRequestWithDefaults instantiates a new MuteChatRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MuteChatRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MuteChatRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MuteChatRequest) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MuteChatRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMute

`func (o *MuteChatRequest) GetMute() bool`

GetMute returns the Mute field if non-nil, zero value otherwise.

### GetMuteOk

`func (o *MuteChatRequest) GetMuteOk() (*bool, bool)`

GetMuteOk returns a tuple with the Mute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMute

`func (o *MuteChatRequest) SetMute(v bool)`

SetMute sets Mute field to given value.

### HasMute

`func (o *MuteChatRequest) HasMute() bool`

HasMute returns a boolean if a field has been set.

### GetFor

`func (o *MuteChatRequest) GetFor() int32`

GetFor returns the For field if non-nil, zero value otherwise.

### GetForOk

`func (o *MuteChatRequest) GetForOk() (*int32, bool)`

GetForOk returns a tuple with the For field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFor

`func (o *MuteChatRequest) SetFor(v int32)`

SetFor sets For field to given value.

### HasFor

`func (o *MuteChatRequest) HasFor() bool`

HasFor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


