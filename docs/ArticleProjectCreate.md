# ArticleProjectCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | Title of article | 
**Description** | Pointer to **string** | The article description. In a publisher case, usually this is the remote article description | [optional] [default to ""]
**Tags** | Pointer to **[]string** | List of tags to be associated with the article. Keywords can be used instead | [optional] 
**Keywords** | Pointer to **[]string** | List of tags to be associated with the article. Tags can be used instead | [optional] 
**References** | Pointer to **[]string** | List of links to be associated with the article (e.g [\&quot;http://link1\&quot;, \&quot;http://link2\&quot;, \&quot;http://link3\&quot;]) | [optional] 
**RelatedMaterials** | Pointer to [**[]RelatedMaterial**](RelatedMaterial.md) | List of related materials; supersedes references and resource DOI/title. | [optional] 
**Categories** | Pointer to **[]int64** | List of category ids to be associated with the article(e.g [1, 23, 33, 66]) | [optional] 
**CategoriesBySourceId** | Pointer to **[]string** | List of category source ids to be associated with the article, supersedes the categories property | [optional] 
**Authors** | Pointer to **[]map[string]interface{}** | List of authors to be associated with the article. The list can contain the following fields: id, name, first_name, last_name, email, orcid_id. If an id is supplied, it will take priority and everything else will be ignored. For adding more authors use the specific authors endpoint. | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** | List of key, values pairs to be associated with the article | [optional] 
**CustomFieldsList** | Pointer to [**[]CustomArticleFieldAdd**](CustomArticleFieldAdd.md) | List of custom fields values, supersedes custom_fields parameter | [optional] 
**DefinedType** | Pointer to **string** | &lt;b&gt;One of:&lt;/b&gt; &lt;code&gt;figure&lt;/code&gt; &lt;code&gt;online resource&lt;/code&gt; &lt;code&gt;preprint&lt;/code&gt; &lt;code&gt;book&lt;/code&gt; &lt;code&gt;conference contribution&lt;/code&gt; &lt;code&gt;media&lt;/code&gt; &lt;code&gt;dataset&lt;/code&gt; &lt;code&gt;poster&lt;/code&gt; &lt;code&gt;journal contribution&lt;/code&gt; &lt;code&gt;presentation&lt;/code&gt; &lt;code&gt;thesis&lt;/code&gt; &lt;code&gt;software&lt;/code&gt; | [optional] 
**Funding** | Pointer to **string** | Grant number or funding authority | [optional] [default to ""]
**FundingList** | Pointer to [**[]FundingCreate**](FundingCreate.md) | Funding creation / update items | [optional] 
**License** | Pointer to **int64** | License id for this article. | [optional] [default to 0]
**Doi** | Pointer to **string** | Not applicable for regular users. In an institutional case, make sure your group supports setting DOIs. This setting is applied by figshare via opening a ticket through our support/helpdesk system. | [optional] [default to ""]
**Handle** | Pointer to **string** | Not applicable for regular users. In an institutional case, make sure your group supports setting Handles. This setting is applied by figshare via opening a ticket through our support/helpdesk system. | [optional] [default to ""]
**ResourceDoi** | Pointer to **string** | Deprecated by related materials. Not applicable to regular users. In a publisher case, this is the publisher article DOI. | [optional] [default to ""]
**ResourceTitle** | Pointer to **string** | Deprecated by related materials. Not applicable to regular users. In a publisher case, this is the publisher article title. | [optional] [default to ""]
**Timeline** | Pointer to [**TimelineUpdate**](TimelineUpdate.md) |  | [optional] 

## Methods

### NewArticleProjectCreate

`func NewArticleProjectCreate(title string, ) *ArticleProjectCreate`

NewArticleProjectCreate instantiates a new ArticleProjectCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArticleProjectCreateWithDefaults

`func NewArticleProjectCreateWithDefaults() *ArticleProjectCreate`

NewArticleProjectCreateWithDefaults instantiates a new ArticleProjectCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *ArticleProjectCreate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ArticleProjectCreate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ArticleProjectCreate) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *ArticleProjectCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ArticleProjectCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ArticleProjectCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ArticleProjectCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *ArticleProjectCreate) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ArticleProjectCreate) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ArticleProjectCreate) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ArticleProjectCreate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetKeywords

