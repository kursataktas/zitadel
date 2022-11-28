package logstore

import (
	"context"
	"net/http"
	"net/textproto"
	"strings"
	"time"

	zitadel_http "github.com/zitadel/zitadel/internal/api/http"
)

type Protocol uint8

const (
	GRPC Protocol = iota
	HTTP
	// TODO: GRPC-Web?
	// TODO: HTTPS?
)

type AccessLogRecord struct {
	Timestamp       time.Time
	Protocol        Protocol
	RequestURL      string
	ResponseStatus  uint32
	RequestHeaders  http.Header
	ResponseHeaders http.Header
	InstanceID      string
	ProjectID       string
	RequestedDomain string
	RequestedHost   string
}

func (a *AccessLogRecord) IsAuthenticated() bool {
	// TODO: tripplecheck
	return a.Protocol == GRPC &&
		len(a.ResponseHeaders[strings.ToLower(zitadel_http.Authorization)]) > 0 ||
		a.Protocol == HTTP &&
			a.RequestHeaders.Get(textproto.CanonicalMIMEHeaderKey(zitadel_http.Authorization)) != "" &&
			a.ResponseStatus != 401 &&
			a.ResponseStatus != 500
}

type StoredAccessLogsReducer interface {
	Reduce(context.Context, []*AccessLogRecord)
}

type StoredAccessLogsReducerFunc func(context.Context, []*AccessLogRecord)

func (s StoredAccessLogsReducerFunc) Reduce(ctx context.Context, records []*AccessLogRecord) {
	s(ctx, records)
}
