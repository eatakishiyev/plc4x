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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// KnxNetIpRouting is the corresponding interface of KnxNetIpRouting
type KnxNetIpRouting interface {
	utils.LengthAware
	utils.Serializable
	ServiceId
	// GetVersion returns Version (property field)
	GetVersion() uint8
}

// KnxNetIpRoutingExactly can be used when we want exactly this type and not a type which fulfills KnxNetIpRouting.
// This is useful for switch cases.
type KnxNetIpRoutingExactly interface {
	isKnxNetIpRouting() bool
}

// _KnxNetIpRouting is the data-structure of this message
type _KnxNetIpRouting struct {
	*_ServiceId
	Version uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_KnxNetIpRouting) GetServiceType() uint8 {
	return 0x05
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_KnxNetIpRouting) InitializeParent(parent ServiceId) {}

func (m *_KnxNetIpRouting) GetParent() ServiceId {
	return m._ServiceId
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_KnxNetIpRouting) GetVersion() uint8 {
	return m.Version
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewKnxNetIpRouting factory function for _KnxNetIpRouting
func NewKnxNetIpRouting(version uint8) *_KnxNetIpRouting {
	_result := &_KnxNetIpRouting{
		Version:    version,
		_ServiceId: NewServiceId(),
	}
	_result._ServiceId._ServiceIdChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastKnxNetIpRouting(structType interface{}) KnxNetIpRouting {
	if casted, ok := structType.(KnxNetIpRouting); ok {
		return casted
	}
	if casted, ok := structType.(*KnxNetIpRouting); ok {
		return *casted
	}
	return nil
}

func (m *_KnxNetIpRouting) GetTypeName() string {
	return "KnxNetIpRouting"
}

func (m *_KnxNetIpRouting) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_KnxNetIpRouting) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (version)
	lengthInBits += 8

	return lengthInBits
}

func (m *_KnxNetIpRouting) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func KnxNetIpRoutingParse(readBuffer utils.ReadBuffer) (KnxNetIpRouting, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("KnxNetIpRouting"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for KnxNetIpRouting")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (version)
	_version, _versionErr := readBuffer.ReadUint8("version", 8)
	if _versionErr != nil {
		return nil, errors.Wrap(_versionErr, "Error parsing 'version' field")
	}
	version := _version

	if closeErr := readBuffer.CloseContext("KnxNetIpRouting"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for KnxNetIpRouting")
	}

	// Create a partially initialized instance
	_child := &_KnxNetIpRouting{
		Version:    version,
		_ServiceId: &_ServiceId{},
	}
	_child._ServiceId._ServiceIdChildRequirements = _child
	return _child, nil
}

func (m *_KnxNetIpRouting) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("KnxNetIpRouting"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for KnxNetIpRouting")
		}

		// Simple Field (version)
		version := uint8(m.GetVersion())
		_versionErr := writeBuffer.WriteUint8("version", 8, (version))
		if _versionErr != nil {
			return errors.Wrap(_versionErr, "Error serializing 'version' field")
		}

		if popErr := writeBuffer.PopContext("KnxNetIpRouting"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for KnxNetIpRouting")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_KnxNetIpRouting) isKnxNetIpRouting() bool {
	return true
}

func (m *_KnxNetIpRouting) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
