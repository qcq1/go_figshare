# CreateOAuthToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** |  | 
**ClientSecret** | **string** |  | 
**Code** | Pointer to **string** | Required if grant_type is &#39;authorization_code&#39; | [optional] 
**GrantType** | **string** |  | 
**Password** | Pointer to **string** | Required if grant_type is &#39;password&#39; | [optional] 
**RefreshToken** | Pointer to **string** | Required if grant_type is &#39;refresh_token&#39; | [optional] 
**Username** | Pointer to **string** | Required if grant_type is &#39;password&#39; | [optional] 

## Methods

### NewCreateOAuthToken

`func NewCreateOAuthToken(clientId string, clientSecret string, grantType string, ) *CreateOAuthToken`

NewCreateOAuthToken instantiates a new CreateOAuthToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOAuthTokenWithDefaults

`func NewCreateOAuthTokenWithDefaults() *CreateOAuthToken`

NewCreateOAuthTokenWithDefaults instantiates a new CreateOAuthToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *CreateOAuthToken) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *CreateOAuthToken) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *CreateOAuthToken) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *CreateOAuthToken) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *CreateOAuthToken) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *CreateOAuthToken) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetCode

`func (o *CreateOAuthToken) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CreateOAuthToken) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CreateOAuthToken) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CreateOAuthToken) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetGrantType

`func (o *CreateOAuthToken) GetGrantType() string`

GetGrantType returns the GrantType field if non-nil, zero value otherwise.

### GetGrantTypeOk

`func (o *CreateOAuthToken) GetGrantTypeOk() (*string, bool)`

GetGrantTypeOk returns a tuple with the GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantType

`func (o *CreateOAuthToken) SetGrantType(v string)`

SetGrantType sets GrantType field to given value.


### GetPassword

`func (o *CreateOAuthToken) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateOAuthToken) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateOAuthToken) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CreateOAuthToken) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetRefreshToken

`func (o *CreateOAuthToken) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *CreateOAuthToken) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *CreateOAuthToken) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *CreateOAuthToken) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### GetUsername

`func (o *CreateOAuthToken) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateOAuthToken) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateOAuthToken) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *CreateOAuthToken) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


