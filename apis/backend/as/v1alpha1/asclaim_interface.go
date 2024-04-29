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

package v1alpha1

import (
	"crypto/sha1"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/henderiw/apiserver-builder/pkg/builder/resource"
	"github.com/henderiw/store"
	"github.com/kuidio/kuid/apis/backend"
	commonv1alpha1 "github.com/kuidio/kuid/apis/common/v1alpha1"
	conditionv1alpha1 "github.com/kuidio/kuid/apis/condition/v1alpha1"
	rresource "github.com/kuidio/kuid/pkg/reconcilers/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/selection"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

const ASClaimPlural = "asclaims"
const ASClaimSingular = "asclaim"

// +k8s:deepcopy-gen=false
var _ resource.Object = &ASClaim{}
var _ resource.ObjectList = &ASClaimList{}

var _ resource.ObjectWithStatusSubResource = &ASClaim{}

func (ASClaimStatus) SubResourceName() string {
	return fmt.Sprintf("%s/%s", ASClaimPlural, "status")
}

func (r ASClaimStatus) CopyTo(obj resource.ObjectWithStatusSubResource) {
	cfg, ok := obj.(*ASClaim)
	if ok {
		cfg.Status = r
	}
}

func (r *ASClaim) GetStatus() resource.StatusSubResource {
	return r.Status
}

// GetListMeta returns the ListMeta
func (r *ASClaimList) GetListMeta() *metav1.ListMeta {
	return &r.ListMeta
}

func (r *ASClaim) GetSingularName() string {
	return ASClaimSingular
}

func (ASClaim) GetGroupVersionResource() schema.GroupVersionResource {
	return schema.GroupVersionResource{
		Group:    SchemeGroupVersion.Group,
		Version:  SchemeGroupVersion.Version,
		Resource: ASClaimPlural,
	}
}

// IsStorageVersion returns true -- v1alpha1.Config is used as the internal version.
// IsStorageVersion implements resource.Object.
func (ASClaim) IsStorageVersion() bool {
	return true
}

// GetObjectMeta implements resource.Object
func (r *ASClaim) GetObjectMeta() *metav1.ObjectMeta {
	return &r.ObjectMeta
}

// NamespaceScoped returns true to indicate Fortune is a namespaced resource.
// NamespaceScoped implements resource.Object.
func (ASClaim) NamespaceScoped() bool {
	return true
}

// New implements resource.Object
func (ASClaim) New() runtime.Object {
	return &ASClaim{}
}

// NewList implements resource.Object
func (ASClaim) NewList() runtime.Object {
	return &ASClaimList{}
}

// GetCondition returns the condition based on the condition kind
func (r *ASClaim) GetCondition(t conditionv1alpha1.ConditionType) conditionv1alpha1.Condition {
	return r.Status.GetCondition(t)
}

// SetConditions sets the conditions on the resource. it allows for 0, 1 or more conditions
// to be set at once
func (r *ASClaim) SetConditions(c ...conditionv1alpha1.Condition) {
	r.Status.SetConditions(c...)
}

// ConvertASClaimFieldSelector is the schema conversion function for normalizing the FieldSelector for ASClaim
func ConvertASClaimFieldSelector(label, value string) (internalLabel, internalValue string, err error) {
	switch label {
	case "metadata.name":
		return label, value, nil
	case "metadata.namespace":
		return label, value, nil
	case "spec.index":
		return label, value, nil
	default:
		return "", "", fmt.Errorf("%q is not a known field selector", label)
	}
}

func (r *ASClaimList) GetItems() []rresource.Object {
	objs := []rresource.Object{}
	for _, r := range r.Items {
		r := r
		objs = append(objs, &r)
	}
	return objs
}

func (r *ASClaim) CalculateHash() ([sha1.Size]byte, error) {
	// Convert the struct to JSON
	jsonData, err := json.Marshal(r)
	if err != nil {
		return [sha1.Size]byte{}, err
	}

	// Calculate SHA-1 hash
	return sha1.Sum(jsonData), nil
}

func (r *ASClaim) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Namespace: r.GetNamespace(),
		Name:      r.GetName(),
	}
}

func (r *ASClaim) GetKey() store.Key {
	return store.KeyFromNSN(types.NamespacedName{Namespace: r.Namespace, Name: r.Spec.Index})
}

