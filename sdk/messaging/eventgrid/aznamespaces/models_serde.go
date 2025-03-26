// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package aznamespaces

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type AcknowledgeResult.
func (a AcknowledgeResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "failedLockTokens", a.FailedLockTokens)
	populate(objectMap, "succeededLockTokens", a.SucceededLockTokens)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type AcknowledgeResult.
func (a *AcknowledgeResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", a, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "failedLockTokens":
			err = unpopulate(val, "FailedLockTokens", &a.FailedLockTokens)
			delete(rawMsg, key)
		case "succeededLockTokens":
			err = unpopulate(val, "SucceededLockTokens", &a.SucceededLockTokens)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", a, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type BrokerProperties.
func (b BrokerProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "deliveryCount", b.DeliveryCount)
	populate(objectMap, "lockToken", b.LockToken)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type BrokerProperties.
func (b *BrokerProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", b, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "deliveryCount":
			err = unpopulate(val, "DeliveryCount", &b.DeliveryCount)
			delete(rawMsg, key)
		case "lockToken":
			err = unpopulate(val, "LockToken", &b.LockToken)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", b, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type CloudEvent.
func (c CloudEvent) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populateAny(objectMap, "data", c.Data)
	populateByteArray(objectMap, "data_base64", c.DataBase64, func() any {
		return runtime.EncodeByteArray(c.DataBase64, runtime.Base64StdFormat)
	})
	populate(objectMap, "datacontenttype", c.Datacontenttype)
	populate(objectMap, "dataschema", c.Dataschema)
	populate(objectMap, "id", c.ID)
	populate(objectMap, "source", c.Source)
	populate(objectMap, "specversion", c.Specversion)
	populate(objectMap, "subject", c.Subject)
	populateDateTimeRFC3339(objectMap, "time", c.Time)
	populate(objectMap, "type", c.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type CloudEvent.
func (c *CloudEvent) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "data":
			err = unpopulate(val, "Data", &c.Data)
			delete(rawMsg, key)
		case "data_base64":
			if val != nil && string(val) != "null" {
				err = runtime.DecodeByteArray(string(val), &c.DataBase64, runtime.Base64StdFormat)
			}
			delete(rawMsg, key)
		case "datacontenttype":
			err = unpopulate(val, "Datacontenttype", &c.Datacontenttype)
			delete(rawMsg, key)
		case "dataschema":
			err = unpopulate(val, "Dataschema", &c.Dataschema)
			delete(rawMsg, key)
		case "id":
			err = unpopulate(val, "ID", &c.ID)
			delete(rawMsg, key)
		case "source":
			err = unpopulate(val, "Source", &c.Source)
			delete(rawMsg, key)
		case "specversion":
			err = unpopulate(val, "Specversion", &c.Specversion)
			delete(rawMsg, key)
		case "subject":
			err = unpopulate(val, "Subject", &c.Subject)
			delete(rawMsg, key)
		case "time":
			err = unpopulateDateTimeRFC3339(val, "Time", &c.Time)
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

// MarshalJSON implements the json.Marshaller interface for type Error.
func (e Error) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "code", e.Code)
	populate(objectMap, "details", e.Details)
	populate(objectMap, "innererror", e.Innererror)
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
		case "innererror":
			err = unpopulate(val, "Innererror", &e.Innererror)
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

// MarshalJSON implements the json.Marshaller interface for type FailedLockToken.
func (f FailedLockToken) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "error", f.Error)
	populate(objectMap, "lockToken", f.LockToken)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type FailedLockToken.
func (f *FailedLockToken) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", f, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "error":
			err = unpopulate(val, "Error", &f.Error)
			delete(rawMsg, key)
		case "lockToken":
			err = unpopulate(val, "LockToken", &f.LockToken)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", f, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type InnerError.
func (i InnerError) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "code", i.Code)
	populate(objectMap, "innererror", i.Innererror)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type InnerError.
func (i *InnerError) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", i, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "code":
			err = unpopulate(val, "Code", &i.Code)
			delete(rawMsg, key)
		case "innererror":
			err = unpopulate(val, "Innererror", &i.Innererror)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", i, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ReceiveDetails.
func (r ReceiveDetails) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "brokerProperties", r.BrokerProperties)
	populate(objectMap, "event", r.Event)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ReceiveDetails.
func (r *ReceiveDetails) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "brokerProperties":
			err = unpopulate(val, "BrokerProperties", &r.BrokerProperties)
			delete(rawMsg, key)
		case "event":
			err = unpopulate(val, "Event", &r.Event)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ReceiveResult.
func (r ReceiveResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "value", r.Details)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ReceiveResult.
func (r *ReceiveResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "value":
			err = unpopulate(val, "Details", &r.Details)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type RejectResult.
func (r RejectResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "failedLockTokens", r.FailedLockTokens)
	populate(objectMap, "succeededLockTokens", r.SucceededLockTokens)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RejectResult.
func (r *RejectResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "failedLockTokens":
			err = unpopulate(val, "FailedLockTokens", &r.FailedLockTokens)
			delete(rawMsg, key)
		case "succeededLockTokens":
			err = unpopulate(val, "SucceededLockTokens", &r.SucceededLockTokens)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ReleaseResult.
func (r ReleaseResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "failedLockTokens", r.FailedLockTokens)
	populate(objectMap, "succeededLockTokens", r.SucceededLockTokens)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ReleaseResult.
func (r *ReleaseResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "failedLockTokens":
			err = unpopulate(val, "FailedLockTokens", &r.FailedLockTokens)
			delete(rawMsg, key)
		case "succeededLockTokens":
			err = unpopulate(val, "SucceededLockTokens", &r.SucceededLockTokens)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type RenewLocksResult.
func (r RenewLocksResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "failedLockTokens", r.FailedLockTokens)
	populate(objectMap, "succeededLockTokens", r.SucceededLockTokens)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RenewLocksResult.
func (r *RenewLocksResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "failedLockTokens":
			err = unpopulate(val, "FailedLockTokens", &r.FailedLockTokens)
			delete(rawMsg, key)
		case "succeededLockTokens":
			err = unpopulate(val, "SucceededLockTokens", &r.SucceededLockTokens)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
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

func populateByteArray[T any](m map[string]any, k string, b []T, convert func() any) {
	if azcore.IsNullValue(b) {
		m[k] = nil
	} else if len(b) == 0 {
		return
	} else {
		m[k] = convert()
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
