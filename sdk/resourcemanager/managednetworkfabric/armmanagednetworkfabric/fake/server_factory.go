// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
	"sync"
)

// ServerFactory is a fake server for instances of the armmanagednetworkfabric.ClientFactory type.
type ServerFactory struct {
	// AccessControlListsServer contains the fakes for client AccessControlListsClient
	AccessControlListsServer AccessControlListsServer

	// ExternalNetworksServer contains the fakes for client ExternalNetworksClient
	ExternalNetworksServer ExternalNetworksServer

	// IPCommunitiesServer contains the fakes for client IPCommunitiesClient
	IPCommunitiesServer IPCommunitiesServer

	// IPExtendedCommunitiesServer contains the fakes for client IPExtendedCommunitiesClient
	IPExtendedCommunitiesServer IPExtendedCommunitiesServer

	// IPPrefixesServer contains the fakes for client IPPrefixesClient
	IPPrefixesServer IPPrefixesServer

	// InternalNetworksServer contains the fakes for client InternalNetworksClient
	InternalNetworksServer InternalNetworksServer

	// InternetGatewayRulesServer contains the fakes for client InternetGatewayRulesClient
	InternetGatewayRulesServer InternetGatewayRulesServer

	// InternetGatewaysServer contains the fakes for client InternetGatewaysClient
	InternetGatewaysServer InternetGatewaysServer

	// L2IsolationDomainsServer contains the fakes for client L2IsolationDomainsClient
	L2IsolationDomainsServer L2IsolationDomainsServer

	// L3IsolationDomainsServer contains the fakes for client L3IsolationDomainsClient
	L3IsolationDomainsServer L3IsolationDomainsServer

	// NeighborGroupsServer contains the fakes for client NeighborGroupsClient
	NeighborGroupsServer NeighborGroupsServer

	// NetworkDeviceSKUsServer contains the fakes for client NetworkDeviceSKUsClient
	NetworkDeviceSKUsServer NetworkDeviceSKUsServer

	// NetworkDevicesServer contains the fakes for client NetworkDevicesClient
	NetworkDevicesServer NetworkDevicesServer

	// NetworkFabricControllersServer contains the fakes for client NetworkFabricControllersClient
	NetworkFabricControllersServer NetworkFabricControllersServer

	// NetworkFabricSKUsServer contains the fakes for client NetworkFabricSKUsClient
	NetworkFabricSKUsServer NetworkFabricSKUsServer

	// NetworkFabricsServer contains the fakes for client NetworkFabricsClient
	NetworkFabricsServer NetworkFabricsServer

	// NetworkInterfacesServer contains the fakes for client NetworkInterfacesClient
	NetworkInterfacesServer NetworkInterfacesServer

	// NetworkPacketBrokersServer contains the fakes for client NetworkPacketBrokersClient
	NetworkPacketBrokersServer NetworkPacketBrokersServer

	// NetworkRacksServer contains the fakes for client NetworkRacksClient
	NetworkRacksServer NetworkRacksServer

	// NetworkTapRulesServer contains the fakes for client NetworkTapRulesClient
	NetworkTapRulesServer NetworkTapRulesServer

	// NetworkTapsServer contains the fakes for client NetworkTapsClient
	NetworkTapsServer NetworkTapsServer

	// NetworkToNetworkInterconnectsServer contains the fakes for client NetworkToNetworkInterconnectsClient
	NetworkToNetworkInterconnectsServer NetworkToNetworkInterconnectsServer

	// OperationsServer contains the fakes for client OperationsClient
	OperationsServer OperationsServer

	// RoutePoliciesServer contains the fakes for client RoutePoliciesClient
	RoutePoliciesServer RoutePoliciesServer
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of armmanagednetworkfabric.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// ServerFactoryTransport connects instances of armmanagednetworkfabric.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                                   *ServerFactory
	trMu                                  sync.Mutex
	trAccessControlListsServer            *AccessControlListsServerTransport
	trExternalNetworksServer              *ExternalNetworksServerTransport
	trIPCommunitiesServer                 *IPCommunitiesServerTransport
	trIPExtendedCommunitiesServer         *IPExtendedCommunitiesServerTransport
	trIPPrefixesServer                    *IPPrefixesServerTransport
	trInternalNetworksServer              *InternalNetworksServerTransport
	trInternetGatewayRulesServer          *InternetGatewayRulesServerTransport
	trInternetGatewaysServer              *InternetGatewaysServerTransport
	trL2IsolationDomainsServer            *L2IsolationDomainsServerTransport
	trL3IsolationDomainsServer            *L3IsolationDomainsServerTransport
	trNeighborGroupsServer                *NeighborGroupsServerTransport
	trNetworkDeviceSKUsServer             *NetworkDeviceSKUsServerTransport
	trNetworkDevicesServer                *NetworkDevicesServerTransport
	trNetworkFabricControllersServer      *NetworkFabricControllersServerTransport
	trNetworkFabricSKUsServer             *NetworkFabricSKUsServerTransport
	trNetworkFabricsServer                *NetworkFabricsServerTransport
	trNetworkInterfacesServer             *NetworkInterfacesServerTransport
	trNetworkPacketBrokersServer          *NetworkPacketBrokersServerTransport
	trNetworkRacksServer                  *NetworkRacksServerTransport
	trNetworkTapRulesServer               *NetworkTapRulesServerTransport
	trNetworkTapsServer                   *NetworkTapsServerTransport
	trNetworkToNetworkInterconnectsServer *NetworkToNetworkInterconnectsServerTransport
	trOperationsServer                    *OperationsServerTransport
	trRoutePoliciesServer                 *RoutePoliciesServerTransport
}