func (r *ASClaim) GetOwnerReference() *commonv1alpha1.OwnerReference {
	return &commonv1alpha1.OwnerReference{
		Group:     SchemeGroupVersion.Group,
		Version:   SchemeGroupVersion.Version,
		Kind:      r.Kind,
		Namespace: r.Namespace,
		Name:      r.Name,
	}
}

func (r *ASClaim) GetClaimRequest() string {
	// we assume validation is already done when calling this
	if r.Spec.ID != nil {
		return strconv.Itoa(int(*r.Spec.ID))
	}
	if r.Spec.Range != nil {
		return *r.Spec.Range
	}
	return ""
}

func (r *ASClaim) GetClaimResponse() string {
	// we assume validation is already done when calling this
	if r.Status.ID != nil {
		return strconv.Itoa(int(*r.Status.ID))
	}
	if r.Status.Range != nil {
		return *r.Status.Range
	}
	return ""
}

func (r *ASClaim) GetClaimType() ASClaimType {
	claimType := ASClaimType_Invalid
	count := 0
	if r.Spec.ID != nil {
		claimType = ASClaimType_StaticID
		count++

	}
	if r.Spec.Range != nil {
		claimType = ASClaimType_Range
		count++

	}
	if count > 1 {
		return ASClaimType_Invalid
	}
	if count == 0 {
		return ASClaimType_DynamicID
	}
	return claimType
}

func (r *ASClaim) ValidateASClaimType() error {
	var sb strings.Builder
	count := 0
	if r.Spec.ID != nil {
		sb.WriteString(fmt.Sprintf("id: %d", *r.Spec.ID))
		count++

	}
	if r.Spec.Range != nil {
		if count > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(fmt.Sprintf("range: %s", *r.Spec.Range))
		count++

	}
	if count > 1 {
		return fmt.Errorf("an ipclaim can only have 1 addressing, got %s", sb.String())
	}
	return nil
}

func validateASID(id int) error {
	if id <= 0 {
		return fmt.Errorf("invalid AS id, got %d", id)
	}
	if id >= 4095 {
		return fmt.Errorf("invalid AS id, got %d", id)
	}
	return nil
}

func (r *ASClaim) ValidateASID() error {
	if r.Spec.ID == nil {
		return fmt.Errorf("no AS id provided")
	}
	if err := validateASID(int(*r.Spec.ID)); err != nil {
		return fmt.Errorf("invalid AS id err %s", err.Error())
	}
	return nil
}

func (r *ASClaim) GetASRange() (int, int) {
	if r.Spec.Range == nil {
		return 0, 0
	}
	parts := strings.SplitN(*r.Spec.Range, "-", 2)
	if len(parts) != 2 {
		return 0, 0
	}
	start, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0
	}
	end, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0
	}
	return start, end
}

func (r *ASClaim) ValidateASRange() error {
	if r.Spec.Range == nil {
		return fmt.Errorf("no AS range provided")
	}
	parts := strings.SplitN(*r.Spec.Range, "-", 2)
	if len(parts) != 2 {
		return fmt.Errorf("invalid AS range, expected <start>-<end>, got: %s", *r.Spec.Range)
	}
	var errm error
	start, err := strconv.Atoi(parts[0])
	if err != nil {
		errm = errors.Join(errm, fmt.Errorf("invalid AS range start, got: %s, err: %s", *r.Spec.Range, err.Error()))
	}
	end, err := strconv.Atoi(parts[1])
	if err != nil {
		errm = errors.Join(errm, fmt.Errorf("invalid AS range end, got: %s, err: %s", *r.Spec.Range, err.Error()))
	}
	if errm != nil {
		return errm
	}
	if start >= end {
		errm = errors.Join(errm, fmt.Errorf("invalid AS range start >= end %s", *r.Spec.Range))
	}
	if err := validateASID(start); err != nil {
		errm = errors.Join(errm, fmt.Errorf("invalid AS start err %s", err.Error()))
	}
	if err := validateASID(end); err != nil {
		errm = errors.Join(errm, fmt.Errorf("invalid AS end err %s", err.Error()))
	}
	return errm
}

