//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armoracledatabase

// AutonomousDatabaseBackupsClientCreateOrUpdateResponse contains the response from method AutonomousDatabaseBackupsClient.BeginCreateOrUpdate.
type AutonomousDatabaseBackupsClientCreateOrUpdateResponse struct {
	// AutonomousDatabaseBackup resource definition
	AutonomousDatabaseBackup
}

// AutonomousDatabaseBackupsClientDeleteResponse contains the response from method AutonomousDatabaseBackupsClient.BeginDelete.
type AutonomousDatabaseBackupsClientDeleteResponse struct {
	// placeholder for future response values
}

// AutonomousDatabaseBackupsClientGetResponse contains the response from method AutonomousDatabaseBackupsClient.Get.
type AutonomousDatabaseBackupsClientGetResponse struct {
	// AutonomousDatabaseBackup resource definition
	AutonomousDatabaseBackup
}

// AutonomousDatabaseBackupsClientListByAutonomousDatabaseResponse contains the response from method AutonomousDatabaseBackupsClient.NewListByAutonomousDatabasePager.
type AutonomousDatabaseBackupsClientListByAutonomousDatabaseResponse struct {
	// The response of a AutonomousDatabaseBackup list operation.
	AutonomousDatabaseBackupListResult
}

// AutonomousDatabaseBackupsClientUpdateResponse contains the response from method AutonomousDatabaseBackupsClient.BeginUpdate.
type AutonomousDatabaseBackupsClientUpdateResponse struct {
	// AutonomousDatabaseBackup resource definition
	AutonomousDatabaseBackup
}

// AutonomousDatabaseCharacterSetsClientGetResponse contains the response from method AutonomousDatabaseCharacterSetsClient.Get.
type AutonomousDatabaseCharacterSetsClientGetResponse struct {
	// AutonomousDatabaseCharacterSets resource definition
	AutonomousDatabaseCharacterSet
}

// AutonomousDatabaseCharacterSetsClientListByLocationResponse contains the response from method AutonomousDatabaseCharacterSetsClient.NewListByLocationPager.
type AutonomousDatabaseCharacterSetsClientListByLocationResponse struct {
	// The response of a AutonomousDatabaseCharacterSet list operation.
	AutonomousDatabaseCharacterSetListResult
}

// AutonomousDatabaseNationalCharacterSetsClientGetResponse contains the response from method AutonomousDatabaseNationalCharacterSetsClient.Get.
type AutonomousDatabaseNationalCharacterSetsClientGetResponse struct {
	// AutonomousDatabaseNationalCharacterSets resource definition
	AutonomousDatabaseNationalCharacterSet
}

// AutonomousDatabaseNationalCharacterSetsClientListByLocationResponse contains the response from method AutonomousDatabaseNationalCharacterSetsClient.NewListByLocationPager.
type AutonomousDatabaseNationalCharacterSetsClientListByLocationResponse struct {
	// The response of a AutonomousDatabaseNationalCharacterSet list operation.
	AutonomousDatabaseNationalCharacterSetListResult
}

// AutonomousDatabaseVersionsClientGetResponse contains the response from method AutonomousDatabaseVersionsClient.Get.
type AutonomousDatabaseVersionsClientGetResponse struct {
	// AutonomousDbVersion resource definition
	AutonomousDbVersion
}

// AutonomousDatabaseVersionsClientListByLocationResponse contains the response from method AutonomousDatabaseVersionsClient.NewListByLocationPager.
type AutonomousDatabaseVersionsClientListByLocationResponse struct {
	// The response of a AutonomousDbVersion list operation.
	AutonomousDbVersionListResult
}

// AutonomousDatabasesClientCreateOrUpdateResponse contains the response from method AutonomousDatabasesClient.BeginCreateOrUpdate.
type AutonomousDatabasesClientCreateOrUpdateResponse struct {
	// Autonomous Database resource model.
	AutonomousDatabase
}

// AutonomousDatabasesClientDeleteResponse contains the response from method AutonomousDatabasesClient.BeginDelete.
type AutonomousDatabasesClientDeleteResponse struct {
	// placeholder for future response values
}

