//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcontainerservicefleet

// FleetMembersClientBeginCreateOptions contains the optional parameters for the FleetMembersClient.BeginCreate method.
type FleetMembersClientBeginCreateOptions struct {
	// The request should only proceed if an entity matches this string.
	IfMatch *string

	// The request should only proceed if no entity matches this string.
	IfNoneMatch *string

	// Resumes the LRO from the provided token.
	ResumeToken string
}

// FleetMembersClientBeginDeleteOptions contains the optional parameters for the FleetMembersClient.BeginDelete method.
type FleetMembersClientBeginDeleteOptions struct {
	// The request should only proceed if an entity matches this string.
	IfMatch *string

	// Resumes the LRO from the provided token.
	ResumeToken string
}

// FleetMembersClientBeginUpdateOptions contains the optional parameters for the FleetMembersClient.BeginUpdate method.
type FleetMembersClientBeginUpdateOptions struct {
	// The request should only proceed if an entity matches this string.
	IfMatch *string

	// Resumes the LRO from the provided token.
	ResumeToken string
}

// FleetMembersClientGetOptions contains the optional parameters for the FleetMembersClient.Get method.
type FleetMembersClientGetOptions struct {
	// placeholder for future optional parameters
}

// FleetMembersClientListByFleetOptions contains the optional parameters for the FleetMembersClient.NewListByFleetPager method.
type FleetMembersClientListByFleetOptions struct {
	// placeholder for future optional parameters
}

// FleetsClientBeginCreateOrUpdateOptions contains the optional parameters for the FleetsClient.BeginCreateOrUpdate method.
type FleetsClientBeginCreateOrUpdateOptions struct {
	// The request should only proceed if an entity matches this string.
	IfMatch *string

	// The request should only proceed if no entity matches this string.
	IfNoneMatch *string

	// Resumes the LRO from the provided token.
	ResumeToken string
}

// FleetsClientBeginDeleteOptions contains the optional parameters for the FleetsClient.BeginDelete method.
type FleetsClientBeginDeleteOptions struct {
	// The request should only proceed if an entity matches this string.
	IfMatch *string

	// Resumes the LRO from the provided token.
	ResumeToken string
}

// FleetsClientBeginUpdateOptions contains the optional parameters for the FleetsClient.BeginUpdate method.
type FleetsClientBeginUpdateOptions struct {
	// The request should only proceed if an entity matches this string.
	IfMatch *string

	// Resumes the LRO from the provided token.
	ResumeToken string
}

// FleetsClientGetOptions contains the optional parameters for the FleetsClient.Get method.
type FleetsClientGetOptions struct {
	// placeholder for future optional parameters
}

// FleetsClientListByResourceGroupOptions contains the optional parameters for the FleetsClient.NewListByResourceGroupPager
// method.
type FleetsClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// FleetsClientListBySubscriptionOptions contains the optional parameters for the FleetsClient.NewListBySubscriptionPager
// method.
type FleetsClientListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// FleetsClientListCredentialsOptions contains the optional parameters for the FleetsClient.ListCredentials method.
type FleetsClientListCredentialsOptions struct {
	// placeholder for future optional parameters
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.NewListPager method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// UpdateRunsClientBeginCreateOrUpdateOptions contains the optional parameters for the UpdateRunsClient.BeginCreateOrUpdate
// method.
type UpdateRunsClientBeginCreateOrUpdateOptions struct {
	// The request should only proceed if an entity matches this string.
	IfMatch *string

	// The request should only proceed if no entity matches this string.
	IfNoneMatch *string

	// Resumes the LRO from the provided token.
	ResumeToken string
}

// UpdateRunsClientBeginDeleteOptions contains the optional parameters for the UpdateRunsClient.BeginDelete method.
type UpdateRunsClientBeginDeleteOptions struct {
	// The request should only proceed if an entity matches this string.
	IfMatch *string

	// Resumes the LRO from the provided token.
	ResumeToken string
}

// UpdateRunsClientBeginStartOptions contains the optional parameters for the UpdateRunsClient.BeginStart method.
type UpdateRunsClientBeginStartOptions struct {
	// The request should only proceed if an entity matches this string.
	IfMatch *string

	// Resumes the LRO from the provided token.
	ResumeToken string
}

// UpdateRunsClientBeginStopOptions contains the optional parameters for the UpdateRunsClient.BeginStop method.
type UpdateRunsClientBeginStopOptions struct {
	// The request should only proceed if an entity matches this string.
	IfMatch *string

	// Resumes the LRO from the provided token.
	ResumeToken string
}

// UpdateRunsClientGetOptions contains the optional parameters for the UpdateRunsClient.Get method.
type UpdateRunsClientGetOptions struct {
	// placeholder for future optional parameters
}

// UpdateRunsClientListByFleetOptions contains the optional parameters for the UpdateRunsClient.NewListByFleetPager method.
type UpdateRunsClientListByFleetOptions struct {
	// placeholder for future optional parameters
}
