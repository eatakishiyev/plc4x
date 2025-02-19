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

// BACnetConstructedDataDatetimeValueAll is the corresponding interface of BACnetConstructedDataDatetimeValueAll
type BACnetConstructedDataDatetimeValueAll interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// IsBACnetConstructedDataDatetimeValueAll is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataDatetimeValueAll()
	// CreateBuilder creates a BACnetConstructedDataDatetimeValueAllBuilder
	CreateBACnetConstructedDataDatetimeValueAllBuilder() BACnetConstructedDataDatetimeValueAllBuilder
}

// _BACnetConstructedDataDatetimeValueAll is the data-structure of this message
type _BACnetConstructedDataDatetimeValueAll struct {
	BACnetConstructedDataContract
}

var _ BACnetConstructedDataDatetimeValueAll = (*_BACnetConstructedDataDatetimeValueAll)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataDatetimeValueAll)(nil)

// NewBACnetConstructedDataDatetimeValueAll factory function for _BACnetConstructedDataDatetimeValueAll
func NewBACnetConstructedDataDatetimeValueAll(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataDatetimeValueAll {
	_result := &_BACnetConstructedDataDatetimeValueAll{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataDatetimeValueAllBuilder is a builder for BACnetConstructedDataDatetimeValueAll
type BACnetConstructedDataDatetimeValueAllBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() BACnetConstructedDataDatetimeValueAllBuilder
	// Build builds the BACnetConstructedDataDatetimeValueAll or returns an error if something is wrong
	Build() (BACnetConstructedDataDatetimeValueAll, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataDatetimeValueAll
}

// NewBACnetConstructedDataDatetimeValueAllBuilder() creates a BACnetConstructedDataDatetimeValueAllBuilder
func NewBACnetConstructedDataDatetimeValueAllBuilder() BACnetConstructedDataDatetimeValueAllBuilder {
	return &_BACnetConstructedDataDatetimeValueAllBuilder{_BACnetConstructedDataDatetimeValueAll: new(_BACnetConstructedDataDatetimeValueAll)}
}

type _BACnetConstructedDataDatetimeValueAllBuilder struct {
	*_BACnetConstructedDataDatetimeValueAll

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataDatetimeValueAllBuilder) = (*_BACnetConstructedDataDatetimeValueAllBuilder)(nil)

func (b *_BACnetConstructedDataDatetimeValueAllBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataDatetimeValueAllBuilder) WithMandatoryFields() BACnetConstructedDataDatetimeValueAllBuilder {
	return b
}

func (b *_BACnetConstructedDataDatetimeValueAllBuilder) Build() (BACnetConstructedDataDatetimeValueAll, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataDatetimeValueAll.deepCopy(), nil
}

func (b *_BACnetConstructedDataDatetimeValueAllBuilder) MustBuild() BACnetConstructedDataDatetimeValueAll {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataDatetimeValueAllBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataDatetimeValueAllBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataDatetimeValueAllBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataDatetimeValueAllBuilder().(*_BACnetConstructedDataDatetimeValueAllBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataDatetimeValueAllBuilder creates a BACnetConstructedDataDatetimeValueAllBuilder
func (b *_BACnetConstructedDataDatetimeValueAll) CreateBACnetConstructedDataDatetimeValueAllBuilder() BACnetConstructedDataDatetimeValueAllBuilder {
	if b == nil {
		return NewBACnetConstructedDataDatetimeValueAllBuilder()
	}
	return &_BACnetConstructedDataDatetimeValueAllBuilder{_BACnetConstructedDataDatetimeValueAll: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataDatetimeValueAll) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_DATETIME_VALUE
}

func (m *_BACnetConstructedDataDatetimeValueAll) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ALL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataDatetimeValueAll) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataDatetimeValueAll(structType any) BACnetConstructedDataDatetimeValueAll {
	if casted, ok := structType.(BACnetConstructedDataDatetimeValueAll); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDatetimeValueAll); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataDatetimeValueAll) GetTypeName() string {
	return "BACnetConstructedDataDatetimeValueAll"
}

func (m *_BACnetConstructedDataDatetimeValueAll) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_BACnetConstructedDataDatetimeValueAll) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataDatetimeValueAll) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataDatetimeValueAll BACnetConstructedDataDatetimeValueAll, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDatetimeValueAll"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataDatetimeValueAll")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((1) == (2))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "All should never occur in context of constructed data. If it does please report"})
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDatetimeValueAll"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataDatetimeValueAll")
	}

	return m, nil
}

func (m *_BACnetConstructedDataDatetimeValueAll) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataDatetimeValueAll) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDatetimeValueAll"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataDatetimeValueAll")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDatetimeValueAll"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataDatetimeValueAll")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataDatetimeValueAll) IsBACnetConstructedDataDatetimeValueAll() {}

func (m *_BACnetConstructedDataDatetimeValueAll) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataDatetimeValueAll) deepCopy() *_BACnetConstructedDataDatetimeValueAll {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataDatetimeValueAllCopy := &_BACnetConstructedDataDatetimeValueAll{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataDatetimeValueAllCopy
}

func (m *_BACnetConstructedDataDatetimeValueAll) String() string {
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
