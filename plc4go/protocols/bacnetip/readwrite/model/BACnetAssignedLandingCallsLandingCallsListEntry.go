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

// BACnetAssignedLandingCallsLandingCallsListEntry is the corresponding interface of BACnetAssignedLandingCallsLandingCallsListEntry
type BACnetAssignedLandingCallsLandingCallsListEntry interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetFloorNumber returns FloorNumber (property field)
	GetFloorNumber() BACnetContextTagUnsignedInteger
	// GetDirection returns Direction (property field)
	GetDirection() BACnetLiftCarDirectionTagged
	// IsBACnetAssignedLandingCallsLandingCallsListEntry is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetAssignedLandingCallsLandingCallsListEntry()
	// CreateBuilder creates a BACnetAssignedLandingCallsLandingCallsListEntryBuilder
	CreateBACnetAssignedLandingCallsLandingCallsListEntryBuilder() BACnetAssignedLandingCallsLandingCallsListEntryBuilder
}

// _BACnetAssignedLandingCallsLandingCallsListEntry is the data-structure of this message
type _BACnetAssignedLandingCallsLandingCallsListEntry struct {
	FloorNumber BACnetContextTagUnsignedInteger
	Direction   BACnetLiftCarDirectionTagged
}

var _ BACnetAssignedLandingCallsLandingCallsListEntry = (*_BACnetAssignedLandingCallsLandingCallsListEntry)(nil)

