// Code generated by Kitex v0.3.2. DO NOT EDIT.

package userservice

import (
	"context"
	"dousheng/kitex_gen/user"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	ContainsName(ctx context.Context, containsNameReq *user.ContainsNameReq, callOptions ...callopt.Option) (r *user.ContainsNameResp, err error)
	CreateUser(ctx context.Context, createUserReq *user.CreateUserReq, callOptions ...callopt.Option) (r *user.CreateUserResp, err error)
	GetUserByName(ctx context.Context, getUserByNameReq *user.GetUserByNameReq, callOptions ...callopt.Option) (r *user.GetUserByNameResp, err error)
	GetUserById(ctx context.Context, getUserByIdReq *user.GetUserByIdReq, callOptions ...callopt.Option) (r *user.GetUserByIdResp, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kUserServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kUserServiceClient struct {
	*kClient
}

func (p *kUserServiceClient) ContainsName(ctx context.Context, containsNameReq *user.ContainsNameReq, callOptions ...callopt.Option) (r *user.ContainsNameResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ContainsName(ctx, containsNameReq)
}

func (p *kUserServiceClient) CreateUser(ctx context.Context, createUserReq *user.CreateUserReq, callOptions ...callopt.Option) (r *user.CreateUserResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateUser(ctx, createUserReq)
}

func (p *kUserServiceClient) GetUserByName(ctx context.Context, getUserByNameReq *user.GetUserByNameReq, callOptions ...callopt.Option) (r *user.GetUserByNameResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUserByName(ctx, getUserByNameReq)
}

func (p *kUserServiceClient) GetUserById(ctx context.Context, getUserByIdReq *user.GetUserByIdReq, callOptions ...callopt.Option) (r *user.GetUserByIdResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUserById(ctx, getUserByIdReq)
}
