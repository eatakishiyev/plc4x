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

// LevelInformation is the corresponding interface of LevelInformation
type LevelInformation interface {
	LevelInformationContract
	LevelInformationRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// LevelInformationContract provides a set of functions which can be overwritten by a sub struct
type LevelInformationContract interface {
	// GetRaw returns Raw (property field)
	GetRaw() uint16
	// GetNibble1 returns Nibble1 (virtual field)
	GetNibble1() uint8
	// GetNibble2 returns Nibble2 (virtual field)
	GetNibble2() uint8
	// GetNibble3 returns Nibble3 (virtual field)
	GetNibble3() uint8
	// GetNibble4 returns Nibble4 (virtual field)
	GetNibble4() uint8
	// GetIsAbsent returns IsAbsent (virtual field)
	GetIsAbsent() bool
	// GetIsCorruptedByNoise returns IsCorruptedByNoise (virtual field)
	GetIsCorruptedByNoise() bool
	// GetIsCorruptedByNoiseOrLevelsDiffer returns IsCorruptedByNoiseOrLevelsDiffer (virtual field)
	GetIsCorruptedByNoiseOrLevelsDiffer() bool
	// GetIsCorrupted returns IsCorrupted (virtual field)
	GetIsCorrupted() bool
}

// LevelInformationRequirements provides a set of functions which need to be implemented by a sub struct
type LevelInformationRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetIsAbsent returns IsAbsent (discriminator field)
	GetIsAbsent() bool
	// GetIsCorrupted returns IsCorrupted (discriminator field)
	GetIsCorrupted() bool
}

// LevelInformationExactly can be used when we want exactly this type and not a type which fulfills LevelInformation.
// This is useful for switch cases.
type LevelInformationExactly interface {
	LevelInformation
	isLevelInformation() bool
}

// _LevelInformation is the data-structure of this message
type _LevelInformation struct {
	_SubType LevelInformation
	Raw      uint16
}

var _ LevelInformationContract = (*_LevelInformation)(nil)

type LevelInformationChild interface {
	utils.Serializable

	GetParent() *LevelInformation

	GetTypeName() string
	LevelInformation
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_LevelInformation) GetRaw() uint16 {
	return m.Raw
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (pm *_LevelInformation) GetNibble1() uint8 {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	return uint8((m.GetRaw() & 0xF000) >> uint8(12))
}

func (pm *_LevelInformation) GetNibble2() uint8 {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	return uint8((m.GetRaw() & 0x0F00) >> uint8(8))
}

func (pm *_LevelInformation) GetNibble3() uint8 {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	return uint8((m.GetRaw() & 0x00F0) >> uint8(4))
}

func (pm *_LevelInformation) GetNibble4() uint8 {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	return uint8((m.GetRaw() & 0x000F) >> uint8(0))
}

func (pm *_LevelInformation) GetIsAbsent() bool {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	return bool(bool(bool(bool(bool((m.GetNibble1()) == (0x0))) && bool(bool((m.GetNibble2()) == (0x0)))) && bool(bool((m.GetNibble3()) == (0x0)))) && bool(bool((m.GetNibble4()) == (0x0))))
}

func (pm *_LevelInformation) GetIsCorruptedByNoise() bool {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	return bool(bool(!(m.GetIsAbsent())) && bool((bool(bool(bool((bool(bool((bool((m.GetNibble1()) < (0x5)))) || bool((bool((m.GetNibble1()) == (0x8))))) || bool((bool((m.GetNibble1()) == (0xC)))))) || bool((bool(bool((bool((m.GetNibble2()) < (0x5)))) || bool((bool((m.GetNibble2()) == (0x8))))) || bool((bool((m.GetNibble2()) == (0xC))))))) || bool((bool(bool((bool((m.GetNibble3()) < (0x5)))) || bool((bool((m.GetNibble3()) == (0x8))))) || bool((bool((m.GetNibble3()) == (0xC))))))) || bool((bool(bool((bool((m.GetNibble4()) < (0x5)))) || bool((bool((m.GetNibble4()) == (0x8))))) || bool((bool((m.GetNibble4()) == (0xC)))))))))
}

func (pm *_LevelInformation) GetIsCorruptedByNoiseOrLevelsDiffer() bool {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	return bool(bool(!(m.GetIsAbsent())) && bool((bool(bool(bool((bool(bool((bool((m.GetNibble1()) == (0x7)))) || bool((bool((m.GetNibble1()) == (0xB))))) || bool((bool((m.GetNibble1()) > (0xC)))))) || bool((bool(bool((bool((m.GetNibble2()) == (0x7)))) || bool((bool((m.GetNibble2()) == (0xB))))) || bool((bool((m.GetNibble2()) > (0xC))))))) || bool((bool(bool((bool((m.GetNibble3()) == (0x7)))) || bool((bool((m.GetNibble3()) == (0xB))))) || bool((bool((m.GetNibble3()) > (0xC))))))) || bool((bool(bool((bool((m.GetNibble4()) == (0x7)))) || bool((bool((m.GetNibble4()) == (0xB))))) || bool((bool((m.GetNibble4()) > (0xC)))))))))
}

func (pm *_LevelInformation) GetIsCorrupted() bool {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	return bool(bool(m.GetIsCorruptedByNoise()) || bool(m.GetIsCorruptedByNoiseOrLevelsDiffer()))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewLevelInformation factory function for _LevelInformation
func NewLevelInformation(raw uint16) *_LevelInformation {
	return &_LevelInformation{Raw: raw}
}

// Deprecated: use the interface for direct cast
func CastLevelInformation(structType any) LevelInformation {
	if casted, ok := structType.(LevelInformation); ok {
		return casted
	}
	if casted, ok := structType.(*LevelInformation); ok {
		return *casted
	}
	return nil
}

func (m *_LevelInformation) GetTypeName() string {
	return "LevelInformation"
}

func (m *_LevelInformation) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_LevelInformation) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func LevelInformationParse[T LevelInformation](ctx context.Context, theBytes []byte) (T, error) {
	return LevelInformationParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes))
}

