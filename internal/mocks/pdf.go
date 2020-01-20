// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import gofpdf "github.com/jung-kurt/gofpdf"
import io "io"
import mock "github.com/stretchr/testify/mock"
import time "time"

// Pdf is an autogenerated mock type for the Pdf type
type Pdf struct {
	mock.Mock
}

// SetUnderlineThickness provides a mock with given fields: thickness
func (_m *Pdf) SetUnderlineThickness(thickness float64) {
	_m.Called(thickness)
}

// AddFont provides a mock function with given fields: familyStr, styleStr, fileStr
func (_m *Pdf) AddFont(familyStr string, styleStr string, fileStr string) {
	_m.Called(familyStr, styleStr, fileStr)
}

// AddFontFromBytes provides a mock function with given fields: familyStr, styleStr, jsonFileBytes, zFileBytes
func (_m *Pdf) AddFontFromBytes(familyStr string, styleStr string, jsonFileBytes []byte, zFileBytes []byte) {
	_m.Called(familyStr, styleStr, jsonFileBytes, zFileBytes)
}

// AddFontFromReader provides a mock function with given fields: familyStr, styleStr, r
func (_m *Pdf) AddFontFromReader(familyStr string, styleStr string, r io.Reader) {
	_m.Called(familyStr, styleStr, r)
}

// AddLayer provides a mock function with given fields: name, visible
func (_m *Pdf) AddLayer(name string, visible bool) int {
	ret := _m.Called(name, visible)

	var r0 int
	if rf, ok := ret.Get(0).(func(string, bool) int); ok {
		r0 = rf(name, visible)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// AddLink provides a mock function with given fields:
func (_m *Pdf) AddLink() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// AddPage provides a mock function with given fields:
func (_m *Pdf) AddPage() {
	_m.Called()
}

// AddPageFormat provides a mock function with given fields: orientationStr, size
func (_m *Pdf) AddPageFormat(orientationStr string, size gofpdf.SizeType) {
	_m.Called(orientationStr, size)
}

// AddSpotColor provides a mock function with given fields: nameStr, c, m, y, k
func (_m *Pdf) AddSpotColor(nameStr string, c byte, m byte, y byte, k byte) {
	_m.Called(nameStr, c, m, y, k)
}

// AliasNbPages provides a mock function with given fields: aliasStr
func (_m *Pdf) AliasNbPages(aliasStr string) {
	_m.Called(aliasStr)
}

// Arc provides a mock function with given fields: x, y, rx, ry, degRotate, degStart, degEnd, styleStr
func (_m *Pdf) Arc(x float64, y float64, rx float64, ry float64, degRotate float64, degStart float64, degEnd float64, styleStr string) {
	_m.Called(x, y, rx, ry, degRotate, degStart, degEnd, styleStr)
}

// ArcTo provides a mock function with given fields: x, y, rx, ry, degRotate, degStart, degEnd
func (_m *Pdf) ArcTo(x float64, y float64, rx float64, ry float64, degRotate float64, degStart float64, degEnd float64) {
	_m.Called(x, y, rx, ry, degRotate, degStart, degEnd)
}

// BeginLayer provides a mock function with given fields: id
func (_m *Pdf) BeginLayer(id int) {
	_m.Called(id)
}

// Beziergon provides a mock function with given fields: points, styleStr
func (_m *Pdf) Beziergon(points []gofpdf.PointType, styleStr string) {
	_m.Called(points, styleStr)
}

// Bookmark provides a mock function with given fields: txtStr, level, y
func (_m *Pdf) Bookmark(txtStr string, level int, y float64) {
	_m.Called(txtStr, level, y)
}

// Cell provides a mock function with given fields: w, h, txtStr
func (_m *Pdf) Cell(w float64, h float64, txtStr string) {
	_m.Called(w, h, txtStr)
}

// CellFormat provides a mock function with given fields: w, h, txtStr, borderStr, ln, alignStr, fill, link, linkStr
func (_m *Pdf) CellFormat(w float64, h float64, txtStr string, borderStr string, ln int, alignStr string, fill bool, link int, linkStr string) {
	_m.Called(int(w), int(h), txtStr, borderStr, ln, alignStr, fill, link, linkStr)
}

// Cellf provides a mock function with given fields: w, h, fmtStr, args
func (_m *Pdf) Cellf(w float64, h float64, fmtStr string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, w, h, fmtStr)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Circle provides a mock function with given fields: x, y, r, styleStr
func (_m *Pdf) Circle(x float64, y float64, r float64, styleStr string) {
	_m.Called(x, y, r, styleStr)
}

// ClearError provides a mock function with given fields:
func (_m *Pdf) ClearError() {
	_m.Called()
}

// ClipCircle provides a mock function with given fields: x, y, r, outline
func (_m *Pdf) ClipCircle(x float64, y float64, r float64, outline bool) {
	_m.Called(x, y, r, outline)
}

// ClipEllipse provides a mock function with given fields: x, y, rx, ry, outline
func (_m *Pdf) ClipEllipse(x float64, y float64, rx float64, ry float64, outline bool) {
	_m.Called(x, y, rx, ry, outline)
}

// ClipEnd provides a mock function with given fields:
func (_m *Pdf) ClipEnd() {
	_m.Called()
}

// ClipPolygon provides a mock function with given fields: points, outline
func (_m *Pdf) ClipPolygon(points []gofpdf.PointType, outline bool) {
	_m.Called(points, outline)
}

// ClipRect provides a mock function with given fields: x, y, w, h, outline
func (_m *Pdf) ClipRect(x float64, y float64, w float64, h float64, outline bool) {
	_m.Called(x, y, w, h, outline)
}

// ClipRoundedRect provides a mock function with given fields: x, y, w, h, r, outline
func (_m *Pdf) ClipRoundedRect(x float64, y float64, w float64, h float64, r float64, outline bool) {
	_m.Called(x, y, w, h, r, outline)
}

// ClipText provides a mock function with given fields: x, y, txtStr, outline
func (_m *Pdf) ClipText(x float64, y float64, txtStr string, outline bool) {
	_m.Called(x, y, txtStr, outline)
}

// Close provides a mock function with given fields:
func (_m *Pdf) Close() {
	_m.Called()
}

// ClosePath provides a mock function with given fields:
func (_m *Pdf) ClosePath() {
	_m.Called()
}

// CreateTemplate provides a mock function with given fields: fn
func (_m *Pdf) CreateTemplate(fn func(*gofpdf.Tpl)) gofpdf.Template {
	ret := _m.Called(fn)

	var r0 gofpdf.Template
	if rf, ok := ret.Get(0).(func(func(*gofpdf.Tpl)) gofpdf.Template); ok {
		r0 = rf(fn)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(gofpdf.Template)
		}
	}

	return r0
}

// CreateTemplateCustom provides a mock function with given fields: corner, size, fn
func (_m *Pdf) CreateTemplateCustom(corner gofpdf.PointType, size gofpdf.SizeType, fn func(*gofpdf.Tpl)) gofpdf.Template {
	ret := _m.Called(corner, size, fn)

	var r0 gofpdf.Template
	if rf, ok := ret.Get(0).(func(gofpdf.PointType, gofpdf.SizeType, func(*gofpdf.Tpl)) gofpdf.Template); ok {
		r0 = rf(corner, size, fn)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(gofpdf.Template)
		}
	}

	return r0
}

