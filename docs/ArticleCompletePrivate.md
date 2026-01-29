# ArticleCompletePrivate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupResourceId** | **NullableString** | Group resource id | 
**AccountId** | **int64** | ID of the account owning the article | 
**CurationStatus** | **string** | Curation status of the article | 
**FigshareUrl** | **string** | Article public url | 
**DownloadDisabled** | **bool** | If true, downloading of files for this article is disabled | 
**Files** | [**[]PublicFile**](PublicFile.md) | List of article files. | 
**FolderStructure** | **map[string]interface{}** | Mapping of file ids to folder paths, if folders are used | 
**Authors** | [**[]Author**](Author.md) | List of article authors | 
**CustomFields** | [**[]CustomArticleField**](CustomArticleField.md) | List of custom fields values | 
**EmbargoOptions** | [**[]GroupEmbargoOptions**](GroupEmbargoOptions.md) | List of embargo options | 
**Citation** | **string** | Article citation | 
**ConfidentialReason** | **string** | Confidentiality reason | 
**EmbargoType** | **NullableString** | Article embargo | 
**IsConfidential** | **bool** | Article Confidentiality | 
**Size** | **int64** | Article size | 
**Funding** | **string** | Article funding | 
**FundingList** | [**[]FundingInformation**](FundingInformation.md) | Full Article funding information | 
**Tags** | **[]string** | List of article tags. Keywords can be used instead | 
**Keywords** | **[]string** | List of article keywords. Tags can be used instead | 
**Version** | **int64** | Article version | 
**IsMetadataRecord** | **bool** | True if article has no files | 
**MetadataReason** | **string** | Article metadata reason | 
**Status** | **string** | Article status | 
**Description** | **string** | Article description | 
**IsEmbargoed** | **bool** | True if article is embargoed | 
**EmbargoDate** | **NullableString** | Date when embargo lifts | 
**IsPublic** | **bool** | True if article is published | 
**ModifiedDate** | **NullableString** | Date when article was last modified | 
**CreatedDate** | **string** | Date when article was created | 
**HasLinkedFile** | **bool** | True if any files are linked to the article | 
**Categories** | [**[]Category**](Category.md) | List of categories selected for the article | 
**License** | [**License**](License.md) |  | 
**EmbargoTitle** | **string** | Title for embargo | 
**EmbargoReason** | **string** | Reason for embargo | 
**References** | **[]string** | List of references | 
**RelatedMaterials** | Pointer to [**[]RelatedMaterial**](RelatedMaterial.md) | List of related materials; supersedes references and resource DOI/title. | [optional] 
**Id** | **int64** | Unique identifier for article | 
**Title** | **string** | Title of article | 
**Doi** | **string** | DOI | 
**Handle** | **string** | Handle | 
**GroupId** | **NullableFloat32** | Group ID | 
**Url** | Pointer to **string** | Api endpoint for article | [optional] 
**UrlPublicHtml** | **string** | Public site endpoint for article | 
**UrlPublicApi** | **string** | Public Api endpoint for article | 
**UrlPrivateHtml** | **string** | Private site endpoint for article | 
**UrlPrivateApi** | **string** | Private Api endpoint for article | 
**PublishedDate** | **NullableString** | Posted date | 
**Timeline** | [**Timeline**](Timeline.md) |  | 
**Thumb** | **string** | Thumbnail image | 
**DefinedType** | **int64** | Type of article identifier | 
**DefinedTypeName** | **string** | Name of the article type identifier | 
**ResourceDoi** | **string** | Deprecated by related materials. Not applicable to regular users. In a publisher case, this is the publisher article DOI. | [default to ""]
**ResourceTitle** | **string** | Deprecated by related materials. Not applicable to regular users. In a publisher case, this is the publisher article title. | [default to ""]

## Methods

### NewArticleCompletePrivate

