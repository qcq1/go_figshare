# ArticleUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authors** | Pointer to **[]map[string]interface{}** | List of authors to be associated with the article. The list can contain the following fields: id, name, first_name, last_name, email, orcid_id. If an id is supplied, it will take priority and everything else will be ignored. For adding more authors use the specific authors endpoint. | [optional] 
**Categories** | Pointer to **[]int64** | List of category ids to be associated with the article(e.g [1, 23, 33, 66]) | [optional] 
**CategoriesBySourceId** | Pointer to **[]string** | List of category source ids to be associated with the article, supersedes the categories property | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** | List of key, values pairs to be associated with the article | [optional] 
**CustomFieldsList** | Pointer to [**[]CustomArticleFieldAdd**](CustomArticleFieldAdd.md) | List of custom fields values, supersedes custom_fields parameter | [optional] 
**DefinedType** | Pointer to **string** | &lt;b&gt;One of:&lt;/b&gt; &lt;code&gt;figure&lt;/code&gt; &lt;code&gt;online resource&lt;/code&gt; &lt;code&gt;preprint&lt;/code&gt; &lt;code&gt;book&lt;/code&gt; &lt;code&gt;conference contribution&lt;/code&gt; &lt;code&gt;media&lt;/code&gt; &lt;code&gt;dataset&lt;/code&gt; &lt;code&gt;poster&lt;/code&gt; &lt;code&gt;journal contribution&lt;/code&gt; &lt;code&gt;presentation&lt;/code&gt; &lt;code&gt;thesis&lt;/code&gt; &lt;code&gt;software&lt;/code&gt; | [optional] 
**Description** | Pointer to **string** | The article description. In a publisher case, usually this is the remote article description | [optional] [default to ""]
**Doi** | Pointer to **string** | Not applicable for regular users. In an institutional case, make sure your group supports setting DOIs. This setting is applied by figshare via opening a ticket through our support/helpdesk system. | [optional] [default to ""]
**DownloadDisabled** | Pointer to **bool** | If true, downloading of files for this article is disabled | [optional] 
**Funding** | Pointer to **string** | Grant number or funding authority | [optional] [default to ""]
**FundingList** | Pointer to [**[]FundingCreate**](FundingCreate.md) | Funding creation / update items | [optional] 
**GroupId** | Pointer to **int64** | Not applicable to regular users. This field is reserved to institutions/publishers with access to assign to specific groups | [optional] 
**Handle** | Pointer to **string** | Not applicable for regular users. In an institutional case, make sure your group supports setting Handles. This setting is applied by figshare via opening a ticket through our support/helpdesk system. | [optional] [default to ""]
**IsMetadataRecord** | Pointer to **bool** | True if article has no files | [optional] 
**Keywords** | Pointer to **[]string** | List of tags to be associated with the article. Tags can be used instead | [optional] 
**License** | Pointer to **int64** | License id for this article. | [optional] [default to 0]
**MetadataReason** | Pointer to **string** | Article metadata reason | [optional] 
**References** | Pointer to **[]string** | List of links to be associated with the article (e.g [\&quot;http://link1\&quot;, \&quot;http://link2\&quot;, \&quot;http://link3\&quot;]) | [optional] 
**RelatedMaterials** | Pointer to [**[]RelatedMaterial**](RelatedMaterial.md) | List of related materials; supersedes references and resource DOI/title. | [optional] 
**ResourceDoi** | Pointer to **string** | Deprecated by related materials. Not applicable to regular users. In a publisher case, this is the publisher article DOI. | [optional] [default to ""]
**ResourceTitle** | Pointer to **string** | Deprecated by related materials. Not applicable to regular users. In a publisher case, this is the publisher article title. | [optional] [default to ""]
**Tags** | Pointer to **[]string** | List of tags to be associated with the article. Keywords can be used instead | [optional] 
**Timeline** | Pointer to [**TimelineUpdate**](TimelineUpdate.md) |  | [optional] 
**Title** | Pointer to **string** | Title of article | [optional] 

## Methods

### NewArticleUpdate

`func NewArticleUpdate() *ArticleUpdate`

NewArticleUpdate instantiates a new ArticleUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArticleUpdateWithDefaults

`func NewArticleUpdateWithDefaults() *ArticleUpdate`

NewArticleUpdateWithDefaults instantiates a new ArticleUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthors

`func (o *ArticleUpdate) GetAuthors() []map[string]interface{}`

GetAuthors returns the Authors field if non-nil, zero value otherwise.

### GetAuthorsOk

`func (o *ArticleUpdate) GetAuthorsOk() (*[]map[string]interface{}, bool)`

GetAuthorsOk returns a tuple with the Authors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthors

`func (o *ArticleUpdate) SetAuthors(v []map[string]interface{})`

SetAuthors sets Authors field to given value.

### HasAuthors

`func (o *ArticleUpdate) HasAuthors() bool`

HasAuthors returns a boolean if a field has been set.

### GetCategories

`func (o *ArticleUpdate) GetCategories() []int64`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *ArticleUpdate) GetCategoriesOk() (*[]int64, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *ArticleUpdate) SetCategories(v []int64)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *ArticleUpdate) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetCategoriesBySourceId

