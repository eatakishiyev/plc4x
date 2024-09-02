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

// SecurityData is the corresponding interface of SecurityData
type SecurityData interface {
	SecurityDataContract
	SecurityDataRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// SecurityDataContract provides a set of functions which can be overwritten by a sub struct
type SecurityDataContract interface {
	// GetCommandTypeContainer returns CommandTypeContainer (property field)
	GetCommandTypeContainer() SecurityCommandTypeContainer
	// GetArgument returns Argument (property field)
	GetArgument() byte
	// GetCommandType returns CommandType (virtual field)
	GetCommandType() SecurityCommandType
}

// SecurityDataRequirements provides a set of functions which need to be implemented by a sub struct
type SecurityDataRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetArgument returns Argument (discriminator field)
	GetArgument() byte
	// GetCommandType returns CommandType (discriminator field)
	GetCommandType() SecurityCommandType
}

// SecurityDataExactly can be used when we want exactly this type and not a type which fulfills SecurityData.
// This is useful for switch cases.
type SecurityDataExactly interface {
	SecurityData
	isSecurityData() bool
}

// _SecurityData is the data-structure of this message
type _SecurityData struct {
	_SubType             SecurityData
	CommandTypeContainer SecurityCommandTypeContainer
	Argument             byte
}

var _ SecurityDataContract = (*_SecurityData)(nil)

type SecurityDataChild interface {
	utils.Serializable

	GetParent() *SecurityData

	GetTypeName() string
	SecurityData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SecurityData) GetCommandTypeContainer() SecurityCommandTypeContainer {
	return m.CommandTypeContainer
}

func (m *_SecurityData) GetArgument() byte {
	return m.Argument
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (pm *_SecurityData) GetCommandType() SecurityCommandType {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	return CastSecurityCommandType(m.GetCommandTypeContainer().CommandType())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSecurityData factory function for _SecurityData
func NewSecurityData(commandTypeContainer SecurityCommandTypeContainer, argument byte) *_SecurityData {
	return &_SecurityData{CommandTypeContainer: commandTypeContainer, Argument: argument}
}

// Deprecated: use the interface for direct cast
func CastSecurityData(structType any) SecurityData {
	if casted, ok := structType.(SecurityData); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityData); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityData) GetTypeName() string {
	return "SecurityData"
}

func (m *_SecurityData) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (commandTypeContainer)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// Simple field (argument)
	lengthInBits += 8

	return lengthInBits
}

func (m *_SecurityData) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func SecurityDataParse[T SecurityData](ctx context.Context, theBytes []byte) (T, error) {
	return SecurityDataParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes))
}

func SecurityDataParseWithBufferProducer[T SecurityData]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := SecurityDataParseWithBuffer[T](ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, err
	}
}

func SecurityDataParseWithBuffer[T SecurityData](ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	v, err := NewSecurityData().parse(ctx, readBuffer)
	if err != nil {
		var zero T
		return zero, err
	}
	return v.(T), err
}

