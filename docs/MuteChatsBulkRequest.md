# MuteChatsBulkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | Pointer to **string** | Entity ID(s), separated by comma. | [optional] 
**All** | Pointer to **bool** | Entity ID(s), separated by comma | [optional] 
**For** | Pointer to **int32** | Mute for N hours. | [optional] 

## Methods

### NewMuteChatsBulkRequest

`func NewMuteChatsBulkRequest() *MuteChatsBulkRequest`

NewMuteChatsBulkRequest instantiates a new MuteChatsBulkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMuteChatsBulkRequestWithDefaults

`func NewMuteChatsBulkRequestWithDefaults() *MuteChatsBulkRequest`

NewMuteChatsBulkRequestWithDefaults instantiates a new MuteChatsBulkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *MuteChatsBulkRequest) GetIds() string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *MuteChatsBulkRequest) GetIdsOk() (*string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *MuteChatsBulkRequest) SetIds(v string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *MuteChatsBulkRequest) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetAll

`func (o *MuteChatsBulkRequest) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *MuteChatsBulkRequest) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *MuteChatsBulkRequest) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *MuteChatsBulkRequest) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetFor

`func (o *MuteChatsBulkRequest) GetFor() int32`

GetFor returns the For field if non-nil, zero value otherwise.

### GetForOk

`func (o *MuteChatsBulkRequest) GetForOk() (*int32, bool)`

GetForOk returns a tuple with the For field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFor

`func (o *MuteChatsBulkRequest) SetFor(v int32)`

SetFor sets For field to given value.

### HasFor

`func (o *MuteChatsBulkRequest) HasFor() bool`

HasFor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


