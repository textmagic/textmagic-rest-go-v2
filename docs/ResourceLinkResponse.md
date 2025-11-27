# ResourceLinkResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Resource ID. | 
**Href** | **string** | A link to this resource. If you want to fetch it, just **GET** this address. | 

## Methods

### NewResourceLinkResponse

`func NewResourceLinkResponse(id int32, href string, ) *ResourceLinkResponse`

NewResourceLinkResponse instantiates a new ResourceLinkResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceLinkResponseWithDefaults

`func NewResourceLinkResponseWithDefaults() *ResourceLinkResponse`

NewResourceLinkResponseWithDefaults instantiates a new ResourceLinkResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResourceLinkResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResourceLinkResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResourceLinkResponse) SetId(v int32)`

SetId sets Id field to given value.


### GetHref

`func (o *ResourceLinkResponse) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ResourceLinkResponse) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ResourceLinkResponse) SetHref(v string)`

SetHref sets Href field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


