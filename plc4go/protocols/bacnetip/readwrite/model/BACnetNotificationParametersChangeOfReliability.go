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

// BACnetNotificationParametersChangeOfReliability is the corresponding interface of BACnetNotificationParametersChangeOfReliability
type BACnetNotificationParametersChangeOfReliability interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetNotificationParameters
	// GetInnerOpeningTag returns InnerOpeningTag (property field)
	GetInnerOpeningTag() BACnetOpeningTag
	// GetReliability returns Reliability (property field)
	GetReliability() BACnetReliabilityTagged
	// GetStatusFlags returns StatusFlags (property field)
	GetStatusFlags() BACnetStatusFlagsTagged
	// GetPropertyValues returns PropertyValues (property field)
	GetPropertyValues() BACnetPropertyValues
	// GetInnerClosingTag returns InnerClosingTag (property field)
	GetInnerClosingTag() BACnetClosingTag
	// IsBACnetNotificationParametersChangeOfReliability is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetNotificationParametersChangeOfReliability()
	// CreateBuilder creates a BACnetNotificationParametersChangeOfReliabilityBuilder
	CreateBACnetNotificationParametersChangeOfReliabilityBuilder() BACnetNotificationParametersChangeOfReliabilityBuilder
}

// _BACnetNotificationParametersChangeOfReliability is the data-structure of this message
type _BACnetNotificationParametersChangeOfReliability struct {
	BACnetNotificationParametersContract
	InnerOpeningTag BACnetOpeningTag
	Reliability     BACnetReliabilityTagged
	StatusFlags     BACnetStatusFlagsTagged
	PropertyValues  BACnetPropertyValues
	InnerClosingTag BACnetClosingTag
}

var _ BACnetNotificationParametersChangeOfReliability = (*_BACnetNotificationParametersChangeOfReliability)(nil)
var _ BACnetNotificationParametersRequirements = (*_BACnetNotificationParametersChangeOfReliability)(nil)

