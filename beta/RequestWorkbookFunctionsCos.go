// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

//
type WorkbookFunctionsCosRequestBuilder struct{ BaseRequestBuilder }

// Cos action undocumented
func (b *WorkbookFunctionsRequestBuilder) Cos(reqObj *WorkbookFunctionsCosRequestParameter) *WorkbookFunctionsCosRequestBuilder {
	bb := &WorkbookFunctionsCosRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/cos"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsCosRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsCosRequestBuilder) Request() *WorkbookFunctionsCosRequest {
	return &WorkbookFunctionsCosRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsCosRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}