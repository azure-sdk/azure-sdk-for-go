// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armeventhub

// ApplicationGroupClientCreateOrUpdateApplicationGroupResponse contains the response from method ApplicationGroupClient.CreateOrUpdateApplicationGroup.
type ApplicationGroupClientCreateOrUpdateApplicationGroupResponse struct {
	// The Application Group object
	ApplicationGroup
}

// ApplicationGroupClientDeleteResponse contains the response from method ApplicationGroupClient.Delete.
type ApplicationGroupClientDeleteResponse struct {
	// placeholder for future response values
}

// ApplicationGroupClientGetResponse contains the response from method ApplicationGroupClient.Get.
type ApplicationGroupClientGetResponse struct {
	// The Application Group object
	ApplicationGroup
}

// ApplicationGroupClientListByNamespaceResponse contains the response from method ApplicationGroupClient.NewListByNamespacePager.
type ApplicationGroupClientListByNamespaceResponse struct {
	// The response from the List Application Groups operation.
	ApplicationGroupListResult
}

// ClustersClientCreateOrUpdateResponse contains the response from method ClustersClient.BeginCreateOrUpdate.
type ClustersClientCreateOrUpdateResponse struct {
	// Single Event Hubs Cluster resource in List or Get operations.
	Cluster
}

// ClustersClientDeleteResponse contains the response from method ClustersClient.BeginDelete.
type ClustersClientDeleteResponse struct {
	// placeholder for future response values
}

// ClustersClientGetResponse contains the response from method ClustersClient.Get.
type ClustersClientGetResponse struct {
	// Single Event Hubs Cluster resource in List or Get operations.
	Cluster
}

// ClustersClientListAvailableClusterRegionResponse contains the response from method ClustersClient.ListAvailableClusterRegion.
type ClustersClientListAvailableClusterRegionResponse struct {
	// The response of the List Available Clusters operation.
	AvailableClustersList
}

// ClustersClientListByResourceGroupResponse contains the response from method ClustersClient.NewListByResourceGroupPager.
type ClustersClientListByResourceGroupResponse struct {
	// The response of the List Event Hubs Clusters operation.
	ClusterListResult
}

// ClustersClientListBySubscriptionResponse contains the response from method ClustersClient.NewListBySubscriptionPager.
type ClustersClientListBySubscriptionResponse struct {
	// The response of the List Event Hubs Clusters operation.
	ClusterListResult
}

// ClustersClientListNamespacesResponse contains the response from method ClustersClient.ListNamespaces.
type ClustersClientListNamespacesResponse struct {
	// The response of the List Namespace IDs operation
	EHNamespaceIDListResult
}

// ClustersClientUpdateResponse contains the response from method ClustersClient.BeginUpdate.
type ClustersClientUpdateResponse struct {
	// Single Event Hubs Cluster resource in List or Get operations.
	Cluster
}

// ConfigurationClientGetResponse contains the response from method ConfigurationClient.Get.
type ConfigurationClientGetResponse struct {
	// Contains all settings for the cluster.
	ClusterQuotaConfigurationProperties
}

// ConfigurationClientPatchResponse contains the response from method ConfigurationClient.Patch.
type ConfigurationClientPatchResponse struct {
	// Contains all settings for the cluster.
	ClusterQuotaConfigurationProperties
}

// ConsumerGroupsClientCreateOrUpdateResponse contains the response from method ConsumerGroupsClient.CreateOrUpdate.
type ConsumerGroupsClientCreateOrUpdateResponse struct {
	// Single item in List or Get Consumer group operation
	ConsumerGroup
}

// ConsumerGroupsClientDeleteResponse contains the response from method ConsumerGroupsClient.Delete.
type ConsumerGroupsClientDeleteResponse struct {
	// placeholder for future response values
}

// ConsumerGroupsClientGetResponse contains the response from method ConsumerGroupsClient.Get.
type ConsumerGroupsClientGetResponse struct {
	// Single item in List or Get Consumer group operation
	ConsumerGroup
}