// AutonomousDatabasesClientFailoverResponse contains the response from method AutonomousDatabasesClient.BeginFailover.
type AutonomousDatabasesClientFailoverResponse struct {
	// Autonomous Database resource model.
	AutonomousDatabase
}

// AutonomousDatabasesClientGenerateWalletResponse contains the response from method AutonomousDatabasesClient.GenerateWallet.
type AutonomousDatabasesClientGenerateWalletResponse struct {
	// Autonomous Database Wallet File resource model.
	AutonomousDatabaseWalletFile
}

// AutonomousDatabasesClientGetResponse contains the response from method AutonomousDatabasesClient.Get.
type AutonomousDatabasesClientGetResponse struct {
	// Autonomous Database resource model.
	AutonomousDatabase
}

// AutonomousDatabasesClientListByResourceGroupResponse contains the response from method AutonomousDatabasesClient.NewListByResourceGroupPager.
type AutonomousDatabasesClientListByResourceGroupResponse struct {
	// The response of a AutonomousDatabase list operation.
	AutonomousDatabaseListResult
}

// AutonomousDatabasesClientListBySubscriptionResponse contains the response from method AutonomousDatabasesClient.NewListBySubscriptionPager.
type AutonomousDatabasesClientListBySubscriptionResponse struct {
	// The response of a AutonomousDatabase list operation.
	AutonomousDatabaseListResult
}

// AutonomousDatabasesClientSwitchoverResponse contains the response from method AutonomousDatabasesClient.BeginSwitchover.
type AutonomousDatabasesClientSwitchoverResponse struct {
	// Autonomous Database resource model.
	AutonomousDatabase
}

// AutonomousDatabasesClientUpdateResponse contains the response from method AutonomousDatabasesClient.BeginUpdate.
type AutonomousDatabasesClientUpdateResponse struct {
	// Autonomous Database resource model.
	AutonomousDatabase
}

// CloudExadataInfrastructuresClientAddStorageCapacityResponse contains the response from method CloudExadataInfrastructuresClient.BeginAddStorageCapacity.
type CloudExadataInfrastructuresClientAddStorageCapacityResponse struct {
	// CloudExadataInfrastructure resource definition
	CloudExadataInfrastructure
}

// CloudExadataInfrastructuresClientCreateOrUpdateResponse contains the response from method CloudExadataInfrastructuresClient.BeginCreateOrUpdate.
type CloudExadataInfrastructuresClientCreateOrUpdateResponse struct {
	// CloudExadataInfrastructure resource definition
	CloudExadataInfrastructure
}

// CloudExadataInfrastructuresClientDeleteResponse contains the response from method CloudExadataInfrastructuresClient.BeginDelete.
type CloudExadataInfrastructuresClientDeleteResponse struct {
	// placeholder for future response values
}

// CloudExadataInfrastructuresClientGetResponse contains the response from method CloudExadataInfrastructuresClient.Get.
type CloudExadataInfrastructuresClientGetResponse struct {
	// CloudExadataInfrastructure resource definition
	CloudExadataInfrastructure
}

// CloudExadataInfrastructuresClientListByResourceGroupResponse contains the response from method CloudExadataInfrastructuresClient.NewListByResourceGroupPager.
type CloudExadataInfrastructuresClientListByResourceGroupResponse struct {
	// The response of a CloudExadataInfrastructure list operation.
	CloudExadataInfrastructureListResult
}

// CloudExadataInfrastructuresClientListBySubscriptionResponse contains the response from method CloudExadataInfrastructuresClient.NewListBySubscriptionPager.
type CloudExadataInfrastructuresClientListBySubscriptionResponse struct {
	// The response of a CloudExadataInfrastructure list operation.
	CloudExadataInfrastructureListResult
}

// CloudExadataInfrastructuresClientUpdateResponse contains the response from method CloudExadataInfrastructuresClient.BeginUpdate.
type CloudExadataInfrastructuresClientUpdateResponse struct {
	// CloudExadataInfrastructure resource definition
	CloudExadataInfrastructure
}

// CloudVMClustersClientAddVMsResponse contains the response from method CloudVMClustersClient.BeginAddVMs.
type CloudVMClustersClientAddVMsResponse struct {
	// CloudVmCluster resource definition
	CloudVMCluster
}

