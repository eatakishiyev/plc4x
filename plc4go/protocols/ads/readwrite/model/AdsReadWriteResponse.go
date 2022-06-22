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

// AdsReadWriteResponse is the corresponding interface of AdsReadWriteResponse
type AdsReadWriteResponse interface {
	utils.LengthAware
	utils.Serializable
	AdsData
	// GetResult returns Result (property field)
	GetResult() ReturnCode
	// GetData returns Data (property field)
	GetData() []byte
}

// AdsReadWriteResponseExactly can be used when we want exactly this type and not a type which fulfills AdsReadWriteResponse.
// This is useful for switch cases.
type AdsReadWriteResponseExactly interface {
	isAdsReadWriteResponse() bool
}

// _AdsReadWriteResponse is the data-structure of this message
type _AdsReadWriteResponse struct {
	*_AdsData
	Result ReturnCode
	Data   []byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AdsReadWriteResponse) GetCommandId() CommandId {
	return CommandId_ADS_READ_WRITE
}

func (m *_AdsReadWriteResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AdsReadWriteResponse) InitializeParent(parent AdsData) {}

func (m *_AdsReadWriteResponse) GetParent() AdsData {
	return m._AdsData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AdsReadWriteResponse) GetResult() ReturnCode {
	return m.Result
}

func (m *_AdsReadWriteResponse) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAdsReadWriteResponse factory function for _AdsReadWriteResponse
func NewAdsReadWriteResponse(result ReturnCode, data []byte) *_AdsReadWriteResponse {
	_result := &_AdsReadWriteResponse{
		Result:   result,
		Data:     data,
		_AdsData: NewAdsData(),
	}
	_result._AdsData._AdsDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAdsReadWriteResponse(structType interface{}) AdsReadWriteResponse {
	if casted, ok := structType.(AdsReadWriteResponse); ok {
		return casted
	}
	if casted, ok := structType.(*AdsReadWriteResponse); ok {
		return *casted
	}
	return nil
}

func (m *_AdsReadWriteResponse) GetTypeName() string {
	return "AdsReadWriteResponse"
}

func (m *_AdsReadWriteResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_AdsReadWriteResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (result)
	lengthInBits += 32

	// Implicit Field (length)
	lengthInBits += 32

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *_AdsReadWriteResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func AdsReadWriteResponseParse(readBuffer utils.ReadBuffer, commandId CommandId, response bool) (AdsReadWriteResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsReadWriteResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsReadWriteResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (result)
	if pullErr := readBuffer.PullContext("result"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for result")
	}
	_result, _resultErr := ReturnCodeParse(readBuffer)
	if _resultErr != nil {
		return nil, errors.Wrap(_resultErr, "Error parsing 'result' field")
	}
	result := _result
	if closeErr := readBuffer.CloseContext("result"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for result")
	}

	// Implicit Field (length) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	length, _lengthErr := readBuffer.ReadUint32("length", 32)
	_ = length
	if _lengthErr != nil {
		return nil, errors.Wrap(_lengthErr, "Error parsing 'length' field")
	}
	// Byte Array field (data)
	numberOfBytesdata := int(length)
	data, _readArrayErr := readBuffer.ReadByteArray("data", numberOfBytesdata)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'data' field")
	}

	if closeErr := readBuffer.CloseContext("AdsReadWriteResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsReadWriteResponse")
	}

	// Create a partially initialized instance
	_child := &_AdsReadWriteResponse{
		Result:   result,
		Data:     data,
		_AdsData: &_AdsData{},
	}
	_child._AdsData._AdsDataChildRequirements = _child
	return _child, nil
}

func (m *_AdsReadWriteResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsReadWriteResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AdsReadWriteResponse")
		}

		// Simple Field (result)
		if pushErr := writeBuffer.PushContext("result"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for result")
		}
		_resultErr := writeBuffer.WriteSerializable(m.GetResult())
		if popErr := writeBuffer.PopContext("result"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for result")
		}
		if _resultErr != nil {
			return errors.Wrap(_resultErr, "Error serializing 'result' field")
		}

		// Implicit Field (length) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		length := uint32(uint32(len(m.GetData())))
		_lengthErr := writeBuffer.WriteUint32("length", 32, (length))
		if _lengthErr != nil {
			return errors.Wrap(_lengthErr, "Error serializing 'length' field")
		}

		// Array Field (data)
		if m.GetData() != nil {
			// Byte Array field (data)
			_writeArrayErr := writeBuffer.WriteByteArray("data", m.GetData())
			if _writeArrayErr != nil {
				return errors.Wrap(_writeArrayErr, "Error serializing 'data' field")
			}
		}

		if popErr := writeBuffer.PopContext("AdsReadWriteResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AdsReadWriteResponse")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_AdsReadWriteResponse) isAdsReadWriteResponse() bool {
	return true
}

func (m *_AdsReadWriteResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
