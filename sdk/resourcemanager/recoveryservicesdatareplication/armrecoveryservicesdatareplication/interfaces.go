// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservicesdatareplication

// EventModelCustomPropertiesClassification provides polymorphic access to related types.
// Call the interface's GetEventModelCustomProperties() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *EventModelCustomProperties, *HyperVToAzStackHCIEventModelCustomProperties, *VMwareToAzStackHCIEventModelCustomProperties
type EventModelCustomPropertiesClassification interface {
	// GetEventModelCustomProperties returns the EventModelCustomProperties content of the underlying type.
	GetEventModelCustomProperties() *EventModelCustomProperties
}

// FabricAgentModelCustomPropertiesClassification provides polymorphic access to related types.
// Call the interface's GetFabricAgentModelCustomProperties() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *FabricAgentModelCustomProperties, *VMwareFabricAgentModelCustomProperties
type FabricAgentModelCustomPropertiesClassification interface {
	// GetFabricAgentModelCustomProperties returns the FabricAgentModelCustomProperties content of the underlying type.
	GetFabricAgentModelCustomProperties() *FabricAgentModelCustomProperties
}

// FabricModelCustomPropertiesClassification provides polymorphic access to related types.
// Call the interface's GetFabricModelCustomProperties() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AzStackHCIFabricModelCustomProperties, *FabricModelCustomProperties, *HyperVMigrateFabricModelCustomProperties, *VMwareMigrateFabricModelCustomProperties
type FabricModelCustomPropertiesClassification interface {
	// GetFabricModelCustomProperties returns the FabricModelCustomProperties content of the underlying type.
	GetFabricModelCustomProperties() *FabricModelCustomProperties
}

// JobModelCustomPropertiesClassification provides polymorphic access to related types.
// Call the interface's GetJobModelCustomProperties() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *FailoverJobModelCustomProperties, *JobModelCustomProperties, *TestFailoverCleanupJobModelCustomProperties, *TestFailoverJobModelCustomProperties
type JobModelCustomPropertiesClassification interface {
	// GetJobModelCustomProperties returns the JobModelCustomProperties content of the underlying type.
	GetJobModelCustomProperties() *JobModelCustomProperties
}

// PlannedFailoverModelCustomPropertiesClassification provides polymorphic access to related types.
// Call the interface's GetPlannedFailoverModelCustomProperties() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *HyperVToAzStackHCIPlannedFailoverModelCustomProperties, *PlannedFailoverModelCustomProperties, *VMwareToAzStackHCIPlannedFailoverModelCustomProperties
type PlannedFailoverModelCustomPropertiesClassification interface {
	// GetPlannedFailoverModelCustomProperties returns the PlannedFailoverModelCustomProperties content of the underlying type.
	GetPlannedFailoverModelCustomProperties() *PlannedFailoverModelCustomProperties
}

// PolicyModelCustomPropertiesClassification provides polymorphic access to related types.
// Call the interface's GetPolicyModelCustomProperties() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *HyperVToAzStackHCIPolicyModelCustomProperties, *PolicyModelCustomProperties, *VMwareToAzStackHCIPolicyModelCustomProperties
type PolicyModelCustomPropertiesClassification interface {
	// GetPolicyModelCustomProperties returns the PolicyModelCustomProperties content of the underlying type.
	GetPolicyModelCustomProperties() *PolicyModelCustomProperties
}

// ProtectedItemModelCustomPropertiesClassification provides polymorphic access to related types.
// Call the interface's GetProtectedItemModelCustomProperties() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *HyperVToAzStackHCIProtectedItemModelCustomProperties, *ProtectedItemModelCustomProperties, *VMwareToAzStackHCIProtectedItemModelCustomProperties
type ProtectedItemModelCustomPropertiesClassification interface {
	// GetProtectedItemModelCustomProperties returns the ProtectedItemModelCustomProperties content of the underlying type.
	GetProtectedItemModelCustomProperties() *ProtectedItemModelCustomProperties
}

// ProtectedItemModelCustomPropertiesUpdateClassification provides polymorphic access to related types.
// Call the interface's GetProtectedItemModelCustomPropertiesUpdate() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *HyperVToAzStackHCIProtectedItemModelCustomPropertiesUpdate, *ProtectedItemModelCustomPropertiesUpdate, *VMwareToAzStackHCIProtectedItemModelCustomPropertiesUpdate
type ProtectedItemModelCustomPropertiesUpdateClassification interface {
	// GetProtectedItemModelCustomPropertiesUpdate returns the ProtectedItemModelCustomPropertiesUpdate content of the underlying type.
	GetProtectedItemModelCustomPropertiesUpdate() *ProtectedItemModelCustomPropertiesUpdate
}

// RecoveryPointModelCustomPropertiesClassification provides polymorphic access to related types.
// Call the interface's GetRecoveryPointModelCustomProperties() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *HyperVToAzStackHCIRecoveryPointModelCustomProperties, *RecoveryPointModelCustomProperties, *VMwareToAzStackHCIRecoveryPointModelCustomProperties
type RecoveryPointModelCustomPropertiesClassification interface {
	// GetRecoveryPointModelCustomProperties returns the RecoveryPointModelCustomProperties content of the underlying type.
	GetRecoveryPointModelCustomProperties() *RecoveryPointModelCustomProperties
}

// ReplicationExtensionModelCustomPropertiesClassification provides polymorphic access to related types.
// Call the interface's GetReplicationExtensionModelCustomProperties() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *HyperVToAzStackHCIReplicationExtensionModelCustomProperties, *ReplicationExtensionModelCustomProperties, *VMwareToAzStackHCIReplicationExtensionModelCustomProperties
type ReplicationExtensionModelCustomPropertiesClassification interface {
	// GetReplicationExtensionModelCustomProperties returns the ReplicationExtensionModelCustomProperties content of the underlying type.
	GetReplicationExtensionModelCustomProperties() *ReplicationExtensionModelCustomProperties
}
