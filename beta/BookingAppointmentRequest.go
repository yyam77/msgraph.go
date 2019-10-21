// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// BookingAppointmentRequestBuilder is request builder for BookingAppointment
type BookingAppointmentRequestBuilder struct{ BaseRequestBuilder }

// Request returns BookingAppointmentRequest
func (b *BookingAppointmentRequestBuilder) Request() *BookingAppointmentRequest {
	return &BookingAppointmentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// BookingAppointmentRequest is request for BookingAppointment
type BookingAppointmentRequest struct{ BaseRequest }

// Do performs HTTP request for BookingAppointment
func (r *BookingAppointmentRequest) Do(method, path string, reqObj interface{}) (resObj *BookingAppointment, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for BookingAppointment
func (r *BookingAppointmentRequest) Get() (*BookingAppointment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for BookingAppointment
func (r *BookingAppointmentRequest) Update(reqObj *BookingAppointment) (*BookingAppointment, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for BookingAppointment
func (r *BookingAppointmentRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}