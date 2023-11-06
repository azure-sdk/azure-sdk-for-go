//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmigrate

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
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClientFactory(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ClientFactory, error) {
	_, err := arm.NewClient(moduleName+".ClientFactory", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	return &ClientFactory{
		subscriptionID: subscriptionID, credential: credential,
		options: options.Clone(),
	}, nil
}

func (c *ClientFactory) NewAssessedMachinesOperationsClient() *AssessedMachinesOperationsClient {
	subClient, _ := NewAssessedMachinesOperationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAssessedSQLDatabaseV2OperationsClient() *AssessedSQLDatabaseV2OperationsClient {
	subClient, _ := NewAssessedSQLDatabaseV2OperationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAssessedSQLInstanceV2OperationsClient() *AssessedSQLInstanceV2OperationsClient {
	subClient, _ := NewAssessedSQLInstanceV2OperationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAssessedSQLMachinesOperationsClient() *AssessedSQLMachinesOperationsClient {
	subClient, _ := NewAssessedSQLMachinesOperationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAssessedSQLRecommendedEntityOperationsClient() *AssessedSQLRecommendedEntityOperationsClient {
	subClient, _ := NewAssessedSQLRecommendedEntityOperationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAssessmentOptionsOperationsClient() *AssessmentOptionsOperationsClient {
	subClient, _ := NewAssessmentOptionsOperationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAssessmentProjectSummaryOperationsClient() *AssessmentProjectSummaryOperationsClient {
	subClient, _ := NewAssessmentProjectSummaryOperationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAssessmentProjectsOperationsClient() *AssessmentProjectsOperationsClient {
	subClient, _ := NewAssessmentProjectsOperationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAssessmentsOperationsClient() *AssessmentsOperationsClient {
	subClient, _ := NewAssessmentsOperationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAvsAssessedMachinesOperationsClient() *AvsAssessedMachinesOperationsClient {
	subClient, _ := NewAvsAssessedMachinesOperationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAvsAssessmentOptionsOperationsClient() *AvsAssessmentOptionsOperationsClient {
	subClient, _ := NewAvsAssessmentOptionsOperationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAvsAssessmentsOperationsClient() *AvsAssessmentsOperationsClient {
	subClient, _ := NewAvsAssessmentsOperationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewGroupsOperationsClient() *GroupsOperationsClient {
	subClient, _ := NewGroupsOperationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewHypervCollectorsOperationsClient() *HypervCollectorsOperationsClient {
	subClient, _ := NewHypervCollectorsOperationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewImportCollectorsOperationsClient() *ImportCollectorsOperationsClient {
	subClient, _ := NewImportCollectorsOperationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewMachinesOperationsClient() *MachinesOperationsClient {
	subClient, _ := NewMachinesOperationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	subClient, _ := NewOperationsClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewPrivateEndpointConnectionOperationsClient() *PrivateEndpointConnectionOperationsClient {
	subClient, _ := NewPrivateEndpointConnectionOperationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewPrivateLinkResourceOperationsClient() *PrivateLinkResourceOperationsClient {
	subClient, _ := NewPrivateLinkResourceOperationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewSQLAssessmentOptionsOperationsClient() *SQLAssessmentOptionsOperationsClient {
	subClient, _ := NewSQLAssessmentOptionsOperationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewSQLAssessmentV2OperationsClient() *SQLAssessmentV2OperationsClient {
	subClient, _ := NewSQLAssessmentV2OperationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewSQLAssessmentV2SummaryOperationsClient() *SQLAssessmentV2SummaryOperationsClient {
	subClient, _ := NewSQLAssessmentV2SummaryOperationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewSQLCollectorOperationsClient() *SQLCollectorOperationsClient {
	subClient, _ := NewSQLCollectorOperationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewServerCollectorsOperationsClient() *ServerCollectorsOperationsClient {
	subClient, _ := NewServerCollectorsOperationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewVmwareCollectorsOperationsClient() *VmwareCollectorsOperationsClient {
	subClient, _ := NewVmwareCollectorsOperationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}
