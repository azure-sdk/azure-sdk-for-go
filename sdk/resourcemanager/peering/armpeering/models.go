//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpeering

import "time"

// BandwidthOffer - The properties that define a peering bandwidth offer.
type BandwidthOffer struct {
	// The name of the bandwidth offer.
	OfferName *string

	// The value of the bandwidth offer in Mbps.
	ValueInMbps *int32
}

// BgpSession - The properties that define a BGP session.
type BgpSession struct {
	// The MD5 authentication key of the session.
	MD5AuthenticationKey *string

	// The maximum number of prefixes advertised over the IPv4 session.
	MaxPrefixesAdvertisedV4 *int32

	// The maximum number of prefixes advertised over the IPv6 session.
	MaxPrefixesAdvertisedV6 *int32

	// The IPv4 session address on Microsoft's end.
	MicrosoftSessionIPv4Address *string

	// The IPv6 session address on Microsoft's end.
	MicrosoftSessionIPv6Address *string

	// The IPv4 session address on peer's end.
	PeerSessionIPv4Address *string

	// The IPv6 session address on peer's end.
	PeerSessionIPv6Address *string

	// The IPv4 prefix that contains both ends' IPv4 addresses.
	SessionPrefixV4 *string

	// The IPv6 prefix that contains both ends' IPv6 addresses.
	SessionPrefixV6 *string

	// READ-ONLY; The state of the IPv4 session.
	SessionStateV4 *SessionStateV4

	// READ-ONLY; The state of the IPv6 session.
	SessionStateV6 *SessionStateV6
}

// CdnPeeringPrefix - The CDN peering prefix
type CdnPeeringPrefix struct {
	// The properties that define a cdn peering prefix.
	Properties *CdnPeeringPrefixProperties

	// READ-ONLY; The ID of the resource.
	ID *string

	// READ-ONLY; The name of the resource.
	Name *string

	// READ-ONLY; The type of the resource.
	Type *string
}

// CdnPeeringPrefixListResult - The paginated list of CDN peering prefixes.
type CdnPeeringPrefixListResult struct {
	// The link to fetch the next page of CDN peering prefixes.
	NextLink *string

	// The list of CDN peering prefixes.
	Value []*CdnPeeringPrefix
}

// CdnPeeringPrefixProperties - The properties that define a CDN peering prefix
type CdnPeeringPrefixProperties struct {
	// READ-ONLY; The Azure region.
	AzureRegion *string

	// READ-ONLY; The Azure service.
	AzureService *string

	// READ-ONLY; The BGP Community
	BgpCommunity *string

	// READ-ONLY; The flag that indicates whether or not this is the primary region.
	IsPrimaryRegion *bool

	// READ-ONLY; The prefix.
	Prefix *string
}

// CheckServiceProviderAvailabilityInput - Class for CheckServiceProviderAvailabilityInput
type CheckServiceProviderAvailabilityInput struct {
	// Gets or sets the peering service location.
	PeeringServiceLocation *string

	// Gets or sets the peering service provider.
	PeeringServiceProvider *string
}

// ConnectionMonitorTest - The Connection Monitor Test class.
type ConnectionMonitorTest struct {
	// The properties that define a Connection Monitor Test.
	Properties *ConnectionMonitorTestProperties

	// READ-ONLY; The ID of the resource.
	ID *string

	// READ-ONLY; The name of the resource.
	Name *string

	// READ-ONLY; The type of the resource.
	Type *string
}

// ConnectionMonitorTestListResult - The paginated list of Connection Monitor Tests.
type ConnectionMonitorTestListResult struct {
	// The link to fetch the next page of Connection Monitor Tests.
	NextLink *string

	// The list of Connection Monitor Tests.
	Value []*ConnectionMonitorTest
}

// ConnectionMonitorTestProperties - The properties that define a Connection Monitor Test.
type ConnectionMonitorTestProperties struct {
	// The Connection Monitor test destination
	Destination *string

	// The Connection Monitor test destination port
	DestinationPort *int32

	// The Connection Monitor test source agent
	SourceAgent *string

	// The Connection Monitor test frequency in seconds
	TestFrequencyInSec *int32

	// READ-ONLY; The flag that indicates if the Connection Monitor test is successful or not.
	IsTestSuccessful *bool

	// READ-ONLY; The path representing the Connection Monitor test.
	Path []*string

	// READ-ONLY; The provisioning state of the resource.
	ProvisioningState *ProvisioningState
}

