# \AuthorsAPI

All URIs are relative to *https://api.figshare.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PrivateAuthorDetails**](AuthorsAPI.md#PrivateAuthorDetails) | **Get** /account/authors/{author_id} | Author details
[**PrivateAuthorsSearch**](AuthorsAPI.md#PrivateAuthorsSearch) | **Post** /account/authors/search | Search Authors



## PrivateAuthorDetails

> AuthorComplete PrivateAuthorDetails(ctx, authorId).Execute()

Author details



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
	authorId := int64(789) // int64 | Author unique identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthorsAPI.PrivateAuthorDetails(context.Background(), authorId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorsAPI.PrivateAuthorDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivateAuthorDetails`: AuthorComplete
	fmt.Fprintf(os.Stdout, "Response from `AuthorsAPI.PrivateAuthorDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorId** | **int64** | Author unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateAuthorDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuthorComplete**](AuthorComplete.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivateAuthorsSearch

> []AuthorComplete PrivateAuthorsSearch(ctx).Search(search).Execute()

Search Authors



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
	search := *openapiclient.NewPrivateAuthorsSearch() // PrivateAuthorsSearch | Search Parameters (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthorsAPI.PrivateAuthorsSearch(context.Background()).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorsAPI.PrivateAuthorsSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivateAuthorsSearch`: []AuthorComplete
	fmt.Fprintf(os.Stdout, "Response from `AuthorsAPI.PrivateAuthorsSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrivateAuthorsSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | [**PrivateAuthorsSearch**](PrivateAuthorsSearch.md) | Search Parameters | 

### Return type

[**[]AuthorComplete**](AuthorComplete.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

