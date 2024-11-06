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

package as

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ASClaim) DeepCopyInto(out *ASClaim) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ASClaim.
func (in *ASClaim) DeepCopy() *ASClaim {
	if in == nil {
		return nil
	}
	out := new(ASClaim)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ASClaim) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ASClaimFilter) DeepCopyInto(out *ASClaimFilter) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ASClaimFilter.
func (in *ASClaimFilter) DeepCopy() *ASClaimFilter {
	if in == nil {
		return nil
	}
	out := new(ASClaimFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ASClaimList) DeepCopyInto(out *ASClaimList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ASClaim, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ASClaimList.
func (in *ASClaimList) DeepCopy() *ASClaimList {
	if in == nil {
		return nil
	}
	out := new(ASClaimList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ASClaimList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ASClaimSpec) DeepCopyInto(out *ASClaimSpec) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(uint32)
		**out = **in
	}
	if in.Range != nil {
		in, out := &in.Range, &out.Range
		*out = new(string)
		**out = **in
	}
	in.ClaimLabels.DeepCopyInto(&out.ClaimLabels)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ASClaimSpec.
func (in *ASClaimSpec) DeepCopy() *ASClaimSpec {
	if in == nil {
		return nil
	}
	out := new(ASClaimSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ASClaimStatus) DeepCopyInto(out *ASClaimStatus) {
	*out = *in
	in.ConditionedStatus.DeepCopyInto(&out.ConditionedStatus)
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(uint32)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ASClaimStatus.
func (in *ASClaimStatus) DeepCopy() *ASClaimStatus {
	if in == nil {
		return nil
	}
	out := new(ASClaimStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ASDynamicIDSyntaxValidator) DeepCopyInto(out *ASDynamicIDSyntaxValidator) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ASDynamicIDSyntaxValidator.
func (in *ASDynamicIDSyntaxValidator) DeepCopy() *ASDynamicIDSyntaxValidator {
	if in == nil {
		return nil
	}
	out := new(ASDynamicIDSyntaxValidator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ASEntry) DeepCopyInto(out *ASEntry) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ASEntry.
func (in *ASEntry) DeepCopy() *ASEntry {
	if in == nil {
		return nil
	}
	out := new(ASEntry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ASEntry) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ASEntryFilter) DeepCopyInto(out *ASEntryFilter) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ASEntryFilter.
func (in *ASEntryFilter) DeepCopy() *ASEntryFilter {
	if in == nil {
		return nil
	}
	out := new(ASEntryFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ASEntryList) DeepCopyInto(out *ASEntryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ASEntry, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ASEntryList.
func (in *ASEntryList) DeepCopy() *ASEntryList {
	if in == nil {
		return nil
	}
	out := new(ASEntryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ASEntryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ASEntrySpec) DeepCopyInto(out *ASEntrySpec) {
	*out = *in
	in.ClaimLabels.DeepCopyInto(&out.ClaimLabels)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ASEntrySpec.
func (in *ASEntrySpec) DeepCopy() *ASEntrySpec {
	if in == nil {
		return nil
	}
	out := new(ASEntrySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ASEntryStatus) DeepCopyInto(out *ASEntryStatus) {
	*out = *in
	in.ConditionedStatus.DeepCopyInto(&out.ConditionedStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ASEntryStatus.
func (in *ASEntryStatus) DeepCopy() *ASEntryStatus {
	if in == nil {
		return nil
	}
	out := new(ASEntryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ASIndex) DeepCopyInto(out *ASIndex) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ASIndex.
func (in *ASIndex) DeepCopy() *ASIndex {
	if in == nil {
		return nil
	}
	out := new(ASIndex)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ASIndex) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ASIndexFilter) DeepCopyInto(out *ASIndexFilter) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ASIndexFilter.
func (in *ASIndexFilter) DeepCopy() *ASIndexFilter {
	if in == nil {
		return nil
	}
	out := new(ASIndexFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ASIndexList) DeepCopyInto(out *ASIndexList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ASIndex, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ASIndexList.
func (in *ASIndexList) DeepCopy() *ASIndexList {
	if in == nil {
		return nil
	}
	out := new(ASIndexList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ASIndexList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ASIndexSpec) DeepCopyInto(out *ASIndexSpec) {
	*out = *in
	if in.MinID != nil {
		in, out := &in.MinID, &out.MinID
		*out = new(uint32)
		**out = **in
	}
	if in.MaxID != nil {
		in, out := &in.MaxID, &out.MaxID
		*out = new(uint32)
		**out = **in
	}
	in.UserDefinedLabels.DeepCopyInto(&out.UserDefinedLabels)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ASIndexSpec.
func (in *ASIndexSpec) DeepCopy() *ASIndexSpec {
	if in == nil {
		return nil
	}
	out := new(ASIndexSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ASIndexStatus) DeepCopyInto(out *ASIndexStatus) {
	*out = *in
	if in.MinID != nil {
		in, out := &in.MinID, &out.MinID
		*out = new(uint32)
		**out = **in
	}
	if in.MaxID != nil {
		in, out := &in.MaxID, &out.MaxID
		*out = new(uint32)
		**out = **in
	}
	in.ConditionedStatus.DeepCopyInto(&out.ConditionedStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ASIndexStatus.
func (in *ASIndexStatus) DeepCopy() *ASIndexStatus {
	if in == nil {
		return nil
	}
	out := new(ASIndexStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ASRangeSyntaxValidator) DeepCopyInto(out *ASRangeSyntaxValidator) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ASRangeSyntaxValidator.
func (in *ASRangeSyntaxValidator) DeepCopy() *ASRangeSyntaxValidator {
	if in == nil {
		return nil
	}
	out := new(ASRangeSyntaxValidator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ASStaticIDSyntaxValidator) DeepCopyInto(out *ASStaticIDSyntaxValidator) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ASStaticIDSyntaxValidator.
func (in *ASStaticIDSyntaxValidator) DeepCopy() *ASStaticIDSyntaxValidator {
	if in == nil {
		return nil
	}
	out := new(ASStaticIDSyntaxValidator)
	in.DeepCopyInto(out)
	return out
}