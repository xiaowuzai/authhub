// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.20.3
// source: openocr/openocr.proto

package openocr

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OcrInvoiceType int32

const (
	OcrInvoiceType_INVOICE_TYPE_UNKNOWN     OcrInvoiceType = 0
	OcrInvoiceType_INVOICE_TYPE_VAT         OcrInvoiceType = 1 //增值税发票
	OcrInvoiceType_INVOICE_TYPE_TRAIN       OcrInvoiceType = 2 //火车票
	OcrInvoiceType_INVOICE_TYPE_TAXI        OcrInvoiceType = 3 //出租车票
	OcrInvoiceType_INVOICE_TYPE_HIGHWAYTOLL OcrInvoiceType = 4 //高速过路费发票
	OcrInvoiceType_INVOICE_TYPE_QUOTA       OcrInvoiceType = 5 //定额发票
	OcrInvoiceType_INVOICE_TYPE_FLIGHT      OcrInvoiceType = 6 //飞机票
	OcrInvoiceType_INVOICE_TYPE_COACH       OcrInvoiceType = 7 //汽车客运发票
	OcrInvoiceType_INVOICE_TYPE_BUS         OcrInvoiceType = 8 //公交车票
	OcrInvoiceType_INVOICE_TYPE_ELE_VAT     OcrInvoiceType = 9 //电子发票
)

// Enum value maps for OcrInvoiceType.
var (
	OcrInvoiceType_name = map[int32]string{
		0: "INVOICE_TYPE_UNKNOWN",
		1: "INVOICE_TYPE_VAT",
		2: "INVOICE_TYPE_TRAIN",
		3: "INVOICE_TYPE_TAXI",
		4: "INVOICE_TYPE_HIGHWAYTOLL",
		5: "INVOICE_TYPE_QUOTA",
		6: "INVOICE_TYPE_FLIGHT",
		7: "INVOICE_TYPE_COACH",
		8: "INVOICE_TYPE_BUS",
		9: "INVOICE_TYPE_ELE_VAT",
	}
	OcrInvoiceType_value = map[string]int32{
		"INVOICE_TYPE_UNKNOWN":     0,
		"INVOICE_TYPE_VAT":         1,
		"INVOICE_TYPE_TRAIN":       2,
		"INVOICE_TYPE_TAXI":        3,
		"INVOICE_TYPE_HIGHWAYTOLL": 4,
		"INVOICE_TYPE_QUOTA":       5,
		"INVOICE_TYPE_FLIGHT":      6,
		"INVOICE_TYPE_COACH":       7,
		"INVOICE_TYPE_BUS":         8,
		"INVOICE_TYPE_ELE_VAT":     9,
	}
)

func (x OcrInvoiceType) Enum() *OcrInvoiceType {
	p := new(OcrInvoiceType)
	*p = x
	return p
}

func (x OcrInvoiceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OcrInvoiceType) Descriptor() protoreflect.EnumDescriptor {
	return file_openocr_openocr_proto_enumTypes[0].Descriptor()
}

func (OcrInvoiceType) Type() protoreflect.EnumType {
	return &file_openocr_openocr_proto_enumTypes[0]
}

func (x OcrInvoiceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OcrInvoiceType.Descriptor instead.
func (OcrInvoiceType) EnumDescriptor() ([]byte, []int) {
	return file_openocr_openocr_proto_rawDescGZIP(), []int{0}
}

type OcrLang int32

const (
	OcrLang_OcrLang_INVALID OcrLang = 0
	OcrLang_OcrLang_LAT     OcrLang = 1
	OcrLang_OcrLang_JAP     OcrLang = 2
	OcrLang_OcrLang_KR      OcrLang = 3
	OcrLang_OcrLang_RU      OcrLang = 4
	OcrLang_OcrLang_ARA     OcrLang = 5
)

// Enum value maps for OcrLang.
var (
	OcrLang_name = map[int32]string{
		0: "OcrLang_INVALID",
		1: "OcrLang_LAT",
		2: "OcrLang_JAP",
		3: "OcrLang_KR",
		4: "OcrLang_RU",
		5: "OcrLang_ARA",
	}
	OcrLang_value = map[string]int32{
		"OcrLang_INVALID": 0,
		"OcrLang_LAT":     1,
		"OcrLang_JAP":     2,
		"OcrLang_KR":      3,
		"OcrLang_RU":      4,
		"OcrLang_ARA":     5,
	}
)

func (x OcrLang) Enum() *OcrLang {
	p := new(OcrLang)
	*p = x
	return p
}

func (x OcrLang) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OcrLang) Descriptor() protoreflect.EnumDescriptor {
	return file_openocr_openocr_proto_enumTypes[1].Descriptor()
}

