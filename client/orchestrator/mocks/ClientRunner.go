// Copyright 2020 Northern.tech AS
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.
package mocks

import context "context"
import mock "github.com/stretchr/testify/mock"
import orchestrator "github.com/mendersoftware/deviceauth/client/orchestrator"

// ClientRunner is an autogenerated mock type for the ClientRunner type
type ClientRunner struct {
	mock.Mock
}

// SubmitDeviceDecommisioningJob provides a mock function with given fields: ctx, req
func (_m *ClientRunner) SubmitDeviceDecommisioningJob(ctx context.Context, req orchestrator.DecommissioningReq) error {
	ret := _m.Called(ctx, req)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, orchestrator.DecommissioningReq) error); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SubmitProvisionDeviceJob provides a mock function with given fields: ctx, req
func (_m *ClientRunner) SubmitProvisionDeviceJob(ctx context.Context, req orchestrator.ProvisionDeviceReq) error {
	ret := _m.Called(ctx, req)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, orchestrator.ProvisionDeviceReq) error); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *ClientRunner) SubmitUpdateDeviceStatusJob(ctx context.Context, req orchestrator.UpdateDeviceStatusReq) error {
	ret := _m.Called(ctx, req)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, orchestrator.UpdateDeviceStatusReq) error); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
