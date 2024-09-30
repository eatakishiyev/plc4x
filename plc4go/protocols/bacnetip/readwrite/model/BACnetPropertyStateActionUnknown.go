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

// BACnetPropertyStateActionUnknown is the corresponding interface of BACnetPropertyStateActionUnknown
type BACnetPropertyStateActionUnknown interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetPropertyStates
	// GetUnknownValue returns UnknownValue (property field)
	GetUnknownValue() BACnetContextTagUnknown
	// IsBACnetPropertyStateActionUnknown is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPropertyStateActionUnknown()
	// CreateBuilder creates a BACnetPropertyStateActionUnknownBuilder
	CreateBACnetPropertyStateActionUnknownBuilder() BACnetPropertyStateActionUnknownBuilder
}

// _BACnetPropertyStateActionUnknown is the data-structure of this message
type _BACnetPropertyStateActionUnknown struct {
	BACnetPropertyStatesContract
	UnknownValue BACnetContextTagUnknown
}

var _ BACnetPropertyStateActionUnknown = (*_BACnetPropertyStateActionUnknown)(nil)
var _ BACnetPropertyStatesRequirements = (*_BACnetPropertyStateActionUnknown)(nil)

// NewBACnetPropertyStateActionUnknown factory function for _BACnetPropertyStateActionUnknown
func NewBACnetPropertyStateActionUnknown(peekedTagHeader BACnetTagHeader, unknownValue BACnetContextTagUnknown) *_BACnetPropertyStateActionUnknown {
	if unknownValue == nil {
		panic("unknownValue of type BACnetContextTagUnknown for BACnetPropertyStateActionUnknown must not be nil")
	}
	_result := &_BACnetPropertyStateActionUnknown{
		BACnetPropertyStatesContract: NewBACnetPropertyStates(peekedTagHeader),
		UnknownValue:                 unknownValue,
	}
	_result.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetPropertyStateActionUnknownBuilder is a builder for BACnetPropertyStateActionUnknown
type BACnetPropertyStateActionUnknownBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(unknownValue BACnetContextTagUnknown) BACnetPropertyStateActionUnknownBuilder
	// WithUnknownValue adds UnknownValue (property field)
	WithUnknownValue(BACnetContextTagUnknown) BACnetPropertyStateActionUnknownBuilder
	// WithUnknownValueBuilder adds UnknownValue (property field) which is build by the builder
	WithUnknownValueBuilder(func(BACnetContextTagUnknownBuilder) BACnetContextTagUnknownBuilder) BACnetPropertyStateActionUnknownBuilder
	// Build builds the BACnetPropertyStateActionUnknown or returns an error if something is wrong
	Build() (BACnetPropertyStateActionUnknown, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetPropertyStateActionUnknown
}

// NewBACnetPropertyStateActionUnknownBuilder() creates a BACnetPropertyStateActionUnknownBuilder
func NewBACnetPropertyStateActionUnknownBuilder() BACnetPropertyStateActionUnknownBuilder {
	return &_BACnetPropertyStateActionUnknownBuilder{_BACnetPropertyStateActionUnknown: new(_BACnetPropertyStateActionUnknown)}
}

type _BACnetPropertyStateActionUnknownBuilder struct {
	*_BACnetPropertyStateActionUnknown

	parentBuilder *_BACnetPropertyStatesBuilder

	err *utils.MultiError
}

var _ (BACnetPropertyStateActionUnknownBuilder) = (*_BACnetPropertyStateActionUnknownBuilder)(nil)

func (b *_BACnetPropertyStateActionUnknownBuilder) setParent(contract BACnetPropertyStatesContract) {
	b.BACnetPropertyStatesContract = contract
}

func (b *_BACnetPropertyStateActionUnknownBuilder) WithMandatoryFields(unknownValue BACnetContextTagUnknown) BACnetPropertyStateActionUnknownBuilder {
	return b.WithUnknownValue(unknownValue)
}

func (b *_BACnetPropertyStateActionUnknownBuilder) WithUnknownValue(unknownValue BACnetContextTagUnknown) BACnetPropertyStateActionUnknownBuilder {
	b.UnknownValue = unknownValue
	return b
}

func (b *_BACnetPropertyStateActionUnknownBuilder) WithUnknownValueBuilder(builderSupplier func(BACnetContextTagUnknownBuilder) BACnetContextTagUnknownBuilder) BACnetPropertyStateActionUnknownBuilder {
	builder := builderSupplier(b.UnknownValue.CreateBACnetContextTagUnknownBuilder())
	var err error
	b.UnknownValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagUnknownBuilder failed"))
	}
	return b
}

