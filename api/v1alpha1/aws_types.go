/*
Copyright 2023.

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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	PendingState = "PENDING"
	ReadyState   = "READY"
	TaintedState = "TAINTED"
)

// AwsSpec defines the desired state of Aws
type AwsSpec struct {
	// Important: Run "make" to regenerate code after modifying this file

	// Name is the name of the account
	Name string `json:"name"`

	// ID is the account identifier
	ID string `json:"id"`

	// ManagementAccountID is the organization's management account id
	ManagementAccountID string `json:"managementAccountId"`

	// Tags is a list of account tags
	Tags map[string]string `json:"tags"`
}

// AwsStatus defines the observed state of Aws
type AwsStatus struct {
	// Important: Run "make" to regenerate code after modifying this file
	State string `json:"state"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:printcolumn:name="ID",type="string",JSONPath=`.spec.id`,description="The ID of the account"
//+kubebuilder:printcolumn:name="Name",type="string",JSONPath=`.spec.name`,description="The name of the account"
//+kubebuilder:printcolumn:name="Management Account ID",type="string",JSONPath=`.spec.managementAccountId`,description="The management account ID of the account"
//+kubebuilder:printcolumn:name="State",type="string",JSONPath=`.status.state`,description="The state of the account"

// Aws is the Schema for the aws API
type Aws struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AwsSpec   `json:"spec,omitempty"`
	Status AwsStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// AwsList contains a list of Aws
type AwsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Aws `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Aws{}, &AwsList{})
}
