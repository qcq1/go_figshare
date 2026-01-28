# Project

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | Api endpoint | 
**PublishedDate** | **NullableString** | Date when project was published | 
**Id** | **int64** | Project id | 
**Title** | **string** | Project title | 
**CreatedDate** | **string** | Date when project was created | 
**ModifiedDate** | **string** | Date when project was last modified | 

## Methods

### NewProject

`func NewProject(url string, publishedDate NullableString, id int64, title string, createdDate string, modifiedDate string, ) *Project`

NewProject instantiates a new Project object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectWithDefaults

`func NewProjectWithDefaults() *Project`

NewProjectWithDefaults instantiates a new Project object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *Project) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Project) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Project) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetPublishedDate

`func (o *Project) GetPublishedDate() string`

GetPublishedDate returns the PublishedDate field if non-nil, zero value otherwise.

### GetPublishedDateOk

`func (o *Project) GetPublishedDateOk() (*string, bool)`

GetPublishedDateOk returns a tuple with the PublishedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedDate

`func (o *Project) SetPublishedDate(v string)`

SetPublishedDate sets PublishedDate field to given value.


### SetPublishedDateNil

`func (o *Project) SetPublishedDateNil(b bool)`

 SetPublishedDateNil sets the value for PublishedDate to be an explicit nil

### UnsetPublishedDate
`func (o *Project) UnsetPublishedDate()`

UnsetPublishedDate ensures that no value is present for PublishedDate, not even an explicit nil
### GetId

`func (o *Project) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Project) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Project) SetId(v int64)`

SetId sets Id field to given value.


### GetTitle

`func (o *Project) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Project) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Project) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetCreatedDate

`func (o *Project) GetCreatedDate() string`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *Project) GetCreatedDateOk() (*string, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *Project) SetCreatedDate(v string)`

SetCreatedDate sets CreatedDate field to given value.


### GetModifiedDate

`func (o *Project) GetModifiedDate() string`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *Project) GetModifiedDateOk() (*string, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *Project) SetModifiedDate(v string)`

SetModifiedDate sets ModifiedDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


