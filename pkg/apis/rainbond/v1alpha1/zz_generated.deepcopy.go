//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// RAINBOND, Application Management Platform
// Copyright (C) 2014-2021 Goodrain Co., Ltd.

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version. For any non-GPL usage of Rainbond,
// one or multiple Commercial Licenses authorized by Goodrain Co., Ltd.
// must be obtained first.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/oam-dev/kubevela/apis/core.oam.dev/common"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentDefinition) DeepCopyInto(out *ComponentDefinition) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentDefinition.
func (in *ComponentDefinition) DeepCopy() *ComponentDefinition {
	if in == nil {
		return nil
	}
	out := new(ComponentDefinition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ComponentDefinition) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentDefinitionList) DeepCopyInto(out *ComponentDefinitionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ComponentDefinition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentDefinitionList.
func (in *ComponentDefinitionList) DeepCopy() *ComponentDefinitionList {
	if in == nil {
		return nil
	}
	out := new(ComponentDefinitionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ComponentDefinitionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentDefinitionSpec) DeepCopyInto(out *ComponentDefinitionSpec) {
	*out = *in
	out.Workload = in.Workload
	if in.ChildResourceKinds != nil {
		in, out := &in.ChildResourceKinds, &out.ChildResourceKinds
		*out = make([]common.ChildResourceKind, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(common.Status)
		**out = **in
	}
	if in.Schematic != nil {
		in, out := &in.Schematic, &out.Schematic
		*out = new(Schematic)
		(*in).DeepCopyInto(*out)
	}
	if in.Extension != nil {
		in, out := &in.Extension, &out.Extension
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentDefinitionSpec.
func (in *ComponentDefinitionSpec) DeepCopy() *ComponentDefinitionSpec {
	if in == nil {
		return nil
	}
	out := new(ComponentDefinitionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentDefinitionStatus) DeepCopyInto(out *ComponentDefinitionStatus) {
	*out = *in
	in.ConditionedStatus.DeepCopyInto(&out.ConditionedStatus)
	if in.LatestRevision != nil {
		in, out := &in.LatestRevision, &out.LatestRevision
		*out = new(common.Revision)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentDefinitionStatus.
func (in *ComponentDefinitionStatus) DeepCopy() *ComponentDefinitionStatus {
	if in == nil {
		return nil
	}
	out := new(ComponentDefinitionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentPort) DeepCopyInto(out *ComponentPort) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentPort.
func (in *ComponentPort) DeepCopy() *ComponentPort {
	if in == nil {
		return nil
	}
	out := new(ComponentPort)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPGetAction) DeepCopyInto(out *HTTPGetAction) {
	*out = *in
	if in.HTTPHeaders != nil {
		in, out := &in.HTTPHeaders, &out.HTTPHeaders
		*out = make([]HTTPHeader, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPGetAction.
func (in *HTTPGetAction) DeepCopy() *HTTPGetAction {
	if in == nil {
		return nil
	}
	out := new(HTTPGetAction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPHeader) DeepCopyInto(out *HTTPHeader) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPHeader.
func (in *HTTPHeader) DeepCopy() *HTTPHeader {
	if in == nil {
		return nil
	}
	out := new(HTTPHeader)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Handler) DeepCopyInto(out *Handler) {
	*out = *in
	if in.HTTPGet != nil {
		in, out := &in.HTTPGet, &out.HTTPGet
		*out = new(HTTPGetAction)
		(*in).DeepCopyInto(*out)
	}
	if in.TCPSocket != nil {
		in, out := &in.TCPSocket, &out.TCPSocket
		*out = new(TCPSocketAction)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Handler.
func (in *Handler) DeepCopy() *Handler {
	if in == nil {
		return nil
	}
	out := new(Handler)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmApp) DeepCopyInto(out *HelmApp) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmApp.
func (in *HelmApp) DeepCopy() *HelmApp {
	if in == nil {
		return nil
	}
	out := new(HelmApp)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HelmApp) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmAppCondition) DeepCopyInto(out *HelmAppCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmAppCondition.
func (in *HelmAppCondition) DeepCopy() *HelmAppCondition {
	if in == nil {
		return nil
	}
	out := new(HelmAppCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmAppList) DeepCopyInto(out *HelmAppList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HelmApp, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmAppList.
func (in *HelmAppList) DeepCopy() *HelmAppList {
	if in == nil {
		return nil
	}
	out := new(HelmAppList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HelmAppList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmAppSpec) DeepCopyInto(out *HelmAppSpec) {
	*out = *in
	if in.AppStore != nil {
		in, out := &in.AppStore, &out.AppStore
		*out = new(HelmAppStore)
		**out = **in
	}
	if in.Overrides != nil {
		in, out := &in.Overrides, &out.Overrides
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmAppSpec.
func (in *HelmAppSpec) DeepCopy() *HelmAppSpec {
	if in == nil {
		return nil
	}
	out := new(HelmAppSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmAppStatus) DeepCopyInto(out *HelmAppStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]HelmAppCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Overrides != nil {
		in, out := &in.Overrides, &out.Overrides
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmAppStatus.
func (in *HelmAppStatus) DeepCopy() *HelmAppStatus {
	if in == nil {
		return nil
	}
	out := new(HelmAppStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmAppStore) DeepCopyInto(out *HelmAppStore) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmAppStore.
func (in *HelmAppStore) DeepCopy() *HelmAppStore {
	if in == nil {
		return nil
	}
	out := new(HelmAppStore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesServiceSource) DeepCopyInto(out *KubernetesServiceSource) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesServiceSource.
func (in *KubernetesServiceSource) DeepCopy() *KubernetesServiceSource {
	if in == nil {
		return nil
	}
	out := new(KubernetesServiceSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Probe) DeepCopyInto(out *Probe) {
	*out = *in
	in.Handler.DeepCopyInto(&out.Handler)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Probe.
func (in *Probe) DeepCopy() *Probe {
	if in == nil {
		return nil
	}
	out := new(Probe)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RBDAbility) DeepCopyInto(out *RBDAbility) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RBDAbility.
func (in *RBDAbility) DeepCopy() *RBDAbility {
	if in == nil {
		return nil
	}
	out := new(RBDAbility)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RBDAbility) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RBDAbilityList) DeepCopyInto(out *RBDAbilityList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RBDAbility, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RBDAbilityList.
func (in *RBDAbilityList) DeepCopy() *RBDAbilityList {
	if in == nil {
		return nil
	}
	out := new(RBDAbilityList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RBDAbilityList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RBDAbilitySpec) DeepCopyInto(out *RBDAbilitySpec) {
	*out = *in
	if in.WatchGroups != nil {
		in, out := &in.WatchGroups, &out.WatchGroups
		*out = make([]WatchGroup, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RBDAbilitySpec.
func (in *RBDAbilitySpec) DeepCopy() *RBDAbilitySpec {
	if in == nil {
		return nil
	}
	out := new(RBDAbilitySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RBDAbilityStatus) DeepCopyInto(out *RBDAbilityStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RBDAbilityStatus.
func (in *RBDAbilityStatus) DeepCopy() *RBDAbilityStatus {
	if in == nil {
		return nil
	}
	out := new(RBDAbilityStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RBDPlugin) DeepCopyInto(out *RBDPlugin) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RBDPlugin.
func (in *RBDPlugin) DeepCopy() *RBDPlugin {
	if in == nil {
		return nil
	}
	out := new(RBDPlugin)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RBDPlugin) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RBDPluginList) DeepCopyInto(out *RBDPluginList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RBDPlugin, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RBDPluginList.
func (in *RBDPluginList) DeepCopy() *RBDPluginList {
	if in == nil {
		return nil
	}
	out := new(RBDPluginList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RBDPluginList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RBDPluginSpec) DeepCopyInto(out *RBDPluginSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RBDPluginSpec.
func (in *RBDPluginSpec) DeepCopy() *RBDPluginSpec {
	if in == nil {
		return nil
	}
	out := new(RBDPluginSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RBDPluginStatus) DeepCopyInto(out *RBDPluginStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RBDPluginStatus.
func (in *RBDPluginStatus) DeepCopy() *RBDPluginStatus {
	if in == nil {
		return nil
	}
	out := new(RBDPluginStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Schematic) DeepCopyInto(out *Schematic) {
	*out = *in
	if in.CUE != nil {
		in, out := &in.CUE, &out.CUE
		*out = new(common.CUE)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Schematic.
func (in *Schematic) DeepCopy() *Schematic {
	if in == nil {
		return nil
	}
	out := new(Schematic)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TCPSocketAction) DeepCopyInto(out *TCPSocketAction) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TCPSocketAction.
func (in *TCPSocketAction) DeepCopy() *TCPSocketAction {
	if in == nil {
		return nil
	}
	out := new(TCPSocketAction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThirdComponent) DeepCopyInto(out *ThirdComponent) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThirdComponent.
func (in *ThirdComponent) DeepCopy() *ThirdComponent {
	if in == nil {
		return nil
	}
	out := new(ThirdComponent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ThirdComponent) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThirdComponentEndpoint) DeepCopyInto(out *ThirdComponentEndpoint) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThirdComponentEndpoint.
func (in *ThirdComponentEndpoint) DeepCopy() *ThirdComponentEndpoint {
	if in == nil {
		return nil
	}
	out := new(ThirdComponentEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThirdComponentEndpointSource) DeepCopyInto(out *ThirdComponentEndpointSource) {
	*out = *in
	if in.StaticEndpoints != nil {
		in, out := &in.StaticEndpoints, &out.StaticEndpoints
		*out = make([]*ThirdComponentEndpoint, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ThirdComponentEndpoint)
				**out = **in
			}
		}
	}
	if in.KubernetesService != nil {
		in, out := &in.KubernetesService, &out.KubernetesService
		*out = new(KubernetesServiceSource)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThirdComponentEndpointSource.
func (in *ThirdComponentEndpointSource) DeepCopy() *ThirdComponentEndpointSource {
	if in == nil {
		return nil
	}
	out := new(ThirdComponentEndpointSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThirdComponentEndpointStatus) DeepCopyInto(out *ThirdComponentEndpointStatus) {
	*out = *in
	if in.TargetRef != nil {
		in, out := &in.TargetRef, &out.TargetRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThirdComponentEndpointStatus.
func (in *ThirdComponentEndpointStatus) DeepCopy() *ThirdComponentEndpointStatus {
	if in == nil {
		return nil
	}
	out := new(ThirdComponentEndpointStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThirdComponentList) DeepCopyInto(out *ThirdComponentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ThirdComponent, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThirdComponentList.
func (in *ThirdComponentList) DeepCopy() *ThirdComponentList {
	if in == nil {
		return nil
	}
	out := new(ThirdComponentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ThirdComponentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThirdComponentSpec) DeepCopyInto(out *ThirdComponentSpec) {
	*out = *in
	if in.Probe != nil {
		in, out := &in.Probe, &out.Probe
		*out = new(Probe)
		(*in).DeepCopyInto(*out)
	}
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]*ComponentPort, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ComponentPort)
				**out = **in
			}
		}
	}
	in.EndpointSource.DeepCopyInto(&out.EndpointSource)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThirdComponentSpec.
func (in *ThirdComponentSpec) DeepCopy() *ThirdComponentSpec {
	if in == nil {
		return nil
	}
	out := new(ThirdComponentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThirdComponentStatus) DeepCopyInto(out *ThirdComponentStatus) {
	*out = *in
	if in.Endpoints != nil {
		in, out := &in.Endpoints, &out.Endpoints
		*out = make([]*ThirdComponentEndpointStatus, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ThirdComponentEndpointStatus)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThirdComponentStatus.
func (in *ThirdComponentStatus) DeepCopy() *ThirdComponentStatus {
	if in == nil {
		return nil
	}
	out := new(ThirdComponentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WatchGroup) DeepCopyInto(out *WatchGroup) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WatchGroup.
func (in *WatchGroup) DeepCopy() *WatchGroup {
	if in == nil {
		return nil
	}
	out := new(WatchGroup)
	in.DeepCopyInto(out)
	return out
}
