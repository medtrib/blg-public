// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.1
// - protoc             v5.28.2
// source: sms/v1/sms.proto

package v1

import (
	v1 "blg-ext/api/ext/v1"
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationSmsSendSms = "/api.sms.v1.Sms/SendSms"

type SmsHTTPServer interface {
	// SendSms 发送短信
	SendSms(context.Context, *SendSmsRequest) (*v1.Reply, error)
}

func RegisterSmsHTTPServer(s *http.Server, srv SmsHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/sms/sendSms", _Sms_SendSms0_HTTP_Handler(srv))
}

func _Sms_SendSms0_HTTP_Handler(srv SmsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SendSmsRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSmsSendSms)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SendSms(ctx, req.(*SendSmsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.Reply)
		return ctx.Result(200, reply)
	}
}

type SmsHTTPClient interface {
	SendSms(ctx context.Context, req *SendSmsRequest, opts ...http.CallOption) (rsp *v1.Reply, err error)
}

type SmsHTTPClientImpl struct {
	cc *http.Client
}

func NewSmsHTTPClient(client *http.Client) SmsHTTPClient {
	return &SmsHTTPClientImpl{client}
}

func (c *SmsHTTPClientImpl) SendSms(ctx context.Context, in *SendSmsRequest, opts ...http.CallOption) (*v1.Reply, error) {
	var out v1.Reply
	pattern := "/v1/sms/sendSms"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSmsSendSms))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}