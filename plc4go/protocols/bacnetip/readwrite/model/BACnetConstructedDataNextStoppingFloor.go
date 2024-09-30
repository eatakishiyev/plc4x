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

// BACnetConstructedDataNextStoppingFloor is the corresponding interface of BACnetConstructedDataNextStoppingFloor
type BACnetConstructedDataNextStoppingFloor interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetNextStoppingFloor returns NextStoppingFloor (property field)
	GetNextStoppingFloor() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataNextStoppingFloor is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataNextStoppingFloor()
	// CreateBuilder creates a BACnetConstructedDataNextStoppingFloorBuilder
	CreateBACnetConstructedDataNextStoppingFloorBuilder() BACnetConstructedDataNextStoppingFloorBuilder
}

// _BACnetConstructedDataNextStoppingFloor is the data-structure of this message
type _BACnetConstructedDataNextStoppingFloor struct {
	BACnetConstructedDataContract
	NextStoppingFloor BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataNextStoppingFloor = (*_BACnetConstructedDataNextStoppingFloor)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataNextStoppingFloor)(nil)

// NewBACnetConstructedDataNextStoppingFloor factory function for _BACnetConstructedDataNextStoppingFloor
func NewBACnetConstructedDataNextStoppingFloor(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, nextStoppingFloor BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataNextStoppingFloor {
	if nextStoppingFloor == nil {
		panic("nextStoppingFloor of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataNextStoppingFloor must not be nil")
	}
	_result := &_BACnetConstructedDataNextStoppingFloor{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		NextStoppingFloor:             nextStoppingFloor,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataNextStoppingFloorBuilder is a builder for BACnetConstructedDataNextStoppingFloor
type BACnetConstructedDataNextStoppingFloorBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(nextStoppingFloor BACnetApplicationTagUnsignedInteger) BACnetConstructedDataNextStoppingFloorBuilder
	// WithNextStoppingFloor adds NextStoppingFloor (property field)
	WithNextStoppingFloor(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataNextStoppingFloorBuilder
	// WithNextStoppingFloorBuilder adds NextStoppingFloor (property field) which is build by the builder
	WithNextStoppingFloorBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataNextStoppingFloorBuilder
	// Build builds the BACnetConstructedDataNextStoppingFloor or returns an error if something is wrong
	Build() (BACnetConstructedDataNextStoppingFloor, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataNextStoppingFloor
}

// NewBACnetConstructedDataNextStoppingFloorBuilder() creates a BACnetConstructedDataNextStoppingFloorBuilder
func NewBACnetConstructedDataNextStoppingFloorBuilder() BACnetConstructedDataNextStoppingFloorBuilder {
	return &_BACnetConstructedDataNextStoppingFloorBuilder{_BACnetConstructedDataNextStoppingFloor: new(_BACnetConstructedDataNextStoppingFloor)}
}

type _BACnetConstructedDataNextStoppingFloorBuilder struct {
	*_BACnetConstructedDataNextStoppingFloor

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataNextStoppingFloorBuilder) = (*_BACnetConstructedDataNextStoppingFloorBuilder)(nil)

func (b *_BACnetConstructedDataNextStoppingFloorBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataNextStoppingFloorBuilder) WithMandatoryFields(nextStoppingFloor BACnetApplicationTagUnsignedInteger) BACnetConstructedDataNextStoppingFloorBuilder {
	return b.WithNextStoppingFloor(nextStoppingFloor)
}

func (b *_BACnetConstructedDataNextStoppingFloorBuilder) WithNextStoppingFloor(nextStoppingFloor BACnetApplicationTagUnsignedInteger) BACnetConstructedDataNextStoppingFloorBuilder {
	b.NextStoppingFloor = nextStoppingFloor
	return b
}

func (b *_BACnetConstructedDataNextStoppingFloorBuilder) WithNextStoppingFloorBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataNextStoppingFloorBuilder {
	builder := builderSupplier(b.NextStoppingFloor.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.NextStoppingFloor, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataNextStoppingFloorBuilder) Build() (BACnetConstructedDataNextStoppingFloor, error) {
	if b.NextStoppingFloor == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'nextStoppingFloor' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataNextStoppingFloor.deepCopy(), nil
}

func (b *_BACnetConstructedDataNextStoppingFloorBuilder) MustBuild() BACnetConstructedDataNextStoppingFloor {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataNextStoppingFloorBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataNextStoppingFloorBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataNextStoppingFloorBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataNextStoppingFloorBuilder().(*_BACnetConstructedDataNextStoppingFloorBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataNextStoppingFloorBuilder creates a BACnetConstructedDataNextStoppingFloorBuilder
func (b *_BACnetConstructedDataNextStoppingFloor) CreateBACnetConstructedDataNextStoppingFloorBuilder() BACnetConstructedDataNextStoppingFloorBuilder {
	if b == nil {
		return NewBACnetConstructedDataNextStoppingFloorBuilder()
	}
	return &_BACnetConstructedDataNextStoppingFloorBuilder{_BACnetConstructedDataNextStoppingFloor: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataNextStoppingFloor) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataNextStoppingFloor) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_NEXT_STOPPING_FLOOR
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataNextStoppingFloor) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataNextStoppingFloor) GetNextStoppingFloor() BACnetApplicationTagUnsignedInteger {
	return m.NextStoppingFloor
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataNextStoppingFloor) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetNextStoppingFloor())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataNextStoppingFloor(structType any) BACnetConstructedDataNextStoppingFloor {
	if casted, ok := structType.(BACnetConstructedDataNextStoppingFloor); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataNextStoppingFloor); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataNextStoppingFloor) GetTypeName() string {
	return "BACnetConstructedDataNextStoppingFloor"
}

func (m *_BACnetConstructedDataNextStoppingFloor) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (nextStoppingFloor)
	lengthInBits += m.NextStoppingFloor.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataNextStoppingFloor) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataNextStoppingFloor) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataNextStoppingFloor BACnetConstructedDataNextStoppingFloor, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataNextStoppingFloor"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataNextStoppingFloor")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	nextStoppingFloor, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "nextStoppingFloor", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'nextStoppingFloor' field"))
	}
	m.NextStoppingFloor = nextStoppingFloor

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), nextStoppingFloor)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataNextStoppingFloor"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataNextStoppingFloor")
	}

	return m, nil
}

func (m *_BACnetConstructedDataNextStoppingFloor) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataNextStoppingFloor) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataNextStoppingFloor"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataNextStoppingFloor")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "nextStoppingFloor", m.GetNextStoppingFloor(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'nextStoppingFloor' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataNextStoppingFloor"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataNextStoppingFloor")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataNextStoppingFloor) IsBACnetConstructedDataNextStoppingFloor() {}

func (m *_BACnetConstructedDataNextStoppingFloor) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataNextStoppingFloor) deepCopy() *_BACnetConstructedDataNextStoppingFloor {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataNextStoppingFloorCopy := &_BACnetConstructedDataNextStoppingFloor{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.NextStoppingFloor.DeepCopy().(BACnetApplicationTagUnsignedInteger),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataNextStoppingFloorCopy
}

func (m *_BACnetConstructedDataNextStoppingFloor) String() string {
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
