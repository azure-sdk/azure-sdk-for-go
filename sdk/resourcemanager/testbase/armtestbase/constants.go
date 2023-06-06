//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armtestbase

const (
	moduleName    = "armtestbase"
	moduleVersion = "v0.7.0"
)

// Action - The action of the command.
type Action string

const (
	ActionClose            Action = "Close"
	ActionCustom           Action = "Custom"
	ActionFlowDrivenCustom Action = "FlowDrivenCustom"
	ActionInstall          Action = "Install"
	ActionLaunch           Action = "Launch"
	ActionUninstall        Action = "Uninstall"
)

// PossibleActionValues returns the possible values for the Action const type.
func PossibleActionValues() []Action {
	return []Action{
		ActionClose,
		ActionCustom,
		ActionFlowDrivenCustom,
		ActionInstall,
		ActionLaunch,
		ActionUninstall,
	}
}

type AnalysisResultName string

const (
	AnalysisResultNameCPURegression     AnalysisResultName = "cpuRegression"
	AnalysisResultNameCPUUtilization    AnalysisResultName = "cpuUtilization"
	AnalysisResultNameMemoryRegression  AnalysisResultName = "memoryRegression"
	AnalysisResultNameMemoryUtilization AnalysisResultName = "memoryUtilization"
	AnalysisResultNameReliability       AnalysisResultName = "reliability"
	AnalysisResultNameScriptExecution   AnalysisResultName = "scriptExecution"
	AnalysisResultNameTestAnalysis      AnalysisResultName = "testAnalysis"
)

// PossibleAnalysisResultNameValues returns the possible values for the AnalysisResultName const type.
func PossibleAnalysisResultNameValues() []AnalysisResultName {
	return []AnalysisResultName{
		AnalysisResultNameCPURegression,
		AnalysisResultNameCPUUtilization,
		AnalysisResultNameMemoryRegression,
		AnalysisResultNameMemoryUtilization,
		AnalysisResultNameReliability,
		AnalysisResultNameScriptExecution,
		AnalysisResultNameTestAnalysis,
	}
}

// AnalysisResultType - Type of the Analysis Result.
type AnalysisResultType string

const (
	AnalysisResultTypeCPURegression     AnalysisResultType = "CPURegression"
	AnalysisResultTypeCPUUtilization    AnalysisResultType = "CPUUtilization"
	AnalysisResultTypeMemoryRegression  AnalysisResultType = "MemoryRegression"
	AnalysisResultTypeMemoryUtilization AnalysisResultType = "MemoryUtilization"
	AnalysisResultTypeReliability       AnalysisResultType = "Reliability"
	AnalysisResultTypeScriptExecution   AnalysisResultType = "ScriptExecution"
	AnalysisResultTypeTestAnalysis      AnalysisResultType = "TestAnalysis"
)

// PossibleAnalysisResultTypeValues returns the possible values for the AnalysisResultType const type.
func PossibleAnalysisResultTypeValues() []AnalysisResultType {
	return []AnalysisResultType{
		AnalysisResultTypeCPURegression,
		AnalysisResultTypeCPUUtilization,
		AnalysisResultTypeMemoryRegression,
		AnalysisResultTypeMemoryUtilization,
		AnalysisResultTypeReliability,
		AnalysisResultTypeScriptExecution,
		AnalysisResultTypeTestAnalysis,
	}
}

// AnalysisStatus - The analysis status.
type AnalysisStatus string

const (
	AnalysisStatusAvailable    AnalysisStatus = "Available"
	AnalysisStatusCompleted    AnalysisStatus = "Completed"
	AnalysisStatusFailed       AnalysisStatus = "Failed"
	AnalysisStatusInProgress   AnalysisStatus = "InProgress"
	AnalysisStatusNone         AnalysisStatus = "None"
	AnalysisStatusNotAvailable AnalysisStatus = "NotAvailable"
	AnalysisStatusSucceeded    AnalysisStatus = "Succeeded"
)

// PossibleAnalysisStatusValues returns the possible values for the AnalysisStatus const type.
func PossibleAnalysisStatusValues() []AnalysisStatus {
	return []AnalysisStatus{
		AnalysisStatusAvailable,
		AnalysisStatusCompleted,
		AnalysisStatusFailed,
		AnalysisStatusInProgress,
		AnalysisStatusNone,
		AnalysisStatusNotAvailable,
		AnalysisStatusSucceeded,
	}
}

