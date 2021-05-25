package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TASPolicy) DeepCopyInto(out *TASPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status

}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TASPolicy.
func (in *TASPolicy) DeepCopy() *TASPolicy {
	if in == nil {
		return nil
	}
	out := new(TASPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TASPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TASPolicyList) DeepCopyInto(out *TASPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TASPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}

}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TASPolicyList.
func (in *TASPolicyList) DeepCopy() *TASPolicyList {
	if in == nil {
		return nil
	}
	out := new(TASPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TASPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TASPolicySpec) DeepCopyInto(out *TASPolicySpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TASPolicySpec.
func (in *TASPolicySpec) DeepCopy() *TASPolicySpec {
	if in == nil {
		return nil
	}
	out := new(TASPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TASPolicyStatus) DeepCopyInto(out *TASPolicyStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TASPolicyStatus.
func (in *TASPolicyStatus) DeepCopy() *TASPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(TASPolicyStatus)
	in.DeepCopyInto(out)
	return out
}