//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsitemanager

import "time"

// Site as ARM Resource
type Site struct {
	// The resource-specific properties for this resource.
	Properties *SiteProperties

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// SiteListResult - The response of a Site list operation.
type SiteListResult struct {
	// REQUIRED; The Site items on this page
	Value []*Site

	// The link to the next page of items
	NextLink *string
}

// SiteProperties - Site properties
type SiteProperties struct {
	// AddressResource ArmId of Site resource
	AddressResourceID *string

	// Description of Site resource
	Description *string

	// displayName of Site resource
	DisplayName *string

	// READ-ONLY; Provisioning state of last operation
	ProvisioningState *ResourceProvisioningState
}

// SiteUpdate - The type used for update operations of the Site.
type SiteUpdate struct {
	// The updatable properties of the Site.
	Properties *SiteUpdateProperties
}

// SiteUpdateProperties - The updatable properties of the Site.
type SiteUpdateProperties struct {
	// AddressResource ArmId of Site resource
	AddressResourceID *string

	// Description of Site resource
	Description *string

	// displayName of Site resource
	DisplayName *string
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time

	// The identity that created the resource.
	CreatedBy *string

	// The type of identity that created the resource.
	CreatedByType *CreatedByType

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time

	// The identity that last modified the resource.
	LastModifiedBy *string

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType
}
