package models

import (
	"reflect"
	"testing"
)

func TestCreatePayment(t *testing.T) {
	type args struct {
		request NewPaymentRequest
	}
	tests := []struct {
		name string
		args args
		want *Payment
	}{
		{
			name: "Incorrect userId",
			args: args{request: NewPaymentRequest{
				UserId:    -15,
				UserEmail: "correct@mail.ru",
				Amount:    15,
				Currency:  "DUB",
			}},
			want: nil,
		},
		{
			name: "Incorrect amount",
			args: args{request: NewPaymentRequest{
				UserId:    15,
				UserEmail: "correct@mail.ru",
				Amount:    -15,
				Currency:  "DUB",
			}},
			want: nil,
		},
		{
			name: "Incorrect email",
			args: args{request: NewPaymentRequest{
				UserId:    15,
				UserEmail: "correctmail.ru",
				Amount:    15,
				Currency:  "DUB",
			}},
			want: nil,
		},
		{
			name: "Incorrect currency",
			args: args{request: NewPaymentRequest{
				UserId:    15,
				UserEmail: "correct@mail.ru",
				Amount:    15,
				Currency:  "",
			}},
			want: nil,
		},
		{
			name: "Everyting correct",
			args: args{request: NewPaymentRequest{
				UserId:    15,
				UserEmail: "correct@mail.ru",
				Amount:    15,
				Currency:  "DUB",
			}},
			want: &Payment{
				UserId:    15,
				UserEmail: "correct@mail.ru",
				Amount:    15,
				Currency:  "DUB",
				Status:    New,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreatePayment(tt.args.request); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreatePayment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsStatus(t *testing.T) {
	type args struct {
		status string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Correct NEW",
			args: args{status: "NEW"},
			want: true,
		},
		{
			name: "Correct ERROR",
			args: args{status: "ERROR"},
			want: true,
		},
		{
			name: "Correct SUCCESS",
			args: args{status: "SUCCESS"},
			want: true,
		},
		{
			name: "Correct FAILURE",
			args: args{status: "FAILURE"},
			want: true,
		},
		{
			name: "Incorrect",
			args: args{status: "gergerhe"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsStatus(tt.args.status); got != tt.want {
				t.Errorf("IsStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}