func (OcrLang) Type() protoreflect.EnumType {
	return &file_openocr_openocr_proto_enumTypes[1]
}

func (x OcrLang) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OcrLang.Descriptor instead.
func (OcrLang) EnumDescriptor() ([]byte, []int) {
	return file_openocr_openocr_proto_rawDescGZIP(), []int{1}
}

type TextBox struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XMin int32  `protobuf:"varint,2,opt,name=x_min,json=xMin,proto3" json:"x_min,omitempty"`
	YMin int32  `protobuf:"varint,3,opt,name=y_min,json=yMin,proto3" json:"y_min,omitempty"`
	XMax int32  `protobuf:"varint,4,opt,name=x_max,json=xMax,proto3" json:"x_max,omitempty"`
	YMax int32  `protobuf:"varint,5,opt,name=y_max,json=yMax,proto3" json:"y_max,omitempty"`
}

func (x *TextBox) Reset() {
	*x = TextBox{}
	if protoimpl.UnsafeEnabled {
		mi := &file_openocr_openocr_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextBox) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextBox) ProtoMessage() {}

func (x *TextBox) ProtoReflect() protoreflect.Message {
	mi := &file_openocr_openocr_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextBox.ProtoReflect.Descriptor instead.
func (*TextBox) Descriptor() ([]byte, []int) {
	return file_openocr_openocr_proto_rawDescGZIP(), []int{0}
}

func (x *TextBox) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *TextBox) GetXMin() int32 {
	if x != nil {
		return x.XMin
	}
	return 0
}

func (x *TextBox) GetYMin() int32 {
	if x != nil {
		return x.YMin
	}
	return 0
}

func (x *TextBox) GetXMax() int32 {
	if x != nil {
		return x.XMax
	}
	return 0
}

func (x *TextBox) GetYMax() int32 {
	if x != nil {
		return x.YMax
	}
	return 0
}

type LineInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text     string     `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	TextBoxs []*TextBox `protobuf:"bytes,6,rep,name=text_boxs,json=textBoxs,proto3" json:"text_boxs,omitempty"`
}

func (x *LineInfo) Reset() {
	*x = LineInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_openocr_openocr_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LineInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LineInfo) ProtoMessage() {}

func (x *LineInfo) ProtoReflect() protoreflect.Message {
	mi := &file_openocr_openocr_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LineInfo.ProtoReflect.Descriptor instead.
func (*LineInfo) Descriptor() ([]byte, []int) {
	return file_openocr_openocr_proto_rawDescGZIP(), []int{1}
}

func (x *LineInfo) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *LineInfo) GetTextBoxs() []*TextBox {
	if x != nil {
		return x.TextBoxs
	}
	return nil
}

type RecognizeTextAutoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RecognizeTextAutoRequest) Reset() {
	*x = RecognizeTextAutoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_openocr_openocr_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecognizeTextAutoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecognizeTextAutoRequest) ProtoMessage() {}

func (x *RecognizeTextAutoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_openocr_openocr_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecognizeTextAutoRequest.ProtoReflect.Descriptor instead.
func (*RecognizeTextAutoRequest) Descriptor() ([]byte, []int) {
	return file_openocr_openocr_proto_rawDescGZIP(), []int{2}
}

func (x *RecognizeTextAutoRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Text struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Text) Reset() {
	*x = Text{}
	if protoimpl.UnsafeEnabled {
		mi := &file_openocr_openocr_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Text) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Text) ProtoMessage() {}

func (x *Text) ProtoReflect() protoreflect.Message {
	mi := &file_openocr_openocr_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Text.ProtoReflect.Descriptor instead.
func (*Text) Descriptor() ([]byte, []int) {
	return file_openocr_openocr_proto_rawDescGZIP(), []int{3}
}

func (x *Text) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type RecognizeTextResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JsonText    *Text       `protobuf:"bytes,1,opt,name=json_text,json=jsonText,proto3" json:"json_text,omitempty"`
	ExcelFileId string      `protobuf:"bytes,2,opt,name=excel_file_id,json=excelFileId,proto3" json:"excel_file_id,omitempty"`
	LineInfos   []*LineInfo `protobuf:"bytes,3,rep,name=line_infos,json=lineInfos,proto3" json:"line_infos,omitempty"`
	Angle       float32     `protobuf:"fixed32,4,opt,name=angle,proto3" json:"angle,omitempty"`
	Scale       float32     `protobuf:"fixed32,5,opt,name=scale,proto3" json:"scale,omitempty"`
}

func (x *RecognizeTextResponse) Reset() {
	*x = RecognizeTextResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_openocr_openocr_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecognizeTextResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecognizeTextResponse) ProtoMessage() {}

func (x *RecognizeTextResponse) ProtoReflect() protoreflect.Message {
	mi := &file_openocr_openocr_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecognizeTextResponse.ProtoReflect.Descriptor instead.
func (*RecognizeTextResponse) Descriptor() ([]byte, []int) {
	return file_openocr_openocr_proto_rawDescGZIP(), []int{4}
}

func (x *RecognizeTextResponse) GetJsonText() *Text {
	if x != nil {
		return x.JsonText
	}
	return nil
}

func (x *RecognizeTextResponse) GetExcelFileId() string {
	if x != nil {
		return x.ExcelFileId
	}
	return ""
}

func (x *RecognizeTextResponse) GetLineInfos() []*LineInfo {
	if x != nil {
		return x.LineInfos
	}
	return nil
}

func (x *RecognizeTextResponse) GetAngle() float32 {
	if x != nil {
		return x.Angle
	}
	return 0
}

func (x *RecognizeTextResponse) GetScale() float32 {
	if x != nil {
		return x.Scale
	}
	return 0
}

type OpenocrRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Lang string `protobuf:"bytes,2,opt,name=lang,proto3" json:"lang,omitempty"`
}

func (x *OpenocrRequest) Reset() {
	*x = OpenocrRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_openocr_openocr_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenocrRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenocrRequest) ProtoMessage() {}

func (x *OpenocrRequest) ProtoReflect() protoreflect.Message {
	mi := &file_openocr_openocr_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenocrRequest.ProtoReflect.Descriptor instead.
func (*OpenocrRequest) Descriptor() ([]byte, []int) {
	return file_openocr_openocr_proto_rawDescGZIP(), []int{5}
}

func (x *OpenocrRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *OpenocrRequest) GetLang() string {
	if x != nil {
		return x.Lang
	}
	return ""
}

type OpenocrResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Boxes []*TextBox `protobuf:"bytes,1,rep,name=boxes,proto3" json:"boxes,omitempty"`
	Id    string     `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *OpenocrResponse) Reset() {
	*x = OpenocrResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_openocr_openocr_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenocrResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenocrResponse) ProtoMessage() {}

