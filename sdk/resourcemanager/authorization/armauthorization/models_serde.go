//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armauthorization

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type ApprovalSettings.
func (a ApprovalSettings) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "approvalMode", a.ApprovalMode)
	populate(objectMap, "approvalStages", a.ApprovalStages)
	populate(objectMap, "isApprovalRequired", a.IsApprovalRequired)
	populate(objectMap, "isApprovalRequiredForExtension", a.IsApprovalRequiredForExtension)
	populate(objectMap, "isRequestorJustificationRequired", a.IsRequestorJustificationRequired)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ApprovalSettings.
func (a *ApprovalSettings) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", a, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "approvalMode":
			err = unpopulate(val, "ApprovalMode", &a.ApprovalMode)
			delete(rawMsg, key)
		case "approvalStages":
			err = unpopulate(val, "ApprovalStages", &a.ApprovalStages)
			delete(rawMsg, key)
		case "isApprovalRequired":
			err = unpopulate(val, "IsApprovalRequired", &a.IsApprovalRequired)
			delete(rawMsg, key)
		case "isApprovalRequiredForExtension":
			err = unpopulate(val, "IsApprovalRequiredForExtension", &a.IsApprovalRequiredForExtension)
			delete(rawMsg, key)
		case "isRequestorJustificationRequired":
			err = unpopulate(val, "IsRequestorJustificationRequired", &a.IsRequestorJustificationRequired)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", a, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ApprovalStage.
func (a ApprovalStage) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "approvalStageTimeOutInDays", a.ApprovalStageTimeOutInDays)
	populate(objectMap, "escalationApprovers", a.EscalationApprovers)
	populate(objectMap, "escalationTimeInMinutes", a.EscalationTimeInMinutes)
	populate(objectMap, "isApproverJustificationRequired", a.IsApproverJustificationRequired)
	populate(objectMap, "isEscalationEnabled", a.IsEscalationEnabled)
	populate(objectMap, "primaryApprovers", a.PrimaryApprovers)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ApprovalStage.
func (a *ApprovalStage) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", a, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "approvalStageTimeOutInDays":
			err = unpopulate(val, "ApprovalStageTimeOutInDays", &a.ApprovalStageTimeOutInDays)
			delete(rawMsg, key)
		case "escalationApprovers":
			err = unpopulate(val, "EscalationApprovers", &a.EscalationApprovers)
			delete(rawMsg, key)
		case "escalationTimeInMinutes":
			err = unpopulate(val, "EscalationTimeInMinutes", &a.EscalationTimeInMinutes)
			delete(rawMsg, key)
		case "isApproverJustificationRequired":
			err = unpopulate(val, "IsApproverJustificationRequired", &a.IsApproverJustificationRequired)
			delete(rawMsg, key)
		case "isEscalationEnabled":
			err = unpopulate(val, "IsEscalationEnabled", &a.IsEscalationEnabled)
			delete(rawMsg, key)
		case "primaryApprovers":
			err = unpopulate(val, "PrimaryApprovers", &a.PrimaryApprovers)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", a, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type PIMOnlyModeSettings.
func (p PIMOnlyModeSettings) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "excludeAllServicePrincipals", p.ExcludeAllServicePrincipals)
	populate(objectMap, "excludes", p.Excludes)
	populate(objectMap, "mode", p.Mode)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PIMOnlyModeSettings.
func (p *PIMOnlyModeSettings) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "excludeAllServicePrincipals":
			err = unpopulate(val, "ExcludeAllServicePrincipals", &p.ExcludeAllServicePrincipals)
			delete(rawMsg, key)
		case "excludes":
			err = unpopulate(val, "Excludes", &p.Excludes)
			delete(rawMsg, key)
		case "mode":
			err = unpopulate(val, "Mode", &p.Mode)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Permission.
func (p Permission) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "actions", p.Actions)
	populate(objectMap, "dataActions", p.DataActions)
	populate(objectMap, "notActions", p.NotActions)
	populate(objectMap, "notDataActions", p.NotDataActions)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Permission.
func (p *Permission) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "actions":
			err = unpopulate(val, "Actions", &p.Actions)
			delete(rawMsg, key)
		case "dataActions":
			err = unpopulate(val, "DataActions", &p.DataActions)
			delete(rawMsg, key)
		case "notActions":
			err = unpopulate(val, "NotActions", &p.NotActions)
			delete(rawMsg, key)
		case "notDataActions":
			err = unpopulate(val, "NotDataActions", &p.NotDataActions)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type PolicyAssignmentProperties.
