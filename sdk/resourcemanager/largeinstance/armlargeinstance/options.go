// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armlargeinstance

// AzureLargeInstancesClientBeginRestartOptions contains the optional parameters for the AzureLargeInstancesClient.BeginRestart
// method.
type AzureLargeInstancesClientBeginRestartOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string

	// When set to 'active', this parameter empowers the server with the ability to forcefully terminate and halt any existing
	// processes that may be running on the server
	ForceParameter *ForceState
}

// AzureLargeInstancesClientBeginShutdownOptions contains the optional parameters for the AzureLargeInstancesClient.BeginShutdown
// method.
type AzureLargeInstancesClientBeginShutdownOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// AzureLargeInstancesClientBeginStartOptions contains the optional parameters for the AzureLargeInstancesClient.BeginStart
// method.
type AzureLargeInstancesClientBeginStartOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// AzureLargeInstancesClientGetOptions contains the optional parameters for the AzureLargeInstancesClient.Get method.
type AzureLargeInstancesClientGetOptions struct {
	// placeholder for future optional parameters
}

// AzureLargeInstancesClientListByResourceGroupOptions contains the optional parameters for the AzureLargeInstancesClient.NewListByResourceGroupPager
// method.
type AzureLargeInstancesClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// AzureLargeInstancesClientListBySubscriptionOptions contains the optional parameters for the AzureLargeInstancesClient.NewListBySubscriptionPager
// method.
type AzureLargeInstancesClientListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// AzureLargeInstancesClientUpdateOptions contains the optional parameters for the AzureLargeInstancesClient.Update method.
type AzureLargeInstancesClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// AzureLargeStorageInstancesClientGetOptions contains the optional parameters for the AzureLargeStorageInstancesClient.Get
// method.
type AzureLargeStorageInstancesClientGetOptions struct {
	// placeholder for future optional parameters
}

// AzureLargeStorageInstancesClientListByResourceGroupOptions contains the optional parameters for the AzureLargeStorageInstancesClient.NewListByResourceGroupPager
// method.
type AzureLargeStorageInstancesClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// AzureLargeStorageInstancesClientListBySubscriptionOptions contains the optional parameters for the AzureLargeStorageInstancesClient.NewListBySubscriptionPager
// method.
type AzureLargeStorageInstancesClientListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// AzureLargeStorageInstancesClientUpdateOptions contains the optional parameters for the AzureLargeStorageInstancesClient.Update
// method.
type AzureLargeStorageInstancesClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.NewListPager method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}
