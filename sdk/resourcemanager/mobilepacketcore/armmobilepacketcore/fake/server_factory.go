// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
	"sync"
)

// ServerFactory is a fake server for instances of the armmobilepacketcore.ClientFactory type.
type ServerFactory struct {
	// AmfDeploymentsServer contains the fakes for client AmfDeploymentsClient
	AmfDeploymentsServer AmfDeploymentsServer

	// ClusterServicesServer contains the fakes for client ClusterServicesClient
	ClusterServicesServer ClusterServicesServer

	// NrfDeploymentsServer contains the fakes for client NrfDeploymentsClient
	NrfDeploymentsServer NrfDeploymentsServer

	// NssfDeploymentsServer contains the fakes for client NssfDeploymentsClient
	NssfDeploymentsServer NssfDeploymentsServer

	// ObservabilityServicesServer contains the fakes for client ObservabilityServicesClient
	ObservabilityServicesServer ObservabilityServicesServer

	// OperationsServer contains the fakes for client OperationsClient
	OperationsServer OperationsServer

	// SmfDeploymentsServer contains the fakes for client SmfDeploymentsClient
	SmfDeploymentsServer SmfDeploymentsServer

	// UpfDeploymentsServer contains the fakes for client UpfDeploymentsClient
	UpfDeploymentsServer UpfDeploymentsServer
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of armmobilepacketcore.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// ServerFactoryTransport connects instances of armmobilepacketcore.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                           *ServerFactory
	trMu                          sync.Mutex
	trAmfDeploymentsServer        *AmfDeploymentsServerTransport
	trClusterServicesServer       *ClusterServicesServerTransport
	trNrfDeploymentsServer        *NrfDeploymentsServerTransport
	trNssfDeploymentsServer       *NssfDeploymentsServerTransport
	trObservabilityServicesServer *ObservabilityServicesServerTransport
	trOperationsServer            *OperationsServerTransport
	trSmfDeploymentsServer        *SmfDeploymentsServerTransport
	trUpfDeploymentsServer        *UpfDeploymentsServerTransport
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
	case "AmfDeploymentsClient":
		initServer(s, &s.trAmfDeploymentsServer, func() *AmfDeploymentsServerTransport {
			return NewAmfDeploymentsServerTransport(&s.srv.AmfDeploymentsServer)
		})
		resp, err = s.trAmfDeploymentsServer.Do(req)
	case "ClusterServicesClient":
		initServer(s, &s.trClusterServicesServer, func() *ClusterServicesServerTransport {
			return NewClusterServicesServerTransport(&s.srv.ClusterServicesServer)
		})
		resp, err = s.trClusterServicesServer.Do(req)
	case "NrfDeploymentsClient":
		initServer(s, &s.trNrfDeploymentsServer, func() *NrfDeploymentsServerTransport {
			return NewNrfDeploymentsServerTransport(&s.srv.NrfDeploymentsServer)
		})
		resp, err = s.trNrfDeploymentsServer.Do(req)
	case "NssfDeploymentsClient":
		initServer(s, &s.trNssfDeploymentsServer, func() *NssfDeploymentsServerTransport {
			return NewNssfDeploymentsServerTransport(&s.srv.NssfDeploymentsServer)
		})
		resp, err = s.trNssfDeploymentsServer.Do(req)
	case "ObservabilityServicesClient":
		initServer(s, &s.trObservabilityServicesServer, func() *ObservabilityServicesServerTransport {
			return NewObservabilityServicesServerTransport(&s.srv.ObservabilityServicesServer)
		})
		resp, err = s.trObservabilityServicesServer.Do(req)
	case "OperationsClient":
		initServer(s, &s.trOperationsServer, func() *OperationsServerTransport { return NewOperationsServerTransport(&s.srv.OperationsServer) })
		resp, err = s.trOperationsServer.Do(req)
	case "SmfDeploymentsClient":
		initServer(s, &s.trSmfDeploymentsServer, func() *SmfDeploymentsServerTransport {
			return NewSmfDeploymentsServerTransport(&s.srv.SmfDeploymentsServer)
		})
		resp, err = s.trSmfDeploymentsServer.Do(req)
	case "UpfDeploymentsClient":
		initServer(s, &s.trUpfDeploymentsServer, func() *UpfDeploymentsServerTransport {
			return NewUpfDeploymentsServerTransport(&s.srv.UpfDeploymentsServer)
		})
		resp, err = s.trUpfDeploymentsServer.Do(req)
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
