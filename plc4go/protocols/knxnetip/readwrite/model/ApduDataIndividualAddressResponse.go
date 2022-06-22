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

// ApduDataIndividualAddressResponse is the corresponding interface of ApduDataIndividualAddressResponse
type ApduDataIndividualAddressResponse interface {
	utils.LengthAware
	utils.Serializable
	ApduData
}

// ApduDataIndividualAddressResponseExactly can be used when we want exactly this type and not a type which fulfills ApduDataIndividualAddressResponse.
// This is useful for switch cases.
type ApduDataIndividualAddressResponseExactly interface {
	isApduDataIndividualAddressResponse() bool
}

// _ApduDataIndividualAddressResponse is the data-structure of this message
type _ApduDataIndividualAddressResponse struct {
	*_ApduData

	// Arguments.
	DataLength uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataIndividualAddressResponse) GetApciType() uint8 {
	return 0x5
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataIndividualAddressResponse) InitializeParent(parent ApduData) {}

func (m *_ApduDataIndividualAddressResponse) GetParent() ApduData {
	return m._ApduData
}

// NewApduDataIndividualAddressResponse factory function for _ApduDataIndividualAddressResponse
func NewApduDataIndividualAddressResponse(dataLength uint8) *_ApduDataIndividualAddressResponse {
	_result := &_ApduDataIndividualAddressResponse{
		_ApduData: NewApduData(dataLength),
	}
	_result._ApduData._ApduDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastApduDataIndividualAddressResponse(structType interface{}) ApduDataIndividualAddressResponse {
	if casted, ok := structType.(ApduDataIndividualAddressResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataIndividualAddressResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataIndividualAddressResponse) GetTypeName() string {
	return "ApduDataIndividualAddressResponse"
}

func (m *_ApduDataIndividualAddressResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ApduDataIndividualAddressResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *_ApduDataIndividualAddressResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ApduDataIndividualAddressResponseParse(readBuffer utils.ReadBuffer, dataLength uint8) (ApduDataIndividualAddressResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataIndividualAddressResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataIndividualAddressResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataIndividualAddressResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataIndividualAddressResponse")
	}

	// Create a partially initialized instance
	_child := &_ApduDataIndividualAddressResponse{
		_ApduData: &_ApduData{},
	}
	_child._ApduData._ApduDataChildRequirements = _child
	return _child, nil
}

func (m *_ApduDataIndividualAddressResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataIndividualAddressResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataIndividualAddressResponse")
		}

		if popErr := writeBuffer.PopContext("ApduDataIndividualAddressResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataIndividualAddressResponse")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_ApduDataIndividualAddressResponse) isApduDataIndividualAddressResponse() bool {
	return true
}

func (m *_ApduDataIndividualAddressResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