func LevelInformationParseWithBufferProducer[T LevelInformation]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := LevelInformationParseWithBuffer[T](ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, err
	}
}

func LevelInformationParseWithBuffer[T LevelInformation](ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	v, err := NewLevelInformation().parse(ctx, readBuffer)
	if err != nil {
		var zero T
		return zero, err
	}
	return v.(T), err
}

func (m *_LevelInformation) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__levelInformation LevelInformation, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("LevelInformation"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for LevelInformation")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	raw, err := ReadPeekField[uint16](ctx, "raw", ReadUnsignedShort(readBuffer, uint8(16)), 0)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'raw' field"))
	}
	m.Raw = raw

	nibble1, err := ReadVirtualField[uint8](ctx, "nibble1", (*uint8)(nil), (raw&0xF000)>>uint8(12))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'nibble1' field"))
	}
	_ = nibble1

	nibble2, err := ReadVirtualField[uint8](ctx, "nibble2", (*uint8)(nil), (raw&0x0F00)>>uint8(8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'nibble2' field"))
	}
	_ = nibble2

	nibble3, err := ReadVirtualField[uint8](ctx, "nibble3", (*uint8)(nil), (raw&0x00F0)>>uint8(4))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'nibble3' field"))
	}
	_ = nibble3

	nibble4, err := ReadVirtualField[uint8](ctx, "nibble4", (*uint8)(nil), (raw&0x000F)>>uint8(0))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'nibble4' field"))
	}
	_ = nibble4

	isAbsent, err := ReadVirtualField[bool](ctx, "isAbsent", (*bool)(nil), bool(bool(bool(bool((nibble1) == (0x0))) && bool(bool((nibble2) == (0x0)))) && bool(bool((nibble3) == (0x0)))) && bool(bool((nibble4) == (0x0))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isAbsent' field"))
	}
	_ = isAbsent

	isCorruptedByNoise, err := ReadVirtualField[bool](ctx, "isCorruptedByNoise", (*bool)(nil), bool(!(isAbsent)) && bool((bool(bool(bool((bool(bool((bool((nibble1) < (0x5)))) || bool((bool((nibble1) == (0x8))))) || bool((bool((nibble1) == (0xC)))))) || bool((bool(bool((bool((nibble2) < (0x5)))) || bool((bool((nibble2) == (0x8))))) || bool((bool((nibble2) == (0xC))))))) || bool((bool(bool((bool((nibble3) < (0x5)))) || bool((bool((nibble3) == (0x8))))) || bool((bool((nibble3) == (0xC))))))) || bool((bool(bool((bool((nibble4) < (0x5)))) || bool((bool((nibble4) == (0x8))))) || bool((bool((nibble4) == (0xC)))))))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isCorruptedByNoise' field"))
	}
	_ = isCorruptedByNoise

	isCorruptedByNoiseOrLevelsDiffer, err := ReadVirtualField[bool](ctx, "isCorruptedByNoiseOrLevelsDiffer", (*bool)(nil), bool(!(isAbsent)) && bool((bool(bool(bool((bool(bool((bool((nibble1) == (0x7)))) || bool((bool((nibble1) == (0xB))))) || bool((bool((nibble1) > (0xC)))))) || bool((bool(bool((bool((nibble2) == (0x7)))) || bool((bool((nibble2) == (0xB))))) || bool((bool((nibble2) > (0xC))))))) || bool((bool(bool((bool((nibble3) == (0x7)))) || bool((bool((nibble3) == (0xB))))) || bool((bool((nibble3) > (0xC))))))) || bool((bool(bool((bool((nibble4) == (0x7)))) || bool((bool((nibble4) == (0xB))))) || bool((bool((nibble4) > (0xC)))))))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isCorruptedByNoiseOrLevelsDiffer' field"))
	}
	_ = isCorruptedByNoiseOrLevelsDiffer

	isCorrupted, err := ReadVirtualField[bool](ctx, "isCorrupted", (*bool)(nil), bool(isCorruptedByNoise) || bool(isCorruptedByNoiseOrLevelsDiffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isCorrupted' field"))
	}
	_ = isCorrupted

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child LevelInformation
	switch {
	case isAbsent == bool(true): // LevelInformationAbsent
		if _child, err = (&_LevelInformationAbsent{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type LevelInformationAbsent for type-switch of LevelInformation")
		}
	case 0 == 0 && isCorrupted == bool(true): // LevelInformationCorrupted
		if _child, err = (&_LevelInformationCorrupted{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type LevelInformationCorrupted for type-switch of LevelInformation")
		}
	case 0 == 0: // LevelInformationNormal
		if _child, err = (&_LevelInformationNormal{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type LevelInformationNormal for type-switch of LevelInformation")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [isAbsent=%v, isCorrupted=%v]", isAbsent, isCorrupted)
	}

	if closeErr := readBuffer.CloseContext("LevelInformation"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for LevelInformation")
	}

	return _child, nil
}

func (pm *_LevelInformation) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child LevelInformation, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("LevelInformation"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for LevelInformation")
	}
	// Virtual field
	nibble1 := m.GetNibble1()
	_ = nibble1
	if _nibble1Err := writeBuffer.WriteVirtual(ctx, "nibble1", m.GetNibble1()); _nibble1Err != nil {
		return errors.Wrap(_nibble1Err, "Error serializing 'nibble1' field")
	}
	// Virtual field
	nibble2 := m.GetNibble2()
	_ = nibble2
	if _nibble2Err := writeBuffer.WriteVirtual(ctx, "nibble2", m.GetNibble2()); _nibble2Err != nil {
		return errors.Wrap(_nibble2Err, "Error serializing 'nibble2' field")
	}
	// Virtual field
	nibble3 := m.GetNibble3()
	_ = nibble3
	if _nibble3Err := writeBuffer.WriteVirtual(ctx, "nibble3", m.GetNibble3()); _nibble3Err != nil {
		return errors.Wrap(_nibble3Err, "Error serializing 'nibble3' field")
	}
	// Virtual field
	nibble4 := m.GetNibble4()
	_ = nibble4
	if _nibble4Err := writeBuffer.WriteVirtual(ctx, "nibble4", m.GetNibble4()); _nibble4Err != nil {
		return errors.Wrap(_nibble4Err, "Error serializing 'nibble4' field")
	}
	// Virtual field
	isAbsent := m.GetIsAbsent()
	_ = isAbsent
	if _isAbsentErr := writeBuffer.WriteVirtual(ctx, "isAbsent", m.GetIsAbsent()); _isAbsentErr != nil {
		return errors.Wrap(_isAbsentErr, "Error serializing 'isAbsent' field")
	}
	// Virtual field
	isCorruptedByNoise := m.GetIsCorruptedByNoise()
	_ = isCorruptedByNoise
	if _isCorruptedByNoiseErr := writeBuffer.WriteVirtual(ctx, "isCorruptedByNoise", m.GetIsCorruptedByNoise()); _isCorruptedByNoiseErr != nil {
		return errors.Wrap(_isCorruptedByNoiseErr, "Error serializing 'isCorruptedByNoise' field")
	}
	// Virtual field
	isCorruptedByNoiseOrLevelsDiffer := m.GetIsCorruptedByNoiseOrLevelsDiffer()
	_ = isCorruptedByNoiseOrLevelsDiffer
	if _isCorruptedByNoiseOrLevelsDifferErr := writeBuffer.WriteVirtual(ctx, "isCorruptedByNoiseOrLevelsDiffer", m.GetIsCorruptedByNoiseOrLevelsDiffer()); _isCorruptedByNoiseOrLevelsDifferErr != nil {
		return errors.Wrap(_isCorruptedByNoiseOrLevelsDifferErr, "Error serializing 'isCorruptedByNoiseOrLevelsDiffer' field")
	}
	// Virtual field
	isCorrupted := m.GetIsCorrupted()
	_ = isCorrupted
	if _isCorruptedErr := writeBuffer.WriteVirtual(ctx, "isCorrupted", m.GetIsCorrupted()); _isCorruptedErr != nil {
		return errors.Wrap(_isCorruptedErr, "Error serializing 'isCorrupted' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("LevelInformation"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for LevelInformation")
	}
	return nil
}

func (m *_LevelInformation) isLevelInformation() bool {
	return true
}