// Architecture - The architecture of an OS or a first party application.
type Architecture string

const (
	ArchitectureArm   Architecture = "arm"
	ArchitectureArm64 Architecture = "arm64"
	ArchitectureIa64  Architecture = "ia64"
	ArchitectureX64   Architecture = "x64"
	ArchitectureX86   Architecture = "x86"
)

// PossibleArchitectureValues returns the possible values for the Architecture const type.
func PossibleArchitectureValues() []Architecture {
	return []Architecture{
		ArchitectureArm,
		ArchitectureArm64,
		ArchitectureIa64,
		ArchitectureX64,
		ArchitectureX86,
	}
}

// Category - The category of the failure.
type Category string

const (
	CategoryInfrastructure Category = "Infrastructure"
	CategoryNone           Category = "None"
	CategoryOSUpdate       Category = "OSUpdate"
	CategoryPackage        Category = "Package"
	CategoryUnidentified   Category = "Unidentified"
)

// PossibleCategoryValues returns the possible values for the Category const type.
func PossibleCategoryValues() []Category {
	return []Category{
		CategoryInfrastructure,
		CategoryNone,
		CategoryOSUpdate,
		CategoryPackage,
		CategoryUnidentified,
	}
}

// ContentType - The type of command content.
type ContentType string

const (
	ContentTypeFile   ContentType = "File"
	ContentTypeInline ContentType = "Inline"
	ContentTypePath   ContentType = "Path"
)

// PossibleContentTypeValues returns the possible values for the ContentType const type.
func PossibleContentTypeValues() []ContentType {
	return []ContentType{
		ContentTypeFile,
		ContentTypeInline,
		ContentTypePath,
	}
}

// CreatedByType - The type of identity that created the resource.
type CreatedByType string

const (
	CreatedByTypeApplication     CreatedByType = "Application"
	CreatedByTypeKey             CreatedByType = "Key"
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	CreatedByTypeUser            CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns the possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{
		CreatedByTypeApplication,
		CreatedByTypeKey,
		CreatedByTypeManagedIdentity,
		CreatedByTypeUser,
	}
}

// DraftPackageSourceType - The source type.
type DraftPackageSourceType string

const (
	DraftPackageSourceTypeIntuneWin       DraftPackageSourceType = "IntuneWin"
	DraftPackageSourceTypeNative          DraftPackageSourceType = "Native"
	DraftPackageSourceTypeTestBasePackage DraftPackageSourceType = "TestBasePackage"
)

// PossibleDraftPackageSourceTypeValues returns the possible values for the DraftPackageSourceType const type.
func PossibleDraftPackageSourceTypeValues() []DraftPackageSourceType {
	return []DraftPackageSourceType{
		DraftPackageSourceTypeIntuneWin,
		DraftPackageSourceTypeNative,
		DraftPackageSourceTypeTestBasePackage,
	}
}

type Engagements string

const (
	EngagementsMAPP  Engagements = "MAPP"
	EngagementsMVI   Engagements = "MVI"
	EngagementsMVP   Engagements = "MVP"
	EngagementsOther Engagements = "Other"
	EngagementsSUVP  Engagements = "SUVP"
)

// PossibleEngagementsValues returns the possible values for the Engagements const type.
func PossibleEngagementsValues() []Engagements {
	return []Engagements{
		EngagementsMAPP,
		EngagementsMVI,
		EngagementsMVP,
		EngagementsOther,
		EngagementsSUVP,
	}
}

// ExecutionStatus - The execution status of a test.
type ExecutionStatus string

const (
	ExecutionStatusCompleted   ExecutionStatus = "Completed"
	ExecutionStatusFailed      ExecutionStatus = "Failed"
	ExecutionStatusInProgress  ExecutionStatus = "InProgress"
	ExecutionStatusIncomplete  ExecutionStatus = "Incomplete"
	ExecutionStatusNone        ExecutionStatus = "None"
	ExecutionStatusNotExecuted ExecutionStatus = "NotExecuted"
	ExecutionStatusProcessing  ExecutionStatus = "Processing"
	ExecutionStatusSucceeded   ExecutionStatus = "Succeeded"
)

