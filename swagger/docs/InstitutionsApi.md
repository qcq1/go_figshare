# \InstitutionsApi

All URIs are relative to *https://api.figshare.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountInstitutionCuration**](InstitutionsApi.md#AccountInstitutionCuration) | **Get** /account/institution/review/{curation_id} | Institution Curation Review
[**AccountInstitutionCurations**](InstitutionsApi.md#AccountInstitutionCurations) | **Get** /account/institution/reviews | Institution Curation Reviews
[**CustomFieldsList**](InstitutionsApi.md#CustomFieldsList) | **Get** /account/institution/custom_fields | Private account institution group custom fields
[**CustomFieldsUpload**](InstitutionsApi.md#CustomFieldsUpload) | **Post** /account/institution/custom_fields/{custom_field_id}/items/upload | Custom fields values files upload
[**GetAccountInstitutionCurationComments**](InstitutionsApi.md#GetAccountInstitutionCurationComments) | **Get** /account/institution/review/{curation_id}/comments | Institution Curation Review Comments
[**InstitutionArticles**](InstitutionsApi.md#InstitutionArticles) | **Get** /institutions/{institution_string_id}/articles/filter-by | Public Institution Articles
[**InstitutionHrfeedUpload**](InstitutionsApi.md#InstitutionHrfeedUpload) | **Post** /institution/hrfeed/upload | Private Institution HRfeed Upload
[**PostAccountInstitutionCurationComments**](InstitutionsApi.md#PostAccountInstitutionCurationComments) | **Post** /account/institution/review/{curation_id}/comments | POST Institution Curation Review Comment
[**PrivateAccountInstitutionUser**](InstitutionsApi.md#PrivateAccountInstitutionUser) | **Get** /account/institution/users/{account_id} | Private Account Institution User
[**PrivateCategoriesList**](InstitutionsApi.md#PrivateCategoriesList) | **Get** /account/categories | Private Account Categories
[**PrivateGroupEmbargoOptionsDetails**](InstitutionsApi.md#PrivateGroupEmbargoOptionsDetails) | **Get** /account/institution/groups/{group_id}/embargo_options | Private Account Institution Group Embargo Options
[**PrivateInstitutionAccount**](InstitutionsApi.md#PrivateInstitutionAccount) | **Get** /account/institution/accounts/{account_id} | Private Institution Account information
[**PrivateInstitutionAccountGroupRoleDelete**](InstitutionsApi.md#PrivateInstitutionAccountGroupRoleDelete) | **Delete** /account/institution/roles/{account_id}/{group_id}/{role_id} | Delete Institution Account Group Role
[**PrivateInstitutionAccountGroupRoles**](InstitutionsApi.md#PrivateInstitutionAccountGroupRoles) | **Get** /account/institution/roles/{account_id} | List Institution Account Group Roles
[**PrivateInstitutionAccountGroupRolesCreate**](InstitutionsApi.md#PrivateInstitutionAccountGroupRolesCreate) | **Post** /account/institution/roles/{account_id} | Add Institution Account Group Roles
[**PrivateInstitutionAccountsCreate**](InstitutionsApi.md#PrivateInstitutionAccountsCreate) | **Post** /account/institution/accounts | Create new Institution Account
[**PrivateInstitutionAccountsList**](InstitutionsApi.md#PrivateInstitutionAccountsList) | **Get** /account/institution/accounts | Private Account Institution Accounts
[**PrivateInstitutionAccountsSearch**](InstitutionsApi.md#PrivateInstitutionAccountsSearch) | **Post** /account/institution/accounts/search | Private Account Institution Accounts Search
[**PrivateInstitutionAccountsUpdate**](InstitutionsApi.md#PrivateInstitutionAccountsUpdate) | **Put** /account/institution/accounts/{account_id} | Update Institution Account
[**PrivateInstitutionArticles**](InstitutionsApi.md#PrivateInstitutionArticles) | **Get** /account/institution/articles | Private Institution Articles
[**PrivateInstitutionDetails**](InstitutionsApi.md#PrivateInstitutionDetails) | **Get** /account/institution | Private Account Institutions
[**PrivateInstitutionEmbargoOptionsDetails**](InstitutionsApi.md#PrivateInstitutionEmbargoOptionsDetails) | **Get** /account/institution/embargo_options | Private Account Institution embargo options
[**PrivateInstitutionGroupsList**](InstitutionsApi.md#PrivateInstitutionGroupsList) | **Get** /account/institution/groups | Private Account Institution Groups
[**PrivateInstitutionRolesList**](InstitutionsApi.md#PrivateInstitutionRolesList) | **Get** /account/institution/roles | Private Account Institution Roles


# **AccountInstitutionCuration**
> CurationDetail AccountInstitutionCuration(ctx, curationId)
Institution Curation Review

Retrieve a certain curation review by its ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **curationId** | **int64**| ID of the curation | 

### Return type

[**CurationDetail**](CurationDetail.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AccountInstitutionCurations**
> Curation AccountInstitutionCurations(ctx, optional)
Institution Curation Reviews

Retrieve a list of curation reviews for this institution

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountInstitutionCurationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountInstitutionCurationsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **optional.Int64**| Filter by the group ID | 
 **articleId** | **optional.Int64**| Retrieve the reviews for this article | 
 **status** | **optional.String**| Filter by the status of the review | 
 **limit** | **optional.Int64**| Number of results included on a page. Used for pagination with query | 
 **offset** | **optional.Int64**| Where to start the listing (the offset of the first result). Used for pagination with limit | 

### Return type

[**Curation**](Curation.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomFieldsList**
> []ShortCustomField CustomFieldsList(ctx, optional)
Private account institution group custom fields

Returns the custom fields in the group the user belongs to, or the ones in the group specified, if the user has access.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CustomFieldsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomFieldsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **optional.Int64**| Group_id | 

### Return type

[**[]ShortCustomField**](ShortCustomField.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomFieldsUpload**
> interface{} CustomFieldsUpload(ctx, customFieldId, optional)
Custom fields values files upload

Uploads a CSV containing values for a specific custom field of type <b>dropdown_large_list</b>. More details in the <a href=\"#custom_fields\">Custom Fields section</a>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customFieldId** | **int64**| Custom field identifier | 
 **optional** | ***CustomFieldsUploadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomFieldsUploadOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **externalFile** | **optional.Interface of *os.File**| CSV file to be uploaded | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountInstitutionCurationComments**
> CurationComment GetAccountInstitutionCurationComments(ctx, curationId, optional)
Institution Curation Review Comments

Retrieve a certain curation review's comments.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **curationId** | **int64**| ID of the curation | 
 **optional** | ***GetAccountInstitutionCurationCommentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAccountInstitutionCurationCommentsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int64**| Number of results included on a page. Used for pagination with query | 
 **offset** | **optional.Int64**| Where to start the listing (the offset of the first result). Used for pagination with limit | 

### Return type

[**CurationComment**](CurationComment.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstitutionArticles**
> []Article InstitutionArticles(ctx, institutionStringId, resourceId, filename)
Public Institution Articles

Returns a list of articles belonging to the institution

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **institutionStringId** | **string**|  | 
  **resourceId** | **string**|  | 
  **filename** | **string**|  | 

### Return type

[**[]Article**](Article.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstitutionHrfeedUpload**
> ResponseMessage InstitutionHrfeedUpload(ctx, optional)
Private Institution HRfeed Upload

More info in the <a href=\"#hr_feed\">HR Feed section</a>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***InstitutionHrfeedUploadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InstitutionHrfeedUploadOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hrfeed** | **optional.Interface of *os.File**| You can find an example in the Hr Feed section | 

### Return type

[**ResponseMessage**](ResponseMessage.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostAccountInstitutionCurationComments**
> PostAccountInstitutionCurationComments(ctx, curationId, curationComment)
POST Institution Curation Review Comment

Add a new comment to the review.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **curationId** | **int64**| ID of the curation | 
  **curationComment** | [**CurationCommentCreate**](CurationCommentCreate.md)| The content/value of the comment. | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateAccountInstitutionUser**
> User PrivateAccountInstitutionUser(ctx, accountId)
Private Account Institution User

Retrieve institution user information using the account_id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**| Account identifier the user is associated to | 

### Return type

[**User**](User.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateCategoriesList**
> []CategoryList PrivateCategoriesList(ctx, )
Private Account Categories

List institution categories (including parent Categories)

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]CategoryList**](CategoryList.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateGroupEmbargoOptionsDetails**
> []GroupEmbargoOptions PrivateGroupEmbargoOptionsDetails(ctx, groupId)
Private Account Institution Group Embargo Options

Account institution group embargo options details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **int64**| Group identifier | 

### Return type

[**[]GroupEmbargoOptions**](GroupEmbargoOptions.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateInstitutionAccount**
> Account PrivateInstitutionAccount(ctx, accountId)
Private Institution Account information

Private Institution Account information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**| Account identifier the user is associated to | 

### Return type

[**Account**](Account.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateInstitutionAccountGroupRoleDelete**
> PrivateInstitutionAccountGroupRoleDelete(ctx, accountId, groupId, roleId)
Delete Institution Account Group Role

Delete Institution Account Group Role

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**| Account identifier for which to remove the role | 
  **groupId** | **int64**| Group identifier for which to remove the role | 
  **roleId** | **int64**| Role identifier | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateInstitutionAccountGroupRoles**
> AccountGroupRoles PrivateInstitutionAccountGroupRoles(ctx, accountId)
List Institution Account Group Roles

List Institution Account Group Roles

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**| Account identifier the user is associated to | 

### Return type

[**AccountGroupRoles**](AccountGroupRoles.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateInstitutionAccountGroupRolesCreate**
> PrivateInstitutionAccountGroupRolesCreate(ctx, accountId, account)
Add Institution Account Group Roles

Add Institution Account Group Roles

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**| Account identifier the user is associated to | 
  **account** | [**AccountGroupRolesCreate**](AccountGroupRolesCreate.md)| Account description | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateInstitutionAccountsCreate**
> AccountCreateResponse PrivateInstitutionAccountsCreate(ctx, account)
Create new Institution Account

Create a new Account by sending account information. When the institution_user_id is provided, no verification email will be sent. The email_verified flag will automatically be set to true. If the institution_user_id is not provided, a verification email will be sent. The email_verified flag will be set to true once the account is created.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **account** | [**AccountCreate**](AccountCreate.md)| Account description | 

### Return type

[**AccountCreateResponse**](AccountCreateResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateInstitutionAccountsList**
> []ShortAccount PrivateInstitutionAccountsList(ctx, optional)
Private Account Institution Accounts

Returns the accounts for which the account has administrative privileges (assigned and inherited).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PrivateInstitutionAccountsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PrivateInstitutionAccountsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int64**| Page number. Used for pagination with page_size | 
 **pageSize** | **optional.Int64**| The number of results included on a page. Used for pagination with page | [default to 10]
 **limit** | **optional.Int64**| Number of results included on a page. Used for pagination with query | 
 **offset** | **optional.Int64**| Where to start the listing (the offset of the first result). Used for pagination with limit | 
 **isActive** | **optional.Int64**| Filter by active status | 
 **institutionUserId** | **optional.String**| Filter by institution_user_id | 
 **email** | **optional.String**| Filter by email | 
 **idLte** | **optional.Int64**| Retrieve accounts with an ID lower or equal to the specified value | 
 **idGte** | **optional.Int64**| Retrieve accounts with an ID greater or equal to the specified value | 

### Return type

[**[]ShortAccount**](ShortAccount.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateInstitutionAccountsSearch**
> []ShortAccount PrivateInstitutionAccountsSearch(ctx, search)
Private Account Institution Accounts Search

Returns the accounts for which the account has administrative privileges (assigned and inherited).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **search** | [**InstitutionAccountsSearch**](InstitutionAccountsSearch.md)| Search Parameters | 

### Return type

[**[]ShortAccount**](ShortAccount.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateInstitutionAccountsUpdate**
> PrivateInstitutionAccountsUpdate(ctx, accountId, account)
Update Institution Account

Update Institution Account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**| Account identifier the user is associated to | 
  **account** | [**AccountUpdate**](AccountUpdate.md)| Account description | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateInstitutionArticles**
> []Article PrivateInstitutionArticles(ctx, optional)
Private Institution Articles

Get Articles from own institution. User must be administrator of the institution

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PrivateInstitutionArticlesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PrivateInstitutionArticlesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int64**| Page number. Used for pagination with page_size | 
 **pageSize** | **optional.Int64**| The number of results included on a page. Used for pagination with page | [default to 10]
 **limit** | **optional.Int64**| Number of results included on a page. Used for pagination with query | 
 **offset** | **optional.Int64**| Where to start the listing (the offset of the first result). Used for pagination with limit | 
 **order** | **optional.String**| The field by which to order. Default varies by endpoint/resource. | [default to published_date]
 **orderDirection** | **optional.String**|  | [default to desc]
 **publishedSince** | **optional.String**| Filter by article publishing date. Will only return articles published after the date. date(ISO 8601) YYYY-MM-DD or date-time(ISO 8601) YYYY-MM-DDTHH:mm:ssZ | 
 **modifiedSince** | **optional.String**| Filter by article modified date. Will only return articles published after the date. date(ISO 8601) YYYY-MM-DD or date-time(ISO 8601) YYYY-MM-DDTHH:mm:ssZ | 
 **status** | **optional.Int64**| only return collections with this status | 
 **resourceDoi** | **optional.String**| only return collections with this resource_doi | 
 **itemType** | **optional.Int64**| Only return articles with the respective type. Mapping for item_type is: 1 - Figure, 2 - Media, 3 - Dataset, 5 - Poster, 6 - Journal contribution, 7 - Presentation, 8 - Thesis, 9 - Software, 11 - Online resource, 12 - Preprint, 13 - Book, 14 - Conference contribution, 15 - Chapter, 16 - Peer review, 17 - Educational resource, 18 - Report, 19 - Standard, 20 - Composition, 21 - Funding, 22 - Physical object, 23 - Data management plan, 24 - Workflow, 25 - Monograph, 26 - Performance, 27 - Event, 28 - Service, 29 - Model | 
 **group** | **optional.Int64**| only return articles from this group | 

### Return type

[**[]Article**](Article.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateInstitutionDetails**
> Institution PrivateInstitutionDetails(ctx, )
Private Account Institutions

Account institution details

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Institution**](Institution.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateInstitutionEmbargoOptionsDetails**
> []GroupEmbargoOptions PrivateInstitutionEmbargoOptionsDetails(ctx, )
Private Account Institution embargo options

Account institution embargo options details

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]GroupEmbargoOptions**](GroupEmbargoOptions.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateInstitutionGroupsList**
> []Group PrivateInstitutionGroupsList(ctx, )
Private Account Institution Groups

Returns the groups for which the account has administrative privileges (assigned and inherited).

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Group**](Group.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateInstitutionRolesList**
> []Role PrivateInstitutionRolesList(ctx, )
Private Account Institution Roles

Returns the roles available for groups and the institution group.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Role**](Role.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

