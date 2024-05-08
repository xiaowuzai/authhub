package backend

import (
	"context"
	"fmt"
	"time"

	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"github.com/jinzhu/copier"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"

	"git.corp.zgcszkw.com/authhub/config"
	"git.corp.zgcszkw.com/authhub/outerpb/openocr"
	"git.corp.zgcszkw.com/authhub/service"
)

var _ service.RecognizeRepo = (*OpenocrClient)(nil)

type OpenocrClient struct {
	conn   *grpc.ClientConn
	client openocr.OpenocrClient
}

func NewOpenocrClient(conf *config.Config) (service.RecognizeRepo, error) {

	target := fmt.Sprintf("%s:%s", conf.SiriusHost, conf.SiriusPort)
	log.Printf("DialBackend Sirius Host: %s", target)

	conn, err := grpc.Dial(
		// fmt.Sprintf("dns:///%s:%s", conf.SiriusHost, conf.SiriusPort),
		target,
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor()),
		// grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)),
		grpc.WithConnectParams(grpc.ConnectParams{
			Backoff: backoff.Config{
				BaseDelay:  3 * time.Second, // 重试基础等待时间
				Multiplier: 1.5,             //乘数因子
				Jitter:     0.2,
			},
			MinConnectTimeout: 2 * time.Second,
		}),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(20000000)),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                time.Second * 20,
			Timeout:             time.Second * 1,
			PermitWithoutStream: true,
		}),
	)
	if err != nil {
		log.Error().
			Err(err).
			Msg("DialBackend open ocr failed.")
		return nil, err
	}

	client := openocr.NewOpenocrClient(conn)
	return &OpenocrClient{
		conn:   conn,
		client: client,
	}, nil
}

func (sc *OpenocrClient) RecognizeTextAuto(ctx context.Context, in *service.RecognizeTextAutoRequest) (*service.RecognizeTextAutoResponse, error) {
	req := &openocr.RecognizeTextAutoRequest{}
	err := copier.Copy(req, in)
	if err != nil {
		return nil, err
	}
	res, err := sc.client.RecognizeTextAuto(ctx, req)
	if err != nil {
		return nil, err
	}

	result := &service.RecognizeTextAutoResponse{}
	if err := copier.Copy(result, res); err != nil {
		return nil, err
	}
	return result, nil
}

func (sc *OpenocrClient) OpenOcr(ctx context.Context, in *service.OpenOcrRequest) (*service.OpenOcrResponse, error) {
	req := &openocr.OpenocrRequest{}
	err := copier.Copy(req, in)
	if err != nil {
		return nil, err
	}
	res, err := sc.client.Openocr(ctx, req)
	if err != nil {
		return nil, err
	}

	result := &service.OpenOcrResponse{}
	if err := copier.Copy(result, res); err != nil {
		return nil, err
	}
	return result, nil
}

func (sc *OpenocrClient) GenerateImg(ctx context.Context, in *service.GenerateImgRequest) (*service.GenerateImgResponse, error) {
	req := &openocr.GenerateImgRequest{}
	err := copier.Copy(req, in)
	if err != nil {
		return nil, err
	}
	res, err := sc.client.GenerateImg(ctx, req)
	if err != nil {
		return nil, err
	}

	result := &service.GenerateImgResponse{
		Id:     res.Id,
		Image:  service.BytesToBase64String(res.Image),
		Height: res.Height,
		Width:  res.Width,
	}
	return result, nil
}
