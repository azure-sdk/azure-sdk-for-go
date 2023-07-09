//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armalertsmanagement

// AlertProcessingRulesClientCreateOrUpdateResponse contains the response from method AlertProcessingRulesClient.CreateOrUpdate.
type AlertProcessingRulesClientCreateOrUpdateResponse struct {
	AlertProcessingRule
	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// AlertProcessingRulesClientDeleteResponse contains the response from method AlertProcessingRulesClient.Delete.
type AlertProcessingRulesClientDeleteResponse struct {
	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// AlertProcessingRulesClientGetByNameResponse contains the response from method AlertProcessingRulesClient.GetByName.
type AlertProcessingRulesClientGetByNameResponse struct {
	AlertProcessingRule
	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// AlertProcessingRulesClientListByResourceGroupResponse contains the response from method AlertProcessingRulesClient.NewListByResourceGroupPager.
type AlertProcessingRulesClientListByResourceGroupResponse struct {
	AlertProcessingRulesList
	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// AlertProcessingRulesClientListBySubscriptionResponse contains the response from method AlertProcessingRulesClient.NewListBySubscriptionPager.
type AlertProcessingRulesClientListBySubscriptionResponse struct {
	AlertProcessingRulesList
	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// AlertProcessingRulesClientUpdateResponse contains the response from method AlertProcessingRulesClient.Update.
type AlertProcessingRulesClientUpdateResponse struct {
	AlertProcessingRule
	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// AlertRuleRecommendationsClientListByResourceResponse contains the response from method AlertRuleRecommendationsClient.NewListByResourcePager.
type AlertRuleRecommendationsClientListByResourceResponse struct {
	AlertRuleRecommendationsListResponse
}

// AlertRuleRecommendationsClientListByTargetTypeResponse contains the response from method AlertRuleRecommendationsClient.NewListByTargetTypePager.
type AlertRuleRecommendationsClientListByTargetTypeResponse struct {
	AlertRuleRecommendationsListResponse
}

// AlertsClientChangeStateResponse contains the response from method AlertsClient.ChangeState.
type AlertsClientChangeStateResponse struct {
	Alert
}

// AlertsClientGetAllResponse contains the response from method AlertsClient.NewGetAllPager.
type AlertsClientGetAllResponse struct {
	AlertsList
}

// AlertsClientGetByIDResponse contains the response from method AlertsClient.GetByID.
type AlertsClientGetByIDResponse struct {
	Alert
}

// AlertsClientGetHistoryResponse contains the response from method AlertsClient.GetHistory.
type AlertsClientGetHistoryResponse struct {
	AlertModification
}

// AlertsClientGetSummaryResponse contains the response from method AlertsClient.GetSummary.
type AlertsClientGetSummaryResponse struct {
	AlertsSummary
}

// AlertsClientMetaDataResponse contains the response from method AlertsClient.MetaData.
type AlertsClientMetaDataResponse struct {
	AlertsMetaData
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	OperationsList
}

// PrometheusRuleGroupsClientCreateOrUpdateResponse contains the response from method PrometheusRuleGroupsClient.CreateOrUpdate.
type PrometheusRuleGroupsClientCreateOrUpdateResponse struct {
	PrometheusRuleGroupResource
}

// PrometheusRuleGroupsClientDeleteResponse contains the response from method PrometheusRuleGroupsClient.Delete.
type PrometheusRuleGroupsClientDeleteResponse struct {
	// placeholder for future response values
}

// PrometheusRuleGroupsClientGetResponse contains the response from method PrometheusRuleGroupsClient.Get.
type PrometheusRuleGroupsClientGetResponse struct {
	PrometheusRuleGroupResource
}

// PrometheusRuleGroupsClientListByResourceGroupResponse contains the response from method PrometheusRuleGroupsClient.NewListByResourceGroupPager.
type PrometheusRuleGroupsClientListByResourceGroupResponse struct {
	PrometheusRuleGroupResourceCollection
}

// PrometheusRuleGroupsClientListBySubscriptionResponse contains the response from method PrometheusRuleGroupsClient.NewListBySubscriptionPager.
type PrometheusRuleGroupsClientListBySubscriptionResponse struct {
	PrometheusRuleGroupResourceCollection
}

// PrometheusRuleGroupsClientUpdateResponse contains the response from method PrometheusRuleGroupsClient.Update.
type PrometheusRuleGroupsClientUpdateResponse struct {
	PrometheusRuleGroupResource
}

// SmartGroupsClientChangeStateResponse contains the response from method SmartGroupsClient.ChangeState.
type SmartGroupsClientChangeStateResponse struct {
	SmartGroup
	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// SmartGroupsClientGetAllResponse contains the response from method SmartGroupsClient.NewGetAllPager.
type SmartGroupsClientGetAllResponse struct {
	SmartGroupsList
}

// SmartGroupsClientGetByIDResponse contains the response from method SmartGroupsClient.GetByID.
type SmartGroupsClientGetByIDResponse struct {
	SmartGroup
	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// SmartGroupsClientGetHistoryResponse contains the response from method SmartGroupsClient.GetHistory.
type SmartGroupsClientGetHistoryResponse struct {
	SmartGroupModification
}
