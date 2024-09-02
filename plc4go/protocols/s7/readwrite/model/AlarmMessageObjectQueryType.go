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

// Constant values.
const AlarmMessageObjectQueryType_VARIABLESPEC uint8 = 0x12

// AlarmMessageObjectQueryType is the corresponding interface of AlarmMessageObjectQueryType
type AlarmMessageObjectQueryType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetLengthDataset returns LengthDataset (property field)
	GetLengthDataset() uint8
	// GetEventState returns EventState (property field)
	GetEventState() State
	// GetAckStateGoing returns AckStateGoing (property field)
	GetAckStateGoing() State
	// GetAckStateComing returns AckStateComing (property field)
	GetAckStateComing() State
	// GetTimeComing returns TimeComing (property field)
	GetTimeComing() DateAndTime
	// GetValueComing returns ValueComing (property field)
	GetValueComing() AssociatedValueType
	// GetTimeGoing returns TimeGoing (property field)
	GetTimeGoing() DateAndTime
	// GetValueGoing returns ValueGoing (property field)
	GetValueGoing() AssociatedValueType
}

// AlarmMessageObjectQueryTypeExactly can be used when we want exactly this type and not a type which fulfills AlarmMessageObjectQueryType.
// This is useful for switch cases.
type AlarmMessageObjectQueryTypeExactly interface {
	AlarmMessageObjectQueryType
	isAlarmMessageObjectQueryType() bool
}

// _AlarmMessageObjectQueryType is the data-structure of this message
type _AlarmMessageObjectQueryType struct {
	LengthDataset  uint8
	EventState     State
	AckStateGoing  State
	AckStateComing State
	TimeComing     DateAndTime
	ValueComing    AssociatedValueType
	TimeGoing      DateAndTime
	ValueGoing     AssociatedValueType
	// Reserved Fields
	reservedField0 *uint16
}

var _ AlarmMessageObjectQueryType = (*_AlarmMessageObjectQueryType)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AlarmMessageObjectQueryType) GetLengthDataset() uint8 {
	return m.LengthDataset
}

func (m *_AlarmMessageObjectQueryType) GetEventState() State {
	return m.EventState
}

func (m *_AlarmMessageObjectQueryType) GetAckStateGoing() State {
	return m.AckStateGoing
}

func (m *_AlarmMessageObjectQueryType) GetAckStateComing() State {
	return m.AckStateComing
}

func (m *_AlarmMessageObjectQueryType) GetTimeComing() DateAndTime {
	return m.TimeComing
}

func (m *_AlarmMessageObjectQueryType) GetValueComing() AssociatedValueType {
	return m.ValueComing
}

func (m *_AlarmMessageObjectQueryType) GetTimeGoing() DateAndTime {
	return m.TimeGoing
}

