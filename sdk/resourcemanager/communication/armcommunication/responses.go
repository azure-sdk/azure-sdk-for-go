//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcommunication

// DomainsClientCancelVerificationResponse contains the response from method DomainsClient.BeginCancelVerification.
type DomainsClientCancelVerificationResponse struct {
	// placeholder for future response values
}

// DomainsClientCreateOrUpdateResponse contains the response from method DomainsClient.BeginCreateOrUpdate.
type DomainsClientCreateOrUpdateResponse struct {
	// A class representing a Domains resource.
	DomainResource
}

// DomainsClientDeleteResponse contains the response from method DomainsClient.BeginDelete.
type DomainsClientDeleteResponse struct {
	// placeholder for future response values
}

// DomainsClientGetResponse contains the response from method DomainsClient.Get.
type DomainsClientGetResponse struct {
	// A class representing a Domains resource.
	DomainResource
}

// DomainsClientInitiateVerificationResponse contains the response from method DomainsClient.BeginInitiateVerification.
type DomainsClientInitiateVerificationResponse struct {
	// placeholder for future response values
}

// DomainsClientListByEmailServiceResourceResponse contains the response from method DomainsClient.NewListByEmailServiceResourcePager.
type DomainsClientListByEmailServiceResourceResponse struct {
	// Object that includes an array of Domains resource and a possible link for next set.
	DomainResourceList
}

// DomainsClientUpdateResponse contains the response from method DomainsClient.BeginUpdate.
type DomainsClientUpdateResponse struct {
	// A class representing a Domains resource.
	DomainResource
}

// EmailServicesClientCreateOrUpdateResponse contains the response from method EmailServicesClient.BeginCreateOrUpdate.
type EmailServicesClientCreateOrUpdateResponse struct {
	// A class representing an EmailService resource.
	EmailServiceResource
}

// EmailServicesClientDeleteResponse contains the response from method EmailServicesClient.BeginDelete.
type EmailServicesClientDeleteResponse struct {
	// placeholder for future response values
}

// EmailServicesClientGetResponse contains the response from method EmailServicesClient.Get.
type EmailServicesClientGetResponse struct {
	// A class representing an EmailService resource.
	EmailServiceResource
}

// EmailServicesClientListByResourceGroupResponse contains the response from method EmailServicesClient.NewListByResourceGroupPager.
type EmailServicesClientListByResourceGroupResponse struct {
	// Object that includes an array of EmailServices and a possible link for next set.
	EmailServiceResourceList
}

// EmailServicesClientListBySubscriptionResponse contains the response from method EmailServicesClient.NewListBySubscriptionPager.
type EmailServicesClientListBySubscriptionResponse struct {
	// Object that includes an array of EmailServices and a possible link for next set.
	EmailServiceResourceList
}

// EmailServicesClientListVerifiedExchangeOnlineDomainsResponse contains the response from method EmailServicesClient.ListVerifiedExchangeOnlineDomains.
type EmailServicesClientListVerifiedExchangeOnlineDomainsResponse struct {
	// List of FQDNs of verified domains in Exchange Online.
	StringArray []*string
}

// EmailServicesClientUpdateResponse contains the response from method EmailServicesClient.BeginUpdate.
type EmailServicesClientUpdateResponse struct {
	// A class representing an EmailService resource.
	EmailServiceResource
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to get the next set of results.
	OperationListResult
}

// SMTPUsernamesClientCreateOrUpdateResponse contains the response from method SMTPUsernamesClient.CreateOrUpdate.
type SMTPUsernamesClientCreateOrUpdateResponse struct {
	// The object describing the smtp username resource.
	SMTPUsernameResource
}

// SMTPUsernamesClientDeleteResponse contains the response from method SMTPUsernamesClient.Delete.
type SMTPUsernamesClientDeleteResponse struct {
	// placeholder for future response values
}

// SMTPUsernamesClientGetResponse contains the response from method SMTPUsernamesClient.Get.
type SMTPUsernamesClientGetResponse struct {
	// The object describing the smtp username resource.
	SMTPUsernameResource
}

// SMTPUsernamesClientListResponse contains the response from method SMTPUsernamesClient.NewListPager.
type SMTPUsernamesClientListResponse struct {
	// Collection of SmtpUsername resources. Response will include a nextLink if response contains more pages.
	SMTPUsernameResourceCollection
}

// SenderUsernamesClientCreateOrUpdateResponse contains the response from method SenderUsernamesClient.CreateOrUpdate.
type SenderUsernamesClientCreateOrUpdateResponse struct {
	// A class representing a SenderUsername resource.
	SenderUsernameResource
}

// SenderUsernamesClientDeleteResponse contains the response from method SenderUsernamesClient.Delete.
type SenderUsernamesClientDeleteResponse struct {
	// placeholder for future response values
}

