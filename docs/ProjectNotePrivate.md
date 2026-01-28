# ProjectNotePrivate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | **string** | Full text of note | 
**Id** | **int64** | Project note id | 
**UserId** | **int64** | User who wrote the note | 
**Abstract** | **string** | Note Abstract - short/truncated content | 
**UserName** | **string** | Username of the one who wrote the note | 
**CreatedDate** | **string** | Date when note was created | 
**ModifiedDate** | **string** | Date when note was last modified | 

## Methods

### NewProjectNotePrivate

`func NewProjectNotePrivate(text string, id int64, userId int64, abstract string, userName string, createdDate string, modifiedDate string, ) *ProjectNotePrivate`

NewProjectNotePrivate instantiates a new ProjectNotePrivate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectNotePrivateWithDefaults

`func NewProjectNotePrivateWithDefaults() *ProjectNotePrivate`

NewProjectNotePrivateWithDefaults instantiates a new ProjectNotePrivate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *ProjectNotePrivate) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ProjectNotePrivate) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ProjectNotePrivate) SetText(v string)`

SetText sets Text field to given value.


### GetId

`func (o *ProjectNotePrivate) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectNotePrivate) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectNotePrivate) SetId(v int64)`

SetId sets Id field to given value.


### GetUserId

`func (o *ProjectNotePrivate) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ProjectNotePrivate) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ProjectNotePrivate) SetUserId(v int64)`

SetUserId sets UserId field to given value.


### GetAbstract

`func (o *ProjectNotePrivate) GetAbstract() string`

GetAbstract returns the Abstract field if non-nil, zero value otherwise.

### GetAbstractOk

`func (o *ProjectNotePrivate) GetAbstractOk() (*string, bool)`

GetAbstractOk returns a tuple with the Abstract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbstract

`func (o *ProjectNotePrivate) SetAbstract(v string)`

SetAbstract sets Abstract field to given value.


### GetUserName

`func (o *ProjectNotePrivate) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *ProjectNotePrivate) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *ProjectNotePrivate) SetUserName(v string)`

SetUserName sets UserName field to given value.


### GetCreatedDate

`func (o *ProjectNotePrivate) GetCreatedDate() string`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *ProjectNotePrivate) GetCreatedDateOk() (*string, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *ProjectNotePrivate) SetCreatedDate(v string)`

SetCreatedDate sets CreatedDate field to given value.


### GetModifiedDate

`func (o *ProjectNotePrivate) GetModifiedDate() string`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *ProjectNotePrivate) GetModifiedDateOk() (*string, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *ProjectNotePrivate) SetModifiedDate(v string)`

SetModifiedDate sets ModifiedDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