// Curve provides a mock function with given fields: x0, y0, cx, cy, x1, y1, styleStr
func (_m *Pdf) Curve(x0 float64, y0 float64, cx float64, cy float64, x1 float64, y1 float64, styleStr string) {
	_m.Called(x0, y0, cx, cy, x1, y1, styleStr)
}

// CurveBezierCubic provides a mock function with given fields: x0, y0, cx0, cy0, cx1, cy1, x1, y1, styleStr
func (_m *Pdf) CurveBezierCubic(x0 float64, y0 float64, cx0 float64, cy0 float64, cx1 float64, cy1 float64, x1 float64, y1 float64, styleStr string) {
	_m.Called(x0, y0, cx0, cy0, cx1, cy1, x1, y1, styleStr)
}

// CurveBezierCubicTo provides a mock function with given fields: cx0, cy0, cx1, cy1, x, y
func (_m *Pdf) CurveBezierCubicTo(cx0 float64, cy0 float64, cx1 float64, cy1 float64, x float64, y float64) {
	_m.Called(cx0, cy0, cx1, cy1, x, y)
}

// CurveCubic provides a mock function with given fields: x0, y0, cx0, cy0, x1, y1, cx1, cy1, styleStr
func (_m *Pdf) CurveCubic(x0 float64, y0 float64, cx0 float64, cy0 float64, x1 float64, y1 float64, cx1 float64, cy1 float64, styleStr string) {
	_m.Called(x0, y0, cx0, cy0, x1, y1, cx1, cy1, styleStr)
}

// CurveTo provides a mock function with given fields: cx, cy, x, y
func (_m *Pdf) CurveTo(cx float64, cy float64, x float64, y float64) {
	_m.Called(cx, cy, x, y)
}

// DrawPath provides a mock function with given fields: styleStr
func (_m *Pdf) DrawPath(styleStr string) {
	_m.Called(styleStr)
}

// Ellipse provides a mock function with given fields: x, y, rx, ry, degRotate, styleStr
func (_m *Pdf) Ellipse(x float64, y float64, rx float64, ry float64, degRotate float64, styleStr string) {
	_m.Called(x, y, rx, ry, degRotate, styleStr)
}

// EndLayer provides a mock function with given fields:
func (_m *Pdf) EndLayer() {
	_m.Called()
}

// Err provides a mock function with given fields:
func (_m *Pdf) Err() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Error provides a mock function with given fields:
func (_m *Pdf) Error() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAlpha provides a mock function with given fields:
func (_m *Pdf) GetAlpha() (float64, string) {
	ret := _m.Called()

	var r0 float64
	if rf, ok := ret.Get(0).(func() float64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(float64)
	}

	var r1 string
	if rf, ok := ret.Get(1).(func() string); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(string)
	}

	return r0, r1
}

// GetAutoPageBreak provides a mock function with given fields:
func (_m *Pdf) GetAutoPageBreak() (bool, float64) {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 float64
	if rf, ok := ret.Get(1).(func() float64); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(float64)
	}

	return r0, r1
}

// GetCellMargin provides a mock function with given fields:
func (_m *Pdf) GetCellMargin() float64 {
	ret := _m.Called()

	var r0 float64
	if rf, ok := ret.Get(0).(func() float64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(float64)
	}

	return r0
}

