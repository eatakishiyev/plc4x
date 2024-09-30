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

// BACnetConstructedDataBaseDeviceSecurityPolicy is the corresponding interface of BACnetConstructedDataBaseDeviceSecurityPolicy
type BACnetConstructedDataBaseDeviceSecurityPolicy interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetBaseDeviceSecurityPolicy returns BaseDeviceSecurityPolicy (property field)
	GetBaseDeviceSecurityPolicy() BACnetSecurityLevelTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetSecurityLevelTagged
	// IsBACnetConstructedDataBaseDeviceSecurityPolicy is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataBaseDeviceSecurityPolicy()
	// CreateBuilder creates a BACnetConstructedDataBaseDeviceSecurityPolicyBuilder
	CreateBACnetConstructedDataBaseDeviceSecurityPolicyBuilder() BACnetConstructedDataBaseDeviceSecurityPolicyBuilder
}

// _BACnetConstructedDataBaseDeviceSecurityPolicy is the data-structure of this message
type _BACnetConstructedDataBaseDeviceSecurityPolicy struct {
	BACnetConstructedDataContract
	BaseDeviceSecurityPolicy BACnetSecurityLevelTagged
}

var _ BACnetConstructedDataBaseDeviceSecurityPolicy = (*_BACnetConstructedDataBaseDeviceSecurityPolicy)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataBaseDeviceSecurityPolicy)(nil)

