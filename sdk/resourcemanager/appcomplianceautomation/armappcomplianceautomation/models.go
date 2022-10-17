//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armappcomplianceautomation

import "time"

// Assessment - A class represent the assessment.
type Assessment struct {
	// READ-ONLY; The description of the assessment.
	Description *string `json:"description,omitempty" azure:"ro"`

	// READ-ONLY; Indicates whether all the resource(s) are compliant.
	IsPass *IsPass `json:"isPass,omitempty" azure:"ro"`

	// READ-ONLY; The name of the assessment.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The policy id mapping to this assessment.
	PolicyID *string `json:"policyId,omitempty" azure:"ro"`

	// READ-ONLY; The remediation of the assessment.
	Remediation *string `json:"remediation,omitempty" azure:"ro"`

	// READ-ONLY; List of resource assessments.
	ResourceList []*AssessmentResource `json:"resourceList,omitempty" azure:"ro"`

	// READ-ONLY; The severity level of this assessment.
	Severity *AssessmentSeverity `json:"severity,omitempty" azure:"ro"`
}

// AssessmentResource - A class represent the assessment resource.
type AssessmentResource struct {
	// READ-ONLY; The reason for the N/A resource.
	Reason *string `json:"reason,omitempty" azure:"ro"`

	// READ-ONLY; The Id of the resource.
	ResourceID *string `json:"resourceId,omitempty" azure:"ro"`

	// READ-ONLY; Resource status.
	ResourceStatus *ResourceStatus `json:"resourceStatus,omitempty" azure:"ro"`

	// READ-ONLY; The status change date for the resource. For unavailable date, set it as N/A.
	StatusChangeDate *string `json:"statusChangeDate,omitempty" azure:"ro"`
}

// Category - A class represent the compliance category.
type Category struct {
	// READ-ONLY; The name of the compliance category. e.g. "Operational Security"
	CategoryName *string `json:"categoryName,omitempty" azure:"ro"`

	// READ-ONLY; Category status.
	CategoryStatus *CategoryStatus `json:"categoryStatus,omitempty" azure:"ro"`

	// READ-ONLY; The category type
	CategoryType *CategoryType `json:"categoryType,omitempty" azure:"ro"`

	// READ-ONLY; List of control families.
	ControlFamilies []*ControlFamily `json:"controlFamilies,omitempty" azure:"ro"`
}

// ComplianceReportItem - Object that includes all the content for single compliance result.
type ComplianceReportItem struct {
	// READ-ONLY; The category name.
	CategoryName *string `json:"categoryName,omitempty" azure:"ro"`

	// READ-ONLY; The compliance result's status.
	ComplianceState *string `json:"complianceState,omitempty" azure:"ro"`

	// READ-ONLY; The control Id.
	ControlID *string `json:"controlId,omitempty" azure:"ro"`

	// READ-ONLY; The control name.
	ControlName *string `json:"controlName,omitempty" azure:"ro"`

	// READ-ONLY; The control type.
	ControlType *ControlType `json:"controlType,omitempty" azure:"ro"`

	// READ-ONLY; The policy's detail description.
	PolicyDescription *string `json:"policyDescription,omitempty" azure:"ro"`

	// READ-ONLY; The policy's display name.
	PolicyDisplayName *string `json:"policyDisplayName,omitempty" azure:"ro"`

	// READ-ONLY; The compliance result mapped policy Id.
	PolicyID *string `json:"policyId,omitempty" azure:"ro"`

	// READ-ONLY; The compliance result mapped resource group.
	ResourceGroup *string `json:"resourceGroup,omitempty" azure:"ro"`

	// READ-ONLY; The compliance result mapped resource Id.
	ResourceID *string `json:"resourceId,omitempty" azure:"ro"`

	// READ-ONLY; The compliance result mapped resource type.
	ResourceType *string `json:"resourceType,omitempty" azure:"ro"`

	// READ-ONLY; The compliance result last changed date. For unavailable date, set it as N/A.
	StatusChangeDate *string `json:"statusChangeDate,omitempty" azure:"ro"`

	// READ-ONLY; The compliance result mapped subscription Id.
	SubscriptionID *string `json:"subscriptionId,omitempty" azure:"ro"`
}

// ComplianceResult - A class represent the compliance result.
type ComplianceResult struct {
	// READ-ONLY; List of categories.
	Categories []*Category `json:"categories,omitempty" azure:"ro"`

	// READ-ONLY; The name of the compliance. e.g. "M365"
	ComplianceName *string `json:"complianceName,omitempty" azure:"ro"`
}

