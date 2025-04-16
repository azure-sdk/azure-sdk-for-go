// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmanagementgroups

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
)

// ClientFactory is a client factory used to create any client in this module.
// Don't use this type directly, use NewClientFactory instead.
type ClientFactory struct {
	internal *arm.Client
}

// NewClientFactory creates a new instance of ClientFactory with the specified values.
// The parameter values will be propagated to any client created from this factory.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClientFactory(credential azcore.TokenCredential, options *arm.ClientOptions) (*ClientFactory, error) {
	internal, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	return &ClientFactory{
		internal: internal,
	}, nil
}

// NewAPIClient creates a new instance of APIClient.
func (c *ClientFactory) NewAPIClient() *APIClient {
	return &APIClient{
		internal: c.internal,
	}
}

// NewClient creates a new instance of Client.
func (c *ClientFactory) NewClient() *Client {
	return &Client{
		internal: c.internal,
	}
}

// NewEntitiesClient creates a new instance of EntitiesClient.
func (c *ClientFactory) NewEntitiesClient() *EntitiesClient {
	return &EntitiesClient{
		internal: c.internal,
	}
}

// NewHierarchySettingsClient creates a new instance of HierarchySettingsClient.
func (c *ClientFactory) NewHierarchySettingsClient() *HierarchySettingsClient {
	return &HierarchySettingsClient{
		internal: c.internal,
	}
}

// NewManagementGroupSubscriptionsClient creates a new instance of ManagementGroupSubscriptionsClient.
func (c *ClientFactory) NewManagementGroupSubscriptionsClient() *ManagementGroupSubscriptionsClient {
	return &ManagementGroupSubscriptionsClient{
		internal: c.internal,
	}
}

// NewOperationsClient creates a new instance of OperationsClient.
func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	return &OperationsClient{
		internal: c.internal,
	}
}
