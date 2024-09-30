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

// BACnetConstructedDataTimerMinPresValue is the corresponding interface of BACnetConstructedDataTimerMinPresValue
type BACnetConstructedDataTimerMinPresValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetMinPresValue returns MinPresValue (property field)
	GetMinPresValue() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataTimerMinPresValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataTimerMinPresValue()
	// CreateBuilder creates a BACnetConstructedDataTimerMinPresValueBuilder
	CreateBACnetConstructedDataTimerMinPresValueBuilder() BACnetConstructedDataTimerMinPresValueBuilder
}

// _BACnetConstructedDataTimerMinPresValue is the data-structure of this message
type _BACnetConstructedDataTimerMinPresValue struct {
	BACnetConstructedDataContract
	MinPresValue BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataTimerMinPresValue = (*_BACnetConstructedDataTimerMinPresValue)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataTimerMinPresValue)(nil)

// NewBACnetConstructedDataTimerMinPresValue factory function for _BACnetConstructedDataTimerMinPresValue
func NewBACnetConstructedDataTimerMinPresValue(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, minPresValue BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataTimerMinPresValue {
	if minPresValue == nil {
		panic("minPresValue of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataTimerMinPresValue must not be nil")
	}
	_result := &_BACnetConstructedDataTimerMinPresValue{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		MinPresValue:                  minPresValue,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataTimerMinPresValueBuilder is a builder for BACnetConstructedDataTimerMinPresValue
type BACnetConstructedDataTimerMinPresValueBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(minPresValue BACnetApplicationTagUnsignedInteger) BACnetConstructedDataTimerMinPresValueBuilder
	// WithMinPresValue adds MinPresValue (property field)
	WithMinPresValue(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataTimerMinPresValueBuilder
	// WithMinPresValueBuilder adds MinPresValue (property field) which is build by the builder
	WithMinPresValueBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataTimerMinPresValueBuilder
	// Build builds the BACnetConstructedDataTimerMinPresValue or returns an error if something is wrong
	Build() (BACnetConstructedDataTimerMinPresValue, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataTimerMinPresValue
}

// NewBACnetConstructedDataTimerMinPresValueBuilder() creates a BACnetConstructedDataTimerMinPresValueBuilder
func NewBACnetConstructedDataTimerMinPresValueBuilder() BACnetConstructedDataTimerMinPresValueBuilder {
	return &_BACnetConstructedDataTimerMinPresValueBuilder{_BACnetConstructedDataTimerMinPresValue: new(_BACnetConstructedDataTimerMinPresValue)}
}

type _BACnetConstructedDataTimerMinPresValueBuilder struct {
	*_BACnetConstructedDataTimerMinPresValue

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataTimerMinPresValueBuilder) = (*_BACnetConstructedDataTimerMinPresValueBuilder)(nil)

func (b *_BACnetConstructedDataTimerMinPresValueBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataTimerMinPresValueBuilder) WithMandatoryFields(minPresValue BACnetApplicationTagUnsignedInteger) BACnetConstructedDataTimerMinPresValueBuilder {
	return b.WithMinPresValue(minPresValue)
}

func (b *_BACnetConstructedDataTimerMinPresValueBuilder) WithMinPresValue(minPresValue BACnetApplicationTagUnsignedInteger) BACnetConstructedDataTimerMinPresValueBuilder {
	b.MinPresValue = minPresValue
	return b
}

func (b *_BACnetConstructedDataTimerMinPresValueBuilder) WithMinPresValueBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataTimerMinPresValueBuilder {
	builder := builderSupplier(b.MinPresValue.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.MinPresValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataTimerMinPresValueBuilder) Build() (BACnetConstructedDataTimerMinPresValue, error) {
	if b.MinPresValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'minPresValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataTimerMinPresValue.deepCopy(), nil
}

func (b *_BACnetConstructedDataTimerMinPresValueBuilder) MustBuild() BACnetConstructedDataTimerMinPresValue {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataTimerMinPresValueBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataTimerMinPresValueBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataTimerMinPresValueBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataTimerMinPresValueBuilder().(*_BACnetConstructedDataTimerMinPresValueBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataTimerMinPresValueBuilder creates a BACnetConstructedDataTimerMinPresValueBuilder
func (b *_BACnetConstructedDataTimerMinPresValue) CreateBACnetConstructedDataTimerMinPresValueBuilder() BACnetConstructedDataTimerMinPresValueBuilder {
	if b == nil {
		return NewBACnetConstructedDataTimerMinPresValueBuilder()
	}
	return &_BACnetConstructedDataTimerMinPresValueBuilder{_BACnetConstructedDataTimerMinPresValue: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataTimerMinPresValue) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_TIMER
}

func (m *_BACnetConstructedDataTimerMinPresValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MIN_PRES_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataTimerMinPresValue) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataTimerMinPresValue) GetMinPresValue() BACnetApplicationTagUnsignedInteger {
	return m.MinPresValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataTimerMinPresValue) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetMinPresValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataTimerMinPresValue(structType any) BACnetConstructedDataTimerMinPresValue {
	if casted, ok := structType.(BACnetConstructedDataTimerMinPresValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataTimerMinPresValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataTimerMinPresValue) GetTypeName() string {
	return "BACnetConstructedDataTimerMinPresValue"
}

func (m *_BACnetConstructedDataTimerMinPresValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (minPresValue)
	lengthInBits += m.MinPresValue.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataTimerMinPresValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataTimerMinPresValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataTimerMinPresValue BACnetConstructedDataTimerMinPresValue, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataTimerMinPresValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataTimerMinPresValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	minPresValue, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "minPresValue", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'minPresValue' field"))
	}
	m.MinPresValue = minPresValue

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), minPresValue)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataTimerMinPresValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataTimerMinPresValue")
	}

	return m, nil
}

func (m *_BACnetConstructedDataTimerMinPresValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataTimerMinPresValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataTimerMinPresValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataTimerMinPresValue")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "minPresValue", m.GetMinPresValue(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'minPresValue' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataTimerMinPresValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataTimerMinPresValue")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataTimerMinPresValue) IsBACnetConstructedDataTimerMinPresValue() {}

func (m *_BACnetConstructedDataTimerMinPresValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataTimerMinPresValue) deepCopy() *_BACnetConstructedDataTimerMinPresValue {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataTimerMinPresValueCopy := &_BACnetConstructedDataTimerMinPresValue{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.MinPresValue.DeepCopy().(BACnetApplicationTagUnsignedInteger),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataTimerMinPresValueCopy
}

func (m *_BACnetConstructedDataTimerMinPresValue) String() string {
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
