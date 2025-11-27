# FavoriteContact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | **int32** | List or Contact ID. | 
**EntityType** | **string** | Entity type which should be marked as **favorite**. | 
**PrimaryLabel** | **string** | Contact first name/last name if entityType is **contact**; List name if entity type is **list**. | 
**SecondaryLabel** | **string** | Phone number if entityType is **contact**; List contacts number if entity type is **list**. | 
**TertiaryLabel** | **NullableString** | Contact country if entityType is **contact**; else, null. | 
**Avatar** | **NullableString** |  | 

## Methods

### NewFavoriteContact

`func NewFavoriteContact(entityId int32, entityType string, primaryLabel string, secondaryLabel string, tertiaryLabel NullableString, avatar NullableString, ) *FavoriteContact`

NewFavoriteContact instantiates a new FavoriteContact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFavoriteContactWithDefaults

`func NewFavoriteContactWithDefaults() *FavoriteContact`

NewFavoriteContactWithDefaults instantiates a new FavoriteContact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityId

`func (o *FavoriteContact) GetEntityId() int32`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *FavoriteContact) GetEntityIdOk() (*int32, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *FavoriteContact) SetEntityId(v int32)`

SetEntityId sets EntityId field to given value.


### GetEntityType

`func (o *FavoriteContact) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *FavoriteContact) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *FavoriteContact) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetPrimaryLabel

`func (o *FavoriteContact) GetPrimaryLabel() string`

GetPrimaryLabel returns the PrimaryLabel field if non-nil, zero value otherwise.

### GetPrimaryLabelOk

`func (o *FavoriteContact) GetPrimaryLabelOk() (*string, bool)`

GetPrimaryLabelOk returns a tuple with the PrimaryLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryLabel

`func (o *FavoriteContact) SetPrimaryLabel(v string)`

SetPrimaryLabel sets PrimaryLabel field to given value.


### GetSecondaryLabel

`func (o *FavoriteContact) GetSecondaryLabel() string`

GetSecondaryLabel returns the SecondaryLabel field if non-nil, zero value otherwise.

### GetSecondaryLabelOk

`func (o *FavoriteContact) GetSecondaryLabelOk() (*string, bool)`

GetSecondaryLabelOk returns a tuple with the SecondaryLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryLabel

`func (o *FavoriteContact) SetSecondaryLabel(v string)`

SetSecondaryLabel sets SecondaryLabel field to given value.


### GetTertiaryLabel

`func (o *FavoriteContact) GetTertiaryLabel() string`

GetTertiaryLabel returns the TertiaryLabel field if non-nil, zero value otherwise.

### GetTertiaryLabelOk

`func (o *FavoriteContact) GetTertiaryLabelOk() (*string, bool)`

GetTertiaryLabelOk returns a tuple with the TertiaryLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTertiaryLabel

`func (o *FavoriteContact) SetTertiaryLabel(v string)`

SetTertiaryLabel sets TertiaryLabel field to given value.


### SetTertiaryLabelNil

`func (o *FavoriteContact) SetTertiaryLabelNil(b bool)`

 SetTertiaryLabelNil sets the value for TertiaryLabel to be an explicit nil

### UnsetTertiaryLabel
`func (o *FavoriteContact) UnsetTertiaryLabel()`

UnsetTertiaryLabel ensures that no value is present for TertiaryLabel, not even an explicit nil
### GetAvatar

`func (o *FavoriteContact) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *FavoriteContact) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *FavoriteContact) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.


### SetAvatarNil

`func (o *FavoriteContact) SetAvatarNil(b bool)`

 SetAvatarNil sets the value for Avatar to be an explicit nil

### UnsetAvatar
`func (o *FavoriteContact) UnsetAvatar()`

UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


