# ContactNote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Contact note ID. | 
**CreatedAt** | **time.Time** | Contact note creation time. | 
**Note** | **string** | Contact note text. | 
**User** | [**NullableUser**](User.md) |  | 

## Methods

### NewContactNote

`func NewContactNote(id int32, createdAt time.Time, note string, user NullableUser, ) *ContactNote`

NewContactNote instantiates a new ContactNote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactNoteWithDefaults

`func NewContactNoteWithDefaults() *ContactNote`

NewContactNoteWithDefaults instantiates a new ContactNote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ContactNote) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContactNote) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContactNote) SetId(v int32)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *ContactNote) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ContactNote) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ContactNote) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetNote

`func (o *ContactNote) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *ContactNote) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *ContactNote) SetNote(v string)`

SetNote sets Note field to given value.


### GetUser

`func (o *ContactNote) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ContactNote) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ContactNote) SetUser(v User)`

SetUser sets User field to given value.


### SetUserNil

`func (o *ContactNote) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *ContactNote) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