`func (o *ArticleProjectCreate) GetKeywords() []string`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *ArticleProjectCreate) GetKeywordsOk() (*[]string, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *ArticleProjectCreate) SetKeywords(v []string)`

SetKeywords sets Keywords field to given value.

### HasKeywords

`func (o *ArticleProjectCreate) HasKeywords() bool`

HasKeywords returns a boolean if a field has been set.

### GetReferences

`func (o *ArticleProjectCreate) GetReferences() []string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *ArticleProjectCreate) GetReferencesOk() (*[]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *ArticleProjectCreate) SetReferences(v []string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *ArticleProjectCreate) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetRelatedMaterials

`func (o *ArticleProjectCreate) GetRelatedMaterials() []RelatedMaterial`

GetRelatedMaterials returns the RelatedMaterials field if non-nil, zero value otherwise.

### GetRelatedMaterialsOk

`func (o *ArticleProjectCreate) GetRelatedMaterialsOk() (*[]RelatedMaterial, bool)`

GetRelatedMaterialsOk returns a tuple with the RelatedMaterials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedMaterials

`func (o *ArticleProjectCreate) SetRelatedMaterials(v []RelatedMaterial)`

SetRelatedMaterials sets RelatedMaterials field to given value.

### HasRelatedMaterials

`func (o *ArticleProjectCreate) HasRelatedMaterials() bool`

HasRelatedMaterials returns a boolean if a field has been set.

### GetCategories

`func (o *ArticleProjectCreate) GetCategories() []int64`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *ArticleProjectCreate) GetCategoriesOk() (*[]int64, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *ArticleProjectCreate) SetCategories(v []int64)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *ArticleProjectCreate) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetCategoriesBySourceId

`func (o *ArticleProjectCreate) GetCategoriesBySourceId() []string`

GetCategoriesBySourceId returns the CategoriesBySourceId field if non-nil, zero value otherwise.

### GetCategoriesBySourceIdOk

`func (o *ArticleProjectCreate) GetCategoriesBySourceIdOk() (*[]string, bool)`

GetCategoriesBySourceIdOk returns a tuple with the CategoriesBySourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoriesBySourceId

`func (o *ArticleProjectCreate) SetCategoriesBySourceId(v []string)`

SetCategoriesBySourceId sets CategoriesBySourceId field to given value.

### HasCategoriesBySourceId

`func (o *ArticleProjectCreate) HasCategoriesBySourceId() bool`

HasCategoriesBySourceId returns a boolean if a field has been set.

### GetAuthors

`func (o *ArticleProjectCreate) GetAuthors() []map[string]interface{}`

GetAuthors returns the Authors field if non-nil, zero value otherwise.

### GetAuthorsOk

`func (o *ArticleProjectCreate) GetAuthorsOk() (*[]map[string]interface{}, bool)`

GetAuthorsOk returns a tuple with the Authors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthors

`func (o *ArticleProjectCreate) SetAuthors(v []map[string]interface{})`

SetAuthors sets Authors field to given value.

### HasAuthors

`func (o *ArticleProjectCreate) HasAuthors() bool`

HasAuthors returns a boolean if a field has been set.

### GetCustomFields