// ConsumerGroupsClientListByEventHubResponse contains the response from method ConsumerGroupsClient.NewListByEventHubPager.
type ConsumerGroupsClientListByEventHubResponse struct {
	// The result to the List Consumer Group operation.
	ConsumerGroupListResult
}

// DisasterRecoveryConfigsClientBreakPairingResponse contains the response from method DisasterRecoveryConfigsClient.BreakPairing.
type DisasterRecoveryConfigsClientBreakPairingResponse struct {
	// placeholder for future response values
}

// DisasterRecoveryConfigsClientCheckNameAvailabilityResponse contains the response from method DisasterRecoveryConfigsClient.CheckNameAvailability.
type DisasterRecoveryConfigsClientCheckNameAvailabilityResponse struct {
	// The Result of the CheckNameAvailability operation
	CheckNameAvailabilityResult
}

// DisasterRecoveryConfigsClientCreateOrUpdateResponse contains the response from method DisasterRecoveryConfigsClient.CreateOrUpdate.
type DisasterRecoveryConfigsClientCreateOrUpdateResponse struct {
	// Single item in List or Get Alias(Disaster Recovery configuration) operation
	ArmDisasterRecovery
}

// DisasterRecoveryConfigsClientDeleteResponse contains the response from method DisasterRecoveryConfigsClient.Delete.
type DisasterRecoveryConfigsClientDeleteResponse struct {
	// placeholder for future response values
}

// DisasterRecoveryConfigsClientFailOverResponse contains the response from method DisasterRecoveryConfigsClient.FailOver.
type DisasterRecoveryConfigsClientFailOverResponse struct {
	// placeholder for future response values
}

// DisasterRecoveryConfigsClientGetAuthorizationRuleResponse contains the response from method DisasterRecoveryConfigsClient.GetAuthorizationRule.
type DisasterRecoveryConfigsClientGetAuthorizationRuleResponse struct {
	// Single item in a List or Get AuthorizationRule operation
	AuthorizationRule
}

// DisasterRecoveryConfigsClientGetResponse contains the response from method DisasterRecoveryConfigsClient.Get.
type DisasterRecoveryConfigsClientGetResponse struct {
	// Single item in List or Get Alias(Disaster Recovery configuration) operation
	ArmDisasterRecovery
}

// DisasterRecoveryConfigsClientListAuthorizationRulesResponse contains the response from method DisasterRecoveryConfigsClient.NewListAuthorizationRulesPager.
type DisasterRecoveryConfigsClientListAuthorizationRulesResponse struct {
	// The response from the List namespace operation.
	AuthorizationRuleListResult
}

// DisasterRecoveryConfigsClientListKeysResponse contains the response from method DisasterRecoveryConfigsClient.ListKeys.
type DisasterRecoveryConfigsClientListKeysResponse struct {
	// Namespace/EventHub Connection String
	AccessKeys
}

// DisasterRecoveryConfigsClientListResponse contains the response from method DisasterRecoveryConfigsClient.NewListPager.
type DisasterRecoveryConfigsClientListResponse struct {
	// The result of the List Alias(Disaster Recovery configuration) operation.
	ArmDisasterRecoveryListResult
}

// EventHubsClientCreateOrUpdateAuthorizationRuleResponse contains the response from method EventHubsClient.CreateOrUpdateAuthorizationRule.
type EventHubsClientCreateOrUpdateAuthorizationRuleResponse struct {
	// Single item in a List or Get AuthorizationRule operation
	AuthorizationRule
}

// EventHubsClientCreateOrUpdateResponse contains the response from method EventHubsClient.CreateOrUpdate.
type EventHubsClientCreateOrUpdateResponse struct {
	// Single item in List or Get Event Hub operation
	Eventhub
}

// EventHubsClientDeleteAuthorizationRuleResponse contains the response from method EventHubsClient.DeleteAuthorizationRule.
type EventHubsClientDeleteAuthorizationRuleResponse struct {
	// placeholder for future response values
}

// EventHubsClientDeleteResponse contains the response from method EventHubsClient.Delete.
type EventHubsClientDeleteResponse struct {
	// placeholder for future response values
}

