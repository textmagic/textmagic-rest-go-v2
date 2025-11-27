# Contact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Contact ID. | 
**Favorited** | **bool** | Is the Contact favorite? [Favorite list](https://docs.textmagic.com/#operation/getFavorites). | 
**Blocked** | **bool** | Is the Contact blocked? [Blocked contacts](https://docs.textmagic.com/#operation/getBlockedContacts). | 
**FirstName** | **NullableString** | Contact first name. | 
**LastName** | **NullableString** | Contact last name. | 
**CompanyName** | **NullableString** | Company name. | 
**Phone** | **NullableString** | Phone number in [E.164 format](https://en.wikipedia.org/wiki/E.164). | 
**Email** | **NullableString** | Contact email address. | 
**Country** | [**NullableCountry**](Country.md) |  | 
**CustomFields** | [**[]CustomFieldListItem**](CustomFieldListItem.md) |  | 
**User** | [**NullableUser**](User.md) |  | 
**Lists** | [**[]List**](List.md) |  | 
**Owner** | Pointer to [**NullableUser**](User.md) |  | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 
**PhoneType** | **NullableString** | Phone number type: * **0** if it is fixed-line; * **1** if it is mobile; * **2** if it is mobile or fixed-line (in case we cannot distingush between fixed-line or mobile); * **3** if it is toll-free; * **4** if it is a premium rate phone; * **5** if it is a shared cost phone; * **6** if it is a VoIP; * **7** if it is a [Personal Number](); * **8** if it is a pager; * **9** if it is a Universal Access Number; * **10** if the phone type is unknown; * **-1** if the phone type is not yet processed or cannot be determined.  | 
**Avatar** | [**NullableContactImage**](ContactImage.md) |  | 
**Notes** | [**[]ContactNote**](ContactNote.md) |  | 
**WhatsappPhone** | Pointer to **NullableString** | Whatsapp phone number in [E.164 format](https://en.wikipedia.org/wiki/E.164). | [optional] 

## Methods

### NewContact

`func NewContact(id int32, favorited bool, blocked bool, firstName NullableString, lastName NullableString, companyName NullableString, phone NullableString, email NullableString, country NullableCountry, customFields []CustomFieldListItem, user NullableUser, lists []List, phoneType NullableString, avatar NullableContactImage, notes []ContactNote, ) *Contact`

NewContact instantiates a new Contact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactWithDefaults

`func NewContactWithDefaults() *Contact`

NewContactWithDefaults instantiates a new Contact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Contact) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Contact) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Contact) SetId(v int32)`

SetId sets Id field to given value.


### GetFavorited

`func (o *Contact) GetFavorited() bool`

GetFavorited returns the Favorited field if non-nil, zero value otherwise.

### GetFavoritedOk

`func (o *Contact) GetFavoritedOk() (*bool, bool)`

GetFavoritedOk returns a tuple with the Favorited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavorited

`func (o *Contact) SetFavorited(v bool)`

SetFavorited sets Favorited field to given value.


### GetBlocked

`func (o *Contact) GetBlocked() bool`

GetBlocked returns the Blocked field if non-nil, zero value otherwise.

### GetBlockedOk

`func (o *Contact) GetBlockedOk() (*bool, bool)`

GetBlockedOk returns a tuple with the Blocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocked

`func (o *Contact) SetBlocked(v bool)`

SetBlocked sets Blocked field to given value.


### GetFirstName

`func (o *Contact) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *Contact) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *Contact) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### SetFirstNameNil

`func (o *Contact) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *Contact) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *Contact) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *Contact) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *Contact) SetLastName(v string)`

SetLastName sets LastName field to given value.


### SetLastNameNil

`func (o *Contact) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *Contact) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetCompanyName

`func (o *Contact) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *Contact) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *Contact) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.


### SetCompanyNameNil

`func (o *Contact) SetCompanyNameNil(b bool)`

 SetCompanyNameNil sets the value for CompanyName to be an explicit nil

### UnsetCompanyName
`func (o *Contact) UnsetCompanyName()`

UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
### GetPhone

`func (o *Contact) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *Contact) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *Contact) SetPhone(v string)`

