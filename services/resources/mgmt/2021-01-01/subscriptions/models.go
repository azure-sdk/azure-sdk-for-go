package subscriptions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2021-01-01/subscriptions"

// AvailabilityZonePeers list of availability zones shared by the subscriptions.
type AvailabilityZonePeers struct {
	// AvailabilityZone - READ-ONLY; The availabilityZone.
	AvailabilityZone *string `json:"availabilityZone,omitempty"`
	// Peers - Details of shared availability zone.
	Peers *[]Peers `json:"peers,omitempty"`
}

// MarshalJSON is the custom marshaler for AvailabilityZonePeers.
func (azp AvailabilityZonePeers) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if azp.Peers != nil {
		objectMap["peers"] = azp.Peers
	}
	return json.Marshal(objectMap)
}

// CheckResourceNameResult resource Name valid if not a reserved word, does not contain a reserved word and
// does not start with a reserved word
type CheckResourceNameResult struct {
	autorest.Response `json:"-"`
	// Name - Name of Resource
	Name *string `json:"name,omitempty"`
	// Type - Type of Resource
	Type *string `json:"type,omitempty"`
	// Status - Is the resource name Allowed or Reserved. Possible values include: 'ResourceNameStatusAllowed', 'ResourceNameStatusReserved'
	Status ResourceNameStatus `json:"status,omitempty"`
}

// CheckZonePeersRequest check zone peers request parameters.
type CheckZonePeersRequest struct {
	// Location - The Microsoft location.
	Location *string `json:"location,omitempty"`
	// SubscriptionIds - The peer Microsoft Azure subscription ID.
	SubscriptionIds *[]string `json:"subscriptionIds,omitempty"`
}

// CheckZonePeersResult result of the Check zone peers operation.
type CheckZonePeersResult struct {
	autorest.Response `json:"-"`
	// SubscriptionID - READ-ONLY; The subscription ID.
	SubscriptionID *string `json:"subscriptionId,omitempty"`
	// Location - the location of the subscription.
	Location *string `json:"location,omitempty"`
	// AvailabilityZonePeers - The Availability Zones shared by the subscriptions.
	AvailabilityZonePeers *[]AvailabilityZonePeers `json:"availabilityZonePeers,omitempty"`
}

// MarshalJSON is the custom marshaler for CheckZonePeersResult.
func (czpr CheckZonePeersResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if czpr.Location != nil {
		objectMap["location"] = czpr.Location
	}
	if czpr.AvailabilityZonePeers != nil {
		objectMap["availabilityZonePeers"] = czpr.AvailabilityZonePeers
	}
	return json.Marshal(objectMap)
}

// CloudError an error response for a resource management request.
type CloudError struct {
	Error *ErrorResponse `json:"error,omitempty"`
}

// ErrorAdditionalInfo the resource management error additional info.
type ErrorAdditionalInfo struct {
	// Type - READ-ONLY; The additional info type.
	Type *string `json:"type,omitempty"`
	// Info - READ-ONLY; The additional info.
	Info interface{} `json:"info,omitempty"`
}

// MarshalJSON is the custom marshaler for ErrorAdditionalInfo.
func (eai ErrorAdditionalInfo) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// ErrorResponse common error response for all Azure Resource Manager APIs to return error details for
// failed operations. (This also follows the OData error response format.)
type ErrorResponse struct {
	// Code - READ-ONLY; The error code.
	Code *string `json:"code,omitempty"`
	// Message - READ-ONLY; The error message.
	Message *string `json:"message,omitempty"`
	// Target - READ-ONLY; The error target.
	Target *string `json:"target,omitempty"`
	// Details - READ-ONLY; The error details.
	Details *[]ErrorResponse `json:"details,omitempty"`
	// AdditionalInfo - READ-ONLY; The error additional info.
	AdditionalInfo *[]ErrorAdditionalInfo `json:"additionalInfo,omitempty"`
}

