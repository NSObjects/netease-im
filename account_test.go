/*
 *
 * account_test.go
 * netease
 *
 * Created by lintao on 2020/6/10 11:11 上午
 * Copyright © 2020-2020 LINTAO. All rights reserved.
 *
 */

package netease

import (
	"fmt"
	"os"
	"testing"
)

var (
	AppKey string
	Secret string
)

//测试用的， 大家自己设置环境变量吧
func init() {
	AppKey = os.Getenv("app_key")
	Secret = os.Getenv("secret")
}

func TestNetEaseIM_CreateAccount(t *testing.T) {
	type fields struct {
		appKey   string
		secret   string
		basePath string
	}
	type args struct {
		account Account
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "创建账号",
			fields: fields{
				appKey: AppKey,
				secret: Secret,
			},
			args: args{Account{
				Accid: "5e82ddaf7da7a00001c12ef8",
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := NewNetEaseIM(tt.fields.appKey, tt.fields.secret)
			got, err := n.CreateAccount(tt.args.account)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(got)
		})
	}
}

func TestNetEaseIM_RefreshToken(t *testing.T) {
	type fields struct {
		appKey   string
		secret   string
		basePath string
	}
	type args struct {
		accid string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "刷新Id",
			fields: fields{
				appKey: AppKey,
				secret: Secret,
			},
			args:    args{accid: "15274894959"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := NewNetEaseIM(tt.fields.appKey, tt.fields.secret)
			got, err := n.RefreshToken(tt.args.accid)
			if (err != nil) != tt.wantErr {
				t.Errorf("RefreshToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(got)
			//if got != tt.want {
			//	t.Errorf("RefreshToken() got = %v, want %v", got, tt.want)
			//}
		})
	}
}
