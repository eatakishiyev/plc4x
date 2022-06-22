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

// BACnetConfirmedServiceRequestSubscribeCOVProperty is the corresponding interface of BACnetConfirmedServiceRequestSubscribeCOVProperty
type BACnetConfirmedServiceRequestSubscribeCOVProperty interface {
	utils.LengthAware
	utils.Serializable
	BACnetConfirmedServiceRequest
	// GetSubscriberProcessIdentifier returns SubscriberProcessIdentifier (property field)
	GetSubscriberProcessIdentifier() BACnetContextTagUnsignedInteger
	// GetMonitoredObjectIdentifier returns MonitoredObjectIdentifier (property field)
	GetMonitoredObjectIdentifier() BACnetContextTagObjectIdentifier
	// GetIssueConfirmedNotifications returns IssueConfirmedNotifications (property field)
	GetIssueConfirmedNotifications() BACnetContextTagBoolean
	// GetLifetime returns Lifetime (property field)
	GetLifetime() BACnetContextTagUnsignedInteger
	// GetMonitoredPropertyIdentifier returns MonitoredPropertyIdentifier (property field)
	GetMonitoredPropertyIdentifier() BACnetPropertyReferenceEnclosed
	// GetCovIncrement returns CovIncrement (property field)
	GetCovIncrement() BACnetContextTagReal
}

// BACnetConfirmedServiceRequestSubscribeCOVPropertyExactly can be used when we want exactly this type and not a type which fulfills BACnetConfirmedServiceRequestSubscribeCOVProperty.
// This is useful for switch cases.
type BACnetConfirmedServiceRequestSubscribeCOVPropertyExactly interface {
	isBACnetConfirmedServiceRequestSubscribeCOVProperty() bool
}

