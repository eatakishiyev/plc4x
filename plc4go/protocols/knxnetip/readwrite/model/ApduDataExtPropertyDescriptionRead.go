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

// ApduDataExtPropertyDescriptionRead is the corresponding interface of ApduDataExtPropertyDescriptionRead
type ApduDataExtPropertyDescriptionRead interface {
	utils.LengthAware
	utils.Serializable
	ApduDataExt
	// GetObjectIndex returns ObjectIndex (property field)
	GetObjectIndex() uint8
	// GetPropertyId returns PropertyId (property field)
	GetPropertyId() uint8
	// GetIndex returns Index (property field)
	GetIndex() uint8
}

// ApduDataExtPropertyDescriptionReadExactly can be used when we want exactly this type and not a type which fulfills ApduDataExtPropertyDescriptionRead.
// This is useful for switch cases.
type ApduDataExtPropertyDescriptionReadExactly interface {
	isApduDataExtPropertyDescriptionRead() bool
}

// _ApduDataExtPropertyDescriptionRead is the data-structure of this message
type _ApduDataExtPropertyDescriptionRead struct {
	*_ApduDataExt
	ObjectIndex uint8
	PropertyId  uint8
	Index       uint8

	// Arguments.
	Length uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataExtPropertyDescriptionRead) GetExtApciType() uint8 {
	return 0x18
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataExtPropertyDescriptionRead) InitializeParent(parent ApduDataExt) {}

func (m *_ApduDataExtPropertyDescriptionRead) GetParent() ApduDataExt {
	return m._ApduDataExt
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ApduDataExtPropertyDescriptionRead) GetObjectIndex() uint8 {
	return m.ObjectIndex
}

func (m *_ApduDataExtPropertyDescriptionRead) GetPropertyId() uint8 {
	return m.PropertyId
}

func (m *_ApduDataExtPropertyDescriptionRead) GetIndex() uint8 {
	return m.Index
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewApduDataExtPropertyDescriptionRead factory function for _ApduDataExtPropertyDescriptionRead
func NewApduDataExtPropertyDescriptionRead(objectIndex uint8, propertyId uint8, index uint8, length uint8) *_ApduDataExtPropertyDescriptionRead {
	_result := &_ApduDataExtPropertyDescriptionRead{
		ObjectIndex:  objectIndex,
		PropertyId:   propertyId,
		Index:        index,
		_ApduDataExt: NewApduDataExt(length),
	}
	_result._ApduDataExt._ApduDataExtChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastApduDataExtPropertyDescriptionRead(structType interface{}) ApduDataExtPropertyDescriptionRead {
	if casted, ok := structType.(ApduDataExtPropertyDescriptionRead); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataExtPropertyDescriptionRead); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataExtPropertyDescriptionRead) GetTypeName() string {
	return "ApduDataExtPropertyDescriptionRead"
}

func (m *_ApduDataExtPropertyDescriptionRead) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ApduDataExtPropertyDescriptionRead) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (objectIndex)
	lengthInBits += 8

	// Simple field (propertyId)
	lengthInBits += 8

	// Simple field (index)
	lengthInBits += 8

	return lengthInBits
}

func (m *_ApduDataExtPropertyDescriptionRead) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ApduDataExtPropertyDescriptionReadParse(readBuffer utils.ReadBuffer, length uint8) (ApduDataExtPropertyDescriptionRead, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtPropertyDescriptionRead"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtPropertyDescriptionRead")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (objectIndex)
	_objectIndex, _objectIndexErr := readBuffer.ReadUint8("objectIndex", 8)
	if _objectIndexErr != nil {
		return nil, errors.Wrap(_objectIndexErr, "Error parsing 'objectIndex' field")
	}
	objectIndex := _objectIndex

	// Simple Field (propertyId)
	_propertyId, _propertyIdErr := readBuffer.ReadUint8("propertyId", 8)
	if _propertyIdErr != nil {
		return nil, errors.Wrap(_propertyIdErr, "Error parsing 'propertyId' field")
	}
	propertyId := _propertyId

	// Simple Field (index)
	_index, _indexErr := readBuffer.ReadUint8("index", 8)
	if _indexErr != nil {
		return nil, errors.Wrap(_indexErr, "Error parsing 'index' field")
	}
	index := _index

	if closeErr := readBuffer.CloseContext("ApduDataExtPropertyDescriptionRead"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtPropertyDescriptionRead")
	}

	// Create a partially initialized instance
	_child := &_ApduDataExtPropertyDescriptionRead{
		ObjectIndex:  objectIndex,
		PropertyId:   propertyId,
		Index:        index,
		_ApduDataExt: &_ApduDataExt{},
	}
	_child._ApduDataExt._ApduDataExtChildRequirements = _child
	return _child, nil
}

func (m *_ApduDataExtPropertyDescriptionRead) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtPropertyDescriptionRead"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtPropertyDescriptionRead")
		}

		// Simple Field (objectIndex)
		objectIndex := uint8(m.GetObjectIndex())
		_objectIndexErr := writeBuffer.WriteUint8("objectIndex", 8, (objectIndex))
		if _objectIndexErr != nil {
			return errors.Wrap(_objectIndexErr, "Error serializing 'objectIndex' field")
		}

		// Simple Field (propertyId)
		propertyId := uint8(m.GetPropertyId())
		_propertyIdErr := writeBuffer.WriteUint8("propertyId", 8, (propertyId))
		if _propertyIdErr != nil {
			return errors.Wrap(_propertyIdErr, "Error serializing 'propertyId' field")
		}

		// Simple Field (index)
		index := uint8(m.GetIndex())
		_indexErr := writeBuffer.WriteUint8("index", 8, (index))
		if _indexErr != nil {
			return errors.Wrap(_indexErr, "Error serializing 'index' field")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtPropertyDescriptionRead"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtPropertyDescriptionRead")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_ApduDataExtPropertyDescriptionRead) isApduDataExtPropertyDescriptionRead() bool {
	return true
}

func (m *_ApduDataExtPropertyDescriptionRead) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
