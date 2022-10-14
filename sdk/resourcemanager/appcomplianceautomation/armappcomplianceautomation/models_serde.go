//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armappcomplianceautomation

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type Assessment.
func (a Assessment) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "description", a.Description)
	populate(objectMap, "isPass", a.IsPass)
	populate(objectMap, "name", a.Name)
	populate(objectMap, "policyId", a.PolicyID)
	populate(objectMap, "remediation", a.Remediation)
	populate(objectMap, "resourceList", a.ResourceList)
	populate(objectMap, "severity", a.Severity)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Assessment.
func (a *Assessment) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", a, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "description":
			err = unpopulate(val, "Description", &a.Description)
			delete(rawMsg, key)
		case "isPass":
			err = unpopulate(val, "IsPass", &a.IsPass)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &a.Name)
			delete(rawMsg, key)
		case "policyId":
			err = unpopulate(val, "PolicyID", &a.PolicyID)
			delete(rawMsg, key)
		case "remediation":
			err = unpopulate(val, "Remediation", &a.Remediation)
			delete(rawMsg, key)
		case "resourceList":
			err = unpopulate(val, "ResourceList", &a.ResourceList)
			delete(rawMsg, key)
		case "severity":
			err = unpopulate(val, "Severity", &a.Severity)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", a, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type AssessmentResource.
func (a AssessmentResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "reason", a.Reason)
	populate(objectMap, "resourceId", a.ResourceID)
	populate(objectMap, "resourceStatus", a.ResourceStatus)
	populate(objectMap, "statusChangeDate", a.StatusChangeDate)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type AssessmentResource.
func (a *AssessmentResource) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", a, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "reason":
			err = unpopulate(val, "Reason", &a.Reason)
			delete(rawMsg, key)
		case "resourceId":
			err = unpopulate(val, "ResourceID", &a.ResourceID)
			delete(rawMsg, key)
		case "resourceStatus":
			err = unpopulate(val, "ResourceStatus", &a.ResourceStatus)
			delete(rawMsg, key)
		case "statusChangeDate":
			err = unpopulate(val, "StatusChangeDate", &a.StatusChangeDate)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", a, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Category.
func (c Category) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "categoryName", c.CategoryName)
	populate(objectMap, "categoryStatus", c.CategoryStatus)
	populate(objectMap, "categoryType", c.CategoryType)
	populate(objectMap, "controlFamilies", c.ControlFamilies)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Category.
func (c *Category) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "categoryName":
			err = unpopulate(val, "CategoryName", &c.CategoryName)
			delete(rawMsg, key)
		case "categoryStatus":
			err = unpopulate(val, "CategoryStatus", &c.CategoryStatus)
			delete(rawMsg, key)
		case "categoryType":
			err = unpopulate(val, "CategoryType", &c.CategoryType)
			delete(rawMsg, key)
		case "controlFamilies":
			err = unpopulate(val, "ControlFamilies", &c.ControlFamilies)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ComplianceReportItem.
