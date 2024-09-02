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

// BACnetEventParameter is the corresponding interface of BACnetEventParameter
type BACnetEventParameter interface {
	BACnetEventParameterContract
	BACnetEventParameterRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// BACnetEventParameterContract provides a set of functions which can be overwritten by a sub struct
type BACnetEventParameterContract interface {
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
}

// BACnetEventParameterRequirements provides a set of functions which need to be implemented by a sub struct
type BACnetEventParameterRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetPeekedTagNumber returns PeekedTagNumber (discriminator field)
	GetPeekedTagNumber() uint8
}

// BACnetEventParameterExactly can be used when we want exactly this type and not a type which fulfills BACnetEventParameter.
// This is useful for switch cases.
type BACnetEventParameterExactly interface {
	BACnetEventParameter
	isBACnetEventParameter() bool
}

// _BACnetEventParameter is the data-structure of this message
type _BACnetEventParameter struct {
	_SubType        BACnetEventParameter
	PeekedTagHeader BACnetTagHeader
}

var _ BACnetEventParameterContract = (*_BACnetEventParameter)(nil)

type BACnetEventParameterChild interface {
	utils.Serializable

	GetParent() *BACnetEventParameter

	GetTypeName() string
	BACnetEventParameter
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventParameter) GetPeekedTagHeader() BACnetTagHeader {
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

func (pm *_BACnetEventParameter) GetPeekedTagNumber() uint8 {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetEventParameter factory function for _BACnetEventParameter
func NewBACnetEventParameter(peekedTagHeader BACnetTagHeader) *_BACnetEventParameter {
	return &_BACnetEventParameter{PeekedTagHeader: peekedTagHeader}
}

// Deprecated: use the interface for direct cast
func CastBACnetEventParameter(structType any) BACnetEventParameter {
	if casted, ok := structType.(BACnetEventParameter); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventParameter); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventParameter) GetTypeName() string {
	return "BACnetEventParameter"
}

func (m *_BACnetEventParameter) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetEventParameter) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func BACnetEventParameterParse[T BACnetEventParameter](ctx context.Context, theBytes []byte) (T, error) {
	return BACnetEventParameterParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetEventParameterParseWithBufferProducer[T BACnetEventParameter]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := BACnetEventParameterParseWithBuffer[T](ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, err
	}
}

func BACnetEventParameterParseWithBuffer[T BACnetEventParameter](ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	v, err := NewBACnetEventParameter().parse(ctx, readBuffer)
	if err != nil {
		var zero T
		return zero, err
	}
	return v.(T), err
}

