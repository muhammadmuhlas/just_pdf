package pdf

import (
	"bytes"
	"github.com/muhammadmuhlas/just_pdf/pkg/color"

	"github.com/muhammadmuhlas/just_pdf/internal"
	"github.com/muhammadmuhlas/just_pdf/pkg/consts"
	"github.com/muhammadmuhlas/just_pdf/pkg/props"
	"github.com/jung-kurt/gofpdf"
)

// JustPdf is the principal abstraction to create a PDF document.
type JustPdf interface {
	// Grid System
	Row(height float64, closure func())
	Col(closure func())
	ColSpace()
	ColSpaces(qtd int)

	// Registers
	RegisterHeader(closure func())
	RegisterFooter(closure func())

	// Helpers
	SetBorder(on bool)
	SetBackgroundColor(color color.Color)
	SetTextColor(color color.Color)
	GetBorder() bool
	GetPageSize() (float64, float64)
	GetCurrentPage() int
	GetCurrentOffset() float64
	SetPageMargins(left, top, right, bottom float64)
	SetLRMargins(left, right float64)
	GetPageMargins() (float64, float64, float64, float64)

	// Outside Col/Row Components
	TableList(header []string, contents [][]string, prop ...props.TableList)
	Line(spaceHeight float64)
	VLine(spaceWidht, spaceHeight float64, color color.Color)

	// Inside Col/Row Components
	Text(text string, prop ...props.Text)
	FileImage(filePathName string, prop ...props.Rect) (err error)
	Base64Image(base64 string, extension consts.Extension, prop ...props.Rect) (err error)
	Barcode(code string, prop ...props.Barcode) error
	QrCode(code string, prop ...props.Rect)
	Signature(label string, prop ...props.Font)

	// File System
	OutputFileAndClose(filePathName string) error
	Output() (bytes.Buffer, error)
}

// PdfJustPdf is the principal structure which implements JustPdf abstraction
type PdfJustPdf struct {
	Pdf                       gofpdf.Pdf
	Math                      internal.Math
	Font                      internal.Font
	TextHelper                internal.Text
	SignHelper                internal.Signature
	Image                     internal.Image
	Code                      internal.Code
	TableListHelper           internal.TableList
	pageIndex                 int
	offsetY                   float64
	rowHeight                 float64
	rowColCount               float64
	backgroundColor           color.Color
	textColor		          color.Color
	colsClosures              []func()
	headerClosure             func()
	footerClosure             func()
	footerHeight              float64
	headerFooterContextActive bool
	calculationMode           bool
	debugMode                 bool
	orientation               consts.Orientation
	pageSize                  consts.PageSize
}

// NewJustPdf create a JustPdf instance returning a pointer to PdfJustPdf
// Receive an Orientation and a PageSize.
func NewJustPdf(orientation consts.Orientation, pageSize consts.PageSize) JustPdf {
	fpdf := gofpdf.New(string(orientation), "mm", string(pageSize), "")
	fpdf.SetMargins(10, 10, 10)

	math := internal.NewMath(fpdf)
	font := internal.NewFont(fpdf, 16, consts.Arial, consts.Bold)
	text := internal.NewText(fpdf, math, font)

	signature := internal.NewSignature(fpdf, math, text)

	image := internal.NewImage(fpdf, math)

	code := internal.NewCode(fpdf, math)

	tableList := internal.NewTableList(text, font)

	justPdf := &PdfJustPdf{
		Pdf:             fpdf,
		Math:            math,
		Font:            font,
		TextHelper:      text,
		SignHelper:      signature,
		Image:           image,
		Code:            code,
		TableListHelper: tableList,
		pageSize:        pageSize,
		orientation:     orientation,
		calculationMode: false,
		backgroundColor: color.NewWhite(),
	}

	justPdf.TableListHelper.BindGrid(justPdf)

	justPdf.Font.SetFamily(consts.Arial)
	justPdf.Font.SetStyle(consts.Bold)
	justPdf.Font.SetSize(16)
	justPdf.debugMode = false

	justPdf.Pdf.AddPage()

	return justPdf
}

