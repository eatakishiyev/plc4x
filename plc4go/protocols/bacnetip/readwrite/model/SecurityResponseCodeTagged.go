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

// SecurityResponseCodeTagged is the corresponding interface of SecurityResponseCodeTagged
type SecurityResponseCodeTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() SecurityResponseCode
}

// SecurityResponseCodeTaggedExactly can be used when we want exactly this type and not a type which fulfills SecurityResponseCodeTagged.
// This is useful for switch cases.
type SecurityResponseCodeTaggedExactly interface {
	SecurityResponseCodeTagged
	isSecurityResponseCodeTagged() bool
}

// _SecurityResponseCodeTagged is the data-structure of this message
type _SecurityResponseCodeTagged struct {
	Header BACnetTagHeader
	Value  SecurityResponseCode

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

var _ SecurityResponseCodeTagged = (*_SecurityResponseCodeTagged)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SecurityResponseCodeTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_SecurityResponseCodeTagged) GetValue() SecurityResponseCode {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSecurityResponseCodeTagged factory function for _SecurityResponseCodeTagged
func NewSecurityResponseCodeTagged(header BACnetTagHeader, value SecurityResponseCode, tagNumber uint8, tagClass TagClass) *_SecurityResponseCodeTagged {
	return &_SecurityResponseCodeTagged{Header: header, Value: value, TagNumber: tagNumber, TagClass: tagClass}
}

// Deprecated: use the interface for direct cast
func CastSecurityResponseCodeTagged(structType any) SecurityResponseCodeTagged {
	if casted, ok := structType.(SecurityResponseCodeTagged); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityResponseCodeTagged); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityResponseCodeTagged) GetTypeName() string {
	return "SecurityResponseCodeTagged"
}

func (m *_SecurityResponseCodeTagged) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits(ctx)

	// Manual Field (value)
	lengthInBits += uint16(int32(m.GetHeader().GetActualLength()) * int32(int32(8)))

	return lengthInBits
}

func (m *_SecurityResponseCodeTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SecurityResponseCodeTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (SecurityResponseCodeTagged, error) {
	return SecurityResponseCodeTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func SecurityResponseCodeTaggedParseWithBufferProducer(tagNumber uint8, tagClass TagClass) func(ctx context.Context, readBuffer utils.ReadBuffer) (SecurityResponseCodeTagged, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (SecurityResponseCodeTagged, error) {
		return SecurityResponseCodeTaggedParseWithBuffer(ctx, readBuffer, tagNumber, tagClass)
	}
}

func SecurityResponseCodeTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (SecurityResponseCodeTagged, error) {
	v, err := NewSecurityResponseCodeTagged(tagNumber, tagClass).parse(ctx, readBuffer, tagNumber, tagClass)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_SecurityResponseCodeTagged) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (__securityResponseCodeTagged SecurityResponseCodeTagged, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityResponseCodeTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityResponseCodeTagged")
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

	value, err := ReadManualField[SecurityResponseCode](ctx, "value", readBuffer, EnsureType[SecurityResponseCode](ReadEnumGenericFailing(ctx, readBuffer, header.GetActualLength(), SecurityResponseCode_SUCCESS)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	if closeErr := readBuffer.CloseContext("SecurityResponseCodeTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityResponseCodeTagged")
	}

	return m, nil
}

func (m *_SecurityResponseCodeTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityResponseCodeTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("SecurityResponseCodeTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for SecurityResponseCodeTagged")
	}

	if err := WriteSimpleField[BACnetTagHeader](ctx, "header", m.GetHeader(), WriteComplex[BACnetTagHeader](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'header' field")
	}

	if err := WriteManualField[SecurityResponseCode](ctx, "value", func(ctx context.Context) error { return WriteEnumGeneric(ctx, writeBuffer, m.GetValue()) }, writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'value' field")
	}

	if popErr := writeBuffer.PopContext("SecurityResponseCodeTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for SecurityResponseCodeTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_SecurityResponseCodeTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_SecurityResponseCodeTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_SecurityResponseCodeTagged) isSecurityResponseCodeTagged() bool {
	return true
}

func (m *_SecurityResponseCodeTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
