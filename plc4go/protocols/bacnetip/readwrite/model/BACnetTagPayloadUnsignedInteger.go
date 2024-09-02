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

// BACnetTagPayloadUnsignedInteger is the corresponding interface of BACnetTagPayloadUnsignedInteger
type BACnetTagPayloadUnsignedInteger interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetValueUint8 returns ValueUint8 (property field)
	GetValueUint8() *uint8
	// GetValueUint16 returns ValueUint16 (property field)
	GetValueUint16() *uint16
	// GetValueUint24 returns ValueUint24 (property field)
	GetValueUint24() *uint32
	// GetValueUint32 returns ValueUint32 (property field)
	GetValueUint32() *uint32
	// GetValueUint40 returns ValueUint40 (property field)
	GetValueUint40() *uint64
	// GetValueUint48 returns ValueUint48 (property field)
	GetValueUint48() *uint64
	// GetValueUint56 returns ValueUint56 (property field)
	GetValueUint56() *uint64
	// GetValueUint64 returns ValueUint64 (property field)
	GetValueUint64() *uint64
	// GetIsUint8 returns IsUint8 (virtual field)
	GetIsUint8() bool
	// GetIsUint16 returns IsUint16 (virtual field)
	GetIsUint16() bool
	// GetIsUint24 returns IsUint24 (virtual field)
	GetIsUint24() bool
	// GetIsUint32 returns IsUint32 (virtual field)
	GetIsUint32() bool
	// GetIsUint40 returns IsUint40 (virtual field)
	GetIsUint40() bool
	// GetIsUint48 returns IsUint48 (virtual field)
	GetIsUint48() bool
	// GetIsUint56 returns IsUint56 (virtual field)
	GetIsUint56() bool
	// GetIsUint64 returns IsUint64 (virtual field)
	GetIsUint64() bool
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() uint64
}

// BACnetTagPayloadUnsignedIntegerExactly can be used when we want exactly this type and not a type which fulfills BACnetTagPayloadUnsignedInteger.
// This is useful for switch cases.
type BACnetTagPayloadUnsignedIntegerExactly interface {
	BACnetTagPayloadUnsignedInteger
	isBACnetTagPayloadUnsignedInteger() bool
}

// _BACnetTagPayloadUnsignedInteger is the data-structure of this message
type _BACnetTagPayloadUnsignedInteger struct {
	ValueUint8  *uint8
	ValueUint16 *uint16
	ValueUint24 *uint32
	ValueUint32 *uint32
	ValueUint40 *uint64
	ValueUint48 *uint64
	ValueUint56 *uint64
	ValueUint64 *uint64

	// Arguments.
	ActualLength uint32
}

var _ BACnetTagPayloadUnsignedInteger = (*_BACnetTagPayloadUnsignedInteger)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetTagPayloadUnsignedInteger) GetValueUint8() *uint8 {
	return m.ValueUint8
}

func (m *_BACnetTagPayloadUnsignedInteger) GetValueUint16() *uint16 {
	return m.ValueUint16
}

func (m *_BACnetTagPayloadUnsignedInteger) GetValueUint24() *uint32 {
	return m.ValueUint24
}

func (m *_BACnetTagPayloadUnsignedInteger) GetValueUint32() *uint32 {
	return m.ValueUint32
}

func (m *_BACnetTagPayloadUnsignedInteger) GetValueUint40() *uint64 {
	return m.ValueUint40
}

func (m *_BACnetTagPayloadUnsignedInteger) GetValueUint48() *uint64 {
	return m.ValueUint48
}

func (m *_BACnetTagPayloadUnsignedInteger) GetValueUint56() *uint64 {
	return m.ValueUint56
}

func (m *_BACnetTagPayloadUnsignedInteger) GetValueUint64() *uint64 {
	return m.ValueUint64
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetTagPayloadUnsignedInteger) GetIsUint8() bool {
	ctx := context.Background()
	_ = ctx
	valueUint8 := m.GetValueUint8()
	_ = valueUint8
	valueUint16 := m.GetValueUint16()
	_ = valueUint16
	valueUint24 := m.GetValueUint24()
	_ = valueUint24
	valueUint32 := m.GetValueUint32()
	_ = valueUint32
	valueUint40 := m.GetValueUint40()
	_ = valueUint40
	valueUint48 := m.GetValueUint48()
	_ = valueUint48
	valueUint56 := m.GetValueUint56()
	_ = valueUint56
	valueUint64 := m.GetValueUint64()
	_ = valueUint64
	return bool(bool((m.GetActualLength()) == (1)))
}

