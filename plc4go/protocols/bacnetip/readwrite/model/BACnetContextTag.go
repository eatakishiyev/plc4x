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

// BACnetContextTag is the corresponding interface of BACnetContextTag
type BACnetContextTag interface {
	BACnetContextTagContract
	BACnetContextTagRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// BACnetContextTagContract provides a set of functions which can be overwritten by a sub struct
type BACnetContextTagContract interface {
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetTagNumber returns TagNumber (virtual field)
	GetTagNumber() uint8
	// GetActualLength returns ActualLength (virtual field)
	GetActualLength() uint32
	// GetTagNumberArgument() returns a parser argument
	GetTagNumberArgument() uint8
}

// BACnetContextTagRequirements provides a set of functions which need to be implemented by a sub struct
type BACnetContextTagRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetDataType returns DataType (discriminator field)
	GetDataType() BACnetDataType
}

// BACnetContextTagExactly can be used when we want exactly this type and not a type which fulfills BACnetContextTag.
// This is useful for switch cases.
type BACnetContextTagExactly interface {
	BACnetContextTag
	isBACnetContextTag() bool
}

// _BACnetContextTag is the data-structure of this message
type _BACnetContextTag struct {
	_SubType BACnetContextTag
	Header   BACnetTagHeader

	// Arguments.
	TagNumberArgument uint8
}

var _ BACnetContextTagContract = (*_BACnetContextTag)(nil)

type BACnetContextTagChild interface {
	utils.Serializable

	GetParent() *BACnetContextTag

	GetTypeName() string
	BACnetContextTag
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetContextTag) GetHeader() BACnetTagHeader {
	return m.Header
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (pm *_BACnetContextTag) GetTagNumber() uint8 {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	return uint8(m.GetHeader().GetTagNumber())
}

func (pm *_BACnetContextTag) GetActualLength() uint32 {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	return uint32(m.GetHeader().GetActualLength())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetContextTag factory function for _BACnetContextTag
func NewBACnetContextTag(header BACnetTagHeader, tagNumberArgument uint8) *_BACnetContextTag {
	return &_BACnetContextTag{Header: header, TagNumberArgument: tagNumberArgument}
}

// Deprecated: use the interface for direct cast
func CastBACnetContextTag(structType any) BACnetContextTag {
	if casted, ok := structType.(BACnetContextTag); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetContextTag); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetContextTag) GetTypeName() string {
	return "BACnetContextTag"
}

func (m *_BACnetContextTag) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetContextTag) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func BACnetContextTagParse[T BACnetContextTag](ctx context.Context, theBytes []byte, tagNumberArgument uint8, dataType BACnetDataType) (T, error) {
	return BACnetContextTagParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes), tagNumberArgument, dataType)
}

func BACnetContextTagParseWithBufferProducer[T BACnetContextTag](tagNumberArgument uint8, dataType BACnetDataType) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := BACnetContextTagParseWithBuffer[T](ctx, readBuffer, tagNumberArgument, dataType)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, err
	}
}

func BACnetContextTagParseWithBuffer[T BACnetContextTag](ctx context.Context, readBuffer utils.ReadBuffer, tagNumberArgument uint8, dataType BACnetDataType) (T, error) {
	v, err := NewBACnetContextTag(tagNumberArgument).parse(ctx, readBuffer, tagNumberArgument, dataType)
	if err != nil {
		var zero T
		return zero, err
	}
	return v.(T), err
}

