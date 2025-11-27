# DeleteScheduledMessagesBulkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | Pointer to **string** | Entity ID(s), separated by comma. | [optional] 
**All** | Pointer to **int32** | Default is 0 (false). If set to 1, all the entities will be removed. | [optional] 
**Status** | Pointer to **string** | Default is an empty string (false). If set, all entities with specified status will be affected. | [optional] [default to ""]

## Methods

### NewDeleteScheduledMessagesBulkRequest

`func NewDeleteScheduledMessagesBulkRequest() *DeleteScheduledMessagesBulkRequest`

NewDeleteScheduledMessagesBulkRequest instantiates a new DeleteScheduledMessagesBulkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteScheduledMessagesBulkRequestWithDefaults

`func NewDeleteScheduledMessagesBulkRequestWithDefaults() *DeleteScheduledMessagesBulkRequest`

NewDeleteScheduledMessagesBulkRequestWithDefaults instantiates a new DeleteScheduledMessagesBulkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *DeleteScheduledMessagesBulkRequest) GetIds() string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *DeleteScheduledMessagesBulkRequest) GetIdsOk() (*string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *DeleteScheduledMessagesBulkRequest) SetIds(v string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *DeleteScheduledMessagesBulkRequest) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetAll

`func (o *DeleteScheduledMessagesBulkRequest) GetAll() int32`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *DeleteScheduledMessagesBulkRequest) GetAllOk() (*int32, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *DeleteScheduledMessagesBulkRequest) SetAll(v int32)`

SetAll sets All field to given value.

### HasAll

`func (o *DeleteScheduledMessagesBulkRequest) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetStatus

`func (o *DeleteScheduledMessagesBulkRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeleteScheduledMessagesBulkRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeleteScheduledMessagesBulkRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeleteScheduledMessagesBulkRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