// GetConversionRatio provides a mock function with given fields:
func (_m *Pdf) GetConversionRatio() float64 {
	ret := _m.Called()

	var r0 float64
	if rf, ok := ret.Get(0).(func() float64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(float64)
	}

	return r0
}

// GetDrawColor provides a mock function with given fields:
func (_m *Pdf) GetDrawColor() (int, int, int) {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 int
	if rf, ok := ret.Get(1).(func() int); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 int
	if rf, ok := ret.Get(2).(func() int); ok {
		r2 = rf()
	} else {
		r2 = ret.Get(2).(int)
	}

	return r0, r1, r2
}

// GetDrawSpotColor provides a mock function with given fields:
func (_m *Pdf) GetDrawSpotColor() (string, byte, byte, byte, byte) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 byte
	if rf, ok := ret.Get(1).(func() byte); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(byte)
	}

	var r2 byte
	if rf, ok := ret.Get(2).(func() byte); ok {
		r2 = rf()
	} else {
		r2 = ret.Get(2).(byte)
	}

	var r3 byte
	if rf, ok := ret.Get(3).(func() byte); ok {
		r3 = rf()
	} else {
		r3 = ret.Get(3).(byte)
	}

	var r4 byte
	if rf, ok := ret.Get(4).(func() byte); ok {
		r4 = rf()
	} else {
		r4 = ret.Get(4).(byte)
	}

	return r0, r1, r2, r3, r4
}

// GetFillColor provides a mock function with given fields:
func (_m *Pdf) GetFillColor() (int, int, int) {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 int
	if rf, ok := ret.Get(1).(func() int); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 int
	if rf, ok := ret.Get(2).(func() int); ok {
		r2 = rf()
	} else {
		r2 = ret.Get(2).(int)
	}

	return r0, r1, r2
}

// GetFillSpotColor provides a mock function with given fields:
func (_m *Pdf) GetFillSpotColor() (string, byte, byte, byte, byte) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 byte
	if rf, ok := ret.Get(1).(func() byte); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(byte)
	}

	var r2 byte
	if rf, ok := ret.Get(2).(func() byte); ok {
		r2 = rf()
	} else {
		r2 = ret.Get(2).(byte)
	}

	var r3 byte
	if rf, ok := ret.Get(3).(func() byte); ok {
		r3 = rf()
	} else {
		r3 = ret.Get(3).(byte)
	}

	var r4 byte
	if rf, ok := ret.Get(4).(func() byte); ok {
		r4 = rf()
	} else {
		r4 = ret.Get(4).(byte)
	}

	return r0, r1, r2, r3, r4
}

// GetFontDesc provides a mock function with given fields: familyStr, styleStr
func (_m *Pdf) GetFontDesc(familyStr string, styleStr string) gofpdf.FontDescType {
	ret := _m.Called(familyStr, styleStr)

	var r0 gofpdf.FontDescType
	if rf, ok := ret.Get(0).(func(string, string) gofpdf.FontDescType); ok {
		r0 = rf(familyStr, styleStr)
	} else {
		r0 = ret.Get(0).(gofpdf.FontDescType)
	}

	return r0
}

// GetFontSize provides a mock function with given fields:
func (_m *Pdf) GetFontSize() (float64, float64) {
	ret := _m.Called()

	var r0 float64
	if rf, ok := ret.Get(0).(func() float64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(float64)
	}

	var r1 float64
	if rf, ok := ret.Get(1).(func() float64); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(float64)
	}

	return r0, r1
}

// GetImageInfo provides a mock function with given fields: imageStr
func (_m *Pdf) GetImageInfo(imageStr string) *gofpdf.ImageInfoType {
	ret := _m.Called(imageStr)

	var r0 *gofpdf.ImageInfoType
	if rf, ok := ret.Get(0).(func(string) *gofpdf.ImageInfoType); ok {
		r0 = rf(imageStr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gofpdf.ImageInfoType)
		}
	}

	return r0
}

// GetLineWidth provides a mock function with given fields:
func (_m *Pdf) GetLineWidth() float64 {
	ret := _m.Called()

	var r0 float64
	if rf, ok := ret.Get(0).(func() float64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(float64)
	}

	return r0
}

// GetMargins provides a mock function with given fields:
func (_m *Pdf) GetMargins() (float64, float64, float64, float64) {
	ret := _m.Called()

	var r0 float64
	if rf, ok := ret.Get(0).(func() float64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(float64)
	}

	var r1 float64
	if rf, ok := ret.Get(1).(func() float64); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(float64)
	}

	var r2 float64
	if rf, ok := ret.Get(2).(func() float64); ok {
		r2 = rf()
	} else {
		r2 = ret.Get(2).(float64)
	}

	var r3 float64
	if rf, ok := ret.Get(3).(func() float64); ok {
		r3 = rf()
	} else {
		r3 = ret.Get(3).(float64)
	}

	return r0, r1, r2, r3
}

// GetPageSize provides a mock function with given fields:
func (_m *Pdf) GetPageSize() (float64, float64) {
	ret := _m.Called()

	var r0 float64
	if rf, ok := ret.Get(0).(func() float64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(float64)
	}

	var r1 float64
	if rf, ok := ret.Get(1).(func() float64); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(float64)
	}

	return r0, r1
}

// GetPageSizeStr provides a mock function with given fields: sizeStr
func (_m *Pdf) GetPageSizeStr(sizeStr string) gofpdf.SizeType {
	ret := _m.Called(sizeStr)

	var r0 gofpdf.SizeType
	if rf, ok := ret.Get(0).(func(string) gofpdf.SizeType); ok {
		r0 = rf(sizeStr)
	} else {
		r0 = ret.Get(0).(gofpdf.SizeType)
	}

	return r0
}

