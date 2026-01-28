# CollectionComplete

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArticlesCount** | **int64** | Number of articles in collection | 
**Authors** | [**[]Author**](Author.md) | List of collection authors | 
**Categories** | [**[]Category**](Category.md) | List of collection categories | 
**Citation** | **string** | Collection citation | 
**CreatedDate** | **string** | Date when collection was created | 
**CustomFields** | [**[]CustomArticleField**](CustomArticleField.md) | Collection custom fields | 
**Description** | **string** | Collection description | 
**Funding** | [**[]FundingInformation**](FundingInformation.md) | Full Collection funding information | 
**GroupId** | **int64** | Collection group | 
**InstitutionId** | **int64** | Collection institution | 
**Keywords** | **[]string** | List of collection keywords. Tags can be used instead | 
**ModifiedDate** | **string** | Date when collection was last modified | 
**Public** | **bool** | True if collection is published | 
**References** | **[]string** | List of collection references | 
**RelatedMaterials** | [**[]RelatedMaterial**](RelatedMaterial.md) | List of related materials; supersedes references and resource DOI/title. | 
**ResourceDoi** | **string** | Collection resource doi | 
**ResourceId** | **string** | Collection resource id | 
**ResourceLink** | **string** | Collection resource link | 
**ResourceTitle** | **string** | Collection resource title | 
**ResourceVersion** | **int64** | Collection resource version | 
**Tags** | **[]string** | List of collection tags. Keywords can be used instead | 
**Timeline** | [**Timeline**](Timeline.md) |  | 
**Version** | **int64** | Collection version | 
**Doi** | **string** | Collection DOI | 
**Handle** | **string** | Collection Handle | 
**Id** | **int64** | Collection id | 
**Title** | **string** | Collection title | 
**Url** | **string** | Api endpoint | 

## Methods

### NewCollectionComplete

`func NewCollectionComplete(articlesCount int64, authors []Author, categories []Category, citation string, createdDate string, customFields []CustomArticleField, description string, funding []FundingInformation, groupId int64, institutionId int64, keywords []string, modifiedDate string, public bool, references []string, relatedMaterials []RelatedMaterial, resourceDoi string, resourceId string, resourceLink string, resourceTitle string, resourceVersion int64, tags []string, timeline Timeline, version int64, doi string, handle string, id int64, title string, url string, ) *CollectionComplete`

NewCollectionComplete instantiates a new CollectionComplete object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionCompleteWithDefaults

`func NewCollectionCompleteWithDefaults() *CollectionComplete`

NewCollectionCompleteWithDefaults instantiates a new CollectionComplete object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArticlesCount

`func (o *CollectionComplete) GetArticlesCount() int64`

GetArticlesCount returns the ArticlesCount field if non-nil, zero value otherwise.

### GetArticlesCountOk

`func (o *CollectionComplete) GetArticlesCountOk() (*int64, bool)`

GetArticlesCountOk returns a tuple with the ArticlesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArticlesCount

`func (o *CollectionComplete) SetArticlesCount(v int64)`

SetArticlesCount sets ArticlesCount field to given value.


### GetAuthors

`func (o *CollectionComplete) GetAuthors() []Author`

GetAuthors returns the Authors field if non-nil, zero value otherwise.

### GetAuthorsOk

`func (o *CollectionComplete) GetAuthorsOk() (*[]Author, bool)`

GetAuthorsOk returns a tuple with the Authors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthors

`func (o *CollectionComplete) SetAuthors(v []Author)`

SetAuthors sets Authors field to given value.


### GetCategories

`func (o *CollectionComplete) GetCategories() []Category`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *CollectionComplete) GetCategoriesOk() (*[]Category, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *CollectionComplete) SetCategories(v []Category)`

SetCategories sets Categories field to given value.


### GetCitation

`func (o *CollectionComplete) GetCitation() string`

GetCitation returns the Citation field if non-nil, zero value otherwise.

### GetCitationOk

`func (o *CollectionComplete) GetCitationOk() (*string, bool)`

GetCitationOk returns a tuple with the Citation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitation

`func (o *CollectionComplete) SetCitation(v string)`

SetCitation sets Citation field to given value.


### GetCreatedDate

`func (o *CollectionComplete) GetCreatedDate() string`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *CollectionComplete) GetCreatedDateOk() (*string, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *CollectionComplete) SetCreatedDate(v string)`

SetCreatedDate sets CreatedDate field to given value.


### GetCustomFields

`func (o *CollectionComplete) GetCustomFields() []CustomArticleField`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *CollectionComplete) GetCustomFieldsOk() (*[]CustomArticleField, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *CollectionComplete) SetCustomFields(v []CustomArticleField)`

SetCustomFields sets CustomFields field to given value.


### GetDescription

`func (o *CollectionComplete) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CollectionComplete) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CollectionComplete) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetFunding

`func (o *CollectionComplete) GetFunding() []FundingInformation`

GetFunding returns the Funding field if non-nil, zero value otherwise.

### GetFundingOk

`func (o *CollectionComplete) GetFundingOk() (*[]FundingInformation, bool)`

GetFundingOk returns a tuple with the Funding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunding

`func (o *CollectionComplete) SetFunding(v []FundingInformation)`

SetFunding sets Funding field to given value.


### GetGroupId

