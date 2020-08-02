// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

//
type WorkbookFunctionsUsdollarRequestBuilder struct{ BaseRequestBuilder }

// Usdollar action undocumented
func (b *WorkbookFunctionsRequestBuilder) Usdollar(reqObj *WorkbookFunctionsUsdollarRequestParameter) *WorkbookFunctionsUsdollarRequestBuilder {
	bb := &WorkbookFunctionsUsdollarRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/usdollar"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsUsdollarRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsUsdollarRequestBuilder) Request() *WorkbookFunctionsUsdollarRequest {
	return &WorkbookFunctionsUsdollarRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsUsdollarRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}