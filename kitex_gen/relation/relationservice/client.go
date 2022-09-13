// Code generated by Kitex v0.3.2. DO NOT EDIT.

package relationservice

import (
	"context"
	"dousheng/kitex_gen/relation"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	CreateRelation(ctx context.Context, createRelationReq *relation.CreateRelationReq, callOptions ...callopt.Option) (r *relation.CreateRelationResp, err error)
	DeleteRelation(ctx context.Context, deleteRelationReq *relation.DeleteRelationReq, callOptions ...callopt.Option) (r *relation.DeleteRelationResp, err error)
	GetFollowers(ctx context.Context, getFollowersReq *relation.GetFollowersReq, callOptions ...callopt.Option) (r *relation.GetFollowersResp, err error)
	GetFollows(ctx context.Context, getFollowsReq *relation.GetFollowsReq, callOptions ...callopt.Option) (r *relation.GetFollowsResp, err error)
	CountFollowers(ctx context.Context, countFollowersReq *relation.CountFollowersReq, callOptions ...callopt.Option) (r *relation.CountFollowersResp, err error)
	CountFollows(ctx context.Context, countFollowsReq *relation.CountFollowsReq, callOptions ...callopt.Option) (r *relation.CountFollowsResp, err error)
	IsFollow(ctx context.Context, isFollowReq *relation.IsFollowReq, callOptions ...callopt.Option) (r *relation.IsFollowResp, err error)
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
	return &kRelationServiceClient{
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

type kRelationServiceClient struct {
	*kClient
}

func (p *kRelationServiceClient) CreateRelation(ctx context.Context, createRelationReq *relation.CreateRelationReq, callOptions ...callopt.Option) (r *relation.CreateRelationResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateRelation(ctx, createRelationReq)
}

func (p *kRelationServiceClient) DeleteRelation(ctx context.Context, deleteRelationReq *relation.DeleteRelationReq, callOptions ...callopt.Option) (r *relation.DeleteRelationResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteRelation(ctx, deleteRelationReq)
}

func (p *kRelationServiceClient) GetFollowers(ctx context.Context, getFollowersReq *relation.GetFollowersReq, callOptions ...callopt.Option) (r *relation.GetFollowersResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFollowers(ctx, getFollowersReq)
}

func (p *kRelationServiceClient) GetFollows(ctx context.Context, getFollowsReq *relation.GetFollowsReq, callOptions ...callopt.Option) (r *relation.GetFollowsResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFollows(ctx, getFollowsReq)
}

func (p *kRelationServiceClient) CountFollowers(ctx context.Context, countFollowersReq *relation.CountFollowersReq, callOptions ...callopt.Option) (r *relation.CountFollowersResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CountFollowers(ctx, countFollowersReq)
}

func (p *kRelationServiceClient) CountFollows(ctx context.Context, countFollowsReq *relation.CountFollowsReq, callOptions ...callopt.Option) (r *relation.CountFollowsResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CountFollows(ctx, countFollowsReq)
}

func (p *kRelationServiceClient) IsFollow(ctx context.Context, isFollowReq *relation.IsFollowReq, callOptions ...callopt.Option) (r *relation.IsFollowResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.IsFollow(ctx, isFollowReq)
}