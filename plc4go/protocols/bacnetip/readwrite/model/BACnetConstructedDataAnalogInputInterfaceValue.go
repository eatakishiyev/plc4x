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

// BACnetConstructedDataAnalogInputInterfaceValue is the corresponding interface of BACnetConstructedDataAnalogInputInterfaceValue
type BACnetConstructedDataAnalogInputInterfaceValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetInterfaceValue returns InterfaceValue (property field)
	GetInterfaceValue() BACnetOptionalREAL
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetOptionalREAL
	// IsBACnetConstructedDataAnalogInputInterfaceValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataAnalogInputInterfaceValue()
	// CreateBuilder creates a BACnetConstructedDataAnalogInputInterfaceValueBuilder
	CreateBACnetConstructedDataAnalogInputInterfaceValueBuilder() BACnetConstructedDataAnalogInputInterfaceValueBuilder
}

// _BACnetConstructedDataAnalogInputInterfaceValue is the data-structure of this message
type _BACnetConstructedDataAnalogInputInterfaceValue struct {
	BACnetConstructedDataContract
	InterfaceValue BACnetOptionalREAL
}

var _ BACnetConstructedDataAnalogInputInterfaceValue = (*_BACnetConstructedDataAnalogInputInterfaceValue)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataAnalogInputInterfaceValue)(nil)

// NewBACnetConstructedDataAnalogInputInterfaceValue factory function for _BACnetConstructedDataAnalogInputInterfaceValue
func NewBACnetConstructedDataAnalogInputInterfaceValue(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, interfaceValue BACnetOptionalREAL, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAnalogInputInterfaceValue {
	if interfaceValue == nil {
		panic("interfaceValue of type BACnetOptionalREAL for BACnetConstructedDataAnalogInputInterfaceValue must not be nil")
	}
	_result := &_BACnetConstructedDataAnalogInputInterfaceValue{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		InterfaceValue:                interfaceValue,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataAnalogInputInterfaceValueBuilder is a builder for BACnetConstructedDataAnalogInputInterfaceValue
type BACnetConstructedDataAnalogInputInterfaceValueBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(interfaceValue BACnetOptionalREAL) BACnetConstructedDataAnalogInputInterfaceValueBuilder
	// WithInterfaceValue adds InterfaceValue (property field)
	WithInterfaceValue(BACnetOptionalREAL) BACnetConstructedDataAnalogInputInterfaceValueBuilder
	// WithInterfaceValueBuilder adds InterfaceValue (property field) which is build by the builder
	WithInterfaceValueBuilder(func(BACnetOptionalREALBuilder) BACnetOptionalREALBuilder) BACnetConstructedDataAnalogInputInterfaceValueBuilder
	// Build builds the BACnetConstructedDataAnalogInputInterfaceValue or returns an error if something is wrong
	Build() (BACnetConstructedDataAnalogInputInterfaceValue, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataAnalogInputInterfaceValue
}

// NewBACnetConstructedDataAnalogInputInterfaceValueBuilder() creates a BACnetConstructedDataAnalogInputInterfaceValueBuilder
func NewBACnetConstructedDataAnalogInputInterfaceValueBuilder() BACnetConstructedDataAnalogInputInterfaceValueBuilder {
	return &_BACnetConstructedDataAnalogInputInterfaceValueBuilder{_BACnetConstructedDataAnalogInputInterfaceValue: new(_BACnetConstructedDataAnalogInputInterfaceValue)}
}

type _BACnetConstructedDataAnalogInputInterfaceValueBuilder struct {
	*_BACnetConstructedDataAnalogInputInterfaceValue

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataAnalogInputInterfaceValueBuilder) = (*_BACnetConstructedDataAnalogInputInterfaceValueBuilder)(nil)

func (b *_BACnetConstructedDataAnalogInputInterfaceValueBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataAnalogInputInterfaceValueBuilder) WithMandatoryFields(interfaceValue BACnetOptionalREAL) BACnetConstructedDataAnalogInputInterfaceValueBuilder {
	return b.WithInterfaceValue(interfaceValue)
}

func (b *_BACnetConstructedDataAnalogInputInterfaceValueBuilder) WithInterfaceValue(interfaceValue BACnetOptionalREAL) BACnetConstructedDataAnalogInputInterfaceValueBuilder {
	b.InterfaceValue = interfaceValue
	return b
}

func (b *_BACnetConstructedDataAnalogInputInterfaceValueBuilder) WithInterfaceValueBuilder(builderSupplier func(BACnetOptionalREALBuilder) BACnetOptionalREALBuilder) BACnetConstructedDataAnalogInputInterfaceValueBuilder {
	builder := builderSupplier(b.InterfaceValue.CreateBACnetOptionalREALBuilder())
	var err error
	b.InterfaceValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetOptionalREALBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataAnalogInputInterfaceValueBuilder) Build() (BACnetConstructedDataAnalogInputInterfaceValue, error) {
	if b.InterfaceValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'interfaceValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataAnalogInputInterfaceValue.deepCopy(), nil
}

func (b *_BACnetConstructedDataAnalogInputInterfaceValueBuilder) MustBuild() BACnetConstructedDataAnalogInputInterfaceValue {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataAnalogInputInterfaceValueBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataAnalogInputInterfaceValueBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataAnalogInputInterfaceValueBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataAnalogInputInterfaceValueBuilder().(*_BACnetConstructedDataAnalogInputInterfaceValueBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataAnalogInputInterfaceValueBuilder creates a BACnetConstructedDataAnalogInputInterfaceValueBuilder
func (b *_BACnetConstructedDataAnalogInputInterfaceValue) CreateBACnetConstructedDataAnalogInputInterfaceValueBuilder() BACnetConstructedDataAnalogInputInterfaceValueBuilder {
	if b == nil {
		return NewBACnetConstructedDataAnalogInputInterfaceValueBuilder()
	}
	return &_BACnetConstructedDataAnalogInputInterfaceValueBuilder{_BACnetConstructedDataAnalogInputInterfaceValue: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAnalogInputInterfaceValue) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_ANALOG_INPUT
}

func (m *_BACnetConstructedDataAnalogInputInterfaceValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_INTERFACE_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAnalogInputInterfaceValue) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAnalogInputInterfaceValue) GetInterfaceValue() BACnetOptionalREAL {
	return m.InterfaceValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAnalogInputInterfaceValue) GetActualValue() BACnetOptionalREAL {
	ctx := context.Background()
	_ = ctx
	return CastBACnetOptionalREAL(m.GetInterfaceValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAnalogInputInterfaceValue(structType any) BACnetConstructedDataAnalogInputInterfaceValue {
	if casted, ok := structType.(BACnetConstructedDataAnalogInputInterfaceValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAnalogInputInterfaceValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAnalogInputInterfaceValue) GetTypeName() string {
	return "BACnetConstructedDataAnalogInputInterfaceValue"
}

func (m *_BACnetConstructedDataAnalogInputInterfaceValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (interfaceValue)
	lengthInBits += m.InterfaceValue.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataAnalogInputInterfaceValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataAnalogInputInterfaceValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataAnalogInputInterfaceValue BACnetConstructedDataAnalogInputInterfaceValue, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAnalogInputInterfaceValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAnalogInputInterfaceValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	interfaceValue, err := ReadSimpleField[BACnetOptionalREAL](ctx, "interfaceValue", ReadComplex[BACnetOptionalREAL](BACnetOptionalREALParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'interfaceValue' field"))
	}
	m.InterfaceValue = interfaceValue

	actualValue, err := ReadVirtualField[BACnetOptionalREAL](ctx, "actualValue", (*BACnetOptionalREAL)(nil), interfaceValue)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAnalogInputInterfaceValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAnalogInputInterfaceValue")
	}

	return m, nil
}

func (m *_BACnetConstructedDataAnalogInputInterfaceValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAnalogInputInterfaceValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAnalogInputInterfaceValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAnalogInputInterfaceValue")
		}

		if err := WriteSimpleField[BACnetOptionalREAL](ctx, "interfaceValue", m.GetInterfaceValue(), WriteComplex[BACnetOptionalREAL](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'interfaceValue' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAnalogInputInterfaceValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAnalogInputInterfaceValue")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAnalogInputInterfaceValue) IsBACnetConstructedDataAnalogInputInterfaceValue() {
}

func (m *_BACnetConstructedDataAnalogInputInterfaceValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataAnalogInputInterfaceValue) deepCopy() *_BACnetConstructedDataAnalogInputInterfaceValue {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataAnalogInputInterfaceValueCopy := &_BACnetConstructedDataAnalogInputInterfaceValue{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.InterfaceValue.DeepCopy().(BACnetOptionalREAL),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataAnalogInputInterfaceValueCopy
}

func (m *_BACnetConstructedDataAnalogInputInterfaceValue) String() string {
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
