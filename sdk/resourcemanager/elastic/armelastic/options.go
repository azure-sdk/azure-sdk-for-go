// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Code generated by Microsoft (R) AutoRest Code Generator.Changes may cause incorrect behavior and will be lost if the code
// is regenerated.
// Code generated by @autorest/go. DO NOT EDIT.

package armelastic

// AllTrafficFiltersClientListOptions contains the optional parameters for the AllTrafficFiltersClient.List method.
type AllTrafficFiltersClientListOptions struct {
	// placeholder for future optional parameters
}

// AssociateTrafficFilterClientBeginAssociateOptions contains the optional parameters for the AssociateTrafficFilterClient.BeginAssociate
// method.
type AssociateTrafficFilterClientBeginAssociateOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string

	// Ruleset Id of the filter
	RulesetID *string
}

// BillingInfoClientGetOptions contains the optional parameters for the BillingInfoClient.Get method.
type BillingInfoClientGetOptions struct {
	// placeholder for future optional parameters
}

// ConnectedPartnerResourcesClientListOptions contains the optional parameters for the ConnectedPartnerResourcesClient.NewListPager
// method.
type ConnectedPartnerResourcesClientListOptions struct {
	// placeholder for future optional parameters
}

// CreateAndAssociateIPFilterClientBeginCreateOptions contains the optional parameters for the CreateAndAssociateIPFilterClient.BeginCreate
// method.
type CreateAndAssociateIPFilterClientBeginCreateOptions struct {
	// List of ips
	IPs *string

	// Name of the traffic filter
	Name *string

	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// CreateAndAssociatePLFilterClientBeginCreateOptions contains the optional parameters for the CreateAndAssociatePLFilterClient.BeginCreate
// method.
type CreateAndAssociatePLFilterClientBeginCreateOptions struct {
	// Name of the traffic filter
	Name *string

	// Guid of the private endpoint
	PrivateEndpointGUID *string

	// Name of the private endpoint
	PrivateEndpointName *string

	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// DeploymentInfoClientListOptions contains the optional parameters for the DeploymentInfoClient.List method.
type DeploymentInfoClientListOptions struct {
	// placeholder for future optional parameters
}

// DetachAndDeleteTrafficFilterClientDeleteOptions contains the optional parameters for the DetachAndDeleteTrafficFilterClient.Delete
// method.
type DetachAndDeleteTrafficFilterClientDeleteOptions struct {
	// Ruleset Id of the filter
	RulesetID *string
}

// DetachTrafficFilterClientBeginUpdateOptions contains the optional parameters for the DetachTrafficFilterClient.BeginUpdate
// method.
type DetachTrafficFilterClientBeginUpdateOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string

	// Ruleset Id of the filter
	RulesetID *string
}

// ExternalUserClientCreateOrUpdateOptions contains the optional parameters for the ExternalUserClient.CreateOrUpdate method.
type ExternalUserClientCreateOrUpdateOptions struct {
	// Elastic External User Creation Parameters
	Body *ExternalUserInfo
}

// ListAssociatedTrafficFiltersClientListOptions contains the optional parameters for the ListAssociatedTrafficFiltersClient.List
// method.
type ListAssociatedTrafficFiltersClientListOptions struct {
	// placeholder for future optional parameters
}

// MonitorClientBeginUpgradeOptions contains the optional parameters for the MonitorClient.BeginUpgrade method.
type MonitorClientBeginUpgradeOptions struct {
	// Elastic Monitor Upgrade Parameters
	Body *MonitorUpgrade

	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// MonitoredResourcesClientListOptions contains the optional parameters for the MonitoredResourcesClient.NewListPager method.
type MonitoredResourcesClientListOptions struct {
	// placeholder for future optional parameters
}

// MonitoredSubscriptionsClientBeginCreateorUpdateOptions contains the optional parameters for the MonitoredSubscriptionsClient.BeginCreateorUpdate
// method.
type MonitoredSubscriptionsClientBeginCreateorUpdateOptions struct {
	Body *MonitoredSubscriptionProperties

	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// MonitoredSubscriptionsClientBeginDeleteOptions contains the optional parameters for the MonitoredSubscriptionsClient.BeginDelete
// method.
type MonitoredSubscriptionsClientBeginDeleteOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// MonitoredSubscriptionsClientBeginUpdateOptions contains the optional parameters for the MonitoredSubscriptionsClient.BeginUpdate
// method.
type MonitoredSubscriptionsClientBeginUpdateOptions struct {
	Body *MonitoredSubscriptionProperties

	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// MonitoredSubscriptionsClientGetOptions contains the optional parameters for the MonitoredSubscriptionsClient.Get method.
type MonitoredSubscriptionsClientGetOptions struct {
	// placeholder for future optional parameters
}

// MonitoredSubscriptionsClientListOptions contains the optional parameters for the MonitoredSubscriptionsClient.NewListPager
// method.
type MonitoredSubscriptionsClientListOptions struct {
	// placeholder for future optional parameters
}

// MonitorsClientBeginCreateOptions contains the optional parameters for the MonitorsClient.BeginCreate method.
type MonitorsClientBeginCreateOptions struct {
	// Elastic monitor resource model
	Body *MonitorResource

	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// MonitorsClientBeginDeleteOptions contains the optional parameters for the MonitorsClient.BeginDelete method.
type MonitorsClientBeginDeleteOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// MonitorsClientBeginUpdateOptions contains the optional parameters for the MonitorsClient.BeginUpdate method.
type MonitorsClientBeginUpdateOptions struct {
	// Elastic resource model update parameters.
	Body *MonitorResourceUpdateParameters

	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// MonitorsClientGetOptions contains the optional parameters for the MonitorsClient.Get method.
type MonitorsClientGetOptions struct {
	// placeholder for future optional parameters
}

// MonitorsClientListByResourceGroupOptions contains the optional parameters for the MonitorsClient.NewListByResourceGroupPager
// method.
type MonitorsClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// MonitorsClientListOptions contains the optional parameters for the MonitorsClient.NewListPager method.
type MonitorsClientListOptions struct {
	// placeholder for future optional parameters
}

// OpenAIClientCreateOrUpdateOptions contains the optional parameters for the OpenAIClient.CreateOrUpdate method.
type OpenAIClientCreateOrUpdateOptions struct {
	Body *OpenAIIntegrationRPModel
}

// OpenAIClientDeleteOptions contains the optional parameters for the OpenAIClient.Delete method.
type OpenAIClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// OpenAIClientGetOptions contains the optional parameters for the OpenAIClient.Get method.
type OpenAIClientGetOptions struct {
	// placeholder for future optional parameters
}

// OpenAIClientGetStatusOptions contains the optional parameters for the OpenAIClient.GetStatus method.
type OpenAIClientGetStatusOptions struct {
	// placeholder for future optional parameters
}

// OpenAIClientListOptions contains the optional parameters for the OpenAIClient.NewListPager method.
type OpenAIClientListOptions struct {
	// placeholder for future optional parameters
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.NewListPager method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// OrganizationsClientBeginResubscribeOptions contains the optional parameters for the OrganizationsClient.BeginResubscribe
// method.
type OrganizationsClientBeginResubscribeOptions struct {
	// Resubscribe Properties
	Body *ResubscribeProperties

	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// OrganizationsClientGetAPIKeyOptions contains the optional parameters for the OrganizationsClient.GetAPIKey method.
type OrganizationsClientGetAPIKeyOptions struct {
	// Email Id parameter of the User Organization, of which the API Key must be returned
	Body *UserEmailID
}

// OrganizationsClientGetElasticToAzureSubscriptionMappingOptions contains the optional parameters for the OrganizationsClient.GetElasticToAzureSubscriptionMapping
// method.
type OrganizationsClientGetElasticToAzureSubscriptionMappingOptions struct {
	// placeholder for future optional parameters
}

// TagRulesClientBeginDeleteOptions contains the optional parameters for the TagRulesClient.BeginDelete method.
type TagRulesClientBeginDeleteOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// TagRulesClientCreateOrUpdateOptions contains the optional parameters for the TagRulesClient.CreateOrUpdate method.
type TagRulesClientCreateOrUpdateOptions struct {
	// request body of MonitoringTagRules
	Body *MonitoringTagRules
}

// TagRulesClientGetOptions contains the optional parameters for the TagRulesClient.Get method.
type TagRulesClientGetOptions struct {
	// placeholder for future optional parameters
}

// TagRulesClientListOptions contains the optional parameters for the TagRulesClient.NewListPager method.
type TagRulesClientListOptions struct {
	// placeholder for future optional parameters
}

// TrafficFiltersClientDeleteOptions contains the optional parameters for the TrafficFiltersClient.Delete method.
type TrafficFiltersClientDeleteOptions struct {
	// Ruleset Id of the filter
	RulesetID *string
}

// UpgradableVersionsClientDetailsOptions contains the optional parameters for the UpgradableVersionsClient.Details method.
type UpgradableVersionsClientDetailsOptions struct {
	// placeholder for future optional parameters
}

// VMCollectionClientUpdateOptions contains the optional parameters for the VMCollectionClient.Update method.
type VMCollectionClientUpdateOptions struct {
	// VM resource Id
	Body *VMCollectionUpdate
}

// VMHostClientListOptions contains the optional parameters for the VMHostClient.NewListPager method.
type VMHostClientListOptions struct {
	// placeholder for future optional parameters
}

// VMIngestionClientDetailsOptions contains the optional parameters for the VMIngestionClient.Details method.
type VMIngestionClientDetailsOptions struct {
	// placeholder for future optional parameters
}

// VersionsClientListOptions contains the optional parameters for the VersionsClient.NewListPager method.
type VersionsClientListOptions struct {
	// placeholder for future optional parameters
}
