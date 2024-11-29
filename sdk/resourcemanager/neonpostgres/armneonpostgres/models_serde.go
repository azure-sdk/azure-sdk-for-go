// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armneonpostgres

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type CompanyDetails.
func (c CompanyDetails) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "businessPhone", c.BusinessPhone)
	populate(objectMap, "companyName", c.CompanyName)
	populate(objectMap, "country", c.Country)
	populate(objectMap, "domain", c.Domain)
	populate(objectMap, "numberOfEmployees", c.NumberOfEmployees)
	populate(objectMap, "officeAddress", c.OfficeAddress)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type CompanyDetails.
func (c *CompanyDetails) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "businessPhone":
				err = unpopulate(val, "BusinessPhone", &c.BusinessPhone)
			delete(rawMsg, key)
		case "companyName":
				err = unpopulate(val, "CompanyName", &c.CompanyName)
			delete(rawMsg, key)
		case "country":
				err = unpopulate(val, "Country", &c.Country)
			delete(rawMsg, key)
		case "domain":
				err = unpopulate(val, "Domain", &c.Domain)
			delete(rawMsg, key)
		case "numberOfEmployees":
				err = unpopulate(val, "NumberOfEmployees", &c.NumberOfEmployees)
			delete(rawMsg, key)
		case "officeAddress":
				err = unpopulate(val, "OfficeAddress", &c.OfficeAddress)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type MarketplaceDetails.
func (m MarketplaceDetails) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "offerDetails", m.OfferDetails)
	populate(objectMap, "subscriptionId", m.SubscriptionID)
	populate(objectMap, "subscriptionStatus", m.SubscriptionStatus)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type MarketplaceDetails.
func (m *MarketplaceDetails) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "offerDetails":
				err = unpopulate(val, "OfferDetails", &m.OfferDetails)
			delete(rawMsg, key)
		case "subscriptionId":
				err = unpopulate(val, "SubscriptionID", &m.SubscriptionID)
			delete(rawMsg, key)
		case "subscriptionStatus":
				err = unpopulate(val, "SubscriptionStatus", &m.SubscriptionStatus)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", m, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type OfferDetails.
func (o OfferDetails) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "offerId", o.OfferID)
	populate(objectMap, "planId", o.PlanID)
	populate(objectMap, "planName", o.PlanName)
	populate(objectMap, "publisherId", o.PublisherID)
	populate(objectMap, "termId", o.TermID)
	populate(objectMap, "termUnit", o.TermUnit)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type OfferDetails.
func (o *OfferDetails) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", o, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "offerId":
				err = unpopulate(val, "OfferID", &o.OfferID)
			delete(rawMsg, key)
		case "planId":
				err = unpopulate(val, "PlanID", &o.PlanID)
			delete(rawMsg, key)
		case "planName":
				err = unpopulate(val, "PlanName", &o.PlanName)
			delete(rawMsg, key)
		case "publisherId":
				err = unpopulate(val, "PublisherID", &o.PublisherID)
			delete(rawMsg, key)
		case "termId":
				err = unpopulate(val, "TermID", &o.TermID)
			delete(rawMsg, key)
		case "termUnit":
				err = unpopulate(val, "TermUnit", &o.TermUnit)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", o, err)
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

// MarshalJSON implements the json.Marshaller interface for type OrganizationProperties.
func (o OrganizationProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "companyDetails", o.CompanyDetails)
	populate(objectMap, "marketplaceDetails", o.MarketplaceDetails)
	populate(objectMap, "partnerOrganizationProperties", o.PartnerOrganizationProperties)
	populate(objectMap, "provisioningState", o.ProvisioningState)
	populate(objectMap, "userDetails", o.UserDetails)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type OrganizationProperties.
