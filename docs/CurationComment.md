# CurationComment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The ID of the comment. | 
**AccountId** | **int64** | The ID of the account which generated this comment. | 
**Type** | **string** | The ID of the account which generated this comment. | 
**Text** | **string** | The value/content of the comment. | 
**CreatedDate** | **string** | The creation date of the comment. | 
**ModifiedDate** | **string** | The date the comment has been modified. | 

## Methods

### NewCurationComment

`func NewCurationComment(id int64, accountId int64, type_ string, text string, createdDate string, modifiedDate string, ) *CurationComment`

NewCurationComment instantiates a new CurationComment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurationCommentWithDefaults

`func NewCurationCommentWithDefaults() *CurationComment`

NewCurationCommentWithDefaults instantiates a new CurationComment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CurationComment) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CurationComment) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CurationComment) SetId(v int64)`

SetId sets Id field to given value.


### GetAccountId

`func (o *CurationComment) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CurationComment) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CurationComment) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### GetType

`func (o *CurationComment) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CurationComment) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CurationComment) SetType(v string)`

SetType sets Type field to given value.


### GetText

`func (o *CurationComment) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *CurationComment) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *CurationComment) SetText(v string)`

SetText sets Text field to given value.


### GetCreatedDate

`func (o *CurationComment) GetCreatedDate() string`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *CurationComment) GetCreatedDateOk() (*string, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *CurationComment) SetCreatedDate(v string)`

SetCreatedDate sets CreatedDate field to given value.


### GetModifiedDate

`func (o *CurationComment) GetModifiedDate() string`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *CurationComment) GetModifiedDateOk() (*string, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *CurationComment) SetModifiedDate(v string)`

SetModifiedDate sets ModifiedDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