`func NewArticleCompletePrivate(groupResourceId NullableString, accountId int64, curationStatus string, figshareUrl string, downloadDisabled bool, files []PublicFile, folderStructure map[string]interface{}, authors []Author, customFields []CustomArticleField, embargoOptions []GroupEmbargoOptions, citation string, confidentialReason string, embargoType NullableString, isConfidential bool, size int64, funding string, fundingList []FundingInformation, tags []string, keywords []string, version int64, isMetadataRecord bool, metadataReason string, status string, description string, isEmbargoed bool, embargoDate NullableString, isPublic bool, modifiedDate NullableString, createdDate string, hasLinkedFile bool, categories []Category, license License, embargoTitle string, embargoReason string, references []string, id int64, title string, doi string, handle string, groupId NullableFloat32, urlPublicHtml string, urlPublicApi string, urlPrivateHtml string, urlPrivateApi string, publishedDate NullableString, timeline Timeline, thumb string, definedType int64, definedTypeName string, resourceDoi string, resourceTitle string, ) *ArticleCompletePrivate`

NewArticleCompletePrivate instantiates a new ArticleCompletePrivate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArticleCompletePrivateWithDefaults

`func NewArticleCompletePrivateWithDefaults() *ArticleCompletePrivate`

NewArticleCompletePrivateWithDefaults instantiates a new ArticleCompletePrivate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupResourceId

`func (o *ArticleCompletePrivate) GetGroupResourceId() string`

GetGroupResourceId returns the GroupResourceId field if non-nil, zero value otherwise.

### GetGroupResourceIdOk

`func (o *ArticleCompletePrivate) GetGroupResourceIdOk() (*string, bool)`

GetGroupResourceIdOk returns a tuple with the GroupResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupResourceId

`func (o *ArticleCompletePrivate) SetGroupResourceId(v string)`

SetGroupResourceId sets GroupResourceId field to given value.


### SetGroupResourceIdNil

`func (o *ArticleCompletePrivate) SetGroupResourceIdNil(b bool)`

 SetGroupResourceIdNil sets the value for GroupResourceId to be an explicit nil

### UnsetGroupResourceId
`func (o *ArticleCompletePrivate) UnsetGroupResourceId()`

UnsetGroupResourceId ensures that no value is present for GroupResourceId, not even an explicit nil
### GetAccountId

`func (o *ArticleCompletePrivate) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ArticleCompletePrivate) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ArticleCompletePrivate) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### GetCurationStatus

`func (o *ArticleCompletePrivate) GetCurationStatus() string`

GetCurationStatus returns the CurationStatus field if non-nil, zero value otherwise.

### GetCurationStatusOk

`func (o *ArticleCompletePrivate) GetCurationStatusOk() (*string, bool)`

GetCurationStatusOk returns a tuple with the CurationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurationStatus

`func (o *ArticleCompletePrivate) SetCurationStatus(v string)`

SetCurationStatus sets CurationStatus field to given value.


### GetFigshareUrl

`func (o *ArticleCompletePrivate) GetFigshareUrl() string`

GetFigshareUrl returns the FigshareUrl field if non-nil, zero value otherwise.

### GetFigshareUrlOk

`func (o *ArticleCompletePrivate) GetFigshareUrlOk() (*string, bool)`

GetFigshareUrlOk returns a tuple with the FigshareUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFigshareUrl

`func (o *ArticleCompletePrivate) SetFigshareUrl(v string)`

SetFigshareUrl sets FigshareUrl field to given value.


### GetDownloadDisabled

`func (o *ArticleCompletePrivate) GetDownloadDisabled() bool`

GetDownloadDisabled returns the DownloadDisabled field if non-nil, zero value otherwise.

### GetDownloadDisabledOk

`func (o *ArticleCompletePrivate) GetDownloadDisabledOk() (*bool, bool)`

GetDownloadDisabledOk returns a tuple with the DownloadDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadDisabled

`func (o *ArticleCompletePrivate) SetDownloadDisabled(v bool)`

SetDownloadDisabled sets DownloadDisabled field to given value.


### GetFiles

`func (o *ArticleCompletePrivate) GetFiles() []PublicFile`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *ArticleCompletePrivate) GetFilesOk() (*[]PublicFile, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *ArticleCompletePrivate) SetFiles(v []PublicFile)`

SetFiles sets Files field to given value.


### GetFolderStructure

`func (o *ArticleCompletePrivate) GetFolderStructure() map[string]interface{}`

GetFolderStructure returns the FolderStructure field if non-nil, zero value otherwise.

### GetFolderStructureOk

