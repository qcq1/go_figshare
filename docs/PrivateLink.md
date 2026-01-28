# PrivateLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Private link id | 
**IsActive** | **bool** | True if private link is active | 
**ExpiresDate** | **string** | Date when link will expire | 
**HtmlLocation** | **string** | HTML url for private link | 

## Methods

### NewPrivateLink

`func NewPrivateLink(id string, isActive bool, expiresDate string, htmlLocation string, ) *PrivateLink`

NewPrivateLink instantiates a new PrivateLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateLinkWithDefaults

`func NewPrivateLinkWithDefaults() *PrivateLink`

NewPrivateLinkWithDefaults instantiates a new PrivateLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PrivateLink) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrivateLink) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrivateLink) SetId(v string)`

SetId sets Id field to given value.


### GetIsActive

`func (o *PrivateLink) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *PrivateLink) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *PrivateLink) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetExpiresDate

`func (o *PrivateLink) GetExpiresDate() string`

GetExpiresDate returns the ExpiresDate field if non-nil, zero value otherwise.

### GetExpiresDateOk

`func (o *PrivateLink) GetExpiresDateOk() (*string, bool)`

GetExpiresDateOk returns a tuple with the ExpiresDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresDate

`func (o *PrivateLink) SetExpiresDate(v string)`

SetExpiresDate sets ExpiresDate field to given value.


### GetHtmlLocation

`func (o *PrivateLink) GetHtmlLocation() string`

GetHtmlLocation returns the HtmlLocation field if non-nil, zero value otherwise.

### GetHtmlLocationOk

`func (o *PrivateLink) GetHtmlLocationOk() (*string, bool)`

GetHtmlLocationOk returns a tuple with the HtmlLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlLocation

`func (o *PrivateLink) SetHtmlLocation(v string)`

SetHtmlLocation sets HtmlLocation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