// CloudVMClustersClientCreateOrUpdateResponse contains the response from method CloudVMClustersClient.BeginCreateOrUpdate.
type CloudVMClustersClientCreateOrUpdateResponse struct {
	// CloudVmCluster resource definition
	CloudVMCluster
}

// CloudVMClustersClientDeleteResponse contains the response from method CloudVMClustersClient.BeginDelete.
type CloudVMClustersClientDeleteResponse struct {
	// placeholder for future response values
}

// CloudVMClustersClientGetResponse contains the response from method CloudVMClustersClient.Get.
type CloudVMClustersClientGetResponse struct {
	// CloudVmCluster resource definition
	CloudVMCluster
}

// CloudVMClustersClientListByResourceGroupResponse contains the response from method CloudVMClustersClient.NewListByResourceGroupPager.
type CloudVMClustersClientListByResourceGroupResponse struct {
	// The response of a CloudVmCluster list operation.
	CloudVMClusterListResult
}

// CloudVMClustersClientListBySubscriptionResponse contains the response from method CloudVMClustersClient.NewListBySubscriptionPager.
type CloudVMClustersClientListBySubscriptionResponse struct {
	// The response of a CloudVmCluster list operation.
	CloudVMClusterListResult
}

// CloudVMClustersClientListPrivateIPAddressesResponse contains the response from method CloudVMClustersClient.ListPrivateIPAddresses.
type CloudVMClustersClientListPrivateIPAddressesResponse struct {
	// Array of PrivateIpAddressProperties
	PrivateIPAddressPropertiesArray []*PrivateIPAddressProperties
}

// CloudVMClustersClientRemoveVMsResponse contains the response from method CloudVMClustersClient.BeginRemoveVMs.
type CloudVMClustersClientRemoveVMsResponse struct {
	// CloudVmCluster resource definition
	CloudVMCluster
}

// CloudVMClustersClientUpdateResponse contains the response from method CloudVMClustersClient.BeginUpdate.
type CloudVMClustersClientUpdateResponse struct {
	// CloudVmCluster resource definition
	CloudVMCluster
}

// DNSPrivateViewsClientGetResponse contains the response from method DNSPrivateViewsClient.Get.
type DNSPrivateViewsClientGetResponse struct {
	// DnsPrivateView resource definition
	DNSPrivateView
}

// DNSPrivateViewsClientListByLocationResponse contains the response from method DNSPrivateViewsClient.NewListByLocationPager.
type DNSPrivateViewsClientListByLocationResponse struct {
	// The response of a DnsPrivateView list operation.
	DNSPrivateViewListResult
}

// DNSPrivateZonesClientGetResponse contains the response from method DNSPrivateZonesClient.Get.
type DNSPrivateZonesClientGetResponse struct {
	// DnsPrivateZone resource definition
	DNSPrivateZone
}

// DNSPrivateZonesClientListByLocationResponse contains the response from method DNSPrivateZonesClient.NewListByLocationPager.
type DNSPrivateZonesClientListByLocationResponse struct {
	// The response of a DnsPrivateZone list operation.
	DNSPrivateZoneListResult
}

// DbNodesClientActionResponse contains the response from method DbNodesClient.BeginAction.
type DbNodesClientActionResponse struct {
	// The DbNode resource belonging to vmCluster
	DbNode
}

// DbNodesClientGetResponse contains the response from method DbNodesClient.Get.
type DbNodesClientGetResponse struct {
	// The DbNode resource belonging to vmCluster
	DbNode
}

// DbNodesClientListByCloudVMClusterResponse contains the response from method DbNodesClient.NewListByCloudVMClusterPager.
type DbNodesClientListByCloudVMClusterResponse struct {
	// The response of a DbNode list operation.
	DbNodeListResult
}

// DbServersClientGetResponse contains the response from method DbServersClient.Get.
type DbServersClientGetResponse struct {
	// DbServer resource model
	DbServer
}

// DbServersClientListByCloudExadataInfrastructureResponse contains the response from method DbServersClient.NewListByCloudExadataInfrastructurePager.
type DbServersClientListByCloudExadataInfrastructureResponse struct {
	// The response of a DbServer list operation.
	DbServerListResult
}

