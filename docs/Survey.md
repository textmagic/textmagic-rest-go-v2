# Survey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Status** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**Receipents** | Pointer to [**[]SurveyRecipient**](SurveyRecipient.md) |  | [optional] 
**Countries** | Pointer to [**[]SurveySenderCountries**](SurveySenderCountries.md) |  | [optional] 

## Methods

### NewSurvey

`func NewSurvey(id int32, name string, status string, createdAt time.Time, updatedAt time.Time, ) *Survey`

NewSurvey instantiates a new Survey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSurveyWithDefaults

`func NewSurveyWithDefaults() *Survey`

NewSurveyWithDefaults instantiates a new Survey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Survey) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Survey) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Survey) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *Survey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Survey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Survey) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *Survey) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Survey) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Survey) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *Survey) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Survey) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Survey) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Survey) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Survey) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Survey) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetReceipents

`func (o *Survey) GetReceipents() []SurveyRecipient`

GetReceipents returns the Receipents field if non-nil, zero value otherwise.

### GetReceipentsOk

`func (o *Survey) GetReceipentsOk() (*[]SurveyRecipient, bool)`

GetReceipentsOk returns a tuple with the Receipents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceipents

`func (o *Survey) SetReceipents(v []SurveyRecipient)`

SetReceipents sets Receipents field to given value.

### HasReceipents

`func (o *Survey) HasReceipents() bool`

HasReceipents returns a boolean if a field has been set.

### GetCountries

`func (o *Survey) GetCountries() []SurveySenderCountries`

GetCountries returns the Countries field if non-nil, zero value otherwise.

### GetCountriesOk

`func (o *Survey) GetCountriesOk() (*[]SurveySenderCountries, bool)`

GetCountriesOk returns a tuple with the Countries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountries

`func (o *Survey) SetCountries(v []SurveySenderCountries)`

SetCountries sets Countries field to given value.

### HasCountries

`func (o *Survey) HasCountries() bool`

HasCountries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


