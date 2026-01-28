# \AuthorsApi

All URIs are relative to *https://api.figshare.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PrivateAuthorDetails**](AuthorsApi.md#PrivateAuthorDetails) | **Get** /account/authors/{author_id} | Author details
[**PrivateAuthorsSearch**](AuthorsApi.md#PrivateAuthorsSearch) | **Post** /account/authors/search | Search Authors


# **PrivateAuthorDetails**
> AuthorComplete PrivateAuthorDetails(ctx, authorId)
Author details

View author details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorId** | **int64**| Author unique identifier | 

### Return type

[**AuthorComplete**](AuthorComplete.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrivateAuthorsSearch**
> []AuthorComplete PrivateAuthorsSearch(ctx, optional)
Search Authors

Search for authors

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PrivateAuthorsSearchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PrivateAuthorsSearchOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | [**optional.Interface of PrivateAuthorsSearch**](PrivateAuthorsSearch.md)| Search Parameters | 

### Return type

[**[]AuthorComplete**](AuthorComplete.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

