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

// HistoryReadDetails is the corresponding interface of HistoryReadDetails
type HistoryReadDetails interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// IsHistoryReadDetails is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsHistoryReadDetails()
	// CreateBuilder creates a HistoryReadDetailsBuilder
	CreateHistoryReadDetailsBuilder() HistoryReadDetailsBuilder
}

// _HistoryReadDetails is the data-structure of this message
type _HistoryReadDetails struct {
	ExtensionObjectDefinitionContract
}

var _ HistoryReadDetails = (*_HistoryReadDetails)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_HistoryReadDetails)(nil)

// NewHistoryReadDetails factory function for _HistoryReadDetails
func NewHistoryReadDetails() *_HistoryReadDetails {
	_result := &_HistoryReadDetails{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// HistoryReadDetailsBuilder is a builder for HistoryReadDetails
type HistoryReadDetailsBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() HistoryReadDetailsBuilder
	// Build builds the HistoryReadDetails or returns an error if something is wrong
	Build() (HistoryReadDetails, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() HistoryReadDetails
}

// NewHistoryReadDetailsBuilder() creates a HistoryReadDetailsBuilder
func NewHistoryReadDetailsBuilder() HistoryReadDetailsBuilder {
	return &_HistoryReadDetailsBuilder{_HistoryReadDetails: new(_HistoryReadDetails)}
}

type _HistoryReadDetailsBuilder struct {
	*_HistoryReadDetails

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (HistoryReadDetailsBuilder) = (*_HistoryReadDetailsBuilder)(nil)

func (b *_HistoryReadDetailsBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_HistoryReadDetailsBuilder) WithMandatoryFields() HistoryReadDetailsBuilder {
	return b
}

func (b *_HistoryReadDetailsBuilder) Build() (HistoryReadDetails, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._HistoryReadDetails.deepCopy(), nil
}

func (b *_HistoryReadDetailsBuilder) MustBuild() HistoryReadDetails {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_HistoryReadDetailsBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_HistoryReadDetailsBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_HistoryReadDetailsBuilder) DeepCopy() any {
	_copy := b.CreateHistoryReadDetailsBuilder().(*_HistoryReadDetailsBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateHistoryReadDetailsBuilder creates a HistoryReadDetailsBuilder
func (b *_HistoryReadDetails) CreateHistoryReadDetailsBuilder() HistoryReadDetailsBuilder {
	if b == nil {
		return NewHistoryReadDetailsBuilder()
	}
	return &_HistoryReadDetailsBuilder{_HistoryReadDetails: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_HistoryReadDetails) GetIdentifier() string {
	return "643"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_HistoryReadDetails) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

// Deprecated: use the interface for direct cast
func CastHistoryReadDetails(structType any) HistoryReadDetails {
	if casted, ok := structType.(HistoryReadDetails); ok {
		return casted
	}
	if casted, ok := structType.(*HistoryReadDetails); ok {
		return *casted
	}
	return nil
}

func (m *_HistoryReadDetails) GetTypeName() string {
	return "HistoryReadDetails"
}

func (m *_HistoryReadDetails) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_HistoryReadDetails) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_HistoryReadDetails) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__historyReadDetails HistoryReadDetails, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("HistoryReadDetails"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for HistoryReadDetails")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("HistoryReadDetails"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for HistoryReadDetails")
	}

	return m, nil
}

func (m *_HistoryReadDetails) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_HistoryReadDetails) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("HistoryReadDetails"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for HistoryReadDetails")
		}

		if popErr := writeBuffer.PopContext("HistoryReadDetails"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for HistoryReadDetails")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_HistoryReadDetails) IsHistoryReadDetails() {}

func (m *_HistoryReadDetails) DeepCopy() any {
	return m.deepCopy()
}

func (m *_HistoryReadDetails) deepCopy() *_HistoryReadDetails {
	if m == nil {
		return nil
	}
	_HistoryReadDetailsCopy := &_HistoryReadDetails{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _HistoryReadDetailsCopy
}

func (m *_HistoryReadDetails) String() string {
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
