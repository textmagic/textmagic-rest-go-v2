# GetContactsAutocompleteResponseItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | **int32** | Id of entity. 0 if object is a reply. | 
**EntityType** | **string** | Entry type: * **contact** if it is related to a contact; * **list** if it is related to a contact list; * **reply** if it is related to an incoming message.  | 
**Value** | **string** | ID of the contact/list if entityType is contact/list OR phone number if entityType is reply. | 
**Label** | **string** | Name of the contact/list if entityType is contact/list OR phone number if entityType is reply. | 
**SharedBy** | **string** | If contact or list was shared by another sub-account, the name of this user will be shown. | 
**IsShared** | **bool** | If contact or list was shared by another sub-account then &#x60;true&#x60; will be set. | 
**Avatar** | **NullableString** | Contact avatar URI. | 
**Favorited** | **bool** | If contact has been marked as favorite. | 
**UserId** | **int32** | Owner ID of the contact/list (if it was shared). | 
**CountryName** | **string** |  | 
**Qposition** | **int32** |  | 
**Rposition** | **int32** |  | 

## Methods

### NewGetContactsAutocompleteResponseItem

`func NewGetContactsAutocompleteResponseItem(entityId int32, entityType string, value string, label string, sharedBy string, isShared bool, avatar NullableString, favorited bool, userId int32, countryName string, qposition int32, rposition int32, ) *GetContactsAutocompleteResponseItem`

NewGetContactsAutocompleteResponseItem instantiates a new GetContactsAutocompleteResponseItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetContactsAutocompleteResponseItemWithDefaults

`func NewGetContactsAutocompleteResponseItemWithDefaults() *GetContactsAutocompleteResponseItem`

NewGetContactsAutocompleteResponseItemWithDefaults instantiates a new GetContactsAutocompleteResponseItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityId

`func (o *GetContactsAutocompleteResponseItem) GetEntityId() int32`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *GetContactsAutocompleteResponseItem) GetEntityIdOk() (*int32, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *GetContactsAutocompleteResponseItem) SetEntityId(v int32)`

SetEntityId sets EntityId field to given value.


### GetEntityType

`func (o *GetContactsAutocompleteResponseItem) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *GetContactsAutocompleteResponseItem) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *GetContactsAutocompleteResponseItem) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetValue

`func (o *GetContactsAutocompleteResponseItem) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GetContactsAutocompleteResponseItem) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GetContactsAutocompleteResponseItem) SetValue(v string)`

SetValue sets Value field to given value.


### GetLabel

`func (o *GetContactsAutocompleteResponseItem) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *GetContactsAutocompleteResponseItem) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *GetContactsAutocompleteResponseItem) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetSharedBy

`func (o *GetContactsAutocompleteResponseItem) GetSharedBy() string`

GetSharedBy returns the SharedBy field if non-nil, zero value otherwise.

### GetSharedByOk

`func (o *GetContactsAutocompleteResponseItem) GetSharedByOk() (*string, bool)`

GetSharedByOk returns a tuple with the SharedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedBy

`func (o *GetContactsAutocompleteResponseItem) SetSharedBy(v string)`

SetSharedBy sets SharedBy field to given value.


### GetIsShared

`func (o *GetContactsAutocompleteResponseItem) GetIsShared() bool`

GetIsShared returns the IsShared field if non-nil, zero value otherwise.

### GetIsSharedOk

`func (o *GetContactsAutocompleteResponseItem) GetIsSharedOk() (*bool, bool)`

GetIsSharedOk returns a tuple with the IsShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShared

`func (o *GetContactsAutocompleteResponseItem) SetIsShared(v bool)`

SetIsShared sets IsShared field to given value.


### GetAvatar

`func (o *GetContactsAutocompleteResponseItem) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *GetContactsAutocompleteResponseItem) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *GetContactsAutocompleteResponseItem) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.


### SetAvatarNil

`func (o *GetContactsAutocompleteResponseItem) SetAvatarNil(b bool)`

 SetAvatarNil sets the value for Avatar to be an explicit nil

### UnsetAvatar
`func (o *GetContactsAutocompleteResponseItem) UnsetAvatar()`

UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil
### GetFavorited

`func (o *GetContactsAutocompleteResponseItem) GetFavorited() bool`

GetFavorited returns the Favorited field if non-nil, zero value otherwise.

### GetFavoritedOk

`func (o *GetContactsAutocompleteResponseItem) GetFavoritedOk() (*bool, bool)`

GetFavoritedOk returns a tuple with the Favorited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavorited

`func (o *GetContactsAutocompleteResponseItem) SetFavorited(v bool)`

SetFavorited sets Favorited field to given value.


### GetUserId

`func (o *GetContactsAutocompleteResponseItem) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *GetContactsAutocompleteResponseItem) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *GetContactsAutocompleteResponseItem) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetCountryName

`func (o *GetContactsAutocompleteResponseItem) GetCountryName() string`

GetCountryName returns the CountryName field if non-nil, zero value otherwise.

### GetCountryNameOk

`func (o *GetContactsAutocompleteResponseItem) GetCountryNameOk() (*string, bool)`

GetCountryNameOk returns a tuple with the CountryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryName

`func (o *GetContactsAutocompleteResponseItem) SetCountryName(v string)`

SetCountryName sets CountryName field to given value.


### GetQposition

`func (o *GetContactsAutocompleteResponseItem) GetQposition() int32`

GetQposition returns the Qposition field if non-nil, zero value otherwise.

### GetQpositionOk

`func (o *GetContactsAutocompleteResponseItem) GetQpositionOk() (*int32, bool)`

GetQpositionOk returns a tuple with the Qposition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQposition

`func (o *GetContactsAutocompleteResponseItem) SetQposition(v int32)`

SetQposition sets Qposition field to given value.


### GetRposition

`func (o *GetContactsAutocompleteResponseItem) GetRposition() int32`

GetRposition returns the Rposition field if non-nil, zero value otherwise.

### GetRpositionOk

`func (o *GetContactsAutocompleteResponseItem) GetRpositionOk() (*int32, bool)`

GetRpositionOk returns a tuple with the Rposition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRposition

`func (o *GetContactsAutocompleteResponseItem) SetRposition(v int32)`

SetRposition sets Rposition field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


