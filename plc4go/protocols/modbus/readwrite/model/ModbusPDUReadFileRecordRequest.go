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

// ModbusPDUReadFileRecordRequest is the corresponding interface of ModbusPDUReadFileRecordRequest
type ModbusPDUReadFileRecordRequest interface {
	utils.LengthAware
	utils.Serializable
	ModbusPDU
	// GetItems returns Items (property field)
	GetItems() []ModbusPDUReadFileRecordRequestItem
}

// ModbusPDUReadFileRecordRequestExactly can be used when we want exactly this type and not a type which fulfills ModbusPDUReadFileRecordRequest.
// This is useful for switch cases.
type ModbusPDUReadFileRecordRequestExactly interface {
	isModbusPDUReadFileRecordRequest() bool
}

// _ModbusPDUReadFileRecordRequest is the data-structure of this message
type _ModbusPDUReadFileRecordRequest struct {
	*_ModbusPDU
	Items []ModbusPDUReadFileRecordRequestItem
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusPDUReadFileRecordRequest) GetErrorFlag() bool {
	return bool(false)
}

func (m *_ModbusPDUReadFileRecordRequest) GetFunctionFlag() uint8 {
	return 0x14
}

func (m *_ModbusPDUReadFileRecordRequest) GetResponse() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusPDUReadFileRecordRequest) InitializeParent(parent ModbusPDU) {}

func (m *_ModbusPDUReadFileRecordRequest) GetParent() ModbusPDU {
	return m._ModbusPDU
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusPDUReadFileRecordRequest) GetItems() []ModbusPDUReadFileRecordRequestItem {
	return m.Items
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewModbusPDUReadFileRecordRequest factory function for _ModbusPDUReadFileRecordRequest
func NewModbusPDUReadFileRecordRequest(items []ModbusPDUReadFileRecordRequestItem) *_ModbusPDUReadFileRecordRequest {
	_result := &_ModbusPDUReadFileRecordRequest{
		Items:      items,
		_ModbusPDU: NewModbusPDU(),
	}
	_result._ModbusPDU._ModbusPDUChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastModbusPDUReadFileRecordRequest(structType interface{}) ModbusPDUReadFileRecordRequest {
	if casted, ok := structType.(ModbusPDUReadFileRecordRequest); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUReadFileRecordRequest); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUReadFileRecordRequest) GetTypeName() string {
	return "ModbusPDUReadFileRecordRequest"
}

func (m *_ModbusPDUReadFileRecordRequest) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ModbusPDUReadFileRecordRequest) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Implicit Field (byteCount)
	lengthInBits += 8

	// Array field
	if len(m.Items) > 0 {
		for _, element := range m.Items {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *_ModbusPDUReadFileRecordRequest) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ModbusPDUReadFileRecordRequestParse(readBuffer utils.ReadBuffer, response bool) (ModbusPDUReadFileRecordRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUReadFileRecordRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUReadFileRecordRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Implicit Field (byteCount) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	byteCount, _byteCountErr := readBuffer.ReadUint8("byteCount", 8)
	_ = byteCount
	if _byteCountErr != nil {
		return nil, errors.Wrap(_byteCountErr, "Error parsing 'byteCount' field")
	}

	// Array field (items)
	if pullErr := readBuffer.PullContext("items", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for items")
	}
	// Length array
	items := make([]ModbusPDUReadFileRecordRequestItem, 0)
	{
		_itemsLength := byteCount
		_itemsEndPos := positionAware.GetPos() + uint16(_itemsLength)
		for positionAware.GetPos() < _itemsEndPos {
			_item, _err := ModbusPDUReadFileRecordRequestItemParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'items' field")
			}
			items = append(items, _item.(ModbusPDUReadFileRecordRequestItem))
		}
	}
	if closeErr := readBuffer.CloseContext("items", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for items")
	}

	if closeErr := readBuffer.CloseContext("ModbusPDUReadFileRecordRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUReadFileRecordRequest")
	}

	// Create a partially initialized instance
	_child := &_ModbusPDUReadFileRecordRequest{
		Items:      items,
		_ModbusPDU: &_ModbusPDU{},
	}
	_child._ModbusPDU._ModbusPDUChildRequirements = _child
	return _child, nil
}

func (m *_ModbusPDUReadFileRecordRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	itemsArraySizeInBytes := func(items []ModbusPDUReadFileRecordRequestItem) uint32 {
		var sizeInBytes uint32 = 0
		for _, v := range items {
			sizeInBytes += uint32(v.GetLengthInBytes())
		}
		return sizeInBytes
	}
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUReadFileRecordRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusPDUReadFileRecordRequest")
		}

		// Implicit Field (byteCount) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		byteCount := uint8(uint8(itemsArraySizeInBytes(m.GetItems())))
		_byteCountErr := writeBuffer.WriteUint8("byteCount", 8, (byteCount))
		if _byteCountErr != nil {
			return errors.Wrap(_byteCountErr, "Error serializing 'byteCount' field")
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

		if popErr := writeBuffer.PopContext("ModbusPDUReadFileRecordRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusPDUReadFileRecordRequest")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_ModbusPDUReadFileRecordRequest) isModbusPDUReadFileRecordRequest() bool {
	return true
}

func (m *_ModbusPDUReadFileRecordRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
