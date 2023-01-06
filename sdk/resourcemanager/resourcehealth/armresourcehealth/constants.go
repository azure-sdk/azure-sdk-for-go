//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armresourcehealth

const (
	moduleName    = "armresourcehealth"
	moduleVersion = "v2.0.0-beta.1"
)

// AvailabilityStateValues - Impacted resource status of the resource.
type AvailabilityStateValues string

const (
	AvailabilityStateValuesAvailable   AvailabilityStateValues = "Available"
	AvailabilityStateValuesDegraded    AvailabilityStateValues = "Degraded"
	AvailabilityStateValuesUnavailable AvailabilityStateValues = "Unavailable"
	AvailabilityStateValuesUnknown     AvailabilityStateValues = "Unknown"
)

// PossibleAvailabilityStateValuesValues returns the possible values for the AvailabilityStateValues const type.
func PossibleAvailabilityStateValuesValues() []AvailabilityStateValues {
	return []AvailabilityStateValues{
		AvailabilityStateValuesAvailable,
		AvailabilityStateValuesDegraded,
		AvailabilityStateValuesUnavailable,
		AvailabilityStateValuesUnknown,
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

// EventLevelValues - Level of event.
type EventLevelValues string

const (
	EventLevelValuesCritical      EventLevelValues = "Critical"
	EventLevelValuesError         EventLevelValues = "Error"
	EventLevelValuesInformational EventLevelValues = "Informational"
	EventLevelValuesWarning       EventLevelValues = "Warning"
)

// PossibleEventLevelValuesValues returns the possible values for the EventLevelValues const type.
func PossibleEventLevelValuesValues() []EventLevelValues {
	return []EventLevelValues{
		EventLevelValuesCritical,
		EventLevelValuesError,
		EventLevelValuesInformational,
		EventLevelValuesWarning,
	}
}

// EventSourceValues - Source of event.
type EventSourceValues string

const (
	EventSourceValuesResourceHealth EventSourceValues = "ResourceHealth"
	EventSourceValuesServiceHealth  EventSourceValues = "ServiceHealth"
)

// PossibleEventSourceValuesValues returns the possible values for the EventSourceValues const type.
func PossibleEventSourceValuesValues() []EventSourceValues {
	return []EventSourceValues{
		EventSourceValuesResourceHealth,
		EventSourceValuesServiceHealth,
	}
}

// EventStatusValues - Current status of event.
type EventStatusValues string

const (
	EventStatusValuesActive   EventStatusValues = "Active"
	EventStatusValuesResolved EventStatusValues = "Resolved"
)

// PossibleEventStatusValuesValues returns the possible values for the EventStatusValues const type.
func PossibleEventStatusValuesValues() []EventStatusValues {
	return []EventStatusValues{
		EventStatusValuesActive,
		EventStatusValuesResolved,
	}
}

// EventTypeValues - Type of event.
type EventTypeValues string

const (
	EventTypeValuesEmergingIssues     EventTypeValues = "EmergingIssues"
	EventTypeValuesHealthAdvisory     EventTypeValues = "HealthAdvisory"
	EventTypeValuesPlannedMaintenance EventTypeValues = "PlannedMaintenance"
	EventTypeValuesRCA                EventTypeValues = "RCA"
	EventTypeValuesSecurityAdvisory   EventTypeValues = "SecurityAdvisory"
	EventTypeValuesServiceIssue       EventTypeValues = "ServiceIssue"
)

// PossibleEventTypeValuesValues returns the possible values for the EventTypeValues const type.
func PossibleEventTypeValuesValues() []EventTypeValues {
	return []EventTypeValues{
		EventTypeValuesEmergingIssues,
		EventTypeValuesHealthAdvisory,
		EventTypeValuesPlannedMaintenance,
		EventTypeValuesRCA,
		EventTypeValuesSecurityAdvisory,
		EventTypeValuesServiceIssue,
	}
}

type IssueNameParameter string

const (
	IssueNameParameterDefault IssueNameParameter = "default"
)

// PossibleIssueNameParameterValues returns the possible values for the IssueNameParameter const type.
func PossibleIssueNameParameterValues() []IssueNameParameter {
	return []IssueNameParameter{
		IssueNameParameterDefault,
	}
}

// LevelValues - Level of insight.
type LevelValues string

const (
	LevelValuesCritical LevelValues = "Critical"
	LevelValuesWarning  LevelValues = "Warning"
)

// PossibleLevelValuesValues returns the possible values for the LevelValues const type.
func PossibleLevelValuesValues() []LevelValues {
	return []LevelValues{
		LevelValuesCritical,
		LevelValuesWarning,
	}
}

// LinkTypeValues - Type of link.
type LinkTypeValues string

const (
	LinkTypeValuesButton    LinkTypeValues = "Button"
	LinkTypeValuesHyperlink LinkTypeValues = "Hyperlink"
)

// PossibleLinkTypeValuesValues returns the possible values for the LinkTypeValues const type.
func PossibleLinkTypeValuesValues() []LinkTypeValues {
	return []LinkTypeValues{
		LinkTypeValuesButton,
		LinkTypeValuesHyperlink,
	}
}

// ReasonChronicityTypes - Chronicity of the availability transition.
type ReasonChronicityTypes string

const (
	ReasonChronicityTypesPersistent ReasonChronicityTypes = "Persistent"
	ReasonChronicityTypesTransient  ReasonChronicityTypes = "Transient"
)

// PossibleReasonChronicityTypesValues returns the possible values for the ReasonChronicityTypes const type.
func PossibleReasonChronicityTypesValues() []ReasonChronicityTypes {
	return []ReasonChronicityTypes{
		ReasonChronicityTypesPersistent,
		ReasonChronicityTypesTransient,
	}
}

// ReasonTypeValues - When the resource's availabilityState is Unavailable, it describes where the health impacting event
// was originated.
type ReasonTypeValues string

const (
	ReasonTypeValuesPlanned       ReasonTypeValues = "Planned"
	ReasonTypeValuesUnplanned     ReasonTypeValues = "Unplanned"
	ReasonTypeValuesUserInitiated ReasonTypeValues = "UserInitiated"
)

// PossibleReasonTypeValuesValues returns the possible values for the ReasonTypeValues const type.
func PossibleReasonTypeValuesValues() []ReasonTypeValues {
	return []ReasonTypeValues{
		ReasonTypeValuesPlanned,
		ReasonTypeValuesUnplanned,
		ReasonTypeValuesUserInitiated,
	}
}

// SeverityValues - The severity level of this active event.
type SeverityValues string

const (
	SeverityValuesError       SeverityValues = "Error"
	SeverityValuesInformation SeverityValues = "Information"
	SeverityValuesWarning     SeverityValues = "Warning"
)

// PossibleSeverityValuesValues returns the possible values for the SeverityValues const type.
func PossibleSeverityValuesValues() []SeverityValues {
	return []SeverityValues{
		SeverityValuesError,
		SeverityValuesInformation,
		SeverityValuesWarning,
	}
}

// StageValues - The stage of this active event.
type StageValues string

const (
	StageValuesActive   StageValues = "Active"
	StageValuesArchived StageValues = "Archived"
	StageValuesResolve  StageValues = "Resolve"
)

// PossibleStageValuesValues returns the possible values for the StageValues const type.
func PossibleStageValuesValues() []StageValues {
	return []StageValues{
		StageValuesActive,
		StageValuesArchived,
		StageValuesResolve,
	}
}
