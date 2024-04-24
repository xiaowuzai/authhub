package service

type Service struct {
	recognizeRepo RecognizeRepo
}

func NewService(recognizeRepo RecognizeRepo) *Service {
	return &Service{
		recognizeRepo: recognizeRepo,
	}
}