func (m *_SecurityData) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__securityData SecurityData, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityData")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(KnowsSecurityCommandTypeContainer(ctx, readBuffer)) {
		return nil, errors.WithStack(utils.ParseAssertError{Message: "no command type could be found"})
	}

	commandTypeContainer, err := ReadEnumField[SecurityCommandTypeContainer](ctx, "commandTypeContainer", "SecurityCommandTypeContainer", ReadEnum(SecurityCommandTypeContainerByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'commandTypeContainer' field"))
	}
	m.CommandTypeContainer = commandTypeContainer

	commandType, err := ReadVirtualField[SecurityCommandType](ctx, "commandType", (*SecurityCommandType)(nil), commandTypeContainer.CommandType())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'commandType' field"))
	}
	_ = commandType

	argument, err := ReadSimpleField(ctx, "argument", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'argument' field"))
	}
	m.Argument = argument

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child SecurityData
	switch {
	case commandType == SecurityCommandType_ON && argument == 0x80: // SecurityDataSystemArmedDisarmed
		if _child, err = (&_SecurityDataSystemArmedDisarmed{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SecurityDataSystemArmedDisarmed for type-switch of SecurityData")
		}
	case commandType == SecurityCommandType_OFF && argument == 0x80: // SecurityDataSystemDisarmed
		if _child, err = (&_SecurityDataSystemDisarmed{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SecurityDataSystemDisarmed for type-switch of SecurityData")
		}
	case commandType == SecurityCommandType_EVENT && argument == 0x81: // SecurityDataExitDelayStarted
		if _child, err = (&_SecurityDataExitDelayStarted{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SecurityDataExitDelayStarted for type-switch of SecurityData")
		}
	case commandType == SecurityCommandType_EVENT && argument == 0x82: // SecurityDataEntryDelayStarted
		if _child, err = (&_SecurityDataEntryDelayStarted{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SecurityDataEntryDelayStarted for type-switch of SecurityData")
		}
	case commandType == SecurityCommandType_ON && argument == 0x83: // SecurityDataAlarmOn
		if _child, err = (&_SecurityDataAlarmOn{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SecurityDataAlarmOn for type-switch of SecurityData")
		}
	case commandType == SecurityCommandType_OFF && argument == 0x83: // SecurityDataAlarmOff
		if _child, err = (&_SecurityDataAlarmOff{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SecurityDataAlarmOff for type-switch of SecurityData")
		}
	case commandType == SecurityCommandType_ON && argument == 0x84: // SecurityDataTamperOn
		if _child, err = (&_SecurityDataTamperOn{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SecurityDataTamperOn for type-switch of SecurityData")
		}
	case commandType == SecurityCommandType_OFF && argument == 0x84: // SecurityDataTamperOff
		if _child, err = (&_SecurityDataTamperOff{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SecurityDataTamperOff for type-switch of SecurityData")
		}
	case commandType == SecurityCommandType_ON && argument == 0x85: // SecurityDataPanicActivated
		if _child, err = (&_SecurityDataPanicActivated{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SecurityDataPanicActivated for type-switch of SecurityData")
		}
	case commandType == SecurityCommandType_OFF && argument == 0x85: // SecurityDataPanicCleared
		if _child, err = (&_SecurityDataPanicCleared{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SecurityDataPanicCleared for type-switch of SecurityData")
		}
	case commandType == SecurityCommandType_EVENT && argument == 0x86: // SecurityDataZoneUnsealed
		if _child, err = (&_SecurityDataZoneUnsealed{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SecurityDataZoneUnsealed for type-switch of SecurityData")
		}
	case commandType == SecurityCommandType_EVENT && argument == 0x87: // SecurityDataZoneSealed
		if _child, err = (&_SecurityDataZoneSealed{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SecurityDataZoneSealed for type-switch of SecurityData")
		}
	case commandType == SecurityCommandType_EVENT && argument == 0x88: // SecurityDataZoneOpen
		if _child, err = (&_SecurityDataZoneOpen{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SecurityDataZoneOpen for type-switch of SecurityData")
		}
	case commandType == SecurityCommandType_EVENT && argument == 0x89: // SecurityDataZoneShort
		if _child, err = (&_SecurityDataZoneShort{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SecurityDataZoneShort for type-switch of SecurityData")
		}
	case commandType == SecurityCommandType_EVENT && argument == 0x89: // SecurityDataZoneIsolated
		if _child, err = (&_SecurityDataZoneIsolated{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SecurityDataZoneIsolated for type-switch of SecurityData")
		}
	case commandType == SecurityCommandType_ON && argument == 0x8B: // SecurityDataLowBatteryDetected
		if _child, err = (&_SecurityDataLowBatteryDetected{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SecurityDataLowBatteryDetected for type-switch of SecurityData")
		}
	case commandType == SecurityCommandType_OFF && argument == 0x8B: // SecurityDataLowBatteryCorrected
		if _child, err = (&_SecurityDataLowBatteryCorrected{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SecurityDataLowBatteryCorrected for type-switch of SecurityData")
		}
	case commandType == SecurityCommandType_EVENT && argument == 0x8C: // SecurityDataLowBatteryCharging
		if _child, err = (&_SecurityDataLowBatteryCharging{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SecurityDataLowBatteryCharging for type-switch of SecurityData")
		}
	case commandType == SecurityCommandType_EVENT && argument == 0x8D: // SecurityDataZoneName
		if _child, err = (&_SecurityDataZoneName{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SecurityDataZoneName for type-switch of SecurityData")
		}
	case commandType == SecurityCommandType_EVENT && argument == 0x8E: // SecurityDataStatusReport1
		if _child, err = (&_SecurityDataStatusReport1{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SecurityDataStatusReport1 for type-switch of SecurityData")
		}
	case commandType == SecurityCommandType_EVENT && argument == 0x8F: // SecurityDataStatusReport2
		if _child, err = (&_SecurityDataStatusReport2{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SecurityDataStatusReport2 for type-switch of SecurityData")
		}
	case commandType == SecurityCommandType_EVENT && argument == 0x90: // SecurityDataPasswordEntryStatus
		if _child, err = (&_SecurityDataPasswordEntryStatus{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SecurityDataPasswordEntryStatus for type-switch of SecurityData")
		}
	case commandType == SecurityCommandType_ON && argument == 0x91: // SecurityDataMainsFailure
		if _child, err = (&_SecurityDataMainsFailure{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SecurityDataMainsFailure for type-switch of SecurityData")
		}
	case commandType == SecurityCommandType_OFF && argument == 0x91: // SecurityDataMainsRestoredOrApplied
		if _child, err = (&_SecurityDataMainsRestoredOrApplied{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SecurityDataMainsRestoredOrApplied for type-switch of SecurityData")
		}
	case commandType == SecurityCommandType_EVENT && argument == 0x92: // SecurityDataArmReadyNotReady
		if _child, err = (&_SecurityDataArmReadyNotReady{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SecurityDataArmReadyNotReady for type-switch of SecurityData")
		}
	case commandType == SecurityCommandType_EVENT && argument == 0x93: // SecurityDataCurrentAlarmType
		if _child, err = (&_SecurityDataCurrentAlarmType{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SecurityDataCurrentAlarmType for type-switch of SecurityData")
		}
	case commandType == SecurityCommandType_ON && argument == 0x94: // SecurityDataLineCutAlarmRaised
		if _child, err = (&_SecurityDataLineCutAlarmRaised{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SecurityDataLineCutAlarmRaised for type-switch of SecurityData")
		}
	case commandType == SecurityCommandType_OFF && argument == 0x94: // SecurityDataLineCutAlarmCleared
		if _child, err = (&_SecurityDataLineCutAlarmCleared{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SecurityDataLineCutAlarmCleared for type-switch of SecurityData")
		}
	case commandType == SecurityCommandType_ON && argument == 0x95: // SecurityDataArmFailedRaised
		if _child, err = (&_SecurityDataArmFailedRaised{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SecurityDataArmFailedRaised for type-switch of SecurityData")
		}
	case commandType == SecurityCommandType_OFF && argument == 0x95: // SecurityDataArmFailedCleared
		if _child, err = (&_SecurityDataArmFailedCleared{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SecurityDataArmFailedCleared for type-switch of SecurityData")
		}
	case commandType == SecurityCommandType_ON && argument == 0x96: // SecurityDataFireAlarmRaised
		if _child, err = (&_SecurityDataFireAlarmRaised{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SecurityDataFireAlarmRaised for type-switch of SecurityData")
		}
	case commandType == SecurityCommandType_OFF && argument == 0x96: // SecurityDataFireAlarmCleared
		if _child, err = (&_SecurityDataFireAlarmCleared{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SecurityDataFireAlarmCleared for type-switch of SecurityData")
		}
	case commandType == SecurityCommandType_ON && argument == 0x97: // SecurityDataGasAlarmRaised
		if _child, err = (&_SecurityDataGasAlarmRaised{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SecurityDataGasAlarmRaised for type-switch of SecurityData")
		}
	case commandType == SecurityCommandType_OFF && argument == 0x97: // SecurityDataGasAlarmCleared
		if _child, err = (&_SecurityDataGasAlarmCleared{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SecurityDataGasAlarmCleared for type-switch of SecurityData")
		}
	case commandType == SecurityCommandType_ON && argument == 0x98: // SecurityDataOtherAlarmRaised
		if _child, err = (&_SecurityDataOtherAlarmRaised{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SecurityDataOtherAlarmRaised for type-switch of SecurityData")
		}
	case commandType == SecurityCommandType_OFF && argument == 0x98: // SecurityDataOtherAlarmCleared
		if _child, err = (&_SecurityDataOtherAlarmCleared{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SecurityDataOtherAlarmCleared for type-switch of SecurityData")
		}
	case commandType == SecurityCommandType_EVENT && argument == 0xA0: // SecurityDataStatus1Request
		if _child, err = (&_SecurityDataStatus1Request{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SecurityDataStatus1Request for type-switch of SecurityData")
		}
	case commandType == SecurityCommandType_EVENT && argument == 0xA1: // SecurityDataStatus2Request
		if _child, err = (&_SecurityDataStatus2Request{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SecurityDataStatus2Request for type-switch of SecurityData")
		}
	case commandType == SecurityCommandType_EVENT && argument == 0xA2: // SecurityDataArmSystem
		if _child, err = (&_SecurityDataArmSystem{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SecurityDataArmSystem for type-switch of SecurityData")
		}
	case commandType == SecurityCommandType_ON && argument == 0xA3: // SecurityDataRaiseTamper
		if _child, err = (&_SecurityDataRaiseTamper{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SecurityDataRaiseTamper for type-switch of SecurityData")
		}
	case commandType == SecurityCommandType_OFF && argument == 0xA3: // SecurityDataDropTamper
		if _child, err = (&_SecurityDataDropTamper{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SecurityDataDropTamper for type-switch of SecurityData")
		}
	case commandType == SecurityCommandType_ON && argument == 0xA4: // SecurityDataRaiseAlarm
		if _child, err = (&_SecurityDataRaiseAlarm{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SecurityDataRaiseAlarm for type-switch of SecurityData")
		}
	case commandType == SecurityCommandType_EVENT && argument == 0xA5: // SecurityDataEmulatedKeypad
		if _child, err = (&_SecurityDataEmulatedKeypad{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SecurityDataEmulatedKeypad for type-switch of SecurityData")
		}
	case commandType == SecurityCommandType_ON && argument == 0xA6: // SecurityDataDisplayMessage
		if _child, err = (&_SecurityDataDisplayMessage{}).parse(ctx, readBuffer, m, commandTypeContainer); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SecurityDataDisplayMessage for type-switch of SecurityData")
		}
	case commandType == SecurityCommandType_EVENT && argument == 0xA7: // SecurityDataRequestZoneName
		if _child, err = (&_SecurityDataRequestZoneName{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SecurityDataRequestZoneName for type-switch of SecurityData")
		}
	case commandType == SecurityCommandType_OFF: // SecurityDataOff
		if _child, err = (&_SecurityDataOff{}).parse(ctx, readBuffer, m, commandTypeContainer); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SecurityDataOff for type-switch of SecurityData")
		}
	case commandType == SecurityCommandType_ON: // SecurityDataOn
		if _child, err = (&_SecurityDataOn{}).parse(ctx, readBuffer, m, commandTypeContainer); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SecurityDataOn for type-switch of SecurityData")
		}
	case commandType == SecurityCommandType_EVENT: // SecurityDataEvent
		if _child, err = (&_SecurityDataEvent{}).parse(ctx, readBuffer, m, commandTypeContainer); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type SecurityDataEvent for type-switch of SecurityData")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [commandType=%v, argument=%v]", commandType, argument)
	}

	if closeErr := readBuffer.CloseContext("SecurityData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityData")
	}

	return _child, nil
}

func (pm *_SecurityData) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child SecurityData, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("SecurityData"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for SecurityData")
	}

	if err := WriteSimpleEnumField[SecurityCommandTypeContainer](ctx, "commandTypeContainer", "SecurityCommandTypeContainer", m.GetCommandTypeContainer(), WriteEnum[SecurityCommandTypeContainer, uint8](SecurityCommandTypeContainer.GetValue, SecurityCommandTypeContainer.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
		return errors.Wrap(err, "Error serializing 'commandTypeContainer' field")
	}
	// Virtual field
	commandType := m.GetCommandType()
	_ = commandType
	if _commandTypeErr := writeBuffer.WriteVirtual(ctx, "commandType", m.GetCommandType()); _commandTypeErr != nil {
		return errors.Wrap(_commandTypeErr, "Error serializing 'commandType' field")
	}

	if err := WriteSimpleField[byte](ctx, "argument", m.GetArgument(), WriteByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'argument' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("SecurityData"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for SecurityData")
	}
	return nil
}

func (m *_SecurityData) isSecurityData() bool {
	return true
}
