//go:build go1.18
// +build go1.18

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

// ServerFactory is a fake server for instances of the armkubernetesruntime.ClientFactory type.
type ServerFactory struct {
	BgpPeersServer      BgpPeersServer
	LoadBalancersServer LoadBalancersServer
	OperationsServer    OperationsServer
	ServicesServer      ServicesServer
	StorageClassServer  StorageClassServer
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of armkubernetesruntime.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// ServerFactoryTransport connects instances of armkubernetesruntime.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                   *ServerFactory
	trMu                  sync.Mutex
	trBgpPeersServer      *BgpPeersServerTransport
	trLoadBalancersServer *LoadBalancersServerTransport
	trOperationsServer    *OperationsServerTransport
	trServicesServer      *ServicesServerTransport
	trStorageClassServer  *StorageClassServerTransport
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
	case "BgpPeersClient":
		initServer(s, &s.trBgpPeersServer, func() *BgpPeersServerTransport { return NewBgpPeersServerTransport(&s.srv.BgpPeersServer) })
		resp, err = s.trBgpPeersServer.Do(req)
	case "LoadBalancersClient":
		initServer(s, &s.trLoadBalancersServer, func() *LoadBalancersServerTransport {
			return NewLoadBalancersServerTransport(&s.srv.LoadBalancersServer)
		})
		resp, err = s.trLoadBalancersServer.Do(req)
	case "OperationsClient":
		initServer(s, &s.trOperationsServer, func() *OperationsServerTransport { return NewOperationsServerTransport(&s.srv.OperationsServer) })
		resp, err = s.trOperationsServer.Do(req)
	case "ServicesClient":
		initServer(s, &s.trServicesServer, func() *ServicesServerTransport { return NewServicesServerTransport(&s.srv.ServicesServer) })
		resp, err = s.trServicesServer.Do(req)
	case "StorageClassClient":
		initServer(s, &s.trStorageClassServer, func() *StorageClassServerTransport { return NewStorageClassServerTransport(&s.srv.StorageClassServer) })
		resp, err = s.trStorageClassServer.Do(req)
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