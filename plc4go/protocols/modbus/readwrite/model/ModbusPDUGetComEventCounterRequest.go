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

// ModbusPDUGetComEventCounterRequest is the corresponding interface of ModbusPDUGetComEventCounterRequest
type ModbusPDUGetComEventCounterRequest interface {
	utils.LengthAware
	utils.Serializable
	ModbusPDU
}

// ModbusPDUGetComEventCounterRequestExactly can be used when we want exactly this type and not a type which fulfills ModbusPDUGetComEventCounterRequest.
// This is useful for switch cases.
type ModbusPDUGetComEventCounterRequestExactly interface {
	isModbusPDUGetComEventCounterRequest() bool
}

// _ModbusPDUGetComEventCounterRequest is the data-structure of this message
type _ModbusPDUGetComEventCounterRequest struct {
	*_ModbusPDU
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusPDUGetComEventCounterRequest) GetErrorFlag() bool {
	return bool(false)
}

func (m *_ModbusPDUGetComEventCounterRequest) GetFunctionFlag() uint8 {
	return 0x0B
}

func (m *_ModbusPDUGetComEventCounterRequest) GetResponse() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusPDUGetComEventCounterRequest) InitializeParent(parent ModbusPDU) {}

func (m *_ModbusPDUGetComEventCounterRequest) GetParent() ModbusPDU {
	return m._ModbusPDU
}

// NewModbusPDUGetComEventCounterRequest factory function for _ModbusPDUGetComEventCounterRequest
func NewModbusPDUGetComEventCounterRequest() *_ModbusPDUGetComEventCounterRequest {
	_result := &_ModbusPDUGetComEventCounterRequest{
		_ModbusPDU: NewModbusPDU(),
	}
	_result._ModbusPDU._ModbusPDUChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastModbusPDUGetComEventCounterRequest(structType interface{}) ModbusPDUGetComEventCounterRequest {
	if casted, ok := structType.(ModbusPDUGetComEventCounterRequest); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUGetComEventCounterRequest); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUGetComEventCounterRequest) GetTypeName() string {
	return "ModbusPDUGetComEventCounterRequest"
}

func (m *_ModbusPDUGetComEventCounterRequest) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ModbusPDUGetComEventCounterRequest) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *_ModbusPDUGetComEventCounterRequest) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ModbusPDUGetComEventCounterRequestParse(readBuffer utils.ReadBuffer, response bool) (ModbusPDUGetComEventCounterRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUGetComEventCounterRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUGetComEventCounterRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ModbusPDUGetComEventCounterRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUGetComEventCounterRequest")
	}

	// Create a partially initialized instance
	_child := &_ModbusPDUGetComEventCounterRequest{
		_ModbusPDU: &_ModbusPDU{},
	}
	_child._ModbusPDU._ModbusPDUChildRequirements = _child
	return _child, nil
}

func (m *_ModbusPDUGetComEventCounterRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUGetComEventCounterRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusPDUGetComEventCounterRequest")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUGetComEventCounterRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusPDUGetComEventCounterRequest")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_ModbusPDUGetComEventCounterRequest) isModbusPDUGetComEventCounterRequest() bool {
	return true
}

func (m *_ModbusPDUGetComEventCounterRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