// EventHubsClientGetAuthorizationRuleResponse contains the response from method EventHubsClient.GetAuthorizationRule.
type EventHubsClientGetAuthorizationRuleResponse struct {
	// Single item in a List or Get AuthorizationRule operation
	AuthorizationRule
}

// EventHubsClientGetResponse contains the response from method EventHubsClient.Get.
type EventHubsClientGetResponse struct {
	// Single item in List or Get Event Hub operation
	Eventhub
}

// EventHubsClientListAuthorizationRulesResponse contains the response from method EventHubsClient.NewListAuthorizationRulesPager.
type EventHubsClientListAuthorizationRulesResponse struct {
	// The response from the List namespace operation.
	AuthorizationRuleListResult
}

// EventHubsClientListByNamespaceResponse contains the response from method EventHubsClient.NewListByNamespacePager.
type EventHubsClientListByNamespaceResponse struct {
	// The result of the List EventHubs operation.
	ListResult
}

// EventHubsClientListKeysResponse contains the response from method EventHubsClient.ListKeys.
type EventHubsClientListKeysResponse struct {
	// Namespace/EventHub Connection String
	AccessKeys
}

// EventHubsClientRegenerateKeysResponse contains the response from method EventHubsClient.RegenerateKeys.
type EventHubsClientRegenerateKeysResponse struct {
	// Namespace/EventHub Connection String
	AccessKeys
}

// NamespacesClientCheckNameAvailabilityResponse contains the response from method NamespacesClient.CheckNameAvailability.
type NamespacesClientCheckNameAvailabilityResponse struct {
	// The Result of the CheckNameAvailability operation
	CheckNameAvailabilityResult
}

// NamespacesClientCreateOrUpdateAuthorizationRuleResponse contains the response from method NamespacesClient.CreateOrUpdateAuthorizationRule.
type NamespacesClientCreateOrUpdateAuthorizationRuleResponse struct {
	// Single item in a List or Get AuthorizationRule operation
	AuthorizationRule
}

// NamespacesClientCreateOrUpdateNetworkRuleSetResponse contains the response from method NamespacesClient.CreateOrUpdateNetworkRuleSet.
type NamespacesClientCreateOrUpdateNetworkRuleSetResponse struct {
	// Description of topic resource.
	NetworkRuleSet
}

// NamespacesClientCreateOrUpdateResponse contains the response from method NamespacesClient.BeginCreateOrUpdate.
type NamespacesClientCreateOrUpdateResponse struct {
	// Single Namespace item in List or Get Operation
	EHNamespace
}

// NamespacesClientDeleteAuthorizationRuleResponse contains the response from method NamespacesClient.DeleteAuthorizationRule.
type NamespacesClientDeleteAuthorizationRuleResponse struct {
	// placeholder for future response values
}

// NamespacesClientDeleteResponse contains the response from method NamespacesClient.BeginDelete.
type NamespacesClientDeleteResponse struct {
	// placeholder for future response values
}

// NamespacesClientFailoverResponse contains the response from method NamespacesClient.BeginFailover.
type NamespacesClientFailoverResponse struct {
	FailOver
}

// NamespacesClientGetAuthorizationRuleResponse contains the response from method NamespacesClient.GetAuthorizationRule.
type NamespacesClientGetAuthorizationRuleResponse struct {
	// Single item in a List or Get AuthorizationRule operation
	AuthorizationRule
}

// NamespacesClientGetNetworkRuleSetResponse contains the response from method NamespacesClient.GetNetworkRuleSet.
type NamespacesClientGetNetworkRuleSetResponse struct {
	// Description of topic resource.
	NetworkRuleSet
}

// NamespacesClientGetResponse contains the response from method NamespacesClient.Get.
type NamespacesClientGetResponse struct {
	// Single Namespace item in List or Get Operation
	EHNamespace
}

// NamespacesClientListAuthorizationRulesResponse contains the response from method NamespacesClient.NewListAuthorizationRulesPager.
type NamespacesClientListAuthorizationRulesResponse struct {
	// The response from the List namespace operation.
	AuthorizationRuleListResult
}

// NamespacesClientListByResourceGroupResponse contains the response from method NamespacesClient.NewListByResourceGroupPager.
type NamespacesClientListByResourceGroupResponse struct {
	// The response of the List Namespace operation
	EHNamespaceListResult
}

