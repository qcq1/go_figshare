# \ProfilesApi

All URIs are relative to *https://api.figshare.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateUserProfile**](ProfilesApi.md#UpdateUserProfile) | **Put** /account/profile | Update public profile
[**UpdateUserProfilePicture**](ProfilesApi.md#UpdateUserProfilePicture) | **Post** /account/profile/{user_id}/picture | Update public profile picture


# **UpdateUserProfile**
> interface{} UpdateUserProfile(ctx, userProfileData, optional)
Update public profile

Updates the fields of the user's public profile.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userProfileData** | [**ProfileUpdateData**](ProfileUpdateData.md)|  | 
 **optional** | ***UpdateUserProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateUserProfileOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **optional.Int64**| User ID | 
 **institutionUserId** | **optional.String**| Institutional user ID | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUserProfilePicture**
> interface{} UpdateUserProfilePicture(ctx, userId, profilePicture)
Update public profile picture

Updates the profile picture of the user's public profile.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **int64**| User ID | 
  **profilePicture** | ***os.File**| User profile picture | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

