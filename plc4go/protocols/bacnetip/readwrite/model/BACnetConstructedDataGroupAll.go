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

// BACnetConstructedDataGroupAll is the corresponding interface of BACnetConstructedDataGroupAll
type BACnetConstructedDataGroupAll interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// IsBACnetConstructedDataGroupAll is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataGroupAll()
	// CreateBuilder creates a BACnetConstructedDataGroupAllBuilder
	CreateBACnetConstructedDataGroupAllBuilder() BACnetConstructedDataGroupAllBuilder
}

// _BACnetConstructedDataGroupAll is the data-structure of this message
type _BACnetConstructedDataGroupAll struct {
	BACnetConstructedDataContract
}

var _ BACnetConstructedDataGroupAll = (*_BACnetConstructedDataGroupAll)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataGroupAll)(nil)

// NewBACnetConstructedDataGroupAll factory function for _BACnetConstructedDataGroupAll
func NewBACnetConstructedDataGroupAll(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataGroupAll {
	_result := &_BACnetConstructedDataGroupAll{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataGroupAllBuilder is a builder for BACnetConstructedDataGroupAll
type BACnetConstructedDataGroupAllBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() BACnetConstructedDataGroupAllBuilder
	// Build builds the BACnetConstructedDataGroupAll or returns an error if something is wrong
	Build() (BACnetConstructedDataGroupAll, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataGroupAll
}

// NewBACnetConstructedDataGroupAllBuilder() creates a BACnetConstructedDataGroupAllBuilder
func NewBACnetConstructedDataGroupAllBuilder() BACnetConstructedDataGroupAllBuilder {
	return &_BACnetConstructedDataGroupAllBuilder{_BACnetConstructedDataGroupAll: new(_BACnetConstructedDataGroupAll)}
}

type _BACnetConstructedDataGroupAllBuilder struct {
	*_BACnetConstructedDataGroupAll

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataGroupAllBuilder) = (*_BACnetConstructedDataGroupAllBuilder)(nil)

func (b *_BACnetConstructedDataGroupAllBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataGroupAllBuilder) WithMandatoryFields() BACnetConstructedDataGroupAllBuilder {
	return b
}

func (b *_BACnetConstructedDataGroupAllBuilder) Build() (BACnetConstructedDataGroupAll, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataGroupAll.deepCopy(), nil
}

func (b *_BACnetConstructedDataGroupAllBuilder) MustBuild() BACnetConstructedDataGroupAll {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataGroupAllBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataGroupAllBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataGroupAllBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataGroupAllBuilder().(*_BACnetConstructedDataGroupAllBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataGroupAllBuilder creates a BACnetConstructedDataGroupAllBuilder
func (b *_BACnetConstructedDataGroupAll) CreateBACnetConstructedDataGroupAllBuilder() BACnetConstructedDataGroupAllBuilder {
	if b == nil {
		return NewBACnetConstructedDataGroupAllBuilder()
	}
	return &_BACnetConstructedDataGroupAllBuilder{_BACnetConstructedDataGroupAll: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataGroupAll) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_GROUP
}

func (m *_BACnetConstructedDataGroupAll) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ALL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataGroupAll) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataGroupAll(structType any) BACnetConstructedDataGroupAll {
	if casted, ok := structType.(BACnetConstructedDataGroupAll); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataGroupAll); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataGroupAll) GetTypeName() string {
	return "BACnetConstructedDataGroupAll"
}

func (m *_BACnetConstructedDataGroupAll) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_BACnetConstructedDataGroupAll) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataGroupAll) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataGroupAll BACnetConstructedDataGroupAll, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataGroupAll"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataGroupAll")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((1) == (2))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "All should never occur in context of constructed data. If it does please report"})
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataGroupAll"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataGroupAll")
	}

	return m, nil
}

func (m *_BACnetConstructedDataGroupAll) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataGroupAll) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataGroupAll"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataGroupAll")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataGroupAll"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataGroupAll")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataGroupAll) IsBACnetConstructedDataGroupAll() {}

func (m *_BACnetConstructedDataGroupAll) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataGroupAll) deepCopy() *_BACnetConstructedDataGroupAll {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataGroupAllCopy := &_BACnetConstructedDataGroupAll{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataGroupAllCopy
}

func (m *_BACnetConstructedDataGroupAll) String() string {
	if m == nil {
		return "<nil>"
	}
	wb := utils.NewWriteBufferBoxBased(
		utils.WithWriteBufferBoxBasedMergeSingleBoxes(),
		utils.WithWriteBufferBoxBasedOmitEmptyBoxes(),
		utils.WithWriteBufferBoxBasedPrintPosLengthFooter(),
	)
	if err := wb.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return wb.GetBox().String()
}
