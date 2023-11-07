//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armtestbase

// AccountsClientBeginCreateOptions contains the optional parameters for the AccountsClient.BeginCreate method.
type AccountsClientBeginCreateOptions struct {
	// The flag indicating if we would like to restore the Test Base Accounts which were soft deleted before.
	Restore *bool

	// Resumes the LRO from the provided token.
	ResumeToken string
}

// AccountsClientBeginDeleteOptions contains the optional parameters for the AccountsClient.BeginDelete method.
type AccountsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// AccountsClientBeginOffboardOptions contains the optional parameters for the AccountsClient.BeginOffboard method.
type AccountsClientBeginOffboardOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// AccountsClientBeginUpdateOptions contains the optional parameters for the AccountsClient.BeginUpdate method.
type AccountsClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// AccountsClientCheckPackageNameAvailabilityOptions contains the optional parameters for the AccountsClient.CheckPackageNameAvailability
// method.
type AccountsClientCheckPackageNameAvailabilityOptions struct {
	// placeholder for future optional parameters
}

// AccountsClientGetFileUploadURLOptions contains the optional parameters for the AccountsClient.GetFileUploadURL method.
type AccountsClientGetFileUploadURLOptions struct {
	// Parameters supplied to the Test Base Account GetFileUploadURL operation.
	Parameters *GetFileUploadURLParameters
}

// AccountsClientGetOptions contains the optional parameters for the AccountsClient.Get method.
type AccountsClientGetOptions struct {
	// placeholder for future optional parameters
}

// AccountsClientListByResourceGroupOptions contains the optional parameters for the AccountsClient.NewListByResourceGroupPager
// method.
type AccountsClientListByResourceGroupOptions struct {
	// The flag indicating if we need to include the Test Base Accounts which were soft deleted before.
	GetDeleted *bool
}

// AnalysisResultsClientGetOptions contains the optional parameters for the AnalysisResultsClient.Get method.
type AnalysisResultsClientGetOptions struct {
	// placeholder for future optional parameters
}

// AnalysisResultsClientListOptions contains the optional parameters for the AnalysisResultsClient.NewListPager method.
type AnalysisResultsClientListOptions struct {
	// placeholder for future optional parameters
}

// AvailableOSClientGetOptions contains the optional parameters for the AvailableOSClient.Get method.
type AvailableOSClientGetOptions struct {
	// placeholder for future optional parameters
}

// AvailableOSClientListOptions contains the optional parameters for the AvailableOSClient.NewListPager method.
type AvailableOSClientListOptions struct {
	// placeholder for future optional parameters
}

// BillingHubServiceClientGetFreeHourBalanceOptions contains the optional parameters for the BillingHubServiceClient.GetFreeHourBalance
// method.
type BillingHubServiceClientGetFreeHourBalanceOptions struct {
	// placeholder for future optional parameters
}

// BillingHubServiceClientGetUsageOptions contains the optional parameters for the BillingHubServiceClient.GetUsage method.
type BillingHubServiceClientGetUsageOptions struct {
	GetUsageRequest *BillingHubGetUsageRequest
}

// CustomerEventsClientBeginCreateOptions contains the optional parameters for the CustomerEventsClient.BeginCreate method.
type CustomerEventsClientBeginCreateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// CustomerEventsClientBeginDeleteOptions contains the optional parameters for the CustomerEventsClient.BeginDelete method.
type CustomerEventsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// CustomerEventsClientGetOptions contains the optional parameters for the CustomerEventsClient.Get method.
type CustomerEventsClientGetOptions struct {
	// placeholder for future optional parameters
}

// CustomerEventsClientListByTestBaseAccountOptions contains the optional parameters for the CustomerEventsClient.NewListByTestBaseAccountPager
// method.
type CustomerEventsClientListByTestBaseAccountOptions struct {
	// placeholder for future optional parameters
}

// EmailEventsClientGetOptions contains the optional parameters for the EmailEventsClient.Get method.
type EmailEventsClientGetOptions struct {
	// placeholder for future optional parameters
}

// EmailEventsClientListOptions contains the optional parameters for the EmailEventsClient.NewListPager method.
type EmailEventsClientListOptions struct {
	// placeholder for future optional parameters
}

// FavoriteProcessesClientCreateOptions contains the optional parameters for the FavoriteProcessesClient.Create method.
type FavoriteProcessesClientCreateOptions struct {
	// placeholder for future optional parameters
}

// FavoriteProcessesClientDeleteOptions contains the optional parameters for the FavoriteProcessesClient.Delete method.
type FavoriteProcessesClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// FavoriteProcessesClientGetOptions contains the optional parameters for the FavoriteProcessesClient.Get method.
type FavoriteProcessesClientGetOptions struct {
	// placeholder for future optional parameters
}

