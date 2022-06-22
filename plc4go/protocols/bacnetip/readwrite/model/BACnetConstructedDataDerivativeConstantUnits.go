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

// BACnetConstructedDataDerivativeConstantUnits is the corresponding interface of BACnetConstructedDataDerivativeConstantUnits
type BACnetConstructedDataDerivativeConstantUnits interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetUnits returns Units (property field)
	GetUnits() BACnetEngineeringUnitsTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetEngineeringUnitsTagged
}

// BACnetConstructedDataDerivativeConstantUnitsExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataDerivativeConstantUnits.
// This is useful for switch cases.
type BACnetConstructedDataDerivativeConstantUnitsExactly interface {
	BACnetConstructedDataDerivativeConstantUnits
	isBACnetConstructedDataDerivativeConstantUnits() bool
}

// _BACnetConstructedDataDerivativeConstantUnits is the data-structure of this message
type _BACnetConstructedDataDerivativeConstantUnits struct {
	*_BACnetConstructedData
	Units BACnetEngineeringUnitsTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataDerivativeConstantUnits) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataDerivativeConstantUnits) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_DERIVATIVE_CONSTANT_UNITS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataDerivativeConstantUnits) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataDerivativeConstantUnits) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataDerivativeConstantUnits) GetUnits() BACnetEngineeringUnitsTagged {
	return m.Units
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataDerivativeConstantUnits) GetActualValue() BACnetEngineeringUnitsTagged {
	return CastBACnetEngineeringUnitsTagged(m.GetUnits())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataDerivativeConstantUnits factory function for _BACnetConstructedDataDerivativeConstantUnits
func NewBACnetConstructedDataDerivativeConstantUnits(units BACnetEngineeringUnitsTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataDerivativeConstantUnits {
	_result := &_BACnetConstructedDataDerivativeConstantUnits{
		Units:                  units,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataDerivativeConstantUnits(structType interface{}) BACnetConstructedDataDerivativeConstantUnits {
	if casted, ok := structType.(BACnetConstructedDataDerivativeConstantUnits); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDerivativeConstantUnits); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataDerivativeConstantUnits) GetTypeName() string {
	return "BACnetConstructedDataDerivativeConstantUnits"
}

func (m *_BACnetConstructedDataDerivativeConstantUnits) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataDerivativeConstantUnits) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (units)
	lengthInBits += m.Units.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataDerivativeConstantUnits) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataDerivativeConstantUnitsParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataDerivativeConstantUnits, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDerivativeConstantUnits"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataDerivativeConstantUnits")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (units)
	if pullErr := readBuffer.PullContext("units"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for units")
	}
	_units, _unitsErr := BACnetEngineeringUnitsTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _unitsErr != nil {
		return nil, errors.Wrap(_unitsErr, "Error parsing 'units' field")
	}
	units := _units.(BACnetEngineeringUnitsTagged)
	if closeErr := readBuffer.CloseContext("units"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for units")
	}

	// Virtual field
	_actualValue := units
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDerivativeConstantUnits"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataDerivativeConstantUnits")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataDerivativeConstantUnits{
		Units: units,
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataDerivativeConstantUnits) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDerivativeConstantUnits"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataDerivativeConstantUnits")
		}

		// Simple Field (units)
		if pushErr := writeBuffer.PushContext("units"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for units")
		}
		_unitsErr := writeBuffer.WriteSerializable(m.GetUnits())
		if popErr := writeBuffer.PopContext("units"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for units")
		}
		if _unitsErr != nil {
			return errors.Wrap(_unitsErr, "Error serializing 'units' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDerivativeConstantUnits"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataDerivativeConstantUnits")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataDerivativeConstantUnits) isBACnetConstructedDataDerivativeConstantUnits() bool {
	return true
}

func (m *_BACnetConstructedDataDerivativeConstantUnits) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
