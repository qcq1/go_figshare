# Timeline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Posted** | **NullableString** | Posted date | 
**Submission** | **NullableString** | Submission date in curation (if curated) | 
**Revision** | **NullableString** | Revision date from curation (if curated) | 
**FirstOnline** | Pointer to **string** | Online posted date | [optional] 
**PublisherPublication** | Pointer to **string** | Publish date | [optional] 
**PublisherAcceptance** | Pointer to **string** | Date when the item was accepted for publication | [optional] 

## Methods

### NewTimeline

`func NewTimeline(posted NullableString, submission NullableString, revision NullableString, ) *Timeline`

NewTimeline instantiates a new Timeline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimelineWithDefaults

`func NewTimelineWithDefaults() *Timeline`

NewTimelineWithDefaults instantiates a new Timeline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPosted

`func (o *Timeline) GetPosted() string`

GetPosted returns the Posted field if non-nil, zero value otherwise.

### GetPostedOk

`func (o *Timeline) GetPostedOk() (*string, bool)`

GetPostedOk returns a tuple with the Posted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosted

`func (o *Timeline) SetPosted(v string)`

SetPosted sets Posted field to given value.


### SetPostedNil

`func (o *Timeline) SetPostedNil(b bool)`

 SetPostedNil sets the value for Posted to be an explicit nil

### UnsetPosted
`func (o *Timeline) UnsetPosted()`

UnsetPosted ensures that no value is present for Posted, not even an explicit nil
### GetSubmission

`func (o *Timeline) GetSubmission() string`

GetSubmission returns the Submission field if non-nil, zero value otherwise.

### GetSubmissionOk

`func (o *Timeline) GetSubmissionOk() (*string, bool)`

GetSubmissionOk returns a tuple with the Submission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmission

`func (o *Timeline) SetSubmission(v string)`

SetSubmission sets Submission field to given value.


### SetSubmissionNil

`func (o *Timeline) SetSubmissionNil(b bool)`

 SetSubmissionNil sets the value for Submission to be an explicit nil

### UnsetSubmission
`func (o *Timeline) UnsetSubmission()`

UnsetSubmission ensures that no value is present for Submission, not even an explicit nil
### GetRevision

`func (o *Timeline) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *Timeline) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *Timeline) SetRevision(v string)`

SetRevision sets Revision field to given value.


### SetRevisionNil

`func (o *Timeline) SetRevisionNil(b bool)`

 SetRevisionNil sets the value for Revision to be an explicit nil

### UnsetRevision
`func (o *Timeline) UnsetRevision()`

UnsetRevision ensures that no value is present for Revision, not even an explicit nil
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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


