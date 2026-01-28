# CollectionUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Funding** | Pointer to **string** | Grant number or funding authority | [optional] [default to ""]
**FundingList** | Pointer to [**[]FundingCreate**](FundingCreate.md) | Funding creation / update items | [optional] 
**Title** | Pointer to **string** | Title of collection | [optional] 
**Description** | Pointer to **string** | The collection description. In a publisher case, usually this is the remote collection description | [optional] [default to ""]
**Articles** | Pointer to **[]int32** | List of articles to be associated with the collection | [optional] 
**Authors** | Pointer to **[]map[string]interface{}** | List of authors to be associated with the collection. The list can contain the following fields: id, name, first_name, last_name, email, orcid_id. If an id is supplied, it will take priority and everything else will be ignored. For adding more authors use the specific authors endpoint. | [optional] 
**Categories** | Pointer to **[]int64** | List of category ids to be associated with the collection (e.g [1, 23, 33, 66]) | [optional] 
**CategoriesBySourceId** | Pointer to **[]string** | List of category source ids to be associated with the article, supersedes the categories property | [optional] 
**Tags** | Pointer to **[]string** | List of tags to be associated with the collection. Keywords can be used instead | [optional] 
**Keywords** | Pointer to **[]string** | List of tags to be associated with the collection. Tags can be used instead | [optional] 
**References** | Pointer to **[]string** | List of links to be associated with the collection (e.g [\&quot;http://link1\&quot;, \&quot;http://link2\&quot;, \&quot;http://link3\&quot;]) | [optional] 
**RelatedMaterials** | Pointer to [**[]RelatedMaterial**](RelatedMaterial.md) | List of related materials; supersedes references and resource DOI/title. | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** | List of key, values pairs to be associated with the collection | [optional] 
**CustomFieldsList** | Pointer to [**[]CustomArticleFieldAdd**](CustomArticleFieldAdd.md) | List of custom fields values, supersedes custom_fields parameter | [optional] 
**Doi** | Pointer to **string** | Not applicable for regular users. In an institutional case, make sure your group supports setting DOIs. This setting is applied by figshare via opening a ticket through our support/helpdesk system. | [optional] [default to ""]
**Handle** | Pointer to **string** | Not applicable for regular users. In an institutional case, make sure your group supports setting Handles. This setting is applied by figshare via opening a ticket through our support/helpdesk system. | [optional] [default to ""]
**ResourceId** | Pointer to **string** | Not applicable to regular users. In a publisher case, this is the publisher article id | [optional] 
**ResourceDoi** | Pointer to **string** | Not applicable to regular users. In a publisher case, this is the publisher article DOI. | [optional] [default to ""]
**ResourceLink** | Pointer to **string** | Not applicable to regular users. In a publisher case, this is the publisher article link | [optional] 
**ResourceTitle** | Pointer to **string** | Not applicable to regular users. In a publisher case, this is the publisher article title. | [optional] [default to ""]
**ResourceVersion** | Pointer to **int32** | Not applicable to regular users. In a publisher case, this is the publisher article version | [optional] 
**GroupId** | Pointer to **int64** | Not applicable to regular users. This field is reserved to institutions/publishers with access to assign to specific groups | [optional] 
**Timeline** | Pointer to [**TimelineUpdate**](TimelineUpdate.md) |  | [optional] 

## Methods

### NewCollectionUpdate

`func NewCollectionUpdate() *CollectionUpdate`

NewCollectionUpdate instantiates a new CollectionUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionUpdateWithDefaults

`func NewCollectionUpdateWithDefaults() *CollectionUpdate`

NewCollectionUpdateWithDefaults instantiates a new CollectionUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunding

`func (o *CollectionUpdate) GetFunding() string`

GetFunding returns the Funding field if non-nil, zero value otherwise.

### GetFundingOk

`func (o *CollectionUpdate) GetFundingOk() (*string, bool)`

GetFundingOk returns a tuple with the Funding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunding

`func (o *CollectionUpdate) SetFunding(v string)`

SetFunding sets Funding field to given value.

### HasFunding

`func (o *CollectionUpdate) HasFunding() bool`

HasFunding returns a boolean if a field has been set.

### GetFundingList

`func (o *CollectionUpdate) GetFundingList() []FundingCreate`

GetFundingList returns the FundingList field if non-nil, zero value otherwise.

### GetFundingListOk

`func (o *CollectionUpdate) GetFundingListOk() (*[]FundingCreate, bool)`

GetFundingListOk returns a tuple with the FundingList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingList

`func (o *CollectionUpdate) SetFundingList(v []FundingCreate)`

SetFundingList sets FundingList field to given value.

### HasFundingList

`func (o *CollectionUpdate) HasFundingList() bool`

HasFundingList returns a boolean if a field has been set.

### GetTitle

`func (o *CollectionUpdate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CollectionUpdate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CollectionUpdate) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CollectionUpdate) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *CollectionUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CollectionUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CollectionUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CollectionUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetArticles

`func (o *CollectionUpdate) GetArticles() []int32`

GetArticles returns the Articles field if non-nil, zero value otherwise.

### GetArticlesOk

`func (o *CollectionUpdate) GetArticlesOk() (*[]int32, bool)`

GetArticlesOk returns a tuple with the Articles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArticles

`func (o *CollectionUpdate) SetArticles(v []int32)`

SetArticles sets Articles field to given value.

### HasArticles

`func (o *CollectionUpdate) HasArticles() bool`

HasArticles returns a boolean if a field has been set.

### GetAuthors

`func (o *CollectionUpdate) GetAuthors() []map[string]interface{}`

GetAuthors returns the Authors field if non-nil, zero value otherwise.

### GetAuthorsOk

`func (o *CollectionUpdate) GetAuthorsOk() (*[]map[string]interface{}, bool)`

GetAuthorsOk returns a tuple with the Authors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthors

`func (o *CollectionUpdate) SetAuthors(v []map[string]interface{})`

SetAuthors sets Authors field to given value.

### HasAuthors

`func (o *CollectionUpdate) HasAuthors() bool`

HasAuthors returns a boolean if a field has been set.

### GetCategories

`func (o *CollectionUpdate) GetCategories() []int64`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *CollectionUpdate) GetCategoriesOk() (*[]int64, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *CollectionUpdate) SetCategories(v []int64)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *CollectionUpdate) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetCategoriesBySourceId

`func (o *CollectionUpdate) GetCategoriesBySourceId() []string`

GetCategoriesBySourceId returns the CategoriesBySourceId field if non-nil, zero value otherwise.

### GetCategoriesBySourceIdOk

`func (o *CollectionUpdate) GetCategoriesBySourceIdOk() (*[]string, bool)`

GetCategoriesBySourceIdOk returns a tuple with the CategoriesBySourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoriesBySourceId

`func (o *CollectionUpdate) SetCategoriesBySourceId(v []string)`

SetCategoriesBySourceId sets CategoriesBySourceId field to given value.

### HasCategoriesBySourceId

`func (o *CollectionUpdate) HasCategoriesBySourceId() bool`

HasCategoriesBySourceId returns a boolean if a field has been set.

### GetTags

`func (o *CollectionUpdate) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CollectionUpdate) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CollectionUpdate) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CollectionUpdate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetKeywords

`func (o *CollectionUpdate) GetKeywords() []string`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *CollectionUpdate) GetKeywordsOk() (*[]string, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *CollectionUpdate) SetKeywords(v []string)`

SetKeywords sets Keywords field to given value.

### HasKeywords

`func (o *CollectionUpdate) HasKeywords() bool`

HasKeywords returns a boolean if a field has been set.

### GetReferences

`func (o *CollectionUpdate) GetReferences() []string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *CollectionUpdate) GetReferencesOk() (*[]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *CollectionUpdate) SetReferences(v []string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *CollectionUpdate) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetRelatedMaterials

`func (o *CollectionUpdate) GetRelatedMaterials() []RelatedMaterial`

GetRelatedMaterials returns the RelatedMaterials field if non-nil, zero value otherwise.

### GetRelatedMaterialsOk

`func (o *CollectionUpdate) GetRelatedMaterialsOk() (*[]RelatedMaterial, bool)`

GetRelatedMaterialsOk returns a tuple with the RelatedMaterials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedMaterials

`func (o *CollectionUpdate) SetRelatedMaterials(v []RelatedMaterial)`

SetRelatedMaterials sets RelatedMaterials field to given value.

### HasRelatedMaterials

`func (o *CollectionUpdate) HasRelatedMaterials() bool`

HasRelatedMaterials returns a boolean if a field has been set.

### GetCustomFields

