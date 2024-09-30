/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// Confirmation is the corresponding interface of Confirmation
type Confirmation interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetAlpha returns Alpha (property field)
	GetAlpha() Alpha
	// GetSecondAlpha returns SecondAlpha (property field)
	GetSecondAlpha() Alpha
	// GetConfirmationType returns ConfirmationType (property field)
	GetConfirmationType() ConfirmationType
	// GetIsSuccess returns IsSuccess (virtual field)
	GetIsSuccess() bool
	// IsConfirmation is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsConfirmation()
	// CreateBuilder creates a ConfirmationBuilder
	CreateConfirmationBuilder() ConfirmationBuilder
}

// _Confirmation is the data-structure of this message
type _Confirmation struct {
	Alpha            Alpha
	SecondAlpha      Alpha
	ConfirmationType ConfirmationType
}

var _ Confirmation = (*_Confirmation)(nil)

// NewConfirmation factory function for _Confirmation
func NewConfirmation(alpha Alpha, secondAlpha Alpha, confirmationType ConfirmationType) *_Confirmation {
	if alpha == nil {
		panic("alpha of type Alpha for Confirmation must not be nil")
	}
	return &_Confirmation{Alpha: alpha, SecondAlpha: secondAlpha, ConfirmationType: confirmationType}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ConfirmationBuilder is a builder for Confirmation
type ConfirmationBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(alpha Alpha, confirmationType ConfirmationType) ConfirmationBuilder
	// WithAlpha adds Alpha (property field)
	WithAlpha(Alpha) ConfirmationBuilder
	// WithAlphaBuilder adds Alpha (property field) which is build by the builder
	WithAlphaBuilder(func(AlphaBuilder) AlphaBuilder) ConfirmationBuilder
	// WithSecondAlpha adds SecondAlpha (property field)
	WithOptionalSecondAlpha(Alpha) ConfirmationBuilder
	// WithOptionalSecondAlphaBuilder adds SecondAlpha (property field) which is build by the builder
	WithOptionalSecondAlphaBuilder(func(AlphaBuilder) AlphaBuilder) ConfirmationBuilder
	// WithConfirmationType adds ConfirmationType (property field)
	WithConfirmationType(ConfirmationType) ConfirmationBuilder
	// Build builds the Confirmation or returns an error if something is wrong
	Build() (Confirmation, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() Confirmation
}

// NewConfirmationBuilder() creates a ConfirmationBuilder
func NewConfirmationBuilder() ConfirmationBuilder {
	return &_ConfirmationBuilder{_Confirmation: new(_Confirmation)}
}

type _ConfirmationBuilder struct {
	*_Confirmation

	err *utils.MultiError
}

var _ (ConfirmationBuilder) = (*_ConfirmationBuilder)(nil)

func (b *_ConfirmationBuilder) WithMandatoryFields(alpha Alpha, confirmationType ConfirmationType) ConfirmationBuilder {
	return b.WithAlpha(alpha).WithConfirmationType(confirmationType)
}

func (b *_ConfirmationBuilder) WithAlpha(alpha Alpha) ConfirmationBuilder {
	b.Alpha = alpha
	return b
}

func (b *_ConfirmationBuilder) WithAlphaBuilder(builderSupplier func(AlphaBuilder) AlphaBuilder) ConfirmationBuilder {
	builder := builderSupplier(b.Alpha.CreateAlphaBuilder())
	var err error
	b.Alpha, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "AlphaBuilder failed"))
	}
	return b
}

func (b *_ConfirmationBuilder) WithOptionalSecondAlpha(secondAlpha Alpha) ConfirmationBuilder {
	b.SecondAlpha = secondAlpha
	return b
}

func (b *_ConfirmationBuilder) WithOptionalSecondAlphaBuilder(builderSupplier func(AlphaBuilder) AlphaBuilder) ConfirmationBuilder {
	builder := builderSupplier(b.SecondAlpha.CreateAlphaBuilder())
	var err error
	b.SecondAlpha, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "AlphaBuilder failed"))
	}
	return b
}

func (b *_ConfirmationBuilder) WithConfirmationType(confirmationType ConfirmationType) ConfirmationBuilder {
	b.ConfirmationType = confirmationType
	return b
}

func (b *_ConfirmationBuilder) Build() (Confirmation, error) {
	if b.Alpha == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'alpha' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._Confirmation.deepCopy(), nil
}

