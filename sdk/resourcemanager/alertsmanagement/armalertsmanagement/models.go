// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armalertsmanagement

import "time"

// Action to be applied.
type Action struct {
	// REQUIRED; Action that should be applied.
	ActionType *ActionType
}

// GetAction implements the ActionClassification interface for type Action.
func (a *Action) GetAction() *Action { return a }

// ActionStatus - Action status
type ActionStatus struct {
	// Value indicating whether alert is suppressed.
	IsSuppressed *bool
}

// AddActionGroups - Add action groups to alert processing rule.
type AddActionGroups struct {
	// REQUIRED; List of action group Ids to add to alert processing rule.
	ActionGroupIDs []*string

	// REQUIRED; Action that should be applied.
	ActionType *ActionType
}

// GetAction implements the ActionClassification interface for type AddActionGroups.
func (a *AddActionGroups) GetAction() *Action {
	return &Action{
		ActionType: a.ActionType,
	}
}

// Alert - An alert created in alert management service.
type Alert struct {
	// Alert property bag
	Properties *AlertProperties

	// READ-ONLY; Azure resource Id
	ID *string

	// READ-ONLY; Azure resource name
	Name *string

	// READ-ONLY; Azure resource type
	Type *string
}

// AlertEnrichmentItem - Alert enrichment item.
type AlertEnrichmentItem struct {
	// REQUIRED; The enrichment description.
	Description *string

	// REQUIRED; The status of the evaluation of the enrichment.
	Status *Status

	// REQUIRED; The enrichment title.
	Title *string

	// REQUIRED; The enrichment type.
	Type *Type

	// The error message. Will be present only if the status is 'Failed'.
	ErrorMessage *string
}

// GetAlertEnrichmentItem implements the AlertEnrichmentItemClassification interface for type AlertEnrichmentItem.
func (a *AlertEnrichmentItem) GetAlertEnrichmentItem() *AlertEnrichmentItem { return a }

// AlertEnrichmentProperties - Properties of the alert enrichment item.
type AlertEnrichmentProperties struct {
	// Enrichment details
	Enrichments []AlertEnrichmentItemClassification

	// READ-ONLY; Unique Id (GUID) of the alert for which the enrichments are being retrieved.
	AlertID *string
}

