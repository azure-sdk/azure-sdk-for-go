// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armimpactreporting

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type ClientIncidentDetails.
func (c ClientIncidentDetails) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "clientIncidentId", c.ClientIncidentID)
	populate(objectMap, "clientIncidentSource", c.ClientIncidentSource)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ClientIncidentDetails.
func (c *ClientIncidentDetails) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "clientIncidentId":
			err = unpopulate(val, "ClientIncidentID", &c.ClientIncidentID)
			delete(rawMsg, key)
		case "clientIncidentSource":
			err = unpopulate(val, "ClientIncidentSource", &c.ClientIncidentSource)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Connectivity.
func (c Connectivity) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "port", c.Port)
	populate(objectMap, "protocol", c.Protocol)
	populate(objectMap, "source", c.Source)
	populate(objectMap, "target", c.Target)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Connectivity.
func (c *Connectivity) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "port":
			err = unpopulate(val, "Port", &c.Port)
			delete(rawMsg, key)
		case "protocol":
			err = unpopulate(val, "Protocol", &c.Protocol)
			delete(rawMsg, key)
		case "source":
			err = unpopulate(val, "Source", &c.Source)
			delete(rawMsg, key)
		case "target":
			err = unpopulate(val, "Target", &c.Target)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Connector.
func (c Connector) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", c.ID)
	populate(objectMap, "name", c.Name)
	populate(objectMap, "properties", c.Properties)
	populate(objectMap, "systemData", c.SystemData)
	populate(objectMap, "type", c.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Connector.
func (c *Connector) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &c.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &c.Name)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &c.Properties)
			delete(rawMsg, key)
		case "systemData":
			err = unpopulate(val, "SystemData", &c.SystemData)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &c.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ConnectorListResult.