// ContactDetail - The contact detail class.
type ContactDetail struct {
	// The e-mail address of the contact.
	Email *string

	// The phone number of the contact.
	Phone *string

	// The role of the contact.
	Role *Role
}

// DirectConnection - The properties that define a direct connection.
type DirectConnection struct {
	// The bandwidth of the connection.
	BandwidthInMbps *int32

	// The BGP session associated with the connection.
	BgpSession *BgpSession

	// The unique identifier (GUID) for the connection.
	ConnectionIdentifier *string

	// The PeeringDB.com ID of the facility at which the connection has to be set up.
	PeeringDBFacilityID *int32

	// The field indicating if Microsoft provides session ip addresses.
	SessionAddressProvider *SessionAddressProvider

	// The flag that indicates whether or not the connection is used for peering service.
	UseForPeeringService *bool

	// READ-ONLY; The state of the connection.
	ConnectionState *ConnectionState

	// READ-ONLY; The error message related to the connection state, if any.
	ErrorMessage *string

	// READ-ONLY; The ID used within Microsoft's peering provisioning system to track the connection
	MicrosoftTrackingID *string

	// READ-ONLY; The bandwidth that is actually provisioned.
	ProvisionedBandwidthInMbps *int32
}

// DirectPeeringFacility - The properties that define a direct peering facility.
type DirectPeeringFacility struct {
	// The address of the direct peering facility.
	Address *string

	// The type of the direct peering.
	DirectPeeringType *DirectPeeringType

	// The PeeringDB.com ID of the facility.
	PeeringDBFacilityID *int32

	// The PeeringDB.com URL of the facility.
	PeeringDBFacilityLink *string
}

// ErrorDetail - The error detail that describes why an operation has failed.
type ErrorDetail struct {
	// READ-ONLY; The error code.
	Code *string

	// READ-ONLY; The error message.
	Message *string
}

// ErrorResponse - The error response that indicates why an operation has failed.
type ErrorResponse struct {
	// The error detail that describes why an operation has failed.
	Error *ErrorDetail
}

// ExchangeConnection - The properties that define an exchange connection.
type ExchangeConnection struct {
	// The BGP session associated with the connection.
	BgpSession *BgpSession

	// The unique identifier (GUID) for the connection.
	ConnectionIdentifier *string

	// The PeeringDB.com ID of the facility at which the connection has to be set up.
	PeeringDBFacilityID *int32

	// READ-ONLY; The state of the connection.
	ConnectionState *ConnectionState

	// READ-ONLY; The error message related to the connection state, if any.
	ErrorMessage *string
}

// ExchangePeeringFacility - The properties that define an exchange peering facility.
type ExchangePeeringFacility struct {
	// The bandwidth of the connection between Microsoft and the exchange peering facility.
	BandwidthInMbps *int32

	// The name of the exchange peering facility.
	ExchangeName *string

	// The IPv4 prefixes associated with the exchange peering facility.
	FacilityIPv4Prefix *string

	// The IPv6 prefixes associated with the exchange peering facility.
	FacilityIPv6Prefix *string

	// The IPv4 address of Microsoft at the exchange peering facility.
	MicrosoftIPv4Address *string

	// The IPv6 address of Microsoft at the exchange peering facility.
	MicrosoftIPv6Address *string

	// The PeeringDB.com ID of the facility.
	PeeringDBFacilityID *int32

	// The PeeringDB.com URL of the facility.
	PeeringDBFacilityLink *string
}

// ListResult - The paginated list of peerings.
type ListResult struct {
	// The link to fetch the next page of peerings.
	NextLink *string

	// The list of peerings.
	Value []*Peering
}

// Location - Peering location is where connectivity could be established to the Microsoft Cloud Edge.
type Location struct {
	// The kind of peering that the peering location supports.
	Kind *Kind

	// The properties that define a peering location.
	Properties *LocationProperties

	// READ-ONLY; The ID of the resource.
	ID *string

	// READ-ONLY; The name of the resource.
	Name *string

	// READ-ONLY; The type of the resource.
	Type *string
}

