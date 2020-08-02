// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

//
type WorkbookFunctionsNominalRequestBuilder struct{ BaseRequestBuilder }

// Nominal action undocumented
func (b *WorkbookFunctionsRequestBuilder) Nominal(reqObj *WorkbookFunctionsNominalRequestParameter) *WorkbookFunctionsNominalRequestBuilder {
	bb := &WorkbookFunctionsNominalRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/nominal"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsNominalRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsNominalRequestBuilder) Request() *WorkbookFunctionsNominalRequest {
	return &WorkbookFunctionsNominalRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsNominalRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}