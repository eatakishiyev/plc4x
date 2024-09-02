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

// BACnetEscalatorOperationDirectionTagged is the corresponding interface of BACnetEscalatorOperationDirectionTagged
type BACnetEscalatorOperationDirectionTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetEscalatorOperationDirection
	// GetProprietaryValue returns ProprietaryValue (property field)
	GetProprietaryValue() uint32
	// GetIsProprietary returns IsProprietary (virtual field)
	GetIsProprietary() bool
}

// BACnetEscalatorOperationDirectionTaggedExactly can be used when we want exactly this type and not a type which fulfills BACnetEscalatorOperationDirectionTagged.
// This is useful for switch cases.
type BACnetEscalatorOperationDirectionTaggedExactly interface {
	BACnetEscalatorOperationDirectionTagged
	isBACnetEscalatorOperationDirectionTagged() bool
}

// _BACnetEscalatorOperationDirectionTagged is the data-structure of this message
type _BACnetEscalatorOperationDirectionTagged struct {
	Header           BACnetTagHeader
	Value            BACnetEscalatorOperationDirection
	ProprietaryValue uint32

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

var _ BACnetEscalatorOperationDirectionTagged = (*_BACnetEscalatorOperationDirectionTagged)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEscalatorOperationDirectionTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetEscalatorOperationDirectionTagged) GetValue() BACnetEscalatorOperationDirection {
	return m.Value
}

func (m *_BACnetEscalatorOperationDirectionTagged) GetProprietaryValue() uint32 {
	return m.ProprietaryValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetEscalatorOperationDirectionTagged) GetIsProprietary() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetValue()) == (BACnetEscalatorOperationDirection_VENDOR_PROPRIETARY_VALUE)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetEscalatorOperationDirectionTagged factory function for _BACnetEscalatorOperationDirectionTagged
func NewBACnetEscalatorOperationDirectionTagged(header BACnetTagHeader, value BACnetEscalatorOperationDirection, proprietaryValue uint32, tagNumber uint8, tagClass TagClass) *_BACnetEscalatorOperationDirectionTagged {
	return &_BACnetEscalatorOperationDirectionTagged{Header: header, Value: value, ProprietaryValue: proprietaryValue, TagNumber: tagNumber, TagClass: tagClass}
}

// Deprecated: use the interface for direct cast
func CastBACnetEscalatorOperationDirectionTagged(structType any) BACnetEscalatorOperationDirectionTagged {
	if casted, ok := structType.(BACnetEscalatorOperationDirectionTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEscalatorOperationDirectionTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEscalatorOperationDirectionTagged) GetTypeName() string {
	return "BACnetEscalatorOperationDirectionTagged"
}

func (m *_BACnetEscalatorOperationDirectionTagged) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits(ctx)

	// Manual Field (value)
	lengthInBits += uint16(utils.InlineIf(m.GetIsProprietary(), func() any { return int32(int32(0)) }, func() any { return int32((int32(m.GetHeader().GetActualLength()) * int32(int32(8)))) }).(int32))

	// A virtual field doesn't have any in- or output.

	// Manual Field (proprietaryValue)
	lengthInBits += uint16(utils.InlineIf(m.GetIsProprietary(), func() any { return int32((int32(m.GetHeader().GetActualLength()) * int32(int32(8)))) }, func() any { return int32(int32(0)) }).(int32))

	return lengthInBits
}

func (m *_BACnetEscalatorOperationDirectionTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetEscalatorOperationDirectionTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetEscalatorOperationDirectionTagged, error) {
	return BACnetEscalatorOperationDirectionTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetEscalatorOperationDirectionTaggedParseWithBufferProducer(tagNumber uint8, tagClass TagClass) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetEscalatorOperationDirectionTagged, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetEscalatorOperationDirectionTagged, error) {
		return BACnetEscalatorOperationDirectionTaggedParseWithBuffer(ctx, readBuffer, tagNumber, tagClass)
	}
}

func BACnetEscalatorOperationDirectionTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetEscalatorOperationDirectionTagged, error) {
	v, err := NewBACnetEscalatorOperationDirectionTagged(tagNumber, tagClass).parse(ctx, readBuffer, tagNumber, tagClass)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_BACnetEscalatorOperationDirectionTagged) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (__bACnetEscalatorOperationDirectionTagged BACnetEscalatorOperationDirectionTagged, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEscalatorOperationDirectionTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEscalatorOperationDirectionTagged")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	header, err := ReadSimpleField[BACnetTagHeader](ctx, "header", ReadComplex[BACnetTagHeader](BACnetTagHeaderParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'header' field"))
	}
	m.Header = header

	// Validation
	if !(bool((header.GetTagClass()) == (tagClass))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "tag class doesn't match"})
	}

	// Validation
	if !(bool((bool((header.GetTagClass()) == (TagClass_APPLICATION_TAGS)))) || bool((bool((header.GetActualTagNumber()) == (tagNumber))))) {
		return nil, errors.WithStack(utils.ParseAssertError{Message: "tagnumber doesn't match"})
	}

	value, err := ReadManualField[BACnetEscalatorOperationDirection](ctx, "value", readBuffer, EnsureType[BACnetEscalatorOperationDirection](ReadEnumGeneric(ctx, readBuffer, header.GetActualLength(), BACnetEscalatorOperationDirection_VENDOR_PROPRIETARY_VALUE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	isProprietary, err := ReadVirtualField[bool](ctx, "isProprietary", (*bool)(nil), bool((value) == (BACnetEscalatorOperationDirection_VENDOR_PROPRIETARY_VALUE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isProprietary' field"))
	}
	_ = isProprietary

	proprietaryValue, err := ReadManualField[uint32](ctx, "proprietaryValue", readBuffer, EnsureType[uint32](ReadProprietaryEnumGeneric(ctx, readBuffer, header.GetActualLength(), isProprietary)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'proprietaryValue' field"))
	}
	m.ProprietaryValue = proprietaryValue

	if closeErr := readBuffer.CloseContext("BACnetEscalatorOperationDirectionTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEscalatorOperationDirectionTagged")
	}

	return m, nil
}

func (m *_BACnetEscalatorOperationDirectionTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetEscalatorOperationDirectionTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetEscalatorOperationDirectionTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetEscalatorOperationDirectionTagged")
	}

	if err := WriteSimpleField[BACnetTagHeader](ctx, "header", m.GetHeader(), WriteComplex[BACnetTagHeader](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'header' field")
	}

	if err := WriteManualField[BACnetEscalatorOperationDirection](ctx, "value", func(ctx context.Context) error { return WriteEnumGeneric(ctx, writeBuffer, m.GetValue()) }, writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'value' field")
	}
	// Virtual field
	isProprietary := m.GetIsProprietary()
	_ = isProprietary
	if _isProprietaryErr := writeBuffer.WriteVirtual(ctx, "isProprietary", m.GetIsProprietary()); _isProprietaryErr != nil {
		return errors.Wrap(_isProprietaryErr, "Error serializing 'isProprietary' field")
	}

	if err := WriteManualField[uint32](ctx, "proprietaryValue", func(ctx context.Context) error {
		return WriteProprietaryEnumGeneric(ctx, writeBuffer, m.GetProprietaryValue(), m.GetIsProprietary())
	}, writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'proprietaryValue' field")
	}

	if popErr := writeBuffer.PopContext("BACnetEscalatorOperationDirectionTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetEscalatorOperationDirectionTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetEscalatorOperationDirectionTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetEscalatorOperationDirectionTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetEscalatorOperationDirectionTagged) isBACnetEscalatorOperationDirectionTagged() bool {
	return true
}

func (m *_BACnetEscalatorOperationDirectionTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
