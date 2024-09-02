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

// SubscribeCOVPropertyMultipleErrorFirstFailedSubscription is the corresponding interface of SubscribeCOVPropertyMultipleErrorFirstFailedSubscription
type SubscribeCOVPropertyMultipleErrorFirstFailedSubscription interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetMonitoredObjectIdentifier returns MonitoredObjectIdentifier (property field)
	GetMonitoredObjectIdentifier() BACnetContextTagObjectIdentifier
	// GetMonitoredPropertyReference returns MonitoredPropertyReference (property field)
	GetMonitoredPropertyReference() BACnetPropertyReferenceEnclosed
	// GetErrorType returns ErrorType (property field)
	GetErrorType() ErrorEnclosed
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionExactly can be used when we want exactly this type and not a type which fulfills SubscribeCOVPropertyMultipleErrorFirstFailedSubscription.
// This is useful for switch cases.
type SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionExactly interface {
	SubscribeCOVPropertyMultipleErrorFirstFailedSubscription
	isSubscribeCOVPropertyMultipleErrorFirstFailedSubscription() bool
}

// _SubscribeCOVPropertyMultipleErrorFirstFailedSubscription is the data-structure of this message
type _SubscribeCOVPropertyMultipleErrorFirstFailedSubscription struct {
	OpeningTag                 BACnetOpeningTag
	MonitoredObjectIdentifier  BACnetContextTagObjectIdentifier
	MonitoredPropertyReference BACnetPropertyReferenceEnclosed
	ErrorType                  ErrorEnclosed
	ClosingTag                 BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

var _ SubscribeCOVPropertyMultipleErrorFirstFailedSubscription = (*_SubscribeCOVPropertyMultipleErrorFirstFailedSubscription)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SubscribeCOVPropertyMultipleErrorFirstFailedSubscription) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_SubscribeCOVPropertyMultipleErrorFirstFailedSubscription) GetMonitoredObjectIdentifier() BACnetContextTagObjectIdentifier {
	return m.MonitoredObjectIdentifier
}

func (m *_SubscribeCOVPropertyMultipleErrorFirstFailedSubscription) GetMonitoredPropertyReference() BACnetPropertyReferenceEnclosed {
	return m.MonitoredPropertyReference
}

func (m *_SubscribeCOVPropertyMultipleErrorFirstFailedSubscription) GetErrorType() ErrorEnclosed {
	return m.ErrorType
}

