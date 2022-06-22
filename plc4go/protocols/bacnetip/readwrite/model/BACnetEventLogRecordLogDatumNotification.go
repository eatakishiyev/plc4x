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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetEventLogRecordLogDatumNotification is the corresponding interface of BACnetEventLogRecordLogDatumNotification
type BACnetEventLogRecordLogDatumNotification interface {
	utils.LengthAware
	utils.Serializable
	BACnetEventLogRecordLogDatum
	// GetInnerOpeningTag returns InnerOpeningTag (property field)
	GetInnerOpeningTag() BACnetOpeningTag
	// GetNotification returns Notification (property field)
	GetNotification() ConfirmedEventNotificationRequest
	// GetInnerClosingTag returns InnerClosingTag (property field)
	GetInnerClosingTag() BACnetClosingTag
}

// BACnetEventLogRecordLogDatumNotificationExactly can be used when we want exactly this type and not a type which fulfills BACnetEventLogRecordLogDatumNotification.
// This is useful for switch cases.
type BACnetEventLogRecordLogDatumNotificationExactly interface {
	isBACnetEventLogRecordLogDatumNotification() bool
}

// _BACnetEventLogRecordLogDatumNotification is the data-structure of this message
type _BACnetEventLogRecordLogDatumNotification struct {
	*_BACnetEventLogRecordLogDatum
	InnerOpeningTag BACnetOpeningTag
	Notification    ConfirmedEventNotificationRequest
	InnerClosingTag BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetEventLogRecordLogDatumNotification) InitializeParent(parent BACnetEventLogRecordLogDatum, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetEventLogRecordLogDatumNotification) GetParent() BACnetEventLogRecordLogDatum {
	return m._BACnetEventLogRecordLogDatum
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventLogRecordLogDatumNotification) GetInnerOpeningTag() BACnetOpeningTag {
	return m.InnerOpeningTag
}

func (m *_BACnetEventLogRecordLogDatumNotification) GetNotification() ConfirmedEventNotificationRequest {
	return m.Notification
}

