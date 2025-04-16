// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcustomerinsights

// AuthorizationPoliciesClientCreateOrUpdateResponse contains the response from method AuthorizationPoliciesClient.CreateOrUpdate.
type AuthorizationPoliciesClientCreateOrUpdateResponse struct {
	// The authorization policy resource format.
	AuthorizationPolicyResourceFormat
}

// AuthorizationPoliciesClientGetResponse contains the response from method AuthorizationPoliciesClient.Get.
type AuthorizationPoliciesClientGetResponse struct {
	// The authorization policy resource format.
	AuthorizationPolicyResourceFormat
}

// AuthorizationPoliciesClientListByHubResponse contains the response from method AuthorizationPoliciesClient.NewListByHubPager.
type AuthorizationPoliciesClientListByHubResponse struct {
	// The response of list authorization policy operation.
	AuthorizationPolicyListResult
}

// AuthorizationPoliciesClientRegeneratePrimaryKeyResponse contains the response from method AuthorizationPoliciesClient.RegeneratePrimaryKey.
type AuthorizationPoliciesClientRegeneratePrimaryKeyResponse struct {
	// The authorization policy.
	AuthorizationPolicy
}

// AuthorizationPoliciesClientRegenerateSecondaryKeyResponse contains the response from method AuthorizationPoliciesClient.RegenerateSecondaryKey.
type AuthorizationPoliciesClientRegenerateSecondaryKeyResponse struct {
	// The authorization policy.
	AuthorizationPolicy
}

// ConnectorMappingsClientCreateOrUpdateResponse contains the response from method ConnectorMappingsClient.CreateOrUpdate.
type ConnectorMappingsClientCreateOrUpdateResponse struct {
	// The connector mapping resource format.
	ConnectorMappingResourceFormat
}

// ConnectorMappingsClientDeleteResponse contains the response from method ConnectorMappingsClient.Delete.
type ConnectorMappingsClientDeleteResponse struct {
	// placeholder for future response values
}

// ConnectorMappingsClientGetResponse contains the response from method ConnectorMappingsClient.Get.
type ConnectorMappingsClientGetResponse struct {
	// The connector mapping resource format.
	ConnectorMappingResourceFormat
}

// ConnectorMappingsClientListByConnectorResponse contains the response from method ConnectorMappingsClient.NewListByConnectorPager.
type ConnectorMappingsClientListByConnectorResponse struct {
	// The response of list connector mapping operation.
	ConnectorMappingListResult
}

// ConnectorsClientCreateOrUpdateResponse contains the response from method ConnectorsClient.BeginCreateOrUpdate.
type ConnectorsClientCreateOrUpdateResponse struct {
	// The connector resource format.
	ConnectorResourceFormat
}

// ConnectorsClientDeleteResponse contains the response from method ConnectorsClient.BeginDelete.
type ConnectorsClientDeleteResponse struct {
	// placeholder for future response values
}

// ConnectorsClientGetResponse contains the response from method ConnectorsClient.Get.
type ConnectorsClientGetResponse struct {
	// The connector resource format.
	ConnectorResourceFormat
}

// ConnectorsClientListByHubResponse contains the response from method ConnectorsClient.NewListByHubPager.
type ConnectorsClientListByHubResponse struct {
	// The response of list connector operation.
	ConnectorListResult
}

// HubsClientCreateOrUpdateResponse contains the response from method HubsClient.CreateOrUpdate.
type HubsClientCreateOrUpdateResponse struct {
	// Hub resource.
	Hub
}

// HubsClientDeleteResponse contains the response from method HubsClient.BeginDelete.
type HubsClientDeleteResponse struct {
	// placeholder for future response values
}

// HubsClientGetResponse contains the response from method HubsClient.Get.
type HubsClientGetResponse struct {
	// Hub resource.
	Hub
}

// HubsClientListByResourceGroupResponse contains the response from method HubsClient.NewListByResourceGroupPager.
type HubsClientListByResourceGroupResponse struct {
	// Response of list hub operation.
	HubListResult
}

// HubsClientListResponse contains the response from method HubsClient.NewListPager.
type HubsClientListResponse struct {
	// Response of list hub operation.
	HubListResult
}

// HubsClientUpdateResponse contains the response from method HubsClient.Update.
type HubsClientUpdateResponse struct {
	// Hub resource.
	Hub
}

// ImagesClientGetUploadURLForDataResponse contains the response from method ImagesClient.GetUploadURLForData.
type ImagesClientGetUploadURLForDataResponse struct {
	// The image definition.
	ImageDefinition
}

// ImagesClientGetUploadURLForEntityTypeResponse contains the response from method ImagesClient.GetUploadURLForEntityType.
type ImagesClientGetUploadURLForEntityTypeResponse struct {
	// The image definition.
	ImageDefinition
}

