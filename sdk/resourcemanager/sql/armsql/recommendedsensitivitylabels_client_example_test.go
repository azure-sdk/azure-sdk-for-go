//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/33c4457b1d13f83965f4fe3367dca4a6df898100/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/SensitivityLabelsRecommendedUpdate.json
func ExampleRecommendedSensitivityLabelsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewRecommendedSensitivityLabelsClient().Update(ctx, "myRG", "myServer", "myDatabase", armsql.RecommendedSensitivityLabelUpdateList{
		Operations: []*armsql.RecommendedSensitivityLabelUpdate{
			{
				Properties: &armsql.RecommendedSensitivityLabelUpdateProperties{
					Schema: to.Ptr("dbo"),
					Column: to.Ptr("column1"),
					Op:     to.Ptr(armsql.RecommendedSensitivityLabelUpdateKindEnable),
					Table:  to.Ptr("table1"),
				},
			},
			{
				Properties: &armsql.RecommendedSensitivityLabelUpdateProperties{
					Schema: to.Ptr("dbo"),
					Column: to.Ptr("column2"),
					Op:     to.Ptr(armsql.RecommendedSensitivityLabelUpdateKindEnable),
					Table:  to.Ptr("table2"),
				},
			},
			{
				Properties: &armsql.RecommendedSensitivityLabelUpdateProperties{
					Schema: to.Ptr("dbo"),
					Column: to.Ptr("column3"),
					Op:     to.Ptr(armsql.RecommendedSensitivityLabelUpdateKindDisable),
					Table:  to.Ptr("table1"),
				},
			}},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}