// Control - A class represent the control.
type Control struct {
	// READ-ONLY; List of assessments.
	Assessments []*Assessment `json:"assessments,omitempty" azure:"ro"`

	// READ-ONLY; The control's description
	ControlDescription *string `json:"controlDescription,omitempty" azure:"ro"`

	// READ-ONLY; The hyper link to the control's description'.
	ControlDescriptionHyperLink *string `json:"controlDescriptionHyperLink,omitempty" azure:"ro"`

	// READ-ONLY; The full name of the control. e.g. "Validate that unsupported operating systems and software components are
	// not in use."
	ControlFullName *string `json:"controlFullName,omitempty" azure:"ro"`

	// READ-ONLY; The Id of the control. e.g. "Operational Security#10"
	ControlID *string `json:"controlId,omitempty" azure:"ro"`

	// READ-ONLY; The short name of the control. e.g. "Unsupported OS and Software."
	ControlShortName *string `json:"controlShortName,omitempty" azure:"ro"`

	// READ-ONLY; Control status.
	ControlStatus *ControlStatus `json:"controlStatus,omitempty" azure:"ro"`

	// READ-ONLY; The control type
	ControlType *ControlType `json:"controlType,omitempty" azure:"ro"`
}

// ControlFamily - A class represent the control family.
type ControlFamily struct {
	// READ-ONLY; List of controls.
	Controls []*Control `json:"controls,omitempty" azure:"ro"`

	// READ-ONLY; The name of the control family. e.g. "Malware Protection - Anti-Virus"
	FamilyName *string `json:"familyName,omitempty" azure:"ro"`

	// READ-ONLY; Control family status.
	FamilyStatus *ControlFamilyStatus `json:"familyStatus,omitempty" azure:"ro"`

	// READ-ONLY; The control family type
	FamilyType *ControlFamilyType `json:"familyType,omitempty" azure:"ro"`
}

// DownloadResponse - Object that includes all the possible response for the download operation.
type DownloadResponse struct {
	// READ-ONLY; compliance detailed pdf report
	ComplianceDetailedPDFReport *DownloadResponseComplianceDetailedPDFReport `json:"complianceDetailedPdfReport,omitempty" azure:"ro"`

	// READ-ONLY; compliance pdf report
	CompliancePDFReport *DownloadResponseCompliancePDFReport `json:"compliancePdfReport,omitempty" azure:"ro"`

	// READ-ONLY; List of the compliance result
	ComplianceReport []*ComplianceReportItem `json:"complianceReport,omitempty" azure:"ro"`

	// READ-ONLY; List of the reports
	ResourceList []*ResourceItem `json:"resourceList,omitempty" azure:"ro"`
}

// DownloadResponseComplianceDetailedPDFReport - compliance detailed pdf report
type DownloadResponseComplianceDetailedPDFReport struct {
	// READ-ONLY; uri of compliance detailed pdf report
	SasURI *string `json:"sasUri,omitempty" azure:"ro"`
}

// DownloadResponseCompliancePDFReport - compliance pdf report
type DownloadResponseCompliancePDFReport struct {
	// READ-ONLY; uri of compliance pdf report
	SasURI *string `json:"sasUri,omitempty" azure:"ro"`
}

// Operation - Details of a REST API operation, returned from the Resource Provider Operations API
type Operation struct {
	// Localized display information for this particular operation.
	Display *OperationDisplay `json:"display,omitempty"`

	// READ-ONLY; Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
	ActionType *ActionType `json:"actionType,omitempty" azure:"ro"`

	// READ-ONLY; Whether the operation applies to data-plane. This is "true" for data-plane operations and "false" for ARM/control-plane
	// operations.
	IsDataAction *bool `json:"isDataAction,omitempty" azure:"ro"`

	// READ-ONLY; The name of the operation, as per Resource-Based Access Control (RBAC). Examples: "Microsoft.Compute/virtualMachines/write",
	// "Microsoft.Compute/virtualMachines/capture/action"
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
	// value is "user,system"
	Origin *Origin `json:"origin,omitempty" azure:"ro"`
}