// GetStringWidth provides a mock function with given fields: s
func (_m *Pdf) GetStringWidth(s string) float64 {
	ret := _m.Called(s)

	var r0 float64
	if rf, ok := ret.Get(0).(func(string) float64); ok {
		r0 = rf(s)
	} else {
		r0 = ret.Get(0).(float64)
	}

	return r0
}

// GetTextColor provides a mock function with given fields:
func (_m *Pdf) GetTextColor() (int, int, int) {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 int
	if rf, ok := ret.Get(1).(func() int); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 int
	if rf, ok := ret.Get(2).(func() int); ok {
		r2 = rf()
	} else {
		r2 = ret.Get(2).(int)
	}

	return r0, r1, r2
}

// GetTextSpotColor provides a mock function with given fields:
func (_m *Pdf) GetTextSpotColor() (string, byte, byte, byte, byte) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 byte
	if rf, ok := ret.Get(1).(func() byte); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(byte)
	}

	var r2 byte
	if rf, ok := ret.Get(2).(func() byte); ok {
		r2 = rf()
	} else {
		r2 = ret.Get(2).(byte)
	}

	var r3 byte
	if rf, ok := ret.Get(3).(func() byte); ok {
		r3 = rf()
	} else {
		r3 = ret.Get(3).(byte)
	}

	var r4 byte
	if rf, ok := ret.Get(4).(func() byte); ok {
		r4 = rf()
	} else {
		r4 = ret.Get(4).(byte)
	}

	return r0, r1, r2, r3, r4
}

// GetX provides a mock function with given fields:
func (_m *Pdf) GetX() float64 {
	ret := _m.Called()

	var r0 float64
	if rf, ok := ret.Get(0).(func() float64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(float64)
	}

	return r0
}

// GetXY provides a mock function with given fields:
func (_m *Pdf) GetXY() (float64, float64) {
	ret := _m.Called()

	var r0 float64
	if rf, ok := ret.Get(0).(func() float64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(float64)
	}

	var r1 float64
	if rf, ok := ret.Get(1).(func() float64); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(float64)
	}

	return r0, r1
}

// GetY provides a mock function with given fields:
func (_m *Pdf) GetY() float64 {
	ret := _m.Called()

	var r0 float64
	if rf, ok := ret.Get(0).(func() float64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(float64)
	}

	return r0
}

// HTMLBasicNew provides a mock function with given fields:
func (_m *Pdf) HTMLBasicNew() gofpdf.HTMLBasicType {
	ret := _m.Called()

	var r0 gofpdf.HTMLBasicType
	if rf, ok := ret.Get(0).(func() gofpdf.HTMLBasicType); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(gofpdf.HTMLBasicType)
	}

	return r0
}

// Image provides a mock function with given fields: imageNameStr, x, y, w, h, flow, tp, link, linkStr
func (_m *Pdf) Image(imageNameStr string, x float64, y float64, w float64, h float64, flow bool, tp string, link int, linkStr string) {
	_m.Called("", int(x), int(y), int(w), int(h))
}

// ImageOptions provides a mock function with given fields: imageNameStr, x, y, w, h, flow, options, link, linkStr
func (_m *Pdf) ImageOptions(imageNameStr string, x float64, y float64, w float64, h float64, flow bool, options gofpdf.ImageOptions, link int, linkStr string) {
	_m.Called(imageNameStr, int(x), int(y), int(w), int(h))
}

