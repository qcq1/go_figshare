# ProjectNote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Abstract** | **string** | Note Abstract - short/truncated content | 
**CreatedDate** | **string** | Date when note was created | 
**Id** | **int64** | Project note id | 
**ModifiedDate** | **string** | Date when note was last modified | 
**UserId** | **int64** | User who wrote the note | 
**UserName** | **string** | Username of the one who wrote the note | 

## Methods

### NewProjectNote

`func NewProjectNote(abstract string, createdDate string, id int64, modifiedDate string, userId int64, userName string, ) *ProjectNote`

NewProjectNote instantiates a new ProjectNote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectNoteWithDefaults

`func NewProjectNoteWithDefaults() *ProjectNote`

NewProjectNoteWithDefaults instantiates a new ProjectNote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbstract

`func (o *ProjectNote) GetAbstract() string`

GetAbstract returns the Abstract field if non-nil, zero value otherwise.

### GetAbstractOk

`func (o *ProjectNote) GetAbstractOk() (*string, bool)`

GetAbstractOk returns a tuple with the Abstract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbstract

`func (o *ProjectNote) SetAbstract(v string)`

SetAbstract sets Abstract field to given value.


### GetCreatedDate

`func (o *ProjectNote) GetCreatedDate() string`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *ProjectNote) GetCreatedDateOk() (*string, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *ProjectNote) SetCreatedDate(v string)`

SetCreatedDate sets CreatedDate field to given value.


### GetId

`func (o *ProjectNote) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectNote) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectNote) SetId(v int64)`

SetId sets Id field to given value.


### GetModifiedDate

`func (o *ProjectNote) GetModifiedDate() string`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *ProjectNote) GetModifiedDateOk() (*string, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *ProjectNote) SetModifiedDate(v string)`

SetModifiedDate sets ModifiedDate field to given value.


### GetUserId

`func (o *ProjectNote) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ProjectNote) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ProjectNote) SetUserId(v int64)`

SetUserId sets UserId field to given value.


### GetUserName

`func (o *ProjectNote) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *ProjectNote) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *ProjectNote) SetUserName(v string)`

SetUserName sets UserName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


