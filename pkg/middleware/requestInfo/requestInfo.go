package requestInfo

import (
	"context"

	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	"github.com/go-kratos/kratos/v2/middleware"
)

const DomainKey = "x-md-global-domain"

// SetRequestInfo 设置Request信息
func SetRequestInfo() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			// 获取 Domain 值
			domain := getDomainFromContext(ctx)

			// 将 Domain 添加到上下文
			ctx = metadata.AppendToClientContext(ctx, DomainKey, domain)
			ctx = context.WithValue(ctx, DomainKey, domain)

			return handler(ctx, req)
		}
	}
}

// 从上下文中获取 Domain 的辅助函数
func getDomainFromContext(ctx context.Context) string {
	// 从元信息中获取 domain
	if md, ok := metadata.FromServerContext(ctx); ok {
		if domain := md.Get(DomainKey); domain != "" {
			return domain
		}
	}

	// 从请求信息中获取 domain
	if tr, ok := transport.FromServerContext(ctx); ok {
		switch tr.Kind() {
		case transport.KindHTTP:
			if domain := tr.(*http.Transport).RequestHeader().Get(DomainKey); domain != "" {
				return domain
			}
		case transport.KindGRPC:
			if domain := tr.(*grpc.Transport).RequestHeader().Get(DomainKey); domain != "" {
				return domain
			}
		}
	}

	return ""
}
