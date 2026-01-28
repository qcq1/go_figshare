# ArticleComplete

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authors** | [**[]Author**](Author.md) | List of article authors | 
**CustomFields** | [**[]CustomArticleField**](CustomArticleField.md) | List of custom fields values | 
**DownloadDisabled** | **bool** | If true, downloading of files for this article is disabled | 
**EmbargoOptions** | [**[]GroupEmbargoOptions**](GroupEmbargoOptions.md) | List of embargo options | 
**FigshareUrl** | **string** | Article public url | 
**Files** | [**[]PublicFile**](PublicFile.md) | List of article files | 
**FolderStructure** | **map[string]interface{}** | Mapping of file ids to folder paths, if folders are used | 
**Categories** | [**[]Category**](Category.md) | List of categories selected for the article | 
**Citation** | **string** | Article citation | 
**ConfidentialReason** | **string** | Confidentiality reason | 
**CreatedDate** | **string** | Date when article was created | 
**Description** | **string** | Article description | 
**EmbargoReason** | **string** | Reason for embargo | 
**EmbargoTitle** | **string** | Title for embargo | 
**Funding** | **string** | Article funding | 
**FundingList** | [**[]FundingInformation**](FundingInformation.md) | Full Article funding information | 
**HasLinkedFile** | **bool** | True if any files are linked to the article | 
**IsConfidential** | **bool** | Article Confidentiality | 
**IsEmbargoed** | **bool** | True if article is embargoed | 
**IsMetadataRecord** | **bool** | True if article has no files | 
**IsPublic** | **bool** | True if article is published | 
**Keywords** | **[]string** | List of article keywords. Tags can be used instead | 
**License** | [**License**](License.md) |  | 
**MetadataReason** | **string** | Article metadata reason | 
**References** | **[]string** | List of references | 
**RelatedMaterials** | Pointer to [**[]RelatedMaterial**](RelatedMaterial.md) | List of related materials; supersedes references and resource DOI/title. | [optional] 
**Size** | **int64** | Article size | 
**Status** | **string** | Article status | 
**Tags** | **[]string** | List of article tags. Keywords can be used instead | 
**Version** | **int64** | Article version | 
**DefinedType** | **int64** | Type of article identifier | 
**DefinedTypeName** | **string** | Name of the article type identifier | 
**Doi** | **string** | DOI | 
**Handle** | **string** | Handle | 
**Id** | **int64** | Unique identifier for article | 
**ResourceDoi** | **string** | Deprecated by related materials. Not applicable to regular users. In a publisher case, this is the publisher article DOI. | [default to ""]
**ResourceTitle** | **string** | Deprecated by related materials. Not applicable to regular users. In a publisher case, this is the publisher article title. | [default to ""]
**Thumb** | **string** | Thumbnail image | 
**Timeline** | [**Timeline**](Timeline.md) |  | 
**Title** | **string** | Title of article | 
**Url** | **string** | Api endpoint for article | 
**UrlPrivateApi** | **string** | Private Api endpoint for article | 
**UrlPrivateHtml** | **string** | Private site endpoint for article | 
**UrlPublicApi** | **string** | Public Api endpoint for article | 
**UrlPublicHtml** | **string** | Public site endpoint for article | 

## Methods

### NewArticleComplete

`func NewArticleComplete(authors []Author, customFields []CustomArticleField, downloadDisabled bool, embargoOptions []GroupEmbargoOptions, figshareUrl string, files []PublicFile, folderStructure map[string]interface{}, categories []Category, citation string, confidentialReason string, createdDate string, description string, embargoReason string, embargoTitle string, funding string, fundingList []FundingInformation, hasLinkedFile bool, isConfidential bool, isEmbargoed bool, isMetadataRecord bool, isPublic bool, keywords []string, license License, metadataReason string, references []string, size int64, status string, tags []string, version int64, definedType int64, definedTypeName string, doi string, handle string, id int64, resourceDoi string, resourceTitle string, thumb string, timeline Timeline, title string, url string, urlPrivateApi string, urlPrivateHtml string, urlPublicApi string, urlPublicHtml string, ) *ArticleComplete`

NewArticleComplete instantiates a new ArticleComplete object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArticleCompleteWithDefaults

`func NewArticleCompleteWithDefaults() *ArticleComplete`

NewArticleCompleteWithDefaults instantiates a new ArticleComplete object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthors

`func (o *ArticleComplete) GetAuthors() []Author`

GetAuthors returns the Authors field if non-nil, zero value otherwise.

### GetAuthorsOk

`func (o *ArticleComplete) GetAuthorsOk() (*[]Author, bool)`

GetAuthorsOk returns a tuple with the Authors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthors

`func (o *ArticleComplete) SetAuthors(v []Author)`

SetAuthors sets Authors field to given value.


### GetCustomFields

`func (o *ArticleComplete) GetCustomFields() []CustomArticleField`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ArticleComplete) GetCustomFieldsOk() (*[]CustomArticleField, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ArticleComplete) SetCustomFields(v []CustomArticleField)`

SetCustomFields sets CustomFields field to given value.


### GetDownloadDisabled

`func (o *ArticleComplete) GetDownloadDisabled() bool`

GetDownloadDisabled returns the DownloadDisabled field if non-nil, zero value otherwise.

### GetDownloadDisabledOk

`func (o *ArticleComplete) GetDownloadDisabledOk() (*bool, bool)`

GetDownloadDisabledOk returns a tuple with the DownloadDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadDisabled

`func (o *ArticleComplete) SetDownloadDisabled(v bool)`

SetDownloadDisabled sets DownloadDisabled field to given value.


### GetEmbargoOptions

`func (o *ArticleComplete) GetEmbargoOptions() []GroupEmbargoOptions`

GetEmbargoOptions returns the EmbargoOptions field if non-nil, zero value otherwise.

### GetEmbargoOptionsOk

`func (o *ArticleComplete) GetEmbargoOptionsOk() (*[]GroupEmbargoOptions, bool)`

GetEmbargoOptionsOk returns a tuple with the EmbargoOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbargoOptions

`func (o *ArticleComplete) SetEmbargoOptions(v []GroupEmbargoOptions)`

SetEmbargoOptions sets EmbargoOptions field to given value.


### GetFigshareUrl

`func (o *ArticleComplete) GetFigshareUrl() string`

GetFigshareUrl returns the FigshareUrl field if non-nil, zero value otherwise.

### GetFigshareUrlOk

`func (o *ArticleComplete) GetFigshareUrlOk() (*string, bool)`

GetFigshareUrlOk returns a tuple with the FigshareUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFigshareUrl

`func (o *ArticleComplete) SetFigshareUrl(v string)`

SetFigshareUrl sets FigshareUrl field to given value.


### GetFiles

`func (o *ArticleComplete) GetFiles() []PublicFile`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *ArticleComplete) GetFilesOk() (*[]PublicFile, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *ArticleComplete) SetFiles(v []PublicFile)`

SetFiles sets Files field to given value.


### GetFolderStructure

`func (o *ArticleComplete) GetFolderStructure() map[string]interface{}`

GetFolderStructure returns the FolderStructure field if non-nil, zero value otherwise.

### GetFolderStructureOk

`func (o *ArticleComplete) GetFolderStructureOk() (*map[string]interface{}, bool)`

GetFolderStructureOk returns a tuple with the FolderStructure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderStructure

`func (o *ArticleComplete) SetFolderStructure(v map[string]interface{})`

SetFolderStructure sets FolderStructure field to given value.


### GetCategories

`func (o *ArticleComplete) GetCategories() []Category`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *ArticleComplete) GetCategoriesOk() (*[]Category, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *ArticleComplete) SetCategories(v []Category)`

SetCategories sets Categories field to given value.


### GetCitation

`func (o *ArticleComplete) GetCitation() string`

GetCitation returns the Citation field if non-nil, zero value otherwise.

### GetCitationOk

`func (o *ArticleComplete) GetCitationOk() (*string, bool)`

GetCitationOk returns a tuple with the Citation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitation

`func (o *ArticleComplete) SetCitation(v string)`

SetCitation sets Citation field to given value.


### GetConfidentialReason

`func (o *ArticleComplete) GetConfidentialReason() string`

GetConfidentialReason returns the ConfidentialReason field if non-nil, zero value otherwise.

### GetConfidentialReasonOk

`func (o *ArticleComplete) GetConfidentialReasonOk() (*string, bool)`

GetConfidentialReasonOk returns a tuple with the ConfidentialReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidentialReason

`func (o *ArticleComplete) SetConfidentialReason(v string)`

SetConfidentialReason sets ConfidentialReason field to given value.


### GetCreatedDate

`func (o *ArticleComplete) GetCreatedDate() string`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *ArticleComplete) GetCreatedDateOk() (*string, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *ArticleComplete) SetCreatedDate(v string)`

