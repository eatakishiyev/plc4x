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

// BVLCBroadcastDistributionTableEntry is the corresponding interface of BVLCBroadcastDistributionTableEntry
type BVLCBroadcastDistributionTableEntry interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetIp returns Ip (property field)
	GetIp() []uint8
	// GetPort returns Port (property field)
	GetPort() uint16
	// GetBroadcastDistributionMap returns BroadcastDistributionMap (property field)
	GetBroadcastDistributionMap() []uint8
}

// BVLCBroadcastDistributionTableEntryExactly can be used when we want exactly this type and not a type which fulfills BVLCBroadcastDistributionTableEntry.
// This is useful for switch cases.
type BVLCBroadcastDistributionTableEntryExactly interface {
	BVLCBroadcastDistributionTableEntry
	isBVLCBroadcastDistributionTableEntry() bool
}

// _BVLCBroadcastDistributionTableEntry is the data-structure of this message
type _BVLCBroadcastDistributionTableEntry struct {
	Ip                       []uint8
	Port                     uint16
	BroadcastDistributionMap []uint8
}

var _ BVLCBroadcastDistributionTableEntry = (*_BVLCBroadcastDistributionTableEntry)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BVLCBroadcastDistributionTableEntry) GetIp() []uint8 {
	return m.Ip
}

func (m *_BVLCBroadcastDistributionTableEntry) GetPort() uint16 {
	return m.Port
}

func (m *_BVLCBroadcastDistributionTableEntry) GetBroadcastDistributionMap() []uint8 {
	return m.BroadcastDistributionMap
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBVLCBroadcastDistributionTableEntry factory function for _BVLCBroadcastDistributionTableEntry
func NewBVLCBroadcastDistributionTableEntry(ip []uint8, port uint16, broadcastDistributionMap []uint8) *_BVLCBroadcastDistributionTableEntry {
	return &_BVLCBroadcastDistributionTableEntry{Ip: ip, Port: port, BroadcastDistributionMap: broadcastDistributionMap}
}

// Deprecated: use the interface for direct cast
func CastBVLCBroadcastDistributionTableEntry(structType any) BVLCBroadcastDistributionTableEntry {
	if casted, ok := structType.(BVLCBroadcastDistributionTableEntry); ok {
		return casted
	}
	if casted, ok := structType.(*BVLCBroadcastDistributionTableEntry); ok {
		return *casted
	}
	return nil
}

func (m *_BVLCBroadcastDistributionTableEntry) GetTypeName() string {
	return "BVLCBroadcastDistributionTableEntry"
}

func (m *_BVLCBroadcastDistributionTableEntry) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Array field
	if len(m.Ip) > 0 {
		lengthInBits += 8 * uint16(len(m.Ip))
	}

	// Simple field (port)
	lengthInBits += 16

	// Array field
	if len(m.BroadcastDistributionMap) > 0 {
		lengthInBits += 8 * uint16(len(m.BroadcastDistributionMap))
	}

	return lengthInBits
}

func (m *_BVLCBroadcastDistributionTableEntry) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BVLCBroadcastDistributionTableEntryParse(ctx context.Context, theBytes []byte) (BVLCBroadcastDistributionTableEntry, error) {
	return BVLCBroadcastDistributionTableEntryParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BVLCBroadcastDistributionTableEntryParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BVLCBroadcastDistributionTableEntry, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BVLCBroadcastDistributionTableEntry, error) {
		return BVLCBroadcastDistributionTableEntryParseWithBuffer(ctx, readBuffer)
	}
}

func BVLCBroadcastDistributionTableEntryParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BVLCBroadcastDistributionTableEntry, error) {
	v, err := NewBVLCBroadcastDistributionTableEntry().parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_BVLCBroadcastDistributionTableEntry) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bVLCBroadcastDistributionTableEntry BVLCBroadcastDistributionTableEntry, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BVLCBroadcastDistributionTableEntry"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BVLCBroadcastDistributionTableEntry")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	ip, err := ReadCountArrayField[uint8](ctx, "ip", ReadUnsignedByte(readBuffer, uint8(8)), uint64(int32(4)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'ip' field"))
	}
	m.Ip = ip

	port, err := ReadSimpleField(ctx, "port", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'port' field"))
	}
	m.Port = port

	broadcastDistributionMap, err := ReadCountArrayField[uint8](ctx, "broadcastDistributionMap", ReadUnsignedByte(readBuffer, uint8(8)), uint64(int32(4)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'broadcastDistributionMap' field"))
	}
	m.BroadcastDistributionMap = broadcastDistributionMap

	if closeErr := readBuffer.CloseContext("BVLCBroadcastDistributionTableEntry"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BVLCBroadcastDistributionTableEntry")
	}

	return m, nil
}

func (m *_BVLCBroadcastDistributionTableEntry) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BVLCBroadcastDistributionTableEntry) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BVLCBroadcastDistributionTableEntry"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BVLCBroadcastDistributionTableEntry")
	}

	if err := WriteSimpleTypeArrayField(ctx, "ip", m.GetIp(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'ip' field")
	}

	if err := WriteSimpleField[uint16](ctx, "port", m.GetPort(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'port' field")
	}

	if err := WriteSimpleTypeArrayField(ctx, "broadcastDistributionMap", m.GetBroadcastDistributionMap(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'broadcastDistributionMap' field")
	}

	if popErr := writeBuffer.PopContext("BVLCBroadcastDistributionTableEntry"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BVLCBroadcastDistributionTableEntry")
	}
	return nil
}

func (m *_BVLCBroadcastDistributionTableEntry) isBVLCBroadcastDistributionTableEntry() bool {
	return true
}

func (m *_BVLCBroadcastDistributionTableEntry) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
