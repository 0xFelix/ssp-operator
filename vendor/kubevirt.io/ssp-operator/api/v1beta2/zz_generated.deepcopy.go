//go:build !ignore_autogenerated

/*


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

package v1beta2

import (
	"github.com/openshift/api/config/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommonInstancetypes) DeepCopyInto(out *CommonInstancetypes) {
	*out = *in
	if in.URL != nil {
		in, out := &in.URL, &out.URL
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommonInstancetypes.
func (in *CommonInstancetypes) DeepCopy() *CommonInstancetypes {
	if in == nil {
		return nil
	}
	out := new(CommonInstancetypes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommonTemplates) DeepCopyInto(out *CommonTemplates) {
	*out = *in
	if in.DataImportCronTemplates != nil {
		in, out := &in.DataImportCronTemplates, &out.DataImportCronTemplates
		*out = make([]DataImportCronTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommonTemplates.
func (in *CommonTemplates) DeepCopy() *CommonTemplates {
	if in == nil {
		return nil
	}
	out := new(CommonTemplates)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataImportCronTemplate) DeepCopyInto(out *DataImportCronTemplate) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataImportCronTemplate.
func (in *DataImportCronTemplate) DeepCopy() *DataImportCronTemplate {
	if in == nil {
		return nil
	}
	out := new(DataImportCronTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureGates) DeepCopyInto(out *FeatureGates) {
	*out = *in
	if in.DeployCommonInstancetypes != nil {
		in, out := &in.DeployCommonInstancetypes, &out.DeployCommonInstancetypes
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureGates.
func (in *FeatureGates) DeepCopy() *FeatureGates {
	if in == nil {
		return nil
	}
	out := new(FeatureGates)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SSP) DeepCopyInto(out *SSP) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SSP.
func (in *SSP) DeepCopy() *SSP {
	if in == nil {
		return nil
	}
	out := new(SSP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SSP) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SSPList) DeepCopyInto(out *SSPList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SSP, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SSPList.
func (in *SSPList) DeepCopy() *SSPList {
	if in == nil {
		return nil
	}
	out := new(SSPList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SSPList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SSPSpec) DeepCopyInto(out *SSPSpec) {
	*out = *in
	if in.TemplateValidator != nil {
		in, out := &in.TemplateValidator, &out.TemplateValidator
		*out = new(TemplateValidator)
		(*in).DeepCopyInto(*out)
	}
	in.CommonTemplates.DeepCopyInto(&out.CommonTemplates)
	if in.TLSSecurityProfile != nil {
		in, out := &in.TLSSecurityProfile, &out.TLSSecurityProfile
		*out = new(v1.TLSSecurityProfile)
		(*in).DeepCopyInto(*out)
	}
	if in.TokenGenerationService != nil {
		in, out := &in.TokenGenerationService, &out.TokenGenerationService
		*out = new(TokenGenerationService)
		**out = **in
	}
	if in.CommonInstancetypes != nil {
		in, out := &in.CommonInstancetypes, &out.CommonInstancetypes
		*out = new(CommonInstancetypes)
		(*in).DeepCopyInto(*out)
	}
	if in.TektonPipelines != nil {
		in, out := &in.TektonPipelines, &out.TektonPipelines
		*out = new(TektonPipelines)
		**out = **in
	}
	if in.TektonTasks != nil {
		in, out := &in.TektonTasks, &out.TektonTasks
		*out = new(TektonTasks)
		**out = **in
	}
	if in.FeatureGates != nil {
		in, out := &in.FeatureGates, &out.FeatureGates
		*out = new(FeatureGates)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SSPSpec.
func (in *SSPSpec) DeepCopy() *SSPSpec {
	if in == nil {
		return nil
	}
	out := new(SSPSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SSPStatus) DeepCopyInto(out *SSPStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SSPStatus.
func (in *SSPStatus) DeepCopy() *SSPStatus {
	if in == nil {
		return nil
	}
	out := new(SSPStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TektonPipelines) DeepCopyInto(out *TektonPipelines) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TektonPipelines.
func (in *TektonPipelines) DeepCopy() *TektonPipelines {
	if in == nil {
		return nil
	}
	out := new(TektonPipelines)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TektonTasks) DeepCopyInto(out *TektonTasks) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TektonTasks.
func (in *TektonTasks) DeepCopy() *TektonTasks {
	if in == nil {
		return nil
	}
	out := new(TektonTasks)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateValidator) DeepCopyInto(out *TemplateValidator) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.Placement != nil {
		in, out := &in.Placement, &out.Placement
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateValidator.
func (in *TemplateValidator) DeepCopy() *TemplateValidator {
	if in == nil {
		return nil
	}
	out := new(TemplateValidator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TokenGenerationService) DeepCopyInto(out *TokenGenerationService) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TokenGenerationService.
func (in *TokenGenerationService) DeepCopy() *TokenGenerationService {
	if in == nil {
		return nil
	}
	out := new(TokenGenerationService)
	in.DeepCopyInto(out)
	return out
}
