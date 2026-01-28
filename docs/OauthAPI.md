# \OauthAPI

All URIs are relative to *https://api.figshare.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateToken**](OauthAPI.md#CreateToken) | **Post** /token | Create OAuth token
[**GetTokenInfo**](OauthAPI.md#GetTokenInfo) | **Get** /token | Get OAuth token information



## CreateToken

> OAuthToken CreateToken(ctx).Body(body).Execute()

Create OAuth token



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
	body := *openapiclient.NewCreateOAuthToken("ClientId_example", "ClientSecret_example", "GrantType_example") // CreateOAuthToken | Create OAuth Token Parameters (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OauthAPI.CreateToken(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OauthAPI.CreateToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateToken`: OAuthToken
	fmt.Fprintf(os.Stdout, "Response from `OauthAPI.CreateToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateOAuthToken**](CreateOAuthToken.md) | Create OAuth Token Parameters | 

### Return type

[**OAuthToken**](OAuthToken.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTokenInfo

> OAuthToken GetTokenInfo(ctx).AccessToken(accessToken).Execute()

Get OAuth token information



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
	accessToken := "accessToken_example" // string | OAuth access token (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OauthAPI.GetTokenInfo(context.Background()).AccessToken(accessToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OauthAPI.GetTokenInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTokenInfo`: OAuthToken
	fmt.Fprintf(os.Stdout, "Response from `OauthAPI.GetTokenInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTokenInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessToken** | **string** | OAuth access token | 

### Return type

[**OAuthToken**](OAuthToken.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

