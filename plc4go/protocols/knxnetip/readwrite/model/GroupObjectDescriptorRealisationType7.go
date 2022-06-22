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

// GroupObjectDescriptorRealisationType7 is the corresponding interface of GroupObjectDescriptorRealisationType7
type GroupObjectDescriptorRealisationType7 interface {
	utils.LengthAware
	utils.Serializable
	// GetDataAddress returns DataAddress (property field)
	GetDataAddress() uint16
	// GetUpdateEnable returns UpdateEnable (property field)
	GetUpdateEnable() bool
	// GetTransmitEnable returns TransmitEnable (property field)
	GetTransmitEnable() bool
	// GetSegmentSelectorEnable returns SegmentSelectorEnable (property field)
	GetSegmentSelectorEnable() bool
	// GetWriteEnable returns WriteEnable (property field)
	GetWriteEnable() bool
	// GetReadEnable returns ReadEnable (property field)
	GetReadEnable() bool
	// GetCommunicationEnable returns CommunicationEnable (property field)
	GetCommunicationEnable() bool
	// GetPriority returns Priority (property field)
	GetPriority() CEMIPriority
	// GetValueType returns ValueType (property field)
	GetValueType() ComObjectValueType
}

// GroupObjectDescriptorRealisationType7Exactly can be used when we want exactly this type and not a type which fulfills GroupObjectDescriptorRealisationType7.
// This is useful for switch cases.
type GroupObjectDescriptorRealisationType7Exactly interface {
	isGroupObjectDescriptorRealisationType7() bool
}

// _GroupObjectDescriptorRealisationType7 is the data-structure of this message
type _GroupObjectDescriptorRealisationType7 struct {
	DataAddress           uint16
	UpdateEnable          bool
	TransmitEnable        bool
	SegmentSelectorEnable bool
	WriteEnable           bool
	ReadEnable            bool
	CommunicationEnable   bool
	Priority              CEMIPriority
	ValueType             ComObjectValueType
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_GroupObjectDescriptorRealisationType7) GetDataAddress() uint16 {
	return m.DataAddress
}

func (m *_GroupObjectDescriptorRealisationType7) GetUpdateEnable() bool {
	return m.UpdateEnable
}

func (m *_GroupObjectDescriptorRealisationType7) GetTransmitEnable() bool {
	return m.TransmitEnable
}

func (m *_GroupObjectDescriptorRealisationType7) GetSegmentSelectorEnable() bool {
	return m.SegmentSelectorEnable
}

func (m *_GroupObjectDescriptorRealisationType7) GetWriteEnable() bool {
	return m.WriteEnable
}

func (m *_GroupObjectDescriptorRealisationType7) GetReadEnable() bool {
	return m.ReadEnable
}

func (m *_GroupObjectDescriptorRealisationType7) GetCommunicationEnable() bool {
	return m.CommunicationEnable
}

func (m *_GroupObjectDescriptorRealisationType7) GetPriority() CEMIPriority {
	return m.Priority
}

func (m *_GroupObjectDescriptorRealisationType7) GetValueType() ComObjectValueType {
	return m.ValueType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewGroupObjectDescriptorRealisationType7 factory function for _GroupObjectDescriptorRealisationType7
func NewGroupObjectDescriptorRealisationType7(dataAddress uint16, updateEnable bool, transmitEnable bool, segmentSelectorEnable bool, writeEnable bool, readEnable bool, communicationEnable bool, priority CEMIPriority, valueType ComObjectValueType) *_GroupObjectDescriptorRealisationType7 {
	return &_GroupObjectDescriptorRealisationType7{DataAddress: dataAddress, UpdateEnable: updateEnable, TransmitEnable: transmitEnable, SegmentSelectorEnable: segmentSelectorEnable, WriteEnable: writeEnable, ReadEnable: readEnable, CommunicationEnable: communicationEnable, Priority: priority, ValueType: valueType}
}

// Deprecated: use the interface for direct cast
func CastGroupObjectDescriptorRealisationType7(structType interface{}) GroupObjectDescriptorRealisationType7 {
	if casted, ok := structType.(GroupObjectDescriptorRealisationType7); ok {
		return casted
	}
	if casted, ok := structType.(*GroupObjectDescriptorRealisationType7); ok {
		return *casted
	}
	return nil
}

func (m *_GroupObjectDescriptorRealisationType7) GetTypeName() string {
	return "GroupObjectDescriptorRealisationType7"
}

func (m *_GroupObjectDescriptorRealisationType7) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_GroupObjectDescriptorRealisationType7) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (dataAddress)
	lengthInBits += 16

	// Simple field (updateEnable)
	lengthInBits += 1

	// Simple field (transmitEnable)
	lengthInBits += 1

	// Simple field (segmentSelectorEnable)
	lengthInBits += 1

	// Simple field (writeEnable)
	lengthInBits += 1

	// Simple field (readEnable)
	lengthInBits += 1

	// Simple field (communicationEnable)
	lengthInBits += 1

	// Simple field (priority)
	lengthInBits += 2

	// Simple field (valueType)
	lengthInBits += 8

	return lengthInBits
}

