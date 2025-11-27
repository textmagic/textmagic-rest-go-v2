# DeleteChatsBulkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | Pointer to **string** | Entity ID(s), separated by comma. | [optional] 
**All** | Pointer to **bool** | Entity ID(s), separated by comma. | [optional] 
**Status** | Pointer to **string** | Default is an empty string. If set, all entities with specified status will be affected. | [optional] 

## Methods

### NewDeleteChatsBulkRequest

`func NewDeleteChatsBulkRequest() *DeleteChatsBulkRequest`

NewDeleteChatsBulkRequest instantiates a new DeleteChatsBulkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteChatsBulkRequestWithDefaults

`func NewDeleteChatsBulkRequestWithDefaults() *DeleteChatsBulkRequest`

NewDeleteChatsBulkRequestWithDefaults instantiates a new DeleteChatsBulkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *DeleteChatsBulkRequest) GetIds() string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *DeleteChatsBulkRequest) GetIdsOk() (*string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *DeleteChatsBulkRequest) SetIds(v string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *DeleteChatsBulkRequest) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetAll

`func (o *DeleteChatsBulkRequest) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *DeleteChatsBulkRequest) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *DeleteChatsBulkRequest) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *DeleteChatsBulkRequest) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetStatus

`func (o *DeleteChatsBulkRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeleteChatsBulkRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeleteChatsBulkRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeleteChatsBulkRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


