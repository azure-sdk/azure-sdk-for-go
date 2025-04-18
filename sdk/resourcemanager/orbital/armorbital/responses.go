// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armorbital

// AvailableGroundStationsClientListByCapabilityResponse contains the response from method AvailableGroundStationsClient.NewListByCapabilityPager.
type AvailableGroundStationsClientListByCapabilityResponse struct {
	// Response for the AvailableGroundStations API service call.
	AvailableGroundStationListResult
}

// ContactProfilesClientCreateOrUpdateResponse contains the response from method ContactProfilesClient.BeginCreateOrUpdate.
type ContactProfilesClientCreateOrUpdateResponse struct {
	// Customer creates a Contact Profile Resource, which will contain all of the configurations required for scheduling a contact.
	ContactProfile
}

// ContactProfilesClientDeleteResponse contains the response from method ContactProfilesClient.BeginDelete.
type ContactProfilesClientDeleteResponse struct {
	// placeholder for future response values
}

// ContactProfilesClientGetResponse contains the response from method ContactProfilesClient.Get.
type ContactProfilesClientGetResponse struct {
	// Customer creates a Contact Profile Resource, which will contain all of the configurations required for scheduling a contact.
	ContactProfile
}

// ContactProfilesClientListBySubscriptionResponse contains the response from method ContactProfilesClient.NewListBySubscriptionPager.
type ContactProfilesClientListBySubscriptionResponse struct {
	// Response for the ListContactProfiles API service call.
	ContactProfileListResult
}

// ContactProfilesClientListResponse contains the response from method ContactProfilesClient.NewListPager.
type ContactProfilesClientListResponse struct {
	// Response for the ListContactProfiles API service call.
	ContactProfileListResult
}

// ContactProfilesClientUpdateTagsResponse contains the response from method ContactProfilesClient.BeginUpdateTags.
type ContactProfilesClientUpdateTagsResponse struct {
	// Customer creates a Contact Profile Resource, which will contain all of the configurations required for scheduling a contact.
	ContactProfile
}

// ContactsClientCreateResponse contains the response from method ContactsClient.BeginCreate.
type ContactsClientCreateResponse struct {
	// Customer creates a contact resource for a spacecraft resource.
	Contact
}

// ContactsClientDeleteResponse contains the response from method ContactsClient.BeginDelete.
type ContactsClientDeleteResponse struct {
	// placeholder for future response values
}

// ContactsClientGetResponse contains the response from method ContactsClient.Get.
type ContactsClientGetResponse struct {
	// Customer creates a contact resource for a spacecraft resource.
	Contact
}

// ContactsClientListResponse contains the response from method ContactsClient.NewListPager.
type ContactsClientListResponse struct {
	// Response for the ListContacts API service call.
	ContactListResult
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to get the next set of results.
	OperationListResult
}

// OperationsResultsClientGetResponse contains the response from method OperationsResultsClient.BeginGet.
type OperationsResultsClientGetResponse struct {
	// Operation Result Entity.
	OperationResult
}

// SpacecraftsClientCreateOrUpdateResponse contains the response from method SpacecraftsClient.BeginCreateOrUpdate.
type SpacecraftsClientCreateOrUpdateResponse struct {
	// Customer creates a spacecraft resource to schedule a contact.
	Spacecraft
}

// SpacecraftsClientDeleteResponse contains the response from method SpacecraftsClient.BeginDelete.
type SpacecraftsClientDeleteResponse struct {
	// placeholder for future response values
}

// SpacecraftsClientGetResponse contains the response from method SpacecraftsClient.Get.
type SpacecraftsClientGetResponse struct {
	// Customer creates a spacecraft resource to schedule a contact.
	Spacecraft
}

// SpacecraftsClientListAvailableContactsResponse contains the response from method SpacecraftsClient.BeginListAvailableContacts.
type SpacecraftsClientListAvailableContactsResponse struct {
	// Response for the ListAvailableContacts API service call.
	AvailableContactsListResult
}

// SpacecraftsClientListBySubscriptionResponse contains the response from method SpacecraftsClient.NewListBySubscriptionPager.
type SpacecraftsClientListBySubscriptionResponse struct {
	// Response for the ListSpacecrafts API service call.
	SpacecraftListResult
}

// SpacecraftsClientListResponse contains the response from method SpacecraftsClient.NewListPager.
type SpacecraftsClientListResponse struct {
	// Response for the ListSpacecrafts API service call.
	SpacecraftListResult
}

// SpacecraftsClientUpdateTagsResponse contains the response from method SpacecraftsClient.BeginUpdateTags.
type SpacecraftsClientUpdateTagsResponse struct {
	// Customer creates a spacecraft resource to schedule a contact.
	Spacecraft
}