// AlertEnrichmentResponse - The alert's enrichments.
type AlertEnrichmentResponse struct {
	// Properties of the alert enrichment item.
	Properties *AlertEnrichmentProperties

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// AlertEnrichmentsList - List the alert's enrichments.
type AlertEnrichmentsList struct {
	// Request URL that can be used to query next page.
	NextLink *string

	// List the alert's enrichments
	Value []*AlertEnrichmentResponse
}

// AlertModification - Alert Modification details
type AlertModification struct {
	// Properties of the alert modification item.
	Properties *AlertModificationProperties

	// READ-ONLY; Azure resource Id
	ID *string

	// READ-ONLY; Azure resource name
	Name *string

	// READ-ONLY; Azure resource type
	Type *string
}

// AlertModificationItem - Alert modification item.
type AlertModificationItem struct {
	// Modification comments
	Comments *string

	// Description of the modification
	Description *string

	// Reason for the modification
	ModificationEvent *AlertModificationEvent

	// Modified date and time
	ModifiedAt *string

	// Modified user details (Principal client name)
	ModifiedBy *string

	// New value
	NewValue *string

	// Old value
	OldValue *string
}

// AlertModificationProperties - Properties of the alert modification item.
type AlertModificationProperties struct {
	// Modification details
	Modifications []*AlertModificationItem

	// READ-ONLY; Unique Id of the alert for which the history is being retrieved
	AlertID *string
}

// AlertProcessingRule - Alert processing rule object containing target scopes, conditions and scheduling logic.
type AlertProcessingRule struct {
	// REQUIRED; Resource location
	Location *string

	// Alert processing rule properties.
	Properties *AlertProcessingRuleProperties

	// Resource tags
	Tags map[string]*string

	// READ-ONLY; Azure resource Id
	ID *string

	// READ-ONLY; Azure resource name
	Name *string

	// READ-ONLY; Alert processing rule system data.
	SystemData *SystemData

	// READ-ONLY; Azure resource type
	Type *string
}

// AlertProcessingRuleProperties - Alert processing rule properties defining scopes, conditions and scheduling logic for alert
// processing rule.
type AlertProcessingRuleProperties struct {
	// REQUIRED; Actions to be applied.
	Actions []ActionClassification

	// REQUIRED; Scopes on which alert processing rule will apply.
	Scopes []*string

	// Conditions on which alerts will be filtered.
	Conditions []*Condition

	// Description of alert processing rule.
	Description *string

	// Indicates if the given alert processing rule is enabled or disabled.
	Enabled *bool

	// Scheduling for alert processing rule.
	Schedule *Schedule
}

// AlertProcessingRulesList - List of alert processing rules.
type AlertProcessingRulesList struct {
	// URL to fetch the next set of alert processing rules.
	NextLink *string

	// List of alert processing rules.
	Value []*AlertProcessingRule
}

// AlertProperties - Alert property bag
type AlertProperties struct {
	// This object contains consistent fields across different monitor services.
	Essentials *Essentials

	// READ-ONLY; Information specific to the monitor service that gives more contextual details about the alert.
	Context any

	// READ-ONLY; Config which would be used for displaying the data in portal.
	EgressConfig any
}

// AlertRuleRecommendationProperties - Describes the format of Alert Rule Recommendations response.
type AlertRuleRecommendationProperties struct {
	// REQUIRED; The recommendation alert rule type.
	AlertRuleType *string

	// REQUIRED; A dictionary that provides the display information for an alert rule recommendation.
	DisplayInformation map[string]*string

	// REQUIRED; A complete ARM template to deploy the alert rules.
	RuleArmTemplate *RuleArmTemplate

	// The recommendation alert rule category.
	Category *string
}

// AlertRuleRecommendationResource - A single alert rule recommendation resource.
type AlertRuleRecommendationResource struct {
	// REQUIRED; recommendation properties.
	Properties *AlertRuleRecommendationProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// AlertRuleRecommendationsListResponse - List of alert rule recommendations.
type AlertRuleRecommendationsListResponse struct {
	// REQUIRED; the values for the alert rule recommendations.
	Value []*AlertRuleRecommendationResource

	// URL to fetch the next set of recommendations.
	NextLink *string
}

// AlertsList - List the alerts.
type AlertsList struct {
	// URL to fetch the next set of alerts.
	NextLink *string

	// List of alerts
	Value []*Alert
}

// AlertsMetaData - alert meta data information.
type AlertsMetaData struct {
	// alert meta data property bag
	Properties AlertsMetaDataPropertiesClassification
}

// AlertsMetaDataProperties - alert meta data property bag
type AlertsMetaDataProperties struct {
	// REQUIRED; Identification of the information to be retrieved by API call
	MetadataIdentifier *MetadataIdentifier
}

// GetAlertsMetaDataProperties implements the AlertsMetaDataPropertiesClassification interface for type AlertsMetaDataProperties.
func (a *AlertsMetaDataProperties) GetAlertsMetaDataProperties() *AlertsMetaDataProperties { return a }

// AlertsSummary - Summary of alerts based on the input filters and 'groupby' parameters.
type AlertsSummary struct {
	// Group the result set.
	Properties *AlertsSummaryGroup

	// READ-ONLY; Azure resource Id
	ID *string

	// READ-ONLY; Azure resource name
	Name *string

	// READ-ONLY; Azure resource type
	Type *string
}

// AlertsSummaryGroup - Group the result set.
type AlertsSummaryGroup struct {
	// Name of the field aggregated
	Groupedby *string

	// Total count of the smart groups.
	SmartGroupsCount *int64

	// Total count of the result set.
	Total *int64

	// List of the items
	Values []*AlertsSummaryGroupItem
}

// AlertsSummaryGroupItem - Alerts summary group item
type AlertsSummaryGroupItem struct {
	// Count of the aggregated field
	Count *int64

	// Name of the field aggregated
	Groupedby *string

	// Value of the aggregated field
	Name *string

	// List of the items
	Values []*AlertsSummaryGroupItem
}

// Comments - Change alert state reason
type Comments struct {
	Comments *string
}

// Condition to trigger an alert processing rule.
type Condition struct {
	// Field for a given condition.
	Field *Field

	// Operator for a given condition.
	Operator *Operator

	// List of values to match for a given condition.
	Values []*string
}

// CorrelationDetails - Correlation details
type CorrelationDetails struct {
	// READ-ONLY; The alert processing rule that was used to correlate this alert. This is an optional field, it will be presented
	// only for a parent alert.
	AlertProcessingRule *string

	// READ-ONLY; The alert's correlation date time in ISO-8601 format.
	CorrelationDateTime *time.Time

	// READ-ONLY; Unique Id (GUID) of the alert that this alert was correlated to.
	ParentAlertID *string
}

// DailyRecurrence - Daily recurrence object.
type DailyRecurrence struct {
	// REQUIRED; Specifies when the recurrence should be applied.
	RecurrenceType *RecurrenceType

	// End time for recurrence.
	EndTime *string

	// Start time for recurrence.
	StartTime *string
}

// GetRecurrence implements the RecurrenceClassification interface for type DailyRecurrence.
func (d *DailyRecurrence) GetRecurrence() *Recurrence {
	return &Recurrence{
		EndTime:        d.EndTime,
		RecurrenceType: d.RecurrenceType,
		StartTime:      d.StartTime,
	}
}

// Essentials - This object contains consistent fields across different monitor services.
type Essentials struct {
	// Action status
	ActionStatus *ActionStatus

	// Correlation details
	CorrelationDetails *CorrelationDetails

	// Alert description.
	Description *string

	// Target ARM resource, on which alert got created.
	TargetResource *string

	// Resource group of target ARM resource, on which alert got created.
	TargetResourceGroup *string

	// Name of the target ARM resource name, on which alert got created.
	TargetResourceName *string

	// Resource type of target ARM resource, on which alert got created.
	TargetResourceType *string

	// READ-ONLY; Rule(monitor) which fired alert instance. Depending on the monitor service, this would be ARM id or name of
	// the rule.
	AlertRule *string

	// READ-ONLY; Alert object state, which can be modified by the user.
	AlertState *AlertState

	// READ-ONLY; Will be presented with the value true only if there are enrichments.
	HasEnrichments *bool

	// READ-ONLY; True if the alert is stateful, and false if it isn't.
	IsStatefulAlert *bool

	// READ-ONLY; Last modification time(ISO-8601 format) of alert instance.
	LastModifiedDateTime *time.Time

	// READ-ONLY; User who last modified the alert, in case of monitor service updates user would be 'system', otherwise name
	// of the user.
	LastModifiedUserName *string

	// READ-ONLY; Condition of the rule at the monitor service. It represents whether the underlying conditions have crossed the
	// defined alert rule thresholds.
	MonitorCondition *MonitorCondition

	// READ-ONLY; Resolved time(ISO-8601 format) of alert instance. This will be updated when monitor service resolves the alert
	// instance because the rule condition is no longer met.
	MonitorConditionResolvedDateTime *time.Time

	// READ-ONLY; Monitor service on which the rule(monitor) is set.
	MonitorService *MonitorService

	// READ-ONLY; Severity of alert Sev0 being highest and Sev4 being lowest.
	Severity *Severity

	// READ-ONLY; The type of signal the alert is based on, which could be metrics, logs or activity logs.
	SignalType *SignalType

	// READ-ONLY; Unique Id of the smart group
	SmartGroupID *string

	// READ-ONLY; Verbose reason describing the reason why this alert instance is added to a smart group
	SmartGroupingReason *string

	// READ-ONLY; Unique Id created by monitor service for each alert instance. This could be used to track the issue at the monitor
	// service, in case of Nagios, Zabbix, SCOM etc.
	SourceCreatedID *string

	// READ-ONLY; Creation time(ISO-8601 format) of alert instance.
	StartDateTime *time.Time
}

// MonitorServiceDetails - Details of a monitor service
type MonitorServiceDetails struct {
	// Monitor service display name
	DisplayName *string

	// Monitor service name
	Name *string
}

// MonitorServiceList - Monitor service details
type MonitorServiceList struct {
	// REQUIRED; Array of operations
	Data []*MonitorServiceDetails

	// REQUIRED; Identification of the information to be retrieved by API call
	MetadataIdentifier *MetadataIdentifier
}

// GetAlertsMetaDataProperties implements the AlertsMetaDataPropertiesClassification interface for type MonitorServiceList.
func (m *MonitorServiceList) GetAlertsMetaDataProperties() *AlertsMetaDataProperties {
	return &AlertsMetaDataProperties{
		MetadataIdentifier: m.MetadataIdentifier,
	}
}

// MonthlyRecurrence - Monthly recurrence object.
type MonthlyRecurrence struct {
	// REQUIRED; Specifies the values for monthly recurrence pattern.
	DaysOfMonth []*int32

	// REQUIRED; Specifies when the recurrence should be applied.
	RecurrenceType *RecurrenceType

	// End time for recurrence.
	EndTime *string

	// Start time for recurrence.
	StartTime *string
}

// GetRecurrence implements the RecurrenceClassification interface for type MonthlyRecurrence.
func (m *MonthlyRecurrence) GetRecurrence() *Recurrence {
	return &Recurrence{
		EndTime:        m.EndTime,
		RecurrenceType: m.RecurrenceType,
		StartTime:      m.StartTime,
	}
}

// Operation provided by provider
type Operation struct {
	// Properties of the operation
	Display *OperationDisplay

	// Name of the operation
	Name *string

	// Origin of the operation
	Origin *string
}

// OperationDisplay - Properties of the operation
type OperationDisplay struct {
	// Description of the operation
	Description *string

	// Operation name
	Operation *string

	// Provider name
	Provider *string

	// Resource name
	Resource *string
}

// OperationsList - Lists the operations available in the AlertsManagement RP.
type OperationsList struct {
	// REQUIRED; Array of operations
	Value []*Operation

	// URL to fetch the next set of alerts.
	NextLink *string
}

// PatchObject - Data contract for patch.
type PatchObject struct {
	// Properties supported by patch operation.
	Properties *PatchProperties

	// Tags to be updated.
	Tags map[string]*string
}

// PatchProperties - Alert processing rule properties supported by patch.
type PatchProperties struct {
	// Indicates if the given alert processing rule is enabled or disabled.
	Enabled *bool
}

// PrometheusEnrichmentItem - Prometheus enrichment object.
type PrometheusEnrichmentItem struct {
	// REQUIRED; An array of the azure monitor workspace resource ids.
	Datasources []*string

	// REQUIRED; The enrichment description.
	Description *string

	// REQUIRED; Partial link to the Grafana explore API.
	GrafanaExplorePath *string

	// REQUIRED; Link to Prometheus query API (Url format).
	LinkToAPI *string

	// REQUIRED; The Prometheus expression query.
	Query *string

	// REQUIRED; The status of the evaluation of the enrichment.
	Status *Status

	// REQUIRED; The enrichment title.
	Title *string

	// REQUIRED; The enrichment type.
	Type *Type

	// The error message. Will be present only if the status is 'Failed'.
	ErrorMessage *string
}

// GetAlertEnrichmentItem implements the AlertEnrichmentItemClassification interface for type PrometheusEnrichmentItem.
func (p *PrometheusEnrichmentItem) GetAlertEnrichmentItem() *AlertEnrichmentItem {
	return &AlertEnrichmentItem{
		Description:  p.Description,
		ErrorMessage: p.ErrorMessage,
		Status:       p.Status,
		Title:        p.Title,
		Type:         p.Type,
	}
}

// GetPrometheusEnrichmentItem implements the PrometheusEnrichmentItemClassification interface for type PrometheusEnrichmentItem.
func (p *PrometheusEnrichmentItem) GetPrometheusEnrichmentItem() *PrometheusEnrichmentItem { return p }

// PrometheusInstantQuery - Prometheus instant query enrichment object.
type PrometheusInstantQuery struct {
	// REQUIRED; An array of the azure monitor workspace resource ids.
	Datasources []*string

	// REQUIRED; The enrichment description.
	Description *string

	// REQUIRED; Partial link to the Grafana explore API.
	GrafanaExplorePath *string

	// REQUIRED; Link to Prometheus query API (Url format).
	LinkToAPI *string

	// REQUIRED; The Prometheus expression query.
	Query *string

	// REQUIRED; The status of the evaluation of the enrichment.
	Status *Status

	// REQUIRED; The date and the time of the evaluation.
	Time *string

	// REQUIRED; The enrichment title.
	Title *string

	// REQUIRED; The enrichment type.
	Type *Type

	// The error message. Will be present only if the status is 'Failed'.
	ErrorMessage *string
}

// GetAlertEnrichmentItem implements the AlertEnrichmentItemClassification interface for type PrometheusInstantQuery.
func (p *PrometheusInstantQuery) GetAlertEnrichmentItem() *AlertEnrichmentItem {
	return &AlertEnrichmentItem{
		Description:  p.Description,
		ErrorMessage: p.ErrorMessage,
		Status:       p.Status,
		Title:        p.Title,
		Type:         p.Type,
	}
}

// GetPrometheusEnrichmentItem implements the PrometheusEnrichmentItemClassification interface for type PrometheusInstantQuery.
func (p *PrometheusInstantQuery) GetPrometheusEnrichmentItem() *PrometheusEnrichmentItem {
	return &PrometheusEnrichmentItem{
		Datasources:        p.Datasources,
		Description:        p.Description,
		ErrorMessage:       p.ErrorMessage,
		GrafanaExplorePath: p.GrafanaExplorePath,
		LinkToAPI:          p.LinkToAPI,
		Query:              p.Query,
		Status:             p.Status,
		Title:              p.Title,
		Type:               p.Type,
	}
}

// PrometheusRangeQuery - Prometheus instant query enrichment object.
type PrometheusRangeQuery struct {
	// REQUIRED; An array of the azure monitor workspace resource ids.
	Datasources []*string

	// REQUIRED; The enrichment description.
	Description *string

	// REQUIRED; The end evaluation date and time in ISO8601 format.
	End *time.Time

	// REQUIRED; Partial link to the Grafana explore API.
	GrafanaExplorePath *string

	// REQUIRED; Link to Prometheus query API (Url format).
	LinkToAPI *string

	// REQUIRED; The Prometheus expression query.
	Query *string

	// REQUIRED; The start evaluation date and time in ISO8601 format.
	Start *time.Time

	// REQUIRED; The status of the evaluation of the enrichment.
	Status *Status

	// REQUIRED; Query resolution step width in ISO8601 format.
	Step *string

	// REQUIRED; The enrichment title.
	Title *string

	// REQUIRED; The enrichment type.
	Type *Type

	// The error message. Will be present only if the status is 'Failed'.
	ErrorMessage *string
}

// GetAlertEnrichmentItem implements the AlertEnrichmentItemClassification interface for type PrometheusRangeQuery.
func (p *PrometheusRangeQuery) GetAlertEnrichmentItem() *AlertEnrichmentItem {
	return &AlertEnrichmentItem{
		Description:  p.Description,
		ErrorMessage: p.ErrorMessage,
		Status:       p.Status,
		Title:        p.Title,
		Type:         p.Type,
	}
}

// GetPrometheusEnrichmentItem implements the PrometheusEnrichmentItemClassification interface for type PrometheusRangeQuery.
func (p *PrometheusRangeQuery) GetPrometheusEnrichmentItem() *PrometheusEnrichmentItem {
	return &PrometheusEnrichmentItem{
		Datasources:        p.Datasources,
		Description:        p.Description,
		ErrorMessage:       p.ErrorMessage,
		GrafanaExplorePath: p.GrafanaExplorePath,
		LinkToAPI:          p.LinkToAPI,
		Query:              p.Query,
		Status:             p.Status,
		Title:              p.Title,
		Type:               p.Type,
	}
}

// PrometheusRule - An Azure Prometheus alerting or recording rule.
type PrometheusRule struct {
	// REQUIRED; The PromQL expression to evaluate. https://prometheus.io/docs/prometheus/latest/querying/basics/. Evaluated periodically
	// as given by 'interval', and the result recorded as a new set of time series
	// with the metric name as given by 'record'.
	Expression *string

	// Actions that are performed when the alert rule becomes active, and when an alert condition is resolved.
	Actions []*PrometheusRuleGroupAction

	// Alert rule name.
	Alert *string

	// The annotations clause specifies a set of informational labels that can be used to store longer additional information
	// such as alert descriptions or runbook links. The annotation values can be
	// templated.
	Annotations map[string]*string

	// Enable/disable rule.
	Enabled *bool

	// The amount of time alert must be active before firing.
	For *string

	// Labels to add or overwrite before storing the result.
	Labels map[string]*string

	// Recorded metrics name.
	Record *string

	// Defines the configuration for resolving fired alerts. Only relevant for alerts.
	ResolveConfiguration *PrometheusRuleResolveConfiguration

	// The severity of the alerts fired by the rule. Must be between 0 and 4.
	Severity *int32
}

// PrometheusRuleGroupAction - An alert action. Only relevant for alerts.
type PrometheusRuleGroupAction struct {
	// The resource id of the action group to use.
	ActionGroupID *string

	// The properties of an action group object.
	ActionProperties map[string]*string
}

// PrometheusRuleGroupProperties - An Azure Prometheus rule group.
type PrometheusRuleGroupProperties struct {
	// REQUIRED; Defines the rules in the Prometheus rule group.
	Rules []*PrometheusRule

	// REQUIRED; Target Azure Monitor workspaces resource ids. This api-version is currently limited to creating with one scope.
	// This may change in future.
	Scopes []*string

	// Apply rule to data from a specific cluster.
	ClusterName *string

	// Rule group description.
	Description *string

	// Enable/disable rule group.
	Enabled *bool

	// The interval in which to run the Prometheus rule group represented in ISO 8601 duration format. Should be between 1 and
	// 15 minutes
	Interval *string
}

// PrometheusRuleGroupResource - The Prometheus rule group resource.
type PrometheusRuleGroupResource struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// REQUIRED; The Prometheus rule group properties of the resource.
	Properties *PrometheusRuleGroupProperties

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// PrometheusRuleGroupResourceCollection - Represents a collection of alert rule resources.
type PrometheusRuleGroupResourceCollection struct {
	// the values for the alert rule resources.
	Value []*PrometheusRuleGroupResource
}

// PrometheusRuleGroupResourcePatchParameters - The Prometheus rule group resource for patch operations.
type PrometheusRuleGroupResourcePatchParameters struct {
	Properties *PrometheusRuleGroupResourcePatchParametersProperties

	// Resource tags
	Tags map[string]*string
}

type PrometheusRuleGroupResourcePatchParametersProperties struct {
	// the flag that indicates whether the Prometheus rule group is enabled.
	Enabled *bool
}

// PrometheusRuleResolveConfiguration - Specifies the Prometheus alert rule configuration.
type PrometheusRuleResolveConfiguration struct {
	// Enable alert auto-resolution.
	AutoResolved *bool

	// Alert auto-resolution timeout.
	TimeToResolve *string
}

// Recurrence object.
type Recurrence struct {
	// REQUIRED; Specifies when the recurrence should be applied.
	RecurrenceType *RecurrenceType

	// End time for recurrence.
	EndTime *string

	// Start time for recurrence.
	StartTime *string
}

// GetRecurrence implements the RecurrenceClassification interface for type Recurrence.
func (r *Recurrence) GetRecurrence() *Recurrence { return r }

// RemoveAllActionGroups - Indicates if all action groups should be removed.
type RemoveAllActionGroups struct {
	// REQUIRED; Action that should be applied.
	ActionType *ActionType
}

// GetAction implements the ActionClassification interface for type RemoveAllActionGroups.
func (r *RemoveAllActionGroups) GetAction() *Action {
	return &Action{
		ActionType: r.ActionType,
	}
}

// RuleArmTemplate - A complete ARM template to deploy the alert rules.
type RuleArmTemplate struct {
	// REQUIRED; A 4 number format for the version number of this template file. For example, 1.0.0.0
	ContentVersion *string

	// REQUIRED; Input parameter definitions
	Parameters any

	// REQUIRED; Alert rule resource definitions
	Resources []any

	// REQUIRED; JSON schema reference
	Schema *string

	// REQUIRED; Variable definitions
	Variables any
}

// Schedule - Scheduling configuration for a given alert processing rule.
type Schedule struct {
	// Scheduling effective from time. Date-Time in ISO-8601 format without timezone suffix.
	EffectiveFrom *string

	// Scheduling effective until time. Date-Time in ISO-8601 format without timezone suffix.
	EffectiveUntil *string

	// List of recurrences.
	Recurrences []RecurrenceClassification

	// Scheduling time zone.
	TimeZone *string
}

// SmartGroup - Set of related alerts grouped together smartly by AMS.
type SmartGroup struct {
	// Properties of smart group.
	Properties *SmartGroupProperties

	// READ-ONLY; Azure resource Id
	ID *string

	// READ-ONLY; Azure resource name
	Name *string

	// READ-ONLY; Azure resource type
	Type *string
}

// SmartGroupAggregatedProperty - Aggregated property of each type
type SmartGroupAggregatedProperty struct {
	// Total number of items of type.
	Count *int64

	// Name of the type.
	Name *string
}

// SmartGroupModification - Alert Modification details
type SmartGroupModification struct {
	// Properties of the smartGroup modification item.
	Properties *SmartGroupModificationProperties

	// READ-ONLY; Azure resource Id
	ID *string

	// READ-ONLY; Azure resource name
	Name *string

	// READ-ONLY; Azure resource type
	Type *string
}

// SmartGroupModificationItem - smartGroup modification item.
type SmartGroupModificationItem struct {
	// Modification comments
	Comments *string

	// Description of the modification
	Description *string

	// Reason for the modification
	ModificationEvent *SmartGroupModificationEvent

	// Modified date and time
	ModifiedAt *string

	// Modified user details (Principal client name)
	ModifiedBy *string

	// New value
	NewValue *string

	// Old value
	OldValue *string
}

// SmartGroupModificationProperties - Properties of the smartGroup modification item.
type SmartGroupModificationProperties struct {
	// Modification details
	Modifications []*SmartGroupModificationItem

	// URL to fetch the next set of results.
	NextLink *string

	// READ-ONLY; Unique Id of the smartGroup for which the history is being retrieved
	SmartGroupID *string
}

// SmartGroupProperties - Properties of smart group.
type SmartGroupProperties struct {
	// Summary of alertSeverities in the smart group
	AlertSeverities []*SmartGroupAggregatedProperty

	// Summary of alertStates in the smart group
	AlertStates []*SmartGroupAggregatedProperty

	// Total number of alerts in smart group
	AlertsCount *int64

	// Summary of monitorConditions in the smart group
	MonitorConditions []*SmartGroupAggregatedProperty

	// Summary of monitorServices in the smart group
	MonitorServices []*SmartGroupAggregatedProperty

	// The URI to fetch the next page of alerts. Call ListNext() with this URI to fetch the next page alerts.
	NextLink *string

	// Summary of target resource groups in the smart group
	ResourceGroups []*SmartGroupAggregatedProperty

	// Summary of target resource types in the smart group
	ResourceTypes []*SmartGroupAggregatedProperty

	// Summary of target resources in the smart group
	Resources []*SmartGroupAggregatedProperty

	// READ-ONLY; Last updated time of smart group. Date-Time in ISO-8601 format.
	LastModifiedDateTime *time.Time

	// READ-ONLY; Last modified by user name.
	LastModifiedUserName *string

	// READ-ONLY; Severity of smart group is the highest(Sev0 >… > Sev4) severity of all the alerts in the group.
	Severity *Severity

	// READ-ONLY; Smart group state
	SmartGroupState *State

	// READ-ONLY; Creation time of smart group. Date-Time in ISO-8601 format.
	StartDateTime *time.Time
}

// SmartGroupsList - List the alerts.
type SmartGroupsList struct {
	// URL to fetch the next set of alerts.
	NextLink *string

	// List of alerts
	Value []*SmartGroup
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time

	// The identity that created the resource.
	CreatedBy *string

	// The type of identity that created the resource.
	CreatedByType *CreatedByType

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time

	// The identity that last modified the resource.
	LastModifiedBy *string

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType
}

// WeeklyRecurrence - Weekly recurrence object.
type WeeklyRecurrence struct {
	// REQUIRED; Specifies the values for weekly recurrence pattern.
	DaysOfWeek []*DaysOfWeek

	// REQUIRED; Specifies when the recurrence should be applied.
	RecurrenceType *RecurrenceType

	// End time for recurrence.
	EndTime *string

	// Start time for recurrence.
	StartTime *string
}

// GetRecurrence implements the RecurrenceClassification interface for type WeeklyRecurrence.
func (w *WeeklyRecurrence) GetRecurrence() *Recurrence {
	return &Recurrence{
		EndTime:        w.EndTime,
		RecurrenceType: w.RecurrenceType,
		StartTime:      w.StartTime,
	}
}
