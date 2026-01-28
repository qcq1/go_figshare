# ProjectComplete

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Funding** | **string** | Project funding | 
**FundingList** | [**[]FundingInformation**](FundingInformation.md) | Full Project funding information | 
**Description** | **string** | Project description | 
**FigshareUrl** | **NullableString** | Project public url | 
**Collaborators** | [**[]Collaborator**](Collaborator.md) | List of project collaborators | 
**CustomFields** | [**[]CustomArticleField**](CustomArticleField.md) | Project custom fields | 
**ModifiedDate** | **string** | Date when project was last modified | 
**CreatedDate** | **string** | Date when project was created | 
**Url** | **string** | Api endpoint | 
**PublishedDate** | **NullableString** | Date when project was published | 
**Id** | **int64** | Project id | 
**Title** | **string** | Project title | 

## Methods

### NewProjectComplete

`func NewProjectComplete(funding string, fundingList []FundingInformation, description string, figshareUrl NullableString, collaborators []Collaborator, customFields []CustomArticleField, modifiedDate string, createdDate string, url string, publishedDate NullableString, id int64, title string, ) *ProjectComplete`

NewProjectComplete instantiates a new ProjectComplete object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectCompleteWithDefaults

`func NewProjectCompleteWithDefaults() *ProjectComplete`

NewProjectCompleteWithDefaults instantiates a new ProjectComplete object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunding

`func (o *ProjectComplete) GetFunding() string`

GetFunding returns the Funding field if non-nil, zero value otherwise.

### GetFundingOk

`func (o *ProjectComplete) GetFundingOk() (*string, bool)`

GetFundingOk returns a tuple with the Funding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunding

`func (o *ProjectComplete) SetFunding(v string)`

SetFunding sets Funding field to given value.


### GetFundingList

`func (o *ProjectComplete) GetFundingList() []FundingInformation`

GetFundingList returns the FundingList field if non-nil, zero value otherwise.

### GetFundingListOk

`func (o *ProjectComplete) GetFundingListOk() (*[]FundingInformation, bool)`

GetFundingListOk returns a tuple with the FundingList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingList

`func (o *ProjectComplete) SetFundingList(v []FundingInformation)`

SetFundingList sets FundingList field to given value.


### GetDescription

`func (o *ProjectComplete) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProjectComplete) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProjectComplete) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetFigshareUrl

`func (o *ProjectComplete) GetFigshareUrl() string`

GetFigshareUrl returns the FigshareUrl field if non-nil, zero value otherwise.

### GetFigshareUrlOk

`func (o *ProjectComplete) GetFigshareUrlOk() (*string, bool)`

GetFigshareUrlOk returns a tuple with the FigshareUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFigshareUrl

`func (o *ProjectComplete) SetFigshareUrl(v string)`

SetFigshareUrl sets FigshareUrl field to given value.


### SetFigshareUrlNil

`func (o *ProjectComplete) SetFigshareUrlNil(b bool)`

 SetFigshareUrlNil sets the value for FigshareUrl to be an explicit nil

### UnsetFigshareUrl
`func (o *ProjectComplete) UnsetFigshareUrl()`

UnsetFigshareUrl ensures that no value is present for FigshareUrl, not even an explicit nil
### GetCollaborators

`func (o *ProjectComplete) GetCollaborators() []Collaborator`

GetCollaborators returns the Collaborators field if non-nil, zero value otherwise.

### GetCollaboratorsOk

`func (o *ProjectComplete) GetCollaboratorsOk() (*[]Collaborator, bool)`

GetCollaboratorsOk returns a tuple with the Collaborators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollaborators

`func (o *ProjectComplete) SetCollaborators(v []Collaborator)`

SetCollaborators sets Collaborators field to given value.


### GetCustomFields

`func (o *ProjectComplete) GetCustomFields() []CustomArticleField`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ProjectComplete) GetCustomFieldsOk() (*[]CustomArticleField, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ProjectComplete) SetCustomFields(v []CustomArticleField)`

SetCustomFields sets CustomFields field to given value.


### GetModifiedDate

`func (o *ProjectComplete) GetModifiedDate() string`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *ProjectComplete) GetModifiedDateOk() (*string, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *ProjectComplete) SetModifiedDate(v string)`

SetModifiedDate sets ModifiedDate field to given value.


### GetCreatedDate

`func (o *ProjectComplete) GetCreatedDate() string`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *ProjectComplete) GetCreatedDateOk() (*string, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *ProjectComplete) SetCreatedDate(v string)`

SetCreatedDate sets CreatedDate field to given value.


### GetUrl

`func (o *ProjectComplete) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ProjectComplete) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ProjectComplete) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetPublishedDate

`func (o *ProjectComplete) GetPublishedDate() string`

GetPublishedDate returns the PublishedDate field if non-nil, zero value otherwise.

### GetPublishedDateOk

`func (o *ProjectComplete) GetPublishedDateOk() (*string, bool)`

GetPublishedDateOk returns a tuple with the PublishedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedDate

`func (o *ProjectComplete) SetPublishedDate(v string)`

SetPublishedDate sets PublishedDate field to given value.


### SetPublishedDateNil

`func (o *ProjectComplete) SetPublishedDateNil(b bool)`

 SetPublishedDateNil sets the value for PublishedDate to be an explicit nil

### UnsetPublishedDate
`func (o *ProjectComplete) UnsetPublishedDate()`

UnsetPublishedDate ensures that no value is present for PublishedDate, not even an explicit nil
### GetId

`func (o *ProjectComplete) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectComplete) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectComplete) SetId(v int64)`

SetId sets Id field to given value.


### GetTitle

`func (o *ProjectComplete) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ProjectComplete) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ProjectComplete) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


