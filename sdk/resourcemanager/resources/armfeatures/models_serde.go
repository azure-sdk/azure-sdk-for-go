// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armfeatures

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type AuthorizationProfile.
func (a AuthorizationProfile) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populateDateTimeRFC3339(objectMap, "approvedTime", a.ApprovedTime)
	populate(objectMap, "approver", a.Approver)
	populateDateTimeRFC3339(objectMap, "requestedTime", a.RequestedTime)
	populate(objectMap, "requester", a.Requester)
	populate(objectMap, "requesterObjectId", a.RequesterObjectID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type AuthorizationProfile.
func (a *AuthorizationProfile) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", a, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "approvedTime":
			err = unpopulateDateTimeRFC3339(val, "ApprovedTime", &a.ApprovedTime)
			delete(rawMsg, key)
		case "approver":
			err = unpopulate(val, "Approver", &a.Approver)
			delete(rawMsg, key)
		case "requestedTime":
			err = unpopulateDateTimeRFC3339(val, "RequestedTime", &a.RequestedTime)
			delete(rawMsg, key)
		case "requester":
			err = unpopulate(val, "Requester", &a.Requester)
			delete(rawMsg, key)
		case "requesterObjectId":
			err = unpopulate(val, "RequesterObjectID", &a.RequesterObjectID)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", a, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ErrorDefinition.
func (e ErrorDefinition) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "code", e.Code)
	populate(objectMap, "details", e.Details)
	populate(objectMap, "message", e.Message)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ErrorDefinition.
func (e *ErrorDefinition) UnmarshalJSON(data []byte) error {
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
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ErrorResponse.
func (e ErrorResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "error", e.Error)
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
		case "error":
			err = unpopulate(val, "Error", &e.Error)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ProxyResource.
func (p ProxyResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", p.ID)
	populate(objectMap, "name", p.Name)
	populate(objectMap, "type", p.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ProxyResource.
func (p *ProxyResource) UnmarshalJSON(data []byte) error {
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
		case "name":
			err = unpopulate(val, "Name", &p.Name)
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

// MarshalJSON implements the json.Marshaller interface for type SubscriptionFeatureRegistration.
func (s SubscriptionFeatureRegistration) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", s.ID)
	populate(objectMap, "name", s.Name)
	populate(objectMap, "properties", s.Properties)
	populate(objectMap, "type", s.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SubscriptionFeatureRegistration.
func (s *SubscriptionFeatureRegistration) UnmarshalJSON(data []byte) error {
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

// MarshalJSON implements the json.Marshaller interface for type SubscriptionFeatureRegistrationList.
func (s SubscriptionFeatureRegistrationList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "nextLink", s.NextLink)
	populate(objectMap, "value", s.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SubscriptionFeatureRegistrationList.
func (s *SubscriptionFeatureRegistrationList) UnmarshalJSON(data []byte) error {
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

// MarshalJSON implements the json.Marshaller interface for type SubscriptionFeatureRegistrationProperties.
func (s SubscriptionFeatureRegistrationProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "approvalType", s.ApprovalType)
	populate(objectMap, "authorizationProfile", s.AuthorizationProfile)
	populate(objectMap, "description", s.Description)
	populate(objectMap, "displayName", s.DisplayName)
	populate(objectMap, "documentationLink", s.DocumentationLink)
	populate(objectMap, "featureName", s.FeatureName)
	populate(objectMap, "metadata", s.Metadata)
	populate(objectMap, "providerNamespace", s.ProviderNamespace)
	populateDateTimeRFC3339(objectMap, "registrationDate", s.RegistrationDate)
	populateDateTimeRFC3339(objectMap, "releaseDate", s.ReleaseDate)
	populate(objectMap, "shouldFeatureDisplayInPortal", s.ShouldFeatureDisplayInPortal)
	populate(objectMap, "state", s.State)
	populate(objectMap, "subscriptionId", s.SubscriptionID)
	populate(objectMap, "tenantId", s.TenantID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SubscriptionFeatureRegistrationProperties.
func (s *SubscriptionFeatureRegistrationProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "approvalType":
			err = unpopulate(val, "ApprovalType", &s.ApprovalType)
			delete(rawMsg, key)
		case "authorizationProfile":
			err = unpopulate(val, "AuthorizationProfile", &s.AuthorizationProfile)
			delete(rawMsg, key)
		case "description":
			err = unpopulate(val, "Description", &s.Description)
			delete(rawMsg, key)
		case "displayName":
			err = unpopulate(val, "DisplayName", &s.DisplayName)
			delete(rawMsg, key)
		case "documentationLink":
			err = unpopulate(val, "DocumentationLink", &s.DocumentationLink)
			delete(rawMsg, key)
		case "featureName":
			err = unpopulate(val, "FeatureName", &s.FeatureName)
			delete(rawMsg, key)
		case "metadata":
			err = unpopulate(val, "Metadata", &s.Metadata)
			delete(rawMsg, key)
		case "providerNamespace":
			err = unpopulate(val, "ProviderNamespace", &s.ProviderNamespace)
			delete(rawMsg, key)
		case "registrationDate":
			err = unpopulateDateTimeRFC3339(val, "RegistrationDate", &s.RegistrationDate)
			delete(rawMsg, key)
		case "releaseDate":
			err = unpopulateDateTimeRFC3339(val, "ReleaseDate", &s.ReleaseDate)
			delete(rawMsg, key)
		case "shouldFeatureDisplayInPortal":
			err = unpopulate(val, "ShouldFeatureDisplayInPortal", &s.ShouldFeatureDisplayInPortal)
			delete(rawMsg, key)
		case "state":
			err = unpopulate(val, "State", &s.State)
			delete(rawMsg, key)
		case "subscriptionId":
			err = unpopulate(val, "SubscriptionID", &s.SubscriptionID)
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
