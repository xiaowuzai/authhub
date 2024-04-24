package service

import "context"

func (s *Service) RecognizeTextAuto(ctx context.Context, in *RecognizeTextAutoRequest) (*RecognizeTextAutoResponse, error) {
	return s.recognizeRepo.RecognizeTextAuto(ctx, in)
}

func (s *Service) OpenOcr(ctx context.Context, in *OpenOcrRequest) (*OpenOcrResponse, error) {
	return s.recognizeRepo.OpenOcr(ctx, in)
}

func (s *Service) GenerateImg(ctx context.Context, in *GenerateImgRequest) (*GenerateImgResponse, error) {
	return s.recognizeRepo.GenerateImg(ctx, in)
}