func (c ComplianceReportItem) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "categoryName", c.CategoryName)
	populate(objectMap, "complianceState", c.ComplianceState)
	populate(objectMap, "controlId", c.ControlID)
	populate(objectMap, "controlName", c.ControlName)
	populate(objectMap, "controlType", c.ControlType)
	populate(objectMap, "policyDescription", c.PolicyDescription)
	populate(objectMap, "policyDisplayName", c.PolicyDisplayName)
	populate(objectMap, "policyId", c.PolicyID)
	populate(objectMap, "resourceGroup", c.ResourceGroup)
	populate(objectMap, "resourceId", c.ResourceID)
	populate(objectMap, "resourceType", c.ResourceType)
	populate(objectMap, "statusChangeDate", c.StatusChangeDate)
	populate(objectMap, "subscriptionId", c.SubscriptionID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ComplianceReportItem.
func (c *ComplianceReportItem) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "categoryName":
			err = unpopulate(val, "CategoryName", &c.CategoryName)
			delete(rawMsg, key)
		case "complianceState":
			err = unpopulate(val, "ComplianceState", &c.ComplianceState)
			delete(rawMsg, key)
		case "controlId":
			err = unpopulate(val, "ControlID", &c.ControlID)
			delete(rawMsg, key)
		case "controlName":
			err = unpopulate(val, "ControlName", &c.ControlName)
			delete(rawMsg, key)
		case "controlType":
			err = unpopulate(val, "ControlType", &c.ControlType)
			delete(rawMsg, key)
		case "policyDescription":
			err = unpopulate(val, "PolicyDescription", &c.PolicyDescription)
			delete(rawMsg, key)
		case "policyDisplayName":
			err = unpopulate(val, "PolicyDisplayName", &c.PolicyDisplayName)
			delete(rawMsg, key)
		case "policyId":
			err = unpopulate(val, "PolicyID", &c.PolicyID)
			delete(rawMsg, key)
		case "resourceGroup":
			err = unpopulate(val, "ResourceGroup", &c.ResourceGroup)
			delete(rawMsg, key)
		case "resourceId":
			err = unpopulate(val, "ResourceID", &c.ResourceID)
			delete(rawMsg, key)
		case "resourceType":
			err = unpopulate(val, "ResourceType", &c.ResourceType)
			delete(rawMsg, key)
		case "statusChangeDate":
			err = unpopulate(val, "StatusChangeDate", &c.StatusChangeDate)
			delete(rawMsg, key)
		case "subscriptionId":
			err = unpopulate(val, "SubscriptionID", &c.SubscriptionID)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ComplianceResult.
func (c ComplianceResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "categories", c.Categories)
	populate(objectMap, "complianceName", c.ComplianceName)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ComplianceResult.
func (c *ComplianceResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "categories":
			err = unpopulate(val, "Categories", &c.Categories)
			delete(rawMsg, key)
		case "complianceName":
			err = unpopulate(val, "ComplianceName", &c.ComplianceName)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Control.
func (c Control) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "assessments", c.Assessments)
	populate(objectMap, "controlDescription", c.ControlDescription)
	populate(objectMap, "controlDescriptionHyperLink", c.ControlDescriptionHyperLink)
	populate(objectMap, "controlFullName", c.ControlFullName)
	populate(objectMap, "controlId", c.ControlID)
	populate(objectMap, "controlShortName", c.ControlShortName)
	populate(objectMap, "controlStatus", c.ControlStatus)
	populate(objectMap, "controlType", c.ControlType)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Control.
func (c *Control) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "assessments":
			err = unpopulate(val, "Assessments", &c.Assessments)
			delete(rawMsg, key)
		case "controlDescription":
			err = unpopulate(val, "ControlDescription", &c.ControlDescription)
			delete(rawMsg, key)
		case "controlDescriptionHyperLink":
			err = unpopulate(val, "ControlDescriptionHyperLink", &c.ControlDescriptionHyperLink)
			delete(rawMsg, key)
		case "controlFullName":
			err = unpopulate(val, "ControlFullName", &c.ControlFullName)
			delete(rawMsg, key)
		case "controlId":
			err = unpopulate(val, "ControlID", &c.ControlID)
			delete(rawMsg, key)
		case "controlShortName":
			err = unpopulate(val, "ControlShortName", &c.ControlShortName)
			delete(rawMsg, key)
		case "controlStatus":
			err = unpopulate(val, "ControlStatus", &c.ControlStatus)
			delete(rawMsg, key)
		case "controlType":
			err = unpopulate(val, "ControlType", &c.ControlType)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ControlFamily.
func (c ControlFamily) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "controls", c.Controls)
	populate(objectMap, "familyName", c.FamilyName)
	populate(objectMap, "familyStatus", c.FamilyStatus)
	populate(objectMap, "familyType", c.FamilyType)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ControlFamily.
func (c *ControlFamily) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "controls":
			err = unpopulate(val, "Controls", &c.Controls)
			delete(rawMsg, key)
		case "familyName":
			err = unpopulate(val, "FamilyName", &c.FamilyName)
			delete(rawMsg, key)
		case "familyStatus":
			err = unpopulate(val, "FamilyStatus", &c.FamilyStatus)
			delete(rawMsg, key)
		case "familyType":
			err = unpopulate(val, "FamilyType", &c.FamilyType)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type DownloadResponse.
func (d DownloadResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "complianceDetailedPdfReport", d.ComplianceDetailedPDFReport)
	populate(objectMap, "compliancePdfReport", d.CompliancePDFReport)
	populate(objectMap, "complianceReport", d.ComplianceReport)
	populate(objectMap, "resourceList", d.ResourceList)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DownloadResponse.
