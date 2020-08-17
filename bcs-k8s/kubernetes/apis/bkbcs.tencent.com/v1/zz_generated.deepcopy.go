// +build !ignore_autogenerated

/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BKDataApiConfig) DeepCopyInto(out *BKDataApiConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BKDataApiConfig.
func (in *BKDataApiConfig) DeepCopy() *BKDataApiConfig {
	if in == nil {
		return nil
	}
	out := new(BKDataApiConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BKDataApiConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BKDataApiConfigList) DeepCopyInto(out *BKDataApiConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BKDataApiConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BKDataApiConfigList.
func (in *BKDataApiConfigList) DeepCopy() *BKDataApiConfigList {
	if in == nil {
		return nil
	}
	out := new(BKDataApiConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BKDataApiConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BKDataApiConfigSpec) DeepCopyInto(out *BKDataApiConfigSpec) {
	*out = *in
	out.AccessDeployPlanConfig = in.AccessDeployPlanConfig
	in.DataCleanStrategyConfig.DeepCopyInto(&out.DataCleanStrategyConfig)
	in.Response.DeepCopyInto(&out.Response)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BKDataApiConfigSpec.
func (in *BKDataApiConfigSpec) DeepCopy() *BKDataApiConfigSpec {
	if in == nil {
		return nil
	}
	out := new(BKDataApiConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BKDataApiResponse) DeepCopyInto(out *BKDataApiResponse) {
	*out = *in
	if in.Errors != nil {
		in, out := &in.Errors, &out.Errors
		*out, _ = DeepCopyMapStr(*in).(map[string]interface{})
	}
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out, _ = DeepCopyMapStr(*in).(map[string]interface{})
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BKDataApiResponse.
func (in *BKDataApiResponse) DeepCopy() *BKDataApiResponse {
	if in == nil {
		return nil
	}
	out := new(BKDataApiResponse)
	in.DeepCopyInto(out)
	return out
}

func DeepCopyMapStr(in interface{}) interface{} {
	switch value := in.(type) {
	case map[string]interface{}:
		mapstr := make(map[string]interface{})
		for k, v := range value {
			mapstr[k] = DeepCopyMapStr(v)
		}
		return mapstr
	case []interface{}:
		l := len(value)
		var arr = make([]interface{}, l)
		for k, v := range value {
			arr[k] = DeepCopyMapStr(v)
		}
		return arr
	default:
		return value
	}
}
