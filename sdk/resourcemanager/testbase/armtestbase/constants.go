//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armtestbase

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/testbase/armtestbase"
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

// ActionType - Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
type ActionType string

const (
	ActionTypeInternal ActionType = "Internal"
)

// PossibleActionTypeValues returns the possible values for the ActionType const type.
func PossibleActionTypeValues() []ActionType {
	return []ActionType{
		ActionTypeInternal,
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

// ApplicationType - The type of a gallery application.
type ApplicationType string

const (
	ApplicationTypeWinget ApplicationType = "Winget"
)

// PossibleApplicationTypeValues returns the possible values for the ApplicationType const type.
func PossibleApplicationTypeValues() []ApplicationType {
	return []ApplicationType{
		ApplicationTypeWinget,
	}
}

// Architecture - The architecture of an OS or a first party application.
type Architecture string

const (
	ArchitectureArm64 Architecture = "arm64"
	ArchitectureX64   Architecture = "x64"
	ArchitectureX86   Architecture = "x86"
)

// PossibleArchitectureValues returns the possible values for the Architecture const type.
func PossibleArchitectureValues() []Architecture {
	return []Architecture{
		ArchitectureArm64,
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

// CredentialType - Credential type.
type CredentialType string

const (
	// CredentialTypeIntuneAccount - Username password credential for intune enrollment.
	CredentialTypeIntuneAccount CredentialType = "IntuneAccount"
)

// PossibleCredentialTypeValues returns the possible values for the CredentialType const type.
func PossibleCredentialTypeValues() []CredentialType {
	return []CredentialType{
		CredentialTypeIntuneAccount,
	}
}

// DraftPackageSourceType - The source type.
type DraftPackageSourceType string

const (
	DraftPackageSourceTypeGalleryApp       DraftPackageSourceType = "GalleryApp"
	DraftPackageSourceTypeIntuneEnrollment DraftPackageSourceType = "IntuneEnrollment"
	DraftPackageSourceTypeIntuneWin        DraftPackageSourceType = "IntuneWin"
	DraftPackageSourceTypeNative           DraftPackageSourceType = "Native"
	DraftPackageSourceTypeTestBasePackage  DraftPackageSourceType = "TestBasePackage"
)

// PossibleDraftPackageSourceTypeValues returns the possible values for the DraftPackageSourceType const type.
func PossibleDraftPackageSourceTypeValues() []DraftPackageSourceType {
	return []DraftPackageSourceType{
		DraftPackageSourceTypeGalleryApp,
		DraftPackageSourceTypeIntuneEnrollment,
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

// FileUploadResourceType - Resource type for file uploading.
type FileUploadResourceType string

const (
	// FileUploadResourceTypePackage - Upload file for package onboarding.
	FileUploadResourceTypePackage FileUploadResourceType = "Package"
	// FileUploadResourceTypeVHD - Upload VHD file for image onboarding.
	FileUploadResourceTypeVHD FileUploadResourceType = "VHD"
)

// PossibleFileUploadResourceTypeValues returns the possible values for the FileUploadResourceType const type.
func PossibleFileUploadResourceTypeValues() []FileUploadResourceType {
	return []FileUploadResourceType{
		FileUploadResourceTypePackage,
		FileUploadResourceTypeVHD,
	}
}

type FreeHourBalanceName string

const (
	FreeHourBalanceNameSubscriptionLevel FreeHourBalanceName = "SubscriptionLevel"
	FreeHourBalanceNameTenantLevel       FreeHourBalanceName = "TenantLevel"
)

// PossibleFreeHourBalanceNameValues returns the possible values for the FreeHourBalanceName const type.
func PossibleFreeHourBalanceNameValues() []FreeHourBalanceName {
	return []FreeHourBalanceName{
		FreeHourBalanceNameSubscriptionLevel,
		FreeHourBalanceNameTenantLevel,
	}
}

type FreeHourStatus string

const (
	FreeHourStatusEnabled   FreeHourStatus = "Enabled"
	FreeHourStatusSuspended FreeHourStatus = "Suspended"
)

// PossibleFreeHourStatusValues returns the possible values for the FreeHourStatus const type.
func PossibleFreeHourStatusValues() []FreeHourStatus {
	return []FreeHourStatus{
		FreeHourStatusEnabled,
		FreeHourStatusSuspended,
	}
}

type FreeHourType string

const (
	FreeHourTypePermanent FreeHourType = "Permanent"
	FreeHourTypeTemporary FreeHourType = "Temporary"
)

// PossibleFreeHourTypeValues returns the possible values for the FreeHourType const type.
func PossibleFreeHourTypeValues() []FreeHourType {
	return []FreeHourType{
		FreeHourTypePermanent,
		FreeHourTypeTemporary,
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

// ImageArchitecture - Custom image architecture.
type ImageArchitecture string

const (
	// ImageArchitectureX64 - 64-bit architecture.
	ImageArchitectureX64 ImageArchitecture = "x64"
)

// PossibleImageArchitectureValues returns the possible values for the ImageArchitecture const type.
func PossibleImageArchitectureValues() []ImageArchitecture {
	return []ImageArchitecture{
		ImageArchitectureX64,
	}
}

// ImageOSState - Custom image OS state.
type ImageOSState string

const (
	// ImageOSStateGeneralized - Sysprep generalization processed.
	ImageOSStateGeneralized ImageOSState = "Generalized"
	// ImageOSStateSpecialized - Fully kept with user specified settings.
	ImageOSStateSpecialized ImageOSState = "Specialized"
)

// PossibleImageOSStateValues returns the possible values for the ImageOSState const type.
func PossibleImageOSStateValues() []ImageOSState {
	return []ImageOSState{
		ImageOSStateGeneralized,
		ImageOSStateSpecialized,
	}
}

// ImageSecurityType - Custom image security type.
type ImageSecurityType string

const (
	// ImageSecurityTypeStandard - Standard security type.
	ImageSecurityTypeStandard ImageSecurityType = "Standard"
	// ImageSecurityTypeTrustedLaunch - Specify higher security level compared to Standard.
	ImageSecurityTypeTrustedLaunch ImageSecurityType = "TrustedLaunch"
)

// PossibleImageSecurityTypeValues returns the possible values for the ImageSecurityType const type.
func PossibleImageSecurityTypeValues() []ImageSecurityType {
	return []ImageSecurityType{
		ImageSecurityTypeStandard,
		ImageSecurityTypeTrustedLaunch,
	}
}

// ImageSource - Custom image source type.
type ImageSource string

const (
	// ImageSourceUnknown - Unknown image source type.
	ImageSourceUnknown ImageSource = "Unknown"
	// ImageSourceVHD - Specify image onboarding through VHD.
	ImageSourceVHD ImageSource = "VHD"
)

// PossibleImageSourceValues returns the possible values for the ImageSource const type.
func PossibleImageSourceValues() []ImageSource {
	return []ImageSource{
		ImageSourceUnknown,
		ImageSourceVHD,
	}
}

// ImageStatus - Status of the custom image.
type ImageStatus string

const (
	// ImageStatusFailed - Failed to onboard or pass validation.
	ImageStatusFailed ImageStatus = "Failed"
	// ImageStatusReady - Completed validation and is ready for use.
	ImageStatusReady ImageStatus = "Ready"
	// ImageStatusUnknown - Unknown image status.
	ImageStatusUnknown ImageStatus = "Unknown"
	// ImageStatusValidating - Succeed to onboard but is in validation process.
	ImageStatusValidating ImageStatus = "Validating"
)

// PossibleImageStatusValues returns the possible values for the ImageStatus const type.
func PossibleImageStatusValues() []ImageStatus {
	return []ImageStatus{
		ImageStatusFailed,
		ImageStatusReady,
		ImageStatusUnknown,
		ImageStatusValidating,
	}
}

// InteropExecutionMode - Specifies how the first party applications should be inter-operated with user's application.
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

type OrderBy string

const (
	OrderByPopularity OrderBy = "popularity"
	OrderByRelevance  OrderBy = "relevance"
)

// PossibleOrderByValues returns the possible values for the OrderBy const type.
func PossibleOrderByValues() []OrderBy {
	return []OrderBy{
		OrderByPopularity,
		OrderByRelevance,
	}
}

// Origin - The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
// value is "user,system"
type Origin string

const (
	OriginSystem     Origin = "system"
	OriginUser       Origin = "user"
	OriginUserSystem Origin = "user,system"
)

// PossibleOriginValues returns the possible values for the Origin const type.
func PossibleOriginValues() []Origin {
	return []Origin{
		OriginSystem,
		OriginUser,
		OriginUserSystem,
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

// SystemAssignedServiceIdentityType - Type of managed service identity (either system assigned, or none).
type SystemAssignedServiceIdentityType string

const (
	SystemAssignedServiceIdentityTypeNone           SystemAssignedServiceIdentityType = "None"
	SystemAssignedServiceIdentityTypeSystemAssigned SystemAssignedServiceIdentityType = "SystemAssigned"
)

// PossibleSystemAssignedServiceIdentityTypeValues returns the possible values for the SystemAssignedServiceIdentityType const type.
func PossibleSystemAssignedServiceIdentityTypeValues() []SystemAssignedServiceIdentityType {
	return []SystemAssignedServiceIdentityType{
		SystemAssignedServiceIdentityTypeNone,
		SystemAssignedServiceIdentityTypeSystemAssigned,
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
	TypeInplaceUpgrade Type = "InplaceUpgrade"
	TypeSecurityUpdate Type = "SecurityUpdate"
)

// PossibleTypeValues returns the possible values for the Type const type.
func PossibleTypeValues() []Type {
	return []Type{
		TypeFeatureUpdate,
		TypeInplaceUpgrade,
		TypeSecurityUpdate,
	}
}

// VHDStatus - The status of the VHD.
type VHDStatus string

const (
	// VHDStatusFailed - Failed to pass VHD verification.
	VHDStatusFailed VHDStatus = "Failed"
	// VHDStatusOccupied - An image is onboarding with this VHD.
	VHDStatusOccupied VHDStatus = "Occupied"
	// VHDStatusReady - Succeed to be upload and pass VHD verification.
	VHDStatusReady VHDStatus = "Ready"
	// VHDStatusUnknown - Unknown VHD status.
	VHDStatusUnknown VHDStatus = "Unknown"
	// VHDStatusVerifying - Processing VHD file checking and malware scanning.
	VHDStatusVerifying VHDStatus = "Verifying"
)

// PossibleVHDStatusValues returns the possible values for the VHDStatus const type.
func PossibleVHDStatusValues() []VHDStatus {
	return []VHDStatus{
		VHDStatusFailed,
		VHDStatusOccupied,
		VHDStatusReady,
		VHDStatusUnknown,
		VHDStatusVerifying,
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

// VerificationStatus - Indicates if the validation or rule checking is passed.
type VerificationStatus string

const (
	// VerificationStatusFailed - Validation or rule checking failed.
	VerificationStatusFailed VerificationStatus = "Failed"
	// VerificationStatusPassed - Validation or rule checking passed.
	VerificationStatusPassed VerificationStatus = "Passed"
)

// PossibleVerificationStatusValues returns the possible values for the VerificationStatus const type.
func PossibleVerificationStatusValues() []VerificationStatus {
	return []VerificationStatus{
		VerificationStatusFailed,
		VerificationStatusPassed,
	}
}
