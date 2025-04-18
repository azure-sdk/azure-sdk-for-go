// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armoep

// EnergyServicesClientBeginAddPartitionOptions contains the optional parameters for the EnergyServicesClient.BeginAddPartition
// method.
type EnergyServicesClientBeginAddPartitionOptions struct {
	// add partition action payload
	Body *DataPartitionAddOrRemoveRequest

	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// EnergyServicesClientBeginCreateOptions contains the optional parameters for the EnergyServicesClient.BeginCreate method.
type EnergyServicesClientBeginCreateOptions struct {
	// Request body.
	Body *EnergyService

	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// EnergyServicesClientBeginDeleteOptions contains the optional parameters for the EnergyServicesClient.BeginDelete method.
type EnergyServicesClientBeginDeleteOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// EnergyServicesClientBeginRemovePartitionOptions contains the optional parameters for the EnergyServicesClient.BeginRemovePartition
// method.
type EnergyServicesClientBeginRemovePartitionOptions struct {
	// remove partition action payload
	Body *DataPartitionAddOrRemoveRequest

	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// EnergyServicesClientGetOptions contains the optional parameters for the EnergyServicesClient.Get method.
type EnergyServicesClientGetOptions struct {
	// placeholder for future optional parameters
}

// EnergyServicesClientListByResourceGroupOptions contains the optional parameters for the EnergyServicesClient.NewListByResourceGroupPager
// method.
type EnergyServicesClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// EnergyServicesClientListBySubscriptionOptions contains the optional parameters for the EnergyServicesClient.NewListBySubscriptionPager
// method.
type EnergyServicesClientListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// EnergyServicesClientListPartitionsOptions contains the optional parameters for the EnergyServicesClient.ListPartitions
// method.
type EnergyServicesClientListPartitionsOptions struct {
	// placeholder for future optional parameters
}

// EnergyServicesClientUpdateOptions contains the optional parameters for the EnergyServicesClient.Update method.
type EnergyServicesClientUpdateOptions struct {
	Body *EnergyResourceUpdate
}

// LocationsClientCheckNameAvailabilityOptions contains the optional parameters for the LocationsClient.CheckNameAvailability
// method.
type LocationsClientCheckNameAvailabilityOptions struct {
	// placeholder for future optional parameters
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.List method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}
