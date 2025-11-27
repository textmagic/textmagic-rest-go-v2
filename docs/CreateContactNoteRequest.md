# CreateContactNoteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Note** | Pointer to **string** | Contact Note text. | [optional] 

## Methods

### NewCreateContactNoteRequest

`func NewCreateContactNoteRequest() *CreateContactNoteRequest`

NewCreateContactNoteRequest instantiates a new CreateContactNoteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateContactNoteRequestWithDefaults

`func NewCreateContactNoteRequestWithDefaults() *CreateContactNoteRequest`

NewCreateContactNoteRequestWithDefaults instantiates a new CreateContactNoteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNote

`func (o *CreateContactNoteRequest) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *CreateContactNoteRequest) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *CreateContactNoteRequest) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *CreateContactNoteRequest) HasNote() bool`

HasNote returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


