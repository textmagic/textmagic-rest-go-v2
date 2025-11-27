# DeleteChatMessagesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InboundIds** | Pointer to **string** | Inbound message IDs to delete. Require when \&quot;all\&quot; is equal to 0 (false). | [optional] 
**SentIds** | Pointer to **string** | Sent message IDs to delete. Require when \&quot;all\&quot; is equal to 0 (false). | [optional] 
**CallsIds** | Pointer to **string** | Calls IDs to delete. Require when \&quot;all\&quot; is equal to 0 (false). | [optional] 
**All** | Pointer to **bool** | Default is 0 (false). If set to 1, all the entities will be removed. | [optional] 

## Methods

### NewDeleteChatMessagesRequest

`func NewDeleteChatMessagesRequest() *DeleteChatMessagesRequest`

NewDeleteChatMessagesRequest instantiates a new DeleteChatMessagesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteChatMessagesRequestWithDefaults

`func NewDeleteChatMessagesRequestWithDefaults() *DeleteChatMessagesRequest`

NewDeleteChatMessagesRequestWithDefaults instantiates a new DeleteChatMessagesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInboundIds

`func (o *DeleteChatMessagesRequest) GetInboundIds() string`

GetInboundIds returns the InboundIds field if non-nil, zero value otherwise.

### GetInboundIdsOk

`func (o *DeleteChatMessagesRequest) GetInboundIdsOk() (*string, bool)`

GetInboundIdsOk returns a tuple with the InboundIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundIds

`func (o *DeleteChatMessagesRequest) SetInboundIds(v string)`

SetInboundIds sets InboundIds field to given value.

### HasInboundIds

`func (o *DeleteChatMessagesRequest) HasInboundIds() bool`

HasInboundIds returns a boolean if a field has been set.

### GetSentIds

`func (o *DeleteChatMessagesRequest) GetSentIds() string`

GetSentIds returns the SentIds field if non-nil, zero value otherwise.

### GetSentIdsOk

`func (o *DeleteChatMessagesRequest) GetSentIdsOk() (*string, bool)`

GetSentIdsOk returns a tuple with the SentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentIds

`func (o *DeleteChatMessagesRequest) SetSentIds(v string)`

SetSentIds sets SentIds field to given value.

### HasSentIds

`func (o *DeleteChatMessagesRequest) HasSentIds() bool`

HasSentIds returns a boolean if a field has been set.

### GetCallsIds

`func (o *DeleteChatMessagesRequest) GetCallsIds() string`

GetCallsIds returns the CallsIds field if non-nil, zero value otherwise.

### GetCallsIdsOk

`func (o *DeleteChatMessagesRequest) GetCallsIdsOk() (*string, bool)`

GetCallsIdsOk returns a tuple with the CallsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallsIds

`func (o *DeleteChatMessagesRequest) SetCallsIds(v string)`

SetCallsIds sets CallsIds field to given value.

### HasCallsIds

`func (o *DeleteChatMessagesRequest) HasCallsIds() bool`

HasCallsIds returns a boolean if a field has been set.

### GetAll

`func (o *DeleteChatMessagesRequest) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *DeleteChatMessagesRequest) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *DeleteChatMessagesRequest) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *DeleteChatMessagesRequest) HasAll() bool`

HasAll returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


