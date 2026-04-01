# ArticleComplete

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FigshareUrl** | **string** | Article public url | 
**DownloadDisabled** | **bool** | If true, downloading of files for this article is disabled | 
**FolderStructure** | Pointer to **map[string]interface{}** | Mapping of file ids to folder paths, if folders are used | [optional] 
**Authors** | [**[]Author**](Author.md) | List of article authors | 
**Citation** | **string** | Article citation | 
**ConfidentialReason** | **string** | Confidentiality reason | 
**EmbargoType** | **string** | Article embargo | 
**IsConfidential** | **bool** | Article Confidentiality | 
**Size** | **int64** | Article size | 
**Funding** | **string** | Article funding | 
**Tags** | **[]string** | List of article tags. Keywords can be used instead | 
**Keywords** | **[]string** | List of article keywords. Tags can be used instead | 
**Version** | **int64** | Article version | 
**IsMetadataRecord** | **bool** | True if article has no files | 
**MetadataReason** | **string** | Article metadata reason | 
**Status** | **string** | Article status | 
**Description** | **string** | Article description | 
**IsEmbargoed** | **bool** | True if article is embargoed | 
**EmbargoDate** | **string** | Date when embargo lifts | 
**IsPublic** | **bool** | True if article is published | 
**ModifiedDate** | **NullableString** | Date when article was last modified | 
**CreatedDate** | **string** | Date when article was created | 
**HasLinkedFile** | **bool** | True if any files are linked to the article | 
**Categories** | [**[]Category**](Category.md) | List of categories selected for the article | 
**License** | [**License**](License.md) | Article selected license | 
**EmbargoTitle** | **string** | Title for embargo | 
**EmbargoReason** | **string** | Reason for embargo | 
**References** | **[]string** | List of references | 
**Id** | **int64** | Unique identifier for article | 
**Title** | **string** | Title of article | 
**Doi** | **string** | DOI | 
**Handle** | **string** | Handle | 
**GroupId** | **NullableFloat32** | Group ID | 
**Url** | **string** | Api endpoint for article | 
**UrlPublicHtml** | **string** | Public site endpoint for article | 
**UrlPublicApi** | **string** | Public Api endpoint for article | 
**UrlPrivateHtml** | **string** | Private site endpoint for article | 
**UrlPrivateApi** | **string** | Private Api endpoint for article | 
**PublishedDate** | **NullableString** | Posted date | 
**Thumb** | **string** | Thumbnail image | 
**DefinedType** | **int64** | Type of article identifier | 
**DefinedTypeName** | **string** | Name of the article type identifier | 
**ResourceDoi** | **string** | Deprecated by related materials. Not applicable to regular users. In a publisher case, this is the publisher article DOI. | [default to ""]
**ResourceTitle** | **string** | Deprecated by related materials. Not applicable to regular users. In a publisher case, this is the publisher article title. | [default to ""]

## Methods

### NewArticleComplete

`func NewArticleComplete(figshareUrl string, downloadDisabled bool, authors []Author, citation string, confidentialReason string, embargoType string, isConfidential bool, size int64, funding string, tags []string, keywords []string, version int64, isMetadataRecord bool, metadataReason string, status string, description string, isEmbargoed bool, embargoDate string, isPublic bool, modifiedDate NullableString, createdDate string, hasLinkedFile bool, categories []Category, license License, embargoTitle string, embargoReason string, references []string, id int64, title string, doi string, handle string, groupId NullableFloat32, url string, urlPublicHtml string, urlPublicApi string, urlPrivateHtml string, urlPrivateApi string, publishedDate NullableString, thumb string, definedType int64, definedTypeName string, resourceDoi string, resourceTitle string, ) *ArticleComplete`

NewArticleComplete instantiates a new ArticleComplete object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArticleCompleteWithDefaults

`func NewArticleCompleteWithDefaults() *ArticleComplete`

NewArticleCompleteWithDefaults instantiates a new ArticleComplete object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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

### HasFolderStructure

`func (o *ArticleComplete) HasFolderStructure() bool`

HasFolderStructure returns a boolean if a field has been set.

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


### GetEmbargoType

`func (o *ArticleComplete) GetEmbargoType() string`

GetEmbargoType returns the EmbargoType field if non-nil, zero value otherwise.

### GetEmbargoTypeOk

`func (o *ArticleComplete) GetEmbargoTypeOk() (*string, bool)`

GetEmbargoTypeOk returns a tuple with the EmbargoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbargoType

`func (o *ArticleComplete) SetEmbargoType(v string)`

SetEmbargoType sets EmbargoType field to given value.


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


### GetEmbargoDate

`func (o *ArticleComplete) GetEmbargoDate() string`

GetEmbargoDate returns the EmbargoDate field if non-nil, zero value otherwise.

### GetEmbargoDateOk

`func (o *ArticleComplete) GetEmbargoDateOk() (*string, bool)`

GetEmbargoDateOk returns a tuple with the EmbargoDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbargoDate

`func (o *ArticleComplete) SetEmbargoDate(v string)`

SetEmbargoDate sets EmbargoDate field to given value.


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


### GetModifiedDate

`func (o *ArticleComplete) GetModifiedDate() string`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *ArticleComplete) GetModifiedDateOk() (*string, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *ArticleComplete) SetModifiedDate(v string)`

SetModifiedDate sets ModifiedDate field to given value.


### SetModifiedDateNil

`func (o *ArticleComplete) SetModifiedDateNil(b bool)`

 SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil

### UnsetModifiedDate
`func (o *ArticleComplete) UnsetModifiedDate()`

UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
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


### GetGroupId

`func (o *ArticleComplete) GetGroupId() float32`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ArticleComplete) GetGroupIdOk() (*float32, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ArticleComplete) SetGroupId(v float32)`

SetGroupId sets GroupId field to given value.


### SetGroupIdNil

`func (o *ArticleComplete) SetGroupIdNil(b bool)`

 SetGroupIdNil sets the value for GroupId to be an explicit nil

### UnsetGroupId
`func (o *ArticleComplete) UnsetGroupId()`

UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
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


### GetPublishedDate

`func (o *ArticleComplete) GetPublishedDate() string`

GetPublishedDate returns the PublishedDate field if non-nil, zero value otherwise.

### GetPublishedDateOk

`func (o *ArticleComplete) GetPublishedDateOk() (*string, bool)`

GetPublishedDateOk returns a tuple with the PublishedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedDate

`func (o *ArticleComplete) SetPublishedDate(v string)`

SetPublishedDate sets PublishedDate field to given value.


### SetPublishedDateNil

`func (o *ArticleComplete) SetPublishedDateNil(b bool)`

 SetPublishedDateNil sets the value for PublishedDate to be an explicit nil

### UnsetPublishedDate
`func (o *ArticleComplete) UnsetPublishedDate()`

UnsetPublishedDate ensures that no value is present for PublishedDate, not even an explicit nil
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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


