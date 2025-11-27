# CustomFieldListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Custom Field ID. | 
**UserCustomFieldId** | **int32** | Old property custom Field ID. | 
**Name** | **string** | Custom Field name. | 
**Value** | **string** | Custom Field value. | 
**CreatedAt** | **time.Time** | Custom field creation time. | 

## Methods

### NewCustomFieldListItem

`func NewCustomFieldListItem(id int32, userCustomFieldId int32, name string, value string, createdAt time.Time, ) *CustomFieldListItem`

NewCustomFieldListItem instantiates a new CustomFieldListItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomFieldListItemWithDefaults

`func NewCustomFieldListItemWithDefaults() *CustomFieldListItem`

NewCustomFieldListItemWithDefaults instantiates a new CustomFieldListItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomFieldListItem) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomFieldListItem) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomFieldListItem) SetId(v int32)`

SetId sets Id field to given value.


### GetUserCustomFieldId

`func (o *CustomFieldListItem) GetUserCustomFieldId() int32`

GetUserCustomFieldId returns the UserCustomFieldId field if non-nil, zero value otherwise.

### GetUserCustomFieldIdOk

`func (o *CustomFieldListItem) GetUserCustomFieldIdOk() (*int32, bool)`

GetUserCustomFieldIdOk returns a tuple with the UserCustomFieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCustomFieldId

`func (o *CustomFieldListItem) SetUserCustomFieldId(v int32)`

SetUserCustomFieldId sets UserCustomFieldId field to given value.


### GetName

`func (o *CustomFieldListItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomFieldListItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomFieldListItem) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *CustomFieldListItem) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CustomFieldListItem) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CustomFieldListItem) SetValue(v string)`

SetValue sets Value field to given value.


### GetCreatedAt

`func (o *CustomFieldListItem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CustomFieldListItem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CustomFieldListItem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


