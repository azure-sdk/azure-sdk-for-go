// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armedgemarketplace

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/edgemarketplace/armedgemarketplace"
	moduleVersion = "v0.1.0"
)

// ActionType - Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
type ActionType string

const (
	ActionTypeInternal ActionType = "Internal"
)

// PossibleActionTypeValues returns the possible values for the ActionType const type.
func PossibleActionTypeValues() []ActionType {
	return []ActionType{
		ActionTypeInternal,
	}
}

// CreatedByType - The type of identity that created the resource.
type CreatedByType string

const (
	CreatedByTypeApplication     CreatedByType = "Application"
	CreatedByTypeKey             CreatedByType = "Key"
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	CreatedByTypeUser            CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns the possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{
		CreatedByTypeApplication,
		CreatedByTypeKey,
		CreatedByTypeManagedIdentity,
		CreatedByTypeUser,
	}
}

// OfferAvailability - Says if the offer is public/private
type OfferAvailability string

const (
	// OfferAvailabilityPrivate - The offer availability is private
	OfferAvailabilityPrivate OfferAvailability = "Private"
	// OfferAvailabilityPublic - The offer availability is public
	OfferAvailabilityPublic OfferAvailability = "Public"
)

// PossibleOfferAvailabilityValues returns the possible values for the OfferAvailability const type.
func PossibleOfferAvailabilityValues() []OfferAvailability {
	return []OfferAvailability{
		OfferAvailabilityPrivate,
		OfferAvailabilityPublic,
	}
}

// OfferReleaseType - Offer release type
type OfferReleaseType string

const (
	// OfferReleaseTypeGA - The offer in GA
	OfferReleaseTypeGA OfferReleaseType = "GA"
	// OfferReleaseTypePreview - The offer in preview
	OfferReleaseTypePreview OfferReleaseType = "Preview"
)

// PossibleOfferReleaseTypeValues returns the possible values for the OfferReleaseType const type.
func PossibleOfferReleaseTypeValues() []OfferReleaseType {
	return []OfferReleaseType{
		OfferReleaseTypeGA,
		OfferReleaseTypePreview,
	}
}

// Origin - The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
// value is "user,system"
type Origin string

const (
	OriginSystem     Origin = "system"
	OriginUser       Origin = "user"
	OriginUserSystem Origin = "user,system"
)

// PossibleOriginValues returns the possible values for the Origin const type.
func PossibleOriginValues() []Origin {
	return []Origin{
		OriginSystem,
		OriginUser,
		OriginUserSystem,
	}
}

// ResourceProvisioningState - The provisioning state of a resource type.
type ResourceProvisioningState string

const (
	// ResourceProvisioningStateCanceled - Resource creation was canceled.
	ResourceProvisioningStateCanceled ResourceProvisioningState = "Canceled"
	// ResourceProvisioningStateFailed - Resource creation failed.
	ResourceProvisioningStateFailed ResourceProvisioningState = "Failed"
	// ResourceProvisioningStateSucceeded - Resource has been created.
	ResourceProvisioningStateSucceeded ResourceProvisioningState = "Succeeded"
)

// PossibleResourceProvisioningStateValues returns the possible values for the ResourceProvisioningState const type.
func PossibleResourceProvisioningStateValues() []ResourceProvisioningState {
	return []ResourceProvisioningState{
		ResourceProvisioningStateCanceled,
		ResourceProvisioningStateFailed,
		ResourceProvisioningStateSucceeded,
	}
}