func (m *_BACnetEventParameter) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetEventParameter BACnetEventParameter, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventParameter"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventParameter")
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
	var _child BACnetEventParameter
	switch {
	case peekedTagNumber == uint8(0): // BACnetEventParameterChangeOfBitstring
		if _child, err = (&_BACnetEventParameterChangeOfBitstring{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetEventParameterChangeOfBitstring for type-switch of BACnetEventParameter")
		}
	case peekedTagNumber == uint8(1): // BACnetEventParameterChangeOfState
		if _child, err = (&_BACnetEventParameterChangeOfState{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetEventParameterChangeOfState for type-switch of BACnetEventParameter")
		}
	case peekedTagNumber == uint8(2): // BACnetEventParameterChangeOfValue
		if _child, err = (&_BACnetEventParameterChangeOfValue{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetEventParameterChangeOfValue for type-switch of BACnetEventParameter")
		}
	case peekedTagNumber == uint8(3): // BACnetEventParameterCommandFailure
		if _child, err = (&_BACnetEventParameterCommandFailure{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetEventParameterCommandFailure for type-switch of BACnetEventParameter")
		}
	case peekedTagNumber == uint8(4): // BACnetEventParameterFloatingLimit
		if _child, err = (&_BACnetEventParameterFloatingLimit{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetEventParameterFloatingLimit for type-switch of BACnetEventParameter")
		}
	case peekedTagNumber == uint8(5): // BACnetEventParameterOutOfRange
		if _child, err = (&_BACnetEventParameterOutOfRange{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetEventParameterOutOfRange for type-switch of BACnetEventParameter")
		}
	case peekedTagNumber == uint8(8): // BACnetEventParameterChangeOfLifeSavety
		if _child, err = (&_BACnetEventParameterChangeOfLifeSavety{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetEventParameterChangeOfLifeSavety for type-switch of BACnetEventParameter")
		}
	case peekedTagNumber == uint8(9): // BACnetEventParameterExtended
		if _child, err = (&_BACnetEventParameterExtended{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetEventParameterExtended for type-switch of BACnetEventParameter")
		}
	case peekedTagNumber == uint8(10): // BACnetEventParameterBufferReady
		if _child, err = (&_BACnetEventParameterBufferReady{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetEventParameterBufferReady for type-switch of BACnetEventParameter")
		}
	case peekedTagNumber == uint8(11): // BACnetEventParameterUnsignedRange
		if _child, err = (&_BACnetEventParameterUnsignedRange{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetEventParameterUnsignedRange for type-switch of BACnetEventParameter")
		}
	case peekedTagNumber == uint8(13): // BACnetEventParameterAccessEvent
		if _child, err = (&_BACnetEventParameterAccessEvent{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetEventParameterAccessEvent for type-switch of BACnetEventParameter")
		}
	case peekedTagNumber == uint8(14): // BACnetEventParameterDoubleOutOfRange
		if _child, err = (&_BACnetEventParameterDoubleOutOfRange{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetEventParameterDoubleOutOfRange for type-switch of BACnetEventParameter")
		}
	case peekedTagNumber == uint8(15): // BACnetEventParameterSignedOutOfRange
		if _child, err = (&_BACnetEventParameterSignedOutOfRange{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetEventParameterSignedOutOfRange for type-switch of BACnetEventParameter")
		}
	case peekedTagNumber == uint8(16): // BACnetEventParameterUnsignedOutOfRange
		if _child, err = (&_BACnetEventParameterUnsignedOutOfRange{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetEventParameterUnsignedOutOfRange for type-switch of BACnetEventParameter")
		}
	case peekedTagNumber == uint8(17): // BACnetEventParameterChangeOfCharacterString
		if _child, err = (&_BACnetEventParameterChangeOfCharacterString{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetEventParameterChangeOfCharacterString for type-switch of BACnetEventParameter")
		}
	case peekedTagNumber == uint8(18): // BACnetEventParameterChangeOfStatusFlags
		if _child, err = (&_BACnetEventParameterChangeOfStatusFlags{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetEventParameterChangeOfStatusFlags for type-switch of BACnetEventParameter")
		}
	case peekedTagNumber == uint8(20): // BACnetEventParameterNone
		if _child, err = (&_BACnetEventParameterNone{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetEventParameterNone for type-switch of BACnetEventParameter")
		}
	case peekedTagNumber == uint8(21): // BACnetEventParameterChangeOfDiscreteValue
		if _child, err = (&_BACnetEventParameterChangeOfDiscreteValue{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetEventParameterChangeOfDiscreteValue for type-switch of BACnetEventParameter")
		}
	case peekedTagNumber == uint8(22): // BACnetEventParameterChangeOfTimer
		if _child, err = (&_BACnetEventParameterChangeOfTimer{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetEventParameterChangeOfTimer for type-switch of BACnetEventParameter")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [peekedTagNumber=%v]", peekedTagNumber)
	}

	if closeErr := readBuffer.CloseContext("BACnetEventParameter"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventParameter")
	}

	return _child, nil
}

func (pm *_BACnetEventParameter) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetEventParameter, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetEventParameter"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetEventParameter")
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

	if popErr := writeBuffer.PopContext("BACnetEventParameter"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetEventParameter")
	}
	return nil
}

func (m *_BACnetEventParameter) isBACnetEventParameter() bool {
	return true
}