func (x *OpenocrResponse) ProtoReflect() protoreflect.Message {
	mi := &file_openocr_openocr_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenocrResponse.ProtoReflect.Descriptor instead.
func (*OpenocrResponse) Descriptor() ([]byte, []int) {
	return file_openocr_openocr_proto_rawDescGZIP(), []int{6}
}

func (x *OpenocrResponse) GetBoxes() []*TextBox {
	if x != nil {
		return x.Boxes
	}
	return nil
}

func (x *OpenocrResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GenerateImgRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Boxes        []*TextBox `protobuf:"bytes,2,rep,name=boxes,proto3" json:"boxes,omitempty"`
	FontFileName string     `protobuf:"bytes,3,opt,name=font_file_name,json=fontFileName,proto3" json:"font_file_name,omitempty"`
}

func (x *GenerateImgRequest) Reset() {
	*x = GenerateImgRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_openocr_openocr_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateImgRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateImgRequest) ProtoMessage() {}

func (x *GenerateImgRequest) ProtoReflect() protoreflect.Message {
	mi := &file_openocr_openocr_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateImgRequest.ProtoReflect.Descriptor instead.
func (*GenerateImgRequest) Descriptor() ([]byte, []int) {
	return file_openocr_openocr_proto_rawDescGZIP(), []int{7}
}

func (x *GenerateImgRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GenerateImgRequest) GetBoxes() []*TextBox {
	if x != nil {
		return x.Boxes
	}
	return nil
}

func (x *GenerateImgRequest) GetFontFileName() string {
	if x != nil {
		return x.FontFileName
	}
	return ""
}

type GenerateImgResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Image  []byte `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
	Height int32  `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	Width  int32  `protobuf:"varint,4,opt,name=width,proto3" json:"width,omitempty"`
}

func (x *GenerateImgResponse) Reset() {
	*x = GenerateImgResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_openocr_openocr_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateImgResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateImgResponse) ProtoMessage() {}

func (x *GenerateImgResponse) ProtoReflect() protoreflect.Message {
	mi := &file_openocr_openocr_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateImgResponse.ProtoReflect.Descriptor instead.
func (*GenerateImgResponse) Descriptor() ([]byte, []int) {
	return file_openocr_openocr_proto_rawDescGZIP(), []int{8}
}

func (x *GenerateImgResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GenerateImgResponse) GetImage() []byte {
	if x != nil {
		return x.Image
	}
	return nil
}

func (x *GenerateImgResponse) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *GenerateImgResponse) GetWidth() int32 {
	if x != nil {
		return x.Width
	}
	return 0
}

var File_openocr_openocr_proto protoreflect.FileDescriptor

var file_openocr_openocr_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6f, 0x70, 0x65, 0x6e, 0x6f, 0x63, 0x72, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x6f, 0x63,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6f, 0x70, 0x65, 0x6e, 0x6f, 0x63, 0x72,
	0x22, 0x71, 0x0a, 0x07, 0x54, 0x65, 0x78, 0x74, 0x42, 0x6f, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12,
	0x13, 0x0a, 0x05, 0x78, 0x5f, 0x6d, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x78, 0x4d, 0x69, 0x6e, 0x12, 0x13, 0x0a, 0x05, 0x79, 0x5f, 0x6d, 0x69, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x79, 0x4d, 0x69, 0x6e, 0x12, 0x13, 0x0a, 0x05, 0x78, 0x5f, 0x6d,
	0x61, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x78, 0x4d, 0x61, 0x78, 0x12, 0x13,
	0x0a, 0x05, 0x79, 0x5f, 0x6d, 0x61, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x79,
	0x4d, 0x61, 0x78, 0x22, 0x4d, 0x0a, 0x08, 0x4c, 0x69, 0x6e, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x12, 0x2d, 0x0a, 0x09, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x62, 0x6f, 0x78, 0x73,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x6f, 0x63, 0x72,
	0x2e, 0x54, 0x65, 0x78, 0x74, 0x42, 0x6f, 0x78, 0x52, 0x08, 0x74, 0x65, 0x78, 0x74, 0x42, 0x6f,
	0x78, 0x73, 0x22, 0x2a, 0x0a, 0x18, 0x52, 0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x7a, 0x65, 0x54,
	0x65, 0x78, 0x74, 0x41, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x20,
	0x0a, 0x04, 0x54, 0x65, 0x78, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x22, 0xc5, 0x01, 0x0a, 0x15, 0x52, 0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x7a, 0x65, 0x54, 0x65,
	0x78, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x09, 0x6a, 0x73,
	0x6f, 0x6e, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x6f, 0x70, 0x65, 0x6e, 0x6f, 0x63, 0x72, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x52, 0x08, 0x6a, 0x73,
	0x6f, 0x6e, 0x54, 0x65, 0x78, 0x74, 0x12, 0x22, 0x0a, 0x0d, 0x65, 0x78, 0x63, 0x65, 0x6c, 0x5f,
	0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65,
	0x78, 0x63, 0x65, 0x6c, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x0a, 0x6c, 0x69,
	0x6e, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x6f, 0x63, 0x72, 0x2e, 0x4c, 0x69, 0x6e, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x09, 0x6c, 0x69, 0x6e, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x61, 0x6e, 0x67, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x61, 0x6e, 0x67,
	0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x05, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x22, 0x34, 0x0a, 0x0e, 0x4f, 0x70, 0x65, 0x6e,
	0x6f, 0x63, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61,
	0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x22, 0x49,
	0x0a, 0x0f, 0x4f, 0x70, 0x65, 0x6e, 0x6f, 0x63, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x26, 0x0a, 0x05, 0x62, 0x6f, 0x78, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x6f, 0x63, 0x72, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x42,
	0x6f, 0x78, 0x52, 0x05, 0x62, 0x6f, 0x78, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x72, 0x0a, 0x12, 0x47, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x49, 0x6d, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x26, 0x0a, 0x05, 0x62, 0x6f, 0x78, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x6f, 0x63, 0x72, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x42, 0x6f, 0x78,
	0x52, 0x05, 0x62, 0x6f, 0x78, 0x65, 0x73, 0x12, 0x24, 0x0a, 0x0e, 0x66, 0x6f, 0x6e, 0x74, 0x5f,
	0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x66, 0x6f, 0x6e, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x69, 0x0a,
	0x13, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x49, 0x6d, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x2a, 0x86, 0x02, 0x0a, 0x0e, 0x4f, 0x63, 0x72,
	0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x14, 0x49,
	0x4e, 0x56, 0x4f, 0x49, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x49, 0x4e, 0x56, 0x4f, 0x49, 0x43, 0x45,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x56, 0x41, 0x54, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x49,
	0x4e, 0x56, 0x4f, 0x49, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x52, 0x41, 0x49,
	0x4e, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x49, 0x4e, 0x56, 0x4f, 0x49, 0x43, 0x45, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x54, 0x41, 0x58, 0x49, 0x10, 0x03, 0x12, 0x1c, 0x0a, 0x18, 0x49, 0x4e,
	0x56, 0x4f, 0x49, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x48, 0x49, 0x47, 0x48, 0x57,
	0x41, 0x59, 0x54, 0x4f, 0x4c, 0x4c, 0x10, 0x04, 0x12, 0x16, 0x0a, 0x12, 0x49, 0x4e, 0x56, 0x4f,
	0x49, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x51, 0x55, 0x4f, 0x54, 0x41, 0x10, 0x05,
	0x12, 0x17, 0x0a, 0x13, 0x49, 0x4e, 0x56, 0x4f, 0x49, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x46, 0x4c, 0x49, 0x47, 0x48, 0x54, 0x10, 0x06, 0x12, 0x16, 0x0a, 0x12, 0x49, 0x4e, 0x56,
	0x4f, 0x49, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f, 0x41, 0x43, 0x48, 0x10,
	0x07, 0x12, 0x14, 0x0a, 0x10, 0x49, 0x4e, 0x56, 0x4f, 0x49, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x42, 0x55, 0x53, 0x10, 0x08, 0x12, 0x18, 0x0a, 0x14, 0x49, 0x4e, 0x56, 0x4f, 0x49,
	0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x4c, 0x45, 0x5f, 0x56, 0x41, 0x54, 0x10,
	0x09, 0x2a, 0x71, 0x0a, 0x07, 0x4f, 0x63, 0x72, 0x4c, 0x61, 0x6e, 0x67, 0x12, 0x13, 0x0a, 0x0f,
	0x4f, 0x63, 0x72, 0x4c, 0x61, 0x6e, 0x67, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10,
	0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x4f, 0x63, 0x72, 0x4c, 0x61, 0x6e, 0x67, 0x5f, 0x4c, 0x41, 0x54,
	0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x4f, 0x63, 0x72, 0x4c, 0x61, 0x6e, 0x67, 0x5f, 0x4a, 0x41,
	0x50, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x4f, 0x63, 0x72, 0x4c, 0x61, 0x6e, 0x67, 0x5f, 0x4b,
	0x52, 0x10, 0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x4f, 0x63, 0x72, 0x4c, 0x61, 0x6e, 0x67, 0x5f, 0x52,
	0x55, 0x10, 0x04, 0x12, 0x0f, 0x0a, 0x0b, 0x4f, 0x63, 0x72, 0x4c, 0x61, 0x6e, 0x67, 0x5f, 0x41,
	0x52, 0x41, 0x10, 0x05, 0x32, 0xef, 0x01, 0x0a, 0x07, 0x4f, 0x70, 0x65, 0x6e, 0x6f, 0x63, 0x72,
	0x12, 0x58, 0x0a, 0x11, 0x52, 0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x7a, 0x65, 0x54, 0x65, 0x78,
	0x74, 0x41, 0x75, 0x74, 0x6f, 0x12, 0x21, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x6f, 0x63, 0x72, 0x2e,
	0x52, 0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x7a, 0x65, 0x54, 0x65, 0x78, 0x74, 0x41, 0x75, 0x74,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x6f,
	0x63, 0x72, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x7a, 0x65, 0x54, 0x65, 0x78, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x07, 0x4f, 0x70,
	0x65, 0x6e, 0x6f, 0x63, 0x72, 0x12, 0x17, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x6f, 0x63, 0x72, 0x2e,
	0x4f, 0x70, 0x65, 0x6e, 0x6f, 0x63, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18,
	0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x6f, 0x63, 0x72, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x6f, 0x63, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x0b, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x49, 0x6d, 0x67, 0x12, 0x1b, 0x2e, 0x6f, 0x70, 0x65, 0x6e,
	0x6f, 0x63, 0x72, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x49, 0x6d, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x6f, 0x63, 0x72,
	0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x49, 0x6d, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x2e, 0x63, 0x6f,
	0x72, 0x70, 0x2e, 0x7a, 0x67, 0x63, 0x73, 0x7a, 0x6b, 0x77, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61,
	0x75, 0x74, 0x68, 0x68, 0x75, 0x62, 0x2f, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x70, 0x62, 0x2f, 0x6f,
	0x70, 0x65, 0x6e, 0x6f, 0x63, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_openocr_openocr_proto_rawDescOnce sync.Once
	file_openocr_openocr_proto_rawDescData = file_openocr_openocr_proto_rawDesc
)

