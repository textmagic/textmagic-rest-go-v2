# SurveySenderCountries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**CountryName** | **string** |  | 
**FromNumber** | **string** |  | 
**AllowDedicated** | **bool** | Is allowed to use a dedicated number? | 

## Methods

### NewSurveySenderCountries

`func NewSurveySenderCountries(id int32, countryName string, fromNumber string, allowDedicated bool, ) *SurveySenderCountries`

NewSurveySenderCountries instantiates a new SurveySenderCountries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSurveySenderCountriesWithDefaults

`func NewSurveySenderCountriesWithDefaults() *SurveySenderCountries`

NewSurveySenderCountriesWithDefaults instantiates a new SurveySenderCountries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SurveySenderCountries) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SurveySenderCountries) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SurveySenderCountries) SetId(v int32)`

SetId sets Id field to given value.


### GetCountryName

`func (o *SurveySenderCountries) GetCountryName() string`

GetCountryName returns the CountryName field if non-nil, zero value otherwise.

### GetCountryNameOk

`func (o *SurveySenderCountries) GetCountryNameOk() (*string, bool)`

GetCountryNameOk returns a tuple with the CountryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryName

`func (o *SurveySenderCountries) SetCountryName(v string)`

SetCountryName sets CountryName field to given value.


### GetFromNumber

`func (o *SurveySenderCountries) GetFromNumber() string`

GetFromNumber returns the FromNumber field if non-nil, zero value otherwise.

### GetFromNumberOk

`func (o *SurveySenderCountries) GetFromNumberOk() (*string, bool)`

GetFromNumberOk returns a tuple with the FromNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromNumber

`func (o *SurveySenderCountries) SetFromNumber(v string)`

SetFromNumber sets FromNumber field to given value.


### GetAllowDedicated

`func (o *SurveySenderCountries) GetAllowDedicated() bool`

GetAllowDedicated returns the AllowDedicated field if non-nil, zero value otherwise.

### GetAllowDedicatedOk

`func (o *SurveySenderCountries) GetAllowDedicatedOk() (*bool, bool)`

GetAllowDedicatedOk returns a tuple with the AllowDedicated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowDedicated

`func (o *SurveySenderCountries) SetAllowDedicated(v bool)`

SetAllowDedicated sets AllowDedicated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


