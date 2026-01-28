# CollectionCompletePrivate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **int64** | ID of the account owning the collection | 
**Funding** | [**[]FundingInformation**](FundingInformation.md) | Full Collection funding information | 
**ResourceId** | **string** | Collection resource id | 
**ResourceDoi** | **string** | Collection resource doi | 
**ResourceTitle** | **string** | Collection resource title | 
**ResourceLink** | **string** | Collection resource link | 
**ResourceVersion** | **int64** | Collection resource version | 
**Version** | **int64** | Collection version | 
**Description** | **string** | Collection description | 
**Categories** | [**[]Category**](Category.md) | List of collection categories | 
**References** | **[]string** | List of collection references | 
**RelatedMaterials** | Pointer to [**[]RelatedMaterial**](RelatedMaterial.md) | List of related materials; supersedes references and resource DOI/title. | [optional] 
**Tags** | **[]string** | List of collection tags. Keywords can be used instead | 
**Keywords** | **[]string** | List of collection keywords. Tags can be used instead | 
**Authors** | [**[]Author**](Author.md) | List of collection authors | 
**InstitutionId** | **int64** | Collection institution | 
**GroupId** | **int64** | Collection group | 
**ArticlesCount** | **int64** | Number of articles in collection | 
**Public** | **bool** | True if collection is published | 
**Citation** | **string** | Collection citation | 
**GroupResourceId** | **NullableString** | Collection group resource id | 
**CustomFields** | [**[]CustomArticleField**](CustomArticleField.md) | Collection custom fields | 
**ModifiedDate** | **string** | Date when collection was last modified | 
**CreatedDate** | **string** | Date when collection was created | 
**Timeline** | [**Timeline**](Timeline.md) |  | 
**Id** | **int64** | Collection id | 
**Title** | **string** | Collection title | 
**Doi** | **string** | Collection DOI | 
**Handle** | **string** | Collection Handle | 
**Url** | **string** | Api endpoint | 
**PublishedDate** | **NullableString** | Date when collection was published  | 

## Methods

### NewCollectionCompletePrivate

`func NewCollectionCompletePrivate(accountId int64, funding []FundingInformation, resourceId string, resourceDoi string, resourceTitle string, resourceLink string, resourceVersion int64, version int64, description string, categories []Category, references []string, tags []string, keywords []string, authors []Author, institutionId int64, groupId int64, articlesCount int64, public bool, citation string, groupResourceId NullableString, customFields []CustomArticleField, modifiedDate string, createdDate string, timeline Timeline, id int64, title string, doi string, handle string, url string, publishedDate NullableString, ) *CollectionCompletePrivate`

NewCollectionCompletePrivate instantiates a new CollectionCompletePrivate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionCompletePrivateWithDefaults

`func NewCollectionCompletePrivateWithDefaults() *CollectionCompletePrivate`

NewCollectionCompletePrivateWithDefaults instantiates a new CollectionCompletePrivate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *CollectionCompletePrivate) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CollectionCompletePrivate) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CollectionCompletePrivate) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### GetFunding

`func (o *CollectionCompletePrivate) GetFunding() []FundingInformation`

GetFunding returns the Funding field if non-nil, zero value otherwise.

### GetFundingOk

`func (o *CollectionCompletePrivate) GetFundingOk() (*[]FundingInformation, bool)`

GetFundingOk returns a tuple with the Funding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunding

`func (o *CollectionCompletePrivate) SetFunding(v []FundingInformation)`

SetFunding sets Funding field to given value.


### GetResourceId

`func (o *CollectionCompletePrivate) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *CollectionCompletePrivate) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *CollectionCompletePrivate) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetResourceDoi

`func (o *CollectionCompletePrivate) GetResourceDoi() string`

GetResourceDoi returns the ResourceDoi field if non-nil, zero value otherwise.

### GetResourceDoiOk

`func (o *CollectionCompletePrivate) GetResourceDoiOk() (*string, bool)`

GetResourceDoiOk returns a tuple with the ResourceDoi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceDoi

`func (o *CollectionCompletePrivate) SetResourceDoi(v string)`

SetResourceDoi sets ResourceDoi field to given value.


### GetResourceTitle

`func (o *CollectionCompletePrivate) GetResourceTitle() string`

GetResourceTitle returns the ResourceTitle field if non-nil, zero value otherwise.

### GetResourceTitleOk

`func (o *CollectionCompletePrivate) GetResourceTitleOk() (*string, bool)`

GetResourceTitleOk returns a tuple with the ResourceTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceTitle

`func (o *CollectionCompletePrivate) SetResourceTitle(v string)`

SetResourceTitle sets ResourceTitle field to given value.


### GetResourceLink

`func (o *CollectionCompletePrivate) GetResourceLink() string`

GetResourceLink returns the ResourceLink field if non-nil, zero value otherwise.

### GetResourceLinkOk

`func (o *CollectionCompletePrivate) GetResourceLinkOk() (*string, bool)`

GetResourceLinkOk returns a tuple with the ResourceLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceLink

`func (o *CollectionCompletePrivate) SetResourceLink(v string)`

SetResourceLink sets ResourceLink field to given value.


### GetResourceVersion

`func (o *CollectionCompletePrivate) GetResourceVersion() int64`

GetResourceVersion returns the ResourceVersion field if non-nil, zero value otherwise.

### GetResourceVersionOk

`func (o *CollectionCompletePrivate) GetResourceVersionOk() (*int64, bool)`

GetResourceVersionOk returns a tuple with the ResourceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceVersion

`func (o *CollectionCompletePrivate) SetResourceVersion(v int64)`

SetResourceVersion sets ResourceVersion field to given value.


### GetVersion

`func (o *CollectionCompletePrivate) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CollectionCompletePrivate) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CollectionCompletePrivate) SetVersion(v int64)`

SetVersion sets Version field to given value.


### GetDescription

`func (o *CollectionCompletePrivate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CollectionCompletePrivate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CollectionCompletePrivate) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCategories

`func (o *CollectionCompletePrivate) GetCategories() []Category`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *CollectionCompletePrivate) GetCategoriesOk() (*[]Category, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *CollectionCompletePrivate) SetCategories(v []Category)`

SetCategories sets Categories field to given value.


### GetReferences

`func (o *CollectionCompletePrivate) GetReferences() []string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *CollectionCompletePrivate) GetReferencesOk() (*[]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *CollectionCompletePrivate) SetReferences(v []string)`

SetReferences sets References field to given value.


### GetRelatedMaterials

`func (o *CollectionCompletePrivate) GetRelatedMaterials() []RelatedMaterial`

GetRelatedMaterials returns the RelatedMaterials field if non-nil, zero value otherwise.

### GetRelatedMaterialsOk

`func (o *CollectionCompletePrivate) GetRelatedMaterialsOk() (*[]RelatedMaterial, bool)`

GetRelatedMaterialsOk returns a tuple with the RelatedMaterials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedMaterials

`func (o *CollectionCompletePrivate) SetRelatedMaterials(v []RelatedMaterial)`

SetRelatedMaterials sets RelatedMaterials field to given value.

### HasRelatedMaterials

`func (o *CollectionCompletePrivate) HasRelatedMaterials() bool`

HasRelatedMaterials returns a boolean if a field has been set.

### GetTags

`func (o *CollectionCompletePrivate) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CollectionCompletePrivate) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CollectionCompletePrivate) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetKeywords

`func (o *CollectionCompletePrivate) GetKeywords() []string`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *CollectionCompletePrivate) GetKeywordsOk() (*[]string, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *CollectionCompletePrivate) SetKeywords(v []string)`

SetKeywords sets Keywords field to given value.


### GetAuthors

`func (o *CollectionCompletePrivate) GetAuthors() []Author`

GetAuthors returns the Authors field if non-nil, zero value otherwise.

### GetAuthorsOk

`func (o *CollectionCompletePrivate) GetAuthorsOk() (*[]Author, bool)`

GetAuthorsOk returns a tuple with the Authors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthors

`func (o *CollectionCompletePrivate) SetAuthors(v []Author)`

SetAuthors sets Authors field to given value.


### GetInstitutionId

`func (o *CollectionCompletePrivate) GetInstitutionId() int64`

GetInstitutionId returns the InstitutionId field if non-nil, zero value otherwise.

### GetInstitutionIdOk

`func (o *CollectionCompletePrivate) GetInstitutionIdOk() (*int64, bool)`

GetInstitutionIdOk returns a tuple with the InstitutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionId

`func (o *CollectionCompletePrivate) SetInstitutionId(v int64)`

SetInstitutionId sets InstitutionId field to given value.


### GetGroupId

`func (o *CollectionCompletePrivate) GetGroupId() int64`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *CollectionCompletePrivate) GetGroupIdOk() (*int64, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *CollectionCompletePrivate) SetGroupId(v int64)`

SetGroupId sets GroupId field to given value.


### GetArticlesCount

`func (o *CollectionCompletePrivate) GetArticlesCount() int64`

GetArticlesCount returns the ArticlesCount field if non-nil, zero value otherwise.

### GetArticlesCountOk

`func (o *CollectionCompletePrivate) GetArticlesCountOk() (*int64, bool)`

GetArticlesCountOk returns a tuple with the ArticlesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArticlesCount

`func (o *CollectionCompletePrivate) SetArticlesCount(v int64)`

SetArticlesCount sets ArticlesCount field to given value.


### GetPublic

`func (o *CollectionCompletePrivate) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *CollectionCompletePrivate) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *CollectionCompletePrivate) SetPublic(v bool)`

