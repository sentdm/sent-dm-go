// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagination

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/sentdm/sent-dm-go/internal/apijson"
	"github.com/sentdm/sent-dm-go/internal/requestconfig"
	"github.com/sentdm/sent-dm-go/packages/param"
	"github.com/sentdm/sent-dm-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type ContactsPageData[T any] struct {
	Contacts   []T                        `json:"contacts"`
	Pagination ContactsPageDataPagination `json:"pagination"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Contacts    respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContactsPageData[T]) RawJSON() string { return r.JSON.raw }
func (r *ContactsPageData[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContactsPageDataPagination struct {
	HasMore bool `json:"has_more"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasMore     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContactsPageDataPagination) RawJSON() string { return r.JSON.raw }
func (r *ContactsPageDataPagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContactsPage[T any] struct {
	Data ContactsPageData[T] `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r ContactsPage[T]) RawJSON() string { return r.JSON.raw }
func (r *ContactsPage[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *ContactsPage[T]) GetNextPage() (res *ContactsPage[T], err error) {
	if len(r.Data.Contacts) == 0 {
		return nil, nil
	}

	if r.Data.Pagination.JSON.HasMore.Valid() && r.Data.Pagination.HasMore == false {
		return nil, nil
	}
	u := r.cfg.Request.URL
	currentPage, err := strconv.ParseInt(u.Query().Get("page"), 10, 64)
	if err != nil {
		currentPage = 1
	}
	cfg := r.cfg.Clone(context.Background())
	query := cfg.Request.URL.Query()
	query.Set("page", fmt.Sprintf("%d", currentPage+1))
	cfg.Request.URL.RawQuery = query.Encode()
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *ContactsPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &ContactsPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type ContactsPageAutoPager[T any] struct {
	page *ContactsPage[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewContactsPageAutoPager[T any](page *ContactsPage[T], err error) *ContactsPageAutoPager[T] {
	return &ContactsPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *ContactsPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Data.Contacts) == 0 {
		return false
	}
	if r.idx >= len(r.page.Data.Contacts) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Data.Contacts) == 0 {
			return false
		}
	}
	r.cur = r.page.Data.Contacts[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *ContactsPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *ContactsPageAutoPager[T]) Err() error {
	return r.err
}

func (r *ContactsPageAutoPager[T]) Index() int {
	return r.run
}

type ConversationsPageData[T any] struct {
	Messages   []T                             `json:"messages"`
	Pagination ConversationsPageDataPagination `json:"pagination"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Messages    respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationsPageData[T]) RawJSON() string { return r.JSON.raw }
func (r *ConversationsPageData[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationsPageDataPagination struct {
	HasMore bool `json:"has_more"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasMore     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationsPageDataPagination) RawJSON() string { return r.JSON.raw }
func (r *ConversationsPageDataPagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationsPage[T any] struct {
	Data ConversationsPageData[T] `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r ConversationsPage[T]) RawJSON() string { return r.JSON.raw }
func (r *ConversationsPage[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *ConversationsPage[T]) GetNextPage() (res *ConversationsPage[T], err error) {
	if len(r.Data.Messages) == 0 {
		return nil, nil
	}

	if r.Data.Pagination.JSON.HasMore.Valid() && r.Data.Pagination.HasMore == false {
		return nil, nil
	}
	u := r.cfg.Request.URL
	currentPage, err := strconv.ParseInt(u.Query().Get("page"), 10, 64)
	if err != nil {
		currentPage = 1
	}
	cfg := r.cfg.Clone(context.Background())
	query := cfg.Request.URL.Query()
	query.Set("page", fmt.Sprintf("%d", currentPage+1))
	cfg.Request.URL.RawQuery = query.Encode()
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *ConversationsPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &ConversationsPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type ConversationsPageAutoPager[T any] struct {
	page *ConversationsPage[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewConversationsPageAutoPager[T any](page *ConversationsPage[T], err error) *ConversationsPageAutoPager[T] {
	return &ConversationsPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *ConversationsPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Data.Messages) == 0 {
		return false
	}
	if r.idx >= len(r.page.Data.Messages) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Data.Messages) == 0 {
			return false
		}
	}
	r.cur = r.page.Data.Messages[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *ConversationsPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *ConversationsPageAutoPager[T]) Err() error {
	return r.err
}

func (r *ConversationsPageAutoPager[T]) Index() int {
	return r.run
}

type TemplatesPageData[T any] struct {
	Pagination TemplatesPageDataPagination `json:"pagination"`
	Templates  []T                         `json:"templates"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Pagination  respjson.Field
		Templates   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TemplatesPageData[T]) RawJSON() string { return r.JSON.raw }
func (r *TemplatesPageData[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TemplatesPageDataPagination struct {
	HasMore bool `json:"has_more"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasMore     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TemplatesPageDataPagination) RawJSON() string { return r.JSON.raw }
func (r *TemplatesPageDataPagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TemplatesPage[T any] struct {
	Data TemplatesPageData[T] `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r TemplatesPage[T]) RawJSON() string { return r.JSON.raw }
func (r *TemplatesPage[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *TemplatesPage[T]) GetNextPage() (res *TemplatesPage[T], err error) {
	if len(r.Data.Templates) == 0 {
		return nil, nil
	}

	if r.Data.Pagination.JSON.HasMore.Valid() && r.Data.Pagination.HasMore == false {
		return nil, nil
	}
	u := r.cfg.Request.URL
	currentPage, err := strconv.ParseInt(u.Query().Get("page"), 10, 64)
	if err != nil {
		currentPage = 1
	}
	cfg := r.cfg.Clone(context.Background())
	query := cfg.Request.URL.Query()
	query.Set("page", fmt.Sprintf("%d", currentPage+1))
	cfg.Request.URL.RawQuery = query.Encode()
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *TemplatesPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &TemplatesPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type TemplatesPageAutoPager[T any] struct {
	page *TemplatesPage[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewTemplatesPageAutoPager[T any](page *TemplatesPage[T], err error) *TemplatesPageAutoPager[T] {
	return &TemplatesPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *TemplatesPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Data.Templates) == 0 {
		return false
	}
	if r.idx >= len(r.page.Data.Templates) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Data.Templates) == 0 {
			return false
		}
	}
	r.cur = r.page.Data.Templates[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *TemplatesPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *TemplatesPageAutoPager[T]) Err() error {
	return r.err
}

func (r *TemplatesPageAutoPager[T]) Index() int {
	return r.run
}

type WebhooksPageData[T any] struct {
	Pagination WebhooksPageDataPagination `json:"pagination"`
	Webhooks   []T                        `json:"webhooks"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Pagination  respjson.Field
		Webhooks    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhooksPageData[T]) RawJSON() string { return r.JSON.raw }
func (r *WebhooksPageData[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhooksPageDataPagination struct {
	HasMore bool `json:"has_more"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasMore     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhooksPageDataPagination) RawJSON() string { return r.JSON.raw }
func (r *WebhooksPageDataPagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhooksPage[T any] struct {
	Data WebhooksPageData[T] `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r WebhooksPage[T]) RawJSON() string { return r.JSON.raw }
func (r *WebhooksPage[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *WebhooksPage[T]) GetNextPage() (res *WebhooksPage[T], err error) {
	if len(r.Data.Webhooks) == 0 {
		return nil, nil
	}

	if r.Data.Pagination.JSON.HasMore.Valid() && r.Data.Pagination.HasMore == false {
		return nil, nil
	}
	u := r.cfg.Request.URL
	currentPage, err := strconv.ParseInt(u.Query().Get("page"), 10, 64)
	if err != nil {
		currentPage = 1
	}
	cfg := r.cfg.Clone(context.Background())
	query := cfg.Request.URL.Query()
	query.Set("page", fmt.Sprintf("%d", currentPage+1))
	cfg.Request.URL.RawQuery = query.Encode()
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *WebhooksPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &WebhooksPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type WebhooksPageAutoPager[T any] struct {
	page *WebhooksPage[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewWebhooksPageAutoPager[T any](page *WebhooksPage[T], err error) *WebhooksPageAutoPager[T] {
	return &WebhooksPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *WebhooksPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Data.Webhooks) == 0 {
		return false
	}
	if r.idx >= len(r.page.Data.Webhooks) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Data.Webhooks) == 0 {
			return false
		}
	}
	r.cur = r.page.Data.Webhooks[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *WebhooksPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *WebhooksPageAutoPager[T]) Err() error {
	return r.err
}

func (r *WebhooksPageAutoPager[T]) Index() int {
	return r.run
}

type WebhookEventsPageData[T any] struct {
	Events     []T                             `json:"events"`
	Pagination WebhookEventsPageDataPagination `json:"pagination"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Events      respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookEventsPageData[T]) RawJSON() string { return r.JSON.raw }
func (r *WebhookEventsPageData[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookEventsPageDataPagination struct {
	HasMore bool `json:"has_more"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasMore     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookEventsPageDataPagination) RawJSON() string { return r.JSON.raw }
func (r *WebhookEventsPageDataPagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookEventsPage[T any] struct {
	Data WebhookEventsPageData[T] `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r WebhookEventsPage[T]) RawJSON() string { return r.JSON.raw }
func (r *WebhookEventsPage[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *WebhookEventsPage[T]) GetNextPage() (res *WebhookEventsPage[T], err error) {
	if len(r.Data.Events) == 0 {
		return nil, nil
	}

	if r.Data.Pagination.JSON.HasMore.Valid() && r.Data.Pagination.HasMore == false {
		return nil, nil
	}
	u := r.cfg.Request.URL
	currentPage, err := strconv.ParseInt(u.Query().Get("page"), 10, 64)
	if err != nil {
		currentPage = 1
	}
	cfg := r.cfg.Clone(context.Background())
	query := cfg.Request.URL.Query()
	query.Set("page", fmt.Sprintf("%d", currentPage+1))
	cfg.Request.URL.RawQuery = query.Encode()
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *WebhookEventsPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &WebhookEventsPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type WebhookEventsPageAutoPager[T any] struct {
	page *WebhookEventsPage[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewWebhookEventsPageAutoPager[T any](page *WebhookEventsPage[T], err error) *WebhookEventsPageAutoPager[T] {
	return &WebhookEventsPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *WebhookEventsPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Data.Events) == 0 {
		return false
	}
	if r.idx >= len(r.page.Data.Events) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Data.Events) == 0 {
			return false
		}
	}
	r.cur = r.page.Data.Events[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *WebhookEventsPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *WebhookEventsPageAutoPager[T]) Err() error {
	return r.err
}

func (r *WebhookEventsPageAutoPager[T]) Index() int {
	return r.run
}
