/*
Tencent is pleased to support the open source community by making Basic Service Configuration Platform available.
Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
Licensed under the MIT License (the "License"); you may not use this file except
in compliance with the License. You may obtain a copy of the License at
http://opensource.org/licenses/MIT
Unless required by applicable law or agreed to in writing, software distributed under
the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
either express or implied. See the License for the specific language governing permissions and
limitations under the License.
*/

package grpcgw

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/protobuf/encoding/protojson"
)

var (
	// MetadataOpt 自定义头部
	MetadataOpt = runtime.WithMetadata(metadataHandler)
	// MarshalerOption自定义返回结构
	MarshalerOpt = runtime.WithMarshalerOption(runtime.MIMEWildcard, &jsonResponse{})

	// JsonMarshalerOpt 序列化
	JsonMarshalerOpt = runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			EmitUnpopulated: true,
			UseProtoNames:   true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})

	// ErrorHandlerOpt 自定义错误处理
	ErrorHandlerOpt = runtime.WithErrorHandler(errorHandler)
)