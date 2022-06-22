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

// S7PayloadWriteVarResponse is the corresponding interface of S7PayloadWriteVarResponse
type S7PayloadWriteVarResponse interface {
	utils.LengthAware
	utils.Serializable
	S7Payload
	// GetItems returns Items (property field)
	GetItems() []S7VarPayloadStatusItem
}

// S7PayloadWriteVarResponseExactly can be used when we want exactly this type and not a type which fulfills S7PayloadWriteVarResponse.
// This is useful for switch cases.
type S7PayloadWriteVarResponseExactly interface {
	isS7PayloadWriteVarResponse() bool
}

// _S7PayloadWriteVarResponse is the data-structure of this message
type _S7PayloadWriteVarResponse struct {
	*_S7Payload
	Items []S7VarPayloadStatusItem

	// Arguments.
	Parameter S7Parameter
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7PayloadWriteVarResponse) GetParameterParameterType() uint8 {
	return 0x05
}

func (m *_S7PayloadWriteVarResponse) GetMessageType() uint8 {
	return 0x03
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7PayloadWriteVarResponse) InitializeParent(parent S7Payload) {}

func (m *_S7PayloadWriteVarResponse) GetParent() S7Payload {
	return m._S7Payload
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7PayloadWriteVarResponse) GetItems() []S7VarPayloadStatusItem {
	return m.Items
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewS7PayloadWriteVarResponse factory function for _S7PayloadWriteVarResponse
func NewS7PayloadWriteVarResponse(items []S7VarPayloadStatusItem, parameter S7Parameter) *_S7PayloadWriteVarResponse {
	_result := &_S7PayloadWriteVarResponse{
		Items:      items,
		_S7Payload: NewS7Payload(parameter),
	}
	_result._S7Payload._S7PayloadChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastS7PayloadWriteVarResponse(structType interface{}) S7PayloadWriteVarResponse {
	if casted, ok := structType.(S7PayloadWriteVarResponse); ok {
		return casted
	}
	if casted, ok := structType.(*S7PayloadWriteVarResponse); ok {
		return *casted
	}
	return nil
}

func (m *_S7PayloadWriteVarResponse) GetTypeName() string {
	return "S7PayloadWriteVarResponse"
}

func (m *_S7PayloadWriteVarResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_S7PayloadWriteVarResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.Items) > 0 {
		for i, element := range m.Items {
			last := i == len(m.Items)-1
			lengthInBits += element.(interface{ GetLengthInBitsConditional(bool) uint16 }).GetLengthInBitsConditional(last)
		}
	}

	return lengthInBits
}

func (m *_S7PayloadWriteVarResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func S7PayloadWriteVarResponseParse(readBuffer utils.ReadBuffer, messageType uint8, parameter S7Parameter) (S7PayloadWriteVarResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7PayloadWriteVarResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7PayloadWriteVarResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (items)
	if pullErr := readBuffer.PullContext("items", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for items")
	}
	// Count array
	items := make([]S7VarPayloadStatusItem, CastS7ParameterWriteVarResponse(parameter).GetNumItems())
	{
		for curItem := uint16(0); curItem < uint16(CastS7ParameterWriteVarResponse(parameter).GetNumItems()); curItem++ {
			_item, _err := S7VarPayloadStatusItemParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'items' field")
			}
			items[curItem] = _item.(S7VarPayloadStatusItem)
		}
	}
	if closeErr := readBuffer.CloseContext("items", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for items")
	}

	if closeErr := readBuffer.CloseContext("S7PayloadWriteVarResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7PayloadWriteVarResponse")
	}

	// Create a partially initialized instance
	_child := &_S7PayloadWriteVarResponse{
		Items:      items,
		_S7Payload: &_S7Payload{},
	}
	_child._S7Payload._S7PayloadChildRequirements = _child
	return _child, nil
}

func (m *_S7PayloadWriteVarResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7PayloadWriteVarResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7PayloadWriteVarResponse")
		}

		// Array Field (items)
		if m.GetItems() != nil {
			if pushErr := writeBuffer.PushContext("items", utils.WithRenderAsList(true)); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for items")
			}
			for _, _element := range m.GetItems() {
				_elementErr := writeBuffer.WriteSerializable(_element)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'items' field")
				}
			}
			if popErr := writeBuffer.PopContext("items", utils.WithRenderAsList(true)); popErr != nil {
				return errors.Wrap(popErr, "Error popping for items")
			}
		}

		if popErr := writeBuffer.PopContext("S7PayloadWriteVarResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7PayloadWriteVarResponse")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_S7PayloadWriteVarResponse) isS7PayloadWriteVarResponse() bool {
	return true
}

func (m *_S7PayloadWriteVarResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
