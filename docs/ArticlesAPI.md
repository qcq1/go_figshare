# \ArticlesAPI

All URIs are relative to *https://api.figshare.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArticleDetails**](ArticlesAPI.md#ArticleDetails) | **Get** /articles/{article_id} | View article details
[**ArticlesList**](ArticlesAPI.md#ArticlesList) | **Get** /articles | Public Articles
[**GetAggregatedCitations**](ArticlesAPI.md#GetAggregatedCitations) | **Get** /viz/data/publication/for/aggregated-citations.json | Get article aggregated citations
[**Total**](ArticlesAPI.md#Total) | **Get** /total/article/{article_id} | View article total downloads, views
[**TotalIns**](ArticlesAPI.md#TotalIns) | **Get** /{institution}/total/article/{article_id} | View article total downloads, views



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

> []Article ArticlesList(ctx).XCursor(xCursor).Page(page).PageSize(pageSize).Limit(limit).Offset(offset).Order(order).OrderDirection(orderDirection).Institution(institution).PublishedSince(publishedSince).ModifiedSince(modifiedSince).Group(group).ResourceDoi(resourceDoi).ItemType(itemType).Doi(doi).Handle(handle).Execute()

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
	xCursor := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique hash used for bypassing the item retrieval limit of 9,000 entities. When using this parameter, please note that the offset parameter will not be available, but the limit parameter will still work as expected. (optional)
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
	resp, r, err := apiClient.ArticlesAPI.ArticlesList(context.Background()).XCursor(xCursor).Page(page).PageSize(pageSize).Limit(limit).Offset(offset).Order(order).OrderDirection(orderDirection).Institution(institution).PublishedSince(publishedSince).ModifiedSince(modifiedSince).Group(group).ResourceDoi(resourceDoi).ItemType(itemType).Doi(doi).Handle(handle).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticlesAPI.ArticlesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ArticlesList`: []Article
	fmt.Fprintf(os.Stdout, "Response from `ArticlesAPI.ArticlesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArticlesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xCursor** | **string** | Unique hash used for bypassing the item retrieval limit of 9,000 entities. When using this parameter, please note that the offset parameter will not be available, but the limit parameter will still work as expected. | 
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

[**[]Article**](Article.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAggregatedCitations

> AggregatedCitations GetAggregatedCitations(ctx).AndSubsetFigshareDoi(andSubsetFigshareDoi).Cookie(cookie).XCsrfToken(xCsrfToken).VizStaggr(vizStaggr).EntityOrder(entityOrder).Np(np).UserAgent(userAgent).Referer(referer).Execute()

Get article aggregated citations



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
	andSubsetFigshareDoi := "andSubsetFigshareDoi_example" // string | Article DOI
	cookie := "cookie_example" // string | 必须包含 session 和 uber_auth_tkt
	xCsrfToken := "xCsrfToken_example" // string | 
	vizStaggr := "vizStaggr_example" // string |  (optional)
	entityOrder := "entityOrder_example" // string |  (optional)
	np := "np_example" // string |  (optional)
	userAgent := "userAgent_example" // string |  (optional)
	referer := "referer_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticlesAPI.GetAggregatedCitations(context.Background()).AndSubsetFigshareDoi(andSubsetFigshareDoi).Cookie(cookie).XCsrfToken(xCsrfToken).VizStaggr(vizStaggr).EntityOrder(entityOrder).Np(np).UserAgent(userAgent).Referer(referer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticlesAPI.GetAggregatedCitations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAggregatedCitations`: AggregatedCitations
	fmt.Fprintf(os.Stdout, "Response from `ArticlesAPI.GetAggregatedCitations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAggregatedCitationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **andSubsetFigshareDoi** | **string** | Article DOI | 
 **cookie** | **string** | 必须包含 session 和 uber_auth_tkt | 
 **xCsrfToken** | **string** |  | 
 **vizStaggr** | **string** |  | 
 **entityOrder** | **string** |  | 
 **np** | **string** |  | 
 **userAgent** | **string** |  | 
 **referer** | **string** |  | 

### Return type

[**AggregatedCitations**](AggregatedCitations.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Total

> TotalArticle Total(ctx, articleId).Execute()

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
	resp, r, err := apiClient.ArticlesAPI.Total(context.Background(), articleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticlesAPI.Total``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Total`: TotalArticle
	fmt.Fprintf(os.Stdout, "Response from `ArticlesAPI.Total`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**articleId** | **int64** | Article Unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiTotalRequest struct via the builder pattern


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


## TotalIns

> TotalArticle TotalIns(ctx, institution, articleId).Execute()

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
	institution := "institution_example" // string | institution
	articleId := int64(789) // int64 | Article Unique identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticlesAPI.TotalIns(context.Background(), institution, articleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticlesAPI.TotalIns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TotalIns`: TotalArticle
	fmt.Fprintf(os.Stdout, "Response from `ArticlesAPI.TotalIns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**institution** | **string** | institution | 
**articleId** | **int64** | Article Unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiTotalInsRequest struct via the builder pattern


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

