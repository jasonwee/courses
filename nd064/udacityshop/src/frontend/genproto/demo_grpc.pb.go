// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package __

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

// ProductCatalogServiceClient is the client API for ProductCatalogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductCatalogServiceClient interface {
	ListProducts(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListProductsResponse, error)
	GetProduct(ctx context.Context, in *GetProductRequest, opts ...grpc.CallOption) (*Product, error)
	SearchProducts(ctx context.Context, in *SearchProductsRequest, opts ...grpc.CallOption) (*SearchProductsResponse, error)
}

type productCatalogServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductCatalogServiceClient(cc grpc.ClientConnInterface) ProductCatalogServiceClient {
	return &productCatalogServiceClient{cc}
}

func (c *productCatalogServiceClient) ListProducts(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListProductsResponse, error) {
	out := new(ListProductsResponse)
	err := c.cc.Invoke(ctx, "/hipstershop.ProductCatalogService/ListProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productCatalogServiceClient) GetProduct(ctx context.Context, in *GetProductRequest, opts ...grpc.CallOption) (*Product, error) {
	out := new(Product)
	err := c.cc.Invoke(ctx, "/hipstershop.ProductCatalogService/GetProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productCatalogServiceClient) SearchProducts(ctx context.Context, in *SearchProductsRequest, opts ...grpc.CallOption) (*SearchProductsResponse, error) {
	out := new(SearchProductsResponse)
	err := c.cc.Invoke(ctx, "/hipstershop.ProductCatalogService/SearchProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductCatalogServiceServer is the server API for ProductCatalogService service.
// All implementations must embed UnimplementedProductCatalogServiceServer
// for forward compatibility
type ProductCatalogServiceServer interface {
	ListProducts(context.Context, *Empty) (*ListProductsResponse, error)
	GetProduct(context.Context, *GetProductRequest) (*Product, error)
	SearchProducts(context.Context, *SearchProductsRequest) (*SearchProductsResponse, error)
	mustEmbedUnimplementedProductCatalogServiceServer()
}

// UnimplementedProductCatalogServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProductCatalogServiceServer struct {
}

func (UnimplementedProductCatalogServiceServer) ListProducts(context.Context, *Empty) (*ListProductsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProducts not implemented")
}
func (UnimplementedProductCatalogServiceServer) GetProduct(context.Context, *GetProductRequest) (*Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProduct not implemented")
}
func (UnimplementedProductCatalogServiceServer) SearchProducts(context.Context, *SearchProductsRequest) (*SearchProductsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchProducts not implemented")
}
func (UnimplementedProductCatalogServiceServer) mustEmbedUnimplementedProductCatalogServiceServer() {}

// UnsafeProductCatalogServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductCatalogServiceServer will
// result in compilation errors.
type UnsafeProductCatalogServiceServer interface {
	mustEmbedUnimplementedProductCatalogServiceServer()
}

func RegisterProductCatalogServiceServer(s grpc.ServiceRegistrar, srv ProductCatalogServiceServer) {
	s.RegisterService(&ProductCatalogService_ServiceDesc, srv)
}

func _ProductCatalogService_ListProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductCatalogServiceServer).ListProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hipstershop.ProductCatalogService/ListProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductCatalogServiceServer).ListProducts(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductCatalogService_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductCatalogServiceServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hipstershop.ProductCatalogService/GetProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductCatalogServiceServer).GetProduct(ctx, req.(*GetProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductCatalogService_SearchProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchProductsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductCatalogServiceServer).SearchProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hipstershop.ProductCatalogService/SearchProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductCatalogServiceServer).SearchProducts(ctx, req.(*SearchProductsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductCatalogService_ServiceDesc is the grpc.ServiceDesc for ProductCatalogService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductCatalogService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hipstershop.ProductCatalogService",
	HandlerType: (*ProductCatalogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListProducts",
			Handler:    _ProductCatalogService_ListProducts_Handler,
		},
		{
			MethodName: "GetProduct",
			Handler:    _ProductCatalogService_GetProduct_Handler,
		},
		{
			MethodName: "SearchProducts",
			Handler:    _ProductCatalogService_SearchProducts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "demo.proto",
}

// CurrencyServiceClient is the client API for CurrencyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CurrencyServiceClient interface {
	GetSupportedCurrencies(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetSupportedCurrenciesResponse, error)
	Convert(ctx context.Context, in *CurrencyConversionRequest, opts ...grpc.CallOption) (*Money, error)
}

type currencyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCurrencyServiceClient(cc grpc.ClientConnInterface) CurrencyServiceClient {
	return &currencyServiceClient{cc}
}

func (c *currencyServiceClient) GetSupportedCurrencies(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetSupportedCurrenciesResponse, error) {
	out := new(GetSupportedCurrenciesResponse)
	err := c.cc.Invoke(ctx, "/hipstershop.CurrencyService/GetSupportedCurrencies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currencyServiceClient) Convert(ctx context.Context, in *CurrencyConversionRequest, opts ...grpc.CallOption) (*Money, error) {
	out := new(Money)
	err := c.cc.Invoke(ctx, "/hipstershop.CurrencyService/Convert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CurrencyServiceServer is the server API for CurrencyService service.
// All implementations must embed UnimplementedCurrencyServiceServer
// for forward compatibility
type CurrencyServiceServer interface {
	GetSupportedCurrencies(context.Context, *Empty) (*GetSupportedCurrenciesResponse, error)
	Convert(context.Context, *CurrencyConversionRequest) (*Money, error)
	mustEmbedUnimplementedCurrencyServiceServer()
}

// UnimplementedCurrencyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCurrencyServiceServer struct {
}

func (UnimplementedCurrencyServiceServer) GetSupportedCurrencies(context.Context, *Empty) (*GetSupportedCurrenciesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSupportedCurrencies not implemented")
}
func (UnimplementedCurrencyServiceServer) Convert(context.Context, *CurrencyConversionRequest) (*Money, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Convert not implemented")
}
func (UnimplementedCurrencyServiceServer) mustEmbedUnimplementedCurrencyServiceServer() {}

// UnsafeCurrencyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CurrencyServiceServer will
// result in compilation errors.
type UnsafeCurrencyServiceServer interface {
	mustEmbedUnimplementedCurrencyServiceServer()
}

func RegisterCurrencyServiceServer(s grpc.ServiceRegistrar, srv CurrencyServiceServer) {
	s.RegisterService(&CurrencyService_ServiceDesc, srv)
}

func _CurrencyService_GetSupportedCurrencies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrencyServiceServer).GetSupportedCurrencies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hipstershop.CurrencyService/GetSupportedCurrencies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrencyServiceServer).GetSupportedCurrencies(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CurrencyService_Convert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CurrencyConversionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrencyServiceServer).Convert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hipstershop.CurrencyService/Convert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrencyServiceServer).Convert(ctx, req.(*CurrencyConversionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CurrencyService_ServiceDesc is the grpc.ServiceDesc for CurrencyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CurrencyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hipstershop.CurrencyService",
	HandlerType: (*CurrencyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSupportedCurrencies",
			Handler:    _CurrencyService_GetSupportedCurrencies_Handler,
		},
		{
			MethodName: "Convert",
			Handler:    _CurrencyService_Convert_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "demo.proto",
}

// AdServiceClient is the client API for AdService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdServiceClient interface {
	GetAds(ctx context.Context, in *AdRequest, opts ...grpc.CallOption) (*AdResponse, error)
}

type adServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdServiceClient(cc grpc.ClientConnInterface) AdServiceClient {
	return &adServiceClient{cc}
}

func (c *adServiceClient) GetAds(ctx context.Context, in *AdRequest, opts ...grpc.CallOption) (*AdResponse, error) {
	out := new(AdResponse)
	err := c.cc.Invoke(ctx, "/hipstershop.AdService/GetAds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdServiceServer is the server API for AdService service.
// All implementations must embed UnimplementedAdServiceServer
// for forward compatibility
type AdServiceServer interface {
	GetAds(context.Context, *AdRequest) (*AdResponse, error)
	mustEmbedUnimplementedAdServiceServer()
}

// UnimplementedAdServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAdServiceServer struct {
}

func (UnimplementedAdServiceServer) GetAds(context.Context, *AdRequest) (*AdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAds not implemented")
}
func (UnimplementedAdServiceServer) mustEmbedUnimplementedAdServiceServer() {}

// UnsafeAdServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdServiceServer will
// result in compilation errors.
type UnsafeAdServiceServer interface {
	mustEmbedUnimplementedAdServiceServer()
}

func RegisterAdServiceServer(s grpc.ServiceRegistrar, srv AdServiceServer) {
	s.RegisterService(&AdService_ServiceDesc, srv)
}

func _AdService_GetAds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdServiceServer).GetAds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hipstershop.AdService/GetAds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdServiceServer).GetAds(ctx, req.(*AdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AdService_ServiceDesc is the grpc.ServiceDesc for AdService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hipstershop.AdService",
	HandlerType: (*AdServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAds",
			Handler:    _AdService_GetAds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "demo.proto",
}

// AdServiceV2Client is the client API for AdServiceV2 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdServiceV2Client interface {
	GetAdsV2(ctx context.Context, in *AdRequest, opts ...grpc.CallOption) (*AdResponse, error)
}

type adServiceV2Client struct {
	cc grpc.ClientConnInterface
}

func NewAdServiceV2Client(cc grpc.ClientConnInterface) AdServiceV2Client {
	return &adServiceV2Client{cc}
}

func (c *adServiceV2Client) GetAdsV2(ctx context.Context, in *AdRequest, opts ...grpc.CallOption) (*AdResponse, error) {
	out := new(AdResponse)
	err := c.cc.Invoke(ctx, "/hipstershop.AdServiceV2/GetAdsV2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdServiceV2Server is the server API for AdServiceV2 service.
// All implementations must embed UnimplementedAdServiceV2Server
// for forward compatibility
type AdServiceV2Server interface {
	GetAdsV2(context.Context, *AdRequest) (*AdResponse, error)
	mustEmbedUnimplementedAdServiceV2Server()
}

// UnimplementedAdServiceV2Server must be embedded to have forward compatible implementations.
type UnimplementedAdServiceV2Server struct {
}

func (UnimplementedAdServiceV2Server) GetAdsV2(context.Context, *AdRequest) (*AdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAdsV2 not implemented")
}
func (UnimplementedAdServiceV2Server) mustEmbedUnimplementedAdServiceV2Server() {}

// UnsafeAdServiceV2Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdServiceV2Server will
// result in compilation errors.
type UnsafeAdServiceV2Server interface {
	mustEmbedUnimplementedAdServiceV2Server()
}

func RegisterAdServiceV2Server(s grpc.ServiceRegistrar, srv AdServiceV2Server) {
	s.RegisterService(&AdServiceV2_ServiceDesc, srv)
}

func _AdServiceV2_GetAdsV2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdServiceV2Server).GetAdsV2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hipstershop.AdServiceV2/GetAdsV2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdServiceV2Server).GetAdsV2(ctx, req.(*AdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AdServiceV2_ServiceDesc is the grpc.ServiceDesc for AdServiceV2 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdServiceV2_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hipstershop.AdServiceV2",
	HandlerType: (*AdServiceV2Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAdsV2",
			Handler:    _AdServiceV2_GetAdsV2_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "demo.proto",
}
