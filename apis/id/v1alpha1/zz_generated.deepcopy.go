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

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterID) DeepCopyInto(out *ClusterID) {
	*out = *in
	out.SiteID = in.SiteID
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterID.
func (in *ClusterID) DeepCopy() *ClusterID {
	if in == nil {
		return nil
	}
	out := new(ClusterID)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointID) DeepCopyInto(out *EndpointID) {
	*out = *in
	out.NodeID = in.NodeID
	if in.ModuleBay != nil {
		in, out := &in.ModuleBay, &out.ModuleBay
		*out = new(int)
		**out = **in
	}
	if in.Module != nil {
		in, out := &in.Module, &out.Module
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointID.
func (in *EndpointID) DeepCopy() *EndpointID {
	if in == nil {
		return nil
	}
	out := new(EndpointID)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeID) DeepCopyInto(out *NodeID) {
	*out = *in
	out.SiteID = in.SiteID
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeID.
func (in *NodeID) DeepCopy() *NodeID {
	if in == nil {
		return nil
	}
	out := new(NodeID)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PartitionAdaptorID) DeepCopyInto(out *PartitionAdaptorID) {
	*out = *in
	in.PartitionPortID.DeepCopyInto(&out.PartitionPortID)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PartitionAdaptorID.
func (in *PartitionAdaptorID) DeepCopy() *PartitionAdaptorID {
	if in == nil {
		return nil
	}
	out := new(PartitionAdaptorID)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PartitionAttachmentID) DeepCopyInto(out *PartitionAttachmentID) {
	*out = *in
	out.SiteID = in.SiteID
	if in.Cluster != nil {
		in, out := &in.Cluster, &out.Cluster
		*out = new(string)
		**out = **in
	}
	if in.Node != nil {
		in, out := &in.Node, &out.Node
		*out = new(string)
		**out = **in
	}
	if in.NodeSet != nil {
		in, out := &in.NodeSet, &out.NodeSet
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PartitionAttachmentID.
func (in *PartitionAttachmentID) DeepCopy() *PartitionAttachmentID {
	if in == nil {
		return nil
	}
	out := new(PartitionAttachmentID)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PartitionClusterID) DeepCopyInto(out *PartitionClusterID) {
	*out = *in
	out.SiteID = in.SiteID
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PartitionClusterID.
func (in *PartitionClusterID) DeepCopy() *PartitionClusterID {
	if in == nil {
		return nil
	}
	out := new(PartitionClusterID)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PartitionEndpointID) DeepCopyInto(out *PartitionEndpointID) {
	*out = *in
	out.NodeID = in.NodeID
	if in.ModuleBay != nil {
		in, out := &in.ModuleBay, &out.ModuleBay
		*out = new(int)
		**out = **in
	}
	if in.Module != nil {
		in, out := &in.Module, &out.Module
		*out = new(int)
		**out = **in
	}
	if in.Adaptor != nil {
		in, out := &in.Adaptor, &out.Adaptor
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PartitionEndpointID.
func (in *PartitionEndpointID) DeepCopy() *PartitionEndpointID {
	if in == nil {
		return nil
	}
	out := new(PartitionEndpointID)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PartitionNodeID) DeepCopyInto(out *PartitionNodeID) {
	*out = *in
	out.SiteID = in.SiteID
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PartitionNodeID.
func (in *PartitionNodeID) DeepCopy() *PartitionNodeID {
	if in == nil {
		return nil
	}
	out := new(PartitionNodeID)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PartitionPortID) DeepCopyInto(out *PartitionPortID) {
	*out = *in
	out.PartitionNodeID = in.PartitionNodeID
	if in.ModuleBay != nil {
		in, out := &in.ModuleBay, &out.ModuleBay
		*out = new(int)
		**out = **in
	}
	if in.Module != nil {
		in, out := &in.Module, &out.Module
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PartitionPortID.
func (in *PartitionPortID) DeepCopy() *PartitionPortID {
	if in == nil {
		return nil
	}
	out := new(PartitionPortID)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SiteID) DeepCopyInto(out *SiteID) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SiteID.
func (in *SiteID) DeepCopy() *SiteID {
	if in == nil {
		return nil
	}
	out := new(SiteID)
	in.DeepCopyInto(out)
	return out
}
