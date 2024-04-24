package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"git.corp.zgcszkw.com/authhub/service"
)

// 识别文件
// @Summary 识别文件
// @Schemes
// @Description 识别文件
// @Tags openocr
// @Accept json
// @Produce json
// @Param message body service.RecognizeTextAutoRequest true "请求参数"
// @Success 200 {object} service.RecognizeTextAutoResponse
// @Router /v1/recognize_text [post]
func (s *Server) RecognizeTextAuto(ctx *gin.Context) {
	req := &service.RecognizeTextAutoRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		log.Error().
			Err(err).
			Msgf("recognize text failed. request: %+v", *req)
		return
	}

	res, err := s.service.RecognizeTextAuto(ctx, req)
	if err != nil {
		log.Error().
			Err(err).
			Msgf("recognize text failed. request: %+v", *req)
		ResponseError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	log.Info().
		Msgf("recognize text success. request: %+v", *req)
	ResponseSuccess(ctx, res)
}

// 外语检测识别
// @Summary 外语检测识别
// @Schemes
// @Description 外语检测识别
// @Tags openocr
// @Accept json
// @Produce json
// @Param message body service.OpenOcrRequest true "请求参数"
// @Success 200 {object} service.OpenOcrResponse
// @Router /v1/open_ocr [post]
func (s *Server) OpenOcr(ctx *gin.Context) {
	req := &service.OpenOcrRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		log.Error().
			Err(err).
			Msgf("open ocr failed. request: %+v", *req)
		return
	}

	res, err := s.service.OpenOcr(ctx, req)
	if err != nil {
		log.Error().
			Err(err).
			Msgf("open ocr failed. request: %+v", *req)
		ResponseError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	log.Info().
		Msgf("open ocr success. request: %+v", *req)
	ResponseSuccess(ctx, res)
}

// 生成图片
// @Summary 生成图片
// @Schemes
// @Description 生成图片
// @Tags openocr
// @Accept json
// @Produce json
// @Param message body service.GenerateImgRequest true "请求参数"
// @Success 200 {object} service.GenerateImgResponse
// @Router /v1/generate_img [post]
func (s *Server) GenerateImg(ctx *gin.Context) {
	req := &service.GenerateImgRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		log.Error().
			Err(err).
			Msgf("generate img failed. request: %+v", *req)
		return
	}

	res, err := s.service.GenerateImg(ctx, req)
	if err != nil {
		log.Error().
			Err(err).
			Msgf("generate img failed. request: %+v", *req)
		ResponseError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	log.Info().
		Msgf("generate img success. request: %+v", *req)
	ResponseSuccess(ctx, res)
}
