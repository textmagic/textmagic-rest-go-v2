# DeleteContactNotesBulkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | Pointer to **string** | Entity ID(s), separated by comma. | [optional] 
**All** | Pointer to **bool** | Entity ID(s), separated by comma. | [optional] 

## Methods

### NewDeleteContactNotesBulkRequest

`func NewDeleteContactNotesBulkRequest() *DeleteContactNotesBulkRequest`

NewDeleteContactNotesBulkRequest instantiates a new DeleteContactNotesBulkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteContactNotesBulkRequestWithDefaults

`func NewDeleteContactNotesBulkRequestWithDefaults() *DeleteContactNotesBulkRequest`

NewDeleteContactNotesBulkRequestWithDefaults instantiates a new DeleteContactNotesBulkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *DeleteContactNotesBulkRequest) GetIds() string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *DeleteContactNotesBulkRequest) GetIdsOk() (*string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *DeleteContactNotesBulkRequest) SetIds(v string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *DeleteContactNotesBulkRequest) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetAll

`func (o *DeleteContactNotesBulkRequest) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *DeleteContactNotesBulkRequest) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *DeleteContactNotesBulkRequest) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *DeleteContactNotesBulkRequest) HasAll() bool`

HasAll returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


