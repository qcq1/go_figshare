# CollectionCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Funding** | Pointer to **string** | Grant number or funding authority | [optional] [default to ""]
**FundingList** | Pointer to [**[]FundingCreate**](FundingCreate.md) | Funding creation / update items | [optional] 
**Title** | **string** | Title of collection | 
**Description** | Pointer to **string** | The collection description. In a publisher case, usually this is the remote collection description | [optional] [default to ""]
**Articles** | Pointer to **[]int32** | List of articles to be associated with the collection | [optional] 
**Authors** | Pointer to **[]map[string]interface{}** | List of authors to be associated with the collection. The list can contain the following fields: id, name, first_name, last_name, email, orcid_id. If an id is supplied, it will take priority and everything else will be ignored. For adding more authors use the specific authors endpoint. | [optional] 
**Categories** | Pointer to **[]int64** | List of category ids to be associated with the collection(e.g [1, 23, 33, 66]) | [optional] 
**CategoriesBySourceId** | Pointer to **[]string** | List of category source ids to be associated with the collection, supersedes the categories property | [optional] 
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

### NewCollectionCreate

`func NewCollectionCreate(title string, ) *CollectionCreate`

NewCollectionCreate instantiates a new CollectionCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionCreateWithDefaults

`func NewCollectionCreateWithDefaults() *CollectionCreate`

NewCollectionCreateWithDefaults instantiates a new CollectionCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunding

`func (o *CollectionCreate) GetFunding() string`

GetFunding returns the Funding field if non-nil, zero value otherwise.

### GetFundingOk

`func (o *CollectionCreate) GetFundingOk() (*string, bool)`

GetFundingOk returns a tuple with the Funding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunding

`func (o *CollectionCreate) SetFunding(v string)`

SetFunding sets Funding field to given value.

### HasFunding

`func (o *CollectionCreate) HasFunding() bool`

HasFunding returns a boolean if a field has been set.

### GetFundingList

`func (o *CollectionCreate) GetFundingList() []FundingCreate`

GetFundingList returns the FundingList field if non-nil, zero value otherwise.

### GetFundingListOk

`func (o *CollectionCreate) GetFundingListOk() (*[]FundingCreate, bool)`

GetFundingListOk returns a tuple with the FundingList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingList

`func (o *CollectionCreate) SetFundingList(v []FundingCreate)`

SetFundingList sets FundingList field to given value.

### HasFundingList

`func (o *CollectionCreate) HasFundingList() bool`

HasFundingList returns a boolean if a field has been set.

### GetTitle

`func (o *CollectionCreate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CollectionCreate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CollectionCreate) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *CollectionCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CollectionCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CollectionCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CollectionCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetArticles

`func (o *CollectionCreate) GetArticles() []int32`

GetArticles returns the Articles field if non-nil, zero value otherwise.

### GetArticlesOk

`func (o *CollectionCreate) GetArticlesOk() (*[]int32, bool)`

GetArticlesOk returns a tuple with the Articles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArticles

`func (o *CollectionCreate) SetArticles(v []int32)`

SetArticles sets Articles field to given value.

### HasArticles

`func (o *CollectionCreate) HasArticles() bool`

HasArticles returns a boolean if a field has been set.

### GetAuthors

`func (o *CollectionCreate) GetAuthors() []map[string]interface{}`

GetAuthors returns the Authors field if non-nil, zero value otherwise.

### GetAuthorsOk

`func (o *CollectionCreate) GetAuthorsOk() (*[]map[string]interface{}, bool)`

GetAuthorsOk returns a tuple with the Authors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthors

`func (o *CollectionCreate) SetAuthors(v []map[string]interface{})`

SetAuthors sets Authors field to given value.

### HasAuthors

`func (o *CollectionCreate) HasAuthors() bool`

HasAuthors returns a boolean if a field has been set.

### GetCategories

`func (o *CollectionCreate) GetCategories() []int64`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *CollectionCreate) GetCategoriesOk() (*[]int64, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *CollectionCreate) SetCategories(v []int64)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *CollectionCreate) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetCategoriesBySourceId

`func (o *CollectionCreate) GetCategoriesBySourceId() []string`

GetCategoriesBySourceId returns the CategoriesBySourceId field if non-nil, zero value otherwise.

### GetCategoriesBySourceIdOk

`func (o *CollectionCreate) GetCategoriesBySourceIdOk() (*[]string, bool)`

GetCategoriesBySourceIdOk returns a tuple with the CategoriesBySourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoriesBySourceId

`func (o *CollectionCreate) SetCategoriesBySourceId(v []string)`

SetCategoriesBySourceId sets CategoriesBySourceId field to given value.

### HasCategoriesBySourceId

`func (o *CollectionCreate) HasCategoriesBySourceId() bool`

HasCategoriesBySourceId returns a boolean if a field has been set.

### GetTags

`func (o *CollectionCreate) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CollectionCreate) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CollectionCreate) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CollectionCreate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetKeywords

`func (o *CollectionCreate) GetKeywords() []string`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *CollectionCreate) GetKeywordsOk() (*[]string, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *CollectionCreate) SetKeywords(v []string)`

SetKeywords sets Keywords field to given value.

### HasKeywords

`func (o *CollectionCreate) HasKeywords() bool`

HasKeywords returns a boolean if a field has been set.

### GetReferences

`func (o *CollectionCreate) GetReferences() []string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *CollectionCreate) GetReferencesOk() (*[]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *CollectionCreate) SetReferences(v []string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *CollectionCreate) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetRelatedMaterials

`func (o *CollectionCreate) GetRelatedMaterials() []RelatedMaterial`

GetRelatedMaterials returns the RelatedMaterials field if non-nil, zero value otherwise.