`func (o *ArticleCompletePrivate) GetFolderStructureOk() (*map[string]interface{}, bool)`

GetFolderStructureOk returns a tuple with the FolderStructure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderStructure

`func (o *ArticleCompletePrivate) SetFolderStructure(v map[string]interface{})`

SetFolderStructure sets FolderStructure field to given value.


### GetAuthors

`func (o *ArticleCompletePrivate) GetAuthors() []Author`

GetAuthors returns the Authors field if non-nil, zero value otherwise.

### GetAuthorsOk

`func (o *ArticleCompletePrivate) GetAuthorsOk() (*[]Author, bool)`

GetAuthorsOk returns a tuple with the Authors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthors

`func (o *ArticleCompletePrivate) SetAuthors(v []Author)`

SetAuthors sets Authors field to given value.


### GetCustomFields

`func (o *ArticleCompletePrivate) GetCustomFields() []CustomArticleField`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ArticleCompletePrivate) GetCustomFieldsOk() (*[]CustomArticleField, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ArticleCompletePrivate) SetCustomFields(v []CustomArticleField)`

SetCustomFields sets CustomFields field to given value.


### GetEmbargoOptions

`func (o *ArticleCompletePrivate) GetEmbargoOptions() []GroupEmbargoOptions`

GetEmbargoOptions returns the EmbargoOptions field if non-nil, zero value otherwise.

### GetEmbargoOptionsOk

`func (o *ArticleCompletePrivate) GetEmbargoOptionsOk() (*[]GroupEmbargoOptions, bool)`

GetEmbargoOptionsOk returns a tuple with the EmbargoOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbargoOptions

`func (o *ArticleCompletePrivate) SetEmbargoOptions(v []GroupEmbargoOptions)`

SetEmbargoOptions sets EmbargoOptions field to given value.


### GetCitation

`func (o *ArticleCompletePrivate) GetCitation() string`

GetCitation returns the Citation field if non-nil, zero value otherwise.

### GetCitationOk

`func (o *ArticleCompletePrivate) GetCitationOk() (*string, bool)`

GetCitationOk returns a tuple with the Citation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitation

`func (o *ArticleCompletePrivate) SetCitation(v string)`

SetCitation sets Citation field to given value.


### GetConfidentialReason

`func (o *ArticleCompletePrivate) GetConfidentialReason() string`

GetConfidentialReason returns the ConfidentialReason field if non-nil, zero value otherwise.

### GetConfidentialReasonOk

`func (o *ArticleCompletePrivate) GetConfidentialReasonOk() (*string, bool)`

GetConfidentialReasonOk returns a tuple with the ConfidentialReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidentialReason

`func (o *ArticleCompletePrivate) SetConfidentialReason(v string)`

SetConfidentialReason sets ConfidentialReason field to given value.


### GetEmbargoType

`func (o *ArticleCompletePrivate) GetEmbargoType() string`

GetEmbargoType returns the EmbargoType field if non-nil, zero value otherwise.

### GetEmbargoTypeOk

`func (o *ArticleCompletePrivate) GetEmbargoTypeOk() (*string, bool)`

GetEmbargoTypeOk returns a tuple with the EmbargoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbargoType

`func (o *ArticleCompletePrivate) SetEmbargoType(v string)`

SetEmbargoType sets EmbargoType field to given value.


### SetEmbargoTypeNil

`func (o *ArticleCompletePrivate) SetEmbargoTypeNil(b bool)`

 SetEmbargoTypeNil sets the value for EmbargoType to be an explicit nil

### UnsetEmbargoType
`func (o *ArticleCompletePrivate) UnsetEmbargoType()`

UnsetEmbargoType ensures that no value is present for EmbargoType, not even an explicit nil
### GetIsConfidential

`func (o *ArticleCompletePrivate) GetIsConfidential() bool`

GetIsConfidential returns the IsConfidential field if non-nil, zero value otherwise.

### GetIsConfidentialOk

`func (o *ArticleCompletePrivate) GetIsConfidentialOk() (*bool, bool)`

GetIsConfidentialOk returns a tuple with the IsConfidential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConfidential

`func (o *ArticleCompletePrivate) SetIsConfidential(v bool)`

SetIsConfidential sets IsConfidential field to given value.


### GetSize

`func (o *ArticleCompletePrivate) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ArticleCompletePrivate) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ArticleCompletePrivate) SetSize(v int64)`

SetSize sets Size field to given value.


### GetFunding

`func (o *ArticleCompletePrivate) GetFunding() string`

GetFunding returns the Funding field if non-nil, zero value otherwise.

### GetFundingOk

`func (o *ArticleCompletePrivate) GetFundingOk() (*string, bool)`

GetFundingOk returns a tuple with the Funding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunding

`func (o *ArticleCompletePrivate) SetFunding(v string)`

SetFunding sets Funding field to given value.


### GetFundingList

`func (o *ArticleCompletePrivate) GetFundingList() []FundingInformation`

GetFundingList returns the FundingList field if non-nil, zero value otherwise.

### GetFundingListOk

`func (o *ArticleCompletePrivate) GetFundingListOk() (*[]FundingInformation, bool)`

GetFundingListOk returns a tuple with the FundingList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingList

`func (o *ArticleCompletePrivate) SetFundingList(v []FundingInformation)`

SetFundingList sets FundingList field to given value.


### GetTags

`func (o *ArticleCompletePrivate) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ArticleCompletePrivate) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ArticleCompletePrivate) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetKeywords

`func (o *ArticleCompletePrivate) GetKeywords() []string`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *ArticleCompletePrivate) GetKeywordsOk() (*[]string, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *ArticleCompletePrivate) SetKeywords(v []string)`

SetKeywords sets Keywords field to given value.


### GetVersion

`func (o *ArticleCompletePrivate) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ArticleCompletePrivate) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ArticleCompletePrivate) SetVersion(v int64)`

SetVersion sets Version field to given value.


### GetIsMetadataRecord

`func (o *ArticleCompletePrivate) GetIsMetadataRecord() bool`

GetIsMetadataRecord returns the IsMetadataRecord field if non-nil, zero value otherwise.

### GetIsMetadataRecordOk

`func (o *ArticleCompletePrivate) GetIsMetadataRecordOk() (*bool, bool)`

GetIsMetadataRecordOk returns a tuple with the IsMetadataRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMetadataRecord

`func (o *ArticleCompletePrivate) SetIsMetadataRecord(v bool)`

SetIsMetadataRecord sets IsMetadataRecord field to given value.


### GetMetadataReason

`func (o *ArticleCompletePrivate) GetMetadataReason() string`

GetMetadataReason returns the MetadataReason field if non-nil, zero value otherwise.

### GetMetadataReasonOk

`func (o *ArticleCompletePrivate) GetMetadataReasonOk() (*string, bool)`

GetMetadataReasonOk returns a tuple with the MetadataReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataReason

`func (o *ArticleCompletePrivate) SetMetadataReason(v string)`

SetMetadataReason sets MetadataReason field to given value.


### GetStatus

`func (o *ArticleCompletePrivate) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ArticleCompletePrivate) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ArticleCompletePrivate) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetDescription

`func (o *ArticleCompletePrivate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ArticleCompletePrivate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ArticleCompletePrivate) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetIsEmbargoed

`func (o *ArticleCompletePrivate) GetIsEmbargoed() bool`

GetIsEmbargoed returns the IsEmbargoed field if non-nil, zero value otherwise.

### GetIsEmbargoedOk

`func (o *ArticleCompletePrivate) GetIsEmbargoedOk() (*bool, bool)`

GetIsEmbargoedOk returns a tuple with the IsEmbargoed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEmbargoed

`func (o *ArticleCompletePrivate) SetIsEmbargoed(v bool)`

SetIsEmbargoed sets IsEmbargoed field to given value.


### GetEmbargoDate

`func (o *ArticleCompletePrivate) GetEmbargoDate() string`

GetEmbargoDate returns the EmbargoDate field if non-nil, zero value otherwise.

### GetEmbargoDateOk

`func (o *ArticleCompletePrivate) GetEmbargoDateOk() (*string, bool)`

GetEmbargoDateOk returns a tuple with the EmbargoDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbargoDate

`func (o *ArticleCompletePrivate) SetEmbargoDate(v string)`

SetEmbargoDate sets EmbargoDate field to given value.


### SetEmbargoDateNil

`func (o *ArticleCompletePrivate) SetEmbargoDateNil(b bool)`

 SetEmbargoDateNil sets the value for EmbargoDate to be an explicit nil

### UnsetEmbargoDate
`func (o *ArticleCompletePrivate) UnsetEmbargoDate()`

UnsetEmbargoDate ensures that no value is present for EmbargoDate, not even an explicit nil
### GetIsPublic

`func (o *ArticleCompletePrivate) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *ArticleCompletePrivate) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *ArticleCompletePrivate) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.


