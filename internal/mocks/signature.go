// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import "github.com/stretchr/testify/mock"
import "github.com/muhammadmuhlas/just_pdf/pkg/props"

// Signature is an autogenerated mock type for the Signature type
type Signature struct {
	mock.Mock
}

// AddSpaceFor provides a mock function with given fields: label, textProp, qtdCols, marginTop, actualCol
func (_m *Signature) AddSpaceFor(label string, textProp props.Text, qtdCols float64, marginTop float64, actualCol float64) {
	_m.Called(label, textProp, qtdCols, marginTop, actualCol)
}