func (c ConnectorListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "nextLink", c.NextLink)
	populate(objectMap, "value", c.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ConnectorListResult.
func (c *ConnectorListResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "nextLink":
			err = unpopulate(val, "NextLink", &c.NextLink)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &c.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ConnectorProperties.
func (c ConnectorProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "connectorId", c.ConnectorID)
	populate(objectMap, "connectorType", c.ConnectorType)
	populateDateTimeRFC3339(objectMap, "lastRunTimeStamp", c.LastRunTimeStamp)
	populate(objectMap, "provisioningState", c.ProvisioningState)
	populate(objectMap, "tenantId", c.TenantID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ConnectorProperties.
func (c *ConnectorProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "connectorId":
			err = unpopulate(val, "ConnectorID", &c.ConnectorID)
			delete(rawMsg, key)
		case "connectorType":
			err = unpopulate(val, "ConnectorType", &c.ConnectorType)
			delete(rawMsg, key)
		case "lastRunTimeStamp":
			err = unpopulateDateTimeRFC3339(val, "LastRunTimeStamp", &c.LastRunTimeStamp)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &c.ProvisioningState)
			delete(rawMsg, key)
		case "tenantId":
			err = unpopulate(val, "TenantID", &c.TenantID)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ConnectorUpdate.
func (c ConnectorUpdate) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "properties", c.Properties)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ConnectorUpdate.
func (c *ConnectorUpdate) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "properties":
			err = unpopulate(val, "Properties", &c.Properties)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ConnectorUpdateProperties.
func (c ConnectorUpdateProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "connectorType", c.ConnectorType)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ConnectorUpdateProperties.
func (c *ConnectorUpdateProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "connectorType":
			err = unpopulate(val, "ConnectorType", &c.ConnectorType)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Content.
func (c Content) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "description", c.Description)
	populate(objectMap, "title", c.Title)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Content.
func (c *Content) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "description":
			err = unpopulate(val, "Description", &c.Description)
			delete(rawMsg, key)
		case "title":
			err = unpopulate(val, "Title", &c.Title)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ErrorDetailProperties.
func (e ErrorDetailProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "errorCode", e.ErrorCode)
	populate(objectMap, "errorMessage", e.ErrorMessage)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ErrorDetailProperties.
func (e *ErrorDetailProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "errorCode":
			err = unpopulate(val, "ErrorCode", &e.ErrorCode)
			delete(rawMsg, key)
		case "errorMessage":
			err = unpopulate(val, "ErrorMessage", &e.ErrorMessage)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ExpectedValueRange.
func (e ExpectedValueRange) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "max", e.Max)
	populate(objectMap, "min", e.Min)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ExpectedValueRange.
func (e *ExpectedValueRange) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "max":
			err = unpopulate(val, "Max", &e.Max)
			delete(rawMsg, key)
		case "min":
			err = unpopulate(val, "Min", &e.Min)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ImpactCategory.
func (i ImpactCategory) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", i.ID)
	populate(objectMap, "name", i.Name)
	populate(objectMap, "properties", i.Properties)
	populate(objectMap, "systemData", i.SystemData)
	populate(objectMap, "type", i.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ImpactCategory.
func (i *ImpactCategory) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", i, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &i.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &i.Name)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &i.Properties)
			delete(rawMsg, key)
		case "systemData":
			err = unpopulate(val, "SystemData", &i.SystemData)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &i.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", i, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ImpactCategoryListResult.
func (i ImpactCategoryListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "nextLink", i.NextLink)
	populate(objectMap, "value", i.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ImpactCategoryListResult.
func (i *ImpactCategoryListResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", i, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "nextLink":
			err = unpopulate(val, "NextLink", &i.NextLink)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &i.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", i, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ImpactCategoryProperties.
func (i ImpactCategoryProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "categoryId", i.CategoryID)
	populate(objectMap, "description", i.Description)
	populate(objectMap, "parentCategoryId", i.ParentCategoryID)
	populate(objectMap, "provisioningState", i.ProvisioningState)
	populate(objectMap, "requiredImpactProperties", i.RequiredImpactProperties)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ImpactCategoryProperties.
func (i *ImpactCategoryProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", i, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "categoryId":
			err = unpopulate(val, "CategoryID", &i.CategoryID)
			delete(rawMsg, key)
		case "description":
			err = unpopulate(val, "Description", &i.Description)
			delete(rawMsg, key)
		case "parentCategoryId":
			err = unpopulate(val, "ParentCategoryID", &i.ParentCategoryID)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &i.ProvisioningState)
			delete(rawMsg, key)
		case "requiredImpactProperties":
			err = unpopulate(val, "RequiredImpactProperties", &i.RequiredImpactProperties)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", i, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ImpactDetails.
func (i ImpactDetails) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populateDateTimeRFC3339(objectMap, "endTime", i.EndTime)
	populate(objectMap, "impactId", i.ImpactID)
	populate(objectMap, "impactedResourceId", i.ImpactedResourceID)
	populateDateTimeRFC3339(objectMap, "startTime", i.StartTime)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ImpactDetails.
func (i *ImpactDetails) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", i, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "endTime":
			err = unpopulateDateTimeRFC3339(val, "EndTime", &i.EndTime)
			delete(rawMsg, key)
		case "impactId":
			err = unpopulate(val, "ImpactID", &i.ImpactID)
			delete(rawMsg, key)
		case "impactedResourceId":
			err = unpopulate(val, "ImpactedResourceID", &i.ImpactedResourceID)
			delete(rawMsg, key)
		case "startTime":
			err = unpopulateDateTimeRFC3339(val, "StartTime", &i.StartTime)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", i, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Insight.
func (i Insight) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", i.ID)
	populate(objectMap, "name", i.Name)
	populate(objectMap, "properties", i.Properties)
	populate(objectMap, "systemData", i.SystemData)
	populate(objectMap, "type", i.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Insight.
func (i *Insight) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", i, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &i.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &i.Name)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &i.Properties)
			delete(rawMsg, key)
		case "systemData":
			err = unpopulate(val, "SystemData", &i.SystemData)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &i.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", i, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type InsightListResult.
func (i InsightListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "nextLink", i.NextLink)
	populate(objectMap, "value", i.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type InsightListResult.
func (i *InsightListResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", i, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "nextLink":
			err = unpopulate(val, "NextLink", &i.NextLink)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &i.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", i, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type InsightProperties.
func (i InsightProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "additionalDetails", i.AdditionalDetails)
	populate(objectMap, "category", i.Category)
	populate(objectMap, "content", i.Content)
	populate(objectMap, "eventId", i.EventID)
	populateDateTimeRFC3339(objectMap, "eventTime", i.EventTime)
	populate(objectMap, "groupId", i.GroupID)
	populate(objectMap, "impact", i.Impact)
	populate(objectMap, "insightUniqueId", i.InsightUniqueID)
	populate(objectMap, "provisioningState", i.ProvisioningState)
	populate(objectMap, "status", i.Status)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type InsightProperties.
func (i *InsightProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", i, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "additionalDetails":
			err = unpopulate(val, "AdditionalDetails", &i.AdditionalDetails)
			delete(rawMsg, key)
		case "category":
			err = unpopulate(val, "Category", &i.Category)
			delete(rawMsg, key)
		case "content":
			err = unpopulate(val, "Content", &i.Content)
			delete(rawMsg, key)
		case "eventId":
			err = unpopulate(val, "EventID", &i.EventID)
			delete(rawMsg, key)
		case "eventTime":
			err = unpopulateDateTimeRFC3339(val, "EventTime", &i.EventTime)
			delete(rawMsg, key)
		case "groupId":
			err = unpopulate(val, "GroupID", &i.GroupID)
			delete(rawMsg, key)
		case "impact":
			err = unpopulate(val, "Impact", &i.Impact)
			delete(rawMsg, key)
		case "insightUniqueId":
			err = unpopulate(val, "InsightUniqueID", &i.InsightUniqueID)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &i.ProvisioningState)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, "Status", &i.Status)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", i, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Operation.
func (o Operation) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
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
	objectMap := make(map[string]any)
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
	objectMap := make(map[string]any)
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

// MarshalJSON implements the json.Marshaller interface for type Performance.
func (p Performance) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "actual", p.Actual)
	populate(objectMap, "expected", p.Expected)
	populate(objectMap, "expectedValueRange", p.ExpectedValueRange)
	populate(objectMap, "metricName", p.MetricName)
	populate(objectMap, "unit", p.Unit)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Performance.
func (p *Performance) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "actual":
			err = unpopulate(val, "Actual", &p.Actual)
			delete(rawMsg, key)
		case "expected":
			err = unpopulate(val, "Expected", &p.Expected)
			delete(rawMsg, key)
		case "expectedValueRange":
			err = unpopulate(val, "ExpectedValueRange", &p.ExpectedValueRange)
			delete(rawMsg, key)
		case "metricName":
			err = unpopulate(val, "MetricName", &p.MetricName)
			delete(rawMsg, key)
		case "unit":
			err = unpopulate(val, "Unit", &p.Unit)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type RequiredImpactProperties.
func (r RequiredImpactProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "allowedValues", r.AllowedValues)
	populate(objectMap, "name", r.Name)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RequiredImpactProperties.
func (r *RequiredImpactProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "allowedValues":
			err = unpopulate(val, "AllowedValues", &r.AllowedValues)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &r.Name)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SourceOrTarget.
func (s SourceOrTarget) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "azureResourceId", s.AzureResourceID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SourceOrTarget.
func (s *SourceOrTarget) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "azureResourceId":
			err = unpopulate(val, "AzureResourceID", &s.AzureResourceID)
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
	objectMap := make(map[string]any)
	populateDateTimeRFC3339(objectMap, "createdAt", s.CreatedAt)
	populate(objectMap, "createdBy", s.CreatedBy)
	populate(objectMap, "createdByType", s.CreatedByType)
	populateDateTimeRFC3339(objectMap, "lastModifiedAt", s.LastModifiedAt)
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
			err = unpopulateDateTimeRFC3339(val, "CreatedAt", &s.CreatedAt)
			delete(rawMsg, key)
		case "createdBy":
			err = unpopulate(val, "CreatedBy", &s.CreatedBy)
			delete(rawMsg, key)
		case "createdByType":
			err = unpopulate(val, "CreatedByType", &s.CreatedByType)
			delete(rawMsg, key)
		case "lastModifiedAt":
			err = unpopulateDateTimeRFC3339(val, "LastModifiedAt", &s.LastModifiedAt)
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

// MarshalJSON implements the json.Marshaller interface for type Workload.
func (w Workload) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "context", w.Context)
	populate(objectMap, "toolset", w.Toolset)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Workload.
func (w *Workload) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", w, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "context":
			err = unpopulate(val, "Context", &w.Context)
			delete(rawMsg, key)
		case "toolset":
			err = unpopulate(val, "Toolset", &w.Toolset)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", w, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type WorkloadImpact.
func (w WorkloadImpact) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", w.ID)
	populate(objectMap, "name", w.Name)
	populate(objectMap, "properties", w.Properties)
	populate(objectMap, "systemData", w.SystemData)
	populate(objectMap, "type", w.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type WorkloadImpact.
func (w *WorkloadImpact) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", w, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &w.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &w.Name)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &w.Properties)
			delete(rawMsg, key)
		case "systemData":
			err = unpopulate(val, "SystemData", &w.SystemData)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &w.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", w, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type WorkloadImpactListResult.
func (w WorkloadImpactListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "nextLink", w.NextLink)
	populate(objectMap, "value", w.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type WorkloadImpactListResult.
func (w *WorkloadImpactListResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", w, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "nextLink":
			err = unpopulate(val, "NextLink", &w.NextLink)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &w.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", w, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type WorkloadImpactProperties.
func (w WorkloadImpactProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "additionalProperties", w.AdditionalProperties)
	populate(objectMap, "armCorrelationIds", w.ArmCorrelationIDs)
	populate(objectMap, "clientIncidentDetails", w.ClientIncidentDetails)
	populate(objectMap, "confidenceLevel", w.ConfidenceLevel)
	populate(objectMap, "connectivity", w.Connectivity)
	populateDateTimeRFC3339(objectMap, "endDateTime", w.EndDateTime)
	populate(objectMap, "errorDetails", w.ErrorDetails)
	populate(objectMap, "impactCategory", w.ImpactCategory)
	populate(objectMap, "impactDescription", w.ImpactDescription)
	populate(objectMap, "impactGroupId", w.ImpactGroupID)
	populate(objectMap, "impactUniqueId", w.ImpactUniqueID)
	populate(objectMap, "impactedResourceId", w.ImpactedResourceID)
	populate(objectMap, "performance", w.Performance)
	populate(objectMap, "provisioningState", w.ProvisioningState)
	populateDateTimeRFC3339(objectMap, "reportedTimeUtc", w.ReportedTimeUTC)
	populateDateTimeRFC3339(objectMap, "startDateTime", w.StartDateTime)
	populate(objectMap, "workload", w.Workload)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type WorkloadImpactProperties.
func (w *WorkloadImpactProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", w, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "additionalProperties":
			err = unpopulate(val, "AdditionalProperties", &w.AdditionalProperties)
			delete(rawMsg, key)
		case "armCorrelationIds":
			err = unpopulate(val, "ArmCorrelationIDs", &w.ArmCorrelationIDs)
			delete(rawMsg, key)
		case "clientIncidentDetails":
			err = unpopulate(val, "ClientIncidentDetails", &w.ClientIncidentDetails)
			delete(rawMsg, key)
		case "confidenceLevel":
			err = unpopulate(val, "ConfidenceLevel", &w.ConfidenceLevel)
			delete(rawMsg, key)
		case "connectivity":
			err = unpopulate(val, "Connectivity", &w.Connectivity)
			delete(rawMsg, key)
		case "endDateTime":
			err = unpopulateDateTimeRFC3339(val, "EndDateTime", &w.EndDateTime)
			delete(rawMsg, key)
		case "errorDetails":
			err = unpopulate(val, "ErrorDetails", &w.ErrorDetails)
			delete(rawMsg, key)
		case "impactCategory":
			err = unpopulate(val, "ImpactCategory", &w.ImpactCategory)
			delete(rawMsg, key)
		case "impactDescription":
			err = unpopulate(val, "ImpactDescription", &w.ImpactDescription)
			delete(rawMsg, key)
		case "impactGroupId":
			err = unpopulate(val, "ImpactGroupID", &w.ImpactGroupID)
			delete(rawMsg, key)
		case "impactUniqueId":
			err = unpopulate(val, "ImpactUniqueID", &w.ImpactUniqueID)
			delete(rawMsg, key)
		case "impactedResourceId":
			err = unpopulate(val, "ImpactedResourceID", &w.ImpactedResourceID)
			delete(rawMsg, key)
		case "performance":
			err = unpopulate(val, "Performance", &w.Performance)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &w.ProvisioningState)
			delete(rawMsg, key)
		case "reportedTimeUtc":
			err = unpopulateDateTimeRFC3339(val, "ReportedTimeUTC", &w.ReportedTimeUTC)
			delete(rawMsg, key)
		case "startDateTime":
			err = unpopulateDateTimeRFC3339(val, "StartDateTime", &w.StartDateTime)
			delete(rawMsg, key)
		case "workload":
			err = unpopulate(val, "Workload", &w.Workload)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", w, err)
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
	if data == nil || string(data) == "null" {
		return nil
	}
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("struct field %s: %v", fn, err)
	}
	return nil
}
