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

// WritePropertyMultipleError is the corresponding interface of WritePropertyMultipleError
type WritePropertyMultipleError interface {
	utils.LengthAware
	utils.Serializable
	BACnetError
	// GetErrorType returns ErrorType (property field)
	GetErrorType() ErrorEnclosed
	// GetFirstFailedWriteAttempt returns FirstFailedWriteAttempt (property field)
	GetFirstFailedWriteAttempt() BACnetObjectPropertyReferenceEnclosed
}

// WritePropertyMultipleErrorExactly can be used when we want exactly this type and not a type which fulfills WritePropertyMultipleError.
// This is useful for switch cases.
type WritePropertyMultipleErrorExactly interface {
	isWritePropertyMultipleError() bool
}

// _WritePropertyMultipleError is the data-structure of this message
type _WritePropertyMultipleError struct {
	*_BACnetError
	ErrorType               ErrorEnclosed
	FirstFailedWriteAttempt BACnetObjectPropertyReferenceEnclosed
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_WritePropertyMultipleError) GetErrorChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_WRITE_PROPERTY_MULTIPLE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_WritePropertyMultipleError) InitializeParent(parent BACnetError) {}

func (m *_WritePropertyMultipleError) GetParent() BACnetError {
	return m._BACnetError
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_WritePropertyMultipleError) GetErrorType() ErrorEnclosed {
	return m.ErrorType
}

func (m *_WritePropertyMultipleError) GetFirstFailedWriteAttempt() BACnetObjectPropertyReferenceEnclosed {
	return m.FirstFailedWriteAttempt
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewWritePropertyMultipleError factory function for _WritePropertyMultipleError
func NewWritePropertyMultipleError(errorType ErrorEnclosed, firstFailedWriteAttempt BACnetObjectPropertyReferenceEnclosed) *_WritePropertyMultipleError {
	_result := &_WritePropertyMultipleError{
		ErrorType:               errorType,
		FirstFailedWriteAttempt: firstFailedWriteAttempt,
		_BACnetError:            NewBACnetError(),
	}
	_result._BACnetError._BACnetErrorChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastWritePropertyMultipleError(structType interface{}) WritePropertyMultipleError {
	if casted, ok := structType.(WritePropertyMultipleError); ok {
		return casted
	}
	if casted, ok := structType.(*WritePropertyMultipleError); ok {
		return *casted
	}
	return nil
}

func (m *_WritePropertyMultipleError) GetTypeName() string {
	return "WritePropertyMultipleError"
}

func (m *_WritePropertyMultipleError) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_WritePropertyMultipleError) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (errorType)
	lengthInBits += m.ErrorType.GetLengthInBits()

	// Simple field (firstFailedWriteAttempt)
	lengthInBits += m.FirstFailedWriteAttempt.GetLengthInBits()

	return lengthInBits
}

func (m *_WritePropertyMultipleError) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func WritePropertyMultipleErrorParse(readBuffer utils.ReadBuffer, errorChoice BACnetConfirmedServiceChoice) (WritePropertyMultipleError, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("WritePropertyMultipleError"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for WritePropertyMultipleError")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (errorType)
	if pullErr := readBuffer.PullContext("errorType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for errorType")
	}
	_errorType, _errorTypeErr := ErrorEnclosedParse(readBuffer, uint8(uint8(0)))
	if _errorTypeErr != nil {
		return nil, errors.Wrap(_errorTypeErr, "Error parsing 'errorType' field")
	}
	errorType := _errorType.(ErrorEnclosed)
	if closeErr := readBuffer.CloseContext("errorType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for errorType")
	}

	// Simple Field (firstFailedWriteAttempt)
	if pullErr := readBuffer.PullContext("firstFailedWriteAttempt"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for firstFailedWriteAttempt")
	}
	_firstFailedWriteAttempt, _firstFailedWriteAttemptErr := BACnetObjectPropertyReferenceEnclosedParse(readBuffer, uint8(uint8(1)))
	if _firstFailedWriteAttemptErr != nil {
		return nil, errors.Wrap(_firstFailedWriteAttemptErr, "Error parsing 'firstFailedWriteAttempt' field")
	}
	firstFailedWriteAttempt := _firstFailedWriteAttempt.(BACnetObjectPropertyReferenceEnclosed)
	if closeErr := readBuffer.CloseContext("firstFailedWriteAttempt"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for firstFailedWriteAttempt")
	}

	if closeErr := readBuffer.CloseContext("WritePropertyMultipleError"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for WritePropertyMultipleError")
	}

	// Create a partially initialized instance
	_child := &_WritePropertyMultipleError{
		ErrorType:               errorType,
		FirstFailedWriteAttempt: firstFailedWriteAttempt,
		_BACnetError:            &_BACnetError{},
	}
	_child._BACnetError._BACnetErrorChildRequirements = _child
	return _child, nil
}

func (m *_WritePropertyMultipleError) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("WritePropertyMultipleError"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for WritePropertyMultipleError")
		}

		// Simple Field (errorType)
		if pushErr := writeBuffer.PushContext("errorType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for errorType")
		}
		_errorTypeErr := writeBuffer.WriteSerializable(m.GetErrorType())
		if popErr := writeBuffer.PopContext("errorType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for errorType")
		}
		if _errorTypeErr != nil {
			return errors.Wrap(_errorTypeErr, "Error serializing 'errorType' field")
		}

		// Simple Field (firstFailedWriteAttempt)
		if pushErr := writeBuffer.PushContext("firstFailedWriteAttempt"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for firstFailedWriteAttempt")
		}
		_firstFailedWriteAttemptErr := writeBuffer.WriteSerializable(m.GetFirstFailedWriteAttempt())
		if popErr := writeBuffer.PopContext("firstFailedWriteAttempt"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for firstFailedWriteAttempt")
		}
		if _firstFailedWriteAttemptErr != nil {
			return errors.Wrap(_firstFailedWriteAttemptErr, "Error serializing 'firstFailedWriteAttempt' field")
		}

		if popErr := writeBuffer.PopContext("WritePropertyMultipleError"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for WritePropertyMultipleError")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_WritePropertyMultipleError) isWritePropertyMultipleError() bool {
	return true
}

func (m *_WritePropertyMultipleError) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
