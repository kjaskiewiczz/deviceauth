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
import inventory "github.com/mendersoftware/deviceauth/client/inventory"
import mock "github.com/stretchr/testify/mock"

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// PatchDeviceV2 provides a mock function with given fields: ctx, did, tid, src, ts, attrs
func (_m *Client) PatchDeviceV2(ctx context.Context, did string, tid string, src string, ts int64, attrs []inventory.Attribute) error {
	ret := _m.Called(ctx, did, tid, src, ts, attrs)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, int64, []inventory.Attribute) error); ok {
		r0 = rf(ctx, did, tid, src, ts, attrs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *Client) SetDeviceStatus(ctx context.Context, tenantId string, deviceIds []string, status string) error {
	ret := _m.Called(ctx, tenantId, deviceIds, status)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, []string, string) error); ok {
		r0 = rf(ctx, tenantId, deviceIds, status)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
