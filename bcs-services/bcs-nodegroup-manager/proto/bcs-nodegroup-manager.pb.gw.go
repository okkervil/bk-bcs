// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: proto/bcs-nodegroup-manager.proto

/*
Package nodegroupmanager is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package nodegroupmanager

import (
	"context"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = metadata.Join

func request_NodegroupManager_GetClusterAutoscalerReview_0(ctx context.Context, marshaler runtime.Marshaler, client NodegroupManagerClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ClusterAutoscalerReview
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.GetClusterAutoscalerReview(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_NodegroupManager_GetClusterAutoscalerReview_0(ctx context.Context, marshaler runtime.Marshaler, server NodegroupManagerServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ClusterAutoscalerReview
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.GetClusterAutoscalerReview(ctx, &protoReq)
	return msg, metadata, err

}

func request_NodegroupManager_CreateNodePoolMgrStrategy_0(ctx context.Context, marshaler runtime.Marshaler, client NodegroupManagerClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq CreateNodePoolMgrStrategyReq
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.CreateNodePoolMgrStrategy(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_NodegroupManager_CreateNodePoolMgrStrategy_0(ctx context.Context, marshaler runtime.Marshaler, server NodegroupManagerServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq CreateNodePoolMgrStrategyReq
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.CreateNodePoolMgrStrategy(ctx, &protoReq)
	return msg, metadata, err

}

func request_NodegroupManager_UpdateNodePoolMgrStrategy_0(ctx context.Context, marshaler runtime.Marshaler, client NodegroupManagerClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq UpdateNodePoolMgrStrategyReq
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.UpdateNodePoolMgrStrategy(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_NodegroupManager_UpdateNodePoolMgrStrategy_0(ctx context.Context, marshaler runtime.Marshaler, server NodegroupManagerServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq UpdateNodePoolMgrStrategyReq
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.UpdateNodePoolMgrStrategy(ctx, &protoReq)
	return msg, metadata, err

}

var (
	filter_NodegroupManager_GetNodePoolMgrStrategy_0 = &utilities.DoubleArray{Encoding: map[string]int{}, Base: []int(nil), Check: []int(nil)}
)

func request_NodegroupManager_GetNodePoolMgrStrategy_0(ctx context.Context, marshaler runtime.Marshaler, client NodegroupManagerClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetNodePoolMgrStrategyReq
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_NodegroupManager_GetNodePoolMgrStrategy_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.GetNodePoolMgrStrategy(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_NodegroupManager_GetNodePoolMgrStrategy_0(ctx context.Context, marshaler runtime.Marshaler, server NodegroupManagerServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetNodePoolMgrStrategyReq
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_NodegroupManager_GetNodePoolMgrStrategy_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.GetNodePoolMgrStrategy(ctx, &protoReq)
	return msg, metadata, err

}

var (
	filter_NodegroupManager_ListNodePoolMgrStrategies_0 = &utilities.DoubleArray{Encoding: map[string]int{}, Base: []int(nil), Check: []int(nil)}
)

func request_NodegroupManager_ListNodePoolMgrStrategies_0(ctx context.Context, marshaler runtime.Marshaler, client NodegroupManagerClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ListNodePoolMgrStrategyReq
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_NodegroupManager_ListNodePoolMgrStrategies_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.ListNodePoolMgrStrategies(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_NodegroupManager_ListNodePoolMgrStrategies_0(ctx context.Context, marshaler runtime.Marshaler, server NodegroupManagerServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ListNodePoolMgrStrategyReq
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_NodegroupManager_ListNodePoolMgrStrategies_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.ListNodePoolMgrStrategies(ctx, &protoReq)
	return msg, metadata, err

}

var (
	filter_NodegroupManager_DeleteNodePoolMgrStrategy_0 = &utilities.DoubleArray{Encoding: map[string]int{}, Base: []int(nil), Check: []int(nil)}
)

func request_NodegroupManager_DeleteNodePoolMgrStrategy_0(ctx context.Context, marshaler runtime.Marshaler, client NodegroupManagerClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq DeleteNodePoolMgrStrategyReq
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_NodegroupManager_DeleteNodePoolMgrStrategy_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.DeleteNodePoolMgrStrategy(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_NodegroupManager_DeleteNodePoolMgrStrategy_0(ctx context.Context, marshaler runtime.Marshaler, server NodegroupManagerServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq DeleteNodePoolMgrStrategyReq
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_NodegroupManager_DeleteNodePoolMgrStrategy_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.DeleteNodePoolMgrStrategy(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterNodegroupManagerGwServer registers the http handlers for service NodegroupManager to "mux".
// UnaryRPC     :call NodegroupManagerServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterNodegroupManagerGwFromEndpoint instead.
func RegisterNodegroupManagerGwServer(ctx context.Context, mux *runtime.ServeMux, server NodegroupManagerServer) error {

	mux.Handle("POST", pattern_NodegroupManager_GetClusterAutoscalerReview_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		ctx, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/nodegroupmanager.NodegroupManager/GetClusterAutoscalerReview", runtime.WithHTTPPathPattern("/nodegroupmanager/v1/autoscaler/review"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_NodegroupManager_GetClusterAutoscalerReview_0(ctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_NodegroupManager_GetClusterAutoscalerReview_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_NodegroupManager_CreateNodePoolMgrStrategy_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		ctx, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/nodegroupmanager.NodegroupManager/CreateNodePoolMgrStrategy", runtime.WithHTTPPathPattern("/nodegroupmanager/v1/nodegroupstrategy"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_NodegroupManager_CreateNodePoolMgrStrategy_0(ctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_NodegroupManager_CreateNodePoolMgrStrategy_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("PUT", pattern_NodegroupManager_UpdateNodePoolMgrStrategy_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		ctx, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/nodegroupmanager.NodegroupManager/UpdateNodePoolMgrStrategy", runtime.WithHTTPPathPattern("/nodegroupmanager/v1/nodegroupstrategy"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_NodegroupManager_UpdateNodePoolMgrStrategy_0(ctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_NodegroupManager_UpdateNodePoolMgrStrategy_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_NodegroupManager_GetNodePoolMgrStrategy_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		ctx, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/nodegroupmanager.NodegroupManager/GetNodePoolMgrStrategy", runtime.WithHTTPPathPattern("/nodegroupmanager/v1/nodegroupstrategy"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_NodegroupManager_GetNodePoolMgrStrategy_0(ctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_NodegroupManager_GetNodePoolMgrStrategy_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_NodegroupManager_ListNodePoolMgrStrategies_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		ctx, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/nodegroupmanager.NodegroupManager/ListNodePoolMgrStrategies", runtime.WithHTTPPathPattern("/nodegroupmanager/v1/nodegroupstrategies"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_NodegroupManager_ListNodePoolMgrStrategies_0(ctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_NodegroupManager_ListNodePoolMgrStrategies_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("DELETE", pattern_NodegroupManager_DeleteNodePoolMgrStrategy_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		ctx, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/nodegroupmanager.NodegroupManager/DeleteNodePoolMgrStrategy", runtime.WithHTTPPathPattern("/nodegroupmanager/v1/nodegroupstrategy"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_NodegroupManager_DeleteNodePoolMgrStrategy_0(ctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_NodegroupManager_DeleteNodePoolMgrStrategy_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterNodegroupManagerGwFromEndpoint is same as RegisterNodegroupManagerGw but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterNodegroupManagerGwFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterNodegroupManagerGw(ctx, mux, conn)
}

// RegisterNodegroupManagerGw registers the http handlers for service NodegroupManager to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterNodegroupManagerGw(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterNodegroupManagerGwClient(ctx, mux, NewNodegroupManagerClient(conn))
}

// RegisterNodegroupManagerGwClient registers the http handlers for service NodegroupManager
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "NodegroupManagerClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "NodegroupManagerClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "NodegroupManagerClient" to call the correct interceptors.
func RegisterNodegroupManagerGwClient(ctx context.Context, mux *runtime.ServeMux, client NodegroupManagerClient) error {

	mux.Handle("POST", pattern_NodegroupManager_GetClusterAutoscalerReview_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		ctx, err = runtime.AnnotateContext(ctx, mux, req, "/nodegroupmanager.NodegroupManager/GetClusterAutoscalerReview", runtime.WithHTTPPathPattern("/nodegroupmanager/v1/autoscaler/review"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_NodegroupManager_GetClusterAutoscalerReview_0(ctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_NodegroupManager_GetClusterAutoscalerReview_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_NodegroupManager_CreateNodePoolMgrStrategy_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		ctx, err = runtime.AnnotateContext(ctx, mux, req, "/nodegroupmanager.NodegroupManager/CreateNodePoolMgrStrategy", runtime.WithHTTPPathPattern("/nodegroupmanager/v1/nodegroupstrategy"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_NodegroupManager_CreateNodePoolMgrStrategy_0(ctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_NodegroupManager_CreateNodePoolMgrStrategy_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("PUT", pattern_NodegroupManager_UpdateNodePoolMgrStrategy_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		ctx, err = runtime.AnnotateContext(ctx, mux, req, "/nodegroupmanager.NodegroupManager/UpdateNodePoolMgrStrategy", runtime.WithHTTPPathPattern("/nodegroupmanager/v1/nodegroupstrategy"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_NodegroupManager_UpdateNodePoolMgrStrategy_0(ctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_NodegroupManager_UpdateNodePoolMgrStrategy_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_NodegroupManager_GetNodePoolMgrStrategy_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		ctx, err = runtime.AnnotateContext(ctx, mux, req, "/nodegroupmanager.NodegroupManager/GetNodePoolMgrStrategy", runtime.WithHTTPPathPattern("/nodegroupmanager/v1/nodegroupstrategy"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_NodegroupManager_GetNodePoolMgrStrategy_0(ctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_NodegroupManager_GetNodePoolMgrStrategy_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_NodegroupManager_ListNodePoolMgrStrategies_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		ctx, err = runtime.AnnotateContext(ctx, mux, req, "/nodegroupmanager.NodegroupManager/ListNodePoolMgrStrategies", runtime.WithHTTPPathPattern("/nodegroupmanager/v1/nodegroupstrategies"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_NodegroupManager_ListNodePoolMgrStrategies_0(ctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_NodegroupManager_ListNodePoolMgrStrategies_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("DELETE", pattern_NodegroupManager_DeleteNodePoolMgrStrategy_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		ctx, err = runtime.AnnotateContext(ctx, mux, req, "/nodegroupmanager.NodegroupManager/DeleteNodePoolMgrStrategy", runtime.WithHTTPPathPattern("/nodegroupmanager/v1/nodegroupstrategy"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_NodegroupManager_DeleteNodePoolMgrStrategy_0(ctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_NodegroupManager_DeleteNodePoolMgrStrategy_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_NodegroupManager_GetClusterAutoscalerReview_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"nodegroupmanager", "v1", "autoscaler", "review"}, ""))

	pattern_NodegroupManager_CreateNodePoolMgrStrategy_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"nodegroupmanager", "v1", "nodegroupstrategy"}, ""))

	pattern_NodegroupManager_UpdateNodePoolMgrStrategy_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"nodegroupmanager", "v1", "nodegroupstrategy"}, ""))

	pattern_NodegroupManager_GetNodePoolMgrStrategy_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"nodegroupmanager", "v1", "nodegroupstrategy"}, ""))

	pattern_NodegroupManager_ListNodePoolMgrStrategies_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"nodegroupmanager", "v1", "nodegroupstrategies"}, ""))

	pattern_NodegroupManager_DeleteNodePoolMgrStrategy_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"nodegroupmanager", "v1", "nodegroupstrategy"}, ""))
)

var (
	forward_NodegroupManager_GetClusterAutoscalerReview_0 = runtime.ForwardResponseMessage

	forward_NodegroupManager_CreateNodePoolMgrStrategy_0 = runtime.ForwardResponseMessage

	forward_NodegroupManager_UpdateNodePoolMgrStrategy_0 = runtime.ForwardResponseMessage

	forward_NodegroupManager_GetNodePoolMgrStrategy_0 = runtime.ForwardResponseMessage

	forward_NodegroupManager_ListNodePoolMgrStrategies_0 = runtime.ForwardResponseMessage

	forward_NodegroupManager_DeleteNodePoolMgrStrategy_0 = runtime.ForwardResponseMessage
)