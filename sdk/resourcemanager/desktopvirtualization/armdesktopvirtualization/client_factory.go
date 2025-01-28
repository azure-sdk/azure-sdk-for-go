//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdesktopvirtualization

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

// NewActiveSessionHostConfigurationsClient creates a new instance of ActiveSessionHostConfigurationsClient.
func (c *ClientFactory) NewActiveSessionHostConfigurationsClient() *ActiveSessionHostConfigurationsClient {
	return &ActiveSessionHostConfigurationsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewAppAttachPackageClient creates a new instance of AppAttachPackageClient.
func (c *ClientFactory) NewAppAttachPackageClient() *AppAttachPackageClient {
	return &AppAttachPackageClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewAppAttachPackageInfoClient creates a new instance of AppAttachPackageInfoClient.
func (c *ClientFactory) NewAppAttachPackageInfoClient() *AppAttachPackageInfoClient {
	return &AppAttachPackageInfoClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewApplicationGroupsClient creates a new instance of ApplicationGroupsClient.
func (c *ClientFactory) NewApplicationGroupsClient() *ApplicationGroupsClient {
	return &ApplicationGroupsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewApplicationsClient creates a new instance of ApplicationsClient.
func (c *ClientFactory) NewApplicationsClient() *ApplicationsClient {
	return &ApplicationsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewControlSessionHostUpdateClient creates a new instance of ControlSessionHostUpdateClient.
func (c *ClientFactory) NewControlSessionHostUpdateClient() *ControlSessionHostUpdateClient {
	return &ControlSessionHostUpdateClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewDesktopsClient creates a new instance of DesktopsClient.
func (c *ClientFactory) NewDesktopsClient() *DesktopsClient {
	return &DesktopsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewHostPoolsClient creates a new instance of HostPoolsClient.
func (c *ClientFactory) NewHostPoolsClient() *HostPoolsClient {
	return &HostPoolsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewInitiateSessionHostUpdateClient creates a new instance of InitiateSessionHostUpdateClient.
func (c *ClientFactory) NewInitiateSessionHostUpdateClient() *InitiateSessionHostUpdateClient {
	return &InitiateSessionHostUpdateClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewMSIXPackagesClient creates a new instance of MSIXPackagesClient.
func (c *ClientFactory) NewMSIXPackagesClient() *MSIXPackagesClient {
	return &MSIXPackagesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewMsixImagesClient creates a new instance of MsixImagesClient.
func (c *ClientFactory) NewMsixImagesClient() *MsixImagesClient {
	return &MsixImagesClient{
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

// NewPrivateEndpointConnectionsClient creates a new instance of PrivateEndpointConnectionsClient.
func (c *ClientFactory) NewPrivateEndpointConnectionsClient() *PrivateEndpointConnectionsClient {
	return &PrivateEndpointConnectionsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewPrivateLinkResourcesClient creates a new instance of PrivateLinkResourcesClient.
func (c *ClientFactory) NewPrivateLinkResourcesClient() *PrivateLinkResourcesClient {
	return &PrivateLinkResourcesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewScalingPlanPersonalSchedulesClient creates a new instance of ScalingPlanPersonalSchedulesClient.
func (c *ClientFactory) NewScalingPlanPersonalSchedulesClient() *ScalingPlanPersonalSchedulesClient {
	return &ScalingPlanPersonalSchedulesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewScalingPlanPooledSchedulesClient creates a new instance of ScalingPlanPooledSchedulesClient.
func (c *ClientFactory) NewScalingPlanPooledSchedulesClient() *ScalingPlanPooledSchedulesClient {
	return &ScalingPlanPooledSchedulesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewScalingPlansClient creates a new instance of ScalingPlansClient.
func (c *ClientFactory) NewScalingPlansClient() *ScalingPlansClient {
	return &ScalingPlansClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSessionHostClient creates a new instance of SessionHostClient.
func (c *ClientFactory) NewSessionHostClient() *SessionHostClient {
	return &SessionHostClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSessionHostConfigurationsClient creates a new instance of SessionHostConfigurationsClient.
func (c *ClientFactory) NewSessionHostConfigurationsClient() *SessionHostConfigurationsClient {
	return &SessionHostConfigurationsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSessionHostManagementsClient creates a new instance of SessionHostManagementsClient.
func (c *ClientFactory) NewSessionHostManagementsClient() *SessionHostManagementsClient {
	return &SessionHostManagementsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSessionHostManagementsUpdateStatusClient creates a new instance of SessionHostManagementsUpdateStatusClient.
func (c *ClientFactory) NewSessionHostManagementsUpdateStatusClient() *SessionHostManagementsUpdateStatusClient {
	return &SessionHostManagementsUpdateStatusClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSessionHostsClient creates a new instance of SessionHostsClient.
func (c *ClientFactory) NewSessionHostsClient() *SessionHostsClient {
	return &SessionHostsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewStartMenuItemsClient creates a new instance of StartMenuItemsClient.
func (c *ClientFactory) NewStartMenuItemsClient() *StartMenuItemsClient {
	return &StartMenuItemsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewUserSessionsClient creates a new instance of UserSessionsClient.
func (c *ClientFactory) NewUserSessionsClient() *UserSessionsClient {
	return &UserSessionsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewWorkspacesClient creates a new instance of WorkspacesClient.
func (c *ClientFactory) NewWorkspacesClient() *WorkspacesClient {
	return &WorkspacesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}
