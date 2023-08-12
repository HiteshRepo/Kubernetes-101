package v1

import "k8s.io/apimachinery/pkg/runtime"

func (in *FlightTicket) DeepCopyInto(out *FlightTicket) {
	out.TypeMeta = in.TypeMeta
	out.ObjectMeta = in.ObjectMeta
	out.Spec = FlightTicketSpec{
		From:   in.Spec.From,
		To:     in.Spec.To,
		Number: in.Spec.Number,
	}
}

func (in *FlightTicket) DeepCopyObject() runtime.Object {
	out := FlightTicket{}
	in.DeepCopyInto(&out)

	return &out
}

func (in *FlightTicket) DeepCopyObject() runtime.Object {
	out := FlightTicketList{}
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta

	if in.Items != nil {
		out.Items = make([]FlightTicket, len(in.Items))
		for i := range in.Items {
			in.Items[i].DeepCopyInto(&out.Items[i])
		}
	}

	return &out
}