// NewBACnetConstructedDataBaseDeviceSecurityPolicy factory function for _BACnetConstructedDataBaseDeviceSecurityPolicy
func NewBACnetConstructedDataBaseDeviceSecurityPolicy(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, baseDeviceSecurityPolicy BACnetSecurityLevelTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataBaseDeviceSecurityPolicy {
	if baseDeviceSecurityPolicy == nil {
		panic("baseDeviceSecurityPolicy of type BACnetSecurityLevelTagged for BACnetConstructedDataBaseDeviceSecurityPolicy must not be nil")
	}
	_result := &_BACnetConstructedDataBaseDeviceSecurityPolicy{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		BaseDeviceSecurityPolicy:      baseDeviceSecurityPolicy,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataBaseDeviceSecurityPolicyBuilder is a builder for BACnetConstructedDataBaseDeviceSecurityPolicy
type BACnetConstructedDataBaseDeviceSecurityPolicyBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(baseDeviceSecurityPolicy BACnetSecurityLevelTagged) BACnetConstructedDataBaseDeviceSecurityPolicyBuilder
	// WithBaseDeviceSecurityPolicy adds BaseDeviceSecurityPolicy (property field)
	WithBaseDeviceSecurityPolicy(BACnetSecurityLevelTagged) BACnetConstructedDataBaseDeviceSecurityPolicyBuilder
	// WithBaseDeviceSecurityPolicyBuilder adds BaseDeviceSecurityPolicy (property field) which is build by the builder
	WithBaseDeviceSecurityPolicyBuilder(func(BACnetSecurityLevelTaggedBuilder) BACnetSecurityLevelTaggedBuilder) BACnetConstructedDataBaseDeviceSecurityPolicyBuilder
	// Build builds the BACnetConstructedDataBaseDeviceSecurityPolicy or returns an error if something is wrong
	Build() (BACnetConstructedDataBaseDeviceSecurityPolicy, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataBaseDeviceSecurityPolicy
}

// NewBACnetConstructedDataBaseDeviceSecurityPolicyBuilder() creates a BACnetConstructedDataBaseDeviceSecurityPolicyBuilder
func NewBACnetConstructedDataBaseDeviceSecurityPolicyBuilder() BACnetConstructedDataBaseDeviceSecurityPolicyBuilder {
	return &_BACnetConstructedDataBaseDeviceSecurityPolicyBuilder{_BACnetConstructedDataBaseDeviceSecurityPolicy: new(_BACnetConstructedDataBaseDeviceSecurityPolicy)}
}

type _BACnetConstructedDataBaseDeviceSecurityPolicyBuilder struct {
	*_BACnetConstructedDataBaseDeviceSecurityPolicy

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataBaseDeviceSecurityPolicyBuilder) = (*_BACnetConstructedDataBaseDeviceSecurityPolicyBuilder)(nil)

func (b *_BACnetConstructedDataBaseDeviceSecurityPolicyBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataBaseDeviceSecurityPolicyBuilder) WithMandatoryFields(baseDeviceSecurityPolicy BACnetSecurityLevelTagged) BACnetConstructedDataBaseDeviceSecurityPolicyBuilder {
	return b.WithBaseDeviceSecurityPolicy(baseDeviceSecurityPolicy)
}

func (b *_BACnetConstructedDataBaseDeviceSecurityPolicyBuilder) WithBaseDeviceSecurityPolicy(baseDeviceSecurityPolicy BACnetSecurityLevelTagged) BACnetConstructedDataBaseDeviceSecurityPolicyBuilder {
	b.BaseDeviceSecurityPolicy = baseDeviceSecurityPolicy
	return b
}

func (b *_BACnetConstructedDataBaseDeviceSecurityPolicyBuilder) WithBaseDeviceSecurityPolicyBuilder(builderSupplier func(BACnetSecurityLevelTaggedBuilder) BACnetSecurityLevelTaggedBuilder) BACnetConstructedDataBaseDeviceSecurityPolicyBuilder {
	builder := builderSupplier(b.BaseDeviceSecurityPolicy.CreateBACnetSecurityLevelTaggedBuilder())
	var err error
	b.BaseDeviceSecurityPolicy, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetSecurityLevelTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataBaseDeviceSecurityPolicyBuilder) Build() (BACnetConstructedDataBaseDeviceSecurityPolicy, error) {
	if b.BaseDeviceSecurityPolicy == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'baseDeviceSecurityPolicy' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataBaseDeviceSecurityPolicy.deepCopy(), nil
}

func (b *_BACnetConstructedDataBaseDeviceSecurityPolicyBuilder) MustBuild() BACnetConstructedDataBaseDeviceSecurityPolicy {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataBaseDeviceSecurityPolicyBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataBaseDeviceSecurityPolicyBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataBaseDeviceSecurityPolicyBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataBaseDeviceSecurityPolicyBuilder().(*_BACnetConstructedDataBaseDeviceSecurityPolicyBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataBaseDeviceSecurityPolicyBuilder creates a BACnetConstructedDataBaseDeviceSecurityPolicyBuilder
func (b *_BACnetConstructedDataBaseDeviceSecurityPolicy) CreateBACnetConstructedDataBaseDeviceSecurityPolicyBuilder() BACnetConstructedDataBaseDeviceSecurityPolicyBuilder {
	if b == nil {
		return NewBACnetConstructedDataBaseDeviceSecurityPolicyBuilder()
	}
	return &_BACnetConstructedDataBaseDeviceSecurityPolicyBuilder{_BACnetConstructedDataBaseDeviceSecurityPolicy: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataBaseDeviceSecurityPolicy) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataBaseDeviceSecurityPolicy) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_BASE_DEVICE_SECURITY_POLICY
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataBaseDeviceSecurityPolicy) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataBaseDeviceSecurityPolicy) GetBaseDeviceSecurityPolicy() BACnetSecurityLevelTagged {
	return m.BaseDeviceSecurityPolicy
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataBaseDeviceSecurityPolicy) GetActualValue() BACnetSecurityLevelTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetSecurityLevelTagged(m.GetBaseDeviceSecurityPolicy())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataBaseDeviceSecurityPolicy(structType any) BACnetConstructedDataBaseDeviceSecurityPolicy {
	if casted, ok := structType.(BACnetConstructedDataBaseDeviceSecurityPolicy); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataBaseDeviceSecurityPolicy); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataBaseDeviceSecurityPolicy) GetTypeName() string {
	return "BACnetConstructedDataBaseDeviceSecurityPolicy"
}

func (m *_BACnetConstructedDataBaseDeviceSecurityPolicy) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (baseDeviceSecurityPolicy)
	lengthInBits += m.BaseDeviceSecurityPolicy.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataBaseDeviceSecurityPolicy) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataBaseDeviceSecurityPolicy) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataBaseDeviceSecurityPolicy BACnetConstructedDataBaseDeviceSecurityPolicy, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataBaseDeviceSecurityPolicy"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataBaseDeviceSecurityPolicy")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	baseDeviceSecurityPolicy, err := ReadSimpleField[BACnetSecurityLevelTagged](ctx, "baseDeviceSecurityPolicy", ReadComplex[BACnetSecurityLevelTagged](BACnetSecurityLevelTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'baseDeviceSecurityPolicy' field"))
	}
	m.BaseDeviceSecurityPolicy = baseDeviceSecurityPolicy

	actualValue, err := ReadVirtualField[BACnetSecurityLevelTagged](ctx, "actualValue", (*BACnetSecurityLevelTagged)(nil), baseDeviceSecurityPolicy)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataBaseDeviceSecurityPolicy"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataBaseDeviceSecurityPolicy")
	}

	return m, nil
}

func (m *_BACnetConstructedDataBaseDeviceSecurityPolicy) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataBaseDeviceSecurityPolicy) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataBaseDeviceSecurityPolicy"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataBaseDeviceSecurityPolicy")
		}

		if err := WriteSimpleField[BACnetSecurityLevelTagged](ctx, "baseDeviceSecurityPolicy", m.GetBaseDeviceSecurityPolicy(), WriteComplex[BACnetSecurityLevelTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'baseDeviceSecurityPolicy' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataBaseDeviceSecurityPolicy"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataBaseDeviceSecurityPolicy")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataBaseDeviceSecurityPolicy) IsBACnetConstructedDataBaseDeviceSecurityPolicy() {
}

func (m *_BACnetConstructedDataBaseDeviceSecurityPolicy) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataBaseDeviceSecurityPolicy) deepCopy() *_BACnetConstructedDataBaseDeviceSecurityPolicy {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataBaseDeviceSecurityPolicyCopy := &_BACnetConstructedDataBaseDeviceSecurityPolicy{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.BaseDeviceSecurityPolicy.DeepCopy().(BACnetSecurityLevelTagged),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataBaseDeviceSecurityPolicyCopy
}

func (m *_BACnetConstructedDataBaseDeviceSecurityPolicy) String() string {
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
