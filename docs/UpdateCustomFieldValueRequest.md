# UpdateCustomFieldValueRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContactId** | Pointer to **int32** | Contact ID. See [Contact](https://docs.textmagic.com/#tag/Contacts).  | [optional] 
**Value** | Pointer to **string** | Custom field value. Note that this value is not parsed in any way; it is stored and used in dynamic fields exactly as you send it. | [optional] 

## Methods

### NewUpdateCustomFieldValueRequest

`func NewUpdateCustomFieldValueRequest() *UpdateCustomFieldValueRequest`

NewUpdateCustomFieldValueRequest instantiates a new UpdateCustomFieldValueRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCustomFieldValueRequestWithDefaults

`func NewUpdateCustomFieldValueRequestWithDefaults() *UpdateCustomFieldValueRequest`

NewUpdateCustomFieldValueRequestWithDefaults instantiates a new UpdateCustomFieldValueRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContactId

`func (o *UpdateCustomFieldValueRequest) GetContactId() int32`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *UpdateCustomFieldValueRequest) GetContactIdOk() (*int32, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *UpdateCustomFieldValueRequest) SetContactId(v int32)`

SetContactId sets ContactId field to given value.

### HasContactId

`func (o *UpdateCustomFieldValueRequest) HasContactId() bool`

HasContactId returns a boolean if a field has been set.

### GetValue

`func (o *UpdateCustomFieldValueRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UpdateCustomFieldValueRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UpdateCustomFieldValueRequest) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *UpdateCustomFieldValueRequest) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


