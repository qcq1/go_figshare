# PrivateProjectArticle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **int64** | ID of the account owning the article | 
**Authors** | [**[]Author**](Author.md) | List of authors | 
**CurationStatus** | **string** | Curation status of the article | 
**CustomFields** | [**[]CustomArticleField**](CustomArticleField.md) | List of custom fields values | 
**DownloadDisabled** | **bool** | If true, downloading of files for this article is disabled | 
**EmbargoOptions** | [**[]GroupEmbargoOptions**](GroupEmbargoOptions.md) | List of embargo options | 
**FigshareUrl** | **string** | Article public url | 
**Files** | [**[]PublicFile**](PublicFile.md) | List of up to 10 article files. | 
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

### NewPrivateProjectArticle

`func NewPrivateProjectArticle(accountId int64, authors []Author, curationStatus string, customFields []CustomArticleField, downloadDisabled bool, embargoOptions []GroupEmbargoOptions, figshareUrl string, files []PublicFile, categories []Category, citation string, confidentialReason string, createdDate string, description string, embargoReason string, embargoTitle string, funding string, fundingList []FundingInformation, hasLinkedFile bool, isConfidential bool, isEmbargoed bool, isMetadataRecord bool, isPublic bool, keywords []string, license License, metadataReason string, references []string, size int64, status string, tags []string, version int64, definedType int64, definedTypeName string, doi string, handle string, id int64, resourceDoi string, resourceTitle string, thumb string, timeline Timeline, title string, url string, urlPrivateApi string, urlPrivateHtml string, urlPublicApi string, urlPublicHtml string, ) *PrivateProjectArticle`

NewPrivateProjectArticle instantiates a new PrivateProjectArticle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateProjectArticleWithDefaults

`func NewPrivateProjectArticleWithDefaults() *PrivateProjectArticle`

NewPrivateProjectArticleWithDefaults instantiates a new PrivateProjectArticle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *PrivateProjectArticle) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *PrivateProjectArticle) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *PrivateProjectArticle) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### GetAuthors

`func (o *PrivateProjectArticle) GetAuthors() []Author`

GetAuthors returns the Authors field if non-nil, zero value otherwise.

### GetAuthorsOk

`func (o *PrivateProjectArticle) GetAuthorsOk() (*[]Author, bool)`

GetAuthorsOk returns a tuple with the Authors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthors

`func (o *PrivateProjectArticle) SetAuthors(v []Author)`

SetAuthors sets Authors field to given value.


### GetCurationStatus

`func (o *PrivateProjectArticle) GetCurationStatus() string`

GetCurationStatus returns the CurationStatus field if non-nil, zero value otherwise.

### GetCurationStatusOk

`func (o *PrivateProjectArticle) GetCurationStatusOk() (*string, bool)`

GetCurationStatusOk returns a tuple with the CurationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurationStatus

`func (o *PrivateProjectArticle) SetCurationStatus(v string)`

SetCurationStatus sets CurationStatus field to given value.


### GetCustomFields

