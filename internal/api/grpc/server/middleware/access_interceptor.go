package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/zitadel/zitadel/internal/api/authz"

	"google.golang.org/grpc/status"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/zitadel/zitadel/internal/logstore"
	"github.com/zitadel/zitadel/internal/logstore/access"
)

func AccessInterceptor(svc *access.Service) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		resp, err := handler(ctx, req)

		var respStatus uint32
		grpcErr, ok := status.FromError(err)
		if ok {
			respStatus = uint32(grpcErr.Code())
		}

		md, _ := metadata.FromIncomingContext(ctx)

		instance := authz.GetInstance(ctx)

		svc.Handle(ctx, &logstore.AccessLogRecord{
			Timestamp:       time.Now(),
			Protocol:        logstore.GRPC,
			RequestURL:      info.FullMethod,
			ResponseStatus:  respStatus,
			RequestHeaders:  nil,
			ResponseHeaders: http.Header(md),
			InstanceID:      instance.InstanceID(),
			ProjectID:       instance.ProjectID(),
			RequestedDomain: instance.RequestedDomain(),
			RequestedHost:   instance.RequestedHost(),
		})
		return resp, err
	}
}
