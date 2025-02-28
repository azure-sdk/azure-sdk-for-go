// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappcomplianceautomation

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

// NewEvidenceClient creates a new instance of EvidenceClient.
func (c *ClientFactory) NewEvidenceClient() *EvidenceClient {
	return &EvidenceClient{
		internal: c.internal,
	}
}

// NewOperationsClient creates a new instance of OperationsClient.
func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	return &OperationsClient{
		internal: c.internal,
	}
}

// NewProviderActionsClient creates a new instance of ProviderActionsClient.
func (c *ClientFactory) NewProviderActionsClient() *ProviderActionsClient {
	return &ProviderActionsClient{
		internal: c.internal,
	}
}

// NewReportClient creates a new instance of ReportClient.
func (c *ClientFactory) NewReportClient() *ReportClient {
	return &ReportClient{
		internal: c.internal,
	}
}

// NewScopingConfigurationClient creates a new instance of ScopingConfigurationClient.
func (c *ClientFactory) NewScopingConfigurationClient() *ScopingConfigurationClient {
	return &ScopingConfigurationClient{
		internal: c.internal,
	}
}

// NewSnapshotClient creates a new instance of SnapshotClient.
func (c *ClientFactory) NewSnapshotClient() *SnapshotClient {
	return &SnapshotClient{
		internal: c.internal,
	}
}

// NewWebhookClient creates a new instance of WebhookClient.
func (c *ClientFactory) NewWebhookClient() *WebhookClient {
	return &WebhookClient{
		internal: c.internal,
	}
}