// ImageTypeFromMime provides a mock function with given fields: mimeStr
func (_m *Pdf) ImageTypeFromMime(mimeStr string) string {
	ret := _m.Called(mimeStr)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(mimeStr)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Line provides a mock function with given fields: x1, y1, x2, y2
func (_m *Pdf) Line(x1 float64, y1 float64, x2 float64, y2 float64) {
	_m.Called(x1, y1, x2, y2)
}

// LineTo provides a mock function with given fields: x, y
func (_m *Pdf) LineTo(x float64, y float64) {
	_m.Called(x, y)
}

// LinearGradient provides a mock function with given fields: x, y, w, h, r1, g1, b1, r2, g2, b2, x1, y1, x2, y2
func (_m *Pdf) LinearGradient(x float64, y float64, w float64, h float64, r1 int, g1 int, b1 int, r2 int, g2 int, b2 int, x1 float64, y1 float64, x2 float64, y2 float64) {
	_m.Called(x, y, w, h, r1, g1, b1, r2, g2, b2, x1, y1, x2, y2)
}

// Link provides a mock function with given fields: x, y, w, h, link
func (_m *Pdf) Link(x float64, y float64, w float64, h float64, link int) {
	_m.Called(x, y, w, h, link)
}

// LinkString provides a mock function with given fields: x, y, w, h, linkStr
func (_m *Pdf) LinkString(x float64, y float64, w float64, h float64, linkStr string) {
	_m.Called(x, y, w, h, linkStr)
}

// Ln provides a mock function with given fields: h
func (_m *Pdf) Ln(h float64) {
	_m.Called(h)
}

// MoveTo provides a mock function with given fields: x, y
func (_m *Pdf) MoveTo(x float64, y float64) {
	_m.Called(x, y)
}

// MultiCell provides a mock function with given fields: w, h, txtStr, borderStr, alignStr, fill
func (_m *Pdf) MultiCell(w float64, h float64, txtStr string, borderStr string, alignStr string, fill bool) {
	_m.Called(w, h, txtStr, borderStr, alignStr, fill)
}

// Ok provides a mock function with given fields:
func (_m *Pdf) Ok() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// OpenLayerPane provides a mock function with given fields:
func (_m *Pdf) OpenLayerPane() {
	_m.Called()
}

// Output provides a mock function with given fields: w
func (_m *Pdf) Output(w io.Writer) error {
	ret := _m.Called(w)

	var r0 error
	if rf, ok := ret.Get(0).(func(io.Writer) error); ok {
		r0 = rf(w)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OutputAndClose provides a mock function with given fields: w
func (_m *Pdf) OutputAndClose(w io.WriteCloser) error {
	ret := _m.Called(w)

	var r0 error
	if rf, ok := ret.Get(0).(func(io.WriteCloser) error); ok {
		r0 = rf(w)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OutputFileAndClose provides a mock function with given fields: fileStr
func (_m *Pdf) OutputFileAndClose(fileStr string) error {
	ret := _m.Called(fileStr)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(fileStr)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PageCount provides a mock function with given fields:
func (_m *Pdf) PageCount() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// PageNo provides a mock function with given fields:
func (_m *Pdf) PageNo() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// PageSize provides a mock function with given fields: pageNum
func (_m *Pdf) PageSize(pageNum int) (float64, float64, string) {
	ret := _m.Called(pageNum)

	var r0 float64
	if rf, ok := ret.Get(0).(func(int) float64); ok {
		r0 = rf(pageNum)
	} else {
		r0 = ret.Get(0).(float64)
	}

	var r1 float64
	if rf, ok := ret.Get(1).(func(int) float64); ok {
		r1 = rf(pageNum)
	} else {
		r1 = ret.Get(1).(float64)
	}

	var r2 string
	if rf, ok := ret.Get(2).(func(int) string); ok {
		r2 = rf(pageNum)
	} else {
		r2 = ret.Get(2).(string)
	}

	return r0, r1, r2
}

// PointConvert provides a mock function with given fields: pt
func (_m *Pdf) PointConvert(pt float64) float64 {
	ret := _m.Called(pt)

	var r0 float64
	if rf, ok := ret.Get(0).(func(float64) float64); ok {
		r0 = rf(pt)
	} else {
		r0 = ret.Get(0).(float64)
	}

	return r0
}

// PointToUnitConvert provides a mock function with given fields: pt
func (_m *Pdf) PointToUnitConvert(pt float64) float64 {
	ret := _m.Called(pt)

	var r0 float64
	if rf, ok := ret.Get(0).(func(float64) float64); ok {
		r0 = rf(pt)
	} else {
		r0 = ret.Get(0).(float64)
	}

	return r0
}

// Polygon provides a mock function with given fields: points, styleStr
func (_m *Pdf) Polygon(points []gofpdf.PointType, styleStr string) {
	_m.Called(points, styleStr)
}

// RadialGradient provides a mock function with given fields: x, y, w, h, r1, g1, b1, r2, g2, b2, x1, y1, x2, y2, r
func (_m *Pdf) RadialGradient(x float64, y float64, w float64, h float64, r1 int, g1 int, b1 int, r2 int, g2 int, b2 int, x1 float64, y1 float64, x2 float64, y2 float64, r float64) {
	_m.Called(x, y, w, h, r1, g1, b1, r2, g2, b2, x1, y1, x2, y2, r)
}

// RawWriteBuf provides a mock function with given fields: r
func (_m *Pdf) RawWriteBuf(r io.Reader) {
	_m.Called(r)
}

// RawWriteStr provides a mock function with given fields: str
func (_m *Pdf) RawWriteStr(str string) {
	_m.Called(str)
}

// Rect provides a mock function with given fields: x, y, w, h, styleStr
func (_m *Pdf) Rect(x float64, y float64, w float64, h float64, styleStr string) {
	_m.Called(x, y, w, h, styleStr)
}

// RegisterAlias provides a mock function with given fields: alias, replacement
func (_m *Pdf) RegisterAlias(alias string, replacement string) {
	_m.Called(alias, replacement)
}

// RegisterImage provides a mock function with given fields: fileStr, tp
func (_m *Pdf) RegisterImage(fileStr string, tp string) *gofpdf.ImageInfoType {
	ret := _m.Called(fileStr, tp)

	var r0 *gofpdf.ImageInfoType
	if rf, ok := ret.Get(0).(func(string, string) *gofpdf.ImageInfoType); ok {
		r0 = rf(fileStr, tp)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gofpdf.ImageInfoType)
		}
	}

	return r0
}

// RegisterImageOptions provides a mock function with given fields: fileStr, options
func (_m *Pdf) RegisterImageOptions(fileStr string, options gofpdf.ImageOptions) *gofpdf.ImageInfoType {
	ret := _m.Called(fileStr, options)

	var r0 *gofpdf.ImageInfoType
	if rf, ok := ret.Get(0).(func(string, gofpdf.ImageOptions) *gofpdf.ImageInfoType); ok {
		r0 = rf(fileStr, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gofpdf.ImageInfoType)
		}
	}

	return r0
}

// RegisterImageOptionsReader provides a mock function with given fields: imgName, options, r
func (_m *Pdf) RegisterImageOptionsReader(imgName string, options gofpdf.ImageOptions, r io.Reader) *gofpdf.ImageInfoType {
	ret := _m.Called("", options, "")

	var r0 *gofpdf.ImageInfoType
	if rf, ok := ret.Get(0).(func(string, gofpdf.ImageOptions, io.Reader) *gofpdf.ImageInfoType); ok {
		r0 = rf(imgName, options, r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gofpdf.ImageInfoType)
		}
	}

	return r0
}

// RegisterImageReader provides a mock function with given fields: imgName, tp, r
func (_m *Pdf) RegisterImageReader(imgName string, tp string, r io.Reader) *gofpdf.ImageInfoType {
	ret := _m.Called(imgName, tp)

	var r0 *gofpdf.ImageInfoType
	if rf, ok := ret.Get(0).(func(string, string, io.Reader) *gofpdf.ImageInfoType); ok {
		r0 = rf(imgName, tp, r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gofpdf.ImageInfoType)
		}
	}

	return r0
}

// SVGBasicWrite provides a mock function with given fields: sb, scale
func (_m *Pdf) SVGBasicWrite(sb *gofpdf.SVGBasicType, scale float64) {
	_m.Called(sb, scale)
}

// SetAcceptPageBreakFunc provides a mock function with given fields: fnc
func (_m *Pdf) SetAcceptPageBreakFunc(fnc func() bool) {
	_m.Called(fnc)
}

// SetAlpha provides a mock function with given fields: alpha, blendModeStr
func (_m *Pdf) SetAlpha(alpha float64, blendModeStr string) {
	_m.Called(alpha, blendModeStr)
}

// SetAuthor provides a mock function with given fields: authorStr, isUTF8
func (_m *Pdf) SetAuthor(authorStr string, isUTF8 bool) {
	_m.Called(authorStr, isUTF8)
}

// SetAutoPageBreak provides a mock function with given fields: auto, margin
func (_m *Pdf) SetAutoPageBreak(auto bool, margin float64) {
	_m.Called(auto, margin)
}

// SetCatalogSort provides a mock function with given fields: flag
func (_m *Pdf) SetCatalogSort(flag bool) {
	_m.Called(flag)
}

// SetCellMargin provides a mock function with given fields: margin
func (_m *Pdf) SetCellMargin(margin float64) {
	_m.Called(margin)
}

// SetCompression provides a mock function with given fields: compress
func (_m *Pdf) SetCompression(compress bool) {
	_m.Called(compress)
}

// SetCreationDate provides a mock function with given fields: tm
func (_m *Pdf) SetCreationDate(tm time.Time) {
	_m.Called(tm)
}

// SetCreator provides a mock function with given fields: creatorStr, isUTF8
func (_m *Pdf) SetCreator(creatorStr string, isUTF8 bool) {
	_m.Called(creatorStr, isUTF8)
}

// SetDashPattern provides a mock function with given fields: dashArray, dashPhase
func (_m *Pdf) SetDashPattern(dashArray []float64, dashPhase float64) {
	_m.Called(dashArray, dashPhase)
}

// SetDisplayMode provides a mock function with given fields: zoomStr, layoutStr
func (_m *Pdf) SetDisplayMode(zoomStr string, layoutStr string) {
	_m.Called(zoomStr, layoutStr)
}

// SetDrawColor provides a mock function with given fields: r, g, b
func (_m *Pdf) SetDrawColor(r int, g int, b int) {
	_m.Called(r, g, b)
}

// SetDrawSpotColor provides a mock function with given fields: nameStr, tint
func (_m *Pdf) SetDrawSpotColor(nameStr string, tint byte) {
	_m.Called(nameStr, tint)
}

// SetError provides a mock function with given fields: err
func (_m *Pdf) SetError(err error) {
	_m.Called(err)
}

// SetErrorf provides a mock function with given fields: fmtStr, args
func (_m *Pdf) SetErrorf(fmtStr string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, fmtStr)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// SetFillColor provides a mock function with given fields: r, g, b
func (_m *Pdf) SetFillColor(r int, g int, b int) {
	_m.Called(r, g, b)
}

// SetFillSpotColor provides a mock function with given fields: nameStr, tint
func (_m *Pdf) SetFillSpotColor(nameStr string, tint byte) {
	_m.Called(nameStr, tint)
}

// SetFont provides a mock function with given fields: familyStr, styleStr, size
func (_m *Pdf) SetFont(familyStr string, styleStr string, size float64) {
	_m.Called(familyStr, styleStr, size)
}

// SetFontLoader provides a mock function with given fields: loader
func (_m *Pdf) SetFontLoader(loader gofpdf.FontLoader) {
	_m.Called(loader)
}

// SetFontLocation provides a mock function with given fields: fontDirStr
func (_m *Pdf) SetFontLocation(fontDirStr string) {
	_m.Called(fontDirStr)
}

// SetFontSize provides a mock function with given fields: size
func (_m *Pdf) SetFontSize(size float64) {
	_m.Called(size)
}

// SetFontStyle provides a mock function with given fields: styleStr
func (_m *Pdf) SetFontStyle(styleStr string) {
	_m.Called(styleStr)
}

// SetFontUnitSize provides a mock function with given fields: size
func (_m *Pdf) SetFontUnitSize(size float64) {
	_m.Called(size)
}

// SetFooterFunc provides a mock function with given fields: fnc
func (_m *Pdf) SetFooterFunc(fnc func()) {
	_m.Called(fnc)
}

// SetFooterFuncLpi provides a mock function with given fields: fnc
func (_m *Pdf) SetFooterFuncLpi(fnc func(bool)) {
	_m.Called(fnc)
}

// SetHeaderFunc provides a mock function with given fields: fnc
func (_m *Pdf) SetHeaderFunc(fnc func()) {
	_m.Called(fnc)
}

// SetHeaderFuncMode provides a mock function with given fields: fnc, homeMode
func (_m *Pdf) SetHeaderFuncMode(fnc func(), homeMode bool) {
	_m.Called(fnc, homeMode)
}

// SetHomeXY provides a mock function with given fields:
func (_m *Pdf) SetHomeXY() {
	_m.Called()
}

// SetJavascript provides a mock function with given fields: script
func (_m *Pdf) SetJavascript(script string) {
	_m.Called(script)
}

// SetKeywords provides a mock function with given fields: keywordsStr, isUTF8
func (_m *Pdf) SetKeywords(keywordsStr string, isUTF8 bool) {
	_m.Called(keywordsStr, isUTF8)
}

// SetLeftMargin provides a mock function with given fields: margin
func (_m *Pdf) SetLeftMargin(margin float64) {
	_m.Called(margin)
}

// SetLineCapStyle provides a mock function with given fields: styleStr
func (_m *Pdf) SetLineCapStyle(styleStr string) {
	_m.Called(styleStr)
}

// SetLineJoinStyle provides a mock function with given fields: styleStr
func (_m *Pdf) SetLineJoinStyle(styleStr string) {
	_m.Called(styleStr)
}

// SetLineWidth provides a mock function with given fields: width
func (_m *Pdf) SetLineWidth(width float64) {
	_m.Called(width)
}

// SetLink provides a mock function with given fields: link, y, page
func (_m *Pdf) SetLink(link int, y float64, page int) {
	_m.Called(link, y, page)
}

// SetMargins provides a mock function with given fields: left, top, right
func (_m *Pdf) SetMargins(left float64, top float64, right float64) {
	_m.Called(left, top, right)
}

// SetPage provides a mock function with given fields: pageNum
func (_m *Pdf) SetPage(pageNum int) {
	_m.Called(pageNum)
}

// SetPageBox provides a mock function with given fields: t, x, y, wd, ht
func (_m *Pdf) SetPageBox(t string, x float64, y float64, wd float64, ht float64) {
	_m.Called(t, x, y, wd, ht)
}

// SetPageBoxRec provides a mock function with given fields: t, pb
func (_m *Pdf) SetPageBoxRec(t string, pb gofpdf.PageBox) {
	_m.Called(t, pb)
}

// SetProtection provides a mock function with given fields: actionFlag, userPassStr, ownerPassStr
func (_m *Pdf) SetProtection(actionFlag byte, userPassStr string, ownerPassStr string) {
	_m.Called(actionFlag, userPassStr, ownerPassStr)
}

// SetRightMargin provides a mock function with given fields: margin
func (_m *Pdf) SetRightMargin(margin float64) {
	_m.Called(margin)
}

// SetSubject provides a mock function with given fields: subjectStr, isUTF8
func (_m *Pdf) SetSubject(subjectStr string, isUTF8 bool) {
	_m.Called(subjectStr, isUTF8)
}

// SetTextColor provides a mock function with given fields: r, g, b
func (_m *Pdf) SetTextColor(r int, g int, b int) {
	_m.Called(r, g, b)
}

// SetTextSpotColor provides a mock function with given fields: nameStr, tint
func (_m *Pdf) SetTextSpotColor(nameStr string, tint byte) {
	_m.Called(nameStr, tint)
}

// SetTitle provides a mock function with given fields: titleStr, isUTF8
func (_m *Pdf) SetTitle(titleStr string, isUTF8 bool) {
	_m.Called(titleStr, isUTF8)
}

// SetTopMargin provides a mock function with given fields: margin
func (_m *Pdf) SetTopMargin(margin float64) {
	_m.Called(margin)
}

// SetX provides a mock function with given fields: x
func (_m *Pdf) SetX(x float64) {
	_m.Called(x)
}

// SetXY provides a mock function with given fields: x, y
func (_m *Pdf) SetXY(x float64, y float64) {
	_m.Called(x, y)
}

// SetXmpMetadata provides a mock function with given fields: xmpStream
func (_m *Pdf) SetXmpMetadata(xmpStream []byte) {
	_m.Called(xmpStream)
}

// SetY provides a mock function with given fields: y
func (_m *Pdf) SetY(y float64) {
	_m.Called(y)
}

// SplitLines provides a mock function with given fields: txt, w
func (_m *Pdf) SplitLines(txt []byte, w float64) [][]byte {
	ret := _m.Called(txt, w)

	var r0 [][]byte
	if rf, ok := ret.Get(0).(func([]byte, float64) [][]byte); ok {
		r0 = rf(txt, w)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([][]byte)
		}
	}

	return r0
}

