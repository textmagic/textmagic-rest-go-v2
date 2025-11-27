# UnsubscribedContact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unsubscribed contact ID. | 
**Phone** | **string** | Phone number in [E.164 format](https://en.wikipedia.org/wiki/E.164). | 
**UnsubscribeTime** | **time.Time** | Time when contact was opted-out. | 
**FirstName** | **NullableString** | Unsubscribed contact first name. | 
**LastName** | **NullableString** | Unsubscribed contact last name. | 

## Methods

### NewUnsubscribedContact

`func NewUnsubscribedContact(id int32, phone string, unsubscribeTime time.Time, firstName NullableString, lastName NullableString, ) *UnsubscribedContact`

NewUnsubscribedContact instantiates a new UnsubscribedContact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnsubscribedContactWithDefaults

`func NewUnsubscribedContactWithDefaults() *UnsubscribedContact`

NewUnsubscribedContactWithDefaults instantiates a new UnsubscribedContact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UnsubscribedContact) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnsubscribedContact) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnsubscribedContact) SetId(v int32)`

SetId sets Id field to given value.


### GetPhone

`func (o *UnsubscribedContact) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *UnsubscribedContact) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *UnsubscribedContact) SetPhone(v string)`

SetPhone sets Phone field to given value.


### GetUnsubscribeTime

`func (o *UnsubscribedContact) GetUnsubscribeTime() time.Time`

GetUnsubscribeTime returns the UnsubscribeTime field if non-nil, zero value otherwise.

### GetUnsubscribeTimeOk

`func (o *UnsubscribedContact) GetUnsubscribeTimeOk() (*time.Time, bool)`

GetUnsubscribeTimeOk returns a tuple with the UnsubscribeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsubscribeTime

`func (o *UnsubscribedContact) SetUnsubscribeTime(v time.Time)`

SetUnsubscribeTime sets UnsubscribeTime field to given value.


### GetFirstName

`func (o *UnsubscribedContact) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UnsubscribedContact) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UnsubscribedContact) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### SetFirstNameNil

`func (o *UnsubscribedContact) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *UnsubscribedContact) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *UnsubscribedContact) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UnsubscribedContact) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UnsubscribedContact) SetLastName(v string)`

SetLastName sets LastName field to given value.


### SetLastNameNil

`func (o *UnsubscribedContact) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *UnsubscribedContact) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


