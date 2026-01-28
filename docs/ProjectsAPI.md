# \ProjectsAPI

All URIs are relative to *https://api.figshare.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PrivateProjectArticleDelete**](ProjectsAPI.md#PrivateProjectArticleDelete) | **Delete** /account/projects/{project_id}/articles/{article_id} | Delete project article
[**PrivateProjectArticleDetails**](ProjectsAPI.md#PrivateProjectArticleDetails) | **Get** /account/projects/{project_id}/articles/{article_id} | Project article details
[**PrivateProjectArticleFile**](ProjectsAPI.md#PrivateProjectArticleFile) | **Get** /account/projects/{project_id}/articles/{article_id}/files/{file_id} | Project article file details
[**PrivateProjectArticleFiles**](ProjectsAPI.md#PrivateProjectArticleFiles) | **Get** /account/projects/{project_id}/articles/{article_id}/files | Project article list files
[**PrivateProjectArticlesCreate**](ProjectsAPI.md#PrivateProjectArticlesCreate) | **Post** /account/projects/{project_id}/articles | Create project article
[**PrivateProjectArticlesList**](ProjectsAPI.md#PrivateProjectArticlesList) | **Get** /account/projects/{project_id}/articles | List project articles
[**PrivateProjectCollaboratorDelete**](ProjectsAPI.md#PrivateProjectCollaboratorDelete) | **Delete** /account/projects/{project_id}/collaborators/{user_id} | Remove project collaborator
[**PrivateProjectCollaboratorsInvite**](ProjectsAPI.md#PrivateProjectCollaboratorsInvite) | **Post** /account/projects/{project_id}/collaborators | Invite project collaborators
[**PrivateProjectCollaboratorsList**](ProjectsAPI.md#PrivateProjectCollaboratorsList) | **Get** /account/projects/{project_id}/collaborators | List project collaborators
[**PrivateProjectCreate**](ProjectsAPI.md#PrivateProjectCreate) | **Post** /account/projects | Create project
[**PrivateProjectDelete**](ProjectsAPI.md#PrivateProjectDelete) | **Delete** /account/projects/{project_id} | Delete project
[**PrivateProjectDetails**](ProjectsAPI.md#PrivateProjectDetails) | **Get** /account/projects/{project_id} | View project details
[**PrivateProjectLeave**](ProjectsAPI.md#PrivateProjectLeave) | **Post** /account/projects/{project_id}/leave | Private Project Leave
[**PrivateProjectNote**](ProjectsAPI.md#PrivateProjectNote) | **Get** /account/projects/{project_id}/notes/{note_id} | Project note details
[**PrivateProjectNoteDelete**](ProjectsAPI.md#PrivateProjectNoteDelete) | **Delete** /account/projects/{project_id}/notes/{note_id} | Delete project note
[**PrivateProjectNoteUpdate**](ProjectsAPI.md#PrivateProjectNoteUpdate) | **Put** /account/projects/{project_id}/notes/{note_id} | Update project note
[**PrivateProjectNotesCreate**](ProjectsAPI.md#PrivateProjectNotesCreate) | **Post** /account/projects/{project_id}/notes | Create project note
[**PrivateProjectNotesList**](ProjectsAPI.md#PrivateProjectNotesList) | **Get** /account/projects/{project_id}/notes | List project notes
[**PrivateProjectPartialUpdate**](ProjectsAPI.md#PrivateProjectPartialUpdate) | **Patch** /account/projects/{project_id} | Partially update project
[**PrivateProjectPublish**](ProjectsAPI.md#PrivateProjectPublish) | **Post** /account/projects/{project_id}/publish | Private Project Publish
[**PrivateProjectUpdate**](ProjectsAPI.md#PrivateProjectUpdate) | **Put** /account/projects/{project_id} | Update project
[**PrivateProjectsList**](ProjectsAPI.md#PrivateProjectsList) | **Get** /account/projects | Private Projects
[**PrivateProjectsSearch**](ProjectsAPI.md#PrivateProjectsSearch) | **Post** /account/projects/search | Private Projects search
[**ProjectArticles**](ProjectsAPI.md#ProjectArticles) | **Get** /projects/{project_id}/articles | Public Project Articles
[**ProjectDetails**](ProjectsAPI.md#ProjectDetails) | **Get** /projects/{project_id} | Public Project
[**ProjectsList**](ProjectsAPI.md#ProjectsList) | **Get** /projects | Public Projects
[**ProjectsSearch**](ProjectsAPI.md#ProjectsSearch) | **Post** /projects/search | Public Projects Search



## PrivateProjectArticleDelete

> PrivateProjectArticleDelete(ctx, projectId, articleId).Execute()

Delete project article



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
	projectId := int64(789) // int64 | Project unique identifier
	articleId := int64(789) // int64 | Project Article unique identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectsAPI.PrivateProjectArticleDelete(context.Background(), projectId, articleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.PrivateProjectArticleDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int64** | Project unique identifier | 
**articleId** | **int64** | Project Article unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateProjectArticleDeleteRequest struct via the builder pattern


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


## PrivateProjectArticleDetails

> ArticleCompletePrivate PrivateProjectArticleDetails(ctx, projectId, articleId).Execute()

Project article details



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
	projectId := int64(789) // int64 | Project unique identifier
	articleId := int64(789) // int64 | Project Article unique identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.PrivateProjectArticleDetails(context.Background(), projectId, articleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.PrivateProjectArticleDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivateProjectArticleDetails`: ArticleCompletePrivate
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.PrivateProjectArticleDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int64** | Project unique identifier | 
**articleId** | **int64** | Project Article unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateProjectArticleDetailsRequest struct via the builder pattern


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


## PrivateProjectArticleFile

> PrivateFile PrivateProjectArticleFile(ctx, projectId, articleId, fileId).Execute()

Project article file details



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
	projectId := int64(789) // int64 | Project unique identifier
	articleId := int64(789) // int64 | Project Article unique identifier
	fileId := int64(789) // int64 | File unique identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.PrivateProjectArticleFile(context.Background(), projectId, articleId, fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.PrivateProjectArticleFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivateProjectArticleFile`: PrivateFile
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.PrivateProjectArticleFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int64** | Project unique identifier | 
**articleId** | **int64** | Project Article unique identifier | 
**fileId** | **int64** | File unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateProjectArticleFileRequest struct via the builder pattern


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


## PrivateProjectArticleFiles

> []PrivateFile PrivateProjectArticleFiles(ctx, projectId, articleId).Execute()

Project article list files



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
	projectId := int64(789) // int64 | Project unique identifier
	articleId := int64(789) // int64 | Project Article unique identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.PrivateProjectArticleFiles(context.Background(), projectId, articleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.PrivateProjectArticleFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivateProjectArticleFiles`: []PrivateFile
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.PrivateProjectArticleFiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int64** | Project unique identifier | 
**articleId** | **int64** | Project Article unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateProjectArticleFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## PrivateProjectArticlesCreate

> Location PrivateProjectArticlesCreate(ctx, projectId).Article(article).Execute()

Create project article



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
	projectId := int64(789) // int64 | Project unique identifier
	article := *openapiclient.NewArticleProjectCreate("Test article title") // ArticleProjectCreate | Article description

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.PrivateProjectArticlesCreate(context.Background(), projectId).Article(article).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.PrivateProjectArticlesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivateProjectArticlesCreate`: Location
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.PrivateProjectArticlesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int64** | Project unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateProjectArticlesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **article** | [**ArticleProjectCreate**](ArticleProjectCreate.md) | Article description | 

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


## PrivateProjectArticlesList

> []Article PrivateProjectArticlesList(ctx, projectId).Execute()

List project articles



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
	projectId := int64(789) // int64 | Project unique identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.PrivateProjectArticlesList(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.PrivateProjectArticlesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivateProjectArticlesList`: []Article
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.PrivateProjectArticlesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int64** | Project unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateProjectArticlesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## PrivateProjectCollaboratorDelete

> PrivateProjectCollaboratorDelete(ctx, projectId, userId).Execute()

Remove project collaborator



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
	projectId := int64(789) // int64 | Project unique identifier
	userId := int64(789) // int64 | User unique identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectsAPI.PrivateProjectCollaboratorDelete(context.Background(), projectId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.PrivateProjectCollaboratorDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int64** | Project unique identifier | 
**userId** | **int64** | User unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateProjectCollaboratorDeleteRequest struct via the builder pattern


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


## PrivateProjectCollaboratorsInvite

> ResponseMessage PrivateProjectCollaboratorsInvite(ctx, projectId).Collaborator(collaborator).Execute()

Invite project collaborators



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
	projectId := int64(789) // int64 | Project unique identifier
	collaborator := *openapiclient.NewProjectCollaboratorInvite("viewer") // ProjectCollaboratorInvite | viewer or collaborator role. User user_id or email of user

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.PrivateProjectCollaboratorsInvite(context.Background(), projectId).Collaborator(collaborator).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.PrivateProjectCollaboratorsInvite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivateProjectCollaboratorsInvite`: ResponseMessage
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.PrivateProjectCollaboratorsInvite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int64** | Project unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateProjectCollaboratorsInviteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **collaborator** | [**ProjectCollaboratorInvite**](ProjectCollaboratorInvite.md) | viewer or collaborator role. User user_id or email of user | 

### Return type

[**ResponseMessage**](ResponseMessage.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivateProjectCollaboratorsList

> []ProjectCollaborator PrivateProjectCollaboratorsList(ctx, projectId).Execute()

List project collaborators



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
	projectId := int64(789) // int64 | Project unique identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.PrivateProjectCollaboratorsList(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.PrivateProjectCollaboratorsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivateProjectCollaboratorsList`: []ProjectCollaborator
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.PrivateProjectCollaboratorsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int64** | Project unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateProjectCollaboratorsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ProjectCollaborator**](ProjectCollaborator.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivateProjectCreate

> CreateProjectResponse PrivateProjectCreate(ctx).Project(project).Execute()

Create project



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
	project := *openapiclient.NewProjectCreate("project title") // ProjectCreate | Project  description

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.PrivateProjectCreate(context.Background()).Project(project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.PrivateProjectCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivateProjectCreate`: CreateProjectResponse
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.PrivateProjectCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrivateProjectCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project** | [**ProjectCreate**](ProjectCreate.md) | Project  description | 

### Return type

[**CreateProjectResponse**](CreateProjectResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivateProjectDelete

> PrivateProjectDelete(ctx, projectId).Execute()

Delete project



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
	projectId := int64(789) // int64 | Project unique identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectsAPI.PrivateProjectDelete(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.PrivateProjectDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int64** | Project unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateProjectDeleteRequest struct via the builder pattern


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


## PrivateProjectDetails

> ProjectCompletePrivate PrivateProjectDetails(ctx, projectId).Execute()

View project details



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
	projectId := int64(789) // int64 | Project unique identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.PrivateProjectDetails(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.PrivateProjectDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivateProjectDetails`: ProjectCompletePrivate
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.PrivateProjectDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int64** | Project unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateProjectDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProjectCompletePrivate**](ProjectCompletePrivate.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivateProjectLeave

> PrivateProjectLeave(ctx, projectId).Execute()

Private Project Leave



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
	projectId := int64(789) // int64 | Project unique identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectsAPI.PrivateProjectLeave(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.PrivateProjectLeave``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int64** | Project unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateProjectLeaveRequest struct via the builder pattern


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


## PrivateProjectNote

> ProjectNotePrivate PrivateProjectNote(ctx, projectId, noteId).Execute()

Project note details

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
	projectId := int64(789) // int64 | Project unique identifier
	noteId := int64(789) // int64 | Note unique identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.PrivateProjectNote(context.Background(), projectId, noteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.PrivateProjectNote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivateProjectNote`: ProjectNotePrivate
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.PrivateProjectNote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int64** | Project unique identifier | 
**noteId** | **int64** | Note unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateProjectNoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ProjectNotePrivate**](ProjectNotePrivate.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivateProjectNoteDelete

> PrivateProjectNoteDelete(ctx, projectId, noteId).Execute()

Delete project note

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
	projectId := int64(789) // int64 | Project unique identifier
	noteId := int64(789) // int64 | Note unique identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectsAPI.PrivateProjectNoteDelete(context.Background(), projectId, noteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.PrivateProjectNoteDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int64** | Project unique identifier | 
**noteId** | **int64** | Note unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateProjectNoteDeleteRequest struct via the builder pattern


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


## PrivateProjectNoteUpdate

> PrivateProjectNoteUpdate(ctx, projectId, noteId).Note(note).Execute()

Update project note

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
	projectId := int64(789) // int64 | Project unique identifier
	noteId := int64(789) // int64 | Note unique identifier
	note := *openapiclient.NewProjectNoteCreate("note to remember") // ProjectNoteCreate | Note message

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectsAPI.PrivateProjectNoteUpdate(context.Background(), projectId, noteId).Note(note).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.PrivateProjectNoteUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int64** | Project unique identifier | 
**noteId** | **int64** | Note unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateProjectNoteUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **note** | [**ProjectNoteCreate**](ProjectNoteCreate.md) | Note message | 

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


## PrivateProjectNotesCreate

> Location PrivateProjectNotesCreate(ctx, projectId).Note(note).Execute()

Create project note



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
	projectId := int64(789) // int64 | Project unique identifier
	note := *openapiclient.NewProjectNoteCreate("note to remember") // ProjectNoteCreate | Note message

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.PrivateProjectNotesCreate(context.Background(), projectId).Note(note).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.PrivateProjectNotesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivateProjectNotesCreate`: Location
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.PrivateProjectNotesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int64** | Project unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateProjectNotesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **note** | [**ProjectNoteCreate**](ProjectNoteCreate.md) | Note message | 

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


## PrivateProjectNotesList

> []ProjectNote PrivateProjectNotesList(ctx, projectId).Page(page).PageSize(pageSize).Limit(limit).Offset(offset).Execute()

List project notes



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
	projectId := int64(789) // int64 | Project unique identifier
	page := int64(789) // int64 | Page number. Used for pagination with page_size (optional)
	pageSize := int64(789) // int64 | The number of results included on a page. Used for pagination with page (optional) (default to 10)
	limit := int64(789) // int64 | Number of results included on a page. Used for pagination with query (optional)
	offset := int64(789) // int64 | Where to start the listing (the offset of the first result). Used for pagination with limit (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.PrivateProjectNotesList(context.Background(), projectId).Page(page).PageSize(pageSize).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.PrivateProjectNotesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivateProjectNotesList`: []ProjectNote
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.PrivateProjectNotesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int64** | Project unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateProjectNotesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int64** | Page number. Used for pagination with page_size | 
 **pageSize** | **int64** | The number of results included on a page. Used for pagination with page | [default to 10]
 **limit** | **int64** | Number of results included on a page. Used for pagination with query | 
 **offset** | **int64** | Where to start the listing (the offset of the first result). Used for pagination with limit | 

### Return type

[**[]ProjectNote**](ProjectNote.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivateProjectPartialUpdate

> PrivateProjectPartialUpdate(ctx, projectId).Project(project).Execute()

Partially update project



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
	projectId := int64(789) // int64 | Project unique identifier
	project := *openapiclient.NewProjectUpdate() // ProjectUpdate | Fields to update (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectsAPI.PrivateProjectPartialUpdate(context.Background(), projectId).Project(project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.PrivateProjectPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int64** | Project unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateProjectPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **project** | [**ProjectUpdate**](ProjectUpdate.md) | Fields to update | 

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


## PrivateProjectPublish

> ResponseMessage PrivateProjectPublish(ctx, projectId).Execute()

Private Project Publish



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
	projectId := int64(789) // int64 | Project unique identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.PrivateProjectPublish(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.PrivateProjectPublish``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivateProjectPublish`: ResponseMessage
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.PrivateProjectPublish`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int64** | Project unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateProjectPublishRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseMessage**](ResponseMessage.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivateProjectUpdate

> PrivateProjectUpdate(ctx, projectId).Project(project).Execute()

Update project



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
	projectId := int64(789) // int64 | Project unique identifier
	project := *openapiclient.NewProjectUpdate() // ProjectUpdate | Project description

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectsAPI.PrivateProjectUpdate(context.Background(), projectId).Project(project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.PrivateProjectUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int64** | Project unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateProjectUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **project** | [**ProjectUpdate**](ProjectUpdate.md) | Project description | 

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


## PrivateProjectsList

> []ProjectPrivate PrivateProjectsList(ctx).Page(page).PageSize(pageSize).Limit(limit).Offset(offset).Order(order).OrderDirection(orderDirection).Storage(storage).Roles(roles).Execute()

Private Projects



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
	order := "order_example" // string | The field by which to order. (optional) (default to "published_date")
	orderDirection := "orderDirection_example" // string |  (optional) (default to "desc")
	storage := "storage_example" // string | only return collections from this institution (optional)
	roles := "roles_example" // string | Any combination of owner, collaborator, viewer separated by comma. Examples: \"owner\" or \"owner,collaborator\". (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.PrivateProjectsList(context.Background()).Page(page).PageSize(pageSize).Limit(limit).Offset(offset).Order(order).OrderDirection(orderDirection).Storage(storage).Roles(roles).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.PrivateProjectsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivateProjectsList`: []ProjectPrivate
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.PrivateProjectsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrivateProjectsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** | Page number. Used for pagination with page_size | 
 **pageSize** | **int64** | The number of results included on a page. Used for pagination with page | [default to 10]
 **limit** | **int64** | Number of results included on a page. Used for pagination with query | 
 **offset** | **int64** | Where to start the listing (the offset of the first result). Used for pagination with limit | 
 **order** | **string** | The field by which to order. | [default to &quot;published_date&quot;]
 **orderDirection** | **string** |  | [default to &quot;desc&quot;]
 **storage** | **string** | only return collections from this institution | 
 **roles** | **string** | Any combination of owner, collaborator, viewer separated by comma. Examples: \&quot;owner\&quot; or \&quot;owner,collaborator\&quot;. | 

### Return type

[**[]ProjectPrivate**](ProjectPrivate.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivateProjectsSearch

> []ProjectPrivate PrivateProjectsSearch(ctx).Search(search).Execute()

Private Projects search



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
	search := *openapiclient.NewProjectsSearch() // ProjectsSearch | Search Parameters (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.PrivateProjectsSearch(context.Background()).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.PrivateProjectsSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivateProjectsSearch`: []ProjectPrivate
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.PrivateProjectsSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrivateProjectsSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | [**ProjectsSearch**](ProjectsSearch.md) | Search Parameters | 

### Return type

[**[]ProjectPrivate**](ProjectPrivate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectArticles

> []Article ProjectArticles(ctx, projectId).Page(page).PageSize(pageSize).Limit(limit).Offset(offset).Execute()

Public Project Articles



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
	projectId := int64(789) // int64 | Project Unique identifier
	page := int64(789) // int64 | Page number. Used for pagination with page_size (optional)
	pageSize := int64(789) // int64 | The number of results included on a page. Used for pagination with page (optional) (default to 10)
	limit := int64(789) // int64 | Number of results included on a page. Used for pagination with query (optional)
	offset := int64(789) // int64 | Where to start the listing (the offset of the first result). Used for pagination with limit (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ProjectArticles(context.Background(), projectId).Page(page).PageSize(pageSize).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectArticles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectArticles`: []Article
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ProjectArticles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int64** | Project Unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectArticlesRequest struct via the builder pattern


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


## ProjectDetails

> ProjectComplete ProjectDetails(ctx, projectId).Execute()

Public Project



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
	projectId := int64(789) // int64 | Project Unique identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ProjectDetails(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectDetails`: ProjectComplete
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ProjectDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int64** | Project Unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProjectComplete**](ProjectComplete.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsList

> []Project ProjectsList(ctx).XCursor(xCursor).Page(page).PageSize(pageSize).Limit(limit).Offset(offset).Order(order).OrderDirection(orderDirection).Institution(institution).PublishedSince(publishedSince).Group(group).Execute()

Public Projects



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
	publishedSince := "publishedSince_example" // string | Filter by article publishing date. Will only return articles published after the date. date(ISO 8601) YYYY-MM-DD (optional)
	group := int64(789) // int64 | only return collections from this group (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ProjectsList(context.Background()).XCursor(xCursor).Page(page).PageSize(pageSize).Limit(limit).Offset(offset).Order(order).OrderDirection(orderDirection).Institution(institution).PublishedSince(publishedSince).Group(group).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsList`: []Project
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ProjectsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectsListRequest struct via the builder pattern


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
 **publishedSince** | **string** | Filter by article publishing date. Will only return articles published after the date. date(ISO 8601) YYYY-MM-DD | 
 **group** | **int64** | only return collections from this group | 

### Return type

[**[]Project**](Project.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsSearch

> []Project ProjectsSearch(ctx).XCursor(xCursor).Search(search).Execute()

Public Projects Search



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
	search := *openapiclient.NewProjectsSearch() // ProjectsSearch | Search Parameters (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ProjectsSearch(context.Background()).XCursor(xCursor).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectsSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsSearch`: []Project
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ProjectsSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectsSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xCursor** | **string** | Unique hash used for bypassing the item retrieval limit of 9,000 entities. When using this parameter, please note that the offset parameter will not be available, but the limit parameter will still work as expected. | 
 **search** | [**ProjectsSearch**](ProjectsSearch.md) | Search Parameters | 

### Return type

[**[]Project**](Project.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

