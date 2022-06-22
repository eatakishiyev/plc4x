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
	"github.com/rs/zerolog/log"
)

// Code generated by code-generation. DO NOT EDIT.

// ExtendedStatusHeader is the corresponding interface of ExtendedStatusHeader
type ExtendedStatusHeader interface {
	utils.LengthAware
	utils.Serializable
	// GetNumberOfCharacterPairs returns NumberOfCharacterPairs (property field)
	GetNumberOfCharacterPairs() uint8
}

// ExtendedStatusHeaderExactly can be used when we want exactly this type and not a type which fulfills ExtendedStatusHeader.
// This is useful for switch cases.
type ExtendedStatusHeaderExactly interface {
	isExtendedStatusHeader() bool
}

// _ExtendedStatusHeader is the data-structure of this message
type _ExtendedStatusHeader struct {
	NumberOfCharacterPairs uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ExtendedStatusHeader) GetNumberOfCharacterPairs() uint8 {
	return m.NumberOfCharacterPairs
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewExtendedStatusHeader factory function for _ExtendedStatusHeader
func NewExtendedStatusHeader(numberOfCharacterPairs uint8) *_ExtendedStatusHeader {
	return &_ExtendedStatusHeader{NumberOfCharacterPairs: numberOfCharacterPairs}
}

// Deprecated: use the interface for direct cast
func CastExtendedStatusHeader(structType interface{}) ExtendedStatusHeader {
	if casted, ok := structType.(ExtendedStatusHeader); ok {
		return casted
	}
	if casted, ok := structType.(*ExtendedStatusHeader); ok {
		return *casted
	}
	return nil
}

func (m *_ExtendedStatusHeader) GetTypeName() string {
	return "ExtendedStatusHeader"
}

func (m *_ExtendedStatusHeader) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ExtendedStatusHeader) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Reserved Field (reserved)
	lengthInBits += 3

	// Simple field (numberOfCharacterPairs)
	lengthInBits += 5

	return lengthInBits
}

func (m *_ExtendedStatusHeader) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ExtendedStatusHeaderParse(readBuffer utils.ReadBuffer) (ExtendedStatusHeader, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ExtendedStatusHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ExtendedStatusHeader")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 3)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != uint8(0x7) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint8(0x7),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	// Simple Field (numberOfCharacterPairs)
	_numberOfCharacterPairs, _numberOfCharacterPairsErr := readBuffer.ReadUint8("numberOfCharacterPairs", 5)
	if _numberOfCharacterPairsErr != nil {
		return nil, errors.Wrap(_numberOfCharacterPairsErr, "Error parsing 'numberOfCharacterPairs' field")
	}
	numberOfCharacterPairs := _numberOfCharacterPairs

	if closeErr := readBuffer.CloseContext("ExtendedStatusHeader"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ExtendedStatusHeader")
	}

	// Create the instance
	return NewExtendedStatusHeader(numberOfCharacterPairs), nil
}

func (m *_ExtendedStatusHeader) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("ExtendedStatusHeader"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ExtendedStatusHeader")
	}

	// Reserved Field (reserved)
	{
		_err := writeBuffer.WriteUint8("reserved", 3, uint8(0x7))
		if _err != nil {
			return errors.Wrap(_err, "Error serializing 'reserved' field")
		}
	}

	// Simple Field (numberOfCharacterPairs)
	numberOfCharacterPairs := uint8(m.GetNumberOfCharacterPairs())
	_numberOfCharacterPairsErr := writeBuffer.WriteUint8("numberOfCharacterPairs", 5, (numberOfCharacterPairs))
	if _numberOfCharacterPairsErr != nil {
		return errors.Wrap(_numberOfCharacterPairsErr, "Error serializing 'numberOfCharacterPairs' field")
	}

	if popErr := writeBuffer.PopContext("ExtendedStatusHeader"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ExtendedStatusHeader")
	}
	return nil
}

func (m *_ExtendedStatusHeader) isExtendedStatusHeader() bool {
	return true
}

func (m *_ExtendedStatusHeader) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
