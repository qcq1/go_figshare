# \ArticlesApi

All URIs are relative to *https://api.figshare.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountArticlePublish**](ArticlesApi.md#AccountArticlePublish) | **Post** /account/articles/{article_id}/publish | Private Article Publish
[**AccountArticleReport**](ArticlesApi.md#AccountArticleReport) | **Get** /account/articles/export | Account Article Report
[**AccountArticleReportGenerate**](ArticlesApi.md#AccountArticleReportGenerate) | **Post** /account/articles/export | Initiate a new Report
[**AccountArticleUnpublish**](ArticlesApi.md#AccountArticleUnpublish) | **Post** /account/articles/{article_id}/unpublish | Public Article Unpublish
[**ArticleDetails**](ArticlesApi.md#ArticleDetails) | **Get** /articles/{article_id} | View article details
[**ArticleFileDetails**](ArticlesApi.md#ArticleFileDetails) | **Get** /articles/{article_id}/files/{file_id} | Article file details
[**ArticleFiles**](ArticlesApi.md#ArticleFiles) | **Get** /articles/{article_id}/files | List article files
[**ArticleVersionConfidentiality**](ArticlesApi.md#ArticleVersionConfidentiality) | **Get** /articles/{article_id}/versions/{version_id}/confidentiality | Public Article Confidentiality for article version
[**ArticleVersionDetails**](ArticlesApi.md#ArticleVersionDetails) | **Get** /articles/{article_id}/versions/{version_id} | Article details for version
[**ArticleVersionEmbargo**](ArticlesApi.md#ArticleVersionEmbargo) | **Get** /articles/{article_id}/versions/{version_id}/embargo | Public Article Embargo for article version
[**ArticleVersionFiles**](ArticlesApi.md#ArticleVersionFiles) | **Get** /articles/{article_id}/versions/{version_id}/files | Article version file details
[**ArticleVersionPartialUpdate**](ArticlesApi.md#ArticleVersionPartialUpdate) | **Patch** /account/articles/{article_id}/versions/{version_id} | Partially update article version
[**ArticleVersionUpdate**](ArticlesApi.md#ArticleVersionUpdate) | **Put** /account/articles/{article_id}/versions/{version_id} | Update article version
[**ArticleVersionUpdateThumb**](ArticlesApi.md#ArticleVersionUpdateThumb) | **Put** /account/articles/{article_id}/versions/{version_id}/update_thumb | Update article version thumbnail
[**ArticleVersions**](ArticlesApi.md#ArticleVersions) | **Get** /articles/{article_id}/versions | List article versions
[**ArticlesList**](ArticlesApi.md#ArticlesList) | **Get** /articles | Public Articles
[**ArticlesSearch**](ArticlesApi.md#ArticlesSearch) | **Post** /articles/search | Public Articles Search
[**PrivateArticleAuthorDelete**](ArticlesApi.md#PrivateArticleAuthorDelete) | **Delete** /account/articles/{article_id}/authors/{author_id} | Delete article author
[**PrivateArticleAuthorsAdd**](ArticlesApi.md#PrivateArticleAuthorsAdd) | **Post** /account/articles/{article_id}/authors | Add article authors
[**PrivateArticleAuthorsList**](ArticlesApi.md#PrivateArticleAuthorsList) | **Get** /account/articles/{article_id}/authors | List article authors
[**PrivateArticleAuthorsReplace**](ArticlesApi.md#PrivateArticleAuthorsReplace) | **Put** /account/articles/{article_id}/authors | Replace article authors
[**PrivateArticleCategoriesAdd**](ArticlesApi.md#PrivateArticleCategoriesAdd) | **Post** /account/articles/{article_id}/categories | Add article categories
[**PrivateArticleCategoriesList**](ArticlesApi.md#PrivateArticleCategoriesList) | **Get** /account/articles/{article_id}/categories | List article categories
[**PrivateArticleCategoriesReplace**](ArticlesApi.md#PrivateArticleCategoriesReplace) | **Put** /account/articles/{article_id}/categories | Replace article categories
[**PrivateArticleCategoryDelete**](ArticlesApi.md#PrivateArticleCategoryDelete) | **Delete** /account/articles/{article_id}/categories/{category_id} | Delete article category
[**PrivateArticleConfidentialityDelete**](ArticlesApi.md#PrivateArticleConfidentialityDelete) | **Delete** /account/articles/{article_id}/confidentiality | Delete article confidentiality
[**PrivateArticleConfidentialityDetails**](ArticlesApi.md#PrivateArticleConfidentialityDetails) | **Get** /account/articles/{article_id}/confidentiality | Article confidentiality details
[**PrivateArticleConfidentialityUpdate**](ArticlesApi.md#PrivateArticleConfidentialityUpdate) | **Put** /account/articles/{article_id}/confidentiality | Update article confidentiality
[**PrivateArticleCreate**](ArticlesApi.md#PrivateArticleCreate) | **Post** /account/articles | Create new Article
[**PrivateArticleDelete**](ArticlesApi.md#PrivateArticleDelete) | **Delete** /account/articles/{article_id} | Delete article
[**PrivateArticleDetails**](ArticlesApi.md#PrivateArticleDetails) | **Get** /account/articles/{article_id} | Article details
[**PrivateArticleDownload**](ArticlesApi.md#PrivateArticleDownload) | **Get** /account/articles/{article_id}/download | Private Article Download
[**PrivateArticleEmbargoDelete**](ArticlesApi.md#PrivateArticleEmbargoDelete) | **Delete** /account/articles/{article_id}/embargo | Delete Article Embargo
[**PrivateArticleEmbargoDetails**](ArticlesApi.md#PrivateArticleEmbargoDetails) | **Get** /account/articles/{article_id}/embargo | Article Embargo Details
[**PrivateArticleEmbargoUpdate**](ArticlesApi.md#PrivateArticleEmbargoUpdate) | **Put** /account/articles/{article_id}/embargo | Update Article Embargo
[**PrivateArticleFile**](ArticlesApi.md#PrivateArticleFile) | **Get** /account/articles/{article_id}/files/{file_id} | Single File
[**PrivateArticleFileDelete**](ArticlesApi.md#PrivateArticleFileDelete) | **Delete** /account/articles/{article_id}/files/{file_id} | File Delete
[**PrivateArticleFilesList**](ArticlesApi.md#PrivateArticleFilesList) | **Get** /account/articles/{article_id}/files | List article files
[**PrivateArticlePartialUpdate**](ArticlesApi.md#PrivateArticlePartialUpdate) | **Patch** /account/articles/{article_id} | Partially update article
[**PrivateArticlePrivateLink**](ArticlesApi.md#PrivateArticlePrivateLink) | **Get** /account/articles/{article_id}/private_links | List private links
[**PrivateArticlePrivateLinkCreate**](ArticlesApi.md#PrivateArticlePrivateLinkCreate) | **Post** /account/articles/{article_id}/private_links | Create private link
[**PrivateArticlePrivateLinkDelete**](ArticlesApi.md#PrivateArticlePrivateLinkDelete) | **Delete** /account/articles/{article_id}/private_links/{link_id} | Disable private link
[**PrivateArticlePrivateLinkUpdate**](ArticlesApi.md#PrivateArticlePrivateLinkUpdate) | **Put** /account/articles/{article_id}/private_links/{link_id} | Update private link
[**PrivateArticleReserveDoi**](ArticlesApi.md#PrivateArticleReserveDoi) | **Post** /account/articles/{article_id}/reserve_doi | Private Article Reserve DOI
[**PrivateArticleReserveHandle**](ArticlesApi.md#PrivateArticleReserveHandle) | **Post** /account/articles/{article_id}/reserve_handle | Private Article Reserve Handle
[**PrivateArticleResource**](ArticlesApi.md#PrivateArticleResource) | **Post** /account/articles/{article_id}/resource | Private Article Resource
[**PrivateArticleUpdate**](ArticlesApi.md#PrivateArticleUpdate) | **Put** /account/articles/{article_id} | Update article
[**PrivateArticleUploadComplete**](ArticlesApi.md#PrivateArticleUploadComplete) | **Post** /account/articles/{article_id}/files/{file_id} | Complete Upload
[**PrivateArticleUploadInitiate**](ArticlesApi.md#PrivateArticleUploadInitiate) | **Post** /account/articles/{article_id}/files | Initiate Upload
[**PrivateArticlesList**](ArticlesApi.md#PrivateArticlesList) | **Get** /account/articles | Private Articles
[**PrivateArticlesSearch**](ArticlesApi.md#PrivateArticlesSearch) | **Post** /account/articles/search | Private Articles search
[**PublicArticleDownload**](ArticlesApi.md#PublicArticleDownload) | **Get** /articles/{article_id}/download | Public Article Download
[**PublicArticleVersionDownload**](ArticlesApi.md#PublicArticleVersionDownload) | **Get** /articles/{article_id}/versions/{version_id}/download | Public Article Version Download


# **AccountArticlePublish**
> Location AccountArticlePublish(ctx, articleId)
Private Article Publish

- If the whole article is under embargo, it will not be published immediately, but when the embargo expires or is lifted. - When an article is published, a new public version will be generated. Any further updates to the article will affect the private article data. In order to make these changes publicly visible, an explicit publish operation is needed.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **articleId** | **int64**| Article unique identifier | 

### Return type

[**Location**](Location.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AccountArticleReport**
> []AccountReport AccountArticleReport(ctx, optional)
Account Article Report

Return status on all reports generated for the account from the oauth credentials

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountArticleReportOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountArticleReportOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **optional.Int64**| A group ID to filter by | 

### Return type

[**[]AccountReport**](AccountReport.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AccountArticleReportGenerate**
> AccountReport AccountArticleReportGenerate(ctx, )
Initiate a new Report

Initiate a new Article Report for this Account. There is a limit of 1 report per day.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**AccountReport**](AccountReport.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AccountArticleUnpublish**
> AccountArticleUnpublish(ctx, articleId, reason)
Public Article Unpublish

Allows authorized users to unpublish an article.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **articleId** | **int64**| Article unique identifier | 
  **reason** | [**ArticleUnpublishData**](ArticleUnpublishData.md)| Article unpublish data | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ArticleDetails**
> ArticleComplete ArticleDetails(ctx, articleId)
View article details

View an article

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **articleId** | **int64**| Article Unique identifier | 

### Return type

[**ArticleComplete**](ArticleComplete.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ArticleFileDetails**
> PublicFile ArticleFileDetails(ctx, articleId, fileId)
Article file details

File by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **articleId** | **int64**| Article Unique identifier | 
  **fileId** | **int64**| File Unique identifier | 

### Return type

[**PublicFile**](PublicFile.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ArticleFiles**
> []PublicFile ArticleFiles(ctx, articleId, optional)
List article files

Files list for article

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **articleId** | **int64**| Article Unique identifier | 
 **optional** | ***ArticleFilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ArticleFilesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int64**| Page number. Used for pagination with page_size | 
 **pageSize** | **optional.Int64**| The number of results included on a page. Used for pagination with page | [default to 10]
 **limit** | **optional.Int64**| Number of results included on a page. Used for pagination with query | 
 **offset** | **optional.Int64**| Where to start the listing (the offset of the first result). Used for pagination with limit | 

### Return type

[**[]PublicFile**](PublicFile.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ArticleVersionConfidentiality**
> ArticleConfidentiality ArticleVersionConfidentiality(ctx, articleId, versionId)
Public Article Confidentiality for article version

Confidentiality for article version. The confidentiality feature is now deprecated. This has been replaced by the new extended embargo functionality and all items that used to be confidential have now been migrated to items with a permanent embargo on files. All API endpoints related to this functionality will remain for backwards compatibility, but will now be attached to the new extended embargo workflows.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **articleId** | **int64**| Article Unique identifier | 
  **versionId** | **int64**| Version Number | 

### Return type

[**ArticleConfidentiality**](ArticleConfidentiality.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ArticleVersionDetails**
> ArticleComplete ArticleVersionDetails(ctx, articleId, versionId)
Article details for version

Article with specified version

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **articleId** | **int64**| Article Unique identifier | 
  **versionId** | **int64**| Article Version Number | 

### Return type

[**ArticleComplete**](ArticleComplete.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ArticleVersionEmbargo**
> ArticleEmbargo ArticleVersionEmbargo(ctx, articleId, versionId)
Public Article Embargo for article version

Embargo for article version

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **articleId** | **int64**| Article Unique identifier | 
  **versionId** | **int64**| Version Number | 

### Return type

[**ArticleEmbargo**](ArticleEmbargo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ArticleVersionFiles**
> []PublicFile ArticleVersionFiles(ctx, articleId, versionId, optional)
Article version file details

File by version id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **articleId** | **int64**| Article Unique identifier | 
  **versionId** | **int64**| Article Version Unique identifier | 
 **optional** | ***ArticleVersionFilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ArticleVersionFilesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int64**| Page number. Used for pagination with page_size | 
 **pageSize** | **optional.Int64**| The number of results included on a page. Used for pagination with page | [default to 10]
 **limit** | **optional.Int64**| Number of results included on a page. Used for pagination with query | 
 **offset** | **optional.Int64**| Where to start the listing (the offset of the first result). Used for pagination with limit | 

### Return type

[**[]PublicFile**](PublicFile.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ArticleVersionPartialUpdate**
> LocationWarningsUpdate ArticleVersionPartialUpdate(ctx, articleId, versionId, article)
Partially update article version

Partially updating an article version by passing only the fields to change.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **articleId** | **int64**| Article unique identifier | 
  **versionId** | **int64**| Article version identifier | 
  **article** | [**ArticleVersionUpdate**](ArticleVersionUpdate.md)| Subset of article version fields to update | 

### Return type

[**LocationWarningsUpdate**](LocationWarningsUpdate.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ArticleVersionUpdate**
> LocationWarningsUpdate ArticleVersionUpdate(ctx, articleId, versionId, article)
Update article version

Updating an article version by passing body parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **articleId** | **int64**| Article unique identifier | 
  **versionId** | **int64**| Article version identifier | 
  **article** | [**ArticleVersionUpdate**](ArticleVersionUpdate.md)| Article description | 

### Return type

[**LocationWarningsUpdate**](LocationWarningsUpdate.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ArticleVersionUpdateThumb**
> ArticleVersionUpdateThumb(ctx, articleId, versionId, fileId)
Update article version thumbnail

For a given public article version update the article thumbnail by choosing one of the associated files

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **articleId** | **int64**| Article unique identifier | 
  **versionId** | **int64**| Article version identifier | 
  **fileId** | [**FileId**](FileId.md)| File ID | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ArticleVersions**
> []ArticleVersions ArticleVersions(ctx, articleId)
List article versions

List public article versions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **articleId** | **int64**| Article Unique identifier | 

### Return type

[**[]ArticleVersions**](ArticleVersions.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ArticlesList**
> []Article ArticlesList(ctx, optional)
Public Articles

Returns a list of public articles

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ArticlesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ArticlesListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xCursor** | [**optional.Interface of string**](.md)| Unique hash used for bypassing the item retrieval limit of 9,000 entities. When using this parameter, please note that the offset parameter will not be available, but the limit parameter will still work as expected. | 
 **page** | **optional.Int64**| Page number. Used for pagination with page_size | 
 **pageSize** | **optional.Int64**| The number of results included on a page. Used for pagination with page | [default to 10]
 **limit** | **optional.Int64**| Number of results included on a page. Used for pagination with query | 
 **offset** | **optional.Int64**| Where to start the listing (the offset of the first result). Used for pagination with limit | 
 **order** | **optional.String**| The field by which to order. Default varies by endpoint/resource. | [default to published_date]
 **orderDirection** | **optional.String**|  | [default to desc]
 **institution** | **optional.Int64**| only return articles from this institution | 
 **publishedSince** | **optional.String**| Filter by article publishing date. Will only return articles published after the date. date(ISO 8601) YYYY-MM-DD or date-time(ISO 8601) YYYY-MM-DDTHH:mm:ssZ | 
 **modifiedSince** | **optional.String**| Filter by article modified date. Will only return articles published after the date. date(ISO 8601) YYYY-MM-DD or date-time(ISO 8601) YYYY-MM-DDTHH:mm:ssZ | 
 **group** | **optional.Int64**| only return articles from this group | 
 **resourceDoi** | **optional.String**| Deprecated by related materials. Only return articles with this resource_doi | 
 **itemType** | **optional.Int64**| Only return articles with the respective type. Mapping for item_type is: 1 - Figure, 2 - Media, 3 - Dataset, 5 - Poster, 6 - Journal contribution, 7 - Presentation, 8 - Thesis, 9 - Software, 11 - Online resource, 12 - Preprint, 13 - Book, 14 - Conference contribution, 15 - Chapter, 16 - Peer review, 17 - Educational resource, 18 - Report, 19 - Standard, 20 - Composition, 21 - Funding, 22 - Physical object, 23 - Data management plan, 24 - Workflow, 25 - Monograph, 26 - Performance, 27 - Event, 28 - Service, 29 - Model | 
 **doi** | **optional.String**| only return articles with this doi | 
 **handle** | **optional.String**| only return articles with this handle | 

### Return type

[**[]Article**](Article.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ArticlesSearch**
> []ArticleWithProject ArticlesSearch(ctx, optional)
Public Articles Search

Returns a list of public articles, filtered by the search parameters

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ArticlesSearchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ArticlesSearchOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xCursor** | [**optional.Interface of string**](.md)| Unique hash used for bypassing the item retrieval limit of 9,000 entities. When using this parameter, please note that the offset parameter will not be available, but the limit parameter will still work as expected. | 
 **search** | [**optional.Interface of ArticleSearch**](ArticleSearch.md)| Search Parameters | 

### Return type

[**[]ArticleWithProject**](ArticleWithProject.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateArticleAuthorDelete**
> PrivateArticleAuthorDelete(ctx, articleId, authorId)
Delete article author

De-associate author from article

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **articleId** | **int64**| Article unique identifier | 
  **authorId** | **int64**| Article Author unique identifier | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateArticleAuthorsAdd**
> PrivateArticleAuthorsAdd(ctx, articleId, authors)
Add article authors

Associate new authors with the article. This will add new authors to the list of already associated authors

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **articleId** | **int64**| Article unique identifier | 
  **authors** | [**AuthorsCreator**](AuthorsCreator.md)| Authors description | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateArticleAuthorsList**
> []Author PrivateArticleAuthorsList(ctx, articleId)
List article authors

List article authors

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **articleId** | **int64**| Article unique identifier | 

### Return type

[**[]Author**](Author.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateArticleAuthorsReplace**
> PrivateArticleAuthorsReplace(ctx, articleId, authors)
Replace article authors

Associate new authors with the article. This will remove all already associated authors and add these new ones

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **articleId** | **int64**| Article unique identifier | 
  **authors** | [**AuthorsCreator**](AuthorsCreator.md)| Authors description | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateArticleCategoriesAdd**
> PrivateArticleCategoriesAdd(ctx, articleId, categories)
Add article categories

Associate new categories with the article. This will add new categories to the list of already associated categories

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **articleId** | **int64**| Article unique identifier | 
  **categories** | [**CategoriesCreator**](CategoriesCreator.md)|  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateArticleCategoriesList**
> []Category PrivateArticleCategoriesList(ctx, articleId)
List article categories

List article categories

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **articleId** | **int64**| Article unique identifier | 

### Return type

[**[]Category**](Category.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateArticleCategoriesReplace**
> PrivateArticleCategoriesReplace(ctx, articleId, categories)
Replace article categories

Associate new categories with the article. This will remove all already associated categories and add these new ones

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **articleId** | **int64**| Article unique identifier | 
  **categories** | [**CategoriesCreator**](CategoriesCreator.md)|  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateArticleCategoryDelete**
> PrivateArticleCategoryDelete(ctx, articleId, categoryId)
Delete article category

De-associate category from article

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **articleId** | **int64**| Article unique identifier | 
  **categoryId** | **int64**| Category unique identifier | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateArticleConfidentialityDelete**
> PrivateArticleConfidentialityDelete(ctx, articleId)
Delete article confidentiality

Delete confidentiality settings. The confidentiality feature is now deprecated. This has been replaced by the new extended embargo functionality and all items that used to be confidential have now been migrated to items with a permanent embargo on files. All API endpoints related to this functionality will remain for backwards compatibility, but will now be attached to the new extended embargo workflows.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **articleId** | **int64**| Article unique identifier | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateArticleConfidentialityDetails**
> ArticleConfidentiality PrivateArticleConfidentialityDetails(ctx, articleId)
Article confidentiality details

View confidentiality settings. The confidentiality feature is now deprecated. This has been replaced by the new extended embargo functionality and all items that used to be confidential have now been migrated to items with a permanent embargo on files. All API endpoints related to this functionality will remain for backwards compatibility, but will now be attached to the new extended embargo workflows.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **articleId** | **int64**| Article unique identifier | 

### Return type

[**ArticleConfidentiality**](ArticleConfidentiality.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateArticleConfidentialityUpdate**
> PrivateArticleConfidentialityUpdate(ctx, articleId, reason)
Update article confidentiality

Update confidentiality settings. The confidentiality feature is now deprecated. This has been replaced by the new extended embargo functionality and all items that used to be confidential have now been migrated to items with a permanent embargo on files. All API endpoints related to this functionality will remain for backwards compatibility, but will now be attached to the new extended embargo workflows.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **articleId** | **int64**| Article unique identifier | 
  **reason** | [**ConfidentialityCreator**](ConfidentialityCreator.md)|  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateArticleCreate**
> LocationWarnings PrivateArticleCreate(ctx, article)
Create new Article

Create a new Article by sending article information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **article** | [**ArticleCreate**](ArticleCreate.md)| Article description | 

### Return type

[**LocationWarnings**](LocationWarnings.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateArticleDelete**
> PrivateArticleDelete(ctx, articleId)
Delete article

Delete an article

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **articleId** | **int64**| Article unique identifier | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateArticleDetails**
> ArticleCompletePrivate PrivateArticleDetails(ctx, articleId)
Article details

View a private article

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **articleId** | **int64**| Article unique identifier | 

### Return type

[**ArticleCompletePrivate**](ArticleCompletePrivate.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateArticleDownload**
> PrivateArticleDownload(ctx, articleId, optional)
Private Article Download

Download files from a private article preserving the folder structure

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **articleId** | **int64**| Article unique identifier | 
 **optional** | ***PrivateArticleDownloadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PrivateArticleDownloadOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **folderPath** | **optional.String**| Folder path to download. If not provided, all files from the article will be downloaded | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateArticleEmbargoDelete**
> PrivateArticleEmbargoDelete(ctx, articleId)
Delete Article Embargo

Will lift the embargo for the specified article

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **articleId** | **int64**| Article unique identifier | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateArticleEmbargoDetails**
> ArticleEmbargo PrivateArticleEmbargoDetails(ctx, articleId)
Article Embargo Details

View a private article embargo details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **articleId** | **int64**| Article unique identifier | 

### Return type

[**ArticleEmbargo**](ArticleEmbargo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateArticleEmbargoUpdate**
> PrivateArticleEmbargoUpdate(ctx, articleId, embargo)
Update Article Embargo

Note: setting an article under whole embargo does not imply that the article will be published when the embargo will expire. You must explicitly call the publish endpoint to enable this functionality.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **articleId** | **int64**| Article unique identifier | 
  **embargo** | [**ArticleEmbargoUpdater**](ArticleEmbargoUpdater.md)| Embargo description | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateArticleFile**
> PrivateFile PrivateArticleFile(ctx, articleId, fileId)
Single File

View details of file for specified article

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **articleId** | **int64**| Article unique identifier | 
  **fileId** | **int64**| File unique identifier | 

### Return type

[**PrivateFile**](PrivateFile.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateArticleFileDelete**
> PrivateArticleFileDelete(ctx, articleId, fileId)
File Delete

Complete file upload

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **articleId** | **int64**| Article unique identifier | 
  **fileId** | **int64**| File unique identifier | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateArticleFilesList**
> []PrivateFile PrivateArticleFilesList(ctx, articleId, optional)
List article files

List private files

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **articleId** | **int64**| Article unique identifier | 
 **optional** | ***PrivateArticleFilesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PrivateArticleFilesListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int64**| Page number. Used for pagination with page_size | 
 **pageSize** | **optional.Int64**| The number of results included on a page. Used for pagination with page | [default to 10]
 **limit** | **optional.Int64**| Number of results included on a page. Used for pagination with query | 
 **offset** | **optional.Int64**| Where to start the listing (the offset of the first result). Used for pagination with limit | 

### Return type

[**[]PrivateFile**](PrivateFile.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateArticlePartialUpdate**
> LocationWarningsUpdate PrivateArticlePartialUpdate(ctx, articleId, article)
Partially update article

Partially update an article by sending only the fields to change.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **articleId** | **int64**| Article unique identifier | 
  **article** | [**ArticleUpdate**](ArticleUpdate.md)| Subset of article fields to update | 

### Return type

[**LocationWarningsUpdate**](LocationWarningsUpdate.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateArticlePrivateLink**
> []PrivateLink PrivateArticlePrivateLink(ctx, articleId)
List private links

List private links

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **articleId** | **int64**| Article unique identifier | 

### Return type

[**[]PrivateLink**](PrivateLink.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateArticlePrivateLinkCreate**
> PrivateLinkResponse PrivateArticlePrivateLinkCreate(ctx, articleId, optional)
Create private link

Create new private link for this article

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **articleId** | **int64**| Article unique identifier | 
 **optional** | ***PrivateArticlePrivateLinkCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PrivateArticlePrivateLinkCreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **privateLink** | [**optional.Interface of PrivateLinkCreator**](PrivateLinkCreator.md)|  | 

### Return type

[**PrivateLinkResponse**](PrivateLinkResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateArticlePrivateLinkDelete**
> PrivateArticlePrivateLinkDelete(ctx, articleId, linkId)
Disable private link

Disable/delete private link for this article

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **articleId** | **int64**| Article unique identifier | 
  **linkId** | **string**| Private link token | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateArticlePrivateLinkUpdate**
> PrivateArticlePrivateLinkUpdate(ctx, articleId, linkId, optional)
Update private link

Update existing private link for this article

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **articleId** | **int64**| Article unique identifier | 
  **linkId** | **string**| Private link token | 
 **optional** | ***PrivateArticlePrivateLinkUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PrivateArticlePrivateLinkUpdateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **privateLink** | [**optional.Interface of PrivateLinkCreator**](PrivateLinkCreator.md)|  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateArticleReserveDoi**
> ArticleDoi PrivateArticleReserveDoi(ctx, articleId)
Private Article Reserve DOI

Reserve DOI for article

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **articleId** | **int64**| Article unique identifier | 

### Return type

[**ArticleDoi**](ArticleDOI.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateArticleReserveHandle**
> ArticleHandle PrivateArticleReserveHandle(ctx, articleId)
Private Article Reserve Handle

Reserve Handle for article

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **articleId** | **int64**| Article unique identifier | 

### Return type

[**ArticleHandle**](ArticleHandle.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateArticleResource**
> Location PrivateArticleResource(ctx, articleId, resource)
Private Article Resource

Edit article resource data.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **articleId** | **int64**| Article unique identifier | 
  **resource** | [**Resource**](Resource.md)| Resource data | 

### Return type

[**Location**](Location.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateArticleUpdate**
> LocationWarningsUpdate PrivateArticleUpdate(ctx, articleId, article)
Update article

Update an article by passing full body parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **articleId** | **int64**| Article unique identifier | 
  **article** | [**ArticleUpdate**](ArticleUpdate.md)| Full article representation | 

### Return type

[**LocationWarningsUpdate**](LocationWarningsUpdate.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateArticleUploadComplete**
> PrivateArticleUploadComplete(ctx, articleId, fileId)
Complete Upload

Complete file upload

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **articleId** | **int64**| Article unique identifier | 
  **fileId** | **int64**| File unique identifier | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateArticleUploadInitiate**
> Location PrivateArticleUploadInitiate(ctx, articleId, file, optional)
Initiate Upload

Initiate a new file upload within the article. Either use the link property to point to an existing file that resides elsewhere and will not be uploaded to Figshare or use the other 3 parameters (md5, name, size).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **articleId** | **int64**| Article unique identifier | 
  **file** | [**FileCreator**](FileCreator.md)|  | 
 **optional** | ***PrivateArticleUploadInitiateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PrivateArticleUploadInitiateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int64**| Page number. Used for pagination with page_size | 
 **pageSize** | **optional.Int64**| The number of results included on a page. Used for pagination with page | [default to 10]
 **limit** | **optional.Int64**| Number of results included on a page. Used for pagination with query | 
 **offset** | **optional.Int64**| Where to start the listing (the offset of the first result). Used for pagination with limit | 

### Return type

[**Location**](Location.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateArticlesList**
> []Article PrivateArticlesList(ctx, optional)
Private Articles

Get Own Articles

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PrivateArticlesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PrivateArticlesListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int64**| Page number. Used for pagination with page_size | 
 **pageSize** | **optional.Int64**| The number of results included on a page. Used for pagination with page | [default to 10]
 **limit** | **optional.Int64**| Number of results included on a page. Used for pagination with query | 
 **offset** | **optional.Int64**| Where to start the listing (the offset of the first result). Used for pagination with limit | 

### Return type

[**[]Article**](Article.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateArticlesSearch**
> []ArticleWithProject PrivateArticlesSearch(ctx, search)
Private Articles search

Returns a list of private articles filtered by the search parameters

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **search** | [**PrivateArticleSearch**](PrivateArticleSearch.md)| Search Parameters | 

### Return type

[**[]ArticleWithProject**](ArticleWithProject.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PublicArticleDownload**
> PublicArticleDownload(ctx, articleId, optional)
Public Article Download

Download files from a public article preserving the folder structure

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **articleId** | **int64**| Article unique identifier | 
 **optional** | ***PublicArticleDownloadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PublicArticleDownloadOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **folderPath** | **optional.String**| Folder path to download. If not provided, all files from the article will be downloaded | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PublicArticleVersionDownload**
> PublicArticleVersionDownload(ctx, articleId, versionId, optional)
Public Article Version Download

Download files from a certain version of an public article preserving the folder structure

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **articleId** | **int64**| Article unique identifier | 
  **versionId** | **int64**| Version Number | 
 **optional** | ***PublicArticleVersionDownloadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PublicArticleVersionDownloadOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **folderPath** | **optional.String**| Folder path to download. If not provided, all files from the article will be downloaded | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

