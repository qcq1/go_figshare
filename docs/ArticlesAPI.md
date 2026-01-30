# \ArticlesAPI

All URIs are relative to *https://api.figshare.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArticleDetails**](ArticlesAPI.md#ArticleDetails) | **Get** /articles/{article_id} | View article details
[**ArticlesList**](ArticlesAPI.md#ArticlesList) | **Get** /articles | Public Articles
[**CrawlTotalArticle**](ArticlesAPI.md#CrawlTotalArticle) | **Get** /total/article/{article_id} | View article total downloads, views



## ArticleDetails

> ArticleComplete ArticleDetails(ctx, articleId).Execute()

View article details



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
	articleId := int64(789) // int64 | Article Unique identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticlesAPI.ArticleDetails(context.Background(), articleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticlesAPI.ArticleDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ArticleDetails`: ArticleComplete
	fmt.Fprintf(os.Stdout, "Response from `ArticlesAPI.ArticleDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**articleId** | **int64** | Article Unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiArticleDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ArticleComplete**](ArticleComplete.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ArticlesList

> Article ArticlesList(ctx).Page(page).PageSize(pageSize).Limit(limit).Offset(offset).Order(order).OrderDirection(orderDirection).Institution(institution).PublishedSince(publishedSince).ModifiedSince(modifiedSince).Group(group).ResourceDoi(resourceDoi).ItemType(itemType).Doi(doi).Handle(handle).Execute()

Public Articles



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
	institution := int64(789) // int64 | only return articles from this institution (optional)
	publishedSince := "publishedSince_example" // string | Filter by article publishing date. Will only return articles published after the date. date(ISO 8601) YYYY-MM-DD or date-time(ISO 8601) YYYY-MM-DDTHH:mm:ssZ (optional)
	modifiedSince := "modifiedSince_example" // string | Filter by article modified date. Will only return articles published after the date. date(ISO 8601) YYYY-MM-DD or date-time(ISO 8601) YYYY-MM-DDTHH:mm:ssZ (optional)
	group := int64(789) // int64 | only return articles from this group (optional)
	resourceDoi := "resourceDoi_example" // string | Deprecated by related materials. Only return articles with this resource_doi (optional)
	itemType := int64(789) // int64 | Only return articles with the respective type. Mapping for item_type is: 1 - Figure, 2 - Media, 3 - Dataset, 5 - Poster, 6 - Journal contribution, 7 - Presentation, 8 - Thesis, 9 - Software, 11 - Online resource, 12 - Preprint, 13 - Book, 14 - Conference contribution, 15 - Chapter, 16 - Peer review, 17 - Educational resource, 18 - Report, 19 - Standard, 20 - Composition, 21 - Funding, 22 - Physical object, 23 - Data management plan, 24 - Workflow, 25 - Monograph, 26 - Performance, 27 - Event, 28 - Service, 29 - Model (optional)
	doi := "doi_example" // string | only return articles with this doi (optional)
	handle := "handle_example" // string | only return articles with this handle (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticlesAPI.ArticlesList(context.Background()).Page(page).PageSize(pageSize).Limit(limit).Offset(offset).Order(order).OrderDirection(orderDirection).Institution(institution).PublishedSince(publishedSince).ModifiedSince(modifiedSince).Group(group).ResourceDoi(resourceDoi).ItemType(itemType).Doi(doi).Handle(handle).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticlesAPI.ArticlesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ArticlesList`: Article
	fmt.Fprintf(os.Stdout, "Response from `ArticlesAPI.ArticlesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArticlesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** | Page number. Used for pagination with page_size | 
 **pageSize** | **int64** | The number of results included on a page. Used for pagination with page | [default to 10]
 **limit** | **int64** | Number of results included on a page. Used for pagination with query | 
 **offset** | **int64** | Where to start the listing (the offset of the first result). Used for pagination with limit | 
 **order** | **string** | The field by which to order. Default varies by endpoint/resource. | [default to &quot;published_date&quot;]
 **orderDirection** | **string** |  | [default to &quot;desc&quot;]
 **institution** | **int64** | only return articles from this institution | 
 **publishedSince** | **string** | Filter by article publishing date. Will only return articles published after the date. date(ISO 8601) YYYY-MM-DD or date-time(ISO 8601) YYYY-MM-DDTHH:mm:ssZ | 
 **modifiedSince** | **string** | Filter by article modified date. Will only return articles published after the date. date(ISO 8601) YYYY-MM-DD or date-time(ISO 8601) YYYY-MM-DDTHH:mm:ssZ | 
 **group** | **int64** | only return articles from this group | 
 **resourceDoi** | **string** | Deprecated by related materials. Only return articles with this resource_doi | 
 **itemType** | **int64** | Only return articles with the respective type. Mapping for item_type is: 1 - Figure, 2 - Media, 3 - Dataset, 5 - Poster, 6 - Journal contribution, 7 - Presentation, 8 - Thesis, 9 - Software, 11 - Online resource, 12 - Preprint, 13 - Book, 14 - Conference contribution, 15 - Chapter, 16 - Peer review, 17 - Educational resource, 18 - Report, 19 - Standard, 20 - Composition, 21 - Funding, 22 - Physical object, 23 - Data management plan, 24 - Workflow, 25 - Monograph, 26 - Performance, 27 - Event, 28 - Service, 29 - Model | 
 **doi** | **string** | only return articles with this doi | 
 **handle** | **string** | only return articles with this handle | 

### Return type

[**Article**](Article.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CrawlTotalArticle

> TotalArticle CrawlTotalArticle(ctx, articleId).Execute()

View article total downloads, views



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
	articleId := int64(789) // int64 | Article Unique identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticlesAPI.CrawlTotalArticle(context.Background(), articleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticlesAPI.CrawlTotalArticle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CrawlTotalArticle`: TotalArticle
	fmt.Fprintf(os.Stdout, "Response from `ArticlesAPI.CrawlTotalArticle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**articleId** | **int64** | Article Unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCrawlTotalArticleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TotalArticle**](TotalArticle.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