`func (o *ArticleUpdate) GetCategoriesBySourceId() []string`

GetCategoriesBySourceId returns the CategoriesBySourceId field if non-nil, zero value otherwise.

### GetCategoriesBySourceIdOk

`func (o *ArticleUpdate) GetCategoriesBySourceIdOk() (*[]string, bool)`

GetCategoriesBySourceIdOk returns a tuple with the CategoriesBySourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoriesBySourceId

`func (o *ArticleUpdate) SetCategoriesBySourceId(v []string)`

SetCategoriesBySourceId sets CategoriesBySourceId field to given value.

### HasCategoriesBySourceId

`func (o *ArticleUpdate) HasCategoriesBySourceId() bool`

HasCategoriesBySourceId returns a boolean if a field has been set.

### GetCustomFields

`func (o *ArticleUpdate) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ArticleUpdate) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ArticleUpdate) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ArticleUpdate) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCustomFieldsList

`func (o *ArticleUpdate) GetCustomFieldsList() []CustomArticleFieldAdd`

GetCustomFieldsList returns the CustomFieldsList field if non-nil, zero value otherwise.

### GetCustomFieldsListOk

`func (o *ArticleUpdate) GetCustomFieldsListOk() (*[]CustomArticleFieldAdd, bool)`

GetCustomFieldsListOk returns a tuple with the CustomFieldsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFieldsList

`func (o *ArticleUpdate) SetCustomFieldsList(v []CustomArticleFieldAdd)`

SetCustomFieldsList sets CustomFieldsList field to given value.

### HasCustomFieldsList

`func (o *ArticleUpdate) HasCustomFieldsList() bool`

HasCustomFieldsList returns a boolean if a field has been set.

### GetDefinedType

`func (o *ArticleUpdate) GetDefinedType() string`

GetDefinedType returns the DefinedType field if non-nil, zero value otherwise.

### GetDefinedTypeOk

`func (o *ArticleUpdate) GetDefinedTypeOk() (*string, bool)`

GetDefinedTypeOk returns a tuple with the DefinedType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinedType

`func (o *ArticleUpdate) SetDefinedType(v string)`

SetDefinedType sets DefinedType field to given value.

### HasDefinedType

`func (o *ArticleUpdate) HasDefinedType() bool`

HasDefinedType returns a boolean if a field has been set.

### GetDescription

`func (o *ArticleUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ArticleUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ArticleUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ArticleUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDoi

`func (o *ArticleUpdate) GetDoi() string`

GetDoi returns the Doi field if non-nil, zero value otherwise.

### GetDoiOk

`func (o *ArticleUpdate) GetDoiOk() (*string, bool)`

GetDoiOk returns a tuple with the Doi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoi

`func (o *ArticleUpdate) SetDoi(v string)`

SetDoi sets Doi field to given value.

### HasDoi

`func (o *ArticleUpdate) HasDoi() bool`

HasDoi returns a boolean if a field has been set.

### GetDownloadDisabled

`func (o *ArticleUpdate) GetDownloadDisabled() bool`

GetDownloadDisabled returns the DownloadDisabled field if non-nil, zero value otherwise.

### GetDownloadDisabledOk

`func (o *ArticleUpdate) GetDownloadDisabledOk() (*bool, bool)`

GetDownloadDisabledOk returns a tuple with the DownloadDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadDisabled

`func (o *ArticleUpdate) SetDownloadDisabled(v bool)`

SetDownloadDisabled sets DownloadDisabled field to given value.

### HasDownloadDisabled

`func (o *ArticleUpdate) HasDownloadDisabled() bool`

HasDownloadDisabled returns a boolean if a field has been set.

### GetFunding

`func (o *ArticleUpdate) GetFunding() string`

GetFunding returns the Funding field if non-nil, zero value otherwise.

### GetFundingOk

`func (o *ArticleUpdate) GetFundingOk() (*string, bool)`

GetFundingOk returns a tuple with the Funding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunding

`func (o *ArticleUpdate) SetFunding(v string)`

SetFunding sets Funding field to given value.

### HasFunding

`func (o *ArticleUpdate) HasFunding() bool`

HasFunding returns a boolean if a field has been set.

### GetFundingList

`func (o *ArticleUpdate) GetFundingList() []FundingCreate`

GetFundingList returns the FundingList field if non-nil, zero value otherwise.

### GetFundingListOk

`func (o *ArticleUpdate) GetFundingListOk() (*[]FundingCreate, bool)`

GetFundingListOk returns a tuple with the FundingList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingList

`func (o *ArticleUpdate) SetFundingList(v []FundingCreate)`

SetFundingList sets FundingList field to given value.

### HasFundingList

`func (o *ArticleUpdate) HasFundingList() bool`

HasFundingList returns a boolean if a field has been set.

### GetGroupId

`func (o *ArticleUpdate) GetGroupId() int64`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ArticleUpdate) GetGroupIdOk() (*int64, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ArticleUpdate) SetGroupId(v int64)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ArticleUpdate) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetHandle

`func (o *ArticleUpdate) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *ArticleUpdate) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *ArticleUpdate) SetHandle(v string)`

SetHandle sets Handle field to given value.

### HasHandle

`func (o *ArticleUpdate) HasHandle() bool`

HasHandle returns a boolean if a field has been set.

### GetIsMetadataRecord

`func (o *ArticleUpdate) GetIsMetadataRecord() bool`

GetIsMetadataRecord returns the IsMetadataRecord field if non-nil, zero value otherwise.

### GetIsMetadataRecordOk

`func (o *ArticleUpdate) GetIsMetadataRecordOk() (*bool, bool)`

GetIsMetadataRecordOk returns a tuple with the IsMetadataRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMetadataRecord

`func (o *ArticleUpdate) SetIsMetadataRecord(v bool)`

SetIsMetadataRecord sets IsMetadataRecord field to given value.

### HasIsMetadataRecord

`func (o *ArticleUpdate) HasIsMetadataRecord() bool`

HasIsMetadataRecord returns a boolean if a field has been set.

### GetKeywords

`func (o *ArticleUpdate) GetKeywords() []string`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *ArticleUpdate) GetKeywordsOk() (*[]string, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *ArticleUpdate) SetKeywords(v []string)`

SetKeywords sets Keywords field to given value.

### HasKeywords

`func (o *ArticleUpdate) HasKeywords() bool`

HasKeywords returns a boolean if a field has been set.

### GetLicense

`func (o *ArticleUpdate) GetLicense() int64`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *ArticleUpdate) GetLicenseOk() (*int64, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *ArticleUpdate) SetLicense(v int64)`