// RegisterHeader define a sequence of Rows, Lines ou TableLists
// which will be added in every new page
func (s *PdfJustPdf) RegisterHeader(closure func()) {
	s.headerClosure = closure
}

// RegisterFooter define a sequence of Rows, Lines ou TableLists
// which will be added in every new page
func (s *PdfJustPdf) RegisterFooter(closure func()) {
	s.footerClosure = closure

	// calculation mode execute all row flow but
	// only to calculate the sum of heights
	s.calculationMode = true
	closure()
	s.calculationMode = false
}

// GetCurrentPage obtain the current page index
// this can be used inside a RegisterFooter/RegisterHeader
// to draw the current page, or to another purposes
func (s *PdfJustPdf) GetCurrentPage() int {
	return s.pageIndex
}

// GetCurrentOffset obtain the current offset in y axis
func (s *PdfJustPdf) GetCurrentOffset() float64 {
	return s.offsetY
}

// SetPageMargins overrides default margins (10,10,10,2(cm))
// the new page margin will affect all PDF pages
func (s *PdfJustPdf) SetPageMargins(left, top, right, bottom float64) {
	if s.pageIndex == 0 {
		s.Pdf.SetY(top)
	}
	s.Pdf.SetTopMargin(top)
	s.Pdf.SetLeftMargin(left)
	s.Pdf.SetRightMargin(right)
	s.Pdf.SetAutoPageBreak(true, bottom)
}

// SetLRMargins only affect left and right margin
func (s *PdfJustPdf) SetLRMargins(left, right float64) {
	s.Pdf.SetLeftMargin(left)
	s.Pdf.SetRightMargin(right)
}

// GetPageMargins returns the set page margins. Comes in order of Left, Top, Right, Bottom
// Default page margins is left: 10, top: 10, right: 10
func (s *PdfJustPdf) GetPageMargins() (left float64, top float64, right float64, bottom float64) {
	return s.Pdf.GetMargins()
}

// Signature add a space for a signature inside a cell,
// the space will have a line and a text below
func (s *PdfJustPdf) Signature(label string, prop ...props.Font) {
	signProp := props.Font{}
	if len(prop) > 0 {
		signProp = prop[0]
	}

	signProp.MakeValid()

	qtdCols := float64(len(s.colsClosures))
	sumOfYOffsets := s.offsetY + s.rowHeight

	s.SignHelper.AddSpaceFor(label, signProp.ToTextProp(consts.Center, 0.0, false, 0), qtdCols, sumOfYOffsets, s.rowColCount)
}

// TableList create a table with multiple rows and columns.
// Headers define the amount of columns from each row.
// Headers have bold style, and localized at the top of table.
// Contents are array of arrays. Each array is one line.
func (s *PdfJustPdf) TableList(header []string, contents [][]string, prop ...props.TableList) {
	s.TableListHelper.Create(header, contents, prop...)
	s.Pdf.PageCount()
}

// SetBorder enable the draw of lines in every cell.
// Draw borders in all columns created.
func (s *PdfJustPdf) SetBorder(on bool) {
	s.debugMode = on
}

// SetBackgroundColor define the background color of the PDF.
// This method can be used to toggle background from rows
func (s *PdfJustPdf) SetBackgroundColor(color color.Color) {
	s.backgroundColor = color
	s.Pdf.SetFillColor(s.backgroundColor.Red, s.backgroundColor.Green, s.backgroundColor.Blue)
}

// SetTextColor define the color of the Text.
func (s *PdfJustPdf) SetTextColor(color color.Color) {
	s.textColor = color
	s.Pdf.SetTextColor(s.backgroundColor.Red, s.backgroundColor.Green, s.backgroundColor.Blue)
}

// GetBorder return the actual border value.
func (s *PdfJustPdf) GetBorder() bool {
	return s.debugMode
}

