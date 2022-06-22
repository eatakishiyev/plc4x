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
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetReadAccessResult is the corresponding interface of BACnetReadAccessResult
type BACnetReadAccessResult interface {
	utils.LengthAware
	utils.Serializable
	// GetObjectIdentifier returns ObjectIdentifier (property field)
	GetObjectIdentifier() BACnetContextTagObjectIdentifier
	// GetListOfResults returns ListOfResults (property field)
	GetListOfResults() BACnetReadAccessResultListOfResults
}

// BACnetReadAccessResultExactly can be used when we want exactly this type and not a type which fulfills BACnetReadAccessResult.
// This is useful for switch cases.
type BACnetReadAccessResultExactly interface {
	isBACnetReadAccessResult() bool
}

// _BACnetReadAccessResult is the data-structure of this message
type _BACnetReadAccessResult struct {
	ObjectIdentifier BACnetContextTagObjectIdentifier
	ListOfResults    BACnetReadAccessResultListOfResults
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetReadAccessResult) GetObjectIdentifier() BACnetContextTagObjectIdentifier {
	return m.ObjectIdentifier
}

func (m *_BACnetReadAccessResult) GetListOfResults() BACnetReadAccessResultListOfResults {
	return m.ListOfResults
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetReadAccessResult factory function for _BACnetReadAccessResult
func NewBACnetReadAccessResult(objectIdentifier BACnetContextTagObjectIdentifier, listOfResults BACnetReadAccessResultListOfResults) *_BACnetReadAccessResult {
	return &_BACnetReadAccessResult{ObjectIdentifier: objectIdentifier, ListOfResults: listOfResults}
}

// Deprecated: use the interface for direct cast
func CastBACnetReadAccessResult(structType interface{}) BACnetReadAccessResult {
	if casted, ok := structType.(BACnetReadAccessResult); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetReadAccessResult); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetReadAccessResult) GetTypeName() string {
	return "BACnetReadAccessResult"
}

func (m *_BACnetReadAccessResult) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetReadAccessResult) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (objectIdentifier)
	lengthInBits += m.ObjectIdentifier.GetLengthInBits()

	// Optional Field (listOfResults)
	if m.ListOfResults != nil {
		lengthInBits += m.ListOfResults.GetLengthInBits()
	}

	return lengthInBits
}

func (m *_BACnetReadAccessResult) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetReadAccessResultParse(readBuffer utils.ReadBuffer) (BACnetReadAccessResult, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetReadAccessResult"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetReadAccessResult")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (objectIdentifier)
	if pullErr := readBuffer.PullContext("objectIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for objectIdentifier")
	}
	_objectIdentifier, _objectIdentifierErr := BACnetContextTagParse(readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_BACNET_OBJECT_IDENTIFIER))
	if _objectIdentifierErr != nil {
		return nil, errors.Wrap(_objectIdentifierErr, "Error parsing 'objectIdentifier' field")
	}
	objectIdentifier := _objectIdentifier.(BACnetContextTagObjectIdentifier)
	if closeErr := readBuffer.CloseContext("objectIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for objectIdentifier")
	}

	// Optional Field (listOfResults) (Can be skipped, if a given expression evaluates to false)
	var listOfResults BACnetReadAccessResultListOfResults = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("listOfResults"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for listOfResults")
		}
		_val, _err := BACnetReadAccessResultListOfResultsParse(readBuffer, uint8(1), objectIdentifier.GetObjectType())
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'listOfResults' field")
		default:
			listOfResults = _val.(BACnetReadAccessResultListOfResults)
			if closeErr := readBuffer.CloseContext("listOfResults"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for listOfResults")
			}
		}
	}

	if closeErr := readBuffer.CloseContext("BACnetReadAccessResult"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetReadAccessResult")
	}

	// Create the instance
	return NewBACnetReadAccessResult(objectIdentifier, listOfResults), nil
}

func (m *_BACnetReadAccessResult) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetReadAccessResult"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetReadAccessResult")
	}

	// Simple Field (objectIdentifier)
	if pushErr := writeBuffer.PushContext("objectIdentifier"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for objectIdentifier")
	}
	_objectIdentifierErr := writeBuffer.WriteSerializable(m.GetObjectIdentifier())
	if popErr := writeBuffer.PopContext("objectIdentifier"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for objectIdentifier")
	}
	if _objectIdentifierErr != nil {
		return errors.Wrap(_objectIdentifierErr, "Error serializing 'objectIdentifier' field")
	}

	// Optional Field (listOfResults) (Can be skipped, if the value is null)
	var listOfResults BACnetReadAccessResultListOfResults = nil
	if m.GetListOfResults() != nil {
		if pushErr := writeBuffer.PushContext("listOfResults"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for listOfResults")
		}
		listOfResults = m.GetListOfResults()
		_listOfResultsErr := writeBuffer.WriteSerializable(listOfResults)
		if popErr := writeBuffer.PopContext("listOfResults"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for listOfResults")
		}
		if _listOfResultsErr != nil {
			return errors.Wrap(_listOfResultsErr, "Error serializing 'listOfResults' field")
		}
	}

	if popErr := writeBuffer.PopContext("BACnetReadAccessResult"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetReadAccessResult")
	}
	return nil
}

func (m *_BACnetReadAccessResult) isBACnetReadAccessResult() bool {
	return true
}

func (m *_BACnetReadAccessResult) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
