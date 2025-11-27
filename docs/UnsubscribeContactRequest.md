# UnsubscribeContactRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phone** | Pointer to **string** | Contact phone number. | [optional] 
**BlockIncoming** | Pointer to **int32** | If set to 1, incoming messages from this number will be blocked. | [optional] 

## Methods

### NewUnsubscribeContactRequest

`func NewUnsubscribeContactRequest() *UnsubscribeContactRequest`

NewUnsubscribeContactRequest instantiates a new UnsubscribeContactRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnsubscribeContactRequestWithDefaults

`func NewUnsubscribeContactRequestWithDefaults() *UnsubscribeContactRequest`

NewUnsubscribeContactRequestWithDefaults instantiates a new UnsubscribeContactRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhone

`func (o *UnsubscribeContactRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *UnsubscribeContactRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *UnsubscribeContactRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *UnsubscribeContactRequest) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetBlockIncoming

`func (o *UnsubscribeContactRequest) GetBlockIncoming() int32`

GetBlockIncoming returns the BlockIncoming field if non-nil, zero value otherwise.

### GetBlockIncomingOk

`func (o *UnsubscribeContactRequest) GetBlockIncomingOk() (*int32, bool)`

GetBlockIncomingOk returns a tuple with the BlockIncoming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockIncoming

`func (o *UnsubscribeContactRequest) SetBlockIncoming(v int32)`

SetBlockIncoming sets BlockIncoming field to given value.

### HasBlockIncoming

`func (o *UnsubscribeContactRequest) HasBlockIncoming() bool`

HasBlockIncoming returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


