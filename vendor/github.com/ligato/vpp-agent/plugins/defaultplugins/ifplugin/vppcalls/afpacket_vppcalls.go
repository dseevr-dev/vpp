// Copyright (c) 2017 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vppcalls

import (
	"fmt"

	"time"

	"github.com/ligato/cn-infra/logging/measure"
	"github.com/ligato/vpp-agent/plugins/defaultplugins/common/bin_api/af_packet"
	intf "github.com/ligato/vpp-agent/plugins/defaultplugins/common/model/interfaces"
)

// AddAfPacketInterface calls AfPacketCreate VPP binary API.
func AddAfPacketInterface(afPacketIntf *intf.Interfaces_Interface_Afpacket, vppChan VPPChannel, timeLog measure.StopWatchEntry) (swIndex uint32, err error) {
	// AfPacketCreate time measurement
	start := time.Now()
	defer func() {
		if timeLog != nil {
			timeLog.LogTimeEntry(time.Since(start))
		}
	}()

	// Prepare the message.
	req := &af_packet.AfPacketCreate{
		HostIfName:      []byte(afPacketIntf.HostIfName),
		UseRandomHwAddr: 1,
	}

	reply := &af_packet.AfPacketCreateReply{}
	err = vppChan.SendRequest(req).ReceiveReply(reply)
	if err != nil {
		return 0, err
	}
	if reply.Retval != 0 {
		return 0, fmt.Errorf("add af_packet interface (%+v) returned %d", afPacketIntf, reply.Retval)
	}

	return reply.SwIfIndex, nil
}

// DeleteAfPacketInterface calls AfPacketDelete VPP binary API.
func DeleteAfPacketInterface(afPacketIntf *intf.Interfaces_Interface_Afpacket, vppChan VPPChannel, timeLog measure.StopWatchEntry) error {
	// AfPacketDelete time measurement
	start := time.Now()
	defer func() {
		if timeLog != nil {
			timeLog.LogTimeEntry(time.Since(start))
		}
	}()

	// Prepare the message.
	req := &af_packet.AfPacketDelete{
		HostIfName: []byte(afPacketIntf.HostIfName),
	}

	reply := &af_packet.AfPacketDeleteReply{}
	err := vppChan.SendRequest(req).ReceiveReply(reply)
	if err != nil {
		return err
	}
	if reply.Retval != 0 {
		return fmt.Errorf("deleting of af_packet interface (%+v) returned %d", afPacketIntf, reply.Retval)
	}

	return nil
}
