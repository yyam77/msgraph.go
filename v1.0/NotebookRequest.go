// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// NotebookRequestBuilder is request builder for Notebook
type NotebookRequestBuilder struct{ BaseRequestBuilder }

// Request returns NotebookRequest
func (b *NotebookRequestBuilder) Request() *NotebookRequest {
	return &NotebookRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// NotebookRequest is request for Notebook
type NotebookRequest struct{ BaseRequest }

// Do performs HTTP request for Notebook
func (r *NotebookRequest) Do(method, path string, reqObj interface{}) (resObj *Notebook, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for Notebook
func (r *NotebookRequest) Get() (*Notebook, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for Notebook
func (r *NotebookRequest) Update(reqObj *Notebook) (*Notebook, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for Notebook
func (r *NotebookRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// SectionGroups returns request builder for SectionGroup collection
func (b *NotebookRequestBuilder) SectionGroups() *NotebookSectionGroupsCollectionRequestBuilder {
	bb := &NotebookSectionGroupsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/sectionGroups"
	return bb
}

// NotebookSectionGroupsCollectionRequestBuilder is request builder for SectionGroup collection
type NotebookSectionGroupsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for SectionGroup collection
func (b *NotebookSectionGroupsCollectionRequestBuilder) Request() *NotebookSectionGroupsCollectionRequest {
	return &NotebookSectionGroupsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for SectionGroup item
func (b *NotebookSectionGroupsCollectionRequestBuilder) ID(id string) *SectionGroupRequestBuilder {
	bb := &SectionGroupRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// NotebookSectionGroupsCollectionRequest is request for SectionGroup collection
type NotebookSectionGroupsCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for SectionGroup collection
func (r *NotebookSectionGroupsCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *SectionGroup, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for SectionGroup collection
func (r *NotebookSectionGroupsCollectionRequest) Paging(method, path string, obj interface{}) ([]SectionGroup, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []SectionGroup
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []SectionGroup
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if len(paging.NextLink) == 0 {
			return values, nil
		}
		res, err = r.client.Get(paging.NextLink)
		if err != nil {
			return nil, err
		}
	}
}

// Get performs GET request for SectionGroup collection
func (r *NotebookSectionGroupsCollectionRequest) Get() ([]SectionGroup, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for SectionGroup collection
func (r *NotebookSectionGroupsCollectionRequest) Add(reqObj *SectionGroup) (*SectionGroup, error) {
	return r.Do("POST", "", reqObj)
}

// Sections returns request builder for OnenoteSection collection
func (b *NotebookRequestBuilder) Sections() *NotebookSectionsCollectionRequestBuilder {
	bb := &NotebookSectionsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/sections"
	return bb
}

// NotebookSectionsCollectionRequestBuilder is request builder for OnenoteSection collection
type NotebookSectionsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for OnenoteSection collection
func (b *NotebookSectionsCollectionRequestBuilder) Request() *NotebookSectionsCollectionRequest {
	return &NotebookSectionsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for OnenoteSection item
func (b *NotebookSectionsCollectionRequestBuilder) ID(id string) *OnenoteSectionRequestBuilder {
	bb := &OnenoteSectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// NotebookSectionsCollectionRequest is request for OnenoteSection collection
type NotebookSectionsCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for OnenoteSection collection
func (r *NotebookSectionsCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *OnenoteSection, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for OnenoteSection collection
func (r *NotebookSectionsCollectionRequest) Paging(method, path string, obj interface{}) ([]OnenoteSection, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []OnenoteSection
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []OnenoteSection
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if len(paging.NextLink) == 0 {
			return values, nil
		}
		res, err = r.client.Get(paging.NextLink)
		if err != nil {
			return nil, err
		}
	}
}

// Get performs GET request for OnenoteSection collection
func (r *NotebookSectionsCollectionRequest) Get() ([]OnenoteSection, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for OnenoteSection collection
func (r *NotebookSectionsCollectionRequest) Add(reqObj *OnenoteSection) (*OnenoteSection, error) {
	return r.Do("POST", "", reqObj)
}