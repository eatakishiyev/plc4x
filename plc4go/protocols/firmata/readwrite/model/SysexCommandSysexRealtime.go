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

// SysexCommandSysexRealtime is the corresponding interface of SysexCommandSysexRealtime
type SysexCommandSysexRealtime interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	SysexCommand
	// IsSysexCommandSysexRealtime is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSysexCommandSysexRealtime()
	// CreateBuilder creates a SysexCommandSysexRealtimeBuilder
	CreateSysexCommandSysexRealtimeBuilder() SysexCommandSysexRealtimeBuilder
}

// _SysexCommandSysexRealtime is the data-structure of this message
type _SysexCommandSysexRealtime struct {
	SysexCommandContract
}

var _ SysexCommandSysexRealtime = (*_SysexCommandSysexRealtime)(nil)
var _ SysexCommandRequirements = (*_SysexCommandSysexRealtime)(nil)

// NewSysexCommandSysexRealtime factory function for _SysexCommandSysexRealtime
func NewSysexCommandSysexRealtime() *_SysexCommandSysexRealtime {
	_result := &_SysexCommandSysexRealtime{
		SysexCommandContract: NewSysexCommand(),
	}
	_result.SysexCommandContract.(*_SysexCommand)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// SysexCommandSysexRealtimeBuilder is a builder for SysexCommandSysexRealtime
type SysexCommandSysexRealtimeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() SysexCommandSysexRealtimeBuilder
	// Build builds the SysexCommandSysexRealtime or returns an error if something is wrong
	Build() (SysexCommandSysexRealtime, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() SysexCommandSysexRealtime
}

// NewSysexCommandSysexRealtimeBuilder() creates a SysexCommandSysexRealtimeBuilder
func NewSysexCommandSysexRealtimeBuilder() SysexCommandSysexRealtimeBuilder {
	return &_SysexCommandSysexRealtimeBuilder{_SysexCommandSysexRealtime: new(_SysexCommandSysexRealtime)}
}

type _SysexCommandSysexRealtimeBuilder struct {
	*_SysexCommandSysexRealtime

	parentBuilder *_SysexCommandBuilder

	err *utils.MultiError
}

var _ (SysexCommandSysexRealtimeBuilder) = (*_SysexCommandSysexRealtimeBuilder)(nil)

func (b *_SysexCommandSysexRealtimeBuilder) setParent(contract SysexCommandContract) {
	b.SysexCommandContract = contract
}

func (b *_SysexCommandSysexRealtimeBuilder) WithMandatoryFields() SysexCommandSysexRealtimeBuilder {
	return b
}

func (b *_SysexCommandSysexRealtimeBuilder) Build() (SysexCommandSysexRealtime, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._SysexCommandSysexRealtime.deepCopy(), nil
}

func (b *_SysexCommandSysexRealtimeBuilder) MustBuild() SysexCommandSysexRealtime {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_SysexCommandSysexRealtimeBuilder) Done() SysexCommandBuilder {
	return b.parentBuilder
}

func (b *_SysexCommandSysexRealtimeBuilder) buildForSysexCommand() (SysexCommand, error) {
	return b.Build()
}

func (b *_SysexCommandSysexRealtimeBuilder) DeepCopy() any {
	_copy := b.CreateSysexCommandSysexRealtimeBuilder().(*_SysexCommandSysexRealtimeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateSysexCommandSysexRealtimeBuilder creates a SysexCommandSysexRealtimeBuilder
func (b *_SysexCommandSysexRealtime) CreateSysexCommandSysexRealtimeBuilder() SysexCommandSysexRealtimeBuilder {
	if b == nil {
		return NewSysexCommandSysexRealtimeBuilder()
	}
	return &_SysexCommandSysexRealtimeBuilder{_SysexCommandSysexRealtime: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SysexCommandSysexRealtime) GetCommandType() uint8 {
	return 0x7F
}

func (m *_SysexCommandSysexRealtime) GetResponse() bool {
	return false
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SysexCommandSysexRealtime) GetParent() SysexCommandContract {
	return m.SysexCommandContract
}

// Deprecated: use the interface for direct cast
func CastSysexCommandSysexRealtime(structType any) SysexCommandSysexRealtime {
	if casted, ok := structType.(SysexCommandSysexRealtime); ok {
		return casted
	}
	if casted, ok := structType.(*SysexCommandSysexRealtime); ok {
		return *casted
	}
	return nil
}

func (m *_SysexCommandSysexRealtime) GetTypeName() string {
	return "SysexCommandSysexRealtime"
}

func (m *_SysexCommandSysexRealtime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SysexCommandContract.(*_SysexCommand).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_SysexCommandSysexRealtime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SysexCommandSysexRealtime) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SysexCommand, response bool) (__sysexCommandSysexRealtime SysexCommandSysexRealtime, err error) {
	m.SysexCommandContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SysexCommandSysexRealtime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SysexCommandSysexRealtime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SysexCommandSysexRealtime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SysexCommandSysexRealtime")
	}

	return m, nil
}

func (m *_SysexCommandSysexRealtime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SysexCommandSysexRealtime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SysexCommandSysexRealtime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SysexCommandSysexRealtime")
		}

		if popErr := writeBuffer.PopContext("SysexCommandSysexRealtime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SysexCommandSysexRealtime")
		}
		return nil
	}
	return m.SysexCommandContract.(*_SysexCommand).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SysexCommandSysexRealtime) IsSysexCommandSysexRealtime() {}

func (m *_SysexCommandSysexRealtime) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SysexCommandSysexRealtime) deepCopy() *_SysexCommandSysexRealtime {
	if m == nil {
		return nil
	}
	_SysexCommandSysexRealtimeCopy := &_SysexCommandSysexRealtime{
		m.SysexCommandContract.(*_SysexCommand).deepCopy(),
	}
	m.SysexCommandContract.(*_SysexCommand)._SubType = m
	return _SysexCommandSysexRealtimeCopy
}

func (m *_SysexCommandSysexRealtime) String() string {
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
