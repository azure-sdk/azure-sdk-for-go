// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcognitiveservices

// ConnectionPropertiesV2Classification provides polymorphic access to related types.
// Call the interface's GetConnectionPropertiesV2() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AADAuthTypeConnectionProperties, *APIKeyAuthConnectionProperties, *AccessKeyAuthTypeConnectionProperties, *AccountKeyAuthTypeConnectionProperties,
// - *ConnectionPropertiesV2, *CustomKeysConnectionProperties, *ManagedIdentityAuthTypeConnectionProperties, *NoneAuthTypeConnectionProperties,
// - *OAuth2AuthTypeConnectionProperties, *PATAuthTypeConnectionProperties, *SASAuthTypeConnectionProperties, *ServicePrincipalAuthTypeConnectionProperties,
// - *UsernamePasswordAuthTypeConnectionProperties
type ConnectionPropertiesV2Classification interface {
	// GetConnectionPropertiesV2 returns the ConnectionPropertiesV2 content of the underlying type.
	GetConnectionPropertiesV2() *ConnectionPropertiesV2
}