// NewBACnetAssignedLandingCallsLandingCallsListEntry factory function for _BACnetAssignedLandingCallsLandingCallsListEntry
func NewBACnetAssignedLandingCallsLandingCallsListEntry(floorNumber BACnetContextTagUnsignedInteger, direction BACnetLiftCarDirectionTagged) *_BACnetAssignedLandingCallsLandingCallsListEntry {
	if floorNumber == nil {
		panic("floorNumber of type BACnetContextTagUnsignedInteger for BACnetAssignedLandingCallsLandingCallsListEntry must not be nil")
	}
	if direction == nil {
		panic("direction of type BACnetLiftCarDirectionTagged for BACnetAssignedLandingCallsLandingCallsListEntry must not be nil")
	}
	return &_BACnetAssignedLandingCallsLandingCallsListEntry{FloorNumber: floorNumber, Direction: direction}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetAssignedLandingCallsLandingCallsListEntryBuilder is a builder for BACnetAssignedLandingCallsLandingCallsListEntry
type BACnetAssignedLandingCallsLandingCallsListEntryBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(floorNumber BACnetContextTagUnsignedInteger, direction BACnetLiftCarDirectionTagged) BACnetAssignedLandingCallsLandingCallsListEntryBuilder
	// WithFloorNumber adds FloorNumber (property field)
	WithFloorNumber(BACnetContextTagUnsignedInteger) BACnetAssignedLandingCallsLandingCallsListEntryBuilder
	// WithFloorNumberBuilder adds FloorNumber (property field) which is build by the builder
	WithFloorNumberBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetAssignedLandingCallsLandingCallsListEntryBuilder
	// WithDirection adds Direction (property field)
	WithDirection(BACnetLiftCarDirectionTagged) BACnetAssignedLandingCallsLandingCallsListEntryBuilder
	// WithDirectionBuilder adds Direction (property field) which is build by the builder
	WithDirectionBuilder(func(BACnetLiftCarDirectionTaggedBuilder) BACnetLiftCarDirectionTaggedBuilder) BACnetAssignedLandingCallsLandingCallsListEntryBuilder
	// Build builds the BACnetAssignedLandingCallsLandingCallsListEntry or returns an error if something is wrong
	Build() (BACnetAssignedLandingCallsLandingCallsListEntry, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetAssignedLandingCallsLandingCallsListEntry
}

// NewBACnetAssignedLandingCallsLandingCallsListEntryBuilder() creates a BACnetAssignedLandingCallsLandingCallsListEntryBuilder
func NewBACnetAssignedLandingCallsLandingCallsListEntryBuilder() BACnetAssignedLandingCallsLandingCallsListEntryBuilder {
	return &_BACnetAssignedLandingCallsLandingCallsListEntryBuilder{_BACnetAssignedLandingCallsLandingCallsListEntry: new(_BACnetAssignedLandingCallsLandingCallsListEntry)}
}

type _BACnetAssignedLandingCallsLandingCallsListEntryBuilder struct {
	*_BACnetAssignedLandingCallsLandingCallsListEntry

	err *utils.MultiError
}

var _ (BACnetAssignedLandingCallsLandingCallsListEntryBuilder) = (*_BACnetAssignedLandingCallsLandingCallsListEntryBuilder)(nil)

func (b *_BACnetAssignedLandingCallsLandingCallsListEntryBuilder) WithMandatoryFields(floorNumber BACnetContextTagUnsignedInteger, direction BACnetLiftCarDirectionTagged) BACnetAssignedLandingCallsLandingCallsListEntryBuilder {
	return b.WithFloorNumber(floorNumber).WithDirection(direction)
}

func (b *_BACnetAssignedLandingCallsLandingCallsListEntryBuilder) WithFloorNumber(floorNumber BACnetContextTagUnsignedInteger) BACnetAssignedLandingCallsLandingCallsListEntryBuilder {
	b.FloorNumber = floorNumber
	return b
}

func (b *_BACnetAssignedLandingCallsLandingCallsListEntryBuilder) WithFloorNumberBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetAssignedLandingCallsLandingCallsListEntryBuilder {
	builder := builderSupplier(b.FloorNumber.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	b.FloorNumber, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetAssignedLandingCallsLandingCallsListEntryBuilder) WithDirection(direction BACnetLiftCarDirectionTagged) BACnetAssignedLandingCallsLandingCallsListEntryBuilder {
	b.Direction = direction
	return b
}

func (b *_BACnetAssignedLandingCallsLandingCallsListEntryBuilder) WithDirectionBuilder(builderSupplier func(BACnetLiftCarDirectionTaggedBuilder) BACnetLiftCarDirectionTaggedBuilder) BACnetAssignedLandingCallsLandingCallsListEntryBuilder {
	builder := builderSupplier(b.Direction.CreateBACnetLiftCarDirectionTaggedBuilder())
	var err error
	b.Direction, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetLiftCarDirectionTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetAssignedLandingCallsLandingCallsListEntryBuilder) Build() (BACnetAssignedLandingCallsLandingCallsListEntry, error) {
	if b.FloorNumber == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'floorNumber' not set"))
	}
	if b.Direction == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'direction' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetAssignedLandingCallsLandingCallsListEntry.deepCopy(), nil
}

func (b *_BACnetAssignedLandingCallsLandingCallsListEntryBuilder) MustBuild() BACnetAssignedLandingCallsLandingCallsListEntry {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetAssignedLandingCallsLandingCallsListEntryBuilder) DeepCopy() any {
	_copy := b.CreateBACnetAssignedLandingCallsLandingCallsListEntryBuilder().(*_BACnetAssignedLandingCallsLandingCallsListEntryBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetAssignedLandingCallsLandingCallsListEntryBuilder creates a BACnetAssignedLandingCallsLandingCallsListEntryBuilder
func (b *_BACnetAssignedLandingCallsLandingCallsListEntry) CreateBACnetAssignedLandingCallsLandingCallsListEntryBuilder() BACnetAssignedLandingCallsLandingCallsListEntryBuilder {
	if b == nil {
		return NewBACnetAssignedLandingCallsLandingCallsListEntryBuilder()
	}
	return &_BACnetAssignedLandingCallsLandingCallsListEntryBuilder{_BACnetAssignedLandingCallsLandingCallsListEntry: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetAssignedLandingCallsLandingCallsListEntry) GetFloorNumber() BACnetContextTagUnsignedInteger {
	return m.FloorNumber
}

func (m *_BACnetAssignedLandingCallsLandingCallsListEntry) GetDirection() BACnetLiftCarDirectionTagged {
	return m.Direction
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetAssignedLandingCallsLandingCallsListEntry(structType any) BACnetAssignedLandingCallsLandingCallsListEntry {
	if casted, ok := structType.(BACnetAssignedLandingCallsLandingCallsListEntry); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetAssignedLandingCallsLandingCallsListEntry); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetAssignedLandingCallsLandingCallsListEntry) GetTypeName() string {
	return "BACnetAssignedLandingCallsLandingCallsListEntry"
}

func (m *_BACnetAssignedLandingCallsLandingCallsListEntry) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (floorNumber)
	lengthInBits += m.FloorNumber.GetLengthInBits(ctx)

	// Simple field (direction)
	lengthInBits += m.Direction.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetAssignedLandingCallsLandingCallsListEntry) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetAssignedLandingCallsLandingCallsListEntryParse(ctx context.Context, theBytes []byte) (BACnetAssignedLandingCallsLandingCallsListEntry, error) {
	return BACnetAssignedLandingCallsLandingCallsListEntryParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetAssignedLandingCallsLandingCallsListEntryParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAssignedLandingCallsLandingCallsListEntry, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAssignedLandingCallsLandingCallsListEntry, error) {
		return BACnetAssignedLandingCallsLandingCallsListEntryParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetAssignedLandingCallsLandingCallsListEntryParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAssignedLandingCallsLandingCallsListEntry, error) {
	v, err := (&_BACnetAssignedLandingCallsLandingCallsListEntry{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetAssignedLandingCallsLandingCallsListEntry) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetAssignedLandingCallsLandingCallsListEntry BACnetAssignedLandingCallsLandingCallsListEntry, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetAssignedLandingCallsLandingCallsListEntry"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetAssignedLandingCallsLandingCallsListEntry")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	floorNumber, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "floorNumber", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'floorNumber' field"))
	}
	m.FloorNumber = floorNumber

	direction, err := ReadSimpleField[BACnetLiftCarDirectionTagged](ctx, "direction", ReadComplex[BACnetLiftCarDirectionTagged](BACnetLiftCarDirectionTaggedParseWithBufferProducer((uint8)(uint8(1)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'direction' field"))
	}
	m.Direction = direction

	if closeErr := readBuffer.CloseContext("BACnetAssignedLandingCallsLandingCallsListEntry"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetAssignedLandingCallsLandingCallsListEntry")
	}

	return m, nil
}

func (m *_BACnetAssignedLandingCallsLandingCallsListEntry) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetAssignedLandingCallsLandingCallsListEntry) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetAssignedLandingCallsLandingCallsListEntry"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetAssignedLandingCallsLandingCallsListEntry")
	}

	if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "floorNumber", m.GetFloorNumber(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'floorNumber' field")
	}

	if err := WriteSimpleField[BACnetLiftCarDirectionTagged](ctx, "direction", m.GetDirection(), WriteComplex[BACnetLiftCarDirectionTagged](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'direction' field")
	}

	if popErr := writeBuffer.PopContext("BACnetAssignedLandingCallsLandingCallsListEntry"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetAssignedLandingCallsLandingCallsListEntry")
	}
	return nil
}

func (m *_BACnetAssignedLandingCallsLandingCallsListEntry) IsBACnetAssignedLandingCallsLandingCallsListEntry() {
}

func (m *_BACnetAssignedLandingCallsLandingCallsListEntry) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetAssignedLandingCallsLandingCallsListEntry) deepCopy() *_BACnetAssignedLandingCallsLandingCallsListEntry {
	if m == nil {
		return nil
	}
	_BACnetAssignedLandingCallsLandingCallsListEntryCopy := &_BACnetAssignedLandingCallsLandingCallsListEntry{
		m.FloorNumber.DeepCopy().(BACnetContextTagUnsignedInteger),
		m.Direction.DeepCopy().(BACnetLiftCarDirectionTagged),
	}
	return _BACnetAssignedLandingCallsLandingCallsListEntryCopy
}

func (m *_BACnetAssignedLandingCallsLandingCallsListEntry) String() string {
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
