# ProfileUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | Pointer to **string** | First name | [optional] 
**LastName** | Pointer to **string** | Last Name | [optional] 
**Orcid** | Pointer to **string** | User ORCID | [optional] 
**JobTitle** | Pointer to **string** | User job title | [optional] 
**FieldsOfInterest** | Pointer to **[]int64** | User fields of interest (category ids) | [optional] 
**FieldsOfInterestBySourceId** | Pointer to **[]string** | User fields of interest (category source IDs), supersedes the fields_of_interest property | [optional] 
**Location** | Pointer to **string** | User location | [optional] 
**Facebook** | Pointer to **string** | User facebook URL | [optional] 
**X** | Pointer to **string** | User X (twitter) URL | [optional] 
**Linkedin** | Pointer to **string** | User linkedin URL | [optional] 
**Bio** | Pointer to **string** | User biographical information | [optional] 
**PersonalProfiles** | Pointer to [**[]ProfileUpdateDataPersonalProfilesInner**](ProfileUpdateDataPersonalProfilesInner.md) | Add up to 10 additional personal profile links | [optional] 

## Methods

### NewProfileUpdateData

`func NewProfileUpdateData() *ProfileUpdateData`

NewProfileUpdateData instantiates a new ProfileUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileUpdateDataWithDefaults

`func NewProfileUpdateDataWithDefaults() *ProfileUpdateData`

NewProfileUpdateDataWithDefaults instantiates a new ProfileUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *ProfileUpdateData) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ProfileUpdateData) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ProfileUpdateData) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *ProfileUpdateData) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *ProfileUpdateData) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ProfileUpdateData) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ProfileUpdateData) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *ProfileUpdateData) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetOrcid

`func (o *ProfileUpdateData) GetOrcid() string`

GetOrcid returns the Orcid field if non-nil, zero value otherwise.

### GetOrcidOk

`func (o *ProfileUpdateData) GetOrcidOk() (*string, bool)`

GetOrcidOk returns a tuple with the Orcid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrcid

`func (o *ProfileUpdateData) SetOrcid(v string)`

SetOrcid sets Orcid field to given value.

### HasOrcid

`func (o *ProfileUpdateData) HasOrcid() bool`

HasOrcid returns a boolean if a field has been set.

### GetJobTitle

`func (o *ProfileUpdateData) GetJobTitle() string`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *ProfileUpdateData) GetJobTitleOk() (*string, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTitle

`func (o *ProfileUpdateData) SetJobTitle(v string)`

SetJobTitle sets JobTitle field to given value.

### HasJobTitle

`func (o *ProfileUpdateData) HasJobTitle() bool`

HasJobTitle returns a boolean if a field has been set.

### GetFieldsOfInterest

`func (o *ProfileUpdateData) GetFieldsOfInterest() []int64`

GetFieldsOfInterest returns the FieldsOfInterest field if non-nil, zero value otherwise.

### GetFieldsOfInterestOk

`func (o *ProfileUpdateData) GetFieldsOfInterestOk() (*[]int64, bool)`

GetFieldsOfInterestOk returns a tuple with the FieldsOfInterest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldsOfInterest

`func (o *ProfileUpdateData) SetFieldsOfInterest(v []int64)`

SetFieldsOfInterest sets FieldsOfInterest field to given value.

### HasFieldsOfInterest

`func (o *ProfileUpdateData) HasFieldsOfInterest() bool`

HasFieldsOfInterest returns a boolean if a field has been set.

### GetFieldsOfInterestBySourceId

`func (o *ProfileUpdateData) GetFieldsOfInterestBySourceId() []string`

GetFieldsOfInterestBySourceId returns the FieldsOfInterestBySourceId field if non-nil, zero value otherwise.

### GetFieldsOfInterestBySourceIdOk

`func (o *ProfileUpdateData) GetFieldsOfInterestBySourceIdOk() (*[]string, bool)`

GetFieldsOfInterestBySourceIdOk returns a tuple with the FieldsOfInterestBySourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldsOfInterestBySourceId

`func (o *ProfileUpdateData) SetFieldsOfInterestBySourceId(v []string)`

SetFieldsOfInterestBySourceId sets FieldsOfInterestBySourceId field to given value.

### HasFieldsOfInterestBySourceId

`func (o *ProfileUpdateData) HasFieldsOfInterestBySourceId() bool`

HasFieldsOfInterestBySourceId returns a boolean if a field has been set.

### GetLocation

`func (o *ProfileUpdateData) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ProfileUpdateData) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ProfileUpdateData) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ProfileUpdateData) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetFacebook

`func (o *ProfileUpdateData) GetFacebook() string`

GetFacebook returns the Facebook field if non-nil, zero value otherwise.

### GetFacebookOk

`func (o *ProfileUpdateData) GetFacebookOk() (*string, bool)`

GetFacebookOk returns a tuple with the Facebook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacebook

`func (o *ProfileUpdateData) SetFacebook(v string)`

SetFacebook sets Facebook field to given value.

### HasFacebook

`func (o *ProfileUpdateData) HasFacebook() bool`

HasFacebook returns a boolean if a field has been set.

### GetX

`func (o *ProfileUpdateData) GetX() string`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *ProfileUpdateData) GetXOk() (*string, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *ProfileUpdateData) SetX(v string)`

SetX sets X field to given value.

### HasX

`func (o *ProfileUpdateData) HasX() bool`

HasX returns a boolean if a field has been set.

### GetLinkedin

`func (o *ProfileUpdateData) GetLinkedin() string`

GetLinkedin returns the Linkedin field if non-nil, zero value otherwise.

### GetLinkedinOk

`func (o *ProfileUpdateData) GetLinkedinOk() (*string, bool)`

GetLinkedinOk returns a tuple with the Linkedin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedin

`func (o *ProfileUpdateData) SetLinkedin(v string)`

SetLinkedin sets Linkedin field to given value.

### HasLinkedin

`func (o *ProfileUpdateData) HasLinkedin() bool`

HasLinkedin returns a boolean if a field has been set.

### GetBio

`func (o *ProfileUpdateData) GetBio() string`

GetBio returns the Bio field if non-nil, zero value otherwise.

### GetBioOk

`func (o *ProfileUpdateData) GetBioOk() (*string, bool)`

GetBioOk returns a tuple with the Bio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBio

`func (o *ProfileUpdateData) SetBio(v string)`

SetBio sets Bio field to given value.

### HasBio

`func (o *ProfileUpdateData) HasBio() bool`

HasBio returns a boolean if a field has been set.

### GetPersonalProfiles

`func (o *ProfileUpdateData) GetPersonalProfiles() []ProfileUpdateDataPersonalProfilesInner`

GetPersonalProfiles returns the PersonalProfiles field if non-nil, zero value otherwise.

### GetPersonalProfilesOk

`func (o *ProfileUpdateData) GetPersonalProfilesOk() (*[]ProfileUpdateDataPersonalProfilesInner, bool)`

GetPersonalProfilesOk returns a tuple with the PersonalProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalProfiles

`func (o *ProfileUpdateData) SetPersonalProfiles(v []ProfileUpdateDataPersonalProfilesInner)`

SetPersonalProfiles sets PersonalProfiles field to given value.

### HasPersonalProfiles

`func (o *ProfileUpdateData) HasPersonalProfiles() bool`

HasPersonalProfiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