// LocationListResult - The paginated list of peering locations.
type LocationListResult struct {
	// The link to fetch the next page of peering locations.
	NextLink *string

	// The list of peering locations.
	Value []*Location
}

// LocationProperties - The properties that define a peering location.
type LocationProperties struct {
	// The Azure region associated with the peering location.
	AzureRegion *string

	// The country in which the peering location exists.
	Country *string

	// The properties that define a direct peering location.
	Direct *LocationPropertiesDirect

	// The properties that define an exchange peering location.
	Exchange *LocationPropertiesExchange

	// The name of the peering location.
	PeeringLocation *string
}

// LocationPropertiesDirect - The properties that define a direct peering location.
type LocationPropertiesDirect struct {
	// The list of bandwidth offers available at the peering location.
	BandwidthOffers []*BandwidthOffer

	// The list of direct peering facilities at the peering location.
	PeeringFacilities []*DirectPeeringFacility
}

// LocationPropertiesExchange - The properties that define an exchange peering location.
type LocationPropertiesExchange struct {
	// The list of exchange peering facilities at the peering location.
	PeeringFacilities []*ExchangePeeringFacility
}

// LogAnalyticsWorkspaceProperties - The properties that define a Log Analytics Workspace.
type LogAnalyticsWorkspaceProperties struct {
	// READ-ONLY; The list of connected agents.
	ConnectedAgents []*string

	// READ-ONLY; The Workspace Key.
	Key *string

	// READ-ONLY; The Workspace ID.
	WorkspaceID *string
}

// LookingGlassOutput - Looking glass output model
type LookingGlassOutput struct {
	// Invoked command
	Command *Command

	// Output of the command
	Output *string
}

// MetricDimension - Dimensions of the metric.
type MetricDimension struct {
	// READ-ONLY; Localized friendly display name of the dimension.
	DisplayName *string

	// READ-ONLY; Name of the dimension.
	Name *string
}

// MetricSpecification - Specifications of the Metrics for Azure Monitoring.
type MetricSpecification struct {
	// READ-ONLY; Aggregation type will be set to one of the values: Average, Minimum, Maximum, Total, Count.
	AggregationType *string

	// READ-ONLY; Dimensions of the metric.
	Dimensions []*MetricDimension

	// READ-ONLY; Localized friendly description of the metric.
	DisplayDescription *string

	// READ-ONLY; Localized friendly display name of the metric.
	DisplayName *string

	// READ-ONLY; Name of the metric.
	Name *string

	// READ-ONLY; Supported time grain types for the metric.
	SupportedTimeGrainTypes []*string

	// READ-ONLY; Unit that makes sense for the metric.
	Unit *string
}

// Operation - The peering API operation.
type Operation struct {
	// READ-ONLY; The information related to the operation.
	Display *OperationDisplayInfo

	// READ-ONLY; The flag that indicates whether the operation applies to data plane.
	IsDataAction *bool

	// READ-ONLY; The name of the operation.
	Name *string

	// READ-ONLY; The properties of the operation.
	Properties *OperationProperties
}

// OperationDisplayInfo - The information related to the operation.
type OperationDisplayInfo struct {
	// READ-ONLY; The description of the operation.
	Description *string

	// READ-ONLY; The name of the operation.
	Operation *string

	// READ-ONLY; The name of the resource provider.
	Provider *string

	// READ-ONLY; The type of the resource.
	Resource *string
}

// OperationListResult - The paginated list of peering API operations.
type OperationListResult struct {
	// The link to fetch the next page of peering API operations.
	NextLink *string

	// The list of peering API operations.
	Value []*Operation
}

// OperationProperties - The properties of the operation.
type OperationProperties struct {
	// READ-ONLY; Service specification payload.
	ServiceSpecification *ServiceSpecification
}

// PeerAsn - The essential information related to the peer's ASN.
type PeerAsn struct {
	// The properties that define a peer's ASN.
	Properties *PeerAsnProperties

	// READ-ONLY; The ID of the resource.
	ID *string

	// READ-ONLY; The name of the resource.
	Name *string

	// READ-ONLY; The type of the resource.
	Type *string
}

// PeerAsnListResult - The paginated list of peer ASNs.
type PeerAsnListResult struct {
	// The link to fetch the next page of peer ASNs.
	NextLink *string

	// The list of peer ASNs.
	Value []*PeerAsn
}