### GetRelatedMaterialsOk

`func (o *CollectionCreate) GetRelatedMaterialsOk() (*[]RelatedMaterial, bool)`

GetRelatedMaterialsOk returns a tuple with the RelatedMaterials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedMaterials

`func (o *CollectionCreate) SetRelatedMaterials(v []RelatedMaterial)`

SetRelatedMaterials sets RelatedMaterials field to given value.

### HasRelatedMaterials

`func (o *CollectionCreate) HasRelatedMaterials() bool`

HasRelatedMaterials returns a boolean if a field has been set.

### GetCustomFields

`func (o *CollectionCreate) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *CollectionCreate) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *CollectionCreate) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *CollectionCreate) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCustomFieldsList

`func (o *CollectionCreate) GetCustomFieldsList() []CustomArticleFieldAdd`

GetCustomFieldsList returns the CustomFieldsList field if non-nil, zero value otherwise.

### GetCustomFieldsListOk

`func (o *CollectionCreate) GetCustomFieldsListOk() (*[]CustomArticleFieldAdd, bool)`

GetCustomFieldsListOk returns a tuple with the CustomFieldsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFieldsList

`func (o *CollectionCreate) SetCustomFieldsList(v []CustomArticleFieldAdd)`

SetCustomFieldsList sets CustomFieldsList field to given value.

### HasCustomFieldsList

`func (o *CollectionCreate) HasCustomFieldsList() bool`

HasCustomFieldsList returns a boolean if a field has been set.

### GetDoi

`func (o *CollectionCreate) GetDoi() string`

GetDoi returns the Doi field if non-nil, zero value otherwise.

### GetDoiOk

`func (o *CollectionCreate) GetDoiOk() (*string, bool)`

GetDoiOk returns a tuple with the Doi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoi

`func (o *CollectionCreate) SetDoi(v string)`

SetDoi sets Doi field to given value.

### HasDoi

`func (o *CollectionCreate) HasDoi() bool`

HasDoi returns a boolean if a field has been set.

### GetHandle

`func (o *CollectionCreate) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *CollectionCreate) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *CollectionCreate) SetHandle(v string)`

SetHandle sets Handle field to given value.

### HasHandle

`func (o *CollectionCreate) HasHandle() bool`

HasHandle returns a boolean if a field has been set.

### GetResourceId

`func (o *CollectionCreate) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *CollectionCreate) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *CollectionCreate) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *CollectionCreate) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetResourceDoi

`func (o *CollectionCreate) GetResourceDoi() string`

GetResourceDoi returns the ResourceDoi field if non-nil, zero value otherwise.

### GetResourceDoiOk

`func (o *CollectionCreate) GetResourceDoiOk() (*string, bool)`

GetResourceDoiOk returns a tuple with the ResourceDoi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceDoi

`func (o *CollectionCreate) SetResourceDoi(v string)`

SetResourceDoi sets ResourceDoi field to given value.

### HasResourceDoi

`func (o *CollectionCreate) HasResourceDoi() bool`

HasResourceDoi returns a boolean if a field has been set.

### GetResourceLink

`func (o *CollectionCreate) GetResourceLink() string`

GetResourceLink returns the ResourceLink field if non-nil, zero value otherwise.

### GetResourceLinkOk

`func (o *CollectionCreate) GetResourceLinkOk() (*string, bool)`

GetResourceLinkOk returns a tuple with the ResourceLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceLink

`func (o *CollectionCreate) SetResourceLink(v string)`

SetResourceLink sets ResourceLink field to given value.

### HasResourceLink

`func (o *CollectionCreate) HasResourceLink() bool`

HasResourceLink returns a boolean if a field has been set.

### GetResourceTitle

`func (o *CollectionCreate) GetResourceTitle() string`

GetResourceTitle returns the ResourceTitle field if non-nil, zero value otherwise.

### GetResourceTitleOk

`func (o *CollectionCreate) GetResourceTitleOk() (*string, bool)`

GetResourceTitleOk returns a tuple with the ResourceTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceTitle

`func (o *CollectionCreate) SetResourceTitle(v string)`

SetResourceTitle sets ResourceTitle field to given value.

### HasResourceTitle

`func (o *CollectionCreate) HasResourceTitle() bool`

HasResourceTitle returns a boolean if a field has been set.

### GetResourceVersion

`func (o *CollectionCreate) GetResourceVersion() int32`

GetResourceVersion returns the ResourceVersion field if non-nil, zero value otherwise.

### GetResourceVersionOk

`func (o *CollectionCreate) GetResourceVersionOk() (*int32, bool)`

GetResourceVersionOk returns a tuple with the ResourceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceVersion

`func (o *CollectionCreate) SetResourceVersion(v int32)`

SetResourceVersion sets ResourceVersion field to given value.

### HasResourceVersion

`func (o *CollectionCreate) HasResourceVersion() bool`

HasResourceVersion returns a boolean if a field has been set.

### GetGroupId

`func (o *CollectionCreate) GetGroupId() int64`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *CollectionCreate) GetGroupIdOk() (*int64, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *CollectionCreate) SetGroupId(v int64)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *CollectionCreate) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetTimeline

`func (o *CollectionCreate) GetTimeline() TimelineUpdate`

GetTimeline returns the Timeline field if non-nil, zero value otherwise.

### GetTimelineOk

`func (o *CollectionCreate) GetTimelineOk() (*TimelineUpdate, bool)`

GetTimelineOk returns a tuple with the Timeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeline

`func (o *CollectionCreate) SetTimeline(v TimelineUpdate)`

SetTimeline sets Timeline field to given value.

### HasTimeline

`func (o *CollectionCreate) HasTimeline() bool`

HasTimeline returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