func (b *_ConfirmationBuilder) MustBuild() Confirmation {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ConfirmationBuilder) DeepCopy() any {
	_copy := b.CreateConfirmationBuilder().(*_ConfirmationBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateConfirmationBuilder creates a ConfirmationBuilder
func (b *_Confirmation) CreateConfirmationBuilder() ConfirmationBuilder {
	if b == nil {
		return NewConfirmationBuilder()
	}
	return &_ConfirmationBuilder{_Confirmation: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_Confirmation) GetAlpha() Alpha {
	return m.Alpha
}

func (m *_Confirmation) GetSecondAlpha() Alpha {
	return m.SecondAlpha
}

func (m *_Confirmation) GetConfirmationType() ConfirmationType {
	return m.ConfirmationType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_Confirmation) GetIsSuccess() bool {
	ctx := context.Background()
	_ = ctx
	secondAlpha := m.GetSecondAlpha()
	_ = secondAlpha
	return bool(bool((m.GetConfirmationType()) == (ConfirmationType_CONFIRMATION_SUCCESSFUL)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastConfirmation(structType any) Confirmation {
	if casted, ok := structType.(Confirmation); ok {
		return casted
	}
	if casted, ok := structType.(*Confirmation); ok {
		return *casted
	}
	return nil
}

func (m *_Confirmation) GetTypeName() string {
	return "Confirmation"
}

func (m *_Confirmation) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (alpha)
	lengthInBits += m.Alpha.GetLengthInBits(ctx)

	// Optional Field (secondAlpha)
	if m.SecondAlpha != nil {
		lengthInBits += m.SecondAlpha.GetLengthInBits(ctx)
	}

	// Simple field (confirmationType)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_Confirmation) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ConfirmationParse(ctx context.Context, theBytes []byte) (Confirmation, error) {
	return ConfirmationParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func ConfirmationParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (Confirmation, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (Confirmation, error) {
		return ConfirmationParseWithBuffer(ctx, readBuffer)
	}
}

func ConfirmationParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (Confirmation, error) {
	v, err := (&_Confirmation{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_Confirmation) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__confirmation Confirmation, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("Confirmation"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for Confirmation")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	alpha, err := ReadSimpleField[Alpha](ctx, "alpha", ReadComplex[Alpha](AlphaParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'alpha' field"))
	}
	m.Alpha = alpha

	var secondAlpha Alpha
	_secondAlpha, err := ReadOptionalField[Alpha](ctx, "secondAlpha", ReadComplex[Alpha](AlphaParseWithBuffer, readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'secondAlpha' field"))
	}
	if _secondAlpha != nil {
		secondAlpha = *_secondAlpha
		m.SecondAlpha = secondAlpha
	}

	confirmationType, err := ReadEnumField[ConfirmationType](ctx, "confirmationType", "ConfirmationType", ReadEnum(ConfirmationTypeByValue, ReadByte(readBuffer, 8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'confirmationType' field"))
	}
	m.ConfirmationType = confirmationType

	isSuccess, err := ReadVirtualField[bool](ctx, "isSuccess", (*bool)(nil), bool((confirmationType) == (ConfirmationType_CONFIRMATION_SUCCESSFUL)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isSuccess' field"))
	}
	_ = isSuccess

	if closeErr := readBuffer.CloseContext("Confirmation"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for Confirmation")
	}

	return m, nil
}

func (m *_Confirmation) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_Confirmation) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("Confirmation"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for Confirmation")
	}

	if err := WriteSimpleField[Alpha](ctx, "alpha", m.GetAlpha(), WriteComplex[Alpha](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'alpha' field")
	}

	if err := WriteOptionalField[Alpha](ctx, "secondAlpha", GetRef(m.GetSecondAlpha()), WriteComplex[Alpha](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'secondAlpha' field")
	}

	if err := WriteSimpleEnumField[ConfirmationType](ctx, "confirmationType", "ConfirmationType", m.GetConfirmationType(), WriteEnum[ConfirmationType, byte](ConfirmationType.GetValue, ConfirmationType.PLC4XEnumName, WriteByte(writeBuffer, 8))); err != nil {
		return errors.Wrap(err, "Error serializing 'confirmationType' field")
	}
	// Virtual field
	isSuccess := m.GetIsSuccess()
	_ = isSuccess
	if _isSuccessErr := writeBuffer.WriteVirtual(ctx, "isSuccess", m.GetIsSuccess()); _isSuccessErr != nil {
		return errors.Wrap(_isSuccessErr, "Error serializing 'isSuccess' field")
	}

	if popErr := writeBuffer.PopContext("Confirmation"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for Confirmation")
	}
	return nil
}

func (m *_Confirmation) IsConfirmation() {}

func (m *_Confirmation) DeepCopy() any {
	return m.deepCopy()
}

func (m *_Confirmation) deepCopy() *_Confirmation {
	if m == nil {
		return nil
	}
	_ConfirmationCopy := &_Confirmation{
		m.Alpha.DeepCopy().(Alpha),
		m.SecondAlpha.DeepCopy().(Alpha),
		m.ConfirmationType,
	}
	return _ConfirmationCopy
}

func (m *_Confirmation) String() string {
	if m == nil {
		return "<nil>"
	}
	wb := utils.NewWriteBufferBoxBased(
		utils.WithWriteBufferBoxBasedMergeSingleBoxes(),
		utils.WithWriteBufferBoxBasedOmitEmptyBoxes(),
		utils.WithWriteBufferBoxBasedPrintPosLengthFooter(),
	)
	if err := wb.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return wb.GetBox().String()
}