// OperationDisplay - Localized display information for this particular operation.
type OperationDisplay struct {
	// READ-ONLY; The short, localized friendly description of the operation; suitable for tool tips and detailed views.
	Description *string `json:"description,omitempty" azure:"ro"`

	// READ-ONLY; The concise, localized friendly name for the operation; suitable for dropdowns. E.g. "Create or Update Virtual
	// Machine", "Restart Virtual Machine".
	Operation *string `json:"operation,omitempty" azure:"ro"`

	// READ-ONLY; The localized friendly form of the resource provider name, e.g. "Microsoft Monitoring Insights" or "Microsoft
	// Compute".
	Provider *string `json:"provider,omitempty" azure:"ro"`

	// READ-ONLY; The localized friendly name of the resource type related to this operation. E.g. "Virtual Machines" or "Job
	// Schedule Collections".
	Resource *string `json:"resource,omitempty" azure:"ro"`
}

// OperationListResult - A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to
// get the next set of results.
type OperationListResult struct {
	// READ-ONLY; URL to get the next set of operation list results (if there are any).
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`

	// READ-ONLY; List of operations supported by the resource provider
	Value []*Operation `json:"value,omitempty" azure:"ro"`
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.List method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// OverviewStatus - The overview of the compliance result for one report.
type OverviewStatus struct {
	// The count of all failed full automation control.
	FailedCount *int32 `json:"failedCount,omitempty"`

	// The count of all manual control.
	ManualCount *int32 `json:"manualCount,omitempty"`

	// The count of all passed full automation control.
	PassedCount *int32 `json:"passedCount,omitempty"`
}

// ReportClientBeginCreateOrUpdateOptions contains the optional parameters for the ReportClient.BeginCreateOrUpdate method.
type ReportClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ReportClientBeginDeleteOptions contains the optional parameters for the ReportClient.BeginDelete method.
type ReportClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ReportClientBeginUpdateOptions contains the optional parameters for the ReportClient.BeginUpdate method.
type ReportClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ReportClientGetOptions contains the optional parameters for the ReportClient.Get method.
type ReportClientGetOptions struct {
	// placeholder for future optional parameters
}

// ReportComplianceStatus - A list which includes all the compliance result for one report.
type ReportComplianceStatus struct {
	// The Microsoft 365 certification name.
	M365 *OverviewStatus `json:"m365,omitempty"`
}

// ReportProperties - Report's properties.
type ReportProperties struct {
	// REQUIRED; List of resource data.
	Resources []*ResourceMetadata `json:"resources,omitempty"`

	// REQUIRED; Report collection trigger time's time zone.
	TimeZone *string `json:"timeZone,omitempty"`

	// REQUIRED; Report collection trigger time.
	TriggerTime *time.Time `json:"triggerTime,omitempty"`

	// Report offer Guid.
	OfferGUID *string `json:"offerGuid,omitempty"`

	// READ-ONLY; Report compliance status.
	ComplianceStatus *ReportComplianceStatus `json:"complianceStatus,omitempty" azure:"ro"`

	// READ-ONLY; Report id in database.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Report last collection trigger time.
	LastTriggerTime *time.Time `json:"lastTriggerTime,omitempty" azure:"ro"`

	// READ-ONLY; Report next collection trigger time.
	NextTriggerTime *time.Time `json:"nextTriggerTime,omitempty" azure:"ro"`

	// READ-ONLY; Azure lifecycle management
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty" azure:"ro"`

	// READ-ONLY; Report name.
	ReportName *string `json:"reportName,omitempty" azure:"ro"`

	// READ-ONLY; Report status.
	Status *ReportStatus `json:"status,omitempty" azure:"ro"`

	// READ-ONLY; List of subscription Ids.
	Subscriptions []*string `json:"subscriptions,omitempty" azure:"ro"`

	// READ-ONLY; Report's tenant id.
	TenantID *string `json:"tenantId,omitempty" azure:"ro"`
}

