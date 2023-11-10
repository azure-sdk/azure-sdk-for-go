//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armazurearcdata

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
)

// ClientFactory is a client factory used to create any client in this module.
// Don't use this type directly, use NewClientFactory instead.
type ClientFactory struct {
	subscriptionID string
	credential     azcore.TokenCredential
	options        *arm.ClientOptions
}

// NewClientFactory creates a new instance of ClientFactory with the specified values.
// The parameter values will be propagated to any client created from this factory.
//   - subscriptionID - The ID of the Azure subscription
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClientFactory(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ClientFactory, error) {
	_, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	return &ClientFactory{
		subscriptionID: subscriptionID, credential: credential,
		options: options.Clone(),
	}, nil
}

// NewActiveDirectoryConnectorsClient creates a new instance of ActiveDirectoryConnectorsClient.
func (c *ClientFactory) NewActiveDirectoryConnectorsClient() *ActiveDirectoryConnectorsClient {
	subClient, _ := NewActiveDirectoryConnectorsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewDataControllersClient creates a new instance of DataControllersClient.
func (c *ClientFactory) NewDataControllersClient() *DataControllersClient {
	subClient, _ := NewDataControllersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewFailoverGroupsClient creates a new instance of FailoverGroupsClient.
func (c *ClientFactory) NewFailoverGroupsClient() *FailoverGroupsClient {
	subClient, _ := NewFailoverGroupsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewOperationsClient creates a new instance of OperationsClient.
func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	subClient, _ := NewOperationsClient(c.credential, c.options)
	return subClient
}

// NewPostgresInstancesClient creates a new instance of PostgresInstancesClient.
func (c *ClientFactory) NewPostgresInstancesClient() *PostgresInstancesClient {
	subClient, _ := NewPostgresInstancesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewSQLManagedInstancesClient creates a new instance of SQLManagedInstancesClient.
func (c *ClientFactory) NewSQLManagedInstancesClient() *SQLManagedInstancesClient {
	subClient, _ := NewSQLManagedInstancesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewSQLServerAvailabilityGroupsClient creates a new instance of SQLServerAvailabilityGroupsClient.
func (c *ClientFactory) NewSQLServerAvailabilityGroupsClient() *SQLServerAvailabilityGroupsClient {
	subClient, _ := NewSQLServerAvailabilityGroupsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewSQLServerDatabasesClient creates a new instance of SQLServerDatabasesClient.
func (c *ClientFactory) NewSQLServerDatabasesClient() *SQLServerDatabasesClient {
	subClient, _ := NewSQLServerDatabasesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewSQLServerInstancesClient creates a new instance of SQLServerInstancesClient.
func (c *ClientFactory) NewSQLServerInstancesClient() *SQLServerInstancesClient {
	subClient, _ := NewSQLServerInstancesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}
