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

// BACnetLogData is the corresponding interface of BACnetLogData
type BACnetLogData interface {
	BACnetLogDataContract
	BACnetLogDataRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// BACnetLogDataContract provides a set of functions which can be overwritten by a sub struct
type BACnetLogDataContract interface {
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
	// GetTagNumber() returns a parser argument
	GetTagNumber() uint8
}

// BACnetLogDataRequirements provides a set of functions which need to be implemented by a sub struct
type BACnetLogDataRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetPeekedTagNumber returns PeekedTagNumber (discriminator field)
	GetPeekedTagNumber() uint8
}

// BACnetLogDataExactly can be used when we want exactly this type and not a type which fulfills BACnetLogData.
// This is useful for switch cases.
type BACnetLogDataExactly interface {
	BACnetLogData
	isBACnetLogData() bool
}

// _BACnetLogData is the data-structure of this message
type _BACnetLogData struct {
	_SubType        BACnetLogData
	OpeningTag      BACnetOpeningTag
	PeekedTagHeader BACnetTagHeader
	ClosingTag      BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

var _ BACnetLogDataContract = (*_BACnetLogData)(nil)

type BACnetLogDataChild interface {
	utils.Serializable

	GetParent() *BACnetLogData

	GetTypeName() string
	BACnetLogData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLogData) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetLogData) GetPeekedTagHeader() BACnetTagHeader {
	return m.PeekedTagHeader
}

func (m *_BACnetLogData) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (pm *_BACnetLogData) GetPeekedTagNumber() uint8 {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetLogData factory function for _BACnetLogData
func NewBACnetLogData(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetLogData {
	return &_BACnetLogData{OpeningTag: openingTag, PeekedTagHeader: peekedTagHeader, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastBACnetLogData(structType any) BACnetLogData {
	if casted, ok := structType.(BACnetLogData); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLogData); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLogData) GetTypeName() string {
	return "BACnetLogData"
}

func (m *_BACnetLogData) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetLogData) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func BACnetLogDataParse[T BACnetLogData](ctx context.Context, theBytes []byte, tagNumber uint8) (T, error) {
	return BACnetLogDataParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetLogDataParseWithBufferProducer[T BACnetLogData](tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := BACnetLogDataParseWithBuffer[T](ctx, readBuffer, tagNumber)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, err
	}
}

func BACnetLogDataParseWithBuffer[T BACnetLogData](ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (T, error) {
	v, err := NewBACnetLogData(tagNumber).parse(ctx, readBuffer, tagNumber)
	if err != nil {
		var zero T
		return zero, err
	}
	return v.(T), err
}

func (m *_BACnetLogData) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (__bACnetLogData BACnetLogData, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLogData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLogData")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	peekedTagHeader, err := ReadPeekField[BACnetTagHeader](ctx, "peekedTagHeader", ReadComplex[BACnetTagHeader](BACnetTagHeaderParseWithBuffer, readBuffer), 0)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedTagHeader' field"))
	}
	m.PeekedTagHeader = peekedTagHeader

	peekedTagNumber, err := ReadVirtualField[uint8](ctx, "peekedTagNumber", (*uint8)(nil), peekedTagHeader.GetActualTagNumber())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedTagNumber' field"))
	}
	_ = peekedTagNumber

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child BACnetLogData
	switch {
	case peekedTagNumber == uint8(0): // BACnetLogDataLogStatus
		if _child, err = (&_BACnetLogDataLogStatus{}).parse(ctx, readBuffer, m, tagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetLogDataLogStatus for type-switch of BACnetLogData")
		}
	case peekedTagNumber == uint8(1): // BACnetLogDataLogData
		if _child, err = (&_BACnetLogDataLogData{}).parse(ctx, readBuffer, m, tagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetLogDataLogData for type-switch of BACnetLogData")
		}
	case peekedTagNumber == uint8(2): // BACnetLogDataLogDataTimeChange
		if _child, err = (&_BACnetLogDataLogDataTimeChange{}).parse(ctx, readBuffer, m, tagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetLogDataLogDataTimeChange for type-switch of BACnetLogData")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [peekedTagNumber=%v]", peekedTagNumber)
	}

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetLogData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLogData")
	}

	return _child, nil
}

func (pm *_BACnetLogData) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetLogData, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetLogData"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetLogData")
	}

	if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'openingTag' field")
	}
	// Virtual field
	peekedTagNumber := m.GetPeekedTagNumber()
	_ = peekedTagNumber
	if _peekedTagNumberErr := writeBuffer.WriteVirtual(ctx, "peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetLogData"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetLogData")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetLogData) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetLogData) isBACnetLogData() bool {
	return true
}
