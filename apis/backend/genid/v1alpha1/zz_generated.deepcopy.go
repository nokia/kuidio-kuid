//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2024 Nokia.

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

package v1alpha1

import (
	commonv1alpha1 "github.com/kuidio/kuid/apis/common/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GENIDClaim) DeepCopyInto(out *GENIDClaim) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GENIDClaim.
func (in *GENIDClaim) DeepCopy() *GENIDClaim {
	if in == nil {
		return nil
	}
	out := new(GENIDClaim)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GENIDClaim) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GENIDClaimList) DeepCopyInto(out *GENIDClaimList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GENIDClaim, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GENIDClaimList.
func (in *GENIDClaimList) DeepCopy() *GENIDClaimList {
	if in == nil {
		return nil
	}
	out := new(GENIDClaimList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GENIDClaimList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GENIDClaimSpec) DeepCopyInto(out *GENIDClaimSpec) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(int64)
		**out = **in
	}
	if in.Range != nil {
		in, out := &in.Range, &out.Range
		*out = new(string)
		**out = **in
	}
	in.ClaimLabels.DeepCopyInto(&out.ClaimLabels)
	if in.Owner != nil {
		in, out := &in.Owner, &out.Owner
		*out = new(commonv1alpha1.OwnerReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GENIDClaimSpec.
func (in *GENIDClaimSpec) DeepCopy() *GENIDClaimSpec {
	if in == nil {
		return nil
	}
	out := new(GENIDClaimSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GENIDClaimStatus) DeepCopyInto(out *GENIDClaimStatus) {
	*out = *in
	in.ConditionedStatus.DeepCopyInto(&out.ConditionedStatus)
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(int64)
		**out = **in
	}
	if in.Range != nil {
		in, out := &in.Range, &out.Range
		*out = new(string)
		**out = **in
	}
	if in.ExpiryTime != nil {
		in, out := &in.ExpiryTime, &out.ExpiryTime
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GENIDClaimStatus.
func (in *GENIDClaimStatus) DeepCopy() *GENIDClaimStatus {
	if in == nil {
		return nil
	}
	out := new(GENIDClaimStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GENIDDynamicIDSyntaxValidator) DeepCopyInto(out *GENIDDynamicIDSyntaxValidator) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GENIDDynamicIDSyntaxValidator.
func (in *GENIDDynamicIDSyntaxValidator) DeepCopy() *GENIDDynamicIDSyntaxValidator {
	if in == nil {
		return nil
	}
	out := new(GENIDDynamicIDSyntaxValidator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GENIDEntry) DeepCopyInto(out *GENIDEntry) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GENIDEntry.
func (in *GENIDEntry) DeepCopy() *GENIDEntry {
	if in == nil {
		return nil
	}
	out := new(GENIDEntry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GENIDEntry) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GENIDEntryList) DeepCopyInto(out *GENIDEntryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GENIDEntry, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GENIDEntryList.
func (in *GENIDEntryList) DeepCopy() *GENIDEntryList {
	if in == nil {
		return nil
	}
	out := new(GENIDEntryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GENIDEntryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GENIDEntrySpec) DeepCopyInto(out *GENIDEntrySpec) {
	*out = *in
	in.ClaimLabels.DeepCopyInto(&out.ClaimLabels)
	if in.Owner != nil {
		in, out := &in.Owner, &out.Owner
		*out = new(commonv1alpha1.OwnerReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GENIDEntrySpec.
func (in *GENIDEntrySpec) DeepCopy() *GENIDEntrySpec {
	if in == nil {
		return nil
	}
	out := new(GENIDEntrySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GENIDEntryStatus) DeepCopyInto(out *GENIDEntryStatus) {
	*out = *in
	in.ConditionedStatus.DeepCopyInto(&out.ConditionedStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GENIDEntryStatus.
func (in *GENIDEntryStatus) DeepCopy() *GENIDEntryStatus {
	if in == nil {
		return nil
	}
	out := new(GENIDEntryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GENIDIndex) DeepCopyInto(out *GENIDIndex) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GENIDIndex.
func (in *GENIDIndex) DeepCopy() *GENIDIndex {
	if in == nil {
		return nil
	}
	out := new(GENIDIndex)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GENIDIndex) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GENIDIndexList) DeepCopyInto(out *GENIDIndexList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GENIDIndex, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GENIDIndexList.
func (in *GENIDIndexList) DeepCopy() *GENIDIndexList {
	if in == nil {
		return nil
	}
	out := new(GENIDIndexList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GENIDIndexList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GENIDIndexSpec) DeepCopyInto(out *GENIDIndexSpec) {
	*out = *in
	if in.MinID != nil {
		in, out := &in.MinID, &out.MinID
		*out = new(int64)
		**out = **in
	}
	if in.MaxID != nil {
		in, out := &in.MaxID, &out.MaxID
		*out = new(int64)
		**out = **in
	}
	in.UserDefinedLabels.DeepCopyInto(&out.UserDefinedLabels)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GENIDIndexSpec.
func (in *GENIDIndexSpec) DeepCopy() *GENIDIndexSpec {
	if in == nil {
		return nil
	}
	out := new(GENIDIndexSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GENIDIndexStatus) DeepCopyInto(out *GENIDIndexStatus) {
	*out = *in
	if in.MinID != nil {
		in, out := &in.MinID, &out.MinID
		*out = new(int64)
		**out = **in
	}
	if in.MaxID != nil {
		in, out := &in.MaxID, &out.MaxID
		*out = new(int64)
		**out = **in
	}
	in.ConditionedStatus.DeepCopyInto(&out.ConditionedStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GENIDIndexStatus.
func (in *GENIDIndexStatus) DeepCopy() *GENIDIndexStatus {
	if in == nil {
		return nil
	}
	out := new(GENIDIndexStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GENIDRangeSyntaxValidator) DeepCopyInto(out *GENIDRangeSyntaxValidator) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GENIDRangeSyntaxValidator.
func (in *GENIDRangeSyntaxValidator) DeepCopy() *GENIDRangeSyntaxValidator {
	if in == nil {
		return nil
	}
	out := new(GENIDRangeSyntaxValidator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GENIDStaticIDSyntaxValidator) DeepCopyInto(out *GENIDStaticIDSyntaxValidator) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GENIDStaticIDSyntaxValidator.
func (in *GENIDStaticIDSyntaxValidator) DeepCopy() *GENIDStaticIDSyntaxValidator {
	if in == nil {
		return nil
	}
	out := new(GENIDStaticIDSyntaxValidator)
	in.DeepCopyInto(out)
	return out
}
