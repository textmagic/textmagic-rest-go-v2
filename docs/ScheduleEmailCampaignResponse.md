# ScheduleEmailCampaignResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Campaign** | [**ScheduledEmailCampaignDetails**](ScheduledEmailCampaignDetails.md) |  | 
**Cost** | **float32** | Estimated cost for sending this campaign. | 

## Methods

### NewScheduleEmailCampaignResponse

`func NewScheduleEmailCampaignResponse(campaign ScheduledEmailCampaignDetails, cost float32, ) *ScheduleEmailCampaignResponse`

NewScheduleEmailCampaignResponse instantiates a new ScheduleEmailCampaignResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleEmailCampaignResponseWithDefaults

`func NewScheduleEmailCampaignResponseWithDefaults() *ScheduleEmailCampaignResponse`

NewScheduleEmailCampaignResponseWithDefaults instantiates a new ScheduleEmailCampaignResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCampaign

`func (o *ScheduleEmailCampaignResponse) GetCampaign() ScheduledEmailCampaignDetails`

GetCampaign returns the Campaign field if non-nil, zero value otherwise.

### GetCampaignOk

`func (o *ScheduleEmailCampaignResponse) GetCampaignOk() (*ScheduledEmailCampaignDetails, bool)`

GetCampaignOk returns a tuple with the Campaign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaign

`func (o *ScheduleEmailCampaignResponse) SetCampaign(v ScheduledEmailCampaignDetails)`

SetCampaign sets Campaign field to given value.


### GetCost

`func (o *ScheduleEmailCampaignResponse) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *ScheduleEmailCampaignResponse) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *ScheduleEmailCampaignResponse) SetCost(v float32)`

SetCost sets Cost field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