SetPublic sets Public field to given value.


### GetCitation

`func (o *CollectionCompletePrivate) GetCitation() string`

GetCitation returns the Citation field if non-nil, zero value otherwise.

### GetCitationOk

`func (o *CollectionCompletePrivate) GetCitationOk() (*string, bool)`

GetCitationOk returns a tuple with the Citation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitation

`func (o *CollectionCompletePrivate) SetCitation(v string)`

SetCitation sets Citation field to given value.


### GetGroupResourceId

`func (o *CollectionCompletePrivate) GetGroupResourceId() string`

GetGroupResourceId returns the GroupResourceId field if non-nil, zero value otherwise.

### GetGroupResourceIdOk

`func (o *CollectionCompletePrivate) GetGroupResourceIdOk() (*string, bool)`

GetGroupResourceIdOk returns a tuple with the GroupResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupResourceId

`func (o *CollectionCompletePrivate) SetGroupResourceId(v string)`

SetGroupResourceId sets GroupResourceId field to given value.


### SetGroupResourceIdNil

`func (o *CollectionCompletePrivate) SetGroupResourceIdNil(b bool)`

 SetGroupResourceIdNil sets the value for GroupResourceId to be an explicit nil

### UnsetGroupResourceId
`func (o *CollectionCompletePrivate) UnsetGroupResourceId()`

UnsetGroupResourceId ensures that no value is present for GroupResourceId, not even an explicit nil
### GetCustomFields

`func (o *CollectionCompletePrivate) GetCustomFields() []CustomArticleField`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *CollectionCompletePrivate) GetCustomFieldsOk() (*[]CustomArticleField, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *CollectionCompletePrivate) SetCustomFields(v []CustomArticleField)`

SetCustomFields sets CustomFields field to given value.


### GetModifiedDate

`func (o *CollectionCompletePrivate) GetModifiedDate() string`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *CollectionCompletePrivate) GetModifiedDateOk() (*string, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *CollectionCompletePrivate) SetModifiedDate(v string)`

SetModifiedDate sets ModifiedDate field to given value.


### GetCreatedDate

`func (o *CollectionCompletePrivate) GetCreatedDate() string`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *CollectionCompletePrivate) GetCreatedDateOk() (*string, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *CollectionCompletePrivate) SetCreatedDate(v string)`

SetCreatedDate sets CreatedDate field to given value.


### GetTimeline

`func (o *CollectionCompletePrivate) GetTimeline() Timeline`

GetTimeline returns the Timeline field if non-nil, zero value otherwise.

### GetTimelineOk

`func (o *CollectionCompletePrivate) GetTimelineOk() (*Timeline, bool)`

GetTimelineOk returns a tuple with the Timeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeline

`func (o *CollectionCompletePrivate) SetTimeline(v Timeline)`

SetTimeline sets Timeline field to given value.


### GetId

`func (o *CollectionCompletePrivate) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CollectionCompletePrivate) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CollectionCompletePrivate) SetId(v int64)`

SetId sets Id field to given value.


### GetTitle

`func (o *CollectionCompletePrivate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CollectionCompletePrivate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CollectionCompletePrivate) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDoi

`func (o *CollectionCompletePrivate) GetDoi() string`

GetDoi returns the Doi field if non-nil, zero value otherwise.

### GetDoiOk

`func (o *CollectionCompletePrivate) GetDoiOk() (*string, bool)`

GetDoiOk returns a tuple with the Doi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoi

`func (o *CollectionCompletePrivate) SetDoi(v string)`

SetDoi sets Doi field to given value.


### GetHandle

`func (o *CollectionCompletePrivate) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *CollectionCompletePrivate) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *CollectionCompletePrivate) SetHandle(v string)`

SetHandle sets Handle field to given value.


### GetUrl

`func (o *CollectionCompletePrivate) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CollectionCompletePrivate) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CollectionCompletePrivate) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetPublishedDate

`func (o *CollectionCompletePrivate) GetPublishedDate() string`

GetPublishedDate returns the PublishedDate field if non-nil, zero value otherwise.

### GetPublishedDateOk

`func (o *CollectionCompletePrivate) GetPublishedDateOk() (*string, bool)`

GetPublishedDateOk returns a tuple with the PublishedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedDate

`func (o *CollectionCompletePrivate) SetPublishedDate(v string)`

SetPublishedDate sets PublishedDate field to given value.


### SetPublishedDateNil

`func (o *CollectionCompletePrivate) SetPublishedDateNil(b bool)`

 SetPublishedDateNil sets the value for PublishedDate to be an explicit nil

### UnsetPublishedDate
`func (o *CollectionCompletePrivate) UnsetPublishedDate()`

UnsetPublishedDate ensures that no value is present for PublishedDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