func (m *_BACnetTagPayloadUnsignedInteger) GetIsUint16() bool {
	ctx := context.Background()
	_ = ctx
	valueUint8 := m.GetValueUint8()
	_ = valueUint8
	valueUint16 := m.GetValueUint16()
	_ = valueUint16
	valueUint24 := m.GetValueUint24()
	_ = valueUint24
	valueUint32 := m.GetValueUint32()
	_ = valueUint32
	valueUint40 := m.GetValueUint40()
	_ = valueUint40
	valueUint48 := m.GetValueUint48()
	_ = valueUint48
	valueUint56 := m.GetValueUint56()
	_ = valueUint56
	valueUint64 := m.GetValueUint64()
	_ = valueUint64
	return bool(bool((m.GetActualLength()) == (2)))
}

func (m *_BACnetTagPayloadUnsignedInteger) GetIsUint24() bool {
	ctx := context.Background()
	_ = ctx
	valueUint8 := m.GetValueUint8()
	_ = valueUint8
	valueUint16 := m.GetValueUint16()
	_ = valueUint16
	valueUint24 := m.GetValueUint24()
	_ = valueUint24
	valueUint32 := m.GetValueUint32()
	_ = valueUint32
	valueUint40 := m.GetValueUint40()
	_ = valueUint40
	valueUint48 := m.GetValueUint48()
	_ = valueUint48
	valueUint56 := m.GetValueUint56()
	_ = valueUint56
	valueUint64 := m.GetValueUint64()
	_ = valueUint64
	return bool(bool((m.GetActualLength()) == (3)))
}

func (m *_BACnetTagPayloadUnsignedInteger) GetIsUint32() bool {
	ctx := context.Background()
	_ = ctx
	valueUint8 := m.GetValueUint8()
	_ = valueUint8
	valueUint16 := m.GetValueUint16()
	_ = valueUint16
	valueUint24 := m.GetValueUint24()
	_ = valueUint24
	valueUint32 := m.GetValueUint32()
	_ = valueUint32
	valueUint40 := m.GetValueUint40()
	_ = valueUint40
	valueUint48 := m.GetValueUint48()
	_ = valueUint48
	valueUint56 := m.GetValueUint56()
	_ = valueUint56
	valueUint64 := m.GetValueUint64()
	_ = valueUint64
	return bool(bool((m.GetActualLength()) == (4)))
}

func (m *_BACnetTagPayloadUnsignedInteger) GetIsUint40() bool {
	ctx := context.Background()
	_ = ctx
	valueUint8 := m.GetValueUint8()
	_ = valueUint8
	valueUint16 := m.GetValueUint16()
	_ = valueUint16
	valueUint24 := m.GetValueUint24()
	_ = valueUint24
	valueUint32 := m.GetValueUint32()
	_ = valueUint32
	valueUint40 := m.GetValueUint40()
	_ = valueUint40
	valueUint48 := m.GetValueUint48()
	_ = valueUint48
	valueUint56 := m.GetValueUint56()
	_ = valueUint56
	valueUint64 := m.GetValueUint64()
	_ = valueUint64
	return bool(bool((m.GetActualLength()) == (5)))
}

func (m *_BACnetTagPayloadUnsignedInteger) GetIsUint48() bool {
	ctx := context.Background()
	_ = ctx
	valueUint8 := m.GetValueUint8()
	_ = valueUint8
	valueUint16 := m.GetValueUint16()
	_ = valueUint16
	valueUint24 := m.GetValueUint24()
	_ = valueUint24
	valueUint32 := m.GetValueUint32()
	_ = valueUint32
	valueUint40 := m.GetValueUint40()
	_ = valueUint40
	valueUint48 := m.GetValueUint48()
	_ = valueUint48
	valueUint56 := m.GetValueUint56()
	_ = valueUint56
	valueUint64 := m.GetValueUint64()
	_ = valueUint64
	return bool(bool((m.GetActualLength()) == (6)))
}

