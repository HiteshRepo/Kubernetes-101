package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	GroupName  = "flights.com"
	GroupVer   = "v1"
	KindName   = "FlightTicket"
	PluralName = "flighttickets"
)

var (
	SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: GroupVer}
	SchemeBuilder      = runtime.NewSchemeBuilder(addKnownTypes)
)

var resource = schema.GroupVersionResource{
	Group:    GroupName,
	Version:  GroupVer,
	Resource: PluralName,
}

var AddToScheme = runtime.NewSchemeBuilder(addKnownTypes).AddToScheme

func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(
		SchemeGroupVersion,
		&FlightTicket{},
		&FlightTicketList{},
	)

	// register the type in the scheme
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
