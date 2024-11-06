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
// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	unsafe "unsafe"

	iputil "github.com/henderiw/iputil"
	asv1alpha1 "github.com/kuidio/kuid/apis/backend/as/v1alpha1"
	ipam "github.com/kuidio/kuid/apis/backend/ipam"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*IPClaim)(nil), (*ipam.IPClaim)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_IPClaim_To_ipam_IPClaim(a.(*IPClaim), b.(*ipam.IPClaim), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ipam.IPClaim)(nil), (*IPClaim)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_ipam_IPClaim_To_v1alpha1_IPClaim(a.(*ipam.IPClaim), b.(*IPClaim), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*IPClaimList)(nil), (*ipam.IPClaimList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_IPClaimList_To_ipam_IPClaimList(a.(*IPClaimList), b.(*ipam.IPClaimList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ipam.IPClaimList)(nil), (*IPClaimList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_ipam_IPClaimList_To_v1alpha1_IPClaimList(a.(*ipam.IPClaimList), b.(*IPClaimList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*IPClaimSpec)(nil), (*ipam.IPClaimSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_IPClaimSpec_To_ipam_IPClaimSpec(a.(*IPClaimSpec), b.(*ipam.IPClaimSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ipam.IPClaimSpec)(nil), (*IPClaimSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_ipam_IPClaimSpec_To_v1alpha1_IPClaimSpec(a.(*ipam.IPClaimSpec), b.(*IPClaimSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*IPClaimStatus)(nil), (*ipam.IPClaimStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_IPClaimStatus_To_ipam_IPClaimStatus(a.(*IPClaimStatus), b.(*ipam.IPClaimStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ipam.IPClaimStatus)(nil), (*IPClaimStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_ipam_IPClaimStatus_To_v1alpha1_IPClaimStatus(a.(*ipam.IPClaimStatus), b.(*IPClaimStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*IPEntry)(nil), (*ipam.IPEntry)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_IPEntry_To_ipam_IPEntry(a.(*IPEntry), b.(*ipam.IPEntry), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ipam.IPEntry)(nil), (*IPEntry)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_ipam_IPEntry_To_v1alpha1_IPEntry(a.(*ipam.IPEntry), b.(*IPEntry), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*IPEntryList)(nil), (*ipam.IPEntryList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_IPEntryList_To_ipam_IPEntryList(a.(*IPEntryList), b.(*ipam.IPEntryList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ipam.IPEntryList)(nil), (*IPEntryList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_ipam_IPEntryList_To_v1alpha1_IPEntryList(a.(*ipam.IPEntryList), b.(*IPEntryList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*IPEntrySpec)(nil), (*ipam.IPEntrySpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_IPEntrySpec_To_ipam_IPEntrySpec(a.(*IPEntrySpec), b.(*ipam.IPEntrySpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ipam.IPEntrySpec)(nil), (*IPEntrySpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_ipam_IPEntrySpec_To_v1alpha1_IPEntrySpec(a.(*ipam.IPEntrySpec), b.(*IPEntrySpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*IPEntryStatus)(nil), (*ipam.IPEntryStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_IPEntryStatus_To_ipam_IPEntryStatus(a.(*IPEntryStatus), b.(*ipam.IPEntryStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ipam.IPEntryStatus)(nil), (*IPEntryStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_ipam_IPEntryStatus_To_v1alpha1_IPEntryStatus(a.(*ipam.IPEntryStatus), b.(*IPEntryStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*IPIndex)(nil), (*ipam.IPIndex)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_IPIndex_To_ipam_IPIndex(a.(*IPIndex), b.(*ipam.IPIndex), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ipam.IPIndex)(nil), (*IPIndex)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_ipam_IPIndex_To_v1alpha1_IPIndex(a.(*ipam.IPIndex), b.(*IPIndex), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*IPIndexList)(nil), (*ipam.IPIndexList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_IPIndexList_To_ipam_IPIndexList(a.(*IPIndexList), b.(*ipam.IPIndexList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ipam.IPIndexList)(nil), (*IPIndexList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_ipam_IPIndexList_To_v1alpha1_IPIndexList(a.(*ipam.IPIndexList), b.(*IPIndexList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*IPIndexSpec)(nil), (*ipam.IPIndexSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_IPIndexSpec_To_ipam_IPIndexSpec(a.(*IPIndexSpec), b.(*ipam.IPIndexSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ipam.IPIndexSpec)(nil), (*IPIndexSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_ipam_IPIndexSpec_To_v1alpha1_IPIndexSpec(a.(*ipam.IPIndexSpec), b.(*IPIndexSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*IPIndexStatus)(nil), (*ipam.IPIndexStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_IPIndexStatus_To_ipam_IPIndexStatus(a.(*IPIndexStatus), b.(*ipam.IPIndexStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ipam.IPIndexStatus)(nil), (*IPIndexStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_ipam_IPIndexStatus_To_v1alpha1_IPIndexStatus(a.(*ipam.IPIndexStatus), b.(*IPIndexStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Prefix)(nil), (*ipam.Prefix)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Prefix_To_ipam_Prefix(a.(*Prefix), b.(*ipam.Prefix), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ipam.Prefix)(nil), (*Prefix)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_ipam_Prefix_To_v1alpha1_Prefix(a.(*ipam.Prefix), b.(*Prefix), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_IPClaim_To_ipam_IPClaim(in *IPClaim, out *ipam.IPClaim, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_IPClaimSpec_To_ipam_IPClaimSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_IPClaimStatus_To_ipam_IPClaimStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_IPClaim_To_ipam_IPClaim is an autogenerated conversion function.
func Convert_v1alpha1_IPClaim_To_ipam_IPClaim(in *IPClaim, out *ipam.IPClaim, s conversion.Scope) error {
	return autoConvert_v1alpha1_IPClaim_To_ipam_IPClaim(in, out, s)
}

func autoConvert_ipam_IPClaim_To_v1alpha1_IPClaim(in *ipam.IPClaim, out *IPClaim, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_ipam_IPClaimSpec_To_v1alpha1_IPClaimSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_ipam_IPClaimStatus_To_v1alpha1_IPClaimStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_ipam_IPClaim_To_v1alpha1_IPClaim is an autogenerated conversion function.
func Convert_ipam_IPClaim_To_v1alpha1_IPClaim(in *ipam.IPClaim, out *IPClaim, s conversion.Scope) error {
	return autoConvert_ipam_IPClaim_To_v1alpha1_IPClaim(in, out, s)
}

func autoConvert_v1alpha1_IPClaimList_To_ipam_IPClaimList(in *IPClaimList, out *ipam.IPClaimList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ipam.IPClaim, len(*in))
		for i := range *in {
			if err := Convert_v1alpha1_IPClaim_To_ipam_IPClaim(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1alpha1_IPClaimList_To_ipam_IPClaimList is an autogenerated conversion function.
func Convert_v1alpha1_IPClaimList_To_ipam_IPClaimList(in *IPClaimList, out *ipam.IPClaimList, s conversion.Scope) error {
	return autoConvert_v1alpha1_IPClaimList_To_ipam_IPClaimList(in, out, s)
}

func autoConvert_ipam_IPClaimList_To_v1alpha1_IPClaimList(in *ipam.IPClaimList, out *IPClaimList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IPClaim, len(*in))
		for i := range *in {
			if err := Convert_ipam_IPClaim_To_v1alpha1_IPClaim(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_ipam_IPClaimList_To_v1alpha1_IPClaimList is an autogenerated conversion function.
func Convert_ipam_IPClaimList_To_v1alpha1_IPClaimList(in *ipam.IPClaimList, out *IPClaimList, s conversion.Scope) error {
	return autoConvert_ipam_IPClaimList_To_v1alpha1_IPClaimList(in, out, s)
}

func autoConvert_v1alpha1_IPClaimSpec_To_ipam_IPClaimSpec(in *IPClaimSpec, out *ipam.IPClaimSpec, s conversion.Scope) error {
	out.Index = in.Index
	out.PrefixType = (*ipam.IPPrefixType)(unsafe.Pointer(in.PrefixType))
	out.Prefix = (*string)(unsafe.Pointer(in.Prefix))
	out.Address = (*string)(unsafe.Pointer(in.Address))
	out.Range = (*string)(unsafe.Pointer(in.Range))
	out.DefaultGateway = (*bool)(unsafe.Pointer(in.DefaultGateway))
	out.CreatePrefix = (*bool)(unsafe.Pointer(in.CreatePrefix))
	out.PrefixLength = (*uint32)(unsafe.Pointer(in.PrefixLength))
	out.AddressFamily = (*iputil.AddressFamily)(unsafe.Pointer(in.AddressFamily))
	out.Idx = (*uint32)(unsafe.Pointer(in.Idx))
	if err := asv1alpha1.Convert_v1alpha1_ClaimLabels_To_common_ClaimLabels(&in.ClaimLabels, &out.ClaimLabels, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_IPClaimSpec_To_ipam_IPClaimSpec is an autogenerated conversion function.
func Convert_v1alpha1_IPClaimSpec_To_ipam_IPClaimSpec(in *IPClaimSpec, out *ipam.IPClaimSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_IPClaimSpec_To_ipam_IPClaimSpec(in, out, s)
}

func autoConvert_ipam_IPClaimSpec_To_v1alpha1_IPClaimSpec(in *ipam.IPClaimSpec, out *IPClaimSpec, s conversion.Scope) error {
	out.Index = in.Index
	out.PrefixType = (*IPPrefixType)(unsafe.Pointer(in.PrefixType))
	out.Prefix = (*string)(unsafe.Pointer(in.Prefix))
	out.Address = (*string)(unsafe.Pointer(in.Address))
	out.Range = (*string)(unsafe.Pointer(in.Range))
	out.DefaultGateway = (*bool)(unsafe.Pointer(in.DefaultGateway))
	out.CreatePrefix = (*bool)(unsafe.Pointer(in.CreatePrefix))
	out.PrefixLength = (*uint32)(unsafe.Pointer(in.PrefixLength))
	out.AddressFamily = (*iputil.AddressFamily)(unsafe.Pointer(in.AddressFamily))
	out.Idx = (*uint32)(unsafe.Pointer(in.Idx))
	if err := asv1alpha1.Convert_common_ClaimLabels_To_v1alpha1_ClaimLabels(&in.ClaimLabels, &out.ClaimLabels, s); err != nil {
		return err
	}
	return nil
}

// Convert_ipam_IPClaimSpec_To_v1alpha1_IPClaimSpec is an autogenerated conversion function.
func Convert_ipam_IPClaimSpec_To_v1alpha1_IPClaimSpec(in *ipam.IPClaimSpec, out *IPClaimSpec, s conversion.Scope) error {
	return autoConvert_ipam_IPClaimSpec_To_v1alpha1_IPClaimSpec(in, out, s)
}

func autoConvert_v1alpha1_IPClaimStatus_To_ipam_IPClaimStatus(in *IPClaimStatus, out *ipam.IPClaimStatus, s conversion.Scope) error {
	if err := asv1alpha1.Convert_v1alpha1_ConditionedStatus_To_condition_ConditionedStatus(&in.ConditionedStatus, &out.ConditionedStatus, s); err != nil {
		return err
	}
	out.Range = (*string)(unsafe.Pointer(in.Range))
	out.Address = (*string)(unsafe.Pointer(in.Address))
	out.Prefix = (*string)(unsafe.Pointer(in.Prefix))
	out.DefaultGateway = (*string)(unsafe.Pointer(in.DefaultGateway))
	out.ExpiryTime = (*string)(unsafe.Pointer(in.ExpiryTime))
	return nil
}

// Convert_v1alpha1_IPClaimStatus_To_ipam_IPClaimStatus is an autogenerated conversion function.
func Convert_v1alpha1_IPClaimStatus_To_ipam_IPClaimStatus(in *IPClaimStatus, out *ipam.IPClaimStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_IPClaimStatus_To_ipam_IPClaimStatus(in, out, s)
}

func autoConvert_ipam_IPClaimStatus_To_v1alpha1_IPClaimStatus(in *ipam.IPClaimStatus, out *IPClaimStatus, s conversion.Scope) error {
	if err := asv1alpha1.Convert_condition_ConditionedStatus_To_v1alpha1_ConditionedStatus(&in.ConditionedStatus, &out.ConditionedStatus, s); err != nil {
		return err
	}
	out.Range = (*string)(unsafe.Pointer(in.Range))
	out.Address = (*string)(unsafe.Pointer(in.Address))
	out.Prefix = (*string)(unsafe.Pointer(in.Prefix))
	out.DefaultGateway = (*string)(unsafe.Pointer(in.DefaultGateway))
	out.ExpiryTime = (*string)(unsafe.Pointer(in.ExpiryTime))
	return nil
}

// Convert_ipam_IPClaimStatus_To_v1alpha1_IPClaimStatus is an autogenerated conversion function.
func Convert_ipam_IPClaimStatus_To_v1alpha1_IPClaimStatus(in *ipam.IPClaimStatus, out *IPClaimStatus, s conversion.Scope) error {
	return autoConvert_ipam_IPClaimStatus_To_v1alpha1_IPClaimStatus(in, out, s)
}

func autoConvert_v1alpha1_IPEntry_To_ipam_IPEntry(in *IPEntry, out *ipam.IPEntry, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_IPEntrySpec_To_ipam_IPEntrySpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_IPEntryStatus_To_ipam_IPEntryStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_IPEntry_To_ipam_IPEntry is an autogenerated conversion function.
func Convert_v1alpha1_IPEntry_To_ipam_IPEntry(in *IPEntry, out *ipam.IPEntry, s conversion.Scope) error {
	return autoConvert_v1alpha1_IPEntry_To_ipam_IPEntry(in, out, s)
}

func autoConvert_ipam_IPEntry_To_v1alpha1_IPEntry(in *ipam.IPEntry, out *IPEntry, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_ipam_IPEntrySpec_To_v1alpha1_IPEntrySpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_ipam_IPEntryStatus_To_v1alpha1_IPEntryStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_ipam_IPEntry_To_v1alpha1_IPEntry is an autogenerated conversion function.
func Convert_ipam_IPEntry_To_v1alpha1_IPEntry(in *ipam.IPEntry, out *IPEntry, s conversion.Scope) error {
	return autoConvert_ipam_IPEntry_To_v1alpha1_IPEntry(in, out, s)
}

func autoConvert_v1alpha1_IPEntryList_To_ipam_IPEntryList(in *IPEntryList, out *ipam.IPEntryList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ipam.IPEntry, len(*in))
		for i := range *in {
			if err := Convert_v1alpha1_IPEntry_To_ipam_IPEntry(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1alpha1_IPEntryList_To_ipam_IPEntryList is an autogenerated conversion function.
func Convert_v1alpha1_IPEntryList_To_ipam_IPEntryList(in *IPEntryList, out *ipam.IPEntryList, s conversion.Scope) error {
	return autoConvert_v1alpha1_IPEntryList_To_ipam_IPEntryList(in, out, s)
}

func autoConvert_ipam_IPEntryList_To_v1alpha1_IPEntryList(in *ipam.IPEntryList, out *IPEntryList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IPEntry, len(*in))
		for i := range *in {
			if err := Convert_ipam_IPEntry_To_v1alpha1_IPEntry(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_ipam_IPEntryList_To_v1alpha1_IPEntryList is an autogenerated conversion function.
func Convert_ipam_IPEntryList_To_v1alpha1_IPEntryList(in *ipam.IPEntryList, out *IPEntryList, s conversion.Scope) error {
	return autoConvert_ipam_IPEntryList_To_v1alpha1_IPEntryList(in, out, s)
}

func autoConvert_v1alpha1_IPEntrySpec_To_ipam_IPEntrySpec(in *IPEntrySpec, out *ipam.IPEntrySpec, s conversion.Scope) error {
	out.Index = in.Index
	out.IndexEntry = in.IndexEntry
	out.PrefixType = (*ipam.IPPrefixType)(unsafe.Pointer(in.PrefixType))
	out.ClaimType = ipam.IPClaimType(in.ClaimType)
	out.Prefix = in.Prefix
	out.DefaultGateway = (*bool)(unsafe.Pointer(in.DefaultGateway))
	out.AddressFamily = (*iputil.AddressFamily)(unsafe.Pointer(in.AddressFamily))
	if err := asv1alpha1.Convert_v1alpha1_UserDefinedLabels_To_common_UserDefinedLabels(&in.UserDefinedLabels, &out.UserDefinedLabels, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_IPEntrySpec_To_ipam_IPEntrySpec is an autogenerated conversion function.
func Convert_v1alpha1_IPEntrySpec_To_ipam_IPEntrySpec(in *IPEntrySpec, out *ipam.IPEntrySpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_IPEntrySpec_To_ipam_IPEntrySpec(in, out, s)
}

func autoConvert_ipam_IPEntrySpec_To_v1alpha1_IPEntrySpec(in *ipam.IPEntrySpec, out *IPEntrySpec, s conversion.Scope) error {
	out.Index = in.Index
	out.IndexEntry = in.IndexEntry
	out.PrefixType = (*IPPrefixType)(unsafe.Pointer(in.PrefixType))
	out.ClaimType = IPClaimType(in.ClaimType)
	out.Prefix = in.Prefix
	out.DefaultGateway = (*bool)(unsafe.Pointer(in.DefaultGateway))
	out.AddressFamily = (*iputil.AddressFamily)(unsafe.Pointer(in.AddressFamily))
	if err := asv1alpha1.Convert_common_UserDefinedLabels_To_v1alpha1_UserDefinedLabels(&in.UserDefinedLabels, &out.UserDefinedLabels, s); err != nil {
		return err
	}
	return nil
}

// Convert_ipam_IPEntrySpec_To_v1alpha1_IPEntrySpec is an autogenerated conversion function.
func Convert_ipam_IPEntrySpec_To_v1alpha1_IPEntrySpec(in *ipam.IPEntrySpec, out *IPEntrySpec, s conversion.Scope) error {
	return autoConvert_ipam_IPEntrySpec_To_v1alpha1_IPEntrySpec(in, out, s)
}

func autoConvert_v1alpha1_IPEntryStatus_To_ipam_IPEntryStatus(in *IPEntryStatus, out *ipam.IPEntryStatus, s conversion.Scope) error {
	if err := asv1alpha1.Convert_v1alpha1_ConditionedStatus_To_condition_ConditionedStatus(&in.ConditionedStatus, &out.ConditionedStatus, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_IPEntryStatus_To_ipam_IPEntryStatus is an autogenerated conversion function.
func Convert_v1alpha1_IPEntryStatus_To_ipam_IPEntryStatus(in *IPEntryStatus, out *ipam.IPEntryStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_IPEntryStatus_To_ipam_IPEntryStatus(in, out, s)
}

func autoConvert_ipam_IPEntryStatus_To_v1alpha1_IPEntryStatus(in *ipam.IPEntryStatus, out *IPEntryStatus, s conversion.Scope) error {
	if err := asv1alpha1.Convert_condition_ConditionedStatus_To_v1alpha1_ConditionedStatus(&in.ConditionedStatus, &out.ConditionedStatus, s); err != nil {
		return err
	}
	return nil
}

// Convert_ipam_IPEntryStatus_To_v1alpha1_IPEntryStatus is an autogenerated conversion function.
func Convert_ipam_IPEntryStatus_To_v1alpha1_IPEntryStatus(in *ipam.IPEntryStatus, out *IPEntryStatus, s conversion.Scope) error {
	return autoConvert_ipam_IPEntryStatus_To_v1alpha1_IPEntryStatus(in, out, s)
}

func autoConvert_v1alpha1_IPIndex_To_ipam_IPIndex(in *IPIndex, out *ipam.IPIndex, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_IPIndexSpec_To_ipam_IPIndexSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_IPIndexStatus_To_ipam_IPIndexStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_IPIndex_To_ipam_IPIndex is an autogenerated conversion function.
func Convert_v1alpha1_IPIndex_To_ipam_IPIndex(in *IPIndex, out *ipam.IPIndex, s conversion.Scope) error {
	return autoConvert_v1alpha1_IPIndex_To_ipam_IPIndex(in, out, s)
}

func autoConvert_ipam_IPIndex_To_v1alpha1_IPIndex(in *ipam.IPIndex, out *IPIndex, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_ipam_IPIndexSpec_To_v1alpha1_IPIndexSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_ipam_IPIndexStatus_To_v1alpha1_IPIndexStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_ipam_IPIndex_To_v1alpha1_IPIndex is an autogenerated conversion function.
func Convert_ipam_IPIndex_To_v1alpha1_IPIndex(in *ipam.IPIndex, out *IPIndex, s conversion.Scope) error {
	return autoConvert_ipam_IPIndex_To_v1alpha1_IPIndex(in, out, s)
}

func autoConvert_v1alpha1_IPIndexList_To_ipam_IPIndexList(in *IPIndexList, out *ipam.IPIndexList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ipam.IPIndex, len(*in))
		for i := range *in {
			if err := Convert_v1alpha1_IPIndex_To_ipam_IPIndex(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1alpha1_IPIndexList_To_ipam_IPIndexList is an autogenerated conversion function.
func Convert_v1alpha1_IPIndexList_To_ipam_IPIndexList(in *IPIndexList, out *ipam.IPIndexList, s conversion.Scope) error {
	return autoConvert_v1alpha1_IPIndexList_To_ipam_IPIndexList(in, out, s)
}

func autoConvert_ipam_IPIndexList_To_v1alpha1_IPIndexList(in *ipam.IPIndexList, out *IPIndexList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IPIndex, len(*in))
		for i := range *in {
			if err := Convert_ipam_IPIndex_To_v1alpha1_IPIndex(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_ipam_IPIndexList_To_v1alpha1_IPIndexList is an autogenerated conversion function.
func Convert_ipam_IPIndexList_To_v1alpha1_IPIndexList(in *ipam.IPIndexList, out *IPIndexList, s conversion.Scope) error {
	return autoConvert_ipam_IPIndexList_To_v1alpha1_IPIndexList(in, out, s)
}

func autoConvert_v1alpha1_IPIndexSpec_To_ipam_IPIndexSpec(in *IPIndexSpec, out *ipam.IPIndexSpec, s conversion.Scope) error {
	if in.Prefixes != nil {
		in, out := &in.Prefixes, &out.Prefixes
		*out = make([]ipam.Prefix, len(*in))
		for i := range *in {
			if err := Convert_v1alpha1_Prefix_To_ipam_Prefix(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Prefixes = nil
	}
	return nil
}

// Convert_v1alpha1_IPIndexSpec_To_ipam_IPIndexSpec is an autogenerated conversion function.
func Convert_v1alpha1_IPIndexSpec_To_ipam_IPIndexSpec(in *IPIndexSpec, out *ipam.IPIndexSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_IPIndexSpec_To_ipam_IPIndexSpec(in, out, s)
}

func autoConvert_ipam_IPIndexSpec_To_v1alpha1_IPIndexSpec(in *ipam.IPIndexSpec, out *IPIndexSpec, s conversion.Scope) error {
	if in.Prefixes != nil {
		in, out := &in.Prefixes, &out.Prefixes
		*out = make([]Prefix, len(*in))
		for i := range *in {
			if err := Convert_ipam_Prefix_To_v1alpha1_Prefix(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Prefixes = nil
	}
	return nil
}

// Convert_ipam_IPIndexSpec_To_v1alpha1_IPIndexSpec is an autogenerated conversion function.
func Convert_ipam_IPIndexSpec_To_v1alpha1_IPIndexSpec(in *ipam.IPIndexSpec, out *IPIndexSpec, s conversion.Scope) error {
	return autoConvert_ipam_IPIndexSpec_To_v1alpha1_IPIndexSpec(in, out, s)
}

func autoConvert_v1alpha1_IPIndexStatus_To_ipam_IPIndexStatus(in *IPIndexStatus, out *ipam.IPIndexStatus, s conversion.Scope) error {
	if err := asv1alpha1.Convert_v1alpha1_ConditionedStatus_To_condition_ConditionedStatus(&in.ConditionedStatus, &out.ConditionedStatus, s); err != nil {
		return err
	}
	if in.Prefixes != nil {
		in, out := &in.Prefixes, &out.Prefixes
		*out = make([]ipam.Prefix, len(*in))
		for i := range *in {
			if err := Convert_v1alpha1_Prefix_To_ipam_Prefix(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Prefixes = nil
	}
	return nil
}

// Convert_v1alpha1_IPIndexStatus_To_ipam_IPIndexStatus is an autogenerated conversion function.
func Convert_v1alpha1_IPIndexStatus_To_ipam_IPIndexStatus(in *IPIndexStatus, out *ipam.IPIndexStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_IPIndexStatus_To_ipam_IPIndexStatus(in, out, s)
}

func autoConvert_ipam_IPIndexStatus_To_v1alpha1_IPIndexStatus(in *ipam.IPIndexStatus, out *IPIndexStatus, s conversion.Scope) error {
	if err := asv1alpha1.Convert_condition_ConditionedStatus_To_v1alpha1_ConditionedStatus(&in.ConditionedStatus, &out.ConditionedStatus, s); err != nil {
		return err
	}
	if in.Prefixes != nil {
		in, out := &in.Prefixes, &out.Prefixes
		*out = make([]Prefix, len(*in))
		for i := range *in {
			if err := Convert_ipam_Prefix_To_v1alpha1_Prefix(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Prefixes = nil
	}
	return nil
}

// Convert_ipam_IPIndexStatus_To_v1alpha1_IPIndexStatus is an autogenerated conversion function.
func Convert_ipam_IPIndexStatus_To_v1alpha1_IPIndexStatus(in *ipam.IPIndexStatus, out *IPIndexStatus, s conversion.Scope) error {
	return autoConvert_ipam_IPIndexStatus_To_v1alpha1_IPIndexStatus(in, out, s)
}

func autoConvert_v1alpha1_Prefix_To_ipam_Prefix(in *Prefix, out *ipam.Prefix, s conversion.Scope) error {
	out.Prefix = in.Prefix
	out.PrefixType = (*ipam.IPPrefixType)(unsafe.Pointer(in.PrefixType))
	if err := asv1alpha1.Convert_v1alpha1_UserDefinedLabels_To_common_UserDefinedLabels(&in.UserDefinedLabels, &out.UserDefinedLabels, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_Prefix_To_ipam_Prefix is an autogenerated conversion function.
func Convert_v1alpha1_Prefix_To_ipam_Prefix(in *Prefix, out *ipam.Prefix, s conversion.Scope) error {
	return autoConvert_v1alpha1_Prefix_To_ipam_Prefix(in, out, s)
}

func autoConvert_ipam_Prefix_To_v1alpha1_Prefix(in *ipam.Prefix, out *Prefix, s conversion.Scope) error {
	out.Prefix = in.Prefix
	out.PrefixType = (*IPPrefixType)(unsafe.Pointer(in.PrefixType))
	if err := asv1alpha1.Convert_common_UserDefinedLabels_To_v1alpha1_UserDefinedLabels(&in.UserDefinedLabels, &out.UserDefinedLabels, s); err != nil {
		return err
	}
	return nil
}

// Convert_ipam_Prefix_To_v1alpha1_Prefix is an autogenerated conversion function.
func Convert_ipam_Prefix_To_v1alpha1_Prefix(in *ipam.Prefix, out *Prefix, s conversion.Scope) error {
	return autoConvert_ipam_Prefix_To_v1alpha1_Prefix(in, out, s)
}