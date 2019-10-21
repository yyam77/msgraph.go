// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// SectionGroupRequestBuilder is request builder for SectionGroup
type SectionGroupRequestBuilder struct{ BaseRequestBuilder }

// Request returns SectionGroupRequest
func (b *SectionGroupRequestBuilder) Request() *SectionGroupRequest {
	return &SectionGroupRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SectionGroupRequest is request for SectionGroup
type SectionGroupRequest struct{ BaseRequest }

// Do performs HTTP request for SectionGroup
func (r *SectionGroupRequest) Do(method, path string, reqObj interface{}) (resObj *SectionGroup, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for SectionGroup
func (r *SectionGroupRequest) Get() (*SectionGroup, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for SectionGroup
func (r *SectionGroupRequest) Update(reqObj *SectionGroup) (*SectionGroup, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for SectionGroup
func (r *SectionGroupRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// ParentNotebook is navigation property
func (b *SectionGroupRequestBuilder) ParentNotebook() *NotebookRequestBuilder {
	bb := &NotebookRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/parentNotebook"
	return bb
}

// ParentSectionGroup is navigation property
func (b *SectionGroupRequestBuilder) ParentSectionGroup() *SectionGroupRequestBuilder {
	bb := &SectionGroupRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/parentSectionGroup"
	return bb
}

// SectionGroups returns request builder for SectionGroup collection
func (b *SectionGroupRequestBuilder) SectionGroups() *SectionGroupSectionGroupsCollectionRequestBuilder {
	bb := &SectionGroupSectionGroupsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/sectionGroups"
	return bb
}

// SectionGroupSectionGroupsCollectionRequestBuilder is request builder for SectionGroup collection
type SectionGroupSectionGroupsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for SectionGroup collection
func (b *SectionGroupSectionGroupsCollectionRequestBuilder) Request() *SectionGroupSectionGroupsCollectionRequest {
	return &SectionGroupSectionGroupsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for SectionGroup item
func (b *SectionGroupSectionGroupsCollectionRequestBuilder) ID(id string) *SectionGroupRequestBuilder {
	bb := &SectionGroupRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// SectionGroupSectionGroupsCollectionRequest is request for SectionGroup collection
type SectionGroupSectionGroupsCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for SectionGroup collection
func (r *SectionGroupSectionGroupsCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *SectionGroup, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for SectionGroup collection
func (r *SectionGroupSectionGroupsCollectionRequest) Paging(method, path string, obj interface{}) ([]SectionGroup, error) {
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
func (r *SectionGroupSectionGroupsCollectionRequest) Get() ([]SectionGroup, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for SectionGroup collection
func (r *SectionGroupSectionGroupsCollectionRequest) Add(reqObj *SectionGroup) (*SectionGroup, error) {
	return r.Do("POST", "", reqObj)
}

// Sections returns request builder for OnenoteSection collection
func (b *SectionGroupRequestBuilder) Sections() *SectionGroupSectionsCollectionRequestBuilder {
	bb := &SectionGroupSectionsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/sections"
	return bb
}

// SectionGroupSectionsCollectionRequestBuilder is request builder for OnenoteSection collection
type SectionGroupSectionsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for OnenoteSection collection
func (b *SectionGroupSectionsCollectionRequestBuilder) Request() *SectionGroupSectionsCollectionRequest {
	return &SectionGroupSectionsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for OnenoteSection item
func (b *SectionGroupSectionsCollectionRequestBuilder) ID(id string) *OnenoteSectionRequestBuilder {
	bb := &OnenoteSectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// SectionGroupSectionsCollectionRequest is request for OnenoteSection collection
type SectionGroupSectionsCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for OnenoteSection collection
func (r *SectionGroupSectionsCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *OnenoteSection, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for OnenoteSection collection
func (r *SectionGroupSectionsCollectionRequest) Paging(method, path string, obj interface{}) ([]OnenoteSection, error) {
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
func (r *SectionGroupSectionsCollectionRequest) Get() ([]OnenoteSection, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for OnenoteSection collection
func (r *SectionGroupSectionsCollectionRequest) Add(reqObj *OnenoteSection) (*OnenoteSection, error) {
	return r.Do("POST", "", reqObj)
}