func (m *_BACnetEventLogRecordLogDatumNotification) GetInnerClosingTag() BACnetClosingTag {
	return m.InnerClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetEventLogRecordLogDatumNotification factory function for _BACnetEventLogRecordLogDatumNotification
func NewBACnetEventLogRecordLogDatumNotification(innerOpeningTag BACnetOpeningTag, notification ConfirmedEventNotificationRequest, innerClosingTag BACnetClosingTag, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetEventLogRecordLogDatumNotification {
	_result := &_BACnetEventLogRecordLogDatumNotification{
		InnerOpeningTag:               innerOpeningTag,
		Notification:                  notification,
		InnerClosingTag:               innerClosingTag,
		_BACnetEventLogRecordLogDatum: NewBACnetEventLogRecordLogDatum(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result._BACnetEventLogRecordLogDatum._BACnetEventLogRecordLogDatumChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetEventLogRecordLogDatumNotification(structType interface{}) BACnetEventLogRecordLogDatumNotification {
	if casted, ok := structType.(BACnetEventLogRecordLogDatumNotification); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventLogRecordLogDatumNotification); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventLogRecordLogDatumNotification) GetTypeName() string {
	return "BACnetEventLogRecordLogDatumNotification"
}

func (m *_BACnetEventLogRecordLogDatumNotification) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetEventLogRecordLogDatumNotification) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (innerOpeningTag)
	lengthInBits += m.InnerOpeningTag.GetLengthInBits()

	// Simple field (notification)
	lengthInBits += m.Notification.GetLengthInBits()

	// Simple field (innerClosingTag)
	lengthInBits += m.InnerClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetEventLogRecordLogDatumNotification) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetEventLogRecordLogDatumNotificationParse(readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetEventLogRecordLogDatumNotification, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventLogRecordLogDatumNotification"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventLogRecordLogDatumNotification")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (innerOpeningTag)
	if pullErr := readBuffer.PullContext("innerOpeningTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for innerOpeningTag")
	}
	_innerOpeningTag, _innerOpeningTagErr := BACnetOpeningTagParse(readBuffer, uint8(uint8(1)))
	if _innerOpeningTagErr != nil {
		return nil, errors.Wrap(_innerOpeningTagErr, "Error parsing 'innerOpeningTag' field")
	}
	innerOpeningTag := _innerOpeningTag.(BACnetOpeningTag)
	if closeErr := readBuffer.CloseContext("innerOpeningTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for innerOpeningTag")
	}

	// Simple Field (notification)
	if pullErr := readBuffer.PullContext("notification"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for notification")
	}
	_notification, _notificationErr := ConfirmedEventNotificationRequestParse(readBuffer)
	if _notificationErr != nil {
		return nil, errors.Wrap(_notificationErr, "Error parsing 'notification' field")
	}
	notification := _notification.(ConfirmedEventNotificationRequest)
	if closeErr := readBuffer.CloseContext("notification"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for notification")
	}

	// Simple Field (innerClosingTag)
	if pullErr := readBuffer.PullContext("innerClosingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for innerClosingTag")
	}
	_innerClosingTag, _innerClosingTagErr := BACnetClosingTagParse(readBuffer, uint8(tagNumber))
	if _innerClosingTagErr != nil {
		return nil, errors.Wrap(_innerClosingTagErr, "Error parsing 'innerClosingTag' field")
	}
	innerClosingTag := _innerClosingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("innerClosingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for innerClosingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetEventLogRecordLogDatumNotification"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventLogRecordLogDatumNotification")
	}

	// Create a partially initialized instance
	_child := &_BACnetEventLogRecordLogDatumNotification{
		InnerOpeningTag:               innerOpeningTag,
		Notification:                  notification,
		InnerClosingTag:               innerClosingTag,
		_BACnetEventLogRecordLogDatum: &_BACnetEventLogRecordLogDatum{},
	}
	_child._BACnetEventLogRecordLogDatum._BACnetEventLogRecordLogDatumChildRequirements = _child
	return _child, nil
}

func (m *_BACnetEventLogRecordLogDatumNotification) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetEventLogRecordLogDatumNotification"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetEventLogRecordLogDatumNotification")
		}

		// Simple Field (innerOpeningTag)
		if pushErr := writeBuffer.PushContext("innerOpeningTag"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for innerOpeningTag")
		}
		_innerOpeningTagErr := writeBuffer.WriteSerializable(m.GetInnerOpeningTag())
		if popErr := writeBuffer.PopContext("innerOpeningTag"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for innerOpeningTag")
		}
		if _innerOpeningTagErr != nil {
			return errors.Wrap(_innerOpeningTagErr, "Error serializing 'innerOpeningTag' field")
		}

		// Simple Field (notification)
		if pushErr := writeBuffer.PushContext("notification"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for notification")
		}
		_notificationErr := writeBuffer.WriteSerializable(m.GetNotification())
		if popErr := writeBuffer.PopContext("notification"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for notification")
		}
		if _notificationErr != nil {
			return errors.Wrap(_notificationErr, "Error serializing 'notification' field")
		}

		// Simple Field (innerClosingTag)
		if pushErr := writeBuffer.PushContext("innerClosingTag"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for innerClosingTag")
		}
		_innerClosingTagErr := writeBuffer.WriteSerializable(m.GetInnerClosingTag())
		if popErr := writeBuffer.PopContext("innerClosingTag"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for innerClosingTag")
		}
		if _innerClosingTagErr != nil {
			return errors.Wrap(_innerClosingTagErr, "Error serializing 'innerClosingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetEventLogRecordLogDatumNotification"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetEventLogRecordLogDatumNotification")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetEventLogRecordLogDatumNotification) isBACnetEventLogRecordLogDatumNotification() bool {
	return true
}

func (m *_BACnetEventLogRecordLogDatumNotification) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