func (p PolicyAssignmentProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "policy", p.Policy)
	populate(objectMap, "roleDefinition", p.RoleDefinition)
	populate(objectMap, "scope", p.Scope)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PolicyAssignmentProperties.
func (p *PolicyAssignmentProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "policy":
			err = unpopulate(val, "Policy", &p.Policy)
			delete(rawMsg, key)
		case "roleDefinition":
			err = unpopulate(val, "RoleDefinition", &p.RoleDefinition)
			delete(rawMsg, key)
		case "scope":
			err = unpopulate(val, "Scope", &p.Scope)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type PolicyAssignmentPropertiesPolicy.
func (p PolicyAssignmentPropertiesPolicy) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", p.ID)
	populate(objectMap, "lastModifiedBy", p.LastModifiedBy)
	populateDateTimeRFC3339(objectMap, "lastModifiedDateTime", p.LastModifiedDateTime)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PolicyAssignmentPropertiesPolicy.
func (p *PolicyAssignmentPropertiesPolicy) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &p.ID)
			delete(rawMsg, key)
		case "lastModifiedBy":
			err = unpopulate(val, "LastModifiedBy", &p.LastModifiedBy)
			delete(rawMsg, key)
		case "lastModifiedDateTime":
			err = unpopulateDateTimeRFC3339(val, "LastModifiedDateTime", &p.LastModifiedDateTime)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type PolicyAssignmentPropertiesRoleDefinition.
