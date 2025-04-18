// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armoracledatabase

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
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
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

// NewAutonomousDatabaseBackupsClient creates a new instance of AutonomousDatabaseBackupsClient.
func (c *ClientFactory) NewAutonomousDatabaseBackupsClient() *AutonomousDatabaseBackupsClient {
	return &AutonomousDatabaseBackupsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewAutonomousDatabaseCharacterSetsClient creates a new instance of AutonomousDatabaseCharacterSetsClient.
func (c *ClientFactory) NewAutonomousDatabaseCharacterSetsClient() *AutonomousDatabaseCharacterSetsClient {
	return &AutonomousDatabaseCharacterSetsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewAutonomousDatabaseNationalCharacterSetsClient creates a new instance of AutonomousDatabaseNationalCharacterSetsClient.
func (c *ClientFactory) NewAutonomousDatabaseNationalCharacterSetsClient() *AutonomousDatabaseNationalCharacterSetsClient {
	return &AutonomousDatabaseNationalCharacterSetsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewAutonomousDatabaseVersionsClient creates a new instance of AutonomousDatabaseVersionsClient.
func (c *ClientFactory) NewAutonomousDatabaseVersionsClient() *AutonomousDatabaseVersionsClient {
	return &AutonomousDatabaseVersionsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewAutonomousDatabasesClient creates a new instance of AutonomousDatabasesClient.
func (c *ClientFactory) NewAutonomousDatabasesClient() *AutonomousDatabasesClient {
	return &AutonomousDatabasesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewCloudExadataInfrastructuresClient creates a new instance of CloudExadataInfrastructuresClient.
func (c *ClientFactory) NewCloudExadataInfrastructuresClient() *CloudExadataInfrastructuresClient {
	return &CloudExadataInfrastructuresClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewCloudVMClustersClient creates a new instance of CloudVMClustersClient.
func (c *ClientFactory) NewCloudVMClustersClient() *CloudVMClustersClient {
	return &CloudVMClustersClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewDNSPrivateViewsClient creates a new instance of DNSPrivateViewsClient.
func (c *ClientFactory) NewDNSPrivateViewsClient() *DNSPrivateViewsClient {
	return &DNSPrivateViewsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewDNSPrivateZonesClient creates a new instance of DNSPrivateZonesClient.
func (c *ClientFactory) NewDNSPrivateZonesClient() *DNSPrivateZonesClient {
	return &DNSPrivateZonesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewDbNodesClient creates a new instance of DbNodesClient.
func (c *ClientFactory) NewDbNodesClient() *DbNodesClient {
	return &DbNodesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewDbServersClient creates a new instance of DbServersClient.
func (c *ClientFactory) NewDbServersClient() *DbServersClient {
	return &DbServersClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewDbSystemShapesClient creates a new instance of DbSystemShapesClient.
func (c *ClientFactory) NewDbSystemShapesClient() *DbSystemShapesClient {
	return &DbSystemShapesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewExadbVMClustersClient creates a new instance of ExadbVMClustersClient.
func (c *ClientFactory) NewExadbVMClustersClient() *ExadbVMClustersClient {
	return &ExadbVMClustersClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewExascaleDbNodesClient creates a new instance of ExascaleDbNodesClient.
func (c *ClientFactory) NewExascaleDbNodesClient() *ExascaleDbNodesClient {
	return &ExascaleDbNodesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewExascaleDbStorageVaultsClient creates a new instance of ExascaleDbStorageVaultsClient.
func (c *ClientFactory) NewExascaleDbStorageVaultsClient() *ExascaleDbStorageVaultsClient {
	return &ExascaleDbStorageVaultsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewFlexComponentsClient creates a new instance of FlexComponentsClient.
func (c *ClientFactory) NewFlexComponentsClient() *FlexComponentsClient {
	return &FlexComponentsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewGiMinorVersionsClient creates a new instance of GiMinorVersionsClient.
func (c *ClientFactory) NewGiMinorVersionsClient() *GiMinorVersionsClient {
	return &GiMinorVersionsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewGiVersionsClient creates a new instance of GiVersionsClient.
func (c *ClientFactory) NewGiVersionsClient() *GiVersionsClient {
	return &GiVersionsClient{
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

// NewOracleSubscriptionsClient creates a new instance of OracleSubscriptionsClient.
func (c *ClientFactory) NewOracleSubscriptionsClient() *OracleSubscriptionsClient {
	return &OracleSubscriptionsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSystemVersionsClient creates a new instance of SystemVersionsClient.
func (c *ClientFactory) NewSystemVersionsClient() *SystemVersionsClient {
	return &SystemVersionsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewVirtualNetworkAddressesClient creates a new instance of VirtualNetworkAddressesClient.
func (c *ClientFactory) NewVirtualNetworkAddressesClient() *VirtualNetworkAddressesClient {
	return &VirtualNetworkAddressesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}