// PossibleExecutionStatusValues returns the possible values for the ExecutionStatus const type.
func PossibleExecutionStatusValues() []ExecutionStatus {
	return []ExecutionStatus{
		ExecutionStatusCompleted,
		ExecutionStatusFailed,
		ExecutionStatusInProgress,
		ExecutionStatusIncomplete,
		ExecutionStatusNone,
		ExecutionStatusNotExecuted,
		ExecutionStatusProcessing,
		ExecutionStatusSucceeded,
	}
}

// ExtractFileType - The type of file to extract.
type ExtractFileType string

const (
	ExtractFileTypeIntuneWinPackage ExtractFileType = "IntuneWinPackage"
	ExtractFileTypeTestBasePackage  ExtractFileType = "TestBasePackage"
)

// PossibleExtractFileTypeValues returns the possible values for the ExtractFileType const type.
func PossibleExtractFileTypeValues() []ExtractFileType {
	return []ExtractFileType{
		ExtractFileTypeIntuneWinPackage,
		ExtractFileTypeTestBasePackage,
	}
}

// Grade - The grade of a test.
type Grade string

const (
	GradeFail         Grade = "Fail"
	GradeNone         Grade = "None"
	GradeNotAvailable Grade = "NotAvailable"
	GradePass         Grade = "Pass"
)

// PossibleGradeValues returns the possible values for the Grade const type.
func PossibleGradeValues() []Grade {
	return []Grade{
		GradeFail,
		GradeNone,
		GradeNotAvailable,
		GradePass,
	}
}

// InteropExecutionMode - Specifies how the first party applications should be interoperated with user's application.
type InteropExecutionMode string

const (
	// InteropExecutionModeFirstPartyApp - User application will test with the first party applications.
	InteropExecutionModeFirstPartyApp InteropExecutionMode = "firstPartyApp"
	// InteropExecutionModeFirstPartyAppWithTests - User application will test with the first party applications. For out-of-box
	// tests, additional test cases for first party applications will also be run.
	InteropExecutionModeFirstPartyAppWithTests InteropExecutionMode = "firstPartyAppWithTests"
)

// PossibleInteropExecutionModeValues returns the possible values for the InteropExecutionMode const type.
func PossibleInteropExecutionModeValues() []InteropExecutionMode {
	return []InteropExecutionMode{
		InteropExecutionModeFirstPartyApp,
		InteropExecutionModeFirstPartyAppWithTests,
	}
}

// IntuneExtractStatus - Extract status.
type IntuneExtractStatus string

const (
	IntuneExtractStatusExtractFailed   IntuneExtractStatus = "ExtractFailed"
	IntuneExtractStatusNoDependencyApp IntuneExtractStatus = "NoDependencyApp"
	IntuneExtractStatusReady           IntuneExtractStatus = "Ready"
	IntuneExtractStatusUploadFailed    IntuneExtractStatus = "UploadFailed"
	IntuneExtractStatusUploading       IntuneExtractStatus = "Uploading"
)

// PossibleIntuneExtractStatusValues returns the possible values for the IntuneExtractStatus const type.
func PossibleIntuneExtractStatusValues() []IntuneExtractStatus {
	return []IntuneExtractStatus{
		IntuneExtractStatusExtractFailed,
		IntuneExtractStatusNoDependencyApp,
		IntuneExtractStatusReady,
		IntuneExtractStatusUploadFailed,
		IntuneExtractStatusUploading,
	}
}

// OsProductState - State of the OS product.
type OsProductState string

const (
	OsProductStateActive   OsProductState = "Active"
	OsProductStateDisabled OsProductState = "Disabled"
)

// PossibleOsProductStateValues returns the possible values for the OsProductState const type.
func PossibleOsProductStateValues() []OsProductState {
	return []OsProductState{
		OsProductStateActive,
		OsProductStateDisabled,
	}
}

// OsUpdateType - Specifies the OS update type to test against.
type OsUpdateType string

const (
	OsUpdateTypeFeatureUpdate  OsUpdateType = "FeatureUpdate"
	OsUpdateTypeInplaceUpgrade OsUpdateType = "InplaceUpgrade"
	OsUpdateTypeSecurityUpdate OsUpdateType = "SecurityUpdate"
)