// String provides a mock function with given fields:
func (_m *Pdf) String() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// TextHelper provides a mock function with given fields: x, y, txtStr
func (_m *Pdf) Text(x float64, y float64, txtStr string) {
	_m.Called(x, y, txtStr)
}

// Transform provides a mock function with given fields: tm
func (_m *Pdf) Transform(tm gofpdf.TransformMatrix) {
	_m.Called(tm)
}

// TransformBegin provides a mock function with given fields:
func (_m *Pdf) TransformBegin() {
	_m.Called()
}

// TransformEnd provides a mock function with given fields:
func (_m *Pdf) TransformEnd() {
	_m.Called()
}

// TransformMirrorHorizontal provides a mock function with given fields: x
func (_m *Pdf) TransformMirrorHorizontal(x float64) {
	_m.Called(x)
}

// TransformMirrorLine provides a mock function with given fields: angle, x, y
func (_m *Pdf) TransformMirrorLine(angle float64, x float64, y float64) {
	_m.Called(angle, x, y)
}

// TransformMirrorPoint provides a mock function with given fields: x, y
func (_m *Pdf) TransformMirrorPoint(x float64, y float64) {
	_m.Called(x, y)
}

// TransformMirrorVertical provides a mock function with given fields: y
func (_m *Pdf) TransformMirrorVertical(y float64) {
	_m.Called(y)
}

