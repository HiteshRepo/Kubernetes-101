package v1

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Define the struct for the custom CRD
type FlightTicket struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FlightTicketSpec `json:"spec"`
}

type FlightTicketSpec struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Number int    `json:"number"`
}

// Define the struct for the list of custom CRDs
type FlightTicketList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FlightTicket `json:"items"`
}