func (m *_GroupObjectDescriptorRealisationType7) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func GroupObjectDescriptorRealisationType7Parse(readBuffer utils.ReadBuffer) (GroupObjectDescriptorRealisationType7, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("GroupObjectDescriptorRealisationType7"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for GroupObjectDescriptorRealisationType7")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (dataAddress)
	_dataAddress, _dataAddressErr := readBuffer.ReadUint16("dataAddress", 16)
	if _dataAddressErr != nil {
		return nil, errors.Wrap(_dataAddressErr, "Error parsing 'dataAddress' field")
	}
	dataAddress := _dataAddress

	// Simple Field (updateEnable)
	_updateEnable, _updateEnableErr := readBuffer.ReadBit("updateEnable")
	if _updateEnableErr != nil {
		return nil, errors.Wrap(_updateEnableErr, "Error parsing 'updateEnable' field")
	}
	updateEnable := _updateEnable

	// Simple Field (transmitEnable)
	_transmitEnable, _transmitEnableErr := readBuffer.ReadBit("transmitEnable")
	if _transmitEnableErr != nil {
		return nil, errors.Wrap(_transmitEnableErr, "Error parsing 'transmitEnable' field")
	}
	transmitEnable := _transmitEnable

	// Simple Field (segmentSelectorEnable)
	_segmentSelectorEnable, _segmentSelectorEnableErr := readBuffer.ReadBit("segmentSelectorEnable")
	if _segmentSelectorEnableErr != nil {
		return nil, errors.Wrap(_segmentSelectorEnableErr, "Error parsing 'segmentSelectorEnable' field")
	}
	segmentSelectorEnable := _segmentSelectorEnable

	// Simple Field (writeEnable)
	_writeEnable, _writeEnableErr := readBuffer.ReadBit("writeEnable")
	if _writeEnableErr != nil {
		return nil, errors.Wrap(_writeEnableErr, "Error parsing 'writeEnable' field")
	}
	writeEnable := _writeEnable

	// Simple Field (readEnable)
	_readEnable, _readEnableErr := readBuffer.ReadBit("readEnable")
	if _readEnableErr != nil {
		return nil, errors.Wrap(_readEnableErr, "Error parsing 'readEnable' field")
	}
	readEnable := _readEnable

	// Simple Field (communicationEnable)
	_communicationEnable, _communicationEnableErr := readBuffer.ReadBit("communicationEnable")
	if _communicationEnableErr != nil {
		return nil, errors.Wrap(_communicationEnableErr, "Error parsing 'communicationEnable' field")
	}
	communicationEnable := _communicationEnable

	// Simple Field (priority)
	if pullErr := readBuffer.PullContext("priority"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for priority")
	}
	_priority, _priorityErr := CEMIPriorityParse(readBuffer)
	if _priorityErr != nil {
		return nil, errors.Wrap(_priorityErr, "Error parsing 'priority' field")
	}
	priority := _priority
	if closeErr := readBuffer.CloseContext("priority"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for priority")
	}

	// Simple Field (valueType)
	if pullErr := readBuffer.PullContext("valueType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for valueType")
	}
	_valueType, _valueTypeErr := ComObjectValueTypeParse(readBuffer)
	if _valueTypeErr != nil {
		return nil, errors.Wrap(_valueTypeErr, "Error parsing 'valueType' field")
	}
	valueType := _valueType
	if closeErr := readBuffer.CloseContext("valueType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for valueType")
	}

	if closeErr := readBuffer.CloseContext("GroupObjectDescriptorRealisationType7"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for GroupObjectDescriptorRealisationType7")
	}

	// Create the instance
	return NewGroupObjectDescriptorRealisationType7(dataAddress, updateEnable, transmitEnable, segmentSelectorEnable, writeEnable, readEnable, communicationEnable, priority, valueType), nil
}

