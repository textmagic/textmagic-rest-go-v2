# MessagesIcs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Schedule ID. | 
**NextSend** | **time.Time** | The next send date in [ISO 8601](https://en.wikipedia.org/?title&#x3D;ISO_8601) format.  | 
**Rrule** | **NullableString** | [iCal RRULE](http://www.kanzaki.com/docs/ical/rrule.html) string.  | 
**Session** | [**NullableMessageSession**](MessageSession.md) |  | 
**LastSent** | **NullableTime** | The date and time when the last message was sent. | 
**ContactName** | **string** | Aggregated contact information. If the message was scheduled to be sent to a single contact, a full name will be returned here. Otherwise, a total amount of contacts will be returned. | 
**Parameters** | [**MessagesIcsParameters**](MessagesIcsParameters.md) |  | 
**Type** | **string** |  | 
**Summary** | **string** | A human-readable summary of the sending schedule. | 
**TextParameters** | [**MessagesIcsTextParameters**](MessagesIcsTextParameters.md) |  | 
**FirstOccurrence** | **NullableTime** | First occurence date. | 
**LastOccurrence** | **NullableTime** | Last occurence date (could be &#x60;null&#x60; if the schedule is endless). | 
**RecipientsCount** | **NullableInt32** | Amount of actual recipients. | 
**Timezone** | **string** | User-friendly timezone name (with spaces replaced by underscores). | 
**Completed** | **bool** | Indicates that scheduling has been completed. | 
**Avatar** | **NullableString** | A relative link to the contact avatar. | 
**CreatedAt** | **time.Time** | Scheduling creation time. | 

## Methods

### NewMessagesIcs

`func NewMessagesIcs(id int32, nextSend time.Time, rrule NullableString, session NullableMessageSession, lastSent NullableTime, contactName string, parameters MessagesIcsParameters, type_ string, summary string, textParameters MessagesIcsTextParameters, firstOccurrence NullableTime, lastOccurrence NullableTime, recipientsCount NullableInt32, timezone string, completed bool, avatar NullableString, createdAt time.Time, ) *MessagesIcs`

NewMessagesIcs instantiates a new MessagesIcs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessagesIcsWithDefaults

`func NewMessagesIcsWithDefaults() *MessagesIcs`

NewMessagesIcsWithDefaults instantiates a new MessagesIcs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MessagesIcs) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MessagesIcs) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MessagesIcs) SetId(v int32)`

SetId sets Id field to given value.


### GetNextSend

`func (o *MessagesIcs) GetNextSend() time.Time`

GetNextSend returns the NextSend field if non-nil, zero value otherwise.

### GetNextSendOk

`func (o *MessagesIcs) GetNextSendOk() (*time.Time, bool)`

GetNextSendOk returns a tuple with the NextSend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextSend

`func (o *MessagesIcs) SetNextSend(v time.Time)`

SetNextSend sets NextSend field to given value.


### GetRrule

`func (o *MessagesIcs) GetRrule() string`

GetRrule returns the Rrule field if non-nil, zero value otherwise.

### GetRruleOk

`func (o *MessagesIcs) GetRruleOk() (*string, bool)`

GetRruleOk returns a tuple with the Rrule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrule

`func (o *MessagesIcs) SetRrule(v string)`

SetRrule sets Rrule field to given value.


### SetRruleNil

`func (o *MessagesIcs) SetRruleNil(b bool)`

 SetRruleNil sets the value for Rrule to be an explicit nil

### UnsetRrule
`func (o *MessagesIcs) UnsetRrule()`

UnsetRrule ensures that no value is present for Rrule, not even an explicit nil
### GetSession

`func (o *MessagesIcs) GetSession() MessageSession`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *MessagesIcs) GetSessionOk() (*MessageSession, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *MessagesIcs) SetSession(v MessageSession)`

SetSession sets Session field to given value.


### SetSessionNil

`func (o *MessagesIcs) SetSessionNil(b bool)`

 SetSessionNil sets the value for Session to be an explicit nil

### UnsetSession
`func (o *MessagesIcs) UnsetSession()`

UnsetSession ensures that no value is present for Session, not even an explicit nil
### GetLastSent

`func (o *MessagesIcs) GetLastSent() time.Time`

GetLastSent returns the LastSent field if non-nil, zero value otherwise.

### GetLastSentOk

`func (o *MessagesIcs) GetLastSentOk() (*time.Time, bool)`

GetLastSentOk returns a tuple with the LastSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSent

`func (o *MessagesIcs) SetLastSent(v time.Time)`

SetLastSent sets LastSent field to given value.


### SetLastSentNil

`func (o *MessagesIcs) SetLastSentNil(b bool)`

 SetLastSentNil sets the value for LastSent to be an explicit nil

### UnsetLastSent
`func (o *MessagesIcs) UnsetLastSent()`

UnsetLastSent ensures that no value is present for LastSent, not even an explicit nil
### GetContactName

`func (o *MessagesIcs) GetContactName() string`

GetContactName returns the ContactName field if non-nil, zero value otherwise.