// GetPageSize return the actual page size
func (s *PdfJustPdf) GetPageSize() (float64, float64) {
	return s.Pdf.GetPageSize()
}

// Line draw a line from margin left to margin right
// in the currently row.
func (s *PdfJustPdf) Line(spaceHeight float64) {
	s.Row(spaceHeight, func() {
		s.Col(func() {
			width, _ := s.Pdf.GetPageSize()
			left, top, right, _ := s.Pdf.GetMargins()

			s.Pdf.Line(left, s.offsetY+top+(spaceHeight/2.0), width-right, s.offsetY+top+(spaceHeight/2.0))
		})
	})
}

// Line draw a line from margin left to margin right
// in the currently row.
func (s *PdfJustPdf) VLine(spaceWidth, spaceHeight float64, color color.Color) {
	_, top, _, _ := s.Pdf.GetMargins()
	s.Pdf.SetDrawColor(color.Red,color.Green,color.Blue)
	s.Pdf.SetLineWidth(spaceWidth)
	s.Pdf.Line(s.Pdf.GetX(), s.offsetY+top+(spaceHeight/2.0)-spaceHeight, s.Pdf.GetX(), s.offsetY+top+(spaceHeight/2.0) + s.rowHeight)
}

// Row define a row and enable add columns inside the row.
func (s *PdfJustPdf) Row(height float64, closure func()) {
	// Used to calculate the height of the footer
	if s.calculationMode {
		s.footerHeight += height
		return
	}

	_, pageHeight := s.Pdf.GetPageSize()
	_, top, _, bottom := s.Pdf.GetMargins()

	totalOffsetY := int(s.offsetY + height + s.footerHeight)
	maxOffsetPage := int(pageHeight - bottom - top)

	// Note: The headerFooterContextActive is needed to avoid recursive
	// calls without end, because footerClosure and headerClosure actually
	// have Row calls too.

	// If the new cell to be added pass the useful space counting the
	// height of the footer, add the footer
	if totalOffsetY > maxOffsetPage {
		if !s.headerFooterContextActive {
			if s.footerClosure != nil {
				s.headerFooterContextActive = true
				s.footerClosure()
				s.headerFooterContextActive = false
			}
			s.offsetY = 0
			s.pageIndex++
		}
	}

	// If is a new page, add the header
	if !s.headerFooterContextActive && s.headerClosure != nil {
		if s.offsetY == 0 {
			s.headerFooterContextActive = true
			s.headerClosure()
			s.headerFooterContextActive = false
		}
	}

	s.rowHeight = height
	s.rowColCount = 0

	// This closure has only JustPdf.Cols, which are
	// not executed firstly, they are added to colsClosures
	// and this enable us to know how many cols will be added
	// and calculate the width from the cells
	closure()

	// Execute the codes inside the Cols
	for _, colClosure := range s.colsClosures {
		colClosure()
	}

	s.colsClosures = nil
	s.offsetY += s.rowHeight
	s.Pdf.Ln(s.rowHeight)
}

// Col create a column inside a row and enable to add
// components inside.
func (s *PdfJustPdf) Col(closure func()) {
	s.colsClosures = append(s.colsClosures, func() {
		widthPerCol := s.Math.GetWidthPerCol(float64(len(s.colsClosures)))
		s.createColSpace(widthPerCol)
		closure()
		s.rowColCount++
	})
}

// ColSpace create an empty column inside a row.
func (s *PdfJustPdf) ColSpace() {
	s.colsClosures = append(s.colsClosures, func() {
		widthPerCol := s.Math.GetWidthPerCol(float64(len(s.colsClosures)))
		s.createColSpace(widthPerCol)
		s.rowColCount++
	})
}

// ColSpaces create some empty columns inside a row.
func (s *PdfJustPdf) ColSpaces(qtd int) {
	for i := 0; i < qtd; i++ {
		s.ColSpace()
	}
}

