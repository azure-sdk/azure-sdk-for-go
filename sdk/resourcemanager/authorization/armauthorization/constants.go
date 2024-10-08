//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armauthorization

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization"
	moduleVersion = "v3.0.0-beta.3"
)

// AccessRecommendationType - The feature- generated recommendation shown to the reviewer.
type AccessRecommendationType string

const (
	AccessRecommendationTypeApprove         AccessRecommendationType = "Approve"
	AccessRecommendationTypeDeny            AccessRecommendationType = "Deny"
	AccessRecommendationTypeNoInfoAvailable AccessRecommendationType = "NoInfoAvailable"
)

// PossibleAccessRecommendationTypeValues returns the possible values for the AccessRecommendationType const type.
func PossibleAccessRecommendationTypeValues() []AccessRecommendationType {
	return []AccessRecommendationType{
		AccessRecommendationTypeApprove,
		AccessRecommendationTypeDeny,
		AccessRecommendationTypeNoInfoAvailable,
	}
}

// AccessReviewActorIdentityType - The identity type : user/servicePrincipal
type AccessReviewActorIdentityType string

const (
	AccessReviewActorIdentityTypeServicePrincipal AccessReviewActorIdentityType = "servicePrincipal"
	AccessReviewActorIdentityTypeUser             AccessReviewActorIdentityType = "user"
)

// PossibleAccessReviewActorIdentityTypeValues returns the possible values for the AccessReviewActorIdentityType const type.
func PossibleAccessReviewActorIdentityTypeValues() []AccessReviewActorIdentityType {
	return []AccessReviewActorIdentityType{
		AccessReviewActorIdentityTypeServicePrincipal,
		AccessReviewActorIdentityTypeUser,
	}
}

// AccessReviewApplyResult - The outcome of applying the decision.
type AccessReviewApplyResult string

const (
	AccessReviewApplyResultAppliedSuccessfully                  AccessReviewApplyResult = "AppliedSuccessfully"
	AccessReviewApplyResultAppliedSuccessfullyButObjectNotFound AccessReviewApplyResult = "AppliedSuccessfullyButObjectNotFound"
	AccessReviewApplyResultAppliedWithUnknownFailure            AccessReviewApplyResult = "AppliedWithUnknownFailure"
	AccessReviewApplyResultApplyNotSupported                    AccessReviewApplyResult = "ApplyNotSupported"
	AccessReviewApplyResultApplying                             AccessReviewApplyResult = "Applying"
	AccessReviewApplyResultNew                                  AccessReviewApplyResult = "New"
)

// PossibleAccessReviewApplyResultValues returns the possible values for the AccessReviewApplyResult const type.
func PossibleAccessReviewApplyResultValues() []AccessReviewApplyResult {
	return []AccessReviewApplyResult{
		AccessReviewApplyResultAppliedSuccessfully,
		AccessReviewApplyResultAppliedSuccessfullyButObjectNotFound,
		AccessReviewApplyResultAppliedWithUnknownFailure,
		AccessReviewApplyResultApplyNotSupported,
		AccessReviewApplyResultApplying,
		AccessReviewApplyResultNew,
	}
}

// AccessReviewDecisionInsightType - The type of insight
type AccessReviewDecisionInsightType string

const (
	AccessReviewDecisionInsightTypeUserSignInInsight AccessReviewDecisionInsightType = "userSignInInsight"
)

// PossibleAccessReviewDecisionInsightTypeValues returns the possible values for the AccessReviewDecisionInsightType const type.
func PossibleAccessReviewDecisionInsightTypeValues() []AccessReviewDecisionInsightType {
	return []AccessReviewDecisionInsightType{
		AccessReviewDecisionInsightTypeUserSignInInsight,
	}
}

type AccessReviewDecisionPrincipalResourceMembershipType string

const (
	AccessReviewDecisionPrincipalResourceMembershipTypeDirect   AccessReviewDecisionPrincipalResourceMembershipType = "direct"
	AccessReviewDecisionPrincipalResourceMembershipTypeIndirect AccessReviewDecisionPrincipalResourceMembershipType = "indirect"
)

