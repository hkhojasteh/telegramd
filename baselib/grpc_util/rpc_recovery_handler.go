/*
 *  Copyright (c) 2017, https://github.com/nebulaim
 *  All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package grpc_util

import (
	"google.golang.org/grpc/codes"
	"context"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc"
	"github.com/nebulaim/telegramd/proto/mtproto"
	"runtime/debug"
	"github.com/golang/glog"
)

func BizUnaryRecoveryHandler(ctx context.Context, p interface{}) (err error) {
	// debug.PrintStack()
	glog.Errorf("BizUnaryRecoveryHandler - %s", debug.Stack())
	// glog.Info(debug.Stack())

	switch p.(type) {
	case *mtproto.TLRpcError:
		code, _ := p.(*mtproto.TLRpcError)
		md, _ := RpcErrorToMD(code)
		grpc.SetTrailer(ctx, md)
		err = status.Errorf(codes.Unknown, "panic triggered rpc_error: {%v}", p)
	default:
		err = status.Errorf(codes.Unknown, "panic unknown triggered: %v", p)
	}
	glog.Errorf("Panic: %v", err.Error())
	return
}

func BizUnaryRecoveryHandler2(ctx context.Context, p interface{}) (err error) {
	switch p.(type) {
	case *mtproto.TLRpcError:
		code, _ := p.(*mtproto.TLRpcError)
		md, _ := RpcErrorToMD(code)
		grpc.SetTrailer(ctx, md)
		err = p.(*mtproto.TLRpcError)
		// err = status.Errorf(codes.Unknown, "panic triggered rpc_error: {%v}", p)
	default:
		err = status.Errorf(codes.Unknown, "panic unknown triggered: %v", p)
	}
	return
}

func BizStreamRecoveryHandler(stream grpc.ServerStream, p interface{}) (err error) {
	return
}

