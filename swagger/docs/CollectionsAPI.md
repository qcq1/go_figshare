# \CollectionsAPI

All URIs are relative to *https://api.figinternal.dev/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CollectionArticles**](CollectionsAPI.md#CollectionArticles) | **Get** /collections/{collection_id}/articles | Public Collection Articles
[**CollectionDetails**](CollectionsAPI.md#CollectionDetails) | **Get** /collections/{collection_id} | Collection details
[**CollectionVersionDetails**](CollectionsAPI.md#CollectionVersionDetails) | **Get** /collections/{collection_id}/versions/{version_id} | Collection Version details
[**CollectionVersions**](CollectionsAPI.md#CollectionVersions) | **Get** /collections/{collection_id}/versions | Collection Versions list
[**CollectionsList**](CollectionsAPI.md#CollectionsList) | **Get** /collections | Public Collections
[**CollectionsSearch**](CollectionsAPI.md#CollectionsSearch) | **Post** /collections/search | Public Collections Search
[**PrivateCollectionArticleDelete**](CollectionsAPI.md#PrivateCollectionArticleDelete) | **Delete** /account/collections/{collection_id}/articles/{article_id} | Delete collection article
[**PrivateCollectionArticlesAdd**](CollectionsAPI.md#PrivateCollectionArticlesAdd) | **Post** /account/collections/{collection_id}/articles | Add collection articles
[**PrivateCollectionArticlesList**](CollectionsAPI.md#PrivateCollectionArticlesList) | **Get** /account/collections/{collection_id}/articles | List collection articles
[**PrivateCollectionArticlesReplace**](CollectionsAPI.md#PrivateCollectionArticlesReplace) | **Put** /account/collections/{collection_id}/articles | Replace collection articles
[**PrivateCollectionAuthorDelete**](CollectionsAPI.md#PrivateCollectionAuthorDelete) | **Delete** /account/collections/{collection_id}/authors/{author_id} | Delete collection author
[**PrivateCollectionAuthorsAdd**](CollectionsAPI.md#PrivateCollectionAuthorsAdd) | **Post** /account/collections/{collection_id}/authors | Add collection authors
[**PrivateCollectionAuthorsList**](CollectionsAPI.md#PrivateCollectionAuthorsList) | **Get** /account/collections/{collection_id}/authors | List collection authors
[**PrivateCollectionAuthorsReplace**](CollectionsAPI.md#PrivateCollectionAuthorsReplace) | **Put** /account/collections/{collection_id}/authors | Replace collection authors
[**PrivateCollectionCategoriesAdd**](CollectionsAPI.md#PrivateCollectionCategoriesAdd) | **Post** /account/collections/{collection_id}/categories | Add collection categories
[**PrivateCollectionCategoriesList**](CollectionsAPI.md#PrivateCollectionCategoriesList) | **Get** /account/collections/{collection_id}/categories | List collection categories
[**PrivateCollectionCategoriesReplace**](CollectionsAPI.md#PrivateCollectionCategoriesReplace) | **Put** /account/collections/{collection_id}/categories | Replace collection categories
[**PrivateCollectionCategoryDelete**](CollectionsAPI.md#PrivateCollectionCategoryDelete) | **Delete** /account/collections/{collection_id}/categories/{category_id} | Delete collection category
[**PrivateCollectionCreate**](CollectionsAPI.md#PrivateCollectionCreate) | **Post** /account/collections | Create collection
[**PrivateCollectionDelete**](CollectionsAPI.md#PrivateCollectionDelete) | **Delete** /account/collections/{collection_id} | Delete collection
[**PrivateCollectionDetails**](CollectionsAPI.md#PrivateCollectionDetails) | **Get** /account/collections/{collection_id} | Collection details
[**PrivateCollectionPatch**](CollectionsAPI.md#PrivateCollectionPatch) | **Patch** /account/collections/{collection_id} | Partially update collection
[**PrivateCollectionPrivateLinkCreate**](CollectionsAPI.md#PrivateCollectionPrivateLinkCreate) | **Post** /account/collections/{collection_id}/private_links | Create collection private link
[**PrivateCollectionPrivateLinkDelete**](CollectionsAPI.md#PrivateCollectionPrivateLinkDelete) | **Delete** /account/collections/{collection_id}/private_links/{link_id} | Disable private link
[**PrivateCollectionPrivateLinkDetails**](CollectionsAPI.md#PrivateCollectionPrivateLinkDetails) | **Get** /account/collections/{collection_id}/private_links/{link_id} | View collection private link
[**PrivateCollectionPrivateLinkUpdate**](CollectionsAPI.md#PrivateCollectionPrivateLinkUpdate) | **Put** /account/collections/{collection_id}/private_links/{link_id} | Update collection private link
[**PrivateCollectionPrivateLinksList**](CollectionsAPI.md#PrivateCollectionPrivateLinksList) | **Get** /account/collections/{collection_id}/private_links | List collection private links
[**PrivateCollectionPublish**](CollectionsAPI.md#PrivateCollectionPublish) | **Post** /account/collections/{collection_id}/publish | Private Collection Publish
[**PrivateCollectionReserveDoi**](CollectionsAPI.md#PrivateCollectionReserveDoi) | **Post** /account/collections/{collection_id}/reserve_doi | Private Collection Reserve DOI
[**PrivateCollectionReserveHandle**](CollectionsAPI.md#PrivateCollectionReserveHandle) | **Post** /account/collections/{collection_id}/reserve_handle | Private Collection Reserve Handle
[**PrivateCollectionResource**](CollectionsAPI.md#PrivateCollectionResource) | **Post** /account/collections/{collection_id}/resource | Private Collection Resource
[**PrivateCollectionUpdate**](CollectionsAPI.md#PrivateCollectionUpdate) | **Put** /account/collections/{collection_id} | Update collection
[**PrivateCollectionsList**](CollectionsAPI.md#PrivateCollectionsList) | **Get** /account/collections | Private Collections List
[**PrivateCollectionsSearch**](CollectionsAPI.md#PrivateCollectionsSearch) | **Post** /account/collections/search | Private Collections Search



## CollectionArticles

> []Article CollectionArticles(ctx, collectionId).Page(page).PageSize(pageSize).Limit(limit).Offset(offset).Execute()

Public Collection Articles



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	collectionId := int64(789) // int64 | Collection Unique identifier
	page := int64(789) // int64 | Page number. Used for pagination with page_size (optional)
	pageSize := int64(789) // int64 | The number of results included on a page. Used for pagination with page (optional) (default to 10)
	limit := int64(789) // int64 | Number of results included on a page. Used for pagination with query (optional)
	offset := int64(789) // int64 | Where to start the listing (the offset of the first result). Used for pagination with limit (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionsAPI.CollectionArticles(context.Background(), collectionId).Page(page).PageSize(pageSize).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.CollectionArticles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CollectionArticles`: []Article
	fmt.Fprintf(os.Stdout, "Response from `CollectionsAPI.CollectionArticles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | Collection Unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCollectionArticlesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int64** | Page number. Used for pagination with page_size | 
 **pageSize** | **int64** | The number of results included on a page. Used for pagination with page | [default to 10]
 **limit** | **int64** | Number of results included on a page. Used for pagination with query | 
 **offset** | **int64** | Where to start the listing (the offset of the first result). Used for pagination with limit | 

### Return type

[**[]Article**](Article.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CollectionDetails

> CollectionComplete CollectionDetails(ctx, collectionId).Execute()

Collection details



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	collectionId := int64(789) // int64 | Collection Unique identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionsAPI.CollectionDetails(context.Background(), collectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.CollectionDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CollectionDetails`: CollectionComplete
	fmt.Fprintf(os.Stdout, "Response from `CollectionsAPI.CollectionDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | Collection Unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCollectionDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CollectionComplete**](CollectionComplete.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CollectionVersionDetails

> CollectionComplete CollectionVersionDetails(ctx, collectionId, versionId).Execute()

Collection Version details



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	collectionId := int64(789) // int64 | Collection Unique identifier
	versionId := int64(789) // int64 | Version Number

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionsAPI.CollectionVersionDetails(context.Background(), collectionId, versionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.CollectionVersionDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CollectionVersionDetails`: CollectionComplete
	fmt.Fprintf(os.Stdout, "Response from `CollectionsAPI.CollectionVersionDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | Collection Unique identifier | 
**versionId** | **int64** | Version Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiCollectionVersionDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CollectionComplete**](CollectionComplete.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CollectionVersions

> []CollectionVersions CollectionVersions(ctx, collectionId).Execute()

Collection Versions list



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	collectionId := int64(789) // int64 | Collection Unique identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionsAPI.CollectionVersions(context.Background(), collectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.CollectionVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CollectionVersions`: []CollectionVersions
	fmt.Fprintf(os.Stdout, "Response from `CollectionsAPI.CollectionVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | Collection Unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCollectionVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]CollectionVersions**](CollectionVersions.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CollectionsList

> []Collection CollectionsList(ctx).XCursor(xCursor).Page(page).PageSize(pageSize).Limit(limit).Offset(offset).Order(order).OrderDirection(orderDirection).Institution(institution).PublishedSince(publishedSince).ModifiedSince(modifiedSince).Group(group).ResourceDoi(resourceDoi).Doi(doi).Handle(handle).Execute()

Public Collections



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	xCursor := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique hash used for bypassing the item retrieval limit of 9,000 entities. When using this parameter, please note that the offset parameter will not be available, but the limit parameter will still work as expected. (optional)
	page := int64(789) // int64 | Page number. Used for pagination with page_size (optional)
	pageSize := int64(789) // int64 | The number of results included on a page. Used for pagination with page (optional) (default to 10)
	limit := int64(789) // int64 | Number of results included on a page. Used for pagination with query (optional)
	offset := int64(789) // int64 | Where to start the listing (the offset of the first result). Used for pagination with limit (optional)
	order := "order_example" // string | The field by which to order. Default varies by endpoint/resource. (optional) (default to "published_date")
	orderDirection := "orderDirection_example" // string |  (optional) (default to "desc")
	institution := int64(789) // int64 | only return collections from this institution (optional)
	publishedSince := "publishedSince_example" // string | Filter by collection publishing date. Will only return collections published after the date. date(ISO 8601) YYYY-MM-DD (optional)
	modifiedSince := "modifiedSince_example" // string | Filter by collection modified date. Will only return collections published after the date. date(ISO 8601) YYYY-MM-DD (optional)
	group := int64(789) // int64 | only return collections from this group (optional)
	resourceDoi := "resourceDoi_example" // string | only return collections with this resource_doi (optional)
	doi := "doi_example" // string | only return collections with this doi (optional)
	handle := "handle_example" // string | only return collections with this handle (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionsAPI.CollectionsList(context.Background()).XCursor(xCursor).Page(page).PageSize(pageSize).Limit(limit).Offset(offset).Order(order).OrderDirection(orderDirection).Institution(institution).PublishedSince(publishedSince).ModifiedSince(modifiedSince).Group(group).ResourceDoi(resourceDoi).Doi(doi).Handle(handle).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.CollectionsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CollectionsList`: []Collection
	fmt.Fprintf(os.Stdout, "Response from `CollectionsAPI.CollectionsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCollectionsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xCursor** | **string** | Unique hash used for bypassing the item retrieval limit of 9,000 entities. When using this parameter, please note that the offset parameter will not be available, but the limit parameter will still work as expected. | 
 **page** | **int64** | Page number. Used for pagination with page_size | 
 **pageSize** | **int64** | The number of results included on a page. Used for pagination with page | [default to 10]
 **limit** | **int64** | Number of results included on a page. Used for pagination with query | 
 **offset** | **int64** | Where to start the listing (the offset of the first result). Used for pagination with limit | 
 **order** | **string** | The field by which to order. Default varies by endpoint/resource. | [default to &quot;published_date&quot;]
 **orderDirection** | **string** |  | [default to &quot;desc&quot;]
 **institution** | **int64** | only return collections from this institution | 
 **publishedSince** | **string** | Filter by collection publishing date. Will only return collections published after the date. date(ISO 8601) YYYY-MM-DD | 
 **modifiedSince** | **string** | Filter by collection modified date. Will only return collections published after the date. date(ISO 8601) YYYY-MM-DD | 
 **group** | **int64** | only return collections from this group | 
 **resourceDoi** | **string** | only return collections with this resource_doi | 
 **doi** | **string** | only return collections with this doi | 
 **handle** | **string** | only return collections with this handle | 

### Return type

[**[]Collection**](Collection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CollectionsSearch

> []Collection CollectionsSearch(ctx).XCursor(xCursor).Search(search).Execute()

Public Collections Search



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	xCursor := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique hash used for bypassing the item retrieval limit of 9,000 entities. When using this parameter, please note that the offset parameter will not be available, but the limit parameter will still work as expected. (optional)
	search := *openapiclient.NewCollectionSearch() // CollectionSearch | Search Parameters (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionsAPI.CollectionsSearch(context.Background()).XCursor(xCursor).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.CollectionsSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CollectionsSearch`: []Collection
	fmt.Fprintf(os.Stdout, "Response from `CollectionsAPI.CollectionsSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCollectionsSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xCursor** | **string** | Unique hash used for bypassing the item retrieval limit of 9,000 entities. When using this parameter, please note that the offset parameter will not be available, but the limit parameter will still work as expected. | 
 **search** | [**CollectionSearch**](CollectionSearch.md) | Search Parameters | 

### Return type

[**[]Collection**](Collection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivateCollectionArticleDelete

> PrivateCollectionArticleDelete(ctx, collectionId, articleId).Execute()

Delete collection article



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	collectionId := int64(789) // int64 | Collection unique identifier
	articleId := int64(789) // int64 | Collection article unique identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CollectionsAPI.PrivateCollectionArticleDelete(context.Background(), collectionId, articleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.PrivateCollectionArticleDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | Collection unique identifier | 
**articleId** | **int64** | Collection article unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateCollectionArticleDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivateCollectionArticlesAdd

> Location PrivateCollectionArticlesAdd(ctx, collectionId).Articles(articles).Execute()

Add collection articles



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	collectionId := int64(789) // int64 | Collection unique identifier
	articles := *openapiclient.NewArticlesCreator([]int64{int64(123)}) // ArticlesCreator | Articles list

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionsAPI.PrivateCollectionArticlesAdd(context.Background(), collectionId).Articles(articles).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.PrivateCollectionArticlesAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivateCollectionArticlesAdd`: Location
	fmt.Fprintf(os.Stdout, "Response from `CollectionsAPI.PrivateCollectionArticlesAdd`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | Collection unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateCollectionArticlesAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **articles** | [**ArticlesCreator**](ArticlesCreator.md) | Articles list | 

### Return type

[**Location**](Location.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivateCollectionArticlesList

> []Article PrivateCollectionArticlesList(ctx, collectionId).Page(page).PageSize(pageSize).Limit(limit).Offset(offset).Execute()

List collection articles



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	collectionId := int64(789) // int64 | Collection unique identifier
	page := int64(789) // int64 | Page number. Used for pagination with page_size (optional)
	pageSize := int64(789) // int64 | The number of results included on a page. Used for pagination with page (optional) (default to 10)
	limit := int64(789) // int64 | Number of results included on a page. Used for pagination with query (optional)
	offset := int64(789) // int64 | Where to start the listing (the offset of the first result). Used for pagination with limit (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionsAPI.PrivateCollectionArticlesList(context.Background(), collectionId).Page(page).PageSize(pageSize).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.PrivateCollectionArticlesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivateCollectionArticlesList`: []Article
	fmt.Fprintf(os.Stdout, "Response from `CollectionsAPI.PrivateCollectionArticlesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | Collection unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateCollectionArticlesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int64** | Page number. Used for pagination with page_size | 
 **pageSize** | **int64** | The number of results included on a page. Used for pagination with page | [default to 10]
 **limit** | **int64** | Number of results included on a page. Used for pagination with query | 
 **offset** | **int64** | Where to start the listing (the offset of the first result). Used for pagination with limit | 

### Return type

[**[]Article**](Article.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivateCollectionArticlesReplace

> PrivateCollectionArticlesReplace(ctx, collectionId).Articles(articles).Execute()

Replace collection articles



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	collectionId := int64(789) // int64 | Collection unique identifier
	articles := *openapiclient.NewArticlesCreator([]int64{int64(123)}) // ArticlesCreator | Articles List

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CollectionsAPI.PrivateCollectionArticlesReplace(context.Background(), collectionId).Articles(articles).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.PrivateCollectionArticlesReplace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | Collection unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateCollectionArticlesReplaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **articles** | [**ArticlesCreator**](ArticlesCreator.md) | Articles List | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivateCollectionAuthorDelete

> PrivateCollectionAuthorDelete(ctx, collectionId, authorId).Execute()

Delete collection author



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	collectionId := int64(789) // int64 | Collection unique identifier
	authorId := int64(789) // int64 | Collection Author unique identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CollectionsAPI.PrivateCollectionAuthorDelete(context.Background(), collectionId, authorId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.PrivateCollectionAuthorDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | Collection unique identifier | 
**authorId** | **int64** | Collection Author unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateCollectionAuthorDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivateCollectionAuthorsAdd

> Location PrivateCollectionAuthorsAdd(ctx, collectionId).Authors(authors).Execute()

Add collection authors



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	collectionId := int64(789) // int64 | Collection unique identifier
	authors := *openapiclient.NewAuthorsCreator([]map[string]interface{}{map[string]interface{}(123)}) // AuthorsCreator | List of authors

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionsAPI.PrivateCollectionAuthorsAdd(context.Background(), collectionId).Authors(authors).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.PrivateCollectionAuthorsAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivateCollectionAuthorsAdd`: Location
	fmt.Fprintf(os.Stdout, "Response from `CollectionsAPI.PrivateCollectionAuthorsAdd`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | Collection unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateCollectionAuthorsAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authors** | [**AuthorsCreator**](AuthorsCreator.md) | List of authors | 

### Return type

[**Location**](Location.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivateCollectionAuthorsList

> []Author PrivateCollectionAuthorsList(ctx, collectionId).Execute()

List collection authors



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	collectionId := int64(789) // int64 | Collection unique identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionsAPI.PrivateCollectionAuthorsList(context.Background(), collectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.PrivateCollectionAuthorsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivateCollectionAuthorsList`: []Author
	fmt.Fprintf(os.Stdout, "Response from `CollectionsAPI.PrivateCollectionAuthorsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | Collection unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateCollectionAuthorsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Author**](Author.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivateCollectionAuthorsReplace

> PrivateCollectionAuthorsReplace(ctx, collectionId).Authors(authors).Execute()

Replace collection authors



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	collectionId := int64(789) // int64 | Collection unique identifier
	authors := *openapiclient.NewAuthorsCreator([]map[string]interface{}{map[string]interface{}(123)}) // AuthorsCreator | List of authors

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CollectionsAPI.PrivateCollectionAuthorsReplace(context.Background(), collectionId).Authors(authors).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.PrivateCollectionAuthorsReplace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | Collection unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateCollectionAuthorsReplaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authors** | [**AuthorsCreator**](AuthorsCreator.md) | List of authors | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivateCollectionCategoriesAdd

> Location PrivateCollectionCategoriesAdd(ctx, collectionId).Categories(categories).Execute()

Add collection categories



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	collectionId := int64(789) // int64 | Collection unique identifier
	categories := *openapiclient.NewCategoriesCreator([]int64{int64(123)}) // CategoriesCreator | Categories list

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionsAPI.PrivateCollectionCategoriesAdd(context.Background(), collectionId).Categories(categories).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.PrivateCollectionCategoriesAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivateCollectionCategoriesAdd`: Location
	fmt.Fprintf(os.Stdout, "Response from `CollectionsAPI.PrivateCollectionCategoriesAdd`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | Collection unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateCollectionCategoriesAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **categories** | [**CategoriesCreator**](CategoriesCreator.md) | Categories list | 

### Return type

[**Location**](Location.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivateCollectionCategoriesList

> []Category PrivateCollectionCategoriesList(ctx, collectionId).Execute()

List collection categories



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	collectionId := int64(789) // int64 | Collection unique identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionsAPI.PrivateCollectionCategoriesList(context.Background(), collectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.PrivateCollectionCategoriesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivateCollectionCategoriesList`: []Category
	fmt.Fprintf(os.Stdout, "Response from `CollectionsAPI.PrivateCollectionCategoriesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | Collection unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateCollectionCategoriesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Category**](Category.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivateCollectionCategoriesReplace

> PrivateCollectionCategoriesReplace(ctx, collectionId).Categories(categories).Execute()

Replace collection categories



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	collectionId := int64(789) // int64 | Collection unique identifier
	categories := *openapiclient.NewCategoriesCreator([]int64{int64(123)}) // CategoriesCreator | Categories list

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CollectionsAPI.PrivateCollectionCategoriesReplace(context.Background(), collectionId).Categories(categories).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.PrivateCollectionCategoriesReplace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | Collection unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateCollectionCategoriesReplaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **categories** | [**CategoriesCreator**](CategoriesCreator.md) | Categories list | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivateCollectionCategoryDelete

> PrivateCollectionCategoryDelete(ctx, collectionId, categoryId).Execute()

Delete collection category



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	collectionId := int64(789) // int64 | Collection unique identifier
	categoryId := int64(789) // int64 | Collection category unique identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CollectionsAPI.PrivateCollectionCategoryDelete(context.Background(), collectionId, categoryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.PrivateCollectionCategoryDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | Collection unique identifier | 
**categoryId** | **int64** | Collection category unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateCollectionCategoryDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivateCollectionCreate

> LocationWarnings PrivateCollectionCreate(ctx).Collection(collection).Execute()

Create collection



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	collection := *openapiclient.NewCollectionCreate("Test collection title") // CollectionCreate | Collection description

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionsAPI.PrivateCollectionCreate(context.Background()).Collection(collection).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.PrivateCollectionCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivateCollectionCreate`: LocationWarnings
	fmt.Fprintf(os.Stdout, "Response from `CollectionsAPI.PrivateCollectionCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrivateCollectionCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **collection** | [**CollectionCreate**](CollectionCreate.md) | Collection description | 

### Return type

[**LocationWarnings**](LocationWarnings.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivateCollectionDelete

> PrivateCollectionDelete(ctx, collectionId).Execute()

Delete collection



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	collectionId := int64(789) // int64 | Collection Unique identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CollectionsAPI.PrivateCollectionDelete(context.Background(), collectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.PrivateCollectionDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | Collection Unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateCollectionDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivateCollectionDetails

> CollectionCompletePrivate PrivateCollectionDetails(ctx, collectionId).Execute()

Collection details



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	collectionId := int64(789) // int64 | Collection Unique identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionsAPI.PrivateCollectionDetails(context.Background(), collectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.PrivateCollectionDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivateCollectionDetails`: CollectionCompletePrivate
	fmt.Fprintf(os.Stdout, "Response from `CollectionsAPI.PrivateCollectionDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | Collection Unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateCollectionDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CollectionCompletePrivate**](CollectionCompletePrivate.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivateCollectionPatch

> LocationWarningsUpdate PrivateCollectionPatch(ctx, collectionId).Collection(collection).Execute()

Partially update collection



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	collectionId := int64(789) // int64 | Collection Unique identifier
	collection := *openapiclient.NewCollectionUpdate() // CollectionUpdate | Subset of collection fields to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionsAPI.PrivateCollectionPatch(context.Background(), collectionId).Collection(collection).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.PrivateCollectionPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivateCollectionPatch`: LocationWarningsUpdate
	fmt.Fprintf(os.Stdout, "Response from `CollectionsAPI.PrivateCollectionPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | Collection Unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateCollectionPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **collection** | [**CollectionUpdate**](CollectionUpdate.md) | Subset of collection fields to update | 

### Return type

[**LocationWarningsUpdate**](LocationWarningsUpdate.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivateCollectionPrivateLinkCreate

> PrivateLinkResponse PrivateCollectionPrivateLinkCreate(ctx, collectionId).PrivateLink(privateLink).Execute()

Create collection private link



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	collectionId := int64(789) // int64 | Collection unique identifier
	privateLink := *openapiclient.NewCollectionPrivateLinkCreator() // CollectionPrivateLinkCreator |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionsAPI.PrivateCollectionPrivateLinkCreate(context.Background(), collectionId).PrivateLink(privateLink).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.PrivateCollectionPrivateLinkCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivateCollectionPrivateLinkCreate`: PrivateLinkResponse
	fmt.Fprintf(os.Stdout, "Response from `CollectionsAPI.PrivateCollectionPrivateLinkCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | Collection unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateCollectionPrivateLinkCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **privateLink** | [**CollectionPrivateLinkCreator**](CollectionPrivateLinkCreator.md) |  | 

### Return type

[**PrivateLinkResponse**](PrivateLinkResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivateCollectionPrivateLinkDelete

> PrivateCollectionPrivateLinkDelete(ctx, collectionId, linkId).Execute()

Disable private link



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	collectionId := int64(789) // int64 | Collection unique identifier
	linkId := "linkId_example" // string | Private link token

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CollectionsAPI.PrivateCollectionPrivateLinkDelete(context.Background(), collectionId, linkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.PrivateCollectionPrivateLinkDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | Collection unique identifier | 
**linkId** | **string** | Private link token | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateCollectionPrivateLinkDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivateCollectionPrivateLinkDetails

> PrivateLink PrivateCollectionPrivateLinkDetails(ctx, collectionId, linkId).Execute()

View collection private link



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	collectionId := int64(789) // int64 | Collection unique identifier
	linkId := "linkId_example" // string | Private link token

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionsAPI.PrivateCollectionPrivateLinkDetails(context.Background(), collectionId, linkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.PrivateCollectionPrivateLinkDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivateCollectionPrivateLinkDetails`: PrivateLink
	fmt.Fprintf(os.Stdout, "Response from `CollectionsAPI.PrivateCollectionPrivateLinkDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | Collection unique identifier | 
**linkId** | **string** | Private link token | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateCollectionPrivateLinkDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PrivateLink**](PrivateLink.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivateCollectionPrivateLinkUpdate

> PrivateCollectionPrivateLinkUpdate(ctx, collectionId, linkId).PrivateLink(privateLink).Execute()

Update collection private link



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	collectionId := int64(789) // int64 | Collection unique identifier
	linkId := "linkId_example" // string | Private link token
	privateLink := *openapiclient.NewCollectionPrivateLinkCreator() // CollectionPrivateLinkCreator |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CollectionsAPI.PrivateCollectionPrivateLinkUpdate(context.Background(), collectionId, linkId).PrivateLink(privateLink).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.PrivateCollectionPrivateLinkUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | Collection unique identifier | 
**linkId** | **string** | Private link token | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateCollectionPrivateLinkUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **privateLink** | [**CollectionPrivateLinkCreator**](CollectionPrivateLinkCreator.md) |  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivateCollectionPrivateLinksList

> []PrivateLink PrivateCollectionPrivateLinksList(ctx, collectionId).Execute()

List collection private links



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	collectionId := int64(789) // int64 | Collection unique identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionsAPI.PrivateCollectionPrivateLinksList(context.Background(), collectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.PrivateCollectionPrivateLinksList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivateCollectionPrivateLinksList`: []PrivateLink
	fmt.Fprintf(os.Stdout, "Response from `CollectionsAPI.PrivateCollectionPrivateLinksList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | Collection unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateCollectionPrivateLinksListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]PrivateLink**](PrivateLink.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivateCollectionPublish

> Location PrivateCollectionPublish(ctx, collectionId).Execute()

Private Collection Publish



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	collectionId := int64(789) // int64 | Collection Unique identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionsAPI.PrivateCollectionPublish(context.Background(), collectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.PrivateCollectionPublish``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivateCollectionPublish`: Location
	fmt.Fprintf(os.Stdout, "Response from `CollectionsAPI.PrivateCollectionPublish`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | Collection Unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateCollectionPublishRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Location**](Location.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivateCollectionReserveDoi

> CollectionDOI PrivateCollectionReserveDoi(ctx, collectionId).Execute()

Private Collection Reserve DOI



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	collectionId := int64(789) // int64 | Collection Unique identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionsAPI.PrivateCollectionReserveDoi(context.Background(), collectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.PrivateCollectionReserveDoi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivateCollectionReserveDoi`: CollectionDOI
	fmt.Fprintf(os.Stdout, "Response from `CollectionsAPI.PrivateCollectionReserveDoi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | Collection Unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateCollectionReserveDoiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CollectionDOI**](CollectionDOI.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivateCollectionReserveHandle

> CollectionHandle PrivateCollectionReserveHandle(ctx, collectionId).Execute()

Private Collection Reserve Handle



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	collectionId := int64(789) // int64 | Collection Unique identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionsAPI.PrivateCollectionReserveHandle(context.Background(), collectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.PrivateCollectionReserveHandle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivateCollectionReserveHandle`: CollectionHandle
	fmt.Fprintf(os.Stdout, "Response from `CollectionsAPI.PrivateCollectionReserveHandle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | Collection Unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateCollectionReserveHandleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CollectionHandle**](CollectionHandle.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivateCollectionResource

> Location PrivateCollectionResource(ctx, collectionId).Resource(resource).Execute()

Private Collection Resource



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	collectionId := int64(789) // int64 | Collection unique identifier
	resource := *openapiclient.NewResource() // Resource | Resource data

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionsAPI.PrivateCollectionResource(context.Background(), collectionId).Resource(resource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.PrivateCollectionResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivateCollectionResource`: Location
	fmt.Fprintf(os.Stdout, "Response from `CollectionsAPI.PrivateCollectionResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | Collection unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateCollectionResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resource** | [**Resource**](Resource.md) | Resource data | 

### Return type

[**Location**](Location.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivateCollectionUpdate

> LocationWarningsUpdate PrivateCollectionUpdate(ctx, collectionId).Collection(collection).Execute()

Update collection



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	collectionId := int64(789) // int64 | Collection Unique identifier
	collection := *openapiclient.NewCollectionUpdate() // CollectionUpdate | Collection description

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionsAPI.PrivateCollectionUpdate(context.Background(), collectionId).Collection(collection).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.PrivateCollectionUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivateCollectionUpdate`: LocationWarningsUpdate
	fmt.Fprintf(os.Stdout, "Response from `CollectionsAPI.PrivateCollectionUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | Collection Unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateCollectionUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **collection** | [**CollectionUpdate**](CollectionUpdate.md) | Collection description | 

### Return type

[**LocationWarningsUpdate**](LocationWarningsUpdate.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivateCollectionsList

> []Collection PrivateCollectionsList(ctx).Page(page).PageSize(pageSize).Limit(limit).Offset(offset).Order(order).OrderDirection(orderDirection).Execute()

Private Collections List



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	page := int64(789) // int64 | Page number. Used for pagination with page_size (optional)
	pageSize := int64(789) // int64 | The number of results included on a page. Used for pagination with page (optional) (default to 10)
	limit := int64(789) // int64 | Number of results included on a page. Used for pagination with query (optional)
	offset := int64(789) // int64 | Where to start the listing (the offset of the first result). Used for pagination with limit (optional)
	order := "order_example" // string | The field by which to order. Default varies by endpoint/resource. (optional) (default to "published_date")
	orderDirection := "orderDirection_example" // string |  (optional) (default to "desc")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionsAPI.PrivateCollectionsList(context.Background()).Page(page).PageSize(pageSize).Limit(limit).Offset(offset).Order(order).OrderDirection(orderDirection).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.PrivateCollectionsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivateCollectionsList`: []Collection
	fmt.Fprintf(os.Stdout, "Response from `CollectionsAPI.PrivateCollectionsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrivateCollectionsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** | Page number. Used for pagination with page_size | 
 **pageSize** | **int64** | The number of results included on a page. Used for pagination with page | [default to 10]
 **limit** | **int64** | Number of results included on a page. Used for pagination with query | 
 **offset** | **int64** | Where to start the listing (the offset of the first result). Used for pagination with limit | 
 **order** | **string** | The field by which to order. Default varies by endpoint/resource. | [default to &quot;published_date&quot;]
 **orderDirection** | **string** |  | [default to &quot;desc&quot;]

### Return type

[**[]Collection**](Collection.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivateCollectionsSearch

> []Collection PrivateCollectionsSearch(ctx).Search(search).Execute()

Private Collections Search



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	search := *openapiclient.NewPrivateCollectionSearch() // PrivateCollectionSearch | Search Parameters

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionsAPI.PrivateCollectionsSearch(context.Background()).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.PrivateCollectionsSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivateCollectionsSearch`: []Collection
	fmt.Fprintf(os.Stdout, "Response from `CollectionsAPI.PrivateCollectionsSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrivateCollectionsSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | [**PrivateCollectionSearch**](PrivateCollectionSearch.md) | Search Parameters | 

### Return type

[**[]Collection**](Collection.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

