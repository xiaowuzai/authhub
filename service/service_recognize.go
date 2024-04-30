package service

import "context"

func (s *Service) RecognizeTextAuto(ctx context.Context, in *RecognizeTextAutoRequest) (*RecognizeTextAutoResponse, error) {
	return s.recognizeRepo.RecognizeTextAuto(ctx, in)
}

func (s *Service) OpenOcr(ctx context.Context, in *OpenOcrRequest) (*OpenOcrResponse, error) {
	res, err := s.recognizeRepo.OpenOcr(ctx, in)
	if err != nil {
		return nil, err
	}

	for i := range res.Boxes {
		res.Boxes[i].Text = StrToBase64(res.Boxes[i].Text)
	}
	return res, nil
}

func (s *Service) GenerateImg(ctx context.Context, in *GenerateImgRequest) (*GenerateImgResponse, error) {

	for i := range in.Boxes {
		text, err := StrFromBase64(in.Boxes[i].Text)
		if err != nil {
			return nil, err
		}

		in.Boxes[i].Text = text
	}

	return s.recognizeRepo.GenerateImg(ctx, in)
}