// PossibleAccessReviewDecisionPrincipalResourceMembershipTypeValues returns the possible values for the AccessReviewDecisionPrincipalResourceMembershipType const type.
func PossibleAccessReviewDecisionPrincipalResourceMembershipTypeValues() []AccessReviewDecisionPrincipalResourceMembershipType {
	return []AccessReviewDecisionPrincipalResourceMembershipType{
		AccessReviewDecisionPrincipalResourceMembershipTypeDirect,
		AccessReviewDecisionPrincipalResourceMembershipTypeIndirect,
	}
}

// AccessReviewHistoryDefinitionStatus - This read-only field specifies the of the requested review history data. This is
// either requested, in-progress, done or error.
type AccessReviewHistoryDefinitionStatus string

const (
	AccessReviewHistoryDefinitionStatusDone       AccessReviewHistoryDefinitionStatus = "Done"
	AccessReviewHistoryDefinitionStatusError      AccessReviewHistoryDefinitionStatus = "Error"
	AccessReviewHistoryDefinitionStatusInProgress AccessReviewHistoryDefinitionStatus = "InProgress"
	AccessReviewHistoryDefinitionStatusRequested  AccessReviewHistoryDefinitionStatus = "Requested"
)

// PossibleAccessReviewHistoryDefinitionStatusValues returns the possible values for the AccessReviewHistoryDefinitionStatus const type.
func PossibleAccessReviewHistoryDefinitionStatusValues() []AccessReviewHistoryDefinitionStatus {
	return []AccessReviewHistoryDefinitionStatus{
		AccessReviewHistoryDefinitionStatusDone,
		AccessReviewHistoryDefinitionStatusError,
		AccessReviewHistoryDefinitionStatusInProgress,
		AccessReviewHistoryDefinitionStatusRequested,
	}
}

// AccessReviewInstanceReviewersType - This field specifies the type of reviewers for a review. Usually for a review, reviewers
// are explicitly assigned. However, in some cases, the reviewers may not be assigned and instead be chosen
// dynamically. For example managers review or self review.
type AccessReviewInstanceReviewersType string

const (
	AccessReviewInstanceReviewersTypeAssigned AccessReviewInstanceReviewersType = "Assigned"
	AccessReviewInstanceReviewersTypeManagers AccessReviewInstanceReviewersType = "Managers"
	AccessReviewInstanceReviewersTypeSelf     AccessReviewInstanceReviewersType = "Self"
)

// PossibleAccessReviewInstanceReviewersTypeValues returns the possible values for the AccessReviewInstanceReviewersType const type.
func PossibleAccessReviewInstanceReviewersTypeValues() []AccessReviewInstanceReviewersType {
	return []AccessReviewInstanceReviewersType{
		AccessReviewInstanceReviewersTypeAssigned,
		AccessReviewInstanceReviewersTypeManagers,
		AccessReviewInstanceReviewersTypeSelf,
	}
}

// AccessReviewInstanceStatus - This read-only field specifies the status of an access review instance.
type AccessReviewInstanceStatus string

const (
	AccessReviewInstanceStatusApplied       AccessReviewInstanceStatus = "Applied"
	AccessReviewInstanceStatusApplying      AccessReviewInstanceStatus = "Applying"
	AccessReviewInstanceStatusAutoReviewed  AccessReviewInstanceStatus = "AutoReviewed"
	AccessReviewInstanceStatusAutoReviewing AccessReviewInstanceStatus = "AutoReviewing"
	AccessReviewInstanceStatusCompleted     AccessReviewInstanceStatus = "Completed"
	AccessReviewInstanceStatusCompleting    AccessReviewInstanceStatus = "Completing"
	AccessReviewInstanceStatusInProgress    AccessReviewInstanceStatus = "InProgress"
	AccessReviewInstanceStatusInitializing  AccessReviewInstanceStatus = "Initializing"
	AccessReviewInstanceStatusNotStarted    AccessReviewInstanceStatus = "NotStarted"
	AccessReviewInstanceStatusScheduled     AccessReviewInstanceStatus = "Scheduled"
	AccessReviewInstanceStatusStarting      AccessReviewInstanceStatus = "Starting"
)

