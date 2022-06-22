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

// BACnetConstructedDataInProcess is the corresponding interface of BACnetConstructedDataInProcess
type BACnetConstructedDataInProcess interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetInProcess returns InProcess (property field)
	GetInProcess() BACnetApplicationTagBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagBoolean
}

// BACnetConstructedDataInProcessExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataInProcess.
// This is useful for switch cases.
type BACnetConstructedDataInProcessExactly interface {
	isBACnetConstructedDataInProcess() bool
}

// _BACnetConstructedDataInProcess is the data-structure of this message
type _BACnetConstructedDataInProcess struct {
	*_BACnetConstructedData
	InProcess BACnetApplicationTagBoolean

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument BACnetTagPayloadUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataInProcess) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataInProcess) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_IN_PROCESS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataInProcess) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataInProcess) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataInProcess) GetInProcess() BACnetApplicationTagBoolean {
	return m.InProcess
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataInProcess) GetActualValue() BACnetApplicationTagBoolean {
	return CastBACnetApplicationTagBoolean(m.GetInProcess())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataInProcess factory function for _BACnetConstructedDataInProcess
func NewBACnetConstructedDataInProcess(inProcess BACnetApplicationTagBoolean, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataInProcess {
	_result := &_BACnetConstructedDataInProcess{
		InProcess:              inProcess,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataInProcess(structType interface{}) BACnetConstructedDataInProcess {
	if casted, ok := structType.(BACnetConstructedDataInProcess); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataInProcess); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataInProcess) GetTypeName() string {
	return "BACnetConstructedDataInProcess"
}

func (m *_BACnetConstructedDataInProcess) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataInProcess) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (inProcess)
	lengthInBits += m.InProcess.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataInProcess) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataInProcessParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataInProcess, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataInProcess"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataInProcess")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (inProcess)
	if pullErr := readBuffer.PullContext("inProcess"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for inProcess")
	}
	_inProcess, _inProcessErr := BACnetApplicationTagParse(readBuffer)
	if _inProcessErr != nil {
		return nil, errors.Wrap(_inProcessErr, "Error parsing 'inProcess' field")
	}
	inProcess := _inProcess.(BACnetApplicationTagBoolean)
	if closeErr := readBuffer.CloseContext("inProcess"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for inProcess")
	}

	// Virtual field
	_actualValue := inProcess
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataInProcess"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataInProcess")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataInProcess{
		InProcess:              inProcess,
		_BACnetConstructedData: &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataInProcess) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataInProcess"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataInProcess")
		}

		// Simple Field (inProcess)
		if pushErr := writeBuffer.PushContext("inProcess"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for inProcess")
		}
		_inProcessErr := writeBuffer.WriteSerializable(m.GetInProcess())
		if popErr := writeBuffer.PopContext("inProcess"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for inProcess")
		}
		if _inProcessErr != nil {
			return errors.Wrap(_inProcessErr, "Error serializing 'inProcess' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataInProcess"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataInProcess")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataInProcess) isBACnetConstructedDataInProcess() bool {
	return true
}

func (m *_BACnetConstructedDataInProcess) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
