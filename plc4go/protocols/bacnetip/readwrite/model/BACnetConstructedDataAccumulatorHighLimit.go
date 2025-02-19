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

// BACnetConstructedDataAccumulatorHighLimit is the corresponding interface of BACnetConstructedDataAccumulatorHighLimit
type BACnetConstructedDataAccumulatorHighLimit interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetHighLimit returns HighLimit (property field)
	GetHighLimit() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataAccumulatorHighLimit is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataAccumulatorHighLimit()
	// CreateBuilder creates a BACnetConstructedDataAccumulatorHighLimitBuilder
	CreateBACnetConstructedDataAccumulatorHighLimitBuilder() BACnetConstructedDataAccumulatorHighLimitBuilder
}

// _BACnetConstructedDataAccumulatorHighLimit is the data-structure of this message
type _BACnetConstructedDataAccumulatorHighLimit struct {
	BACnetConstructedDataContract
	HighLimit BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataAccumulatorHighLimit = (*_BACnetConstructedDataAccumulatorHighLimit)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataAccumulatorHighLimit)(nil)

// NewBACnetConstructedDataAccumulatorHighLimit factory function for _BACnetConstructedDataAccumulatorHighLimit
func NewBACnetConstructedDataAccumulatorHighLimit(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, highLimit BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAccumulatorHighLimit {
	if highLimit == nil {
		panic("highLimit of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataAccumulatorHighLimit must not be nil")
	}
	_result := &_BACnetConstructedDataAccumulatorHighLimit{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		HighLimit:                     highLimit,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataAccumulatorHighLimitBuilder is a builder for BACnetConstructedDataAccumulatorHighLimit
type BACnetConstructedDataAccumulatorHighLimitBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(highLimit BACnetApplicationTagUnsignedInteger) BACnetConstructedDataAccumulatorHighLimitBuilder
	// WithHighLimit adds HighLimit (property field)
	WithHighLimit(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataAccumulatorHighLimitBuilder
	// WithHighLimitBuilder adds HighLimit (property field) which is build by the builder
	WithHighLimitBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataAccumulatorHighLimitBuilder
	// Build builds the BACnetConstructedDataAccumulatorHighLimit or returns an error if something is wrong
	Build() (BACnetConstructedDataAccumulatorHighLimit, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataAccumulatorHighLimit
}

// NewBACnetConstructedDataAccumulatorHighLimitBuilder() creates a BACnetConstructedDataAccumulatorHighLimitBuilder
func NewBACnetConstructedDataAccumulatorHighLimitBuilder() BACnetConstructedDataAccumulatorHighLimitBuilder {
	return &_BACnetConstructedDataAccumulatorHighLimitBuilder{_BACnetConstructedDataAccumulatorHighLimit: new(_BACnetConstructedDataAccumulatorHighLimit)}
}

type _BACnetConstructedDataAccumulatorHighLimitBuilder struct {
	*_BACnetConstructedDataAccumulatorHighLimit

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataAccumulatorHighLimitBuilder) = (*_BACnetConstructedDataAccumulatorHighLimitBuilder)(nil)

func (b *_BACnetConstructedDataAccumulatorHighLimitBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataAccumulatorHighLimitBuilder) WithMandatoryFields(highLimit BACnetApplicationTagUnsignedInteger) BACnetConstructedDataAccumulatorHighLimitBuilder {
	return b.WithHighLimit(highLimit)
}

func (b *_BACnetConstructedDataAccumulatorHighLimitBuilder) WithHighLimit(highLimit BACnetApplicationTagUnsignedInteger) BACnetConstructedDataAccumulatorHighLimitBuilder {
	b.HighLimit = highLimit
	return b
}

func (b *_BACnetConstructedDataAccumulatorHighLimitBuilder) WithHighLimitBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataAccumulatorHighLimitBuilder {
	builder := builderSupplier(b.HighLimit.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.HighLimit, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataAccumulatorHighLimitBuilder) Build() (BACnetConstructedDataAccumulatorHighLimit, error) {
	if b.HighLimit == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'highLimit' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataAccumulatorHighLimit.deepCopy(), nil
}

func (b *_BACnetConstructedDataAccumulatorHighLimitBuilder) MustBuild() BACnetConstructedDataAccumulatorHighLimit {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataAccumulatorHighLimitBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataAccumulatorHighLimitBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataAccumulatorHighLimitBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataAccumulatorHighLimitBuilder().(*_BACnetConstructedDataAccumulatorHighLimitBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataAccumulatorHighLimitBuilder creates a BACnetConstructedDataAccumulatorHighLimitBuilder
func (b *_BACnetConstructedDataAccumulatorHighLimit) CreateBACnetConstructedDataAccumulatorHighLimitBuilder() BACnetConstructedDataAccumulatorHighLimitBuilder {
	if b == nil {
		return NewBACnetConstructedDataAccumulatorHighLimitBuilder()
	}
	return &_BACnetConstructedDataAccumulatorHighLimitBuilder{_BACnetConstructedDataAccumulatorHighLimit: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAccumulatorHighLimit) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_ACCUMULATOR
}

func (m *_BACnetConstructedDataAccumulatorHighLimit) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_HIGH_LIMIT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAccumulatorHighLimit) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAccumulatorHighLimit) GetHighLimit() BACnetApplicationTagUnsignedInteger {
	return m.HighLimit
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAccumulatorHighLimit) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetHighLimit())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAccumulatorHighLimit(structType any) BACnetConstructedDataAccumulatorHighLimit {
	if casted, ok := structType.(BACnetConstructedDataAccumulatorHighLimit); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAccumulatorHighLimit); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAccumulatorHighLimit) GetTypeName() string {
	return "BACnetConstructedDataAccumulatorHighLimit"
}

func (m *_BACnetConstructedDataAccumulatorHighLimit) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (highLimit)
	lengthInBits += m.HighLimit.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataAccumulatorHighLimit) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataAccumulatorHighLimit) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataAccumulatorHighLimit BACnetConstructedDataAccumulatorHighLimit, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAccumulatorHighLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAccumulatorHighLimit")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	highLimit, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "highLimit", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'highLimit' field"))
	}
	m.HighLimit = highLimit

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), highLimit)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAccumulatorHighLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAccumulatorHighLimit")
	}

	return m, nil
}

func (m *_BACnetConstructedDataAccumulatorHighLimit) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAccumulatorHighLimit) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAccumulatorHighLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAccumulatorHighLimit")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "highLimit", m.GetHighLimit(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'highLimit' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAccumulatorHighLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAccumulatorHighLimit")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAccumulatorHighLimit) IsBACnetConstructedDataAccumulatorHighLimit() {}

func (m *_BACnetConstructedDataAccumulatorHighLimit) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataAccumulatorHighLimit) deepCopy() *_BACnetConstructedDataAccumulatorHighLimit {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataAccumulatorHighLimitCopy := &_BACnetConstructedDataAccumulatorHighLimit{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.HighLimit.DeepCopy().(BACnetApplicationTagUnsignedInteger),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataAccumulatorHighLimitCopy
}

func (m *_BACnetConstructedDataAccumulatorHighLimit) String() string {
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
