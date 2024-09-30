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

// ParameterValueInterfaceOptions2 is the corresponding interface of ParameterValueInterfaceOptions2
type ParameterValueInterfaceOptions2 interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ParameterValue
	// GetValue returns Value (property field)
	GetValue() InterfaceOptions2
	// GetData returns Data (property field)
	GetData() []byte
	// IsParameterValueInterfaceOptions2 is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsParameterValueInterfaceOptions2()
	// CreateBuilder creates a ParameterValueInterfaceOptions2Builder
	CreateParameterValueInterfaceOptions2Builder() ParameterValueInterfaceOptions2Builder
}

// _ParameterValueInterfaceOptions2 is the data-structure of this message
type _ParameterValueInterfaceOptions2 struct {
	ParameterValueContract
	Value InterfaceOptions2
	Data  []byte
}

var _ ParameterValueInterfaceOptions2 = (*_ParameterValueInterfaceOptions2)(nil)
var _ ParameterValueRequirements = (*_ParameterValueInterfaceOptions2)(nil)

// NewParameterValueInterfaceOptions2 factory function for _ParameterValueInterfaceOptions2
func NewParameterValueInterfaceOptions2(value InterfaceOptions2, data []byte, numBytes uint8) *_ParameterValueInterfaceOptions2 {
	if value == nil {
		panic("value of type InterfaceOptions2 for ParameterValueInterfaceOptions2 must not be nil")
	}
	_result := &_ParameterValueInterfaceOptions2{
		ParameterValueContract: NewParameterValue(numBytes),
		Value:                  value,
		Data:                   data,
	}
	_result.ParameterValueContract.(*_ParameterValue)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ParameterValueInterfaceOptions2Builder is a builder for ParameterValueInterfaceOptions2
type ParameterValueInterfaceOptions2Builder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(value InterfaceOptions2, data []byte) ParameterValueInterfaceOptions2Builder
	// WithValue adds Value (property field)
	WithValue(InterfaceOptions2) ParameterValueInterfaceOptions2Builder
	// WithValueBuilder adds Value (property field) which is build by the builder
	WithValueBuilder(func(InterfaceOptions2Builder) InterfaceOptions2Builder) ParameterValueInterfaceOptions2Builder
	// WithData adds Data (property field)
	WithData(...byte) ParameterValueInterfaceOptions2Builder
	// Build builds the ParameterValueInterfaceOptions2 or returns an error if something is wrong
	Build() (ParameterValueInterfaceOptions2, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ParameterValueInterfaceOptions2
}

// NewParameterValueInterfaceOptions2Builder() creates a ParameterValueInterfaceOptions2Builder
func NewParameterValueInterfaceOptions2Builder() ParameterValueInterfaceOptions2Builder {
	return &_ParameterValueInterfaceOptions2Builder{_ParameterValueInterfaceOptions2: new(_ParameterValueInterfaceOptions2)}
}

type _ParameterValueInterfaceOptions2Builder struct {
	*_ParameterValueInterfaceOptions2

	parentBuilder *_ParameterValueBuilder

	err *utils.MultiError
}

var _ (ParameterValueInterfaceOptions2Builder) = (*_ParameterValueInterfaceOptions2Builder)(nil)

func (b *_ParameterValueInterfaceOptions2Builder) setParent(contract ParameterValueContract) {
	b.ParameterValueContract = contract
}

func (b *_ParameterValueInterfaceOptions2Builder) WithMandatoryFields(value InterfaceOptions2, data []byte) ParameterValueInterfaceOptions2Builder {
	return b.WithValue(value).WithData(data...)
}

func (b *_ParameterValueInterfaceOptions2Builder) WithValue(value InterfaceOptions2) ParameterValueInterfaceOptions2Builder {
	b.Value = value
	return b
}

func (b *_ParameterValueInterfaceOptions2Builder) WithValueBuilder(builderSupplier func(InterfaceOptions2Builder) InterfaceOptions2Builder) ParameterValueInterfaceOptions2Builder {
	builder := builderSupplier(b.Value.CreateInterfaceOptions2Builder())
	var err error
	b.Value, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "InterfaceOptions2Builder failed"))
	}
	return b
}

func (b *_ParameterValueInterfaceOptions2Builder) WithData(data ...byte) ParameterValueInterfaceOptions2Builder {
	b.Data = data
	return b
}

func (b *_ParameterValueInterfaceOptions2Builder) Build() (ParameterValueInterfaceOptions2, error) {
	if b.Value == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'value' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ParameterValueInterfaceOptions2.deepCopy(), nil
}