func (m *_SubscribeCOVPropertyMultipleErrorFirstFailedSubscription) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSubscribeCOVPropertyMultipleErrorFirstFailedSubscription factory function for _SubscribeCOVPropertyMultipleErrorFirstFailedSubscription
func NewSubscribeCOVPropertyMultipleErrorFirstFailedSubscription(openingTag BACnetOpeningTag, monitoredObjectIdentifier BACnetContextTagObjectIdentifier, monitoredPropertyReference BACnetPropertyReferenceEnclosed, errorType ErrorEnclosed, closingTag BACnetClosingTag, tagNumber uint8) *_SubscribeCOVPropertyMultipleErrorFirstFailedSubscription {
	return &_SubscribeCOVPropertyMultipleErrorFirstFailedSubscription{OpeningTag: openingTag, MonitoredObjectIdentifier: monitoredObjectIdentifier, MonitoredPropertyReference: monitoredPropertyReference, ErrorType: errorType, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastSubscribeCOVPropertyMultipleErrorFirstFailedSubscription(structType any) SubscribeCOVPropertyMultipleErrorFirstFailedSubscription {
	if casted, ok := structType.(SubscribeCOVPropertyMultipleErrorFirstFailedSubscription); ok {
		return casted
	}
	if casted, ok := structType.(*SubscribeCOVPropertyMultipleErrorFirstFailedSubscription); ok {
		return *casted
	}
	return nil
}

func (m *_SubscribeCOVPropertyMultipleErrorFirstFailedSubscription) GetTypeName() string {
	return "SubscribeCOVPropertyMultipleErrorFirstFailedSubscription"
}

func (m *_SubscribeCOVPropertyMultipleErrorFirstFailedSubscription) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (monitoredObjectIdentifier)
	lengthInBits += m.MonitoredObjectIdentifier.GetLengthInBits(ctx)

	// Simple field (monitoredPropertyReference)
	lengthInBits += m.MonitoredPropertyReference.GetLengthInBits(ctx)

	// Simple field (errorType)
	lengthInBits += m.ErrorType.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_SubscribeCOVPropertyMultipleErrorFirstFailedSubscription) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionParse(ctx context.Context, theBytes []byte, tagNumber uint8) (SubscribeCOVPropertyMultipleErrorFirstFailedSubscription, error) {
	return SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionParseWithBufferProducer(tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (SubscribeCOVPropertyMultipleErrorFirstFailedSubscription, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (SubscribeCOVPropertyMultipleErrorFirstFailedSubscription, error) {
		return SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionParseWithBuffer(ctx, readBuffer, tagNumber)
	}
}

func SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (SubscribeCOVPropertyMultipleErrorFirstFailedSubscription, error) {
	v, err := NewSubscribeCOVPropertyMultipleErrorFirstFailedSubscription(tagNumber).parse(ctx, readBuffer, tagNumber)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_SubscribeCOVPropertyMultipleErrorFirstFailedSubscription) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (__subscribeCOVPropertyMultipleErrorFirstFailedSubscription SubscribeCOVPropertyMultipleErrorFirstFailedSubscription, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SubscribeCOVPropertyMultipleErrorFirstFailedSubscription"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SubscribeCOVPropertyMultipleErrorFirstFailedSubscription")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	monitoredObjectIdentifier, err := ReadSimpleField[BACnetContextTagObjectIdentifier](ctx, "monitoredObjectIdentifier", ReadComplex[BACnetContextTagObjectIdentifier](BACnetContextTagParseWithBufferProducer[BACnetContextTagObjectIdentifier]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_BACNET_OBJECT_IDENTIFIER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'monitoredObjectIdentifier' field"))
	}
	m.MonitoredObjectIdentifier = monitoredObjectIdentifier

	monitoredPropertyReference, err := ReadSimpleField[BACnetPropertyReferenceEnclosed](ctx, "monitoredPropertyReference", ReadComplex[BACnetPropertyReferenceEnclosed](BACnetPropertyReferenceEnclosedParseWithBufferProducer((uint8)(uint8(1))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'monitoredPropertyReference' field"))
	}
	m.MonitoredPropertyReference = monitoredPropertyReference

	errorType, err := ReadSimpleField[ErrorEnclosed](ctx, "errorType", ReadComplex[ErrorEnclosed](ErrorEnclosedParseWithBufferProducer((uint8)(uint8(2))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'errorType' field"))
	}
	m.ErrorType = errorType

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("SubscribeCOVPropertyMultipleErrorFirstFailedSubscription"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SubscribeCOVPropertyMultipleErrorFirstFailedSubscription")
	}

	return m, nil
}

func (m *_SubscribeCOVPropertyMultipleErrorFirstFailedSubscription) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SubscribeCOVPropertyMultipleErrorFirstFailedSubscription) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("SubscribeCOVPropertyMultipleErrorFirstFailedSubscription"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for SubscribeCOVPropertyMultipleErrorFirstFailedSubscription")
	}

	if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'openingTag' field")
	}

	if err := WriteSimpleField[BACnetContextTagObjectIdentifier](ctx, "monitoredObjectIdentifier", m.GetMonitoredObjectIdentifier(), WriteComplex[BACnetContextTagObjectIdentifier](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'monitoredObjectIdentifier' field")
	}

	if err := WriteSimpleField[BACnetPropertyReferenceEnclosed](ctx, "monitoredPropertyReference", m.GetMonitoredPropertyReference(), WriteComplex[BACnetPropertyReferenceEnclosed](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'monitoredPropertyReference' field")
	}

	if err := WriteSimpleField[ErrorEnclosed](ctx, "errorType", m.GetErrorType(), WriteComplex[ErrorEnclosed](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'errorType' field")
	}

	if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("SubscribeCOVPropertyMultipleErrorFirstFailedSubscription"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for SubscribeCOVPropertyMultipleErrorFirstFailedSubscription")
	}
	return nil
}

////
// Arguments Getter

func (m *_SubscribeCOVPropertyMultipleErrorFirstFailedSubscription) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_SubscribeCOVPropertyMultipleErrorFirstFailedSubscription) isSubscribeCOVPropertyMultipleErrorFirstFailedSubscription() bool {
	return true
}

func (m *_SubscribeCOVPropertyMultipleErrorFirstFailedSubscription) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
