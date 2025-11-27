# UnblockContactsBulkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | Pointer to **string** | Entity ID(s), separated by comma. | [optional] 
**All** | Pointer to **int32** | Default is 0 (false). If set to 1, all entities will be removed. | [optional] 

## Methods

### NewUnblockContactsBulkRequest

`func NewUnblockContactsBulkRequest() *UnblockContactsBulkRequest`

NewUnblockContactsBulkRequest instantiates a new UnblockContactsBulkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnblockContactsBulkRequestWithDefaults

`func NewUnblockContactsBulkRequestWithDefaults() *UnblockContactsBulkRequest`

NewUnblockContactsBulkRequestWithDefaults instantiates a new UnblockContactsBulkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *UnblockContactsBulkRequest) GetIds() string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *UnblockContactsBulkRequest) GetIdsOk() (*string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *UnblockContactsBulkRequest) SetIds(v string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *UnblockContactsBulkRequest) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetAll

`func (o *UnblockContactsBulkRequest) GetAll() int32`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *UnblockContactsBulkRequest) GetAllOk() (*int32, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *UnblockContactsBulkRequest) SetAll(v int32)`

SetAll sets All field to given value.

### HasAll

`func (o *UnblockContactsBulkRequest) HasAll() bool`

HasAll returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


