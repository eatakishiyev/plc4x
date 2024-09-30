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

// AxisInformation is the corresponding interface of AxisInformation
type AxisInformation interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetEngineeringUnits returns EngineeringUnits (property field)
	GetEngineeringUnits() ExtensionObjectDefinition
	// GetEURange returns EURange (property field)
	GetEURange() ExtensionObjectDefinition
	// GetTitle returns Title (property field)
	GetTitle() LocalizedText
	// GetAxisScaleType returns AxisScaleType (property field)
	GetAxisScaleType() AxisScaleEnumeration
	// GetNoOfAxisSteps returns NoOfAxisSteps (property field)
	GetNoOfAxisSteps() int32
	// GetAxisSteps returns AxisSteps (property field)
	GetAxisSteps() []float64
	// IsAxisInformation is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAxisInformation()
	// CreateBuilder creates a AxisInformationBuilder
	CreateAxisInformationBuilder() AxisInformationBuilder
}

// _AxisInformation is the data-structure of this message
type _AxisInformation struct {
	ExtensionObjectDefinitionContract
	EngineeringUnits ExtensionObjectDefinition
	EURange          ExtensionObjectDefinition
	Title            LocalizedText
	AxisScaleType    AxisScaleEnumeration
	NoOfAxisSteps    int32
	AxisSteps        []float64
}

var _ AxisInformation = (*_AxisInformation)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_AxisInformation)(nil)

// NewAxisInformation factory function for _AxisInformation
func NewAxisInformation(engineeringUnits ExtensionObjectDefinition, eURange ExtensionObjectDefinition, title LocalizedText, axisScaleType AxisScaleEnumeration, noOfAxisSteps int32, axisSteps []float64) *_AxisInformation {
	if engineeringUnits == nil {
		panic("engineeringUnits of type ExtensionObjectDefinition for AxisInformation must not be nil")
	}
	if eURange == nil {
		panic("eURange of type ExtensionObjectDefinition for AxisInformation must not be nil")
	}
	if title == nil {
		panic("title of type LocalizedText for AxisInformation must not be nil")
	}
	_result := &_AxisInformation{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		EngineeringUnits:                  engineeringUnits,
		EURange:                           eURange,
		Title:                             title,
		AxisScaleType:                     axisScaleType,
		NoOfAxisSteps:                     noOfAxisSteps,
		AxisSteps:                         axisSteps,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// AxisInformationBuilder is a builder for AxisInformation
type AxisInformationBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(engineeringUnits ExtensionObjectDefinition, eURange ExtensionObjectDefinition, title LocalizedText, axisScaleType AxisScaleEnumeration, noOfAxisSteps int32, axisSteps []float64) AxisInformationBuilder
	// WithEngineeringUnits adds EngineeringUnits (property field)
	WithEngineeringUnits(ExtensionObjectDefinition) AxisInformationBuilder
	// WithEngineeringUnitsBuilder adds EngineeringUnits (property field) which is build by the builder
	WithEngineeringUnitsBuilder(func(ExtensionObjectDefinitionBuilder) ExtensionObjectDefinitionBuilder) AxisInformationBuilder
	// WithEURange adds EURange (property field)
	WithEURange(ExtensionObjectDefinition) AxisInformationBuilder
	// WithEURangeBuilder adds EURange (property field) which is build by the builder
	WithEURangeBuilder(func(ExtensionObjectDefinitionBuilder) ExtensionObjectDefinitionBuilder) AxisInformationBuilder
	// WithTitle adds Title (property field)
	WithTitle(LocalizedText) AxisInformationBuilder
	// WithTitleBuilder adds Title (property field) which is build by the builder
	WithTitleBuilder(func(LocalizedTextBuilder) LocalizedTextBuilder) AxisInformationBuilder
	// WithAxisScaleType adds AxisScaleType (property field)
	WithAxisScaleType(AxisScaleEnumeration) AxisInformationBuilder
	// WithNoOfAxisSteps adds NoOfAxisSteps (property field)
	WithNoOfAxisSteps(int32) AxisInformationBuilder
	// WithAxisSteps adds AxisSteps (property field)
	WithAxisSteps(...float64) AxisInformationBuilder
	// Build builds the AxisInformation or returns an error if something is wrong
	Build() (AxisInformation, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() AxisInformation
}

// NewAxisInformationBuilder() creates a AxisInformationBuilder
func NewAxisInformationBuilder() AxisInformationBuilder {
	return &_AxisInformationBuilder{_AxisInformation: new(_AxisInformation)}
}

type _AxisInformationBuilder struct {
	*_AxisInformation

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (AxisInformationBuilder) = (*_AxisInformationBuilder)(nil)

func (b *_AxisInformationBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_AxisInformationBuilder) WithMandatoryFields(engineeringUnits ExtensionObjectDefinition, eURange ExtensionObjectDefinition, title LocalizedText, axisScaleType AxisScaleEnumeration, noOfAxisSteps int32, axisSteps []float64) AxisInformationBuilder {
	return b.WithEngineeringUnits(engineeringUnits).WithEURange(eURange).WithTitle(title).WithAxisScaleType(axisScaleType).WithNoOfAxisSteps(noOfAxisSteps).WithAxisSteps(axisSteps...)
}

func (b *_AxisInformationBuilder) WithEngineeringUnits(engineeringUnits ExtensionObjectDefinition) AxisInformationBuilder {
	b.EngineeringUnits = engineeringUnits
	return b
}

func (b *_AxisInformationBuilder) WithEngineeringUnitsBuilder(builderSupplier func(ExtensionObjectDefinitionBuilder) ExtensionObjectDefinitionBuilder) AxisInformationBuilder {
	builder := builderSupplier(b.EngineeringUnits.CreateExtensionObjectDefinitionBuilder())
	var err error
	b.EngineeringUnits, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ExtensionObjectDefinitionBuilder failed"))
	}
	return b
}

func (b *_AxisInformationBuilder) WithEURange(eURange ExtensionObjectDefinition) AxisInformationBuilder {
	b.EURange = eURange
	return b
}

func (b *_AxisInformationBuilder) WithEURangeBuilder(builderSupplier func(ExtensionObjectDefinitionBuilder) ExtensionObjectDefinitionBuilder) AxisInformationBuilder {
	builder := builderSupplier(b.EURange.CreateExtensionObjectDefinitionBuilder())
	var err error
	b.EURange, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ExtensionObjectDefinitionBuilder failed"))
	}
	return b
}

func (b *_AxisInformationBuilder) WithTitle(title LocalizedText) AxisInformationBuilder {
	b.Title = title
	return b
}

func (b *_AxisInformationBuilder) WithTitleBuilder(builderSupplier func(LocalizedTextBuilder) LocalizedTextBuilder) AxisInformationBuilder {
	builder := builderSupplier(b.Title.CreateLocalizedTextBuilder())
	var err error
	b.Title, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "LocalizedTextBuilder failed"))
	}
	return b
}

