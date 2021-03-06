package monitoringapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/synapse/2019-11-01-preview/monitoring"
)

// ClientAPI contains the set of methods on the Client type.
type ClientAPI interface {
	GetSparkJobList(ctx context.Context, APIVersion string, xMsClientRequestID string) (result monitoring.SparkJobListViewResponse, err error)
	GetSQLJobQueryString(ctx context.Context, APIVersion string, xMsClientRequestID string, filter string, orderby string, skip string) (result monitoring.SQLQueryStringDataModel, err error)
}

var _ ClientAPI = (*monitoring.Client)(nil)
