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
	"github.com/rs/zerolog/log"
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConfirmedServiceRequestDeviceCommunicationControl is the corresponding interface of BACnetConfirmedServiceRequestDeviceCommunicationControl
type BACnetConfirmedServiceRequestDeviceCommunicationControl interface {
	utils.LengthAware
	utils.Serializable
	BACnetConfirmedServiceRequest
	// GetTimeDuration returns TimeDuration (property field)
	GetTimeDuration() BACnetContextTagUnsignedInteger
	// GetEnableDisable returns EnableDisable (property field)
	GetEnableDisable() BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged
	// GetPassword returns Password (property field)
	GetPassword() BACnetContextTagCharacterString
}

// BACnetConfirmedServiceRequestDeviceCommunicationControlExactly can be used when we want exactly this type and not a type which fulfills BACnetConfirmedServiceRequestDeviceCommunicationControl.
// This is useful for switch cases.
type BACnetConfirmedServiceRequestDeviceCommunicationControlExactly interface {
	BACnetConfirmedServiceRequestDeviceCommunicationControl
	isBACnetConfirmedServiceRequestDeviceCommunicationControl() bool
}

// _BACnetConfirmedServiceRequestDeviceCommunicationControl is the data-structure of this message
type _BACnetConfirmedServiceRequestDeviceCommunicationControl struct {
	*_BACnetConfirmedServiceRequest
	TimeDuration  BACnetContextTagUnsignedInteger
	EnableDisable BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged
	Password      BACnetContextTagCharacterString
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConfirmedServiceRequestDeviceCommunicationControl) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_DEVICE_COMMUNICATION_CONTROL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestDeviceCommunicationControl) InitializeParent(parent BACnetConfirmedServiceRequest) {
}

func (m *_BACnetConfirmedServiceRequestDeviceCommunicationControl) GetParent() BACnetConfirmedServiceRequest {
	return m._BACnetConfirmedServiceRequest
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestDeviceCommunicationControl) GetTimeDuration() BACnetContextTagUnsignedInteger {
	return m.TimeDuration
}

func (m *_BACnetConfirmedServiceRequestDeviceCommunicationControl) GetEnableDisable() BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged {
	return m.EnableDisable
}

