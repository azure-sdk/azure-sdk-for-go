//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcostmanagement

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
)

// ClientFactory is a client factory used to create any client in this module.
// Don't use this type directly, use NewClientFactory instead.
type ClientFactory struct {
	filter     *string
	ifMatch    *string
	credential azcore.TokenCredential
	options    *arm.ClientOptions
}

// NewClientFactory creates a new instance of ClientFactory with the specified values.
// The parameter values will be propagated to any client created from this factory.
//   - filter - OData filter option. May be used to filter budgets by properties/category. The filter supports 'eq' only.
//   - ifMatch - ETag of the Entity. Not required when creating an entity, but required when updating an entity.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClientFactory(filter *string, ifMatch *string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ClientFactory, error) {
	_, err := arm.NewClient(moduleName+".ClientFactory", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	return &ClientFactory{
		filter: filter, ifMatch: ifMatch, credential: credential,
		options: options.Clone(),
	}, nil
}

func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	subClient, _ := NewOperationsClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewBenefitRecommendationsClient() *BenefitRecommendationsClient {
	subClient, _ := NewBenefitRecommendationsClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewBenefitUtilizationSummariesClient() *BenefitUtilizationSummariesClient {
	subClient, _ := NewBenefitUtilizationSummariesClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewBudgetsClient() *BudgetsClient {
	subClient, _ := NewBudgetsClient(c.filter, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewExportsClient() *ExportsClient {
	subClient, _ := NewExportsClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewGenerateCostDetailsReportClient() *GenerateCostDetailsReportClient {
	subClient, _ := NewGenerateCostDetailsReportClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewGenerateDetailedCostReportClient() *GenerateDetailedCostReportClient {
	subClient, _ := NewGenerateDetailedCostReportClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewGenerateDetailedCostReportOperationResultsClient() *GenerateDetailedCostReportOperationResultsClient {
	subClient, _ := NewGenerateDetailedCostReportOperationResultsClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewGenerateDetailedCostReportOperationStatusClient() *GenerateDetailedCostReportOperationStatusClient {
	subClient, _ := NewGenerateDetailedCostReportOperationStatusClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewViewsClient() *ViewsClient {
	subClient, _ := NewViewsClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAlertsClient() *AlertsClient {
	subClient, _ := NewAlertsClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewForecastClient() *ForecastClient {
	subClient, _ := NewForecastClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewDimensionsClient() *DimensionsClient {
	subClient, _ := NewDimensionsClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewQueryClient() *QueryClient {
	subClient, _ := NewQueryClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewGenerateReservationDetailsReportClient() *GenerateReservationDetailsReportClient {
	subClient, _ := NewGenerateReservationDetailsReportClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewPriceSheetClient() *PriceSheetClient {
	subClient, _ := NewPriceSheetClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewScheduledActionsClient() *ScheduledActionsClient {
	subClient, _ := NewScheduledActionsClient(c.ifMatch, c.credential, c.options)
	return subClient
}
