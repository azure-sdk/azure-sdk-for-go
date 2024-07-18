// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armmobilepacketcore

// ClusterServiceClusterTypeSpecificDataClassification provides polymorphic access to related types.
// Call the interface's GetClusterServiceClusterTypeSpecificData() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *ClusterServiceAksClusterData, *ClusterServiceClusterTypeSpecificData, *ClusterServiceNexusAksClusterData
type ClusterServiceClusterTypeSpecificDataClassification interface {
	// GetClusterServiceClusterTypeSpecificData returns the ClusterServiceClusterTypeSpecificData content of the underlying type.
	GetClusterServiceClusterTypeSpecificData() *ClusterServiceClusterTypeSpecificData
}
