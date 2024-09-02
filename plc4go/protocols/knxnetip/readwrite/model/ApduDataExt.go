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

// ApduDataExt is the corresponding interface of ApduDataExt
type ApduDataExt interface {
	ApduDataExtContract
	ApduDataExtRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// ApduDataExtContract provides a set of functions which can be overwritten by a sub struct
type ApduDataExtContract interface {
	// GetLength() returns a parser argument
	GetLength() uint8
}

// ApduDataExtRequirements provides a set of functions which need to be implemented by a sub struct
type ApduDataExtRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetExtApciType returns ExtApciType (discriminator field)
	GetExtApciType() uint8
}

// ApduDataExtExactly can be used when we want exactly this type and not a type which fulfills ApduDataExt.
// This is useful for switch cases.
type ApduDataExtExactly interface {
	ApduDataExt
	isApduDataExt() bool
}

// _ApduDataExt is the data-structure of this message
type _ApduDataExt struct {
	_SubType ApduDataExt

	// Arguments.
	Length uint8
}

var _ ApduDataExtContract = (*_ApduDataExt)(nil)

type ApduDataExtChild interface {
	utils.Serializable

	GetParent() *ApduDataExt

	GetTypeName() string
	ApduDataExt
}

// NewApduDataExt factory function for _ApduDataExt
func NewApduDataExt(length uint8) *_ApduDataExt {
	return &_ApduDataExt{Length: length}
}

// Deprecated: use the interface for direct cast
func CastApduDataExt(structType any) ApduDataExt {
	if casted, ok := structType.(ApduDataExt); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataExt); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataExt) GetTypeName() string {
	return "ApduDataExt"
}

func (m *_ApduDataExt) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (extApciType)
	lengthInBits += 6

	return lengthInBits
}

func (m *_ApduDataExt) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func ApduDataExtParse[T ApduDataExt](ctx context.Context, theBytes []byte, length uint8) (T, error) {
	return ApduDataExtParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes), length)
}

func ApduDataExtParseWithBufferProducer[T ApduDataExt](length uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := ApduDataExtParseWithBuffer[T](ctx, readBuffer, length)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, err
	}
}

func ApduDataExtParseWithBuffer[T ApduDataExt](ctx context.Context, readBuffer utils.ReadBuffer, length uint8) (T, error) {
	v, err := NewApduDataExt(length).parse(ctx, readBuffer, length)
	if err != nil {
		var zero T
		return zero, err
	}
	return v.(T), err
}

