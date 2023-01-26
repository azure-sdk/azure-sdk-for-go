//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsubscription

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type ErrorResponse.
func (e ErrorResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "code", e.Code)
	populate(objectMap, "message", e.Message)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ErrorResponse.
func (e *ErrorResponse) UnmarshalJSON(data []byte) error {
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
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ErrorResponseBody.
func (e ErrorResponseBody) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "code", e.Code)
	populate(objectMap, "error", e.Error)
	populate(objectMap, "message", e.Message)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ErrorResponseBody.
func (e *ErrorResponseBody) UnmarshalJSON(data []byte) error {
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
		case "error":
			err = unpopulate(val, "Error", &e.Error)
			delete(rawMsg, key)
		case "message":
			err = unpopulate(val, "Message", &e.Message)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ListResult.
func (l ListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "nextLink", l.NextLink)
	populate(objectMap, "value", l.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ListResult.
func (l *ListResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", l, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "nextLink":
			err = unpopulate(val, "NextLink", &l.NextLink)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &l.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", l, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Location.
func (l Location) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "displayName", l.DisplayName)
	populate(objectMap, "id", l.ID)
	populate(objectMap, "latitude", l.Latitude)
	populate(objectMap, "longitude", l.Longitude)
	populate(objectMap, "name", l.Name)
	populate(objectMap, "subscriptionId", l.SubscriptionID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Location.
func (l *Location) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", l, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "displayName":
			err = unpopulate(val, "DisplayName", &l.DisplayName)
			delete(rawMsg, key)
		case "id":
			err = unpopulate(val, "ID", &l.ID)
			delete(rawMsg, key)
		case "latitude":
			err = unpopulate(val, "Latitude", &l.Latitude)
			delete(rawMsg, key)
		case "longitude":
			err = unpopulate(val, "Longitude", &l.Longitude)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &l.Name)
			delete(rawMsg, key)
		case "subscriptionId":
			err = unpopulate(val, "SubscriptionID", &l.SubscriptionID)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", l, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type LocationListResult.
func (l LocationListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "value", l.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type LocationListResult.
func (l *LocationListResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", l, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "value":
			err = unpopulate(val, "Value", &l.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", l, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Policies.
func (p Policies) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "locationPlacementId", p.LocationPlacementID)
	populate(objectMap, "quotaId", p.QuotaID)
	populate(objectMap, "spendingLimit", p.SpendingLimit)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Policies.
func (p *Policies) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "locationPlacementId":
			err = unpopulate(val, "LocationPlacementID", &p.LocationPlacementID)
			delete(rawMsg, key)
		case "quotaId":
			err = unpopulate(val, "QuotaID", &p.QuotaID)
			delete(rawMsg, key)
		case "spendingLimit":
			err = unpopulate(val, "SpendingLimit", &p.SpendingLimit)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Subscription.
func (s Subscription) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "authorizationSource", s.AuthorizationSource)
	populate(objectMap, "displayName", s.DisplayName)
	populate(objectMap, "id", s.ID)
	populate(objectMap, "state", s.State)
	populate(objectMap, "subscriptionId", s.SubscriptionID)
	populate(objectMap, "subscriptionPolicies", s.SubscriptionPolicies)
	populate(objectMap, "tags", s.Tags)
	populate(objectMap, "tenantId", s.TenantID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Subscription.
func (s *Subscription) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "authorizationSource":
			err = unpopulate(val, "AuthorizationSource", &s.AuthorizationSource)
			delete(rawMsg, key)
		case "displayName":
			err = unpopulate(val, "DisplayName", &s.DisplayName)
			delete(rawMsg, key)
		case "id":
			err = unpopulate(val, "ID", &s.ID)
			delete(rawMsg, key)
		case "state":
			err = unpopulate(val, "State", &s.State)
			delete(rawMsg, key)
		case "subscriptionId":
			err = unpopulate(val, "SubscriptionID", &s.SubscriptionID)
			delete(rawMsg, key)
		case "subscriptionPolicies":
			err = unpopulate(val, "SubscriptionPolicies", &s.SubscriptionPolicies)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, "Tags", &s.Tags)
			delete(rawMsg, key)
		case "tenantId":
			err = unpopulate(val, "TenantID", &s.TenantID)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type TenantIDDescription.
func (t TenantIDDescription) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "country", t.Country)
	populate(objectMap, "countryCode", t.CountryCode)
	populate(objectMap, "defaultDomain", t.DefaultDomain)
	populate(objectMap, "displayName", t.DisplayName)
	populate(objectMap, "domains", t.Domains)
	populate(objectMap, "id", t.ID)
	populate(objectMap, "tenantCategory", t.TenantCategory)
	populate(objectMap, "tenantId", t.TenantID)
	populate(objectMap, "tenantType", t.TenantType)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type TenantIDDescription.
func (t *TenantIDDescription) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", t, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "country":
			err = unpopulate(val, "Country", &t.Country)
			delete(rawMsg, key)
		case "countryCode":
			err = unpopulate(val, "CountryCode", &t.CountryCode)
			delete(rawMsg, key)
		case "defaultDomain":
			err = unpopulate(val, "DefaultDomain", &t.DefaultDomain)
			delete(rawMsg, key)
		case "displayName":
			err = unpopulate(val, "DisplayName", &t.DisplayName)
			delete(rawMsg, key)
		case "domains":
			err = unpopulate(val, "Domains", &t.Domains)
			delete(rawMsg, key)
		case "id":
			err = unpopulate(val, "ID", &t.ID)
			delete(rawMsg, key)
		case "tenantCategory":
			err = unpopulate(val, "TenantCategory", &t.TenantCategory)
			delete(rawMsg, key)
		case "tenantId":
			err = unpopulate(val, "TenantID", &t.TenantID)
			delete(rawMsg, key)
		case "tenantType":
			err = unpopulate(val, "TenantType", &t.TenantType)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", t, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type TenantListResult.
func (t TenantListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "nextLink", t.NextLink)
	populate(objectMap, "value", t.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type TenantListResult.
func (t *TenantListResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", t, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "nextLink":
			err = unpopulate(val, "NextLink", &t.NextLink)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &t.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", t, err)
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
