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

// BACnetConstructedDataCarMode is the corresponding interface of BACnetConstructedDataCarMode
type BACnetConstructedDataCarMode interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetCarMode returns CarMode (property field)
	GetCarMode() BACnetLiftCarModeTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetLiftCarModeTagged
	// IsBACnetConstructedDataCarMode is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataCarMode()
	// CreateBuilder creates a BACnetConstructedDataCarModeBuilder
	CreateBACnetConstructedDataCarModeBuilder() BACnetConstructedDataCarModeBuilder
}

// _BACnetConstructedDataCarMode is the data-structure of this message
type _BACnetConstructedDataCarMode struct {
	BACnetConstructedDataContract
	CarMode BACnetLiftCarModeTagged
}

var _ BACnetConstructedDataCarMode = (*_BACnetConstructedDataCarMode)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataCarMode)(nil)

// NewBACnetConstructedDataCarMode factory function for _BACnetConstructedDataCarMode
func NewBACnetConstructedDataCarMode(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, carMode BACnetLiftCarModeTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataCarMode {
	if carMode == nil {
		panic("carMode of type BACnetLiftCarModeTagged for BACnetConstructedDataCarMode must not be nil")
	}
	_result := &_BACnetConstructedDataCarMode{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		CarMode:                       carMode,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataCarModeBuilder is a builder for BACnetConstructedDataCarMode
type BACnetConstructedDataCarModeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(carMode BACnetLiftCarModeTagged) BACnetConstructedDataCarModeBuilder
	// WithCarMode adds CarMode (property field)
	WithCarMode(BACnetLiftCarModeTagged) BACnetConstructedDataCarModeBuilder
	// WithCarModeBuilder adds CarMode (property field) which is build by the builder
	WithCarModeBuilder(func(BACnetLiftCarModeTaggedBuilder) BACnetLiftCarModeTaggedBuilder) BACnetConstructedDataCarModeBuilder
	// Build builds the BACnetConstructedDataCarMode or returns an error if something is wrong
	Build() (BACnetConstructedDataCarMode, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataCarMode
}

// NewBACnetConstructedDataCarModeBuilder() creates a BACnetConstructedDataCarModeBuilder
func NewBACnetConstructedDataCarModeBuilder() BACnetConstructedDataCarModeBuilder {
	return &_BACnetConstructedDataCarModeBuilder{_BACnetConstructedDataCarMode: new(_BACnetConstructedDataCarMode)}
}

type _BACnetConstructedDataCarModeBuilder struct {
	*_BACnetConstructedDataCarMode

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataCarModeBuilder) = (*_BACnetConstructedDataCarModeBuilder)(nil)

func (b *_BACnetConstructedDataCarModeBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataCarModeBuilder) WithMandatoryFields(carMode BACnetLiftCarModeTagged) BACnetConstructedDataCarModeBuilder {
	return b.WithCarMode(carMode)
}

func (b *_BACnetConstructedDataCarModeBuilder) WithCarMode(carMode BACnetLiftCarModeTagged) BACnetConstructedDataCarModeBuilder {
	b.CarMode = carMode
	return b
}

func (b *_BACnetConstructedDataCarModeBuilder) WithCarModeBuilder(builderSupplier func(BACnetLiftCarModeTaggedBuilder) BACnetLiftCarModeTaggedBuilder) BACnetConstructedDataCarModeBuilder {
	builder := builderSupplier(b.CarMode.CreateBACnetLiftCarModeTaggedBuilder())
	var err error
	b.CarMode, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetLiftCarModeTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataCarModeBuilder) Build() (BACnetConstructedDataCarMode, error) {
	if b.CarMode == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'carMode' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataCarMode.deepCopy(), nil
}

func (b *_BACnetConstructedDataCarModeBuilder) MustBuild() BACnetConstructedDataCarMode {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataCarModeBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataCarModeBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataCarModeBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataCarModeBuilder().(*_BACnetConstructedDataCarModeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataCarModeBuilder creates a BACnetConstructedDataCarModeBuilder
func (b *_BACnetConstructedDataCarMode) CreateBACnetConstructedDataCarModeBuilder() BACnetConstructedDataCarModeBuilder {
	if b == nil {
		return NewBACnetConstructedDataCarModeBuilder()
	}
	return &_BACnetConstructedDataCarModeBuilder{_BACnetConstructedDataCarMode: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataCarMode) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataCarMode) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_CAR_MODE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataCarMode) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataCarMode) GetCarMode() BACnetLiftCarModeTagged {
	return m.CarMode
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataCarMode) GetActualValue() BACnetLiftCarModeTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetLiftCarModeTagged(m.GetCarMode())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataCarMode(structType any) BACnetConstructedDataCarMode {
	if casted, ok := structType.(BACnetConstructedDataCarMode); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataCarMode); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataCarMode) GetTypeName() string {
	return "BACnetConstructedDataCarMode"
}

func (m *_BACnetConstructedDataCarMode) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (carMode)
	lengthInBits += m.CarMode.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataCarMode) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataCarMode) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataCarMode BACnetConstructedDataCarMode, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataCarMode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataCarMode")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	carMode, err := ReadSimpleField[BACnetLiftCarModeTagged](ctx, "carMode", ReadComplex[BACnetLiftCarModeTagged](BACnetLiftCarModeTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'carMode' field"))
	}
	m.CarMode = carMode

	actualValue, err := ReadVirtualField[BACnetLiftCarModeTagged](ctx, "actualValue", (*BACnetLiftCarModeTagged)(nil), carMode)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataCarMode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataCarMode")
	}

	return m, nil
}

func (m *_BACnetConstructedDataCarMode) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataCarMode) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataCarMode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataCarMode")
		}

		if err := WriteSimpleField[BACnetLiftCarModeTagged](ctx, "carMode", m.GetCarMode(), WriteComplex[BACnetLiftCarModeTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'carMode' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataCarMode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataCarMode")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataCarMode) IsBACnetConstructedDataCarMode() {}

func (m *_BACnetConstructedDataCarMode) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataCarMode) deepCopy() *_BACnetConstructedDataCarMode {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataCarModeCopy := &_BACnetConstructedDataCarMode{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.CarMode.DeepCopy().(BACnetLiftCarModeTagged),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataCarModeCopy
}

func (m *_BACnetConstructedDataCarMode) String() string {
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
