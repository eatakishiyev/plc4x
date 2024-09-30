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

// BACnetConfirmedServiceRequestReadProperty is the corresponding interface of BACnetConfirmedServiceRequestReadProperty
type BACnetConfirmedServiceRequestReadProperty interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConfirmedServiceRequest
	// GetObjectIdentifier returns ObjectIdentifier (property field)
	GetObjectIdentifier() BACnetContextTagObjectIdentifier
	// GetPropertyIdentifier returns PropertyIdentifier (property field)
	GetPropertyIdentifier() BACnetPropertyIdentifierTagged
	// GetArrayIndex returns ArrayIndex (property field)
	GetArrayIndex() BACnetContextTagUnsignedInteger
	// IsBACnetConfirmedServiceRequestReadProperty is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConfirmedServiceRequestReadProperty()
	// CreateBuilder creates a BACnetConfirmedServiceRequestReadPropertyBuilder
	CreateBACnetConfirmedServiceRequestReadPropertyBuilder() BACnetConfirmedServiceRequestReadPropertyBuilder
}

// _BACnetConfirmedServiceRequestReadProperty is the data-structure of this message
type _BACnetConfirmedServiceRequestReadProperty struct {
	BACnetConfirmedServiceRequestContract
	ObjectIdentifier   BACnetContextTagObjectIdentifier
	PropertyIdentifier BACnetPropertyIdentifierTagged
	ArrayIndex         BACnetContextTagUnsignedInteger
}

var _ BACnetConfirmedServiceRequestReadProperty = (*_BACnetConfirmedServiceRequestReadProperty)(nil)
var _ BACnetConfirmedServiceRequestRequirements = (*_BACnetConfirmedServiceRequestReadProperty)(nil)

// NewBACnetConfirmedServiceRequestReadProperty factory function for _BACnetConfirmedServiceRequestReadProperty
func NewBACnetConfirmedServiceRequestReadProperty(objectIdentifier BACnetContextTagObjectIdentifier, propertyIdentifier BACnetPropertyIdentifierTagged, arrayIndex BACnetContextTagUnsignedInteger, serviceRequestLength uint32) *_BACnetConfirmedServiceRequestReadProperty {
	if objectIdentifier == nil {
		panic("objectIdentifier of type BACnetContextTagObjectIdentifier for BACnetConfirmedServiceRequestReadProperty must not be nil")
	}
	if propertyIdentifier == nil {
		panic("propertyIdentifier of type BACnetPropertyIdentifierTagged for BACnetConfirmedServiceRequestReadProperty must not be nil")
	}
	_result := &_BACnetConfirmedServiceRequestReadProperty{
		BACnetConfirmedServiceRequestContract: NewBACnetConfirmedServiceRequest(serviceRequestLength),
		ObjectIdentifier:                      objectIdentifier,
		PropertyIdentifier:                    propertyIdentifier,
		ArrayIndex:                            arrayIndex,
	}
	_result.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConfirmedServiceRequestReadPropertyBuilder is a builder for BACnetConfirmedServiceRequestReadProperty
type BACnetConfirmedServiceRequestReadPropertyBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(objectIdentifier BACnetContextTagObjectIdentifier, propertyIdentifier BACnetPropertyIdentifierTagged) BACnetConfirmedServiceRequestReadPropertyBuilder
	// WithObjectIdentifier adds ObjectIdentifier (property field)
	WithObjectIdentifier(BACnetContextTagObjectIdentifier) BACnetConfirmedServiceRequestReadPropertyBuilder
	// WithObjectIdentifierBuilder adds ObjectIdentifier (property field) which is build by the builder
	WithObjectIdentifierBuilder(func(BACnetContextTagObjectIdentifierBuilder) BACnetContextTagObjectIdentifierBuilder) BACnetConfirmedServiceRequestReadPropertyBuilder
	// WithPropertyIdentifier adds PropertyIdentifier (property field)
	WithPropertyIdentifier(BACnetPropertyIdentifierTagged) BACnetConfirmedServiceRequestReadPropertyBuilder
	// WithPropertyIdentifierBuilder adds PropertyIdentifier (property field) which is build by the builder
	WithPropertyIdentifierBuilder(func(BACnetPropertyIdentifierTaggedBuilder) BACnetPropertyIdentifierTaggedBuilder) BACnetConfirmedServiceRequestReadPropertyBuilder
	// WithArrayIndex adds ArrayIndex (property field)
	WithOptionalArrayIndex(BACnetContextTagUnsignedInteger) BACnetConfirmedServiceRequestReadPropertyBuilder
	// WithOptionalArrayIndexBuilder adds ArrayIndex (property field) which is build by the builder
	WithOptionalArrayIndexBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetConfirmedServiceRequestReadPropertyBuilder
	// Build builds the BACnetConfirmedServiceRequestReadProperty or returns an error if something is wrong
	Build() (BACnetConfirmedServiceRequestReadProperty, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConfirmedServiceRequestReadProperty
}

// NewBACnetConfirmedServiceRequestReadPropertyBuilder() creates a BACnetConfirmedServiceRequestReadPropertyBuilder
func NewBACnetConfirmedServiceRequestReadPropertyBuilder() BACnetConfirmedServiceRequestReadPropertyBuilder {
	return &_BACnetConfirmedServiceRequestReadPropertyBuilder{_BACnetConfirmedServiceRequestReadProperty: new(_BACnetConfirmedServiceRequestReadProperty)}
}

type _BACnetConfirmedServiceRequestReadPropertyBuilder struct {
	*_BACnetConfirmedServiceRequestReadProperty

	parentBuilder *_BACnetConfirmedServiceRequestBuilder

	err *utils.MultiError
}

var _ (BACnetConfirmedServiceRequestReadPropertyBuilder) = (*_BACnetConfirmedServiceRequestReadPropertyBuilder)(nil)

func (b *_BACnetConfirmedServiceRequestReadPropertyBuilder) setParent(contract BACnetConfirmedServiceRequestContract) {
	b.BACnetConfirmedServiceRequestContract = contract
}

func (b *_BACnetConfirmedServiceRequestReadPropertyBuilder) WithMandatoryFields(objectIdentifier BACnetContextTagObjectIdentifier, propertyIdentifier BACnetPropertyIdentifierTagged) BACnetConfirmedServiceRequestReadPropertyBuilder {
	return b.WithObjectIdentifier(objectIdentifier).WithPropertyIdentifier(propertyIdentifier)
}

func (b *_BACnetConfirmedServiceRequestReadPropertyBuilder) WithObjectIdentifier(objectIdentifier BACnetContextTagObjectIdentifier) BACnetConfirmedServiceRequestReadPropertyBuilder {
	b.ObjectIdentifier = objectIdentifier
	return b
}

func (b *_BACnetConfirmedServiceRequestReadPropertyBuilder) WithObjectIdentifierBuilder(builderSupplier func(BACnetContextTagObjectIdentifierBuilder) BACnetContextTagObjectIdentifierBuilder) BACnetConfirmedServiceRequestReadPropertyBuilder {
	builder := builderSupplier(b.ObjectIdentifier.CreateBACnetContextTagObjectIdentifierBuilder())
	var err error
	b.ObjectIdentifier, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagObjectIdentifierBuilder failed"))
	}
	return b
}