func (m *_GroupObjectDescriptorRealisationType7) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("GroupObjectDescriptorRealisationType7"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for GroupObjectDescriptorRealisationType7")
	}

	// Simple Field (dataAddress)
	dataAddress := uint16(m.GetDataAddress())
	_dataAddressErr := writeBuffer.WriteUint16("dataAddress", 16, (dataAddress))
	if _dataAddressErr != nil {
		return errors.Wrap(_dataAddressErr, "Error serializing 'dataAddress' field")
	}

	// Simple Field (updateEnable)
	updateEnable := bool(m.GetUpdateEnable())
	_updateEnableErr := writeBuffer.WriteBit("updateEnable", (updateEnable))
	if _updateEnableErr != nil {
		return errors.Wrap(_updateEnableErr, "Error serializing 'updateEnable' field")
	}

	// Simple Field (transmitEnable)
	transmitEnable := bool(m.GetTransmitEnable())
	_transmitEnableErr := writeBuffer.WriteBit("transmitEnable", (transmitEnable))
	if _transmitEnableErr != nil {
		return errors.Wrap(_transmitEnableErr, "Error serializing 'transmitEnable' field")
	}

	// Simple Field (segmentSelectorEnable)
	segmentSelectorEnable := bool(m.GetSegmentSelectorEnable())
	_segmentSelectorEnableErr := writeBuffer.WriteBit("segmentSelectorEnable", (segmentSelectorEnable))
	if _segmentSelectorEnableErr != nil {
		return errors.Wrap(_segmentSelectorEnableErr, "Error serializing 'segmentSelectorEnable' field")
	}

	// Simple Field (writeEnable)
	writeEnable := bool(m.GetWriteEnable())
	_writeEnableErr := writeBuffer.WriteBit("writeEnable", (writeEnable))
	if _writeEnableErr != nil {
		return errors.Wrap(_writeEnableErr, "Error serializing 'writeEnable' field")
	}

	// Simple Field (readEnable)
	readEnable := bool(m.GetReadEnable())
	_readEnableErr := writeBuffer.WriteBit("readEnable", (readEnable))
	if _readEnableErr != nil {
		return errors.Wrap(_readEnableErr, "Error serializing 'readEnable' field")
	}

	// Simple Field (communicationEnable)
	communicationEnable := bool(m.GetCommunicationEnable())
	_communicationEnableErr := writeBuffer.WriteBit("communicationEnable", (communicationEnable))
	if _communicationEnableErr != nil {
		return errors.Wrap(_communicationEnableErr, "Error serializing 'communicationEnable' field")
	}

	// Simple Field (priority)
	if pushErr := writeBuffer.PushContext("priority"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for priority")
	}
	_priorityErr := writeBuffer.WriteSerializable(m.GetPriority())
	if popErr := writeBuffer.PopContext("priority"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for priority")
	}
	if _priorityErr != nil {
		return errors.Wrap(_priorityErr, "Error serializing 'priority' field")
	}

	// Simple Field (valueType)
	if pushErr := writeBuffer.PushContext("valueType"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for valueType")
	}
	_valueTypeErr := writeBuffer.WriteSerializable(m.GetValueType())
	if popErr := writeBuffer.PopContext("valueType"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for valueType")
	}
	if _valueTypeErr != nil {
		return errors.Wrap(_valueTypeErr, "Error serializing 'valueType' field")
	}

	if popErr := writeBuffer.PopContext("GroupObjectDescriptorRealisationType7"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for GroupObjectDescriptorRealisationType7")
	}
	return nil
}

func (m *_GroupObjectDescriptorRealisationType7) isGroupObjectDescriptorRealisationType7() bool {
	return true
}

func (m *_GroupObjectDescriptorRealisationType7) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