func (b *_ParameterValueInterfaceOptions2Builder) MustBuild() ParameterValueInterfaceOptions2 {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_ParameterValueInterfaceOptions2Builder) Done() ParameterValueBuilder {
	return b.parentBuilder
}

func (b *_ParameterValueInterfaceOptions2Builder) buildForParameterValue() (ParameterValue, error) {
	return b.Build()
}

func (b *_ParameterValueInterfaceOptions2Builder) DeepCopy() any {
	_copy := b.CreateParameterValueInterfaceOptions2Builder().(*_ParameterValueInterfaceOptions2Builder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateParameterValueInterfaceOptions2Builder creates a ParameterValueInterfaceOptions2Builder
func (b *_ParameterValueInterfaceOptions2) CreateParameterValueInterfaceOptions2Builder() ParameterValueInterfaceOptions2Builder {
	if b == nil {
		return NewParameterValueInterfaceOptions2Builder()
	}
	return &_ParameterValueInterfaceOptions2Builder{_ParameterValueInterfaceOptions2: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ParameterValueInterfaceOptions2) GetParameterType() ParameterType {
	return ParameterType_INTERFACE_OPTIONS_2
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ParameterValueInterfaceOptions2) GetParent() ParameterValueContract {
	return m.ParameterValueContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ParameterValueInterfaceOptions2) GetValue() InterfaceOptions2 {
	return m.Value
}

func (m *_ParameterValueInterfaceOptions2) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastParameterValueInterfaceOptions2(structType any) ParameterValueInterfaceOptions2 {
	if casted, ok := structType.(ParameterValueInterfaceOptions2); ok {
		return casted
	}
	if casted, ok := structType.(*ParameterValueInterfaceOptions2); ok {
		return *casted
	}
	return nil
}

func (m *_ParameterValueInterfaceOptions2) GetTypeName() string {
	return "ParameterValueInterfaceOptions2"
}

func (m *_ParameterValueInterfaceOptions2) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ParameterValueContract.(*_ParameterValue).getLengthInBits(ctx))

	// Simple field (value)
	lengthInBits += m.Value.GetLengthInBits(ctx)

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *_ParameterValueInterfaceOptions2) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ParameterValueInterfaceOptions2) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ParameterValue, parameterType ParameterType, numBytes uint8) (__parameterValueInterfaceOptions2 ParameterValueInterfaceOptions2, err error) {
	m.ParameterValueContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ParameterValueInterfaceOptions2"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ParameterValueInterfaceOptions2")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((numBytes) >= (1))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "InterfaceOptions2 has exactly one byte"})
	}

	value, err := ReadSimpleField[InterfaceOptions2](ctx, "value", ReadComplex[InterfaceOptions2](InterfaceOptions2ParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	data, err := readBuffer.ReadByteArray("data", int(int32(numBytes)-int32(int32(1))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'data' field"))
	}
	m.Data = data

	if closeErr := readBuffer.CloseContext("ParameterValueInterfaceOptions2"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ParameterValueInterfaceOptions2")
	}

	return m, nil
}

func (m *_ParameterValueInterfaceOptions2) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ParameterValueInterfaceOptions2) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ParameterValueInterfaceOptions2"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ParameterValueInterfaceOptions2")
		}

		if err := WriteSimpleField[InterfaceOptions2](ctx, "value", m.GetValue(), WriteComplex[InterfaceOptions2](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'value' field")
		}

		if err := WriteByteArrayField(ctx, "data", m.GetData(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("ParameterValueInterfaceOptions2"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ParameterValueInterfaceOptions2")
		}
		return nil
	}
	return m.ParameterValueContract.(*_ParameterValue).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ParameterValueInterfaceOptions2) IsParameterValueInterfaceOptions2() {}

func (m *_ParameterValueInterfaceOptions2) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ParameterValueInterfaceOptions2) deepCopy() *_ParameterValueInterfaceOptions2 {
	if m == nil {
		return nil
	}
	_ParameterValueInterfaceOptions2Copy := &_ParameterValueInterfaceOptions2{
		m.ParameterValueContract.(*_ParameterValue).deepCopy(),
		m.Value.DeepCopy().(InterfaceOptions2),
		utils.DeepCopySlice[byte, byte](m.Data),
	}
	m.ParameterValueContract.(*_ParameterValue)._SubType = m
	return _ParameterValueInterfaceOptions2Copy
}

func (m *_ParameterValueInterfaceOptions2) String() string {
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