### GetModifiedDate

`func (o *ArticleCompletePrivate) GetModifiedDate() string`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *ArticleCompletePrivate) GetModifiedDateOk() (*string, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *ArticleCompletePrivate) SetModifiedDate(v string)`

SetModifiedDate sets ModifiedDate field to given value.


### SetModifiedDateNil

`func (o *ArticleCompletePrivate) SetModifiedDateNil(b bool)`

 SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil

### UnsetModifiedDate
`func (o *ArticleCompletePrivate) UnsetModifiedDate()`

UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
### GetCreatedDate

`func (o *ArticleCompletePrivate) GetCreatedDate() string`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *ArticleCompletePrivate) GetCreatedDateOk() (*string, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *ArticleCompletePrivate) SetCreatedDate(v string)`

SetCreatedDate sets CreatedDate field to given value.


### GetHasLinkedFile

`func (o *ArticleCompletePrivate) GetHasLinkedFile() bool`

GetHasLinkedFile returns the HasLinkedFile field if non-nil, zero value otherwise.

### GetHasLinkedFileOk

`func (o *ArticleCompletePrivate) GetHasLinkedFileOk() (*bool, bool)`

GetHasLinkedFileOk returns a tuple with the HasLinkedFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasLinkedFile

`func (o *ArticleCompletePrivate) SetHasLinkedFile(v bool)`

SetHasLinkedFile sets HasLinkedFile field to given value.


### GetCategories

`func (o *ArticleCompletePrivate) GetCategories() []Category`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *ArticleCompletePrivate) GetCategoriesOk() (*[]Category, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *ArticleCompletePrivate) SetCategories(v []Category)`

SetCategories sets Categories field to given value.


### GetLicense

`func (o *ArticleCompletePrivate) GetLicense() License`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *ArticleCompletePrivate) GetLicenseOk() (*License, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *ArticleCompletePrivate) SetLicense(v License)`

SetLicense sets License field to given value.


### GetEmbargoTitle

`func (o *ArticleCompletePrivate) GetEmbargoTitle() string`

GetEmbargoTitle returns the EmbargoTitle field if non-nil, zero value otherwise.

### GetEmbargoTitleOk

`func (o *ArticleCompletePrivate) GetEmbargoTitleOk() (*string, bool)`

GetEmbargoTitleOk returns a tuple with the EmbargoTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbargoTitle

`func (o *ArticleCompletePrivate) SetEmbargoTitle(v string)`

SetEmbargoTitle sets EmbargoTitle field to given value.


### GetEmbargoReason

`func (o *ArticleCompletePrivate) GetEmbargoReason() string`

GetEmbargoReason returns the EmbargoReason field if non-nil, zero value otherwise.

### GetEmbargoReasonOk

`func (o *ArticleCompletePrivate) GetEmbargoReasonOk() (*string, bool)`

GetEmbargoReasonOk returns a tuple with the EmbargoReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbargoReason

`func (o *ArticleCompletePrivate) SetEmbargoReason(v string)`

SetEmbargoReason sets EmbargoReason field to given value.


### GetReferences

`func (o *ArticleCompletePrivate) GetReferences() []string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *ArticleCompletePrivate) GetReferencesOk() (*[]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *ArticleCompletePrivate) SetReferences(v []string)`

SetReferences sets References field to given value.


### GetRelatedMaterials

`func (o *ArticleCompletePrivate) GetRelatedMaterials() []RelatedMaterial`

GetRelatedMaterials returns the RelatedMaterials field if non-nil, zero value otherwise.

### GetRelatedMaterialsOk

`func (o *ArticleCompletePrivate) GetRelatedMaterialsOk() (*[]RelatedMaterial, bool)`

