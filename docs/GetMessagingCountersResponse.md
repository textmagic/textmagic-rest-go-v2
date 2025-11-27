# GetMessagingCountersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contacts** | **int32** | Total contacts amount. | 
**Sent** | **int32** | Total sent messages amount. | 
**Received** | **int32** | Total received messages amount. | 

## Methods

### NewGetMessagingCountersResponse

`func NewGetMessagingCountersResponse(contacts int32, sent int32, received int32, ) *GetMessagingCountersResponse`

NewGetMessagingCountersResponse instantiates a new GetMessagingCountersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMessagingCountersResponseWithDefaults

`func NewGetMessagingCountersResponseWithDefaults() *GetMessagingCountersResponse`

NewGetMessagingCountersResponseWithDefaults instantiates a new GetMessagingCountersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContacts

`func (o *GetMessagingCountersResponse) GetContacts() int32`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *GetMessagingCountersResponse) GetContactsOk() (*int32, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *GetMessagingCountersResponse) SetContacts(v int32)`

SetContacts sets Contacts field to given value.


### GetSent

`func (o *GetMessagingCountersResponse) GetSent() int32`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *GetMessagingCountersResponse) GetSentOk() (*int32, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *GetMessagingCountersResponse) SetSent(v int32)`

SetSent sets Sent field to given value.


### GetReceived

`func (o *GetMessagingCountersResponse) GetReceived() int32`

GetReceived returns the Received field if non-nil, zero value otherwise.

### GetReceivedOk

`func (o *GetMessagingCountersResponse) GetReceivedOk() (*int32, bool)`

GetReceivedOk returns a tuple with the Received field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceived

`func (o *GetMessagingCountersResponse) SetReceived(v int32)`

SetReceived sets Received field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


