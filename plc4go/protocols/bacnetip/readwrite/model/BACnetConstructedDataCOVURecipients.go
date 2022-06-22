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

// BACnetConstructedDataCOVURecipients is the corresponding interface of BACnetConstructedDataCOVURecipients
type BACnetConstructedDataCOVURecipients interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetCovuRecipients returns CovuRecipients (property field)
	GetCovuRecipients() []BACnetRecipient
}

// BACnetConstructedDataCOVURecipientsExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataCOVURecipients.
// This is useful for switch cases.
type BACnetConstructedDataCOVURecipientsExactly interface {
	isBACnetConstructedDataCOVURecipients() bool
}

// _BACnetConstructedDataCOVURecipients is the data-structure of this message
type _BACnetConstructedDataCOVURecipients struct {
	*_BACnetConstructedData
	CovuRecipients []BACnetRecipient

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument BACnetTagPayloadUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataCOVURecipients) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataCOVURecipients) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_COVU_RECIPIENTS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataCOVURecipients) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataCOVURecipients) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataCOVURecipients) GetCovuRecipients() []BACnetRecipient {
	return m.CovuRecipients
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataCOVURecipients factory function for _BACnetConstructedDataCOVURecipients
func NewBACnetConstructedDataCOVURecipients(covuRecipients []BACnetRecipient, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataCOVURecipients {
	_result := &_BACnetConstructedDataCOVURecipients{
		CovuRecipients:         covuRecipients,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataCOVURecipients(structType interface{}) BACnetConstructedDataCOVURecipients {
	if casted, ok := structType.(BACnetConstructedDataCOVURecipients); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataCOVURecipients); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataCOVURecipients) GetTypeName() string {
	return "BACnetConstructedDataCOVURecipients"
}

func (m *_BACnetConstructedDataCOVURecipients) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataCOVURecipients) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.CovuRecipients) > 0 {
		for _, element := range m.CovuRecipients {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataCOVURecipients) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataCOVURecipientsParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataCOVURecipients, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataCOVURecipients"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataCOVURecipients")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (covuRecipients)
	if pullErr := readBuffer.PullContext("covuRecipients", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for covuRecipients")
	}
	// Terminated array
	covuRecipients := make([]BACnetRecipient, 0)
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetRecipientParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'covuRecipients' field")
			}
			covuRecipients = append(covuRecipients, _item.(BACnetRecipient))

		}
	}
	if closeErr := readBuffer.CloseContext("covuRecipients", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for covuRecipients")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataCOVURecipients"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataCOVURecipients")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataCOVURecipients{
		CovuRecipients:         covuRecipients,
		_BACnetConstructedData: &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataCOVURecipients) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataCOVURecipients"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataCOVURecipients")
		}

		// Array Field (covuRecipients)
		if m.GetCovuRecipients() != nil {
			if pushErr := writeBuffer.PushContext("covuRecipients", utils.WithRenderAsList(true)); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for covuRecipients")
			}
			for _, _element := range m.GetCovuRecipients() {
				_elementErr := writeBuffer.WriteSerializable(_element)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'covuRecipients' field")
				}
			}
			if popErr := writeBuffer.PopContext("covuRecipients", utils.WithRenderAsList(true)); popErr != nil {
				return errors.Wrap(popErr, "Error popping for covuRecipients")
			}
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataCOVURecipients"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataCOVURecipients")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataCOVURecipients) isBACnetConstructedDataCOVURecipients() bool {
	return true
}

func (m *_BACnetConstructedDataCOVURecipients) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