// NamespacesClientListKeysResponse contains the response from method NamespacesClient.ListKeys.
type NamespacesClientListKeysResponse struct {
	// Namespace/EventHub Connection String
	AccessKeys
}

// NamespacesClientListNetworkRuleSetResponse contains the response from method NamespacesClient.ListNetworkRuleSet.
type NamespacesClientListNetworkRuleSetResponse struct {
	// The response of the List NetworkRuleSet operation
	NetworkRuleSetListResult
}

// NamespacesClientListResponse contains the response from method NamespacesClient.NewListPager.
type NamespacesClientListResponse struct {
	// The response of the List Namespace operation
	EHNamespaceListResult
}

// NamespacesClientRegenerateKeysResponse contains the response from method NamespacesClient.RegenerateKeys.
type NamespacesClientRegenerateKeysResponse struct {
	// Namespace/EventHub Connection String
	AccessKeys
}

// NamespacesClientUpdateResponse contains the response from method NamespacesClient.Update.
type NamespacesClientUpdateResponse struct {
	// Single Namespace item in List or Get Operation
	EHNamespace
}

// NetworkSecurityPerimeterConfigurationClientListResponse contains the response from method NetworkSecurityPerimeterConfigurationClient.List.
type NetworkSecurityPerimeterConfigurationClientListResponse struct {
	// Result of the List NetworkSecurityPerimeterConfiguration operation.
	NetworkSecurityPerimeterConfigurationList
}

// NetworkSecurityPerimeterConfigurationsClientCreateOrUpdateResponse contains the response from method NetworkSecurityPerimeterConfigurationsClient.BeginCreateOrUpdate.
type NetworkSecurityPerimeterConfigurationsClientCreateOrUpdateResponse struct {
	// placeholder for future response values
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// Result of the request to list Event Hub operations. It contains a list of operations and a URL link to get the next set
	// of results.
	OperationListResult
}

// PrivateEndpointConnectionsClientCreateOrUpdateResponse contains the response from method PrivateEndpointConnectionsClient.CreateOrUpdate.
type PrivateEndpointConnectionsClientCreateOrUpdateResponse struct {
	// Properties of the PrivateEndpointConnection.
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientDeleteResponse contains the response from method PrivateEndpointConnectionsClient.BeginDelete.
type PrivateEndpointConnectionsClientDeleteResponse struct {
	// placeholder for future response values
}

// PrivateEndpointConnectionsClientGetResponse contains the response from method PrivateEndpointConnectionsClient.Get.
type PrivateEndpointConnectionsClientGetResponse struct {
	// Properties of the PrivateEndpointConnection.
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientListResponse contains the response from method PrivateEndpointConnectionsClient.NewListPager.
type PrivateEndpointConnectionsClientListResponse struct {
	// Result of the list of all private endpoint connections operation.
	PrivateEndpointConnectionListResult
}

// PrivateLinkResourcesClientGetResponse contains the response from method PrivateLinkResourcesClient.Get.
type PrivateLinkResourcesClientGetResponse struct {
	// Result of the List private link resources operation.
	PrivateLinkResourcesListResult
}

// SchemaRegistryClientCreateOrUpdateResponse contains the response from method SchemaRegistryClient.CreateOrUpdate.
type SchemaRegistryClientCreateOrUpdateResponse struct {
	// Single item in List or Get Schema Group operation
	SchemaGroup
}

// SchemaRegistryClientDeleteResponse contains the response from method SchemaRegistryClient.Delete.
type SchemaRegistryClientDeleteResponse struct {
	// placeholder for future response values
}

// SchemaRegistryClientGetResponse contains the response from method SchemaRegistryClient.Get.
type SchemaRegistryClientGetResponse struct {
	// Single item in List or Get Schema Group operation
	SchemaGroup
}

// SchemaRegistryClientListByNamespaceResponse contains the response from method SchemaRegistryClient.NewListByNamespacePager.
type SchemaRegistryClientListByNamespaceResponse struct {
	// The result of the List SchemaGroup operation.
	SchemaGroupListResult
}