func (r *ASClaim) ValidateSyntax() field.ErrorList {
	var allErrs field.ErrorList

	gv, err := schema.ParseGroupVersion(r.APIVersion)
	if err != nil {
		allErrs = append(allErrs, field.Invalid(
			field.NewPath("apiVersion"),
			r,
			fmt.Errorf("invalid apiVersion: err: %s", err.Error()).Error(),
		))
		return allErrs
	}

	// this is for user convenience
	if r.Spec.Owner == nil {
		r.Spec.Owner = &commonv1alpha1.OwnerReference{
			Group:     gv.Group,
			Version:   gv.Version,
			Kind:      r.Kind,
			Namespace: r.Namespace,
			Name:      r.Name,
		}
	}

	if err := r.ValidateASClaimType(); err != nil {
		allErrs = append(allErrs, field.Invalid(
			field.NewPath(""),
			r,
			err.Error(),
		))
		return allErrs
	}
	var v SyntaxValidator
	claimType := r.GetClaimType()
	switch claimType {
	case ASClaimType_DynamicID:
		v = &ASDynamicIDSyntaxValidator{name: string(claimType)}
	case ASClaimType_StaticID:
		v = &ASStaticIDSyntaxValidator{name: string(claimType)}
	case ASClaimType_Range:
		v = &ASRangeSyntaxValidator{name: string(claimType)}
	default:
		return allErrs
	}
	return v.Validate(r)
}

func (r *ASClaim) ValidateOwner(labels labels.Set) error {
	routeClaimName := labels[backend.KuidClaimNameKey]
	routeOwner := commonv1alpha1.OwnerReference{
		Group:     labels[backend.KuidOwnerGroupKey],
		Version:   labels[backend.KuidOwnerVersionKey],
		Kind:      labels[backend.KuidOwnerKindKey],
		Namespace: labels[backend.KuidOwnerNamespaceKey],
		Name:      labels[backend.KuidOwnerNameKey],
	}
	if (r.Spec.Owner != nil && *r.Spec.Owner != routeOwner) || r.Name != routeClaimName {
		return fmt.Errorf("route owned by different claim got name %s/%s owner %s/%s",
			r.Name,
			routeClaimName,
			r.Spec.Owner.String(),
			routeOwner.String(),
		)
	}
	return nil
}

// GetLabelSelector returns a labels selector based on the label selector
func (r *ASClaim) GetLabelSelector() (labels.Selector, error) {
	return r.Spec.GetLabelSelector()
}

func (r *ASClaim) GetClaimLabels() labels.Set {
	labels := r.Spec.GetUserDefinedLabels()
	// system defined labels
	labels[backend.KuidASClaimTypeKey] = string(r.GetClaimType())
	labels[backend.KuidClaimNameKey] = r.Name
	labels[backend.KuidOwnerGroupKey] = r.Spec.Owner.Group
	labels[backend.KuidOwnerVersionKey] = r.Spec.Owner.Version
	labels[backend.KuidOwnerKindKey] = r.Spec.Owner.Kind
	labels[backend.KuidOwnerNamespaceKey] = r.Spec.Owner.Namespace
	labels[backend.KuidOwnerNameKey] = r.Spec.Owner.Name
	return labels
}

// GetOwnerSelector returns a label selector to select the owner of the claim in the backend
func (r *ASClaim) GetOwnerSelector() (labels.Selector, error) {
	l := map[string]string{
		backend.KuidOwnerGroupKey:     r.Spec.Owner.Group,
		backend.KuidOwnerVersionKey:   r.Spec.Owner.Version,
		backend.KuidOwnerKindKey:      r.Spec.Owner.Kind,
		backend.KuidOwnerNamespaceKey: r.Spec.Owner.Namespace,
		backend.KuidOwnerNameKey:      r.Spec.Owner.Name,
		backend.KuidClaimNameKey:      r.Name,
	}

	fullselector := labels.NewSelector()
	for k, v := range l {
		req, err := labels.NewRequirement(k, selection.Equals, []string{v})
		if err != nil {
			return nil, err
		}
		fullselector = fullselector.Add(*req)
	}
	return fullselector, nil
}

// BuildASClaim returns a reource from a client Object a Spec/Status
func BuildASClaim(meta metav1.ObjectMeta, spec *ASClaimSpec, status *ASClaimStatus) *ASClaim {
	aspec := ASClaimSpec{}
	if spec != nil {
		aspec = *spec
	}
	astatus := ASClaimStatus{}
	if status != nil {
		astatus = *status
	}
	return &ASClaim{
		TypeMeta: metav1.TypeMeta{
			APIVersion: SchemeGroupVersion.Identifier(),
			Kind:       ASClaimKind,
		},
		ObjectMeta: meta,
		Spec:       aspec,
		Status:     astatus,
	}
}