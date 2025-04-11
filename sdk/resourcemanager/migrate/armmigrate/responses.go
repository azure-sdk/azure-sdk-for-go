// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armmigrate

// AssessedMachinesV2OperationsClientGetResponse contains the response from method AssessedMachinesV2OperationsClient.Get.
type AssessedMachinesV2OperationsClientGetResponse struct {
	// Machine assessment V2 Assessed Machine resource.
	AssessedMachineV2
}

// AssessedMachinesV2OperationsClientListByParentResponse contains the response from method AssessedMachinesV2OperationsClient.NewListByParentPager.
type AssessedMachinesV2OperationsClientListByParentResponse struct {
	// The response of a AssessedMachineV2 list operation.
	AssessedMachineV2ListResult
}

// MachineAssessmentV2SummaryOperationsClientGetResponse contains the response from method MachineAssessmentV2SummaryOperationsClient.Get.
type MachineAssessmentV2SummaryOperationsClientGetResponse struct {
	// Machine Assessment REST resource.
	MachineAssessmentV2Summary
}

// MachineAssessmentV2SummaryOperationsClientListByParentResponse contains the response from method MachineAssessmentV2SummaryOperationsClient.NewListByParentPager.
type MachineAssessmentV2SummaryOperationsClientListByParentResponse struct {
	// The response of a MachineAssessmentV2Summary list operation.
	MachineAssessmentV2SummaryListResult
}

// MachineAssessmentsV2OperationsClientCreateResponse contains the response from method MachineAssessmentsV2OperationsClient.BeginCreate.
type MachineAssessmentsV2OperationsClientCreateResponse struct {
	// Machine assessment V2 resource.
	MachineAssessmentV2
}

// MachineAssessmentsV2OperationsClientDeleteResponse contains the response from method MachineAssessmentsV2OperationsClient.Delete.
type MachineAssessmentsV2OperationsClientDeleteResponse struct {
	// placeholder for future response values
}

// MachineAssessmentsV2OperationsClientDownloadURLResponse contains the response from method MachineAssessmentsV2OperationsClient.BeginDownloadURL.
type MachineAssessmentsV2OperationsClientDownloadURLResponse struct {
	// Data model of Download URL for assessment report.
	DownloadURL
}

// MachineAssessmentsV2OperationsClientGetResponse contains the response from method MachineAssessmentsV2OperationsClient.Get.
type MachineAssessmentsV2OperationsClientGetResponse struct {
	// Machine assessment V2 resource.
	MachineAssessmentV2
}

// MachineAssessmentsV2OperationsClientListByParentResponse contains the response from method MachineAssessmentsV2OperationsClient.NewListByParentPager.
type MachineAssessmentsV2OperationsClientListByParentResponse struct {
	// The response of a MachineAssessmentV2 list operation.
	MachineAssessmentV2ListResult
}

// MachineGraphAssessmentOptionsOperationsClientGetResponse contains the response from method MachineGraphAssessmentOptionsOperationsClient.Get.
type MachineGraphAssessmentOptionsOperationsClientGetResponse struct {
	// Machine Assessment REST resource.
	MachineGraphAssessmentOptions
}

// MachineGraphAssessmentOptionsOperationsClientListByParentResponse contains the response from method MachineGraphAssessmentOptionsOperationsClient.NewListByParentPager.
type MachineGraphAssessmentOptionsOperationsClientListByParentResponse struct {
	// The response of a MachineGraphAssessmentOptions list operation.
	MachineGraphAssessmentOptionsListResult
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to get the next set of results.
	OperationListResult
}
