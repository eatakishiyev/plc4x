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

// BACnetObjectPropertyReferenceEnclosed is the corresponding interface of BACnetObjectPropertyReferenceEnclosed
type BACnetObjectPropertyReferenceEnclosed interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetObjectPropertyReference returns ObjectPropertyReference (property field)
	GetObjectPropertyReference() BACnetObjectPropertyReference
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// BACnetObjectPropertyReferenceEnclosedExactly can be used when we want exactly this type and not a type which fulfills BACnetObjectPropertyReferenceEnclosed.
// This is useful for switch cases.
type BACnetObjectPropertyReferenceEnclosedExactly interface {
	BACnetObjectPropertyReferenceEnclosed
	isBACnetObjectPropertyReferenceEnclosed() bool
}

// _BACnetObjectPropertyReferenceEnclosed is the data-structure of this message
type _BACnetObjectPropertyReferenceEnclosed struct {
	OpeningTag              BACnetOpeningTag
	ObjectPropertyReference BACnetObjectPropertyReference
	ClosingTag              BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

var _ BACnetObjectPropertyReferenceEnclosed = (*_BACnetObjectPropertyReferenceEnclosed)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetObjectPropertyReferenceEnclosed) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetObjectPropertyReferenceEnclosed) GetObjectPropertyReference() BACnetObjectPropertyReference {
	return m.ObjectPropertyReference
}

func (m *_BACnetObjectPropertyReferenceEnclosed) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetObjectPropertyReferenceEnclosed factory function for _BACnetObjectPropertyReferenceEnclosed
func NewBACnetObjectPropertyReferenceEnclosed(openingTag BACnetOpeningTag, objectPropertyReference BACnetObjectPropertyReference, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetObjectPropertyReferenceEnclosed {
	return &_BACnetObjectPropertyReferenceEnclosed{OpeningTag: openingTag, ObjectPropertyReference: objectPropertyReference, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastBACnetObjectPropertyReferenceEnclosed(structType any) BACnetObjectPropertyReferenceEnclosed {
	if casted, ok := structType.(BACnetObjectPropertyReferenceEnclosed); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetObjectPropertyReferenceEnclosed); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetObjectPropertyReferenceEnclosed) GetTypeName() string {
	return "BACnetObjectPropertyReferenceEnclosed"
}

func (m *_BACnetObjectPropertyReferenceEnclosed) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (objectPropertyReference)
	lengthInBits += m.ObjectPropertyReference.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetObjectPropertyReferenceEnclosed) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetObjectPropertyReferenceEnclosedParse(ctx context.Context, theBytes []byte, tagNumber uint8) (BACnetObjectPropertyReferenceEnclosed, error) {
	return BACnetObjectPropertyReferenceEnclosedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetObjectPropertyReferenceEnclosedParseWithBufferProducer(tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetObjectPropertyReferenceEnclosed, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetObjectPropertyReferenceEnclosed, error) {
		return BACnetObjectPropertyReferenceEnclosedParseWithBuffer(ctx, readBuffer, tagNumber)
	}
}

func BACnetObjectPropertyReferenceEnclosedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetObjectPropertyReferenceEnclosed, error) {
	v, err := NewBACnetObjectPropertyReferenceEnclosed(tagNumber).parse(ctx, readBuffer, tagNumber)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_BACnetObjectPropertyReferenceEnclosed) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (__bACnetObjectPropertyReferenceEnclosed BACnetObjectPropertyReferenceEnclosed, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetObjectPropertyReferenceEnclosed"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetObjectPropertyReferenceEnclosed")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	objectPropertyReference, err := ReadSimpleField[BACnetObjectPropertyReference](ctx, "objectPropertyReference", ReadComplex[BACnetObjectPropertyReference](BACnetObjectPropertyReferenceParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'objectPropertyReference' field"))
	}
	m.ObjectPropertyReference = objectPropertyReference

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetObjectPropertyReferenceEnclosed"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetObjectPropertyReferenceEnclosed")
	}

	return m, nil
}

func (m *_BACnetObjectPropertyReferenceEnclosed) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetObjectPropertyReferenceEnclosed) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetObjectPropertyReferenceEnclosed"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetObjectPropertyReferenceEnclosed")
	}

	if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'openingTag' field")
	}

	if err := WriteSimpleField[BACnetObjectPropertyReference](ctx, "objectPropertyReference", m.GetObjectPropertyReference(), WriteComplex[BACnetObjectPropertyReference](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'objectPropertyReference' field")
	}

	if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetObjectPropertyReferenceEnclosed"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetObjectPropertyReferenceEnclosed")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetObjectPropertyReferenceEnclosed) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetObjectPropertyReferenceEnclosed) isBACnetObjectPropertyReferenceEnclosed() bool {
	return true
}

func (m *_BACnetObjectPropertyReferenceEnclosed) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