// Text create a text inside a cell.
func (s *PdfJustPdf) Text(text string, prop ...props.Text) {
	textProp := props.Text{}
	if len(prop) > 0 {
		textProp = prop[0]
	}

	textProp.MakeValid()

	if textProp.Top > s.rowHeight {
		textProp.Top = s.rowHeight
	}

	sumOfYOffsets := textProp.Top + s.offsetY

	s.TextHelper.Add(text, textProp, sumOfYOffsets, s.rowColCount, float64(len(s.colsClosures)))
}

// FileImage add an Image reading from disk inside a cell.
// Defining Image properties.
func (s *PdfJustPdf) FileImage(filePathName string, prop ...props.Rect) error {
	rectProp := props.Rect{}
	if len(prop) > 0 {
		rectProp = prop[0]
	}

	rectProp.MakeValid()

	qtdCols := float64(len(s.colsClosures))
	sumOfyOffsets := s.offsetY + rectProp.Top

	return s.Image.AddFromFile(filePathName, sumOfyOffsets, s.rowColCount, qtdCols, s.rowHeight, rectProp)
}

// Base64Image add an Image reading byte slices inside a cell.
// Defining Image properties.
func (s *PdfJustPdf) Base64Image(base64 string, extension consts.Extension, prop ...props.Rect) error {
	rectProp := props.Rect{}
	if len(prop) > 0 {
		rectProp = prop[0]
	}

	rectProp.MakeValid()

	qtdCols := float64(len(s.colsClosures))
	sumOfyOffsets := s.offsetY + rectProp.Top

	return s.Image.AddFromBase64(base64, sumOfyOffsets, s.rowColCount, qtdCols, s.rowHeight, rectProp, extension)
}

// OutputFileAndClose save pdf in disk.
func (s *PdfJustPdf) OutputFileAndClose(filePathName string) (err error) {
	s.drawLastFooter()
	err = s.Pdf.OutputFileAndClose(filePathName)

	return
}

// Output extract PDF in byte slices
func (s *PdfJustPdf) Output() (bytes.Buffer, error) {
	s.drawLastFooter()
	var buffer bytes.Buffer
	err := s.Pdf.Output(&buffer)
	return buffer, err
}

// Barcode create an barcode inside a cell.
func (s *PdfJustPdf) Barcode(code string, prop ...props.Barcode) (err error) {
	barcodeProp := props.Barcode{}
	if len(prop) > 0 {
		barcodeProp = prop[0]
	}

	barcodeProp.MakeValid()

	qtdCols := float64(len(s.colsClosures))
	sumOfyOffsets := s.offsetY + barcodeProp.Top

	err = s.Code.AddBar(code, sumOfyOffsets, s.rowColCount, qtdCols, s.rowHeight, barcodeProp)

	return
}

// QrCode create a qrcode inside a cell.
func (s *PdfJustPdf) QrCode(code string, prop ...props.Rect) {
	rectProp := props.Rect{}
	if len(prop) > 0 {
		rectProp = prop[0]
	}

	rectProp.MakeValid()

	qtdCols := float64(len(s.colsClosures))
	sumOfyOffsets := s.offsetY + rectProp.Top
	s.Code.AddQr(code, sumOfyOffsets, s.rowColCount, qtdCols, s.rowHeight, rectProp)
}

func (s *PdfJustPdf) createColSpace(actualWidthPerCol float64) {
	border := ""

	if s.debugMode {
		border = "1"
	}

	s.Pdf.CellFormat(actualWidthPerCol, s.rowHeight, "", border, 0.0, "C", !s.backgroundColor.IsWhite(), 0.0, "")
}

func (s *PdfJustPdf) drawLastFooter() {
	if s.footerClosure != nil {
		_, pageHeight := s.Pdf.GetPageSize()
		_, top, _, bottom := s.Pdf.GetMargins()

		if s.offsetY+s.footerHeight < pageHeight-bottom-top {
			s.headerFooterContextActive = true
			s.footerClosure()
			s.headerFooterContextActive = false
		}
	}
}
