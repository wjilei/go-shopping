// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: catalog.proto

package catalog

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Catalog service

type CatalogService interface {
	GetProductDetails(ctx context.Context, in *DetailRequest, opts ...client.CallOption) (*DetailResponse, error)
	GetProductCategories(ctx context.Context, in *AllCategoriesRequest, opts ...client.CallOption) (*AllCategoriesResponse, error)
	GetProductsInCategory(ctx context.Context, in *CategoryProductsRequest, opts ...client.CallOption) (*CategoryProductsResponse, error)
	ProductSearch(ctx context.Context, in *SearchRequest, opts ...client.CallOption) (*SearchResponse, error)
}

type catalogService struct {
	c    client.Client
	name string
}

func NewCatalogService(name string, c client.Client) CatalogService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "catalog"
	}
	return &catalogService{
		c:    c,
		name: name,
	}
}

func (c *catalogService) GetProductDetails(ctx context.Context, in *DetailRequest, opts ...client.CallOption) (*DetailResponse, error) {
	req := c.c.NewRequest(c.name, "Catalog.GetProductDetails", in)
	out := new(DetailResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogService) GetProductCategories(ctx context.Context, in *AllCategoriesRequest, opts ...client.CallOption) (*AllCategoriesResponse, error) {
	req := c.c.NewRequest(c.name, "Catalog.GetProductCategories", in)
	out := new(AllCategoriesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogService) GetProductsInCategory(ctx context.Context, in *CategoryProductsRequest, opts ...client.CallOption) (*CategoryProductsResponse, error) {
	req := c.c.NewRequest(c.name, "Catalog.GetProductsInCategory", in)
	out := new(CategoryProductsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogService) ProductSearch(ctx context.Context, in *SearchRequest, opts ...client.CallOption) (*SearchResponse, error) {
	req := c.c.NewRequest(c.name, "Catalog.ProductSearch", in)
	out := new(SearchResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Catalog service

type CatalogHandler interface {
	GetProductDetails(context.Context, *DetailRequest, *DetailResponse) error
	GetProductCategories(context.Context, *AllCategoriesRequest, *AllCategoriesResponse) error
	GetProductsInCategory(context.Context, *CategoryProductsRequest, *CategoryProductsResponse) error
	ProductSearch(context.Context, *SearchRequest, *SearchResponse) error
}

func RegisterCatalogHandler(s server.Server, hdlr CatalogHandler, opts ...server.HandlerOption) error {
	type catalog interface {
		GetProductDetails(ctx context.Context, in *DetailRequest, out *DetailResponse) error
		GetProductCategories(ctx context.Context, in *AllCategoriesRequest, out *AllCategoriesResponse) error
		GetProductsInCategory(ctx context.Context, in *CategoryProductsRequest, out *CategoryProductsResponse) error
		ProductSearch(ctx context.Context, in *SearchRequest, out *SearchResponse) error
	}
	type Catalog struct {
		catalog
	}
	h := &catalogHandler{hdlr}
	return s.Handle(s.NewHandler(&Catalog{h}, opts...))
}

type catalogHandler struct {
	CatalogHandler
}

func (h *catalogHandler) GetProductDetails(ctx context.Context, in *DetailRequest, out *DetailResponse) error {
	return h.CatalogHandler.GetProductDetails(ctx, in, out)
}

func (h *catalogHandler) GetProductCategories(ctx context.Context, in *AllCategoriesRequest, out *AllCategoriesResponse) error {
	return h.CatalogHandler.GetProductCategories(ctx, in, out)
}

func (h *catalogHandler) GetProductsInCategory(ctx context.Context, in *CategoryProductsRequest, out *CategoryProductsResponse) error {
	return h.CatalogHandler.GetProductsInCategory(ctx, in, out)
}

func (h *catalogHandler) ProductSearch(ctx context.Context, in *SearchRequest, out *SearchResponse) error {
	return h.CatalogHandler.ProductSearch(ctx, in, out)
}