### GetContactNameOk

`func (o *MessagesIcs) GetContactNameOk() (*string, bool)`

GetContactNameOk returns a tuple with the ContactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactName

`func (o *MessagesIcs) SetContactName(v string)`

SetContactName sets ContactName field to given value.


### GetParameters

`func (o *MessagesIcs) GetParameters() MessagesIcsParameters`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *MessagesIcs) GetParametersOk() (*MessagesIcsParameters, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *MessagesIcs) SetParameters(v MessagesIcsParameters)`

SetParameters sets Parameters field to given value.


### GetType

`func (o *MessagesIcs) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MessagesIcs) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MessagesIcs) SetType(v string)`

SetType sets Type field to given value.


### GetSummary

`func (o *MessagesIcs) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *MessagesIcs) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *MessagesIcs) SetSummary(v string)`

SetSummary sets Summary field to given value.


### GetTextParameters

`func (o *MessagesIcs) GetTextParameters() MessagesIcsTextParameters`

GetTextParameters returns the TextParameters field if non-nil, zero value otherwise.

### GetTextParametersOk

`func (o *MessagesIcs) GetTextParametersOk() (*MessagesIcsTextParameters, bool)`

GetTextParametersOk returns a tuple with the TextParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextParameters

`func (o *MessagesIcs) SetTextParameters(v MessagesIcsTextParameters)`

SetTextParameters sets TextParameters field to given value.


### GetFirstOccurrence

`func (o *MessagesIcs) GetFirstOccurrence() time.Time`

GetFirstOccurrence returns the FirstOccurrence field if non-nil, zero value otherwise.

### GetFirstOccurrenceOk

`func (o *MessagesIcs) GetFirstOccurrenceOk() (*time.Time, bool)`

GetFirstOccurrenceOk returns a tuple with the FirstOccurrence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstOccurrence

`func (o *MessagesIcs) SetFirstOccurrence(v time.Time)`

SetFirstOccurrence sets FirstOccurrence field to given value.


### SetFirstOccurrenceNil

`func (o *MessagesIcs) SetFirstOccurrenceNil(b bool)`

 SetFirstOccurrenceNil sets the value for FirstOccurrence to be an explicit nil

### UnsetFirstOccurrence
`func (o *MessagesIcs) UnsetFirstOccurrence()`

UnsetFirstOccurrence ensures that no value is present for FirstOccurrence, not even an explicit nil
### GetLastOccurrence

`func (o *MessagesIcs) GetLastOccurrence() time.Time`

GetLastOccurrence returns the LastOccurrence field if non-nil, zero value otherwise.

### GetLastOccurrenceOk

`func (o *MessagesIcs) GetLastOccurrenceOk() (*time.Time, bool)`

GetLastOccurrenceOk returns a tuple with the LastOccurrence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastOccurrence

`func (o *MessagesIcs) SetLastOccurrence(v time.Time)`

SetLastOccurrence sets LastOccurrence field to given value.


### SetLastOccurrenceNil

`func (o *MessagesIcs) SetLastOccurrenceNil(b bool)`

 SetLastOccurrenceNil sets the value for LastOccurrence to be an explicit nil

### UnsetLastOccurrence
`func (o *MessagesIcs) UnsetLastOccurrence()`

UnsetLastOccurrence ensures that no value is present for LastOccurrence, not even an explicit nil
### GetRecipientsCount

`func (o *MessagesIcs) GetRecipientsCount() int32`

GetRecipientsCount returns the RecipientsCount field if non-nil, zero value otherwise.

### GetRecipientsCountOk

`func (o *MessagesIcs) GetRecipientsCountOk() (*int32, bool)`

GetRecipientsCountOk returns a tuple with the RecipientsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientsCount

`func (o *MessagesIcs) SetRecipientsCount(v int32)`

SetRecipientsCount sets RecipientsCount field to given value.


### SetRecipientsCountNil

`func (o *MessagesIcs) SetRecipientsCountNil(b bool)`

 SetRecipientsCountNil sets the value for RecipientsCount to be an explicit nil

### UnsetRecipientsCount
`func (o *MessagesIcs) UnsetRecipientsCount()`

UnsetRecipientsCount ensures that no value is present for RecipientsCount, not even an explicit nil
### GetTimezone

`func (o *MessagesIcs) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *MessagesIcs) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *MessagesIcs) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetCompleted

`func (o *MessagesIcs) GetCompleted() bool`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *MessagesIcs) GetCompletedOk() (*bool, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *MessagesIcs) SetCompleted(v bool)`

SetCompleted sets Completed field to given value.


### GetAvatar

`func (o *MessagesIcs) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *MessagesIcs) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *MessagesIcs) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.


### SetAvatarNil

`func (o *MessagesIcs) SetAvatarNil(b bool)`

 SetAvatarNil sets the value for Avatar to be an explicit nil

### UnsetAvatar
`func (o *MessagesIcs) UnsetAvatar()`

UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil
### GetCreatedAt

`func (o *MessagesIcs) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MessagesIcs) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MessagesIcs) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


