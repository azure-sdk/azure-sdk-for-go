//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armalertsmanagement

const (
	moduleName    = "armalertsmanagement"
	moduleVersion = "v0.5.1"
)

// ActionType - Action that should be applied.
type ActionType string

const (
	ActionTypeAddActionGroups       ActionType = "AddActionGroups"
	ActionTypeRemoveAllActionGroups ActionType = "RemoveAllActionGroups"
)

// PossibleActionTypeValues returns the possible values for the ActionType const type.
func PossibleActionTypeValues() []ActionType {
	return []ActionType{
		ActionTypeAddActionGroups,
		ActionTypeRemoveAllActionGroups,
	}
}

// AlertModificationEvent - Reason for the modification
type AlertModificationEvent string

const (
	AlertModificationEventAlertCreated           AlertModificationEvent = "AlertCreated"
	AlertModificationEventStateChange            AlertModificationEvent = "StateChange"
	AlertModificationEventMonitorConditionChange AlertModificationEvent = "MonitorConditionChange"
	AlertModificationEventSeverityChange         AlertModificationEvent = "SeverityChange"
	AlertModificationEventActionRuleTriggered    AlertModificationEvent = "ActionRuleTriggered"
	AlertModificationEventActionRuleSuppressed   AlertModificationEvent = "ActionRuleSuppressed"
	AlertModificationEventActionsTriggered       AlertModificationEvent = "ActionsTriggered"
	AlertModificationEventActionsSuppressed      AlertModificationEvent = "ActionsSuppressed"
	AlertModificationEventActionsFailed          AlertModificationEvent = "ActionsFailed"
)

// PossibleAlertModificationEventValues returns the possible values for the AlertModificationEvent const type.
func PossibleAlertModificationEventValues() []AlertModificationEvent {
	return []AlertModificationEvent{
		AlertModificationEventAlertCreated,
		AlertModificationEventStateChange,
		AlertModificationEventMonitorConditionChange,
		AlertModificationEventSeverityChange,
		AlertModificationEventActionRuleTriggered,
		AlertModificationEventActionRuleSuppressed,
		AlertModificationEventActionsTriggered,
		AlertModificationEventActionsSuppressed,
		AlertModificationEventActionsFailed,
	}
}

type AlertState string

const (
	AlertStateAcknowledged AlertState = "Acknowledged"
	AlertStateClosed       AlertState = "Closed"
	AlertStateNew          AlertState = "New"
)

// PossibleAlertStateValues returns the possible values for the AlertState const type.
func PossibleAlertStateValues() []AlertState {
	return []AlertState{
		AlertStateAcknowledged,
		AlertStateClosed,
		AlertStateNew,
	}
}

type AlertsSortByFields string

const (
	AlertsSortByFieldsAlertState           AlertsSortByFields = "alertState"
	AlertsSortByFieldsLastModifiedDateTime AlertsSortByFields = "lastModifiedDateTime"
	AlertsSortByFieldsMonitorCondition     AlertsSortByFields = "monitorCondition"
	AlertsSortByFieldsName                 AlertsSortByFields = "name"
	AlertsSortByFieldsSeverity             AlertsSortByFields = "severity"
	AlertsSortByFieldsStartDateTime        AlertsSortByFields = "startDateTime"
	AlertsSortByFieldsTargetResource       AlertsSortByFields = "targetResource"
	AlertsSortByFieldsTargetResourceGroup  AlertsSortByFields = "targetResourceGroup"
	AlertsSortByFieldsTargetResourceName   AlertsSortByFields = "targetResourceName"
	AlertsSortByFieldsTargetResourceType   AlertsSortByFields = "targetResourceType"
)

// PossibleAlertsSortByFieldsValues returns the possible values for the AlertsSortByFields const type.
func PossibleAlertsSortByFieldsValues() []AlertsSortByFields {
	return []AlertsSortByFields{
		AlertsSortByFieldsAlertState,
		AlertsSortByFieldsLastModifiedDateTime,
		AlertsSortByFieldsMonitorCondition,
		AlertsSortByFieldsName,
		AlertsSortByFieldsSeverity,
		AlertsSortByFieldsStartDateTime,
		AlertsSortByFieldsTargetResource,
		AlertsSortByFieldsTargetResourceGroup,
		AlertsSortByFieldsTargetResourceName,
		AlertsSortByFieldsTargetResourceType,
	}
}

