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

// BACnetConstructedDataLargeAnalogValueHighLimit is the corresponding interface of BACnetConstructedDataLargeAnalogValueHighLimit
type BACnetConstructedDataLargeAnalogValueHighLimit interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetHighLimit returns HighLimit (property field)
	GetHighLimit() BACnetApplicationTagDouble
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagDouble
	// IsBACnetConstructedDataLargeAnalogValueHighLimit is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataLargeAnalogValueHighLimit()
	// CreateBuilder creates a BACnetConstructedDataLargeAnalogValueHighLimitBuilder
	CreateBACnetConstructedDataLargeAnalogValueHighLimitBuilder() BACnetConstructedDataLargeAnalogValueHighLimitBuilder
}

// _BACnetConstructedDataLargeAnalogValueHighLimit is the data-structure of this message
type _BACnetConstructedDataLargeAnalogValueHighLimit struct {
	BACnetConstructedDataContract
	HighLimit BACnetApplicationTagDouble
}

var _ BACnetConstructedDataLargeAnalogValueHighLimit = (*_BACnetConstructedDataLargeAnalogValueHighLimit)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataLargeAnalogValueHighLimit)(nil)

// NewBACnetConstructedDataLargeAnalogValueHighLimit factory function for _BACnetConstructedDataLargeAnalogValueHighLimit
func NewBACnetConstructedDataLargeAnalogValueHighLimit(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, highLimit BACnetApplicationTagDouble, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLargeAnalogValueHighLimit {
	if highLimit == nil {
		panic("highLimit of type BACnetApplicationTagDouble for BACnetConstructedDataLargeAnalogValueHighLimit must not be nil")
	}
	_result := &_BACnetConstructedDataLargeAnalogValueHighLimit{
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

// BACnetConstructedDataLargeAnalogValueHighLimitBuilder is a builder for BACnetConstructedDataLargeAnalogValueHighLimit
type BACnetConstructedDataLargeAnalogValueHighLimitBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(highLimit BACnetApplicationTagDouble) BACnetConstructedDataLargeAnalogValueHighLimitBuilder
	// WithHighLimit adds HighLimit (property field)
	WithHighLimit(BACnetApplicationTagDouble) BACnetConstructedDataLargeAnalogValueHighLimitBuilder
	// WithHighLimitBuilder adds HighLimit (property field) which is build by the builder
	WithHighLimitBuilder(func(BACnetApplicationTagDoubleBuilder) BACnetApplicationTagDoubleBuilder) BACnetConstructedDataLargeAnalogValueHighLimitBuilder
	// Build builds the BACnetConstructedDataLargeAnalogValueHighLimit or returns an error if something is wrong
	Build() (BACnetConstructedDataLargeAnalogValueHighLimit, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataLargeAnalogValueHighLimit
}

// NewBACnetConstructedDataLargeAnalogValueHighLimitBuilder() creates a BACnetConstructedDataLargeAnalogValueHighLimitBuilder
func NewBACnetConstructedDataLargeAnalogValueHighLimitBuilder() BACnetConstructedDataLargeAnalogValueHighLimitBuilder {
	return &_BACnetConstructedDataLargeAnalogValueHighLimitBuilder{_BACnetConstructedDataLargeAnalogValueHighLimit: new(_BACnetConstructedDataLargeAnalogValueHighLimit)}
}

type _BACnetConstructedDataLargeAnalogValueHighLimitBuilder struct {
	*_BACnetConstructedDataLargeAnalogValueHighLimit

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataLargeAnalogValueHighLimitBuilder) = (*_BACnetConstructedDataLargeAnalogValueHighLimitBuilder)(nil)

func (b *_BACnetConstructedDataLargeAnalogValueHighLimitBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataLargeAnalogValueHighLimitBuilder) WithMandatoryFields(highLimit BACnetApplicationTagDouble) BACnetConstructedDataLargeAnalogValueHighLimitBuilder {
	return b.WithHighLimit(highLimit)
}

func (b *_BACnetConstructedDataLargeAnalogValueHighLimitBuilder) WithHighLimit(highLimit BACnetApplicationTagDouble) BACnetConstructedDataLargeAnalogValueHighLimitBuilder {
	b.HighLimit = highLimit
	return b
}

func (b *_BACnetConstructedDataLargeAnalogValueHighLimitBuilder) WithHighLimitBuilder(builderSupplier func(BACnetApplicationTagDoubleBuilder) BACnetApplicationTagDoubleBuilder) BACnetConstructedDataLargeAnalogValueHighLimitBuilder {
	builder := builderSupplier(b.HighLimit.CreateBACnetApplicationTagDoubleBuilder())
	var err error
	b.HighLimit, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagDoubleBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataLargeAnalogValueHighLimitBuilder) Build() (BACnetConstructedDataLargeAnalogValueHighLimit, error) {
	if b.HighLimit == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'highLimit' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataLargeAnalogValueHighLimit.deepCopy(), nil
}

func (b *_BACnetConstructedDataLargeAnalogValueHighLimitBuilder) MustBuild() BACnetConstructedDataLargeAnalogValueHighLimit {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataLargeAnalogValueHighLimitBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataLargeAnalogValueHighLimitBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataLargeAnalogValueHighLimitBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataLargeAnalogValueHighLimitBuilder().(*_BACnetConstructedDataLargeAnalogValueHighLimitBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataLargeAnalogValueHighLimitBuilder creates a BACnetConstructedDataLargeAnalogValueHighLimitBuilder
func (b *_BACnetConstructedDataLargeAnalogValueHighLimit) CreateBACnetConstructedDataLargeAnalogValueHighLimitBuilder() BACnetConstructedDataLargeAnalogValueHighLimitBuilder {
	if b == nil {
		return NewBACnetConstructedDataLargeAnalogValueHighLimitBuilder()
	}
	return &_BACnetConstructedDataLargeAnalogValueHighLimitBuilder{_BACnetConstructedDataLargeAnalogValueHighLimit: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLargeAnalogValueHighLimit) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_LARGE_ANALOG_VALUE
}

func (m *_BACnetConstructedDataLargeAnalogValueHighLimit) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_HIGH_LIMIT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLargeAnalogValueHighLimit) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLargeAnalogValueHighLimit) GetHighLimit() BACnetApplicationTagDouble {
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

func (m *_BACnetConstructedDataLargeAnalogValueHighLimit) GetActualValue() BACnetApplicationTagDouble {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagDouble(m.GetHighLimit())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLargeAnalogValueHighLimit(structType any) BACnetConstructedDataLargeAnalogValueHighLimit {
	if casted, ok := structType.(BACnetConstructedDataLargeAnalogValueHighLimit); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLargeAnalogValueHighLimit); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLargeAnalogValueHighLimit) GetTypeName() string {
	return "BACnetConstructedDataLargeAnalogValueHighLimit"
}

func (m *_BACnetConstructedDataLargeAnalogValueHighLimit) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (highLimit)
	lengthInBits += m.HighLimit.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLargeAnalogValueHighLimit) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataLargeAnalogValueHighLimit) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataLargeAnalogValueHighLimit BACnetConstructedDataLargeAnalogValueHighLimit, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLargeAnalogValueHighLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLargeAnalogValueHighLimit")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	highLimit, err := ReadSimpleField[BACnetApplicationTagDouble](ctx, "highLimit", ReadComplex[BACnetApplicationTagDouble](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagDouble](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'highLimit' field"))
	}
	m.HighLimit = highLimit

	actualValue, err := ReadVirtualField[BACnetApplicationTagDouble](ctx, "actualValue", (*BACnetApplicationTagDouble)(nil), highLimit)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLargeAnalogValueHighLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLargeAnalogValueHighLimit")
	}

	return m, nil
}

func (m *_BACnetConstructedDataLargeAnalogValueHighLimit) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLargeAnalogValueHighLimit) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLargeAnalogValueHighLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLargeAnalogValueHighLimit")
		}

		if err := WriteSimpleField[BACnetApplicationTagDouble](ctx, "highLimit", m.GetHighLimit(), WriteComplex[BACnetApplicationTagDouble](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'highLimit' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLargeAnalogValueHighLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLargeAnalogValueHighLimit")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLargeAnalogValueHighLimit) IsBACnetConstructedDataLargeAnalogValueHighLimit() {
}

func (m *_BACnetConstructedDataLargeAnalogValueHighLimit) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataLargeAnalogValueHighLimit) deepCopy() *_BACnetConstructedDataLargeAnalogValueHighLimit {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataLargeAnalogValueHighLimitCopy := &_BACnetConstructedDataLargeAnalogValueHighLimit{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.HighLimit.DeepCopy().(BACnetApplicationTagDouble),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataLargeAnalogValueHighLimitCopy
}

func (m *_BACnetConstructedDataLargeAnalogValueHighLimit) String() string {
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
