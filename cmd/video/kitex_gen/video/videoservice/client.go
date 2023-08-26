// Code generated by Kitex v0.6.2. DO NOT EDIT.

package videoservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	video "github.com/ozline/tiktok/cmd/video/kitex_gen/video"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Feed(ctx context.Context, Req *video.FeedRequest, callOptions ...callopt.Option) (r *video.FeedResponse, err error)
	PutVideo(ctx context.Context, Req *video.PutVideoRequest, callOptions ...callopt.Option) (r *video.PutVideoResponse, err error)
	GetFavoriteVideoInfo(ctx context.Context, Req *video.GetFavoriteVideoInfoRequest, callOptions ...callopt.Option) (r *video.GetFavoriteVideoInfoResponse, err error)
	GetPublishList(ctx context.Context, Req *video.GetPublishListRequest, callOptions ...callopt.Option) (r *video.GetPublishListResponse, err error)
	GetWorkCount(ctx context.Context, Req *video.GetWorkCountRequest, callOptions ...callopt.Option) (r *video.GetWorkCountResponse, err error)
	GetVideoIDByUid(ctx context.Context, Req *video.GetVideoIDByUidRequset, callOptions ...callopt.Option) (r *video.GetVideoIDByUidResponse, err error)
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
	return &kVideoServiceClient{
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

type kVideoServiceClient struct {
	*kClient
}

func (p *kVideoServiceClient) Feed(ctx context.Context, Req *video.FeedRequest, callOptions ...callopt.Option) (r *video.FeedResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Feed(ctx, Req)
}

func (p *kVideoServiceClient) PutVideo(ctx context.Context, Req *video.PutVideoRequest, callOptions ...callopt.Option) (r *video.PutVideoResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.PutVideo(ctx, Req)
}

func (p *kVideoServiceClient) GetFavoriteVideoInfo(ctx context.Context, Req *video.GetFavoriteVideoInfoRequest, callOptions ...callopt.Option) (r *video.GetFavoriteVideoInfoResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFavoriteVideoInfo(ctx, Req)
}

func (p *kVideoServiceClient) GetPublishList(ctx context.Context, Req *video.GetPublishListRequest, callOptions ...callopt.Option) (r *video.GetPublishListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetPublishList(ctx, Req)
}

func (p *kVideoServiceClient) GetWorkCount(ctx context.Context, Req *video.GetWorkCountRequest, callOptions ...callopt.Option) (r *video.GetWorkCountResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetWorkCount(ctx, Req)
}

func (p *kVideoServiceClient) GetVideoIDByUid(ctx context.Context, Req *video.GetVideoIDByUidRequset, callOptions ...callopt.Option) (r *video.GetVideoIDByUidResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetVideoIDByUid(ctx, Req)
}