GetRelatedMaterialsOk returns a tuple with the RelatedMaterials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedMaterials

`func (o *ArticleCompletePrivate) SetRelatedMaterials(v []RelatedMaterial)`

SetRelatedMaterials sets RelatedMaterials field to given value.

### HasRelatedMaterials

`func (o *ArticleCompletePrivate) HasRelatedMaterials() bool`

HasRelatedMaterials returns a boolean if a field has been set.

### GetId

`func (o *ArticleCompletePrivate) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArticleCompletePrivate) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArticleCompletePrivate) SetId(v int64)`

SetId sets Id field to given value.


### GetTitle

`func (o *ArticleCompletePrivate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ArticleCompletePrivate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ArticleCompletePrivate) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDoi

`func (o *ArticleCompletePrivate) GetDoi() string`

GetDoi returns the Doi field if non-nil, zero value otherwise.

### GetDoiOk

`func (o *ArticleCompletePrivate) GetDoiOk() (*string, bool)`

GetDoiOk returns a tuple with the Doi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoi

`func (o *ArticleCompletePrivate) SetDoi(v string)`

SetDoi sets Doi field to given value.


### GetHandle

`func (o *ArticleCompletePrivate) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *ArticleCompletePrivate) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *ArticleCompletePrivate) SetHandle(v string)`

SetHandle sets Handle field to given value.


### GetGroupId

`func (o *ArticleCompletePrivate) GetGroupId() float32`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ArticleCompletePrivate) GetGroupIdOk() (*float32, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ArticleCompletePrivate) SetGroupId(v float32)`

SetGroupId sets GroupId field to given value.


### SetGroupIdNil

`func (o *ArticleCompletePrivate) SetGroupIdNil(b bool)`

 SetGroupIdNil sets the value for GroupId to be an explicit nil

### UnsetGroupId
`func (o *ArticleCompletePrivate) UnsetGroupId()`

UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
### GetUrl

