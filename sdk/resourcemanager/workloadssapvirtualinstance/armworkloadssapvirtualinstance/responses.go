//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armworkloadssapvirtualinstance

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to get the next set of results.
	OperationListResult
}

// SapApplicationServerInstancesClientCreateResponse contains the response from method SapApplicationServerInstancesClient.BeginCreate.
type SapApplicationServerInstancesClientCreateResponse struct {
	// Define the SAP Application Server Instance resource.
	SAPApplicationServerInstance
}

// SapApplicationServerInstancesClientDeleteResponse contains the response from method SapApplicationServerInstancesClient.BeginDelete.
type SapApplicationServerInstancesClientDeleteResponse struct {
	// placeholder for future response values
}

// SapApplicationServerInstancesClientGetResponse contains the response from method SapApplicationServerInstancesClient.Get.
type SapApplicationServerInstancesClientGetResponse struct {
	// Define the SAP Application Server Instance resource.
	SAPApplicationServerInstance
}

// SapApplicationServerInstancesClientListResponse contains the response from method SapApplicationServerInstancesClient.NewListPager.
type SapApplicationServerInstancesClientListResponse struct {
	// The response of a SAPApplicationServerInstance list operation.
	SAPApplicationServerInstanceListResult
}

// SapApplicationServerInstancesClientStartResponse contains the response from method SapApplicationServerInstancesClient.BeginStart.
type SapApplicationServerInstancesClientStartResponse struct {
	// The current status of an async operation.
	OperationStatusResult
}

// SapApplicationServerInstancesClientStopResponse contains the response from method SapApplicationServerInstancesClient.BeginStop.
type SapApplicationServerInstancesClientStopResponse struct {
	// The current status of an async operation.
	OperationStatusResult
}

// SapApplicationServerInstancesClientUpdateResponse contains the response from method SapApplicationServerInstancesClient.Update.
type SapApplicationServerInstancesClientUpdateResponse struct {
	// Define the SAP Application Server Instance resource.
	SAPApplicationServerInstance
}

// SapCentralServerInstancesClientCreateResponse contains the response from method SapCentralServerInstancesClient.BeginCreate.
type SapCentralServerInstancesClientCreateResponse struct {
	// Define the SAP Central Services Instance resource.
	SAPCentralServerInstance
}

// SapCentralServerInstancesClientDeleteResponse contains the response from method SapCentralServerInstancesClient.BeginDelete.
type SapCentralServerInstancesClientDeleteResponse struct {
	// placeholder for future response values
}

// SapCentralServerInstancesClientGetResponse contains the response from method SapCentralServerInstancesClient.Get.
type SapCentralServerInstancesClientGetResponse struct {
	// Define the SAP Central Services Instance resource.
	SAPCentralServerInstance
}

// SapCentralServerInstancesClientListResponse contains the response from method SapCentralServerInstancesClient.NewListPager.
type SapCentralServerInstancesClientListResponse struct {
	// The response of a SAPCentralServerInstance list operation.
	SAPCentralServerInstanceListResult
}

// SapCentralServerInstancesClientStartResponse contains the response from method SapCentralServerInstancesClient.BeginStart.
type SapCentralServerInstancesClientStartResponse struct {
	// The current status of an async operation.
	OperationStatusResult
}

// SapCentralServerInstancesClientStopResponse contains the response from method SapCentralServerInstancesClient.BeginStop.
type SapCentralServerInstancesClientStopResponse struct {
	// The current status of an async operation.
	OperationStatusResult
}

// SapCentralServerInstancesClientUpdateResponse contains the response from method SapCentralServerInstancesClient.Update.
type SapCentralServerInstancesClientUpdateResponse struct {
	// Define the SAP Central Services Instance resource.
	SAPCentralServerInstance
}

// SapDatabaseInstancesClientCreateResponse contains the response from method SapDatabaseInstancesClient.BeginCreate.
type SapDatabaseInstancesClientCreateResponse struct {
	// Define the Database resource.
	SAPDatabaseInstance
}

// SapDatabaseInstancesClientDeleteResponse contains the response from method SapDatabaseInstancesClient.BeginDelete.
type SapDatabaseInstancesClientDeleteResponse struct {
	// placeholder for future response values
}

// SapDatabaseInstancesClientGetResponse contains the response from method SapDatabaseInstancesClient.Get.
type SapDatabaseInstancesClientGetResponse struct {
	// Define the Database resource.
	SAPDatabaseInstance
}

// SapDatabaseInstancesClientListResponse contains the response from method SapDatabaseInstancesClient.NewListPager.
type SapDatabaseInstancesClientListResponse struct {
	// The response of a SAPDatabaseInstance list operation.
	SAPDatabaseInstanceListResult
}