`func (o *PrivateProjectArticle) GetCustomFields() []CustomArticleField`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PrivateProjectArticle) GetCustomFieldsOk() (*[]CustomArticleField, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PrivateProjectArticle) SetCustomFields(v []CustomArticleField)`

SetCustomFields sets CustomFields field to given value.


### GetDownloadDisabled

`func (o *PrivateProjectArticle) GetDownloadDisabled() bool`

GetDownloadDisabled returns the DownloadDisabled field if non-nil, zero value otherwise.

### GetDownloadDisabledOk

`func (o *PrivateProjectArticle) GetDownloadDisabledOk() (*bool, bool)`

GetDownloadDisabledOk returns a tuple with the DownloadDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadDisabled

`func (o *PrivateProjectArticle) SetDownloadDisabled(v bool)`

SetDownloadDisabled sets DownloadDisabled field to given value.


### GetEmbargoOptions

`func (o *PrivateProjectArticle) GetEmbargoOptions() []GroupEmbargoOptions`

GetEmbargoOptions returns the EmbargoOptions field if non-nil, zero value otherwise.

### GetEmbargoOptionsOk

`func (o *PrivateProjectArticle) GetEmbargoOptionsOk() (*[]GroupEmbargoOptions, bool)`

GetEmbargoOptionsOk returns a tuple with the EmbargoOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbargoOptions

`func (o *PrivateProjectArticle) SetEmbargoOptions(v []GroupEmbargoOptions)`

SetEmbargoOptions sets EmbargoOptions field to given value.


### GetFigshareUrl

`func (o *PrivateProjectArticle) GetFigshareUrl() string`

GetFigshareUrl returns the FigshareUrl field if non-nil, zero value otherwise.

### GetFigshareUrlOk

`func (o *PrivateProjectArticle) GetFigshareUrlOk() (*string, bool)`

GetFigshareUrlOk returns a tuple with the FigshareUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFigshareUrl

`func (o *PrivateProjectArticle) SetFigshareUrl(v string)`

SetFigshareUrl sets FigshareUrl field to given value.


### GetFiles

`func (o *PrivateProjectArticle) GetFiles() []PublicFile`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *PrivateProjectArticle) GetFilesOk() (*[]PublicFile, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *PrivateProjectArticle) SetFiles(v []PublicFile)`

SetFiles sets Files field to given value.


### GetCategories

`func (o *PrivateProjectArticle) GetCategories() []Category`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *PrivateProjectArticle) GetCategoriesOk() (*[]Category, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *PrivateProjectArticle) SetCategories(v []Category)`

SetCategories sets Categories field to given value.


### GetCitation

`func (o *PrivateProjectArticle) GetCitation() string`

GetCitation returns the Citation field if non-nil, zero value otherwise.

### GetCitationOk

`func (o *PrivateProjectArticle) GetCitationOk() (*string, bool)`

GetCitationOk returns a tuple with the Citation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitation

`func (o *PrivateProjectArticle) SetCitation(v string)`

SetCitation sets Citation field to given value.


### GetConfidentialReason

`func (o *PrivateProjectArticle) GetConfidentialReason() string`

GetConfidentialReason returns the ConfidentialReason field if non-nil, zero value otherwise.

### GetConfidentialReasonOk

`func (o *PrivateProjectArticle) GetConfidentialReasonOk() (*string, bool)`

GetConfidentialReasonOk returns a tuple with the ConfidentialReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidentialReason

`func (o *PrivateProjectArticle) SetConfidentialReason(v string)`

SetConfidentialReason sets ConfidentialReason field to given value.


### GetCreatedDate

`func (o *PrivateProjectArticle) GetCreatedDate() string`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *PrivateProjectArticle) GetCreatedDateOk() (*string, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *PrivateProjectArticle) SetCreatedDate(v string)`

SetCreatedDate sets CreatedDate field to given value.


### GetDescription