func (m *_AlarmMessageObjectQueryType) GetValueGoing() AssociatedValueType {
	return m.ValueGoing
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_AlarmMessageObjectQueryType) GetVariableSpec() uint8 {
	return AlarmMessageObjectQueryType_VARIABLESPEC
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAlarmMessageObjectQueryType factory function for _AlarmMessageObjectQueryType
func NewAlarmMessageObjectQueryType(lengthDataset uint8, eventState State, ackStateGoing State, ackStateComing State, timeComing DateAndTime, valueComing AssociatedValueType, timeGoing DateAndTime, valueGoing AssociatedValueType) *_AlarmMessageObjectQueryType {
	return &_AlarmMessageObjectQueryType{LengthDataset: lengthDataset, EventState: eventState, AckStateGoing: ackStateGoing, AckStateComing: ackStateComing, TimeComing: timeComing, ValueComing: valueComing, TimeGoing: timeGoing, ValueGoing: valueGoing}
}

// Deprecated: use the interface for direct cast
func CastAlarmMessageObjectQueryType(structType any) AlarmMessageObjectQueryType {
	if casted, ok := structType.(AlarmMessageObjectQueryType); ok {
		return casted
	}
	if casted, ok := structType.(*AlarmMessageObjectQueryType); ok {
		return *casted
	}
	return nil
}

func (m *_AlarmMessageObjectQueryType) GetTypeName() string {
	return "AlarmMessageObjectQueryType"
}

func (m *_AlarmMessageObjectQueryType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (lengthDataset)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 16

	// Const Field (variableSpec)
	lengthInBits += 8

	// Simple field (eventState)
	lengthInBits += m.EventState.GetLengthInBits(ctx)

	// Simple field (ackStateGoing)
	lengthInBits += m.AckStateGoing.GetLengthInBits(ctx)

	// Simple field (ackStateComing)
	lengthInBits += m.AckStateComing.GetLengthInBits(ctx)

	// Simple field (timeComing)
	lengthInBits += m.TimeComing.GetLengthInBits(ctx)

	// Simple field (valueComing)
	lengthInBits += m.ValueComing.GetLengthInBits(ctx)

	// Simple field (timeGoing)
	lengthInBits += m.TimeGoing.GetLengthInBits(ctx)

	// Simple field (valueGoing)
	lengthInBits += m.ValueGoing.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_AlarmMessageObjectQueryType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AlarmMessageObjectQueryTypeParse(ctx context.Context, theBytes []byte) (AlarmMessageObjectQueryType, error) {
	return AlarmMessageObjectQueryTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func AlarmMessageObjectQueryTypeParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (AlarmMessageObjectQueryType, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (AlarmMessageObjectQueryType, error) {
		return AlarmMessageObjectQueryTypeParseWithBuffer(ctx, readBuffer)
	}
}

func AlarmMessageObjectQueryTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AlarmMessageObjectQueryType, error) {
	v, err := NewAlarmMessageObjectQueryType().parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_AlarmMessageObjectQueryType) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__alarmMessageObjectQueryType AlarmMessageObjectQueryType, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AlarmMessageObjectQueryType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AlarmMessageObjectQueryType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	lengthDataset, err := ReadSimpleField(ctx, "lengthDataset", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lengthDataset' field"))
	}
	m.LengthDataset = lengthDataset

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedShort(readBuffer, uint8(16)), uint16(0x0000))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	variableSpec, err := ReadConstField[uint8](ctx, "variableSpec", ReadUnsignedByte(readBuffer, uint8(8)), AlarmMessageObjectQueryType_VARIABLESPEC)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'variableSpec' field"))
	}
	_ = variableSpec

	eventState, err := ReadSimpleField[State](ctx, "eventState", ReadComplex[State](StateParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'eventState' field"))
	}
	m.EventState = eventState

	ackStateGoing, err := ReadSimpleField[State](ctx, "ackStateGoing", ReadComplex[State](StateParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'ackStateGoing' field"))
	}
	m.AckStateGoing = ackStateGoing

	ackStateComing, err := ReadSimpleField[State](ctx, "ackStateComing", ReadComplex[State](StateParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'ackStateComing' field"))
	}
	m.AckStateComing = ackStateComing

	timeComing, err := ReadSimpleField[DateAndTime](ctx, "timeComing", ReadComplex[DateAndTime](DateAndTimeParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeComing' field"))
	}
	m.TimeComing = timeComing

	valueComing, err := ReadSimpleField[AssociatedValueType](ctx, "valueComing", ReadComplex[AssociatedValueType](AssociatedValueTypeParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'valueComing' field"))
	}
	m.ValueComing = valueComing

	timeGoing, err := ReadSimpleField[DateAndTime](ctx, "timeGoing", ReadComplex[DateAndTime](DateAndTimeParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeGoing' field"))
	}
	m.TimeGoing = timeGoing

	valueGoing, err := ReadSimpleField[AssociatedValueType](ctx, "valueGoing", ReadComplex[AssociatedValueType](AssociatedValueTypeParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'valueGoing' field"))
	}
	m.ValueGoing = valueGoing

	if closeErr := readBuffer.CloseContext("AlarmMessageObjectQueryType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AlarmMessageObjectQueryType")
	}

	return m, nil
}

func (m *_AlarmMessageObjectQueryType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AlarmMessageObjectQueryType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("AlarmMessageObjectQueryType"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for AlarmMessageObjectQueryType")
	}

	if err := WriteSimpleField[uint8](ctx, "lengthDataset", m.GetLengthDataset(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'lengthDataset' field")
	}

	if err := WriteReservedField[uint16](ctx, "reserved", uint16(0x0000), WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'reserved' field number 1")
	}

	if err := WriteConstField(ctx, "variableSpec", AlarmMessageObjectQueryType_VARIABLESPEC, WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'variableSpec' field")
	}

	if err := WriteSimpleField[State](ctx, "eventState", m.GetEventState(), WriteComplex[State](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'eventState' field")
	}

	if err := WriteSimpleField[State](ctx, "ackStateGoing", m.GetAckStateGoing(), WriteComplex[State](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'ackStateGoing' field")
	}

	if err := WriteSimpleField[State](ctx, "ackStateComing", m.GetAckStateComing(), WriteComplex[State](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'ackStateComing' field")
	}

	if err := WriteSimpleField[DateAndTime](ctx, "timeComing", m.GetTimeComing(), WriteComplex[DateAndTime](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'timeComing' field")
	}

	if err := WriteSimpleField[AssociatedValueType](ctx, "valueComing", m.GetValueComing(), WriteComplex[AssociatedValueType](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'valueComing' field")
	}

	if err := WriteSimpleField[DateAndTime](ctx, "timeGoing", m.GetTimeGoing(), WriteComplex[DateAndTime](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'timeGoing' field")
	}

	if err := WriteSimpleField[AssociatedValueType](ctx, "valueGoing", m.GetValueGoing(), WriteComplex[AssociatedValueType](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'valueGoing' field")
	}

	if popErr := writeBuffer.PopContext("AlarmMessageObjectQueryType"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for AlarmMessageObjectQueryType")
	}
	return nil
}

func (m *_AlarmMessageObjectQueryType) isAlarmMessageObjectQueryType() bool {
	return true
}

func (m *_AlarmMessageObjectQueryType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