// PossibleOsUpdateTypeValues returns the possible values for the OsUpdateType const type.
func PossibleOsUpdateTypeValues() []OsUpdateType {
	return []OsUpdateType{
		OsUpdateTypeFeatureUpdate,
		OsUpdateTypeInplaceUpgrade,
		OsUpdateTypeSecurityUpdate,
	}
}

// PackageStatus - The status of the package.
type PackageStatus string

const (
	PackageStatusDeleted                   PackageStatus = "Deleted"
	PackageStatusError                     PackageStatus = "Error"
	PackageStatusPreValidationCheckPass    PackageStatus = "PreValidationCheckPass"
	PackageStatusReady                     PackageStatus = "Ready"
	PackageStatusRegistered                PackageStatus = "Registered"
	PackageStatusUnknown                   PackageStatus = "Unknown"
	PackageStatusValidatingPackage         PackageStatus = "ValidatingPackage"
	PackageStatusValidationLongerThanUsual PackageStatus = "ValidationLongerThanUsual"
	PackageStatusVerifyingPackage          PackageStatus = "VerifyingPackage"
)

// PossiblePackageStatusValues returns the possible values for the PackageStatus const type.
func PossiblePackageStatusValues() []PackageStatus {
	return []PackageStatus{
		PackageStatusDeleted,
		PackageStatusError,
		PackageStatusPreValidationCheckPass,
		PackageStatusReady,
		PackageStatusRegistered,
		PackageStatusUnknown,
		PackageStatusValidatingPackage,
		PackageStatusValidationLongerThanUsual,
		PackageStatusVerifyingPackage,
	}
}

// PackageStudioTabs - Specifies the tabs when creating / cloning / editing a package.
type PackageStudioTabs string

const (
	PackageStudioTabsBasicsTab          PackageStudioTabs = "BasicsTab"
	PackageStudioTabsConfigureTestTab   PackageStudioTabs = "ConfigureTestTab"
	PackageStudioTabsEditPackageTab     PackageStudioTabs = "EditPackageTab"
	PackageStudioTabsReviewAndCreateTab PackageStudioTabs = "ReviewAndCreateTab"
	PackageStudioTabsTagsTab            PackageStudioTabs = "TagsTab"
	PackageStudioTabsTestMatrixTab      PackageStudioTabs = "TestMatrixTab"
	PackageStudioTabsUnspecified        PackageStudioTabs = "Unspecified"
)

// PossiblePackageStudioTabsValues returns the possible values for the PackageStudioTabs const type.
func PossiblePackageStudioTabsValues() []PackageStudioTabs {
	return []PackageStudioTabs{
		PackageStudioTabsBasicsTab,
		PackageStudioTabsConfigureTestTab,
		PackageStudioTabsEditPackageTab,
		PackageStudioTabsReviewAndCreateTab,
		PackageStudioTabsTagsTab,
		PackageStudioTabsTestMatrixTab,
		PackageStudioTabsUnspecified,
	}
}

// ProvisioningState - ARM provisioning state, see https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/Addendum.md#provisioningstate-property
type ProvisioningState string

const (
	ProvisioningStateCancelled ProvisioningState = "Cancelled"
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateUpdating  ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateCancelled,
		ProvisioningStateCreating,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateSucceeded,
		ProvisioningStateUpdating,
	}
}

// Reason - The reason for unavailability of a name. Required if nameAvailable == false.
type Reason string

const (
	ReasonAlreadyExists Reason = "AlreadyExists"
	ReasonInvalid       Reason = "Invalid"
)

// PossibleReasonValues returns the possible values for the Reason const type.
func PossibleReasonValues() []Reason {
	return []Reason{
		ReasonAlreadyExists,
		ReasonInvalid,
	}
}

type RequestStatus string

const (
	RequestStatusApproved RequestStatus = "Approved"
	RequestStatusDeclined RequestStatus = "Declined"
	RequestStatusInReview RequestStatus = "InReview"
)

// PossibleRequestStatusValues returns the possible values for the RequestStatus const type.
func PossibleRequestStatusValues() []RequestStatus {
	return []RequestStatus{
		RequestStatusApproved,
		RequestStatusDeclined,
		RequestStatusInReview,
	}
}