// _BACnetConfirmedServiceRequestSubscribeCOVProperty is the data-structure of this message
type _BACnetConfirmedServiceRequestSubscribeCOVProperty struct {
	*_BACnetConfirmedServiceRequest
	SubscriberProcessIdentifier BACnetContextTagUnsignedInteger
	MonitoredObjectIdentifier   BACnetContextTagObjectIdentifier
	IssueConfirmedNotifications BACnetContextTagBoolean
	Lifetime                    BACnetContextTagUnsignedInteger
	MonitoredPropertyIdentifier BACnetPropertyReferenceEnclosed
	CovIncrement                BACnetContextTagReal

	// Arguments.
	ServiceRequestLength uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConfirmedServiceRequestSubscribeCOVProperty) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_SUBSCRIBE_COV_PROPERTY
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestSubscribeCOVProperty) InitializeParent(parent BACnetConfirmedServiceRequest) {
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVProperty) GetParent() BACnetConfirmedServiceRequest {
	return m._BACnetConfirmedServiceRequest
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestSubscribeCOVProperty) GetSubscriberProcessIdentifier() BACnetContextTagUnsignedInteger {
	return m.SubscriberProcessIdentifier
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVProperty) GetMonitoredObjectIdentifier() BACnetContextTagObjectIdentifier {
	return m.MonitoredObjectIdentifier
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVProperty) GetIssueConfirmedNotifications() BACnetContextTagBoolean {
	return m.IssueConfirmedNotifications
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVProperty) GetLifetime() BACnetContextTagUnsignedInteger {
	return m.Lifetime
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVProperty) GetMonitoredPropertyIdentifier() BACnetPropertyReferenceEnclosed {
	return m.MonitoredPropertyIdentifier
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVProperty) GetCovIncrement() BACnetContextTagReal {
	return m.CovIncrement
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestSubscribeCOVProperty factory function for _BACnetConfirmedServiceRequestSubscribeCOVProperty
func NewBACnetConfirmedServiceRequestSubscribeCOVProperty(subscriberProcessIdentifier BACnetContextTagUnsignedInteger, monitoredObjectIdentifier BACnetContextTagObjectIdentifier, issueConfirmedNotifications BACnetContextTagBoolean, lifetime BACnetContextTagUnsignedInteger, monitoredPropertyIdentifier BACnetPropertyReferenceEnclosed, covIncrement BACnetContextTagReal, serviceRequestLength uint16) *_BACnetConfirmedServiceRequestSubscribeCOVProperty {
	_result := &_BACnetConfirmedServiceRequestSubscribeCOVProperty{
		SubscriberProcessIdentifier:    subscriberProcessIdentifier,
		MonitoredObjectIdentifier:      monitoredObjectIdentifier,
		IssueConfirmedNotifications:    issueConfirmedNotifications,
		Lifetime:                       lifetime,
		MonitoredPropertyIdentifier:    monitoredPropertyIdentifier,
		CovIncrement:                   covIncrement,
		_BACnetConfirmedServiceRequest: NewBACnetConfirmedServiceRequest(serviceRequestLength),
	}
	_result._BACnetConfirmedServiceRequest._BACnetConfirmedServiceRequestChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestSubscribeCOVProperty(structType interface{}) BACnetConfirmedServiceRequestSubscribeCOVProperty {
	if casted, ok := structType.(BACnetConfirmedServiceRequestSubscribeCOVProperty); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestSubscribeCOVProperty); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVProperty) GetTypeName() string {
	return "BACnetConfirmedServiceRequestSubscribeCOVProperty"
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVProperty) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVProperty) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (subscriberProcessIdentifier)
	lengthInBits += m.SubscriberProcessIdentifier.GetLengthInBits()

	// Simple field (monitoredObjectIdentifier)
	lengthInBits += m.MonitoredObjectIdentifier.GetLengthInBits()

	// Optional Field (issueConfirmedNotifications)
	if m.IssueConfirmedNotifications != nil {
		lengthInBits += m.IssueConfirmedNotifications.GetLengthInBits()
	}

	// Optional Field (lifetime)
	if m.Lifetime != nil {
		lengthInBits += m.Lifetime.GetLengthInBits()
	}

	// Simple field (monitoredPropertyIdentifier)
	lengthInBits += m.MonitoredPropertyIdentifier.GetLengthInBits()

	// Optional Field (covIncrement)
	if m.CovIncrement != nil {
		lengthInBits += m.CovIncrement.GetLengthInBits()
	}

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVProperty) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConfirmedServiceRequestSubscribeCOVPropertyParse(readBuffer utils.ReadBuffer, serviceRequestLength uint16) (BACnetConfirmedServiceRequestSubscribeCOVProperty, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestSubscribeCOVProperty"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestSubscribeCOVProperty")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (subscriberProcessIdentifier)
	if pullErr := readBuffer.PullContext("subscriberProcessIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for subscriberProcessIdentifier")
	}
	_subscriberProcessIdentifier, _subscriberProcessIdentifierErr := BACnetContextTagParse(readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _subscriberProcessIdentifierErr != nil {
		return nil, errors.Wrap(_subscriberProcessIdentifierErr, "Error parsing 'subscriberProcessIdentifier' field")
	}
	subscriberProcessIdentifier := _subscriberProcessIdentifier.(BACnetContextTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("subscriberProcessIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for subscriberProcessIdentifier")
	}

	// Simple Field (monitoredObjectIdentifier)
	if pullErr := readBuffer.PullContext("monitoredObjectIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for monitoredObjectIdentifier")
	}
	_monitoredObjectIdentifier, _monitoredObjectIdentifierErr := BACnetContextTagParse(readBuffer, uint8(uint8(1)), BACnetDataType(BACnetDataType_BACNET_OBJECT_IDENTIFIER))
	if _monitoredObjectIdentifierErr != nil {
		return nil, errors.Wrap(_monitoredObjectIdentifierErr, "Error parsing 'monitoredObjectIdentifier' field")
	}
	monitoredObjectIdentifier := _monitoredObjectIdentifier.(BACnetContextTagObjectIdentifier)
	if closeErr := readBuffer.CloseContext("monitoredObjectIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for monitoredObjectIdentifier")
	}

	// Optional Field (issueConfirmedNotifications) (Can be skipped, if a given expression evaluates to false)
	var issueConfirmedNotifications BACnetContextTagBoolean = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("issueConfirmedNotifications"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for issueConfirmedNotifications")
		}
		_val, _err := BACnetContextTagParse(readBuffer, uint8(2), BACnetDataType_BOOLEAN)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'issueConfirmedNotifications' field")
		default:
			issueConfirmedNotifications = _val.(BACnetContextTagBoolean)
			if closeErr := readBuffer.CloseContext("issueConfirmedNotifications"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for issueConfirmedNotifications")
			}
		}
	}

	// Optional Field (lifetime) (Can be skipped, if a given expression evaluates to false)
	var lifetime BACnetContextTagUnsignedInteger = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("lifetime"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for lifetime")
		}
		_val, _err := BACnetContextTagParse(readBuffer, uint8(3), BACnetDataType_UNSIGNED_INTEGER)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'lifetime' field")
		default:
			lifetime = _val.(BACnetContextTagUnsignedInteger)
			if closeErr := readBuffer.CloseContext("lifetime"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for lifetime")
			}
		}
	}

	// Simple Field (monitoredPropertyIdentifier)
	if pullErr := readBuffer.PullContext("monitoredPropertyIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for monitoredPropertyIdentifier")
	}
	_monitoredPropertyIdentifier, _monitoredPropertyIdentifierErr := BACnetPropertyReferenceEnclosedParse(readBuffer, uint8(uint8(4)))
	if _monitoredPropertyIdentifierErr != nil {
		return nil, errors.Wrap(_monitoredPropertyIdentifierErr, "Error parsing 'monitoredPropertyIdentifier' field")
	}
	monitoredPropertyIdentifier := _monitoredPropertyIdentifier.(BACnetPropertyReferenceEnclosed)
	if closeErr := readBuffer.CloseContext("monitoredPropertyIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for monitoredPropertyIdentifier")
	}

	// Optional Field (covIncrement) (Can be skipped, if a given expression evaluates to false)
	var covIncrement BACnetContextTagReal = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("covIncrement"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for covIncrement")
		}
		_val, _err := BACnetContextTagParse(readBuffer, uint8(5), BACnetDataType_REAL)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'covIncrement' field")
		default:
			covIncrement = _val.(BACnetContextTagReal)
			if closeErr := readBuffer.CloseContext("covIncrement"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for covIncrement")
			}
		}
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestSubscribeCOVProperty"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestSubscribeCOVProperty")
	}

	// Create a partially initialized instance
	_child := &_BACnetConfirmedServiceRequestSubscribeCOVProperty{
		SubscriberProcessIdentifier:    subscriberProcessIdentifier,
		MonitoredObjectIdentifier:      monitoredObjectIdentifier,
		IssueConfirmedNotifications:    issueConfirmedNotifications,
		Lifetime:                       lifetime,
		MonitoredPropertyIdentifier:    monitoredPropertyIdentifier,
		CovIncrement:                   covIncrement,
		_BACnetConfirmedServiceRequest: &_BACnetConfirmedServiceRequest{},
	}
	_child._BACnetConfirmedServiceRequest._BACnetConfirmedServiceRequestChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVProperty) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestSubscribeCOVProperty"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestSubscribeCOVProperty")
		}

		// Simple Field (subscriberProcessIdentifier)
		if pushErr := writeBuffer.PushContext("subscriberProcessIdentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for subscriberProcessIdentifier")
		}
		_subscriberProcessIdentifierErr := writeBuffer.WriteSerializable(m.GetSubscriberProcessIdentifier())
		if popErr := writeBuffer.PopContext("subscriberProcessIdentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for subscriberProcessIdentifier")
		}
		if _subscriberProcessIdentifierErr != nil {
			return errors.Wrap(_subscriberProcessIdentifierErr, "Error serializing 'subscriberProcessIdentifier' field")
		}

		// Simple Field (monitoredObjectIdentifier)
		if pushErr := writeBuffer.PushContext("monitoredObjectIdentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for monitoredObjectIdentifier")
		}
		_monitoredObjectIdentifierErr := writeBuffer.WriteSerializable(m.GetMonitoredObjectIdentifier())
		if popErr := writeBuffer.PopContext("monitoredObjectIdentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for monitoredObjectIdentifier")
		}
		if _monitoredObjectIdentifierErr != nil {
			return errors.Wrap(_monitoredObjectIdentifierErr, "Error serializing 'monitoredObjectIdentifier' field")
		}

		// Optional Field (issueConfirmedNotifications) (Can be skipped, if the value is null)
		var issueConfirmedNotifications BACnetContextTagBoolean = nil
		if m.GetIssueConfirmedNotifications() != nil {
			if pushErr := writeBuffer.PushContext("issueConfirmedNotifications"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for issueConfirmedNotifications")
			}
			issueConfirmedNotifications = m.GetIssueConfirmedNotifications()
			_issueConfirmedNotificationsErr := writeBuffer.WriteSerializable(issueConfirmedNotifications)
			if popErr := writeBuffer.PopContext("issueConfirmedNotifications"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for issueConfirmedNotifications")
			}
			if _issueConfirmedNotificationsErr != nil {
				return errors.Wrap(_issueConfirmedNotificationsErr, "Error serializing 'issueConfirmedNotifications' field")
			}
		}

		// Optional Field (lifetime) (Can be skipped, if the value is null)
		var lifetime BACnetContextTagUnsignedInteger = nil
		if m.GetLifetime() != nil {
			if pushErr := writeBuffer.PushContext("lifetime"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for lifetime")
			}
			lifetime = m.GetLifetime()
			_lifetimeErr := writeBuffer.WriteSerializable(lifetime)
			if popErr := writeBuffer.PopContext("lifetime"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for lifetime")
			}
			if _lifetimeErr != nil {
				return errors.Wrap(_lifetimeErr, "Error serializing 'lifetime' field")
			}
		}

		// Simple Field (monitoredPropertyIdentifier)
		if pushErr := writeBuffer.PushContext("monitoredPropertyIdentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for monitoredPropertyIdentifier")
		}
		_monitoredPropertyIdentifierErr := writeBuffer.WriteSerializable(m.GetMonitoredPropertyIdentifier())
		if popErr := writeBuffer.PopContext("monitoredPropertyIdentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for monitoredPropertyIdentifier")
		}
		if _monitoredPropertyIdentifierErr != nil {
			return errors.Wrap(_monitoredPropertyIdentifierErr, "Error serializing 'monitoredPropertyIdentifier' field")
		}

		// Optional Field (covIncrement) (Can be skipped, if the value is null)
		var covIncrement BACnetContextTagReal = nil
		if m.GetCovIncrement() != nil {
			if pushErr := writeBuffer.PushContext("covIncrement"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for covIncrement")
			}
			covIncrement = m.GetCovIncrement()
			_covIncrementErr := writeBuffer.WriteSerializable(covIncrement)
			if popErr := writeBuffer.PopContext("covIncrement"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for covIncrement")
			}
			if _covIncrementErr != nil {
				return errors.Wrap(_covIncrementErr, "Error serializing 'covIncrement' field")
			}
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestSubscribeCOVProperty"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestSubscribeCOVProperty")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVProperty) isBACnetConfirmedServiceRequestSubscribeCOVProperty() bool {
	return true
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVProperty) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