// PossibleAccessReviewInstanceStatusValues returns the possible values for the AccessReviewInstanceStatus const type.
func PossibleAccessReviewInstanceStatusValues() []AccessReviewInstanceStatus {
	return []AccessReviewInstanceStatus{
		AccessReviewInstanceStatusApplied,
		AccessReviewInstanceStatusApplying,
		AccessReviewInstanceStatusAutoReviewed,
		AccessReviewInstanceStatusAutoReviewing,
		AccessReviewInstanceStatusCompleted,
		AccessReviewInstanceStatusCompleting,
		AccessReviewInstanceStatusInProgress,
		AccessReviewInstanceStatusInitializing,
		AccessReviewInstanceStatusNotStarted,
		AccessReviewInstanceStatusScheduled,
		AccessReviewInstanceStatusStarting,
	}
}

// AccessReviewRecurrencePatternType - The recurrence type : weekly, monthly, etc.
type AccessReviewRecurrencePatternType string

const (
	AccessReviewRecurrencePatternTypeAbsoluteMonthly AccessReviewRecurrencePatternType = "absoluteMonthly"
	AccessReviewRecurrencePatternTypeWeekly          AccessReviewRecurrencePatternType = "weekly"
)

// PossibleAccessReviewRecurrencePatternTypeValues returns the possible values for the AccessReviewRecurrencePatternType const type.
func PossibleAccessReviewRecurrencePatternTypeValues() []AccessReviewRecurrencePatternType {
	return []AccessReviewRecurrencePatternType{
		AccessReviewRecurrencePatternTypeAbsoluteMonthly,
		AccessReviewRecurrencePatternTypeWeekly,
	}
}

// AccessReviewRecurrenceRangeType - The recurrence range type. The possible values are: endDate, noEnd, numbered.
type AccessReviewRecurrenceRangeType string

const (
	AccessReviewRecurrenceRangeTypeEndDate  AccessReviewRecurrenceRangeType = "endDate"
	AccessReviewRecurrenceRangeTypeNoEnd    AccessReviewRecurrenceRangeType = "noEnd"
	AccessReviewRecurrenceRangeTypeNumbered AccessReviewRecurrenceRangeType = "numbered"
)

// PossibleAccessReviewRecurrenceRangeTypeValues returns the possible values for the AccessReviewRecurrenceRangeType const type.
func PossibleAccessReviewRecurrenceRangeTypeValues() []AccessReviewRecurrenceRangeType {
	return []AccessReviewRecurrenceRangeType{
		AccessReviewRecurrenceRangeTypeEndDate,
		AccessReviewRecurrenceRangeTypeNoEnd,
		AccessReviewRecurrenceRangeTypeNumbered,
	}
}

// AccessReviewResult - Represents a reviewer's decision for a given review
type AccessReviewResult string

const (
	AccessReviewResultApprove     AccessReviewResult = "Approve"
	AccessReviewResultDeny        AccessReviewResult = "Deny"
	AccessReviewResultDontKnow    AccessReviewResult = "DontKnow"
	AccessReviewResultNotNotified AccessReviewResult = "NotNotified"
	AccessReviewResultNotReviewed AccessReviewResult = "NotReviewed"
)

// PossibleAccessReviewResultValues returns the possible values for the AccessReviewResult const type.
func PossibleAccessReviewResultValues() []AccessReviewResult {
	return []AccessReviewResult{
		AccessReviewResultApprove,
		AccessReviewResultDeny,
		AccessReviewResultDontKnow,
		AccessReviewResultNotNotified,
		AccessReviewResultNotReviewed,
	}
}

// AccessReviewReviewerType - The identity type : user/servicePrincipal
type AccessReviewReviewerType string

const (
	AccessReviewReviewerTypeServicePrincipal AccessReviewReviewerType = "servicePrincipal"
	AccessReviewReviewerTypeUser             AccessReviewReviewerType = "user"
)

// PossibleAccessReviewReviewerTypeValues returns the possible values for the AccessReviewReviewerType const type.
func PossibleAccessReviewReviewerTypeValues() []AccessReviewReviewerType {
	return []AccessReviewReviewerType{
		AccessReviewReviewerTypeServicePrincipal,
		AccessReviewReviewerTypeUser,
	}
}

// AccessReviewScheduleDefinitionReviewersType - This field specifies the type of reviewers for a review. Usually for a review,
// reviewers are explicitly assigned. However, in some cases, the reviewers may not be assigned and instead be chosen
// dynamically. For example managers review or self review.
type AccessReviewScheduleDefinitionReviewersType string