// PeerAsnProperties - The properties that define a peer's ASN.
type PeerAsnProperties struct {
	// The Autonomous System Number (ASN) of the peer.
	PeerAsn *int32

	// The contact details of the peer.
	PeerContactDetail []*ContactDetail

	// The name of the peer.
	PeerName *string

	// READ-ONLY; The error message for the validation state
	ErrorMessage *string

	// READ-ONLY; The validation state of the ASN associated with the peer.
	ValidationState *ValidationState
}

// Peering is a logical representation of a set of connections to the Microsoft Cloud Edge at a location.
type Peering struct {
	// REQUIRED; The kind of the peering.
	Kind *Kind

	// REQUIRED; The location of the resource.
	Location *string

	// REQUIRED; The SKU that defines the tier and kind of the peering.
	SKU *SKU

	// The properties that define a peering.
	Properties *Properties

	// The resource tags.
	Tags map[string]*string

	// READ-ONLY; The ID of the resource.
	ID *string

	// READ-ONLY; The name of the resource.
	Name *string

	// READ-ONLY; The type of the resource.
	Type *string
}

// Properties - The properties that define connectivity to the Microsoft Cloud Edge.
type Properties struct {
	// The properties that define a direct peering.
	Direct *PropertiesDirect

	// The properties that define an exchange peering.
	Exchange *PropertiesExchange

	// The location of the peering.
	PeeringLocation *string

	// READ-ONLY; The provisioning state of the resource.
	ProvisioningState *ProvisioningState
}

// PropertiesDirect - The properties that define a direct peering.
type PropertiesDirect struct {
	// The set of connections that constitute a direct peering.
	Connections []*DirectConnection

	// The type of direct peering.
	DirectPeeringType *DirectPeeringType

	// The reference of the peer ASN.
	PeerAsn *SubResource

	// READ-ONLY; The flag that indicates whether or not the peering is used for peering service.
	UseForPeeringService *bool
}

// PropertiesExchange - The properties that define an exchange peering.
type PropertiesExchange struct {
	// The set of connections that constitute an exchange peering.
	Connections []*ExchangeConnection

	// The reference of the peer ASN.
	PeerAsn *SubResource
}

// ReceivedRoute - The properties that define a received route.
type ReceivedRoute struct {
	// READ-ONLY; The AS path for the prefix.
	AsPath *string

	// READ-ONLY; The next hop for the prefix.
	NextHop *string

	// READ-ONLY; The origin AS change information for the prefix.
	OriginAsValidationState *string

	// READ-ONLY; The prefix.
	Prefix *string

	// READ-ONLY; The received timestamp associated with the prefix.
	ReceivedTimestamp *string

	// READ-ONLY; The RPKI validation state for the prefix and origin AS that's listed in the AS path.
	RpkiValidationState *string

	// READ-ONLY; The authority which holds the Route Origin Authorization record for the prefix, if any.
	TrustAnchor *string
}

// ReceivedRouteListResult - The paginated list of received routes for the peering.
type ReceivedRouteListResult struct {
	// The link to fetch the next page of received routes for the peering.
	NextLink *string

	// The list of received routes for the peering.
	Value []*ReceivedRoute
}

// RegisteredAsn - The customer's ASN that is registered by the peering service provider.
type RegisteredAsn struct {
	// The properties that define a registered ASN.
	Properties *RegisteredAsnProperties

	// READ-ONLY; The ID of the resource.
	ID *string

	// READ-ONLY; The name of the resource.
	Name *string

	// READ-ONLY; The type of the resource.
	Type *string
}

// RegisteredAsnListResult - The paginated list of peering registered ASNs.
type RegisteredAsnListResult struct {
	// The link to fetch the next page of peering registered ASNs.
	NextLink *string

	// The list of peering registered ASNs.
	Value []*RegisteredAsn
}

// RegisteredAsnProperties - The properties that define a registered ASN.
type RegisteredAsnProperties struct {
	// The customer's ASN from which traffic originates.
	Asn *int32

	// READ-ONLY; The peering service prefix key that is to be shared with the customer.
	PeeringServicePrefixKey *string

	// READ-ONLY; The provisioning state of the resource.
	ProvisioningState *ProvisioningState
}