func (b *_AxisInformationBuilder) WithAxisScaleType(axisScaleType AxisScaleEnumeration) AxisInformationBuilder {
	b.AxisScaleType = axisScaleType
	return b
}

func (b *_AxisInformationBuilder) WithNoOfAxisSteps(noOfAxisSteps int32) AxisInformationBuilder {
	b.NoOfAxisSteps = noOfAxisSteps
	return b
}

func (b *_AxisInformationBuilder) WithAxisSteps(axisSteps ...float64) AxisInformationBuilder {
	b.AxisSteps = axisSteps
	return b
}

func (b *_AxisInformationBuilder) Build() (AxisInformation, error) {
	if b.EngineeringUnits == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'engineeringUnits' not set"))
	}
	if b.EURange == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'eURange' not set"))
	}
	if b.Title == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'title' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._AxisInformation.deepCopy(), nil
}

func (b *_AxisInformationBuilder) MustBuild() AxisInformation {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_AxisInformationBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_AxisInformationBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_AxisInformationBuilder) DeepCopy() any {
	_copy := b.CreateAxisInformationBuilder().(*_AxisInformationBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateAxisInformationBuilder creates a AxisInformationBuilder
func (b *_AxisInformation) CreateAxisInformationBuilder() AxisInformationBuilder {
	if b == nil {
		return NewAxisInformationBuilder()
	}
	return &_AxisInformationBuilder{_AxisInformation: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AxisInformation) GetIdentifier() string {
	return "12081"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AxisInformation) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AxisInformation) GetEngineeringUnits() ExtensionObjectDefinition {
	return m.EngineeringUnits
}

func (m *_AxisInformation) GetEURange() ExtensionObjectDefinition {
	return m.EURange
}

func (m *_AxisInformation) GetTitle() LocalizedText {
	return m.Title
}

func (m *_AxisInformation) GetAxisScaleType() AxisScaleEnumeration {
	return m.AxisScaleType
}

func (m *_AxisInformation) GetNoOfAxisSteps() int32 {
	return m.NoOfAxisSteps
}

func (m *_AxisInformation) GetAxisSteps() []float64 {
	return m.AxisSteps
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastAxisInformation(structType any) AxisInformation {
	if casted, ok := structType.(AxisInformation); ok {
		return casted
	}
	if casted, ok := structType.(*AxisInformation); ok {
		return *casted
	}
	return nil
}

func (m *_AxisInformation) GetTypeName() string {
	return "AxisInformation"
}

func (m *_AxisInformation) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (engineeringUnits)
	lengthInBits += m.EngineeringUnits.GetLengthInBits(ctx)

	// Simple field (eURange)
	lengthInBits += m.EURange.GetLengthInBits(ctx)

	// Simple field (title)
	lengthInBits += m.Title.GetLengthInBits(ctx)

	// Simple field (axisScaleType)
	lengthInBits += 32

	// Simple field (noOfAxisSteps)
	lengthInBits += 32

	// Array field
	if len(m.AxisSteps) > 0 {
		lengthInBits += 64 * uint16(len(m.AxisSteps))
	}

	return lengthInBits
}

func (m *_AxisInformation) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AxisInformation) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__axisInformation AxisInformation, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AxisInformation"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AxisInformation")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	engineeringUnits, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "engineeringUnits", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("889")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'engineeringUnits' field"))
	}
	m.EngineeringUnits = engineeringUnits

	eURange, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "eURange", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("886")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'eURange' field"))
	}
	m.EURange = eURange

	title, err := ReadSimpleField[LocalizedText](ctx, "title", ReadComplex[LocalizedText](LocalizedTextParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'title' field"))
	}
	m.Title = title

	axisScaleType, err := ReadEnumField[AxisScaleEnumeration](ctx, "axisScaleType", "AxisScaleEnumeration", ReadEnum(AxisScaleEnumerationByValue, ReadUnsignedInt(readBuffer, uint8(32))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'axisScaleType' field"))
	}
	m.AxisScaleType = axisScaleType

	noOfAxisSteps, err := ReadSimpleField(ctx, "noOfAxisSteps", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfAxisSteps' field"))
	}
	m.NoOfAxisSteps = noOfAxisSteps

	axisSteps, err := ReadCountArrayField[float64](ctx, "axisSteps", ReadDouble(readBuffer, uint8(64)), uint64(noOfAxisSteps))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'axisSteps' field"))
	}
	m.AxisSteps = axisSteps

	if closeErr := readBuffer.CloseContext("AxisInformation"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AxisInformation")
	}

	return m, nil
}