type AlertsSummaryGroupByFields string

const (
	AlertsSummaryGroupByFieldsAlertRule        AlertsSummaryGroupByFields = "alertRule"
	AlertsSummaryGroupByFieldsAlertState       AlertsSummaryGroupByFields = "alertState"
	AlertsSummaryGroupByFieldsMonitorCondition AlertsSummaryGroupByFields = "monitorCondition"
	AlertsSummaryGroupByFieldsMonitorService   AlertsSummaryGroupByFields = "monitorService"
	AlertsSummaryGroupByFieldsSeverity         AlertsSummaryGroupByFields = "severity"
	AlertsSummaryGroupByFieldsSignalType       AlertsSummaryGroupByFields = "signalType"
)

// PossibleAlertsSummaryGroupByFieldsValues returns the possible values for the AlertsSummaryGroupByFields const type.
func PossibleAlertsSummaryGroupByFieldsValues() []AlertsSummaryGroupByFields {
	return []AlertsSummaryGroupByFields{
		AlertsSummaryGroupByFieldsAlertRule,
		AlertsSummaryGroupByFieldsAlertState,
		AlertsSummaryGroupByFieldsMonitorCondition,
		AlertsSummaryGroupByFieldsMonitorService,
		AlertsSummaryGroupByFieldsSeverity,
		AlertsSummaryGroupByFieldsSignalType,
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

// DaysOfWeek - Days of week.
type DaysOfWeek string

const (
	DaysOfWeekFriday    DaysOfWeek = "Friday"
	DaysOfWeekMonday    DaysOfWeek = "Monday"
	DaysOfWeekSaturday  DaysOfWeek = "Saturday"
	DaysOfWeekSunday    DaysOfWeek = "Sunday"
	DaysOfWeekThursday  DaysOfWeek = "Thursday"
	DaysOfWeekTuesday   DaysOfWeek = "Tuesday"
	DaysOfWeekWednesday DaysOfWeek = "Wednesday"
)

// PossibleDaysOfWeekValues returns the possible values for the DaysOfWeek const type.
func PossibleDaysOfWeekValues() []DaysOfWeek {
	return []DaysOfWeek{
		DaysOfWeekFriday,
		DaysOfWeekMonday,
		DaysOfWeekSaturday,
		DaysOfWeekSunday,
		DaysOfWeekThursday,
		DaysOfWeekTuesday,
		DaysOfWeekWednesday,
	}
}

// Field - Field for a given condition.
type Field string

const (
	FieldAlertContext        Field = "AlertContext"
	FieldAlertRuleID         Field = "AlertRuleId"
	FieldAlertRuleName       Field = "AlertRuleName"
	FieldDescription         Field = "Description"
	FieldMonitorCondition    Field = "MonitorCondition"
	FieldMonitorService      Field = "MonitorService"
	FieldSeverity            Field = "Severity"
	FieldSignalType          Field = "SignalType"
	FieldTargetResource      Field = "TargetResource"
	FieldTargetResourceGroup Field = "TargetResourceGroup"
	FieldTargetResourceType  Field = "TargetResourceType"
)

// PossibleFieldValues returns the possible values for the Field const type.
func PossibleFieldValues() []Field {
	return []Field{
		FieldAlertContext,
		FieldAlertRuleID,
		FieldAlertRuleName,
		FieldDescription,
		FieldMonitorCondition,
		FieldMonitorService,
		FieldSeverity,
		FieldSignalType,
		FieldTargetResource,
		FieldTargetResourceGroup,
		FieldTargetResourceType,
	}
}

type Identifier string

const (
	IdentifierMonitorServiceList Identifier = "MonitorServiceList"
)

// PossibleIdentifierValues returns the possible values for the Identifier const type.
func PossibleIdentifierValues() []Identifier {
	return []Identifier{
		IdentifierMonitorServiceList,
	}
}

// MetadataIdentifier - Identification of the information to be retrieved by API call
type MetadataIdentifier string

const (
	MetadataIdentifierMonitorServiceList MetadataIdentifier = "MonitorServiceList"
)

// PossibleMetadataIdentifierValues returns the possible values for the MetadataIdentifier const type.
func PossibleMetadataIdentifierValues() []MetadataIdentifier {
	return []MetadataIdentifier{
		MetadataIdentifierMonitorServiceList,
	}
}

type MonitorCondition string

const (
	MonitorConditionFired    MonitorCondition = "Fired"
	MonitorConditionResolved MonitorCondition = "Resolved"
)

// PossibleMonitorConditionValues returns the possible values for the MonitorCondition const type.
func PossibleMonitorConditionValues() []MonitorCondition {
	return []MonitorCondition{
		MonitorConditionFired,
		MonitorConditionResolved,
	}
}

type MonitorService string

const (
	MonitorServiceActivityLogAdministrative MonitorService = "ActivityLog Administrative"
	MonitorServiceActivityLogAutoscale      MonitorService = "ActivityLog Autoscale"
	MonitorServiceActivityLogPolicy         MonitorService = "ActivityLog Policy"
	MonitorServiceActivityLogRecommendation MonitorService = "ActivityLog Recommendation"
	MonitorServiceActivityLogSecurity       MonitorService = "ActivityLog Security"
	MonitorServiceApplicationInsights       MonitorService = "Application Insights"
	MonitorServiceLogAnalytics              MonitorService = "Log Analytics"
	MonitorServiceNagios                    MonitorService = "Nagios"
	MonitorServicePlatform                  MonitorService = "Platform"
	MonitorServiceSCOM                      MonitorService = "SCOM"
	MonitorServiceServiceHealth             MonitorService = "ServiceHealth"
	MonitorServiceSmartDetector             MonitorService = "SmartDetector"
	MonitorServiceVMInsights                MonitorService = "VM Insights"
	MonitorServiceZabbix                    MonitorService = "Zabbix"
)

// PossibleMonitorServiceValues returns the possible values for the MonitorService const type.
func PossibleMonitorServiceValues() []MonitorService {
	return []MonitorService{
		MonitorServiceActivityLogAdministrative,
		MonitorServiceActivityLogAutoscale,
		MonitorServiceActivityLogPolicy,
		MonitorServiceActivityLogRecommendation,
		MonitorServiceActivityLogSecurity,
		MonitorServiceApplicationInsights,
		MonitorServiceLogAnalytics,
		MonitorServiceNagios,
		MonitorServicePlatform,
		MonitorServiceSCOM,
		MonitorServiceServiceHealth,
		MonitorServiceSmartDetector,
		MonitorServiceVMInsights,
		MonitorServiceZabbix,
	}
}

// Operator - Operator for a given condition.
type Operator string

const (
	OperatorContains       Operator = "Contains"
	OperatorDoesNotContain Operator = "DoesNotContain"
	OperatorEquals         Operator = "Equals"
	OperatorNotEquals      Operator = "NotEquals"
)

// PossibleOperatorValues returns the possible values for the Operator const type.
func PossibleOperatorValues() []Operator {
	return []Operator{
		OperatorContains,
		OperatorDoesNotContain,
		OperatorEquals,
		OperatorNotEquals,
	}
}

// RecurrenceType - Specifies when the recurrence should be applied.
type RecurrenceType string

const (
	RecurrenceTypeDaily   RecurrenceType = "Daily"
	RecurrenceTypeMonthly RecurrenceType = "Monthly"
	RecurrenceTypeWeekly  RecurrenceType = "Weekly"
)

// PossibleRecurrenceTypeValues returns the possible values for the RecurrenceType const type.
func PossibleRecurrenceTypeValues() []RecurrenceType {
	return []RecurrenceType{
		RecurrenceTypeDaily,
		RecurrenceTypeMonthly,
		RecurrenceTypeWeekly,
	}
}

type Severity string

const (
	SeveritySev0 Severity = "Sev0"
	SeveritySev1 Severity = "Sev1"
	SeveritySev2 Severity = "Sev2"
	SeveritySev3 Severity = "Sev3"
	SeveritySev4 Severity = "Sev4"
)

// PossibleSeverityValues returns the possible values for the Severity const type.
func PossibleSeverityValues() []Severity {
	return []Severity{
		SeveritySev0,
		SeveritySev1,
		SeveritySev2,
		SeveritySev3,
		SeveritySev4,
	}
}

// SignalType - The type of signal the alert is based on, which could be metrics, logs or activity logs.
type SignalType string

const (
	SignalTypeLog     SignalType = "Log"
	SignalTypeMetric  SignalType = "Metric"
	SignalTypeUnknown SignalType = "Unknown"
)

// PossibleSignalTypeValues returns the possible values for the SignalType const type.
func PossibleSignalTypeValues() []SignalType {
	return []SignalType{
		SignalTypeLog,
		SignalTypeMetric,
		SignalTypeUnknown,
	}
}

// SmartGroupModificationEvent - Reason for the modification
type SmartGroupModificationEvent string

const (
	SmartGroupModificationEventSmartGroupCreated SmartGroupModificationEvent = "SmartGroupCreated"
	SmartGroupModificationEventStateChange       SmartGroupModificationEvent = "StateChange"
	SmartGroupModificationEventAlertAdded        SmartGroupModificationEvent = "AlertAdded"
	SmartGroupModificationEventAlertRemoved      SmartGroupModificationEvent = "AlertRemoved"
)

// PossibleSmartGroupModificationEventValues returns the possible values for the SmartGroupModificationEvent const type.
func PossibleSmartGroupModificationEventValues() []SmartGroupModificationEvent {
	return []SmartGroupModificationEvent{
		SmartGroupModificationEventSmartGroupCreated,
		SmartGroupModificationEventStateChange,
		SmartGroupModificationEventAlertAdded,
		SmartGroupModificationEventAlertRemoved,
	}
}

type SmartGroupsSortByFields string

const (
	SmartGroupsSortByFieldsAlertsCount          SmartGroupsSortByFields = "alertsCount"
	SmartGroupsSortByFieldsLastModifiedDateTime SmartGroupsSortByFields = "lastModifiedDateTime"
	SmartGroupsSortByFieldsSeverity             SmartGroupsSortByFields = "severity"
	SmartGroupsSortByFieldsStartDateTime        SmartGroupsSortByFields = "startDateTime"
	SmartGroupsSortByFieldsState                SmartGroupsSortByFields = "state"
)

// PossibleSmartGroupsSortByFieldsValues returns the possible values for the SmartGroupsSortByFields const type.
func PossibleSmartGroupsSortByFieldsValues() []SmartGroupsSortByFields {
	return []SmartGroupsSortByFields{
		SmartGroupsSortByFieldsAlertsCount,
		SmartGroupsSortByFieldsLastModifiedDateTime,
		SmartGroupsSortByFieldsSeverity,
		SmartGroupsSortByFieldsStartDateTime,
		SmartGroupsSortByFieldsState,
	}
}

type SortOrder string

const (
	SortOrderAsc  SortOrder = "asc"
	SortOrderDesc SortOrder = "desc"
)

// PossibleSortOrderValues returns the possible values for the SortOrder const type.
func PossibleSortOrderValues() []SortOrder {
	return []SortOrder{
		SortOrderAsc,
		SortOrderDesc,
	}
}

// State - Smart group state
type State string

const (
	StateAcknowledged State = "Acknowledged"
	StateClosed       State = "Closed"
	StateNew          State = "New"
)

// PossibleStateValues returns the possible values for the State const type.
func PossibleStateValues() []State {
	return []State{
		StateAcknowledged,
		StateClosed,
		StateNew,
	}
}

type TimeRange string

const (
	TimeRangeOneD    TimeRange = "1d"
	TimeRangeOneH    TimeRange = "1h"
	TimeRangeSevenD  TimeRange = "7d"
	TimeRangeThirtyD TimeRange = "30d"
)

// PossibleTimeRangeValues returns the possible values for the TimeRange const type.
func PossibleTimeRangeValues() []TimeRange {
	return []TimeRange{
		TimeRangeOneD,
		TimeRangeOneH,
		TimeRangeSevenD,
		TimeRangeThirtyD,
	}
}