func (d *DownloadResponse) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "complianceDetailedPdfReport":
			err = unpopulate(val, "ComplianceDetailedPDFReport", &d.ComplianceDetailedPDFReport)
			delete(rawMsg, key)
		case "compliancePdfReport":
			err = unpopulate(val, "CompliancePDFReport", &d.CompliancePDFReport)
			delete(rawMsg, key)
		case "complianceReport":
			err = unpopulate(val, "ComplianceReport", &d.ComplianceReport)
			delete(rawMsg, key)
		case "resourceList":
			err = unpopulate(val, "ResourceList", &d.ResourceList)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", d, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type DownloadResponseComplianceDetailedPDFReport.
func (d DownloadResponseComplianceDetailedPDFReport) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "sasUri", d.SasURI)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DownloadResponseComplianceDetailedPDFReport.
func (d *DownloadResponseComplianceDetailedPDFReport) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "sasUri":
			err = unpopulate(val, "SasURI", &d.SasURI)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", d, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type DownloadResponseCompliancePDFReport.
func (d DownloadResponseCompliancePDFReport) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "sasUri", d.SasURI)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DownloadResponseCompliancePDFReport.
func (d *DownloadResponseCompliancePDFReport) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "sasUri":
			err = unpopulate(val, "SasURI", &d.SasURI)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", d, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Operation.
func (o Operation) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "actionType", o.ActionType)
	populate(objectMap, "display", o.Display)
	populate(objectMap, "isDataAction", o.IsDataAction)
	populate(objectMap, "name", o.Name)
	populate(objectMap, "origin", o.Origin)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Operation.
func (o *Operation) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", o, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "actionType":
			err = unpopulate(val, "ActionType", &o.ActionType)
			delete(rawMsg, key)
		case "display":
			err = unpopulate(val, "Display", &o.Display)
			delete(rawMsg, key)
		case "isDataAction":
			err = unpopulate(val, "IsDataAction", &o.IsDataAction)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &o.Name)
			delete(rawMsg, key)
		case "origin":
			err = unpopulate(val, "Origin", &o.Origin)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", o, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type OperationDisplay.
func (o OperationDisplay) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "description", o.Description)
	populate(objectMap, "operation", o.Operation)
	populate(objectMap, "provider", o.Provider)
	populate(objectMap, "resource", o.Resource)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type OperationDisplay.
func (o *OperationDisplay) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", o, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "description":
			err = unpopulate(val, "Description", &o.Description)
			delete(rawMsg, key)
		case "operation":
			err = unpopulate(val, "Operation", &o.Operation)
			delete(rawMsg, key)
		case "provider":
			err = unpopulate(val, "Provider", &o.Provider)
			delete(rawMsg, key)
		case "resource":
			err = unpopulate(val, "Resource", &o.Resource)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", o, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type OperationListResult.
func (o OperationListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", o.NextLink)
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type OperationListResult.
func (o *OperationListResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", o, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "nextLink":
			err = unpopulate(val, "NextLink", &o.NextLink)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &o.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", o, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type OverviewStatus.
func (o OverviewStatus) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "failed", o.Failed)
	populate(objectMap, "manual", o.Manual)
	populate(objectMap, "passed", o.Passed)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type OverviewStatus.
func (o *OverviewStatus) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", o, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "failed":
			err = unpopulate(val, "Failed", &o.Failed)
			delete(rawMsg, key)
		case "manual":
			err = unpopulate(val, "Manual", &o.Manual)
			delete(rawMsg, key)
		case "passed":
			err = unpopulate(val, "Passed", &o.Passed)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", o, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ReportComplianceStatus.
func (r ReportComplianceStatus) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "m365", r.M365)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ReportComplianceStatus.
func (r *ReportComplianceStatus) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "m365":
			err = unpopulate(val, "M365", &r.M365)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ReportProperties.
