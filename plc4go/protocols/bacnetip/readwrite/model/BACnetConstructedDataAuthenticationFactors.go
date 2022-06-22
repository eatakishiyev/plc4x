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

// BACnetConstructedDataAuthenticationFactors is the corresponding interface of BACnetConstructedDataAuthenticationFactors
type BACnetConstructedDataAuthenticationFactors interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetNumberOfDataElements returns NumberOfDataElements (property field)
	GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger
	// GetAuthenticationFactors returns AuthenticationFactors (property field)
	GetAuthenticationFactors() []BACnetCredentialAuthenticationFactor
	// GetZero returns Zero (virtual field)
	GetZero() uint64
}

// BACnetConstructedDataAuthenticationFactorsExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataAuthenticationFactors.
// This is useful for switch cases.
type BACnetConstructedDataAuthenticationFactorsExactly interface {
	isBACnetConstructedDataAuthenticationFactors() bool
}

// _BACnetConstructedDataAuthenticationFactors is the data-structure of this message
type _BACnetConstructedDataAuthenticationFactors struct {
	*_BACnetConstructedData
	NumberOfDataElements  BACnetApplicationTagUnsignedInteger
	AuthenticationFactors []BACnetCredentialAuthenticationFactor

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument BACnetTagPayloadUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAuthenticationFactors) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataAuthenticationFactors) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_AUTHENTICATION_FACTORS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAuthenticationFactors) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataAuthenticationFactors) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAuthenticationFactors) GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger {
	return m.NumberOfDataElements
}

func (m *_BACnetConstructedDataAuthenticationFactors) GetAuthenticationFactors() []BACnetCredentialAuthenticationFactor {
	return m.AuthenticationFactors
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAuthenticationFactors) GetZero() uint64 {
	numberOfDataElements := m.NumberOfDataElements
	_ = numberOfDataElements
	return uint64(uint64(0))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataAuthenticationFactors factory function for _BACnetConstructedDataAuthenticationFactors
func NewBACnetConstructedDataAuthenticationFactors(numberOfDataElements BACnetApplicationTagUnsignedInteger, authenticationFactors []BACnetCredentialAuthenticationFactor, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAuthenticationFactors {
	_result := &_BACnetConstructedDataAuthenticationFactors{
		NumberOfDataElements:   numberOfDataElements,
		AuthenticationFactors:  authenticationFactors,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAuthenticationFactors(structType interface{}) BACnetConstructedDataAuthenticationFactors {
	if casted, ok := structType.(BACnetConstructedDataAuthenticationFactors); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAuthenticationFactors); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAuthenticationFactors) GetTypeName() string {
	return "BACnetConstructedDataAuthenticationFactors"
}

func (m *_BACnetConstructedDataAuthenticationFactors) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataAuthenticationFactors) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// A virtual field doesn't have any in- or output.

	// Optional Field (numberOfDataElements)
	if m.NumberOfDataElements != nil {
		lengthInBits += m.NumberOfDataElements.GetLengthInBits()
	}

	// Array field
	if len(m.AuthenticationFactors) > 0 {
		for _, element := range m.AuthenticationFactors {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataAuthenticationFactors) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataAuthenticationFactorsParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataAuthenticationFactors, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAuthenticationFactors"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAuthenticationFactors")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Virtual field
	_zero := uint64(0)
	zero := uint64(_zero)
	_ = zero

	// Optional Field (numberOfDataElements) (Can be skipped, if a given expression evaluates to false)
	var numberOfDataElements BACnetApplicationTagUnsignedInteger = nil
	if bool(bool((arrayIndexArgument) != (nil))) && bool(bool((arrayIndexArgument.GetActualValue()) == (zero))) {
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("numberOfDataElements"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for numberOfDataElements")
		}
		_val, _err := BACnetApplicationTagParse(readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'numberOfDataElements' field")
		default:
			numberOfDataElements = _val.(BACnetApplicationTagUnsignedInteger)
			if closeErr := readBuffer.CloseContext("numberOfDataElements"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for numberOfDataElements")
			}
		}
	}

	// Array field (authenticationFactors)
	if pullErr := readBuffer.PullContext("authenticationFactors", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for authenticationFactors")
	}
	// Terminated array
	authenticationFactors := make([]BACnetCredentialAuthenticationFactor, 0)
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetCredentialAuthenticationFactorParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'authenticationFactors' field")
			}
			authenticationFactors = append(authenticationFactors, _item.(BACnetCredentialAuthenticationFactor))

		}
	}
	if closeErr := readBuffer.CloseContext("authenticationFactors", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for authenticationFactors")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAuthenticationFactors"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAuthenticationFactors")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataAuthenticationFactors{
		NumberOfDataElements:   numberOfDataElements,
		AuthenticationFactors:  authenticationFactors,
		_BACnetConstructedData: &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataAuthenticationFactors) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAuthenticationFactors"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAuthenticationFactors")
		}
		// Virtual field
		if _zeroErr := writeBuffer.WriteVirtual("zero", m.GetZero()); _zeroErr != nil {
			return errors.Wrap(_zeroErr, "Error serializing 'zero' field")
		}

		// Optional Field (numberOfDataElements) (Can be skipped, if the value is null)
		var numberOfDataElements BACnetApplicationTagUnsignedInteger = nil
		if m.GetNumberOfDataElements() != nil {
			if pushErr := writeBuffer.PushContext("numberOfDataElements"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for numberOfDataElements")
			}
			numberOfDataElements = m.GetNumberOfDataElements()
			_numberOfDataElementsErr := writeBuffer.WriteSerializable(numberOfDataElements)
			if popErr := writeBuffer.PopContext("numberOfDataElements"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for numberOfDataElements")
			}
			if _numberOfDataElementsErr != nil {
				return errors.Wrap(_numberOfDataElementsErr, "Error serializing 'numberOfDataElements' field")
			}
		}

		// Array Field (authenticationFactors)
		if m.GetAuthenticationFactors() != nil {
			if pushErr := writeBuffer.PushContext("authenticationFactors", utils.WithRenderAsList(true)); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for authenticationFactors")
			}
			for _, _element := range m.GetAuthenticationFactors() {
				_elementErr := writeBuffer.WriteSerializable(_element)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'authenticationFactors' field")
				}
			}
			if popErr := writeBuffer.PopContext("authenticationFactors", utils.WithRenderAsList(true)); popErr != nil {
				return errors.Wrap(popErr, "Error popping for authenticationFactors")
			}
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAuthenticationFactors"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAuthenticationFactors")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAuthenticationFactors) isBACnetConstructedDataAuthenticationFactors() bool {
	return true
}

func (m *_BACnetConstructedDataAuthenticationFactors) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