func (m *_ApduDataExt) parse(ctx context.Context, readBuffer utils.ReadBuffer, length uint8) (__apduDataExt ApduDataExt, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExt"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExt")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	extApciType, err := ReadDiscriminatorField[uint8](ctx, "extApciType", ReadUnsignedByte(readBuffer, uint8(6)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'extApciType' field"))
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child ApduDataExt
	switch {
	case extApciType == 0x00: // ApduDataExtOpenRoutingTableRequest
		if _child, err = (&_ApduDataExtOpenRoutingTableRequest{}).parse(ctx, readBuffer, m, length); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataExtOpenRoutingTableRequest for type-switch of ApduDataExt")
		}
	case extApciType == 0x01: // ApduDataExtReadRoutingTableRequest
		if _child, err = (&_ApduDataExtReadRoutingTableRequest{}).parse(ctx, readBuffer, m, length); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataExtReadRoutingTableRequest for type-switch of ApduDataExt")
		}
	case extApciType == 0x02: // ApduDataExtReadRoutingTableResponse
		if _child, err = (&_ApduDataExtReadRoutingTableResponse{}).parse(ctx, readBuffer, m, length); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataExtReadRoutingTableResponse for type-switch of ApduDataExt")
		}
	case extApciType == 0x03: // ApduDataExtWriteRoutingTableRequest
		if _child, err = (&_ApduDataExtWriteRoutingTableRequest{}).parse(ctx, readBuffer, m, length); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataExtWriteRoutingTableRequest for type-switch of ApduDataExt")
		}
	case extApciType == 0x08: // ApduDataExtReadRouterMemoryRequest
		if _child, err = (&_ApduDataExtReadRouterMemoryRequest{}).parse(ctx, readBuffer, m, length); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataExtReadRouterMemoryRequest for type-switch of ApduDataExt")
		}
	case extApciType == 0x09: // ApduDataExtReadRouterMemoryResponse
		if _child, err = (&_ApduDataExtReadRouterMemoryResponse{}).parse(ctx, readBuffer, m, length); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataExtReadRouterMemoryResponse for type-switch of ApduDataExt")
		}
	case extApciType == 0x0A: // ApduDataExtWriteRouterMemoryRequest
		if _child, err = (&_ApduDataExtWriteRouterMemoryRequest{}).parse(ctx, readBuffer, m, length); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataExtWriteRouterMemoryRequest for type-switch of ApduDataExt")
		}
	case extApciType == 0x0D: // ApduDataExtReadRouterStatusRequest
		if _child, err = (&_ApduDataExtReadRouterStatusRequest{}).parse(ctx, readBuffer, m, length); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataExtReadRouterStatusRequest for type-switch of ApduDataExt")
		}
	case extApciType == 0x0E: // ApduDataExtReadRouterStatusResponse
		if _child, err = (&_ApduDataExtReadRouterStatusResponse{}).parse(ctx, readBuffer, m, length); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataExtReadRouterStatusResponse for type-switch of ApduDataExt")
		}
	case extApciType == 0x0F: // ApduDataExtWriteRouterStatusRequest
		if _child, err = (&_ApduDataExtWriteRouterStatusRequest{}).parse(ctx, readBuffer, m, length); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataExtWriteRouterStatusRequest for type-switch of ApduDataExt")
		}
	case extApciType == 0x10: // ApduDataExtMemoryBitWrite
		if _child, err = (&_ApduDataExtMemoryBitWrite{}).parse(ctx, readBuffer, m, length); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataExtMemoryBitWrite for type-switch of ApduDataExt")
		}
	case extApciType == 0x11: // ApduDataExtAuthorizeRequest
		if _child, err = (&_ApduDataExtAuthorizeRequest{}).parse(ctx, readBuffer, m, length); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataExtAuthorizeRequest for type-switch of ApduDataExt")
		}
	case extApciType == 0x12: // ApduDataExtAuthorizeResponse
		if _child, err = (&_ApduDataExtAuthorizeResponse{}).parse(ctx, readBuffer, m, length); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataExtAuthorizeResponse for type-switch of ApduDataExt")
		}
	case extApciType == 0x13: // ApduDataExtKeyWrite
		if _child, err = (&_ApduDataExtKeyWrite{}).parse(ctx, readBuffer, m, length); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataExtKeyWrite for type-switch of ApduDataExt")
		}
	case extApciType == 0x14: // ApduDataExtKeyResponse
		if _child, err = (&_ApduDataExtKeyResponse{}).parse(ctx, readBuffer, m, length); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataExtKeyResponse for type-switch of ApduDataExt")
		}
	case extApciType == 0x15: // ApduDataExtPropertyValueRead
		if _child, err = (&_ApduDataExtPropertyValueRead{}).parse(ctx, readBuffer, m, length); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataExtPropertyValueRead for type-switch of ApduDataExt")
		}
	case extApciType == 0x16: // ApduDataExtPropertyValueResponse
		if _child, err = (&_ApduDataExtPropertyValueResponse{}).parse(ctx, readBuffer, m, length); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataExtPropertyValueResponse for type-switch of ApduDataExt")
		}
	case extApciType == 0x17: // ApduDataExtPropertyValueWrite
		if _child, err = (&_ApduDataExtPropertyValueWrite{}).parse(ctx, readBuffer, m, length); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataExtPropertyValueWrite for type-switch of ApduDataExt")
		}
	case extApciType == 0x18: // ApduDataExtPropertyDescriptionRead
		if _child, err = (&_ApduDataExtPropertyDescriptionRead{}).parse(ctx, readBuffer, m, length); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataExtPropertyDescriptionRead for type-switch of ApduDataExt")
		}
	case extApciType == 0x19: // ApduDataExtPropertyDescriptionResponse
		if _child, err = (&_ApduDataExtPropertyDescriptionResponse{}).parse(ctx, readBuffer, m, length); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataExtPropertyDescriptionResponse for type-switch of ApduDataExt")
		}
	case extApciType == 0x1A: // ApduDataExtNetworkParameterRead
		if _child, err = (&_ApduDataExtNetworkParameterRead{}).parse(ctx, readBuffer, m, length); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataExtNetworkParameterRead for type-switch of ApduDataExt")
		}
	case extApciType == 0x1B: // ApduDataExtNetworkParameterResponse
		if _child, err = (&_ApduDataExtNetworkParameterResponse{}).parse(ctx, readBuffer, m, length); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataExtNetworkParameterResponse for type-switch of ApduDataExt")
		}
	case extApciType == 0x1C: // ApduDataExtIndividualAddressSerialNumberRead
		if _child, err = (&_ApduDataExtIndividualAddressSerialNumberRead{}).parse(ctx, readBuffer, m, length); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataExtIndividualAddressSerialNumberRead for type-switch of ApduDataExt")
		}
	case extApciType == 0x1D: // ApduDataExtIndividualAddressSerialNumberResponse
		if _child, err = (&_ApduDataExtIndividualAddressSerialNumberResponse{}).parse(ctx, readBuffer, m, length); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataExtIndividualAddressSerialNumberResponse for type-switch of ApduDataExt")
		}
	case extApciType == 0x1E: // ApduDataExtIndividualAddressSerialNumberWrite
		if _child, err = (&_ApduDataExtIndividualAddressSerialNumberWrite{}).parse(ctx, readBuffer, m, length); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataExtIndividualAddressSerialNumberWrite for type-switch of ApduDataExt")
		}
	case extApciType == 0x20: // ApduDataExtDomainAddressWrite
		if _child, err = (&_ApduDataExtDomainAddressWrite{}).parse(ctx, readBuffer, m, length); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataExtDomainAddressWrite for type-switch of ApduDataExt")
		}
	case extApciType == 0x21: // ApduDataExtDomainAddressRead
		if _child, err = (&_ApduDataExtDomainAddressRead{}).parse(ctx, readBuffer, m, length); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataExtDomainAddressRead for type-switch of ApduDataExt")
		}
	case extApciType == 0x22: // ApduDataExtDomainAddressResponse
		if _child, err = (&_ApduDataExtDomainAddressResponse{}).parse(ctx, readBuffer, m, length); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataExtDomainAddressResponse for type-switch of ApduDataExt")
		}
	case extApciType == 0x23: // ApduDataExtDomainAddressSelectiveRead
		if _child, err = (&_ApduDataExtDomainAddressSelectiveRead{}).parse(ctx, readBuffer, m, length); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataExtDomainAddressSelectiveRead for type-switch of ApduDataExt")
		}
	case extApciType == 0x24: // ApduDataExtNetworkParameterWrite
		if _child, err = (&_ApduDataExtNetworkParameterWrite{}).parse(ctx, readBuffer, m, length); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataExtNetworkParameterWrite for type-switch of ApduDataExt")
		}
	case extApciType == 0x25: // ApduDataExtLinkRead
		if _child, err = (&_ApduDataExtLinkRead{}).parse(ctx, readBuffer, m, length); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataExtLinkRead for type-switch of ApduDataExt")
		}
	case extApciType == 0x26: // ApduDataExtLinkResponse
		if _child, err = (&_ApduDataExtLinkResponse{}).parse(ctx, readBuffer, m, length); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataExtLinkResponse for type-switch of ApduDataExt")
		}
	case extApciType == 0x27: // ApduDataExtLinkWrite
		if _child, err = (&_ApduDataExtLinkWrite{}).parse(ctx, readBuffer, m, length); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataExtLinkWrite for type-switch of ApduDataExt")
		}
	case extApciType == 0x28: // ApduDataExtGroupPropertyValueRead
		if _child, err = (&_ApduDataExtGroupPropertyValueRead{}).parse(ctx, readBuffer, m, length); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataExtGroupPropertyValueRead for type-switch of ApduDataExt")
		}
	case extApciType == 0x29: // ApduDataExtGroupPropertyValueResponse
		if _child, err = (&_ApduDataExtGroupPropertyValueResponse{}).parse(ctx, readBuffer, m, length); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataExtGroupPropertyValueResponse for type-switch of ApduDataExt")
		}
	case extApciType == 0x2A: // ApduDataExtGroupPropertyValueWrite
		if _child, err = (&_ApduDataExtGroupPropertyValueWrite{}).parse(ctx, readBuffer, m, length); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataExtGroupPropertyValueWrite for type-switch of ApduDataExt")
		}
	case extApciType == 0x2B: // ApduDataExtGroupPropertyValueInfoReport
		if _child, err = (&_ApduDataExtGroupPropertyValueInfoReport{}).parse(ctx, readBuffer, m, length); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataExtGroupPropertyValueInfoReport for type-switch of ApduDataExt")
		}
	case extApciType == 0x2C: // ApduDataExtDomainAddressSerialNumberRead
		if _child, err = (&_ApduDataExtDomainAddressSerialNumberRead{}).parse(ctx, readBuffer, m, length); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataExtDomainAddressSerialNumberRead for type-switch of ApduDataExt")
		}
	case extApciType == 0x2D: // ApduDataExtDomainAddressSerialNumberResponse
		if _child, err = (&_ApduDataExtDomainAddressSerialNumberResponse{}).parse(ctx, readBuffer, m, length); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataExtDomainAddressSerialNumberResponse for type-switch of ApduDataExt")
		}
	case extApciType == 0x2E: // ApduDataExtDomainAddressSerialNumberWrite
		if _child, err = (&_ApduDataExtDomainAddressSerialNumberWrite{}).parse(ctx, readBuffer, m, length); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataExtDomainAddressSerialNumberWrite for type-switch of ApduDataExt")
		}
	case extApciType == 0x30: // ApduDataExtFileStreamInfoReport
		if _child, err = (&_ApduDataExtFileStreamInfoReport{}).parse(ctx, readBuffer, m, length); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataExtFileStreamInfoReport for type-switch of ApduDataExt")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [extApciType=%v]", extApciType)
	}

	if closeErr := readBuffer.CloseContext("ApduDataExt"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExt")
	}

	return _child, nil
}

func (pm *_ApduDataExt) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child ApduDataExt, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("ApduDataExt"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ApduDataExt")
	}

	if err := WriteDiscriminatorField(ctx, "extApciType", m.GetExtApciType(), WriteUnsignedByte(writeBuffer, 6)); err != nil {
		return errors.Wrap(err, "Error serializing 'extApciType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("ApduDataExt"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ApduDataExt")
	}
	return nil
}

////
// Arguments Getter

func (m *_ApduDataExt) GetLength() uint8 {
	return m.Length
}

//
////

func (m *_ApduDataExt) isApduDataExt() bool {
	return true
}
