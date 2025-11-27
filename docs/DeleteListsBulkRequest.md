# DeleteListsBulkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | Pointer to **string** | Entity ID(s), separated by comma. | [optional] 
**All** | Pointer to **int32** | Default is 0 (false). If set to 1, all the entities will be removed. | [optional] 

## Methods

### NewDeleteListsBulkRequest

`func NewDeleteListsBulkRequest() *DeleteListsBulkRequest`

NewDeleteListsBulkRequest instantiates a new DeleteListsBulkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteListsBulkRequestWithDefaults

`func NewDeleteListsBulkRequestWithDefaults() *DeleteListsBulkRequest`

NewDeleteListsBulkRequestWithDefaults instantiates a new DeleteListsBulkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *DeleteListsBulkRequest) GetIds() string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *DeleteListsBulkRequest) GetIdsOk() (*string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *DeleteListsBulkRequest) SetIds(v string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *DeleteListsBulkRequest) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetAll

`func (o *DeleteListsBulkRequest) GetAll() int32`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *DeleteListsBulkRequest) GetAllOk() (*int32, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *DeleteListsBulkRequest) SetAll(v int32)`

SetAll sets All field to given value.

### HasAll

`func (o *DeleteListsBulkRequest) HasAll() bool`

HasAll returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


