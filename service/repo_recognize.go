package service

import (
	"context"
)

type RecognizeRepo interface {
	RecognizeTextAuto(ctx context.Context, in *RecognizeTextAutoRequest) (*RecognizeTextAutoResponse, error)
	OpenOcr(ctx context.Context, in *OpenOcrRequest) (*OpenOcrResponse, error)
	GenerateImg(ctx context.Context, in *GenerateImgRequest) (*GenerateImgResponse, error)
}

type RecognizeTextAutoRequest struct {
	Id string `json:"id,omitempty" binding:"required"`
}
type RecognizeTextAutoResponse struct {
	JsonText    *Text       `json:"json_text,omitempty"`
	ExcelFileId string      `json:"excel_file_id,omitempty"`
	LineInfos   []*LineInfo `json:"line_infos,omitempty"`
	Angle       float32     `json:"angle,omitempty"`
	Scale       float32     `json:"scale,omitempty"`
}

type OpenOcrRequest struct {
	Id   string `json:"id,omitempty" binding:"required"`
	Lang string `json:"lang,omitempty" binding:"required"`
}
type OpenOcrResponse struct {
	Boxes []*TextBox `json:"boxes,omitempty"`
	Id    string     `json:"id,omitempty"`
}

type GenerateImgRequest struct {
	Id           string     `json:"id,omitempty" binding:"required"`
	Boxes        []*TextBox `json:"boxes,omitempty" binding:"required"`
	FontFileName string     `json:"font_file_name,omitempty" binding:"required"`
}
type GenerateImgResponse struct {
	Id     string `json:"id,omitempty"`
	Image  []byte `json:"image,omitempty"`
	Height int32  `json:"height,omitempty"`
	Width  int32  `json:"width,omitempty"`
}

type TextBox struct {
	Text string `json:"text,omitempty" binding:"required"` // 需要 base64
	XMin int32  `json:"x_min,omitempty" binding:"required"`
	YMin int32  `json:"y_min,omitempty" binding:"required"`
	XMax int32  `json:"x_max,omitempty" binding:"required"`
	YMax int32  `json:"y_max,omitempty" binding:"required"`
}

type LineInfo struct {
	Text     string     `json:"text,omitempty"`
	TextBoxs []*TextBox `json:"text_boxs,omitempty"`
}

type Text struct {
	Content string `json:"content,omitempty"`
}