func (b *_BACnetConfirmedServiceRequestReadPropertyBuilder) WithPropertyIdentifier(propertyIdentifier BACnetPropertyIdentifierTagged) BACnetConfirmedServiceRequestReadPropertyBuilder {
	b.PropertyIdentifier = propertyIdentifier
	return b
}

func (b *_BACnetConfirmedServiceRequestReadPropertyBuilder) WithPropertyIdentifierBuilder(builderSupplier func(BACnetPropertyIdentifierTaggedBuilder) BACnetPropertyIdentifierTaggedBuilder) BACnetConfirmedServiceRequestReadPropertyBuilder {
	builder := builderSupplier(b.PropertyIdentifier.CreateBACnetPropertyIdentifierTaggedBuilder())
	var err error
	b.PropertyIdentifier, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetPropertyIdentifierTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetConfirmedServiceRequestReadPropertyBuilder) WithOptionalArrayIndex(arrayIndex BACnetContextTagUnsignedInteger) BACnetConfirmedServiceRequestReadPropertyBuilder {
	b.ArrayIndex = arrayIndex
	return b
}

func (b *_BACnetConfirmedServiceRequestReadPropertyBuilder) WithOptionalArrayIndexBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetConfirmedServiceRequestReadPropertyBuilder {
	builder := builderSupplier(b.ArrayIndex.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	b.ArrayIndex, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConfirmedServiceRequestReadPropertyBuilder) Build() (BACnetConfirmedServiceRequestReadProperty, error) {
	if b.ObjectIdentifier == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'objectIdentifier' not set"))
	}
	if b.PropertyIdentifier == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'propertyIdentifier' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConfirmedServiceRequestReadProperty.deepCopy(), nil
}

func (b *_BACnetConfirmedServiceRequestReadPropertyBuilder) MustBuild() BACnetConfirmedServiceRequestReadProperty {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConfirmedServiceRequestReadPropertyBuilder) Done() BACnetConfirmedServiceRequestBuilder {
	return b.parentBuilder
}

func (b *_BACnetConfirmedServiceRequestReadPropertyBuilder) buildForBACnetConfirmedServiceRequest() (BACnetConfirmedServiceRequest, error) {
	return b.Build()
}

