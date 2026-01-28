# \InstitutionsAPI

All URIs are relative to *https://api.figshare.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountInstitutionCuration**](InstitutionsAPI.md#AccountInstitutionCuration) | **Get** /account/institution/review/{curation_id} | Institution Curation Review
[**AccountInstitutionCurations**](InstitutionsAPI.md#AccountInstitutionCurations) | **Get** /account/institution/reviews | Institution Curation Reviews
[**CustomFieldsList**](InstitutionsAPI.md#CustomFieldsList) | **Get** /account/institution/custom_fields | Private account institution group custom fields
[**CustomFieldsUpload**](InstitutionsAPI.md#CustomFieldsUpload) | **Post** /account/institution/custom_fields/{custom_field_id}/items/upload | Custom fields values files upload
[**GetAccountInstitutionCurationComments**](InstitutionsAPI.md#GetAccountInstitutionCurationComments) | **Get** /account/institution/review/{curation_id}/comments | Institution Curation Review Comments
[**InstitutionArticles**](InstitutionsAPI.md#InstitutionArticles) | **Get** /institutions/{institution_string_id}/articles/filter-by | Public Institution Articles
[**InstitutionHrfeedUpload**](InstitutionsAPI.md#InstitutionHrfeedUpload) | **Post** /institution/hrfeed/upload | Private Institution HRfeed Upload
[**PostAccountInstitutionCurationComments**](InstitutionsAPI.md#PostAccountInstitutionCurationComments) | **Post** /account/institution/review/{curation_id}/comments | POST Institution Curation Review Comment
[**PrivateAccountInstitutionUser**](InstitutionsAPI.md#PrivateAccountInstitutionUser) | **Get** /account/institution/users/{account_id} | Private Account Institution User
[**PrivateCategoriesList**](InstitutionsAPI.md#PrivateCategoriesList) | **Get** /account/categories | Private Account Categories
[**PrivateGroupEmbargoOptionsDetails**](InstitutionsAPI.md#PrivateGroupEmbargoOptionsDetails) | **Get** /account/institution/groups/{group_id}/embargo_options | Private Account Institution Group Embargo Options
[**PrivateInstitutionAccount**](InstitutionsAPI.md#PrivateInstitutionAccount) | **Get** /account/institution/accounts/{account_id} | Private Institution Account information
[**PrivateInstitutionAccountGroupRoleDelete**](InstitutionsAPI.md#PrivateInstitutionAccountGroupRoleDelete) | **Delete** /account/institution/roles/{account_id}/{group_id}/{role_id} | Delete Institution Account Group Role
[**PrivateInstitutionAccountGroupRoles**](InstitutionsAPI.md#PrivateInstitutionAccountGroupRoles) | **Get** /account/institution/roles/{account_id} | List Institution Account Group Roles
[**PrivateInstitutionAccountGroupRolesCreate**](InstitutionsAPI.md#PrivateInstitutionAccountGroupRolesCreate) | **Post** /account/institution/roles/{account_id} | Add Institution Account Group Roles
[**PrivateInstitutionAccountsCreate**](InstitutionsAPI.md#PrivateInstitutionAccountsCreate) | **Post** /account/institution/accounts | Create new Institution Account
[**PrivateInstitutionAccountsList**](InstitutionsAPI.md#PrivateInstitutionAccountsList) | **Get** /account/institution/accounts | Private Account Institution Accounts
[**PrivateInstitutionAccountsSearch**](InstitutionsAPI.md#PrivateInstitutionAccountsSearch) | **Post** /account/institution/accounts/search | Private Account Institution Accounts Search
[**PrivateInstitutionAccountsUpdate**](InstitutionsAPI.md#PrivateInstitutionAccountsUpdate) | **Put** /account/institution/accounts/{account_id} | Update Institution Account
[**PrivateInstitutionArticles**](InstitutionsAPI.md#PrivateInstitutionArticles) | **Get** /account/institution/articles | Private Institution Articles
[**PrivateInstitutionDetails**](InstitutionsAPI.md#PrivateInstitutionDetails) | **Get** /account/institution | Private Account Institutions
[**PrivateInstitutionEmbargoOptionsDetails**](InstitutionsAPI.md#PrivateInstitutionEmbargoOptionsDetails) | **Get** /account/institution/embargo_options | Private Account Institution embargo options
[**PrivateInstitutionGroupsList**](InstitutionsAPI.md#PrivateInstitutionGroupsList) | **Get** /account/institution/groups | Private Account Institution Groups
[**PrivateInstitutionRolesList**](InstitutionsAPI.md#PrivateInstitutionRolesList) | **Get** /account/institution/roles | Private Account Institution Roles



## AccountInstitutionCuration

> CurationDetail AccountInstitutionCuration(ctx, curationId).Execute()

Institution Curation Review



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
	curationId := int64(789) // int64 | ID of the curation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstitutionsAPI.AccountInstitutionCuration(context.Background(), curationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstitutionsAPI.AccountInstitutionCuration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccountInstitutionCuration`: CurationDetail
	fmt.Fprintf(os.Stdout, "Response from `InstitutionsAPI.AccountInstitutionCuration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**curationId** | **int64** | ID of the curation | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountInstitutionCurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CurationDetail**](CurationDetail.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountInstitutionCurations

> Curation AccountInstitutionCurations(ctx).GroupId(groupId).ArticleId(articleId).Status(status).Limit(limit).Offset(offset).Execute()

Institution Curation Reviews



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
	groupId := int64(789) // int64 | Filter by the group ID (optional)
	articleId := int64(789) // int64 | Retrieve the reviews for this article (optional)
	status := "status_example" // string | Filter by the status of the review (optional)
	limit := int64(789) // int64 | Number of results included on a page. Used for pagination with query (optional)
	offset := int64(789) // int64 | Where to start the listing (the offset of the first result). Used for pagination with limit (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstitutionsAPI.AccountInstitutionCurations(context.Background()).GroupId(groupId).ArticleId(articleId).Status(status).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstitutionsAPI.AccountInstitutionCurations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccountInstitutionCurations`: Curation
	fmt.Fprintf(os.Stdout, "Response from `InstitutionsAPI.AccountInstitutionCurations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccountInstitutionCurationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **int64** | Filter by the group ID | 
 **articleId** | **int64** | Retrieve the reviews for this article | 
 **status** | **string** | Filter by the status of the review | 
 **limit** | **int64** | Number of results included on a page. Used for pagination with query | 
 **offset** | **int64** | Where to start the listing (the offset of the first result). Used for pagination with limit | 

### Return type

[**Curation**](Curation.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomFieldsList

> []ShortCustomField CustomFieldsList(ctx).GroupId(groupId).Execute()

Private account institution group custom fields



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
	groupId := int64(789) // int64 | Group_id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstitutionsAPI.CustomFieldsList(context.Background()).GroupId(groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstitutionsAPI.CustomFieldsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomFieldsList`: []ShortCustomField
	fmt.Fprintf(os.Stdout, "Response from `InstitutionsAPI.CustomFieldsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomFieldsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **int64** | Group_id | 

### Return type

[**[]ShortCustomField**](ShortCustomField.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomFieldsUpload

> map[string]interface{} CustomFieldsUpload(ctx, customFieldId).ExternalFile(externalFile).Execute()

Custom fields values files upload



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
	customFieldId := int64(789) // int64 | Custom field identifier
	externalFile := os.NewFile(1234, "some_file") // *os.File | CSV file to be uploaded (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstitutionsAPI.CustomFieldsUpload(context.Background(), customFieldId).ExternalFile(externalFile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstitutionsAPI.CustomFieldsUpload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomFieldsUpload`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `InstitutionsAPI.CustomFieldsUpload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customFieldId** | **int64** | Custom field identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomFieldsUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **externalFile** | ***os.File** | CSV file to be uploaded | 

### Return type

**map[string]interface{}**

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountInstitutionCurationComments

> CurationComment GetAccountInstitutionCurationComments(ctx, curationId).Limit(limit).Offset(offset).Execute()

Institution Curation Review Comments



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
	curationId := int64(789) // int64 | ID of the curation
	limit := int64(789) // int64 | Number of results included on a page. Used for pagination with query (optional)
	offset := int64(789) // int64 | Where to start the listing (the offset of the first result). Used for pagination with limit (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstitutionsAPI.GetAccountInstitutionCurationComments(context.Background(), curationId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstitutionsAPI.GetAccountInstitutionCurationComments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountInstitutionCurationComments`: CurationComment
	fmt.Fprintf(os.Stdout, "Response from `InstitutionsAPI.GetAccountInstitutionCurationComments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**curationId** | **int64** | ID of the curation | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountInstitutionCurationCommentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int64** | Number of results included on a page. Used for pagination with query | 
 **offset** | **int64** | Where to start the listing (the offset of the first result). Used for pagination with limit | 

### Return type

[**CurationComment**](CurationComment.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstitutionArticles

> []Article InstitutionArticles(ctx, institutionStringId).ResourceId(resourceId).Filename(filename).Execute()

Public Institution Articles



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
	institutionStringId := "institutionStringId_example" // string | 
	resourceId := "resourceId_example" // string | 
	filename := "filename_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstitutionsAPI.InstitutionArticles(context.Background(), institutionStringId).ResourceId(resourceId).Filename(filename).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstitutionsAPI.InstitutionArticles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InstitutionArticles`: []Article
	fmt.Fprintf(os.Stdout, "Response from `InstitutionsAPI.InstitutionArticles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**institutionStringId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstitutionArticlesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resourceId** | **string** |  | 
 **filename** | **string** |  | 

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


## InstitutionHrfeedUpload

> ResponseMessage InstitutionHrfeedUpload(ctx).Hrfeed(hrfeed).Execute()

Private Institution HRfeed Upload



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
	hrfeed := os.NewFile(1234, "some_file") // *os.File | You can find an example in the Hr Feed section (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstitutionsAPI.InstitutionHrfeedUpload(context.Background()).Hrfeed(hrfeed).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstitutionsAPI.InstitutionHrfeedUpload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InstitutionHrfeedUpload`: ResponseMessage
	fmt.Fprintf(os.Stdout, "Response from `InstitutionsAPI.InstitutionHrfeedUpload`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInstitutionHrfeedUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hrfeed** | ***os.File** | You can find an example in the Hr Feed section | 

### Return type

[**ResponseMessage**](ResponseMessage.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAccountInstitutionCurationComments

> PostAccountInstitutionCurationComments(ctx, curationId).CurationComment(curationComment).Execute()

POST Institution Curation Review Comment



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
	curationId := int64(789) // int64 | ID of the curation
	curationComment := *openapiclient.NewCurationCommentCreate("Text_example") // CurationCommentCreate | The content/value of the comment.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InstitutionsAPI.PostAccountInstitutionCurationComments(context.Background(), curationId).CurationComment(curationComment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstitutionsAPI.PostAccountInstitutionCurationComments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**curationId** | **int64** | ID of the curation | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAccountInstitutionCurationCommentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **curationComment** | [**CurationCommentCreate**](CurationCommentCreate.md) | The content/value of the comment. | 

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


## PrivateAccountInstitutionUser

> User PrivateAccountInstitutionUser(ctx, accountId).Execute()

Private Account Institution User



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
	accountId := int64(789) // int64 | Account identifier the user is associated to

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstitutionsAPI.PrivateAccountInstitutionUser(context.Background(), accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstitutionsAPI.PrivateAccountInstitutionUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivateAccountInstitutionUser`: User
	fmt.Fprintf(os.Stdout, "Response from `InstitutionsAPI.PrivateAccountInstitutionUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int64** | Account identifier the user is associated to | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateAccountInstitutionUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**User**](User.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivateCategoriesList

> []CategoryList PrivateCategoriesList(ctx).Execute()

Private Account Categories



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
	resp, r, err := apiClient.InstitutionsAPI.PrivateCategoriesList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstitutionsAPI.PrivateCategoriesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivateCategoriesList`: []CategoryList
	fmt.Fprintf(os.Stdout, "Response from `InstitutionsAPI.PrivateCategoriesList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateCategoriesListRequest struct via the builder pattern


### Return type

[**[]CategoryList**](CategoryList.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivateGroupEmbargoOptionsDetails

> []GroupEmbargoOptions PrivateGroupEmbargoOptionsDetails(ctx, groupId).Execute()

Private Account Institution Group Embargo Options



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
	groupId := int64(789) // int64 | Group identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstitutionsAPI.PrivateGroupEmbargoOptionsDetails(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstitutionsAPI.PrivateGroupEmbargoOptionsDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivateGroupEmbargoOptionsDetails`: []GroupEmbargoOptions
	fmt.Fprintf(os.Stdout, "Response from `InstitutionsAPI.PrivateGroupEmbargoOptionsDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int64** | Group identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateGroupEmbargoOptionsDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GroupEmbargoOptions**](GroupEmbargoOptions.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivateInstitutionAccount

> Account PrivateInstitutionAccount(ctx, accountId).Execute()

Private Institution Account information



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
	accountId := int64(789) // int64 | Account identifier the user is associated to

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstitutionsAPI.PrivateInstitutionAccount(context.Background(), accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstitutionsAPI.PrivateInstitutionAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivateInstitutionAccount`: Account
	fmt.Fprintf(os.Stdout, "Response from `InstitutionsAPI.PrivateInstitutionAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int64** | Account identifier the user is associated to | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateInstitutionAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Account**](Account.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivateInstitutionAccountGroupRoleDelete

> PrivateInstitutionAccountGroupRoleDelete(ctx, accountId, groupId, roleId).Execute()

Delete Institution Account Group Role



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
	accountId := int64(789) // int64 | Account identifier for which to remove the role
	groupId := int64(789) // int64 | Group identifier for which to remove the role
	roleId := int64(789) // int64 | Role identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InstitutionsAPI.PrivateInstitutionAccountGroupRoleDelete(context.Background(), accountId, groupId, roleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstitutionsAPI.PrivateInstitutionAccountGroupRoleDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int64** | Account identifier for which to remove the role | 
**groupId** | **int64** | Group identifier for which to remove the role | 
**roleId** | **int64** | Role identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateInstitutionAccountGroupRoleDeleteRequest struct via the builder pattern


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


## PrivateInstitutionAccountGroupRoles

> map[string]interface{} PrivateInstitutionAccountGroupRoles(ctx, accountId).Execute()

List Institution Account Group Roles



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
	accountId := int64(789) // int64 | Account identifier the user is associated to

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstitutionsAPI.PrivateInstitutionAccountGroupRoles(context.Background(), accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstitutionsAPI.PrivateInstitutionAccountGroupRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivateInstitutionAccountGroupRoles`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `InstitutionsAPI.PrivateInstitutionAccountGroupRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int64** | Account identifier the user is associated to | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateInstitutionAccountGroupRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivateInstitutionAccountGroupRolesCreate

> PrivateInstitutionAccountGroupRolesCreate(ctx, accountId).Account(account).Execute()

Add Institution Account Group Roles



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
	accountId := int64(789) // int64 | Account identifier the user is associated to
	account := map[string]interface{}{ ... } // map[string]interface{} | Account description

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InstitutionsAPI.PrivateInstitutionAccountGroupRolesCreate(context.Background(), accountId).Account(account).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstitutionsAPI.PrivateInstitutionAccountGroupRolesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int64** | Account identifier the user is associated to | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateInstitutionAccountGroupRolesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **account** | **map[string]interface{}** | Account description | 

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


## PrivateInstitutionAccountsCreate

> AccountCreateResponse PrivateInstitutionAccountsCreate(ctx).Account(account).Execute()

Create new Institution Account



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
	account := *openapiclient.NewAccountCreate("johndoe@example.com", "Doe") // AccountCreate | Account description

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstitutionsAPI.PrivateInstitutionAccountsCreate(context.Background()).Account(account).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstitutionsAPI.PrivateInstitutionAccountsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivateInstitutionAccountsCreate`: AccountCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `InstitutionsAPI.PrivateInstitutionAccountsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrivateInstitutionAccountsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | [**AccountCreate**](AccountCreate.md) | Account description | 

### Return type

[**AccountCreateResponse**](AccountCreateResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivateInstitutionAccountsList

> []ShortAccount PrivateInstitutionAccountsList(ctx).Page(page).PageSize(pageSize).Limit(limit).Offset(offset).IsActive(isActive).InstitutionUserId(institutionUserId).Email(email).IdLte(idLte).IdGte(idGte).Execute()

Private Account Institution Accounts



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
	isActive := int64(789) // int64 | Filter by active status (optional)
	institutionUserId := "institutionUserId_example" // string | Filter by institution_user_id (optional)
	email := "email_example" // string | Filter by email (optional)
	idLte := int64(789) // int64 | Retrieve accounts with an ID lower or equal to the specified value (optional)
	idGte := int64(789) // int64 | Retrieve accounts with an ID greater or equal to the specified value (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstitutionsAPI.PrivateInstitutionAccountsList(context.Background()).Page(page).PageSize(pageSize).Limit(limit).Offset(offset).IsActive(isActive).InstitutionUserId(institutionUserId).Email(email).IdLte(idLte).IdGte(idGte).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstitutionsAPI.PrivateInstitutionAccountsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivateInstitutionAccountsList`: []ShortAccount
	fmt.Fprintf(os.Stdout, "Response from `InstitutionsAPI.PrivateInstitutionAccountsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrivateInstitutionAccountsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** | Page number. Used for pagination with page_size | 
 **pageSize** | **int64** | The number of results included on a page. Used for pagination with page | [default to 10]
 **limit** | **int64** | Number of results included on a page. Used for pagination with query | 
 **offset** | **int64** | Where to start the listing (the offset of the first result). Used for pagination with limit | 
 **isActive** | **int64** | Filter by active status | 
 **institutionUserId** | **string** | Filter by institution_user_id | 
 **email** | **string** | Filter by email | 
 **idLte** | **int64** | Retrieve accounts with an ID lower or equal to the specified value | 
 **idGte** | **int64** | Retrieve accounts with an ID greater or equal to the specified value | 

### Return type

[**[]ShortAccount**](ShortAccount.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivateInstitutionAccountsSearch

> []ShortAccount PrivateInstitutionAccountsSearch(ctx).Search(search).Execute()

Private Account Institution Accounts Search



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
	search := *openapiclient.NewInstitutionAccountsSearch() // InstitutionAccountsSearch | Search Parameters

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstitutionsAPI.PrivateInstitutionAccountsSearch(context.Background()).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstitutionsAPI.PrivateInstitutionAccountsSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivateInstitutionAccountsSearch`: []ShortAccount
	fmt.Fprintf(os.Stdout, "Response from `InstitutionsAPI.PrivateInstitutionAccountsSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrivateInstitutionAccountsSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | [**InstitutionAccountsSearch**](InstitutionAccountsSearch.md) | Search Parameters | 

### Return type

[**[]ShortAccount**](ShortAccount.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivateInstitutionAccountsUpdate

> PrivateInstitutionAccountsUpdate(ctx, accountId).Account(account).Execute()

Update Institution Account



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
	accountId := int64(789) // int64 | Account identifier the user is associated to
	account := *openapiclient.NewAccountUpdate(int64(123), false) // AccountUpdate | Account description

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InstitutionsAPI.PrivateInstitutionAccountsUpdate(context.Background(), accountId).Account(account).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstitutionsAPI.PrivateInstitutionAccountsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int64** | Account identifier the user is associated to | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateInstitutionAccountsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **account** | [**AccountUpdate**](AccountUpdate.md) | Account description | 

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


## PrivateInstitutionArticles

> []Article PrivateInstitutionArticles(ctx).Page(page).PageSize(pageSize).Limit(limit).Offset(offset).Order(order).OrderDirection(orderDirection).PublishedSince(publishedSince).ModifiedSince(modifiedSince).Status(status).ResourceDoi(resourceDoi).ItemType(itemType).Group(group).Execute()

Private Institution Articles



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
	publishedSince := "publishedSince_example" // string | Filter by article publishing date. Will only return articles published after the date. date(ISO 8601) YYYY-MM-DD or date-time(ISO 8601) YYYY-MM-DDTHH:mm:ssZ (optional)
	modifiedSince := "modifiedSince_example" // string | Filter by article modified date. Will only return articles published after the date. date(ISO 8601) YYYY-MM-DD or date-time(ISO 8601) YYYY-MM-DDTHH:mm:ssZ (optional)
	status := int64(789) // int64 | only return collections with this status (optional)
	resourceDoi := "resourceDoi_example" // string | only return collections with this resource_doi (optional)
	itemType := int64(789) // int64 | Only return articles with the respective type. Mapping for item_type is: 1 - Figure, 2 - Media, 3 - Dataset, 5 - Poster, 6 - Journal contribution, 7 - Presentation, 8 - Thesis, 9 - Software, 11 - Online resource, 12 - Preprint, 13 - Book, 14 - Conference contribution, 15 - Chapter, 16 - Peer review, 17 - Educational resource, 18 - Report, 19 - Standard, 20 - Composition, 21 - Funding, 22 - Physical object, 23 - Data management plan, 24 - Workflow, 25 - Monograph, 26 - Performance, 27 - Event, 28 - Service, 29 - Model (optional)
	group := int64(789) // int64 | only return articles from this group (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstitutionsAPI.PrivateInstitutionArticles(context.Background()).Page(page).PageSize(pageSize).Limit(limit).Offset(offset).Order(order).OrderDirection(orderDirection).PublishedSince(publishedSince).ModifiedSince(modifiedSince).Status(status).ResourceDoi(resourceDoi).ItemType(itemType).Group(group).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstitutionsAPI.PrivateInstitutionArticles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivateInstitutionArticles`: []Article
	fmt.Fprintf(os.Stdout, "Response from `InstitutionsAPI.PrivateInstitutionArticles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrivateInstitutionArticlesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** | Page number. Used for pagination with page_size | 
 **pageSize** | **int64** | The number of results included on a page. Used for pagination with page | [default to 10]
 **limit** | **int64** | Number of results included on a page. Used for pagination with query | 
 **offset** | **int64** | Where to start the listing (the offset of the first result). Used for pagination with limit | 
 **order** | **string** | The field by which to order. Default varies by endpoint/resource. | [default to &quot;published_date&quot;]
 **orderDirection** | **string** |  | [default to &quot;desc&quot;]
 **publishedSince** | **string** | Filter by article publishing date. Will only return articles published after the date. date(ISO 8601) YYYY-MM-DD or date-time(ISO 8601) YYYY-MM-DDTHH:mm:ssZ | 
 **modifiedSince** | **string** | Filter by article modified date. Will only return articles published after the date. date(ISO 8601) YYYY-MM-DD or date-time(ISO 8601) YYYY-MM-DDTHH:mm:ssZ | 
 **status** | **int64** | only return collections with this status | 
 **resourceDoi** | **string** | only return collections with this resource_doi | 
 **itemType** | **int64** | Only return articles with the respective type. Mapping for item_type is: 1 - Figure, 2 - Media, 3 - Dataset, 5 - Poster, 6 - Journal contribution, 7 - Presentation, 8 - Thesis, 9 - Software, 11 - Online resource, 12 - Preprint, 13 - Book, 14 - Conference contribution, 15 - Chapter, 16 - Peer review, 17 - Educational resource, 18 - Report, 19 - Standard, 20 - Composition, 21 - Funding, 22 - Physical object, 23 - Data management plan, 24 - Workflow, 25 - Monograph, 26 - Performance, 27 - Event, 28 - Service, 29 - Model | 
 **group** | **int64** | only return articles from this group | 

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


## PrivateInstitutionDetails

> Institution PrivateInstitutionDetails(ctx).Execute()

Private Account Institutions



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
	resp, r, err := apiClient.InstitutionsAPI.PrivateInstitutionDetails(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstitutionsAPI.PrivateInstitutionDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivateInstitutionDetails`: Institution
	fmt.Fprintf(os.Stdout, "Response from `InstitutionsAPI.PrivateInstitutionDetails`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateInstitutionDetailsRequest struct via the builder pattern


### Return type

[**Institution**](Institution.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivateInstitutionEmbargoOptionsDetails

> []GroupEmbargoOptions PrivateInstitutionEmbargoOptionsDetails(ctx).Execute()

Private Account Institution embargo options



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
	resp, r, err := apiClient.InstitutionsAPI.PrivateInstitutionEmbargoOptionsDetails(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstitutionsAPI.PrivateInstitutionEmbargoOptionsDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivateInstitutionEmbargoOptionsDetails`: []GroupEmbargoOptions
	fmt.Fprintf(os.Stdout, "Response from `InstitutionsAPI.PrivateInstitutionEmbargoOptionsDetails`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateInstitutionEmbargoOptionsDetailsRequest struct via the builder pattern


### Return type

[**[]GroupEmbargoOptions**](GroupEmbargoOptions.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivateInstitutionGroupsList

> []Group PrivateInstitutionGroupsList(ctx).Execute()

Private Account Institution Groups



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
	resp, r, err := apiClient.InstitutionsAPI.PrivateInstitutionGroupsList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstitutionsAPI.PrivateInstitutionGroupsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivateInstitutionGroupsList`: []Group
	fmt.Fprintf(os.Stdout, "Response from `InstitutionsAPI.PrivateInstitutionGroupsList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateInstitutionGroupsListRequest struct via the builder pattern


### Return type

[**[]Group**](Group.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivateInstitutionRolesList

> []Role PrivateInstitutionRolesList(ctx).Execute()

Private Account Institution Roles



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
	resp, r, err := apiClient.InstitutionsAPI.PrivateInstitutionRolesList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstitutionsAPI.PrivateInstitutionRolesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivateInstitutionRolesList`: []Role
	fmt.Fprintf(os.Stdout, "Response from `InstitutionsAPI.PrivateInstitutionRolesList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateInstitutionRolesListRequest struct via the builder pattern


### Return type

[**[]Role**](Role.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