type RequestTypes string

const (
	RequestTypesPreReleaseAccess RequestTypes = "PreReleaseAccess"
)

// PossibleRequestTypesValues returns the possible values for the RequestTypes const type.
func PossibleRequestTypesValues() []RequestTypes {
	return []RequestTypes{
		RequestTypesPreReleaseAccess,
	}
}

// TestAnalysisStatus - The status of the analysis.
type TestAnalysisStatus string

const (
	TestAnalysisStatusAnalyzing TestAnalysisStatus = "Analyzing"
	TestAnalysisStatusCompleted TestAnalysisStatus = "Completed"
	TestAnalysisStatusFailed    TestAnalysisStatus = "Failed"
	TestAnalysisStatusNone      TestAnalysisStatus = "None"
)

// PossibleTestAnalysisStatusValues returns the possible values for the TestAnalysisStatus const type.
func PossibleTestAnalysisStatusValues() []TestAnalysisStatus {
	return []TestAnalysisStatus{
		TestAnalysisStatusAnalyzing,
		TestAnalysisStatusCompleted,
		TestAnalysisStatusFailed,
		TestAnalysisStatusNone,
	}
}

// TestStatus - The status of a test.
type TestStatus string

const (
	TestStatusCompleted               TestStatus = "Completed"
	TestStatusDataProcessing          TestStatus = "DataProcessing"
	TestStatusInfrastructureFailure   TestStatus = "InfrastructureFailure"
	TestStatusNone                    TestStatus = "None"
	TestStatusTestAndUpdateFailure    TestStatus = "TestAndUpdateFailure"
	TestStatusTestExecutionInProgress TestStatus = "TestExecutionInProgress"
	TestStatusTestFailure             TestStatus = "TestFailure"
	TestStatusUpdateFailure           TestStatus = "UpdateFailure"
)

// PossibleTestStatusValues returns the possible values for the TestStatus const type.
func PossibleTestStatusValues() []TestStatus {
	return []TestStatus{
		TestStatusCompleted,
		TestStatusDataProcessing,
		TestStatusInfrastructureFailure,
		TestStatusNone,
		TestStatusTestAndUpdateFailure,
		TestStatusTestExecutionInProgress,
		TestStatusTestFailure,
		TestStatusUpdateFailure,
	}
}

// TestType - The test type.
type TestType string

const (
	TestTypeFlowDrivenTest TestType = "FlowDrivenTest"
	TestTypeFunctionalTest TestType = "FunctionalTest"
	TestTypeOutOfBoxTest   TestType = "OutOfBoxTest"
)

// PossibleTestTypeValues returns the possible values for the TestType const type.
func PossibleTestTypeValues() []TestType {
	return []TestType{
		TestTypeFlowDrivenTest,
		TestTypeFunctionalTest,
		TestTypeOutOfBoxTest,
	}
}

// Tier - The tier of this particular SKU.
type Tier string

const (
	TierStandard Tier = "Standard"
)

// PossibleTierValues returns the possible values for the Tier const type.
func PossibleTierValues() []Tier {
	return []Tier{
		TierStandard,
	}
}

// Type - The type of this release (OS update).
type Type string

const (
	TypeFeatureUpdate  Type = "FeatureUpdate"
	TypeSecurityUpdate Type = "SecurityUpdate"
)

// PossibleTypeValues returns the possible values for the Type const type.
func PossibleTypeValues() []Type {
	return []Type{
		TypeFeatureUpdate,
		TypeSecurityUpdate,
	}
}

// ValidationRunStatus - The status of the validation run of the package.
type ValidationRunStatus string

const (
	ValidationRunStatusFailed  ValidationRunStatus = "Failed"
	ValidationRunStatusPassed  ValidationRunStatus = "Passed"
	ValidationRunStatusPending ValidationRunStatus = "Pending"
	ValidationRunStatusUnknown ValidationRunStatus = "Unknown"
)

// PossibleValidationRunStatusValues returns the possible values for the ValidationRunStatus const type.
func PossibleValidationRunStatusValues() []ValidationRunStatus {
	return []ValidationRunStatus{
		ValidationRunStatusFailed,
		ValidationRunStatusPassed,
		ValidationRunStatusPending,
		ValidationRunStatusUnknown,
	}
}
