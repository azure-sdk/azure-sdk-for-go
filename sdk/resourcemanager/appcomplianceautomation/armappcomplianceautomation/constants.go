// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappcomplianceautomation

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcomplianceautomation/armappcomplianceautomation"
	moduleVersion = "v1.0.1"
)

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

// CategoryStatus - Indicates the category status.
type CategoryStatus string

const (
	// CategoryStatusFailed - The category is failed.
	CategoryStatusFailed CategoryStatus = "Failed"
	// CategoryStatusNotApplicable - The category is not applicable.
	CategoryStatusNotApplicable CategoryStatus = "NotApplicable"
	// CategoryStatusPassed - The category is passed.
	CategoryStatusPassed CategoryStatus = "Passed"
	// CategoryStatusPendingApproval - The category is pending for approval.
	CategoryStatusPendingApproval CategoryStatus = "PendingApproval"
)

// PossibleCategoryStatusValues returns the possible values for the CategoryStatus const type.
func PossibleCategoryStatusValues() []CategoryStatus {
	return []CategoryStatus{
		CategoryStatusFailed,
		CategoryStatusNotApplicable,
		CategoryStatusPassed,
		CategoryStatusPendingApproval,
	}
}

// CheckNameAvailabilityReason - The reason why the given name is not available.
type CheckNameAvailabilityReason string

const (
	CheckNameAvailabilityReasonAlreadyExists CheckNameAvailabilityReason = "AlreadyExists"
	CheckNameAvailabilityReasonInvalid       CheckNameAvailabilityReason = "Invalid"
)

// PossibleCheckNameAvailabilityReasonValues returns the possible values for the CheckNameAvailabilityReason const type.
func PossibleCheckNameAvailabilityReasonValues() []CheckNameAvailabilityReason {
	return []CheckNameAvailabilityReason{
		CheckNameAvailabilityReasonAlreadyExists,
		CheckNameAvailabilityReasonInvalid,
	}
}

// ContentType - content type
type ContentType string

const (
	// ContentTypeApplicationJSON - The content type is application/json.
	ContentTypeApplicationJSON ContentType = "application/json"
)

// PossibleContentTypeValues returns the possible values for the ContentType const type.
func PossibleContentTypeValues() []ContentType {
	return []ContentType{
		ContentTypeApplicationJSON,
	}
}

// ControlFamilyStatus - Indicates the control family status.
type ControlFamilyStatus string

const (
	// ControlFamilyStatusFailed - The control family is failed.
	ControlFamilyStatusFailed ControlFamilyStatus = "Failed"
	// ControlFamilyStatusNotApplicable - The control family is not applicable.
	ControlFamilyStatusNotApplicable ControlFamilyStatus = "NotApplicable"
	// ControlFamilyStatusPassed - The control family is passed.
	ControlFamilyStatusPassed ControlFamilyStatus = "Passed"
	// ControlFamilyStatusPendingApproval - The control family is pending for approval.
	ControlFamilyStatusPendingApproval ControlFamilyStatus = "PendingApproval"
)

// PossibleControlFamilyStatusValues returns the possible values for the ControlFamilyStatus const type.
func PossibleControlFamilyStatusValues() []ControlFamilyStatus {
	return []ControlFamilyStatus{
		ControlFamilyStatusFailed,
		ControlFamilyStatusNotApplicable,
		ControlFamilyStatusPassed,
		ControlFamilyStatusPendingApproval,
	}
}

// ControlStatus - Indicates the control status.
type ControlStatus string

const (
	// ControlStatusFailed - The control is failed.
	ControlStatusFailed ControlStatus = "Failed"
	// ControlStatusNotApplicable - The control is not applicable.
	ControlStatusNotApplicable ControlStatus = "NotApplicable"
	// ControlStatusPassed - The control is passed.
	ControlStatusPassed ControlStatus = "Passed"
	// ControlStatusPendingApproval - The control is pending for approval.
	ControlStatusPendingApproval ControlStatus = "PendingApproval"
)

