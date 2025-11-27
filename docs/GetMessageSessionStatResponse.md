# GetMessageSessionStatResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Failed** | **int32** | Amount of failed messages. | 
**Delivered** | **int32** | Amount of delivered messages. | 
**Accepted** | **int32** | Amount of accepted messages. | 
**Rejected** | **int32** | Amount of rejected messages. | 
**Scheduled** | **int32** | Amount of scheduled messages. | 
**All** | **int32** | Total amount of messages. | 
**Sent** | **int32** | Amount of sent but not yet delivered messages. | 

## Methods

### NewGetMessageSessionStatResponse

`func NewGetMessageSessionStatResponse(failed int32, delivered int32, accepted int32, rejected int32, scheduled int32, all int32, sent int32, ) *GetMessageSessionStatResponse`

NewGetMessageSessionStatResponse instantiates a new GetMessageSessionStatResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMessageSessionStatResponseWithDefaults

`func NewGetMessageSessionStatResponseWithDefaults() *GetMessageSessionStatResponse`

NewGetMessageSessionStatResponseWithDefaults instantiates a new GetMessageSessionStatResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailed

`func (o *GetMessageSessionStatResponse) GetFailed() int32`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *GetMessageSessionStatResponse) GetFailedOk() (*int32, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *GetMessageSessionStatResponse) SetFailed(v int32)`

SetFailed sets Failed field to given value.


### GetDelivered

`func (o *GetMessageSessionStatResponse) GetDelivered() int32`

GetDelivered returns the Delivered field if non-nil, zero value otherwise.

### GetDeliveredOk

`func (o *GetMessageSessionStatResponse) GetDeliveredOk() (*int32, bool)`

GetDeliveredOk returns a tuple with the Delivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivered

`func (o *GetMessageSessionStatResponse) SetDelivered(v int32)`

SetDelivered sets Delivered field to given value.


### GetAccepted

`func (o *GetMessageSessionStatResponse) GetAccepted() int32`

GetAccepted returns the Accepted field if non-nil, zero value otherwise.

### GetAcceptedOk

`func (o *GetMessageSessionStatResponse) GetAcceptedOk() (*int32, bool)`

GetAcceptedOk returns a tuple with the Accepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccepted

`func (o *GetMessageSessionStatResponse) SetAccepted(v int32)`

SetAccepted sets Accepted field to given value.


### GetRejected

`func (o *GetMessageSessionStatResponse) GetRejected() int32`

GetRejected returns the Rejected field if non-nil, zero value otherwise.

### GetRejectedOk

`func (o *GetMessageSessionStatResponse) GetRejectedOk() (*int32, bool)`

GetRejectedOk returns a tuple with the Rejected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejected

`func (o *GetMessageSessionStatResponse) SetRejected(v int32)`

SetRejected sets Rejected field to given value.


### GetScheduled

`func (o *GetMessageSessionStatResponse) GetScheduled() int32`

GetScheduled returns the Scheduled field if non-nil, zero value otherwise.

### GetScheduledOk

`func (o *GetMessageSessionStatResponse) GetScheduledOk() (*int32, bool)`

GetScheduledOk returns a tuple with the Scheduled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduled

`func (o *GetMessageSessionStatResponse) SetScheduled(v int32)`

SetScheduled sets Scheduled field to given value.


### GetAll

`func (o *GetMessageSessionStatResponse) GetAll() int32`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *GetMessageSessionStatResponse) GetAllOk() (*int32, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *GetMessageSessionStatResponse) SetAll(v int32)`

SetAll sets All field to given value.


### GetSent

`func (o *GetMessageSessionStatResponse) GetSent() int32`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *GetMessageSessionStatResponse) GetSentOk() (*int32, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *GetMessageSessionStatResponse) SetSent(v int32)`

SetSent sets Sent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


