//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsecurity

import "time"

// Baseline details.
type Baseline struct {
	// Expected results.
	ExpectedResults [][]*string `json:"expectedResults,omitempty"`

	// Baseline update time (UTC).
	UpdatedTime *time.Time `json:"updatedTime,omitempty"`
}

// BaselineAdjustedResult - The rule result adjusted with baseline.
type BaselineAdjustedResult struct {
	// Baseline details.
	Baseline *Baseline `json:"baseline,omitempty"`

	// Results the are not in baseline.
	ResultsNotInBaseline [][]*string `json:"resultsNotInBaseline,omitempty"`

	// Results the are in baseline.
	ResultsOnlyInBaseline [][]*string `json:"resultsOnlyInBaseline,omitempty"`

	// The rule result status.
	Status *RuleStatus `json:"status,omitempty"`
}

// BenchmarkReference - The benchmark references.
type BenchmarkReference struct {
	// The benchmark name.
	Benchmark *string `json:"benchmark,omitempty"`

	// The benchmark reference.
	Reference *string `json:"reference,omitempty"`
}

// ErrorAdditionalInfo - The resource management error additional info.
type ErrorAdditionalInfo struct {
	// READ-ONLY; The additional info.
	Info any `json:"info,omitempty" azure:"ro"`

	// READ-ONLY; The additional info type.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// QueryCheck - The rule query details.
type QueryCheck struct {
	// Column names of expected result.
	ColumnNames []*string `json:"columnNames,omitempty"`

	// Expected result.
	ExpectedResult [][]*string `json:"expectedResult,omitempty"`

	// The rule query.
	Query *string `json:"query,omitempty"`
}

// Remediation details.
type Remediation struct {
	// Is remediation automated.
	Automated *bool `json:"automated,omitempty"`

	// Remediation description.
	Description *string `json:"description,omitempty"`

	// Optional link to remediate in Azure Portal.
	PortalLink *string `json:"portalLink,omitempty"`

	// Remediation script.
	Scripts []*string `json:"scripts,omitempty"`
}

// Resource - Describes an Azure resource.
type Resource struct {
	// READ-ONLY; Resource Id
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Resource name
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Resource type
	Type *string `json:"type,omitempty" azure:"ro"`
}

// RuleResults - Rule results.
type RuleResults struct {
	// Rule results properties.
	Properties *RuleResultsProperties `json:"properties,omitempty"`

	// READ-ONLY; Resource Id
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Resource name
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Resource type
	Type *string `json:"type,omitempty" azure:"ro"`
}

// RuleResultsInput - Rule results input.
type RuleResultsInput struct {
	// Take results from latest scan.
	LatestScan *bool `json:"latestScan,omitempty"`

	// Expected results to be inserted into the baseline. Leave this field empty it LatestScan == true.
	Results [][]*string `json:"results,omitempty"`
}

// RuleResultsProperties - Rule results properties.
type RuleResultsProperties struct {
	// Expected results in the baseline.
	Results [][]*string `json:"results,omitempty"`
}

// RulesResults - A list of rules results.
type RulesResults struct {
	// List of rule results.
	Value []*RuleResults `json:"value,omitempty"`
}

// RulesResultsInput - Rules results input.
type RulesResultsInput struct {
	// Take results from latest scan.
	LatestScan *bool `json:"latestScan,omitempty"`

	// Expected results to be inserted into the baseline. Leave this field empty it LatestScan == true.
	Results map[string][][]*string `json:"results,omitempty"`
}

// SQLVulnerabilityAssessmentBaselineRulesClientAddOptions contains the optional parameters for the SQLVulnerabilityAssessmentBaselineRulesClient.Add
// method.
type SQLVulnerabilityAssessmentBaselineRulesClientAddOptions struct {
	// The baseline rules.
	Body *RulesResultsInput
}

// SQLVulnerabilityAssessmentBaselineRulesClientCreateOrUpdateOptions contains the optional parameters for the SQLVulnerabilityAssessmentBaselineRulesClient.CreateOrUpdate
// method.
type SQLVulnerabilityAssessmentBaselineRulesClientCreateOrUpdateOptions struct {
	// The baseline results for this rule.
	Body *RuleResultsInput
}

// SQLVulnerabilityAssessmentBaselineRulesClientDeleteOptions contains the optional parameters for the SQLVulnerabilityAssessmentBaselineRulesClient.Delete
// method.
type SQLVulnerabilityAssessmentBaselineRulesClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// SQLVulnerabilityAssessmentBaselineRulesClientGetOptions contains the optional parameters for the SQLVulnerabilityAssessmentBaselineRulesClient.Get
// method.
type SQLVulnerabilityAssessmentBaselineRulesClientGetOptions struct {
	// placeholder for future optional parameters
}

// SQLVulnerabilityAssessmentBaselineRulesClientListOptions contains the optional parameters for the SQLVulnerabilityAssessmentBaselineRulesClient.List
// method.
type SQLVulnerabilityAssessmentBaselineRulesClientListOptions struct {
	// placeholder for future optional parameters
}

// SQLVulnerabilityAssessmentScanResultsClientGetOptions contains the optional parameters for the SQLVulnerabilityAssessmentScanResultsClient.Get
// method.
type SQLVulnerabilityAssessmentScanResultsClientGetOptions struct {
	// placeholder for future optional parameters
}

// SQLVulnerabilityAssessmentScanResultsClientListOptions contains the optional parameters for the SQLVulnerabilityAssessmentScanResultsClient.List
// method.
type SQLVulnerabilityAssessmentScanResultsClientListOptions struct {
	// placeholder for future optional parameters
}

// SQLVulnerabilityAssessmentScansClientGetOptions contains the optional parameters for the SQLVulnerabilityAssessmentScansClient.Get
// method.
type SQLVulnerabilityAssessmentScansClientGetOptions struct {
	// placeholder for future optional parameters
}

// SQLVulnerabilityAssessmentScansClientListOptions contains the optional parameters for the SQLVulnerabilityAssessmentScansClient.List
// method.
type SQLVulnerabilityAssessmentScansClientListOptions struct {
	// placeholder for future optional parameters
}

// ScanPropertiesV2 - A vulnerability assessment scan record properties.
type ScanPropertiesV2 struct {
	// The database name.
	Database *string `json:"database,omitempty"`

	// Scan results are valid until end time (UTC).
	EndTime *time.Time `json:"endTime,omitempty"`

	// The number of failed rules with high severity.
	HighSeverityFailedRulesCount *int32 `json:"highSeverityFailedRulesCount,omitempty"`

	// Baseline created for this database, and has one or more rules.
	IsBaselineApplied *bool `json:"isBaselineApplied,omitempty"`

	// Last scan time.
	LastScanTime *time.Time `json:"lastScanTime,omitempty"`

	// The number of failed rules with low severity.
	LowSeverityFailedRulesCount *int32 `json:"lowSeverityFailedRulesCount,omitempty"`

	// The number of failed rules with medium severity.
	MediumSeverityFailedRulesCount *int32 `json:"mediumSeverityFailedRulesCount,omitempty"`

	// The SQL version.
	SQLVersion *string `json:"sqlVersion,omitempty"`

	// The server name.
	Server *string `json:"server,omitempty"`

	// The scan start time (UTC).
	StartTime *time.Time `json:"startTime,omitempty"`

	// The scan status.
	State *ScanState `json:"state,omitempty"`

	// The number of total failed rules.
	TotalFailedRulesCount *int32 `json:"totalFailedRulesCount,omitempty"`

	// The number of total passed rules.
	TotalPassedRulesCount *int32 `json:"totalPassedRulesCount,omitempty"`

	// The number of total rules assessed.
	TotalRulesCount *int32 `json:"totalRulesCount,omitempty"`

	// The scan trigger type.
	TriggerType *ScanTriggerType `json:"triggerType,omitempty"`
}

// ScanResult - A vulnerability assessment scan result for a single rule.
type ScanResult struct {
	// A vulnerability assessment scan result properties for a single rule.
	Properties *ScanResultProperties `json:"properties,omitempty"`

	// READ-ONLY; Resource Id
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Resource name
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Resource type
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ScanResultProperties - A vulnerability assessment scan result properties for a single rule.
type ScanResultProperties struct {
	// The rule result adjusted with baseline.
	BaselineAdjustedResult *BaselineAdjustedResult `json:"baselineAdjustedResult,omitempty"`

	// Indicated whether the results specified here are trimmed.
	IsTrimmed *bool `json:"isTrimmed,omitempty"`

	// The results of the query that was run.
	QueryResults [][]*string `json:"queryResults,omitempty"`

	// Remediation details.
	Remediation *Remediation `json:"remediation,omitempty"`

	// The rule Id.
	RuleID *string `json:"ruleId,omitempty"`

	// vulnerability assessment rule metadata details.
	RuleMetadata *VaRule `json:"ruleMetadata,omitempty"`

	// The rule result status.
	Status *RuleStatus `json:"status,omitempty"`
}

// ScanResults - A list of vulnerability assessment scan results.
type ScanResults struct {
	// List of vulnerability assessment scan results.
	Value []*ScanResult `json:"value,omitempty"`
}

// ScanV2 - A vulnerability assessment scan record.
type ScanV2 struct {
	// A vulnerability assessment scan record properties.
	Properties *ScanPropertiesV2 `json:"properties,omitempty"`

	// READ-ONLY; Resource Id
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Resource name
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Resource type
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ScansV2 - A list of vulnerability assessment scan records.
type ScansV2 struct {
	// List of vulnerability assessment scan records.
	Value []*ScanV2 `json:"value,omitempty"`
}

// VaRule - vulnerability assessment rule metadata details.
type VaRule struct {
	// The benchmark references.
	BenchmarkReferences []*BenchmarkReference `json:"benchmarkReferences,omitempty"`

	// The rule category.
	Category *string `json:"category,omitempty"`

	// The rule description.
	Description *string `json:"description,omitempty"`

	// The rule query details.
	QueryCheck *QueryCheck `json:"queryCheck,omitempty"`

	// The rule rationale.
	Rationale *string `json:"rationale,omitempty"`

	// The rule Id.
	RuleID *string `json:"ruleId,omitempty"`

	// The rule type.
	RuleType *RuleType `json:"ruleType,omitempty"`

	// The rule severity.
	Severity *RuleSeverity `json:"severity,omitempty"`

	// The rule title.
	Title *string `json:"title,omitempty"`
}
