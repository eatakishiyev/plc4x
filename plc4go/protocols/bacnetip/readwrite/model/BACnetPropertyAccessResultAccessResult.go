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
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetPropertyAccessResultAccessResult is the corresponding interface of BACnetPropertyAccessResultAccessResult
type BACnetPropertyAccessResultAccessResult interface {
	BACnetPropertyAccessResultAccessResultContract
	BACnetPropertyAccessResultAccessResultRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// BACnetPropertyAccessResultAccessResultContract provides a set of functions which can be overwritten by a sub struct
type BACnetPropertyAccessResultAccessResultContract interface {
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
	// GetObjectTypeArgument() returns a parser argument
	GetObjectTypeArgument() BACnetObjectType
	// GetPropertyIdentifierArgument() returns a parser argument
	GetPropertyIdentifierArgument() BACnetPropertyIdentifier
	// GetPropertyArrayIndexArgument() returns a parser argument
	GetPropertyArrayIndexArgument() BACnetTagPayloadUnsignedInteger
}

// BACnetPropertyAccessResultAccessResultRequirements provides a set of functions which need to be implemented by a sub struct
type BACnetPropertyAccessResultAccessResultRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetPeekedTagNumber returns PeekedTagNumber (discriminator field)
	GetPeekedTagNumber() uint8
}

// BACnetPropertyAccessResultAccessResultExactly can be used when we want exactly this type and not a type which fulfills BACnetPropertyAccessResultAccessResult.
// This is useful for switch cases.
type BACnetPropertyAccessResultAccessResultExactly interface {
	BACnetPropertyAccessResultAccessResult
	isBACnetPropertyAccessResultAccessResult() bool
}

// _BACnetPropertyAccessResultAccessResult is the data-structure of this message
type _BACnetPropertyAccessResultAccessResult struct {
	_SubType        BACnetPropertyAccessResultAccessResult
	PeekedTagHeader BACnetTagHeader

	// Arguments.
	ObjectTypeArgument         BACnetObjectType
	PropertyIdentifierArgument BACnetPropertyIdentifier
	PropertyArrayIndexArgument BACnetTagPayloadUnsignedInteger
}

var _ BACnetPropertyAccessResultAccessResultContract = (*_BACnetPropertyAccessResultAccessResult)(nil)

type BACnetPropertyAccessResultAccessResultChild interface {
	utils.Serializable

	GetParent() *BACnetPropertyAccessResultAccessResult

	GetTypeName() string
	BACnetPropertyAccessResultAccessResult
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyAccessResultAccessResult) GetPeekedTagHeader() BACnetTagHeader {
	return m.PeekedTagHeader
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (pm *_BACnetPropertyAccessResultAccessResult) GetPeekedTagNumber() uint8 {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyAccessResultAccessResult factory function for _BACnetPropertyAccessResultAccessResult
func NewBACnetPropertyAccessResultAccessResult(peekedTagHeader BACnetTagHeader, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, propertyArrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetPropertyAccessResultAccessResult {
	return &_BACnetPropertyAccessResultAccessResult{PeekedTagHeader: peekedTagHeader, ObjectTypeArgument: objectTypeArgument, PropertyIdentifierArgument: propertyIdentifierArgument, PropertyArrayIndexArgument: propertyArrayIndexArgument}
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyAccessResultAccessResult(structType any) BACnetPropertyAccessResultAccessResult {
	if casted, ok := structType.(BACnetPropertyAccessResultAccessResult); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyAccessResultAccessResult); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyAccessResultAccessResult) GetTypeName() string {
	return "BACnetPropertyAccessResultAccessResult"
}

func (m *_BACnetPropertyAccessResultAccessResult) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetPropertyAccessResultAccessResult) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func BACnetPropertyAccessResultAccessResultParse[T BACnetPropertyAccessResultAccessResult](ctx context.Context, theBytes []byte, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, propertyArrayIndexArgument BACnetTagPayloadUnsignedInteger) (T, error) {
	return BACnetPropertyAccessResultAccessResultParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes), objectTypeArgument, propertyIdentifierArgument, propertyArrayIndexArgument)
}

