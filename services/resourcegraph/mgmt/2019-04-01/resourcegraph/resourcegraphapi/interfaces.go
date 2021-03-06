package resourcegraphapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/resourcegraph/mgmt/2019-04-01/resourcegraph"
)

// BaseClientAPI contains the set of methods on the BaseClient type.
type BaseClientAPI interface {
	Resources(ctx context.Context, query resourcegraph.QueryRequest) (result resourcegraph.QueryResponse, err error)
}

var _ BaseClientAPI = (*resourcegraph.BaseClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result resourcegraph.OperationListResult, err error)
}

var _ OperationsClientAPI = (*resourcegraph.OperationsClient)(nil)
