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

package basetypes

import (
	. "github.com/apache/plc4x/plc4go/internal/bacnetip/bacgopes/comp"
	. "github.com/apache/plc4x/plc4go/internal/bacnetip/bacgopes/constructeddata"
	. "github.com/apache/plc4x/plc4go/internal/bacnetip/bacgopes/primitivedata"
)

type COVMultipleSubscription struct {
	*Sequence
	sequenceElements []Element
}

func NewCOVMultipleSubscription(arg Arg) (*COVMultipleSubscription, error) {
	s := &COVMultipleSubscription{
		sequenceElements: []Element{
			NewElement("recipient", V2E(NewRecipientProcess), WithElementContext(0)),
			NewElement("issueConfirmedNotifications", V2E(NewBoolean), WithElementContext(1)),
			NewElement("timeRemaining", V2E(NewUnsigned), WithElementContext(2)),
			NewElement("maxNotificationDelay", V2E(NewUnsigned), WithElementContext(3)),
			NewElement("listOfCOVSubscriptionSpecifications", SequenceOfE(NewCOVMultipleSubscriptionList), WithElementContext(4)),
		},
	}
	panic("implementchoice")
	return s, nil
}