// TransformRotate provides a mock function with given fields: angle, x, y
func (_m *Pdf) TransformRotate(angle float64, x float64, y float64) {
	_m.Called(angle, x, y)
}

// TransformScale provides a mock function with given fields: scaleWd, scaleHt, x, y
func (_m *Pdf) TransformScale(scaleWd float64, scaleHt float64, x float64, y float64) {
	_m.Called(scaleWd, scaleHt, x, y)
}

// TransformScaleX provides a mock function with given fields: scaleWd, x, y
func (_m *Pdf) TransformScaleX(scaleWd float64, x float64, y float64) {
	_m.Called(scaleWd, x, y)
}

// TransformScaleXY provides a mock function with given fields: s, x, y
func (_m *Pdf) TransformScaleXY(s float64, x float64, y float64) {
	_m.Called(s, x, y)
}

// TransformScaleY provides a mock function with given fields: scaleHt, x, y
func (_m *Pdf) TransformScaleY(scaleHt float64, x float64, y float64) {
	_m.Called(scaleHt, x, y)
}

// TransformSkew provides a mock function with given fields: angleX, angleY, x, y
func (_m *Pdf) TransformSkew(angleX float64, angleY float64, x float64, y float64) {
	_m.Called(angleX, angleY, x, y)
}

// TransformSkewX provides a mock function with given fields: angleX, x, y
func (_m *Pdf) TransformSkewX(angleX float64, x float64, y float64) {
	_m.Called(angleX, x, y)
}

