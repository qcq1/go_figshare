# OAuthToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **string** | The OAuth access token | 
**Token** | Pointer to **string** | The OAuth access token | [optional] 
**TokenType** | **string** | The type of token issued | 
**ExpiresIn** | **int64** | Token expiration time in seconds | 
**RefreshToken** | Pointer to **string** | The refresh token (only provided for certain grant types) | [optional] 
**Scope** | Pointer to **string** | The scope of the access token | [optional] 

## Methods

### NewOAuthToken

`func NewOAuthToken(accessToken string, tokenType string, expiresIn int64, ) *OAuthToken`

NewOAuthToken instantiates a new OAuthToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuthTokenWithDefaults

`func NewOAuthTokenWithDefaults() *OAuthToken`

NewOAuthTokenWithDefaults instantiates a new OAuthToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *OAuthToken) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *OAuthToken) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *OAuthToken) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetToken

`func (o *OAuthToken) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *OAuthToken) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *OAuthToken) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *OAuthToken) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTokenType

`func (o *OAuthToken) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *OAuthToken) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *OAuthToken) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.


### GetExpiresIn

`func (o *OAuthToken) GetExpiresIn() int64`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *OAuthToken) GetExpiresInOk() (*int64, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *OAuthToken) SetExpiresIn(v int64)`

SetExpiresIn sets ExpiresIn field to given value.


### GetRefreshToken

`func (o *OAuthToken) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *OAuthToken) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *OAuthToken) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *OAuthToken) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### GetScope

`func (o *OAuthToken) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *OAuthToken) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *OAuthToken) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *OAuthToken) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