// RegisteredPrefix - The customer's prefix that is registered by the peering service provider.
type RegisteredPrefix struct {
	// The properties that define a registered prefix.
	Properties *RegisteredPrefixProperties

	// READ-ONLY; The ID of the resource.
	ID *string

	// READ-ONLY; The name of the resource.
	Name *string

	// READ-ONLY; The type of the resource.
	Type *string
}

// RegisteredPrefixListResult - The paginated list of peering registered prefixes.
type RegisteredPrefixListResult struct {
	// The link to fetch the next page of peering registered prefixes.
	NextLink *string

	// The list of peering registered prefixes.
	Value []*RegisteredPrefix
}

// RegisteredPrefixProperties - The properties that define a registered prefix.
type RegisteredPrefixProperties struct {
	// The customer's prefix from which traffic originates.
	Prefix *string

	// READ-ONLY; The error message associated with the validation state, if any.
	ErrorMessage *string

	// READ-ONLY; The peering service prefix key that is to be shared with the customer.
	PeeringServicePrefixKey *string

	// READ-ONLY; The prefix validation state.
	PrefixValidationState *PrefixValidationState

	// READ-ONLY; The provisioning state of the resource.
	ProvisioningState *ProvisioningState
}

// Resource - The ARM resource class.
type Resource struct {
	// READ-ONLY; The ID of the resource.
	ID *string

	// READ-ONLY; The name of the resource.
	Name *string

	// READ-ONLY; The type of the resource.
	Type *string
}

// ResourceTags - The resource tags.
type ResourceTags struct {
	// Gets or sets the tags, a dictionary of descriptors arm object
	Tags map[string]*string
}

// RpUnbilledPrefix - The Routing Preference unbilled prefix
type RpUnbilledPrefix struct {
	// READ-ONLY; The Azure region.
	AzureRegion *string

	// READ-ONLY; The peer ASN.
	PeerAsn *int32

	// READ-ONLY; The prefix.
	Prefix *string
}

// RpUnbilledPrefixListResult - The paginated list of RP unbilled prefixes.
type RpUnbilledPrefixListResult struct {
	// The link to fetch the next page of RP unbilled prefixes.
	NextLink *string

	// The list of RP unbilled prefixes.
	Value []*RpUnbilledPrefix
}

// SKU - The SKU that defines the tier and kind of the peering.
type SKU struct {
	// The name of the peering SKU.
	Name *string

	// READ-ONLY; The family of the peering SKU.
	Family *Family

	// READ-ONLY; The size of the peering SKU.
	Size *Size

	// READ-ONLY; The tier of the peering SKU.
	Tier *Tier
}

// Service - Peering Service
type Service struct {
	// REQUIRED; The location of the resource.
	Location *string

	// The properties that define a peering service.
	Properties *ServiceProperties

	// The SKU that defines the type of the peering service.
	SKU *ServiceSKU

	// The resource tags.
	Tags map[string]*string

	// READ-ONLY; The ID of the resource.
	ID *string

	// READ-ONLY; The name of the resource.
	Name *string

	// READ-ONLY; The type of the resource.
	Type *string
}

// ServiceCountry - The peering service country.
type ServiceCountry struct {
	// READ-ONLY; The ID of the resource.
	ID *string

	// READ-ONLY; The name of the resource.
	Name *string

	// READ-ONLY; The type of the resource.
	Type *string
}

// ServiceCountryListResult - The paginated list of peering service countries.
type ServiceCountryListResult struct {
	// The link to fetch the next page of peering service countries.
	NextLink *string

	// The list of peering service countries.
	Value []*ServiceCountry
}

// ServiceListResult - The paginated list of peering services.
type ServiceListResult struct {
	// The link to fetch the next page of peering services.
	NextLink *string

	// The list of peering services.
	Value []*Service
}

// ServiceLocation - The peering service location.
type ServiceLocation struct {
	// The properties that define a peering service location.
	Properties *ServiceLocationProperties

	// READ-ONLY; The ID of the resource.
	ID *string

	// READ-ONLY; The name of the resource.
	Name *string

	// READ-ONLY; The type of the resource.
	Type *string
}