func (p PolicyAssignmentPropertiesRoleDefinition) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "displayName", p.DisplayName)
	populate(objectMap, "id", p.ID)
	populate(objectMap, "type", p.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PolicyAssignmentPropertiesRoleDefinition.
func (p *PolicyAssignmentPropertiesRoleDefinition) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "displayName":
			err = unpopulate(val, "DisplayName", &p.DisplayName)
			delete(rawMsg, key)
		case "id":
			err = unpopulate(val, "ID", &p.ID)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &p.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type PolicyAssignmentPropertiesScope.
func (p PolicyAssignmentPropertiesScope) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "displayName", p.DisplayName)
	populate(objectMap, "id", p.ID)
	populate(objectMap, "type", p.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PolicyAssignmentPropertiesScope.
func (p *PolicyAssignmentPropertiesScope) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "displayName":
			err = unpopulate(val, "DisplayName", &p.DisplayName)
			delete(rawMsg, key)
		case "id":
			err = unpopulate(val, "ID", &p.ID)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &p.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type PolicyProperties.
func (p PolicyProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "scope", p.Scope)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PolicyProperties.
func (p *PolicyProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "scope":
			err = unpopulate(val, "Scope", &p.Scope)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type PolicyPropertiesScope.
func (p PolicyPropertiesScope) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "displayName", p.DisplayName)
	populate(objectMap, "id", p.ID)
	populate(objectMap, "type", p.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PolicyPropertiesScope.
func (p *PolicyPropertiesScope) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "displayName":
			err = unpopulate(val, "DisplayName", &p.DisplayName)
			delete(rawMsg, key)
		case "id":
			err = unpopulate(val, "ID", &p.ID)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &p.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Principal.
func (p Principal) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "displayName", p.DisplayName)
	populate(objectMap, "email", p.Email)
	populate(objectMap, "id", p.ID)
	populate(objectMap, "type", p.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Principal.
func (p *Principal) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "displayName":
			err = unpopulate(val, "DisplayName", &p.DisplayName)
			delete(rawMsg, key)
		case "email":
			err = unpopulate(val, "Email", &p.Email)
			delete(rawMsg, key)
		case "id":
			err = unpopulate(val, "ID", &p.ID)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &p.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type RoleManagementPolicy.
func (r RoleManagementPolicy) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", r.ID)
	populate(objectMap, "name", r.Name)
	populate(objectMap, "properties", r.Properties)
	populate(objectMap, "type", r.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RoleManagementPolicy.
func (r *RoleManagementPolicy) UnmarshalJSON(data []byte) error {
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

// MarshalJSON implements the json.Marshaller interface for type RoleManagementPolicyApprovalRule.
func (r RoleManagementPolicyApprovalRule) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", r.ID)
	objectMap["ruleType"] = RoleManagementPolicyRuleTypeRoleManagementPolicyApprovalRule
	populate(objectMap, "setting", r.Setting)
	populate(objectMap, "target", r.Target)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RoleManagementPolicyApprovalRule.
func (r *RoleManagementPolicyApprovalRule) UnmarshalJSON(data []byte) error {
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
		case "ruleType":
			err = unpopulate(val, "RuleType", &r.RuleType)
			delete(rawMsg, key)
		case "setting":
			err = unpopulate(val, "Setting", &r.Setting)
			delete(rawMsg, key)
		case "target":
			err = unpopulate(val, "Target", &r.Target)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type RoleManagementPolicyAssignment.
func (r RoleManagementPolicyAssignment) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", r.ID)
	populate(objectMap, "name", r.Name)
	populate(objectMap, "properties", r.Properties)
	populate(objectMap, "type", r.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RoleManagementPolicyAssignment.
func (r *RoleManagementPolicyAssignment) UnmarshalJSON(data []byte) error {
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

// MarshalJSON implements the json.Marshaller interface for type RoleManagementPolicyAssignmentListResult.
func (r RoleManagementPolicyAssignmentListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "nextLink", r.NextLink)
	populate(objectMap, "value", r.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RoleManagementPolicyAssignmentListResult.
func (r *RoleManagementPolicyAssignmentListResult) UnmarshalJSON(data []byte) error {
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

// MarshalJSON implements the json.Marshaller interface for type RoleManagementPolicyAssignmentProperties.
func (r RoleManagementPolicyAssignmentProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "effectiveRules", r.EffectiveRules)
	populate(objectMap, "policyAssignmentProperties", r.PolicyAssignmentProperties)
	populate(objectMap, "policyId", r.PolicyID)
	populate(objectMap, "roleDefinitionId", r.RoleDefinitionID)
	populate(objectMap, "scope", r.Scope)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RoleManagementPolicyAssignmentProperties.
func (r *RoleManagementPolicyAssignmentProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "effectiveRules":
			r.EffectiveRules, err = unmarshalRoleManagementPolicyRuleClassificationArray(val)
			delete(rawMsg, key)
		case "policyAssignmentProperties":
			err = unpopulate(val, "PolicyAssignmentProperties", &r.PolicyAssignmentProperties)
			delete(rawMsg, key)
		case "policyId":
			err = unpopulate(val, "PolicyID", &r.PolicyID)
			delete(rawMsg, key)
		case "roleDefinitionId":
			err = unpopulate(val, "RoleDefinitionID", &r.RoleDefinitionID)
			delete(rawMsg, key)
		case "scope":
			err = unpopulate(val, "Scope", &r.Scope)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type RoleManagementPolicyAuthenticationContextRule.
func (r RoleManagementPolicyAuthenticationContextRule) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "claimValue", r.ClaimValue)
	populate(objectMap, "id", r.ID)
	populate(objectMap, "isEnabled", r.IsEnabled)
	objectMap["ruleType"] = RoleManagementPolicyRuleTypeRoleManagementPolicyAuthenticationContextRule
	populate(objectMap, "target", r.Target)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RoleManagementPolicyAuthenticationContextRule.
func (r *RoleManagementPolicyAuthenticationContextRule) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "claimValue":
			err = unpopulate(val, "ClaimValue", &r.ClaimValue)
			delete(rawMsg, key)
		case "id":
			err = unpopulate(val, "ID", &r.ID)
			delete(rawMsg, key)
		case "isEnabled":
			err = unpopulate(val, "IsEnabled", &r.IsEnabled)
			delete(rawMsg, key)
		case "ruleType":
			err = unpopulate(val, "RuleType", &r.RuleType)
			delete(rawMsg, key)
		case "target":
			err = unpopulate(val, "Target", &r.Target)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type RoleManagementPolicyEnablementRule.
func (r RoleManagementPolicyEnablementRule) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "enabledRules", r.EnabledRules)
	populate(objectMap, "id", r.ID)
	objectMap["ruleType"] = RoleManagementPolicyRuleTypeRoleManagementPolicyEnablementRule
	populate(objectMap, "target", r.Target)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RoleManagementPolicyEnablementRule.
func (r *RoleManagementPolicyEnablementRule) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "enabledRules":
			err = unpopulate(val, "EnabledRules", &r.EnabledRules)
			delete(rawMsg, key)
		case "id":
			err = unpopulate(val, "ID", &r.ID)
			delete(rawMsg, key)
		case "ruleType":
			err = unpopulate(val, "RuleType", &r.RuleType)
			delete(rawMsg, key)
		case "target":
			err = unpopulate(val, "Target", &r.Target)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type RoleManagementPolicyExpirationRule.
func (r RoleManagementPolicyExpirationRule) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", r.ID)
	populate(objectMap, "isExpirationRequired", r.IsExpirationRequired)
	populate(objectMap, "maximumDuration", r.MaximumDuration)
	objectMap["ruleType"] = RoleManagementPolicyRuleTypeRoleManagementPolicyExpirationRule
	populate(objectMap, "target", r.Target)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RoleManagementPolicyExpirationRule.
func (r *RoleManagementPolicyExpirationRule) UnmarshalJSON(data []byte) error {
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
		case "isExpirationRequired":
			err = unpopulate(val, "IsExpirationRequired", &r.IsExpirationRequired)
			delete(rawMsg, key)
		case "maximumDuration":
			err = unpopulate(val, "MaximumDuration", &r.MaximumDuration)
			delete(rawMsg, key)
		case "ruleType":
			err = unpopulate(val, "RuleType", &r.RuleType)
			delete(rawMsg, key)
		case "target":
			err = unpopulate(val, "Target", &r.Target)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type RoleManagementPolicyListResult.
func (r RoleManagementPolicyListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "nextLink", r.NextLink)
	populate(objectMap, "value", r.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RoleManagementPolicyListResult.
func (r *RoleManagementPolicyListResult) UnmarshalJSON(data []byte) error {
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

// MarshalJSON implements the json.Marshaller interface for type RoleManagementPolicyNotificationRule.
func (r RoleManagementPolicyNotificationRule) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", r.ID)
	populate(objectMap, "isDefaultRecipientsEnabled", r.IsDefaultRecipientsEnabled)
	populate(objectMap, "notificationLevel", r.NotificationLevel)
	populate(objectMap, "notificationRecipients", r.NotificationRecipients)
	populate(objectMap, "notificationType", r.NotificationType)
	populate(objectMap, "recipientType", r.RecipientType)
	objectMap["ruleType"] = RoleManagementPolicyRuleTypeRoleManagementPolicyNotificationRule
	populate(objectMap, "target", r.Target)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RoleManagementPolicyNotificationRule.
func (r *RoleManagementPolicyNotificationRule) UnmarshalJSON(data []byte) error {
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
		case "isDefaultRecipientsEnabled":
			err = unpopulate(val, "IsDefaultRecipientsEnabled", &r.IsDefaultRecipientsEnabled)
			delete(rawMsg, key)
		case "notificationLevel":
			err = unpopulate(val, "NotificationLevel", &r.NotificationLevel)
			delete(rawMsg, key)
		case "notificationRecipients":
			err = unpopulate(val, "NotificationRecipients", &r.NotificationRecipients)
			delete(rawMsg, key)
		case "notificationType":
			err = unpopulate(val, "NotificationType", &r.NotificationType)
			delete(rawMsg, key)
		case "recipientType":
			err = unpopulate(val, "RecipientType", &r.RecipientType)
			delete(rawMsg, key)
		case "ruleType":
			err = unpopulate(val, "RuleType", &r.RuleType)
			delete(rawMsg, key)
		case "target":
			err = unpopulate(val, "Target", &r.Target)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type RoleManagementPolicyPimOnlyModeRule.
func (r RoleManagementPolicyPimOnlyModeRule) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", r.ID)
	populate(objectMap, "pimOnlyModeSettings", r.PimOnlyModeSettings)
	objectMap["ruleType"] = RoleManagementPolicyRuleTypeRoleManagementPolicyPimOnlyModeRule
	populate(objectMap, "target", r.Target)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RoleManagementPolicyPimOnlyModeRule.
func (r *RoleManagementPolicyPimOnlyModeRule) UnmarshalJSON(data []byte) error {
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
		case "pimOnlyModeSettings":
			err = unpopulate(val, "PimOnlyModeSettings", &r.PimOnlyModeSettings)
			delete(rawMsg, key)
		case "ruleType":
			err = unpopulate(val, "RuleType", &r.RuleType)
			delete(rawMsg, key)
		case "target":
			err = unpopulate(val, "Target", &r.Target)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type RoleManagementPolicyProperties.
func (r RoleManagementPolicyProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "description", r.Description)
	populate(objectMap, "displayName", r.DisplayName)
	populate(objectMap, "effectiveRules", r.EffectiveRules)
	populate(objectMap, "isOrganizationDefault", r.IsOrganizationDefault)
	populate(objectMap, "lastModifiedBy", r.LastModifiedBy)
	populateDateTimeRFC3339(objectMap, "lastModifiedDateTime", r.LastModifiedDateTime)
	populate(objectMap, "policyProperties", r.PolicyProperties)
	populate(objectMap, "rules", r.Rules)
	populate(objectMap, "scope", r.Scope)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RoleManagementPolicyProperties.
func (r *RoleManagementPolicyProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "description":
			err = unpopulate(val, "Description", &r.Description)
			delete(rawMsg, key)
		case "displayName":
			err = unpopulate(val, "DisplayName", &r.DisplayName)
			delete(rawMsg, key)
		case "effectiveRules":
			r.EffectiveRules, err = unmarshalRoleManagementPolicyRuleClassificationArray(val)
			delete(rawMsg, key)
		case "isOrganizationDefault":
			err = unpopulate(val, "IsOrganizationDefault", &r.IsOrganizationDefault)
			delete(rawMsg, key)
		case "lastModifiedBy":
			err = unpopulate(val, "LastModifiedBy", &r.LastModifiedBy)
			delete(rawMsg, key)
		case "lastModifiedDateTime":
			err = unpopulateDateTimeRFC3339(val, "LastModifiedDateTime", &r.LastModifiedDateTime)
			delete(rawMsg, key)
		case "policyProperties":
			err = unpopulate(val, "PolicyProperties", &r.PolicyProperties)
			delete(rawMsg, key)
		case "rules":
			r.Rules, err = unmarshalRoleManagementPolicyRuleClassificationArray(val)
			delete(rawMsg, key)
		case "scope":
			err = unpopulate(val, "Scope", &r.Scope)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type RoleManagementPolicyRule.
func (r RoleManagementPolicyRule) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", r.ID)
	objectMap["ruleType"] = r.RuleType
	populate(objectMap, "target", r.Target)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RoleManagementPolicyRule.
func (r *RoleManagementPolicyRule) UnmarshalJSON(data []byte) error {
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
		case "ruleType":
			err = unpopulate(val, "RuleType", &r.RuleType)
			delete(rawMsg, key)
		case "target":
			err = unpopulate(val, "Target", &r.Target)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type RoleManagementPolicyRuleTarget.
func (r RoleManagementPolicyRuleTarget) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "caller", r.Caller)
	populate(objectMap, "enforcedSettings", r.EnforcedSettings)
	populate(objectMap, "inheritableSettings", r.InheritableSettings)
	populate(objectMap, "level", r.Level)
	populate(objectMap, "operations", r.Operations)
	populate(objectMap, "targetObjects", r.TargetObjects)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RoleManagementPolicyRuleTarget.
func (r *RoleManagementPolicyRuleTarget) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "caller":
			err = unpopulate(val, "Caller", &r.Caller)
			delete(rawMsg, key)
		case "enforcedSettings":
			err = unpopulate(val, "EnforcedSettings", &r.EnforcedSettings)
			delete(rawMsg, key)
		case "inheritableSettings":
			err = unpopulate(val, "InheritableSettings", &r.InheritableSettings)
			delete(rawMsg, key)
		case "level":
			err = unpopulate(val, "Level", &r.Level)
			delete(rawMsg, key)
		case "operations":
			err = unpopulate(val, "Operations", &r.Operations)
			delete(rawMsg, key)
		case "targetObjects":
			err = unpopulate(val, "TargetObjects", &r.TargetObjects)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type UserSet.
func (u UserSet) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "description", u.Description)
	populate(objectMap, "id", u.ID)
	populate(objectMap, "isBackup", u.IsBackup)
	populate(objectMap, "userType", u.UserType)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type UserSet.
func (u *UserSet) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", u, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "description":
			err = unpopulate(val, "Description", &u.Description)
			delete(rawMsg, key)
		case "id":
			err = unpopulate(val, "ID", &u.ID)
			delete(rawMsg, key)
		case "isBackup":
			err = unpopulate(val, "IsBackup", &u.IsBackup)
			delete(rawMsg, key)
		case "userType":
			err = unpopulate(val, "UserType", &u.UserType)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", u, err)
		}
	}
	return nil
}

func populate(m map[string]any, k string, v any) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, fn string, v any) error {
	if data == nil {
		return nil
	}
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("struct field %s: %v", fn, err)
	}
	return nil
}