// MarshalJSON is the custom marshaler for ErrorResponse.
func (er ErrorResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// ListResult subscription list operation response.
type ListResult struct {
	autorest.Response `json:"-"`
	// Value - An array of subscriptions.
	Value *[]Subscription `json:"value,omitempty"`
	// NextLink - The URL to get the next set of results.
	NextLink *string `json:"nextLink,omitempty"`
}

// ListResultIterator provides access to a complete listing of Subscription values.
type ListResultIterator struct {
	i    int
	page ListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *ListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ListResultIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *ListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter ListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter ListResultIterator) Response() ListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter ListResultIterator) Value() Subscription {
	if !iter.page.NotDone() {
		return Subscription{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the ListResultIterator type.
func NewListResultIterator(page ListResultPage) ListResultIterator {
	return ListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (lr ListResult) IsEmpty() bool {
	return lr.Value == nil || len(*lr.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (lr ListResult) hasNextLink() bool {
	return lr.NextLink != nil && len(*lr.NextLink) != 0
}

// listResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (lr ListResult) listResultPreparer(ctx context.Context) (*http.Request, error) {
	if !lr.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(lr.NextLink)))
}

// ListResultPage contains a page of Subscription values.
type ListResultPage struct {
	fn func(context.Context, ListResult) (ListResult, error)
	lr ListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *ListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.lr)
		if err != nil {
			return err
		}
		page.lr = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *ListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page ListResultPage) NotDone() bool {
	return !page.lr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page ListResultPage) Response() ListResult {
	return page.lr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page ListResultPage) Values() []Subscription {
	if page.lr.IsEmpty() {
		return nil
	}
	return *page.lr.Value
}

// Creates a new instance of the ListResultPage type.
func NewListResultPage(cur ListResult, getNextPage func(context.Context, ListResult) (ListResult, error)) ListResultPage {
	return ListResultPage{
		fn: getNextPage,
		lr: cur,
	}
}

// Location location information.
type Location struct {
	// ID - READ-ONLY; The fully qualified ID of the location. For example, /subscriptions/00000000-0000-0000-0000-000000000000/locations/westus.
	ID *string `json:"id,omitempty"`
	// SubscriptionID - READ-ONLY; The subscription ID.
	SubscriptionID *string `json:"subscriptionId,omitempty"`
	// Name - READ-ONLY; The location name.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The location type. Possible values include: 'LocationTypeRegion', 'LocationTypeEdgeZone'
	Type LocationType `json:"type,omitempty"`
	// DisplayName - READ-ONLY; The display name of the location.
	DisplayName *string `json:"displayName,omitempty"`
	// RegionalDisplayName - READ-ONLY; The display name of the location and its region.
	RegionalDisplayName *string `json:"regionalDisplayName,omitempty"`
	// Metadata - Metadata of the location, such as lat/long, paired region, and others.
	Metadata *LocationMetadata `json:"metadata,omitempty"`
}

// MarshalJSON is the custom marshaler for Location.
func (l Location) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if l.Metadata != nil {
		objectMap["metadata"] = l.Metadata
	}
	return json.Marshal(objectMap)
}

// LocationListResult location list operation response.
type LocationListResult struct {
	autorest.Response `json:"-"`
	// Value - An array of locations.
	Value *[]Location `json:"value,omitempty"`
}

// LocationMetadata location metadata information
type LocationMetadata struct {
	// RegionType - READ-ONLY; The type of the region. Possible values include: 'RegionTypePhysical', 'RegionTypeLogical'
	RegionType RegionType `json:"regionType,omitempty"`
	// RegionCategory - READ-ONLY; The category of the region. Possible values include: 'RegionCategoryRecommended', 'RegionCategoryExtended', 'RegionCategoryOther'
	RegionCategory RegionCategory `json:"regionCategory,omitempty"`
	// GeographyGroup - READ-ONLY; The geography group of the location.
	GeographyGroup *string `json:"geographyGroup,omitempty"`
	// Longitude - READ-ONLY; The longitude of the location.
	Longitude *string `json:"longitude,omitempty"`
	// Latitude - READ-ONLY; The latitude of the location.
	Latitude *string `json:"latitude,omitempty"`
	// PhysicalLocation - READ-ONLY; The physical location of the Azure location.
	PhysicalLocation *string `json:"physicalLocation,omitempty"`
	// PairedRegion - The regions paired to this region.
	PairedRegion *[]PairedRegion `json:"pairedRegion,omitempty"`
	// HomeLocation - READ-ONLY; The home location of an edge zone.
	HomeLocation *string `json:"homeLocation,omitempty"`
}

// MarshalJSON is the custom marshaler for LocationMetadata.
func (lm LocationMetadata) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if lm.PairedRegion != nil {
		objectMap["pairedRegion"] = lm.PairedRegion
	}
	return json.Marshal(objectMap)
}

// ManagedByTenant information about a tenant managing the subscription.
type ManagedByTenant struct {
	// TenantID - READ-ONLY; The tenant ID of the managing tenant. This is a GUID.
	TenantID *string `json:"tenantId,omitempty"`
}

// MarshalJSON is the custom marshaler for ManagedByTenant.
func (mbt ManagedByTenant) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// Operation microsoft.Resources operation
type Operation struct {
	// Name - Operation name: {provider}/{resource}/{operation}
	Name *string `json:"name,omitempty"`
	// Display - The object that represents the operation.
	Display *OperationDisplay `json:"display,omitempty"`
}

