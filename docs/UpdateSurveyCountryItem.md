# UpdateSurveyCountryItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Country** | **string** | The 2-letter ISO country code. | 
**UserInboundId** | **int32** | User inbound phone ID. | 

## Methods

### NewUpdateSurveyCountryItem

`func NewUpdateSurveyCountryItem(country string, userInboundId int32, ) *UpdateSurveyCountryItem`

NewUpdateSurveyCountryItem instantiates a new UpdateSurveyCountryItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSurveyCountryItemWithDefaults

`func NewUpdateSurveyCountryItemWithDefaults() *UpdateSurveyCountryItem`

NewUpdateSurveyCountryItemWithDefaults instantiates a new UpdateSurveyCountryItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountry

`func (o *UpdateSurveyCountryItem) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *UpdateSurveyCountryItem) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *UpdateSurveyCountryItem) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetUserInboundId

`func (o *UpdateSurveyCountryItem) GetUserInboundId() int32`

GetUserInboundId returns the UserInboundId field if non-nil, zero value otherwise.

### GetUserInboundIdOk

`func (o *UpdateSurveyCountryItem) GetUserInboundIdOk() (*int32, bool)`

GetUserInboundIdOk returns a tuple with the UserInboundId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInboundId

`func (o *UpdateSurveyCountryItem) SetUserInboundId(v int32)`

SetUserInboundId sets UserInboundId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