func (m *_AxisInformation) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AxisInformation) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AxisInformation"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AxisInformation")
		}

		if err := WriteSimpleField[ExtensionObjectDefinition](ctx, "engineeringUnits", m.GetEngineeringUnits(), WriteComplex[ExtensionObjectDefinition](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'engineeringUnits' field")
		}

		if err := WriteSimpleField[ExtensionObjectDefinition](ctx, "eURange", m.GetEURange(), WriteComplex[ExtensionObjectDefinition](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'eURange' field")
		}

		if err := WriteSimpleField[LocalizedText](ctx, "title", m.GetTitle(), WriteComplex[LocalizedText](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'title' field")
		}

		if err := WriteSimpleEnumField[AxisScaleEnumeration](ctx, "axisScaleType", "AxisScaleEnumeration", m.GetAxisScaleType(), WriteEnum[AxisScaleEnumeration, uint32](AxisScaleEnumeration.GetValue, AxisScaleEnumeration.PLC4XEnumName, WriteUnsignedInt(writeBuffer, 32))); err != nil {
			return errors.Wrap(err, "Error serializing 'axisScaleType' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfAxisSteps", m.GetNoOfAxisSteps(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfAxisSteps' field")
		}

		if err := WriteSimpleTypeArrayField(ctx, "axisSteps", m.GetAxisSteps(), WriteDouble(writeBuffer, 64)); err != nil {
			return errors.Wrap(err, "Error serializing 'axisSteps' field")
		}

		if popErr := writeBuffer.PopContext("AxisInformation"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AxisInformation")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AxisInformation) IsAxisInformation() {}

func (m *_AxisInformation) DeepCopy() any {
	return m.deepCopy()
}

func (m *_AxisInformation) deepCopy() *_AxisInformation {
	if m == nil {
		return nil
	}
	_AxisInformationCopy := &_AxisInformation{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.EngineeringUnits.DeepCopy().(ExtensionObjectDefinition),
		m.EURange.DeepCopy().(ExtensionObjectDefinition),
		m.Title.DeepCopy().(LocalizedText),
		m.AxisScaleType,
		m.NoOfAxisSteps,
		utils.DeepCopySlice[float64, float64](m.AxisSteps),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _AxisInformationCopy
}

func (m *_AxisInformation) String() string {
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
