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

// Constant values.
const CBusConstants_CBUSTCPDEFAULTPORT uint16 = uint16(10001)

// CBusConstants is the corresponding interface of CBusConstants
type CBusConstants interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// CBusConstantsExactly can be used when we want exactly this type and not a type which fulfills CBusConstants.
// This is useful for switch cases.
type CBusConstantsExactly interface {
	CBusConstants
	isCBusConstants() bool
}

// _CBusConstants is the data-structure of this message
type _CBusConstants struct {
}

var _ CBusConstants = (*_CBusConstants)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_CBusConstants) GetCbusTcpDefaultPort() uint16 {
	return CBusConstants_CBUSTCPDEFAULTPORT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCBusConstants factory function for _CBusConstants
func NewCBusConstants() *_CBusConstants {
	return &_CBusConstants{}
}

// Deprecated: use the interface for direct cast
func CastCBusConstants(structType any) CBusConstants {
	if casted, ok := structType.(CBusConstants); ok {
		return casted
	}
	if casted, ok := structType.(*CBusConstants); ok {
		return *casted
	}
	return nil
}

func (m *_CBusConstants) GetTypeName() string {
	return "CBusConstants"
}

func (m *_CBusConstants) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Const Field (cbusTcpDefaultPort)
	lengthInBits += 16

	return lengthInBits
}

func (m *_CBusConstants) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func CBusConstantsParse(ctx context.Context, theBytes []byte) (CBusConstants, error) {
	return CBusConstantsParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func CBusConstantsParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (CBusConstants, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (CBusConstants, error) {
		return CBusConstantsParseWithBuffer(ctx, readBuffer)
	}
}

func CBusConstantsParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (CBusConstants, error) {
	v, err := NewCBusConstants().parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_CBusConstants) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__cBusConstants CBusConstants, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CBusConstants"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CBusConstants")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	cbusTcpDefaultPort, err := ReadConstField[uint16](ctx, "cbusTcpDefaultPort", ReadUnsignedShort(readBuffer, uint8(16)), CBusConstants_CBUSTCPDEFAULTPORT)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'cbusTcpDefaultPort' field"))
	}
	_ = cbusTcpDefaultPort

	if closeErr := readBuffer.CloseContext("CBusConstants"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CBusConstants")
	}

	return m, nil
}

func (m *_CBusConstants) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CBusConstants) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("CBusConstants"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for CBusConstants")
	}

	if err := WriteConstField(ctx, "cbusTcpDefaultPort", CBusConstants_CBUSTCPDEFAULTPORT, WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'cbusTcpDefaultPort' field")
	}

	if popErr := writeBuffer.PopContext("CBusConstants"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for CBusConstants")
	}
	return nil
}

func (m *_CBusConstants) isCBusConstants() bool {
	return true
}

func (m *_CBusConstants) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