func (m *_BACnetTagPayloadUnsignedInteger) GetIsUint56() bool {
	ctx := context.Background()
	_ = ctx
	valueUint8 := m.GetValueUint8()
	_ = valueUint8
	valueUint16 := m.GetValueUint16()
	_ = valueUint16
	valueUint24 := m.GetValueUint24()
	_ = valueUint24
	valueUint32 := m.GetValueUint32()
	_ = valueUint32
	valueUint40 := m.GetValueUint40()
	_ = valueUint40
	valueUint48 := m.GetValueUint48()
	_ = valueUint48
	valueUint56 := m.GetValueUint56()
	_ = valueUint56
	valueUint64 := m.GetValueUint64()
	_ = valueUint64
	return bool(bool((m.GetActualLength()) == (7)))
}

func (m *_BACnetTagPayloadUnsignedInteger) GetIsUint64() bool {
	ctx := context.Background()
	_ = ctx
	valueUint8 := m.GetValueUint8()
	_ = valueUint8
	valueUint16 := m.GetValueUint16()
	_ = valueUint16
	valueUint24 := m.GetValueUint24()
	_ = valueUint24
	valueUint32 := m.GetValueUint32()
	_ = valueUint32
	valueUint40 := m.GetValueUint40()
	_ = valueUint40
	valueUint48 := m.GetValueUint48()
	_ = valueUint48
	valueUint56 := m.GetValueUint56()
	_ = valueUint56
	valueUint64 := m.GetValueUint64()
	_ = valueUint64
	return bool(bool((m.GetActualLength()) == (8)))
}

func (m *_BACnetTagPayloadUnsignedInteger) GetActualValue() uint64 {
	ctx := context.Background()
	_ = ctx
	valueUint8 := m.GetValueUint8()
	_ = valueUint8
	valueUint16 := m.GetValueUint16()
	_ = valueUint16
	valueUint24 := m.GetValueUint24()
	_ = valueUint24
	valueUint32 := m.GetValueUint32()
	_ = valueUint32
	valueUint40 := m.GetValueUint40()
	_ = valueUint40
	valueUint48 := m.GetValueUint48()
	_ = valueUint48
	valueUint56 := m.GetValueUint56()
	_ = valueUint56
	valueUint64 := m.GetValueUint64()
	_ = valueUint64
	return uint64(utils.InlineIf(m.GetIsUint8(), func() any { return uint64((*m.GetValueUint8())) }, func() any {
		return uint64((utils.InlineIf(m.GetIsUint16(), func() any { return uint64((*m.GetValueUint16())) }, func() any {
			return uint64((utils.InlineIf(m.GetIsUint24(), func() any { return uint64((*m.GetValueUint24())) }, func() any {
				return uint64((utils.InlineIf(m.GetIsUint32(), func() any { return uint64((*m.GetValueUint32())) }, func() any {
					return uint64((utils.InlineIf(m.GetIsUint40(), func() any { return uint64((*m.GetValueUint40())) }, func() any {
						return uint64((utils.InlineIf(m.GetIsUint48(), func() any { return uint64((*m.GetValueUint48())) }, func() any {
							return uint64((utils.InlineIf(m.GetIsUint56(), func() any { return uint64((*m.GetValueUint56())) }, func() any { return uint64((*m.GetValueUint64())) }).(uint64)))
						}).(uint64)))
					}).(uint64)))
				}).(uint64)))
			}).(uint64)))
		}).(uint64)))
	}).(uint64))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetTagPayloadUnsignedInteger factory function for _BACnetTagPayloadUnsignedInteger
func NewBACnetTagPayloadUnsignedInteger(valueUint8 *uint8, valueUint16 *uint16, valueUint24 *uint32, valueUint32 *uint32, valueUint40 *uint64, valueUint48 *uint64, valueUint56 *uint64, valueUint64 *uint64, actualLength uint32) *_BACnetTagPayloadUnsignedInteger {
	return &_BACnetTagPayloadUnsignedInteger{ValueUint8: valueUint8, ValueUint16: valueUint16, ValueUint24: valueUint24, ValueUint32: valueUint32, ValueUint40: valueUint40, ValueUint48: valueUint48, ValueUint56: valueUint56, ValueUint64: valueUint64, ActualLength: actualLength}
}

// Deprecated: use the interface for direct cast
func CastBACnetTagPayloadUnsignedInteger(structType any) BACnetTagPayloadUnsignedInteger {
	if casted, ok := structType.(BACnetTagPayloadUnsignedInteger); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetTagPayloadUnsignedInteger); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetTagPayloadUnsignedInteger) GetTypeName() string {
	return "BACnetTagPayloadUnsignedInteger"
}