// OperationDisplay the object that represents the operation.
type OperationDisplay struct {
	// Provider - Service provider: Microsoft.Resources
	Provider *string `json:"provider,omitempty"`
	// Resource - Resource on which the operation is performed: Profile, endpoint, etc.
	Resource *string `json:"resource,omitempty"`
	// Operation - Operation type: Read, write, delete, etc.
	Operation *string `json:"operation,omitempty"`
	// Description - Description of the operation.
	Description *string `json:"description,omitempty"`
}

// OperationListResult result of the request to list Microsoft.Resources operations. It contains a list of
// operations and a URL link to get the next set of results.
type OperationListResult struct {
	// Value - List of Microsoft.Resources operations.
	Value *[]Operation `json:"value,omitempty"`
	// NextLink - URL to get the next set of operation list results if there are any.
	NextLink *string `json:"nextLink,omitempty"`
}

// PairedRegion information regarding paired region.
type PairedRegion struct {
	// Name - READ-ONLY; The name of the paired region.
	Name *string `json:"name,omitempty"`
	// ID - READ-ONLY; The fully qualified ID of the location. For example, /subscriptions/00000000-0000-0000-0000-000000000000/locations/westus.
	ID *string `json:"id,omitempty"`
	// SubscriptionID - READ-ONLY; The subscription ID.
	SubscriptionID *string `json:"subscriptionId,omitempty"`
}

// MarshalJSON is the custom marshaler for PairedRegion.
func (pr PairedRegion) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// Peers information about shared availability zone.
type Peers struct {
	// SubscriptionID - READ-ONLY; The subscription ID.
	SubscriptionID *string `json:"subscriptionId,omitempty"`
	// AvailabilityZone - READ-ONLY; The availabilityZone.
	AvailabilityZone *string `json:"availabilityZone,omitempty"`
}

// MarshalJSON is the custom marshaler for Peers.
func (p Peers) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// Policies subscription policies.
type Policies struct {
	// LocationPlacementID - READ-ONLY; The subscription location placement ID. The ID indicates which regions are visible for a subscription. For example, a subscription with a location placement Id of Public_2014-09-01 has access to Azure public regions.
	LocationPlacementID *string `json:"locationPlacementId,omitempty"`
	// QuotaID - READ-ONLY; The subscription quota ID.
	QuotaID *string `json:"quotaId,omitempty"`
	// SpendingLimit - READ-ONLY; The subscription spending limit. Possible values include: 'SpendingLimitOn', 'SpendingLimitOff', 'SpendingLimitCurrentPeriodOff'
	SpendingLimit SpendingLimit `json:"spendingLimit,omitempty"`
}

// MarshalJSON is the custom marshaler for Policies.
func (p Policies) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// ResourceName name and Type of the Resource
type ResourceName struct {
	// Name - Name of the resource
	Name *string `json:"name,omitempty"`
	// Type - The type of the resource
	Type *string `json:"type,omitempty"`
}

// Subscription subscription information.
type Subscription struct {
	autorest.Response `json:"-"`
	// ID - READ-ONLY; The fully qualified ID for the subscription. For example, /subscriptions/00000000-0000-0000-0000-000000000000.
	ID *string `json:"id,omitempty"`
	// SubscriptionID - READ-ONLY; The subscription ID.
	SubscriptionID *string `json:"subscriptionId,omitempty"`
	// DisplayName - READ-ONLY; The subscription display name.
	DisplayName *string `json:"displayName,omitempty"`
	// TenantID - READ-ONLY; The subscription tenant ID.
	TenantID *string `json:"tenantId,omitempty"`
	// State - READ-ONLY; The subscription state. Possible values are Enabled, Warned, PastDue, Disabled, and Deleted. Possible values include: 'StateEnabled', 'StateWarned', 'StatePastDue', 'StateDisabled', 'StateDeleted'
	State State `json:"state,omitempty"`
	// SubscriptionPolicies - The subscription policies.
	SubscriptionPolicies *Policies `json:"subscriptionPolicies,omitempty"`
	// AuthorizationSource - The authorization source of the request. Valid values are one or more combinations of Legacy, RoleBased, Bypassed, Direct and Management. For example, 'Legacy, RoleBased'.
	AuthorizationSource *string `json:"authorizationSource,omitempty"`
	// ManagedByTenants - An array containing the tenants managing the subscription.
	ManagedByTenants *[]ManagedByTenant `json:"managedByTenants,omitempty"`
	// Tags - The tags attached to the subscription.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for Subscription.
func (s Subscription) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if s.SubscriptionPolicies != nil {
		objectMap["subscriptionPolicies"] = s.SubscriptionPolicies
	}
	if s.AuthorizationSource != nil {
		objectMap["authorizationSource"] = s.AuthorizationSource
	}
	if s.ManagedByTenants != nil {
		objectMap["managedByTenants"] = s.ManagedByTenants
	}
	if s.Tags != nil {
		objectMap["tags"] = s.Tags
	}
	return json.Marshal(objectMap)
}

