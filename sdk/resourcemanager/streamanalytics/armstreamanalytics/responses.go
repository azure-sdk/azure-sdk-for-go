//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstreamanalytics

// ClustersClientCreateOrUpdateResponse contains the response from method ClustersClient.BeginCreateOrUpdate.
type ClustersClientCreateOrUpdateResponse struct {
	// A Stream Analytics Cluster object
	Cluster
}

// ClustersClientDeleteResponse contains the response from method ClustersClient.BeginDelete.
type ClustersClientDeleteResponse struct {
	// placeholder for future response values
}

// ClustersClientGetResponse contains the response from method ClustersClient.Get.
type ClustersClientGetResponse struct {
	// A Stream Analytics Cluster object
	Cluster
}

// ClustersClientListByResourceGroupResponse contains the response from method ClustersClient.NewListByResourceGroupPager.
type ClustersClientListByResourceGroupResponse struct {
	// A list of clusters populated by a 'list' operation.
	ClusterListResult
}

// ClustersClientListBySubscriptionResponse contains the response from method ClustersClient.NewListBySubscriptionPager.
type ClustersClientListBySubscriptionResponse struct {
	// A list of clusters populated by a 'list' operation.
	ClusterListResult
}

// ClustersClientListStreamingJobsResponse contains the response from method ClustersClient.NewListStreamingJobsPager.
type ClustersClientListStreamingJobsResponse struct {
	// A list of streaming jobs. Populated by a List operation.
	ClusterJobListResult
}

// ClustersClientUpdateResponse contains the response from method ClustersClient.BeginUpdate.
type ClustersClientUpdateResponse struct {
	// A Stream Analytics Cluster object
	Cluster
}

// FunctionsClientCreateOrReplaceResponse contains the response from method FunctionsClient.CreateOrReplace.
type FunctionsClientCreateOrReplaceResponse struct {
	// A function object, containing all information associated with the named function. All functions are contained under a streaming
	// job.
	Function

	// ETag contains the information returned from the ETag header response.
	ETag *string
}

// FunctionsClientDeleteResponse contains the response from method FunctionsClient.Delete.
type FunctionsClientDeleteResponse struct {
	// placeholder for future response values
}

// FunctionsClientGetResponse contains the response from method FunctionsClient.Get.
type FunctionsClientGetResponse struct {
	// A function object, containing all information associated with the named function. All functions are contained under a streaming
	// job.
	Function

	// ETag contains the information returned from the ETag header response.
	ETag *string
}

// FunctionsClientListByStreamingJobResponse contains the response from method FunctionsClient.NewListByStreamingJobPager.
type FunctionsClientListByStreamingJobResponse struct {
	// Object containing a list of functions under a streaming job.
	FunctionListResult
}

// FunctionsClientRetrieveDefaultDefinitionResponse contains the response from method FunctionsClient.RetrieveDefaultDefinition.
type FunctionsClientRetrieveDefaultDefinitionResponse struct {
	// A function object, containing all information associated with the named function. All functions are contained under a streaming
	// job.
	Function
}

// FunctionsClientTestResponse contains the response from method FunctionsClient.BeginTest.
type FunctionsClientTestResponse struct {
	// Describes the status of the test operation along with error information, if applicable.
	ResourceTestStatus
}

// FunctionsClientUpdateResponse contains the response from method FunctionsClient.Update.
type FunctionsClientUpdateResponse struct {
	// A function object, containing all information associated with the named function. All functions are contained under a streaming
	// job.
	Function

	// ETag contains the information returned from the ETag header response.
	ETag *string
}

// InputsClientCreateOrReplaceResponse contains the response from method InputsClient.CreateOrReplace.
type InputsClientCreateOrReplaceResponse struct {
	// An input object, containing all information associated with the named input. All inputs are contained under a streaming
	// job.
	Input

	// ETag contains the information returned from the ETag header response.
	ETag *string
}

// InputsClientDeleteResponse contains the response from method InputsClient.Delete.
type InputsClientDeleteResponse struct {
	// placeholder for future response values
}

// InputsClientGetResponse contains the response from method InputsClient.Get.
type InputsClientGetResponse struct {
	// An input object, containing all information associated with the named input. All inputs are contained under a streaming
	// job.
	Input

	// ETag contains the information returned from the ETag header response.
	ETag *string
}

// InputsClientListByStreamingJobResponse contains the response from method InputsClient.NewListByStreamingJobPager.
type InputsClientListByStreamingJobResponse struct {
	// Object containing a list of inputs under a streaming job.
	InputListResult
}

// InputsClientTestResponse contains the response from method InputsClient.BeginTest.
type InputsClientTestResponse struct {
	// Describes the status of the test operation along with error information, if applicable.
	ResourceTestStatus
}

// InputsClientUpdateResponse contains the response from method InputsClient.Update.
type InputsClientUpdateResponse struct {
	// An input object, containing all information associated with the named input. All inputs are contained under a streaming
	// job.
	Input

	// ETag contains the information returned from the ETag header response.
	ETag *string
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// Result of the request to list Stream Analytics operations. It contains a list of operations and a URL link to get the next
	// set of results.
	OperationListResult
}

// OutputsClientCreateOrReplaceResponse contains the response from method OutputsClient.CreateOrReplace.
type OutputsClientCreateOrReplaceResponse struct {
	// An output object, containing all information associated with the named output. All outputs are contained under a streaming
	// job.
	Output

	// ETag contains the information returned from the ETag header response.
	ETag *string
}

