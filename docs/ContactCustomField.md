# ContactCustomField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** |  | [optional] 
**UserCustomField** | Pointer to [**UserCustomField**](UserCustomField.md) |  | [optional] 

## Methods

### NewContactCustomField

`func NewContactCustomField() *ContactCustomField`

NewContactCustomField instantiates a new ContactCustomField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactCustomFieldWithDefaults

`func NewContactCustomFieldWithDefaults() *ContactCustomField`

NewContactCustomFieldWithDefaults instantiates a new ContactCustomField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *ContactCustomField) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ContactCustomField) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ContactCustomField) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ContactCustomField) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetUserCustomField

`func (o *ContactCustomField) GetUserCustomField() UserCustomField`

GetUserCustomField returns the UserCustomField field if non-nil, zero value otherwise.

### GetUserCustomFieldOk

`func (o *ContactCustomField) GetUserCustomFieldOk() (*UserCustomField, bool)`

GetUserCustomFieldOk returns a tuple with the UserCustomField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCustomField

`func (o *ContactCustomField) SetUserCustomField(v UserCustomField)`

SetUserCustomField sets UserCustomField field to given value.

### HasUserCustomField

`func (o *ContactCustomField) HasUserCustomField() bool`

HasUserCustomField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


