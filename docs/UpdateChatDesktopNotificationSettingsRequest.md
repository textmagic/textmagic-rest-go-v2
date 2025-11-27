# UpdateChatDesktopNotificationSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlaySound** | Pointer to **bool** | Enable notification sound? | [optional] 
**ShowNotifications** | Pointer to **bool** | Show desktop notifications about new messages. | [optional] 
**ShowText** | Pointer to **bool** | Incoming message text will be displayed in desktop notifications. | [optional] 
**SoundId** | Pointer to **int32** | Sound Id of a notification. | [optional] 

## Methods

### NewUpdateChatDesktopNotificationSettingsRequest

`func NewUpdateChatDesktopNotificationSettingsRequest() *UpdateChatDesktopNotificationSettingsRequest`

NewUpdateChatDesktopNotificationSettingsRequest instantiates a new UpdateChatDesktopNotificationSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateChatDesktopNotificationSettingsRequestWithDefaults

`func NewUpdateChatDesktopNotificationSettingsRequestWithDefaults() *UpdateChatDesktopNotificationSettingsRequest`

NewUpdateChatDesktopNotificationSettingsRequestWithDefaults instantiates a new UpdateChatDesktopNotificationSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlaySound

`func (o *UpdateChatDesktopNotificationSettingsRequest) GetPlaySound() bool`

GetPlaySound returns the PlaySound field if non-nil, zero value otherwise.

### GetPlaySoundOk

`func (o *UpdateChatDesktopNotificationSettingsRequest) GetPlaySoundOk() (*bool, bool)`

GetPlaySoundOk returns a tuple with the PlaySound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaySound

`func (o *UpdateChatDesktopNotificationSettingsRequest) SetPlaySound(v bool)`

SetPlaySound sets PlaySound field to given value.

### HasPlaySound

`func (o *UpdateChatDesktopNotificationSettingsRequest) HasPlaySound() bool`

HasPlaySound returns a boolean if a field has been set.

### GetShowNotifications

`func (o *UpdateChatDesktopNotificationSettingsRequest) GetShowNotifications() bool`

GetShowNotifications returns the ShowNotifications field if non-nil, zero value otherwise.

### GetShowNotificationsOk

`func (o *UpdateChatDesktopNotificationSettingsRequest) GetShowNotificationsOk() (*bool, bool)`

GetShowNotificationsOk returns a tuple with the ShowNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowNotifications

`func (o *UpdateChatDesktopNotificationSettingsRequest) SetShowNotifications(v bool)`

SetShowNotifications sets ShowNotifications field to given value.

### HasShowNotifications

`func (o *UpdateChatDesktopNotificationSettingsRequest) HasShowNotifications() bool`

HasShowNotifications returns a boolean if a field has been set.

### GetShowText

`func (o *UpdateChatDesktopNotificationSettingsRequest) GetShowText() bool`

GetShowText returns the ShowText field if non-nil, zero value otherwise.

### GetShowTextOk

`func (o *UpdateChatDesktopNotificationSettingsRequest) GetShowTextOk() (*bool, bool)`

GetShowTextOk returns a tuple with the ShowText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowText

`func (o *UpdateChatDesktopNotificationSettingsRequest) SetShowText(v bool)`

SetShowText sets ShowText field to given value.

### HasShowText

`func (o *UpdateChatDesktopNotificationSettingsRequest) HasShowText() bool`

HasShowText returns a boolean if a field has been set.

### GetSoundId

`func (o *UpdateChatDesktopNotificationSettingsRequest) GetSoundId() int32`

GetSoundId returns the SoundId field if non-nil, zero value otherwise.

### GetSoundIdOk

`func (o *UpdateChatDesktopNotificationSettingsRequest) GetSoundIdOk() (*int32, bool)`

GetSoundIdOk returns a tuple with the SoundId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoundId

`func (o *UpdateChatDesktopNotificationSettingsRequest) SetSoundId(v int32)`

SetSoundId sets SoundId field to given value.

### HasSoundId

`func (o *UpdateChatDesktopNotificationSettingsRequest) HasSoundId() bool`

HasSoundId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


