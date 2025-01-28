//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmigrationhub

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
)

// ClientFactory is a client factory used to create any client in this module.
// Don't use this type directly, use NewClientFactory instead.
type ClientFactory struct {
	subscriptionID string
	internal       *arm.Client
}

// NewClientFactory creates a new instance of ClientFactory with the specified values.
// The parameter values will be propagated to any client created from this factory.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClientFactory(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ClientFactory, error) {
	internal, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	return &ClientFactory{
		subscriptionID: subscriptionID,
		internal:       internal,
	}, nil
}

// NewDatabaseInstancesControllerClient creates a new instance of DatabaseInstancesControllerClient.
func (c *ClientFactory) NewDatabaseInstancesControllerClient() *DatabaseInstancesControllerClient {
	return &DatabaseInstancesControllerClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewDatabasesControllerClient creates a new instance of DatabasesControllerClient.
func (c *ClientFactory) NewDatabasesControllerClient() *DatabasesControllerClient {
	return &DatabasesControllerClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewEventsControllerClient creates a new instance of EventsControllerClient.
func (c *ClientFactory) NewEventsControllerClient() *EventsControllerClient {
	return &EventsControllerClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewMachinesControllerClient creates a new instance of MachinesControllerClient.
func (c *ClientFactory) NewMachinesControllerClient() *MachinesControllerClient {
	return &MachinesControllerClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewMigrateProjectsControllerClient creates a new instance of MigrateProjectsControllerClient.
func (c *ClientFactory) NewMigrateProjectsControllerClient() *MigrateProjectsControllerClient {
	return &MigrateProjectsControllerClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewOperationsClient creates a new instance of OperationsClient.
func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	return &OperationsClient{
		internal: c.internal,
	}
}

// NewPrivateEndpointConnectionControllerClient creates a new instance of PrivateEndpointConnectionControllerClient.
func (c *ClientFactory) NewPrivateEndpointConnectionControllerClient() *PrivateEndpointConnectionControllerClient {
	return &PrivateEndpointConnectionControllerClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewPrivateEndpointConnectionProxyControllerClient creates a new instance of PrivateEndpointConnectionProxyControllerClient.
func (c *ClientFactory) NewPrivateEndpointConnectionProxyControllerClient() *PrivateEndpointConnectionProxyControllerClient {
	return &PrivateEndpointConnectionProxyControllerClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewPrivateEndpointConnectionsControllerClient creates a new instance of PrivateEndpointConnectionsControllerClient.
func (c *ClientFactory) NewPrivateEndpointConnectionsControllerClient() *PrivateEndpointConnectionsControllerClient {
	return &PrivateEndpointConnectionsControllerClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewPrivateLinkResourceControllerClient creates a new instance of PrivateLinkResourceControllerClient.
func (c *ClientFactory) NewPrivateLinkResourceControllerClient() *PrivateLinkResourceControllerClient {
	return &PrivateLinkResourceControllerClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewProjectsClient creates a new instance of ProjectsClient.
func (c *ClientFactory) NewProjectsClient() *ProjectsClient {
	return &ProjectsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSolutionsControllerClient creates a new instance of SolutionsControllerClient.
func (c *ClientFactory) NewSolutionsControllerClient() *SolutionsControllerClient {
	return &SolutionsControllerClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewVirtualDesktopUserControllerClient creates a new instance of VirtualDesktopUserControllerClient.
func (c *ClientFactory) NewVirtualDesktopUserControllerClient() *VirtualDesktopUserControllerClient {
	return &VirtualDesktopUserControllerClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewWebServersControllerClient creates a new instance of WebServersControllerClient.
func (c *ClientFactory) NewWebServersControllerClient() *WebServersControllerClient {
	return &WebServersControllerClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewWebSitesControllerClient creates a new instance of WebSitesControllerClient.
func (c *ClientFactory) NewWebSitesControllerClient() *WebSitesControllerClient {
	return &WebSitesControllerClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}