// ReportResource - A class represent a AppComplianceAutomation report resource.
type ReportResource struct {
	// REQUIRED; Report property.
	Properties *ReportProperties `json:"properties,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ReportResourceList - Object that includes an array of resources and a possible link for next set.
type ReportResourceList struct {
	// The URL the client should use to fetch the next page (per server side paging). It's null for now, added for future use.
	NextLink *string `json:"nextLink,omitempty"`

	// READ-ONLY; List of the reports
	Value []*ReportResource `json:"value,omitempty" azure:"ro"`
}

// ReportResourcePatch - A class represent a AppComplianceAutomation report resource update properties.
type ReportResourcePatch struct {
	// Report property.
	Properties *ReportProperties `json:"properties,omitempty"`
}

// ReportsClientListOptions contains the optional parameters for the ReportsClient.List method.
type ReportsClientListOptions struct {
	// The offerGuid which mapping to the reports.
	OfferGUID *string
	// The tenant id of the report creator.
	ReportCreatorTenantID *string
	// OData Select statement. Limits the properties on each entry to just those requested, e.g. ?$select=reportName,id.
	Select *string
	// Skip over when retrieving results.
	SkipToken *string
	// Number of elements to return when retrieving results.
	Top *int32
}

// ResourceItem - Resource Id.
type ResourceItem struct {
	// READ-ONLY; The resource group name of this resource.
	ResourceGroup *string `json:"resourceGroup,omitempty" azure:"ro"`

	// READ-ONLY; The resource Id.
	ResourceID *string `json:"resourceId,omitempty" azure:"ro"`

	// READ-ONLY; The resource type of this resource.
	ResourceType *string `json:"resourceType,omitempty" azure:"ro"`

	// READ-ONLY; The subscription Id of this resource.
	SubscriptionID *string `json:"subscriptionId,omitempty" azure:"ro"`
}

// ResourceMetadata - Single resource Id's metadata.
type ResourceMetadata struct {
	// REQUIRED; Resource Id.
	ResourceID *string `json:"resourceId,omitempty"`

	// Resource kind.
	ResourceKind *string `json:"resourceKind,omitempty"`

	// Resource name.
	ResourceName *string `json:"resourceName,omitempty"`

	// Resource type.
	ResourceType *string `json:"resourceType,omitempty"`

	// Resource's tag type.
	Tags map[string]*string `json:"tags,omitempty"`
}

// SnapshotClientBeginDownloadOptions contains the optional parameters for the SnapshotClient.BeginDownload method.
type SnapshotClientBeginDownloadOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// SnapshotClientGetOptions contains the optional parameters for the SnapshotClient.Get method.
type SnapshotClientGetOptions struct {
	// placeholder for future optional parameters
}

// SnapshotDownloadProperties - Snapshot's download query properties.
type SnapshotDownloadProperties struct {
	// REQUIRED; Indicates the download type.
	DownloadType *DownloadType `json:"downloadType,omitempty"`

	// The offerGuid which mapping to the reports.
	OfferGUID *string `json:"offerGuid,omitempty"`

	// Tenant id.
	ReportCreatorTenantID *string `json:"reportCreatorTenantId,omitempty"`
}

// SnapshotDownloadRequest - Snapshot's download request.
type SnapshotDownloadRequest struct {
	// Snapshot's download properties
	Properties *SnapshotDownloadProperties `json:"properties,omitempty"`
}

// SnapshotProperties - Snapshot's properties.
type SnapshotProperties struct {
	// READ-ONLY; List of compliance results.
	ComplianceResults []*ComplianceResult `json:"complianceResults,omitempty" azure:"ro"`

	// READ-ONLY; The timestamp of resource creation (UTC).
	CreatedAt *time.Time `json:"createdAt,omitempty" azure:"ro"`

	// READ-ONLY; Snapshot id in the database.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Azure lifecycle management
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty" azure:"ro"`

	// READ-ONLY; The report essential info.
	ReportProperties *ReportProperties `json:"reportProperties,omitempty" azure:"ro"`

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource.
	ReportSystemData *SystemData `json:"reportSystemData,omitempty" azure:"ro"`

	// READ-ONLY; Snapshot name.
	SnapshotName *string `json:"snapshotName,omitempty" azure:"ro"`
}

// SnapshotResource - A class represent a AppComplianceAutomation snapshot resource.
type SnapshotResource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Snapshot's property'.
	Properties *SnapshotProperties `json:"properties,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// SnapshotResourceList - Object that includes an array of resources and a possible link for next set.
type SnapshotResourceList struct {
	// The URL the client should use to fetch the next page (per server side paging). It's null for now, added for future use.
	NextLink *string `json:"nextLink,omitempty"`

	// READ-ONLY; List of the snapshots
	Value []*SnapshotResource `json:"value,omitempty" azure:"ro"`
}

// SnapshotsClientListOptions contains the optional parameters for the SnapshotsClient.List method.
type SnapshotsClientListOptions struct {
	// The offerGuid which mapping to the reports.
	OfferGUID *string
	// The tenant id of the report creator.
	ReportCreatorTenantID *string
	// OData Select statement. Limits the properties on each entry to just those requested, e.g. ?$select=reportName,id.
	Select *string
	// Skip over when retrieving results.
	SkipToken *string
	// Number of elements to return when retrieving results.
	Top *int32
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// The type of identity that created the resource.
	CreatedByType *CreatedByType `json:"createdByType,omitempty"`

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time `json:"lastModifiedAt,omitempty"`

	// The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType `json:"lastModifiedByType,omitempty"`
}
