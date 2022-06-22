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

// KnxNetIpCore is the corresponding interface of KnxNetIpCore
type KnxNetIpCore interface {
	utils.LengthAware
	utils.Serializable
	ServiceId
	// GetVersion returns Version (property field)
	GetVersion() uint8
}

// KnxNetIpCoreExactly can be used when we want exactly this type and not a type which fulfills KnxNetIpCore.
// This is useful for switch cases.
type KnxNetIpCoreExactly interface {
	isKnxNetIpCore() bool
}

// _KnxNetIpCore is the data-structure of this message
type _KnxNetIpCore struct {
	*_ServiceId
	Version uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_KnxNetIpCore) GetServiceType() uint8 {
	return 0x02
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_KnxNetIpCore) InitializeParent(parent ServiceId) {}

func (m *_KnxNetIpCore) GetParent() ServiceId {
	return m._ServiceId
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_KnxNetIpCore) GetVersion() uint8 {
	return m.Version
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewKnxNetIpCore factory function for _KnxNetIpCore
func NewKnxNetIpCore(version uint8) *_KnxNetIpCore {
	_result := &_KnxNetIpCore{
		Version:    version,
		_ServiceId: NewServiceId(),
	}
	_result._ServiceId._ServiceIdChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastKnxNetIpCore(structType interface{}) KnxNetIpCore {
	if casted, ok := structType.(KnxNetIpCore); ok {
		return casted
	}
	if casted, ok := structType.(*KnxNetIpCore); ok {
		return *casted
	}
	return nil
}

func (m *_KnxNetIpCore) GetTypeName() string {
	return "KnxNetIpCore"
}

func (m *_KnxNetIpCore) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_KnxNetIpCore) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (version)
	lengthInBits += 8

	return lengthInBits
}

func (m *_KnxNetIpCore) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func KnxNetIpCoreParse(readBuffer utils.ReadBuffer) (KnxNetIpCore, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("KnxNetIpCore"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for KnxNetIpCore")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (version)
	_version, _versionErr := readBuffer.ReadUint8("version", 8)
	if _versionErr != nil {
		return nil, errors.Wrap(_versionErr, "Error parsing 'version' field")
	}
	version := _version

	if closeErr := readBuffer.CloseContext("KnxNetIpCore"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for KnxNetIpCore")
	}

	// Create a partially initialized instance
	_child := &_KnxNetIpCore{
		Version:    version,
		_ServiceId: &_ServiceId{},
	}
	_child._ServiceId._ServiceIdChildRequirements = _child
	return _child, nil
}

func (m *_KnxNetIpCore) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("KnxNetIpCore"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for KnxNetIpCore")
		}

		// Simple Field (version)
		version := uint8(m.GetVersion())
		_versionErr := writeBuffer.WriteUint8("version", 8, (version))
		if _versionErr != nil {
			return errors.Wrap(_versionErr, "Error serializing 'version' field")
		}

		if popErr := writeBuffer.PopContext("KnxNetIpCore"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for KnxNetIpCore")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_KnxNetIpCore) isKnxNetIpCore() bool {
	return true
}

func (m *_KnxNetIpCore) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
