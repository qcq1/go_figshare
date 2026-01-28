# PrivateLinkCreator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresDate** | Pointer to **string** | Date when this private link should expire - optional. By default private links expire in 365 days. | [optional] 

## Methods

### NewPrivateLinkCreator

`func NewPrivateLinkCreator() *PrivateLinkCreator`

NewPrivateLinkCreator instantiates a new PrivateLinkCreator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateLinkCreatorWithDefaults

`func NewPrivateLinkCreatorWithDefaults() *PrivateLinkCreator`

NewPrivateLinkCreatorWithDefaults instantiates a new PrivateLinkCreator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresDate

`func (o *PrivateLinkCreator) GetExpiresDate() string`

GetExpiresDate returns the ExpiresDate field if non-nil, zero value otherwise.

### GetExpiresDateOk

`func (o *PrivateLinkCreator) GetExpiresDateOk() (*string, bool)`

GetExpiresDateOk returns a tuple with the ExpiresDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresDate

`func (o *PrivateLinkCreator) SetExpiresDate(v string)`

SetExpiresDate sets ExpiresDate field to given value.

### HasExpiresDate

`func (o *PrivateLinkCreator) HasExpiresDate() bool`

HasExpiresDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