func (b *_BACnetConfirmedServiceRequestReadPropertyBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConfirmedServiceRequestReadPropertyBuilder().(*_BACnetConfirmedServiceRequestReadPropertyBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConfirmedServiceRequestReadPropertyBuilder creates a BACnetConfirmedServiceRequestReadPropertyBuilder
func (b *_BACnetConfirmedServiceRequestReadProperty) CreateBACnetConfirmedServiceRequestReadPropertyBuilder() BACnetConfirmedServiceRequestReadPropertyBuilder {
	if b == nil {
		return NewBACnetConfirmedServiceRequestReadPropertyBuilder()
	}
	return &_BACnetConfirmedServiceRequestReadPropertyBuilder{_BACnetConfirmedServiceRequestReadProperty: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConfirmedServiceRequestReadProperty) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_READ_PROPERTY
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestReadProperty) GetParent() BACnetConfirmedServiceRequestContract {
	return m.BACnetConfirmedServiceRequestContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestReadProperty) GetObjectIdentifier() BACnetContextTagObjectIdentifier {
	return m.ObjectIdentifier
}

func (m *_BACnetConfirmedServiceRequestReadProperty) GetPropertyIdentifier() BACnetPropertyIdentifierTagged {
	return m.PropertyIdentifier
}

func (m *_BACnetConfirmedServiceRequestReadProperty) GetArrayIndex() BACnetContextTagUnsignedInteger {
	return m.ArrayIndex
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestReadProperty(structType any) BACnetConfirmedServiceRequestReadProperty {
	if casted, ok := structType.(BACnetConfirmedServiceRequestReadProperty); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestReadProperty); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestReadProperty) GetTypeName() string {
	return "BACnetConfirmedServiceRequestReadProperty"
}

func (m *_BACnetConfirmedServiceRequestReadProperty) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).getLengthInBits(ctx))

	// Simple field (objectIdentifier)
	lengthInBits += m.ObjectIdentifier.GetLengthInBits(ctx)

	// Simple field (propertyIdentifier)
	lengthInBits += m.PropertyIdentifier.GetLengthInBits(ctx)

	// Optional Field (arrayIndex)
	if m.ArrayIndex != nil {
		lengthInBits += m.ArrayIndex.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestReadProperty) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConfirmedServiceRequestReadProperty) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConfirmedServiceRequest, serviceRequestLength uint32) (__bACnetConfirmedServiceRequestReadProperty BACnetConfirmedServiceRequestReadProperty, err error) {
	m.BACnetConfirmedServiceRequestContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestReadProperty"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestReadProperty")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	objectIdentifier, err := ReadSimpleField[BACnetContextTagObjectIdentifier](ctx, "objectIdentifier", ReadComplex[BACnetContextTagObjectIdentifier](BACnetContextTagParseWithBufferProducer[BACnetContextTagObjectIdentifier]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_BACNET_OBJECT_IDENTIFIER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'objectIdentifier' field"))
	}
	m.ObjectIdentifier = objectIdentifier

	propertyIdentifier, err := ReadSimpleField[BACnetPropertyIdentifierTagged](ctx, "propertyIdentifier", ReadComplex[BACnetPropertyIdentifierTagged](BACnetPropertyIdentifierTaggedParseWithBufferProducer((uint8)(uint8(1)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'propertyIdentifier' field"))
	}
	m.PropertyIdentifier = propertyIdentifier

	var arrayIndex BACnetContextTagUnsignedInteger
	_arrayIndex, err := ReadOptionalField[BACnetContextTagUnsignedInteger](ctx, "arrayIndex", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(2)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'arrayIndex' field"))
	}
	if _arrayIndex != nil {
		arrayIndex = *_arrayIndex
		m.ArrayIndex = arrayIndex
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestReadProperty"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestReadProperty")
	}

	return m, nil
}

func (m *_BACnetConfirmedServiceRequestReadProperty) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestReadProperty) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestReadProperty"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestReadProperty")
		}

		if err := WriteSimpleField[BACnetContextTagObjectIdentifier](ctx, "objectIdentifier", m.GetObjectIdentifier(), WriteComplex[BACnetContextTagObjectIdentifier](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'objectIdentifier' field")
		}

		if err := WriteSimpleField[BACnetPropertyIdentifierTagged](ctx, "propertyIdentifier", m.GetPropertyIdentifier(), WriteComplex[BACnetPropertyIdentifierTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'propertyIdentifier' field")
		}

		if err := WriteOptionalField[BACnetContextTagUnsignedInteger](ctx, "arrayIndex", GetRef(m.GetArrayIndex()), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'arrayIndex' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestReadProperty"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestReadProperty")
		}
		return nil
	}
	return m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConfirmedServiceRequestReadProperty) IsBACnetConfirmedServiceRequestReadProperty() {}

func (m *_BACnetConfirmedServiceRequestReadProperty) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConfirmedServiceRequestReadProperty) deepCopy() *_BACnetConfirmedServiceRequestReadProperty {
	if m == nil {
		return nil
	}
	_BACnetConfirmedServiceRequestReadPropertyCopy := &_BACnetConfirmedServiceRequestReadProperty{
		m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).deepCopy(),
		m.ObjectIdentifier.DeepCopy().(BACnetContextTagObjectIdentifier),
		m.PropertyIdentifier.DeepCopy().(BACnetPropertyIdentifierTagged),
		m.ArrayIndex.DeepCopy().(BACnetContextTagUnsignedInteger),
	}
	m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest)._SubType = m
	return _BACnetConfirmedServiceRequestReadPropertyCopy
}

func (m *_BACnetConfirmedServiceRequestReadProperty) String() string {
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
