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

// ApduDataDeviceDescriptorRead is the corresponding interface of ApduDataDeviceDescriptorRead
type ApduDataDeviceDescriptorRead interface {
	utils.LengthAware
	utils.Serializable
	ApduData
	// GetDescriptorType returns DescriptorType (property field)
	GetDescriptorType() uint8
}

// ApduDataDeviceDescriptorReadExactly can be used when we want exactly this type and not a type which fulfills ApduDataDeviceDescriptorRead.
// This is useful for switch cases.
type ApduDataDeviceDescriptorReadExactly interface {
	ApduDataDeviceDescriptorRead
	isApduDataDeviceDescriptorRead() bool
}

// _ApduDataDeviceDescriptorRead is the data-structure of this message
type _ApduDataDeviceDescriptorRead struct {
	*_ApduData
	DescriptorType uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataDeviceDescriptorRead) GetApciType() uint8 {
	return 0xC
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataDeviceDescriptorRead) InitializeParent(parent ApduData) {}

func (m *_ApduDataDeviceDescriptorRead) GetParent() ApduData {
	return m._ApduData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ApduDataDeviceDescriptorRead) GetDescriptorType() uint8 {
	return m.DescriptorType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewApduDataDeviceDescriptorRead factory function for _ApduDataDeviceDescriptorRead
func NewApduDataDeviceDescriptorRead(descriptorType uint8, dataLength uint8) *_ApduDataDeviceDescriptorRead {
	_result := &_ApduDataDeviceDescriptorRead{
		DescriptorType: descriptorType,
		_ApduData:      NewApduData(dataLength),
	}
	_result._ApduData._ApduDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastApduDataDeviceDescriptorRead(structType interface{}) ApduDataDeviceDescriptorRead {
	if casted, ok := structType.(ApduDataDeviceDescriptorRead); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataDeviceDescriptorRead); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataDeviceDescriptorRead) GetTypeName() string {
	return "ApduDataDeviceDescriptorRead"
}

func (m *_ApduDataDeviceDescriptorRead) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ApduDataDeviceDescriptorRead) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (descriptorType)
	lengthInBits += 6

	return lengthInBits
}

func (m *_ApduDataDeviceDescriptorRead) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ApduDataDeviceDescriptorReadParse(readBuffer utils.ReadBuffer, dataLength uint8) (ApduDataDeviceDescriptorRead, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataDeviceDescriptorRead"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataDeviceDescriptorRead")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (descriptorType)
	_descriptorType, _descriptorTypeErr := readBuffer.ReadUint8("descriptorType", 6)
	if _descriptorTypeErr != nil {
		return nil, errors.Wrap(_descriptorTypeErr, "Error parsing 'descriptorType' field")
	}
	descriptorType := _descriptorType

	if closeErr := readBuffer.CloseContext("ApduDataDeviceDescriptorRead"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataDeviceDescriptorRead")
	}

	// Create a partially initialized instance
	_child := &_ApduDataDeviceDescriptorRead{
		DescriptorType: descriptorType,
		_ApduData: &_ApduData{
			DataLength: dataLength,
		},
	}
	_child._ApduData._ApduDataChildRequirements = _child
	return _child, nil
}

func (m *_ApduDataDeviceDescriptorRead) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataDeviceDescriptorRead"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataDeviceDescriptorRead")
		}

		// Simple Field (descriptorType)
		descriptorType := uint8(m.GetDescriptorType())
		_descriptorTypeErr := writeBuffer.WriteUint8("descriptorType", 6, (descriptorType))
		if _descriptorTypeErr != nil {
			return errors.Wrap(_descriptorTypeErr, "Error serializing 'descriptorType' field")
		}

		if popErr := writeBuffer.PopContext("ApduDataDeviceDescriptorRead"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataDeviceDescriptorRead")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_ApduDataDeviceDescriptorRead) isApduDataDeviceDescriptorRead() bool {
	return true
}

func (m *_ApduDataDeviceDescriptorRead) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