`func (o *ArticleProjectCreate) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ArticleProjectCreate) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ArticleProjectCreate) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ArticleProjectCreate) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCustomFieldsList

`func (o *ArticleProjectCreate) GetCustomFieldsList() []CustomArticleFieldAdd`

GetCustomFieldsList returns the CustomFieldsList field if non-nil, zero value otherwise.

### GetCustomFieldsListOk

`func (o *ArticleProjectCreate) GetCustomFieldsListOk() (*[]CustomArticleFieldAdd, bool)`

GetCustomFieldsListOk returns a tuple with the CustomFieldsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFieldsList

`func (o *ArticleProjectCreate) SetCustomFieldsList(v []CustomArticleFieldAdd)`

SetCustomFieldsList sets CustomFieldsList field to given value.

### HasCustomFieldsList

`func (o *ArticleProjectCreate) HasCustomFieldsList() bool`

HasCustomFieldsList returns a boolean if a field has been set.

### GetDefinedType

`func (o *ArticleProjectCreate) GetDefinedType() string`

GetDefinedType returns the DefinedType field if non-nil, zero value otherwise.

### GetDefinedTypeOk

`func (o *ArticleProjectCreate) GetDefinedTypeOk() (*string, bool)`

GetDefinedTypeOk returns a tuple with the DefinedType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinedType

`func (o *ArticleProjectCreate) SetDefinedType(v string)`

SetDefinedType sets DefinedType field to given value.

### HasDefinedType

`func (o *ArticleProjectCreate) HasDefinedType() bool`

HasDefinedType returns a boolean if a field has been set.

### GetFunding

`func (o *ArticleProjectCreate) GetFunding() string`

GetFunding returns the Funding field if non-nil, zero value otherwise.

### GetFundingOk

`func (o *ArticleProjectCreate) GetFundingOk() (*string, bool)`

GetFundingOk returns a tuple with the Funding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunding

`func (o *ArticleProjectCreate) SetFunding(v string)`

SetFunding sets Funding field to given value.

### HasFunding

`func (o *ArticleProjectCreate) HasFunding() bool`

HasFunding returns a boolean if a field has been set.

### GetFundingList

`func (o *ArticleProjectCreate) GetFundingList() []FundingCreate`

GetFundingList returns the FundingList field if non-nil, zero value otherwise.

### GetFundingListOk

`func (o *ArticleProjectCreate) GetFundingListOk() (*[]FundingCreate, bool)`

GetFundingListOk returns a tuple with the FundingList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingList

`func (o *ArticleProjectCreate) SetFundingList(v []FundingCreate)`

SetFundingList sets FundingList field to given value.

### HasFundingList

`func (o *ArticleProjectCreate) HasFundingList() bool`

HasFundingList returns a boolean if a field has been set.

### GetLicense

`func (o *ArticleProjectCreate) GetLicense() int64`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *ArticleProjectCreate) GetLicenseOk() (*int64, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *ArticleProjectCreate) SetLicense(v int64)`

SetLicense sets License field to given value.

### HasLicense

`func (o *ArticleProjectCreate) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetDoi

`func (o *ArticleProjectCreate) GetDoi() string`

GetDoi returns the Doi field if non-nil, zero value otherwise.

### GetDoiOk

`func (o *ArticleProjectCreate) GetDoiOk() (*string, bool)`

GetDoiOk returns a tuple with the Doi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoi

`func (o *ArticleProjectCreate) SetDoi(v string)`

SetDoi sets Doi field to given value.

### HasDoi

`func (o *ArticleProjectCreate) HasDoi() bool`

HasDoi returns a boolean if a field has been set.

### GetHandle

`func (o *ArticleProjectCreate) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *ArticleProjectCreate) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *ArticleProjectCreate) SetHandle(v string)`

SetHandle sets Handle field to given value.

### HasHandle

`func (o *ArticleProjectCreate) HasHandle() bool`

HasHandle returns a boolean if a field has been set.

### GetResourceDoi

`func (o *ArticleProjectCreate) GetResourceDoi() string`

GetResourceDoi returns the ResourceDoi field if non-nil, zero value otherwise.

### GetResourceDoiOk

`func (o *ArticleProjectCreate) GetResourceDoiOk() (*string, bool)`

GetResourceDoiOk returns a tuple with the ResourceDoi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceDoi

`func (o *ArticleProjectCreate) SetResourceDoi(v string)`

SetResourceDoi sets ResourceDoi field to given value.

### HasResourceDoi

`func (o *ArticleProjectCreate) HasResourceDoi() bool`

HasResourceDoi returns a boolean if a field has been set.

### GetResourceTitle

`func (o *ArticleProjectCreate) GetResourceTitle() string`

GetResourceTitle returns the ResourceTitle field if non-nil, zero value otherwise.

### GetResourceTitleOk

`func (o *ArticleProjectCreate) GetResourceTitleOk() (*string, bool)`

GetResourceTitleOk returns a tuple with the ResourceTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceTitle

`func (o *ArticleProjectCreate) SetResourceTitle(v string)`

SetResourceTitle sets ResourceTitle field to given value.

### HasResourceTitle

`func (o *ArticleProjectCreate) HasResourceTitle() bool`

HasResourceTitle returns a boolean if a field has been set.

### GetTimeline

`func (o *ArticleProjectCreate) GetTimeline() TimelineUpdate`

GetTimeline returns the Timeline field if non-nil, zero value otherwise.

### GetTimelineOk

`func (o *ArticleProjectCreate) GetTimelineOk() (*TimelineUpdate, bool)`

GetTimelineOk returns a tuple with the Timeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeline

`func (o *ArticleProjectCreate) SetTimeline(v TimelineUpdate)`

SetTimeline sets Timeline field to given value.

### HasTimeline

`func (o *ArticleProjectCreate) HasTimeline() bool`

HasTimeline returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