func (o *OrganizationProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", o, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "companyDetails":
				err = unpopulate(val, "CompanyDetails", &o.CompanyDetails)
			delete(rawMsg, key)
		case "marketplaceDetails":
				err = unpopulate(val, "MarketplaceDetails", &o.MarketplaceDetails)
			delete(rawMsg, key)
		case "partnerOrganizationProperties":
				err = unpopulate(val, "PartnerOrganizationProperties", &o.PartnerOrganizationProperties)
			delete(rawMsg, key)
		case "provisioningState":
				err = unpopulate(val, "ProvisioningState", &o.ProvisioningState)
			delete(rawMsg, key)
		case "userDetails":
				err = unpopulate(val, "UserDetails", &o.UserDetails)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", o, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type OrganizationResource.
func (o OrganizationResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", o.ID)
	populate(objectMap, "location", o.Location)
	populate(objectMap, "name", o.Name)
	populate(objectMap, "properties", o.Properties)
	populate(objectMap, "systemData", o.SystemData)
	populate(objectMap, "tags", o.Tags)
	populate(objectMap, "type", o.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type OrganizationResource.
func (o *OrganizationResource) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", o, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
				err = unpopulate(val, "ID", &o.ID)
			delete(rawMsg, key)
		case "location":
				err = unpopulate(val, "Location", &o.Location)
			delete(rawMsg, key)
		case "name":
				err = unpopulate(val, "Name", &o.Name)
			delete(rawMsg, key)
		case "properties":
				err = unpopulate(val, "Properties", &o.Properties)
			delete(rawMsg, key)
		case "systemData":
				err = unpopulate(val, "SystemData", &o.SystemData)
			delete(rawMsg, key)
		case "tags":
				err = unpopulate(val, "Tags", &o.Tags)
			delete(rawMsg, key)
		case "type":
				err = unpopulate(val, "Type", &o.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", o, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type OrganizationResourceListResult.
func (o OrganizationResourceListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "nextLink", o.NextLink)
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type OrganizationResourceListResult.
func (o *OrganizationResourceListResult) UnmarshalJSON(data []byte) error {
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

// MarshalJSON implements the json.Marshaller interface for type PartnerOrganizationProperties.
func (p PartnerOrganizationProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "organizationId", p.OrganizationID)
	populate(objectMap, "organizationName", p.OrganizationName)
	populate(objectMap, "singleSignOnProperties", p.SingleSignOnProperties)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PartnerOrganizationProperties.
func (p *PartnerOrganizationProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "organizationId":
				err = unpopulate(val, "OrganizationID", &p.OrganizationID)
			delete(rawMsg, key)
		case "organizationName":
				err = unpopulate(val, "OrganizationName", &p.OrganizationName)
			delete(rawMsg, key)
		case "singleSignOnProperties":
				err = unpopulate(val, "SingleSignOnProperties", &p.SingleSignOnProperties)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SingleSignOnProperties.
func (s SingleSignOnProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "aadDomains", s.AADDomains)
	populate(objectMap, "enterpriseAppId", s.EnterpriseAppID)
	populate(objectMap, "singleSignOnState", s.SingleSignOnState)
	populate(objectMap, "singleSignOnUrl", s.SingleSignOnURL)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SingleSignOnProperties.
func (s *SingleSignOnProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "aadDomains":
				err = unpopulate(val, "AADDomains", &s.AADDomains)
			delete(rawMsg, key)
		case "enterpriseAppId":
				err = unpopulate(val, "EnterpriseAppID", &s.EnterpriseAppID)
			delete(rawMsg, key)
		case "singleSignOnState":
				err = unpopulate(val, "SingleSignOnState", &s.SingleSignOnState)
			delete(rawMsg, key)
		case "singleSignOnUrl":
				err = unpopulate(val, "SingleSignOnURL", &s.SingleSignOnURL)
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

// MarshalJSON implements the json.Marshaller interface for type UserDetails.
func (u UserDetails) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "emailAddress", u.EmailAddress)
	populate(objectMap, "firstName", u.FirstName)
	populate(objectMap, "lastName", u.LastName)
	populate(objectMap, "phoneNumber", u.PhoneNumber)
	populate(objectMap, "upn", u.Upn)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type UserDetails.
func (u *UserDetails) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", u, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "emailAddress":
				err = unpopulate(val, "EmailAddress", &u.EmailAddress)
			delete(rawMsg, key)
		case "firstName":
				err = unpopulate(val, "FirstName", &u.FirstName)
			delete(rawMsg, key)
		case "lastName":
				err = unpopulate(val, "LastName", &u.LastName)
			delete(rawMsg, key)
		case "phoneNumber":
				err = unpopulate(val, "PhoneNumber", &u.PhoneNumber)
			delete(rawMsg, key)
		case "upn":
				err = unpopulate(val, "Upn", &u.Upn)
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
	if data == nil || string(data) == "null" {
		return nil
	}
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("struct field %s: %v", fn, err)
	}
	return nil
}

