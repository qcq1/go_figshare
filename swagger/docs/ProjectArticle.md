# ProjectArticle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewProjectArticle

`func NewProjectArticle(categories []Category, citation string, confidentialReason string, createdDate string, description string, embargoReason string, embargoTitle string, funding string, fundingList []FundingInformation, hasLinkedFile bool, isConfidential bool, isEmbargoed bool, isMetadataRecord bool, isPublic bool, keywords []string, license License, metadataReason string, references []string, size int64, status string, tags []string, version int64, definedType int64, definedTypeName string, doi string, handle string, id int64, resourceDoi string, resourceTitle string, thumb string, timeline Timeline, title string, url string, urlPrivateApi string, urlPrivateHtml string, urlPublicApi string, urlPublicHtml string, ) *ProjectArticle`

NewProjectArticle instantiates a new ProjectArticle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectArticleWithDefaults

`func NewProjectArticleWithDefaults() *ProjectArticle`

NewProjectArticleWithDefaults instantiates a new ProjectArticle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategories

`func (o *ProjectArticle) GetCategories() []Category`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *ProjectArticle) GetCategoriesOk() (*[]Category, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *ProjectArticle) SetCategories(v []Category)`

SetCategories sets Categories field to given value.


### GetCitation

`func (o *ProjectArticle) GetCitation() string`

GetCitation returns the Citation field if non-nil, zero value otherwise.

### GetCitationOk

`func (o *ProjectArticle) GetCitationOk() (*string, bool)`

GetCitationOk returns a tuple with the Citation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitation

`func (o *ProjectArticle) SetCitation(v string)`

SetCitation sets Citation field to given value.


### GetConfidentialReason

`func (o *ProjectArticle) GetConfidentialReason() string`

GetConfidentialReason returns the ConfidentialReason field if non-nil, zero value otherwise.

### GetConfidentialReasonOk

`func (o *ProjectArticle) GetConfidentialReasonOk() (*string, bool)`

GetConfidentialReasonOk returns a tuple with the ConfidentialReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidentialReason

`func (o *ProjectArticle) SetConfidentialReason(v string)`

SetConfidentialReason sets ConfidentialReason field to given value.


### GetCreatedDate

`func (o *ProjectArticle) GetCreatedDate() string`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *ProjectArticle) GetCreatedDateOk() (*string, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *ProjectArticle) SetCreatedDate(v string)`

SetCreatedDate sets CreatedDate field to given value.


### GetDescription

`func (o *ProjectArticle) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProjectArticle) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProjectArticle) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEmbargoReason

`func (o *ProjectArticle) GetEmbargoReason() string`

GetEmbargoReason returns the EmbargoReason field if non-nil, zero value otherwise.

### GetEmbargoReasonOk

`func (o *ProjectArticle) GetEmbargoReasonOk() (*string, bool)`

GetEmbargoReasonOk returns a tuple with the EmbargoReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbargoReason

`func (o *ProjectArticle) SetEmbargoReason(v string)`

SetEmbargoReason sets EmbargoReason field to given value.


### GetEmbargoTitle

`func (o *ProjectArticle) GetEmbargoTitle() string`

GetEmbargoTitle returns the EmbargoTitle field if non-nil, zero value otherwise.

### GetEmbargoTitleOk

`func (o *ProjectArticle) GetEmbargoTitleOk() (*string, bool)`

GetEmbargoTitleOk returns a tuple with the EmbargoTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbargoTitle

`func (o *ProjectArticle) SetEmbargoTitle(v string)`

SetEmbargoTitle sets EmbargoTitle field to given value.


### GetFunding

`func (o *ProjectArticle) GetFunding() string`

GetFunding returns the Funding field if non-nil, zero value otherwise.

### GetFundingOk

`func (o *ProjectArticle) GetFundingOk() (*string, bool)`

GetFundingOk returns a tuple with the Funding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunding

`func (o *ProjectArticle) SetFunding(v string)`

SetFunding sets Funding field to given value.


### GetFundingList

`func (o *ProjectArticle) GetFundingList() []FundingInformation`

GetFundingList returns the FundingList field if non-nil, zero value otherwise.

### GetFundingListOk

`func (o *ProjectArticle) GetFundingListOk() (*[]FundingInformation, bool)`

GetFundingListOk returns a tuple with the FundingList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingList

`func (o *ProjectArticle) SetFundingList(v []FundingInformation)`

SetFundingList sets FundingList field to given value.


### GetHasLinkedFile

`func (o *ProjectArticle) GetHasLinkedFile() bool`

GetHasLinkedFile returns the HasLinkedFile field if non-nil, zero value otherwise.

### GetHasLinkedFileOk

`func (o *ProjectArticle) GetHasLinkedFileOk() (*bool, bool)`

GetHasLinkedFileOk returns a tuple with the HasLinkedFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasLinkedFile

`func (o *ProjectArticle) SetHasLinkedFile(v bool)`

SetHasLinkedFile sets HasLinkedFile field to given value.


### GetIsConfidential

`func (o *ProjectArticle) GetIsConfidential() bool`

GetIsConfidential returns the IsConfidential field if non-nil, zero value otherwise.

### GetIsConfidentialOk

`func (o *ProjectArticle) GetIsConfidentialOk() (*bool, bool)`

GetIsConfidentialOk returns a tuple with the IsConfidential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConfidential

`func (o *ProjectArticle) SetIsConfidential(v bool)`

SetIsConfidential sets IsConfidential field to given value.


### GetIsEmbargoed

`func (o *ProjectArticle) GetIsEmbargoed() bool`

GetIsEmbargoed returns the IsEmbargoed field if non-nil, zero value otherwise.

### GetIsEmbargoedOk

`func (o *ProjectArticle) GetIsEmbargoedOk() (*bool, bool)`

GetIsEmbargoedOk returns a tuple with the IsEmbargoed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEmbargoed

`func (o *ProjectArticle) SetIsEmbargoed(v bool)`

SetIsEmbargoed sets IsEmbargoed field to given value.


### GetIsMetadataRecord

`func (o *ProjectArticle) GetIsMetadataRecord() bool`

GetIsMetadataRecord returns the IsMetadataRecord field if non-nil, zero value otherwise.

### GetIsMetadataRecordOk

`func (o *ProjectArticle) GetIsMetadataRecordOk() (*bool, bool)`

GetIsMetadataRecordOk returns a tuple with the IsMetadataRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMetadataRecord

`func (o *ProjectArticle) SetIsMetadataRecord(v bool)`

SetIsMetadataRecord sets IsMetadataRecord field to given value.


### GetIsPublic

`func (o *ProjectArticle) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *ProjectArticle) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *ProjectArticle) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.


