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

// ServerFactory is a fake server for instances of the armpeering.ClientFactory type.
type ServerFactory struct {
	// CdnPeeringPrefixesServer contains the fakes for client CdnPeeringPrefixesClient
	CdnPeeringPrefixesServer CdnPeeringPrefixesServer

	// ConnectionMonitorTestsServer contains the fakes for client ConnectionMonitorTestsClient
	ConnectionMonitorTestsServer ConnectionMonitorTestsServer

	// LegacyPeeringsServer contains the fakes for client LegacyPeeringsClient
	LegacyPeeringsServer LegacyPeeringsServer

	// LocationsServer contains the fakes for client LocationsClient
	LocationsServer LocationsServer

	// LookingGlassServer contains the fakes for client LookingGlassClient
	LookingGlassServer LookingGlassServer

	// ManagementServer contains the fakes for client ManagementClient
	ManagementServer ManagementServer

	// OperationsServer contains the fakes for client OperationsClient
	OperationsServer OperationsServer

	// PeerAsnsServer contains the fakes for client PeerAsnsClient
	PeerAsnsServer PeerAsnsServer

	// PeeringsServer contains the fakes for client PeeringsClient
	PeeringsServer PeeringsServer

	// PrefixesServer contains the fakes for client PrefixesClient
	PrefixesServer PrefixesServer

	// ReceivedRoutesServer contains the fakes for client ReceivedRoutesClient
	ReceivedRoutesServer ReceivedRoutesServer

	// RegisteredAsnsServer contains the fakes for client RegisteredAsnsClient
	RegisteredAsnsServer RegisteredAsnsServer

	// RegisteredPrefixesServer contains the fakes for client RegisteredPrefixesClient
	RegisteredPrefixesServer RegisteredPrefixesServer

	// RpUnbilledPrefixesServer contains the fakes for client RpUnbilledPrefixesClient
	RpUnbilledPrefixesServer RpUnbilledPrefixesServer

	// ServiceCountriesServer contains the fakes for client ServiceCountriesClient
	ServiceCountriesServer ServiceCountriesServer

	// ServiceLocationsServer contains the fakes for client ServiceLocationsClient
	ServiceLocationsServer ServiceLocationsServer

	// ServiceProvidersServer contains the fakes for client ServiceProvidersClient
	ServiceProvidersServer ServiceProvidersServer

	// ServicesServer contains the fakes for client ServicesClient
	ServicesServer ServicesServer
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of armpeering.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// ServerFactoryTransport connects instances of armpeering.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                            *ServerFactory
	trMu                           sync.Mutex
	trCdnPeeringPrefixesServer     *CdnPeeringPrefixesServerTransport
	trConnectionMonitorTestsServer *ConnectionMonitorTestsServerTransport
	trLegacyPeeringsServer         *LegacyPeeringsServerTransport
	trLocationsServer              *LocationsServerTransport
	trLookingGlassServer           *LookingGlassServerTransport
	trManagementServer             *ManagementServerTransport
	trOperationsServer             *OperationsServerTransport
	trPeerAsnsServer               *PeerAsnsServerTransport
	trPeeringsServer               *PeeringsServerTransport
	trPrefixesServer               *PrefixesServerTransport
	trReceivedRoutesServer         *ReceivedRoutesServerTransport
	trRegisteredAsnsServer         *RegisteredAsnsServerTransport
	trRegisteredPrefixesServer     *RegisteredPrefixesServerTransport
	trRpUnbilledPrefixesServer     *RpUnbilledPrefixesServerTransport
	trServiceCountriesServer       *ServiceCountriesServerTransport
	trServiceLocationsServer       *ServiceLocationsServerTransport
	trServiceProvidersServer       *ServiceProvidersServerTransport
	trServicesServer               *ServicesServerTransport
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
	case "CdnPeeringPrefixesClient":
		initServer(s, &s.trCdnPeeringPrefixesServer, func() *CdnPeeringPrefixesServerTransport {
			return NewCdnPeeringPrefixesServerTransport(&s.srv.CdnPeeringPrefixesServer)
		})
		resp, err = s.trCdnPeeringPrefixesServer.Do(req)
	case "ConnectionMonitorTestsClient":
		initServer(s, &s.trConnectionMonitorTestsServer, func() *ConnectionMonitorTestsServerTransport {
			return NewConnectionMonitorTestsServerTransport(&s.srv.ConnectionMonitorTestsServer)
		})
		resp, err = s.trConnectionMonitorTestsServer.Do(req)
	case "LegacyPeeringsClient":
		initServer(s, &s.trLegacyPeeringsServer, func() *LegacyPeeringsServerTransport {
			return NewLegacyPeeringsServerTransport(&s.srv.LegacyPeeringsServer)
		})
		resp, err = s.trLegacyPeeringsServer.Do(req)
	case "LocationsClient":
		initServer(s, &s.trLocationsServer, func() *LocationsServerTransport { return NewLocationsServerTransport(&s.srv.LocationsServer) })
		resp, err = s.trLocationsServer.Do(req)
	case "LookingGlassClient":
		initServer(s, &s.trLookingGlassServer, func() *LookingGlassServerTransport { return NewLookingGlassServerTransport(&s.srv.LookingGlassServer) })
		resp, err = s.trLookingGlassServer.Do(req)
	case "ManagementClient":
		initServer(s, &s.trManagementServer, func() *ManagementServerTransport { return NewManagementServerTransport(&s.srv.ManagementServer) })
		resp, err = s.trManagementServer.Do(req)
	case "OperationsClient":
		initServer(s, &s.trOperationsServer, func() *OperationsServerTransport { return NewOperationsServerTransport(&s.srv.OperationsServer) })
		resp, err = s.trOperationsServer.Do(req)
	case "PeerAsnsClient":
		initServer(s, &s.trPeerAsnsServer, func() *PeerAsnsServerTransport { return NewPeerAsnsServerTransport(&s.srv.PeerAsnsServer) })
		resp, err = s.trPeerAsnsServer.Do(req)
	case "PeeringsClient":
		initServer(s, &s.trPeeringsServer, func() *PeeringsServerTransport { return NewPeeringsServerTransport(&s.srv.PeeringsServer) })
		resp, err = s.trPeeringsServer.Do(req)
	case "PrefixesClient":
		initServer(s, &s.trPrefixesServer, func() *PrefixesServerTransport { return NewPrefixesServerTransport(&s.srv.PrefixesServer) })
		resp, err = s.trPrefixesServer.Do(req)
	case "ReceivedRoutesClient":
		initServer(s, &s.trReceivedRoutesServer, func() *ReceivedRoutesServerTransport {
			return NewReceivedRoutesServerTransport(&s.srv.ReceivedRoutesServer)
		})
		resp, err = s.trReceivedRoutesServer.Do(req)
	case "RegisteredAsnsClient":
		initServer(s, &s.trRegisteredAsnsServer, func() *RegisteredAsnsServerTransport {
			return NewRegisteredAsnsServerTransport(&s.srv.RegisteredAsnsServer)
		})
		resp, err = s.trRegisteredAsnsServer.Do(req)
	case "RegisteredPrefixesClient":
		initServer(s, &s.trRegisteredPrefixesServer, func() *RegisteredPrefixesServerTransport {
			return NewRegisteredPrefixesServerTransport(&s.srv.RegisteredPrefixesServer)
		})
		resp, err = s.trRegisteredPrefixesServer.Do(req)
	case "RpUnbilledPrefixesClient":
		initServer(s, &s.trRpUnbilledPrefixesServer, func() *RpUnbilledPrefixesServerTransport {
			return NewRpUnbilledPrefixesServerTransport(&s.srv.RpUnbilledPrefixesServer)
		})
		resp, err = s.trRpUnbilledPrefixesServer.Do(req)
	case "ServiceCountriesClient":
		initServer(s, &s.trServiceCountriesServer, func() *ServiceCountriesServerTransport {
			return NewServiceCountriesServerTransport(&s.srv.ServiceCountriesServer)
		})
		resp, err = s.trServiceCountriesServer.Do(req)
	case "ServiceLocationsClient":
		initServer(s, &s.trServiceLocationsServer, func() *ServiceLocationsServerTransport {
			return NewServiceLocationsServerTransport(&s.srv.ServiceLocationsServer)
		})
		resp, err = s.trServiceLocationsServer.Do(req)
	case "ServiceProvidersClient":
		initServer(s, &s.trServiceProvidersServer, func() *ServiceProvidersServerTransport {
			return NewServiceProvidersServerTransport(&s.srv.ServiceProvidersServer)
		})
		resp, err = s.trServiceProvidersServer.Do(req)
	case "ServicesClient":
		initServer(s, &s.trServicesServer, func() *ServicesServerTransport { return NewServicesServerTransport(&s.srv.ServicesServer) })
		resp, err = s.trServicesServer.Do(req)
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