// NewBACnetNotificationParametersChangeOfReliability factory function for _BACnetNotificationParametersChangeOfReliability
func NewBACnetNotificationParametersChangeOfReliability(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, innerOpeningTag BACnetOpeningTag, reliability BACnetReliabilityTagged, statusFlags BACnetStatusFlagsTagged, propertyValues BACnetPropertyValues, innerClosingTag BACnetClosingTag, tagNumber uint8, objectTypeArgument BACnetObjectType) *_BACnetNotificationParametersChangeOfReliability {
	if innerOpeningTag == nil {
		panic("innerOpeningTag of type BACnetOpeningTag for BACnetNotificationParametersChangeOfReliability must not be nil")
	}
	if reliability == nil {
		panic("reliability of type BACnetReliabilityTagged for BACnetNotificationParametersChangeOfReliability must not be nil")
	}
	if statusFlags == nil {
		panic("statusFlags of type BACnetStatusFlagsTagged for BACnetNotificationParametersChangeOfReliability must not be nil")
	}
	if propertyValues == nil {
		panic("propertyValues of type BACnetPropertyValues for BACnetNotificationParametersChangeOfReliability must not be nil")
	}
	if innerClosingTag == nil {
		panic("innerClosingTag of type BACnetClosingTag for BACnetNotificationParametersChangeOfReliability must not be nil")
	}
	_result := &_BACnetNotificationParametersChangeOfReliability{
		BACnetNotificationParametersContract: NewBACnetNotificationParameters(openingTag, peekedTagHeader, closingTag, tagNumber, objectTypeArgument),
		InnerOpeningTag:                      innerOpeningTag,
		Reliability:                          reliability,
		StatusFlags:                          statusFlags,
		PropertyValues:                       propertyValues,
		InnerClosingTag:                      innerClosingTag,
	}
	_result.BACnetNotificationParametersContract.(*_BACnetNotificationParameters)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetNotificationParametersChangeOfReliabilityBuilder is a builder for BACnetNotificationParametersChangeOfReliability
type BACnetNotificationParametersChangeOfReliabilityBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(innerOpeningTag BACnetOpeningTag, reliability BACnetReliabilityTagged, statusFlags BACnetStatusFlagsTagged, propertyValues BACnetPropertyValues, innerClosingTag BACnetClosingTag) BACnetNotificationParametersChangeOfReliabilityBuilder
	// WithInnerOpeningTag adds InnerOpeningTag (property field)
	WithInnerOpeningTag(BACnetOpeningTag) BACnetNotificationParametersChangeOfReliabilityBuilder
	// WithInnerOpeningTagBuilder adds InnerOpeningTag (property field) which is build by the builder
	WithInnerOpeningTagBuilder(func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetNotificationParametersChangeOfReliabilityBuilder
	// WithReliability adds Reliability (property field)
	WithReliability(BACnetReliabilityTagged) BACnetNotificationParametersChangeOfReliabilityBuilder
	// WithReliabilityBuilder adds Reliability (property field) which is build by the builder
	WithReliabilityBuilder(func(BACnetReliabilityTaggedBuilder) BACnetReliabilityTaggedBuilder) BACnetNotificationParametersChangeOfReliabilityBuilder
	// WithStatusFlags adds StatusFlags (property field)
	WithStatusFlags(BACnetStatusFlagsTagged) BACnetNotificationParametersChangeOfReliabilityBuilder
	// WithStatusFlagsBuilder adds StatusFlags (property field) which is build by the builder
	WithStatusFlagsBuilder(func(BACnetStatusFlagsTaggedBuilder) BACnetStatusFlagsTaggedBuilder) BACnetNotificationParametersChangeOfReliabilityBuilder
	// WithPropertyValues adds PropertyValues (property field)
	WithPropertyValues(BACnetPropertyValues) BACnetNotificationParametersChangeOfReliabilityBuilder
	// WithPropertyValuesBuilder adds PropertyValues (property field) which is build by the builder
	WithPropertyValuesBuilder(func(BACnetPropertyValuesBuilder) BACnetPropertyValuesBuilder) BACnetNotificationParametersChangeOfReliabilityBuilder
	// WithInnerClosingTag adds InnerClosingTag (property field)
	WithInnerClosingTag(BACnetClosingTag) BACnetNotificationParametersChangeOfReliabilityBuilder
	// WithInnerClosingTagBuilder adds InnerClosingTag (property field) which is build by the builder
	WithInnerClosingTagBuilder(func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetNotificationParametersChangeOfReliabilityBuilder
	// Build builds the BACnetNotificationParametersChangeOfReliability or returns an error if something is wrong
	Build() (BACnetNotificationParametersChangeOfReliability, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetNotificationParametersChangeOfReliability
}

// NewBACnetNotificationParametersChangeOfReliabilityBuilder() creates a BACnetNotificationParametersChangeOfReliabilityBuilder
func NewBACnetNotificationParametersChangeOfReliabilityBuilder() BACnetNotificationParametersChangeOfReliabilityBuilder {
	return &_BACnetNotificationParametersChangeOfReliabilityBuilder{_BACnetNotificationParametersChangeOfReliability: new(_BACnetNotificationParametersChangeOfReliability)}
}

type _BACnetNotificationParametersChangeOfReliabilityBuilder struct {
	*_BACnetNotificationParametersChangeOfReliability

	parentBuilder *_BACnetNotificationParametersBuilder

	err *utils.MultiError
}

var _ (BACnetNotificationParametersChangeOfReliabilityBuilder) = (*_BACnetNotificationParametersChangeOfReliabilityBuilder)(nil)

func (b *_BACnetNotificationParametersChangeOfReliabilityBuilder) setParent(contract BACnetNotificationParametersContract) {
	b.BACnetNotificationParametersContract = contract
}

func (b *_BACnetNotificationParametersChangeOfReliabilityBuilder) WithMandatoryFields(innerOpeningTag BACnetOpeningTag, reliability BACnetReliabilityTagged, statusFlags BACnetStatusFlagsTagged, propertyValues BACnetPropertyValues, innerClosingTag BACnetClosingTag) BACnetNotificationParametersChangeOfReliabilityBuilder {
	return b.WithInnerOpeningTag(innerOpeningTag).WithReliability(reliability).WithStatusFlags(statusFlags).WithPropertyValues(propertyValues).WithInnerClosingTag(innerClosingTag)
}

func (b *_BACnetNotificationParametersChangeOfReliabilityBuilder) WithInnerOpeningTag(innerOpeningTag BACnetOpeningTag) BACnetNotificationParametersChangeOfReliabilityBuilder {
	b.InnerOpeningTag = innerOpeningTag
	return b
}

func (b *_BACnetNotificationParametersChangeOfReliabilityBuilder) WithInnerOpeningTagBuilder(builderSupplier func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetNotificationParametersChangeOfReliabilityBuilder {
	builder := builderSupplier(b.InnerOpeningTag.CreateBACnetOpeningTagBuilder())
	var err error
	b.InnerOpeningTag, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetOpeningTagBuilder failed"))
	}
	return b
}

func (b *_BACnetNotificationParametersChangeOfReliabilityBuilder) WithReliability(reliability BACnetReliabilityTagged) BACnetNotificationParametersChangeOfReliabilityBuilder {
	b.Reliability = reliability
	return b
}

func (b *_BACnetNotificationParametersChangeOfReliabilityBuilder) WithReliabilityBuilder(builderSupplier func(BACnetReliabilityTaggedBuilder) BACnetReliabilityTaggedBuilder) BACnetNotificationParametersChangeOfReliabilityBuilder {
	builder := builderSupplier(b.Reliability.CreateBACnetReliabilityTaggedBuilder())
	var err error
	b.Reliability, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetReliabilityTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetNotificationParametersChangeOfReliabilityBuilder) WithStatusFlags(statusFlags BACnetStatusFlagsTagged) BACnetNotificationParametersChangeOfReliabilityBuilder {
	b.StatusFlags = statusFlags
	return b
}

func (b *_BACnetNotificationParametersChangeOfReliabilityBuilder) WithStatusFlagsBuilder(builderSupplier func(BACnetStatusFlagsTaggedBuilder) BACnetStatusFlagsTaggedBuilder) BACnetNotificationParametersChangeOfReliabilityBuilder {
	builder := builderSupplier(b.StatusFlags.CreateBACnetStatusFlagsTaggedBuilder())
	var err error
	b.StatusFlags, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetStatusFlagsTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetNotificationParametersChangeOfReliabilityBuilder) WithPropertyValues(propertyValues BACnetPropertyValues) BACnetNotificationParametersChangeOfReliabilityBuilder {
	b.PropertyValues = propertyValues
	return b
}

func (b *_BACnetNotificationParametersChangeOfReliabilityBuilder) WithPropertyValuesBuilder(builderSupplier func(BACnetPropertyValuesBuilder) BACnetPropertyValuesBuilder) BACnetNotificationParametersChangeOfReliabilityBuilder {
	builder := builderSupplier(b.PropertyValues.CreateBACnetPropertyValuesBuilder())
	var err error
	b.PropertyValues, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetPropertyValuesBuilder failed"))
	}
	return b
}

func (b *_BACnetNotificationParametersChangeOfReliabilityBuilder) WithInnerClosingTag(innerClosingTag BACnetClosingTag) BACnetNotificationParametersChangeOfReliabilityBuilder {
	b.InnerClosingTag = innerClosingTag
	return b
}

func (b *_BACnetNotificationParametersChangeOfReliabilityBuilder) WithInnerClosingTagBuilder(builderSupplier func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetNotificationParametersChangeOfReliabilityBuilder {
	builder := builderSupplier(b.InnerClosingTag.CreateBACnetClosingTagBuilder())
	var err error
	b.InnerClosingTag, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetClosingTagBuilder failed"))
	}
	return b
}

func (b *_BACnetNotificationParametersChangeOfReliabilityBuilder) Build() (BACnetNotificationParametersChangeOfReliability, error) {
	if b.InnerOpeningTag == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'innerOpeningTag' not set"))
	}
	if b.Reliability == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'reliability' not set"))
	}
	if b.StatusFlags == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'statusFlags' not set"))
	}
	if b.PropertyValues == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'propertyValues' not set"))
	}
	if b.InnerClosingTag == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'innerClosingTag' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetNotificationParametersChangeOfReliability.deepCopy(), nil
}

