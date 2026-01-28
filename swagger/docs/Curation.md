# Curation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **int64** | The ID of the account of the owner of the article of this review. | 
**ArticleId** | **int64** | The ID of the article of this review. | 
**AssignedTo** | **int64** | The ID of the account to which this review is assigned. | 
**CommentsCount** | **int64** | The number of comments in the review. | 
**CreatedDate** | **string** | The creation date of the review. | 
**GroupId** | **int64** | The group in which the article is present. | 
**Id** | **int64** | The review id | 
**ModifiedDate** | **string** | The date the review has been modified. | 
**RequestNumber** | **int64** | The request number of the review. | 
**ResolutionComment** | **string** | The resolution comment of the review. | 
**Status** | **string** | The status of the review. | 
**Version** | **int64** | The Version number of the article in review. | 

## Methods

### NewCuration

`func NewCuration(accountId int64, articleId int64, assignedTo int64, commentsCount int64, createdDate string, groupId int64, id int64, modifiedDate string, requestNumber int64, resolutionComment string, status string, version int64, ) *Curation`

NewCuration instantiates a new Curation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurationWithDefaults

`func NewCurationWithDefaults() *Curation`

NewCurationWithDefaults instantiates a new Curation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *Curation) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Curation) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Curation) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### GetArticleId

`func (o *Curation) GetArticleId() int64`

GetArticleId returns the ArticleId field if non-nil, zero value otherwise.

### GetArticleIdOk

`func (o *Curation) GetArticleIdOk() (*int64, bool)`

GetArticleIdOk returns a tuple with the ArticleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArticleId

`func (o *Curation) SetArticleId(v int64)`

SetArticleId sets ArticleId field to given value.


### GetAssignedTo

`func (o *Curation) GetAssignedTo() int64`

GetAssignedTo returns the AssignedTo field if non-nil, zero value otherwise.

### GetAssignedToOk

`func (o *Curation) GetAssignedToOk() (*int64, bool)`

GetAssignedToOk returns a tuple with the AssignedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedTo

`func (o *Curation) SetAssignedTo(v int64)`

SetAssignedTo sets AssignedTo field to given value.


### GetCommentsCount

`func (o *Curation) GetCommentsCount() int64`

GetCommentsCount returns the CommentsCount field if non-nil, zero value otherwise.

### GetCommentsCountOk

`func (o *Curation) GetCommentsCountOk() (*int64, bool)`

GetCommentsCountOk returns a tuple with the CommentsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentsCount

`func (o *Curation) SetCommentsCount(v int64)`

SetCommentsCount sets CommentsCount field to given value.


### GetCreatedDate

`func (o *Curation) GetCreatedDate() string`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *Curation) GetCreatedDateOk() (*string, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *Curation) SetCreatedDate(v string)`

SetCreatedDate sets CreatedDate field to given value.


### GetGroupId

`func (o *Curation) GetGroupId() int64`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *Curation) GetGroupIdOk() (*int64, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *Curation) SetGroupId(v int64)`

SetGroupId sets GroupId field to given value.


### GetId

`func (o *Curation) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Curation) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Curation) SetId(v int64)`

SetId sets Id field to given value.


### GetModifiedDate

`func (o *Curation) GetModifiedDate() string`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *Curation) GetModifiedDateOk() (*string, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *Curation) SetModifiedDate(v string)`

SetModifiedDate sets ModifiedDate field to given value.


### GetRequestNumber

`func (o *Curation) GetRequestNumber() int64`

GetRequestNumber returns the RequestNumber field if non-nil, zero value otherwise.

### GetRequestNumberOk

`func (o *Curation) GetRequestNumberOk() (*int64, bool)`

GetRequestNumberOk returns a tuple with the RequestNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestNumber

`func (o *Curation) SetRequestNumber(v int64)`

SetRequestNumber sets RequestNumber field to given value.


### GetResolutionComment

`func (o *Curation) GetResolutionComment() string`

GetResolutionComment returns the ResolutionComment field if non-nil, zero value otherwise.

### GetResolutionCommentOk

`func (o *Curation) GetResolutionCommentOk() (*string, bool)`

GetResolutionCommentOk returns a tuple with the ResolutionComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionComment

`func (o *Curation) SetResolutionComment(v string)`

SetResolutionComment sets ResolutionComment field to given value.


### GetStatus

`func (o *Curation) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Curation) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Curation) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetVersion

`func (o *Curation) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Curation) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Curation) SetVersion(v int64)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