// TransformSkewY provides a mock function with given fields: angleY, x, y
func (_m *Pdf) TransformSkewY(angleY float64, x float64, y float64) {
	_m.Called(angleY, x, y)
}

// TransformTranslate provides a mock function with given fields: tx, ty
func (_m *Pdf) TransformTranslate(tx float64, ty float64) {
	_m.Called(tx, ty)
}

// TransformTranslateX provides a mock function with given fields: tx
func (_m *Pdf) TransformTranslateX(tx float64) {
	_m.Called(tx)
}

// TransformTranslateY provides a mock function with given fields: ty
func (_m *Pdf) TransformTranslateY(ty float64) {
	_m.Called(ty)
}

// UnicodeTranslatorFromDescriptor provides a mock function with given fields: cpStr
func (_m *Pdf) UnicodeTranslatorFromDescriptor(cpStr string) func(string) string {
	ret := _m.Called(cpStr)

	var r0 func(string) string
	if rf, ok := ret.Get(0).(func(string) func(string) string); ok {
		r0 = rf(cpStr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(func(string) string)
		}
	}

	return r0
}

// UnitToPointConvert provides a mock function with given fields: u
func (_m *Pdf) UnitToPointConvert(u float64) float64 {
	ret := _m.Called(u)

	var r0 float64
	if rf, ok := ret.Get(0).(func(float64) float64); ok {
		r0 = rf(u)
	} else {
		r0 = ret.Get(0).(float64)
	}

	return r0
}

// UseTemplate provides a mock function with given fields: t
func (_m *Pdf) UseTemplate(t gofpdf.Template) {
	_m.Called(t)
}

// UseTemplateScaled provides a mock function with given fields: t, corner, size
func (_m *Pdf) UseTemplateScaled(t gofpdf.Template, corner gofpdf.PointType, size gofpdf.SizeType) {
	_m.Called(t, corner, size)
}

// Write provides a mock function with given fields: h, txtStr
func (_m *Pdf) Write(h float64, txtStr string) {
	_m.Called(h, txtStr)
}

// WriteAligned provides a mock function with given fields: width, lineHeight, textStr, alignStr
func (_m *Pdf) WriteAligned(width float64, lineHeight float64, textStr string, alignStr string) {
	_m.Called(width, lineHeight, textStr, alignStr)
}

// WriteLinkID provides a mock function with given fields: h, displayStr, linkID
func (_m *Pdf) WriteLinkID(h float64, displayStr string, linkID int) {
	_m.Called(h, displayStr, linkID)
}

// WriteLinkString provides a mock function with given fields: h, displayStr, targetStr
func (_m *Pdf) WriteLinkString(h float64, displayStr string, targetStr string) {
	_m.Called(h, displayStr, targetStr)
}

// Writef provides a mock function with given fields: h, fmtStr, args
func (_m *Pdf) Writef(h float64, fmtStr string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, h, fmtStr)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}