const (
	AccessReviewScheduleDefinitionReviewersTypeAssigned AccessReviewScheduleDefinitionReviewersType = "Assigned"
	AccessReviewScheduleDefinitionReviewersTypeManagers AccessReviewScheduleDefinitionReviewersType = "Managers"
	AccessReviewScheduleDefinitionReviewersTypeSelf     AccessReviewScheduleDefinitionReviewersType = "Self"
)

// PossibleAccessReviewScheduleDefinitionReviewersTypeValues returns the possible values for the AccessReviewScheduleDefinitionReviewersType const type.
func PossibleAccessReviewScheduleDefinitionReviewersTypeValues() []AccessReviewScheduleDefinitionReviewersType {
	return []AccessReviewScheduleDefinitionReviewersType{
		AccessReviewScheduleDefinitionReviewersTypeAssigned,
		AccessReviewScheduleDefinitionReviewersTypeManagers,
		AccessReviewScheduleDefinitionReviewersTypeSelf,
	}
}

// AccessReviewScheduleDefinitionStatus - This read-only field specifies the status of an accessReview.
type AccessReviewScheduleDefinitionStatus string

const (
	AccessReviewScheduleDefinitionStatusApplied       AccessReviewScheduleDefinitionStatus = "Applied"
	AccessReviewScheduleDefinitionStatusApplying      AccessReviewScheduleDefinitionStatus = "Applying"
	AccessReviewScheduleDefinitionStatusAutoReviewed  AccessReviewScheduleDefinitionStatus = "AutoReviewed"
	AccessReviewScheduleDefinitionStatusAutoReviewing AccessReviewScheduleDefinitionStatus = "AutoReviewing"
	AccessReviewScheduleDefinitionStatusCompleted     AccessReviewScheduleDefinitionStatus = "Completed"
	AccessReviewScheduleDefinitionStatusCompleting    AccessReviewScheduleDefinitionStatus = "Completing"
	AccessReviewScheduleDefinitionStatusInProgress    AccessReviewScheduleDefinitionStatus = "InProgress"
	AccessReviewScheduleDefinitionStatusInitializing  AccessReviewScheduleDefinitionStatus = "Initializing"
	AccessReviewScheduleDefinitionStatusNotStarted    AccessReviewScheduleDefinitionStatus = "NotStarted"
	AccessReviewScheduleDefinitionStatusScheduled     AccessReviewScheduleDefinitionStatus = "Scheduled"
	AccessReviewScheduleDefinitionStatusStarting      AccessReviewScheduleDefinitionStatus = "Starting"
)

// PossibleAccessReviewScheduleDefinitionStatusValues returns the possible values for the AccessReviewScheduleDefinitionStatus const type.
func PossibleAccessReviewScheduleDefinitionStatusValues() []AccessReviewScheduleDefinitionStatus {
	return []AccessReviewScheduleDefinitionStatus{
		AccessReviewScheduleDefinitionStatusApplied,
		AccessReviewScheduleDefinitionStatusApplying,
		AccessReviewScheduleDefinitionStatusAutoReviewed,
		AccessReviewScheduleDefinitionStatusAutoReviewing,
		AccessReviewScheduleDefinitionStatusCompleted,
		AccessReviewScheduleDefinitionStatusCompleting,
		AccessReviewScheduleDefinitionStatusInProgress,
		AccessReviewScheduleDefinitionStatusInitializing,
		AccessReviewScheduleDefinitionStatusNotStarted,
		AccessReviewScheduleDefinitionStatusScheduled,
		AccessReviewScheduleDefinitionStatusStarting,
	}
}

// AccessReviewScopeAssignmentState - The role assignment state eligible/active to review
type AccessReviewScopeAssignmentState string

const (
	AccessReviewScopeAssignmentStateActive   AccessReviewScopeAssignmentState = "active"
	AccessReviewScopeAssignmentStateEligible AccessReviewScopeAssignmentState = "eligible"
)

// PossibleAccessReviewScopeAssignmentStateValues returns the possible values for the AccessReviewScopeAssignmentState const type.
func PossibleAccessReviewScopeAssignmentStateValues() []AccessReviewScopeAssignmentState {
	return []AccessReviewScopeAssignmentState{
		AccessReviewScopeAssignmentStateActive,
		AccessReviewScopeAssignmentStateEligible,
	}
}

