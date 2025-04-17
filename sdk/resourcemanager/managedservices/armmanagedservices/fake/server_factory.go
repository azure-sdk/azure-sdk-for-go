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

// ServerFactory is a fake server for instances of the armmanagedservices.ClientFactory type.
type ServerFactory struct {
	// MarketplaceRegistrationDefinitionsServer contains the fakes for client MarketplaceRegistrationDefinitionsClient
	MarketplaceRegistrationDefinitionsServer MarketplaceRegistrationDefinitionsServer

	// MarketplaceRegistrationDefinitionsWithoutScopeServer contains the fakes for client MarketplaceRegistrationDefinitionsWithoutScopeClient
	MarketplaceRegistrationDefinitionsWithoutScopeServer MarketplaceRegistrationDefinitionsWithoutScopeServer

	// OperationsServer contains the fakes for client OperationsClient
	OperationsServer OperationsServer

	// OperationsWithScopeServer contains the fakes for client OperationsWithScopeClient
	OperationsWithScopeServer OperationsWithScopeServer

	// RegistrationAssignmentsServer contains the fakes for client RegistrationAssignmentsClient
	RegistrationAssignmentsServer RegistrationAssignmentsServer

	// RegistrationDefinitionsServer contains the fakes for client RegistrationDefinitionsClient
	RegistrationDefinitionsServer RegistrationDefinitionsServer
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of armmanagedservices.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// ServerFactoryTransport connects instances of armmanagedservices.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                                                    *ServerFactory
	trMu                                                   sync.Mutex
	trMarketplaceRegistrationDefinitionsServer             *MarketplaceRegistrationDefinitionsServerTransport
	trMarketplaceRegistrationDefinitionsWithoutScopeServer *MarketplaceRegistrationDefinitionsWithoutScopeServerTransport
	trOperationsServer                                     *OperationsServerTransport
	trOperationsWithScopeServer                            *OperationsWithScopeServerTransport
	trRegistrationAssignmentsServer                        *RegistrationAssignmentsServerTransport
	trRegistrationDefinitionsServer                        *RegistrationDefinitionsServerTransport
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
	case "MarketplaceRegistrationDefinitionsClient":
		initServer(s, &s.trMarketplaceRegistrationDefinitionsServer, func() *MarketplaceRegistrationDefinitionsServerTransport {
			return NewMarketplaceRegistrationDefinitionsServerTransport(&s.srv.MarketplaceRegistrationDefinitionsServer)
		})
		resp, err = s.trMarketplaceRegistrationDefinitionsServer.Do(req)
	case "MarketplaceRegistrationDefinitionsWithoutScopeClient":
		initServer(s, &s.trMarketplaceRegistrationDefinitionsWithoutScopeServer, func() *MarketplaceRegistrationDefinitionsWithoutScopeServerTransport {
			return NewMarketplaceRegistrationDefinitionsWithoutScopeServerTransport(&s.srv.MarketplaceRegistrationDefinitionsWithoutScopeServer)
		})
		resp, err = s.trMarketplaceRegistrationDefinitionsWithoutScopeServer.Do(req)
	case "OperationsClient":
		initServer(s, &s.trOperationsServer, func() *OperationsServerTransport { return NewOperationsServerTransport(&s.srv.OperationsServer) })
		resp, err = s.trOperationsServer.Do(req)
	case "OperationsWithScopeClient":
		initServer(s, &s.trOperationsWithScopeServer, func() *OperationsWithScopeServerTransport {
			return NewOperationsWithScopeServerTransport(&s.srv.OperationsWithScopeServer)
		})
		resp, err = s.trOperationsWithScopeServer.Do(req)
	case "RegistrationAssignmentsClient":
		initServer(s, &s.trRegistrationAssignmentsServer, func() *RegistrationAssignmentsServerTransport {
			return NewRegistrationAssignmentsServerTransport(&s.srv.RegistrationAssignmentsServer)
		})
		resp, err = s.trRegistrationAssignmentsServer.Do(req)
	case "RegistrationDefinitionsClient":
		initServer(s, &s.trRegistrationDefinitionsServer, func() *RegistrationDefinitionsServerTransport {
			return NewRegistrationDefinitionsServerTransport(&s.srv.RegistrationDefinitionsServer)
		})
		resp, err = s.trRegistrationDefinitionsServer.Do(req)
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