### GetKeywords

`func (o *ProjectArticle) GetKeywords() []string`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *ProjectArticle) GetKeywordsOk() (*[]string, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *ProjectArticle) SetKeywords(v []string)`

SetKeywords sets Keywords field to given value.


### GetLicense

`func (o *ProjectArticle) GetLicense() License`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *ProjectArticle) GetLicenseOk() (*License, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *ProjectArticle) SetLicense(v License)`

SetLicense sets License field to given value.


### GetMetadataReason

`func (o *ProjectArticle) GetMetadataReason() string`

GetMetadataReason returns the MetadataReason field if non-nil, zero value otherwise.

### GetMetadataReasonOk

`func (o *ProjectArticle) GetMetadataReasonOk() (*string, bool)`

GetMetadataReasonOk returns a tuple with the MetadataReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataReason

`func (o *ProjectArticle) SetMetadataReason(v string)`

SetMetadataReason sets MetadataReason field to given value.


### GetReferences

`func (o *ProjectArticle) GetReferences() []string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *ProjectArticle) GetReferencesOk() (*[]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *ProjectArticle) SetReferences(v []string)`

SetReferences sets References field to given value.


### GetRelatedMaterials

`func (o *ProjectArticle) GetRelatedMaterials() []RelatedMaterial`

GetRelatedMaterials returns the RelatedMaterials field if non-nil, zero value otherwise.

### GetRelatedMaterialsOk

`func (o *ProjectArticle) GetRelatedMaterialsOk() (*[]RelatedMaterial, bool)`

GetRelatedMaterialsOk returns a tuple with the RelatedMaterials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedMaterials

`func (o *ProjectArticle) SetRelatedMaterials(v []RelatedMaterial)`

SetRelatedMaterials sets RelatedMaterials field to given value.

### HasRelatedMaterials

`func (o *ProjectArticle) HasRelatedMaterials() bool`

HasRelatedMaterials returns a boolean if a field has been set.

### GetSize

`func (o *ProjectArticle) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ProjectArticle) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ProjectArticle) SetSize(v int64)`

SetSize sets Size field to given value.


### GetStatus

`func (o *ProjectArticle) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProjectArticle) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProjectArticle) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTags

`func (o *ProjectArticle) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ProjectArticle) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ProjectArticle) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetVersion

`func (o *ProjectArticle) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ProjectArticle) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ProjectArticle) SetVersion(v int64)`

SetVersion sets Version field to given value.


### GetDefinedType

`func (o *ProjectArticle) GetDefinedType() int64`

GetDefinedType returns the DefinedType field if non-nil, zero value otherwise.

### GetDefinedTypeOk

`func (o *ProjectArticle) GetDefinedTypeOk() (*int64, bool)`

GetDefinedTypeOk returns a tuple with the DefinedType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinedType

`func (o *ProjectArticle) SetDefinedType(v int64)`

SetDefinedType sets DefinedType field to given value.


### GetDefinedTypeName

`func (o *ProjectArticle) GetDefinedTypeName() string`

GetDefinedTypeName returns the DefinedTypeName field if non-nil, zero value otherwise.

### GetDefinedTypeNameOk

`func (o *ProjectArticle) GetDefinedTypeNameOk() (*string, bool)`

GetDefinedTypeNameOk returns a tuple with the DefinedTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinedTypeName

`func (o *ProjectArticle) SetDefinedTypeName(v string)`

