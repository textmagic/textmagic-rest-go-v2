# Timezone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Internal timezone ID. | 
**Area** | **string** | Timezone area. | 
**Dst** | **int32** | Is daylight saving time used in this timezone? | 
**Offset** | **int32** | Offset from UTC time in seconds. In this example, it is 21600/60/60&#x3D;6 hours. | 
**Timezone** | **string** | User-friendly timezone name (with spaces replaced by underscores). | 

## Methods

### NewTimezone

`func NewTimezone(id int32, area string, dst int32, offset int32, timezone string, ) *Timezone`

NewTimezone instantiates a new Timezone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimezoneWithDefaults

`func NewTimezoneWithDefaults() *Timezone`

NewTimezoneWithDefaults instantiates a new Timezone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Timezone) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Timezone) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Timezone) SetId(v int32)`

SetId sets Id field to given value.


### GetArea

`func (o *Timezone) GetArea() string`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *Timezone) GetAreaOk() (*string, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *Timezone) SetArea(v string)`

SetArea sets Area field to given value.


### GetDst

`func (o *Timezone) GetDst() int32`

GetDst returns the Dst field if non-nil, zero value otherwise.

### GetDstOk

`func (o *Timezone) GetDstOk() (*int32, bool)`

GetDstOk returns a tuple with the Dst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDst

`func (o *Timezone) SetDst(v int32)`

SetDst sets Dst field to given value.


### GetOffset

`func (o *Timezone) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *Timezone) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *Timezone) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetTimezone

`func (o *Timezone) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *Timezone) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *Timezone) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


