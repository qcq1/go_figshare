# \CollectionsApi

All URIs are relative to *https://api.figshare.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CollectionArticles**](CollectionsApi.md#CollectionArticles) | **Get** /collections/{collection_id}/articles | Public Collection Articles
[**CollectionDetails**](CollectionsApi.md#CollectionDetails) | **Get** /collections/{collection_id} | Collection details
[**CollectionVersionDetails**](CollectionsApi.md#CollectionVersionDetails) | **Get** /collections/{collection_id}/versions/{version_id} | Collection Version details
[**CollectionVersions**](CollectionsApi.md#CollectionVersions) | **Get** /collections/{collection_id}/versions | Collection Versions list
[**CollectionsList**](CollectionsApi.md#CollectionsList) | **Get** /collections | Public Collections
[**CollectionsSearch**](CollectionsApi.md#CollectionsSearch) | **Post** /collections/search | Public Collections Search
[**PrivateCollectionArticleDelete**](CollectionsApi.md#PrivateCollectionArticleDelete) | **Delete** /account/collections/{collection_id}/articles/{article_id} | Delete collection article
[**PrivateCollectionArticlesAdd**](CollectionsApi.md#PrivateCollectionArticlesAdd) | **Post** /account/collections/{collection_id}/articles | Add collection articles
[**PrivateCollectionArticlesList**](CollectionsApi.md#PrivateCollectionArticlesList) | **Get** /account/collections/{collection_id}/articles | List collection articles
[**PrivateCollectionArticlesReplace**](CollectionsApi.md#PrivateCollectionArticlesReplace) | **Put** /account/collections/{collection_id}/articles | Replace collection articles
[**PrivateCollectionAuthorDelete**](CollectionsApi.md#PrivateCollectionAuthorDelete) | **Delete** /account/collections/{collection_id}/authors/{author_id} | Delete collection author
[**PrivateCollectionAuthorsAdd**](CollectionsApi.md#PrivateCollectionAuthorsAdd) | **Post** /account/collections/{collection_id}/authors | Add collection authors
[**PrivateCollectionAuthorsList**](CollectionsApi.md#PrivateCollectionAuthorsList) | **Get** /account/collections/{collection_id}/authors | List collection authors
[**PrivateCollectionAuthorsReplace**](CollectionsApi.md#PrivateCollectionAuthorsReplace) | **Put** /account/collections/{collection_id}/authors | Replace collection authors
[**PrivateCollectionCategoriesAdd**](CollectionsApi.md#PrivateCollectionCategoriesAdd) | **Post** /account/collections/{collection_id}/categories | Add collection categories
[**PrivateCollectionCategoriesList**](CollectionsApi.md#PrivateCollectionCategoriesList) | **Get** /account/collections/{collection_id}/categories | List collection categories
[**PrivateCollectionCategoriesReplace**](CollectionsApi.md#PrivateCollectionCategoriesReplace) | **Put** /account/collections/{collection_id}/categories | Replace collection categories
[**PrivateCollectionCategoryDelete**](CollectionsApi.md#PrivateCollectionCategoryDelete) | **Delete** /account/collections/{collection_id}/categories/{category_id} | Delete collection category
[**PrivateCollectionCreate**](CollectionsApi.md#PrivateCollectionCreate) | **Post** /account/collections | Create collection
[**PrivateCollectionDelete**](CollectionsApi.md#PrivateCollectionDelete) | **Delete** /account/collections/{collection_id} | Delete collection
[**PrivateCollectionDetails**](CollectionsApi.md#PrivateCollectionDetails) | **Get** /account/collections/{collection_id} | Collection details
[**PrivateCollectionPatch**](CollectionsApi.md#PrivateCollectionPatch) | **Patch** /account/collections/{collection_id} | Partially update collection
[**PrivateCollectionPrivateLinkCreate**](CollectionsApi.md#PrivateCollectionPrivateLinkCreate) | **Post** /account/collections/{collection_id}/private_links | Create collection private link
[**PrivateCollectionPrivateLinkDelete**](CollectionsApi.md#PrivateCollectionPrivateLinkDelete) | **Delete** /account/collections/{collection_id}/private_links/{link_id} | Disable private link
[**PrivateCollectionPrivateLinkDetails**](CollectionsApi.md#PrivateCollectionPrivateLinkDetails) | **Get** /account/collections/{collection_id}/private_links/{link_id} | View collection private link
[**PrivateCollectionPrivateLinkUpdate**](CollectionsApi.md#PrivateCollectionPrivateLinkUpdate) | **Put** /account/collections/{collection_id}/private_links/{link_id} | Update collection private link
[**PrivateCollectionPrivateLinksList**](CollectionsApi.md#PrivateCollectionPrivateLinksList) | **Get** /account/collections/{collection_id}/private_links | List collection private links
[**PrivateCollectionPublish**](CollectionsApi.md#PrivateCollectionPublish) | **Post** /account/collections/{collection_id}/publish | Private Collection Publish
[**PrivateCollectionReserveDoi**](CollectionsApi.md#PrivateCollectionReserveDoi) | **Post** /account/collections/{collection_id}/reserve_doi | Private Collection Reserve DOI
[**PrivateCollectionReserveHandle**](CollectionsApi.md#PrivateCollectionReserveHandle) | **Post** /account/collections/{collection_id}/reserve_handle | Private Collection Reserve Handle
[**PrivateCollectionResource**](CollectionsApi.md#PrivateCollectionResource) | **Post** /account/collections/{collection_id}/resource | Private Collection Resource
[**PrivateCollectionUpdate**](CollectionsApi.md#PrivateCollectionUpdate) | **Put** /account/collections/{collection_id} | Update collection
[**PrivateCollectionsList**](CollectionsApi.md#PrivateCollectionsList) | **Get** /account/collections | Private Collections List
[**PrivateCollectionsSearch**](CollectionsApi.md#PrivateCollectionsSearch) | **Post** /account/collections/search | Private Collections Search


# **CollectionArticles**
> []Article CollectionArticles(ctx, collectionId, optional)
Public Collection Articles

Returns a list of public collection articles

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **collectionId** | **int64**| Collection Unique identifier | 
 **optional** | ***CollectionArticlesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionArticlesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int64**| Page number. Used for pagination with page_size | 
 **pageSize** | **optional.Int64**| The number of results included on a page. Used for pagination with page | [default to 10]
 **limit** | **optional.Int64**| Number of results included on a page. Used for pagination with query | 
 **offset** | **optional.Int64**| Where to start the listing (the offset of the first result). Used for pagination with limit | 

### Return type

[**[]Article**](Article.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CollectionDetails**
> CollectionComplete CollectionDetails(ctx, collectionId)
Collection details

View a collection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **collectionId** | **int64**| Collection Unique identifier | 

### Return type

[**CollectionComplete**](CollectionComplete.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CollectionVersionDetails**
> CollectionComplete CollectionVersionDetails(ctx, collectionId, versionId)
Collection Version details

View details for a certain version of a collection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **collectionId** | **int64**| Collection Unique identifier | 
  **versionId** | **int64**| Version Number | 

### Return type

[**CollectionComplete**](CollectionComplete.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CollectionVersions**
> []CollectionVersions CollectionVersions(ctx, collectionId)
Collection Versions list

Returns a list of public collection Versions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **collectionId** | **int64**| Collection Unique identifier | 

### Return type

[**[]CollectionVersions**](CollectionVersions.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CollectionsList**
> []Collection CollectionsList(ctx, optional)
Public Collections

Returns a list of public collections

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CollectionsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xCursor** | [**optional.Interface of string**](.md)| Unique hash used for bypassing the item retrieval limit of 9,000 entities. When using this parameter, please note that the offset parameter will not be available, but the limit parameter will still work as expected. | 
 **page** | **optional.Int64**| Page number. Used for pagination with page_size | 
 **pageSize** | **optional.Int64**| The number of results included on a page. Used for pagination with page | [default to 10]
 **limit** | **optional.Int64**| Number of results included on a page. Used for pagination with query | 
 **offset** | **optional.Int64**| Where to start the listing (the offset of the first result). Used for pagination with limit | 
 **order** | **optional.String**| The field by which to order. Default varies by endpoint/resource. | [default to published_date]
 **orderDirection** | **optional.String**|  | [default to desc]
 **institution** | **optional.Int64**| only return collections from this institution | 
 **publishedSince** | **optional.String**| Filter by collection publishing date. Will only return collections published after the date. date(ISO 8601) YYYY-MM-DD | 
 **modifiedSince** | **optional.String**| Filter by collection modified date. Will only return collections published after the date. date(ISO 8601) YYYY-MM-DD | 
 **group** | **optional.Int64**| only return collections from this group | 
 **resourceDoi** | **optional.String**| only return collections with this resource_doi | 
 **doi** | **optional.String**| only return collections with this doi | 
 **handle** | **optional.String**| only return collections with this handle | 

### Return type

[**[]Collection**](Collection.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CollectionsSearch**
> []Collection CollectionsSearch(ctx, optional)
Public Collections Search

Returns a list of public collections

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CollectionsSearchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionsSearchOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xCursor** | [**optional.Interface of string**](.md)| Unique hash used for bypassing the item retrieval limit of 9,000 entities. When using this parameter, please note that the offset parameter will not be available, but the limit parameter will still work as expected. | 
 **search** | [**optional.Interface of CollectionSearch**](CollectionSearch.md)| Search Parameters | 

### Return type

[**[]Collection**](Collection.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateCollectionArticleDelete**
> PrivateCollectionArticleDelete(ctx, collectionId, articleId)
Delete collection article

De-associate article from collection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **collectionId** | **int64**| Collection unique identifier | 
  **articleId** | **int64**| Collection article unique identifier | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateCollectionArticlesAdd**
> Location PrivateCollectionArticlesAdd(ctx, collectionId, articles)
Add collection articles

Associate new articles with the collection. This will add new articles to the list of already associated articles

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **collectionId** | **int64**| Collection unique identifier | 
  **articles** | [**ArticlesCreator**](ArticlesCreator.md)| Articles list | 

### Return type

[**Location**](Location.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateCollectionArticlesList**
> []Article PrivateCollectionArticlesList(ctx, collectionId, optional)
List collection articles

List collection articles

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **collectionId** | **int64**| Collection unique identifier | 
 **optional** | ***PrivateCollectionArticlesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PrivateCollectionArticlesListOpts struct

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

# **PrivateCollectionArticlesReplace**
> PrivateCollectionArticlesReplace(ctx, collectionId, articles)
Replace collection articles

Associate new articles with the collection. This will remove all already associated articles and add these new ones

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **collectionId** | **int64**| Collection unique identifier | 
  **articles** | [**ArticlesCreator**](ArticlesCreator.md)| Articles List | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateCollectionAuthorDelete**
> PrivateCollectionAuthorDelete(ctx, collectionId, authorId)
Delete collection author

Delete collection author

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **collectionId** | **int64**| Collection unique identifier | 
  **authorId** | **int64**| Collection Author unique identifier | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateCollectionAuthorsAdd**
> Location PrivateCollectionAuthorsAdd(ctx, collectionId, authors)
Add collection authors

Associate new authors with the collection. This will add new authors to the list of already associated authors

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **collectionId** | **int64**| Collection unique identifier | 
  **authors** | [**AuthorsCreator**](AuthorsCreator.md)| List of authors | 

### Return type

[**Location**](Location.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateCollectionAuthorsList**
> []Author PrivateCollectionAuthorsList(ctx, collectionId)
List collection authors

List collection authors

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **collectionId** | **int64**| Collection unique identifier | 

### Return type

[**[]Author**](Author.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateCollectionAuthorsReplace**
> PrivateCollectionAuthorsReplace(ctx, collectionId, authors)
Replace collection authors

Associate new authors with the collection. This will remove all already associated authors and add these new ones

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **collectionId** | **int64**| Collection unique identifier | 
  **authors** | [**AuthorsCreator**](AuthorsCreator.md)| List of authors | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateCollectionCategoriesAdd**
> Location PrivateCollectionCategoriesAdd(ctx, collectionId, categories)
Add collection categories

Associate new categories with the collection. This will add new categories to the list of already associated categories

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **collectionId** | **int64**| Collection unique identifier | 
  **categories** | [**CategoriesCreator**](CategoriesCreator.md)| Categories list | 

### Return type

[**Location**](Location.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateCollectionCategoriesList**
> []Category PrivateCollectionCategoriesList(ctx, collectionId)
List collection categories

List collection categories

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **collectionId** | **int64**| Collection unique identifier | 

### Return type

[**[]Category**](Category.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateCollectionCategoriesReplace**
> PrivateCollectionCategoriesReplace(ctx, collectionId, categories)
Replace collection categories

Associate new categories with the collection. This will remove all already associated categories and add these new ones

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **collectionId** | **int64**| Collection unique identifier | 
  **categories** | [**CategoriesCreator**](CategoriesCreator.md)| Categories list | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateCollectionCategoryDelete**
> PrivateCollectionCategoryDelete(ctx, collectionId, categoryId)
Delete collection category

De-associate category from collection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **collectionId** | **int64**| Collection unique identifier | 
  **categoryId** | **int64**| Collection category unique identifier | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateCollectionCreate**
> LocationWarnings PrivateCollectionCreate(ctx, collection)
Create collection

Create a new Collection by sending collection information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **collection** | [**CollectionCreate**](CollectionCreate.md)| Collection description | 

### Return type

[**LocationWarnings**](LocationWarnings.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateCollectionDelete**
> PrivateCollectionDelete(ctx, collectionId)
Delete collection

Delete n collection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **collectionId** | **int64**| Collection Unique identifier | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateCollectionDetails**
> CollectionCompletePrivate PrivateCollectionDetails(ctx, collectionId)
Collection details

View a collection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **collectionId** | **int64**| Collection Unique identifier | 

### Return type

[**CollectionCompletePrivate**](CollectionCompletePrivate.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateCollectionPatch**
> LocationWarningsUpdate PrivateCollectionPatch(ctx, collectionId, collection)
Partially update collection

Partially update a collection by sending only the fields to change.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **collectionId** | **int64**| Collection Unique identifier | 
  **collection** | [**CollectionUpdate**](CollectionUpdate.md)| Subset of collection fields to update | 

### Return type

[**LocationWarningsUpdate**](LocationWarningsUpdate.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateCollectionPrivateLinkCreate**
> PrivateLinkResponse PrivateCollectionPrivateLinkCreate(ctx, collectionId, optional)
Create collection private link

Create new private link

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **collectionId** | **int64**| Collection unique identifier | 
 **optional** | ***PrivateCollectionPrivateLinkCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PrivateCollectionPrivateLinkCreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **privateLink** | [**optional.Interface of CollectionPrivateLinkCreator**](CollectionPrivateLinkCreator.md)|  | 

### Return type

[**PrivateLinkResponse**](PrivateLinkResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateCollectionPrivateLinkDelete**
> PrivateCollectionPrivateLinkDelete(ctx, collectionId, linkId)
Disable private link

Disable/delete private link for this collection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **collectionId** | **int64**| Collection unique identifier | 
  **linkId** | **string**| Private link token | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateCollectionPrivateLinkDetails**
> PrivateLink PrivateCollectionPrivateLinkDetails(ctx, collectionId, linkId)
View collection private link

View existing private link for this collection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **collectionId** | **int64**| Collection unique identifier | 
  **linkId** | **string**| Private link token | 

### Return type

[**PrivateLink**](PrivateLink.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateCollectionPrivateLinkUpdate**
> PrivateCollectionPrivateLinkUpdate(ctx, collectionId, linkId, optional)
Update collection private link

Update existing private link for this collection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **collectionId** | **int64**| Collection unique identifier | 
  **linkId** | **string**| Private link token | 
 **optional** | ***PrivateCollectionPrivateLinkUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PrivateCollectionPrivateLinkUpdateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **privateLink** | [**optional.Interface of CollectionPrivateLinkCreator**](CollectionPrivateLinkCreator.md)|  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateCollectionPrivateLinksList**
> []PrivateLink PrivateCollectionPrivateLinksList(ctx, collectionId)
List collection private links

List article private links

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **collectionId** | **int64**| Collection unique identifier | 

### Return type

[**[]PrivateLink**](PrivateLink.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateCollectionPublish**
> Location PrivateCollectionPublish(ctx, collectionId)
Private Collection Publish

When a collection is published, a new public version will be generated. Any further updates to the collection will affect the private collection data. In order to make these changes publicly visible, an explicit publish operation is needed.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **collectionId** | **int64**| Collection Unique identifier | 

### Return type

[**Location**](Location.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateCollectionReserveDoi**
> CollectionDoi PrivateCollectionReserveDoi(ctx, collectionId)
Private Collection Reserve DOI

Reserve DOI for collection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **collectionId** | **int64**| Collection Unique identifier | 

### Return type

[**CollectionDoi**](CollectionDOI.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateCollectionReserveHandle**
> CollectionHandle PrivateCollectionReserveHandle(ctx, collectionId)
Private Collection Reserve Handle

Reserve Handle for collection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **collectionId** | **int64**| Collection Unique identifier | 

### Return type

[**CollectionHandle**](CollectionHandle.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateCollectionResource**
> Location PrivateCollectionResource(ctx, collectionId, resource)
Private Collection Resource

Edit collection resource data.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **collectionId** | **int64**| Collection unique identifier | 
  **resource** | [**Resource**](Resource.md)| Resource data | 

### Return type

[**Location**](Location.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateCollectionUpdate**
> LocationWarningsUpdate PrivateCollectionUpdate(ctx, collectionId, collection)
Update collection

Update a collection by passing full body parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **collectionId** | **int64**| Collection Unique identifier | 
  **collection** | [**CollectionUpdate**](CollectionUpdate.md)| Collection description | 

### Return type

[**LocationWarningsUpdate**](LocationWarningsUpdate.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateCollectionsList**
> []Collection PrivateCollectionsList(ctx, optional)
Private Collections List

List private collections

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PrivateCollectionsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PrivateCollectionsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int64**| Page number. Used for pagination with page_size | 
 **pageSize** | **optional.Int64**| The number of results included on a page. Used for pagination with page | [default to 10]
 **limit** | **optional.Int64**| Number of results included on a page. Used for pagination with query | 
 **offset** | **optional.Int64**| Where to start the listing (the offset of the first result). Used for pagination with limit | 
 **order** | **optional.String**| The field by which to order. Default varies by endpoint/resource. | [default to published_date]
 **orderDirection** | **optional.String**|  | [default to desc]

### Return type

[**[]Collection**](Collection.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateCollectionsSearch**
> []Collection PrivateCollectionsSearch(ctx, search)
Private Collections Search

Returns a list of private Collections

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **search** | [**PrivateCollectionSearch**](PrivateCollectionSearch.md)| Search Parameters | 

### Return type

[**[]Collection**](Collection.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

