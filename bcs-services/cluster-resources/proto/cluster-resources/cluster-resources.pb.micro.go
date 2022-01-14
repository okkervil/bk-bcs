// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/cluster-resources/cluster-resources.proto

package cluster_resources

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/struct"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for ClusterResources service

func NewClusterResourcesEndpoints() []*api.Endpoint {
	return []*api.Endpoint{
		&api.Endpoint{
			Name:    "ClusterResources.Echo",
			Path:    []string{"/clusterresources/v1/echo"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "ClusterResources.Ping",
			Path:    []string{"/clusterresources/v1/ping"},
			Method:  []string{"GET"},
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "ClusterResources.Healthz",
			Path:    []string{"/clusterresources/v1/healthz"},
			Method:  []string{"GET"},
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "ClusterResources.Version",
			Path:    []string{"/clusterresources/v1/version"},
			Method:  []string{"GET"},
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "ClusterResources.ListDeploy",
			Path:    []string{"/clusterresources/v1/projects/{projectID}/clusters/{clusterID}/namespaces/{namespace}/workloads/deployments"},
			Method:  []string{"GET"},
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "ClusterResources.GetDeploy",
			Path:    []string{"/clusterresources/v1/projects/{projectID}/clusters/{clusterID}/namespaces/{namespace}/workloads/deployments/{name}"},
			Method:  []string{"GET"},
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "ClusterResources.CreateDeploy",
			Path:    []string{"/clusterresources/v1/projects/{projectID}/clusters/{clusterID}/workloads/deployments"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "ClusterResources.UpdateDeploy",
			Path:    []string{"/clusterresources/v1/projects/{projectID}/clusters/{clusterID}/namespaces/{namespace}/workloads/deployments/{name}"},
			Method:  []string{"PUT"},
			Body:    "*",
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "ClusterResources.DeleteDeploy",
			Path:    []string{"/clusterresources/v1/projects/{projectID}/clusters/{clusterID}/namespaces/{namespace}/workloads/deployments/{name}"},
			Method:  []string{"DELETE"},
			Body:    "",
			Handler: "rpc",
		},
	}
}

// Client API for ClusterResources service

type ClusterResourcesService interface {
	// 基础类接口
	Echo(ctx context.Context, in *EchoReq, opts ...client.CallOption) (*EchoResp, error)
	Ping(ctx context.Context, in *PingReq, opts ...client.CallOption) (*PingResp, error)
	Healthz(ctx context.Context, in *HealthzReq, opts ...client.CallOption) (*HealthzResp, error)
	Version(ctx context.Context, in *VersionReq, opts ...client.CallOption) (*VersionResp, error)
	// 工作负载类接口
	ListDeploy(ctx context.Context, in *NamespaceScopedResListReq, opts ...client.CallOption) (*CommonResp, error)
	GetDeploy(ctx context.Context, in *NamespaceScopedResGetReq, opts ...client.CallOption) (*CommonResp, error)
	CreateDeploy(ctx context.Context, in *NamespaceScopedResCreateReq, opts ...client.CallOption) (*CommonResp, error)
	UpdateDeploy(ctx context.Context, in *NamespaceScopedResUpdateReq, opts ...client.CallOption) (*CommonResp, error)
	DeleteDeploy(ctx context.Context, in *NamespaceScopedResDeleteReq, opts ...client.CallOption) (*CommonResp, error)
}

type clusterResourcesService struct {
	c    client.Client
	name string
}

func NewClusterResourcesService(name string, c client.Client) ClusterResourcesService {
	return &clusterResourcesService{
		c:    c,
		name: name,
	}
}

func (c *clusterResourcesService) Echo(ctx context.Context, in *EchoReq, opts ...client.CallOption) (*EchoResp, error) {
	req := c.c.NewRequest(c.name, "ClusterResources.Echo", in)
	out := new(EchoResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterResourcesService) Ping(ctx context.Context, in *PingReq, opts ...client.CallOption) (*PingResp, error) {
	req := c.c.NewRequest(c.name, "ClusterResources.Ping", in)
	out := new(PingResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterResourcesService) Healthz(ctx context.Context, in *HealthzReq, opts ...client.CallOption) (*HealthzResp, error) {
	req := c.c.NewRequest(c.name, "ClusterResources.Healthz", in)
	out := new(HealthzResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterResourcesService) Version(ctx context.Context, in *VersionReq, opts ...client.CallOption) (*VersionResp, error) {
	req := c.c.NewRequest(c.name, "ClusterResources.Version", in)
	out := new(VersionResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterResourcesService) ListDeploy(ctx context.Context, in *NamespaceScopedResListReq, opts ...client.CallOption) (*CommonResp, error) {
	req := c.c.NewRequest(c.name, "ClusterResources.ListDeploy", in)
	out := new(CommonResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterResourcesService) GetDeploy(ctx context.Context, in *NamespaceScopedResGetReq, opts ...client.CallOption) (*CommonResp, error) {
	req := c.c.NewRequest(c.name, "ClusterResources.GetDeploy", in)
	out := new(CommonResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterResourcesService) CreateDeploy(ctx context.Context, in *NamespaceScopedResCreateReq, opts ...client.CallOption) (*CommonResp, error) {
	req := c.c.NewRequest(c.name, "ClusterResources.CreateDeploy", in)
	out := new(CommonResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterResourcesService) UpdateDeploy(ctx context.Context, in *NamespaceScopedResUpdateReq, opts ...client.CallOption) (*CommonResp, error) {
	req := c.c.NewRequest(c.name, "ClusterResources.UpdateDeploy", in)
	out := new(CommonResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterResourcesService) DeleteDeploy(ctx context.Context, in *NamespaceScopedResDeleteReq, opts ...client.CallOption) (*CommonResp, error) {
	req := c.c.NewRequest(c.name, "ClusterResources.DeleteDeploy", in)
	out := new(CommonResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ClusterResources service

type ClusterResourcesHandler interface {
	// 基础类接口
	Echo(context.Context, *EchoReq, *EchoResp) error
	Ping(context.Context, *PingReq, *PingResp) error
	Healthz(context.Context, *HealthzReq, *HealthzResp) error
	Version(context.Context, *VersionReq, *VersionResp) error
	// 工作负载类接口
	ListDeploy(context.Context, *NamespaceScopedResListReq, *CommonResp) error
	GetDeploy(context.Context, *NamespaceScopedResGetReq, *CommonResp) error
	CreateDeploy(context.Context, *NamespaceScopedResCreateReq, *CommonResp) error
	UpdateDeploy(context.Context, *NamespaceScopedResUpdateReq, *CommonResp) error
	DeleteDeploy(context.Context, *NamespaceScopedResDeleteReq, *CommonResp) error
}

func RegisterClusterResourcesHandler(s server.Server, hdlr ClusterResourcesHandler, opts ...server.HandlerOption) error {
	type clusterResources interface {
		Echo(ctx context.Context, in *EchoReq, out *EchoResp) error
		Ping(ctx context.Context, in *PingReq, out *PingResp) error
		Healthz(ctx context.Context, in *HealthzReq, out *HealthzResp) error
		Version(ctx context.Context, in *VersionReq, out *VersionResp) error
		ListDeploy(ctx context.Context, in *NamespaceScopedResListReq, out *CommonResp) error
		GetDeploy(ctx context.Context, in *NamespaceScopedResGetReq, out *CommonResp) error
		CreateDeploy(ctx context.Context, in *NamespaceScopedResCreateReq, out *CommonResp) error
		UpdateDeploy(ctx context.Context, in *NamespaceScopedResUpdateReq, out *CommonResp) error
		DeleteDeploy(ctx context.Context, in *NamespaceScopedResDeleteReq, out *CommonResp) error
	}
	type ClusterResources struct {
		clusterResources
	}
	h := &clusterResourcesHandler{hdlr}
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "ClusterResources.Echo",
		Path:    []string{"/clusterresources/v1/echo"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "ClusterResources.Ping",
		Path:    []string{"/clusterresources/v1/ping"},
		Method:  []string{"GET"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "ClusterResources.Healthz",
		Path:    []string{"/clusterresources/v1/healthz"},
		Method:  []string{"GET"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "ClusterResources.Version",
		Path:    []string{"/clusterresources/v1/version"},
		Method:  []string{"GET"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "ClusterResources.ListDeploy",
		Path:    []string{"/clusterresources/v1/projects/{projectID}/clusters/{clusterID}/namespaces/{namespace}/workloads/deployments"},
		Method:  []string{"GET"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "ClusterResources.GetDeploy",
		Path:    []string{"/clusterresources/v1/projects/{projectID}/clusters/{clusterID}/namespaces/{namespace}/workloads/deployments/{name}"},
		Method:  []string{"GET"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "ClusterResources.CreateDeploy",
		Path:    []string{"/clusterresources/v1/projects/{projectID}/clusters/{clusterID}/workloads/deployments"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "ClusterResources.UpdateDeploy",
		Path:    []string{"/clusterresources/v1/projects/{projectID}/clusters/{clusterID}/namespaces/{namespace}/workloads/deployments/{name}"},
		Method:  []string{"PUT"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "ClusterResources.DeleteDeploy",
		Path:    []string{"/clusterresources/v1/projects/{projectID}/clusters/{clusterID}/namespaces/{namespace}/workloads/deployments/{name}"},
		Method:  []string{"DELETE"},
		Body:    "",
		Handler: "rpc",
	}))
	return s.Handle(s.NewHandler(&ClusterResources{h}, opts...))
}

type clusterResourcesHandler struct {
	ClusterResourcesHandler
}

func (h *clusterResourcesHandler) Echo(ctx context.Context, in *EchoReq, out *EchoResp) error {
	return h.ClusterResourcesHandler.Echo(ctx, in, out)
}

func (h *clusterResourcesHandler) Ping(ctx context.Context, in *PingReq, out *PingResp) error {
	return h.ClusterResourcesHandler.Ping(ctx, in, out)
}

func (h *clusterResourcesHandler) Healthz(ctx context.Context, in *HealthzReq, out *HealthzResp) error {
	return h.ClusterResourcesHandler.Healthz(ctx, in, out)
}

func (h *clusterResourcesHandler) Version(ctx context.Context, in *VersionReq, out *VersionResp) error {
	return h.ClusterResourcesHandler.Version(ctx, in, out)
}

func (h *clusterResourcesHandler) ListDeploy(ctx context.Context, in *NamespaceScopedResListReq, out *CommonResp) error {
	return h.ClusterResourcesHandler.ListDeploy(ctx, in, out)
}

func (h *clusterResourcesHandler) GetDeploy(ctx context.Context, in *NamespaceScopedResGetReq, out *CommonResp) error {
	return h.ClusterResourcesHandler.GetDeploy(ctx, in, out)
}

func (h *clusterResourcesHandler) CreateDeploy(ctx context.Context, in *NamespaceScopedResCreateReq, out *CommonResp) error {
	return h.ClusterResourcesHandler.CreateDeploy(ctx, in, out)
}

func (h *clusterResourcesHandler) UpdateDeploy(ctx context.Context, in *NamespaceScopedResUpdateReq, out *CommonResp) error {
	return h.ClusterResourcesHandler.UpdateDeploy(ctx, in, out)
}

func (h *clusterResourcesHandler) DeleteDeploy(ctx context.Context, in *NamespaceScopedResDeleteReq, out *CommonResp) error {
	return h.ClusterResourcesHandler.DeleteDeploy(ctx, in, out)
}
