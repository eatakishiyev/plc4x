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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// SemanticVersionString is the corresponding interface of SemanticVersionString
type SemanticVersionString interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// SemanticVersionStringExactly can be used when we want exactly this type and not a type which fulfills SemanticVersionString.
// This is useful for switch cases.
type SemanticVersionStringExactly interface {
	SemanticVersionString
	isSemanticVersionString() bool
}

// _SemanticVersionString is the data-structure of this message
type _SemanticVersionString struct {
}

var _ SemanticVersionString = (*_SemanticVersionString)(nil)

// NewSemanticVersionString factory function for _SemanticVersionString
func NewSemanticVersionString() *_SemanticVersionString {
	return &_SemanticVersionString{}
}

// Deprecated: use the interface for direct cast
func CastSemanticVersionString(structType any) SemanticVersionString {
	if casted, ok := structType.(SemanticVersionString); ok {
		return casted
	}
	if casted, ok := structType.(*SemanticVersionString); ok {
		return *casted
	}
	return nil
}

func (m *_SemanticVersionString) GetTypeName() string {
	return "SemanticVersionString"
}

func (m *_SemanticVersionString) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *_SemanticVersionString) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SemanticVersionStringParse(ctx context.Context, theBytes []byte) (SemanticVersionString, error) {
	return SemanticVersionStringParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func SemanticVersionStringParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (SemanticVersionString, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (SemanticVersionString, error) {
		return SemanticVersionStringParseWithBuffer(ctx, readBuffer)
	}
}

func SemanticVersionStringParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (SemanticVersionString, error) {
	v, err := NewSemanticVersionString().parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_SemanticVersionString) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__semanticVersionString SemanticVersionString, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SemanticVersionString"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SemanticVersionString")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SemanticVersionString"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SemanticVersionString")
	}

	return m, nil
}

func (m *_SemanticVersionString) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SemanticVersionString) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("SemanticVersionString"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for SemanticVersionString")
	}

	if popErr := writeBuffer.PopContext("SemanticVersionString"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for SemanticVersionString")
	}
	return nil
}

func (m *_SemanticVersionString) isSemanticVersionString() bool {
	return true
}

func (m *_SemanticVersionString) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
