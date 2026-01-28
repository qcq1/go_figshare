# AuthorsCreator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authors** | **[]map[string]interface{}** | List of authors to be associated with the article. The list can contain the following fields: id, name, first_name, last_name, email, orcid_id. If an id is supplied, it will take priority and everything else will be ignored. For adding more authors use the specific authors endpoint. | 

## Methods

### NewAuthorsCreator

`func NewAuthorsCreator(authors []map[string]interface{}, ) *AuthorsCreator`

NewAuthorsCreator instantiates a new AuthorsCreator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorsCreatorWithDefaults

`func NewAuthorsCreatorWithDefaults() *AuthorsCreator`

NewAuthorsCreatorWithDefaults instantiates a new AuthorsCreator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthors

`func (o *AuthorsCreator) GetAuthors() []map[string]interface{}`

GetAuthors returns the Authors field if non-nil, zero value otherwise.

### GetAuthorsOk

`func (o *AuthorsCreator) GetAuthorsOk() (*[]map[string]interface{}, bool)`

GetAuthorsOk returns a tuple with the Authors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthors

`func (o *AuthorsCreator) SetAuthors(v []map[string]interface{})`

SetAuthors sets Authors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