func (b *_BACnetNotificationParametersChangeOfReliabilityBuilder) MustBuild() BACnetNotificationParametersChangeOfReliability {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetNotificationParametersChangeOfReliabilityBuilder) Done() BACnetNotificationParametersBuilder {
	return b.parentBuilder
}

func (b *_BACnetNotificationParametersChangeOfReliabilityBuilder) buildForBACnetNotificationParameters() (BACnetNotificationParameters, error) {
	return b.Build()
}

func (b *_BACnetNotificationParametersChangeOfReliabilityBuilder) DeepCopy() any {
	_copy := b.CreateBACnetNotificationParametersChangeOfReliabilityBuilder().(*_BACnetNotificationParametersChangeOfReliabilityBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetNotificationParametersChangeOfReliabilityBuilder creates a BACnetNotificationParametersChangeOfReliabilityBuilder
func (b *_BACnetNotificationParametersChangeOfReliability) CreateBACnetNotificationParametersChangeOfReliabilityBuilder() BACnetNotificationParametersChangeOfReliabilityBuilder {
	if b == nil {
		return NewBACnetNotificationParametersChangeOfReliabilityBuilder()
	}
	return &_BACnetNotificationParametersChangeOfReliabilityBuilder{_BACnetNotificationParametersChangeOfReliability: b.deepCopy()}
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

func (m *_BACnetNotificationParametersChangeOfReliability) GetParent() BACnetNotificationParametersContract {
	return m.BACnetNotificationParametersContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetNotificationParametersChangeOfReliability) GetInnerOpeningTag() BACnetOpeningTag {
	return m.InnerOpeningTag
}

func (m *_BACnetNotificationParametersChangeOfReliability) GetReliability() BACnetReliabilityTagged {
	return m.Reliability
}

func (m *_BACnetNotificationParametersChangeOfReliability) GetStatusFlags() BACnetStatusFlagsTagged {
	return m.StatusFlags
}

func (m *_BACnetNotificationParametersChangeOfReliability) GetPropertyValues() BACnetPropertyValues {
	return m.PropertyValues
}

func (m *_BACnetNotificationParametersChangeOfReliability) GetInnerClosingTag() BACnetClosingTag {
	return m.InnerClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetNotificationParametersChangeOfReliability(structType any) BACnetNotificationParametersChangeOfReliability {
	if casted, ok := structType.(BACnetNotificationParametersChangeOfReliability); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetNotificationParametersChangeOfReliability); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetNotificationParametersChangeOfReliability) GetTypeName() string {
	return "BACnetNotificationParametersChangeOfReliability"
}

func (m *_BACnetNotificationParametersChangeOfReliability) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetNotificationParametersContract.(*_BACnetNotificationParameters).getLengthInBits(ctx))

	// Simple field (innerOpeningTag)
	lengthInBits += m.InnerOpeningTag.GetLengthInBits(ctx)

	// Simple field (reliability)
	lengthInBits += m.Reliability.GetLengthInBits(ctx)

	// Simple field (statusFlags)
	lengthInBits += m.StatusFlags.GetLengthInBits(ctx)

	// Simple field (propertyValues)
	lengthInBits += m.PropertyValues.GetLengthInBits(ctx)

	// Simple field (innerClosingTag)
	lengthInBits += m.InnerClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetNotificationParametersChangeOfReliability) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetNotificationParametersChangeOfReliability) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetNotificationParameters, peekedTagNumber uint8, tagNumber uint8, objectTypeArgument BACnetObjectType) (__bACnetNotificationParametersChangeOfReliability BACnetNotificationParametersChangeOfReliability, err error) {
	m.BACnetNotificationParametersContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersChangeOfReliability"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetNotificationParametersChangeOfReliability")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	innerOpeningTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "innerOpeningTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(peekedTagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'innerOpeningTag' field"))
	}
	m.InnerOpeningTag = innerOpeningTag

	reliability, err := ReadSimpleField[BACnetReliabilityTagged](ctx, "reliability", ReadComplex[BACnetReliabilityTagged](BACnetReliabilityTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'reliability' field"))
	}
	m.Reliability = reliability

	statusFlags, err := ReadSimpleField[BACnetStatusFlagsTagged](ctx, "statusFlags", ReadComplex[BACnetStatusFlagsTagged](BACnetStatusFlagsTaggedParseWithBufferProducer((uint8)(uint8(1)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'statusFlags' field"))
	}
	m.StatusFlags = statusFlags

	propertyValues, err := ReadSimpleField[BACnetPropertyValues](ctx, "propertyValues", ReadComplex[BACnetPropertyValues](BACnetPropertyValuesParseWithBufferProducer((uint8)(uint8(2)), (BACnetObjectType)(objectTypeArgument)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'propertyValues' field"))
	}
	m.PropertyValues = propertyValues

	innerClosingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "innerClosingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(peekedTagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'innerClosingTag' field"))
	}
	m.InnerClosingTag = innerClosingTag

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersChangeOfReliability"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNotificationParametersChangeOfReliability")
	}

	return m, nil
}

