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

// BACnetSetpointReference is the corresponding interface of BACnetSetpointReference
type BACnetSetpointReference interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetSetPointReference returns SetPointReference (property field)
	GetSetPointReference() BACnetObjectPropertyReferenceEnclosed
}

// BACnetSetpointReferenceExactly can be used when we want exactly this type and not a type which fulfills BACnetSetpointReference.
// This is useful for switch cases.
type BACnetSetpointReferenceExactly interface {
	BACnetSetpointReference
	isBACnetSetpointReference() bool
}

// _BACnetSetpointReference is the data-structure of this message
type _BACnetSetpointReference struct {
	SetPointReference BACnetObjectPropertyReferenceEnclosed
}

var _ BACnetSetpointReference = (*_BACnetSetpointReference)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetSetpointReference) GetSetPointReference() BACnetObjectPropertyReferenceEnclosed {
	return m.SetPointReference
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetSetpointReference factory function for _BACnetSetpointReference
func NewBACnetSetpointReference(setPointReference BACnetObjectPropertyReferenceEnclosed) *_BACnetSetpointReference {
	return &_BACnetSetpointReference{SetPointReference: setPointReference}
}

// Deprecated: use the interface for direct cast
func CastBACnetSetpointReference(structType any) BACnetSetpointReference {
	if casted, ok := structType.(BACnetSetpointReference); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetSetpointReference); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetSetpointReference) GetTypeName() string {
	return "BACnetSetpointReference"
}

func (m *_BACnetSetpointReference) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Optional Field (setPointReference)
	if m.SetPointReference != nil {
		lengthInBits += m.SetPointReference.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_BACnetSetpointReference) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetSetpointReferenceParse(ctx context.Context, theBytes []byte) (BACnetSetpointReference, error) {
	return BACnetSetpointReferenceParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetSetpointReferenceParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetSetpointReference, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetSetpointReference, error) {
		return BACnetSetpointReferenceParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetSetpointReferenceParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetSetpointReference, error) {
	v, err := NewBACnetSetpointReference().parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_BACnetSetpointReference) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetSetpointReference BACnetSetpointReference, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetSetpointReference"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetSetpointReference")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	var setPointReference BACnetObjectPropertyReferenceEnclosed
	_setPointReference, err := ReadOptionalField[BACnetObjectPropertyReferenceEnclosed](ctx, "setPointReference", ReadComplex[BACnetObjectPropertyReferenceEnclosed](BACnetObjectPropertyReferenceEnclosedParseWithBufferProducer((uint8)(uint8(0))), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'setPointReference' field"))
	}
	if _setPointReference != nil {
		setPointReference = *_setPointReference
		m.SetPointReference = setPointReference
	}

	if closeErr := readBuffer.CloseContext("BACnetSetpointReference"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetSetpointReference")
	}

	return m, nil
}

func (m *_BACnetSetpointReference) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetSetpointReference) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetSetpointReference"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetSetpointReference")
	}

	if err := WriteOptionalField[BACnetObjectPropertyReferenceEnclosed](ctx, "setPointReference", GetRef(m.GetSetPointReference()), WriteComplex[BACnetObjectPropertyReferenceEnclosed](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'setPointReference' field")
	}

	if popErr := writeBuffer.PopContext("BACnetSetpointReference"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetSetpointReference")
	}
	return nil
}

func (m *_BACnetSetpointReference) isBACnetSetpointReference() bool {
	return true
}

func (m *_BACnetSetpointReference) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
