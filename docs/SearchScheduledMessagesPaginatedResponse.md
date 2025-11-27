# SearchScheduledMessagesPaginatedResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | **int32** |  | 
**PageCount** | **int32** | The total number of pages. | 
**Limit** | **int32** | The number of results per page. | 
**Resources** | [**[]MessagesIcs**](MessagesIcs.md) |  | 

## Methods

### NewSearchScheduledMessagesPaginatedResponse

`func NewSearchScheduledMessagesPaginatedResponse(page int32, pageCount int32, limit int32, resources []MessagesIcs, ) *SearchScheduledMessagesPaginatedResponse`

NewSearchScheduledMessagesPaginatedResponse instantiates a new SearchScheduledMessagesPaginatedResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchScheduledMessagesPaginatedResponseWithDefaults

`func NewSearchScheduledMessagesPaginatedResponseWithDefaults() *SearchScheduledMessagesPaginatedResponse`

NewSearchScheduledMessagesPaginatedResponseWithDefaults instantiates a new SearchScheduledMessagesPaginatedResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *SearchScheduledMessagesPaginatedResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *SearchScheduledMessagesPaginatedResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *SearchScheduledMessagesPaginatedResponse) SetPage(v int32)`

SetPage sets Page field to given value.


### GetPageCount

`func (o *SearchScheduledMessagesPaginatedResponse) GetPageCount() int32`

GetPageCount returns the PageCount field if non-nil, zero value otherwise.

### GetPageCountOk

`func (o *SearchScheduledMessagesPaginatedResponse) GetPageCountOk() (*int32, bool)`

GetPageCountOk returns a tuple with the PageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageCount

`func (o *SearchScheduledMessagesPaginatedResponse) SetPageCount(v int32)`

SetPageCount sets PageCount field to given value.


### GetLimit

`func (o *SearchScheduledMessagesPaginatedResponse) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *SearchScheduledMessagesPaginatedResponse) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *SearchScheduledMessagesPaginatedResponse) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetResources

`func (o *SearchScheduledMessagesPaginatedResponse) GetResources() []MessagesIcs`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *SearchScheduledMessagesPaginatedResponse) GetResourcesOk() (*[]MessagesIcs, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *SearchScheduledMessagesPaginatedResponse) SetResources(v []MessagesIcs)`

SetResources sets Resources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


