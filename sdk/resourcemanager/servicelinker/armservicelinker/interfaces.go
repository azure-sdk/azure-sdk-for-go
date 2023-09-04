//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armservicelinker

// AuthInfoBaseClassification provides polymorphic access to related types.
// Call the interface's GetAuthInfoBase() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AccessKeyInfoBase, *AuthInfoBase, *SecretAuthInfo, *ServicePrincipalCertificateAuthInfo, *ServicePrincipalSecretAuthInfo,
// - *SystemAssignedIdentityAuthInfo, *UserAccountAuthInfo, *UserAssignedIdentityAuthInfo
type AuthInfoBaseClassification interface {
	// GetAuthInfoBase returns the AuthInfoBase content of the underlying type.
	GetAuthInfoBase() *AuthInfoBase
}

// AzureResourcePropertiesBaseClassification provides polymorphic access to related types.
// Call the interface's GetAzureResourcePropertiesBase() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AzureKeyVaultProperties, *AzureResourcePropertiesBase
type AzureResourcePropertiesBaseClassification interface {
	// GetAzureResourcePropertiesBase returns the AzureResourcePropertiesBase content of the underlying type.
	GetAzureResourcePropertiesBase() *AzureResourcePropertiesBase
}

// DryrunParametersClassification provides polymorphic access to related types.
// Call the interface's GetDryrunParameters() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *CreateOrUpdateDryrunParameters, *DryrunParameters
type DryrunParametersClassification interface {
	// GetDryrunParameters returns the DryrunParameters content of the underlying type.
	GetDryrunParameters() *DryrunParameters
}

// DryrunPrerequisiteResultClassification provides polymorphic access to related types.
// Call the interface's GetDryrunPrerequisiteResult() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *BasicErrorDryrunPrerequisiteResult, *DryrunPrerequisiteResult, *PermissionsMissingDryrunPrerequisiteResult
type DryrunPrerequisiteResultClassification interface {
	// GetDryrunPrerequisiteResult returns the DryrunPrerequisiteResult content of the underlying type.
	GetDryrunPrerequisiteResult() *DryrunPrerequisiteResult
}

// SecretInfoBaseClassification provides polymorphic access to related types.
// Call the interface's GetSecretInfoBase() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *KeyVaultSecretReferenceSecretInfo, *KeyVaultSecretURISecretInfo, *SecretInfoBase, *ValueSecretInfo
type SecretInfoBaseClassification interface {
	// GetSecretInfoBase returns the SecretInfoBase content of the underlying type.
	GetSecretInfoBase() *SecretInfoBase
}

// TargetServiceBaseClassification provides polymorphic access to related types.
// Call the interface's GetTargetServiceBase() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AzureResource, *ConfluentBootstrapServer, *ConfluentSchemaRegistry, *SelfHostedServer, *TargetServiceBase
type TargetServiceBaseClassification interface {
	// GetTargetServiceBase returns the TargetServiceBase content of the underlying type.
	GetTargetServiceBase() *TargetServiceBase
}
