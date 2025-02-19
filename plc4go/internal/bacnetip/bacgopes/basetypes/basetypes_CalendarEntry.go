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

type CalendarEntry struct {
	*Choice
	choiceElements []Element
}

func NewCalendarEntry(arg Arg) (*CalendarEntry, error) {
	s := &CalendarEntry{
		choiceElements: []Element{
			NewElement("date", Vs2E(NewDate), WithElementContext(0)),
			NewElement("dateRange", V2E(NewDateRange), WithElementContext(1)),
			NewElement("weekNDay", V2E(NewWeekNDay), WithElementContext(2)),
		},
	}
	panic("implementchoice")
	return s, nil
}