// InteractionsClientCreateOrUpdateResponse contains the response from method InteractionsClient.BeginCreateOrUpdate.
type InteractionsClientCreateOrUpdateResponse struct {
	// The interaction resource format.
	InteractionResourceFormat
}

// InteractionsClientGetResponse contains the response from method InteractionsClient.Get.
type InteractionsClientGetResponse struct {
	// The interaction resource format.
	InteractionResourceFormat
}

// InteractionsClientListByHubResponse contains the response from method InteractionsClient.NewListByHubPager.
type InteractionsClientListByHubResponse struct {
	// The response of list interaction operation.
	InteractionListResult
}

// InteractionsClientSuggestRelationshipLinksResponse contains the response from method InteractionsClient.SuggestRelationshipLinks.
type InteractionsClientSuggestRelationshipLinksResponse struct {
	// The response of suggest relationship links operation.
	SuggestRelationshipLinksResponse
}

// KpiClientCreateOrUpdateResponse contains the response from method KpiClient.BeginCreateOrUpdate.
type KpiClientCreateOrUpdateResponse struct {
	// The KPI resource format.
	KpiResourceFormat
}

// KpiClientDeleteResponse contains the response from method KpiClient.BeginDelete.
type KpiClientDeleteResponse struct {
	// placeholder for future response values
}

// KpiClientGetResponse contains the response from method KpiClient.Get.
type KpiClientGetResponse struct {
	// The KPI resource format.
	KpiResourceFormat
}

// KpiClientListByHubResponse contains the response from method KpiClient.NewListByHubPager.
type KpiClientListByHubResponse struct {
	// The response of list KPI operation.
	KpiListResult
}

// KpiClientReprocessResponse contains the response from method KpiClient.Reprocess.
type KpiClientReprocessResponse struct {
	// placeholder for future response values
}

// LinksClientCreateOrUpdateResponse contains the response from method LinksClient.BeginCreateOrUpdate.
type LinksClientCreateOrUpdateResponse struct {
	// The link resource format.
	LinkResourceFormat
}

// LinksClientDeleteResponse contains the response from method LinksClient.Delete.
type LinksClientDeleteResponse struct {
	// placeholder for future response values
}

// LinksClientGetResponse contains the response from method LinksClient.Get.
type LinksClientGetResponse struct {
	// The link resource format.
	LinkResourceFormat
}

// LinksClientListByHubResponse contains the response from method LinksClient.NewListByHubPager.
type LinksClientListByHubResponse struct {
	// The response of list link operation.
	LinkListResult
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// Result of the request to list Customer Insights operations. It contains a list of operations and a URL link to get the
	// next set of results.
	OperationListResult
}

// PredictionsClientCreateOrUpdateResponse contains the response from method PredictionsClient.BeginCreateOrUpdate.
type PredictionsClientCreateOrUpdateResponse struct {
	// The prediction resource format.
	PredictionResourceFormat
}

// PredictionsClientDeleteResponse contains the response from method PredictionsClient.BeginDelete.
type PredictionsClientDeleteResponse struct {
	// placeholder for future response values
}

// PredictionsClientGetModelStatusResponse contains the response from method PredictionsClient.GetModelStatus.
type PredictionsClientGetModelStatusResponse struct {
	// The prediction model status.
	PredictionModelStatus
}

// PredictionsClientGetResponse contains the response from method PredictionsClient.Get.
type PredictionsClientGetResponse struct {
	// The prediction resource format.
	PredictionResourceFormat
}

// PredictionsClientGetTrainingResultsResponse contains the response from method PredictionsClient.GetTrainingResults.
type PredictionsClientGetTrainingResultsResponse struct {
	// The training results of the prediction.
	PredictionTrainingResults
}

// PredictionsClientListByHubResponse contains the response from method PredictionsClient.NewListByHubPager.
type PredictionsClientListByHubResponse struct {
	// The response of list predictions operation.
	PredictionListResult
}

// PredictionsClientModelStatusResponse contains the response from method PredictionsClient.ModelStatus.
type PredictionsClientModelStatusResponse struct {
	// placeholder for future response values
}

// ProfilesClientCreateOrUpdateResponse contains the response from method ProfilesClient.BeginCreateOrUpdate.
type ProfilesClientCreateOrUpdateResponse struct {
	// The profile resource format.
	ProfileResourceFormat
}

// ProfilesClientDeleteResponse contains the response from method ProfilesClient.BeginDelete.
type ProfilesClientDeleteResponse struct {
	// placeholder for future response values
}