func (m *_BACnetConfirmedServiceRequestDeviceCommunicationControl) GetPassword() BACnetContextTagCharacterString {
	return m.Password
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestDeviceCommunicationControl factory function for _BACnetConfirmedServiceRequestDeviceCommunicationControl
func NewBACnetConfirmedServiceRequestDeviceCommunicationControl(timeDuration BACnetContextTagUnsignedInteger, enableDisable BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged, password BACnetContextTagCharacterString, serviceRequestLength uint16) *_BACnetConfirmedServiceRequestDeviceCommunicationControl {
	_result := &_BACnetConfirmedServiceRequestDeviceCommunicationControl{
		TimeDuration:                   timeDuration,
		EnableDisable:                  enableDisable,
		Password:                       password,
		_BACnetConfirmedServiceRequest: NewBACnetConfirmedServiceRequest(serviceRequestLength),
	}
	_result._BACnetConfirmedServiceRequest._BACnetConfirmedServiceRequestChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestDeviceCommunicationControl(structType interface{}) BACnetConfirmedServiceRequestDeviceCommunicationControl {
	if casted, ok := structType.(BACnetConfirmedServiceRequestDeviceCommunicationControl); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestDeviceCommunicationControl); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestDeviceCommunicationControl) GetTypeName() string {
	return "BACnetConfirmedServiceRequestDeviceCommunicationControl"
}

func (m *_BACnetConfirmedServiceRequestDeviceCommunicationControl) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConfirmedServiceRequestDeviceCommunicationControl) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Optional Field (timeDuration)
	if m.TimeDuration != nil {
		lengthInBits += m.TimeDuration.GetLengthInBits()
	}

	// Simple field (enableDisable)
	lengthInBits += m.EnableDisable.GetLengthInBits()

	// Optional Field (password)
	if m.Password != nil {
		lengthInBits += m.Password.GetLengthInBits()
	}

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestDeviceCommunicationControl) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConfirmedServiceRequestDeviceCommunicationControlParse(readBuffer utils.ReadBuffer, serviceRequestLength uint16) (BACnetConfirmedServiceRequestDeviceCommunicationControl, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestDeviceCommunicationControl"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestDeviceCommunicationControl")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Optional Field (timeDuration) (Can be skipped, if a given expression evaluates to false)
	var timeDuration BACnetContextTagUnsignedInteger = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("timeDuration"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for timeDuration")
		}
		_val, _err := BACnetContextTagParse(readBuffer, uint8(0), BACnetDataType_UNSIGNED_INTEGER)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'timeDuration' field")
		default:
			timeDuration = _val.(BACnetContextTagUnsignedInteger)
			if closeErr := readBuffer.CloseContext("timeDuration"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for timeDuration")
			}
		}
	}

	// Simple Field (enableDisable)
	if pullErr := readBuffer.PullContext("enableDisable"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for enableDisable")
	}
	_enableDisable, _enableDisableErr := BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTaggedParse(readBuffer, uint8(uint8(1)), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _enableDisableErr != nil {
		return nil, errors.Wrap(_enableDisableErr, "Error parsing 'enableDisable' field")
	}
	enableDisable := _enableDisable.(BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged)
	if closeErr := readBuffer.CloseContext("enableDisable"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for enableDisable")
	}

	// Optional Field (password) (Can be skipped, if a given expression evaluates to false)
	var password BACnetContextTagCharacterString = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("password"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for password")
		}
		_val, _err := BACnetContextTagParse(readBuffer, uint8(2), BACnetDataType_CHARACTER_STRING)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'password' field")
		default:
			password = _val.(BACnetContextTagCharacterString)
			if closeErr := readBuffer.CloseContext("password"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for password")
			}
		}
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestDeviceCommunicationControl"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestDeviceCommunicationControl")
	}

	// Create a partially initialized instance
	_child := &_BACnetConfirmedServiceRequestDeviceCommunicationControl{
		TimeDuration:  timeDuration,
		EnableDisable: enableDisable,
		Password:      password,
		_BACnetConfirmedServiceRequest: &_BACnetConfirmedServiceRequest{
			ServiceRequestLength: serviceRequestLength,
		},
	}
	_child._BACnetConfirmedServiceRequest._BACnetConfirmedServiceRequestChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConfirmedServiceRequestDeviceCommunicationControl) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestDeviceCommunicationControl"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestDeviceCommunicationControl")
		}

		// Optional Field (timeDuration) (Can be skipped, if the value is null)
		var timeDuration BACnetContextTagUnsignedInteger = nil
		if m.GetTimeDuration() != nil {
			if pushErr := writeBuffer.PushContext("timeDuration"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for timeDuration")
			}
			timeDuration = m.GetTimeDuration()
			_timeDurationErr := writeBuffer.WriteSerializable(timeDuration)
			if popErr := writeBuffer.PopContext("timeDuration"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for timeDuration")
			}
			if _timeDurationErr != nil {
				return errors.Wrap(_timeDurationErr, "Error serializing 'timeDuration' field")
			}
		}

		// Simple Field (enableDisable)
		if pushErr := writeBuffer.PushContext("enableDisable"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for enableDisable")
		}
		_enableDisableErr := writeBuffer.WriteSerializable(m.GetEnableDisable())
		if popErr := writeBuffer.PopContext("enableDisable"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for enableDisable")
		}
		if _enableDisableErr != nil {
			return errors.Wrap(_enableDisableErr, "Error serializing 'enableDisable' field")
		}

		// Optional Field (password) (Can be skipped, if the value is null)
		var password BACnetContextTagCharacterString = nil
		if m.GetPassword() != nil {
			if pushErr := writeBuffer.PushContext("password"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for password")
			}
			password = m.GetPassword()
			_passwordErr := writeBuffer.WriteSerializable(password)
			if popErr := writeBuffer.PopContext("password"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for password")
			}
			if _passwordErr != nil {
				return errors.Wrap(_passwordErr, "Error serializing 'password' field")
			}
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestDeviceCommunicationControl"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestDeviceCommunicationControl")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConfirmedServiceRequestDeviceCommunicationControl) isBACnetConfirmedServiceRequestDeviceCommunicationControl() bool {
	return true
}

func (m *_BACnetConfirmedServiceRequestDeviceCommunicationControl) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
