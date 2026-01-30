# AggregatedCitations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Navigation** | Pointer to [**AggregatedCitationsNavigation**](AggregatedCitationsNavigation.md) |  | [optional] 
**Results** | Pointer to [**[]AggregatedCitationsResultsInner**](AggregatedCitationsResultsInner.md) |  | [optional] 

## Methods

### NewAggregatedCitations

`func NewAggregatedCitations() *AggregatedCitations`

NewAggregatedCitations instantiates a new AggregatedCitations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregatedCitationsWithDefaults

`func NewAggregatedCitationsWithDefaults() *AggregatedCitations`

NewAggregatedCitationsWithDefaults instantiates a new AggregatedCitations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNavigation

`func (o *AggregatedCitations) GetNavigation() AggregatedCitationsNavigation`

GetNavigation returns the Navigation field if non-nil, zero value otherwise.

### GetNavigationOk

`func (o *AggregatedCitations) GetNavigationOk() (*AggregatedCitationsNavigation, bool)`

GetNavigationOk returns a tuple with the Navigation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNavigation

`func (o *AggregatedCitations) SetNavigation(v AggregatedCitationsNavigation)`

SetNavigation sets Navigation field to given value.

### HasNavigation

`func (o *AggregatedCitations) HasNavigation() bool`

HasNavigation returns a boolean if a field has been set.

### GetResults

`func (o *AggregatedCitations) GetResults() []AggregatedCitationsResultsInner`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *AggregatedCitations) GetResultsOk() (*[]AggregatedCitationsResultsInner, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *AggregatedCitations) SetResults(v []AggregatedCitationsResultsInner)`

SetResults sets Results field to given value.

### HasResults

`func (o *AggregatedCitations) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