func (b *_BACnetPropertyStateActionUnknownBuilder) Build() (BACnetPropertyStateActionUnknown, error) {
	if b.UnknownValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'unknownValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetPropertyStateActionUnknown.deepCopy(), nil
}

func (b *_BACnetPropertyStateActionUnknownBuilder) MustBuild() BACnetPropertyStateActionUnknown {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetPropertyStateActionUnknownBuilder) Done() BACnetPropertyStatesBuilder {
	return b.parentBuilder
}

func (b *_BACnetPropertyStateActionUnknownBuilder) buildForBACnetPropertyStates() (BACnetPropertyStates, error) {
	return b.Build()
}

func (b *_BACnetPropertyStateActionUnknownBuilder) DeepCopy() any {
	_copy := b.CreateBACnetPropertyStateActionUnknownBuilder().(*_BACnetPropertyStateActionUnknownBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetPropertyStateActionUnknownBuilder creates a BACnetPropertyStateActionUnknownBuilder
func (b *_BACnetPropertyStateActionUnknown) CreateBACnetPropertyStateActionUnknownBuilder() BACnetPropertyStateActionUnknownBuilder {
	if b == nil {
		return NewBACnetPropertyStateActionUnknownBuilder()
	}
	return &_BACnetPropertyStateActionUnknownBuilder{_BACnetPropertyStateActionUnknown: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStateActionUnknown) GetParent() BACnetPropertyStatesContract {
	return m.BACnetPropertyStatesContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStateActionUnknown) GetUnknownValue() BACnetContextTagUnknown {
	return m.UnknownValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStateActionUnknown(structType any) BACnetPropertyStateActionUnknown {
	if casted, ok := structType.(BACnetPropertyStateActionUnknown); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStateActionUnknown); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStateActionUnknown) GetTypeName() string {
	return "BACnetPropertyStateActionUnknown"
}

func (m *_BACnetPropertyStateActionUnknown) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).getLengthInBits(ctx))

	// Simple field (unknownValue)
	lengthInBits += m.UnknownValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStateActionUnknown) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPropertyStateActionUnknown) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPropertyStates, peekedTagNumber uint8) (__bACnetPropertyStateActionUnknown BACnetPropertyStateActionUnknown, err error) {
	m.BACnetPropertyStatesContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStateActionUnknown"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStateActionUnknown")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	unknownValue, err := ReadSimpleField[BACnetContextTagUnknown](ctx, "unknownValue", ReadComplex[BACnetContextTagUnknown](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnknown]((uint8)(peekedTagNumber), (BACnetDataType)(BACnetDataType_UNKNOWN)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'unknownValue' field"))
	}
	m.UnknownValue = unknownValue

	if closeErr := readBuffer.CloseContext("BACnetPropertyStateActionUnknown"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStateActionUnknown")
	}

	return m, nil
}

func (m *_BACnetPropertyStateActionUnknown) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStateActionUnknown) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStateActionUnknown"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStateActionUnknown")
		}

		if err := WriteSimpleField[BACnetContextTagUnknown](ctx, "unknownValue", m.GetUnknownValue(), WriteComplex[BACnetContextTagUnknown](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'unknownValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStateActionUnknown"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStateActionUnknown")
		}
		return nil
	}
	return m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStateActionUnknown) IsBACnetPropertyStateActionUnknown() {}

func (m *_BACnetPropertyStateActionUnknown) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetPropertyStateActionUnknown) deepCopy() *_BACnetPropertyStateActionUnknown {
	if m == nil {
		return nil
	}
	_BACnetPropertyStateActionUnknownCopy := &_BACnetPropertyStateActionUnknown{
		m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).deepCopy(),
		m.UnknownValue.DeepCopy().(BACnetContextTagUnknown),
	}
	m.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = m
	return _BACnetPropertyStateActionUnknownCopy
}

func (m *_BACnetPropertyStateActionUnknown) String() string {
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
