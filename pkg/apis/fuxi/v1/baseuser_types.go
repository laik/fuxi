package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// BaseUserSpec defines the desired state of BaseUser
type BaseUserSpec struct {
	Name          *string `json:"name"`
	Password      *string `json:"password"`
	RoleId        int     `json:"role_id"`
	DepartmentId  int     `json:"department_id"`
	IsDelete      bool    `json:"is_delete"`
	Email         string  `json:"email"`
	Display       string  `json:"display"`
	CreatorId     int     `json:"creator_id"`
	LastLoginTime string  `json:"last_login_time"`
}

// BaseUserStatus defines the observed state of BaseUser
type BaseUserStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BaseUser is the Schema for the baseusers API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=baseusers,scope=Namespaced
type BaseUser struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BaseUserSpec   `json:"spec,omitempty"`
	Status BaseUserStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BaseUserList contains a list of BaseUser
type BaseUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BaseUser `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BaseUser{}, &BaseUserList{})
}