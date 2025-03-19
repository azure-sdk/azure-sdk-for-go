// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armfeatures

import "time"

// AuthorizationProfile - Authorization Profile
type AuthorizationProfile struct {
	// READ-ONLY; The approved time
	ApprovedTime *time.Time

	// READ-ONLY; The approver
	Approver *string

	// READ-ONLY; The requested time
	RequestedTime *time.Time

	// READ-ONLY; The requester
	Requester *string

	// READ-ONLY; The requester object id
	RequesterObjectID *string
}

// ErrorDefinition - Error definition.
type ErrorDefinition struct {
	// Internal error details.
	Details []*ErrorDefinition

	// READ-ONLY; Service specific error code which serves as the substatus for the HTTP error code.
	Code *string

	// READ-ONLY; Description of the error.
	Message *string
}

// ErrorResponse - Error response indicates that the service is not able to process the incoming request.
type ErrorResponse struct {
	// The error details.
	Error *ErrorDefinition
}

// FeatureOperationsListResult - List of previewed features.
type FeatureOperationsListResult struct {
	// The URL to use for getting the next set of results.
	NextLink *string

	// The array of features.
	Value []*FeatureResult
}

// FeatureProperties - Information about feature.
type FeatureProperties struct {
	// The registration state of the feature for the subscription.
	State *string
}

// FeatureResult - Previewed feature information.
type FeatureResult struct {
	// The resource ID of the feature.
	ID *string

	// The name of the feature.
	Name *string

	// Properties of the previewed feature.
	Properties *FeatureProperties

	// The resource type of the feature.
	Type *string
}

// Operation - Microsoft.Features operation
type Operation struct {
	// The object that represents the operation.
	Display *OperationDisplay

	// Operation name: {provider}/{resource}/{operation}
	Name *string
}

// OperationDisplay - The object that represents the operation.
type OperationDisplay struct {
	// Operation type: Read, write, delete, etc.
	Operation *string

	// Service provider: Microsoft.Features
	Provider *string

	// Resource on which the operation is performed: Profile, endpoint, etc.
	Resource *string
}

// OperationListResult - Result of the request to list Microsoft.Features operations. It contains a list of operations and
// a URL link to get the next set of results.
type OperationListResult struct {
	// URL to get the next set of operation list results if there are any.
	NextLink *string

	// List of Microsoft.Features operations.
	Value []*Operation
}

// ProxyResource - An Azure proxy resource.
type ProxyResource struct {
	// READ-ONLY; Azure resource Id.
	ID *string

	// READ-ONLY; Azure resource name.
	Name *string

	// READ-ONLY; Azure resource type.
	Type *string
}

// SubscriptionFeatureRegistration - Subscription feature registration details
type SubscriptionFeatureRegistration struct {
	Properties *SubscriptionFeatureRegistrationProperties

	// READ-ONLY; Azure resource Id.
	ID *string

	// READ-ONLY; Azure resource name.
	Name *string

	// READ-ONLY; Azure resource type.
	Type *string
}

// SubscriptionFeatureRegistrationList - The list of subscription feature registrations.
type SubscriptionFeatureRegistrationList struct {
	// The link used to get the next page of subscription feature registrations list.
	NextLink *string

	// The list of subscription feature registrations.
	Value []*SubscriptionFeatureRegistration
}

type SubscriptionFeatureRegistrationProperties struct {
	// Authorization Profile
	AuthorizationProfile *AuthorizationProfile

	// The feature description.
	Description *string

	// Key-value pairs for meta data.
	Metadata map[string]*string

	// Indicates whether feature should be displayed in Portal.
	ShouldFeatureDisplayInPortal *bool

	// The state.
	State *SubscriptionFeatureRegistrationState

	// READ-ONLY; The feature approval type.
	ApprovalType *SubscriptionFeatureRegistrationApprovalType

	// READ-ONLY; The featureDisplayName.
	DisplayName *string

	// READ-ONLY; The feature documentation link.
	DocumentationLink *string

	// READ-ONLY; The featureName.
	FeatureName *string

	// READ-ONLY; The providerNamespace.
	ProviderNamespace *string

	// READ-ONLY; The feature registration date.
	RegistrationDate *time.Time

	// READ-ONLY; The feature release date.
	ReleaseDate *time.Time

	// READ-ONLY; The subscriptionId.
	SubscriptionID *string

	// READ-ONLY; The tenantId.
	TenantID *string
}
