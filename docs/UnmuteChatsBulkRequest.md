# UnmuteChatsBulkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | Pointer to **string** | Entity ID(s), separated by comma. | [optional] 
**All** | Pointer to **bool** | Entity ID(s), separated by comma | [optional] 

## Methods

### NewUnmuteChatsBulkRequest

`func NewUnmuteChatsBulkRequest() *UnmuteChatsBulkRequest`

NewUnmuteChatsBulkRequest instantiates a new UnmuteChatsBulkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnmuteChatsBulkRequestWithDefaults

`func NewUnmuteChatsBulkRequestWithDefaults() *UnmuteChatsBulkRequest`

NewUnmuteChatsBulkRequestWithDefaults instantiates a new UnmuteChatsBulkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *UnmuteChatsBulkRequest) GetIds() string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *UnmuteChatsBulkRequest) GetIdsOk() (*string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *UnmuteChatsBulkRequest) SetIds(v string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *UnmuteChatsBulkRequest) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetAll

`func (o *UnmuteChatsBulkRequest) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *UnmuteChatsBulkRequest) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *UnmuteChatsBulkRequest) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *UnmuteChatsBulkRequest) HasAll() bool`

HasAll returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