func (m *_BACnetTagPayloadUnsignedInteger) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	// Optional Field (valueUint8)
	if m.ValueUint8 != nil {
		lengthInBits += 8
	}

	// A virtual field doesn't have any in- or output.

	// Optional Field (valueUint16)
	if m.ValueUint16 != nil {
		lengthInBits += 16
	}

	// A virtual field doesn't have any in- or output.

	// Optional Field (valueUint24)
	if m.ValueUint24 != nil {
		lengthInBits += 24
	}

	// A virtual field doesn't have any in- or output.

	// Optional Field (valueUint32)
	if m.ValueUint32 != nil {
		lengthInBits += 32
	}

	// A virtual field doesn't have any in- or output.

	// Optional Field (valueUint40)
	if m.ValueUint40 != nil {
		lengthInBits += 40
	}

	// A virtual field doesn't have any in- or output.

	// Optional Field (valueUint48)
	if m.ValueUint48 != nil {
		lengthInBits += 48
	}

	// A virtual field doesn't have any in- or output.

	// Optional Field (valueUint56)
	if m.ValueUint56 != nil {
		lengthInBits += 56
	}

	// A virtual field doesn't have any in- or output.

	// Optional Field (valueUint64)
	if m.ValueUint64 != nil {
		lengthInBits += 64
	}

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetTagPayloadUnsignedInteger) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetTagPayloadUnsignedIntegerParse(ctx context.Context, theBytes []byte, actualLength uint32) (BACnetTagPayloadUnsignedInteger, error) {
	return BACnetTagPayloadUnsignedIntegerParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), actualLength)
}

func BACnetTagPayloadUnsignedIntegerParseWithBufferProducer(actualLength uint32) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetTagPayloadUnsignedInteger, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetTagPayloadUnsignedInteger, error) {
		return BACnetTagPayloadUnsignedIntegerParseWithBuffer(ctx, readBuffer, actualLength)
	}
}

func BACnetTagPayloadUnsignedIntegerParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, actualLength uint32) (BACnetTagPayloadUnsignedInteger, error) {
	v, err := NewBACnetTagPayloadUnsignedInteger(actualLength).parse(ctx, readBuffer, actualLength)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_BACnetTagPayloadUnsignedInteger) parse(ctx context.Context, readBuffer utils.ReadBuffer, actualLength uint32) (__bACnetTagPayloadUnsignedInteger BACnetTagPayloadUnsignedInteger, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetTagPayloadUnsignedInteger"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetTagPayloadUnsignedInteger")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	isUint8, err := ReadVirtualField[bool](ctx, "isUint8", (*bool)(nil), bool((actualLength) == (1)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isUint8' field"))
	}
	_ = isUint8

	var valueUint8 *uint8
	valueUint8, err = ReadOptionalField[uint8](ctx, "valueUint8", ReadUnsignedByte(readBuffer, uint8(8)), isUint8)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'valueUint8' field"))
	}
	m.ValueUint8 = valueUint8

	isUint16, err := ReadVirtualField[bool](ctx, "isUint16", (*bool)(nil), bool((actualLength) == (2)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isUint16' field"))
	}
	_ = isUint16

	var valueUint16 *uint16
	valueUint16, err = ReadOptionalField[uint16](ctx, "valueUint16", ReadUnsignedShort(readBuffer, uint8(16)), isUint16)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'valueUint16' field"))
	}
	m.ValueUint16 = valueUint16

	isUint24, err := ReadVirtualField[bool](ctx, "isUint24", (*bool)(nil), bool((actualLength) == (3)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isUint24' field"))
	}
	_ = isUint24

	var valueUint24 *uint32
	valueUint24, err = ReadOptionalField[uint32](ctx, "valueUint24", ReadUnsignedInt(readBuffer, uint8(24)), isUint24)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'valueUint24' field"))
	}
	m.ValueUint24 = valueUint24

	isUint32, err := ReadVirtualField[bool](ctx, "isUint32", (*bool)(nil), bool((actualLength) == (4)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isUint32' field"))
	}
	_ = isUint32

	var valueUint32 *uint32
	valueUint32, err = ReadOptionalField[uint32](ctx, "valueUint32", ReadUnsignedInt(readBuffer, uint8(32)), isUint32)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'valueUint32' field"))
	}
	m.ValueUint32 = valueUint32

	isUint40, err := ReadVirtualField[bool](ctx, "isUint40", (*bool)(nil), bool((actualLength) == (5)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isUint40' field"))
	}
	_ = isUint40

	var valueUint40 *uint64
	valueUint40, err = ReadOptionalField[uint64](ctx, "valueUint40", ReadUnsignedLong(readBuffer, uint8(40)), isUint40)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'valueUint40' field"))
	}
	m.ValueUint40 = valueUint40

	isUint48, err := ReadVirtualField[bool](ctx, "isUint48", (*bool)(nil), bool((actualLength) == (6)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isUint48' field"))
	}
	_ = isUint48

	var valueUint48 *uint64
	valueUint48, err = ReadOptionalField[uint64](ctx, "valueUint48", ReadUnsignedLong(readBuffer, uint8(48)), isUint48)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'valueUint48' field"))
	}
	m.ValueUint48 = valueUint48

	isUint56, err := ReadVirtualField[bool](ctx, "isUint56", (*bool)(nil), bool((actualLength) == (7)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isUint56' field"))
	}
	_ = isUint56

	var valueUint56 *uint64
	valueUint56, err = ReadOptionalField[uint64](ctx, "valueUint56", ReadUnsignedLong(readBuffer, uint8(56)), isUint56)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'valueUint56' field"))
	}
	m.ValueUint56 = valueUint56

	isUint64, err := ReadVirtualField[bool](ctx, "isUint64", (*bool)(nil), bool((actualLength) == (8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isUint64' field"))
	}
	_ = isUint64

	var valueUint64 *uint64
	valueUint64, err = ReadOptionalField[uint64](ctx, "valueUint64", ReadUnsignedLong(readBuffer, uint8(64)), isUint64)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'valueUint64' field"))
	}
	m.ValueUint64 = valueUint64

	// Validation
	if !(bool(bool(bool(bool(bool(bool(bool(isUint8) || bool(isUint16)) || bool(isUint24)) || bool(isUint32)) || bool(isUint40)) || bool(isUint48)) || bool(isUint56)) || bool(isUint64)) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "unmapped integer length"})
	}

	actualValue, err := ReadVirtualField[uint64](ctx, "actualValue", (*uint64)(nil), utils.InlineIf(isUint8, func() any { return uint64((*valueUint8)) }, func() any {
		return uint64((utils.InlineIf(isUint16, func() any { return uint64((*valueUint16)) }, func() any {
			return uint64((utils.InlineIf(isUint24, func() any { return uint64((*valueUint24)) }, func() any {
				return uint64((utils.InlineIf(isUint32, func() any { return uint64((*valueUint32)) }, func() any {
					return uint64((utils.InlineIf(isUint40, func() any { return uint64((*valueUint40)) }, func() any {
						return uint64((utils.InlineIf(isUint48, func() any { return uint64((*valueUint48)) }, func() any {
							return uint64((utils.InlineIf(isUint56, func() any { return uint64((*valueUint56)) }, func() any { return uint64((*valueUint64)) }).(uint64)))
						}).(uint64)))
					}).(uint64)))
				}).(uint64)))
			}).(uint64)))
		}).(uint64)))
	}).(uint64))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetTagPayloadUnsignedInteger"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetTagPayloadUnsignedInteger")
	}

	return m, nil
}

