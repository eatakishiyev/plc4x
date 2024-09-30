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

// BrowsePathResult is the corresponding interface of BrowsePathResult
type BrowsePathResult interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetStatusCode returns StatusCode (property field)
	GetStatusCode() StatusCode
	// GetNoOfTargets returns NoOfTargets (property field)
	GetNoOfTargets() int32
	// GetTargets returns Targets (property field)
	GetTargets() []ExtensionObjectDefinition
	// IsBrowsePathResult is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBrowsePathResult()
	// CreateBuilder creates a BrowsePathResultBuilder
	CreateBrowsePathResultBuilder() BrowsePathResultBuilder
}

// _BrowsePathResult is the data-structure of this message
type _BrowsePathResult struct {
	ExtensionObjectDefinitionContract
	StatusCode  StatusCode
	NoOfTargets int32
	Targets     []ExtensionObjectDefinition
}

var _ BrowsePathResult = (*_BrowsePathResult)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_BrowsePathResult)(nil)

// NewBrowsePathResult factory function for _BrowsePathResult
func NewBrowsePathResult(statusCode StatusCode, noOfTargets int32, targets []ExtensionObjectDefinition) *_BrowsePathResult {
	if statusCode == nil {
		panic("statusCode of type StatusCode for BrowsePathResult must not be nil")
	}
	_result := &_BrowsePathResult{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		StatusCode:                        statusCode,
		NoOfTargets:                       noOfTargets,
		Targets:                           targets,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BrowsePathResultBuilder is a builder for BrowsePathResult
type BrowsePathResultBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(statusCode StatusCode, noOfTargets int32, targets []ExtensionObjectDefinition) BrowsePathResultBuilder
	// WithStatusCode adds StatusCode (property field)
	WithStatusCode(StatusCode) BrowsePathResultBuilder
	// WithStatusCodeBuilder adds StatusCode (property field) which is build by the builder
	WithStatusCodeBuilder(func(StatusCodeBuilder) StatusCodeBuilder) BrowsePathResultBuilder
	// WithNoOfTargets adds NoOfTargets (property field)
	WithNoOfTargets(int32) BrowsePathResultBuilder
	// WithTargets adds Targets (property field)
	WithTargets(...ExtensionObjectDefinition) BrowsePathResultBuilder
	// Build builds the BrowsePathResult or returns an error if something is wrong
	Build() (BrowsePathResult, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BrowsePathResult
}

// NewBrowsePathResultBuilder() creates a BrowsePathResultBuilder
func NewBrowsePathResultBuilder() BrowsePathResultBuilder {
	return &_BrowsePathResultBuilder{_BrowsePathResult: new(_BrowsePathResult)}
}

type _BrowsePathResultBuilder struct {
	*_BrowsePathResult

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (BrowsePathResultBuilder) = (*_BrowsePathResultBuilder)(nil)

func (b *_BrowsePathResultBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_BrowsePathResultBuilder) WithMandatoryFields(statusCode StatusCode, noOfTargets int32, targets []ExtensionObjectDefinition) BrowsePathResultBuilder {
	return b.WithStatusCode(statusCode).WithNoOfTargets(noOfTargets).WithTargets(targets...)
}

func (b *_BrowsePathResultBuilder) WithStatusCode(statusCode StatusCode) BrowsePathResultBuilder {
	b.StatusCode = statusCode
	return b
}

func (b *_BrowsePathResultBuilder) WithStatusCodeBuilder(builderSupplier func(StatusCodeBuilder) StatusCodeBuilder) BrowsePathResultBuilder {
	builder := builderSupplier(b.StatusCode.CreateStatusCodeBuilder())
	var err error
	b.StatusCode, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "StatusCodeBuilder failed"))
	}
	return b
}

func (b *_BrowsePathResultBuilder) WithNoOfTargets(noOfTargets int32) BrowsePathResultBuilder {
	b.NoOfTargets = noOfTargets
	return b
}