// AccessReviewScopePrincipalType - The identity type user/servicePrincipal to review
type AccessReviewScopePrincipalType string

const (
	AccessReviewScopePrincipalTypeGuestUser         AccessReviewScopePrincipalType = "guestUser"
	AccessReviewScopePrincipalTypeRedeemedGuestUser AccessReviewScopePrincipalType = "redeemedGuestUser"
	AccessReviewScopePrincipalTypeServicePrincipal  AccessReviewScopePrincipalType = "servicePrincipal"
	AccessReviewScopePrincipalTypeUser              AccessReviewScopePrincipalType = "user"
	AccessReviewScopePrincipalTypeUserGroup         AccessReviewScopePrincipalType = "user,group"
)

// PossibleAccessReviewScopePrincipalTypeValues returns the possible values for the AccessReviewScopePrincipalType const type.
func PossibleAccessReviewScopePrincipalTypeValues() []AccessReviewScopePrincipalType {
	return []AccessReviewScopePrincipalType{
		AccessReviewScopePrincipalTypeGuestUser,
		AccessReviewScopePrincipalTypeRedeemedGuestUser,
		AccessReviewScopePrincipalTypeServicePrincipal,
		AccessReviewScopePrincipalTypeUser,
		AccessReviewScopePrincipalTypeUserGroup,
	}
}

// ApprovalMode - The type of rule
type ApprovalMode string

const (
	ApprovalModeNoApproval  ApprovalMode = "NoApproval"
	ApprovalModeParallel    ApprovalMode = "Parallel"
	ApprovalModeSerial      ApprovalMode = "Serial"
	ApprovalModeSingleStage ApprovalMode = "SingleStage"
)

// PossibleApprovalModeValues returns the possible values for the ApprovalMode const type.
func PossibleApprovalModeValues() []ApprovalMode {
	return []ApprovalMode{
		ApprovalModeNoApproval,
		ApprovalModeParallel,
		ApprovalModeSerial,
		ApprovalModeSingleStage,
	}
}

// AssignmentType - Assignment type of the role assignment schedule
type AssignmentType string

const (
	AssignmentTypeActivated AssignmentType = "Activated"
	AssignmentTypeAssigned  AssignmentType = "Assigned"
)

// PossibleAssignmentTypeValues returns the possible values for the AssignmentType const type.
func PossibleAssignmentTypeValues() []AssignmentType {
	return []AssignmentType{
		AssignmentTypeActivated,
		AssignmentTypeAssigned,
	}
}

// DecisionResourceType - The type of resource
type DecisionResourceType string

const (
	DecisionResourceTypeAzureRole DecisionResourceType = "azureRole"
)

// PossibleDecisionResourceTypeValues returns the possible values for the DecisionResourceType const type.
func PossibleDecisionResourceTypeValues() []DecisionResourceType {
	return []DecisionResourceType{
		DecisionResourceTypeAzureRole,
	}
}

// DecisionTargetType - The type of decision target : User/ServicePrincipal
type DecisionTargetType string

const (
	DecisionTargetTypeServicePrincipal DecisionTargetType = "servicePrincipal"
	DecisionTargetTypeUser             DecisionTargetType = "user"
)

// PossibleDecisionTargetTypeValues returns the possible values for the DecisionTargetType const type.
func PossibleDecisionTargetTypeValues() []DecisionTargetType {
	return []DecisionTargetType{
		DecisionTargetTypeServicePrincipal,
		DecisionTargetTypeUser,
	}
}

// DefaultDecisionType - This specifies the behavior for the autoReview feature when an access review completes.
type DefaultDecisionType string

const (
	DefaultDecisionTypeApprove        DefaultDecisionType = "Approve"
	DefaultDecisionTypeDeny           DefaultDecisionType = "Deny"
	DefaultDecisionTypeRecommendation DefaultDecisionType = "Recommendation"
)

// PossibleDefaultDecisionTypeValues returns the possible values for the DefaultDecisionType const type.
func PossibleDefaultDecisionTypeValues() []DefaultDecisionType {
	return []DefaultDecisionType{
		DefaultDecisionTypeApprove,
		DefaultDecisionTypeDeny,
		DefaultDecisionTypeRecommendation,
	}
}

