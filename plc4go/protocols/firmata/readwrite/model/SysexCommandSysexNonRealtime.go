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

// SysexCommandSysexNonRealtime is the corresponding interface of SysexCommandSysexNonRealtime
type SysexCommandSysexNonRealtime interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	SysexCommand
	// IsSysexCommandSysexNonRealtime is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSysexCommandSysexNonRealtime()
	// CreateBuilder creates a SysexCommandSysexNonRealtimeBuilder
	CreateSysexCommandSysexNonRealtimeBuilder() SysexCommandSysexNonRealtimeBuilder
}

// _SysexCommandSysexNonRealtime is the data-structure of this message
type _SysexCommandSysexNonRealtime struct {
	SysexCommandContract
}

var _ SysexCommandSysexNonRealtime = (*_SysexCommandSysexNonRealtime)(nil)
var _ SysexCommandRequirements = (*_SysexCommandSysexNonRealtime)(nil)

// NewSysexCommandSysexNonRealtime factory function for _SysexCommandSysexNonRealtime
func NewSysexCommandSysexNonRealtime() *_SysexCommandSysexNonRealtime {
	_result := &_SysexCommandSysexNonRealtime{
		SysexCommandContract: NewSysexCommand(),
	}
	_result.SysexCommandContract.(*_SysexCommand)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// SysexCommandSysexNonRealtimeBuilder is a builder for SysexCommandSysexNonRealtime
type SysexCommandSysexNonRealtimeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() SysexCommandSysexNonRealtimeBuilder
	// Build builds the SysexCommandSysexNonRealtime or returns an error if something is wrong
	Build() (SysexCommandSysexNonRealtime, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() SysexCommandSysexNonRealtime
}

// NewSysexCommandSysexNonRealtimeBuilder() creates a SysexCommandSysexNonRealtimeBuilder
func NewSysexCommandSysexNonRealtimeBuilder() SysexCommandSysexNonRealtimeBuilder {
	return &_SysexCommandSysexNonRealtimeBuilder{_SysexCommandSysexNonRealtime: new(_SysexCommandSysexNonRealtime)}
}

type _SysexCommandSysexNonRealtimeBuilder struct {
	*_SysexCommandSysexNonRealtime

	parentBuilder *_SysexCommandBuilder

	err *utils.MultiError
}

var _ (SysexCommandSysexNonRealtimeBuilder) = (*_SysexCommandSysexNonRealtimeBuilder)(nil)

func (b *_SysexCommandSysexNonRealtimeBuilder) setParent(contract SysexCommandContract) {
	b.SysexCommandContract = contract
}

func (b *_SysexCommandSysexNonRealtimeBuilder) WithMandatoryFields() SysexCommandSysexNonRealtimeBuilder {
	return b
}

func (b *_SysexCommandSysexNonRealtimeBuilder) Build() (SysexCommandSysexNonRealtime, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._SysexCommandSysexNonRealtime.deepCopy(), nil
}

func (b *_SysexCommandSysexNonRealtimeBuilder) MustBuild() SysexCommandSysexNonRealtime {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_SysexCommandSysexNonRealtimeBuilder) Done() SysexCommandBuilder {
	return b.parentBuilder
}

func (b *_SysexCommandSysexNonRealtimeBuilder) buildForSysexCommand() (SysexCommand, error) {
	return b.Build()
}

func (b *_SysexCommandSysexNonRealtimeBuilder) DeepCopy() any {
	_copy := b.CreateSysexCommandSysexNonRealtimeBuilder().(*_SysexCommandSysexNonRealtimeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateSysexCommandSysexNonRealtimeBuilder creates a SysexCommandSysexNonRealtimeBuilder
func (b *_SysexCommandSysexNonRealtime) CreateSysexCommandSysexNonRealtimeBuilder() SysexCommandSysexNonRealtimeBuilder {
	if b == nil {
		return NewSysexCommandSysexNonRealtimeBuilder()
	}
	return &_SysexCommandSysexNonRealtimeBuilder{_SysexCommandSysexNonRealtime: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SysexCommandSysexNonRealtime) GetCommandType() uint8 {
	return 0x7E
}

func (m *_SysexCommandSysexNonRealtime) GetResponse() bool {
	return false
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SysexCommandSysexNonRealtime) GetParent() SysexCommandContract {
	return m.SysexCommandContract
}

// Deprecated: use the interface for direct cast
func CastSysexCommandSysexNonRealtime(structType any) SysexCommandSysexNonRealtime {
	if casted, ok := structType.(SysexCommandSysexNonRealtime); ok {
		return casted
	}
	if casted, ok := structType.(*SysexCommandSysexNonRealtime); ok {
		return *casted
	}
	return nil
}

func (m *_SysexCommandSysexNonRealtime) GetTypeName() string {
	return "SysexCommandSysexNonRealtime"
}

func (m *_SysexCommandSysexNonRealtime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SysexCommandContract.(*_SysexCommand).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_SysexCommandSysexNonRealtime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SysexCommandSysexNonRealtime) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SysexCommand, response bool) (__sysexCommandSysexNonRealtime SysexCommandSysexNonRealtime, err error) {
	m.SysexCommandContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SysexCommandSysexNonRealtime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SysexCommandSysexNonRealtime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SysexCommandSysexNonRealtime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SysexCommandSysexNonRealtime")
	}

	return m, nil
}

func (m *_SysexCommandSysexNonRealtime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SysexCommandSysexNonRealtime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SysexCommandSysexNonRealtime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SysexCommandSysexNonRealtime")
		}

		if popErr := writeBuffer.PopContext("SysexCommandSysexNonRealtime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SysexCommandSysexNonRealtime")
		}
		return nil
	}
	return m.SysexCommandContract.(*_SysexCommand).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SysexCommandSysexNonRealtime) IsSysexCommandSysexNonRealtime() {}

func (m *_SysexCommandSysexNonRealtime) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SysexCommandSysexNonRealtime) deepCopy() *_SysexCommandSysexNonRealtime {
	if m == nil {
		return nil
	}
	_SysexCommandSysexNonRealtimeCopy := &_SysexCommandSysexNonRealtime{
		m.SysexCommandContract.(*_SysexCommand).deepCopy(),
	}
	m.SysexCommandContract.(*_SysexCommand)._SubType = m
	return _SysexCommandSysexNonRealtimeCopy
}

func (m *_SysexCommandSysexNonRealtime) String() string {
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
