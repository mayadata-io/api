// +build !ignore_autogenerated

/*
Copyright 2020 The OpenEBS Authors

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CStorPoolCluster) DeepCopyInto(out *CStorPoolCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	in.VersionDetails.DeepCopyInto(&out.VersionDetails)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CStorPoolCluster.
func (in *CStorPoolCluster) DeepCopy() *CStorPoolCluster {
	if in == nil {
		return nil
	}
	out := new(CStorPoolCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CStorPoolCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CStorPoolClusterBlockDevice) DeepCopyInto(out *CStorPoolClusterBlockDevice) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CStorPoolClusterBlockDevice.
func (in *CStorPoolClusterBlockDevice) DeepCopy() *CStorPoolClusterBlockDevice {
	if in == nil {
		return nil
	}
	out := new(CStorPoolClusterBlockDevice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CStorPoolClusterCondition) DeepCopyInto(out *CStorPoolClusterCondition) {
	*out = *in
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CStorPoolClusterCondition.
func (in *CStorPoolClusterCondition) DeepCopy() *CStorPoolClusterCondition {
	if in == nil {
		return nil
	}
	out := new(CStorPoolClusterCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CStorPoolClusterList) DeepCopyInto(out *CStorPoolClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CStorPoolCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CStorPoolClusterList.
func (in *CStorPoolClusterList) DeepCopy() *CStorPoolClusterList {
	if in == nil {
		return nil
	}
	out := new(CStorPoolClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CStorPoolClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CStorPoolClusterSpec) DeepCopyInto(out *CStorPoolClusterSpec) {
	*out = *in
	if in.Pools != nil {
		in, out := &in.Pools, &out.Pools
		*out = make([]PoolSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DefaultResources != nil {
		in, out := &in.DefaultResources, &out.DefaultResources
		*out = new(corev1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.DefaultAuxResources != nil {
		in, out := &in.DefaultAuxResources, &out.DefaultAuxResources
		*out = new(corev1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]corev1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CStorPoolClusterSpec.
func (in *CStorPoolClusterSpec) DeepCopy() *CStorPoolClusterSpec {
	if in == nil {
		return nil
	}
	out := new(CStorPoolClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CStorPoolClusterStatus) DeepCopyInto(out *CStorPoolClusterStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]CStorPoolClusterCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CStorPoolClusterStatus.
func (in *CStorPoolClusterStatus) DeepCopy() *CStorPoolClusterStatus {
	if in == nil {
		return nil
	}
	out := new(CStorPoolClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CStorPoolInstacneCapacity) DeepCopyInto(out *CStorPoolInstacneCapacity) {
	*out = *in
	out.Total = in.Total.DeepCopy()
	out.Free = in.Free.DeepCopy()
	out.Used = in.Used.DeepCopy()
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CStorPoolInstacneCapacity.
func (in *CStorPoolInstacneCapacity) DeepCopy() *CStorPoolInstacneCapacity {
	if in == nil {
		return nil
	}
	out := new(CStorPoolInstacneCapacity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CStorPoolInstance) DeepCopyInto(out *CStorPoolInstance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	in.VersionDetails.DeepCopyInto(&out.VersionDetails)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CStorPoolInstance.
func (in *CStorPoolInstance) DeepCopy() *CStorPoolInstance {
	if in == nil {
		return nil
	}
	out := new(CStorPoolInstance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CStorPoolInstance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CStorPoolInstanceCondition) DeepCopyInto(out *CStorPoolInstanceCondition) {
	*out = *in
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CStorPoolInstanceCondition.
func (in *CStorPoolInstanceCondition) DeepCopy() *CStorPoolInstanceCondition {
	if in == nil {
		return nil
	}
	out := new(CStorPoolInstanceCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CStorPoolInstanceList) DeepCopyInto(out *CStorPoolInstanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CStorPoolInstance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CStorPoolInstanceList.
func (in *CStorPoolInstanceList) DeepCopy() *CStorPoolInstanceList {
	if in == nil {
		return nil
	}
	out := new(CStorPoolInstanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CStorPoolInstanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CStorPoolInstanceSpec) DeepCopyInto(out *CStorPoolInstanceSpec) {
	*out = *in
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.PoolConfig.DeepCopyInto(&out.PoolConfig)
	if in.RaidGroups != nil {
		in, out := &in.RaidGroups, &out.RaidGroups
		*out = make([]RaidGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CStorPoolInstanceSpec.
func (in *CStorPoolInstanceSpec) DeepCopy() *CStorPoolInstanceSpec {
	if in == nil {
		return nil
	}
	out := new(CStorPoolInstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CStorPoolInstanceStatus) DeepCopyInto(out *CStorPoolInstanceStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]CStorPoolInstanceCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Capacity.DeepCopyInto(&out.Capacity)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CStorPoolInstanceStatus.
func (in *CStorPoolInstanceStatus) DeepCopy() *CStorPoolInstanceStatus {
	if in == nil {
		return nil
	}
	out := new(CStorPoolInstanceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CStorVolume) DeepCopyInto(out *CStorVolume) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	in.VersionDetails.DeepCopyInto(&out.VersionDetails)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CStorVolume.
func (in *CStorVolume) DeepCopy() *CStorVolume {
	if in == nil {
		return nil
	}
	out := new(CStorVolume)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CStorVolume) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CStorVolumeCapacityAttr) DeepCopyInto(out *CStorVolumeCapacityAttr) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CStorVolumeCapacityAttr.
func (in *CStorVolumeCapacityAttr) DeepCopy() *CStorVolumeCapacityAttr {
	if in == nil {
		return nil
	}
	out := new(CStorVolumeCapacityAttr)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CStorVolumeClaim) DeepCopyInto(out *CStorVolumeClaim) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Publish = in.Publish
	in.Status.DeepCopyInto(&out.Status)
	in.VersionDetails.DeepCopyInto(&out.VersionDetails)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CStorVolumeClaim.
func (in *CStorVolumeClaim) DeepCopy() *CStorVolumeClaim {
	if in == nil {
		return nil
	}
	out := new(CStorVolumeClaim)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CStorVolumeClaim) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CStorVolumeClaimCondition) DeepCopyInto(out *CStorVolumeClaimCondition) {
	*out = *in
	in.LastProbeTime.DeepCopyInto(&out.LastProbeTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CStorVolumeClaimCondition.
func (in *CStorVolumeClaimCondition) DeepCopy() *CStorVolumeClaimCondition {
	if in == nil {
		return nil
	}
	out := new(CStorVolumeClaimCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CStorVolumeClaimList) DeepCopyInto(out *CStorVolumeClaimList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CStorVolumeClaim, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CStorVolumeClaimList.
func (in *CStorVolumeClaimList) DeepCopy() *CStorVolumeClaimList {
	if in == nil {
		return nil
	}
	out := new(CStorVolumeClaimList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CStorVolumeClaimList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CStorVolumeClaimPublish) DeepCopyInto(out *CStorVolumeClaimPublish) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CStorVolumeClaimPublish.
func (in *CStorVolumeClaimPublish) DeepCopy() *CStorVolumeClaimPublish {
	if in == nil {
		return nil
	}
	out := new(CStorVolumeClaimPublish)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CStorVolumeClaimSpec) DeepCopyInto(out *CStorVolumeClaimSpec) {
	*out = *in
	if in.Capacity != nil {
		in, out := &in.Capacity, &out.Capacity
		*out = make(corev1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	if in.CStorVolumeRef != nil {
		in, out := &in.CStorVolumeRef, &out.CStorVolumeRef
		*out = new(corev1.ObjectReference)
		**out = **in
	}
	in.Policy.DeepCopyInto(&out.Policy)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CStorVolumeClaimSpec.
func (in *CStorVolumeClaimSpec) DeepCopy() *CStorVolumeClaimSpec {
	if in == nil {
		return nil
	}
	out := new(CStorVolumeClaimSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CStorVolumeClaimStatus) DeepCopyInto(out *CStorVolumeClaimStatus) {
	*out = *in
	if in.Capacity != nil {
		in, out := &in.Capacity, &out.Capacity
		*out = make(corev1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]CStorVolumeClaimCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CStorVolumeClaimStatus.
func (in *CStorVolumeClaimStatus) DeepCopy() *CStorVolumeClaimStatus {
	if in == nil {
		return nil
	}
	out := new(CStorVolumeClaimStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CStorVolumeCondition) DeepCopyInto(out *CStorVolumeCondition) {
	*out = *in
	in.LastProbeTime.DeepCopyInto(&out.LastProbeTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CStorVolumeCondition.
func (in *CStorVolumeCondition) DeepCopy() *CStorVolumeCondition {
	if in == nil {
		return nil
	}
	out := new(CStorVolumeCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CStorVolumeList) DeepCopyInto(out *CStorVolumeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CStorVolume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CStorVolumeList.
func (in *CStorVolumeList) DeepCopy() *CStorVolumeList {
	if in == nil {
		return nil
	}
	out := new(CStorVolumeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CStorVolumeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CStorVolumePolicy) DeepCopyInto(out *CStorVolumePolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CStorVolumePolicy.
func (in *CStorVolumePolicy) DeepCopy() *CStorVolumePolicy {
	if in == nil {
		return nil
	}
	out := new(CStorVolumePolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CStorVolumePolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CStorVolumePolicyList) DeepCopyInto(out *CStorVolumePolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CStorVolumePolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CStorVolumePolicyList.
func (in *CStorVolumePolicyList) DeepCopy() *CStorVolumePolicyList {
	if in == nil {
		return nil
	}
	out := new(CStorVolumePolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CStorVolumePolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CStorVolumePolicySpec) DeepCopyInto(out *CStorVolumePolicySpec) {
	*out = *in
	out.Provision = in.Provision
	in.Target.DeepCopyInto(&out.Target)
	in.Replica.DeepCopyInto(&out.Replica)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CStorVolumePolicySpec.
func (in *CStorVolumePolicySpec) DeepCopy() *CStorVolumePolicySpec {
	if in == nil {
		return nil
	}
	out := new(CStorVolumePolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CStorVolumePolicyStatus) DeepCopyInto(out *CStorVolumePolicyStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CStorVolumePolicyStatus.
func (in *CStorVolumePolicyStatus) DeepCopy() *CStorVolumePolicyStatus {
	if in == nil {
		return nil
	}
	out := new(CStorVolumePolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CStorVolumeReplica) DeepCopyInto(out *CStorVolumeReplica) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	in.VersionDetails.DeepCopyInto(&out.VersionDetails)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CStorVolumeReplica.
func (in *CStorVolumeReplica) DeepCopy() *CStorVolumeReplica {
	if in == nil {
		return nil
	}
	out := new(CStorVolumeReplica)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CStorVolumeReplica) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CStorVolumeReplicaDetails) DeepCopyInto(out *CStorVolumeReplicaDetails) {
	*out = *in
	if in.KnownReplicas != nil {
		in, out := &in.KnownReplicas, &out.KnownReplicas
		*out = make(map[ReplicaID]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CStorVolumeReplicaDetails.
func (in *CStorVolumeReplicaDetails) DeepCopy() *CStorVolumeReplicaDetails {
	if in == nil {
		return nil
	}
	out := new(CStorVolumeReplicaDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CStorVolumeReplicaList) DeepCopyInto(out *CStorVolumeReplicaList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CStorVolumeReplica, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CStorVolumeReplicaList.
func (in *CStorVolumeReplicaList) DeepCopy() *CStorVolumeReplicaList {
	if in == nil {
		return nil
	}
	out := new(CStorVolumeReplicaList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CStorVolumeReplicaList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CStorVolumeReplicaSpec) DeepCopyInto(out *CStorVolumeReplicaSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CStorVolumeReplicaSpec.
func (in *CStorVolumeReplicaSpec) DeepCopy() *CStorVolumeReplicaSpec {
	if in == nil {
		return nil
	}
	out := new(CStorVolumeReplicaSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CStorVolumeReplicaStatus) DeepCopyInto(out *CStorVolumeReplicaStatus) {
	*out = *in
	out.Capacity = in.Capacity
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CStorVolumeReplicaStatus.
func (in *CStorVolumeReplicaStatus) DeepCopy() *CStorVolumeReplicaStatus {
	if in == nil {
		return nil
	}
	out := new(CStorVolumeReplicaStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CStorVolumeSpec) DeepCopyInto(out *CStorVolumeSpec) {
	*out = *in
	out.Capacity = in.Capacity.DeepCopy()
	in.ReplicaDetails.DeepCopyInto(&out.ReplicaDetails)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CStorVolumeSpec.
func (in *CStorVolumeSpec) DeepCopy() *CStorVolumeSpec {
	if in == nil {
		return nil
	}
	out := new(CStorVolumeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CStorVolumeStatus) DeepCopyInto(out *CStorVolumeStatus) {
	*out = *in
	if in.ReplicaStatuses != nil {
		in, out := &in.ReplicaStatuses, &out.ReplicaStatuses
		*out = make([]ReplicaStatus, len(*in))
		copy(*out, *in)
	}
	out.Capacity = in.Capacity.DeepCopy()
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]CStorVolumeCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.ReplicaDetails.DeepCopyInto(&out.ReplicaDetails)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CStorVolumeStatus.
func (in *CStorVolumeStatus) DeepCopy() *CStorVolumeStatus {
	if in == nil {
		return nil
	}
	out := new(CStorVolumeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CVStatus) DeepCopyInto(out *CVStatus) {
	*out = *in
	if in.ReplicaStatuses != nil {
		in, out := &in.ReplicaStatuses, &out.ReplicaStatuses
		*out = make([]ReplicaStatus, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CVStatus.
func (in *CVStatus) DeepCopy() *CVStatus {
	if in == nil {
		return nil
	}
	out := new(CVStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CVStatusResponse) DeepCopyInto(out *CVStatusResponse) {
	*out = *in
	if in.CVStatuses != nil {
		in, out := &in.CVStatuses, &out.CVStatuses
		*out = make([]CVStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CVStatusResponse.
func (in *CVStatusResponse) DeepCopy() *CVStatusResponse {
	if in == nil {
		return nil
	}
	out := new(CVStatusResponse)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PoolConfig) DeepCopyInto(out *PoolConfig) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(corev1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.AuxResources != nil {
		in, out := &in.AuxResources, &out.AuxResources
		*out = new(corev1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]corev1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PoolConfig.
func (in *PoolConfig) DeepCopy() *PoolConfig {
	if in == nil {
		return nil
	}
	out := new(PoolConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PoolSpec) DeepCopyInto(out *PoolSpec) {
	*out = *in
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.DataRaidGroups != nil {
		in, out := &in.DataRaidGroups, &out.DataRaidGroups
		*out = make([]RaidGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.WriteCacheRaidGroups != nil {
		in, out := &in.WriteCacheRaidGroups, &out.WriteCacheRaidGroups
		*out = make([]RaidGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.PoolConfig.DeepCopyInto(&out.PoolConfig)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PoolSpec.
func (in *PoolSpec) DeepCopy() *PoolSpec {
	if in == nil {
		return nil
	}
	out := new(PoolSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Provision) DeepCopyInto(out *Provision) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Provision.
func (in *Provision) DeepCopy() *Provision {
	if in == nil {
		return nil
	}
	out := new(Provision)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RaidGroup) DeepCopyInto(out *RaidGroup) {
	*out = *in
	if in.BlockDevices != nil {
		in, out := &in.BlockDevices, &out.BlockDevices
		*out = make([]CStorPoolClusterBlockDevice, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RaidGroup.
func (in *RaidGroup) DeepCopy() *RaidGroup {
	if in == nil {
		return nil
	}
	out := new(RaidGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicaSpec) DeepCopyInto(out *ReplicaSpec) {
	*out = *in
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(corev1.PodAffinity)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicaSpec.
func (in *ReplicaSpec) DeepCopy() *ReplicaSpec {
	if in == nil {
		return nil
	}
	out := new(ReplicaSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicaStatus) DeepCopyInto(out *ReplicaStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicaStatus.
func (in *ReplicaStatus) DeepCopy() *ReplicaStatus {
	if in == nil {
		return nil
	}
	out := new(ReplicaStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetSpec) DeepCopyInto(out *TargetSpec) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(corev1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.AuxResources != nil {
		in, out := &in.AuxResources, &out.AuxResources
		*out = new(corev1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]corev1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(corev1.PodAffinity)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetSpec.
func (in *TargetSpec) DeepCopy() *TargetSpec {
	if in == nil {
		return nil
	}
	out := new(TargetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VersionDetails) DeepCopyInto(out *VersionDetails) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VersionDetails.
func (in *VersionDetails) DeepCopy() *VersionDetails {
	if in == nil {
		return nil
	}
	out := new(VersionDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VersionStatus) DeepCopyInto(out *VersionStatus) {
	*out = *in
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VersionStatus.
func (in *VersionStatus) DeepCopy() *VersionStatus {
	if in == nil {
		return nil
	}
	out := new(VersionStatus)
	in.DeepCopyInto(out)
	return out
}