// EnablementRules - The type of enablement rule
type EnablementRules string

const (
	EnablementRulesJustification             EnablementRules = "Justification"
	EnablementRulesMultiFactorAuthentication EnablementRules = "MultiFactorAuthentication"
	EnablementRulesTicketing                 EnablementRules = "Ticketing"
)

// PossibleEnablementRulesValues returns the possible values for the EnablementRules const type.
func PossibleEnablementRulesValues() []EnablementRules {
	return []EnablementRules{
		EnablementRulesJustification,
		EnablementRulesMultiFactorAuthentication,
		EnablementRulesTicketing,
	}
}

type ExcludedPrincipalTypes string

const (
	ExcludedPrincipalTypesServicePrincipalsAsRequestor ExcludedPrincipalTypes = "ServicePrincipalsAsRequestor"
	ExcludedPrincipalTypesServicePrincipalsAsTarget    ExcludedPrincipalTypes = "ServicePrincipalsAsTarget"
)

// PossibleExcludedPrincipalTypesValues returns the possible values for the ExcludedPrincipalTypes const type.
func PossibleExcludedPrincipalTypesValues() []ExcludedPrincipalTypes {
	return []ExcludedPrincipalTypes{
		ExcludedPrincipalTypesServicePrincipalsAsRequestor,
		ExcludedPrincipalTypesServicePrincipalsAsTarget,
	}
}

// MemberType - Membership type of the role assignment schedule
type MemberType string

const (
	MemberTypeDirect    MemberType = "Direct"
	MemberTypeGroup     MemberType = "Group"
	MemberTypeInherited MemberType = "Inherited"
)

// PossibleMemberTypeValues returns the possible values for the MemberType const type.
func PossibleMemberTypeValues() []MemberType {
	return []MemberType{
		MemberTypeDirect,
		MemberTypeGroup,
		MemberTypeInherited,
	}
}

// NotificationDeliveryMechanism - The type of notification.
type NotificationDeliveryMechanism string

const (
	NotificationDeliveryMechanismEmail NotificationDeliveryMechanism = "Email"
)

// PossibleNotificationDeliveryMechanismValues returns the possible values for the NotificationDeliveryMechanism const type.
func PossibleNotificationDeliveryMechanismValues() []NotificationDeliveryMechanism {
	return []NotificationDeliveryMechanism{
		NotificationDeliveryMechanismEmail,
	}
}

// NotificationLevel - The notification level.
type NotificationLevel string

const (
	NotificationLevelAll      NotificationLevel = "All"
	NotificationLevelCritical NotificationLevel = "Critical"
	NotificationLevelNone     NotificationLevel = "None"
)

// PossibleNotificationLevelValues returns the possible values for the NotificationLevel const type.
func PossibleNotificationLevelValues() []NotificationLevel {
	return []NotificationLevel{
		NotificationLevelAll,
		NotificationLevelCritical,
		NotificationLevelNone,
	}
}

// PIMOnlyMode - Determines whether the setting is enabled, disabled or report only.
type PIMOnlyMode string

const (
	PIMOnlyModeDisabled   PIMOnlyMode = "Disabled"
	PIMOnlyModeEnabled    PIMOnlyMode = "Enabled"
	PIMOnlyModeReportOnly PIMOnlyMode = "ReportOnly"
)

// PossiblePIMOnlyModeValues returns the possible values for the PIMOnlyMode const type.
func PossiblePIMOnlyModeValues() []PIMOnlyMode {
	return []PIMOnlyMode{
		PIMOnlyModeDisabled,
		PIMOnlyModeEnabled,
		PIMOnlyModeReportOnly,
	}
}

// PrincipalType - The principal type of the assigned principal ID.
type PrincipalType string

const (
	PrincipalTypeDevice           PrincipalType = "Device"
	PrincipalTypeForeignGroup     PrincipalType = "ForeignGroup"
	PrincipalTypeGroup            PrincipalType = "Group"
	PrincipalTypeServicePrincipal PrincipalType = "ServicePrincipal"
	PrincipalTypeUser             PrincipalType = "User"
)

