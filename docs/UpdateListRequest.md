# UpdateListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | List name. | [optional] 
**Shared** | Pointer to **bool** | Make this list shared or not? | [optional] [default to false]
**Favorited** | Pointer to **bool** | Is list favorited. | [optional] [default to false]
**IsDefault** | Pointer to **bool** | Is list default for new contacts (web only). | [optional] [default to false]

## Methods

### NewUpdateListRequest

`func NewUpdateListRequest() *UpdateListRequest`

NewUpdateListRequest instantiates a new UpdateListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateListRequestWithDefaults

`func NewUpdateListRequestWithDefaults() *UpdateListRequest`

NewUpdateListRequestWithDefaults instantiates a new UpdateListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateListRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateListRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateListRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateListRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetShared

`func (o *UpdateListRequest) GetShared() bool`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *UpdateListRequest) GetSharedOk() (*bool, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *UpdateListRequest) SetShared(v bool)`

SetShared sets Shared field to given value.

### HasShared

`func (o *UpdateListRequest) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetFavorited

`func (o *UpdateListRequest) GetFavorited() bool`

GetFavorited returns the Favorited field if non-nil, zero value otherwise.

### GetFavoritedOk

`func (o *UpdateListRequest) GetFavoritedOk() (*bool, bool)`

GetFavoritedOk returns a tuple with the Favorited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavorited

`func (o *UpdateListRequest) SetFavorited(v bool)`

SetFavorited sets Favorited field to given value.

### HasFavorited

`func (o *UpdateListRequest) HasFavorited() bool`

HasFavorited returns a boolean if a field has been set.

### GetIsDefault

`func (o *UpdateListRequest) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *UpdateListRequest) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *UpdateListRequest) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *UpdateListRequest) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


