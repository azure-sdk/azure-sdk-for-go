// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armscheduler

// HTTPAuthenticationClassification provides polymorphic access to related types.
// Call the interface's GetHTTPAuthentication() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *BasicAuthentication, *ClientCertAuthentication, *HTTPAuthentication, *OAuthAuthentication
type HTTPAuthenticationClassification interface {
	// GetHTTPAuthentication returns the HTTPAuthentication content of the underlying type.
	GetHTTPAuthentication() *HTTPAuthentication
}