// ProfilesClientGetEnrichingKpisResponse contains the response from method ProfilesClient.GetEnrichingKpis.
type ProfilesClientGetEnrichingKpisResponse struct {
	// Array of KpiDefinition
	KpiDefinitionArray []*KpiDefinition
}

// ProfilesClientGetResponse contains the response from method ProfilesClient.Get.
type ProfilesClientGetResponse struct {
	// The profile resource format.
	ProfileResourceFormat
}

// ProfilesClientListByHubResponse contains the response from method ProfilesClient.NewListByHubPager.
type ProfilesClientListByHubResponse struct {
	// The response of list profile operation.
	ProfileListResult
}

// RelationshipLinksClientCreateOrUpdateResponse contains the response from method RelationshipLinksClient.BeginCreateOrUpdate.
type RelationshipLinksClientCreateOrUpdateResponse struct {
	// The relationship link resource format.
	RelationshipLinkResourceFormat
}

// RelationshipLinksClientDeleteResponse contains the response from method RelationshipLinksClient.BeginDelete.
type RelationshipLinksClientDeleteResponse struct {
	// placeholder for future response values
}

// RelationshipLinksClientGetResponse contains the response from method RelationshipLinksClient.Get.
type RelationshipLinksClientGetResponse struct {
	// The relationship link resource format.
	RelationshipLinkResourceFormat
}

// RelationshipLinksClientListByHubResponse contains the response from method RelationshipLinksClient.NewListByHubPager.
type RelationshipLinksClientListByHubResponse struct {
	// The response of list relationship link operation.
	RelationshipLinkListResult
}

// RelationshipsClientCreateOrUpdateResponse contains the response from method RelationshipsClient.BeginCreateOrUpdate.
type RelationshipsClientCreateOrUpdateResponse struct {
	// The relationship resource format.
	RelationshipResourceFormat
}

// RelationshipsClientDeleteResponse contains the response from method RelationshipsClient.BeginDelete.
type RelationshipsClientDeleteResponse struct {
	// placeholder for future response values
}

// RelationshipsClientGetResponse contains the response from method RelationshipsClient.Get.
type RelationshipsClientGetResponse struct {
	// The relationship resource format.
	RelationshipResourceFormat
}

// RelationshipsClientListByHubResponse contains the response from method RelationshipsClient.NewListByHubPager.
type RelationshipsClientListByHubResponse struct {
	// The response of list relationship operation.
	RelationshipListResult
}

// RoleAssignmentsClientCreateOrUpdateResponse contains the response from method RoleAssignmentsClient.BeginCreateOrUpdate.
type RoleAssignmentsClientCreateOrUpdateResponse struct {
	// The Role Assignment resource format.
	RoleAssignmentResourceFormat
}

// RoleAssignmentsClientDeleteResponse contains the response from method RoleAssignmentsClient.Delete.
type RoleAssignmentsClientDeleteResponse struct {
	// placeholder for future response values
}

// RoleAssignmentsClientGetResponse contains the response from method RoleAssignmentsClient.Get.
type RoleAssignmentsClientGetResponse struct {
	// The Role Assignment resource format.
	RoleAssignmentResourceFormat
}

// RoleAssignmentsClientListByHubResponse contains the response from method RoleAssignmentsClient.NewListByHubPager.
type RoleAssignmentsClientListByHubResponse struct {
	// The response of list role assignment operation.
	RoleAssignmentListResult
}

// RolesClientListByHubResponse contains the response from method RolesClient.NewListByHubPager.
type RolesClientListByHubResponse struct {
	// The response of list role assignment operation.
	RoleListResult
}

// ViewsClientCreateOrUpdateResponse contains the response from method ViewsClient.CreateOrUpdate.
type ViewsClientCreateOrUpdateResponse struct {
	// The view resource format.
	ViewResourceFormat
}

// ViewsClientDeleteResponse contains the response from method ViewsClient.Delete.
type ViewsClientDeleteResponse struct {
	// placeholder for future response values
}

// ViewsClientGetResponse contains the response from method ViewsClient.Get.
type ViewsClientGetResponse struct {
	// The view resource format.
	ViewResourceFormat
}

// ViewsClientListByHubResponse contains the response from method ViewsClient.NewListByHubPager.
type ViewsClientListByHubResponse struct {
	// The response of list view operation.
	ViewListResult
}

// WidgetTypesClientGetResponse contains the response from method WidgetTypesClient.Get.
type WidgetTypesClientGetResponse struct {
	// The WidgetTypeResourceFormat
	WidgetTypeResourceFormat
}

// WidgetTypesClientListByHubResponse contains the response from method WidgetTypesClient.NewListByHubPager.
type WidgetTypesClientListByHubResponse struct {
	// The response of list widget type operation.
	WidgetTypeListResult
}