`func (o *CollectionComplete) GetGroupId() int64`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *CollectionComplete) GetGroupIdOk() (*int64, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *CollectionComplete) SetGroupId(v int64)`

SetGroupId sets GroupId field to given value.


### GetInstitutionId

`func (o *CollectionComplete) GetInstitutionId() int64`

GetInstitutionId returns the InstitutionId field if non-nil, zero value otherwise.

### GetInstitutionIdOk

`func (o *CollectionComplete) GetInstitutionIdOk() (*int64, bool)`

GetInstitutionIdOk returns a tuple with the InstitutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionId

`func (o *CollectionComplete) SetInstitutionId(v int64)`

SetInstitutionId sets InstitutionId field to given value.


### GetKeywords

`func (o *CollectionComplete) GetKeywords() []string`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *CollectionComplete) GetKeywordsOk() (*[]string, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *CollectionComplete) SetKeywords(v []string)`

SetKeywords sets Keywords field to given value.


### GetModifiedDate

`func (o *CollectionComplete) GetModifiedDate() string`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *CollectionComplete) GetModifiedDateOk() (*string, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *CollectionComplete) SetModifiedDate(v string)`

SetModifiedDate sets ModifiedDate field to given value.


### GetPublic

`func (o *CollectionComplete) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *CollectionComplete) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *CollectionComplete) SetPublic(v bool)`

SetPublic sets Public field to given value.


### GetReferences

`func (o *CollectionComplete) GetReferences() []string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *CollectionComplete) GetReferencesOk() (*[]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *CollectionComplete) SetReferences(v []string)`

SetReferences sets References field to given value.


### GetRelatedMaterials

`func (o *CollectionComplete) GetRelatedMaterials() []RelatedMaterial`

GetRelatedMaterials returns the RelatedMaterials field if non-nil, zero value otherwise.

### GetRelatedMaterialsOk

`func (o *CollectionComplete) GetRelatedMaterialsOk() (*[]RelatedMaterial, bool)`

GetRelatedMaterialsOk returns a tuple with the RelatedMaterials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedMaterials

`func (o *CollectionComplete) SetRelatedMaterials(v []RelatedMaterial)`

SetRelatedMaterials sets RelatedMaterials field to given value.


### GetResourceDoi

`func (o *CollectionComplete) GetResourceDoi() string`

GetResourceDoi returns the ResourceDoi field if non-nil, zero value otherwise.

### GetResourceDoiOk

`func (o *CollectionComplete) GetResourceDoiOk() (*string, bool)`

GetResourceDoiOk returns a tuple with the ResourceDoi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceDoi

`func (o *CollectionComplete) SetResourceDoi(v string)`

SetResourceDoi sets ResourceDoi field to given value.


### GetResourceId

`func (o *CollectionComplete) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *CollectionComplete) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *CollectionComplete) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetResourceLink

`func (o *CollectionComplete) GetResourceLink() string`

GetResourceLink returns the ResourceLink field if non-nil, zero value otherwise.

### GetResourceLinkOk

`func (o *CollectionComplete) GetResourceLinkOk() (*string, bool)`

GetResourceLinkOk returns a tuple with the ResourceLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceLink

`func (o *CollectionComplete) SetResourceLink(v string)`

SetResourceLink sets ResourceLink field to given value.


### GetResourceTitle

`func (o *CollectionComplete) GetResourceTitle() string`

GetResourceTitle returns the ResourceTitle field if non-nil, zero value otherwise.

### GetResourceTitleOk

`func (o *CollectionComplete) GetResourceTitleOk() (*string, bool)`

GetResourceTitleOk returns a tuple with the ResourceTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceTitle

`func (o *CollectionComplete) SetResourceTitle(v string)`

SetResourceTitle sets ResourceTitle field to given value.


### GetResourceVersion

`func (o *CollectionComplete) GetResourceVersion() int64`

GetResourceVersion returns the ResourceVersion field if non-nil, zero value otherwise.

### GetResourceVersionOk

`func (o *CollectionComplete) GetResourceVersionOk() (*int64, bool)`

GetResourceVersionOk returns a tuple with the ResourceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceVersion

`func (o *CollectionComplete) SetResourceVersion(v int64)`

SetResourceVersion sets ResourceVersion field to given value.


### GetTags

`func (o *CollectionComplete) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CollectionComplete) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CollectionComplete) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetTimeline

`func (o *CollectionComplete) GetTimeline() Timeline`

GetTimeline returns the Timeline field if non-nil, zero value otherwise.

### GetTimelineOk

`func (o *CollectionComplete) GetTimelineOk() (*Timeline, bool)`

GetTimelineOk returns a tuple with the Timeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeline

`func (o *CollectionComplete) SetTimeline(v Timeline)`

SetTimeline sets Timeline field to given value.


### GetVersion

`func (o *CollectionComplete) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CollectionComplete) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CollectionComplete) SetVersion(v int64)`

SetVersion sets Version field to given value.


### GetDoi

`func (o *CollectionComplete) GetDoi() string`

GetDoi returns the Doi field if non-nil, zero value otherwise.

### GetDoiOk

`func (o *CollectionComplete) GetDoiOk() (*string, bool)`

GetDoiOk returns a tuple with the Doi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoi

`func (o *CollectionComplete) SetDoi(v string)`

SetDoi sets Doi field to given value.


### GetHandle

`func (o *CollectionComplete) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *CollectionComplete) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *CollectionComplete) SetHandle(v string)`

SetHandle sets Handle field to given value.


### GetId

`func (o *CollectionComplete) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CollectionComplete) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CollectionComplete) SetId(v int64)`

SetId sets Id field to given value.


### GetTitle

`func (o *CollectionComplete) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CollectionComplete) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CollectionComplete) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetUrl

`func (o *CollectionComplete) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CollectionComplete) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CollectionComplete) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


