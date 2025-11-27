# SurveyNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Label** | **string** |  | 
**Body** | **string** |  | 
**NodeType** | **string** |  | 
**IsEndNode** | **bool** |  | 
**SendDelay** | **int32** |  | 
**StartNodes** | **[]string** |  | 
**EndNodes** | **[]string** |  | 

## Methods

### NewSurveyNode

`func NewSurveyNode(id int32, label string, body string, nodeType string, isEndNode bool, sendDelay int32, startNodes []string, endNodes []string, ) *SurveyNode`

NewSurveyNode instantiates a new SurveyNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSurveyNodeWithDefaults

`func NewSurveyNodeWithDefaults() *SurveyNode`

NewSurveyNodeWithDefaults instantiates a new SurveyNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SurveyNode) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SurveyNode) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SurveyNode) SetId(v int32)`

SetId sets Id field to given value.


### GetLabel

`func (o *SurveyNode) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *SurveyNode) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *SurveyNode) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetBody

`func (o *SurveyNode) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *SurveyNode) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *SurveyNode) SetBody(v string)`

SetBody sets Body field to given value.


### GetNodeType

`func (o *SurveyNode) GetNodeType() string`

GetNodeType returns the NodeType field if non-nil, zero value otherwise.

### GetNodeTypeOk

`func (o *SurveyNode) GetNodeTypeOk() (*string, bool)`

GetNodeTypeOk returns a tuple with the NodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeType

`func (o *SurveyNode) SetNodeType(v string)`

SetNodeType sets NodeType field to given value.


### GetIsEndNode

`func (o *SurveyNode) GetIsEndNode() bool`

GetIsEndNode returns the IsEndNode field if non-nil, zero value otherwise.

### GetIsEndNodeOk

`func (o *SurveyNode) GetIsEndNodeOk() (*bool, bool)`

GetIsEndNodeOk returns a tuple with the IsEndNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEndNode

`func (o *SurveyNode) SetIsEndNode(v bool)`

SetIsEndNode sets IsEndNode field to given value.


### GetSendDelay

`func (o *SurveyNode) GetSendDelay() int32`

GetSendDelay returns the SendDelay field if non-nil, zero value otherwise.

### GetSendDelayOk

`func (o *SurveyNode) GetSendDelayOk() (*int32, bool)`

GetSendDelayOk returns a tuple with the SendDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendDelay

`func (o *SurveyNode) SetSendDelay(v int32)`

SetSendDelay sets SendDelay field to given value.


### GetStartNodes

`func (o *SurveyNode) GetStartNodes() []string`

GetStartNodes returns the StartNodes field if non-nil, zero value otherwise.

### GetStartNodesOk

`func (o *SurveyNode) GetStartNodesOk() (*[]string, bool)`

GetStartNodesOk returns a tuple with the StartNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartNodes

`func (o *SurveyNode) SetStartNodes(v []string)`

SetStartNodes sets StartNodes field to given value.


### SetStartNodesNil

`func (o *SurveyNode) SetStartNodesNil(b bool)`

 SetStartNodesNil sets the value for StartNodes to be an explicit nil

### UnsetStartNodes
`func (o *SurveyNode) UnsetStartNodes()`

UnsetStartNodes ensures that no value is present for StartNodes, not even an explicit nil
### GetEndNodes

`func (o *SurveyNode) GetEndNodes() []string`

GetEndNodes returns the EndNodes field if non-nil, zero value otherwise.

### GetEndNodesOk

`func (o *SurveyNode) GetEndNodesOk() (*[]string, bool)`

GetEndNodesOk returns a tuple with the EndNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndNodes

`func (o *SurveyNode) SetEndNodes(v []string)`

SetEndNodes sets EndNodes field to given value.


### SetEndNodesNil

`func (o *SurveyNode) SetEndNodesNil(b bool)`

 SetEndNodesNil sets the value for EndNodes to be an explicit nil

### UnsetEndNodes
`func (o *SurveyNode) UnsetEndNodes()`

UnsetEndNodes ensures that no value is present for EndNodes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