// TenantIDDescription tenant Id information.
type TenantIDDescription struct {
	// ID - READ-ONLY; The fully qualified ID of the tenant. For example, /tenants/00000000-0000-0000-0000-000000000000.
	ID *string `json:"id,omitempty"`
	// TenantID - READ-ONLY; The tenant ID. For example, 00000000-0000-0000-0000-000000000000.
	TenantID *string `json:"tenantId,omitempty"`
	// TenantCategory - READ-ONLY; Category of the tenant. Possible values include: 'TenantCategoryHome', 'TenantCategoryProjectedBy', 'TenantCategoryManagedBy'
	TenantCategory TenantCategory `json:"tenantCategory,omitempty"`
	// Country - READ-ONLY; Country/region name of the address for the tenant.
	Country *string `json:"country,omitempty"`
	// CountryCode - READ-ONLY; Country/region abbreviation for the tenant.
	CountryCode *string `json:"countryCode,omitempty"`
	// DisplayName - READ-ONLY; The display name of the tenant.
	DisplayName *string `json:"displayName,omitempty"`
	// Domains - READ-ONLY; The list of domains for the tenant.
	Domains *[]string `json:"domains,omitempty"`
	// DefaultDomain - READ-ONLY; The default domain for the tenant.
	DefaultDomain *string `json:"defaultDomain,omitempty"`
	// TenantType - READ-ONLY; The tenant type. Only available for 'Home' tenant category.
	TenantType *string `json:"tenantType,omitempty"`
	// TenantBrandingLogoURL - READ-ONLY; The tenant's branding logo URL. Only available for 'Home' tenant category.
	TenantBrandingLogoURL *string `json:"tenantBrandingLogoUrl,omitempty"`
}

// MarshalJSON is the custom marshaler for TenantIDDescription.
func (tid TenantIDDescription) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// TenantListResult tenant Ids information.
type TenantListResult struct {
	autorest.Response `json:"-"`
	// Value - An array of tenants.
	Value *[]TenantIDDescription `json:"value,omitempty"`
	// NextLink - The URL to use for getting the next set of results.
	NextLink *string `json:"nextLink,omitempty"`
}

// TenantListResultIterator provides access to a complete listing of TenantIDDescription values.
type TenantListResultIterator struct {
	i    int
	page TenantListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *TenantListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TenantListResultIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *TenantListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter TenantListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter TenantListResultIterator) Response() TenantListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter TenantListResultIterator) Value() TenantIDDescription {
	if !iter.page.NotDone() {
		return TenantIDDescription{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the TenantListResultIterator type.
func NewTenantListResultIterator(page TenantListResultPage) TenantListResultIterator {
	return TenantListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (tlr TenantListResult) IsEmpty() bool {
	return tlr.Value == nil || len(*tlr.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (tlr TenantListResult) hasNextLink() bool {
	return tlr.NextLink != nil && len(*tlr.NextLink) != 0
}

// tenantListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (tlr TenantListResult) tenantListResultPreparer(ctx context.Context) (*http.Request, error) {
	if !tlr.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(tlr.NextLink)))
}

// TenantListResultPage contains a page of TenantIDDescription values.
type TenantListResultPage struct {
	fn  func(context.Context, TenantListResult) (TenantListResult, error)
	tlr TenantListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *TenantListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TenantListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.tlr)
		if err != nil {
			return err
		}
		page.tlr = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *TenantListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page TenantListResultPage) NotDone() bool {
	return !page.tlr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page TenantListResultPage) Response() TenantListResult {
	return page.tlr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page TenantListResultPage) Values() []TenantIDDescription {
	if page.tlr.IsEmpty() {
		return nil
	}
	return *page.tlr.Value
}

// Creates a new instance of the TenantListResultPage type.
func NewTenantListResultPage(cur TenantListResult, getNextPage func(context.Context, TenantListResult) (TenantListResult, error)) TenantListResultPage {
	return TenantListResultPage{
		fn:  getNextPage,
		tlr: cur,
	}
}
