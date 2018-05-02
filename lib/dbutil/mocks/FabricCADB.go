/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

                 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by mockery v1.0.0

package mocks

import dbutil "github.com/hyperledger/fabric-ca/lib/dbutil"
import mock "github.com/stretchr/testify/mock"
import sql "database/sql"
import sqlx "github.com/jmoiron/sqlx"

// FabricCADB is an autogenerated mock type for the FabricCADB type
type FabricCADB struct {
	mock.Mock
}

// BeginTx provides a mock function with given fields:
func (_m *FabricCADB) BeginTx() dbutil.FabricCATx {
	ret := _m.Called()

	var r0 dbutil.FabricCATx
	if rf, ok := ret.Get(0).(func() dbutil.FabricCATx); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(dbutil.FabricCATx)
		}
	}

	return r0
}

// MustBegin provides a mock function with given fields:
func (_m *FabricCADB) MustBegin() *sqlx.Tx {
	ret := _m.Called()

	var r0 *sqlx.Tx
	if rf, ok := ret.Get(0).(func() *sqlx.Tx); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sqlx.Tx)
		}
	}

	return r0
}

// NamedExec provides a mock function with given fields: query, arg
func (_m *FabricCADB) NamedExec(query string, arg interface{}) (sql.Result, error) {
	ret := _m.Called(query, arg)

	var r0 sql.Result
	if rf, ok := ret.Get(0).(func(string, interface{}) sql.Result); ok {
		r0 = rf(query, arg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(sql.Result)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, interface{}) error); ok {
		r1 = rf(query, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Rebind provides a mock function with given fields: query
func (_m *FabricCADB) Rebind(query string) string {
	ret := _m.Called(query)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(query)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Select provides a mock function with given fields: dest, query, args
func (_m *FabricCADB) Select(dest interface{}, query string, args ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, dest, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, string, ...interface{}) error); ok {
		r0 = rf(dest, query, args...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}