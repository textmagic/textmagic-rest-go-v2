# NullableUserPersonalInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | User ID. | [optional] 
**FirstName** | Pointer to **NullableString** | User&#39;s first name. | [optional] 
**LastName** | Pointer to **NullableString** | User&#39;s last name. | [optional] 
**AvatarUrl** | Pointer to **NullableString** | URL to user&#39;s avatar image. | [optional] 
**Email** | Pointer to **NullableString** | User&#39;s email address. | [optional] 

## Methods

### NewNullableUserPersonalInfo

`func NewNullableUserPersonalInfo() *NullableUserPersonalInfo`

NewNullableUserPersonalInfo instantiates a new NullableUserPersonalInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNullableUserPersonalInfoWithDefaults

`func NewNullableUserPersonalInfoWithDefaults() *NullableUserPersonalInfo`

NewNullableUserPersonalInfoWithDefaults instantiates a new NullableUserPersonalInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NullableUserPersonalInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NullableUserPersonalInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NullableUserPersonalInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *NullableUserPersonalInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFirstName

`func (o *NullableUserPersonalInfo) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *NullableUserPersonalInfo) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *NullableUserPersonalInfo) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *NullableUserPersonalInfo) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *NullableUserPersonalInfo) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *NullableUserPersonalInfo) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *NullableUserPersonalInfo) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *NullableUserPersonalInfo) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *NullableUserPersonalInfo) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *NullableUserPersonalInfo) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *NullableUserPersonalInfo) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *NullableUserPersonalInfo) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetAvatarUrl

`func (o *NullableUserPersonalInfo) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *NullableUserPersonalInfo) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *NullableUserPersonalInfo) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *NullableUserPersonalInfo) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### SetAvatarUrlNil

`func (o *NullableUserPersonalInfo) SetAvatarUrlNil(b bool)`

 SetAvatarUrlNil sets the value for AvatarUrl to be an explicit nil

### UnsetAvatarUrl
`func (o *NullableUserPersonalInfo) UnsetAvatarUrl()`

UnsetAvatarUrl ensures that no value is present for AvatarUrl, not even an explicit nil
### GetEmail

`func (o *NullableUserPersonalInfo) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *NullableUserPersonalInfo) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *NullableUserPersonalInfo) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *NullableUserPersonalInfo) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *NullableUserPersonalInfo) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *NullableUserPersonalInfo) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


