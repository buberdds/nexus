// Tooling for response pagination.
package common

import (
	"net/http"
	"strconv"
)

const (
	LimitKey  = "limit"
	OffsetKey = "offset"

	// By default, just order by the first returned column so
	// we always have a deterministic ordering.
	DefaultOrder  = "1"
	DefaultLimit  = uint64(100)
	DefaultOffset = uint64(0)

	MaximumLimit = uint64(1000)

	DefaultBucketSizeSeconds = uint64(3600)
)

// Pagination is used to define parameters for pagination.
type Pagination struct {
	Limit  uint64
	Offset uint64
	Order  *string
}

// BucketedStatsParams are used to parametrize stats queries that return
// time-bucketed results.
type BucketedStatsParams struct {
	BucketSizeSeconds uint32
}

// NewPagination extracts pagination parameters from an http request.
func NewPagination(r *http.Request) (p Pagination, err error) {
	values := r.URL.Query()

	limit := DefaultLimit
	if v := values.Get(LimitKey); v != "" {
		limit, err = strconv.ParseUint(v, 10, 64)
	}
	if limit > MaximumLimit {
		limit = MaximumLimit
	}

	offset := DefaultOffset
	if v := values.Get(OffsetKey); v != "" {
		offset, err = strconv.ParseUint(v, 10, 64)
	}

	order := DefaultOrder

	p = Pagination{
		Limit:  limit,
		Offset: offset,
		Order:  &order,
	}
	return
}

// BucketedStatsParams extracts bucket size parameters from an http request.
func NewBucketedStatsParams(r *http.Request) (b BucketedStatsParams, err error) {
	values := r.URL.Query()

	bucketSizeSeconds := DefaultBucketSizeSeconds
	if v := values.Get("bucket_size_seconds"); v != "" {
		bucketSizeSeconds, err = strconv.ParseUint(v, 10, 32)
	}

	b = BucketedStatsParams{
		BucketSizeSeconds: uint32(bucketSizeSeconds), // safe cast because of how the "32" param to ParseUint
	}
	return
}
