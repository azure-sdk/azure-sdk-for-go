// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armimpactreporting_test

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/impactreporting/armimpactreporting"
	"log"
	"time"
)

// Generated from example definition: 2024-05-01-preview/Insights_Create.json
func ExampleInsightsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armimpactreporting.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewInsightsClient().Create(ctx, "impactid22", "insightId12", armimpactreporting.Insight{
		Properties: &armimpactreporting.InsightProperties{
			Content: &armimpactreporting.Content{
				Title:       to.Ptr("Impact Has been correlated to an outage"),
				Description: to.Ptr("At 2018-11-08T00:00:00Z UTC, your services dependent on these resources <link href=”…”>VM1</link> may have experienced an issue. <br/><div>We have identified an outage that affected these resources(s). You can look at outage information on <link href=\"https:// portal.azure.com/#view/Microsoft_Azure_Health/AzureHealthBrowseBlade/~/serviceIssues/trackingId/NL2W-VCZ\">NL2W-VCZ</link> link.<div>"),
			},
			Category:        to.Ptr("repair"),
			Status:          to.Ptr("resolved"),
			EventTime:       to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-15T04:00:00.009223Z"); return t }()),
			InsightUniqueID: to.Ptr("00000000-0000-0000-0000-000000000000"),
			Impact: &armimpactreporting.ImpactDetails{
				ImpactedResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource-rg/providers/Microsoft.Sql/sqlserver/dbservername"),
				StartTime:          to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-15T01:00:00.009223Z"); return t }()),
				ImpactID:           to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/microsoft.Impact/workloadImpacts/impactid22"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armimpactreporting.InsightsClientCreateResponse{
	// 	Insight: &armimpactreporting.Insight{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/microsoft.Impact/workloadImpacts/impactid22/insights/insightId12"),
	// 		Name: to.Ptr("insightId12"),
	// 		Type: to.Ptr("Microsoft.Impact/insights"),
	// 		Properties: &armimpactreporting.InsightProperties{
	// 			EventTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-15T04:00:00.009223Z"); return t}()),
	// 			Content: &armimpactreporting.Content{
	// 				Title: to.Ptr("Impact Has been correlated to an outage"),
	// 				Description: to.Ptr("At 2018-11-08T00:00:00Z UTC, your services dependent on these resources <link href=”…”>VM1</link> may have experienced an issue. <br/><div>We have identified an outage that affected these resources(s). You can look at outage information on <link href=\"https:// portal.azure.com/#view/Microsoft_Azure_Health/AzureHealthBrowseBlade/~/serviceIssues/trackingId/NL2W-VCZ\">NL2W-VCZ</link> link.<div>"),
	// 			},
	// 			Category: to.Ptr("repair"),
	// 			Status: to.Ptr("resolved"),
	// 			InsightUniqueID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			Impact: &armimpactreporting.ImpactDetails{
	// 				ImpactedResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource-rg/providers/Microsoft.Sql/sqlserver/dbservername"),
	// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-15T01:00:00.009223Z"); return t}()),
	// 				ImpactID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/microsoft.Impact/workloadImpacts/impactid22"),
	// 			},
	// 		},
	// 	},
	// }
}

// Generated from example definition: 2024-05-01-preview/Insights_Delete.json
func ExampleInsightsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armimpactreporting.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewInsightsClient().Delete(ctx, "impactid22", "insightId12", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armimpactreporting.InsightsClientDeleteResponse{
	// }
}

// Generated from example definition: 2024-05-01-preview/Insights_Get_diagnostics.json
func ExampleInsightsClient_Get_getInsightSampleForDiagnosticsCategory() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armimpactreporting.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewInsightsClient().Get(ctx, "impactid", "insight1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armimpactreporting.InsightsClientGetResponse{
	// 	Insight: &armimpactreporting.Insight{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Impact/workloadimpacts/impactid/insights/insight1"),
	// 		Name: to.Ptr("insight1"),
	// 		Type: to.Ptr("Microsoft.impact/workloadimpacts/insights"),
	// 		Properties: &armimpactreporting.InsightProperties{
	// 			Category: to.Ptr("Diagnostics"),
	// 			Impact: &armimpactreporting.ImpactDetails{
	// 				ImpactID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Impact/workloadimpacts/impactid"),
	// 				ImpactedResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Compute/virtualamchine/vm"),
	// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-11-08T00:00:00Z"); return t}()),
	// 				EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-11-08T00:00:00Z"); return t}()),
	// 			},
	// 			InsightUniqueID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			Content: &armimpactreporting.Content{
	// 				Title: to.Ptr("We ran diagnostics on your resource and found an issue"),
	// 				Description: to.Ptr("<!--issueDescription-->\n<p>The physical host node where your VM is running had a networking stack update. This might result in a brief connectivity loss.</p>\n<br/><table><tr><th align=\"left\">Resource</th><th align=\"left\">Impact Start Time</th><th align=\"left\">Impact End Time</th><th align=\"left\">Impact Duration(Timespan)</th></tr><tr><td align=\"left\">west-eur-VM</td><td align=\"left\">2024-02-18 01:09:31 UTC</td><td align=\"left\">2024-02-18 01:09:35 UTC</td><td align=\"left\">00:00:04.2690000</td></tr></table>\n<!--/issueDescription-->\n<p>Azure performs updates to improve reliability, performance, and security of the VMs. Azure chooses the least impactful method, which might result in a brief connectivity loss. We are continuously working to improve and reduce impact of our updates, and we apologize for any inconvenience this may have caused you.</p>\n<!--recommendedActions-->\n<h2><strong>Recommended Documents</strong></h2>\n<ul>\n<li>To prepare for VM maintenance events and reduce its impact, try using <a href=\"https://docs.microsoft.com/azure/virtual-machines/windows/scheduled-events\">Scheduled Events for Windows</a> or <a href=\"https://docs.microsoft.com/azure/virtual-machines/linux/scheduled-events\">Linux</a></li>\n<li>Learn more about Azure maintenance and configuring for high availability:\n<ul>\n<li><a href=\"https://docs.microsoft.com/azure/virtual-machines/maintenance-and-updates\">Maintenance and updates for virtual machines in Azure</a></li>\n<li><a href=\"https://docs.microsoft.com/azure/virtual-machines/windows/tutorial-availability-sets\">Configure availability of virtual machines</a></li>\n</ul>\n</li>\n<li>To troubleshoot this scenario in the future, see <a href=\"https://docs.microsoft.com/azure/resource-health/resource-health-overview\">Resource Health Center</a></li>\n</ul>\n"),
	// 			},
	// 		},
	// 	},
	// }
}

// Generated from example definition: 2024-05-01-preview/Insights_Get_mitigationAction.json
func ExampleInsightsClient_Get_getInsightSampleForMitigationActionCategory() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armimpactreporting.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewInsightsClient().Get(ctx, "impactId", "HPCUASucceeded", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armimpactreporting.InsightsClientGetResponse{
	// 	Insight: &armimpactreporting.Insight{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Impact/workloadimpacts/impactId/insights/HPCUASucceeded"),
	// 		Name: to.Ptr("HPCUASucceeded"),
	// 		Type: to.Ptr("Microsoft.impact/workloadimpacts/insights"),
	// 		Properties: &armimpactreporting.InsightProperties{
	// 			Category: to.Ptr("MitigationAction"),
	// 			EventTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-15T04:00:00.009223Z"); return t}()),
	// 			InsightUniqueID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			Impact: &armimpactreporting.ImpactDetails{
	// 				ImpactID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Impact/workloadimpacts/impactId"),
	// 				ImpactedResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Compute/virtualMachine/vm"),
	// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-11-08T00:00:00Z"); return t}()),
	// 				EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-11-08T00:00:00Z"); return t}()),
	// 			},
	// 			Content: &armimpactreporting.Content{
	// 				Title: to.Ptr("Node was flagged for inspection"),
	// 				Description: to.Ptr("At 2024-02-16 21:05:07 (UTC) an impact was reported on west-eur-VM indicating a potential disruption from Azure platform. <strong>Action Taken</strong> The hardware your VM was running on was flagged for inspection. We will conduct our investigation and take corrective action to prevent further impact on your workloads. We apologize for any disruption. Microsoft Azure Team"),
	// 			},
	// 			AdditionalDetails: &armimpactreporting.InsightPropertiesAdditionalDetails{
	// 			},
	// 		},
	// 	},
	// }
}

// Generated from example definition: 2024-05-01-preview/Insights_Get_servicehealth.json
func ExampleInsightsClient_Get_getInsightSampleForServiceHealthCategory() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armimpactreporting.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewInsightsClient().Get(ctx, "impactid", "insightname", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armimpactreporting.InsightsClientGetResponse{
	// 	Insight: &armimpactreporting.Insight{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/microsoft.Impact/workloadImpacts/impactname/insights/insightname"),
	// 		Name: to.Ptr("insightname"),
	// 		Type: to.Ptr("Microsoft.Impact/insights"),
	// 		Properties: &armimpactreporting.InsightProperties{
	// 			Category: to.Ptr("ServiceHealth"),
	// 			EventID: to.Ptr("ABC-123"),
	// 			EventTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-15T04:00:00.009223Z"); return t}()),
	// 			Impact: &armimpactreporting.ImpactDetails{
	// 				ImpactID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Impact/workloadimpacts/impactid"),
	// 				ImpactedResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.compute/virtualmachines/vm1"),
	// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-11-08T00:00:00Z"); return t}()),
	// 				EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-11-08T00:00:00Z"); return t}()),
	// 			},
	// 			InsightUniqueID: to.Ptr("a3d91a07-698b-4044-a230-e918252c4c59"),
	// 			Content: &armimpactreporting.Content{
	// 				Title: to.Ptr("Impact Has been correlated to an outage"),
	// 				Description: to.Ptr("On November 8, 2018, at 00:00 UTC, there may have been disruptions to your services linked to the resource <a href='...'>VM1<a>. We have pinpointed a service outage impacting these resources. For details on this outage, please refer to the service health information available at <a href='...'>service health</a>."),
	// 			},
	// 		},
	// 	},
	// }
}

// Generated from example definition: 2024-05-01-preview/Insights_ListBySubscription.json
func ExampleInsightsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armimpactreporting.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewInsightsClient().NewListBySubscriptionPager("impactid22", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page = armimpactreporting.InsightsClientListBySubscriptionResponse{
		// 	InsightListResult: armimpactreporting.InsightListResult{
		// 		Value: []*armimpactreporting.Insight{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/microsoft.Impact/workloadImpacts/impactid22/insights/insightId12"),
		// 				Name: to.Ptr("insightId12"),
		// 				Type: to.Ptr("Microsoft.Impact/insights"),
		// 				Properties: &armimpactreporting.InsightProperties{
		// 					EventTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-15T04:00:00.009223Z"); return t}()),
		// 					Content: &armimpactreporting.Content{
		// 						Title: to.Ptr("Impact Has been correlated to an outage"),
		// 						Description: to.Ptr("At 2018-11-08T00:00:00Z UTC, your services dependent on these resources <link href=”…”>VM1</link> may have experienced an issue. <br/><div>We have identified an outage that affected these resources(s). You can look at outage information on <link href=\"https:// portal.azure.com/#view/Microsoft_Azure_Health/AzureHealthBrowseBlade/~/serviceIssues/trackingId/NL2W-VCZ\">NL2W-VCZ</link> link.<div>"),
		// 					},
		// 					Category: to.Ptr("repair"),
		// 					Status: to.Ptr("resolved"),
		// 					InsightUniqueID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 					Impact: &armimpactreporting.ImpactDetails{
		// 						ImpactedResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource-rg/providers/Microsoft.Sql/sqlserver/dbservername"),
		// 						StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-15T01:00:00.009223Z"); return t}()),
		// 						ImpactID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/microsoft.Impact/workloadImpacts/impactid22"),
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
