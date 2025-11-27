# CreateListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | List name. | [optional] 
**Shared** | Pointer to **bool** | Should the new list be **shared** among all the sub-accounts? | [optional] [default to false]
**Favorited** | Pointer to **bool** | Is the list favorited? Default is false. | [optional] [default to false]
**IsDefault** | Pointer to **bool** | Is the list default for new contacts (web only)? | [optional] [default to false]

## Methods

### NewCreateListRequest

`func NewCreateListRequest() *CreateListRequest`

NewCreateListRequest instantiates a new CreateListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateListRequestWithDefaults

`func NewCreateListRequestWithDefaults() *CreateListRequest`

NewCreateListRequestWithDefaults instantiates a new CreateListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateListRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateListRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateListRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateListRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetShared

`func (o *CreateListRequest) GetShared() bool`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *CreateListRequest) GetSharedOk() (*bool, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *CreateListRequest) SetShared(v bool)`

SetShared sets Shared field to given value.

### HasShared

`func (o *CreateListRequest) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetFavorited

`func (o *CreateListRequest) GetFavorited() bool`

GetFavorited returns the Favorited field if non-nil, zero value otherwise.

### GetFavoritedOk

`func (o *CreateListRequest) GetFavoritedOk() (*bool, bool)`

GetFavoritedOk returns a tuple with the Favorited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavorited

`func (o *CreateListRequest) SetFavorited(v bool)`

SetFavorited sets Favorited field to given value.

### HasFavorited

`func (o *CreateListRequest) HasFavorited() bool`

HasFavorited returns a boolean if a field has been set.

### GetIsDefault

`func (o *CreateListRequest) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *CreateListRequest) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *CreateListRequest) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *CreateListRequest) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