SetCreatedDate sets CreatedDate field to given value.


### GetDescription

`func (o *ArticleComplete) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ArticleComplete) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ArticleComplete) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEmbargoReason

`func (o *ArticleComplete) GetEmbargoReason() string`

GetEmbargoReason returns the EmbargoReason field if non-nil, zero value otherwise.

### GetEmbargoReasonOk

`func (o *ArticleComplete) GetEmbargoReasonOk() (*string, bool)`

GetEmbargoReasonOk returns a tuple with the EmbargoReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbargoReason

`func (o *ArticleComplete) SetEmbargoReason(v string)`

SetEmbargoReason sets EmbargoReason field to given value.


### GetEmbargoTitle

`func (o *ArticleComplete) GetEmbargoTitle() string`

GetEmbargoTitle returns the EmbargoTitle field if non-nil, zero value otherwise.

### GetEmbargoTitleOk

`func (o *ArticleComplete) GetEmbargoTitleOk() (*string, bool)`

GetEmbargoTitleOk returns a tuple with the EmbargoTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbargoTitle

`func (o *ArticleComplete) SetEmbargoTitle(v string)`

SetEmbargoTitle sets EmbargoTitle field to given value.


### GetFunding

`func (o *ArticleComplete) GetFunding() string`

GetFunding returns the Funding field if non-nil, zero value otherwise.

### GetFundingOk

`func (o *ArticleComplete) GetFundingOk() (*string, bool)`

GetFundingOk returns a tuple with the Funding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunding

`func (o *ArticleComplete) SetFunding(v string)`

SetFunding sets Funding field to given value.


### GetFundingList

`func (o *ArticleComplete) GetFundingList() []FundingInformation`

GetFundingList returns the FundingList field if non-nil, zero value otherwise.

### GetFundingListOk

`func (o *ArticleComplete) GetFundingListOk() (*[]FundingInformation, bool)`

GetFundingListOk returns a tuple with the FundingList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingList

`func (o *ArticleComplete) SetFundingList(v []FundingInformation)`

SetFundingList sets FundingList field to given value.


### GetHasLinkedFile

`func (o *ArticleComplete) GetHasLinkedFile() bool`

GetHasLinkedFile returns the HasLinkedFile field if non-nil, zero value otherwise.

### GetHasLinkedFileOk

`func (o *ArticleComplete) GetHasLinkedFileOk() (*bool, bool)`

GetHasLinkedFileOk returns a tuple with the HasLinkedFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasLinkedFile

`func (o *ArticleComplete) SetHasLinkedFile(v bool)`

SetHasLinkedFile sets HasLinkedFile field to given value.


### GetIsConfidential

`func (o *ArticleComplete) GetIsConfidential() bool`

GetIsConfidential returns the IsConfidential field if non-nil, zero value otherwise.

### GetIsConfidentialOk

`func (o *ArticleComplete) GetIsConfidentialOk() (*bool, bool)`

GetIsConfidentialOk returns a tuple with the IsConfidential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConfidential

`func (o *ArticleComplete) SetIsConfidential(v bool)`

SetIsConfidential sets IsConfidential field to given value.


### GetIsEmbargoed

`func (o *ArticleComplete) GetIsEmbargoed() bool`

GetIsEmbargoed returns the IsEmbargoed field if non-nil, zero value otherwise.

### GetIsEmbargoedOk

`func (o *ArticleComplete) GetIsEmbargoedOk() (*bool, bool)`

GetIsEmbargoedOk returns a tuple with the IsEmbargoed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEmbargoed

`func (o *ArticleComplete) SetIsEmbargoed(v bool)`

SetIsEmbargoed sets IsEmbargoed field to given value.


### GetIsMetadataRecord

`func (o *ArticleComplete) GetIsMetadataRecord() bool`

GetIsMetadataRecord returns the IsMetadataRecord field if non-nil, zero value otherwise.

### GetIsMetadataRecordOk

`func (o *ArticleComplete) GetIsMetadataRecordOk() (*bool, bool)`

GetIsMetadataRecordOk returns a tuple with the IsMetadataRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMetadataRecord

`func (o *ArticleComplete) SetIsMetadataRecord(v bool)`

SetIsMetadataRecord sets IsMetadataRecord field to given value.


### GetIsPublic

`func (o *ArticleComplete) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *ArticleComplete) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *ArticleComplete) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.


