package service

import (
	"reflect"
	"testing"

	"github.com/Dshepett/payment-service/internal/models"
	"github.com/Dshepett/payment-service/internal/storage"
	"github.com/Dshepett/payment-service/internal/storage/local"
)

func TestService_AddPayment(t *testing.T) {
	type fields struct {
		storage       storage.Storage
		adminUsername string
		adminPassword string
	}
	type args struct {
		payment *models.Payment
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "It just works!",
			fields: fields{
				storage:       &local.Storage{PaymentRepository: &local.PaymentRepository{}},
				adminUsername: "any",
				adminPassword: "any",
			},
			args:    args{&models.Payment{}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				storage:       tt.fields.storage,
				adminUsername: tt.fields.adminUsername,
				adminPassword: tt.fields.adminPassword,
			}
			if err := s.AddPayment(tt.args.payment); (err != nil) != tt.wantErr {
				t.Errorf("AddPayment() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_ChangePaymentStatus(t *testing.T) {
	type fields struct {
		storage       storage.Storage
		adminUsername string
		adminPassword string
	}
	type args struct {
		id     int
		status string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Incorrect status",
			fields: fields{
				storage: &local.Storage{
					PaymentRepository: &local.PaymentRepository{
						Payments: []models.Payment{
							{
								Id:     1,
								Status: models.New,
							},
						},
					},
				},
				adminUsername: "any",
				adminPassword: "any",
			},
			args: args{
				id:     1,
				status: "incorrect_status",
			},
			wantErr: true,
		},
		{
			name: "Status could not be changed",
			fields: fields{
				storage: &local.Storage{
					PaymentRepository: &local.PaymentRepository{
						Payments: []models.Payment{
							{
								Id:     1,
								Status: models.Success,
							},
						},
					},
				},
				adminUsername: "any",
				adminPassword: "any",
			},
			args: args{
				id:     1,
				status: "FAILURE",
			},
			wantErr: true,
		},
		{
			name: "Payment does not exist",
			fields: fields{
				storage: &local.Storage{
					PaymentRepository: &local.PaymentRepository{
						Payments: []models.Payment{
							{
								Id:     1,
								Status: models.New,
							},
						},
					},
				},
				adminUsername: "any",
				adminPassword: "any",
			},
			args: args{
				id:     2,
				status: "FAILURE",
			},
			wantErr: true,
		},
		{
			name: "Everything is correct",
			fields: fields{
				storage: &local.Storage{
					PaymentRepository: &local.PaymentRepository{
						Payments: []models.Payment{
							{
								Id:     1,
								Status: models.New,
							},
						},
					},
				},
				adminUsername: "any",
				adminPassword: "any",
			},
			args: args{
				id:     1,
				status: "FAILURE",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				storage:       tt.fields.storage,
				adminUsername: tt.fields.adminUsername,
				adminPassword: tt.fields.adminPassword,
			}
			if err := s.ChangePaymentStatus(tt.args.id, tt.args.status); (err != nil) != tt.wantErr {
				t.Errorf("ChangePaymentStatus() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_DenyPayment(t *testing.T) {
	type fields struct {
		storage       storage.Storage
		adminUsername string
		adminPassword string
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Payment exist with available to be denied status and will by denyed",
			fields: fields{
				storage: &local.Storage{
					PaymentRepository: &local.PaymentRepository{
						Payments: []models.Payment{
							{
								Id:     1,
								Status: models.New,
							},
						},
					},
				},
				adminUsername: "any",
				adminPassword: "any",
			},
			args: args{
				id: 1,
			},
			wantErr: false,
		},
		{
			name: "Payment exist with not available to be denied status and will by denyed",
			fields: fields{
				storage: &local.Storage{
					PaymentRepository: &local.PaymentRepository{
						Payments: []models.Payment{
							{
								Id:     1,
								Status: models.Success,
							},
						},
					},
				},
				adminUsername: "any",
				adminPassword: "any",
			},
			args: args{
				id: 1,
			},
			wantErr: true,
		},
		{
			name: "Payment does not exist",
			fields: fields{
				storage: &local.Storage{
					PaymentRepository: &local.PaymentRepository{
						Payments: []models.Payment{
							{
								Id:     1,
								Status: models.New,
							},
						},
					},
				},
				adminUsername: "any",
				adminPassword: "any",
			},
			args: args{
				id: 2,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				storage:       tt.fields.storage,
				adminUsername: tt.fields.adminUsername,
				adminPassword: tt.fields.adminPassword,
			}
			if err := s.DenyPayment(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DenyPayment() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_GenerateToken(t *testing.T) {
	type fields struct {
		storage       storage.Storage
		adminUsername string
		adminPassword string
	}
	type args struct {
		username string
		password string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Correct data",
			fields: fields{
				storage:       nil,
				adminUsername: "User",
				adminPassword: "123",
			},
			args: args{
				username: "User",
				password: "123",
			},
			wantErr: false,
		},
		{
			name: "Incorrect data(password)",
			fields: fields{
				storage:       nil,
				adminUsername: "User",
				adminPassword: "123",
			},
			args: args{
				username: "User",
				password: "1234",
			},
			wantErr: true,
		},
		{
			name: "Incorrect data(username)",
			fields: fields{
				storage:       nil,
				adminUsername: "User",
				adminPassword: "123",
			},
			args: args{
				username: "User1",
				password: "123",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				storage:       tt.fields.storage,
				adminUsername: tt.fields.adminUsername,
				adminPassword: tt.fields.adminPassword,
			}
			_, err := s.GenerateToken(tt.args.username, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestService_PaymentStatus(t *testing.T) {
	type fields struct {
		storage       storage.Storage
		adminUsername string
		adminPassword string
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Payment exist",
			fields: fields{
				storage: &local.Storage{
					PaymentRepository: &local.PaymentRepository{
						Payments: []models.Payment{
							{
								Id:     1,
								Status: models.New,
							},
						},
					},
				},
				adminUsername: "any",
				adminPassword: "any",
			},
			args: args{
				id: 1,
			},
			want:    "NEW",
			wantErr: false,
		},
		{
			name: "Payment does not exist",
			fields: fields{
				storage: &local.Storage{
					PaymentRepository: &local.PaymentRepository{
						Payments: []models.Payment{
							{
								Id:     1,
								Status: models.New,
							},
						},
					},
				},
				adminUsername: "any",
				adminPassword: "any",
			},
			args: args{
				id: 2,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Payment exist with SUCCESS status",
			fields: fields{
				storage: &local.Storage{
					PaymentRepository: &local.PaymentRepository{
						Payments: []models.Payment{
							{
								Id:     1,
								Status: models.Success,
							},
						},
					},
				},
				adminUsername: "any",
				adminPassword: "any",
			},
			args: args{
				id: 1,
			},
			want:    "SUCCESS",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				storage:       tt.fields.storage,
				adminUsername: tt.fields.adminUsername,
				adminPassword: tt.fields.adminPassword,
			}
			got, err := s.PaymentStatus(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("PaymentStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PaymentStatus() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_PaymentsByEmail(t *testing.T) {
	type fields struct {
		storage       storage.Storage
		adminUsername string
		adminPassword string
	}
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []models.Payment
		wantErr bool
	}{
		{
			name: "Return 2 elements",
			fields: fields{
				storage: &local.Storage{
					PaymentRepository: &local.PaymentRepository{
						Payments: []models.Payment{
							{
								Id:        1,
								Status:    models.New,
								UserEmail: "123",
							},
							{
								Id:        2,
								Status:    models.New,
								UserEmail: "1234",
							},
							{
								Id:        3,
								Status:    models.New,
								UserEmail: "123",
							},
						},
					},
				},
				adminUsername: "any",
				adminPassword: "any",
			},
			args: args{
				email: "123",
			},
			want: []models.Payment{
				{
					Id:        1,
					Status:    models.New,
					UserEmail: "123",
				},
				{
					Id:        3,
					Status:    models.New,
					UserEmail: "123",
				},
			},
			wantErr: false,
		},
		{
			name: "Returns empty slice",
			fields: fields{
				storage: &local.Storage{
					PaymentRepository: &local.PaymentRepository{
						Payments: []models.Payment{
							{
								Id:        1,
								Status:    models.New,
								UserEmail: "123",
							},
							{
								Id:        2,
								Status:    models.New,
								UserEmail: "1234",
							},
							{
								Id:        3,
								Status:    models.New,
								UserEmail: "123",
							},
						},
					},
				},
				adminUsername: "any",
				adminPassword: "any",
			},
			args: args{
				email: "12345",
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				storage:       tt.fields.storage,
				adminUsername: tt.fields.adminUsername,
				adminPassword: tt.fields.adminPassword,
			}
			got, err := s.PaymentsByEmail(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("PaymentsByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PaymentsByEmail() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_PaymentsByUserId(t *testing.T) {
	type fields struct {
		storage       storage.Storage
		adminUsername string
		adminPassword string
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []models.Payment
		wantErr bool
	}{
		{
			name: "Return 2 elements",
			fields: fields{
				storage: &local.Storage{
					PaymentRepository: &local.PaymentRepository{
						Payments: []models.Payment{
							{
								Id:     1,
								Status: models.New,
								UserId: 123,
							},
							{
								Id:     2,
								Status: models.New,
								UserId: 1234,
							},
							{
								Id:     3,
								Status: models.New,
								UserId: 123,
							},
						},
					},
				},
				adminUsername: "any",
				adminPassword: "any",
			},
			args: args{
				id: 123,
			},
			want: []models.Payment{
				{
					Id:     1,
					Status: models.New,
					UserId: 123,
				},
				{
					Id:     3,
					Status: models.New,
					UserId: 123,
				},
			},
			wantErr: false,
		},
		{
			name: "Returns empty slice",
			fields: fields{
				storage: &local.Storage{
					PaymentRepository: &local.PaymentRepository{
						Payments: []models.Payment{
							{
								Id:     1,
								Status: models.New,
								UserId: 123,
							},
							{
								Id:     2,
								Status: models.New,
								UserId: 1234,
							},
							{
								Id:     3,
								Status: models.New,
								UserId: 123,
							},
						},
					},
				},
				adminUsername: "any",
				adminPassword: "any",
			},
			args: args{
				id: 12345,
			},
			want:    nil,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				storage:       tt.fields.storage,
				adminUsername: tt.fields.adminUsername,
				adminPassword: tt.fields.adminPassword,
			}
			got, err := s.PaymentsByUserId(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("PaymentsByUserId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PaymentsByUserId() got = %v, want %v", got, tt.want)
			}
		})
	}
}
