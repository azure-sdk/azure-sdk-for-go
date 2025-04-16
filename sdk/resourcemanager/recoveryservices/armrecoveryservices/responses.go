// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservices

// ClientCapabilitiesResponse contains the response from method Client.Capabilities.
type ClientCapabilitiesResponse struct {
	// Capabilities response for Microsoft.RecoveryServices
	CapabilitiesResponse
}

// ClientCheckNameAvailabilityResponse contains the response from method Client.CheckNameAvailability.
type ClientCheckNameAvailabilityResponse struct {
	// Response for check name availability API. Resource provider will set availability as true | false.
	CheckNameAvailabilityResult
}

// OperationsClientGetOperationResultResponse contains the response from method OperationsClient.GetOperationResult.
type OperationsClientGetOperationResultResponse struct {
	// Resource information, as returned by the resource provider.
	Vault
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// Operations List response which contains list of available APIs.
	ClientDiscoveryResponse
}

// OperationsClientOperationStatusGetResponse contains the response from method OperationsClient.OperationStatusGet.
type OperationsClientOperationStatusGetResponse struct {
	// Operation Resource
	OperationResource
}

// PrivateLinkResourcesClientGetResponse contains the response from method PrivateLinkResourcesClient.Get.
type PrivateLinkResourcesClientGetResponse struct {
	// Information of the private link resource.
	PrivateLinkResource
}

// PrivateLinkResourcesClientListResponse contains the response from method PrivateLinkResourcesClient.NewListPager.
type PrivateLinkResourcesClientListResponse struct {
	// Class which represent the stamps associated with the vault.
	PrivateLinkResources
}

// RegisteredIdentitiesClientDeleteResponse contains the response from method RegisteredIdentitiesClient.Delete.
type RegisteredIdentitiesClientDeleteResponse struct {
	// placeholder for future response values
}

// ReplicationUsagesClientListResponse contains the response from method ReplicationUsagesClient.NewListPager.
type ReplicationUsagesClientListResponse struct {
	// Replication usages for vault.
	ReplicationUsageList
}

// UsagesClientListByVaultsResponse contains the response from method UsagesClient.NewListByVaultsPager.
type UsagesClientListByVaultsResponse struct {
	// Usage for vault.
	VaultUsageList
}

// VaultCertificatesClientCreateResponse contains the response from method VaultCertificatesClient.Create.
type VaultCertificatesClientCreateResponse struct {
	// Certificate corresponding to a vault that can be used by clients to register themselves with the vault.
	VaultCertificateResponse
}

// VaultExtendedInfoClientCreateOrUpdateResponse contains the response from method VaultExtendedInfoClient.CreateOrUpdate.
type VaultExtendedInfoClientCreateOrUpdateResponse struct {
	// Vault extended information.
	VaultExtendedInfoResource
}

// VaultExtendedInfoClientGetResponse contains the response from method VaultExtendedInfoClient.Get.
type VaultExtendedInfoClientGetResponse struct {
	// Vault extended information.
	VaultExtendedInfoResource
}

// VaultExtendedInfoClientUpdateResponse contains the response from method VaultExtendedInfoClient.Update.
type VaultExtendedInfoClientUpdateResponse struct {
	// Vault extended information.
	VaultExtendedInfoResource
}

// VaultsClientCreateOrUpdateResponse contains the response from method VaultsClient.BeginCreateOrUpdate.
type VaultsClientCreateOrUpdateResponse struct {
	// Resource information, as returned by the resource provider.
	Vault
}

// VaultsClientDeleteResponse contains the response from method VaultsClient.BeginDelete.
type VaultsClientDeleteResponse struct {
	// placeholder for future response values
}

// VaultsClientGetResponse contains the response from method VaultsClient.Get.
type VaultsClientGetResponse struct {
	// Resource information, as returned by the resource provider.
	Vault
}

// VaultsClientListByResourceGroupResponse contains the response from method VaultsClient.NewListByResourceGroupPager.
type VaultsClientListByResourceGroupResponse struct {
	// The response model for a list of Vaults.
	VaultList
}

// VaultsClientListBySubscriptionIDResponse contains the response from method VaultsClient.NewListBySubscriptionIDPager.
type VaultsClientListBySubscriptionIDResponse struct {
	// The response model for a list of Vaults.
	VaultList
}

// VaultsClientUpdateResponse contains the response from method VaultsClient.BeginUpdate.
type VaultsClientUpdateResponse struct {
	// Resource information, as returned by the resource provider.
	Vault
}
