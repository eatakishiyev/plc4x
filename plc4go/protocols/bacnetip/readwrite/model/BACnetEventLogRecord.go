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

// BACnetEventLogRecord is the corresponding interface of BACnetEventLogRecord
type BACnetEventLogRecord interface {
	utils.LengthAware
	utils.Serializable
	// GetTimestamp returns Timestamp (property field)
	GetTimestamp() BACnetDateTimeEnclosed
	// GetLogDatum returns LogDatum (property field)
	GetLogDatum() BACnetEventLogRecordLogDatum
}

// BACnetEventLogRecordExactly can be used when we want exactly this type and not a type which fulfills BACnetEventLogRecord.
// This is useful for switch cases.
type BACnetEventLogRecordExactly interface {
	isBACnetEventLogRecord() bool
}

// _BACnetEventLogRecord is the data-structure of this message
type _BACnetEventLogRecord struct {
	Timestamp BACnetDateTimeEnclosed
	LogDatum  BACnetEventLogRecordLogDatum
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventLogRecord) GetTimestamp() BACnetDateTimeEnclosed {
	return m.Timestamp
}

func (m *_BACnetEventLogRecord) GetLogDatum() BACnetEventLogRecordLogDatum {
	return m.LogDatum
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetEventLogRecord factory function for _BACnetEventLogRecord
func NewBACnetEventLogRecord(timestamp BACnetDateTimeEnclosed, logDatum BACnetEventLogRecordLogDatum) *_BACnetEventLogRecord {
	return &_BACnetEventLogRecord{Timestamp: timestamp, LogDatum: logDatum}
}

// Deprecated: use the interface for direct cast
func CastBACnetEventLogRecord(structType interface{}) BACnetEventLogRecord {
	if casted, ok := structType.(BACnetEventLogRecord); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventLogRecord); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventLogRecord) GetTypeName() string {
	return "BACnetEventLogRecord"
}

func (m *_BACnetEventLogRecord) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetEventLogRecord) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (timestamp)
	lengthInBits += m.Timestamp.GetLengthInBits()

	// Simple field (logDatum)
	lengthInBits += m.LogDatum.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetEventLogRecord) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetEventLogRecordParse(readBuffer utils.ReadBuffer) (BACnetEventLogRecord, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventLogRecord"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventLogRecord")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (timestamp)
	if pullErr := readBuffer.PullContext("timestamp"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for timestamp")
	}
	_timestamp, _timestampErr := BACnetDateTimeEnclosedParse(readBuffer, uint8(uint8(0)))
	if _timestampErr != nil {
		return nil, errors.Wrap(_timestampErr, "Error parsing 'timestamp' field")
	}
	timestamp := _timestamp.(BACnetDateTimeEnclosed)
	if closeErr := readBuffer.CloseContext("timestamp"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for timestamp")
	}

	// Simple Field (logDatum)
	if pullErr := readBuffer.PullContext("logDatum"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for logDatum")
	}
	_logDatum, _logDatumErr := BACnetEventLogRecordLogDatumParse(readBuffer, uint8(uint8(1)))
	if _logDatumErr != nil {
		return nil, errors.Wrap(_logDatumErr, "Error parsing 'logDatum' field")
	}
	logDatum := _logDatum.(BACnetEventLogRecordLogDatum)
	if closeErr := readBuffer.CloseContext("logDatum"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for logDatum")
	}

	if closeErr := readBuffer.CloseContext("BACnetEventLogRecord"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventLogRecord")
	}

	// Create the instance
	return NewBACnetEventLogRecord(timestamp, logDatum), nil
}

func (m *_BACnetEventLogRecord) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetEventLogRecord"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetEventLogRecord")
	}

	// Simple Field (timestamp)
	if pushErr := writeBuffer.PushContext("timestamp"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for timestamp")
	}
	_timestampErr := writeBuffer.WriteSerializable(m.GetTimestamp())
	if popErr := writeBuffer.PopContext("timestamp"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for timestamp")
	}
	if _timestampErr != nil {
		return errors.Wrap(_timestampErr, "Error serializing 'timestamp' field")
	}

	// Simple Field (logDatum)
	if pushErr := writeBuffer.PushContext("logDatum"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for logDatum")
	}
	_logDatumErr := writeBuffer.WriteSerializable(m.GetLogDatum())
	if popErr := writeBuffer.PopContext("logDatum"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for logDatum")
	}
	if _logDatumErr != nil {
		return errors.Wrap(_logDatumErr, "Error serializing 'logDatum' field")
	}

	if popErr := writeBuffer.PopContext("BACnetEventLogRecord"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetEventLogRecord")
	}
	return nil
}

func (m *_BACnetEventLogRecord) isBACnetEventLogRecord() bool {
	return true
}

func (m *_BACnetEventLogRecord) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
