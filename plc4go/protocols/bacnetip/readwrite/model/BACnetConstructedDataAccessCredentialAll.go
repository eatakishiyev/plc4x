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

// BACnetConstructedDataAccessCredentialAll is the corresponding interface of BACnetConstructedDataAccessCredentialAll
type BACnetConstructedDataAccessCredentialAll interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// IsBACnetConstructedDataAccessCredentialAll is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataAccessCredentialAll()
	// CreateBuilder creates a BACnetConstructedDataAccessCredentialAllBuilder
	CreateBACnetConstructedDataAccessCredentialAllBuilder() BACnetConstructedDataAccessCredentialAllBuilder
}

// _BACnetConstructedDataAccessCredentialAll is the data-structure of this message
type _BACnetConstructedDataAccessCredentialAll struct {
	BACnetConstructedDataContract
}

var _ BACnetConstructedDataAccessCredentialAll = (*_BACnetConstructedDataAccessCredentialAll)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataAccessCredentialAll)(nil)

// NewBACnetConstructedDataAccessCredentialAll factory function for _BACnetConstructedDataAccessCredentialAll
func NewBACnetConstructedDataAccessCredentialAll(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAccessCredentialAll {
	_result := &_BACnetConstructedDataAccessCredentialAll{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataAccessCredentialAllBuilder is a builder for BACnetConstructedDataAccessCredentialAll
type BACnetConstructedDataAccessCredentialAllBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() BACnetConstructedDataAccessCredentialAllBuilder
	// Build builds the BACnetConstructedDataAccessCredentialAll or returns an error if something is wrong
	Build() (BACnetConstructedDataAccessCredentialAll, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataAccessCredentialAll
}

// NewBACnetConstructedDataAccessCredentialAllBuilder() creates a BACnetConstructedDataAccessCredentialAllBuilder
func NewBACnetConstructedDataAccessCredentialAllBuilder() BACnetConstructedDataAccessCredentialAllBuilder {
	return &_BACnetConstructedDataAccessCredentialAllBuilder{_BACnetConstructedDataAccessCredentialAll: new(_BACnetConstructedDataAccessCredentialAll)}
}

type _BACnetConstructedDataAccessCredentialAllBuilder struct {
	*_BACnetConstructedDataAccessCredentialAll

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataAccessCredentialAllBuilder) = (*_BACnetConstructedDataAccessCredentialAllBuilder)(nil)

func (b *_BACnetConstructedDataAccessCredentialAllBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataAccessCredentialAllBuilder) WithMandatoryFields() BACnetConstructedDataAccessCredentialAllBuilder {
	return b
}

func (b *_BACnetConstructedDataAccessCredentialAllBuilder) Build() (BACnetConstructedDataAccessCredentialAll, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataAccessCredentialAll.deepCopy(), nil
}

func (b *_BACnetConstructedDataAccessCredentialAllBuilder) MustBuild() BACnetConstructedDataAccessCredentialAll {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataAccessCredentialAllBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataAccessCredentialAllBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataAccessCredentialAllBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataAccessCredentialAllBuilder().(*_BACnetConstructedDataAccessCredentialAllBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataAccessCredentialAllBuilder creates a BACnetConstructedDataAccessCredentialAllBuilder
func (b *_BACnetConstructedDataAccessCredentialAll) CreateBACnetConstructedDataAccessCredentialAllBuilder() BACnetConstructedDataAccessCredentialAllBuilder {
	if b == nil {
		return NewBACnetConstructedDataAccessCredentialAllBuilder()
	}
	return &_BACnetConstructedDataAccessCredentialAllBuilder{_BACnetConstructedDataAccessCredentialAll: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAccessCredentialAll) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_ACCESS_CREDENTIAL
}

func (m *_BACnetConstructedDataAccessCredentialAll) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ALL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAccessCredentialAll) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAccessCredentialAll(structType any) BACnetConstructedDataAccessCredentialAll {
	if casted, ok := structType.(BACnetConstructedDataAccessCredentialAll); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAccessCredentialAll); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAccessCredentialAll) GetTypeName() string {
	return "BACnetConstructedDataAccessCredentialAll"
}

func (m *_BACnetConstructedDataAccessCredentialAll) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_BACnetConstructedDataAccessCredentialAll) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataAccessCredentialAll) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataAccessCredentialAll BACnetConstructedDataAccessCredentialAll, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAccessCredentialAll"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAccessCredentialAll")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((1) == (2))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "All should never occur in context of constructed data. If it does please report"})
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAccessCredentialAll"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAccessCredentialAll")
	}

	return m, nil
}

func (m *_BACnetConstructedDataAccessCredentialAll) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAccessCredentialAll) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAccessCredentialAll"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAccessCredentialAll")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAccessCredentialAll"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAccessCredentialAll")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAccessCredentialAll) IsBACnetConstructedDataAccessCredentialAll() {}

func (m *_BACnetConstructedDataAccessCredentialAll) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataAccessCredentialAll) deepCopy() *_BACnetConstructedDataAccessCredentialAll {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataAccessCredentialAllCopy := &_BACnetConstructedDataAccessCredentialAll{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataAccessCredentialAllCopy
}

func (m *_BACnetConstructedDataAccessCredentialAll) String() string {
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
