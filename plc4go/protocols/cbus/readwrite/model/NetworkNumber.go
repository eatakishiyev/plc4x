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

// NetworkNumber is the corresponding interface of NetworkNumber
type NetworkNumber interface {
	utils.LengthAware
	utils.Serializable
	// GetNumber returns Number (property field)
	GetNumber() uint8
}

// NetworkNumberExactly can be used when we want exactly this type and not a type which fulfills NetworkNumber.
// This is useful for switch cases.
type NetworkNumberExactly interface {
	isNetworkNumber() bool
}

// _NetworkNumber is the data-structure of this message
type _NetworkNumber struct {
	Number uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NetworkNumber) GetNumber() uint8 {
	return m.Number
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewNetworkNumber factory function for _NetworkNumber
func NewNetworkNumber(number uint8) *_NetworkNumber {
	return &_NetworkNumber{Number: number}
}

// Deprecated: use the interface for direct cast
func CastNetworkNumber(structType interface{}) NetworkNumber {
	if casted, ok := structType.(NetworkNumber); ok {
		return casted
	}
	if casted, ok := structType.(*NetworkNumber); ok {
		return *casted
	}
	return nil
}

func (m *_NetworkNumber) GetTypeName() string {
	return "NetworkNumber"
}

func (m *_NetworkNumber) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_NetworkNumber) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (number)
	lengthInBits += 8

	return lengthInBits
}

func (m *_NetworkNumber) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func NetworkNumberParse(readBuffer utils.ReadBuffer) (NetworkNumber, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NetworkNumber"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NetworkNumber")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (number)
	_number, _numberErr := readBuffer.ReadUint8("number", 8)
	if _numberErr != nil {
		return nil, errors.Wrap(_numberErr, "Error parsing 'number' field")
	}
	number := _number

	if closeErr := readBuffer.CloseContext("NetworkNumber"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NetworkNumber")
	}

	// Create the instance
	return NewNetworkNumber(number), nil
}

func (m *_NetworkNumber) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("NetworkNumber"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for NetworkNumber")
	}

	// Simple Field (number)
	number := uint8(m.GetNumber())
	_numberErr := writeBuffer.WriteUint8("number", 8, (number))
	if _numberErr != nil {
		return errors.Wrap(_numberErr, "Error serializing 'number' field")
	}

	if popErr := writeBuffer.PopContext("NetworkNumber"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for NetworkNumber")
	}
	return nil
}

func (m *_NetworkNumber) isNetworkNumber() bool {
	return true
}

func (m *_NetworkNumber) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