`func (o *PrivateProjectArticle) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PrivateProjectArticle) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PrivateProjectArticle) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEmbargoReason

`func (o *PrivateProjectArticle) GetEmbargoReason() string`

GetEmbargoReason returns the EmbargoReason field if non-nil, zero value otherwise.

### GetEmbargoReasonOk

`func (o *PrivateProjectArticle) GetEmbargoReasonOk() (*string, bool)`

GetEmbargoReasonOk returns a tuple with the EmbargoReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbargoReason

`func (o *PrivateProjectArticle) SetEmbargoReason(v string)`

SetEmbargoReason sets EmbargoReason field to given value.


### GetEmbargoTitle

`func (o *PrivateProjectArticle) GetEmbargoTitle() string`

GetEmbargoTitle returns the EmbargoTitle field if non-nil, zero value otherwise.

### GetEmbargoTitleOk

`func (o *PrivateProjectArticle) GetEmbargoTitleOk() (*string, bool)`

GetEmbargoTitleOk returns a tuple with the EmbargoTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbargoTitle

`func (o *PrivateProjectArticle) SetEmbargoTitle(v string)`

SetEmbargoTitle sets EmbargoTitle field to given value.


### GetFunding

`func (o *PrivateProjectArticle) GetFunding() string`

GetFunding returns the Funding field if non-nil, zero value otherwise.

### GetFundingOk

`func (o *PrivateProjectArticle) GetFundingOk() (*string, bool)`

GetFundingOk returns a tuple with the Funding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunding

`func (o *PrivateProjectArticle) SetFunding(v string)`

SetFunding sets Funding field to given value.


### GetFundingList

`func (o *PrivateProjectArticle) GetFundingList() []FundingInformation`

GetFundingList returns the FundingList field if non-nil, zero value otherwise.

### GetFundingListOk

`func (o *PrivateProjectArticle) GetFundingListOk() (*[]FundingInformation, bool)`

GetFundingListOk returns a tuple with the FundingList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingList

`func (o *PrivateProjectArticle) SetFundingList(v []FundingInformation)`

SetFundingList sets FundingList field to given value.


### GetHasLinkedFile

`func (o *PrivateProjectArticle) GetHasLinkedFile() bool`

GetHasLinkedFile returns the HasLinkedFile field if non-nil, zero value otherwise.

### GetHasLinkedFileOk

`func (o *PrivateProjectArticle) GetHasLinkedFileOk() (*bool, bool)`

GetHasLinkedFileOk returns a tuple with the HasLinkedFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasLinkedFile

`func (o *PrivateProjectArticle) SetHasLinkedFile(v bool)`

SetHasLinkedFile sets HasLinkedFile field to given value.


### GetIsConfidential

`func (o *PrivateProjectArticle) GetIsConfidential() bool`

GetIsConfidential returns the IsConfidential field if non-nil, zero value otherwise.

### GetIsConfidentialOk

`func (o *PrivateProjectArticle) GetIsConfidentialOk() (*bool, bool)`

GetIsConfidentialOk returns a tuple with the IsConfidential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConfidential

`func (o *PrivateProjectArticle) SetIsConfidential(v bool)`

SetIsConfidential sets IsConfidential field to given value.


### GetIsEmbargoed

`func (o *PrivateProjectArticle) GetIsEmbargoed() bool`

GetIsEmbargoed returns the IsEmbargoed field if non-nil, zero value otherwise.

### GetIsEmbargoedOk

`func (o *PrivateProjectArticle) GetIsEmbargoedOk() (*bool, bool)`

GetIsEmbargoedOk returns a tuple with the IsEmbargoed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEmbargoed

`func (o *PrivateProjectArticle) SetIsEmbargoed(v bool)`

SetIsEmbargoed sets IsEmbargoed field to given value.


### GetIsMetadataRecord

`func (o *PrivateProjectArticle) GetIsMetadataRecord() bool`

GetIsMetadataRecord returns the IsMetadataRecord field if non-nil, zero value otherwise.

### GetIsMetadataRecordOk

`func (o *PrivateProjectArticle) GetIsMetadataRecordOk() (*bool, bool)`

GetIsMetadataRecordOk returns a tuple with the IsMetadataRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMetadataRecord

`func (o *PrivateProjectArticle) SetIsMetadataRecord(v bool)`

SetIsMetadataRecord sets IsMetadataRecord field to given value.


### GetIsPublic

`func (o *PrivateProjectArticle) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *PrivateProjectArticle) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *PrivateProjectArticle) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.


### GetKeywords

`func (o *PrivateProjectArticle) GetKeywords() []string`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *PrivateProjectArticle) GetKeywordsOk() (*[]string, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *PrivateProjectArticle) SetKeywords(v []string)`

SetKeywords sets Keywords field to given value.


### GetLicense

`func (o *PrivateProjectArticle) GetLicense() License`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *PrivateProjectArticle) GetLicenseOk() (*License, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *PrivateProjectArticle) SetLicense(v License)`

SetLicense sets License field to given value.


### GetMetadataReason

`func (o *PrivateProjectArticle) GetMetadataReason() string`

GetMetadataReason returns the MetadataReason field if non-nil, zero value otherwise.

### GetMetadataReasonOk

`func (o *PrivateProjectArticle) GetMetadataReasonOk() (*string, bool)`

GetMetadataReasonOk returns a tuple with the MetadataReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataReason

`func (o *PrivateProjectArticle) SetMetadataReason(v string)`

SetMetadataReason sets MetadataReason field to given value.


### GetReferences

`func (o *PrivateProjectArticle) GetReferences() []string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *PrivateProjectArticle) GetReferencesOk() (*[]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *PrivateProjectArticle) SetReferences(v []string)`

SetReferences sets References field to given value.


### GetRelatedMaterials

`func (o *PrivateProjectArticle) GetRelatedMaterials() []RelatedMaterial`

GetRelatedMaterials returns the RelatedMaterials field if non-nil, zero value otherwise.

### GetRelatedMaterialsOk

`func (o *PrivateProjectArticle) GetRelatedMaterialsOk() (*[]RelatedMaterial, bool)`

GetRelatedMaterialsOk returns a tuple with the RelatedMaterials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedMaterials

`func (o *PrivateProjectArticle) SetRelatedMaterials(v []RelatedMaterial)`

SetRelatedMaterials sets RelatedMaterials field to given value.

### HasRelatedMaterials

`func (o *PrivateProjectArticle) HasRelatedMaterials() bool`

HasRelatedMaterials returns a boolean if a field has been set.

### GetSize

`func (o *PrivateProjectArticle) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PrivateProjectArticle) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PrivateProjectArticle) SetSize(v int64)`

SetSize sets Size field to given value.


### GetStatus

`func (o *PrivateProjectArticle) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PrivateProjectArticle) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PrivateProjectArticle) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTags

`func (o *PrivateProjectArticle) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PrivateProjectArticle) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PrivateProjectArticle) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetVersion

`func (o *PrivateProjectArticle) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PrivateProjectArticle) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PrivateProjectArticle) SetVersion(v int64)`

SetVersion sets Version field to given value.


### GetDefinedType

`func (o *PrivateProjectArticle) GetDefinedType() int64`

GetDefinedType returns the DefinedType field if non-nil, zero value otherwise.

### GetDefinedTypeOk

`func (o *PrivateProjectArticle) GetDefinedTypeOk() (*int64, bool)`

GetDefinedTypeOk returns a tuple with the DefinedType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinedType

`func (o *PrivateProjectArticle) SetDefinedType(v int64)`

SetDefinedType sets DefinedType field to given value.


### GetDefinedTypeName

`func (o *PrivateProjectArticle) GetDefinedTypeName() string`

GetDefinedTypeName returns the DefinedTypeName field if non-nil, zero value otherwise.

### GetDefinedTypeNameOk

`func (o *PrivateProjectArticle) GetDefinedTypeNameOk() (*string, bool)`

GetDefinedTypeNameOk returns a tuple with the DefinedTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinedTypeName

`func (o *PrivateProjectArticle) SetDefinedTypeName(v string)`

SetDefinedTypeName sets DefinedTypeName field to given value.


### GetDoi

`func (o *PrivateProjectArticle) GetDoi() string`

GetDoi returns the Doi field if non-nil, zero value otherwise.

### GetDoiOk

`func (o *PrivateProjectArticle) GetDoiOk() (*string, bool)`

GetDoiOk returns a tuple with the Doi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoi

`func (o *PrivateProjectArticle) SetDoi(v string)`

SetDoi sets Doi field to given value.


### GetHandle

`func (o *PrivateProjectArticle) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *PrivateProjectArticle) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *PrivateProjectArticle) SetHandle(v string)`

SetHandle sets Handle field to given value.


### GetId

`func (o *PrivateProjectArticle) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrivateProjectArticle) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrivateProjectArticle) SetId(v int64)`

SetId sets Id field to given value.


### GetResourceDoi

`func (o *PrivateProjectArticle) GetResourceDoi() string`

GetResourceDoi returns the ResourceDoi field if non-nil, zero value otherwise.

### GetResourceDoiOk

`func (o *PrivateProjectArticle) GetResourceDoiOk() (*string, bool)`

GetResourceDoiOk returns a tuple with the ResourceDoi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceDoi

`func (o *PrivateProjectArticle) SetResourceDoi(v string)`