SetLicense sets License field to given value.

### HasLicense

`func (o *ArticleUpdate) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetMetadataReason

`func (o *ArticleUpdate) GetMetadataReason() string`

GetMetadataReason returns the MetadataReason field if non-nil, zero value otherwise.

### GetMetadataReasonOk

`func (o *ArticleUpdate) GetMetadataReasonOk() (*string, bool)`

GetMetadataReasonOk returns a tuple with the MetadataReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataReason

`func (o *ArticleUpdate) SetMetadataReason(v string)`

SetMetadataReason sets MetadataReason field to given value.

### HasMetadataReason

`func (o *ArticleUpdate) HasMetadataReason() bool`

HasMetadataReason returns a boolean if a field has been set.

### GetReferences

`func (o *ArticleUpdate) GetReferences() []string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *ArticleUpdate) GetReferencesOk() (*[]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *ArticleUpdate) SetReferences(v []string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *ArticleUpdate) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetRelatedMaterials

`func (o *ArticleUpdate) GetRelatedMaterials() []RelatedMaterial`

GetRelatedMaterials returns the RelatedMaterials field if non-nil, zero value otherwise.

### GetRelatedMaterialsOk

`func (o *ArticleUpdate) GetRelatedMaterialsOk() (*[]RelatedMaterial, bool)`

GetRelatedMaterialsOk returns a tuple with the RelatedMaterials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedMaterials

`func (o *ArticleUpdate) SetRelatedMaterials(v []RelatedMaterial)`

SetRelatedMaterials sets RelatedMaterials field to given value.

### HasRelatedMaterials

`func (o *ArticleUpdate) HasRelatedMaterials() bool`

HasRelatedMaterials returns a boolean if a field has been set.

### GetResourceDoi

`func (o *ArticleUpdate) GetResourceDoi() string`

GetResourceDoi returns the ResourceDoi field if non-nil, zero value otherwise.

### GetResourceDoiOk

`func (o *ArticleUpdate) GetResourceDoiOk() (*string, bool)`

GetResourceDoiOk returns a tuple with the ResourceDoi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceDoi

`func (o *ArticleUpdate) SetResourceDoi(v string)`

SetResourceDoi sets ResourceDoi field to given value.

### HasResourceDoi

`func (o *ArticleUpdate) HasResourceDoi() bool`

HasResourceDoi returns a boolean if a field has been set.

### GetResourceTitle

`func (o *ArticleUpdate) GetResourceTitle() string`

GetResourceTitle returns the ResourceTitle field if non-nil, zero value otherwise.

### GetResourceTitleOk

`func (o *ArticleUpdate) GetResourceTitleOk() (*string, bool)`

GetResourceTitleOk returns a tuple with the ResourceTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceTitle

`func (o *ArticleUpdate) SetResourceTitle(v string)`

SetResourceTitle sets ResourceTitle field to given value.

### HasResourceTitle

`func (o *ArticleUpdate) HasResourceTitle() bool`

HasResourceTitle returns a boolean if a field has been set.

### GetTags

`func (o *ArticleUpdate) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ArticleUpdate) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ArticleUpdate) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ArticleUpdate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTimeline

`func (o *ArticleUpdate) GetTimeline() TimelineUpdate`

GetTimeline returns the Timeline field if non-nil, zero value otherwise.

### GetTimelineOk

`func (o *ArticleUpdate) GetTimelineOk() (*TimelineUpdate, bool)`

GetTimelineOk returns a tuple with the Timeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeline

`func (o *ArticleUpdate) SetTimeline(v TimelineUpdate)`

SetTimeline sets Timeline field to given value.

### HasTimeline

`func (o *ArticleUpdate) HasTimeline() bool`

HasTimeline returns a boolean if a field has been set.

### GetTitle

`func (o *ArticleUpdate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ArticleUpdate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ArticleUpdate) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ArticleUpdate) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


