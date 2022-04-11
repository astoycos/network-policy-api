//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdminNetworkPolicy) DeepCopyInto(out *AdminNetworkPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdminNetworkPolicy.
func (in *AdminNetworkPolicy) DeepCopy() *AdminNetworkPolicy {
	if in == nil {
		return nil
	}
	out := new(AdminNetworkPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AdminNetworkPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdminNetworkPolicyEgressRule) DeepCopyInto(out *AdminNetworkPolicyEgressRule) {
	*out = *in
	in.Ports.DeepCopyInto(&out.Ports)
	if in.To != nil {
		in, out := &in.To, &out.To
		*out = make([]AdminNetworkPolicyPeer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdminNetworkPolicyEgressRule.
func (in *AdminNetworkPolicyEgressRule) DeepCopy() *AdminNetworkPolicyEgressRule {
	if in == nil {
		return nil
	}
	out := new(AdminNetworkPolicyEgressRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdminNetworkPolicyIngressRule) DeepCopyInto(out *AdminNetworkPolicyIngressRule) {
	*out = *in
	in.Ports.DeepCopyInto(&out.Ports)
	if in.From != nil {
		in, out := &in.From, &out.From
		*out = make([]AdminNetworkPolicyPeer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdminNetworkPolicyIngressRule.
func (in *AdminNetworkPolicyIngressRule) DeepCopy() *AdminNetworkPolicyIngressRule {
	if in == nil {
		return nil
	}
	out := new(AdminNetworkPolicyIngressRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdminNetworkPolicyList) DeepCopyInto(out *AdminNetworkPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AdminNetworkPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdminNetworkPolicyList.
func (in *AdminNetworkPolicyList) DeepCopy() *AdminNetworkPolicyList {
	if in == nil {
		return nil
	}
	out := new(AdminNetworkPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AdminNetworkPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdminNetworkPolicyPeer) DeepCopyInto(out *AdminNetworkPolicyPeer) {
	*out = *in
	if in.Namespaces != nil {
		in, out := &in.Namespaces, &out.Namespaces
		*out = new(NamespaceSet)
		(*in).DeepCopyInto(*out)
	}
	if in.NamespacedPods != nil {
		in, out := &in.NamespacedPods, &out.NamespacedPods
		*out = new(NamespaceAndPodSet)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdminNetworkPolicyPeer.
func (in *AdminNetworkPolicyPeer) DeepCopy() *AdminNetworkPolicyPeer {
	if in == nil {
		return nil
	}
	out := new(AdminNetworkPolicyPeer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdminNetworkPolicyPort) DeepCopyInto(out *AdminNetworkPolicyPort) {
	*out = *in
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(corev1.Protocol)
		**out = **in
	}
	out.Port = in.Port
	if in.EndPort != nil {
		in, out := &in.EndPort, &out.EndPort
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdminNetworkPolicyPort.
func (in *AdminNetworkPolicyPort) DeepCopy() *AdminNetworkPolicyPort {
	if in == nil {
		return nil
	}
	out := new(AdminNetworkPolicyPort)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdminNetworkPolicyPorts) DeepCopyInto(out *AdminNetworkPolicyPorts) {
	*out = *in
	if in.AllPorts != nil {
		in, out := &in.AllPorts, &out.AllPorts
		*out = new(bool)
		**out = **in
	}
	if in.List != nil {
		in, out := &in.List, &out.List
		*out = make([]AdminNetworkPolicyPort, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdminNetworkPolicyPorts.
func (in *AdminNetworkPolicyPorts) DeepCopy() *AdminNetworkPolicyPorts {
	if in == nil {
		return nil
	}
	out := new(AdminNetworkPolicyPorts)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdminNetworkPolicySpec) DeepCopyInto(out *AdminNetworkPolicySpec) {
	*out = *in
	if in.Priority != nil {
		in, out := &in.Priority, &out.Priority
		*out = new(int32)
		**out = **in
	}
	in.Subject.DeepCopyInto(&out.Subject)
	if in.Ingress != nil {
		in, out := &in.Ingress, &out.Ingress
		*out = make([]AdminNetworkPolicyIngressRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Egress != nil {
		in, out := &in.Egress, &out.Egress
		*out = make([]AdminNetworkPolicyEgressRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdminNetworkPolicySpec.
func (in *AdminNetworkPolicySpec) DeepCopy() *AdminNetworkPolicySpec {
	if in == nil {
		return nil
	}
	out := new(AdminNetworkPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdminNetworkPolicyStatus) DeepCopyInto(out *AdminNetworkPolicyStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdminNetworkPolicyStatus.
func (in *AdminNetworkPolicyStatus) DeepCopy() *AdminNetworkPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(AdminNetworkPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdminNetworkPolicySubject) DeepCopyInto(out *AdminNetworkPolicySubject) {
	*out = *in
	if in.NamespaceSelector != nil {
		in, out := &in.NamespaceSelector, &out.NamespaceSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.NamespaceAndPodSelector != nil {
		in, out := &in.NamespaceAndPodSelector, &out.NamespaceAndPodSelector
		*out = new(NamespacedPodSubject)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdminNetworkPolicySubject.
func (in *AdminNetworkPolicySubject) DeepCopy() *AdminNetworkPolicySubject {
	if in == nil {
		return nil
	}
	out := new(AdminNetworkPolicySubject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceAndPodSet) DeepCopyInto(out *NamespaceAndPodSet) {
	*out = *in
	in.Namespaces.DeepCopyInto(&out.Namespaces)
	in.Pods.DeepCopyInto(&out.Pods)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceAndPodSet.
func (in *NamespaceAndPodSet) DeepCopy() *NamespaceAndPodSet {
	if in == nil {
		return nil
	}
	out := new(NamespaceAndPodSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceSet) DeepCopyInto(out *NamespaceSet) {
	*out = *in
	if in.Self != nil {
		in, out := &in.Self, &out.Self
		*out = new(bool)
		**out = **in
	}
	if in.NotSelf != nil {
		in, out := &in.NotSelf, &out.NotSelf
		*out = new(bool)
		**out = **in
	}
	if in.NamespaceSelector != nil {
		in, out := &in.NamespaceSelector, &out.NamespaceSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.SameLabels != nil {
		in, out := &in.SameLabels, &out.SameLabels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NotSameLabels != nil {
		in, out := &in.NotSameLabels, &out.NotSameLabels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceSet.
func (in *NamespaceSet) DeepCopy() *NamespaceSet {
	if in == nil {
		return nil
	}
	out := new(NamespaceSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespacedPodSubject) DeepCopyInto(out *NamespacedPodSubject) {
	*out = *in
	in.NamespaceSelector.DeepCopyInto(&out.NamespaceSelector)
	in.PodSelector.DeepCopyInto(&out.PodSelector)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespacedPodSubject.
func (in *NamespacedPodSubject) DeepCopy() *NamespacedPodSubject {
	if in == nil {
		return nil
	}
	out := new(NamespacedPodSubject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodSet) DeepCopyInto(out *PodSet) {
	*out = *in
	if in.PodSelector != nil {
		in, out := &in.PodSelector, &out.PodSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodSet.
func (in *PodSet) DeepCopy() *PodSet {
	if in == nil {
		return nil
	}
	out := new(PodSet)
	in.DeepCopyInto(out)
	return out
}
