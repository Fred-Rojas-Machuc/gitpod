// Copyright (c) 2022 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: idemetrics.proto

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MetricsServiceClient is the client API for MetricsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MetricsServiceClient interface {
	AddCounter(ctx context.Context, in *AddCounterRequest, opts ...grpc.CallOption) (*AddCounterResponse, error)
	ObserveHistogram(ctx context.Context, in *ObserveHistogramRequest, opts ...grpc.CallOption) (*ObserveHistogramResponse, error)
}

type metricsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMetricsServiceClient(cc grpc.ClientConnInterface) MetricsServiceClient {
	return &metricsServiceClient{cc}
}

func (c *metricsServiceClient) AddCounter(ctx context.Context, in *AddCounterRequest, opts ...grpc.CallOption) (*AddCounterResponse, error) {
	out := new(AddCounterResponse)
	err := c.cc.Invoke(ctx, "/ide_metrics_api.MetricsService/AddCounter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsServiceClient) ObserveHistogram(ctx context.Context, in *ObserveHistogramRequest, opts ...grpc.CallOption) (*ObserveHistogramResponse, error) {
	out := new(ObserveHistogramResponse)
	err := c.cc.Invoke(ctx, "/ide_metrics_api.MetricsService/ObserveHistogram", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetricsServiceServer is the server API for MetricsService service.
// All implementations must embed UnimplementedMetricsServiceServer
// for forward compatibility
type MetricsServiceServer interface {
	AddCounter(context.Context, *AddCounterRequest) (*AddCounterResponse, error)
	ObserveHistogram(context.Context, *ObserveHistogramRequest) (*ObserveHistogramResponse, error)
	mustEmbedUnimplementedMetricsServiceServer()
}

// UnimplementedMetricsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMetricsServiceServer struct {
}

func (UnimplementedMetricsServiceServer) AddCounter(context.Context, *AddCounterRequest) (*AddCounterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCounter not implemented")
}
func (UnimplementedMetricsServiceServer) ObserveHistogram(context.Context, *ObserveHistogramRequest) (*ObserveHistogramResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ObserveHistogram not implemented")
}
func (UnimplementedMetricsServiceServer) mustEmbedUnimplementedMetricsServiceServer() {}

// UnsafeMetricsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MetricsServiceServer will
// result in compilation errors.
type UnsafeMetricsServiceServer interface {
	mustEmbedUnimplementedMetricsServiceServer()
}

func RegisterMetricsServiceServer(s grpc.ServiceRegistrar, srv MetricsServiceServer) {
	s.RegisterService(&MetricsService_ServiceDesc, srv)
}

func _MetricsService_AddCounter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCounterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServiceServer).AddCounter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ide_metrics_api.MetricsService/AddCounter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServiceServer).AddCounter(ctx, req.(*AddCounterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricsService_ObserveHistogram_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ObserveHistogramRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServiceServer).ObserveHistogram(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ide_metrics_api.MetricsService/ObserveHistogram",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServiceServer).ObserveHistogram(ctx, req.(*ObserveHistogramRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MetricsService_ServiceDesc is the grpc.ServiceDesc for MetricsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MetricsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ide_metrics_api.MetricsService",
	HandlerType: (*MetricsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddCounter",
			Handler:    _MetricsService_AddCounter_Handler,
		},
		{
			MethodName: "ObserveHistogram",
			Handler:    _MetricsService_ObserveHistogram_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "idemetrics.proto",
}