func (m *_BACnetTagPayloadUnsignedInteger) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetTagPayloadUnsignedInteger) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetTagPayloadUnsignedInteger"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetTagPayloadUnsignedInteger")
	}
	// Virtual field
	isUint8 := m.GetIsUint8()
	_ = isUint8
	if _isUint8Err := writeBuffer.WriteVirtual(ctx, "isUint8", m.GetIsUint8()); _isUint8Err != nil {
		return errors.Wrap(_isUint8Err, "Error serializing 'isUint8' field")
	}

	if err := WriteOptionalField[uint8](ctx, "valueUint8", m.GetValueUint8(), WriteUnsignedByte(writeBuffer, 8), true); err != nil {
		return errors.Wrap(err, "Error serializing 'valueUint8' field")
	}
	// Virtual field
	isUint16 := m.GetIsUint16()
	_ = isUint16
	if _isUint16Err := writeBuffer.WriteVirtual(ctx, "isUint16", m.GetIsUint16()); _isUint16Err != nil {
		return errors.Wrap(_isUint16Err, "Error serializing 'isUint16' field")
	}

	if err := WriteOptionalField[uint16](ctx, "valueUint16", m.GetValueUint16(), WriteUnsignedShort(writeBuffer, 16), true); err != nil {
		return errors.Wrap(err, "Error serializing 'valueUint16' field")
	}
	// Virtual field
	isUint24 := m.GetIsUint24()
	_ = isUint24
	if _isUint24Err := writeBuffer.WriteVirtual(ctx, "isUint24", m.GetIsUint24()); _isUint24Err != nil {
		return errors.Wrap(_isUint24Err, "Error serializing 'isUint24' field")
	}

	if err := WriteOptionalField[uint32](ctx, "valueUint24", m.GetValueUint24(), WriteUnsignedInt(writeBuffer, 24), true); err != nil {
		return errors.Wrap(err, "Error serializing 'valueUint24' field")
	}
	// Virtual field
	isUint32 := m.GetIsUint32()
	_ = isUint32
	if _isUint32Err := writeBuffer.WriteVirtual(ctx, "isUint32", m.GetIsUint32()); _isUint32Err != nil {
		return errors.Wrap(_isUint32Err, "Error serializing 'isUint32' field")
	}

	if err := WriteOptionalField[uint32](ctx, "valueUint32", m.GetValueUint32(), WriteUnsignedInt(writeBuffer, 32), true); err != nil {
		return errors.Wrap(err, "Error serializing 'valueUint32' field")
	}
	// Virtual field
	isUint40 := m.GetIsUint40()
	_ = isUint40
	if _isUint40Err := writeBuffer.WriteVirtual(ctx, "isUint40", m.GetIsUint40()); _isUint40Err != nil {
		return errors.Wrap(_isUint40Err, "Error serializing 'isUint40' field")
	}

	if err := WriteOptionalField[uint64](ctx, "valueUint40", m.GetValueUint40(), WriteUnsignedLong(writeBuffer, 40), true); err != nil {
		return errors.Wrap(err, "Error serializing 'valueUint40' field")
	}
	// Virtual field
	isUint48 := m.GetIsUint48()
	_ = isUint48
	if _isUint48Err := writeBuffer.WriteVirtual(ctx, "isUint48", m.GetIsUint48()); _isUint48Err != nil {
		return errors.Wrap(_isUint48Err, "Error serializing 'isUint48' field")
	}

	if err := WriteOptionalField[uint64](ctx, "valueUint48", m.GetValueUint48(), WriteUnsignedLong(writeBuffer, 48), true); err != nil {
		return errors.Wrap(err, "Error serializing 'valueUint48' field")
	}
	// Virtual field
	isUint56 := m.GetIsUint56()
	_ = isUint56
	if _isUint56Err := writeBuffer.WriteVirtual(ctx, "isUint56", m.GetIsUint56()); _isUint56Err != nil {
		return errors.Wrap(_isUint56Err, "Error serializing 'isUint56' field")
	}

	if err := WriteOptionalField[uint64](ctx, "valueUint56", m.GetValueUint56(), WriteUnsignedLong(writeBuffer, 56), true); err != nil {
		return errors.Wrap(err, "Error serializing 'valueUint56' field")
	}
	// Virtual field
	isUint64 := m.GetIsUint64()
	_ = isUint64
	if _isUint64Err := writeBuffer.WriteVirtual(ctx, "isUint64", m.GetIsUint64()); _isUint64Err != nil {
		return errors.Wrap(_isUint64Err, "Error serializing 'isUint64' field")
	}

	if err := WriteOptionalField[uint64](ctx, "valueUint64", m.GetValueUint64(), WriteUnsignedLong(writeBuffer, 64), true); err != nil {
		return errors.Wrap(err, "Error serializing 'valueUint64' field")
	}
	// Virtual field
	actualValue := m.GetActualValue()
	_ = actualValue
	if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
		return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
	}

	if popErr := writeBuffer.PopContext("BACnetTagPayloadUnsignedInteger"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetTagPayloadUnsignedInteger")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetTagPayloadUnsignedInteger) GetActualLength() uint32 {
	return m.ActualLength
}

//
////

func (m *_BACnetTagPayloadUnsignedInteger) isBACnetTagPayloadUnsignedInteger() bool {
	return true
}

func (m *_BACnetTagPayloadUnsignedInteger) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