SetPhone sets Phone field to given value.


### SetPhoneNil

`func (o *Contact) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *Contact) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetEmail

`func (o *Contact) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Contact) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Contact) SetEmail(v string)`

SetEmail sets Email field to given value.


### SetEmailNil

`func (o *Contact) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *Contact) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetCountry

`func (o *Contact) GetCountry() Country`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Contact) GetCountryOk() (*Country, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Contact) SetCountry(v Country)`

SetCountry sets Country field to given value.


### SetCountryNil

`func (o *Contact) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *Contact) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetCustomFields

`func (o *Contact) GetCustomFields() []CustomFieldListItem`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Contact) GetCustomFieldsOk() (*[]CustomFieldListItem, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Contact) SetCustomFields(v []CustomFieldListItem)`

SetCustomFields sets CustomFields field to given value.


### GetUser

`func (o *Contact) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Contact) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Contact) SetUser(v User)`

SetUser sets User field to given value.


### SetUserNil

`func (o *Contact) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *Contact) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetLists

`func (o *Contact) GetLists() []List`

GetLists returns the Lists field if non-nil, zero value otherwise.

### GetListsOk

`func (o *Contact) GetListsOk() (*[]List, bool)`

GetListsOk returns a tuple with the Lists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLists

`func (o *Contact) SetLists(v []List)`

SetLists sets Lists field to given value.


### GetOwner

`func (o *Contact) GetOwner() User`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Contact) GetOwnerOk() (*User, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Contact) SetOwner(v User)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Contact) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *Contact) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *Contact) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetTags

`func (o *Contact) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Contact) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Contact) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Contact) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetPhoneType

`func (o *Contact) GetPhoneType() string`

GetPhoneType returns the PhoneType field if non-nil, zero value otherwise.

### GetPhoneTypeOk

`func (o *Contact) GetPhoneTypeOk() (*string, bool)`

GetPhoneTypeOk returns a tuple with the PhoneType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneType

`func (o *Contact) SetPhoneType(v string)`

SetPhoneType sets PhoneType field to given value.


### SetPhoneTypeNil

`func (o *Contact) SetPhoneTypeNil(b bool)`

 SetPhoneTypeNil sets the value for PhoneType to be an explicit nil

### UnsetPhoneType
`func (o *Contact) UnsetPhoneType()`

UnsetPhoneType ensures that no value is present for PhoneType, not even an explicit nil
### GetAvatar

`func (o *Contact) GetAvatar() ContactImage`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *Contact) GetAvatarOk() (*ContactImage, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *Contact) SetAvatar(v ContactImage)`

SetAvatar sets Avatar field to given value.


### SetAvatarNil

`func (o *Contact) SetAvatarNil(b bool)`

 SetAvatarNil sets the value for Avatar to be an explicit nil

### UnsetAvatar
`func (o *Contact) UnsetAvatar()`

UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil
### GetNotes

`func (o *Contact) GetNotes() []ContactNote`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *Contact) GetNotesOk() (*[]ContactNote, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *Contact) SetNotes(v []ContactNote)`

SetNotes sets Notes field to given value.


### GetWhatsappPhone

`func (o *Contact) GetWhatsappPhone() string`

GetWhatsappPhone returns the WhatsappPhone field if non-nil, zero value otherwise.

### GetWhatsappPhoneOk

`func (o *Contact) GetWhatsappPhoneOk() (*string, bool)`

GetWhatsappPhoneOk returns a tuple with the WhatsappPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhatsappPhone

`func (o *Contact) SetWhatsappPhone(v string)`

SetWhatsappPhone sets WhatsappPhone field to given value.

### HasWhatsappPhone

`func (o *Contact) HasWhatsappPhone() bool`

HasWhatsappPhone returns a boolean if a field has been set.

### SetWhatsappPhoneNil

`func (o *Contact) SetWhatsappPhoneNil(b bool)`

 SetWhatsappPhoneNil sets the value for WhatsappPhone to be an explicit nil

### UnsetWhatsappPhone
`func (o *Contact) UnsetWhatsappPhone()`

UnsetWhatsappPhone ensures that no value is present for WhatsappPhone, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


