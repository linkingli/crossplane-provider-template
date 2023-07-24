/*
Copyright 2022 The Crossplane Authors.

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
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// redParameters are the configurable fields of a red.
type redParameters struct {
	ConfigurableField string `json:"configurableField"`
}

// redObservation are the observable fields of a red.
type redObservation struct {
	ObservableField string `json:"observableField,omitempty"`
}

// A redSpec defines the desired state of a red.
type redSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       redParameters `json:"forProvider"`
}

// A redStatus represents the observed state of a red.
type redStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          redObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A red is an example API type.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,jelly}
type red struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   redSpec   `json:"spec"`
	Status redStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// redList contains a list of red
type redList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []red `json:"items"`
}

// red type metadata.
var (
	redKind             = reflect.TypeOf(red{}).Name()
	redGroupKind        = schema.GroupKind{Group: Group, Kind: redKind}.String()
	redKindAPIVersion   = redKind + "." + SchemeGroupVersion.String()
	redGroupVersionKind = SchemeGroupVersion.WithKind(redKind)
)

func init() {
	SchemeBuilder.Register(&red{}, &redList{})
}