SetResourceDoi sets ResourceDoi field to given value.


### GetResourceTitle

`func (o *PrivateProjectArticle) GetResourceTitle() string`

GetResourceTitle returns the ResourceTitle field if non-nil, zero value otherwise.

### GetResourceTitleOk

`func (o *PrivateProjectArticle) GetResourceTitleOk() (*string, bool)`

GetResourceTitleOk returns a tuple with the ResourceTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceTitle

`func (o *PrivateProjectArticle) SetResourceTitle(v string)`

SetResourceTitle sets ResourceTitle field to given value.


### GetThumb

`func (o *PrivateProjectArticle) GetThumb() string`

GetThumb returns the Thumb field if non-nil, zero value otherwise.

### GetThumbOk

`func (o *PrivateProjectArticle) GetThumbOk() (*string, bool)`

GetThumbOk returns a tuple with the Thumb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumb

`func (o *PrivateProjectArticle) SetThumb(v string)`

SetThumb sets Thumb field to given value.


### GetTimeline

`func (o *PrivateProjectArticle) GetTimeline() Timeline`

GetTimeline returns the Timeline field if non-nil, zero value otherwise.

### GetTimelineOk

`func (o *PrivateProjectArticle) GetTimelineOk() (*Timeline, bool)`

GetTimelineOk returns a tuple with the Timeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeline

`func (o *PrivateProjectArticle) SetTimeline(v Timeline)`

SetTimeline sets Timeline field to given value.


### GetTitle

`func (o *PrivateProjectArticle) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PrivateProjectArticle) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PrivateProjectArticle) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetUrl

`func (o *PrivateProjectArticle) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PrivateProjectArticle) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PrivateProjectArticle) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetUrlPrivateApi

`func (o *PrivateProjectArticle) GetUrlPrivateApi() string`

GetUrlPrivateApi returns the UrlPrivateApi field if non-nil, zero value otherwise.

### GetUrlPrivateApiOk

`func (o *PrivateProjectArticle) GetUrlPrivateApiOk() (*string, bool)`

GetUrlPrivateApiOk returns a tuple with the UrlPrivateApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlPrivateApi

`func (o *PrivateProjectArticle) SetUrlPrivateApi(v string)`

SetUrlPrivateApi sets UrlPrivateApi field to given value.


### GetUrlPrivateHtml

`func (o *PrivateProjectArticle) GetUrlPrivateHtml() string`

GetUrlPrivateHtml returns the UrlPrivateHtml field if non-nil, zero value otherwise.

### GetUrlPrivateHtmlOk

`func (o *PrivateProjectArticle) GetUrlPrivateHtmlOk() (*string, bool)`

GetUrlPrivateHtmlOk returns a tuple with the UrlPrivateHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlPrivateHtml

`func (o *PrivateProjectArticle) SetUrlPrivateHtml(v string)`

SetUrlPrivateHtml sets UrlPrivateHtml field to given value.


### GetUrlPublicApi

`func (o *PrivateProjectArticle) GetUrlPublicApi() string`

GetUrlPublicApi returns the UrlPublicApi field if non-nil, zero value otherwise.

### GetUrlPublicApiOk

`func (o *PrivateProjectArticle) GetUrlPublicApiOk() (*string, bool)`

GetUrlPublicApiOk returns a tuple with the UrlPublicApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlPublicApi

`func (o *PrivateProjectArticle) SetUrlPublicApi(v string)`

SetUrlPublicApi sets UrlPublicApi field to given value.


### GetUrlPublicHtml

`func (o *PrivateProjectArticle) GetUrlPublicHtml() string`

GetUrlPublicHtml returns the UrlPublicHtml field if non-nil, zero value otherwise.

### GetUrlPublicHtmlOk

`func (o *PrivateProjectArticle) GetUrlPublicHtmlOk() (*string, bool)`

GetUrlPublicHtmlOk returns a tuple with the UrlPublicHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlPublicHtml

`func (o *PrivateProjectArticle) SetUrlPublicHtml(v string)`

SetUrlPublicHtml sets UrlPublicHtml field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


