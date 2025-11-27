# FileMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | File metadata type. | 
**Height** | **int32** |  | 
**Width** | **int32** |  | 
**Preview** | [**NullableFileMetadataPreview**](FileMetadataPreview.md) |  | 

## Methods

### NewFileMetadata

`func NewFileMetadata(type_ string, height int32, width int32, preview NullableFileMetadataPreview, ) *FileMetadata`

NewFileMetadata instantiates a new FileMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileMetadataWithDefaults

`func NewFileMetadataWithDefaults() *FileMetadata`

NewFileMetadataWithDefaults instantiates a new FileMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FileMetadata) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FileMetadata) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FileMetadata) SetType(v string)`

SetType sets Type field to given value.


### GetHeight

`func (o *FileMetadata) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *FileMetadata) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *FileMetadata) SetHeight(v int32)`

SetHeight sets Height field to given value.


### GetWidth

`func (o *FileMetadata) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *FileMetadata) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *FileMetadata) SetWidth(v int32)`

SetWidth sets Width field to given value.


### GetPreview

`func (o *FileMetadata) GetPreview() FileMetadataPreview`

GetPreview returns the Preview field if non-nil, zero value otherwise.

### GetPreviewOk

`func (o *FileMetadata) GetPreviewOk() (*FileMetadataPreview, bool)`

GetPreviewOk returns a tuple with the Preview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreview

`func (o *FileMetadata) SetPreview(v FileMetadataPreview)`

SetPreview sets Preview field to given value.


### SetPreviewNil

`func (o *FileMetadata) SetPreviewNil(b bool)`

 SetPreviewNil sets the value for Preview to be an explicit nil

### UnsetPreview
`func (o *FileMetadata) UnsetPreview()`

UnsetPreview ensures that no value is present for Preview, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


