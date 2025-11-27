# DeleteOutboundMessagesBulkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | Pointer to **string** | Entity ID(s), separated by comma. | [optional] 
**All** | Pointer to **int32** | Default is 0 (false). If set to 1, all the entities will be removed. | [optional] 

## Methods

### NewDeleteOutboundMessagesBulkRequest

`func NewDeleteOutboundMessagesBulkRequest() *DeleteOutboundMessagesBulkRequest`

NewDeleteOutboundMessagesBulkRequest instantiates a new DeleteOutboundMessagesBulkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteOutboundMessagesBulkRequestWithDefaults

`func NewDeleteOutboundMessagesBulkRequestWithDefaults() *DeleteOutboundMessagesBulkRequest`

NewDeleteOutboundMessagesBulkRequestWithDefaults instantiates a new DeleteOutboundMessagesBulkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *DeleteOutboundMessagesBulkRequest) GetIds() string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *DeleteOutboundMessagesBulkRequest) GetIdsOk() (*string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *DeleteOutboundMessagesBulkRequest) SetIds(v string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *DeleteOutboundMessagesBulkRequest) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetAll

`func (o *DeleteOutboundMessagesBulkRequest) GetAll() int32`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *DeleteOutboundMessagesBulkRequest) GetAllOk() (*int32, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *DeleteOutboundMessagesBulkRequest) SetAll(v int32)`

SetAll sets All field to given value.

### HasAll

`func (o *DeleteOutboundMessagesBulkRequest) HasAll() bool`

HasAll returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


