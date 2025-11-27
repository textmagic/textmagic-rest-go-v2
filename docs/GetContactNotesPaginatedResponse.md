# GetContactNotesPaginatedResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | **int32** |  | 
**PageCount** | **int32** | The total number of pages. | 
**Limit** | **int32** | The number of results per page. | 
**Resources** | [**[]ContactNote**](ContactNote.md) |  | 

## Methods

### NewGetContactNotesPaginatedResponse

`func NewGetContactNotesPaginatedResponse(page int32, pageCount int32, limit int32, resources []ContactNote, ) *GetContactNotesPaginatedResponse`

NewGetContactNotesPaginatedResponse instantiates a new GetContactNotesPaginatedResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetContactNotesPaginatedResponseWithDefaults

`func NewGetContactNotesPaginatedResponseWithDefaults() *GetContactNotesPaginatedResponse`

NewGetContactNotesPaginatedResponseWithDefaults instantiates a new GetContactNotesPaginatedResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *GetContactNotesPaginatedResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *GetContactNotesPaginatedResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *GetContactNotesPaginatedResponse) SetPage(v int32)`

SetPage sets Page field to given value.


### GetPageCount

`func (o *GetContactNotesPaginatedResponse) GetPageCount() int32`

GetPageCount returns the PageCount field if non-nil, zero value otherwise.

### GetPageCountOk

`func (o *GetContactNotesPaginatedResponse) GetPageCountOk() (*int32, bool)`

GetPageCountOk returns a tuple with the PageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageCount

`func (o *GetContactNotesPaginatedResponse) SetPageCount(v int32)`

SetPageCount sets PageCount field to given value.


### GetLimit

`func (o *GetContactNotesPaginatedResponse) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *GetContactNotesPaginatedResponse) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *GetContactNotesPaginatedResponse) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetResources

`func (o *GetContactNotesPaginatedResponse) GetResources() []ContactNote`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *GetContactNotesPaginatedResponse) GetResourcesOk() (*[]ContactNote, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *GetContactNotesPaginatedResponse) SetResources(v []ContactNote)`

SetResources sets Resources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