func (r ReportProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "complianceStatus", r.ComplianceStatus)
	populate(objectMap, "id", r.ID)
	populateTimeRFC3339(objectMap, "lastTriggerTime", r.LastTriggerTime)
	populateTimeRFC3339(objectMap, "nextTriggerTime", r.NextTriggerTime)
	populate(objectMap, "offerGuid", r.OfferGUID)
	populate(objectMap, "provisioningState", r.ProvisioningState)
	populate(objectMap, "reportName", r.ReportName)
	populate(objectMap, "resources", r.Resources)
	populate(objectMap, "status", r.Status)
	populate(objectMap, "subscriptions", r.Subscriptions)
	populate(objectMap, "tenantId", r.TenantID)
	populate(objectMap, "timeZone", r.TimeZone)
	populateTimeRFC3339(objectMap, "triggerTime", r.TriggerTime)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ReportProperties.
func (r *ReportProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "complianceStatus":
			err = unpopulate(val, "ComplianceStatus", &r.ComplianceStatus)
			delete(rawMsg, key)
		case "id":
			err = unpopulate(val, "ID", &r.ID)
			delete(rawMsg, key)
		case "lastTriggerTime":
			err = unpopulateTimeRFC3339(val, "LastTriggerTime", &r.LastTriggerTime)
			delete(rawMsg, key)
		case "nextTriggerTime":
			err = unpopulateTimeRFC3339(val, "NextTriggerTime", &r.NextTriggerTime)
			delete(rawMsg, key)
		case "offerGuid":
			err = unpopulate(val, "OfferGUID", &r.OfferGUID)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &r.ProvisioningState)
			delete(rawMsg, key)
		case "reportName":
			err = unpopulate(val, "ReportName", &r.ReportName)
			delete(rawMsg, key)
		case "resources":
			err = unpopulate(val, "Resources", &r.Resources)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, "Status", &r.Status)
			delete(rawMsg, key)
		case "subscriptions":
			err = unpopulate(val, "Subscriptions", &r.Subscriptions)
			delete(rawMsg, key)
		case "tenantId":
			err = unpopulate(val, "TenantID", &r.TenantID)
			delete(rawMsg, key)
		case "timeZone":
			err = unpopulate(val, "TimeZone", &r.TimeZone)
			delete(rawMsg, key)
		case "triggerTime":
			err = unpopulateTimeRFC3339(val, "TriggerTime", &r.TriggerTime)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ReportResource.
func (r ReportResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", r.ID)
	populate(objectMap, "name", r.Name)
	populate(objectMap, "properties", r.Properties)
	populate(objectMap, "systemData", r.SystemData)
	populate(objectMap, "type", r.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ReportResource.
func (r *ReportResource) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &r.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &r.Name)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &r.Properties)
			delete(rawMsg, key)
		case "systemData":
			err = unpopulate(val, "SystemData", &r.SystemData)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &r.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ReportResourceList.
func (r ReportResourceList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", r.NextLink)
	populate(objectMap, "value", r.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ReportResourceList.
func (r *ReportResourceList) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "nextLink":
			err = unpopulate(val, "NextLink", &r.NextLink)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &r.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ReportResourcePatch.
func (r ReportResourcePatch) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "properties", r.Properties)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ReportResourcePatch.
func (r *ReportResourcePatch) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "properties":
			err = unpopulate(val, "Properties", &r.Properties)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ResourceItem.
func (r ResourceItem) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "resourceGroup", r.ResourceGroup)
	populate(objectMap, "resourceId", r.ResourceID)
	populate(objectMap, "resourceType", r.ResourceType)
	populate(objectMap, "subscriptionId", r.SubscriptionID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ResourceItem.
func (r *ResourceItem) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "resourceGroup":
			err = unpopulate(val, "ResourceGroup", &r.ResourceGroup)
			delete(rawMsg, key)
		case "resourceId":
			err = unpopulate(val, "ResourceID", &r.ResourceID)
			delete(rawMsg, key)
		case "resourceType":
			err = unpopulate(val, "ResourceType", &r.ResourceType)
			delete(rawMsg, key)
		case "subscriptionId":
			err = unpopulate(val, "SubscriptionID", &r.SubscriptionID)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ResourceMetadata.
func (r ResourceMetadata) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "resourceId", r.ResourceID)
	populate(objectMap, "resourceKind", r.ResourceKind)
	populate(objectMap, "resourceName", r.ResourceName)
	populate(objectMap, "resourceType", r.ResourceType)
	populate(objectMap, "tags", r.Tags)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ResourceMetadata.