// SenderUsernamesClientGetResponse contains the response from method SenderUsernamesClient.Get.
type SenderUsernamesClientGetResponse struct {
	// A class representing a SenderUsername resource.
	SenderUsernameResource
}

// SenderUsernamesClientListByDomainsResponse contains the response from method SenderUsernamesClient.NewListByDomainsPager.
type SenderUsernamesClientListByDomainsResponse struct {
	// A class representing a Domains SenderUsernames collection.
	SenderUsernameResourceCollection
}

// ServicesClientCheckNameAvailabilityResponse contains the response from method ServicesClient.CheckNameAvailability.
type ServicesClientCheckNameAvailabilityResponse struct {
	// The check availability result.
	CheckNameAvailabilityResponse
}

// ServicesClientCreateOrUpdateResponse contains the response from method ServicesClient.BeginCreateOrUpdate.
type ServicesClientCreateOrUpdateResponse struct {
	// A class representing a CommunicationService resource.
	ServiceResource
}

// ServicesClientDeleteResponse contains the response from method ServicesClient.BeginDelete.
type ServicesClientDeleteResponse struct {
	// placeholder for future response values
}

// ServicesClientGetResponse contains the response from method ServicesClient.Get.
type ServicesClientGetResponse struct {
	// A class representing a CommunicationService resource.
	ServiceResource
}

// ServicesClientLinkNotificationHubResponse contains the response from method ServicesClient.LinkNotificationHub.
type ServicesClientLinkNotificationHubResponse struct {
	// A notification hub that has been linked to the communication service
	LinkedNotificationHub
}

// ServicesClientListByResourceGroupResponse contains the response from method ServicesClient.NewListByResourceGroupPager.
type ServicesClientListByResourceGroupResponse struct {
	// Object that includes an array of CommunicationServices and a possible link for next set.
	ServiceResourceList
}

// ServicesClientListBySubscriptionResponse contains the response from method ServicesClient.NewListBySubscriptionPager.
type ServicesClientListBySubscriptionResponse struct {
	// Object that includes an array of CommunicationServices and a possible link for next set.
	ServiceResourceList
}

// ServicesClientListKeysResponse contains the response from method ServicesClient.ListKeys.
type ServicesClientListKeysResponse struct {
	// A class representing the access keys of a CommunicationService.
	ServiceKeys
}

// ServicesClientRegenerateKeyResponse contains the response from method ServicesClient.RegenerateKey.
type ServicesClientRegenerateKeyResponse struct {
	// A class representing the access keys of a CommunicationService.
	ServiceKeys
}

// ServicesClientUpdateResponse contains the response from method ServicesClient.Update.
type ServicesClientUpdateResponse struct {
	// A class representing a CommunicationService resource.
	ServiceResource
}

// SuppressionListAddressesClientCreateOrUpdateResponse contains the response from method SuppressionListAddressesClient.CreateOrUpdate.
type SuppressionListAddressesClientCreateOrUpdateResponse struct {
	// A object that represents a SuppressionList record.
	SuppressionListAddressResource
}

// SuppressionListAddressesClientDeleteResponse contains the response from method SuppressionListAddressesClient.Delete.
type SuppressionListAddressesClientDeleteResponse struct {
	// placeholder for future response values
}

// SuppressionListAddressesClientGetResponse contains the response from method SuppressionListAddressesClient.Get.
type SuppressionListAddressesClientGetResponse struct {
	// A object that represents a SuppressionList record.
	SuppressionListAddressResource
}

// SuppressionListAddressesClientListResponse contains the response from method SuppressionListAddressesClient.NewListPager.
type SuppressionListAddressesClientListResponse struct {
	// Collection of addresses in a suppression list. Response will include a nextLink if response contains more pages.
	SuppressionListAddressResourceCollection
}

// SuppressionListsClientCreateOrUpdateResponse contains the response from method SuppressionListsClient.CreateOrUpdate.
type SuppressionListsClientCreateOrUpdateResponse struct {
	// A class representing a SuppressionList resource.
	SuppressionListResource
}

// SuppressionListsClientDeleteResponse contains the response from method SuppressionListsClient.Delete.
type SuppressionListsClientDeleteResponse struct {
	// placeholder for future response values
}

// SuppressionListsClientGetResponse contains the response from method SuppressionListsClient.Get.
type SuppressionListsClientGetResponse struct {
	// A class representing a SuppressionList resource.
	SuppressionListResource
}

// SuppressionListsClientListByDomainResponse contains the response from method SuppressionListsClient.NewListByDomainPager.
type SuppressionListsClientListByDomainResponse struct {
	// A class representing a Domains SuppressionListResource collection.
	SuppressionListResourceCollection
}