func file_openocr_openocr_proto_rawDescGZIP() []byte {
	file_openocr_openocr_proto_rawDescOnce.Do(func() {
		file_openocr_openocr_proto_rawDescData = protoimpl.X.CompressGZIP(file_openocr_openocr_proto_rawDescData)
	})
	return file_openocr_openocr_proto_rawDescData
}

var file_openocr_openocr_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_openocr_openocr_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_openocr_openocr_proto_goTypes = []interface{}{
	(OcrInvoiceType)(0),              // 0: openocr.OcrInvoiceType
	(OcrLang)(0),                     // 1: openocr.OcrLang
	(*TextBox)(nil),                  // 2: openocr.TextBox
	(*LineInfo)(nil),                 // 3: openocr.LineInfo
	(*RecognizeTextAutoRequest)(nil), // 4: openocr.RecognizeTextAutoRequest
	(*Text)(nil),                     // 5: openocr.Text
	(*RecognizeTextResponse)(nil),    // 6: openocr.RecognizeTextResponse
	(*OpenocrRequest)(nil),           // 7: openocr.OpenocrRequest
	(*OpenocrResponse)(nil),          // 8: openocr.OpenocrResponse
	(*GenerateImgRequest)(nil),       // 9: openocr.GenerateImgRequest
	(*GenerateImgResponse)(nil),      // 10: openocr.GenerateImgResponse
}
var file_openocr_openocr_proto_depIdxs = []int32{
	2,  // 0: openocr.LineInfo.text_boxs:type_name -> openocr.TextBox
	5,  // 1: openocr.RecognizeTextResponse.json_text:type_name -> openocr.Text
	3,  // 2: openocr.RecognizeTextResponse.line_infos:type_name -> openocr.LineInfo
	2,  // 3: openocr.OpenocrResponse.boxes:type_name -> openocr.TextBox
	2,  // 4: openocr.GenerateImgRequest.boxes:type_name -> openocr.TextBox
	4,  // 5: openocr.Openocr.RecognizeTextAuto:input_type -> openocr.RecognizeTextAutoRequest
	7,  // 6: openocr.Openocr.Openocr:input_type -> openocr.OpenocrRequest
	9,  // 7: openocr.Openocr.GenerateImg:input_type -> openocr.GenerateImgRequest
	6,  // 8: openocr.Openocr.RecognizeTextAuto:output_type -> openocr.RecognizeTextResponse
	8,  // 9: openocr.Openocr.Openocr:output_type -> openocr.OpenocrResponse
	10, // 10: openocr.Openocr.GenerateImg:output_type -> openocr.GenerateImgResponse
	8,  // [8:11] is the sub-list for method output_type
	5,  // [5:8] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_openocr_openocr_proto_init() }
func file_openocr_openocr_proto_init() {
	if File_openocr_openocr_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_openocr_openocr_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TextBox); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_openocr_openocr_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LineInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_openocr_openocr_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecognizeTextAutoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_openocr_openocr_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Text); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_openocr_openocr_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecognizeTextResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_openocr_openocr_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenocrRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_openocr_openocr_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenocrResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_openocr_openocr_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateImgRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_openocr_openocr_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateImgResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_openocr_openocr_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_openocr_openocr_proto_goTypes,
		DependencyIndexes: file_openocr_openocr_proto_depIdxs,
		EnumInfos:         file_openocr_openocr_proto_enumTypes,
		MessageInfos:      file_openocr_openocr_proto_msgTypes,
	}.Build()
	File_openocr_openocr_proto = out.File
	file_openocr_openocr_proto_rawDesc = nil
	file_openocr_openocr_proto_goTypes = nil
	file_openocr_openocr_proto_depIdxs = nil
}
