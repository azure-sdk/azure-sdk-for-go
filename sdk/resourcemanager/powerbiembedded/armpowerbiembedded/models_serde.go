//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpowerbiembedded

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type AzureSKU.
func (a AzureSKU) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "name", a.Name)
	populate(objectMap, "tier", a.Tier)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type AzureSKU.
func (a *AzureSKU) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", a, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "name":
			err = unpopulate(val, "Name", &a.Name)
			delete(rawMsg, key)
		case "tier":
			err = unpopulate(val, "Tier", &a.Tier)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", a, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type CheckNameRequest.
func (c CheckNameRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "name", c.Name)
	populate(objectMap, "type", c.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type CheckNameRequest.
func (c *CheckNameRequest) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "name":
			err = unpopulate(val, "Name", &c.Name)
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

// MarshalJSON implements the json.Marshaller interface for type CheckNameResponse.
func (c CheckNameResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "message", c.Message)
	populate(objectMap, "nameAvailable", c.NameAvailable)
	populate(objectMap, "reason", c.Reason)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type CheckNameResponse.
func (c *CheckNameResponse) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "message":
			err = unpopulate(val, "Message", &c.Message)
			delete(rawMsg, key)
		case "nameAvailable":
			err = unpopulate(val, "NameAvailable", &c.NameAvailable)
			delete(rawMsg, key)
		case "reason":
			err = unpopulate(val, "Reason", &c.Reason)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type CreateWorkspaceCollectionRequest.
func (c CreateWorkspaceCollectionRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "location", c.Location)
	populate(objectMap, "sku", c.SKU)
	populate(objectMap, "tags", c.Tags)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type CreateWorkspaceCollectionRequest.
func (c *CreateWorkspaceCollectionRequest) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "location":
			err = unpopulate(val, "Location", &c.Location)
			delete(rawMsg, key)
		case "sku":
			err = unpopulate(val, "SKU", &c.SKU)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, "Tags", &c.Tags)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Display.
func (d Display) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "description", d.Description)
	populate(objectMap, "operation", d.Operation)
	populate(objectMap, "origin", d.Origin)
	populate(objectMap, "provider", d.Provider)
	populate(objectMap, "resource", d.Resource)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Display.
func (d *Display) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "description":
			err = unpopulate(val, "Description", &d.Description)
			delete(rawMsg, key)
		case "operation":
			err = unpopulate(val, "Operation", &d.Operation)
			delete(rawMsg, key)
		case "origin":
			err = unpopulate(val, "Origin", &d.Origin)
			delete(rawMsg, key)
		case "provider":
			err = unpopulate(val, "Provider", &d.Provider)
			delete(rawMsg, key)
		case "resource":
			err = unpopulate(val, "Resource", &d.Resource)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", d, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Error.
func (e Error) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "code", e.Code)
	populate(objectMap, "details", e.Details)
	populate(objectMap, "message", e.Message)
	populate(objectMap, "target", e.Target)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Error.
func (e *Error) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "code":
			err = unpopulate(val, "Code", &e.Code)
			delete(rawMsg, key)
		case "details":
			err = unpopulate(val, "Details", &e.Details)
			delete(rawMsg, key)
		case "message":
			err = unpopulate(val, "Message", &e.Message)
			delete(rawMsg, key)
		case "target":
			err = unpopulate(val, "Target", &e.Target)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ErrorDetail.
func (e ErrorDetail) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "code", e.Code)
	populate(objectMap, "message", e.Message)
	populate(objectMap, "target", e.Target)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ErrorDetail.
func (e *ErrorDetail) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "code":
			err = unpopulate(val, "Code", &e.Code)
			delete(rawMsg, key)
		case "message":
			err = unpopulate(val, "Message", &e.Message)
			delete(rawMsg, key)
		case "target":
			err = unpopulate(val, "Target", &e.Target)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type MigrateWorkspaceCollectionRequest.
func (m MigrateWorkspaceCollectionRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "resources", m.Resources)
	populate(objectMap, "targetResourceGroup", m.TargetResourceGroup)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type MigrateWorkspaceCollectionRequest.
func (m *MigrateWorkspaceCollectionRequest) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "resources":
			err = unpopulate(val, "Resources", &m.Resources)
			delete(rawMsg, key)
		case "targetResourceGroup":
			err = unpopulate(val, "TargetResourceGroup", &m.TargetResourceGroup)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", m, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Operation.
func (o Operation) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "display", o.Display)
	populate(objectMap, "name", o.Name)
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
		case "display":
			err = unpopulate(val, "Display", &o.Display)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &o.Name)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", o, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type OperationList.
func (o OperationList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type OperationList.
func (o *OperationList) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", o, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
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

// MarshalJSON implements the json.Marshaller interface for type UpdateWorkspaceCollectionRequest.
func (u UpdateWorkspaceCollectionRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "sku", u.SKU)
	populate(objectMap, "tags", u.Tags)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type UpdateWorkspaceCollectionRequest.
func (u *UpdateWorkspaceCollectionRequest) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", u, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "sku":
			err = unpopulate(val, "SKU", &u.SKU)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, "Tags", &u.Tags)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", u, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Workspace.
func (w Workspace) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", w.ID)
	populate(objectMap, "name", w.Name)
	populateAny(objectMap, "properties", w.Properties)
	populate(objectMap, "type", w.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Workspace.
func (w *Workspace) UnmarshalJSON(data []byte) error {
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

// MarshalJSON implements the json.Marshaller interface for type WorkspaceCollection.
func (w WorkspaceCollection) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", w.ID)
	populate(objectMap, "location", w.Location)
	populate(objectMap, "name", w.Name)
	populateAny(objectMap, "properties", w.Properties)
	populate(objectMap, "sku", w.SKU)
	populate(objectMap, "tags", w.Tags)
	populate(objectMap, "type", w.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type WorkspaceCollection.
func (w *WorkspaceCollection) UnmarshalJSON(data []byte) error {
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
		case "location":
			err = unpopulate(val, "Location", &w.Location)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &w.Name)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &w.Properties)
			delete(rawMsg, key)
		case "sku":
			err = unpopulate(val, "SKU", &w.SKU)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, "Tags", &w.Tags)
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

// MarshalJSON implements the json.Marshaller interface for type WorkspaceCollectionAccessKey.
func (w WorkspaceCollectionAccessKey) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "keyName", w.KeyName)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type WorkspaceCollectionAccessKey.
func (w *WorkspaceCollectionAccessKey) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", w, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "keyName":
			err = unpopulate(val, "KeyName", &w.KeyName)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", w, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type WorkspaceCollectionAccessKeys.
func (w WorkspaceCollectionAccessKeys) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "key1", w.Key1)
	populate(objectMap, "key2", w.Key2)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type WorkspaceCollectionAccessKeys.
func (w *WorkspaceCollectionAccessKeys) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", w, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "key1":
			err = unpopulate(val, "Key1", &w.Key1)
			delete(rawMsg, key)
		case "key2":
			err = unpopulate(val, "Key2", &w.Key2)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", w, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type WorkspaceCollectionList.
func (w WorkspaceCollectionList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "value", w.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type WorkspaceCollectionList.
func (w *WorkspaceCollectionList) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", w, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
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

// MarshalJSON implements the json.Marshaller interface for type WorkspaceList.
func (w WorkspaceList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "value", w.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type WorkspaceList.
func (w *WorkspaceList) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", w, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
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

func populate(m map[string]any, k string, v any) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func populateAny(m map[string]any, k string, v any) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else {
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
