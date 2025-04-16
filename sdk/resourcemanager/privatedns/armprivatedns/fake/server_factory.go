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

// ServerFactory is a fake server for instances of the armprivatedns.ClientFactory type.
type ServerFactory struct {
	// PrivateZonesServer contains the fakes for client PrivateZonesClient
	PrivateZonesServer PrivateZonesServer

	// RecordSetsServer contains the fakes for client RecordSetsClient
	RecordSetsServer RecordSetsServer

	// VirtualNetworkLinksServer contains the fakes for client VirtualNetworkLinksClient
	VirtualNetworkLinksServer VirtualNetworkLinksServer
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of armprivatedns.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// ServerFactoryTransport connects instances of armprivatedns.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                         *ServerFactory
	trMu                        sync.Mutex
	trPrivateZonesServer        *PrivateZonesServerTransport
	trRecordSetsServer          *RecordSetsServerTransport
	trVirtualNetworkLinksServer *VirtualNetworkLinksServerTransport
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
	case "PrivateZonesClient":
		initServer(s, &s.trPrivateZonesServer, func() *PrivateZonesServerTransport { return NewPrivateZonesServerTransport(&s.srv.PrivateZonesServer) })
		resp, err = s.trPrivateZonesServer.Do(req)
	case "RecordSetsClient":
		initServer(s, &s.trRecordSetsServer, func() *RecordSetsServerTransport { return NewRecordSetsServerTransport(&s.srv.RecordSetsServer) })
		resp, err = s.trRecordSetsServer.Do(req)
	case "VirtualNetworkLinksClient":
		initServer(s, &s.trVirtualNetworkLinksServer, func() *VirtualNetworkLinksServerTransport {
			return NewVirtualNetworkLinksServerTransport(&s.srv.VirtualNetworkLinksServer)
		})
		resp, err = s.trVirtualNetworkLinksServer.Do(req)
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
