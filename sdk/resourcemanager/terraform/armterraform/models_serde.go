//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armterraform

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type BaseExportModel.
func (b BaseExportModel) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "fullProperties", b.FullProperties)
	populate(objectMap, "maskSensitive", b.MaskSensitive)
	populate(objectMap, "targetProvider", b.TargetProvider)
	objectMap["type"] = b.Type
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type BaseExportModel.
func (b *BaseExportModel) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", b, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "fullProperties":
			err = unpopulate(val, "FullProperties", &b.FullProperties)
			delete(rawMsg, key)
		case "maskSensitive":
			err = unpopulate(val, "MaskSensitive", &b.MaskSensitive)
			delete(rawMsg, key)
		case "targetProvider":
			err = unpopulate(val, "TargetProvider", &b.TargetProvider)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &b.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", b, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ExportQuery.
func (e ExportQuery) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "fullProperties", e.FullProperties)
	populate(objectMap, "maskSensitive", e.MaskSensitive)
	populate(objectMap, "namePattern", e.NamePattern)
	populate(objectMap, "query", e.Query)
	populate(objectMap, "recursive", e.Recursive)
	populate(objectMap, "targetProvider", e.TargetProvider)
	objectMap["type"] = TypeExportQuery
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ExportQuery.
func (e *ExportQuery) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "fullProperties":
			err = unpopulate(val, "FullProperties", &e.FullProperties)
			delete(rawMsg, key)
		case "maskSensitive":
			err = unpopulate(val, "MaskSensitive", &e.MaskSensitive)
			delete(rawMsg, key)
		case "namePattern":
			err = unpopulate(val, "NamePattern", &e.NamePattern)
			delete(rawMsg, key)
		case "query":
			err = unpopulate(val, "Query", &e.Query)
			delete(rawMsg, key)
		case "recursive":
			err = unpopulate(val, "Recursive", &e.Recursive)
			delete(rawMsg, key)
		case "targetProvider":
			err = unpopulate(val, "TargetProvider", &e.TargetProvider)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &e.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ExportResource.
func (e ExportResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "fullProperties", e.FullProperties)
	populate(objectMap, "maskSensitive", e.MaskSensitive)
	populate(objectMap, "namePattern", e.NamePattern)
	populate(objectMap, "resourceIds", e.ResourceIDs)
	populate(objectMap, "resourceName", e.ResourceName)
	populate(objectMap, "resourceType", e.ResourceType)
	populate(objectMap, "targetProvider", e.TargetProvider)
	objectMap["type"] = TypeExportResource
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ExportResource.
func (e *ExportResource) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "fullProperties":
			err = unpopulate(val, "FullProperties", &e.FullProperties)
			delete(rawMsg, key)
		case "maskSensitive":
			err = unpopulate(val, "MaskSensitive", &e.MaskSensitive)
			delete(rawMsg, key)
		case "namePattern":
			err = unpopulate(val, "NamePattern", &e.NamePattern)
			delete(rawMsg, key)
		case "resourceIds":
			err = unpopulate(val, "ResourceIDs", &e.ResourceIDs)
			delete(rawMsg, key)
		case "resourceName":
			err = unpopulate(val, "ResourceName", &e.ResourceName)
			delete(rawMsg, key)
		case "resourceType":
			err = unpopulate(val, "ResourceType", &e.ResourceType)
			delete(rawMsg, key)
		case "targetProvider":
			err = unpopulate(val, "TargetProvider", &e.TargetProvider)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &e.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ExportResourceGroup.
func (e ExportResourceGroup) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "fullProperties", e.FullProperties)
	populate(objectMap, "maskSensitive", e.MaskSensitive)
	populate(objectMap, "namePattern", e.NamePattern)
	populate(objectMap, "resourceGroupName", e.ResourceGroupName)
	populate(objectMap, "targetProvider", e.TargetProvider)
	objectMap["type"] = TypeExportResourceGroup
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ExportResourceGroup.
func (e *ExportResourceGroup) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "fullProperties":
			err = unpopulate(val, "FullProperties", &e.FullProperties)
			delete(rawMsg, key)
		case "maskSensitive":
			err = unpopulate(val, "MaskSensitive", &e.MaskSensitive)
			delete(rawMsg, key)
		case "namePattern":
			err = unpopulate(val, "NamePattern", &e.NamePattern)
			delete(rawMsg, key)
		case "resourceGroupName":
			err = unpopulate(val, "ResourceGroupName", &e.ResourceGroupName)
			delete(rawMsg, key)
		case "targetProvider":
			err = unpopulate(val, "TargetProvider", &e.TargetProvider)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &e.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
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
