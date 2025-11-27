# UploadMessageAttachmentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chars** | **int32** | &#x60;href&#x60; field characters count.  | 
**Href** | **string** | This is a relative link to your file. To construct a full link, just add “[https://my.textmagic.com/”](https://my.textmagic.com/%E2%80%9D) to the beginning (like this: [https://my.textmagic.com/click/Zwcj9](https://my.textmagic.com/click/Zwcj9)). For most modern devices, you can omit the “https://” part and write just [my.textmagic.com/click/Zwcj9](https://my.textmagic.com/click/Zwcj9), which will save you 8 characters.  | 
**Name** | **string** | File name of the uploaded file.  | 
**Size** | **int32** | Attachment size in bytes. | 
**Resource** | **string** | Internal file name | 

## Methods

### NewUploadMessageAttachmentResponse

`func NewUploadMessageAttachmentResponse(chars int32, href string, name string, size int32, resource string, ) *UploadMessageAttachmentResponse`

NewUploadMessageAttachmentResponse instantiates a new UploadMessageAttachmentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadMessageAttachmentResponseWithDefaults

`func NewUploadMessageAttachmentResponseWithDefaults() *UploadMessageAttachmentResponse`

NewUploadMessageAttachmentResponseWithDefaults instantiates a new UploadMessageAttachmentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChars

`func (o *UploadMessageAttachmentResponse) GetChars() int32`

GetChars returns the Chars field if non-nil, zero value otherwise.

### GetCharsOk

`func (o *UploadMessageAttachmentResponse) GetCharsOk() (*int32, bool)`

GetCharsOk returns a tuple with the Chars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChars

`func (o *UploadMessageAttachmentResponse) SetChars(v int32)`

SetChars sets Chars field to given value.


### GetHref

`func (o *UploadMessageAttachmentResponse) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *UploadMessageAttachmentResponse) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *UploadMessageAttachmentResponse) SetHref(v string)`

SetHref sets Href field to given value.


### GetName

`func (o *UploadMessageAttachmentResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UploadMessageAttachmentResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UploadMessageAttachmentResponse) SetName(v string)`

SetName sets Name field to given value.


### GetSize

`func (o *UploadMessageAttachmentResponse) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *UploadMessageAttachmentResponse) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *UploadMessageAttachmentResponse) SetSize(v int32)`

SetSize sets Size field to given value.


### GetResource

`func (o *UploadMessageAttachmentResponse) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *UploadMessageAttachmentResponse) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *UploadMessageAttachmentResponse) SetResource(v string)`

SetResource sets Resource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


