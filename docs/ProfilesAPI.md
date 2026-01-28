# \ProfilesAPI

All URIs are relative to *https://api.figshare.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateUserProfile**](ProfilesAPI.md#UpdateUserProfile) | **Put** /account/profile | Update public profile
[**UpdateUserProfilePicture**](ProfilesAPI.md#UpdateUserProfilePicture) | **Post** /account/profile/{user_id}/picture | Update public profile picture



## UpdateUserProfile

> map[string]interface{} UpdateUserProfile(ctx).UserProfileData(userProfileData).UserId(userId).InstitutionUserId(institutionUserId).Execute()

Update public profile



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
	userProfileData := *openapiclient.NewProfileUpdateData() // ProfileUpdateData | 
	userId := int64(789) // int64 | User ID (optional)
	institutionUserId := "institutionUserId_example" // string | Institutional user ID (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.UpdateUserProfile(context.Background()).UserProfileData(userProfileData).UserId(userId).InstitutionUserId(institutionUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.UpdateUserProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUserProfile`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.UpdateUserProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userProfileData** | [**ProfileUpdateData**](ProfileUpdateData.md) |  | 
 **userId** | **int64** | User ID | 
 **institutionUserId** | **string** | Institutional user ID | 

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


## UpdateUserProfilePicture

> map[string]interface{} UpdateUserProfilePicture(ctx, userId).ProfilePicture(profilePicture).Execute()

Update public profile picture



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
	userId := int64(789) // int64 | User ID
	profilePicture := os.NewFile(1234, "some_file") // *os.File | User profile picture

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.UpdateUserProfilePicture(context.Background(), userId).ProfilePicture(profilePicture).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.UpdateUserProfilePicture``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUserProfilePicture`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.UpdateUserProfilePicture`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserProfilePictureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **profilePicture** | ***os.File** | User profile picture | 

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