// SapDatabaseInstancesClientStartResponse contains the response from method SapDatabaseInstancesClient.BeginStart.
type SapDatabaseInstancesClientStartResponse struct {
	// The current status of an async operation.
	OperationStatusResult
}

// SapDatabaseInstancesClientStopResponse contains the response from method SapDatabaseInstancesClient.BeginStop.
type SapDatabaseInstancesClientStopResponse struct {
	// The current status of an async operation.
	OperationStatusResult
}

// SapDatabaseInstancesClientUpdateResponse contains the response from method SapDatabaseInstancesClient.Update.
type SapDatabaseInstancesClientUpdateResponse struct {
	// Define the Database resource.
	SAPDatabaseInstance
}

// SapVirtualInstancesClientCreateResponse contains the response from method SapVirtualInstancesClient.BeginCreate.
type SapVirtualInstancesClientCreateResponse struct {
	// Define the Virtual Instance for SAP solutions resource.
	SAPVirtualInstance
}

// SapVirtualInstancesClientDeleteResponse contains the response from method SapVirtualInstancesClient.BeginDelete.
type SapVirtualInstancesClientDeleteResponse struct {
	// placeholder for future response values
}

// SapVirtualInstancesClientGetAvailabilityZoneDetailsResponse contains the response from method SapVirtualInstancesClient.GetAvailabilityZoneDetails.
type SapVirtualInstancesClientGetAvailabilityZoneDetailsResponse struct {
	// The list of supported availability zone pairs which are part of SAP HA deployment.
	SAPAvailabilityZoneDetailsResult
}

// SapVirtualInstancesClientGetDiskConfigurationsResponse contains the response from method SapVirtualInstancesClient.GetDiskConfigurations.
type SapVirtualInstancesClientGetDiskConfigurationsResponse struct {
	// The list of disk configuration for vmSku which are part of SAP deployment.
	SAPDiskConfigurationsResult
}

// SapVirtualInstancesClientGetResponse contains the response from method SapVirtualInstancesClient.Get.
type SapVirtualInstancesClientGetResponse struct {
	// Define the Virtual Instance for SAP solutions resource.
	SAPVirtualInstance
}

// SapVirtualInstancesClientGetSapSupportedSKUResponse contains the response from method SapVirtualInstancesClient.GetSapSupportedSKU.
type SapVirtualInstancesClientGetSapSupportedSKUResponse struct {
	// The list of supported SKUs for different resources which are part of SAP deployment.
	SAPSupportedResourceSKUsResult
}

// SapVirtualInstancesClientGetSizingRecommendationsResponse contains the response from method SapVirtualInstancesClient.GetSizingRecommendations.
type SapVirtualInstancesClientGetSizingRecommendationsResponse struct {
	// The SAP sizing recommendation result.
	SAPSizingRecommendationResultClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SapVirtualInstancesClientGetSizingRecommendationsResponse.
func (s *SapVirtualInstancesClientGetSizingRecommendationsResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalSAPSizingRecommendationResultClassification(data)
	if err != nil {
		return err
	}
	s.SAPSizingRecommendationResultClassification = res
	return nil
}

// SapVirtualInstancesClientListByResourceGroupResponse contains the response from method SapVirtualInstancesClient.NewListByResourceGroupPager.
type SapVirtualInstancesClientListByResourceGroupResponse struct {
	// The response of a SAPVirtualInstance list operation.
	SAPVirtualInstanceListResult
}

// SapVirtualInstancesClientListBySubscriptionResponse contains the response from method SapVirtualInstancesClient.NewListBySubscriptionPager.
type SapVirtualInstancesClientListBySubscriptionResponse struct {
	// The response of a SAPVirtualInstance list operation.
	SAPVirtualInstanceListResult
}

// SapVirtualInstancesClientStartResponse contains the response from method SapVirtualInstancesClient.BeginStart.
type SapVirtualInstancesClientStartResponse struct {
	// The current status of an async operation.
	OperationStatusResult
}

// SapVirtualInstancesClientStopResponse contains the response from method SapVirtualInstancesClient.BeginStop.
type SapVirtualInstancesClientStopResponse struct {
	// The current status of an async operation.
	OperationStatusResult
}

// SapVirtualInstancesClientUpdateResponse contains the response from method SapVirtualInstancesClient.BeginUpdate.
type SapVirtualInstancesClientUpdateResponse struct {
	// Define the Virtual Instance for SAP solutions resource.
	SAPVirtualInstance
}
