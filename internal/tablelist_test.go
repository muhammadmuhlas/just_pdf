package internal_test

import (
	"fmt"
	"github.com/muhammadmuhlas/just_pdf/internal"
	"github.com/muhammadmuhlas/just_pdf/internal/mocks"
	"github.com/muhammadmuhlas/just_pdf/pkg/color"
	"github.com/muhammadmuhlas/just_pdf/pkg/consts"
	"github.com/muhammadmuhlas/just_pdf/pkg/props"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestNewTableList(t *testing.T) {
	// Act
	tableList := internal.NewTableList(nil, nil)

	// Assert
	assert.NotNil(t, tableList)
	assert.Equal(t, fmt.Sprintf("%T", tableList), "*internal.tableList")
}

func TestTableList_Create_WhenHeaderIsNil(t *testing.T) {
	// Arrange
	text := &mocks.Text{}
	text.On("GetLinesQuantity", mock.Anything, mock.Anything, mock.Anything).Return(1)
	sut := internal.NewTableList(text, nil)

	_, contents := getContents()

	// Act
	sut.Create(nil, contents)

	// Assert
	text.AssertNotCalled(t, "GetLinesQuantity")
}

func TestTableList_Create_WhenHeaderIsEmpty(t *testing.T) {
	// Arrange
	text := &mocks.Text{}
	text.On("GetLinesQuantity", mock.Anything, mock.Anything, mock.Anything).Return(1)
	sut := internal.NewTableList(text, nil)

	_, contents := getContents()

	// Act
	sut.Create([]string{}, contents)

	// Assert
	text.AssertNotCalled(t, "GetLinesQuantity")
}

func TestTableList_Create_WhenContentIsNil(t *testing.T) {
	// Arrange
	text := &mocks.Text{}
	text.On("GetLinesQuantity", mock.Anything, mock.Anything, mock.Anything).Return(1)
	sut := internal.NewTableList(text, nil)

	headers, _ := getContents()

	// Act
	sut.Create(headers, nil)

	// Assert
	text.AssertNotCalled(t, "GetLinesQuantity")
}

func TestTableList_Create_WhenContentIsEmpty(t *testing.T) {
	// Arrange
	text := &mocks.Text{}
	text.On("GetLinesQuantity", mock.Anything, mock.Anything, mock.Anything).Return(1)
	sut := internal.NewTableList(text, nil)

	headers, _ := getContents()

	// Act
	sut.Create(headers, [][]string{})

	// Assert
	text.AssertNotCalled(t, "GetLinesQuantity")
}

func TestTableList_Create_Happy(t *testing.T) {
	// Arrange
	text := &mocks.Text{}
	text.On("GetLinesQuantity", mock.Anything, mock.Anything, mock.Anything).Return(1)

	font := &mocks.Font{}
	font.On("GetFont").Return(consts.Arial, consts.Bold, 1.0)
	font.On("GetScaleFactor").Return(1.5)

	justPdfGrid := &mocks.JustPdf{}
	justPdfGrid.On("Row", mock.Anything, mock.Anything).Return(nil)
	justPdfGrid.On("Line", mock.Anything).Return(nil)
	justPdfGrid.On("SetBackgroundColor", mock.Anything).Return(nil)

	sut := internal.NewTableList(text, font)
	sut.BindGrid(justPdfGrid)

	headers, contents := getContents()

	// Act
	sut.Create(headers, contents, props.TableList{
		Line: true,
	})

	// Assert
	text.AssertNotCalled(t, "GetLinesQuantity")
	text.AssertNumberOfCalls(t, "GetLinesQuantity", 84)

	font.AssertCalled(t, "GetFont")
	font.AssertNumberOfCalls(t, "GetFont", 21)

	justPdfGrid.AssertCalled(t, "Row", mock.Anything, mock.Anything)
	justPdfGrid.AssertNumberOfCalls(t, "Row", 21)
	justPdfGrid.AssertNumberOfCalls(t, "Line", 20)
	justPdfGrid.AssertNotCalled(t, "SetBackgroundColor")
}

func TestTableList_Create_HappyWithBackgroundColor(t *testing.T) {
	// Arrange
	text := &mocks.Text{}
	text.On("GetLinesQuantity", mock.Anything, mock.Anything, mock.Anything).Return(1)

	font := &mocks.Font{}
	font.On("GetFont").Return(consts.Arial, consts.Bold, 1.0)
	font.On("GetScaleFactor").Return(1.5)

	justPdfGrid := &mocks.JustPdf{}
	justPdfGrid.On("Row", mock.Anything, mock.Anything).Return(nil)
	justPdfGrid.On("Line", mock.Anything).Return(nil)
	justPdfGrid.On("SetBackgroundColor", mock.Anything).Return(nil)

	sut := internal.NewTableList(text, font)
	sut.BindGrid(justPdfGrid)

	headers, contents := getContents()
	color := color.Color{
		Red:   200,
		Green: 200,
		Blue:  200,
	}

	// Act
	sut.Create(headers, contents, props.TableList{
		AlternatedBackground: &color,
	})

	// Assert
	text.AssertNotCalled(t, "GetLinesQuantity")
	text.AssertNumberOfCalls(t, "GetLinesQuantity", 84)

	font.AssertCalled(t, "GetFont")
	font.AssertNumberOfCalls(t, "GetFont", 21)

	justPdfGrid.AssertCalled(t, "Row", mock.Anything, mock.Anything)
	justPdfGrid.AssertNumberOfCalls(t, "Row", 21)

	justPdfGrid.AssertNotCalled(t, "Line")

	//justPdfGrid.AssertCalled(t, "SetBackgroundColor", color)
	justPdfGrid.AssertNumberOfCalls(t, "SetBackgroundColor", 20)
}

func TestTableList_Create_Happy_Without_Line(t *testing.T) {
	// Arrange
	text := &mocks.Text{}
	text.On("GetLinesQuantity", mock.Anything, mock.Anything, mock.Anything).Return(1)

	font := &mocks.Font{}
	font.On("GetFont").Return(consts.Arial, consts.Bold, 1.0)
	font.On("GetScaleFactor").Return(1.5)

	justPdfGrid := &mocks.JustPdf{}
	justPdfGrid.On("Row", mock.Anything, mock.Anything).Return(nil)
	justPdfGrid.On("Line", mock.Anything).Return(nil)
	justPdfGrid.On("SetBackgroundColor", mock.Anything).Return(nil)

	sut := internal.NewTableList(text, font)
	sut.BindGrid(justPdfGrid)

	headers, contents := getContents()

	// Act

	sut.Create(headers, contents)

	// Assert
	text.AssertNotCalled(t, "GetLinesQuantity")
	text.AssertNumberOfCalls(t, "GetLinesQuantity", 84)

	font.AssertCalled(t, "GetFont")
	font.AssertNumberOfCalls(t, "GetFont", 21)

	justPdfGrid.AssertCalled(t, "Row", mock.Anything, mock.Anything)
	justPdfGrid.AssertNumberOfCalls(t, "Row", 21)
	justPdfGrid.AssertNumberOfCalls(t, "Line", 0)
}

func TestTableList_Create_WhenContentIsEmptyWithLine(t *testing.T) {
	// Arrange
	text := &mocks.Text{}
	text.On("GetLinesQuantity", mock.Anything, mock.Anything, mock.Anything).Return(1)
	sut := internal.NewTableList(text, nil)

	justPdfGrid := mocks.JustPdf{}
	justPdfGrid.On("Line", mock.Anything).Return(nil)

	headers, _ := getContents()

	// Act
	sut.Create(headers, [][]string{}, props.TableList{
		Line: true,
	})

	// Assert
	text.AssertNotCalled(t, "GetLinesQuantity")
	justPdfGrid.AssertNotCalled(t, "Line")
}

func getContents() ([]string, [][]string) {
	header := []string{"j = 0", "j = 1", "j = 2", "j = 4"}

	contents := [][]string{}
	for i := 0; i < 20; i++ {
		content := []string{}
		for j := 0; j < 4; j++ {
			content = append(content, fmt.Sprintf("i = %d, j = %d", i, j))
		}
		contents = append(contents, content)
	}

	return header, contents
}