func (r *ResourceMetadata) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "resourceId":
			err = unpopulate(val, "ResourceID", &r.ResourceID)
			delete(rawMsg, key)
		case "resourceKind":
			err = unpopulate(val, "ResourceKind", &r.ResourceKind)
			delete(rawMsg, key)
		case "resourceName":
			err = unpopulate(val, "ResourceName", &r.ResourceName)
			delete(rawMsg, key)
		case "resourceType":
			err = unpopulate(val, "ResourceType", &r.ResourceType)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, "Tags", &r.Tags)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SnapshotProperties.
func (s SnapshotProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "complianceResults", s.ComplianceResults)
	populateTimeRFC3339(objectMap, "createdAt", s.CreatedAt)
	populate(objectMap, "id", s.ID)
	populate(objectMap, "provisioningState", s.ProvisioningState)
	populate(objectMap, "reportProperties", s.ReportProperties)
	populate(objectMap, "reportSystemData", s.ReportSystemData)
	populate(objectMap, "snapshotName", s.SnapshotName)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SnapshotProperties.
func (s *SnapshotProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "complianceResults":
			err = unpopulate(val, "ComplianceResults", &s.ComplianceResults)
			delete(rawMsg, key)
		case "createdAt":
			err = unpopulateTimeRFC3339(val, "CreatedAt", &s.CreatedAt)
			delete(rawMsg, key)
		case "id":
			err = unpopulate(val, "ID", &s.ID)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &s.ProvisioningState)
			delete(rawMsg, key)
		case "reportProperties":
			err = unpopulate(val, "ReportProperties", &s.ReportProperties)
			delete(rawMsg, key)
		case "reportSystemData":
			err = unpopulate(val, "ReportSystemData", &s.ReportSystemData)
			delete(rawMsg, key)
		case "snapshotName":
			err = unpopulate(val, "SnapshotName", &s.SnapshotName)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SnapshotResource.
func (s SnapshotResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", s.ID)
	populate(objectMap, "name", s.Name)
	populate(objectMap, "properties", s.Properties)
	populate(objectMap, "systemData", s.SystemData)
	populate(objectMap, "type", s.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SnapshotResource.
func (s *SnapshotResource) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &s.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &s.Name)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &s.Properties)
			delete(rawMsg, key)
		case "systemData":
			err = unpopulate(val, "SystemData", &s.SystemData)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &s.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SnapshotResourceList.
func (s SnapshotResourceList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", s.NextLink)
	populate(objectMap, "value", s.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SnapshotResourceList.
func (s *SnapshotResourceList) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "nextLink":
			err = unpopulate(val, "NextLink", &s.NextLink)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &s.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SystemData.
func (s SystemData) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "createdAt", s.CreatedAt)
	populate(objectMap, "createdBy", s.CreatedBy)
	populate(objectMap, "createdByType", s.CreatedByType)
	populateTimeRFC3339(objectMap, "lastModifiedAt", s.LastModifiedAt)
	populate(objectMap, "lastModifiedBy", s.LastModifiedBy)
	populate(objectMap, "lastModifiedByType", s.LastModifiedByType)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SystemData.
func (s *SystemData) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "createdAt":
			err = unpopulateTimeRFC3339(val, "CreatedAt", &s.CreatedAt)
			delete(rawMsg, key)
		case "createdBy":
			err = unpopulate(val, "CreatedBy", &s.CreatedBy)
			delete(rawMsg, key)
		case "createdByType":
			err = unpopulate(val, "CreatedByType", &s.CreatedByType)
			delete(rawMsg, key)
		case "lastModifiedAt":
			err = unpopulateTimeRFC3339(val, "LastModifiedAt", &s.LastModifiedAt)
			delete(rawMsg, key)
		case "lastModifiedBy":
			err = unpopulate(val, "LastModifiedBy", &s.LastModifiedBy)
			delete(rawMsg, key)
		case "lastModifiedByType":
			err = unpopulate(val, "LastModifiedByType", &s.LastModifiedByType)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

func populate(m map[string]interface{}, k string, v interface{}) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, fn string, v interface{}) error {
	if data == nil {
		return nil
	}
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("struct field %s: %v", fn, err)
	}
	return nil
}
