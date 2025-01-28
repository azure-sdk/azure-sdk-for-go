//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmigrationmodernization

// MigrateAgentModelCustomPropertiesClassification provides polymorphic access to related types.
// Call the interface's GetMigrateAgentModelCustomProperties() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *MigrateAgentModelCustomProperties, *VMwareMigrateAgentModelCustomProperties
type MigrateAgentModelCustomPropertiesClassification interface {
	// GetMigrateAgentModelCustomProperties returns the MigrateAgentModelCustomProperties content of the underlying type.
	GetMigrateAgentModelCustomProperties() *MigrateAgentModelCustomProperties
}

// WorkflowModelCustomPropertiesClassification provides polymorphic access to related types.
// Call the interface's GetWorkflowModelCustomProperties() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *BuildContainerImageWorkflowModelCustomProperties, *EnableReplicationWorkflowModelCustomProperties, *MigrateWorkflowModelCustomProperties,
// - *TestMigrateCleanupWorkflowModelCustomProperties, *TestMigrateWorkflowModelCustomProperties, *WorkflowModelCustomProperties
type WorkflowModelCustomPropertiesClassification interface {
	// GetWorkflowModelCustomProperties returns the WorkflowModelCustomProperties content of the underlying type.
	GetWorkflowModelCustomProperties() *WorkflowModelCustomProperties
}

// WorkloadDeploymentModelCustomPropertiesClassification provides polymorphic access to related types.
// Call the interface's GetWorkloadDeploymentModelCustomProperties() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *ApacheTomcatAKSWorkloadDeploymentModelCustomProperties, *IISAKSWorkloadDeploymentModelCustomProperties, *WorkloadDeploymentModelCustomProperties
type WorkloadDeploymentModelCustomPropertiesClassification interface {
	// GetWorkloadDeploymentModelCustomProperties returns the WorkloadDeploymentModelCustomProperties content of the underlying type.
	GetWorkloadDeploymentModelCustomProperties() *WorkloadDeploymentModelCustomProperties
}

// WorkloadInstanceModelCustomPropertiesClassification provides polymorphic access to related types.
// Call the interface's GetWorkloadInstanceModelCustomProperties() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *ApacheTomcatWorkloadInstanceModelCustomProperties, *IISWorkloadInstanceModelCustomProperties, *WorkloadInstanceModelCustomProperties
type WorkloadInstanceModelCustomPropertiesClassification interface {
	// GetWorkloadInstanceModelCustomProperties returns the WorkloadInstanceModelCustomProperties content of the underlying type.
	GetWorkloadInstanceModelCustomProperties() *WorkloadInstanceModelCustomProperties
}
