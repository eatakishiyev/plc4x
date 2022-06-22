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

// SzlId is the corresponding interface of SzlId
type SzlId interface {
	utils.LengthAware
	utils.Serializable
	// GetTypeClass returns TypeClass (property field)
	GetTypeClass() SzlModuleTypeClass
	// GetSublistExtract returns SublistExtract (property field)
	GetSublistExtract() uint8
	// GetSublistList returns SublistList (property field)
	GetSublistList() SzlSublist
}

// SzlIdExactly can be used when we want exactly this type and not a type which fulfills SzlId.
// This is useful for switch cases.
type SzlIdExactly interface {
	isSzlId() bool
}

// _SzlId is the data-structure of this message
type _SzlId struct {
	TypeClass      SzlModuleTypeClass
	SublistExtract uint8
	SublistList    SzlSublist
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SzlId) GetTypeClass() SzlModuleTypeClass {
	return m.TypeClass
}

func (m *_SzlId) GetSublistExtract() uint8 {
	return m.SublistExtract
}

func (m *_SzlId) GetSublistList() SzlSublist {
	return m.SublistList
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSzlId factory function for _SzlId
func NewSzlId(typeClass SzlModuleTypeClass, sublistExtract uint8, sublistList SzlSublist) *_SzlId {
	return &_SzlId{TypeClass: typeClass, SublistExtract: sublistExtract, SublistList: sublistList}
}

// Deprecated: use the interface for direct cast
func CastSzlId(structType interface{}) SzlId {
	if casted, ok := structType.(SzlId); ok {
		return casted
	}
	if casted, ok := structType.(*SzlId); ok {
		return *casted
	}
	return nil
}

func (m *_SzlId) GetTypeName() string {
	return "SzlId"
}

func (m *_SzlId) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_SzlId) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (typeClass)
	lengthInBits += 4

	// Simple field (sublistExtract)
	lengthInBits += 4

	// Simple field (sublistList)
	lengthInBits += 8

	return lengthInBits
}

func (m *_SzlId) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func SzlIdParse(readBuffer utils.ReadBuffer) (SzlId, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SzlId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SzlId")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (typeClass)
	if pullErr := readBuffer.PullContext("typeClass"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for typeClass")
	}
	_typeClass, _typeClassErr := SzlModuleTypeClassParse(readBuffer)
	if _typeClassErr != nil {
		return nil, errors.Wrap(_typeClassErr, "Error parsing 'typeClass' field")
	}
	typeClass := _typeClass
	if closeErr := readBuffer.CloseContext("typeClass"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for typeClass")
	}

	// Simple Field (sublistExtract)
	_sublistExtract, _sublistExtractErr := readBuffer.ReadUint8("sublistExtract", 4)
	if _sublistExtractErr != nil {
		return nil, errors.Wrap(_sublistExtractErr, "Error parsing 'sublistExtract' field")
	}
	sublistExtract := _sublistExtract

	// Simple Field (sublistList)
	if pullErr := readBuffer.PullContext("sublistList"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for sublistList")
	}
	_sublistList, _sublistListErr := SzlSublistParse(readBuffer)
	if _sublistListErr != nil {
		return nil, errors.Wrap(_sublistListErr, "Error parsing 'sublistList' field")
	}
	sublistList := _sublistList
	if closeErr := readBuffer.CloseContext("sublistList"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for sublistList")
	}

	if closeErr := readBuffer.CloseContext("SzlId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SzlId")
	}

	// Create the instance
	return NewSzlId(typeClass, sublistExtract, sublistList), nil
}

func (m *_SzlId) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("SzlId"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for SzlId")
	}

	// Simple Field (typeClass)
	if pushErr := writeBuffer.PushContext("typeClass"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for typeClass")
	}
	_typeClassErr := writeBuffer.WriteSerializable(m.GetTypeClass())
	if popErr := writeBuffer.PopContext("typeClass"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for typeClass")
	}
	if _typeClassErr != nil {
		return errors.Wrap(_typeClassErr, "Error serializing 'typeClass' field")
	}

	// Simple Field (sublistExtract)
	sublistExtract := uint8(m.GetSublistExtract())
	_sublistExtractErr := writeBuffer.WriteUint8("sublistExtract", 4, (sublistExtract))
	if _sublistExtractErr != nil {
		return errors.Wrap(_sublistExtractErr, "Error serializing 'sublistExtract' field")
	}

	// Simple Field (sublistList)
	if pushErr := writeBuffer.PushContext("sublistList"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for sublistList")
	}
	_sublistListErr := writeBuffer.WriteSerializable(m.GetSublistList())
	if popErr := writeBuffer.PopContext("sublistList"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for sublistList")
	}
	if _sublistListErr != nil {
		return errors.Wrap(_sublistListErr, "Error serializing 'sublistList' field")
	}

	if popErr := writeBuffer.PopContext("SzlId"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for SzlId")
	}
	return nil
}

func (m *_SzlId) isSzlId() bool {
	return true
}

func (m *_SzlId) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
