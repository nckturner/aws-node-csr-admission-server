package main

import (
	"github.com/openshift/generic-admission-server/pkg/cmd"
	admissionv1beta1 "k8s.io/api/admission/v1beta1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/rest"
)

type awsNodeCSRAdmissionHook struct {
}

func main() {
	cmd.RunAdmissionServer(&awsNodeCSRAdmissionHook{})
}

// where to host it
func (a *awsNodeCSRAdmissionHook) ValidatingResource() (plural schema.GroupVersionResource, singular string) {
	return schema.GroupVersionResource{}, ""
}

// your business logic
func (a *awsNodeCSRAdmissionHook) Validate(admissionSpec *admissionv1beta1.AdmissionRequest) *admissionv1beta1.AdmissionResponse {
	return nil
}

// any special initialization goes here
func (a *awsNodeCSRAdmissionHook) Initialize(kubeClientConfig *rest.Config, stopCh <-chan struct{}) error {
	return nil
}
