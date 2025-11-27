# ClearAndAssignContactsToListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contacts** | Pointer to **string** | Contact ID(s), separated by a comma or \&quot;all\&quot; to add all contacts belonging to the current user. | [optional] 

## Methods

### NewClearAndAssignContactsToListRequest

`func NewClearAndAssignContactsToListRequest() *ClearAndAssignContactsToListRequest`

NewClearAndAssignContactsToListRequest instantiates a new ClearAndAssignContactsToListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClearAndAssignContactsToListRequestWithDefaults

`func NewClearAndAssignContactsToListRequestWithDefaults() *ClearAndAssignContactsToListRequest`

NewClearAndAssignContactsToListRequestWithDefaults instantiates a new ClearAndAssignContactsToListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContacts

`func (o *ClearAndAssignContactsToListRequest) GetContacts() string`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *ClearAndAssignContactsToListRequest) GetContactsOk() (*string, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *ClearAndAssignContactsToListRequest) SetContacts(v string)`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *ClearAndAssignContactsToListRequest) HasContacts() bool`

HasContacts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


