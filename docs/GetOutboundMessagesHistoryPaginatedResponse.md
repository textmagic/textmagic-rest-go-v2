# GetOutboundMessagesHistoryPaginatedResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastId** | **NullableInt32** |  | 
**NextLastId** | **int32** |  | 
**Limit** | **int32** | The number of results per page. | 
**Resources** | [**[]MessageOut**](MessageOut.md) |  | 

## Methods

### NewGetOutboundMessagesHistoryPaginatedResponse

`func NewGetOutboundMessagesHistoryPaginatedResponse(lastId NullableInt32, nextLastId int32, limit int32, resources []MessageOut, ) *GetOutboundMessagesHistoryPaginatedResponse`

NewGetOutboundMessagesHistoryPaginatedResponse instantiates a new GetOutboundMessagesHistoryPaginatedResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOutboundMessagesHistoryPaginatedResponseWithDefaults

`func NewGetOutboundMessagesHistoryPaginatedResponseWithDefaults() *GetOutboundMessagesHistoryPaginatedResponse`

NewGetOutboundMessagesHistoryPaginatedResponseWithDefaults instantiates a new GetOutboundMessagesHistoryPaginatedResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastId

`func (o *GetOutboundMessagesHistoryPaginatedResponse) GetLastId() int32`

GetLastId returns the LastId field if non-nil, zero value otherwise.

### GetLastIdOk

`func (o *GetOutboundMessagesHistoryPaginatedResponse) GetLastIdOk() (*int32, bool)`

GetLastIdOk returns a tuple with the LastId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastId

`func (o *GetOutboundMessagesHistoryPaginatedResponse) SetLastId(v int32)`

SetLastId sets LastId field to given value.


### SetLastIdNil

`func (o *GetOutboundMessagesHistoryPaginatedResponse) SetLastIdNil(b bool)`

 SetLastIdNil sets the value for LastId to be an explicit nil

### UnsetLastId
`func (o *GetOutboundMessagesHistoryPaginatedResponse) UnsetLastId()`

UnsetLastId ensures that no value is present for LastId, not even an explicit nil
### GetNextLastId

`func (o *GetOutboundMessagesHistoryPaginatedResponse) GetNextLastId() int32`

GetNextLastId returns the NextLastId field if non-nil, zero value otherwise.

### GetNextLastIdOk

`func (o *GetOutboundMessagesHistoryPaginatedResponse) GetNextLastIdOk() (*int32, bool)`

GetNextLastIdOk returns a tuple with the NextLastId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextLastId

`func (o *GetOutboundMessagesHistoryPaginatedResponse) SetNextLastId(v int32)`

SetNextLastId sets NextLastId field to given value.


### GetLimit

`func (o *GetOutboundMessagesHistoryPaginatedResponse) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *GetOutboundMessagesHistoryPaginatedResponse) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *GetOutboundMessagesHistoryPaginatedResponse) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetResources

`func (o *GetOutboundMessagesHistoryPaginatedResponse) GetResources() []MessageOut`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *GetOutboundMessagesHistoryPaginatedResponse) GetResourcesOk() (*[]MessageOut, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *GetOutboundMessagesHistoryPaginatedResponse) SetResources(v []MessageOut)`

SetResources sets Resources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