// FavoriteProcessesClientListOptions contains the optional parameters for the FavoriteProcessesClient.NewListPager method.
type FavoriteProcessesClientListOptions struct {
	// placeholder for future optional parameters
}

// FlightingRingsClientGetOptions contains the optional parameters for the FlightingRingsClient.Get method.
type FlightingRingsClientGetOptions struct {
	// placeholder for future optional parameters
}

// FlightingRingsClientListOptions contains the optional parameters for the FlightingRingsClient.NewListPager method.
type FlightingRingsClientListOptions struct {
	// placeholder for future optional parameters
}

// OSUpdatesClientGetOptions contains the optional parameters for the OSUpdatesClient.Get method.
type OSUpdatesClientGetOptions struct {
	// placeholder for future optional parameters
}

// OSUpdatesClientListOptions contains the optional parameters for the OSUpdatesClient.NewListPager method.
type OSUpdatesClientListOptions struct {
	// placeholder for future optional parameters
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.NewListPager method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// PackagesClientBeginCreateOptions contains the optional parameters for the PackagesClient.BeginCreate method.
type PackagesClientBeginCreateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// PackagesClientBeginDeleteOptions contains the optional parameters for the PackagesClient.BeginDelete method.
type PackagesClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// PackagesClientBeginHardDeleteOptions contains the optional parameters for the PackagesClient.BeginHardDelete method.
type PackagesClientBeginHardDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// PackagesClientBeginUpdateOptions contains the optional parameters for the PackagesClient.BeginUpdate method.
type PackagesClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// PackagesClientGetDownloadURLOptions contains the optional parameters for the PackagesClient.GetDownloadURL method.
type PackagesClientGetDownloadURLOptions struct {
	// placeholder for future optional parameters
}

// PackagesClientGetOptions contains the optional parameters for the PackagesClient.Get method.
type PackagesClientGetOptions struct {
	// placeholder for future optional parameters
}

// PackagesClientListByTestBaseAccountOptions contains the optional parameters for the PackagesClient.NewListByTestBaseAccountPager
// method.
type PackagesClientListByTestBaseAccountOptions struct {
	// placeholder for future optional parameters
}

// PackagesClientRunTestOptions contains the optional parameters for the PackagesClient.RunTest method.
type PackagesClientRunTestOptions struct {
	// The parameters supplied to the Test Base Package to start a Test Run.
	Parameters *PackageRunTestParameters
}

// SKUsClientListOptions contains the optional parameters for the SKUsClient.NewListPager method.
type SKUsClientListOptions struct {
	// placeholder for future optional parameters
}

// TestResultsClientGetConsoleLogDownloadURLOptions contains the optional parameters for the TestResultsClient.GetConsoleLogDownloadURL
// method.
type TestResultsClientGetConsoleLogDownloadURLOptions struct {
	// placeholder for future optional parameters
}

// TestResultsClientGetDownloadURLOptions contains the optional parameters for the TestResultsClient.GetDownloadURL method.
type TestResultsClientGetDownloadURLOptions struct {
	// placeholder for future optional parameters
}

// TestResultsClientGetOptions contains the optional parameters for the TestResultsClient.Get method.
type TestResultsClientGetOptions struct {
	// placeholder for future optional parameters
}

// TestResultsClientGetVideoDownloadURLOptions contains the optional parameters for the TestResultsClient.GetVideoDownloadURL
// method.
type TestResultsClientGetVideoDownloadURLOptions struct {
	// placeholder for future optional parameters
}

// TestResultsClientListOptions contains the optional parameters for the TestResultsClient.NewListPager method.
type TestResultsClientListOptions struct {
	// Odata filter
	Filter *string
}

// TestSummariesClientGetOptions contains the optional parameters for the TestSummariesClient.Get method.
type TestSummariesClientGetOptions struct {
	// placeholder for future optional parameters
}

// TestSummariesClientListOptions contains the optional parameters for the TestSummariesClient.NewListPager method.
type TestSummariesClientListOptions struct {
	// placeholder for future optional parameters
}

// TestTypesClientGetOptions contains the optional parameters for the TestTypesClient.Get method.
type TestTypesClientGetOptions struct {
	// placeholder for future optional parameters
}

// TestTypesClientListOptions contains the optional parameters for the TestTypesClient.NewListPager method.
type TestTypesClientListOptions struct {
	// placeholder for future optional parameters
}

// UsageClientListOptions contains the optional parameters for the UsageClient.NewListPager method.
type UsageClientListOptions struct {
	// Odata filter
	Filter *string
}
