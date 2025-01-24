//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpolicyinsights

import "time"

// Attestation - An attestation resource.
type Attestation struct {
	// REQUIRED; Properties for the attestation.
	Properties *AttestationProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// AttestationEvidence - A piece of evidence supporting the compliance state set in the attestation.
type AttestationEvidence struct {
	// The description for this piece of evidence.
	Description *string

	// The URI location of the evidence.
	SourceURI *string
}

// AttestationListResult - List of attestations.
type AttestationListResult struct {
	// READ-ONLY; The URL to get the next set of results.
	NextLink *string

	// READ-ONLY; Array of attestation definitions.
	Value []*Attestation
}

// AttestationProperties - The properties of an attestation resource.
type AttestationProperties struct {
	// REQUIRED; The resource ID of the policy assignment that the attestation is setting the state for.
	PolicyAssignmentID *string

	// The time the evidence was assessed
	AssessmentDate *time.Time

	// Comments describing why this attestation was created.
	Comments *string

	// The compliance state that should be set on the resource.
	ComplianceState *ComplianceState

	// The evidence supporting the compliance state set in this attestation.
	Evidence []*AttestationEvidence

	// The time the compliance state should expire.
	ExpiresOn *time.Time

	// Additional metadata for this attestation
	Metadata any

	// The person responsible for setting the state of the resource. This value is typically an Azure Active Directory object
	// ID.
	Owner *string

	// The policy definition reference ID from a policy set definition that the attestation is setting the state for. If the policy
	// assignment assigns a policy set definition the attestation can choose a
	// definition within the set definition with this property or omit this and set the state for the entire set definition.
	PolicyDefinitionReferenceID *string

	// READ-ONLY; The time the compliance state was last changed in this attestation.
	LastComplianceStateChangeAt *time.Time

	// READ-ONLY; The status of the attestation.
	ProvisioningState *string
}

// CheckManagementGroupRestrictionsRequest - The check policy restrictions parameters describing the resource that is being
// evaluated.
type CheckManagementGroupRestrictionsRequest struct {
	// The list of fields and values that should be evaluated for potential restrictions.
	PendingFields []*PendingField

	// The information about the resource that will be evaluated.
	ResourceDetails *CheckRestrictionsResourceDetails
}

// CheckRestrictionEvaluationDetails - Policy evaluation details.
type CheckRestrictionEvaluationDetails struct {
	// Details of the evaluated expressions.
	EvaluatedExpressions []*ExpressionEvaluationDetails

	// Evaluation details of IfNotExists effect.
	IfNotExistsDetails *IfNotExistsEvaluationDetails

	// READ-ONLY; The reason for the evaluation result.
	Reason *string
}

// CheckRestrictionsRequest - The check policy restrictions parameters describing the resource that is being evaluated.
type CheckRestrictionsRequest struct {
	// REQUIRED; The information about the resource that will be evaluated.
	ResourceDetails *CheckRestrictionsResourceDetails

	// Whether to include policies with the 'audit' effect in the results. Defaults to false.
	IncludeAuditEffect *bool

	// The list of fields and values that should be evaluated for potential restrictions.
	PendingFields []*PendingField
}

// CheckRestrictionsResourceDetails - The information about the resource that will be evaluated.
type CheckRestrictionsResourceDetails struct {
	// REQUIRED; The resource content. This should include whatever properties are already known and can be a partial set of all
	// resource properties.
	ResourceContent any

	// The api-version of the resource content.
	APIVersion *string

	// The scope where the resource is being created. For example, if the resource is a child resource this would be the parent
	// resource's resource ID.
	Scope *string
}

// CheckRestrictionsResult - The result of a check policy restrictions evaluation on a resource.
type CheckRestrictionsResult struct {
	// READ-ONLY; Evaluation results for the provided partial resource content.
	ContentEvaluationResult *CheckRestrictionsResultContentEvaluationResult

	// READ-ONLY; The restrictions that will be placed on various fields in the resource by policy.
	FieldRestrictions []*FieldRestrictions
}

// CheckRestrictionsResultContentEvaluationResult - Evaluation results for the provided partial resource content.
type CheckRestrictionsResultContentEvaluationResult struct {
	// Policy evaluation results against the given resource content. This will indicate if the partial content that was provided
	// will be denied as-is.
	PolicyEvaluations []*PolicyEvaluationResult
}

// ComplianceDetail - The compliance state rollup.
type ComplianceDetail struct {
	// The compliance state.
	ComplianceState *string

	// Summarized count value for this compliance state.
	Count *int32
}

// ComponentEventDetails - Component event details.
type ComponentEventDetails struct {
	// OPTIONAL; Contains additional key/value pairs not defined in the schema.
	AdditionalProperties map[string]any

	// Component Id.
	ID *string

	// Component name.
	Name *string

	// Policy definition action, i.e. effect.
	PolicyDefinitionAction *string

	// Principal object ID for the user who initiated the resource component operation that triggered the policy event.
	PrincipalOid *string

	// Tenant ID for the policy event record.
	TenantID *string

	// Timestamp for component policy event record.
	Timestamp *time.Time

	// Component type.
	Type *string
}

// ComponentExpressionEvaluationDetails - Evaluation details of policy language expressions.
type ComponentExpressionEvaluationDetails struct {
	// Evaluation result.
	Result *string

	// READ-ONLY; Expression evaluated.
	Expression *string

	// READ-ONLY; The kind of expression that was evaluated.
	ExpressionKind *string

	// READ-ONLY; Value of the expression.
	ExpressionValue any

	// READ-ONLY; Operator to compare the expression value and the target value.
	Operator *string

	// READ-ONLY; Property path if the expression is a field or an alias.
	Path *string

	// READ-ONLY; Target value to be compared with the expression value.
	TargetValue any
}

// ComponentPolicyEvaluationDetails - Policy evaluation details.
type ComponentPolicyEvaluationDetails struct {
	// Additional textual reason for the evaluation outcome.
	Reason *string

	// READ-ONLY; Details of the evaluated expressions.
	EvaluatedExpressions []*ComponentExpressionEvaluationDetails
}

// ComponentPolicyState - Component Policy State record.
type ComponentPolicyState struct {
	// OPTIONAL; Contains additional key/value pairs not defined in the schema.
	AdditionalProperties map[string]any

	// Policy evaluation details. This is only included in the response if the request contains $expand=PolicyEvaluationDetails.
	PolicyEvaluationDetails *ComponentPolicyEvaluationDetails

	// READ-ONLY; Compliance state of the resource.
	ComplianceState *string

	// READ-ONLY; Component Id.
	ComponentID *string

	// READ-ONLY; Component name.
	ComponentName *string

	// READ-ONLY; Component type.
	ComponentType *string

	// READ-ONLY; OData context string; used by OData clients to resolve type information based on metadata.
	ODataContext *string

	// READ-ONLY; OData entity ID; always set to null since component policy state records do not have an entity ID.
	ODataID *string

	// READ-ONLY; Policy assignment ID.
	PolicyAssignmentID *string

	// READ-ONLY; Policy assignment name.
	PolicyAssignmentName *string

	// READ-ONLY; Policy assignment owner.
	PolicyAssignmentOwner *string

	// READ-ONLY; Policy assignment parameters.
	PolicyAssignmentParameters *string

	// READ-ONLY; Policy assignment scope.
	PolicyAssignmentScope *string

	// READ-ONLY; Evaluated policy assignment version.
	PolicyAssignmentVersion *string

	// READ-ONLY; Policy definition action, i.e. effect.
	PolicyDefinitionAction *string

	// READ-ONLY; Policy definition category.
	PolicyDefinitionCategory *string

	// READ-ONLY; Policy definition group names.
	PolicyDefinitionGroupNames []*string

	// READ-ONLY; Policy definition ID.
	PolicyDefinitionID *string

	// READ-ONLY; Policy definition name.
	PolicyDefinitionName *string

	// READ-ONLY; Reference ID for the policy definition inside the policy set, if the policy assignment is for a policy set.
	PolicyDefinitionReferenceID *string

	// READ-ONLY; Evaluated policy definition version.
	PolicyDefinitionVersion *string

	// READ-ONLY; Policy set definition category, if the policy assignment is for a policy set.
	PolicySetDefinitionCategory *string

	// READ-ONLY; Policy set definition ID, if the policy assignment is for a policy set.
	PolicySetDefinitionID *string

	// READ-ONLY; Policy set definition name, if the policy assignment is for a policy set.
	PolicySetDefinitionName *string

	// READ-ONLY; Policy set definition owner, if the policy assignment is for a policy set.
	PolicySetDefinitionOwner *string

	// READ-ONLY; Policy set definition parameters, if the policy assignment is for a policy set.
	PolicySetDefinitionParameters *string

	// READ-ONLY; Evaluated policy set definition version.
	PolicySetDefinitionVersion *string

	// READ-ONLY; Resource group name.
	ResourceGroup *string

	// READ-ONLY; Resource ID.
	ResourceID *string

	// READ-ONLY; Resource location.
	ResourceLocation *string

	// READ-ONLY; Resource type.
	ResourceType *string

	// READ-ONLY; Subscription ID.
	SubscriptionID *string

	// READ-ONLY; Timestamp for the component policy state record.
	Timestamp *time.Time
}

// ComponentPolicyStatesQueryResults - Query results.
type ComponentPolicyStatesQueryResults struct {
	// OData context string; used by OData clients to resolve type information based on metadata.
	ODataContext *string

	// OData entity count; represents the number of policy state records returned.
	ODataCount *int32

	// Query results.
	Value []*ComponentPolicyState
}

// ComponentStateDetails - Component state details.
type ComponentStateDetails struct {
	// OPTIONAL; Contains additional key/value pairs not defined in the schema.
	AdditionalProperties map[string]any

	// Component compliance state.
	ComplianceState *string

	// Component Id.
	ID *string

	// Component name.
	Name *string

	// Component compliance evaluation timestamp.
	Timestamp *time.Time

	// Component type.
	Type *string
}

// ErrorDefinition - Error definition.
type ErrorDefinition struct {
	// READ-ONLY; Additional scenario specific error details.
	AdditionalInfo []*TypedErrorInfo

	// READ-ONLY; Service specific error code which serves as the substatus for the HTTP error code.
	Code *string

	// READ-ONLY; Internal error details.
	Details []*ErrorDefinition

	// READ-ONLY; Description of the error.
	Message *string

	// READ-ONLY; The target of the error.
	Target *string
}

// ErrorDefinitionAutoGenerated - Error definition.
type ErrorDefinitionAutoGenerated struct {
	// READ-ONLY; Additional scenario specific error details.
	AdditionalInfo []*TypedErrorInfo

	// READ-ONLY; Service specific error code which serves as the substatus for the HTTP error code.
	Code *string

	// READ-ONLY; Internal error details.
	Details []*ErrorDefinitionAutoGenerated

	// READ-ONLY; Description of the error.
	Message *string

	// READ-ONLY; The target of the error.
	Target *string
}

// ErrorDefinitionAutoGenerated2 - Error definition.
type ErrorDefinitionAutoGenerated2 struct {
	// READ-ONLY; Additional scenario specific error details.
	AdditionalInfo []*TypedErrorInfo

	// READ-ONLY; Service specific error code which serves as the substatus for the HTTP error code.
	Code *string

	// READ-ONLY; Internal error details.
	Details []*ErrorDefinitionAutoGenerated2

	// READ-ONLY; Description of the error.
	Message *string

	// READ-ONLY; The target of the error.
	Target *string
}

// ErrorResponse - Error response.
type ErrorResponse struct {
	// The error details.
	Error *ErrorDefinition
}

// ErrorResponseAutoGenerated - Error response.
type ErrorResponseAutoGenerated struct {
	// The error details.
	Error *ErrorDefinitionAutoGenerated
}

// ErrorResponseAutoGenerated2 - Error response.
type ErrorResponseAutoGenerated2 struct {
	// The error details.
	Error *ErrorDefinitionAutoGenerated2
}

// ExpressionEvaluationDetails - Evaluation details of policy language expressions.
type ExpressionEvaluationDetails struct {
	// Expression evaluated.
	Expression *string

	// Value of the expression.
	ExpressionValue any

	// Operator to compare the expression value and the target value.
	Operator *string

	// Property path if the expression is a field or an alias.
	Path *string

	// Evaluation result.
	Result *string

	// Target value to be compared with the expression value.
	TargetValue any

	// READ-ONLY; The kind of expression that was evaluated.
	ExpressionKind *string
}

// FieldRestriction - The restrictions on a field imposed by a specific policy.
type FieldRestriction struct {
	// READ-ONLY; The value that policy will set for the field if the user does not provide a value.
	DefaultValue *string

	// READ-ONLY; The details of the policy that is causing the field restriction.
	Policy *PolicyReference

	// READ-ONLY; The effect of the policy that is causing the field restriction. http://aka.ms/policyeffects
	PolicyEffect *string

	// READ-ONLY; The reason for the restriction.
	Reason *string

	// READ-ONLY; The type of restriction that is imposed on the field.
	Result *FieldRestrictionResult

	// READ-ONLY; The values that policy either requires or denies for the field.
	Values []*string
}

// FieldRestrictions - The restrictions that will be placed on a field in the resource by policy.
type FieldRestrictions struct {
	// The restrictions placed on that field by policy.
	Restrictions []*FieldRestriction

	// READ-ONLY; The name of the field. This can be a top-level property like 'name' or 'type' or an Azure Policy field alias.
	Field *string
}

// IfNotExistsEvaluationDetails - Evaluation details of IfNotExists effect.
type IfNotExistsEvaluationDetails struct {
	// ID of the last evaluated resource for IfNotExists effect.
	ResourceID *string

	// Total number of resources to which the existence condition is applicable.
	TotalResources *int32
}

// Operation definition.
type Operation struct {
	// Display metadata associated with the operation.
	Display *OperationDisplay

	// Indicates whether the operation is a data action
	IsDataAction *bool

	// Operation name.
	Name *string
}

// OperationDisplay - Display metadata associated with the operation.
type OperationDisplay struct {
	// Operation description.
	Description *string

	// Operation name.
	Operation *string

	// Resource provider name.
	Provider *string

	// Resource name on which the operation is performed.
	Resource *string
}

// OperationsListResults - List of available operations.
type OperationsListResults struct {
	// OData entity count; represents the number of operations returned.
	ODataCount *int32

	// List of available operations.
	Value []*Operation
}

// PendingField - A field that should be evaluated against Azure Policy to determine restrictions.
type PendingField struct {
	// REQUIRED; The name of the field. This can be a top-level property like 'name' or 'type' or an Azure Policy field alias.
	Field *string

	// The list of potential values for the field that should be evaluated against Azure Policy.
	Values []*string
}

// PolicyAssignmentSummary - Policy assignment summary.
type PolicyAssignmentSummary struct {
	// Policy assignment ID.
	PolicyAssignmentID *string

	// Policy definitions summary.
	PolicyDefinitions []*PolicyDefinitionSummary

	// Policy definition group summary.
	PolicyGroups []*PolicyGroupSummary

	// Policy set definition ID, if the policy assignment is for a policy set.
	PolicySetDefinitionID *string

	// Compliance summary for the policy assignment.
	Results *SummaryResults
}

// PolicyDefinitionSummary - Policy definition summary.
type PolicyDefinitionSummary struct {
	// Policy effect, i.e. policy definition action.
	Effect *string

	// Policy definition group names.
	PolicyDefinitionGroupNames []*string

	// Policy definition ID.
	PolicyDefinitionID *string

	// Policy definition reference ID.
	PolicyDefinitionReferenceID *string

	// Compliance summary for the policy definition.
	Results *SummaryResults
}

// PolicyDetails - The policy details.
type PolicyDetails struct {
	// READ-ONLY; The display name of the policy assignment.
	PolicyAssignmentDisplayName *string

	// READ-ONLY; The ID of the policy assignment.
	PolicyAssignmentID *string

	// READ-ONLY; The scope of the policy assignment.
	PolicyAssignmentScope *string

	// READ-ONLY; The ID of the policy definition.
	PolicyDefinitionID *string

	// READ-ONLY; The policy definition reference ID within the policy set definition.
	PolicyDefinitionReferenceID *string

	// READ-ONLY; The ID of the policy set definition.
	PolicySetDefinitionID *string
}

// PolicyEffectDetails - The details of the effect that was applied to the resource.
type PolicyEffectDetails struct {
	// READ-ONLY; The effect that was applied to the resource. http://aka.ms/policyeffects
	PolicyEffect *string
}

// PolicyEvaluationDetails - Policy evaluation details.
type PolicyEvaluationDetails struct {
	// Details of the evaluated expressions.
	EvaluatedExpressions []*ExpressionEvaluationDetails

	// Evaluation details of IfNotExists effect.
	IfNotExistsDetails *IfNotExistsEvaluationDetails
}

// PolicyEvaluationResult - The result of a non-compliant policy evaluation against the given resource content.
type PolicyEvaluationResult struct {
	// READ-ONLY; The details of the effect that was applied to the resource.
	EffectDetails *PolicyEffectDetails

	// READ-ONLY; The detailed results of the policy expressions and values that were evaluated.
	EvaluationDetails *CheckRestrictionEvaluationDetails

	// READ-ONLY; The result of the policy evaluation against the resource. This will typically be 'NonCompliant' but may contain
	// other values if errors were encountered.
	EvaluationResult *string

	// READ-ONLY; The details of the policy that was evaluated.
	PolicyInfo *PolicyReference
}

// PolicyEvent - Policy event record.
type PolicyEvent struct {
	// OPTIONAL; Contains additional key/value pairs not defined in the schema.
	AdditionalProperties map[string]any

	// Compliance state of the resource.
	ComplianceState *string

	// Components events records populated only when URL contains $expand=components clause.
	Components []*ComponentEventDetails

	// Effective parameters for the policy assignment.
	EffectiveParameters *string

	// Flag which states whether the resource is compliant against the policy assignment it was evaluated against.
	IsCompliant *bool

	// Comma separated list of management group IDs, which represent the hierarchy of the management groups the resource is under.
	ManagementGroupIDs *string

	// OData context string; used by OData clients to resolve type information based on metadata.
	ODataContext *string

	// OData entity ID; always set to null since policy event records do not have an entity ID.
	ODataID *string

	// Policy assignment ID.
	PolicyAssignmentID *string

	// Policy assignment name.
	PolicyAssignmentName *string

	// Policy assignment owner.
	PolicyAssignmentOwner *string

	// Policy assignment parameters.
	PolicyAssignmentParameters *string

	// Policy assignment scope.
	PolicyAssignmentScope *string

	// Policy definition action, i.e. effect.
	PolicyDefinitionAction *string

	// Policy definition category.
	PolicyDefinitionCategory *string

	// Policy definition ID.
	PolicyDefinitionID *string

	// Policy definition name.
	PolicyDefinitionName *string

	// Reference ID for the policy definition inside the policy set, if the policy assignment is for a policy set.
	PolicyDefinitionReferenceID *string

	// Policy set definition category, if the policy assignment is for a policy set.
	PolicySetDefinitionCategory *string

	// Policy set definition ID, if the policy assignment is for a policy set.
	PolicySetDefinitionID *string

	// Policy set definition name, if the policy assignment is for a policy set.
	PolicySetDefinitionName *string

	// Policy set definition owner, if the policy assignment is for a policy set.
	PolicySetDefinitionOwner *string

	// Policy set definition parameters, if the policy assignment is for a policy set.
	PolicySetDefinitionParameters *string

	// Principal object ID for the user who initiated the resource operation that triggered the policy event.
	PrincipalOid *string

	// Resource group name.
	ResourceGroup *string

	// Resource ID.
	ResourceID *string

	// Resource location.
	ResourceLocation *string

	// List of resource tags.
	ResourceTags *string

	// Resource type.
	ResourceType *string

	// Subscription ID.
	SubscriptionID *string

	// Tenant ID for the policy event record.
	TenantID *string

	// Timestamp for the policy event record.
	Timestamp *time.Time
}

// PolicyEventsQueryResults - Query results.
type PolicyEventsQueryResults struct {
	// OData context string; used by OData clients to resolve type information based on metadata.
	ODataContext *string

	// OData entity count; represents the number of policy event records returned.
	ODataCount *int32

	// Odata next link; URL to get the next set of results.
	ODataNextLink *string

	// Query results.
	Value []*PolicyEvent
}

// PolicyGroupSummary - Policy definition group summary.
type PolicyGroupSummary struct {
	// Policy group name.
	PolicyGroupName *string

	// Compliance summary for the policy definition group.
	Results *SummaryResults
}

// PolicyMetadata - Policy metadata resource definition.
type PolicyMetadata struct {
	// Properties of the policy metadata.
	Properties *PolicyMetadataProperties

	// READ-ONLY; The ID of the policy metadata.
	ID *string

	// READ-ONLY; The name of the policy metadata.
	Name *string

	// READ-ONLY; The type of the policy metadata.
	Type *string
}

// PolicyMetadataCollection - Collection of policy metadata resources.
type PolicyMetadataCollection struct {
	// READ-ONLY; The URL to get the next set of results.
	NextLink *string

	// READ-ONLY; Array of policy metadata definitions.
	Value []*SlimPolicyMetadata
}

// PolicyMetadataProperties - The properties of the policy metadata.
type PolicyMetadataProperties struct {
	// READ-ONLY; Url for getting additional content about the resource metadata.
	AdditionalContentURL *string

	// READ-ONLY; The category of the policy metadata.
	Category *string

	// READ-ONLY; The description of the policy metadata.
	Description *string

	// READ-ONLY; Additional metadata.
	Metadata any

	// READ-ONLY; The policy metadata identifier.
	MetadataID *string

	// READ-ONLY; The owner of the policy metadata.
	Owner *string

	// READ-ONLY; The requirements of the policy metadata.
	Requirements *string

	// READ-ONLY; The title of the policy metadata.
	Title *string
}

// PolicyMetadataSlimProperties - The properties of the policy metadata, excluding properties containing large strings
type PolicyMetadataSlimProperties struct {
	// READ-ONLY; Url for getting additional content about the resource metadata.
	AdditionalContentURL *string

	// READ-ONLY; The category of the policy metadata.
	Category *string

	// READ-ONLY; Additional metadata.
	Metadata any

	// READ-ONLY; The policy metadata identifier.
	MetadataID *string

	// READ-ONLY; The owner of the policy metadata.
	Owner *string

	// READ-ONLY; The title of the policy metadata.
	Title *string
}

// PolicyReference - Resource identifiers for a policy.
type PolicyReference struct {
	// READ-ONLY; The resource identifier of the policy assignment.
	PolicyAssignmentID *string

	// READ-ONLY; The resource identifier of the policy definition.
	PolicyDefinitionID *string

	// READ-ONLY; The reference identifier of a specific policy definition within a policy set definition.
	PolicyDefinitionReferenceID *string

	// READ-ONLY; The resource identifier of the policy set definition.
	PolicySetDefinitionID *string
}

// PolicyState - Policy state record.
type PolicyState struct {
	// OPTIONAL; Contains additional key/value pairs not defined in the schema.
	AdditionalProperties map[string]any

	// Compliance state of the resource.
	ComplianceState *string

	// Components state compliance records populated only when URL contains $expand=components clause.
	Components []*ComponentStateDetails

	// Effective parameters for the policy assignment.
	EffectiveParameters *string

	// Flag which states whether the resource is compliant against the policy assignment it was evaluated against. This property
	// is deprecated; please use ComplianceState instead.
	IsCompliant *bool

	// Comma separated list of management group IDs, which represent the hierarchy of the management groups the resource is under.
	ManagementGroupIDs *string

	// OData context string; used by OData clients to resolve type information based on metadata.
	ODataContext *string

	// OData entity ID; always set to null since policy state records do not have an entity ID.
	ODataID *string

	// Policy assignment ID.
	PolicyAssignmentID *string

	// Policy assignment name.
	PolicyAssignmentName *string

	// Policy assignment owner.
	PolicyAssignmentOwner *string

	// Policy assignment parameters.
	PolicyAssignmentParameters *string

	// Policy assignment scope.
	PolicyAssignmentScope *string

	// Policy definition action, i.e. effect.
	PolicyDefinitionAction *string

	// Policy definition category.
	PolicyDefinitionCategory *string

	// Policy definition group names.
	PolicyDefinitionGroupNames []*string

	// Policy definition ID.
	PolicyDefinitionID *string

	// Policy definition name.
	PolicyDefinitionName *string

	// Reference ID for the policy definition inside the policy set, if the policy assignment is for a policy set.
	PolicyDefinitionReferenceID *string

	// Policy evaluation details.
	PolicyEvaluationDetails *PolicyEvaluationDetails

	// Policy set definition category, if the policy assignment is for a policy set.
	PolicySetDefinitionCategory *string

	// Policy set definition ID, if the policy assignment is for a policy set.
	PolicySetDefinitionID *string

	// Policy set definition name, if the policy assignment is for a policy set.
	PolicySetDefinitionName *string

	// Policy set definition owner, if the policy assignment is for a policy set.
	PolicySetDefinitionOwner *string

	// Policy set definition parameters, if the policy assignment is for a policy set.
	PolicySetDefinitionParameters *string

	// Resource group name.
	ResourceGroup *string

	// Resource ID.
	ResourceID *string

	// Resource location.
	ResourceLocation *string

	// List of resource tags.
	ResourceTags *string

	// Resource type.
	ResourceType *string

	// Subscription ID.
	SubscriptionID *string

	// Timestamp for the policy state record.
	Timestamp *time.Time

	// READ-ONLY; Evaluated policy assignment version.
	PolicyAssignmentVersion *string

	// READ-ONLY; Evaluated policy definition version.
	PolicyDefinitionVersion *string

	// READ-ONLY; Evaluated policy set definition version.
	PolicySetDefinitionVersion *string
}

// PolicyStatesQueryResults - Query results.
type PolicyStatesQueryResults struct {
	// OData context string; used by OData clients to resolve type information based on metadata.
	ODataContext *string

	// OData entity count; represents the number of policy state records returned.
	ODataCount *int32

	// Odata next link; URL to get the next set of results.
	ODataNextLink *string

	// Query results.
	Value []*PolicyState
}

// PolicyTrackedResource - Policy tracked resource record.
type PolicyTrackedResource struct {
	// READ-ONLY; The details of the policy triggered deployment that created the tracked resource.
	CreatedBy *TrackedResourceModificationDetails

	// READ-ONLY; The details of the policy triggered deployment that modified the tracked resource.
	LastModifiedBy *TrackedResourceModificationDetails

	// READ-ONLY; Timestamp of the last update to the tracked resource.
	LastUpdateUTC *time.Time

	// READ-ONLY; The details of the policy that require the tracked resource.
	PolicyDetails *PolicyDetails

	// READ-ONLY; The ID of the policy tracked resource.
	TrackedResourceID *string
}

// PolicyTrackedResourcesQueryResults - Query results.
type PolicyTrackedResourcesQueryResults struct {
	// READ-ONLY; The URL to get the next set of results.
	NextLink *string

	// READ-ONLY; Query results.
	Value []*PolicyTrackedResource
}

// QueryFailure - Error response.
type QueryFailure struct {
	// Error definition.
	Error *QueryFailureError
}

// QueryFailureError - Error definition.
type QueryFailureError struct {
	// READ-ONLY; Service specific error code which serves as the substatus for the HTTP error code.
	Code *string

	// READ-ONLY; Description of the error.
	Message *string
}

// Remediation - The remediation definition.
type Remediation struct {
	// Properties for the remediation.
	Properties *RemediationProperties

	// READ-ONLY; The ID of the remediation.
	ID *string

	// READ-ONLY; The name of the remediation.
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the remediation.
	Type *string
}

// RemediationDeployment - Details of a single deployment created by the remediation.
type RemediationDeployment struct {
	// READ-ONLY; The time at which the remediation was created.
	CreatedOn *time.Time

	// READ-ONLY; Resource ID of the template deployment that will remediate the resource.
	DeploymentID *string

	// READ-ONLY; Error encountered while remediated the resource.
	Error *ErrorDefinition

	// READ-ONLY; The time at which the remediation deployment was last updated.
	LastUpdatedOn *time.Time

	// READ-ONLY; Resource ID of the resource that is being remediated by the deployment.
	RemediatedResourceID *string

	// READ-ONLY; Location of the resource that is being remediated.
	ResourceLocation *string

	// READ-ONLY; Status of the remediation deployment.
	Status *string
}

// RemediationDeploymentSummary - The deployment status summary for all deployments created by the remediation.
type RemediationDeploymentSummary struct {
	// READ-ONLY; The number of deployments required by the remediation that have failed.
	FailedDeployments *int32

	// READ-ONLY; The number of deployments required by the remediation that have succeeded.
	SuccessfulDeployments *int32

	// READ-ONLY; The number of deployments required by the remediation.
	TotalDeployments *int32
}

// RemediationDeploymentsListResult - List of deployments for a remediation.
type RemediationDeploymentsListResult struct {
	// READ-ONLY; The URL to get the next set of results.
	NextLink *string

	// READ-ONLY; Array of deployments for the remediation.
	Value []*RemediationDeployment
}

// RemediationFilters - The filters that will be applied to determine which resources to remediate.
type RemediationFilters struct {
	// The resource locations that will be remediated.
	Locations []*string

	// The IDs of the resources that will be remediated. Can specify at most 100 IDs. This filter cannot be used when ReEvaluateCompliance
	// is set to ReEvaluateCompliance, and cannot be empty if provided.
	ResourceIDs []*string
}

// RemediationListResult - List of remediations.
type RemediationListResult struct {
	// READ-ONLY; The URL to get the next set of results.
	NextLink *string

	// READ-ONLY; Array of remediation definitions.
	Value []*Remediation
}

// RemediationProperties - The remediation properties.
type RemediationProperties struct {
	// The remediation failure threshold settings
	FailureThreshold *RemediationPropertiesFailureThreshold

	// The filters that will be applied to determine which resources to remediate.
	Filters *RemediationFilters

	// Determines how many resources to remediate at any given time. Can be used to increase or reduce the pace of the remediation.
	// If not provided, the default parallel deployments value is used.
	ParallelDeployments *int32

	// The resource ID of the policy assignment that should be remediated.
	PolicyAssignmentID *string

	// The policy definition reference ID of the individual definition that should be remediated. Required when the policy assignment
	// being remediated assigns a policy set definition.
	PolicyDefinitionReferenceID *string

	// Determines the max number of resources that can be remediated by the remediation job. If not provided, the default resource
	// count is used.
	ResourceCount *int32

	// The way resources to remediate are discovered. Defaults to ExistingNonCompliant if not specified.
	ResourceDiscoveryMode *ResourceDiscoveryMode

	// READ-ONLY; The remediation correlation Id. Can be used to find events related to the remediation in the activity log.
	CorrelationID *string

	// READ-ONLY; The time at which the remediation was created.
	CreatedOn *time.Time

	// READ-ONLY; The deployment status summary for all deployments created by the remediation.
	DeploymentStatus *RemediationDeploymentSummary

	// READ-ONLY; The time at which the remediation was last updated.
	LastUpdatedOn *time.Time

	// READ-ONLY; The status of the remediation. This refers to the entire remediation task, not individual deployments. Allowed
	// values are Evaluating, Canceled, Cancelling, Failed, Complete, or Succeeded.
	ProvisioningState *string

	// READ-ONLY; The remediation status message. Provides additional details regarding the state of the remediation.
	StatusMessage *string
}

// RemediationPropertiesFailureThreshold - The remediation failure threshold settings
type RemediationPropertiesFailureThreshold struct {
	// A number between 0.0 to 1.0 representing the percentage failure threshold. The remediation will fail if the percentage
	// of failed remediation operations (i.e. failed deployments) exceeds this
	// threshold.
	Percentage *float32
}

// Resource - Common fields that are returned in the response for all Azure Resource Manager resources
type Resource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// SlimPolicyMetadata - Slim version of policy metadata resource definition, excluding properties with large strings
type SlimPolicyMetadata struct {
	// Properties of the policy metadata.
	Properties *PolicyMetadataSlimProperties

	// READ-ONLY; The ID of the policy metadata.
	ID *string

	// READ-ONLY; The name of the policy metadata.
	Name *string

	// READ-ONLY; The type of the policy metadata.
	Type *string
}

// SummarizeResults - Summarize action results.
type SummarizeResults struct {
	// OData context string; used by OData clients to resolve type information based on metadata.
	ODataContext *string

	// OData entity count; represents the number of summaries returned; always set to 1.
	ODataCount *int32

	// Summarize action results.
	Value []*Summary
}

// Summary results.
type Summary struct {
	// OData context string; used by OData clients to resolve type information based on metadata.
	ODataContext *string

	// OData entity ID; always set to null since summaries do not have an entity ID.
	ODataID *string

	// Policy assignments summary.
	PolicyAssignments []*PolicyAssignmentSummary

	// Compliance summary for all policy assignments.
	Results *SummaryResults
}

// SummaryResults - Compliance summary on a particular summary level.
type SummaryResults struct {
	// Number of non-compliant policies.
	NonCompliantPolicies *int32

	// Number of non-compliant resources.
	NonCompliantResources *int32

	// The policy artifact summary at this level. For query scope level, it represents policy assignment summary. For policy assignment
	// level, it represents policy definitions summary.
	PolicyDetails []*ComplianceDetail

	// The policy definition group summary at this level.
	PolicyGroupDetails []*ComplianceDetail

	// HTTP POST URI for queryResults action on Microsoft.PolicyInsights to retrieve raw results for the compliance summary. This
	// property will not be available by default in future API versions, but could
	// be queried explicitly.
	QueryResultsURI *string

	// The resources summary at this level.
	ResourceDetails []*ComplianceDetail
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

// TrackedResourceModificationDetails - The details of the policy triggered deployment that created or modified the tracked
// resource.
type TrackedResourceModificationDetails struct {
	// READ-ONLY; The ID of the deployment that created or modified the tracked resource.
	DeploymentID *string

	// READ-ONLY; Timestamp of the deployment that created or modified the tracked resource.
	DeploymentTime *time.Time

	// READ-ONLY; The details of the policy that created or modified the tracked resource.
	PolicyDetails *PolicyDetails
}

// TypedErrorInfo - Scenario specific error details.
type TypedErrorInfo struct {
	// READ-ONLY; The scenario specific error details.
	Info any

	// READ-ONLY; The type of included error details.
	Type *string
}