SetDefinedTypeName sets DefinedTypeName field to given value.


### GetDoi

`func (o *ProjectArticle) GetDoi() string`

GetDoi returns the Doi field if non-nil, zero value otherwise.

### GetDoiOk

`func (o *ProjectArticle) GetDoiOk() (*string, bool)`

GetDoiOk returns a tuple with the Doi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoi

`func (o *ProjectArticle) SetDoi(v string)`

SetDoi sets Doi field to given value.


### GetHandle

`func (o *ProjectArticle) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *ProjectArticle) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *ProjectArticle) SetHandle(v string)`

SetHandle sets Handle field to given value.


### GetId

`func (o *ProjectArticle) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectArticle) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectArticle) SetId(v int64)`

SetId sets Id field to given value.


### GetResourceDoi

`func (o *ProjectArticle) GetResourceDoi() string`

GetResourceDoi returns the ResourceDoi field if non-nil, zero value otherwise.

### GetResourceDoiOk

`func (o *ProjectArticle) GetResourceDoiOk() (*string, bool)`

GetResourceDoiOk returns a tuple with the ResourceDoi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceDoi

`func (o *ProjectArticle) SetResourceDoi(v string)`

SetResourceDoi sets ResourceDoi field to given value.


### GetResourceTitle

`func (o *ProjectArticle) GetResourceTitle() string`

GetResourceTitle returns the ResourceTitle field if non-nil, zero value otherwise.

### GetResourceTitleOk

`func (o *ProjectArticle) GetResourceTitleOk() (*string, bool)`

GetResourceTitleOk returns a tuple with the ResourceTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceTitle

`func (o *ProjectArticle) SetResourceTitle(v string)`

SetResourceTitle sets ResourceTitle field to given value.


### GetThumb

`func (o *ProjectArticle) GetThumb() string`

GetThumb returns the Thumb field if non-nil, zero value otherwise.

### GetThumbOk

`func (o *ProjectArticle) GetThumbOk() (*string, bool)`

GetThumbOk returns a tuple with the Thumb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumb

`func (o *ProjectArticle) SetThumb(v string)`

SetThumb sets Thumb field to given value.


### GetTimeline

`func (o *ProjectArticle) GetTimeline() Timeline`

GetTimeline returns the Timeline field if non-nil, zero value otherwise.

### GetTimelineOk

`func (o *ProjectArticle) GetTimelineOk() (*Timeline, bool)`

GetTimelineOk returns a tuple with the Timeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeline

`func (o *ProjectArticle) SetTimeline(v Timeline)`

SetTimeline sets Timeline field to given value.


### GetTitle

`func (o *ProjectArticle) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ProjectArticle) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ProjectArticle) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetUrl

`func (o *ProjectArticle) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ProjectArticle) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ProjectArticle) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetUrlPrivateApi

`func (o *ProjectArticle) GetUrlPrivateApi() string`

GetUrlPrivateApi returns the UrlPrivateApi field if non-nil, zero value otherwise.

### GetUrlPrivateApiOk

`func (o *ProjectArticle) GetUrlPrivateApiOk() (*string, bool)`

GetUrlPrivateApiOk returns a tuple with the UrlPrivateApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlPrivateApi

`func (o *ProjectArticle) SetUrlPrivateApi(v string)`

SetUrlPrivateApi sets UrlPrivateApi field to given value.


### GetUrlPrivateHtml

`func (o *ProjectArticle) GetUrlPrivateHtml() string`

GetUrlPrivateHtml returns the UrlPrivateHtml field if non-nil, zero value otherwise.

### GetUrlPrivateHtmlOk

`func (o *ProjectArticle) GetUrlPrivateHtmlOk() (*string, bool)`

GetUrlPrivateHtmlOk returns a tuple with the UrlPrivateHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlPrivateHtml

`func (o *ProjectArticle) SetUrlPrivateHtml(v string)`

SetUrlPrivateHtml sets UrlPrivateHtml field to given value.


### GetUrlPublicApi

`func (o *ProjectArticle) GetUrlPublicApi() string`

GetUrlPublicApi returns the UrlPublicApi field if non-nil, zero value otherwise.

### GetUrlPublicApiOk

`func (o *ProjectArticle) GetUrlPublicApiOk() (*string, bool)`

GetUrlPublicApiOk returns a tuple with the UrlPublicApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlPublicApi

`func (o *ProjectArticle) SetUrlPublicApi(v string)`

SetUrlPublicApi sets UrlPublicApi field to given value.


### GetUrlPublicHtml

`func (o *ProjectArticle) GetUrlPublicHtml() string`

GetUrlPublicHtml returns the UrlPublicHtml field if non-nil, zero value otherwise.

### GetUrlPublicHtmlOk

`func (o *ProjectArticle) GetUrlPublicHtmlOk() (*string, bool)`

GetUrlPublicHtmlOk returns a tuple with the UrlPublicHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlPublicHtml

`func (o *ProjectArticle) SetUrlPublicHtml(v string)`

SetUrlPublicHtml sets UrlPublicHtml field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


