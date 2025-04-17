// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcommerce

// RateCardClientGetResponse contains the response from method RateCardClient.Get.
type RateCardClientGetResponse struct {
	// Price and Metadata information for resources
	ResourceRateCardInfo
}

// UsageAggregatesClientListResponse contains the response from method UsageAggregatesClient.NewListPager.
type UsageAggregatesClientListResponse struct {
	// Possible types are UsageAggregationListResult, ErrorObjectResponse
	Value any
}