// OutputsClientDeleteResponse contains the response from method OutputsClient.Delete.
type OutputsClientDeleteResponse struct {
	// placeholder for future response values
}

// OutputsClientGetResponse contains the response from method OutputsClient.Get.
type OutputsClientGetResponse struct {
	// An output object, containing all information associated with the named output. All outputs are contained under a streaming
	// job.
	Output

	// ETag contains the information returned from the ETag header response.
	ETag *string
}

// OutputsClientListByStreamingJobResponse contains the response from method OutputsClient.NewListByStreamingJobPager.
type OutputsClientListByStreamingJobResponse struct {
	// Object containing a list of outputs under a streaming job.
	OutputListResult
}

// OutputsClientTestResponse contains the response from method OutputsClient.BeginTest.
type OutputsClientTestResponse struct {
	// Describes the status of the test operation along with error information, if applicable.
	ResourceTestStatus
}

// OutputsClientUpdateResponse contains the response from method OutputsClient.Update.
type OutputsClientUpdateResponse struct {
	// An output object, containing all information associated with the named output. All outputs are contained under a streaming
	// job.
	Output

	// ETag contains the information returned from the ETag header response.
	ETag *string
}

// PrivateEndpointsClientCreateOrUpdateResponse contains the response from method PrivateEndpointsClient.CreateOrUpdate.
type PrivateEndpointsClientCreateOrUpdateResponse struct {
	// Complete information about the private endpoint.
	PrivateEndpoint
}

// PrivateEndpointsClientDeleteResponse contains the response from method PrivateEndpointsClient.BeginDelete.
type PrivateEndpointsClientDeleteResponse struct {
	// placeholder for future response values
}

// PrivateEndpointsClientGetResponse contains the response from method PrivateEndpointsClient.Get.
type PrivateEndpointsClientGetResponse struct {
	// Complete information about the private endpoint.
	PrivateEndpoint
}

// PrivateEndpointsClientListByClusterResponse contains the response from method PrivateEndpointsClient.NewListByClusterPager.
type PrivateEndpointsClientListByClusterResponse struct {
	// A list of private endpoints.
	PrivateEndpointListResult
}

// StreamingJobsClientCreateOrReplaceResponse contains the response from method StreamingJobsClient.BeginCreateOrReplace.
type StreamingJobsClientCreateOrReplaceResponse struct {
	// A streaming job object, containing all information associated with the named streaming job.
	StreamingJob
}

// StreamingJobsClientDeleteResponse contains the response from method StreamingJobsClient.BeginDelete.
type StreamingJobsClientDeleteResponse struct {
	// placeholder for future response values
}

// StreamingJobsClientGetResponse contains the response from method StreamingJobsClient.Get.
type StreamingJobsClientGetResponse struct {
	// A streaming job object, containing all information associated with the named streaming job.
	StreamingJob

	// ETag contains the information returned from the ETag header response.
	ETag *string
}

// StreamingJobsClientListByResourceGroupResponse contains the response from method StreamingJobsClient.NewListByResourceGroupPager.
type StreamingJobsClientListByResourceGroupResponse struct {
	// Object containing a list of streaming jobs.
	StreamingJobListResult
}

// StreamingJobsClientListResponse contains the response from method StreamingJobsClient.NewListPager.
type StreamingJobsClientListResponse struct {
	// Object containing a list of streaming jobs.
	StreamingJobListResult
}

// StreamingJobsClientScaleResponse contains the response from method StreamingJobsClient.BeginScale.
type StreamingJobsClientScaleResponse struct {
	// placeholder for future response values
}

// StreamingJobsClientStartResponse contains the response from method StreamingJobsClient.BeginStart.
type StreamingJobsClientStartResponse struct {
	// placeholder for future response values
}

// StreamingJobsClientStopResponse contains the response from method StreamingJobsClient.BeginStop.
type StreamingJobsClientStopResponse struct {
	// placeholder for future response values
}

// StreamingJobsClientUpdateResponse contains the response from method StreamingJobsClient.Update.
type StreamingJobsClientUpdateResponse struct {
	// A streaming job object, containing all information associated with the named streaming job.
	StreamingJob

	// ETag contains the information returned from the ETag header response.
	ETag *string
}

// SubscriptionsClientListQuotasResponse contains the response from method SubscriptionsClient.ListQuotas.
type SubscriptionsClientListQuotasResponse struct {
	// Result of the GetQuotas operation. It contains a list of quotas for the subscription in a particular region.
	SubscriptionQuotasListResult
}

// TransformationsClientCreateOrReplaceResponse contains the response from method TransformationsClient.CreateOrReplace.
type TransformationsClientCreateOrReplaceResponse struct {
	// A transformation object, containing all information associated with the named transformation. All transformations are contained
	// under a streaming job.
	Transformation

	// ETag contains the information returned from the ETag header response.
	ETag *string
}

// TransformationsClientGetResponse contains the response from method TransformationsClient.Get.
type TransformationsClientGetResponse struct {
	// A transformation object, containing all information associated with the named transformation. All transformations are contained
	// under a streaming job.
	Transformation

	// ETag contains the information returned from the ETag header response.
	ETag *string
}

// TransformationsClientUpdateResponse contains the response from method TransformationsClient.Update.
type TransformationsClientUpdateResponse struct {
	// A transformation object, containing all information associated with the named transformation. All transformations are contained
	// under a streaming job.
	Transformation

	// ETag contains the information returned from the ETag header response.
	ETag *string
}
