# \ArticlesAPI

All URIs are relative to *https://api.figshare.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountArticlePublish**](ArticlesAPI.md#AccountArticlePublish) | **Post** /account/articles/{article_id}/publish | Private Article Publish
[**AccountArticleReport**](ArticlesAPI.md#AccountArticleReport) | **Get** /account/articles/export | Account Article Report
[**AccountArticleReportGenerate**](ArticlesAPI.md#AccountArticleReportGenerate) | **Post** /account/articles/export | Initiate a new Report
[**AccountArticleUnpublish**](ArticlesAPI.md#AccountArticleUnpublish) | **Post** /account/articles/{article_id}/unpublish | Public Article Unpublish
[**ArticleDetails**](ArticlesAPI.md#ArticleDetails) | **Get** /articles/{article_id} | View article details
[**ArticleFileDetails**](ArticlesAPI.md#ArticleFileDetails) | **Get** /articles/{article_id}/files/{file_id} | Article file details
[**ArticleFiles**](ArticlesAPI.md#ArticleFiles) | **Get** /articles/{article_id}/files | List article files
[**ArticleVersionConfidentiality**](ArticlesAPI.md#ArticleVersionConfidentiality) | **Get** /articles/{article_id}/versions/{version_id}/confidentiality | Public Article Confidentiality for article version
[**ArticleVersionDetails**](ArticlesAPI.md#ArticleVersionDetails) | **Get** /articles/{article_id}/versions/{version_id} | Article details for version
[**ArticleVersionEmbargo**](ArticlesAPI.md#ArticleVersionEmbargo) | **Get** /articles/{article_id}/versions/{version_id}/embargo | Public Article Embargo for article version
[**ArticleVersionFiles**](ArticlesAPI.md#ArticleVersionFiles) | **Get** /articles/{article_id}/versions/{version_id}/files | Article version file details
[**ArticleVersionPartialUpdate**](ArticlesAPI.md#ArticleVersionPartialUpdate) | **Patch** /account/articles/{article_id}/versions/{version_id} | Partially update article version
[**ArticleVersionUpdate**](ArticlesAPI.md#ArticleVersionUpdate) | **Put** /account/articles/{article_id}/versions/{version_id} | Update article version
[**ArticleVersionUpdateThumb**](ArticlesAPI.md#ArticleVersionUpdateThumb) | **Put** /account/articles/{article_id}/versions/{version_id}/update_thumb | Update article version thumbnail
[**ArticleVersions**](ArticlesAPI.md#ArticleVersions) | **Get** /articles/{article_id}/versions | List article versions
[**ArticlesList**](ArticlesAPI.md#ArticlesList) | **Get** /articles | Public Articles
[**ArticlesSearch**](ArticlesAPI.md#ArticlesSearch) | **Post** /articles/search | Public Articles Search
[**PrivateArticleAuthorDelete**](ArticlesAPI.md#PrivateArticleAuthorDelete) | **Delete** /account/articles/{article_id}/authors/{author_id} | Delete article author
[**PrivateArticleAuthorsAdd**](ArticlesAPI.md#PrivateArticleAuthorsAdd) | **Post** /account/articles/{article_id}/authors | Add article authors
[**PrivateArticleAuthorsList**](ArticlesAPI.md#PrivateArticleAuthorsList) | **Get** /account/articles/{article_id}/authors | List article authors
[**PrivateArticleAuthorsReplace**](ArticlesAPI.md#PrivateArticleAuthorsReplace) | **Put** /account/articles/{article_id}/authors | Replace article authors
[**PrivateArticleCategoriesAdd**](ArticlesAPI.md#PrivateArticleCategoriesAdd) | **Post** /account/articles/{article_id}/categories | Add article categories
[**PrivateArticleCategoriesList**](ArticlesAPI.md#PrivateArticleCategoriesList) | **Get** /account/articles/{article_id}/categories | List article categories
[**PrivateArticleCategoriesReplace**](ArticlesAPI.md#PrivateArticleCategoriesReplace) | **Put** /account/articles/{article_id}/categories | Replace article categories
[**PrivateArticleCategoryDelete**](ArticlesAPI.md#PrivateArticleCategoryDelete) | **Delete** /account/articles/{article_id}/categories/{category_id} | Delete article category
[**PrivateArticleConfidentialityDelete**](ArticlesAPI.md#PrivateArticleConfidentialityDelete) | **Delete** /account/articles/{article_id}/confidentiality | Delete article confidentiality
[**PrivateArticleConfidentialityDetails**](ArticlesAPI.md#PrivateArticleConfidentialityDetails) | **Get** /account/articles/{article_id}/confidentiality | Article confidentiality details
[**PrivateArticleConfidentialityUpdate**](ArticlesAPI.md#PrivateArticleConfidentialityUpdate) | **Put** /account/articles/{article_id}/confidentiality | Update article confidentiality
[**PrivateArticleCreate**](ArticlesAPI.md#PrivateArticleCreate) | **Post** /account/articles | Create new Article
[**PrivateArticleDelete**](ArticlesAPI.md#PrivateArticleDelete) | **Delete** /account/articles/{article_id} | Delete article
[**PrivateArticleDetails**](ArticlesAPI.md#PrivateArticleDetails) | **Get** /account/articles/{article_id} | Article details
[**PrivateArticleDownload**](ArticlesAPI.md#PrivateArticleDownload) | **Get** /account/articles/{article_id}/download | Private Article Download
[**PrivateArticleEmbargoDelete**](ArticlesAPI.md#PrivateArticleEmbargoDelete) | **Delete** /account/articles/{article_id}/embargo | Delete Article Embargo
[**PrivateArticleEmbargoDetails**](ArticlesAPI.md#PrivateArticleEmbargoDetails) | **Get** /account/articles/{article_id}/embargo | Article Embargo Details
[**PrivateArticleEmbargoUpdate**](ArticlesAPI.md#PrivateArticleEmbargoUpdate) | **Put** /account/articles/{article_id}/embargo | Update Article Embargo
[**PrivateArticleFile**](ArticlesAPI.md#PrivateArticleFile) | **Get** /account/articles/{article_id}/files/{file_id} | Single File
[**PrivateArticleFileDelete**](ArticlesAPI.md#PrivateArticleFileDelete) | **Delete** /account/articles/{article_id}/files/{file_id} | File Delete
[**PrivateArticleFilesList**](ArticlesAPI.md#PrivateArticleFilesList) | **Get** /account/articles/{article_id}/files | List article files
[**PrivateArticlePartialUpdate**](ArticlesAPI.md#PrivateArticlePartialUpdate) | **Patch** /account/articles/{article_id} | Partially update article
[**PrivateArticlePrivateLink**](ArticlesAPI.md#PrivateArticlePrivateLink) | **Get** /account/articles/{article_id}/private_links | List private links
[**PrivateArticlePrivateLinkCreate**](ArticlesAPI.md#PrivateArticlePrivateLinkCreate) | **Post** /account/articles/{article_id}/private_links | Create private link
[**PrivateArticlePrivateLinkDelete**](ArticlesAPI.md#PrivateArticlePrivateLinkDelete) | **Delete** /account/articles/{article_id}/private_links/{link_id} | Disable private link
[**PrivateArticlePrivateLinkUpdate**](ArticlesAPI.md#PrivateArticlePrivateLinkUpdate) | **Put** /account/articles/{article_id}/private_links/{link_id} | Update private link
[**PrivateArticleReserveDoi**](ArticlesAPI.md#PrivateArticleReserveDoi) | **Post** /account/articles/{article_id}/reserve_doi | Private Article Reserve DOI
[**PrivateArticleReserveHandle**](ArticlesAPI.md#PrivateArticleReserveHandle) | **Post** /account/articles/{article_id}/reserve_handle | Private Article Reserve Handle
[**PrivateArticleResource**](ArticlesAPI.md#PrivateArticleResource) | **Post** /account/articles/{article_id}/resource | Private Article Resource
[**PrivateArticleUpdate**](ArticlesAPI.md#PrivateArticleUpdate) | **Put** /account/articles/{article_id} | Update article
[**PrivateArticleUploadComplete**](ArticlesAPI.md#PrivateArticleUploadComplete) | **Post** /account/articles/{article_id}/files/{file_id} | Complete Upload
[**PrivateArticleUploadInitiate**](ArticlesAPI.md#PrivateArticleUploadInitiate) | **Post** /account/articles/{article_id}/files | Initiate Upload
[**PrivateArticlesList**](ArticlesAPI.md#PrivateArticlesList) | **Get** /account/articles | Private Articles
[**PrivateArticlesSearch**](ArticlesAPI.md#PrivateArticlesSearch) | **Post** /account/articles/search | Private Articles search
[**PublicArticleDownload**](ArticlesAPI.md#PublicArticleDownload) | **Get** /articles/{article_id}/download | Public Article Download
[**PublicArticleVersionDownload**](ArticlesAPI.md#PublicArticleVersionDownload) | **Get** /articles/{article_id}/versions/{version_id}/download | Public Article Version Download



## AccountArticlePublish

> Location AccountArticlePublish(ctx, articleId).Execute()

Private Article Publish



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
	articleId := int64(789) // int64 | Article unique identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticlesAPI.AccountArticlePublish(context.Background(), articleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticlesAPI.AccountArticlePublish``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccountArticlePublish`: Location
	fmt.Fprintf(os.Stdout, "Response from `ArticlesAPI.AccountArticlePublish`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**articleId** | **int64** | Article unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountArticlePublishRequest struct via the builder pattern


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


## AccountArticleReport

> []AccountReport AccountArticleReport(ctx).GroupId(groupId).Execute()

Account Article Report



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
	groupId := int64(789) // int64 | A group ID to filter by (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticlesAPI.AccountArticleReport(context.Background()).GroupId(groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticlesAPI.AccountArticleReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccountArticleReport`: []AccountReport
	fmt.Fprintf(os.Stdout, "Response from `ArticlesAPI.AccountArticleReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccountArticleReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **int64** | A group ID to filter by | 

### Return type

[**[]AccountReport**](AccountReport.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountArticleReportGenerate

> AccountReport AccountArticleReportGenerate(ctx).Execute()

Initiate a new Report



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticlesAPI.AccountArticleReportGenerate(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticlesAPI.AccountArticleReportGenerate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccountArticleReportGenerate`: AccountReport
	fmt.Fprintf(os.Stdout, "Response from `ArticlesAPI.AccountArticleReportGenerate`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAccountArticleReportGenerateRequest struct via the builder pattern


### Return type

[**AccountReport**](AccountReport.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountArticleUnpublish

> AccountArticleUnpublish(ctx, articleId).Reason(reason).Execute()

Public Article Unpublish



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
	articleId := int64(789) // int64 | Article unique identifier
	reason := *openapiclient.NewArticleUnpublishData("Unpublish article reason") // ArticleUnpublishData | Article unpublish data

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ArticlesAPI.AccountArticleUnpublish(context.Background(), articleId).Reason(reason).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticlesAPI.AccountArticleUnpublish``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**articleId** | **int64** | Article unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountArticleUnpublishRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reason** | [**ArticleUnpublishData**](ArticleUnpublishData.md) | Article unpublish data | 

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


## ArticleFileDetails

> PublicFile ArticleFileDetails(ctx, articleId, fileId).Execute()

Article file details



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
	fileId := int64(789) // int64 | File Unique identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticlesAPI.ArticleFileDetails(context.Background(), articleId, fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticlesAPI.ArticleFileDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ArticleFileDetails`: PublicFile
	fmt.Fprintf(os.Stdout, "Response from `ArticlesAPI.ArticleFileDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**articleId** | **int64** | Article Unique identifier | 
**fileId** | **int64** | File Unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiArticleFileDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PublicFile**](PublicFile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ArticleFiles

> []PublicFile ArticleFiles(ctx, articleId).Page(page).PageSize(pageSize).Limit(limit).Offset(offset).Execute()

List article files



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
	page := int64(789) // int64 | Page number. Used for pagination with page_size (optional)
	pageSize := int64(789) // int64 | The number of results included on a page. Used for pagination with page (optional) (default to 10)
	limit := int64(789) // int64 | Number of results included on a page. Used for pagination with query (optional)
	offset := int64(789) // int64 | Where to start the listing (the offset of the first result). Used for pagination with limit (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticlesAPI.ArticleFiles(context.Background(), articleId).Page(page).PageSize(pageSize).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticlesAPI.ArticleFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ArticleFiles`: []PublicFile
	fmt.Fprintf(os.Stdout, "Response from `ArticlesAPI.ArticleFiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**articleId** | **int64** | Article Unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiArticleFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int64** | Page number. Used for pagination with page_size | 
 **pageSize** | **int64** | The number of results included on a page. Used for pagination with page | [default to 10]
 **limit** | **int64** | Number of results included on a page. Used for pagination with query | 
 **offset** | **int64** | Where to start the listing (the offset of the first result). Used for pagination with limit | 

### Return type

[**[]PublicFile**](PublicFile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ArticleVersionConfidentiality

> ArticleConfidentiality ArticleVersionConfidentiality(ctx, articleId, versionId).Execute()

Public Article Confidentiality for article version



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
	versionId := int64(789) // int64 | Version Number

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticlesAPI.ArticleVersionConfidentiality(context.Background(), articleId, versionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticlesAPI.ArticleVersionConfidentiality``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ArticleVersionConfidentiality`: ArticleConfidentiality
	fmt.Fprintf(os.Stdout, "Response from `ArticlesAPI.ArticleVersionConfidentiality`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**articleId** | **int64** | Article Unique identifier | 
**versionId** | **int64** | Version Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiArticleVersionConfidentialityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ArticleConfidentiality**](ArticleConfidentiality.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ArticleVersionDetails

> ArticleComplete ArticleVersionDetails(ctx, articleId, versionId).Execute()

Article details for version



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
	versionId := int64(789) // int64 | Article Version Number

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticlesAPI.ArticleVersionDetails(context.Background(), articleId, versionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticlesAPI.ArticleVersionDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ArticleVersionDetails`: ArticleComplete
	fmt.Fprintf(os.Stdout, "Response from `ArticlesAPI.ArticleVersionDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**articleId** | **int64** | Article Unique identifier | 
**versionId** | **int64** | Article Version Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiArticleVersionDetailsRequest struct via the builder pattern


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


## ArticleVersionEmbargo

> ArticleEmbargo ArticleVersionEmbargo(ctx, articleId, versionId).Execute()

Public Article Embargo for article version



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
	versionId := int64(789) // int64 | Version Number

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticlesAPI.ArticleVersionEmbargo(context.Background(), articleId, versionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticlesAPI.ArticleVersionEmbargo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ArticleVersionEmbargo`: ArticleEmbargo
	fmt.Fprintf(os.Stdout, "Response from `ArticlesAPI.ArticleVersionEmbargo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**articleId** | **int64** | Article Unique identifier | 
**versionId** | **int64** | Version Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiArticleVersionEmbargoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ArticleEmbargo**](ArticleEmbargo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ArticleVersionFiles

> []PublicFile ArticleVersionFiles(ctx, articleId, versionId).Page(page).PageSize(pageSize).Limit(limit).Offset(offset).Execute()

Article version file details



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
	versionId := int64(789) // int64 | Article Version Unique identifier
	page := int64(789) // int64 | Page number. Used for pagination with page_size (optional)
	pageSize := int64(789) // int64 | The number of results included on a page. Used for pagination with page (optional) (default to 10)
	limit := int64(789) // int64 | Number of results included on a page. Used for pagination with query (optional)
	offset := int64(789) // int64 | Where to start the listing (the offset of the first result). Used for pagination with limit (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticlesAPI.ArticleVersionFiles(context.Background(), articleId, versionId).Page(page).PageSize(pageSize).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticlesAPI.ArticleVersionFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ArticleVersionFiles`: []PublicFile
	fmt.Fprintf(os.Stdout, "Response from `ArticlesAPI.ArticleVersionFiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**articleId** | **int64** | Article Unique identifier | 
**versionId** | **int64** | Article Version Unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiArticleVersionFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int64** | Page number. Used for pagination with page_size | 
 **pageSize** | **int64** | The number of results included on a page. Used for pagination with page | [default to 10]
 **limit** | **int64** | Number of results included on a page. Used for pagination with query | 
 **offset** | **int64** | Where to start the listing (the offset of the first result). Used for pagination with limit | 

### Return type

[**[]PublicFile**](PublicFile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ArticleVersionPartialUpdate

> LocationWarningsUpdate ArticleVersionPartialUpdate(ctx, articleId, versionId).Article(article).Execute()

Partially update article version



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
	articleId := int64(789) // int64 | Article unique identifier
	versionId := int64(789) // int64 | Article version identifier
	article := *openapiclient.NewArticleVersionUpdate() // ArticleVersionUpdate | Subset of article version fields to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticlesAPI.ArticleVersionPartialUpdate(context.Background(), articleId, versionId).Article(article).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticlesAPI.ArticleVersionPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ArticleVersionPartialUpdate`: LocationWarningsUpdate
	fmt.Fprintf(os.Stdout, "Response from `ArticlesAPI.ArticleVersionPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**articleId** | **int64** | Article unique identifier | 
**versionId** | **int64** | Article version identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiArticleVersionPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **article** | [**ArticleVersionUpdate**](ArticleVersionUpdate.md) | Subset of article version fields to update | 

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


## ArticleVersionUpdate

> LocationWarningsUpdate ArticleVersionUpdate(ctx, articleId, versionId).Article(article).Execute()

Update article version



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
	articleId := int64(789) // int64 | Article unique identifier
	versionId := int64(789) // int64 | Article version identifier
	article := *openapiclient.NewArticleVersionUpdate() // ArticleVersionUpdate | Article description

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticlesAPI.ArticleVersionUpdate(context.Background(), articleId, versionId).Article(article).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticlesAPI.ArticleVersionUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ArticleVersionUpdate`: LocationWarningsUpdate
	fmt.Fprintf(os.Stdout, "Response from `ArticlesAPI.ArticleVersionUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**articleId** | **int64** | Article unique identifier | 
**versionId** | **int64** | Article version identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiArticleVersionUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **article** | [**ArticleVersionUpdate**](ArticleVersionUpdate.md) | Article description | 

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


## ArticleVersionUpdateThumb

> ArticleVersionUpdateThumb(ctx, articleId, versionId).FileId(fileId).Execute()

Update article version thumbnail



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
	articleId := int64(789) // int64 | Article unique identifier
	versionId := int64(789) // int64 | Article version identifier
	fileId := *openapiclient.NewFileId() // FileId | File ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ArticlesAPI.ArticleVersionUpdateThumb(context.Background(), articleId, versionId).FileId(fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticlesAPI.ArticleVersionUpdateThumb``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**articleId** | **int64** | Article unique identifier | 
**versionId** | **int64** | Article version identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiArticleVersionUpdateThumbRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fileId** | [**FileId**](FileId.md) | File ID | 

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


## ArticleVersions

> []ArticleVersions ArticleVersions(ctx, articleId).Execute()

List article versions



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
	resp, r, err := apiClient.ArticlesAPI.ArticleVersions(context.Background(), articleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticlesAPI.ArticleVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ArticleVersions`: []ArticleVersions
	fmt.Fprintf(os.Stdout, "Response from `ArticlesAPI.ArticleVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**articleId** | **int64** | Article Unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiArticleVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ArticleVersions**](ArticleVersions.md)

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


## ArticlesSearch

> []ArticleWithProject ArticlesSearch(ctx).XCursor(xCursor).Search(search).Execute()

Public Articles Search



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
	search := *openapiclient.NewArticleSearch() // ArticleSearch | Search Parameters (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticlesAPI.ArticlesSearch(context.Background()).XCursor(xCursor).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticlesAPI.ArticlesSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ArticlesSearch`: []ArticleWithProject
	fmt.Fprintf(os.Stdout, "Response from `ArticlesAPI.ArticlesSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArticlesSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xCursor** | **string** | Unique hash used for bypassing the item retrieval limit of 9,000 entities. When using this parameter, please note that the offset parameter will not be available, but the limit parameter will still work as expected. | 
 **search** | [**ArticleSearch**](ArticleSearch.md) | Search Parameters | 

### Return type

[**[]ArticleWithProject**](ArticleWithProject.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivateArticleAuthorDelete

> PrivateArticleAuthorDelete(ctx, articleId, authorId).Execute()

Delete article author



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
	articleId := int64(789) // int64 | Article unique identifier
	authorId := int64(789) // int64 | Article Author unique identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ArticlesAPI.PrivateArticleAuthorDelete(context.Background(), articleId, authorId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticlesAPI.PrivateArticleAuthorDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**articleId** | **int64** | Article unique identifier | 
**authorId** | **int64** | Article Author unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateArticleAuthorDeleteRequest struct via the builder pattern


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


## PrivateArticleAuthorsAdd

> PrivateArticleAuthorsAdd(ctx, articleId).Authors(authors).Execute()

Add article authors



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
	articleId := int64(789) // int64 | Article unique identifier
	authors := *openapiclient.NewAuthorsCreator([]map[string]interface{}{map[string]interface{}(123)}) // AuthorsCreator | Authors description

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ArticlesAPI.PrivateArticleAuthorsAdd(context.Background(), articleId).Authors(authors).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticlesAPI.PrivateArticleAuthorsAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**articleId** | **int64** | Article unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateArticleAuthorsAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authors** | [**AuthorsCreator**](AuthorsCreator.md) | Authors description | 

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


## PrivateArticleAuthorsList

> []Author PrivateArticleAuthorsList(ctx, articleId).Execute()

List article authors



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
	articleId := int64(789) // int64 | Article unique identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticlesAPI.PrivateArticleAuthorsList(context.Background(), articleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticlesAPI.PrivateArticleAuthorsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivateArticleAuthorsList`: []Author
	fmt.Fprintf(os.Stdout, "Response from `ArticlesAPI.PrivateArticleAuthorsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**articleId** | **int64** | Article unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateArticleAuthorsListRequest struct via the builder pattern


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


## PrivateArticleAuthorsReplace

> PrivateArticleAuthorsReplace(ctx, articleId).Authors(authors).Execute()

Replace article authors



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
	articleId := int64(789) // int64 | Article unique identifier
	authors := *openapiclient.NewAuthorsCreator([]map[string]interface{}{map[string]interface{}(123)}) // AuthorsCreator | Authors description

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ArticlesAPI.PrivateArticleAuthorsReplace(context.Background(), articleId).Authors(authors).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticlesAPI.PrivateArticleAuthorsReplace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**articleId** | **int64** | Article unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateArticleAuthorsReplaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authors** | [**AuthorsCreator**](AuthorsCreator.md) | Authors description | 

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


## PrivateArticleCategoriesAdd

> PrivateArticleCategoriesAdd(ctx, articleId).Categories(categories).Execute()

Add article categories



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
	articleId := int64(789) // int64 | Article unique identifier
	categories := *openapiclient.NewCategoriesCreator([]int64{int64(123)}) // CategoriesCreator | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ArticlesAPI.PrivateArticleCategoriesAdd(context.Background(), articleId).Categories(categories).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticlesAPI.PrivateArticleCategoriesAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**articleId** | **int64** | Article unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateArticleCategoriesAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **categories** | [**CategoriesCreator**](CategoriesCreator.md) |  | 

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


## PrivateArticleCategoriesList

> []Category PrivateArticleCategoriesList(ctx, articleId).Execute()

List article categories



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
	articleId := int64(789) // int64 | Article unique identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticlesAPI.PrivateArticleCategoriesList(context.Background(), articleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticlesAPI.PrivateArticleCategoriesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivateArticleCategoriesList`: []Category
	fmt.Fprintf(os.Stdout, "Response from `ArticlesAPI.PrivateArticleCategoriesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**articleId** | **int64** | Article unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateArticleCategoriesListRequest struct via the builder pattern


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


## PrivateArticleCategoriesReplace

> PrivateArticleCategoriesReplace(ctx, articleId).Categories(categories).Execute()

Replace article categories



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
	articleId := int64(789) // int64 | Article unique identifier
	categories := *openapiclient.NewCategoriesCreator([]int64{int64(123)}) // CategoriesCreator | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ArticlesAPI.PrivateArticleCategoriesReplace(context.Background(), articleId).Categories(categories).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticlesAPI.PrivateArticleCategoriesReplace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**articleId** | **int64** | Article unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateArticleCategoriesReplaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **categories** | [**CategoriesCreator**](CategoriesCreator.md) |  | 

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


## PrivateArticleCategoryDelete

> PrivateArticleCategoryDelete(ctx, articleId, categoryId).Execute()

Delete article category



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
	articleId := int64(789) // int64 | Article unique identifier
	categoryId := int64(789) // int64 | Category unique identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ArticlesAPI.PrivateArticleCategoryDelete(context.Background(), articleId, categoryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticlesAPI.PrivateArticleCategoryDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**articleId** | **int64** | Article unique identifier | 
**categoryId** | **int64** | Category unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateArticleCategoryDeleteRequest struct via the builder pattern


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


## PrivateArticleConfidentialityDelete

> PrivateArticleConfidentialityDelete(ctx, articleId).Execute()

Delete article confidentiality



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
	articleId := int64(789) // int64 | Article unique identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ArticlesAPI.PrivateArticleConfidentialityDelete(context.Background(), articleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticlesAPI.PrivateArticleConfidentialityDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**articleId** | **int64** | Article unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateArticleConfidentialityDeleteRequest struct via the builder pattern


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


## PrivateArticleConfidentialityDetails

> ArticleConfidentiality PrivateArticleConfidentialityDetails(ctx, articleId).Execute()

Article confidentiality details



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
	articleId := int64(789) // int64 | Article unique identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticlesAPI.PrivateArticleConfidentialityDetails(context.Background(), articleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticlesAPI.PrivateArticleConfidentialityDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivateArticleConfidentialityDetails`: ArticleConfidentiality
	fmt.Fprintf(os.Stdout, "Response from `ArticlesAPI.PrivateArticleConfidentialityDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**articleId** | **int64** | Article unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateArticleConfidentialityDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ArticleConfidentiality**](ArticleConfidentiality.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivateArticleConfidentialityUpdate

> PrivateArticleConfidentialityUpdate(ctx, articleId).Reason(reason).Execute()

Update article confidentiality



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
	articleId := int64(789) // int64 | Article unique identifier
	reason := *openapiclient.NewConfidentialityCreator("Reason_example") // ConfidentialityCreator | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ArticlesAPI.PrivateArticleConfidentialityUpdate(context.Background(), articleId).Reason(reason).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticlesAPI.PrivateArticleConfidentialityUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**articleId** | **int64** | Article unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateArticleConfidentialityUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reason** | [**ConfidentialityCreator**](ConfidentialityCreator.md) |  | 

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


## PrivateArticleCreate

> LocationWarnings PrivateArticleCreate(ctx).Article(article).Execute()

Create new Article



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
	article := *openapiclient.NewArticleCreate("Test article title") // ArticleCreate | Article description

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticlesAPI.PrivateArticleCreate(context.Background()).Article(article).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticlesAPI.PrivateArticleCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivateArticleCreate`: LocationWarnings
	fmt.Fprintf(os.Stdout, "Response from `ArticlesAPI.PrivateArticleCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrivateArticleCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **article** | [**ArticleCreate**](ArticleCreate.md) | Article description | 

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


## PrivateArticleDelete

> PrivateArticleDelete(ctx, articleId).Execute()

Delete article



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
	articleId := int64(789) // int64 | Article unique identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ArticlesAPI.PrivateArticleDelete(context.Background(), articleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticlesAPI.PrivateArticleDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**articleId** | **int64** | Article unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateArticleDeleteRequest struct via the builder pattern


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


## PrivateArticleDetails

> ArticleCompletePrivate PrivateArticleDetails(ctx, articleId).Execute()

Article details



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
	articleId := int64(789) // int64 | Article unique identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticlesAPI.PrivateArticleDetails(context.Background(), articleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticlesAPI.PrivateArticleDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivateArticleDetails`: ArticleCompletePrivate
	fmt.Fprintf(os.Stdout, "Response from `ArticlesAPI.PrivateArticleDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**articleId** | **int64** | Article unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateArticleDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ArticleCompletePrivate**](ArticleCompletePrivate.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivateArticleDownload

> PrivateArticleDownload(ctx, articleId).FolderPath(folderPath).Execute()

Private Article Download



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
	articleId := int64(789) // int64 | Article unique identifier
	folderPath := "folderPath_example" // string | Folder path to download. If not provided, all files from the article will be downloaded (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ArticlesAPI.PrivateArticleDownload(context.Background(), articleId).FolderPath(folderPath).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticlesAPI.PrivateArticleDownload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**articleId** | **int64** | Article unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateArticleDownloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **folderPath** | **string** | Folder path to download. If not provided, all files from the article will be downloaded | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivateArticleEmbargoDelete

> PrivateArticleEmbargoDelete(ctx, articleId).Execute()

Delete Article Embargo



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
	articleId := int64(789) // int64 | Article unique identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ArticlesAPI.PrivateArticleEmbargoDelete(context.Background(), articleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticlesAPI.PrivateArticleEmbargoDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**articleId** | **int64** | Article unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateArticleEmbargoDeleteRequest struct via the builder pattern


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


## PrivateArticleEmbargoDetails

> ArticleEmbargo PrivateArticleEmbargoDetails(ctx, articleId).Execute()

Article Embargo Details



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
	articleId := int64(789) // int64 | Article unique identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticlesAPI.PrivateArticleEmbargoDetails(context.Background(), articleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticlesAPI.PrivateArticleEmbargoDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivateArticleEmbargoDetails`: ArticleEmbargo
	fmt.Fprintf(os.Stdout, "Response from `ArticlesAPI.PrivateArticleEmbargoDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**articleId** | **int64** | Article unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateArticleEmbargoDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ArticleEmbargo**](ArticleEmbargo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivateArticleEmbargoUpdate

> PrivateArticleEmbargoUpdate(ctx, articleId).Embargo(embargo).Execute()

Update Article Embargo



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
	articleId := int64(789) // int64 | Article unique identifier
	embargo := *openapiclient.NewArticleEmbargoUpdater("2018-05-22T04:04:04", "file", true) // ArticleEmbargoUpdater | Embargo description

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ArticlesAPI.PrivateArticleEmbargoUpdate(context.Background(), articleId).Embargo(embargo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticlesAPI.PrivateArticleEmbargoUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**articleId** | **int64** | Article unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateArticleEmbargoUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **embargo** | [**ArticleEmbargoUpdater**](ArticleEmbargoUpdater.md) | Embargo description | 

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


## PrivateArticleFile

> PrivateFile PrivateArticleFile(ctx, articleId, fileId).Execute()

Single File



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
	articleId := int64(789) // int64 | Article unique identifier
	fileId := int64(789) // int64 | File unique identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticlesAPI.PrivateArticleFile(context.Background(), articleId, fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticlesAPI.PrivateArticleFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivateArticleFile`: PrivateFile
	fmt.Fprintf(os.Stdout, "Response from `ArticlesAPI.PrivateArticleFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**articleId** | **int64** | Article unique identifier | 
**fileId** | **int64** | File unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateArticleFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PrivateFile**](PrivateFile.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivateArticleFileDelete

> PrivateArticleFileDelete(ctx, articleId, fileId).Execute()

File Delete



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
	articleId := int64(789) // int64 | Article unique identifier
	fileId := int64(789) // int64 | File unique identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ArticlesAPI.PrivateArticleFileDelete(context.Background(), articleId, fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticlesAPI.PrivateArticleFileDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**articleId** | **int64** | Article unique identifier | 
**fileId** | **int64** | File unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateArticleFileDeleteRequest struct via the builder pattern


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


## PrivateArticleFilesList

> []PrivateFile PrivateArticleFilesList(ctx, articleId).Page(page).PageSize(pageSize).Limit(limit).Offset(offset).Execute()

List article files



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
	articleId := int64(789) // int64 | Article unique identifier
	page := int64(789) // int64 | Page number. Used for pagination with page_size (optional)
	pageSize := int64(789) // int64 | The number of results included on a page. Used for pagination with page (optional) (default to 10)
	limit := int64(789) // int64 | Number of results included on a page. Used for pagination with query (optional)
	offset := int64(789) // int64 | Where to start the listing (the offset of the first result). Used for pagination with limit (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticlesAPI.PrivateArticleFilesList(context.Background(), articleId).Page(page).PageSize(pageSize).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticlesAPI.PrivateArticleFilesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivateArticleFilesList`: []PrivateFile
	fmt.Fprintf(os.Stdout, "Response from `ArticlesAPI.PrivateArticleFilesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**articleId** | **int64** | Article unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateArticleFilesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int64** | Page number. Used for pagination with page_size | 
 **pageSize** | **int64** | The number of results included on a page. Used for pagination with page | [default to 10]
 **limit** | **int64** | Number of results included on a page. Used for pagination with query | 
 **offset** | **int64** | Where to start the listing (the offset of the first result). Used for pagination with limit | 

### Return type

[**[]PrivateFile**](PrivateFile.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivateArticlePartialUpdate

> LocationWarningsUpdate PrivateArticlePartialUpdate(ctx, articleId).Article(article).Execute()

Partially update article



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
	articleId := int64(789) // int64 | Article unique identifier
	article := *openapiclient.NewArticleUpdate() // ArticleUpdate | Subset of article fields to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticlesAPI.PrivateArticlePartialUpdate(context.Background(), articleId).Article(article).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticlesAPI.PrivateArticlePartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivateArticlePartialUpdate`: LocationWarningsUpdate
	fmt.Fprintf(os.Stdout, "Response from `ArticlesAPI.PrivateArticlePartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**articleId** | **int64** | Article unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateArticlePartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **article** | [**ArticleUpdate**](ArticleUpdate.md) | Subset of article fields to update | 

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


## PrivateArticlePrivateLink

> []PrivateLink PrivateArticlePrivateLink(ctx, articleId).Execute()

List private links



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
	articleId := int64(789) // int64 | Article unique identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticlesAPI.PrivateArticlePrivateLink(context.Background(), articleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticlesAPI.PrivateArticlePrivateLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivateArticlePrivateLink`: []PrivateLink
	fmt.Fprintf(os.Stdout, "Response from `ArticlesAPI.PrivateArticlePrivateLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**articleId** | **int64** | Article unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateArticlePrivateLinkRequest struct via the builder pattern


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


## PrivateArticlePrivateLinkCreate

> PrivateLinkResponse PrivateArticlePrivateLinkCreate(ctx, articleId).PrivateLink(privateLink).Execute()

Create private link



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
	articleId := int64(789) // int64 | Article unique identifier
	privateLink := *openapiclient.NewPrivateLinkCreator() // PrivateLinkCreator |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticlesAPI.PrivateArticlePrivateLinkCreate(context.Background(), articleId).PrivateLink(privateLink).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticlesAPI.PrivateArticlePrivateLinkCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivateArticlePrivateLinkCreate`: PrivateLinkResponse
	fmt.Fprintf(os.Stdout, "Response from `ArticlesAPI.PrivateArticlePrivateLinkCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**articleId** | **int64** | Article unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateArticlePrivateLinkCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **privateLink** | [**PrivateLinkCreator**](PrivateLinkCreator.md) |  | 

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


## PrivateArticlePrivateLinkDelete

> PrivateArticlePrivateLinkDelete(ctx, articleId, linkId).Execute()

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
	articleId := int64(789) // int64 | Article unique identifier
	linkId := "linkId_example" // string | Private link token

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ArticlesAPI.PrivateArticlePrivateLinkDelete(context.Background(), articleId, linkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticlesAPI.PrivateArticlePrivateLinkDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**articleId** | **int64** | Article unique identifier | 
**linkId** | **string** | Private link token | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateArticlePrivateLinkDeleteRequest struct via the builder pattern


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


## PrivateArticlePrivateLinkUpdate

> PrivateArticlePrivateLinkUpdate(ctx, articleId, linkId).PrivateLink(privateLink).Execute()

Update private link



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
	articleId := int64(789) // int64 | Article unique identifier
	linkId := "linkId_example" // string | Private link token
	privateLink := *openapiclient.NewPrivateLinkCreator() // PrivateLinkCreator |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ArticlesAPI.PrivateArticlePrivateLinkUpdate(context.Background(), articleId, linkId).PrivateLink(privateLink).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticlesAPI.PrivateArticlePrivateLinkUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**articleId** | **int64** | Article unique identifier | 
**linkId** | **string** | Private link token | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateArticlePrivateLinkUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **privateLink** | [**PrivateLinkCreator**](PrivateLinkCreator.md) |  | 

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


## PrivateArticleReserveDoi

> ArticleDOI PrivateArticleReserveDoi(ctx, articleId).Execute()

Private Article Reserve DOI



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
	articleId := int64(789) // int64 | Article unique identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticlesAPI.PrivateArticleReserveDoi(context.Background(), articleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticlesAPI.PrivateArticleReserveDoi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivateArticleReserveDoi`: ArticleDOI
	fmt.Fprintf(os.Stdout, "Response from `ArticlesAPI.PrivateArticleReserveDoi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**articleId** | **int64** | Article unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateArticleReserveDoiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ArticleDOI**](ArticleDOI.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivateArticleReserveHandle

> ArticleHandle PrivateArticleReserveHandle(ctx, articleId).Execute()

Private Article Reserve Handle



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
	articleId := int64(789) // int64 | Article unique identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticlesAPI.PrivateArticleReserveHandle(context.Background(), articleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticlesAPI.PrivateArticleReserveHandle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivateArticleReserveHandle`: ArticleHandle
	fmt.Fprintf(os.Stdout, "Response from `ArticlesAPI.PrivateArticleReserveHandle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**articleId** | **int64** | Article unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateArticleReserveHandleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ArticleHandle**](ArticleHandle.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivateArticleResource

> Location PrivateArticleResource(ctx, articleId).Resource(resource).Execute()

Private Article Resource



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
	articleId := int64(789) // int64 | Article unique identifier
	resource := *openapiclient.NewResource() // Resource | Resource data

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticlesAPI.PrivateArticleResource(context.Background(), articleId).Resource(resource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticlesAPI.PrivateArticleResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivateArticleResource`: Location
	fmt.Fprintf(os.Stdout, "Response from `ArticlesAPI.PrivateArticleResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**articleId** | **int64** | Article unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateArticleResourceRequest struct via the builder pattern


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


## PrivateArticleUpdate

> LocationWarningsUpdate PrivateArticleUpdate(ctx, articleId).Article(article).Execute()

Update article



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
	articleId := int64(789) // int64 | Article unique identifier
	article := *openapiclient.NewArticleUpdate() // ArticleUpdate | Full article representation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticlesAPI.PrivateArticleUpdate(context.Background(), articleId).Article(article).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticlesAPI.PrivateArticleUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivateArticleUpdate`: LocationWarningsUpdate
	fmt.Fprintf(os.Stdout, "Response from `ArticlesAPI.PrivateArticleUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**articleId** | **int64** | Article unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateArticleUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **article** | [**ArticleUpdate**](ArticleUpdate.md) | Full article representation | 

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


## PrivateArticleUploadComplete

> PrivateArticleUploadComplete(ctx, articleId, fileId).Execute()

Complete Upload



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
	articleId := int64(789) // int64 | Article unique identifier
	fileId := int64(789) // int64 | File unique identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ArticlesAPI.PrivateArticleUploadComplete(context.Background(), articleId, fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticlesAPI.PrivateArticleUploadComplete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**articleId** | **int64** | Article unique identifier | 
**fileId** | **int64** | File unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateArticleUploadCompleteRequest struct via the builder pattern


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


## PrivateArticleUploadInitiate

> Location PrivateArticleUploadInitiate(ctx, articleId).File(file).Page(page).PageSize(pageSize).Limit(limit).Offset(offset).Execute()

Initiate Upload



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
	articleId := int64(789) // int64 | Article unique identifier
	file := *openapiclient.NewFileCreator() // FileCreator | 
	page := int64(789) // int64 | Page number. Used for pagination with page_size (optional)
	pageSize := int64(789) // int64 | The number of results included on a page. Used for pagination with page (optional) (default to 10)
	limit := int64(789) // int64 | Number of results included on a page. Used for pagination with query (optional)
	offset := int64(789) // int64 | Where to start the listing (the offset of the first result). Used for pagination with limit (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticlesAPI.PrivateArticleUploadInitiate(context.Background(), articleId).File(file).Page(page).PageSize(pageSize).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticlesAPI.PrivateArticleUploadInitiate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivateArticleUploadInitiate`: Location
	fmt.Fprintf(os.Stdout, "Response from `ArticlesAPI.PrivateArticleUploadInitiate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**articleId** | **int64** | Article unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateArticleUploadInitiateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | [**FileCreator**](FileCreator.md) |  | 
 **page** | **int64** | Page number. Used for pagination with page_size | 
 **pageSize** | **int64** | The number of results included on a page. Used for pagination with page | [default to 10]
 **limit** | **int64** | Number of results included on a page. Used for pagination with query | 
 **offset** | **int64** | Where to start the listing (the offset of the first result). Used for pagination with limit | 

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


## PrivateArticlesList

> []Article PrivateArticlesList(ctx).Page(page).PageSize(pageSize).Limit(limit).Offset(offset).Execute()

Private Articles



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticlesAPI.PrivateArticlesList(context.Background()).Page(page).PageSize(pageSize).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticlesAPI.PrivateArticlesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivateArticlesList`: []Article
	fmt.Fprintf(os.Stdout, "Response from `ArticlesAPI.PrivateArticlesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrivateArticlesListRequest struct via the builder pattern


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


## PrivateArticlesSearch

> []ArticleWithProject PrivateArticlesSearch(ctx).Search(search).Execute()

Private Articles search



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
	search := *openapiclient.NewPrivateArticleSearch() // PrivateArticleSearch | Search Parameters

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticlesAPI.PrivateArticlesSearch(context.Background()).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticlesAPI.PrivateArticlesSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivateArticlesSearch`: []ArticleWithProject
	fmt.Fprintf(os.Stdout, "Response from `ArticlesAPI.PrivateArticlesSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrivateArticlesSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | [**PrivateArticleSearch**](PrivateArticleSearch.md) | Search Parameters | 

### Return type

[**[]ArticleWithProject**](ArticleWithProject.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PublicArticleDownload

> PublicArticleDownload(ctx, articleId).FolderPath(folderPath).Execute()

Public Article Download



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
	articleId := int64(789) // int64 | Article unique identifier
	folderPath := "folderPath_example" // string | Folder path to download. If not provided, all files from the article will be downloaded (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ArticlesAPI.PublicArticleDownload(context.Background(), articleId).FolderPath(folderPath).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticlesAPI.PublicArticleDownload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**articleId** | **int64** | Article unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublicArticleDownloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **folderPath** | **string** | Folder path to download. If not provided, all files from the article will be downloaded | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PublicArticleVersionDownload

> PublicArticleVersionDownload(ctx, articleId, versionId).FolderPath(folderPath).Execute()

Public Article Version Download



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
	articleId := int64(789) // int64 | Article unique identifier
	versionId := int64(789) // int64 | Version Number
	folderPath := "folderPath_example" // string | Folder path to download. If not provided, all files from the article will be downloaded (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ArticlesAPI.PublicArticleVersionDownload(context.Background(), articleId, versionId).FolderPath(folderPath).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticlesAPI.PublicArticleVersionDownload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**articleId** | **int64** | Article unique identifier | 
**versionId** | **int64** | Version Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublicArticleVersionDownloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **folderPath** | **string** | Folder path to download. If not provided, all files from the article will be downloaded | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

