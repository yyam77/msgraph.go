// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ManagedAppPolicyRequestBuilder is request builder for ManagedAppPolicy
type ManagedAppPolicyRequestBuilder struct{ BaseRequestBuilder }

// Request returns ManagedAppPolicyRequest
func (b *ManagedAppPolicyRequestBuilder) Request() *ManagedAppPolicyRequest {
	return &ManagedAppPolicyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ManagedAppPolicyRequest is request for ManagedAppPolicy
type ManagedAppPolicyRequest struct{ BaseRequest }

// Do performs HTTP request for ManagedAppPolicy
func (r *ManagedAppPolicyRequest) Do(method, path string, reqObj interface{}) (resObj *ManagedAppPolicy, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for ManagedAppPolicy
func (r *ManagedAppPolicyRequest) Get() (*ManagedAppPolicy, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for ManagedAppPolicy
func (r *ManagedAppPolicyRequest) Update(reqObj *ManagedAppPolicy) (*ManagedAppPolicy, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for ManagedAppPolicy
func (r *ManagedAppPolicyRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}