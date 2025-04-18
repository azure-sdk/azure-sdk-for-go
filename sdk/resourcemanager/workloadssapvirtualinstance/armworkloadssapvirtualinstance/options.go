// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armworkloadssapvirtualinstance

// OperationsClientListOptions contains the optional parameters for the OperationsClient.NewListPager method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// SapApplicationServerInstancesClientBeginCreateOptions contains the optional parameters for the SapApplicationServerInstancesClient.BeginCreate
// method.
type SapApplicationServerInstancesClientBeginCreateOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// SapApplicationServerInstancesClientBeginDeleteOptions contains the optional parameters for the SapApplicationServerInstancesClient.BeginDelete
// method.
type SapApplicationServerInstancesClientBeginDeleteOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// SapApplicationServerInstancesClientBeginStartOptions contains the optional parameters for the SapApplicationServerInstancesClient.BeginStart
// method.
type SapApplicationServerInstancesClientBeginStartOptions struct {
	// SAP Application server instance start request body.
	Body *StartRequest

	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// SapApplicationServerInstancesClientBeginStopOptions contains the optional parameters for the SapApplicationServerInstancesClient.BeginStop
// method.
type SapApplicationServerInstancesClientBeginStopOptions struct {
	// SAP Application server instance stop request body.
	Body *StopRequest

	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// SapApplicationServerInstancesClientGetOptions contains the optional parameters for the SapApplicationServerInstancesClient.Get
// method.
type SapApplicationServerInstancesClientGetOptions struct {
	// placeholder for future optional parameters
}

// SapApplicationServerInstancesClientListOptions contains the optional parameters for the SapApplicationServerInstancesClient.NewListPager
// method.
type SapApplicationServerInstancesClientListOptions struct {
	// placeholder for future optional parameters
}

// SapApplicationServerInstancesClientUpdateOptions contains the optional parameters for the SapApplicationServerInstancesClient.Update
// method.
type SapApplicationServerInstancesClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// SapCentralServerInstancesClientBeginCreateOptions contains the optional parameters for the SapCentralServerInstancesClient.BeginCreate
// method.
type SapCentralServerInstancesClientBeginCreateOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// SapCentralServerInstancesClientBeginDeleteOptions contains the optional parameters for the SapCentralServerInstancesClient.BeginDelete
// method.
type SapCentralServerInstancesClientBeginDeleteOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// SapCentralServerInstancesClientBeginStartOptions contains the optional parameters for the SapCentralServerInstancesClient.BeginStart
// method.
type SapCentralServerInstancesClientBeginStartOptions struct {
	// SAP Central Services instance start request body.
	Body *StartRequest

	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// SapCentralServerInstancesClientBeginStopOptions contains the optional parameters for the SapCentralServerInstancesClient.BeginStop
// method.
type SapCentralServerInstancesClientBeginStopOptions struct {
	// SAP Central Services instance stop request body.
	Body *StopRequest

	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// SapCentralServerInstancesClientGetOptions contains the optional parameters for the SapCentralServerInstancesClient.Get
// method.
type SapCentralServerInstancesClientGetOptions struct {
	// placeholder for future optional parameters
}

// SapCentralServerInstancesClientListOptions contains the optional parameters for the SapCentralServerInstancesClient.NewListPager
// method.
type SapCentralServerInstancesClientListOptions struct {
	// placeholder for future optional parameters
}

// SapCentralServerInstancesClientUpdateOptions contains the optional parameters for the SapCentralServerInstancesClient.Update
// method.
type SapCentralServerInstancesClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// SapDatabaseInstancesClientBeginCreateOptions contains the optional parameters for the SapDatabaseInstancesClient.BeginCreate
// method.
type SapDatabaseInstancesClientBeginCreateOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// SapDatabaseInstancesClientBeginDeleteOptions contains the optional parameters for the SapDatabaseInstancesClient.BeginDelete
// method.
type SapDatabaseInstancesClientBeginDeleteOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// SapDatabaseInstancesClientBeginStartOptions contains the optional parameters for the SapDatabaseInstancesClient.BeginStart
// method.
type SapDatabaseInstancesClientBeginStartOptions struct {
	// SAP Database server instance start request body.
	Body *StartRequest

	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// SapDatabaseInstancesClientBeginStopOptions contains the optional parameters for the SapDatabaseInstancesClient.BeginStop
// method.
type SapDatabaseInstancesClientBeginStopOptions struct {
	// Stop request for the database instance of the SAP system.
	Body *StopRequest

	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// SapDatabaseInstancesClientGetOptions contains the optional parameters for the SapDatabaseInstancesClient.Get method.
type SapDatabaseInstancesClientGetOptions struct {
	// placeholder for future optional parameters
}

// SapDatabaseInstancesClientListOptions contains the optional parameters for the SapDatabaseInstancesClient.NewListPager
// method.
type SapDatabaseInstancesClientListOptions struct {
	// placeholder for future optional parameters
}

// SapDatabaseInstancesClientUpdateOptions contains the optional parameters for the SapDatabaseInstancesClient.Update method.
type SapDatabaseInstancesClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// SapVirtualInstancesClientBeginCreateOptions contains the optional parameters for the SapVirtualInstancesClient.BeginCreate
// method.
type SapVirtualInstancesClientBeginCreateOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// SapVirtualInstancesClientBeginDeleteOptions contains the optional parameters for the SapVirtualInstancesClient.BeginDelete
// method.
type SapVirtualInstancesClientBeginDeleteOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// SapVirtualInstancesClientBeginStartOptions contains the optional parameters for the SapVirtualInstancesClient.BeginStart
// method.
type SapVirtualInstancesClientBeginStartOptions struct {
	// The Virtual Instance for SAP solutions resource start request body.
	Body *StartRequest

	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// SapVirtualInstancesClientBeginStopOptions contains the optional parameters for the SapVirtualInstancesClient.BeginStop
// method.
type SapVirtualInstancesClientBeginStopOptions struct {
	// The Virtual Instance for SAP solutions resource stop request body.
	Body *StopRequest

	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// SapVirtualInstancesClientBeginUpdateOptions contains the optional parameters for the SapVirtualInstancesClient.BeginUpdate
// method.
type SapVirtualInstancesClientBeginUpdateOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// SapVirtualInstancesClientGetOptions contains the optional parameters for the SapVirtualInstancesClient.Get method.
type SapVirtualInstancesClientGetOptions struct {
	// placeholder for future optional parameters
}

// SapVirtualInstancesClientInvokeAvailabilityZoneDetailsOptions contains the optional parameters for the SapVirtualInstancesClient.InvokeAvailabilityZoneDetails
// method.
type SapVirtualInstancesClientInvokeAvailabilityZoneDetailsOptions struct {
	// placeholder for future optional parameters
}

// SapVirtualInstancesClientInvokeDiskConfigurationsOptions contains the optional parameters for the SapVirtualInstancesClient.InvokeDiskConfigurations
// method.
type SapVirtualInstancesClientInvokeDiskConfigurationsOptions struct {
	// placeholder for future optional parameters
}

// SapVirtualInstancesClientInvokeSapSupportedSKUOptions contains the optional parameters for the SapVirtualInstancesClient.InvokeSapSupportedSKU
// method.
type SapVirtualInstancesClientInvokeSapSupportedSKUOptions struct {
	// placeholder for future optional parameters
}

// SapVirtualInstancesClientInvokeSizingRecommendationsOptions contains the optional parameters for the SapVirtualInstancesClient.InvokeSizingRecommendations
// method.
type SapVirtualInstancesClientInvokeSizingRecommendationsOptions struct {
	// placeholder for future optional parameters
}

// SapVirtualInstancesClientListByResourceGroupOptions contains the optional parameters for the SapVirtualInstancesClient.NewListByResourceGroupPager
// method.
type SapVirtualInstancesClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// SapVirtualInstancesClientListBySubscriptionOptions contains the optional parameters for the SapVirtualInstancesClient.NewListBySubscriptionPager
// method.
type SapVirtualInstancesClientListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}
