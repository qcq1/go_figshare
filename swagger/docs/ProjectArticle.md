# ProjectArticle

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedDate** | **string** | Date when article was created | [default to null]
**DefinedType** | **int64** | Type of article identifier | [default to null]
**DefinedTypeName** | **string** | Name of the article type identifier | [default to null]
**Doi** | **string** | DOI | [default to null]
**Handle** | **string** | Handle | [default to null]
**Id** | **int64** | Unique identifier for article | [default to null]
**ResourceDoi** | **string** | Deprecated by related materials. Not applicable to regular users. In a publisher case, this is the publisher article DOI. | [default to null]
**ResourceTitle** | **string** | Deprecated by related materials. Not applicable to regular users. In a publisher case, this is the publisher article title. | [default to null]
**Thumb** | **string** | Thumbnail image | [default to null]
**Timeline** | [***Timeline**](Timeline.md) | Various timeline dates | [default to null]
**Title** | **string** | Title of article | [default to null]
**Url** | **string** | Api endpoint for article | [default to null]
**UrlPrivateApi** | **string** | Private Api endpoint for article | [default to null]
**UrlPrivateHtml** | **string** | Private site endpoint for article | [default to null]
**UrlPublicApi** | **string** | Public Api endpoint for article | [default to null]
**UrlPublicHtml** | **string** | Public site endpoint for article | [default to null]
**Categories** | [**[]Category**](Category.md) | List of categories selected for the article | [optional] [default to null]
**Citation** | **string** | Article citation | [optional] [default to null]
**ConfidentialReason** | **string** | Confidentiality reason | [optional] [default to null]
**Description** | **string** | Article description | [optional] [default to null]
**EmbargoReason** | **string** | Reason for embargo | [optional] [default to null]
**EmbargoTitle** | **string** | Title for embargo | [optional] [default to null]
**Funding** | **string** | Article funding | [optional] [default to null]
**FundingList** | [**[]FundingInformation**](FundingInformation.md) | Full Article funding information | [optional] [default to null]
**HasLinkedFile** | **bool** | True if any files are linked to the article | [optional] [default to null]
**IsConfidential** | **bool** | Article Confidentiality | [optional] [default to null]
**IsEmbargoed** | **bool** | True if article is embargoed | [optional] [default to null]
**IsMetadataRecord** | **bool** | True if article has no files | [optional] [default to null]
**IsPublic** | **bool** | True if article is published | [optional] [default to null]
**Keywords** | **[]string** | List of article keywords. Tags can be used instead | [optional] [default to null]
**License** | [***License**](License.md) | Article selected license | [optional] [default to null]
**MetadataReason** | **string** | Article metadata reason | [optional] [default to null]
**References** | **[]string** | List of references | [optional] [default to null]
**RelatedMaterials** | [**[]RelatedMaterial**](RelatedMaterial.md) | List of related materials; supersedes references and resource DOI/title. | [optional] [default to null]
**Size** | **int64** | Article size | [optional] [default to null]
**Status** | **string** | Article status | [optional] [default to null]
**Tags** | **[]string** | List of article tags. Keywords can be used instead | [optional] [default to null]
**Version** | **int64** | Article version | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