// PossibleControlStatusValues returns the possible values for the ControlStatus const type.
func PossibleControlStatusValues() []ControlStatus {
	return []ControlStatus{
		ControlStatusFailed,
		ControlStatusNotApplicable,
		ControlStatusPassed,
		ControlStatusPendingApproval,
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

// DeliveryStatus - webhook deliveryStatus
type DeliveryStatus string

const (
	// DeliveryStatusFailed - The webhook is failed to deliver.
	DeliveryStatusFailed DeliveryStatus = "Failed"
	// DeliveryStatusNotStarted - The webhook is not delivered.
	DeliveryStatusNotStarted DeliveryStatus = "NotStarted"
	// DeliveryStatusSucceeded - The webhook is delivered successfully.
	DeliveryStatusSucceeded DeliveryStatus = "Succeeded"
)

// PossibleDeliveryStatusValues returns the possible values for the DeliveryStatus const type.
func PossibleDeliveryStatusValues() []DeliveryStatus {
	return []DeliveryStatus{
		DeliveryStatusFailed,
		DeliveryStatusNotStarted,
		DeliveryStatusSucceeded,
	}
}

// DownloadType - Indicates the download type.
type DownloadType string

const (
	// DownloadTypeComplianceDetailedPDFReport - Download the detailed compliance pdf report.
	DownloadTypeComplianceDetailedPDFReport DownloadType = "ComplianceDetailedPdfReport"
	// DownloadTypeCompliancePDFReport - Download the compliance pdf report.
	DownloadTypeCompliancePDFReport DownloadType = "CompliancePdfReport"
	// DownloadTypeComplianceReport - Download the compliance report.
	DownloadTypeComplianceReport DownloadType = "ComplianceReport"
	// DownloadTypeResourceList - Download the resource list of the report.
	DownloadTypeResourceList DownloadType = "ResourceList"
)

// PossibleDownloadTypeValues returns the possible values for the DownloadType const type.
func PossibleDownloadTypeValues() []DownloadType {
	return []DownloadType{
		DownloadTypeComplianceDetailedPDFReport,
		DownloadTypeCompliancePDFReport,
		DownloadTypeComplianceReport,
		DownloadTypeResourceList,
	}
}

// EnableSSLVerification - whether to enable ssl verification
type EnableSSLVerification string

const (
	// EnableSSLVerificationFalse - The ssl verification is not enabled.
	EnableSSLVerificationFalse EnableSSLVerification = "false"
	// EnableSSLVerificationTrue - The ssl verification is enabled.
	EnableSSLVerificationTrue EnableSSLVerification = "true"
)

// PossibleEnableSSLVerificationValues returns the possible values for the EnableSSLVerification const type.
func PossibleEnableSSLVerificationValues() []EnableSSLVerification {
	return []EnableSSLVerification{
		EnableSSLVerificationFalse,
		EnableSSLVerificationTrue,
	}
}

// EvidenceType - Evidence type
type EvidenceType string

const (
	// EvidenceTypeAutoCollectedEvidence - The evidence auto collected by App Compliance Automation.
	EvidenceTypeAutoCollectedEvidence EvidenceType = "AutoCollectedEvidence"
	// EvidenceTypeData - The evidence is data.
	EvidenceTypeData EvidenceType = "Data"
	// EvidenceTypeFile - The evidence is a file.
	EvidenceTypeFile EvidenceType = "File"
)

// PossibleEvidenceTypeValues returns the possible values for the EvidenceType const type.
func PossibleEvidenceTypeValues() []EvidenceType {
	return []EvidenceType{
		EvidenceTypeAutoCollectedEvidence,
		EvidenceTypeData,
		EvidenceTypeFile,
	}
}

// InputType - Question input type.
type InputType string

const (
	// InputTypeBoolean - The input content should be a boolean.
	InputTypeBoolean InputType = "Boolean"
	// InputTypeDate - The input content should be a date.
	InputTypeDate InputType = "Date"
	// InputTypeEmail - The input content should be an email address.
	InputTypeEmail InputType = "Email"
	// InputTypeGroup - The input content is a group of answers.
	InputTypeGroup InputType = "Group"
	// InputTypeMultiSelectCheckbox - The input content are multiple results seleted from the checkboxes.
	InputTypeMultiSelectCheckbox InputType = "MultiSelectCheckbox"
	// InputTypeMultiSelectDropdown - The input content are multiple results seleted from the dropdown options.
	InputTypeMultiSelectDropdown InputType = "MultiSelectDropdown"
	// InputTypeMultiSelectDropdownCustom - The input content are result seleted from the custom dropdown options.
	InputTypeMultiSelectDropdownCustom InputType = "MultiSelectDropdownCustom"
	// InputTypeMultilineText - The input content should be multiline text.
	InputTypeMultilineText InputType = "MultilineText"
	// InputTypeNone - The input type is a text box.
	InputTypeNone InputType = "None"
	// InputTypeNumber - The input content should be a number.
	InputTypeNumber InputType = "Number"
	// InputTypeSingleSelectDropdown - The input content is a single result seleted from the dropdown options.
	InputTypeSingleSelectDropdown InputType = "SingleSelectDropdown"
	// InputTypeSingleSelection - The input content is a single result seleted from the options.
	InputTypeSingleSelection InputType = "SingleSelection"
	// InputTypeTelephone - The input content should be a telephone number.
	InputTypeTelephone InputType = "Telephone"
	// InputTypeText - The input content is text string.
	InputTypeText InputType = "Text"
	// InputTypeURL - The input content should be a URL.
	InputTypeURL InputType = "Url"
	// InputTypeUpload - The input content is a uploaded file.
	InputTypeUpload InputType = "Upload"
	// InputTypeYearPicker - The input content is a Year, pick from the dropdown list.
	InputTypeYearPicker InputType = "YearPicker"
	// InputTypeYesNoNa - The input content should be Yes, No or Na.
	InputTypeYesNoNa InputType = "YesNoNa"
)

// PossibleInputTypeValues returns the possible values for the InputType const type.
func PossibleInputTypeValues() []InputType {
	return []InputType{
		InputTypeBoolean,
		InputTypeDate,
		InputTypeEmail,
		InputTypeGroup,
		InputTypeMultiSelectCheckbox,
		InputTypeMultiSelectDropdown,
		InputTypeMultiSelectDropdownCustom,
		InputTypeMultilineText,
		InputTypeNone,
		InputTypeNumber,
		InputTypeSingleSelectDropdown,
		InputTypeSingleSelection,
		InputTypeTelephone,
		InputTypeText,
		InputTypeURL,
		InputTypeUpload,
		InputTypeYearPicker,
		InputTypeYesNoNa,
	}
}

// IsRecommendSolution - Indicates whether this solution is the recommended.
type IsRecommendSolution string

const (
	// IsRecommendSolutionFalse - This solution is not the recommended.
	IsRecommendSolutionFalse IsRecommendSolution = "false"
	// IsRecommendSolutionTrue - This solution is the recommended.
	IsRecommendSolutionTrue IsRecommendSolution = "true"
)

// PossibleIsRecommendSolutionValues returns the possible values for the IsRecommendSolution const type.
func PossibleIsRecommendSolutionValues() []IsRecommendSolution {
	return []IsRecommendSolution{
		IsRecommendSolutionFalse,
		IsRecommendSolutionTrue,
	}
}

// NotificationEvent - notification event.
type NotificationEvent string

const (
	// NotificationEventAssessmentFailure - The subscribed report failed while collecting the assessments.
	NotificationEventAssessmentFailure NotificationEvent = "assessment_failure"
	// NotificationEventGenerateSnapshotFailed - The subscribed report's snapshot is failed to generate.
	NotificationEventGenerateSnapshotFailed NotificationEvent = "generate_snapshot_failed"
	// NotificationEventGenerateSnapshotSuccess - The subscribed report's snapshot is successfully generated.
	NotificationEventGenerateSnapshotSuccess NotificationEvent = "generate_snapshot_success"
	// NotificationEventReportConfigurationChanges - The subscribed report's configuration is changed.
	NotificationEventReportConfigurationChanges NotificationEvent = "report_configuration_changes"
	// NotificationEventReportDeletion - The subscribed report is deleted.
	NotificationEventReportDeletion NotificationEvent = "report_deletion"
)

// PossibleNotificationEventValues returns the possible values for the NotificationEvent const type.
func PossibleNotificationEventValues() []NotificationEvent {
	return []NotificationEvent{
		NotificationEventAssessmentFailure,
		NotificationEventGenerateSnapshotFailed,
		NotificationEventGenerateSnapshotSuccess,
		NotificationEventReportConfigurationChanges,
		NotificationEventReportDeletion,
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

// ProvisioningState - Resource provisioning states.
type ProvisioningState string

const (
	// ProvisioningStateCanceled - The provision is canceled.
	ProvisioningStateCanceled ProvisioningState = "Canceled"
	// ProvisioningStateCreating - The creation is in progress.
	ProvisioningStateCreating ProvisioningState = "Creating"
	// ProvisioningStateDeleting - The deletion is in progress.
	ProvisioningStateDeleting ProvisioningState = "Deleting"
	// ProvisioningStateFailed - The provision is failed.
	ProvisioningStateFailed ProvisioningState = "Failed"
	// ProvisioningStateFixing - The fix of the resource in progress.
	ProvisioningStateFixing ProvisioningState = "Fixing"
	// ProvisioningStateSucceeded - The provision is succeeded.
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	// ProvisioningStateUpdating - The update of the resource in progress.
	ProvisioningStateUpdating ProvisioningState = "Updating"
	// ProvisioningStateVerifying - The verification of the resource in progress.
	ProvisioningStateVerifying ProvisioningState = "Verifying"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateCanceled,
		ProvisioningStateCreating,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateFixing,
		ProvisioningStateSucceeded,
		ProvisioningStateUpdating,
		ProvisioningStateVerifying,
	}
}

// ReportStatus - Report status.
type ReportStatus string

const (
	// ReportStatusActive - The report is active.
	ReportStatusActive ReportStatus = "Active"
	// ReportStatusDisabled - The report is disabled.
	ReportStatusDisabled ReportStatus = "Disabled"
	// ReportStatusFailed - The report is failed.
	ReportStatusFailed ReportStatus = "Failed"
	// ReportStatusReviewing - The report is under reviewing.
	ReportStatusReviewing ReportStatus = "Reviewing"
)

// PossibleReportStatusValues returns the possible values for the ReportStatus const type.
func PossibleReportStatusValues() []ReportStatus {
	return []ReportStatus{
		ReportStatusActive,
		ReportStatusDisabled,
		ReportStatusFailed,
		ReportStatusReviewing,
	}
}

// ResourceOrigin - Resource Origin.
type ResourceOrigin string

const (
	// ResourceOriginAWS - The resource is from AWS.
	ResourceOriginAWS ResourceOrigin = "AWS"
	// ResourceOriginAzure - The resource is from Azure.
	ResourceOriginAzure ResourceOrigin = "Azure"
	// ResourceOriginGCP - The resource is from GCP.
	ResourceOriginGCP ResourceOrigin = "GCP"
)

// PossibleResourceOriginValues returns the possible values for the ResourceOrigin const type.
func PossibleResourceOriginValues() []ResourceOrigin {
	return []ResourceOrigin{
		ResourceOriginAWS,
		ResourceOriginAzure,
		ResourceOriginGCP,
	}
}

// ResourceStatus - Indicates the resource status.
type ResourceStatus string

const (
	// ResourceStatusHealthy - The resource is healthy.
	ResourceStatusHealthy ResourceStatus = "Healthy"
	// ResourceStatusUnhealthy - The resource is unhealthy.
	ResourceStatusUnhealthy ResourceStatus = "Unhealthy"
)

// PossibleResourceStatusValues returns the possible values for the ResourceStatus const type.
func PossibleResourceStatusValues() []ResourceStatus {
	return []ResourceStatus{
		ResourceStatusHealthy,
		ResourceStatusUnhealthy,
	}
}

// ResponsibilityEnvironment - Indicates the customer responsibility supported cloud environment.
type ResponsibilityEnvironment string

const (
	// ResponsibilityEnvironmentAWS - The responsibility is supported in AWS.
	ResponsibilityEnvironmentAWS ResponsibilityEnvironment = "AWS"
	// ResponsibilityEnvironmentAzure - The responsibility is supported in Azure.
	ResponsibilityEnvironmentAzure ResponsibilityEnvironment = "Azure"
	// ResponsibilityEnvironmentGCP - The responsibility is supported in GCP.
	ResponsibilityEnvironmentGCP ResponsibilityEnvironment = "GCP"
	// ResponsibilityEnvironmentGeneral - The responsibility is general requirement of all environment.
	ResponsibilityEnvironmentGeneral ResponsibilityEnvironment = "General"
)

// PossibleResponsibilityEnvironmentValues returns the possible values for the ResponsibilityEnvironment const type.
func PossibleResponsibilityEnvironmentValues() []ResponsibilityEnvironment {
	return []ResponsibilityEnvironment{
		ResponsibilityEnvironmentAWS,
		ResponsibilityEnvironmentAzure,
		ResponsibilityEnvironmentGCP,
		ResponsibilityEnvironmentGeneral,
	}
}

// ResponsibilitySeverity - Indicates the customer responsibility severity.
type ResponsibilitySeverity string

const (
	// ResponsibilitySeverityHigh - The responsibility is high severity.
	ResponsibilitySeverityHigh ResponsibilitySeverity = "High"
	// ResponsibilitySeverityLow - The responsibility is low severity.
	ResponsibilitySeverityLow ResponsibilitySeverity = "Low"
	// ResponsibilitySeverityMedium - The responsibility is medium severity.
	ResponsibilitySeverityMedium ResponsibilitySeverity = "Medium"
)

// PossibleResponsibilitySeverityValues returns the possible values for the ResponsibilitySeverity const type.
func PossibleResponsibilitySeverityValues() []ResponsibilitySeverity {
	return []ResponsibilitySeverity{
		ResponsibilitySeverityHigh,
		ResponsibilitySeverityLow,
		ResponsibilitySeverityMedium,
	}
}

// ResponsibilityStatus - Indicates the customer responsibility status.
type ResponsibilityStatus string

const (
	// ResponsibilityStatusFailed - The responsibility is failed.
	ResponsibilityStatusFailed ResponsibilityStatus = "Failed"
	// ResponsibilityStatusNotApplicable - The responsibility is not applicable.
	ResponsibilityStatusNotApplicable ResponsibilityStatus = "NotApplicable"
	// ResponsibilityStatusPassed - The responsibility is passed.
	ResponsibilityStatusPassed ResponsibilityStatus = "Passed"
	// ResponsibilityStatusPendingApproval - The responsibility is pending for approval.
	ResponsibilityStatusPendingApproval ResponsibilityStatus = "PendingApproval"
)

// PossibleResponsibilityStatusValues returns the possible values for the ResponsibilityStatus const type.
func PossibleResponsibilityStatusValues() []ResponsibilityStatus {
	return []ResponsibilityStatus{
		ResponsibilityStatusFailed,
		ResponsibilityStatusNotApplicable,
		ResponsibilityStatusPassed,
		ResponsibilityStatusPendingApproval,
	}
}

// ResponsibilityType - Indicates the customer responsibility type.
type ResponsibilityType string

const (
	// ResponsibilityTypeAutomated - The responsibility is automated.
	ResponsibilityTypeAutomated ResponsibilityType = "Automated"
	// ResponsibilityTypeManual - The responsibility is manual.
	ResponsibilityTypeManual ResponsibilityType = "Manual"
	// ResponsibilityTypeScopedManual - The responsibility is scoped manual.
	ResponsibilityTypeScopedManual ResponsibilityType = "ScopedManual"
)

// PossibleResponsibilityTypeValues returns the possible values for the ResponsibilityType const type.
func PossibleResponsibilityTypeValues() []ResponsibilityType {
	return []ResponsibilityType{
		ResponsibilityTypeAutomated,
		ResponsibilityTypeManual,
		ResponsibilityTypeScopedManual,
	}
}

// Result - Indicates whether the fix action is Succeeded or Failed.
type Result string

const (
	// ResultFailed - The result is failed.
	ResultFailed Result = "Failed"
	// ResultSucceeded - The result is succeeded.
	ResultSucceeded Result = "Succeeded"
)

// PossibleResultValues returns the possible values for the Result const type.
func PossibleResultValues() []Result {
	return []Result{
		ResultFailed,
		ResultSucceeded,
	}
}

// Rule - Scoping question rule.
type Rule string

const (
	// RuleAzureApplication - The question answer should be an AzureApplication.
	RuleAzureApplication Rule = "AzureApplication"
	// RuleCharLength - The question answer length is limited.
	RuleCharLength Rule = "CharLength"
	// RuleCreditCardPCI - The question answer should be a CreditCardPCI.
	RuleCreditCardPCI Rule = "CreditCardPCI"
	// RuleDomains - The question answer should be domains.
	RuleDomains Rule = "Domains"
	// RuleDynamicDropdown - The question answer should be dynamic dropdown.
	RuleDynamicDropdown Rule = "DynamicDropdown"
	// RulePreventNonEnglishChar - The question answer should prevent non-english char.
	RulePreventNonEnglishChar Rule = "PreventNonEnglishChar"
	// RulePublicSOX - The question answer should be a PublicSOX.
	RulePublicSOX Rule = "PublicSOX"
	// RulePublisherVerification - The question answer should be publisher verification.
	RulePublisherVerification Rule = "PublisherVerification"
	// RuleRequired - The question is required to answer.
	RuleRequired Rule = "Required"
	// RuleURL - The question answer should be an Url.
	RuleURL Rule = "Url"
	// RuleUSPrivacyShield - The question answer should be a UsPrivacyShield.
	RuleUSPrivacyShield Rule = "USPrivacyShield"
	// RuleUrls - The question answer should be Urls.
	RuleUrls Rule = "Urls"
	// RuleValidEmail - The question answer should be a valid email.
	RuleValidEmail Rule = "ValidEmail"
	// RuleValidGUID - The question answer should be a valid guid.
	RuleValidGUID Rule = "ValidGuid"
)

// PossibleRuleValues returns the possible values for the Rule const type.
func PossibleRuleValues() []Rule {
	return []Rule{
		RuleAzureApplication,
		RuleCharLength,
		RuleCreditCardPCI,
		RuleDomains,
		RuleDynamicDropdown,
		RulePreventNonEnglishChar,
		RulePublicSOX,
		RulePublisherVerification,
		RuleRequired,
		RuleURL,
		RuleUSPrivacyShield,
		RuleUrls,
		RuleValidEmail,
		RuleValidGUID,
	}
}

// SendAllEvents - whether to send notification under any event.
type SendAllEvents string

const (
	// SendAllEventsFalse - No need to send notification under any event.
	SendAllEventsFalse SendAllEvents = "false"
	// SendAllEventsTrue - Need send notification under any event.
	SendAllEventsTrue SendAllEvents = "true"
)

// PossibleSendAllEventsValues returns the possible values for the SendAllEvents const type.
func PossibleSendAllEventsValues() []SendAllEvents {
	return []SendAllEvents{
		SendAllEventsFalse,
		SendAllEventsTrue,
	}
}

// UpdateWebhookKey - whether to update webhookKey.
type UpdateWebhookKey string

const (
	// UpdateWebhookKeyFalse - No need to update the webhook key.
	UpdateWebhookKeyFalse UpdateWebhookKey = "false"
	// UpdateWebhookKeyTrue - Need update the webhook key.
	UpdateWebhookKeyTrue UpdateWebhookKey = "true"
)

// PossibleUpdateWebhookKeyValues returns the possible values for the UpdateWebhookKey const type.
func PossibleUpdateWebhookKeyValues() []UpdateWebhookKey {
	return []UpdateWebhookKey{
		UpdateWebhookKeyFalse,
		UpdateWebhookKeyTrue,
	}
}

// WebhookKeyEnabled - whether webhookKey is enabled.
type WebhookKeyEnabled string

const (
	// WebhookKeyEnabledFalse - The webhookKey is not enabled.
	WebhookKeyEnabledFalse WebhookKeyEnabled = "false"
	// WebhookKeyEnabledTrue - The webhookKey is enabled.
	WebhookKeyEnabledTrue WebhookKeyEnabled = "true"
)

// PossibleWebhookKeyEnabledValues returns the possible values for the WebhookKeyEnabled const type.
func PossibleWebhookKeyEnabledValues() []WebhookKeyEnabled {
	return []WebhookKeyEnabled{
		WebhookKeyEnabledFalse,
		WebhookKeyEnabledTrue,
	}
}

// WebhookStatus - Webhook status.
type WebhookStatus string

const (
	// WebhookStatusDisabled - The webhook is disabled.
	WebhookStatusDisabled WebhookStatus = "Disabled"
	// WebhookStatusEnabled - The webhook is enabled.
	WebhookStatusEnabled WebhookStatus = "Enabled"
)

// PossibleWebhookStatusValues returns the possible values for the WebhookStatus const type.
func PossibleWebhookStatusValues() []WebhookStatus {
	return []WebhookStatus{
		WebhookStatusDisabled,
		WebhookStatusEnabled,
	}
}
