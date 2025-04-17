// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armproviderhub

import "time"

type AuthorizationActionMapping struct {
	Desired  *string
	Original *string
}

type CanaryTrafficRegionRolloutConfiguration struct {
	Regions     []*string
	SkipRegions []*string
}

type CheckNameAvailabilitySpecifications struct {
	EnableDefaultValidation           *bool
	ResourceTypesWithCustomValidation []*string
}

type CheckinManifestInfo struct {
	// REQUIRED
	IsCheckedIn *bool

	// REQUIRED
	StatusMessage *string
	CommitID      *string
	PullRequest   *string
}

type CheckinManifestParams struct {
	// REQUIRED; The baseline ARM manifest location supplied to the checkin manifest operation.
	BaselineArmManifestLocation *string

	// REQUIRED; The environment supplied to the checkin manifest operation.
	Environment *string
}

// CustomRollout - Rollout details.
type CustomRollout struct {
	// REQUIRED; Properties of the rollout.
	Properties *CustomRolloutProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

type CustomRolloutArrayResponseWithContinuation struct {
	// The URL to get to the next set of results, if there are any.
	NextLink *string
	Value    []*CustomRollout
}

// CustomRolloutProperties - Properties of the rollout.
type CustomRolloutProperties struct {
	// REQUIRED
	Specification     *CustomRolloutPropertiesSpecification
	ProvisioningState *ProvisioningState
	Status            *CustomRolloutPropertiesStatus
}

type CustomRolloutPropertiesAutoGenerated struct {
	// REQUIRED
	Specification     *CustomRolloutPropertiesSpecification
	ProvisioningState *ProvisioningState
	Status            *CustomRolloutPropertiesStatus
}

type CustomRolloutPropertiesSpecification struct {
	// REQUIRED
	Canary                    *CustomRolloutSpecificationCanary
	ProviderRegistration      *CustomRolloutSpecificationProviderRegistration
	ResourceTypeRegistrations []*ResourceTypeRegistration
}

type CustomRolloutPropertiesStatus struct {
	CompletedRegions []*string

	// Dictionary of
	FailedOrSkippedRegions map[string]*ExtendedErrorInfo
}

type CustomRolloutSpecification struct {
	// REQUIRED
	Canary                    *CustomRolloutSpecificationCanary
	ProviderRegistration      *CustomRolloutSpecificationProviderRegistration
	ResourceTypeRegistrations []*ResourceTypeRegistration
}

type CustomRolloutSpecificationCanary struct {
	Regions []*string
}

type CustomRolloutSpecificationProviderRegistration struct {
	Properties *ProviderRegistrationProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

type CustomRolloutStatus struct {
	CompletedRegions []*string

	// Dictionary of
	FailedOrSkippedRegions map[string]*ExtendedErrorInfo
}

// DefaultRollout - Default rollout definition.
type DefaultRollout struct {
	// Properties of the rollout.
	Properties *DefaultRolloutProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

type DefaultRolloutArrayResponseWithContinuation struct {
	// The URL to get to the next set of results, if there are any.
	NextLink *string
	Value    []*DefaultRollout
}

// DefaultRolloutProperties - Properties of the rollout.
type DefaultRolloutProperties struct {
	ProvisioningState *ProvisioningState
	Specification     *DefaultRolloutPropertiesSpecification
	Status            *DefaultRolloutPropertiesStatus
}

type DefaultRolloutPropertiesAutoGenerated struct {
	ProvisioningState *ProvisioningState
	Specification     *DefaultRolloutPropertiesSpecification
	Status            *DefaultRolloutPropertiesStatus
}

type DefaultRolloutPropertiesSpecification struct {
	Canary                    *DefaultRolloutSpecificationCanary
	HighTraffic               *DefaultRolloutSpecificationHighTraffic
	LowTraffic                *DefaultRolloutSpecificationLowTraffic
	MediumTraffic             *DefaultRolloutSpecificationMediumTraffic
	ProviderRegistration      *DefaultRolloutSpecificationProviderRegistration
	ResourceTypeRegistrations []*ResourceTypeRegistration
	RestOfTheWorldGroupOne    *DefaultRolloutSpecificationRestOfTheWorldGroupOne
	RestOfTheWorldGroupTwo    *DefaultRolloutSpecificationRestOfTheWorldGroupTwo
}

type DefaultRolloutPropertiesStatus struct {
	CompletedRegions []*string

	// Dictionary of
	FailedOrSkippedRegions           map[string]*ExtendedErrorInfo
	NextTrafficRegion                *TrafficRegionCategory
	NextTrafficRegionScheduledTime   *time.Time
	SubscriptionReregistrationResult *SubscriptionReregistrationResult
}

type DefaultRolloutSpecification struct {
	Canary                    *DefaultRolloutSpecificationCanary
	HighTraffic               *DefaultRolloutSpecificationHighTraffic
	LowTraffic                *DefaultRolloutSpecificationLowTraffic
	MediumTraffic             *DefaultRolloutSpecificationMediumTraffic
	ProviderRegistration      *DefaultRolloutSpecificationProviderRegistration
	ResourceTypeRegistrations []*ResourceTypeRegistration
	RestOfTheWorldGroupOne    *DefaultRolloutSpecificationRestOfTheWorldGroupOne
	RestOfTheWorldGroupTwo    *DefaultRolloutSpecificationRestOfTheWorldGroupTwo
}

type DefaultRolloutSpecificationCanary struct {
	Regions     []*string
	SkipRegions []*string
}

type DefaultRolloutSpecificationHighTraffic struct {
	Regions      []*string
	WaitDuration *string
}

type DefaultRolloutSpecificationLowTraffic struct {
	Regions      []*string
	WaitDuration *string
}

type DefaultRolloutSpecificationMediumTraffic struct {
	Regions      []*string
	WaitDuration *string
}

type DefaultRolloutSpecificationProviderRegistration struct {
	Properties *ProviderRegistrationProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

type DefaultRolloutSpecificationRestOfTheWorldGroupOne struct {
	Regions      []*string
	WaitDuration *string
}

type DefaultRolloutSpecificationRestOfTheWorldGroupTwo struct {
	Regions      []*string
	WaitDuration *string
}

type DefaultRolloutStatus struct {
	CompletedRegions []*string

	// Dictionary of
	FailedOrSkippedRegions           map[string]*ExtendedErrorInfo
	NextTrafficRegion                *TrafficRegionCategory
	NextTrafficRegionScheduledTime   *time.Time
	SubscriptionReregistrationResult *SubscriptionReregistrationResult
}

// Error - Standard error object.
type Error struct {
	// READ-ONLY; Server-defined set of error codes.
	Code *string

	// READ-ONLY; Array of details about specific errors that led to this reported error.
	Details []*Error

	// READ-ONLY; Object containing more specific information than the current object about the error.
	InnerError *ErrorInnerError

	// READ-ONLY; Human-readable representation of the error.
	Message *string

	// READ-ONLY; Target of the error.
	Target *string
}

// ErrorInnerError - Object containing more specific information than the current object about the error.
type ErrorInnerError struct {
	// OPTIONAL; Contains additional key/value pairs not defined in the schema.
	AdditionalProperties map[string]any

	// READ-ONLY; Specific error code than was provided by the containing error.
	Code *string

	// READ-ONLY; Object containing more specific information than the current object about the error.
	InnerError any
}

// ErrorResponse - Standard error response.
type ErrorResponse struct {
	// Standard error object.
	Error *ErrorResponseError
}

// ErrorResponseError - Standard error object.
type ErrorResponseError struct {
	// READ-ONLY; Server-defined set of error codes.
	Code *string

	// READ-ONLY; Array of details about specific errors that led to this reported error.
	Details []*Error

	// READ-ONLY; Object containing more specific information than the current object about the error.
	InnerError *ErrorInnerError

	// READ-ONLY; Human-readable representation of the error.
	Message *string

	// READ-ONLY; Target of the error.
	Target *string
}

type ExtendedErrorInfo struct {
	AdditionalInfo []*TypedErrorInfo
	Code           *string
	Details        []*ExtendedErrorInfo
	Message        *string
	Target         *string
}

type ExtendedLocationOptions struct {
	SupportedPolicy *string
	Type            *string
}

type ExtensionOptions struct {
	Request  []*ExtensionOptionType
	Response []*ExtensionOptionType
}

type FeaturesRule struct {
	// REQUIRED
	RequiredFeaturesPolicy *FeaturesPolicy
}

type IdentityManagement struct {
	Type *IdentityManagementTypes
}

type IdentityManagementProperties struct {
	ApplicationID *string
	Type          *IdentityManagementTypes
}

// InnerError - Inner error containing list of errors.
type InnerError struct {
	// OPTIONAL; Contains additional key/value pairs not defined in the schema.
	AdditionalProperties map[string]any

	// READ-ONLY; Specific error code than was provided by the containing error.
	Code *string

	// READ-ONLY; Object containing more specific information than the current object about the error.
	InnerError any
}

type LightHouseAuthorization struct {
	// REQUIRED
	PrincipalID *string

	// REQUIRED
	RoleDefinitionID *string
}

type LinkedAccessCheck struct {
	ActionName       *string
	LinkedAction     *string
	LinkedActionVerb *string
	LinkedProperty   *string
	LinkedType       *string
}

type LinkedOperationRule struct {
	// REQUIRED
	LinkedAction *LinkedAction

	// REQUIRED
	LinkedOperation *LinkedOperation
}

type LoggingHiddenPropertyPath struct {
	HiddenPathsOnRequest  []*string
	HiddenPathsOnResponse []*string
}

type LoggingRule struct {
	// REQUIRED
	Action *string

	// REQUIRED
	DetailLevel *LoggingDetails

	// REQUIRED
	Direction           *LoggingDirections
	HiddenPropertyPaths *LoggingRuleHiddenPropertyPaths
}

type LoggingRuleHiddenPropertyPaths struct {
	HiddenPathsOnRequest  []*string
	HiddenPathsOnResponse []*string
}

type Metadata struct {
	ProviderAuthentication          *MetadataProviderAuthentication
	ProviderAuthorizations          []*ResourceProviderAuthorization
	ThirdPartyProviderAuthorization *MetadataThirdPartyProviderAuthorization
}

type MetadataProviderAuthentication struct {
	// REQUIRED
	AllowedAudiences []*string
}

type MetadataThirdPartyProviderAuthorization struct {
	Authorizations    []*LightHouseAuthorization
	ManagedByTenantID *string
}

type NotificationEndpoint struct {
	Locations               []*string
	NotificationDestination *string
}

// NotificationRegistration - The notification registration definition.
type NotificationRegistration struct {
	Properties *NotificationRegistrationProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

type NotificationRegistrationArrayResponseWithContinuation struct {
	// The URL to get to the next set of results, if there are any.
	NextLink *string
	Value    []*NotificationRegistration
}

type NotificationRegistrationProperties struct {
	IncludedEvents        []*string
	MessageScope          *MessageScope
	NotificationEndpoints []*NotificationEndpoint
	NotificationMode      *NotificationMode
	ProvisioningState     *ProvisioningState
}

type NotificationRegistrationPropertiesAutoGenerated struct {
	IncludedEvents        []*string
	MessageScope          *MessageScope
	NotificationEndpoints []*NotificationEndpoint
	NotificationMode      *NotificationMode
	ProvisioningState     *ProvisioningState
}

type OperationsContent struct {
	// Operations content.
	Properties *OperationsDefinition

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// OperationsDefinition - Properties of an Operation.
type OperationsDefinition struct {
	// REQUIRED; Display information of the operation.
	Display *OperationsDefinitionDisplay

	// REQUIRED; Name of the operation.
	Name       *string
	ActionType *OperationActionType

	// Indicates whether the operation applies to data-plane.
	IsDataAction *bool
	Origin       *OperationOrigins

	// Anything
	Properties any
}

type OperationsDefinitionArrayResponseWithContinuation struct {
	// The URL to get to the next set of results, if there are any.
	NextLink *string
	Value    []*OperationsDefinition
}

// OperationsDefinitionDisplay - Display information of the operation.
type OperationsDefinitionDisplay struct {
	// REQUIRED
	Description *string

	// REQUIRED
	Operation *string

	// REQUIRED
	Provider *string

	// REQUIRED
	Resource *string
}

type OperationsDisplayDefinition struct {
	// REQUIRED
	Description *string

	// REQUIRED
	Operation *string

	// REQUIRED
	Provider *string

	// REQUIRED
	Resource *string
}

type OperationsPutContent struct {
	// REQUIRED
	Contents []*OperationsDefinition
}

type ProviderRegistration struct {
	Properties *ProviderRegistrationProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

type ProviderRegistrationArrayResponseWithContinuation struct {
	// The URL to get to the next set of results, if there are any.
	NextLink *string
	Value    []*ProviderRegistration
}

type ProviderRegistrationProperties struct {
	Capabilities []*ResourceProviderCapabilities
	FeaturesRule *ResourceProviderManifestPropertiesFeaturesRule
	Management   *ResourceProviderManifestPropertiesManagement

	// Anything
	Metadata                                        any
	Namespace                                       *string
	ProviderAuthentication                          *ResourceProviderManifestPropertiesProviderAuthentication
	ProviderAuthorizations                          []*ResourceProviderAuthorization
	ProviderHubMetadata                             *ProviderRegistrationPropertiesProviderHubMetadata
	ProviderType                                    *ResourceProviderType
	ProviderVersion                                 *string
	ProvisioningState                               *ProvisioningState
	RequestHeaderOptions                            *ResourceProviderManifestPropertiesRequestHeaderOptions
	RequiredFeatures                                []*string
	SubscriptionLifecycleNotificationSpecifications *ProviderRegistrationPropertiesSubscriptionLifecycleNotificationSpecifications
	TemplateDeploymentOptions                       *ResourceProviderManifestPropertiesTemplateDeploymentOptions
}

type ProviderRegistrationPropertiesAutoGenerated struct {
	Capabilities []*ResourceProviderCapabilities
	FeaturesRule *ResourceProviderManifestPropertiesFeaturesRule
	Management   *ResourceProviderManifestPropertiesManagement

	// Anything
	Metadata                                        any
	Namespace                                       *string
	ProviderAuthentication                          *ResourceProviderManifestPropertiesProviderAuthentication
	ProviderAuthorizations                          []*ResourceProviderAuthorization
	ProviderHubMetadata                             *ProviderRegistrationPropertiesProviderHubMetadata
	ProviderType                                    *ResourceProviderType
	ProviderVersion                                 *string
	ProvisioningState                               *ProvisioningState
	RequestHeaderOptions                            *ResourceProviderManifestPropertiesRequestHeaderOptions
	RequiredFeatures                                []*string
	SubscriptionLifecycleNotificationSpecifications *ProviderRegistrationPropertiesSubscriptionLifecycleNotificationSpecifications
	TemplateDeploymentOptions                       *ResourceProviderManifestPropertiesTemplateDeploymentOptions
}

type ProviderRegistrationPropertiesProviderHubMetadata struct {
	ProviderAuthentication          *MetadataProviderAuthentication
	ProviderAuthorizations          []*ResourceProviderAuthorization
	ThirdPartyProviderAuthorization *MetadataThirdPartyProviderAuthorization
}

type ProviderRegistrationPropertiesSubscriptionLifecycleNotificationSpecifications struct {
	SoftDeleteTTL                    *string
	SubscriptionStateOverrideActions []*SubscriptionStateOverrideAction
}

// ProxyResource - The resource model definition for a Azure Resource Manager proxy resource. It will not have tags and a
// location
type ProxyResource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

type ReRegisterSubscriptionMetadata struct {
	// REQUIRED
	Enabled          *bool
	ConcurrencyLimit *int32
}

type RequestHeaderOptions struct {
	OptInHeaders *OptInHeaderType
}

// Resource - Common fields that are returned in the response for all Azure Resource Manager resources
type Resource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

type ResourceMovePolicy struct {
	CrossResourceGroupMoveEnabled *bool
	CrossSubscriptionMoveEnabled  *bool
	ValidationRequired            *bool
}

type ResourceProviderAuthentication struct {
	// REQUIRED
	AllowedAudiences []*string
}

type ResourceProviderAuthorization struct {
	ApplicationID             *string
	ManagedByRoleDefinitionID *string
	RoleDefinitionID          *string
}

type ResourceProviderCapabilities struct {
	// REQUIRED
	Effect *ResourceProviderCapabilitiesEffect

	// REQUIRED
	QuotaID          *string
	RequiredFeatures []*string
}

type ResourceProviderEndpoint struct {
	APIVersions      []*string
	Enabled          *bool
	EndpointURI      *string
	FeaturesRule     *ResourceProviderEndpointFeaturesRule
	Locations        []*string
	RequiredFeatures []*string
	Timeout          *string
}

type ResourceProviderEndpointFeaturesRule struct {
	// REQUIRED
	RequiredFeaturesPolicy *FeaturesPolicy
}

type ResourceProviderManagement struct {
	IncidentContactEmail   *string
	IncidentRoutingService *string
	IncidentRoutingTeam    *string
	ManifestOwners         []*string
	ResourceAccessPolicy   *ResourceAccessPolicy
	ResourceAccessRoles    []any
	SchemaOwners           []*string
	ServiceTreeInfos       []*ServiceTreeInfo
}

type ResourceProviderManifest struct {
	Capabilities                []*ResourceProviderCapabilities
	FeaturesRule                *ResourceProviderManifestFeaturesRule
	GlobalNotificationEndpoints []*ResourceProviderEndpoint
	Management                  *ResourceProviderManifestManagement

	// Anything
	Metadata                       any
	Namespace                      *string
	ProviderAuthentication         *ResourceProviderManifestProviderAuthentication
	ProviderAuthorizations         []*ResourceProviderAuthorization
	ProviderType                   *ResourceProviderType
	ProviderVersion                *string
	ReRegisterSubscriptionMetadata *ResourceProviderManifestReRegisterSubscriptionMetadata
	RequestHeaderOptions           *ResourceProviderManifestRequestHeaderOptions
	RequiredFeatures               []*string
	ResourceTypes                  []*ResourceType
}

type ResourceProviderManifestFeaturesRule struct {
	// REQUIRED
	RequiredFeaturesPolicy *FeaturesPolicy
}

type ResourceProviderManifestManagement struct {
	IncidentContactEmail   *string
	IncidentRoutingService *string
	IncidentRoutingTeam    *string
	ManifestOwners         []*string
	ResourceAccessPolicy   *ResourceAccessPolicy
	ResourceAccessRoles    []any
	SchemaOwners           []*string
	ServiceTreeInfos       []*ServiceTreeInfo
}

type ResourceProviderManifestProperties struct {
	Capabilities []*ResourceProviderCapabilities
	FeaturesRule *ResourceProviderManifestPropertiesFeaturesRule
	Management   *ResourceProviderManifestPropertiesManagement

	// Anything
	Metadata                  any
	Namespace                 *string
	ProviderAuthentication    *ResourceProviderManifestPropertiesProviderAuthentication
	ProviderAuthorizations    []*ResourceProviderAuthorization
	ProviderType              *ResourceProviderType
	ProviderVersion           *string
	RequestHeaderOptions      *ResourceProviderManifestPropertiesRequestHeaderOptions
	RequiredFeatures          []*string
	TemplateDeploymentOptions *ResourceProviderManifestPropertiesTemplateDeploymentOptions
}

type ResourceProviderManifestPropertiesFeaturesRule struct {
	// REQUIRED
	RequiredFeaturesPolicy *FeaturesPolicy
}

type ResourceProviderManifestPropertiesManagement struct {
	IncidentContactEmail   *string
	IncidentRoutingService *string
	IncidentRoutingTeam    *string
	ManifestOwners         []*string
	ResourceAccessPolicy   *ResourceAccessPolicy
	ResourceAccessRoles    []any
	SchemaOwners           []*string
	ServiceTreeInfos       []*ServiceTreeInfo
}

type ResourceProviderManifestPropertiesProviderAuthentication struct {
	// REQUIRED
	AllowedAudiences []*string
}

type ResourceProviderManifestPropertiesRequestHeaderOptions struct {
	OptInHeaders *OptInHeaderType
}

type ResourceProviderManifestPropertiesTemplateDeploymentOptions struct {
	PreflightOptions   []*PreflightOption
	PreflightSupported *bool
}

type ResourceProviderManifestProviderAuthentication struct {
	// REQUIRED
	AllowedAudiences []*string
}

type ResourceProviderManifestReRegisterSubscriptionMetadata struct {
	// REQUIRED
	Enabled          *bool
	ConcurrencyLimit *int32
}

type ResourceProviderManifestRequestHeaderOptions struct {
	OptInHeaders *OptInHeaderType
}

type ResourceType struct {
	AllowedUnauthorizedActions  []*string
	AuthorizationActionMappings []*AuthorizationActionMapping
	DefaultAPIVersion           *string
	DisallowedActionVerbs       []*string
	Endpoints                   []*ResourceProviderEndpoint
	ExtendedLocations           []*ExtendedLocationOptions
	FeaturesRule                *ResourceTypeFeaturesRule
	IdentityManagement          *ResourceTypeIdentityManagement
	LinkedAccessChecks          []*LinkedAccessCheck
	LinkedOperationRules        []*LinkedOperationRule
	LoggingRules                []*LoggingRule
	MarketplaceType             *MarketplaceType

	// Anything
	Metadata                 any
	Name                     *string
	RequestHeaderOptions     *ResourceTypeRequestHeaderOptions
	RequiredFeatures         []*string
	ResourceDeletionPolicy   *ManifestResourceDeletionPolicy
	ResourceValidation       *ResourceValidation
	RoutingType              *RoutingType
	SKULink                  *string
	ServiceTreeInfos         []*ServiceTreeInfo
	SubscriptionStateRules   []*SubscriptionStateRule
	TemplateDeploymentPolicy *ResourceTypeTemplateDeploymentPolicy
	ThrottlingRules          []*ThrottlingRule
}

type ResourceTypeEndpoint struct {
	APIVersions      []*string
	Enabled          *bool
	Extensions       []*ResourceTypeExtension
	FeaturesRule     *ResourceTypeEndpointFeaturesRule
	Locations        []*string
	RequiredFeatures []*string
	Timeout          *string
}

type ResourceTypeEndpointFeaturesRule struct {
	// REQUIRED
	RequiredFeaturesPolicy *FeaturesPolicy
}

type ResourceTypeExtension struct {
	EndpointURI         *string
	ExtensionCategories []*ExtensionCategory
	Timeout             *string
}

type ResourceTypeExtensionOptions struct {
	ResourceCreationBegin *ResourceTypeExtensionOptionsResourceCreationBegin
}

type ResourceTypeExtensionOptionsResourceCreationBegin struct {
	Request  []*ExtensionOptionType
	Response []*ExtensionOptionType
}

type ResourceTypeFeaturesRule struct {
	// REQUIRED
	RequiredFeaturesPolicy *FeaturesPolicy
}

type ResourceTypeIdentityManagement struct {
	Type *IdentityManagementTypes
}

type ResourceTypeRegistration struct {
	Properties *ResourceTypeRegistrationProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

type ResourceTypeRegistrationArrayResponseWithContinuation struct {
	// The URL to get to the next set of results, if there are any.
	NextLink *string
	Value    []*ResourceTypeRegistration
}

type ResourceTypeRegistrationProperties struct {
	AllowedUnauthorizedActions                      []*string
	AuthorizationActionMappings                     []*AuthorizationActionMapping
	CheckNameAvailabilitySpecifications             *ResourceTypeRegistrationPropertiesCheckNameAvailabilitySpecifications
	DefaultAPIVersion                               *string
	DisallowedActionVerbs                           []*string
	EnableAsyncOperation                            *bool
	EnableThirdPartyS2S                             *bool
	Endpoints                                       []*ResourceTypeEndpoint
	ExtendedLocations                               []*ExtendedLocationOptions
	ExtensionOptions                                *ResourceTypeRegistrationPropertiesExtensionOptions
	FeaturesRule                                    *ResourceTypeRegistrationPropertiesFeaturesRule
	IdentityManagement                              *ResourceTypeRegistrationPropertiesIdentityManagement
	IsPureProxy                                     *bool
	LinkedAccessChecks                              []*LinkedAccessCheck
	LoggingRules                                    []*LoggingRule
	MarketplaceType                                 *MarketplaceType
	ProvisioningState                               *ProvisioningState
	Regionality                                     *Regionality
	RequestHeaderOptions                            *ResourceTypeRegistrationPropertiesRequestHeaderOptions
	RequiredFeatures                                []*string
	ResourceDeletionPolicy                          *ResourceDeletionPolicy
	ResourceMovePolicy                              *ResourceTypeRegistrationPropertiesResourceMovePolicy
	RoutingType                                     *RoutingType
	ServiceTreeInfos                                []*ServiceTreeInfo
	SubscriptionLifecycleNotificationSpecifications *ResourceTypeRegistrationPropertiesSubscriptionLifecycleNotificationSpecifications
	SubscriptionStateRules                          []*SubscriptionStateRule
	SwaggerSpecifications                           []*SwaggerSpecification
	TemplateDeploymentOptions                       *ResourceTypeRegistrationPropertiesTemplateDeploymentOptions
	ThrottlingRules                                 []*ThrottlingRule
}

type ResourceTypeRegistrationPropertiesAutoGenerated struct {
	AllowedUnauthorizedActions                      []*string
	AuthorizationActionMappings                     []*AuthorizationActionMapping
	CheckNameAvailabilitySpecifications             *ResourceTypeRegistrationPropertiesCheckNameAvailabilitySpecifications
	DefaultAPIVersion                               *string
	DisallowedActionVerbs                           []*string
	EnableAsyncOperation                            *bool
	EnableThirdPartyS2S                             *bool
	Endpoints                                       []*ResourceTypeEndpoint
	ExtendedLocations                               []*ExtendedLocationOptions
	ExtensionOptions                                *ResourceTypeRegistrationPropertiesExtensionOptions
	FeaturesRule                                    *ResourceTypeRegistrationPropertiesFeaturesRule
	IdentityManagement                              *ResourceTypeRegistrationPropertiesIdentityManagement
	IsPureProxy                                     *bool
	LinkedAccessChecks                              []*LinkedAccessCheck
	LoggingRules                                    []*LoggingRule
	MarketplaceType                                 *MarketplaceType
	ProvisioningState                               *ProvisioningState
	Regionality                                     *Regionality
	RequestHeaderOptions                            *ResourceTypeRegistrationPropertiesRequestHeaderOptions
	RequiredFeatures                                []*string
	ResourceDeletionPolicy                          *ResourceDeletionPolicy
	ResourceMovePolicy                              *ResourceTypeRegistrationPropertiesResourceMovePolicy
	RoutingType                                     *RoutingType
	ServiceTreeInfos                                []*ServiceTreeInfo
	SubscriptionLifecycleNotificationSpecifications *ResourceTypeRegistrationPropertiesSubscriptionLifecycleNotificationSpecifications
	SubscriptionStateRules                          []*SubscriptionStateRule
	SwaggerSpecifications                           []*SwaggerSpecification
	TemplateDeploymentOptions                       *ResourceTypeRegistrationPropertiesTemplateDeploymentOptions
	ThrottlingRules                                 []*ThrottlingRule
}

type ResourceTypeRegistrationPropertiesCheckNameAvailabilitySpecifications struct {
	EnableDefaultValidation           *bool
	ResourceTypesWithCustomValidation []*string
}

type ResourceTypeRegistrationPropertiesExtensionOptions struct {
	ResourceCreationBegin *ResourceTypeExtensionOptionsResourceCreationBegin
}

type ResourceTypeRegistrationPropertiesFeaturesRule struct {
	// REQUIRED
	RequiredFeaturesPolicy *FeaturesPolicy
}

type ResourceTypeRegistrationPropertiesIdentityManagement struct {
	ApplicationID *string
	Type          *IdentityManagementTypes
}

type ResourceTypeRegistrationPropertiesRequestHeaderOptions struct {
	OptInHeaders *OptInHeaderType
}

type ResourceTypeRegistrationPropertiesResourceMovePolicy struct {
	CrossResourceGroupMoveEnabled *bool
	CrossSubscriptionMoveEnabled  *bool
	ValidationRequired            *bool
}

type ResourceTypeRegistrationPropertiesSubscriptionLifecycleNotificationSpecifications struct {
	SoftDeleteTTL                    *string
	SubscriptionStateOverrideActions []*SubscriptionStateOverrideAction
}

type ResourceTypeRegistrationPropertiesTemplateDeploymentOptions struct {
	PreflightOptions   []*PreflightOption
	PreflightSupported *bool
}

type ResourceTypeRequestHeaderOptions struct {
	OptInHeaders *OptInHeaderType
}

type ResourceTypeSKU struct {
	// REQUIRED
	SKUSettings       []*SKUSetting
	ProvisioningState *ProvisioningState
}

type ResourceTypeTemplateDeploymentPolicy struct {
	// REQUIRED
	Capabilities *TemplateDeploymentCapabilities

	// REQUIRED
	PreflightOptions *TemplateDeploymentPreflightOptions
}

type RolloutStatusBase struct {
	CompletedRegions []*string

	// Dictionary of
	FailedOrSkippedRegions map[string]*ExtendedErrorInfo
}

type SKUCapability struct {
	// REQUIRED
	Name *string

	// REQUIRED
	Value *string
}

type SKUCapacity struct {
	// REQUIRED
	Minimum   *int32
	Default   *int32
	Maximum   *int32
	ScaleType *SKUScaleType
}

type SKUCost struct {
	// REQUIRED
	MeterID      *string
	ExtendedUnit *string
	Quantity     *int32
}

type SKULocationInfo struct {
	// REQUIRED
	Location          *string
	ExtendedLocations []*string
	Type              *ExtendedLocationType
	ZoneDetails       []*SKUZoneDetail
	Zones             []*string
}

type SKUResource struct {
	Properties *SKUResourceProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

type SKUResourceArrayResponseWithContinuation struct {
	// The URL to get to the next set of results, if there are any.
	NextLink *string
	Value    []*SKUResource
}

type SKUResourceProperties struct {
	// REQUIRED
	SKUSettings       []*SKUSetting
	ProvisioningState *ProvisioningState
}

type SKUSetting struct {
	// REQUIRED
	Name             *string
	Capabilities     []*SKUCapability
	Capacity         *SKUSettingCapacity
	Costs            []*SKUCost
	Family           *string
	Kind             *string
	LocationInfo     []*SKULocationInfo
	Locations        []*string
	RequiredFeatures []*string
	RequiredQuotaIDs []*string
	Size             *string
	Tier             *string
}

type SKUSettingCapacity struct {
	// REQUIRED
	Minimum   *int32
	Default   *int32
	Maximum   *int32
	ScaleType *SKUScaleType
}

type SKUZoneDetail struct {
	Capabilities []*SKUCapability
	Name         []*string
}

type ServiceTreeInfo struct {
	ComponentID *string
	ServiceID   *string
}

type SubscriptionLifecycleNotificationSpecifications struct {
	SoftDeleteTTL                    *string
	SubscriptionStateOverrideActions []*SubscriptionStateOverrideAction
}

type SubscriptionStateOverrideAction struct {
	// REQUIRED
	Action *SubscriptionNotificationOperation

	// REQUIRED
	State *SubscriptionTransitioningState
}

type SubscriptionStateRule struct {
	AllowedActions []*string
	State          *SubscriptionState
}

type SwaggerSpecification struct {
	APIVersions          []*string
	SwaggerSpecFolderURI *string
}

type TemplateDeploymentOptions struct {
	PreflightOptions   []*PreflightOption
	PreflightSupported *bool
}

type TemplateDeploymentPolicy struct {
	// REQUIRED
	Capabilities *TemplateDeploymentCapabilities

	// REQUIRED
	PreflightOptions *TemplateDeploymentPreflightOptions
}

type ThirdPartyProviderAuthorization struct {
	Authorizations    []*LightHouseAuthorization
	ManagedByTenantID *string
}

type ThrottlingMetric struct {
	// REQUIRED
	Limit *int64

	// REQUIRED
	Type     *ThrottlingMetricType
	Interval *string
}

type ThrottlingRule struct {
	// REQUIRED
	Action *string

	// REQUIRED
	Metrics          []*ThrottlingMetric
	RequiredFeatures []*string
}

type TrafficRegionRolloutConfiguration struct {
	Regions      []*string
	WaitDuration *string
}

type TrafficRegions struct {
	Regions []*string
}

type TypedErrorInfo struct {
	// REQUIRED
	Type *string

	// READ-ONLY; Anything
	Info any
}
