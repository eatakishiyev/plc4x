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

// COTPParameterChecksum is the corresponding interface of COTPParameterChecksum
type COTPParameterChecksum interface {
	utils.LengthAware
	utils.Serializable
	COTPParameter
	// GetCrc returns Crc (property field)
	GetCrc() uint8
}

// COTPParameterChecksumExactly can be used when we want exactly this type and not a type which fulfills COTPParameterChecksum.
// This is useful for switch cases.
type COTPParameterChecksumExactly interface {
	isCOTPParameterChecksum() bool
}

// _COTPParameterChecksum is the data-structure of this message
type _COTPParameterChecksum struct {
	*_COTPParameter
	Crc uint8

	// Arguments.
	Rest uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_COTPParameterChecksum) GetParameterType() uint8 {
	return 0xC3
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_COTPParameterChecksum) InitializeParent(parent COTPParameter) {}

func (m *_COTPParameterChecksum) GetParent() COTPParameter {
	return m._COTPParameter
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_COTPParameterChecksum) GetCrc() uint8 {
	return m.Crc
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCOTPParameterChecksum factory function for _COTPParameterChecksum
func NewCOTPParameterChecksum(crc uint8, rest uint8) *_COTPParameterChecksum {
	_result := &_COTPParameterChecksum{
		Crc:            crc,
		_COTPParameter: NewCOTPParameter(rest),
	}
	_result._COTPParameter._COTPParameterChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCOTPParameterChecksum(structType interface{}) COTPParameterChecksum {
	if casted, ok := structType.(COTPParameterChecksum); ok {
		return casted
	}
	if casted, ok := structType.(*COTPParameterChecksum); ok {
		return *casted
	}
	return nil
}

func (m *_COTPParameterChecksum) GetTypeName() string {
	return "COTPParameterChecksum"
}

func (m *_COTPParameterChecksum) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_COTPParameterChecksum) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (crc)
	lengthInBits += 8

	return lengthInBits
}

func (m *_COTPParameterChecksum) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func COTPParameterChecksumParse(readBuffer utils.ReadBuffer, rest uint8) (COTPParameterChecksum, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("COTPParameterChecksum"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for COTPParameterChecksum")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (crc)
	_crc, _crcErr := readBuffer.ReadUint8("crc", 8)
	if _crcErr != nil {
		return nil, errors.Wrap(_crcErr, "Error parsing 'crc' field")
	}
	crc := _crc

	if closeErr := readBuffer.CloseContext("COTPParameterChecksum"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for COTPParameterChecksum")
	}

	// Create a partially initialized instance
	_child := &_COTPParameterChecksum{
		Crc:            crc,
		_COTPParameter: &_COTPParameter{},
	}
	_child._COTPParameter._COTPParameterChildRequirements = _child
	return _child, nil
}

func (m *_COTPParameterChecksum) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("COTPParameterChecksum"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for COTPParameterChecksum")
		}

		// Simple Field (crc)
		crc := uint8(m.GetCrc())
		_crcErr := writeBuffer.WriteUint8("crc", 8, (crc))
		if _crcErr != nil {
			return errors.Wrap(_crcErr, "Error serializing 'crc' field")
		}

		if popErr := writeBuffer.PopContext("COTPParameterChecksum"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for COTPParameterChecksum")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_COTPParameterChecksum) isCOTPParameterChecksum() bool {
	return true
}

func (m *_COTPParameterChecksum) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