func (b *_BrowsePathResultBuilder) WithTargets(targets ...ExtensionObjectDefinition) BrowsePathResultBuilder {
	b.Targets = targets
	return b
}

func (b *_BrowsePathResultBuilder) Build() (BrowsePathResult, error) {
	if b.StatusCode == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'statusCode' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BrowsePathResult.deepCopy(), nil
}

func (b *_BrowsePathResultBuilder) MustBuild() BrowsePathResult {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BrowsePathResultBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_BrowsePathResultBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_BrowsePathResultBuilder) DeepCopy() any {
	_copy := b.CreateBrowsePathResultBuilder().(*_BrowsePathResultBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBrowsePathResultBuilder creates a BrowsePathResultBuilder
func (b *_BrowsePathResult) CreateBrowsePathResultBuilder() BrowsePathResultBuilder {
	if b == nil {
		return NewBrowsePathResultBuilder()
	}
	return &_BrowsePathResultBuilder{_BrowsePathResult: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BrowsePathResult) GetIdentifier() string {
	return "551"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BrowsePathResult) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BrowsePathResult) GetStatusCode() StatusCode {
	return m.StatusCode
}

func (m *_BrowsePathResult) GetNoOfTargets() int32 {
	return m.NoOfTargets
}

func (m *_BrowsePathResult) GetTargets() []ExtensionObjectDefinition {
	return m.Targets
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBrowsePathResult(structType any) BrowsePathResult {
	if casted, ok := structType.(BrowsePathResult); ok {
		return casted
	}
	if casted, ok := structType.(*BrowsePathResult); ok {
		return *casted
	}
	return nil
}

func (m *_BrowsePathResult) GetTypeName() string {
	return "BrowsePathResult"
}

func (m *_BrowsePathResult) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (statusCode)
	lengthInBits += m.StatusCode.GetLengthInBits(ctx)

	// Simple field (noOfTargets)
	lengthInBits += 32

	// Array field
	if len(m.Targets) > 0 {
		for _curItem, element := range m.Targets {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Targets), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_BrowsePathResult) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BrowsePathResult) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__browsePathResult BrowsePathResult, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BrowsePathResult"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BrowsePathResult")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	statusCode, err := ReadSimpleField[StatusCode](ctx, "statusCode", ReadComplex[StatusCode](StatusCodeParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'statusCode' field"))
	}
	m.StatusCode = statusCode

	noOfTargets, err := ReadSimpleField(ctx, "noOfTargets", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfTargets' field"))
	}
	m.NoOfTargets = noOfTargets

	targets, err := ReadCountArrayField[ExtensionObjectDefinition](ctx, "targets", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("548")), readBuffer), uint64(noOfTargets))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'targets' field"))
	}
	m.Targets = targets

	if closeErr := readBuffer.CloseContext("BrowsePathResult"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BrowsePathResult")
	}

	return m, nil
}

func (m *_BrowsePathResult) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BrowsePathResult) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BrowsePathResult"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BrowsePathResult")
		}

		if err := WriteSimpleField[StatusCode](ctx, "statusCode", m.GetStatusCode(), WriteComplex[StatusCode](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'statusCode' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfTargets", m.GetNoOfTargets(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfTargets' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "targets", m.GetTargets(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'targets' field")
		}

		if popErr := writeBuffer.PopContext("BrowsePathResult"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BrowsePathResult")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BrowsePathResult) IsBrowsePathResult() {}

func (m *_BrowsePathResult) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BrowsePathResult) deepCopy() *_BrowsePathResult {
	if m == nil {
		return nil
	}
	_BrowsePathResultCopy := &_BrowsePathResult{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.StatusCode.DeepCopy().(StatusCode),
		m.NoOfTargets,
		utils.DeepCopySlice[ExtensionObjectDefinition, ExtensionObjectDefinition](m.Targets),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _BrowsePathResultCopy
}

func (m *_BrowsePathResult) String() string {
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
