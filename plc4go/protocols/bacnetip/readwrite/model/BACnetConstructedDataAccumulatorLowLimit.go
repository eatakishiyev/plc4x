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

// BACnetConstructedDataAccumulatorLowLimit is the corresponding interface of BACnetConstructedDataAccumulatorLowLimit
type BACnetConstructedDataAccumulatorLowLimit interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetLowLimit returns LowLimit (property field)
	GetLowLimit() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataAccumulatorLowLimit is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataAccumulatorLowLimit()
	// CreateBuilder creates a BACnetConstructedDataAccumulatorLowLimitBuilder
	CreateBACnetConstructedDataAccumulatorLowLimitBuilder() BACnetConstructedDataAccumulatorLowLimitBuilder
}

// _BACnetConstructedDataAccumulatorLowLimit is the data-structure of this message
type _BACnetConstructedDataAccumulatorLowLimit struct {
	BACnetConstructedDataContract
	LowLimit BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataAccumulatorLowLimit = (*_BACnetConstructedDataAccumulatorLowLimit)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataAccumulatorLowLimit)(nil)

// NewBACnetConstructedDataAccumulatorLowLimit factory function for _BACnetConstructedDataAccumulatorLowLimit
func NewBACnetConstructedDataAccumulatorLowLimit(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, lowLimit BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAccumulatorLowLimit {
	if lowLimit == nil {
		panic("lowLimit of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataAccumulatorLowLimit must not be nil")
	}
	_result := &_BACnetConstructedDataAccumulatorLowLimit{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		LowLimit:                      lowLimit,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataAccumulatorLowLimitBuilder is a builder for BACnetConstructedDataAccumulatorLowLimit
type BACnetConstructedDataAccumulatorLowLimitBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(lowLimit BACnetApplicationTagUnsignedInteger) BACnetConstructedDataAccumulatorLowLimitBuilder
	// WithLowLimit adds LowLimit (property field)
	WithLowLimit(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataAccumulatorLowLimitBuilder
	// WithLowLimitBuilder adds LowLimit (property field) which is build by the builder
	WithLowLimitBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataAccumulatorLowLimitBuilder
	// Build builds the BACnetConstructedDataAccumulatorLowLimit or returns an error if something is wrong
	Build() (BACnetConstructedDataAccumulatorLowLimit, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataAccumulatorLowLimit
}

// NewBACnetConstructedDataAccumulatorLowLimitBuilder() creates a BACnetConstructedDataAccumulatorLowLimitBuilder
func NewBACnetConstructedDataAccumulatorLowLimitBuilder() BACnetConstructedDataAccumulatorLowLimitBuilder {
	return &_BACnetConstructedDataAccumulatorLowLimitBuilder{_BACnetConstructedDataAccumulatorLowLimit: new(_BACnetConstructedDataAccumulatorLowLimit)}
}

type _BACnetConstructedDataAccumulatorLowLimitBuilder struct {
	*_BACnetConstructedDataAccumulatorLowLimit

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataAccumulatorLowLimitBuilder) = (*_BACnetConstructedDataAccumulatorLowLimitBuilder)(nil)

func (b *_BACnetConstructedDataAccumulatorLowLimitBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataAccumulatorLowLimitBuilder) WithMandatoryFields(lowLimit BACnetApplicationTagUnsignedInteger) BACnetConstructedDataAccumulatorLowLimitBuilder {
	return b.WithLowLimit(lowLimit)
}

func (b *_BACnetConstructedDataAccumulatorLowLimitBuilder) WithLowLimit(lowLimit BACnetApplicationTagUnsignedInteger) BACnetConstructedDataAccumulatorLowLimitBuilder {
	b.LowLimit = lowLimit
	return b
}

func (b *_BACnetConstructedDataAccumulatorLowLimitBuilder) WithLowLimitBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataAccumulatorLowLimitBuilder {
	builder := builderSupplier(b.LowLimit.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.LowLimit, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataAccumulatorLowLimitBuilder) Build() (BACnetConstructedDataAccumulatorLowLimit, error) {
	if b.LowLimit == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'lowLimit' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataAccumulatorLowLimit.deepCopy(), nil
}

func (b *_BACnetConstructedDataAccumulatorLowLimitBuilder) MustBuild() BACnetConstructedDataAccumulatorLowLimit {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataAccumulatorLowLimitBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataAccumulatorLowLimitBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataAccumulatorLowLimitBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataAccumulatorLowLimitBuilder().(*_BACnetConstructedDataAccumulatorLowLimitBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataAccumulatorLowLimitBuilder creates a BACnetConstructedDataAccumulatorLowLimitBuilder
func (b *_BACnetConstructedDataAccumulatorLowLimit) CreateBACnetConstructedDataAccumulatorLowLimitBuilder() BACnetConstructedDataAccumulatorLowLimitBuilder {
	if b == nil {
		return NewBACnetConstructedDataAccumulatorLowLimitBuilder()
	}
	return &_BACnetConstructedDataAccumulatorLowLimitBuilder{_BACnetConstructedDataAccumulatorLowLimit: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAccumulatorLowLimit) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_ACCUMULATOR
}

func (m *_BACnetConstructedDataAccumulatorLowLimit) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LOW_LIMIT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAccumulatorLowLimit) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAccumulatorLowLimit) GetLowLimit() BACnetApplicationTagUnsignedInteger {
	return m.LowLimit
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAccumulatorLowLimit) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetLowLimit())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAccumulatorLowLimit(structType any) BACnetConstructedDataAccumulatorLowLimit {
	if casted, ok := structType.(BACnetConstructedDataAccumulatorLowLimit); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAccumulatorLowLimit); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAccumulatorLowLimit) GetTypeName() string {
	return "BACnetConstructedDataAccumulatorLowLimit"
}

func (m *_BACnetConstructedDataAccumulatorLowLimit) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (lowLimit)
	lengthInBits += m.LowLimit.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataAccumulatorLowLimit) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataAccumulatorLowLimit) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataAccumulatorLowLimit BACnetConstructedDataAccumulatorLowLimit, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAccumulatorLowLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAccumulatorLowLimit")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	lowLimit, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "lowLimit", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lowLimit' field"))
	}
	m.LowLimit = lowLimit

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), lowLimit)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAccumulatorLowLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAccumulatorLowLimit")
	}

	return m, nil
}

func (m *_BACnetConstructedDataAccumulatorLowLimit) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAccumulatorLowLimit) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAccumulatorLowLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAccumulatorLowLimit")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "lowLimit", m.GetLowLimit(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'lowLimit' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAccumulatorLowLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAccumulatorLowLimit")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAccumulatorLowLimit) IsBACnetConstructedDataAccumulatorLowLimit() {}

func (m *_BACnetConstructedDataAccumulatorLowLimit) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataAccumulatorLowLimit) deepCopy() *_BACnetConstructedDataAccumulatorLowLimit {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataAccumulatorLowLimitCopy := &_BACnetConstructedDataAccumulatorLowLimit{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.LowLimit.DeepCopy().(BACnetApplicationTagUnsignedInteger),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataAccumulatorLowLimitCopy
}

func (m *_BACnetConstructedDataAccumulatorLowLimit) String() string {
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