func BACnetPropertyAccessResultAccessResultParseWithBufferProducer[T BACnetPropertyAccessResultAccessResult](objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, propertyArrayIndexArgument BACnetTagPayloadUnsignedInteger) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := BACnetPropertyAccessResultAccessResultParseWithBuffer[T](ctx, readBuffer, objectTypeArgument, propertyIdentifierArgument, propertyArrayIndexArgument)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, err
	}
}

func BACnetPropertyAccessResultAccessResultParseWithBuffer[T BACnetPropertyAccessResultAccessResult](ctx context.Context, readBuffer utils.ReadBuffer, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, propertyArrayIndexArgument BACnetTagPayloadUnsignedInteger) (T, error) {
	v, err := NewBACnetPropertyAccessResultAccessResult(objectTypeArgument, propertyIdentifierArgument, propertyArrayIndexArgument).parse(ctx, readBuffer, objectTypeArgument, propertyIdentifierArgument, propertyArrayIndexArgument)
	if err != nil {
		var zero T
		return zero, err
	}
	return v.(T), err
}

func (m *_BACnetPropertyAccessResultAccessResult) parse(ctx context.Context, readBuffer utils.ReadBuffer, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, propertyArrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetPropertyAccessResultAccessResult BACnetPropertyAccessResultAccessResult, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyAccessResultAccessResult"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyAccessResultAccessResult")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	peekedTagHeader, err := ReadPeekField[BACnetTagHeader](ctx, "peekedTagHeader", ReadComplex[BACnetTagHeader](BACnetTagHeaderParseWithBuffer, readBuffer), 0)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedTagHeader' field"))
	}
	m.PeekedTagHeader = peekedTagHeader

	peekedTagNumber, err := ReadVirtualField[uint8](ctx, "peekedTagNumber", (*uint8)(nil), peekedTagHeader.GetActualTagNumber())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedTagNumber' field"))
	}
	_ = peekedTagNumber

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child BACnetPropertyAccessResultAccessResult
	switch {
	case peekedTagNumber == uint8(4): // BACnetPropertyAccessResultAccessResultPropertyValue
		if _child, err = (&_BACnetPropertyAccessResultAccessResultPropertyValue{}).parse(ctx, readBuffer, m, objectTypeArgument, propertyIdentifierArgument, propertyArrayIndexArgument); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyAccessResultAccessResultPropertyValue for type-switch of BACnetPropertyAccessResultAccessResult")
		}
	case peekedTagNumber == uint8(5): // BACnetPropertyAccessResultAccessResultPropertyAccessError
		if _child, err = (&_BACnetPropertyAccessResultAccessResultPropertyAccessError{}).parse(ctx, readBuffer, m, objectTypeArgument, propertyIdentifierArgument, propertyArrayIndexArgument); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyAccessResultAccessResultPropertyAccessError for type-switch of BACnetPropertyAccessResultAccessResult")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [peekedTagNumber=%v]", peekedTagNumber)
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyAccessResultAccessResult"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyAccessResultAccessResult")
	}

	return _child, nil
}

func (pm *_BACnetPropertyAccessResultAccessResult) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetPropertyAccessResultAccessResult, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetPropertyAccessResultAccessResult"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetPropertyAccessResultAccessResult")
	}
	// Virtual field
	peekedTagNumber := m.GetPeekedTagNumber()
	_ = peekedTagNumber
	if _peekedTagNumberErr := writeBuffer.WriteVirtual(ctx, "peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("BACnetPropertyAccessResultAccessResult"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetPropertyAccessResultAccessResult")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetPropertyAccessResultAccessResult) GetObjectTypeArgument() BACnetObjectType {
	return m.ObjectTypeArgument
}
func (m *_BACnetPropertyAccessResultAccessResult) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return m.PropertyIdentifierArgument
}
func (m *_BACnetPropertyAccessResultAccessResult) GetPropertyArrayIndexArgument() BACnetTagPayloadUnsignedInteger {
	return m.PropertyArrayIndexArgument
}

//
////

func (m *_BACnetPropertyAccessResultAccessResult) isBACnetPropertyAccessResultAccessResult() bool {
	return true
}
