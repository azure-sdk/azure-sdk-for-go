// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armblueprint

// ArtifactClassification provides polymorphic access to related types.
// Call the interface's GetArtifact() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *Artifact, *PolicyAssignmentArtifact, *RoleAssignmentArtifact, *TemplateArtifact
type ArtifactClassification interface {
	// GetArtifact returns the Artifact content of the underlying type.
	GetArtifact() *Artifact
}