`func (o *CollectionUpdate) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *CollectionUpdate) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *CollectionUpdate) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *CollectionUpdate) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCustomFieldsList

`func (o *CollectionUpdate) GetCustomFieldsList() []CustomArticleFieldAdd`

GetCustomFieldsList returns the CustomFieldsList field if non-nil, zero value otherwise.

### GetCustomFieldsListOk

`func (o *CollectionUpdate) GetCustomFieldsListOk() (*[]CustomArticleFieldAdd, bool)`

GetCustomFieldsListOk returns a tuple with the CustomFieldsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFieldsList

`func (o *CollectionUpdate) SetCustomFieldsList(v []CustomArticleFieldAdd)`

SetCustomFieldsList sets CustomFieldsList field to given value.

### HasCustomFieldsList

`func (o *CollectionUpdate) HasCustomFieldsList() bool`

HasCustomFieldsList returns a boolean if a field has been set.

### GetDoi

`func (o *CollectionUpdate) GetDoi() string`

GetDoi returns the Doi field if non-nil, zero value otherwise.

### GetDoiOk

`func (o *CollectionUpdate) GetDoiOk() (*string, bool)`

GetDoiOk returns a tuple with the Doi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoi

`func (o *CollectionUpdate) SetDoi(v string)`

SetDoi sets Doi field to given value.

### HasDoi

`func (o *CollectionUpdate) HasDoi() bool`

HasDoi returns a boolean if a field has been set.

### GetHandle

`func (o *CollectionUpdate) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *CollectionUpdate) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *CollectionUpdate) SetHandle(v string)`

SetHandle sets Handle field to given value.

### HasHandle

`func (o *CollectionUpdate) HasHandle() bool`

HasHandle returns a boolean if a field has been set.

### GetResourceId

`func (o *CollectionUpdate) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *CollectionUpdate) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *CollectionUpdate) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *CollectionUpdate) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetResourceDoi

`func (o *CollectionUpdate) GetResourceDoi() string`

GetResourceDoi returns the ResourceDoi field if non-nil, zero value otherwise.

### GetResourceDoiOk

`func (o *CollectionUpdate) GetResourceDoiOk() (*string, bool)`

GetResourceDoiOk returns a tuple with the ResourceDoi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceDoi

`func (o *CollectionUpdate) SetResourceDoi(v string)`

SetResourceDoi sets ResourceDoi field to given value.

### HasResourceDoi

`func (o *CollectionUpdate) HasResourceDoi() bool`

HasResourceDoi returns a boolean if a field has been set.

### GetResourceLink

`func (o *CollectionUpdate) GetResourceLink() string`

GetResourceLink returns the ResourceLink field if non-nil, zero value otherwise.

### GetResourceLinkOk

`func (o *CollectionUpdate) GetResourceLinkOk() (*string, bool)`

GetResourceLinkOk returns a tuple with the ResourceLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceLink

`func (o *CollectionUpdate) SetResourceLink(v string)`

SetResourceLink sets ResourceLink field to given value.

### HasResourceLink

`func (o *CollectionUpdate) HasResourceLink() bool`

HasResourceLink returns a boolean if a field has been set.

### GetResourceTitle

`func (o *CollectionUpdate) GetResourceTitle() string`

GetResourceTitle returns the ResourceTitle field if non-nil, zero value otherwise.

### GetResourceTitleOk

`func (o *CollectionUpdate) GetResourceTitleOk() (*string, bool)`

GetResourceTitleOk returns a tuple with the ResourceTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceTitle

`func (o *CollectionUpdate) SetResourceTitle(v string)`

SetResourceTitle sets ResourceTitle field to given value.

### HasResourceTitle

`func (o *CollectionUpdate) HasResourceTitle() bool`

HasResourceTitle returns a boolean if a field has been set.

### GetResourceVersion

`func (o *CollectionUpdate) GetResourceVersion() int32`

GetResourceVersion returns the ResourceVersion field if non-nil, zero value otherwise.

### GetResourceVersionOk

`func (o *CollectionUpdate) GetResourceVersionOk() (*int32, bool)`

GetResourceVersionOk returns a tuple with the ResourceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceVersion

`func (o *CollectionUpdate) SetResourceVersion(v int32)`

SetResourceVersion sets ResourceVersion field to given value.

### HasResourceVersion

`func (o *CollectionUpdate) HasResourceVersion() bool`

HasResourceVersion returns a boolean if a field has been set.

### GetGroupId

`func (o *CollectionUpdate) GetGroupId() int64`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *CollectionUpdate) GetGroupIdOk() (*int64, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *CollectionUpdate) SetGroupId(v int64)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *CollectionUpdate) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetTimeline

`func (o *CollectionUpdate) GetTimeline() TimelineUpdate`

GetTimeline returns the Timeline field if non-nil, zero value otherwise.

### GetTimelineOk

`func (o *CollectionUpdate) GetTimelineOk() (*TimelineUpdate, bool)`

GetTimelineOk returns a tuple with the Timeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeline

`func (o *CollectionUpdate) SetTimeline(v TimelineUpdate)`

SetTimeline sets Timeline field to given value.

### HasTimeline

`func (o *CollectionUpdate) HasTimeline() bool`

HasTimeline returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