### GetKeywords

`func (o *ArticleComplete) GetKeywords() []string`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *ArticleComplete) GetKeywordsOk() (*[]string, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *ArticleComplete) SetKeywords(v []string)`

SetKeywords sets Keywords field to given value.


### GetLicense

`func (o *ArticleComplete) GetLicense() License`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *ArticleComplete) GetLicenseOk() (*License, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *ArticleComplete) SetLicense(v License)`

SetLicense sets License field to given value.


### GetMetadataReason

`func (o *ArticleComplete) GetMetadataReason() string`

GetMetadataReason returns the MetadataReason field if non-nil, zero value otherwise.

### GetMetadataReasonOk

`func (o *ArticleComplete) GetMetadataReasonOk() (*string, bool)`

GetMetadataReasonOk returns a tuple with the MetadataReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataReason

`func (o *ArticleComplete) SetMetadataReason(v string)`

SetMetadataReason sets MetadataReason field to given value.


### GetReferences

`func (o *ArticleComplete) GetReferences() []string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *ArticleComplete) GetReferencesOk() (*[]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *ArticleComplete) SetReferences(v []string)`

SetReferences sets References field to given value.


### GetRelatedMaterials

`func (o *ArticleComplete) GetRelatedMaterials() []RelatedMaterial`

GetRelatedMaterials returns the RelatedMaterials field if non-nil, zero value otherwise.

### GetRelatedMaterialsOk

`func (o *ArticleComplete) GetRelatedMaterialsOk() (*[]RelatedMaterial, bool)`

GetRelatedMaterialsOk returns a tuple with the RelatedMaterials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedMaterials

`func (o *ArticleComplete) SetRelatedMaterials(v []RelatedMaterial)`

SetRelatedMaterials sets RelatedMaterials field to given value.

### HasRelatedMaterials

`func (o *ArticleComplete) HasRelatedMaterials() bool`

HasRelatedMaterials returns a boolean if a field has been set.

### GetSize

`func (o *ArticleComplete) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ArticleComplete) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ArticleComplete) SetSize(v int64)`

SetSize sets Size field to given value.


### GetStatus

`func (o *ArticleComplete) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ArticleComplete) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ArticleComplete) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTags

`func (o *ArticleComplete) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ArticleComplete) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ArticleComplete) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetVersion

`func (o *ArticleComplete) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ArticleComplete) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ArticleComplete) SetVersion(v int64)`

SetVersion sets Version field to given value.


### GetDefinedType

`func (o *ArticleComplete) GetDefinedType() int64`

GetDefinedType returns the DefinedType field if non-nil, zero value otherwise.

### GetDefinedTypeOk

`func (o *ArticleComplete) GetDefinedTypeOk() (*int64, bool)`

GetDefinedTypeOk returns a tuple with the DefinedType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinedType

`func (o *ArticleComplete) SetDefinedType(v int64)`

SetDefinedType sets DefinedType field to given value.


### GetDefinedTypeName

`func (o *ArticleComplete) GetDefinedTypeName() string`

GetDefinedTypeName returns the DefinedTypeName field if non-nil, zero value otherwise.

### GetDefinedTypeNameOk

`func (o *ArticleComplete) GetDefinedTypeNameOk() (*string, bool)`

GetDefinedTypeNameOk returns a tuple with the DefinedTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinedTypeName

`func (o *ArticleComplete) SetDefinedTypeName(v string)`

SetDefinedTypeName sets DefinedTypeName field to given value.


### GetDoi

`func (o *ArticleComplete) GetDoi() string`

GetDoi returns the Doi field if non-nil, zero value otherwise.

### GetDoiOk

`func (o *ArticleComplete) GetDoiOk() (*string, bool)`

GetDoiOk returns a tuple with the Doi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoi

`func (o *ArticleComplete) SetDoi(v string)`

SetDoi sets Doi field to given value.


### GetHandle

`func (o *ArticleComplete) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *ArticleComplete) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *ArticleComplete) SetHandle(v string)`

SetHandle sets Handle field to given value.


### GetId

`func (o *ArticleComplete) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArticleComplete) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArticleComplete) SetId(v int64)`

