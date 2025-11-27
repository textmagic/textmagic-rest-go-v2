# ScheduleEmailCampaignRequestScheduleParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDateTime** | **time.Time** | When to start sending the campaign (ISO 8601 format). | 
**Timezone** | **string** | Timezone for the schedule (e.g., \&quot;America/New_York\&quot;). | 
**Rrule** | Pointer to **NullableString** | RFC 5545 recurrence rule for recurring campaigns. | [optional] 

## Methods

### NewScheduleEmailCampaignRequestScheduleParams

`func NewScheduleEmailCampaignRequestScheduleParams(startDateTime time.Time, timezone string, ) *ScheduleEmailCampaignRequestScheduleParams`

NewScheduleEmailCampaignRequestScheduleParams instantiates a new ScheduleEmailCampaignRequestScheduleParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleEmailCampaignRequestScheduleParamsWithDefaults

`func NewScheduleEmailCampaignRequestScheduleParamsWithDefaults() *ScheduleEmailCampaignRequestScheduleParams`

NewScheduleEmailCampaignRequestScheduleParamsWithDefaults instantiates a new ScheduleEmailCampaignRequestScheduleParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartDateTime

`func (o *ScheduleEmailCampaignRequestScheduleParams) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *ScheduleEmailCampaignRequestScheduleParams) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *ScheduleEmailCampaignRequestScheduleParams) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.


### GetTimezone

`func (o *ScheduleEmailCampaignRequestScheduleParams) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *ScheduleEmailCampaignRequestScheduleParams) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *ScheduleEmailCampaignRequestScheduleParams) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetRrule

`func (o *ScheduleEmailCampaignRequestScheduleParams) GetRrule() string`

GetRrule returns the Rrule field if non-nil, zero value otherwise.

### GetRruleOk

`func (o *ScheduleEmailCampaignRequestScheduleParams) GetRruleOk() (*string, bool)`

GetRruleOk returns a tuple with the Rrule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrule

`func (o *ScheduleEmailCampaignRequestScheduleParams) SetRrule(v string)`

SetRrule sets Rrule field to given value.

### HasRrule

`func (o *ScheduleEmailCampaignRequestScheduleParams) HasRrule() bool`

HasRrule returns a boolean if a field has been set.

### SetRruleNil

`func (o *ScheduleEmailCampaignRequestScheduleParams) SetRruleNil(b bool)`

 SetRruleNil sets the value for Rrule to be an explicit nil

### UnsetRrule
`func (o *ScheduleEmailCampaignRequestScheduleParams) UnsetRrule()`

UnsetRrule ensures that no value is present for Rrule, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


