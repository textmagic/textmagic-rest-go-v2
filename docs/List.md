# List

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | List ID. | 
**Name** | **string** | List name. | 
**Description** | **NullableString** | Description of the list. | 
**Favorited** | **bool** | Is the List favorited? See [Favorites list](https://docs.textmagic.com/#operation/getFavourites). | 
**MembersCount** | **int32** | List members count. | 
**User** | [**NullableUser**](User.md) |  | 
**Service** | **bool** | Internal service field. | 
**Shared** | **NullableBool** | Is the list **shared** among all sub-accounts? | 
**Avatar** | [**NullableListImage**](ListImage.md) |  | 
**IsDefault** | **NullableBool** | Indicates that List is used as a default. All new contacts added via the Web-app will be added in this List by default. | 

## Methods

### NewList

`func NewList(id int32, name string, description NullableString, favorited bool, membersCount int32, user NullableUser, service bool, shared NullableBool, avatar NullableListImage, isDefault NullableBool, ) *List`

NewList instantiates a new List object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWithDefaults

`func NewListWithDefaults() *List`

NewListWithDefaults instantiates a new List object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *List) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *List) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *List) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *List) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *List) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *List) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *List) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *List) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *List) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *List) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *List) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetFavorited

`func (o *List) GetFavorited() bool`

GetFavorited returns the Favorited field if non-nil, zero value otherwise.

### GetFavoritedOk

`func (o *List) GetFavoritedOk() (*bool, bool)`

GetFavoritedOk returns a tuple with the Favorited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavorited

`func (o *List) SetFavorited(v bool)`

SetFavorited sets Favorited field to given value.


### GetMembersCount

`func (o *List) GetMembersCount() int32`

GetMembersCount returns the MembersCount field if non-nil, zero value otherwise.

### GetMembersCountOk

`func (o *List) GetMembersCountOk() (*int32, bool)`

GetMembersCountOk returns a tuple with the MembersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersCount

`func (o *List) SetMembersCount(v int32)`

SetMembersCount sets MembersCount field to given value.


### GetUser

`func (o *List) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *List) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *List) SetUser(v User)`

SetUser sets User field to given value.


### SetUserNil

`func (o *List) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *List) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetService

`func (o *List) GetService() bool`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *List) GetServiceOk() (*bool, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *List) SetService(v bool)`

SetService sets Service field to given value.


### GetShared

`func (o *List) GetShared() bool`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *List) GetSharedOk() (*bool, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *List) SetShared(v bool)`

SetShared sets Shared field to given value.


### SetSharedNil

`func (o *List) SetSharedNil(b bool)`

 SetSharedNil sets the value for Shared to be an explicit nil

### UnsetShared
`func (o *List) UnsetShared()`

UnsetShared ensures that no value is present for Shared, not even an explicit nil
### GetAvatar

`func (o *List) GetAvatar() ListImage`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *List) GetAvatarOk() (*ListImage, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *List) SetAvatar(v ListImage)`

SetAvatar sets Avatar field to given value.


### SetAvatarNil

`func (o *List) SetAvatarNil(b bool)`

 SetAvatarNil sets the value for Avatar to be an explicit nil

### UnsetAvatar
`func (o *List) UnsetAvatar()`

UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil
### GetIsDefault

`func (o *List) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *List) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *List) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### SetIsDefaultNil

`func (o *List) SetIsDefaultNil(b bool)`

 SetIsDefaultNil sets the value for IsDefault to be an explicit nil

### UnsetIsDefault
`func (o *List) UnsetIsDefault()`

UnsetIsDefault ensures that no value is present for IsDefault, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