SetId sets Id field to given value.


### GetResourceDoi

`func (o *ArticleComplete) GetResourceDoi() string`

GetResourceDoi returns the ResourceDoi field if non-nil, zero value otherwise.

### GetResourceDoiOk

`func (o *ArticleComplete) GetResourceDoiOk() (*string, bool)`

GetResourceDoiOk returns a tuple with the ResourceDoi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceDoi

`func (o *ArticleComplete) SetResourceDoi(v string)`

SetResourceDoi sets ResourceDoi field to given value.


### GetResourceTitle

`func (o *ArticleComplete) GetResourceTitle() string`

GetResourceTitle returns the ResourceTitle field if non-nil, zero value otherwise.

### GetResourceTitleOk

`func (o *ArticleComplete) GetResourceTitleOk() (*string, bool)`

GetResourceTitleOk returns a tuple with the ResourceTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceTitle

`func (o *ArticleComplete) SetResourceTitle(v string)`

SetResourceTitle sets ResourceTitle field to given value.


### GetThumb

`func (o *ArticleComplete) GetThumb() string`

GetThumb returns the Thumb field if non-nil, zero value otherwise.

### GetThumbOk

`func (o *ArticleComplete) GetThumbOk() (*string, bool)`

GetThumbOk returns a tuple with the Thumb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumb

`func (o *ArticleComplete) SetThumb(v string)`

SetThumb sets Thumb field to given value.


### GetTimeline

`func (o *ArticleComplete) GetTimeline() Timeline`

GetTimeline returns the Timeline field if non-nil, zero value otherwise.

### GetTimelineOk

`func (o *ArticleComplete) GetTimelineOk() (*Timeline, bool)`

GetTimelineOk returns a tuple with the Timeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeline

`func (o *ArticleComplete) SetTimeline(v Timeline)`

SetTimeline sets Timeline field to given value.


### GetTitle

`func (o *ArticleComplete) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ArticleComplete) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ArticleComplete) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetUrl

`func (o *ArticleComplete) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ArticleComplete) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ArticleComplete) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetUrlPrivateApi

`func (o *ArticleComplete) GetUrlPrivateApi() string`

GetUrlPrivateApi returns the UrlPrivateApi field if non-nil, zero value otherwise.

### GetUrlPrivateApiOk

`func (o *ArticleComplete) GetUrlPrivateApiOk() (*string, bool)`

GetUrlPrivateApiOk returns a tuple with the UrlPrivateApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlPrivateApi

`func (o *ArticleComplete) SetUrlPrivateApi(v string)`

SetUrlPrivateApi sets UrlPrivateApi field to given value.


### GetUrlPrivateHtml

`func (o *ArticleComplete) GetUrlPrivateHtml() string`

GetUrlPrivateHtml returns the UrlPrivateHtml field if non-nil, zero value otherwise.

### GetUrlPrivateHtmlOk

`func (o *ArticleComplete) GetUrlPrivateHtmlOk() (*string, bool)`

GetUrlPrivateHtmlOk returns a tuple with the UrlPrivateHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlPrivateHtml

`func (o *ArticleComplete) SetUrlPrivateHtml(v string)`

SetUrlPrivateHtml sets UrlPrivateHtml field to given value.


### GetUrlPublicApi

`func (o *ArticleComplete) GetUrlPublicApi() string`

GetUrlPublicApi returns the UrlPublicApi field if non-nil, zero value otherwise.

### GetUrlPublicApiOk

`func (o *ArticleComplete) GetUrlPublicApiOk() (*string, bool)`

GetUrlPublicApiOk returns a tuple with the UrlPublicApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlPublicApi

`func (o *ArticleComplete) SetUrlPublicApi(v string)`

SetUrlPublicApi sets UrlPublicApi field to given value.


### GetUrlPublicHtml

`func (o *ArticleComplete) GetUrlPublicHtml() string`

GetUrlPublicHtml returns the UrlPublicHtml field if non-nil, zero value otherwise.

### GetUrlPublicHtmlOk

`func (o *ArticleComplete) GetUrlPublicHtmlOk() (*string, bool)`

GetUrlPublicHtmlOk returns a tuple with the UrlPublicHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlPublicHtml

`func (o *ArticleComplete) SetUrlPublicHtml(v string)`

SetUrlPublicHtml sets UrlPublicHtml field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