// ServiceLocationListResult - The paginated list of peering service locations.
type ServiceLocationListResult struct {
	// The link to fetch the next page of peering service locations.
	NextLink *string

	// The list of peering service locations.
	Value []*ServiceLocation
}

// ServiceLocationProperties - The properties that define connectivity to the Peering Service Location.
type ServiceLocationProperties struct {
	// Azure region for the location
	AzureRegion *string

	// Country of the customer
	Country *string

	// State of the customer
	State *string
}

// ServicePrefix - The peering service prefix class.
type ServicePrefix struct {
	// Gets or sets the peering prefix properties.
	Properties *ServicePrefixProperties

	// READ-ONLY; The ID of the resource.
	ID *string

	// READ-ONLY; The name of the resource.
	Name *string

	// READ-ONLY; The type of the resource.
	Type *string
}

// ServicePrefixEvent - The details of the event associated with a prefix.
type ServicePrefixEvent struct {
	// READ-ONLY; The description of the event associated with a prefix.
	EventDescription *string

	// READ-ONLY; The level of the event associated with a prefix.
	EventLevel *string

	// READ-ONLY; The summary of the event associated with a prefix.
	EventSummary *string

	// READ-ONLY; The timestamp of the event associated with a prefix.
	EventTimestamp *time.Time

	// READ-ONLY; The type of the event associated with a prefix.
	EventType *string
}

// ServicePrefixListResult - The paginated list of peering service prefixes.
type ServicePrefixListResult struct {
	// The link to fetch the next page of peering service prefixes.
	NextLink *string

	// The list of peering service prefixes.
	Value []*ServicePrefix
}

// ServicePrefixProperties - The peering service prefix properties class.
type ServicePrefixProperties struct {
	// The peering service prefix key
	PeeringServicePrefixKey *string

	// The prefix from which your traffic originates.
	Prefix *string

	// READ-ONLY; The error message for validation state
	ErrorMessage *string

	// READ-ONLY; The list of events for peering service prefix
	Events []*ServicePrefixEvent

	// READ-ONLY; The prefix learned type
	LearnedType *LearnedType

	// READ-ONLY; The prefix validation state
	PrefixValidationState *PrefixValidationState

	// READ-ONLY; The provisioning state of the resource.
	ProvisioningState *ProvisioningState
}

// ServiceProperties - The properties that define connectivity to the Peering Service.
type ServiceProperties struct {
	// The Log Analytics Workspace Properties
	LogAnalyticsWorkspaceProperties *LogAnalyticsWorkspaceProperties

	// The location (state/province) of the customer.
	PeeringServiceLocation *string

	// The name of the service provider.
	PeeringServiceProvider *string

	// The backup peering (Microsoft/service provider) location to be used for customer traffic.
	ProviderBackupPeeringLocation *string

	// The primary peering (Microsoft/service provider) location to be used for customer traffic.
	ProviderPrimaryPeeringLocation *string

	// READ-ONLY; The provisioning state of the resource.
	ProvisioningState *ProvisioningState
}

// ServiceProvider - PeeringService provider
type ServiceProvider struct {
	// The properties that define a peering service provider.
	Properties *ServiceProviderProperties

	// READ-ONLY; The ID of the resource.
	ID *string

	// READ-ONLY; The name of the resource.
	Name *string

	// READ-ONLY; The type of the resource.
	Type *string
}

// ServiceProviderListResult - The paginated list of peering service providers.
type ServiceProviderListResult struct {
	// The link to fetch the next page of peering service providers.
	NextLink *string

	// The list of peering service providers.
	Value []*ServiceProvider
}

// ServiceProviderProperties - The properties that define connectivity to the Peering Service Provider.
type ServiceProviderProperties struct {
	// The list of locations at which the service provider peers with Microsoft.
	PeeringLocations []*string

	// The name of the service provider.
	ServiceProviderName *string
}

// ServiceSKU - The SKU that defines the type of the peering service.
type ServiceSKU struct {
	// The name of the peering service SKU.
	Name *string
}

// ServiceSpecification - Service specification payload.
type ServiceSpecification struct {
	// READ-ONLY; Specifications of the Metrics for Azure Monitoring.
	MetricSpecifications []*MetricSpecification
}

// SubResource - The sub resource.
type SubResource struct {
	// The identifier of the referenced resource.
	ID *string
}