// DbSystemShapesClientGetResponse contains the response from method DbSystemShapesClient.Get.
type DbSystemShapesClientGetResponse struct {
	// DbSystemShape resource definition
	DbSystemShape
}

// DbSystemShapesClientListByLocationResponse contains the response from method DbSystemShapesClient.NewListByLocationPager.
type DbSystemShapesClientListByLocationResponse struct {
	// The response of a DbSystemShape list operation.
	DbSystemShapeListResult
}

// GiVersionsClientGetResponse contains the response from method GiVersionsClient.Get.
type GiVersionsClientGetResponse struct {
	// GiVersion resource definition
	GiVersion
}

// GiVersionsClientListByLocationResponse contains the response from method GiVersionsClient.NewListByLocationPager.
type GiVersionsClientListByLocationResponse struct {
	// The response of a GiVersion list operation.
	GiVersionListResult
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to get the next set of results.
	OperationListResult
}

// OracleSubscriptionsClientCreateOrUpdateResponse contains the response from method OracleSubscriptionsClient.BeginCreateOrUpdate.
type OracleSubscriptionsClientCreateOrUpdateResponse struct {
	// OracleSubscription resource definition
	OracleSubscription
}

// OracleSubscriptionsClientDeleteResponse contains the response from method OracleSubscriptionsClient.BeginDelete.
type OracleSubscriptionsClientDeleteResponse struct {
	// placeholder for future response values
}

// OracleSubscriptionsClientGetResponse contains the response from method OracleSubscriptionsClient.Get.
type OracleSubscriptionsClientGetResponse struct {
	// OracleSubscription resource definition
	OracleSubscription
}

// OracleSubscriptionsClientListActivationLinksResponse contains the response from method OracleSubscriptionsClient.BeginListActivationLinks.
type OracleSubscriptionsClientListActivationLinksResponse struct {
	// Activation Links model
	ActivationLinks
}

// OracleSubscriptionsClientListBySubscriptionResponse contains the response from method OracleSubscriptionsClient.NewListBySubscriptionPager.
type OracleSubscriptionsClientListBySubscriptionResponse struct {
	// The response of a OracleSubscription list operation.
	OracleSubscriptionListResult
}

// OracleSubscriptionsClientListCloudAccountDetailsResponse contains the response from method OracleSubscriptionsClient.BeginListCloudAccountDetails.
type OracleSubscriptionsClientListCloudAccountDetailsResponse struct {
	// Cloud Account Details model
	CloudAccountDetails
}

// OracleSubscriptionsClientListSaasSubscriptionDetailsResponse contains the response from method OracleSubscriptionsClient.BeginListSaasSubscriptionDetails.
type OracleSubscriptionsClientListSaasSubscriptionDetailsResponse struct {
	// SaaS Subscription Details model
	SaasSubscriptionDetails
}

// OracleSubscriptionsClientUpdateResponse contains the response from method OracleSubscriptionsClient.BeginUpdate.
type OracleSubscriptionsClientUpdateResponse struct {
	// OracleSubscription resource definition
	OracleSubscription
}

// VirtualNetworkAddressesClientCreateOrUpdateResponse contains the response from method VirtualNetworkAddressesClient.BeginCreateOrUpdate.
type VirtualNetworkAddressesClientCreateOrUpdateResponse struct {
	// Virtual IP resource belonging to a vm cluster resource.
	VirtualNetworkAddress
}

// VirtualNetworkAddressesClientDeleteResponse contains the response from method VirtualNetworkAddressesClient.BeginDelete.
type VirtualNetworkAddressesClientDeleteResponse struct {
	// placeholder for future response values
}

// VirtualNetworkAddressesClientGetResponse contains the response from method VirtualNetworkAddressesClient.Get.
type VirtualNetworkAddressesClientGetResponse struct {
	// Virtual IP resource belonging to a vm cluster resource.
	VirtualNetworkAddress
}

// VirtualNetworkAddressesClientListByCloudVMClusterResponse contains the response from method VirtualNetworkAddressesClient.NewListByCloudVMClusterPager.
type VirtualNetworkAddressesClientListByCloudVMClusterResponse struct {
	// The response of a VirtualNetworkAddress list operation.
	VirtualNetworkAddressListResult
}
