# \ProjectsApi

All URIs are relative to *https://api.figshare.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PrivateProjectArticleDelete**](ProjectsApi.md#PrivateProjectArticleDelete) | **Delete** /account/projects/{project_id}/articles/{article_id} | Delete project article
[**PrivateProjectArticleDetails**](ProjectsApi.md#PrivateProjectArticleDetails) | **Get** /account/projects/{project_id}/articles/{article_id} | Project article details
[**PrivateProjectArticleFile**](ProjectsApi.md#PrivateProjectArticleFile) | **Get** /account/projects/{project_id}/articles/{article_id}/files/{file_id} | Project article file details
[**PrivateProjectArticleFiles**](ProjectsApi.md#PrivateProjectArticleFiles) | **Get** /account/projects/{project_id}/articles/{article_id}/files | Project article list files
[**PrivateProjectArticlesCreate**](ProjectsApi.md#PrivateProjectArticlesCreate) | **Post** /account/projects/{project_id}/articles | Create project article
[**PrivateProjectArticlesList**](ProjectsApi.md#PrivateProjectArticlesList) | **Get** /account/projects/{project_id}/articles | List project articles
[**PrivateProjectCollaboratorDelete**](ProjectsApi.md#PrivateProjectCollaboratorDelete) | **Delete** /account/projects/{project_id}/collaborators/{user_id} | Remove project collaborator
[**PrivateProjectCollaboratorsInvite**](ProjectsApi.md#PrivateProjectCollaboratorsInvite) | **Post** /account/projects/{project_id}/collaborators | Invite project collaborators
[**PrivateProjectCollaboratorsList**](ProjectsApi.md#PrivateProjectCollaboratorsList) | **Get** /account/projects/{project_id}/collaborators | List project collaborators
[**PrivateProjectCreate**](ProjectsApi.md#PrivateProjectCreate) | **Post** /account/projects | Create project
[**PrivateProjectDelete**](ProjectsApi.md#PrivateProjectDelete) | **Delete** /account/projects/{project_id} | Delete project
[**PrivateProjectDetails**](ProjectsApi.md#PrivateProjectDetails) | **Get** /account/projects/{project_id} | View project details
[**PrivateProjectLeave**](ProjectsApi.md#PrivateProjectLeave) | **Post** /account/projects/{project_id}/leave | Private Project Leave
[**PrivateProjectNote**](ProjectsApi.md#PrivateProjectNote) | **Get** /account/projects/{project_id}/notes/{note_id} | Project note details
[**PrivateProjectNoteDelete**](ProjectsApi.md#PrivateProjectNoteDelete) | **Delete** /account/projects/{project_id}/notes/{note_id} | Delete project note
[**PrivateProjectNoteUpdate**](ProjectsApi.md#PrivateProjectNoteUpdate) | **Put** /account/projects/{project_id}/notes/{note_id} | Update project note
[**PrivateProjectNotesCreate**](ProjectsApi.md#PrivateProjectNotesCreate) | **Post** /account/projects/{project_id}/notes | Create project note
[**PrivateProjectNotesList**](ProjectsApi.md#PrivateProjectNotesList) | **Get** /account/projects/{project_id}/notes | List project notes
[**PrivateProjectPartialUpdate**](ProjectsApi.md#PrivateProjectPartialUpdate) | **Patch** /account/projects/{project_id} | Partially update project
[**PrivateProjectPublish**](ProjectsApi.md#PrivateProjectPublish) | **Post** /account/projects/{project_id}/publish | Private Project Publish
[**PrivateProjectUpdate**](ProjectsApi.md#PrivateProjectUpdate) | **Put** /account/projects/{project_id} | Update project
[**PrivateProjectsList**](ProjectsApi.md#PrivateProjectsList) | **Get** /account/projects | Private Projects
[**PrivateProjectsSearch**](ProjectsApi.md#PrivateProjectsSearch) | **Post** /account/projects/search | Private Projects search
[**ProjectArticles**](ProjectsApi.md#ProjectArticles) | **Get** /projects/{project_id}/articles | Public Project Articles
[**ProjectDetails**](ProjectsApi.md#ProjectDetails) | **Get** /projects/{project_id} | Public Project
[**ProjectsList**](ProjectsApi.md#ProjectsList) | **Get** /projects | Public Projects
[**ProjectsSearch**](ProjectsApi.md#ProjectsSearch) | **Post** /projects/search | Public Projects Search


# **PrivateProjectArticleDelete**
> PrivateProjectArticleDelete(ctx, projectId, articleId)
Delete project article

Delete project article

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| Project unique identifier | 
  **articleId** | **int64**| Project Article unique identifier | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateProjectArticleDetails**
> ArticleCompletePrivate PrivateProjectArticleDetails(ctx, projectId, articleId)
Project article details

Project article details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| Project unique identifier | 
  **articleId** | **int64**| Project Article unique identifier | 

### Return type

[**ArticleCompletePrivate**](ArticleCompletePrivate.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateProjectArticleFile**
> PrivateFile PrivateProjectArticleFile(ctx, projectId, articleId, fileId)
Project article file details

Project article file details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| Project unique identifier | 
  **articleId** | **int64**| Project Article unique identifier | 
  **fileId** | **int64**| File unique identifier | 

### Return type

[**PrivateFile**](PrivateFile.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateProjectArticleFiles**
> []PrivateFile PrivateProjectArticleFiles(ctx, projectId, articleId)
Project article list files

List article files

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| Project unique identifier | 
  **articleId** | **int64**| Project Article unique identifier | 

### Return type

[**[]PrivateFile**](PrivateFile.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateProjectArticlesCreate**
> Location PrivateProjectArticlesCreate(ctx, projectId, article)
Create project article

Create a new Article and associate it with this project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| Project unique identifier | 
  **article** | [**ArticleProjectCreate**](ArticleProjectCreate.md)| Article description | 

### Return type

[**Location**](Location.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateProjectArticlesList**
> []Article PrivateProjectArticlesList(ctx, projectId)
List project articles

List project articles

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| Project unique identifier | 

### Return type

[**[]Article**](Article.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateProjectCollaboratorDelete**
> PrivateProjectCollaboratorDelete(ctx, projectId, userId)
Remove project collaborator

Remove project collaborator

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| Project unique identifier | 
  **userId** | **int64**| User unique identifier | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateProjectCollaboratorsInvite**
> ResponseMessage PrivateProjectCollaboratorsInvite(ctx, projectId, collaborator)
Invite project collaborators

Invite users to collaborate on project or view the project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| Project unique identifier | 
  **collaborator** | [**ProjectCollaboratorInvite**](ProjectCollaboratorInvite.md)| viewer or collaborator role. User user_id or email of user | 

### Return type

[**ResponseMessage**](ResponseMessage.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateProjectCollaboratorsList**
> []ProjectCollaborator PrivateProjectCollaboratorsList(ctx, projectId)
List project collaborators

List Project collaborators and invited users

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| Project unique identifier | 

### Return type

[**[]ProjectCollaborator**](ProjectCollaborator.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateProjectCreate**
> CreateProjectResponse PrivateProjectCreate(ctx, project)
Create project

Create a new project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **project** | [**ProjectCreate**](ProjectCreate.md)| Project  description | 

### Return type

[**CreateProjectResponse**](CreateProjectResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateProjectDelete**
> PrivateProjectDelete(ctx, projectId)
Delete project

A project can be deleted only if: - it is not public - it does not have public articles.  When an individual project is deleted, all the articles are moved to my data of each owner.  When a group project is deleted, all the articles and files are deleted as well. Only project owner, group admin and above can delete a project. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| Project unique identifier | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateProjectDetails**
> ProjectCompletePrivate PrivateProjectDetails(ctx, projectId)
View project details

View a private project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| Project unique identifier | 

### Return type

[**ProjectCompletePrivate**](ProjectCompletePrivate.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateProjectLeave**
> PrivateProjectLeave(ctx, projectId)
Private Project Leave

Please note: project's owner cannot leave the project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| Project unique identifier | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateProjectNote**
> ProjectNotePrivate PrivateProjectNote(ctx, projectId, noteId)
Project note details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| Project unique identifier | 
  **noteId** | **int64**| Note unique identifier | 

### Return type

[**ProjectNotePrivate**](ProjectNotePrivate.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateProjectNoteDelete**
> PrivateProjectNoteDelete(ctx, projectId, noteId)
Delete project note

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| Project unique identifier | 
  **noteId** | **int64**| Note unique identifier | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateProjectNoteUpdate**
> PrivateProjectNoteUpdate(ctx, projectId, noteId, note)
Update project note

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| Project unique identifier | 
  **noteId** | **int64**| Note unique identifier | 
  **note** | [**ProjectNoteCreate**](ProjectNoteCreate.md)| Note message | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateProjectNotesCreate**
> Location PrivateProjectNotesCreate(ctx, projectId, note)
Create project note

Create a new project note

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| Project unique identifier | 
  **note** | [**ProjectNoteCreate**](ProjectNoteCreate.md)| Note message | 

### Return type

[**Location**](Location.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateProjectNotesList**
> []ProjectNote PrivateProjectNotesList(ctx, projectId, optional)
List project notes

List project notes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| Project unique identifier | 
 **optional** | ***PrivateProjectNotesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PrivateProjectNotesListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int64**| Page number. Used for pagination with page_size | 
 **pageSize** | **optional.Int64**| The number of results included on a page. Used for pagination with page | [default to 10]
 **limit** | **optional.Int64**| Number of results included on a page. Used for pagination with query | 
 **offset** | **optional.Int64**| Where to start the listing (the offset of the first result). Used for pagination with limit | 

### Return type

[**[]ProjectNote**](ProjectNote.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateProjectPartialUpdate**
> PrivateProjectPartialUpdate(ctx, projectId, optional)
Partially update project

Partially update a project; only provided fields will be changed.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| Project unique identifier | 
 **optional** | ***PrivateProjectPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PrivateProjectPartialUpdateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **project** | [**optional.Interface of ProjectUpdate**](ProjectUpdate.md)| Fields to update | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateProjectPublish**
> ResponseMessage PrivateProjectPublish(ctx, projectId)
Private Project Publish

Publish a project. Possible after all items inside it are public

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| Project unique identifier | 

### Return type

[**ResponseMessage**](ResponseMessage.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateProjectUpdate**
> PrivateProjectUpdate(ctx, projectId, project)
Update project

Updating an project by passing body parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| Project unique identifier | 
  **project** | [**ProjectUpdate**](ProjectUpdate.md)| Project description | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateProjectsList**
> []ProjectPrivate PrivateProjectsList(ctx, optional)
Private Projects

List private projects

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PrivateProjectsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PrivateProjectsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int64**| Page number. Used for pagination with page_size | 
 **pageSize** | **optional.Int64**| The number of results included on a page. Used for pagination with page | [default to 10]
 **limit** | **optional.Int64**| Number of results included on a page. Used for pagination with query | 
 **offset** | **optional.Int64**| Where to start the listing (the offset of the first result). Used for pagination with limit | 
 **order** | **optional.String**| The field by which to order. | [default to published_date]
 **orderDirection** | **optional.String**|  | [default to desc]
 **storage** | **optional.String**| only return collections from this institution | 
 **roles** | **optional.String**| Any combination of owner, collaborator, viewer separated by comma. Examples: \&quot;owner\&quot; or \&quot;owner,collaborator\&quot;. | 

### Return type

[**[]ProjectPrivate**](ProjectPrivate.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateProjectsSearch**
> []ProjectPrivate PrivateProjectsSearch(ctx, optional)
Private Projects search

Search inside the private projects

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PrivateProjectsSearchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PrivateProjectsSearchOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | [**optional.Interface of ProjectsSearch**](ProjectsSearch.md)| Search Parameters | 

### Return type

[**[]ProjectPrivate**](ProjectPrivate.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectArticles**
> []Article ProjectArticles(ctx, projectId, optional)
Public Project Articles

List articles in project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| Project Unique identifier | 
 **optional** | ***ProjectArticlesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectArticlesOpts struct

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

# **ProjectDetails**
> ProjectComplete ProjectDetails(ctx, projectId)
Public Project

View a project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| Project Unique identifier | 

### Return type

[**ProjectComplete**](ProjectComplete.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsList**
> []Project ProjectsList(ctx, optional)
Public Projects

Returns a list of public projects

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProjectsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsListOpts struct

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
 **publishedSince** | **optional.String**| Filter by article publishing date. Will only return articles published after the date. date(ISO 8601) YYYY-MM-DD | 
 **group** | **optional.Int64**| only return collections from this group | 

### Return type

[**[]Project**](Project.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsSearch**
> []Project ProjectsSearch(ctx, optional)
Public Projects Search

Returns a list of public articles

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProjectsSearchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsSearchOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xCursor** | [**optional.Interface of string**](.md)| Unique hash used for bypassing the item retrieval limit of 9,000 entities. When using this parameter, please note that the offset parameter will not be available, but the limit parameter will still work as expected. | 
 **search** | [**optional.Interface of ProjectsSearch**](ProjectsSearch.md)| Search Parameters | 

### Return type

[**[]Project**](Project.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