// Do implements the policy.Transporter interface for ServerFactoryTransport.
func (s *ServerFactoryTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	client := method[:strings.Index(method, ".")]
	var resp *http.Response
	var err error

	switch client {
	case "AccessControlListsClient":
		initServer(s, &s.trAccessControlListsServer, func() *AccessControlListsServerTransport {
			return NewAccessControlListsServerTransport(&s.srv.AccessControlListsServer)
		})
		resp, err = s.trAccessControlListsServer.Do(req)
	case "ExternalNetworksClient":
		initServer(s, &s.trExternalNetworksServer, func() *ExternalNetworksServerTransport {
			return NewExternalNetworksServerTransport(&s.srv.ExternalNetworksServer)
		})
		resp, err = s.trExternalNetworksServer.Do(req)
	case "IPCommunitiesClient":
		initServer(s, &s.trIPCommunitiesServer, func() *IPCommunitiesServerTransport {
			return NewIPCommunitiesServerTransport(&s.srv.IPCommunitiesServer)
		})
		resp, err = s.trIPCommunitiesServer.Do(req)
	case "IPExtendedCommunitiesClient":
		initServer(s, &s.trIPExtendedCommunitiesServer, func() *IPExtendedCommunitiesServerTransport {
			return NewIPExtendedCommunitiesServerTransport(&s.srv.IPExtendedCommunitiesServer)
		})
		resp, err = s.trIPExtendedCommunitiesServer.Do(req)
	case "IPPrefixesClient":
		initServer(s, &s.trIPPrefixesServer, func() *IPPrefixesServerTransport { return NewIPPrefixesServerTransport(&s.srv.IPPrefixesServer) })
		resp, err = s.trIPPrefixesServer.Do(req)
	case "InternalNetworksClient":
		initServer(s, &s.trInternalNetworksServer, func() *InternalNetworksServerTransport {
			return NewInternalNetworksServerTransport(&s.srv.InternalNetworksServer)
		})
		resp, err = s.trInternalNetworksServer.Do(req)
	case "InternetGatewayRulesClient":
		initServer(s, &s.trInternetGatewayRulesServer, func() *InternetGatewayRulesServerTransport {
			return NewInternetGatewayRulesServerTransport(&s.srv.InternetGatewayRulesServer)
		})
		resp, err = s.trInternetGatewayRulesServer.Do(req)
	case "InternetGatewaysClient":
		initServer(s, &s.trInternetGatewaysServer, func() *InternetGatewaysServerTransport {
			return NewInternetGatewaysServerTransport(&s.srv.InternetGatewaysServer)
		})
		resp, err = s.trInternetGatewaysServer.Do(req)
	case "L2IsolationDomainsClient":
		initServer(s, &s.trL2IsolationDomainsServer, func() *L2IsolationDomainsServerTransport {
			return NewL2IsolationDomainsServerTransport(&s.srv.L2IsolationDomainsServer)
		})
		resp, err = s.trL2IsolationDomainsServer.Do(req)
	case "L3IsolationDomainsClient":
		initServer(s, &s.trL3IsolationDomainsServer, func() *L3IsolationDomainsServerTransport {
			return NewL3IsolationDomainsServerTransport(&s.srv.L3IsolationDomainsServer)
		})
		resp, err = s.trL3IsolationDomainsServer.Do(req)
	case "NeighborGroupsClient":
		initServer(s, &s.trNeighborGroupsServer, func() *NeighborGroupsServerTransport {
			return NewNeighborGroupsServerTransport(&s.srv.NeighborGroupsServer)
		})
		resp, err = s.trNeighborGroupsServer.Do(req)
	case "NetworkDeviceSKUsClient":
		initServer(s, &s.trNetworkDeviceSKUsServer, func() *NetworkDeviceSKUsServerTransport {
			return NewNetworkDeviceSKUsServerTransport(&s.srv.NetworkDeviceSKUsServer)
		})
		resp, err = s.trNetworkDeviceSKUsServer.Do(req)
	case "NetworkDevicesClient":
		initServer(s, &s.trNetworkDevicesServer, func() *NetworkDevicesServerTransport {
			return NewNetworkDevicesServerTransport(&s.srv.NetworkDevicesServer)
		})
		resp, err = s.trNetworkDevicesServer.Do(req)
	case "NetworkFabricControllersClient":
		initServer(s, &s.trNetworkFabricControllersServer, func() *NetworkFabricControllersServerTransport {
			return NewNetworkFabricControllersServerTransport(&s.srv.NetworkFabricControllersServer)
		})
		resp, err = s.trNetworkFabricControllersServer.Do(req)
	case "NetworkFabricSKUsClient":
		initServer(s, &s.trNetworkFabricSKUsServer, func() *NetworkFabricSKUsServerTransport {
			return NewNetworkFabricSKUsServerTransport(&s.srv.NetworkFabricSKUsServer)
		})
		resp, err = s.trNetworkFabricSKUsServer.Do(req)
	case "NetworkFabricsClient":
		initServer(s, &s.trNetworkFabricsServer, func() *NetworkFabricsServerTransport {
			return NewNetworkFabricsServerTransport(&s.srv.NetworkFabricsServer)
		})
		resp, err = s.trNetworkFabricsServer.Do(req)
	case "NetworkInterfacesClient":
		initServer(s, &s.trNetworkInterfacesServer, func() *NetworkInterfacesServerTransport {
			return NewNetworkInterfacesServerTransport(&s.srv.NetworkInterfacesServer)
		})
		resp, err = s.trNetworkInterfacesServer.Do(req)
	case "NetworkPacketBrokersClient":
		initServer(s, &s.trNetworkPacketBrokersServer, func() *NetworkPacketBrokersServerTransport {
			return NewNetworkPacketBrokersServerTransport(&s.srv.NetworkPacketBrokersServer)
		})
		resp, err = s.trNetworkPacketBrokersServer.Do(req)
	case "NetworkRacksClient":
		initServer(s, &s.trNetworkRacksServer, func() *NetworkRacksServerTransport { return NewNetworkRacksServerTransport(&s.srv.NetworkRacksServer) })
		resp, err = s.trNetworkRacksServer.Do(req)
	case "NetworkTapRulesClient":
		initServer(s, &s.trNetworkTapRulesServer, func() *NetworkTapRulesServerTransport {
			return NewNetworkTapRulesServerTransport(&s.srv.NetworkTapRulesServer)
		})
		resp, err = s.trNetworkTapRulesServer.Do(req)
	case "NetworkTapsClient":
		initServer(s, &s.trNetworkTapsServer, func() *NetworkTapsServerTransport { return NewNetworkTapsServerTransport(&s.srv.NetworkTapsServer) })
		resp, err = s.trNetworkTapsServer.Do(req)
	case "NetworkToNetworkInterconnectsClient":
		initServer(s, &s.trNetworkToNetworkInterconnectsServer, func() *NetworkToNetworkInterconnectsServerTransport {
			return NewNetworkToNetworkInterconnectsServerTransport(&s.srv.NetworkToNetworkInterconnectsServer)
		})
		resp, err = s.trNetworkToNetworkInterconnectsServer.Do(req)
	case "OperationsClient":
		initServer(s, &s.trOperationsServer, func() *OperationsServerTransport { return NewOperationsServerTransport(&s.srv.OperationsServer) })
		resp, err = s.trOperationsServer.Do(req)
	case "RoutePoliciesClient":
		initServer(s, &s.trRoutePoliciesServer, func() *RoutePoliciesServerTransport {
			return NewRoutePoliciesServerTransport(&s.srv.RoutePoliciesServer)
		})
		resp, err = s.trRoutePoliciesServer.Do(req)
	default:
		err = fmt.Errorf("unhandled client %s", client)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func initServer[T any](s *ServerFactoryTransport, dst **T, src func() *T) {
	s.trMu.Lock()
	if *dst == nil {
		*dst = src()
	}
	s.trMu.Unlock()
}
