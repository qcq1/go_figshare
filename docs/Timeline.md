# Timeline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstOnline** | Pointer to **string** | Online posted date | [optional] 
**PublisherAcceptance** | Pointer to **string** | Date when the item was accepted for publication | [optional] 
**PublisherPublication** | Pointer to **string** | Publish date | [optional] 

## Methods

### NewTimeline

`func NewTimeline() *Timeline`

NewTimeline instantiates a new Timeline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimelineWithDefaults

`func NewTimelineWithDefaults() *Timeline`

NewTimelineWithDefaults instantiates a new Timeline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstOnline

`func (o *Timeline) GetFirstOnline() string`

GetFirstOnline returns the FirstOnline field if non-nil, zero value otherwise.

### GetFirstOnlineOk

`func (o *Timeline) GetFirstOnlineOk() (*string, bool)`

GetFirstOnlineOk returns a tuple with the FirstOnline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstOnline

`func (o *Timeline) SetFirstOnline(v string)`

SetFirstOnline sets FirstOnline field to given value.

### HasFirstOnline

`func (o *Timeline) HasFirstOnline() bool`

HasFirstOnline returns a boolean if a field has been set.

### GetPublisherAcceptance

`func (o *Timeline) GetPublisherAcceptance() string`

GetPublisherAcceptance returns the PublisherAcceptance field if non-nil, zero value otherwise.

### GetPublisherAcceptanceOk

`func (o *Timeline) GetPublisherAcceptanceOk() (*string, bool)`

GetPublisherAcceptanceOk returns a tuple with the PublisherAcceptance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisherAcceptance

`func (o *Timeline) SetPublisherAcceptance(v string)`

SetPublisherAcceptance sets PublisherAcceptance field to given value.

### HasPublisherAcceptance

`func (o *Timeline) HasPublisherAcceptance() bool`

HasPublisherAcceptance returns a boolean if a field has been set.

### GetPublisherPublication

`func (o *Timeline) GetPublisherPublication() string`

GetPublisherPublication returns the PublisherPublication field if non-nil, zero value otherwise.

### GetPublisherPublicationOk

`func (o *Timeline) GetPublisherPublicationOk() (*string, bool)`

GetPublisherPublicationOk returns a tuple with the PublisherPublication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisherPublication

`func (o *Timeline) SetPublisherPublication(v string)`

SetPublisherPublication sets PublisherPublication field to given value.

### HasPublisherPublication

`func (o *Timeline) HasPublisherPublication() bool`

HasPublisherPublication returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