`func (o *ArticleCompletePrivate) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ArticleCompletePrivate) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ArticleCompletePrivate) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ArticleCompletePrivate) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUrlPublicHtml

`func (o *ArticleCompletePrivate) GetUrlPublicHtml() string`

GetUrlPublicHtml returns the UrlPublicHtml field if non-nil, zero value otherwise.

### GetUrlPublicHtmlOk

`func (o *ArticleCompletePrivate) GetUrlPublicHtmlOk() (*string, bool)`

GetUrlPublicHtmlOk returns a tuple with the UrlPublicHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlPublicHtml

`func (o *ArticleCompletePrivate) SetUrlPublicHtml(v string)`

SetUrlPublicHtml sets UrlPublicHtml field to given value.


### GetUrlPublicApi

`func (o *ArticleCompletePrivate) GetUrlPublicApi() string`

GetUrlPublicApi returns the UrlPublicApi field if non-nil, zero value otherwise.

### GetUrlPublicApiOk

`func (o *ArticleCompletePrivate) GetUrlPublicApiOk() (*string, bool)`

GetUrlPublicApiOk returns a tuple with the UrlPublicApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlPublicApi

`func (o *ArticleCompletePrivate) SetUrlPublicApi(v string)`

SetUrlPublicApi sets UrlPublicApi field to given value.


### GetUrlPrivateHtml

`func (o *ArticleCompletePrivate) GetUrlPrivateHtml() string`

GetUrlPrivateHtml returns the UrlPrivateHtml field if non-nil, zero value otherwise.

### GetUrlPrivateHtmlOk

`func (o *ArticleCompletePrivate) GetUrlPrivateHtmlOk() (*string, bool)`

GetUrlPrivateHtmlOk returns a tuple with the UrlPrivateHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlPrivateHtml

`func (o *ArticleCompletePrivate) SetUrlPrivateHtml(v string)`

SetUrlPrivateHtml sets UrlPrivateHtml field to given value.


### GetUrlPrivateApi

`func (o *ArticleCompletePrivate) GetUrlPrivateApi() string`

GetUrlPrivateApi returns the UrlPrivateApi field if non-nil, zero value otherwise.

### GetUrlPrivateApiOk

`func (o *ArticleCompletePrivate) GetUrlPrivateApiOk() (*string, bool)`

GetUrlPrivateApiOk returns a tuple with the UrlPrivateApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlPrivateApi

`func (o *ArticleCompletePrivate) SetUrlPrivateApi(v string)`

SetUrlPrivateApi sets UrlPrivateApi field to given value.


### GetPublishedDate

`func (o *ArticleCompletePrivate) GetPublishedDate() string`

GetPublishedDate returns the PublishedDate field if non-nil, zero value otherwise.

### GetPublishedDateOk

`func (o *ArticleCompletePrivate) GetPublishedDateOk() (*string, bool)`

GetPublishedDateOk returns a tuple with the PublishedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedDate

`func (o *ArticleCompletePrivate) SetPublishedDate(v string)`

SetPublishedDate sets PublishedDate field to given value.


### SetPublishedDateNil

`func (o *ArticleCompletePrivate) SetPublishedDateNil(b bool)`

 SetPublishedDateNil sets the value for PublishedDate to be an explicit nil

### UnsetPublishedDate
`func (o *ArticleCompletePrivate) UnsetPublishedDate()`

UnsetPublishedDate ensures that no value is present for PublishedDate, not even an explicit nil
### GetTimeline

`func (o *ArticleCompletePrivate) GetTimeline() Timeline`

GetTimeline returns the Timeline field if non-nil, zero value otherwise.

### GetTimelineOk

`func (o *ArticleCompletePrivate) GetTimelineOk() (*Timeline, bool)`

GetTimelineOk returns a tuple with the Timeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeline

`func (o *ArticleCompletePrivate) SetTimeline(v Timeline)`

SetTimeline sets Timeline field to given value.


### GetThumb

`func (o *ArticleCompletePrivate) GetThumb() string`

GetThumb returns the Thumb field if non-nil, zero value otherwise.

### GetThumbOk

`func (o *ArticleCompletePrivate) GetThumbOk() (*string, bool)`

GetThumbOk returns a tuple with the Thumb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumb

`func (o *ArticleCompletePrivate) SetThumb(v string)`

SetThumb sets Thumb field to given value.


### GetDefinedType

`func (o *ArticleCompletePrivate) GetDefinedType() int64`

GetDefinedType returns the DefinedType field if non-nil, zero value otherwise.

### GetDefinedTypeOk

`func (o *ArticleCompletePrivate) GetDefinedTypeOk() (*int64, bool)`

GetDefinedTypeOk returns a tuple with the DefinedType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinedType

`func (o *ArticleCompletePrivate) SetDefinedType(v int64)`

SetDefinedType sets DefinedType field to given value.


### GetDefinedTypeName

`func (o *ArticleCompletePrivate) GetDefinedTypeName() string`

GetDefinedTypeName returns the DefinedTypeName field if non-nil, zero value otherwise.

### GetDefinedTypeNameOk

`func (o *ArticleCompletePrivate) GetDefinedTypeNameOk() (*string, bool)`

GetDefinedTypeNameOk returns a tuple with the DefinedTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinedTypeName

`func (o *ArticleCompletePrivate) SetDefinedTypeName(v string)`

SetDefinedTypeName sets DefinedTypeName field to given value.


### GetResourceDoi

`func (o *ArticleCompletePrivate) GetResourceDoi() string`

GetResourceDoi returns the ResourceDoi field if non-nil, zero value otherwise.

### GetResourceDoiOk

`func (o *ArticleCompletePrivate) GetResourceDoiOk() (*string, bool)`

GetResourceDoiOk returns a tuple with the ResourceDoi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceDoi

`func (o *ArticleCompletePrivate) SetResourceDoi(v string)`

SetResourceDoi sets ResourceDoi field to given value.


### GetResourceTitle

`func (o *ArticleCompletePrivate) GetResourceTitle() string`

GetResourceTitle returns the ResourceTitle field if non-nil, zero value otherwise.

### GetResourceTitleOk

`func (o *ArticleCompletePrivate) GetResourceTitleOk() (*string, bool)`

GetResourceTitleOk returns a tuple with the ResourceTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceTitle

`func (o *ArticleCompletePrivate) SetResourceTitle(v string)`

SetResourceTitle sets ResourceTitle field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


