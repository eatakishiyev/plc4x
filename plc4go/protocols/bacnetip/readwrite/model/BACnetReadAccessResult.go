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

// BACnetReadAccessResult is the corresponding interface of BACnetReadAccessResult
type BACnetReadAccessResult interface {
	fmt.Stringer
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
	BACnetReadAccessResult
	isBACnetReadAccessResult() bool
}

// _BACnetReadAccessResult is the data-structure of this message
type _BACnetReadAccessResult struct {
	ObjectIdentifier BACnetContextTagObjectIdentifier
	ListOfResults    BACnetReadAccessResultListOfResults
}

var _ BACnetReadAccessResult = (*_BACnetReadAccessResult)(nil)

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
func CastBACnetReadAccessResult(structType any) BACnetReadAccessResult {
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

func (m *_BACnetReadAccessResult) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (objectIdentifier)
	lengthInBits += m.ObjectIdentifier.GetLengthInBits(ctx)

	// Optional Field (listOfResults)
	if m.ListOfResults != nil {
		lengthInBits += m.ListOfResults.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_BACnetReadAccessResult) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetReadAccessResultParse(ctx context.Context, theBytes []byte) (BACnetReadAccessResult, error) {
	return BACnetReadAccessResultParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetReadAccessResultParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetReadAccessResult, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetReadAccessResult, error) {
		return BACnetReadAccessResultParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetReadAccessResultParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetReadAccessResult, error) {
	v, err := NewBACnetReadAccessResult().parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_BACnetReadAccessResult) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetReadAccessResult BACnetReadAccessResult, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetReadAccessResult"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetReadAccessResult")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	objectIdentifier, err := ReadSimpleField[BACnetContextTagObjectIdentifier](ctx, "objectIdentifier", ReadComplex[BACnetContextTagObjectIdentifier](BACnetContextTagParseWithBufferProducer[BACnetContextTagObjectIdentifier]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_BACNET_OBJECT_IDENTIFIER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'objectIdentifier' field"))
	}
	m.ObjectIdentifier = objectIdentifier

	var listOfResults BACnetReadAccessResultListOfResults
	_listOfResults, err := ReadOptionalField[BACnetReadAccessResultListOfResults](ctx, "listOfResults", ReadComplex[BACnetReadAccessResultListOfResults](BACnetReadAccessResultListOfResultsParseWithBufferProducer((uint8)(uint8(1)), (BACnetObjectType)(objectIdentifier.GetObjectType())), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'listOfResults' field"))
	}
	if _listOfResults != nil {
		listOfResults = *_listOfResults
		m.ListOfResults = listOfResults
	}

	if closeErr := readBuffer.CloseContext("BACnetReadAccessResult"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetReadAccessResult")
	}

	return m, nil
}

func (m *_BACnetReadAccessResult) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetReadAccessResult) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetReadAccessResult"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetReadAccessResult")
	}

	if err := WriteSimpleField[BACnetContextTagObjectIdentifier](ctx, "objectIdentifier", m.GetObjectIdentifier(), WriteComplex[BACnetContextTagObjectIdentifier](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'objectIdentifier' field")
	}

	if err := WriteOptionalField[BACnetReadAccessResultListOfResults](ctx, "listOfResults", GetRef(m.GetListOfResults()), WriteComplex[BACnetReadAccessResultListOfResults](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'listOfResults' field")
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
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