func (m *_BACnetContextTag) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumberArgument uint8, dataType BACnetDataType) (__bACnetContextTag BACnetContextTag, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetContextTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetContextTag")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	header, err := ReadSimpleField[BACnetTagHeader](ctx, "header", ReadComplex[BACnetTagHeader](BACnetTagHeaderParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'header' field"))
	}
	m.Header = header

	// Validation
	if !(bool((header.GetActualTagNumber()) == (tagNumberArgument))) {
		return nil, errors.WithStack(utils.ParseAssertError{Message: "tagnumber doesn't match"})
	}

	// Validation
	if !(bool((header.GetTagClass()) == (TagClass_CONTEXT_SPECIFIC_TAGS))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "should be a context tag"})
	}

	tagNumber, err := ReadVirtualField[uint8](ctx, "tagNumber", (*uint8)(nil), header.GetTagNumber())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'tagNumber' field"))
	}
	_ = tagNumber

	actualLength, err := ReadVirtualField[uint32](ctx, "actualLength", (*uint32)(nil), header.GetActualLength())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualLength' field"))
	}
	_ = actualLength

	// Validation
	if !(bool(bool((header.GetLengthValueType()) != (6))) && bool(bool((header.GetLengthValueType()) != (7)))) {
		return nil, errors.WithStack(utils.ParseAssertError{Message: "length 6 and 7 reserved for opening and closing tag"})
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child BACnetContextTag
	switch {
	case dataType == BACnetDataType_NULL: // BACnetContextTagNull
		if _child, err = (&_BACnetContextTagNull{}).parse(ctx, readBuffer, m, header, tagNumberArgument, dataType); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetContextTagNull for type-switch of BACnetContextTag")
		}
	case dataType == BACnetDataType_BOOLEAN: // BACnetContextTagBoolean
		if _child, err = (&_BACnetContextTagBoolean{}).parse(ctx, readBuffer, m, header, tagNumberArgument, dataType); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetContextTagBoolean for type-switch of BACnetContextTag")
		}
	case dataType == BACnetDataType_UNSIGNED_INTEGER: // BACnetContextTagUnsignedInteger
		if _child, err = (&_BACnetContextTagUnsignedInteger{}).parse(ctx, readBuffer, m, header, tagNumberArgument, dataType); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetContextTagUnsignedInteger for type-switch of BACnetContextTag")
		}
	case dataType == BACnetDataType_SIGNED_INTEGER: // BACnetContextTagSignedInteger
		if _child, err = (&_BACnetContextTagSignedInteger{}).parse(ctx, readBuffer, m, header, tagNumberArgument, dataType); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetContextTagSignedInteger for type-switch of BACnetContextTag")
		}
	case dataType == BACnetDataType_REAL: // BACnetContextTagReal
		if _child, err = (&_BACnetContextTagReal{}).parse(ctx, readBuffer, m, tagNumberArgument, dataType); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetContextTagReal for type-switch of BACnetContextTag")
		}
	case dataType == BACnetDataType_DOUBLE: // BACnetContextTagDouble
		if _child, err = (&_BACnetContextTagDouble{}).parse(ctx, readBuffer, m, tagNumberArgument, dataType); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetContextTagDouble for type-switch of BACnetContextTag")
		}
	case dataType == BACnetDataType_OCTET_STRING: // BACnetContextTagOctetString
		if _child, err = (&_BACnetContextTagOctetString{}).parse(ctx, readBuffer, m, header, tagNumberArgument, dataType); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetContextTagOctetString for type-switch of BACnetContextTag")
		}
	case dataType == BACnetDataType_CHARACTER_STRING: // BACnetContextTagCharacterString
		if _child, err = (&_BACnetContextTagCharacterString{}).parse(ctx, readBuffer, m, header, tagNumberArgument, dataType); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetContextTagCharacterString for type-switch of BACnetContextTag")
		}
	case dataType == BACnetDataType_BIT_STRING: // BACnetContextTagBitString
		if _child, err = (&_BACnetContextTagBitString{}).parse(ctx, readBuffer, m, header, tagNumberArgument, dataType); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetContextTagBitString for type-switch of BACnetContextTag")
		}
	case dataType == BACnetDataType_ENUMERATED: // BACnetContextTagEnumerated
		if _child, err = (&_BACnetContextTagEnumerated{}).parse(ctx, readBuffer, m, header, tagNumberArgument, dataType); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetContextTagEnumerated for type-switch of BACnetContextTag")
		}
	case dataType == BACnetDataType_DATE: // BACnetContextTagDate
		if _child, err = (&_BACnetContextTagDate{}).parse(ctx, readBuffer, m, tagNumberArgument, dataType); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetContextTagDate for type-switch of BACnetContextTag")
		}
	case dataType == BACnetDataType_TIME: // BACnetContextTagTime
		if _child, err = (&_BACnetContextTagTime{}).parse(ctx, readBuffer, m, tagNumberArgument, dataType); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetContextTagTime for type-switch of BACnetContextTag")
		}
	case dataType == BACnetDataType_BACNET_OBJECT_IDENTIFIER: // BACnetContextTagObjectIdentifier
		if _child, err = (&_BACnetContextTagObjectIdentifier{}).parse(ctx, readBuffer, m, tagNumberArgument, dataType); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetContextTagObjectIdentifier for type-switch of BACnetContextTag")
		}
	case dataType == BACnetDataType_UNKNOWN: // BACnetContextTagUnknown
		if _child, err = (&_BACnetContextTagUnknown{}).parse(ctx, readBuffer, m, actualLength, tagNumberArgument, dataType); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetContextTagUnknown for type-switch of BACnetContextTag")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [dataType=%v]", dataType)
	}

	if closeErr := readBuffer.CloseContext("BACnetContextTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetContextTag")
	}

	return _child, nil
}

func (pm *_BACnetContextTag) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetContextTag, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetContextTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetContextTag")
	}

	if err := WriteSimpleField[BACnetTagHeader](ctx, "header", m.GetHeader(), WriteComplex[BACnetTagHeader](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'header' field")
	}
	// Virtual field
	tagNumber := m.GetTagNumber()
	_ = tagNumber
	if _tagNumberErr := writeBuffer.WriteVirtual(ctx, "tagNumber", m.GetTagNumber()); _tagNumberErr != nil {
		return errors.Wrap(_tagNumberErr, "Error serializing 'tagNumber' field")
	}
	// Virtual field
	actualLength := m.GetActualLength()
	_ = actualLength
	if _actualLengthErr := writeBuffer.WriteVirtual(ctx, "actualLength", m.GetActualLength()); _actualLengthErr != nil {
		return errors.Wrap(_actualLengthErr, "Error serializing 'actualLength' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("BACnetContextTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetContextTag")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetContextTag) GetTagNumberArgument() uint8 {
	return m.TagNumberArgument
}

//
////

func (m *_BACnetContextTag) isBACnetContextTag() bool {
	return true
}