func (m *_BACnetNotificationParametersChangeOfReliability) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetNotificationParametersChangeOfReliability) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetNotificationParametersChangeOfReliability"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetNotificationParametersChangeOfReliability")
		}

		if err := WriteSimpleField[BACnetOpeningTag](ctx, "innerOpeningTag", m.GetInnerOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'innerOpeningTag' field")
		}

		if err := WriteSimpleField[BACnetReliabilityTagged](ctx, "reliability", m.GetReliability(), WriteComplex[BACnetReliabilityTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'reliability' field")
		}

		if err := WriteSimpleField[BACnetStatusFlagsTagged](ctx, "statusFlags", m.GetStatusFlags(), WriteComplex[BACnetStatusFlagsTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'statusFlags' field")
		}

		if err := WriteSimpleField[BACnetPropertyValues](ctx, "propertyValues", m.GetPropertyValues(), WriteComplex[BACnetPropertyValues](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'propertyValues' field")
		}

		if err := WriteSimpleField[BACnetClosingTag](ctx, "innerClosingTag", m.GetInnerClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'innerClosingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetNotificationParametersChangeOfReliability"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetNotificationParametersChangeOfReliability")
		}
		return nil
	}
	return m.BACnetNotificationParametersContract.(*_BACnetNotificationParameters).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetNotificationParametersChangeOfReliability) IsBACnetNotificationParametersChangeOfReliability() {
}

func (m *_BACnetNotificationParametersChangeOfReliability) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetNotificationParametersChangeOfReliability) deepCopy() *_BACnetNotificationParametersChangeOfReliability {
	if m == nil {
		return nil
	}
	_BACnetNotificationParametersChangeOfReliabilityCopy := &_BACnetNotificationParametersChangeOfReliability{
		m.BACnetNotificationParametersContract.(*_BACnetNotificationParameters).deepCopy(),
		m.InnerOpeningTag.DeepCopy().(BACnetOpeningTag),
		m.Reliability.DeepCopy().(BACnetReliabilityTagged),
		m.StatusFlags.DeepCopy().(BACnetStatusFlagsTagged),
		m.PropertyValues.DeepCopy().(BACnetPropertyValues),
		m.InnerClosingTag.DeepCopy().(BACnetClosingTag),
	}
	m.BACnetNotificationParametersContract.(*_BACnetNotificationParameters)._SubType = m
	return _BACnetNotificationParametersChangeOfReliabilityCopy
}

func (m *_BACnetNotificationParametersChangeOfReliability) String() string {
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
