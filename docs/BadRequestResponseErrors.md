# BadRequestResponseErrors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Common** | Pointer to **[]string** | Array of messages with errors related to the entire request. For example, you did not specify either the **text** or the **templateId** when [sending the message](https://docs.textmagic.com/#tag/Outbound-Messages).  | [optional] 
**Fields** | Pointer to **map[string]interface{}** | Associative array. The keys are the POST/PUT parameter names and the values are arrays with error messages for these parameters.  | [optional] 

## Methods

### NewBadRequestResponseErrors

`func NewBadRequestResponseErrors() *BadRequestResponseErrors`

NewBadRequestResponseErrors instantiates a new BadRequestResponseErrors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBadRequestResponseErrorsWithDefaults

`func NewBadRequestResponseErrorsWithDefaults() *BadRequestResponseErrors`

NewBadRequestResponseErrorsWithDefaults instantiates a new BadRequestResponseErrors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommon

`func (o *BadRequestResponseErrors) GetCommon() []string`

GetCommon returns the Common field if non-nil, zero value otherwise.

### GetCommonOk

`func (o *BadRequestResponseErrors) GetCommonOk() (*[]string, bool)`

GetCommonOk returns a tuple with the Common field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommon

`func (o *BadRequestResponseErrors) SetCommon(v []string)`

SetCommon sets Common field to given value.

### HasCommon

`func (o *BadRequestResponseErrors) HasCommon() bool`

HasCommon returns a boolean if a field has been set.

### GetFields

`func (o *BadRequestResponseErrors) GetFields() map[string]interface{}`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *BadRequestResponseErrors) GetFieldsOk() (*map[string]interface{}, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *BadRequestResponseErrors) SetFields(v map[string]interface{})`

SetFields sets Fields field to given value.

### HasFields

`func (o *BadRequestResponseErrors) HasFields() bool`

HasFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


