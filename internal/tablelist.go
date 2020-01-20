package internal

import (
	"github.com/muhammadmuhlas/just_pdf/pkg/color"
	"github.com/muhammadmuhlas/just_pdf/pkg/props"
)

// JustPdfGridPart is the abstraction to deal with the gris system inside the table list
type JustPdfGridPart interface {
	// Grid System
	Row(height float64, closure func())
	SetLRMargins(left, right float64)
	Col(closure func())
	ColSpace()

	// Helpers
	SetBackgroundColor(color color.Color)
	GetCurrentOffset() float64

	// Outside Col/Row Components
	Line(spaceHeight float64)
}

// TableList is the abstraction to create a table with header and contents
type TableList interface {
	Create(header []string, contents [][]string, prop ...props.TableList)
	BindGrid(part JustPdfGridPart)
}

type tableList struct {
	pdf  JustPdfGridPart
	text Text
	font Font
}

// NewTableList create a TableList
func NewTableList(text Text, font Font) *tableList {
	return &tableList{
		text: text,
		font: font,
	}
}

// BindGrid bind the grid system to TableList
func (s *tableList) BindGrid(pdf JustPdfGridPart) {
	s.pdf = pdf
}

// Create create a header section with a list of strings and
// create many rows with contents
func (s *tableList) Create(header []string, contents [][]string, prop ...props.TableList) {
	if len(header) == 0 {
		return
	}

	if len(contents) == 0 {
		return
	}

	tableProp := props.TableList{}

	if len(prop) > 0 {
		tableProp = prop[0]
	}

	tableProp.MakeValid()

	qtdCols := float64(len(header))

	headerTextProp := tableProp.HeaderProp.ToTextProp(tableProp.Align, 0.0, false, 1.0)
	headerHeight := s.calcLinesHeight(header, headerTextProp, qtdCols)

	// Draw header
	s.pdf.Row(headerHeight, func() {
		headerMarginTop := 2.0

		for i, h := range header {
			hs := h
			is := i

			s.pdf.SetBackgroundColor(*tableProp.HeaderColor)
			s.pdf.Col(func() {
				if headerMarginTop > headerHeight {
					headerMarginTop = headerHeight
				}

				reason := hs

				sumOyYOffesets := headerMarginTop + s.pdf.GetCurrentOffset() + 2.5
				headerTextProp.Align = tableProp.CustomAlign[is]
				headerTextProp.Color = color.NewWhite()
				s.text.Add(reason, headerTextProp, sumOyYOffesets, float64(is), qtdCols)
			})
		}
	})

	contentMarginTop := 0.7

	// Draw contents
	for index, content := range contents {
		contentTextProp := tableProp.ContentProp.ToTextProp(tableProp.Align, 0.0, false, 0.2)
		contentHeight := s.calcLinesContentHeight(content, contentTextProp, qtdCols)

		s.pdf.Row(contentHeight, func() {
			if (tableProp.AlternatedBackground != nil || tableProp.AlternatedOddBackground != nil) && index%2 == 0 {
				s.pdf.SetBackgroundColor(*tableProp.AlternatedBackground)
			} else {
				s.pdf.SetBackgroundColor(*tableProp.AlternatedOddBackground)
			}
			for j, c := range content {
				cs := c
				js := j
				hs := float64(len(header))

				sumOyYOffesets := contentMarginTop + s.pdf.GetCurrentOffset() + 2.0
				s.pdf.Col(func() {
					contentTextProp.Align = tableProp.CustomAlign[js]
					contentTextProp.Color = *tableProp.ContentFontColor
					s.text.Add(cs, contentTextProp, sumOyYOffesets, float64(js), hs)
				})
			}
		})
		s.pdf.SetBackgroundColor(color.NewWhite())

		if tableProp.Line {
			s.pdf.Line(1.0)
		}
	}
}

func (s *tableList) calcLinesHeight(textList []string, textProp props.Text, qtdCols float64) float64 {
	maxLines := 2.0

	for _, text := range textList {
		qtdLines := float64(s.text.GetLinesQuantity(text, textProp, qtdCols))
		if qtdLines > maxLines {
			maxLines = qtdLines
		}
	}

	_, _, fontSize := s.font.GetFont()

	// Font size corrected by the scale factor from "mm" inside gofpdf f.k
	fontHeight := fontSize / s.font.GetScaleFactor()

	return fontHeight*maxLines + 3.0
}

func (s *tableList) calcLinesContentHeight(textList []string, textProp props.Text, qtdCols float64) float64 {
	maxLines := 2.0

	for _, text := range textList {
		qtdLines := float64(s.text.GetLinesQuantity(text, textProp, qtdCols))
		if qtdLines > maxLines {
			maxLines = qtdLines
		}
	}

	_, _, fontSize := s.font.GetFont()

	// Font size corrected by the scale factor from "mm" inside gofpdf f.k
	fontHeight := fontSize / s.font.GetScaleFactor()

	return fontHeight*maxLines + 3.0
}
