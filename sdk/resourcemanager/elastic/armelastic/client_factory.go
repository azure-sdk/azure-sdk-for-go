//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Code generated by Microsoft (R) AutoRest Code Generator.Changes may cause incorrect behavior and will be lost if the code
// is regenerated.
// Code generated by @autorest/go. DO NOT EDIT.

package armelastic

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
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
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

// NewAllTrafficFiltersClient creates a new instance of AllTrafficFiltersClient.
func (c *ClientFactory) NewAllTrafficFiltersClient() *AllTrafficFiltersClient {
	subClient, _ := NewAllTrafficFiltersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewAssociateTrafficFilterClient creates a new instance of AssociateTrafficFilterClient.
func (c *ClientFactory) NewAssociateTrafficFilterClient() *AssociateTrafficFilterClient {
	subClient, _ := NewAssociateTrafficFilterClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewBillingInfoClient creates a new instance of BillingInfoClient.
func (c *ClientFactory) NewBillingInfoClient() *BillingInfoClient {
	subClient, _ := NewBillingInfoClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewConnectedPartnerResourcesClient creates a new instance of ConnectedPartnerResourcesClient.
func (c *ClientFactory) NewConnectedPartnerResourcesClient() *ConnectedPartnerResourcesClient {
	subClient, _ := NewConnectedPartnerResourcesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewCreateAndAssociateIPFilterClient creates a new instance of CreateAndAssociateIPFilterClient.
func (c *ClientFactory) NewCreateAndAssociateIPFilterClient() *CreateAndAssociateIPFilterClient {
	subClient, _ := NewCreateAndAssociateIPFilterClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewCreateAndAssociatePLFilterClient creates a new instance of CreateAndAssociatePLFilterClient.
func (c *ClientFactory) NewCreateAndAssociatePLFilterClient() *CreateAndAssociatePLFilterClient {
	subClient, _ := NewCreateAndAssociatePLFilterClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewDeploymentInfoClient creates a new instance of DeploymentInfoClient.
func (c *ClientFactory) NewDeploymentInfoClient() *DeploymentInfoClient {
	subClient, _ := NewDeploymentInfoClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewDetachAndDeleteTrafficFilterClient creates a new instance of DetachAndDeleteTrafficFilterClient.
func (c *ClientFactory) NewDetachAndDeleteTrafficFilterClient() *DetachAndDeleteTrafficFilterClient {
	subClient, _ := NewDetachAndDeleteTrafficFilterClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewDetachTrafficFilterClient creates a new instance of DetachTrafficFilterClient.
func (c *ClientFactory) NewDetachTrafficFilterClient() *DetachTrafficFilterClient {
	subClient, _ := NewDetachTrafficFilterClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewExternalUserClient creates a new instance of ExternalUserClient.
func (c *ClientFactory) NewExternalUserClient() *ExternalUserClient {
	subClient, _ := NewExternalUserClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewListAssociatedTrafficFiltersClient creates a new instance of ListAssociatedTrafficFiltersClient.
func (c *ClientFactory) NewListAssociatedTrafficFiltersClient() *ListAssociatedTrafficFiltersClient {
	subClient, _ := NewListAssociatedTrafficFiltersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewMonitorClient creates a new instance of MonitorClient.
func (c *ClientFactory) NewMonitorClient() *MonitorClient {
	subClient, _ := NewMonitorClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewMonitoredResourcesClient creates a new instance of MonitoredResourcesClient.
func (c *ClientFactory) NewMonitoredResourcesClient() *MonitoredResourcesClient {
	subClient, _ := NewMonitoredResourcesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewMonitorsClient creates a new instance of MonitorsClient.
func (c *ClientFactory) NewMonitorsClient() *MonitorsClient {
	subClient, _ := NewMonitorsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewOpenAIClient creates a new instance of OpenAIClient.
func (c *ClientFactory) NewOpenAIClient() *OpenAIClient {
	subClient, _ := NewOpenAIClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewOperationsClient creates a new instance of OperationsClient.
func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	subClient, _ := NewOperationsClient(c.credential, c.options)
	return subClient
}

// NewOrganizationsClient creates a new instance of OrganizationsClient.
func (c *ClientFactory) NewOrganizationsClient() *OrganizationsClient {
	subClient, _ := NewOrganizationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewTagRulesClient creates a new instance of TagRulesClient.
func (c *ClientFactory) NewTagRulesClient() *TagRulesClient {
	subClient, _ := NewTagRulesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewTrafficFiltersClient creates a new instance of TrafficFiltersClient.
func (c *ClientFactory) NewTrafficFiltersClient() *TrafficFiltersClient {
	subClient, _ := NewTrafficFiltersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewUpgradableVersionsClient creates a new instance of UpgradableVersionsClient.
func (c *ClientFactory) NewUpgradableVersionsClient() *UpgradableVersionsClient {
	subClient, _ := NewUpgradableVersionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewVMCollectionClient creates a new instance of VMCollectionClient.
func (c *ClientFactory) NewVMCollectionClient() *VMCollectionClient {
	subClient, _ := NewVMCollectionClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewVMHostClient creates a new instance of VMHostClient.
func (c *ClientFactory) NewVMHostClient() *VMHostClient {
	subClient, _ := NewVMHostClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewVMIngestionClient creates a new instance of VMIngestionClient.
func (c *ClientFactory) NewVMIngestionClient() *VMIngestionClient {
	subClient, _ := NewVMIngestionClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewVersionsClient creates a new instance of VersionsClient.
func (c *ClientFactory) NewVersionsClient() *VersionsClient {
	subClient, _ := NewVersionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}
