# \OtherApi

All URIs are relative to *https://api.figshare.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CategoriesList**](OtherApi.md#CategoriesList) | **Get** /categories | Public Categories
[**FileDownload**](OtherApi.md#FileDownload) | **Get** /file/download/{file_id} | Public File Download
[**ItemTypesList**](OtherApi.md#ItemTypesList) | **Get** /item_types | Item Types
[**LicensesList**](OtherApi.md#LicensesList) | **Get** /licenses | Public Licenses
[**PrivateAccount**](OtherApi.md#PrivateAccount) | **Get** /account | Private Account information
[**PrivateFundingSearch**](OtherApi.md#PrivateFundingSearch) | **Post** /account/funding/search | Search Funding
[**PrivateLicensesList**](OtherApi.md#PrivateLicensesList) | **Get** /account/licenses | Private Account Licenses


# **CategoriesList**
> []CategoryList CategoriesList(ctx, )
Public Categories

Returns a list of public categories

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]CategoryList**](CategoryList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FileDownload**
> FileDownload(ctx, fileId)
Public File Download

Starts the download of a file

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fileId** | **int64**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/force-download

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ItemTypesList**
> []ItemType ItemTypesList(ctx, optional)
Item Types

Returns the list of Item Types of the requested group. If no user is authenticated, returns the item types available for Figshare.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ItemTypesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ItemTypesListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **optional.Int64**| Identifier of the group for which the item types are requested | [default to 0]

### Return type

[**[]ItemType**](ItemType.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LicensesList**
> []License LicensesList(ctx, )
Public Licenses

Returns a list of public licenses

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]License**](License.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateAccount**
> Account PrivateAccount(ctx, )
Private Account information

Account information for token/personal token

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Account**](Account.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateFundingSearch**
> []FundingInformation PrivateFundingSearch(ctx, optional)
Search Funding

Search for fundings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PrivateFundingSearchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PrivateFundingSearchOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | [**optional.Interface of FundingSearch**](FundingSearch.md)| Search Parameters | 

### Return type

[**[]FundingInformation**](FundingInformation.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateLicensesList**
> []License PrivateLicensesList(ctx, )
Private Account Licenses

This is a private endpoint that requires OAuth. It will return a list with figshare public licenses AND licenses defined for account's institution.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]License**](License.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