// PossiblePrincipalTypeValues returns the possible values for the PrincipalType const type.
func PossiblePrincipalTypeValues() []PrincipalType {
	return []PrincipalType{
		PrincipalTypeDevice,
		PrincipalTypeForeignGroup,
		PrincipalTypeGroup,
		PrincipalTypeServicePrincipal,
		PrincipalTypeUser,
	}
}

// RecipientType - The recipient type.
type RecipientType string

const (
	RecipientTypeAdmin     RecipientType = "Admin"
	RecipientTypeApprover  RecipientType = "Approver"
	RecipientTypeRequestor RecipientType = "Requestor"
)

// PossibleRecipientTypeValues returns the possible values for the RecipientType const type.
func PossibleRecipientTypeValues() []RecipientType {
	return []RecipientType{
		RecipientTypeAdmin,
		RecipientTypeApprover,
		RecipientTypeRequestor,
	}
}

// RecordAllDecisionsResult - The decision to make. Approvers can take action of Approve/Deny
type RecordAllDecisionsResult string

const (
	RecordAllDecisionsResultApprove RecordAllDecisionsResult = "Approve"
	RecordAllDecisionsResultDeny    RecordAllDecisionsResult = "Deny"
)

// PossibleRecordAllDecisionsResultValues returns the possible values for the RecordAllDecisionsResult const type.
func PossibleRecordAllDecisionsResultValues() []RecordAllDecisionsResult {
	return []RecordAllDecisionsResult{
		RecordAllDecisionsResultApprove,
		RecordAllDecisionsResultDeny,
	}
}

// RequestType - The type of the role assignment schedule request. Eg: SelfActivate, AdminAssign etc
type RequestType string

const (
	RequestTypeAdminAssign    RequestType = "AdminAssign"
	RequestTypeAdminExtend    RequestType = "AdminExtend"
	RequestTypeAdminRemove    RequestType = "AdminRemove"
	RequestTypeAdminRenew     RequestType = "AdminRenew"
	RequestTypeAdminUpdate    RequestType = "AdminUpdate"
	RequestTypeSelfActivate   RequestType = "SelfActivate"
	RequestTypeSelfDeactivate RequestType = "SelfDeactivate"
	RequestTypeSelfExtend     RequestType = "SelfExtend"
	RequestTypeSelfRenew      RequestType = "SelfRenew"
)

// PossibleRequestTypeValues returns the possible values for the RequestType const type.
func PossibleRequestTypeValues() []RequestType {
	return []RequestType{
		RequestTypeAdminAssign,
		RequestTypeAdminExtend,
		RequestTypeAdminRemove,
		RequestTypeAdminRenew,
		RequestTypeAdminUpdate,
		RequestTypeSelfActivate,
		RequestTypeSelfDeactivate,
		RequestTypeSelfExtend,
		RequestTypeSelfRenew,
	}
}

// RoleManagementPolicyRuleType - The type of rule
type RoleManagementPolicyRuleType string

const (
	RoleManagementPolicyRuleTypeRoleManagementPolicyApprovalRule              RoleManagementPolicyRuleType = "RoleManagementPolicyApprovalRule"
	RoleManagementPolicyRuleTypeRoleManagementPolicyAuthenticationContextRule RoleManagementPolicyRuleType = "RoleManagementPolicyAuthenticationContextRule"
	RoleManagementPolicyRuleTypeRoleManagementPolicyEnablementRule            RoleManagementPolicyRuleType = "RoleManagementPolicyEnablementRule"
	RoleManagementPolicyRuleTypeRoleManagementPolicyExpirationRule            RoleManagementPolicyRuleType = "RoleManagementPolicyExpirationRule"
	RoleManagementPolicyRuleTypeRoleManagementPolicyNotificationRule          RoleManagementPolicyRuleType = "RoleManagementPolicyNotificationRule"
	RoleManagementPolicyRuleTypeRoleManagementPolicyPimOnlyModeRule           RoleManagementPolicyRuleType = "RoleManagementPolicyPimOnlyModeRule"
)

