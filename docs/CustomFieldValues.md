# CustomFieldValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldId** | Pointer to **int32** |  | [optional] 
**FieldTitle** | Pointer to **string** |  | [optional] 
**FieldType** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCustomFieldValues

`func NewCustomFieldValues() *CustomFieldValues`

NewCustomFieldValues instantiates a new CustomFieldValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomFieldValuesWithDefaults

`func NewCustomFieldValuesWithDefaults() *CustomFieldValues`

NewCustomFieldValuesWithDefaults instantiates a new CustomFieldValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldId

`func (o *CustomFieldValues) GetFieldId() int32`

GetFieldId returns the FieldId field if non-nil, zero value otherwise.

### GetFieldIdOk

`func (o *CustomFieldValues) GetFieldIdOk() (*int32, bool)`

GetFieldIdOk returns a tuple with the FieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldId

`func (o *CustomFieldValues) SetFieldId(v int32)`

SetFieldId sets FieldId field to given value.

### HasFieldId

`func (o *CustomFieldValues) HasFieldId() bool`

HasFieldId returns a boolean if a field has been set.

### GetFieldTitle

`func (o *CustomFieldValues) GetFieldTitle() string`

GetFieldTitle returns the FieldTitle field if non-nil, zero value otherwise.

### GetFieldTitleOk

`func (o *CustomFieldValues) GetFieldTitleOk() (*string, bool)`

GetFieldTitleOk returns a tuple with the FieldTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldTitle

`func (o *CustomFieldValues) SetFieldTitle(v string)`

SetFieldTitle sets FieldTitle field to given value.

### HasFieldTitle

`func (o *CustomFieldValues) HasFieldTitle() bool`

HasFieldTitle returns a boolean if a field has been set.

### GetFieldType

`func (o *CustomFieldValues) GetFieldType() string`

GetFieldType returns the FieldType field if non-nil, zero value otherwise.

### GetFieldTypeOk

`func (o *CustomFieldValues) GetFieldTypeOk() (*string, bool)`

GetFieldTypeOk returns a tuple with the FieldType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldType

`func (o *CustomFieldValues) SetFieldType(v string)`

SetFieldType sets FieldType field to given value.

### HasFieldType

`func (o *CustomFieldValues) HasFieldType() bool`

HasFieldType returns a boolean if a field has been set.

### GetValue

`func (o *CustomFieldValues) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CustomFieldValues) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CustomFieldValues) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *CustomFieldValues) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


