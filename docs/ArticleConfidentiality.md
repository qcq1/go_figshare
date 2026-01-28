# ArticleConfidentiality

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsConfidential** | **bool** | True if article is confidential | 
**Reason** | **string** | Reason for confidentiality | 

## Methods

### NewArticleConfidentiality

`func NewArticleConfidentiality(isConfidential bool, reason string, ) *ArticleConfidentiality`

NewArticleConfidentiality instantiates a new ArticleConfidentiality object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArticleConfidentialityWithDefaults

`func NewArticleConfidentialityWithDefaults() *ArticleConfidentiality`

NewArticleConfidentialityWithDefaults instantiates a new ArticleConfidentiality object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsConfidential

`func (o *ArticleConfidentiality) GetIsConfidential() bool`

GetIsConfidential returns the IsConfidential field if non-nil, zero value otherwise.

### GetIsConfidentialOk

`func (o *ArticleConfidentiality) GetIsConfidentialOk() (*bool, bool)`

GetIsConfidentialOk returns a tuple with the IsConfidential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConfidential

`func (o *ArticleConfidentiality) SetIsConfidential(v bool)`

SetIsConfidential sets IsConfidential field to given value.


### GetReason

`func (o *ArticleConfidentiality) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ArticleConfidentiality) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ArticleConfidentiality) SetReason(v string)`

SetReason sets Reason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