// PossibleRoleManagementPolicyRuleTypeValues returns the possible values for the RoleManagementPolicyRuleType const type.
func PossibleRoleManagementPolicyRuleTypeValues() []RoleManagementPolicyRuleType {
	return []RoleManagementPolicyRuleType{
		RoleManagementPolicyRuleTypeRoleManagementPolicyApprovalRule,
		RoleManagementPolicyRuleTypeRoleManagementPolicyAuthenticationContextRule,
		RoleManagementPolicyRuleTypeRoleManagementPolicyEnablementRule,
		RoleManagementPolicyRuleTypeRoleManagementPolicyExpirationRule,
		RoleManagementPolicyRuleTypeRoleManagementPolicyNotificationRule,
		RoleManagementPolicyRuleTypeRoleManagementPolicyPimOnlyModeRule,
	}
}

// SeverityLevel - Severity level of the alert.
type SeverityLevel string

const (
	SeverityLevelHigh   SeverityLevel = "High"
	SeverityLevelLow    SeverityLevel = "Low"
	SeverityLevelMedium SeverityLevel = "Medium"
)

// PossibleSeverityLevelValues returns the possible values for the SeverityLevel const type.
func PossibleSeverityLevelValues() []SeverityLevel {
	return []SeverityLevel{
		SeverityLevelHigh,
		SeverityLevelLow,
		SeverityLevelMedium,
	}
}

// Status - The status of the role assignment schedule.
type Status string

const (
	StatusAccepted                    Status = "Accepted"
	StatusAdminApproved               Status = "AdminApproved"
	StatusAdminDenied                 Status = "AdminDenied"
	StatusCanceled                    Status = "Canceled"
	StatusDenied                      Status = "Denied"
	StatusFailed                      Status = "Failed"
	StatusFailedAsResourceIsLocked    Status = "FailedAsResourceIsLocked"
	StatusGranted                     Status = "Granted"
	StatusInvalid                     Status = "Invalid"
	StatusPendingAdminDecision        Status = "PendingAdminDecision"
	StatusPendingApproval             Status = "PendingApproval"
	StatusPendingApprovalProvisioning Status = "PendingApprovalProvisioning"
	StatusPendingEvaluation           Status = "PendingEvaluation"
	StatusPendingExternalProvisioning Status = "PendingExternalProvisioning"
	StatusPendingProvisioning         Status = "PendingProvisioning"
	StatusPendingRevocation           Status = "PendingRevocation"
	StatusPendingScheduleCreation     Status = "PendingScheduleCreation"
	StatusProvisioned                 Status = "Provisioned"
	StatusProvisioningStarted         Status = "ProvisioningStarted"
	StatusRevoked                     Status = "Revoked"
	StatusScheduleCreated             Status = "ScheduleCreated"
	StatusTimedOut                    Status = "TimedOut"
)

// PossibleStatusValues returns the possible values for the Status const type.
func PossibleStatusValues() []Status {
	return []Status{
		StatusAccepted,
		StatusAdminApproved,
		StatusAdminDenied,
		StatusCanceled,
		StatusDenied,
		StatusFailed,
		StatusFailedAsResourceIsLocked,
		StatusGranted,
		StatusInvalid,
		StatusPendingAdminDecision,
		StatusPendingApproval,
		StatusPendingApprovalProvisioning,
		StatusPendingEvaluation,
		StatusPendingExternalProvisioning,
		StatusPendingProvisioning,
		StatusPendingRevocation,
		StatusPendingScheduleCreation,
		StatusProvisioned,
		StatusProvisioningStarted,
		StatusRevoked,
		StatusScheduleCreated,
		StatusTimedOut,
	}
}

// Type - Type of the role assignment schedule expiration
type Type string

const (
	TypeAfterDateTime Type = "AfterDateTime"
	TypeAfterDuration Type = "AfterDuration"
	TypeNoExpiration  Type = "NoExpiration"
)

// PossibleTypeValues returns the possible values for the Type const type.
func PossibleTypeValues() []Type {
	return []Type{
		TypeAfterDateTime,
		TypeAfterDuration,
		TypeNoExpiration,
	}
}

// UserType - The type of user.
type UserType string

const (
	UserTypeGroup            UserType = "Group"
	UserTypeServicePrincipal UserType = "ServicePrincipal"
	UserTypeUser             UserType = "User"
)

// PossibleUserTypeValues returns the possible values for the UserType const type.
func PossibleUserTypeValues() []UserType {
	return []UserType{
		UserTypeGroup,
		UserTypeServicePrincipal,
		UserTypeUser,
	}
}
