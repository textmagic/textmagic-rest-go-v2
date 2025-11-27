# EmailSenderItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique email sender identifier. | 
**DomainId** | **int32** | ID of the associated domain. | 
**Email** | **string** | Email address of the sender. | 
**CreatedAt** | **time.Time** | When the email sender was created. | 
**DomainStatus** | **string** | Current verification status of the associated domain. | 
**FromName** | Pointer to **NullableString** | Display name for the sender. | [optional] 
**ReplyTo** | Pointer to **NullableString** | Reply-to email address. | [optional] 

## Methods

### NewEmailSenderItem

`func NewEmailSenderItem(id int32, domainId int32, email string, createdAt time.Time, domainStatus string, ) *EmailSenderItem`

NewEmailSenderItem instantiates a new EmailSenderItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailSenderItemWithDefaults

`func NewEmailSenderItemWithDefaults() *EmailSenderItem`

NewEmailSenderItemWithDefaults instantiates a new EmailSenderItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EmailSenderItem) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmailSenderItem) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmailSenderItem) SetId(v int32)`

SetId sets Id field to given value.


### GetDomainId

`func (o *EmailSenderItem) GetDomainId() int32`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *EmailSenderItem) GetDomainIdOk() (*int32, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *EmailSenderItem) SetDomainId(v int32)`

SetDomainId sets DomainId field to given value.


### GetEmail

`func (o *EmailSenderItem) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *EmailSenderItem) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *EmailSenderItem) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetCreatedAt

`func (o *EmailSenderItem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EmailSenderItem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EmailSenderItem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDomainStatus

`func (o *EmailSenderItem) GetDomainStatus() string`

GetDomainStatus returns the DomainStatus field if non-nil, zero value otherwise.

### GetDomainStatusOk

`func (o *EmailSenderItem) GetDomainStatusOk() (*string, bool)`

GetDomainStatusOk returns a tuple with the DomainStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainStatus

`func (o *EmailSenderItem) SetDomainStatus(v string)`

SetDomainStatus sets DomainStatus field to given value.


### GetFromName

`func (o *EmailSenderItem) GetFromName() string`

GetFromName returns the FromName field if non-nil, zero value otherwise.

### GetFromNameOk

`func (o *EmailSenderItem) GetFromNameOk() (*string, bool)`

GetFromNameOk returns a tuple with the FromName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromName

`func (o *EmailSenderItem) SetFromName(v string)`

SetFromName sets FromName field to given value.

### HasFromName

`func (o *EmailSenderItem) HasFromName() bool`

HasFromName returns a boolean if a field has been set.

### SetFromNameNil

`func (o *EmailSenderItem) SetFromNameNil(b bool)`

 SetFromNameNil sets the value for FromName to be an explicit nil

### UnsetFromName
`func (o *EmailSenderItem) UnsetFromName()`

UnsetFromName ensures that no value is present for FromName, not even an explicit nil
### GetReplyTo

`func (o *EmailSenderItem) GetReplyTo() string`

GetReplyTo returns the ReplyTo field if non-nil, zero value otherwise.

### GetReplyToOk

`func (o *EmailSenderItem) GetReplyToOk() (*string, bool)`

GetReplyToOk returns a tuple with the ReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyTo

`func (o *EmailSenderItem) SetReplyTo(v string)`

SetReplyTo sets ReplyTo field to given value.

### HasReplyTo

`func (o *EmailSenderItem) HasReplyTo() bool`

HasReplyTo returns a boolean if a field has been set.

### SetReplyToNil

`func (o *EmailSenderItem) SetReplyToNil(b bool)`

 SetReplyToNil sets the value for ReplyTo to be an explicit nil

### UnsetReplyTo
`func (o *EmailSenderItem) UnsetReplyTo()`

UnsetReplyTo ensures that no value is present for ReplyTo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


