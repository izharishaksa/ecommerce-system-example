package customer

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCustomer(t *testing.T) {
	type args struct {
		name  string
		email string
	}
	tests := []struct {
		name    string
		args    args
		want    *Customer
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				name:  "customer name",
				email: "izharishaksa@gmail.com",
			},
			want: &Customer{
				Id:    uuid.New(),
				Name:  "customer name",
				Email: "izharishaksa@gmail.com",
			},
			wantErr: false,
		},
		{
			name: "failed_name",
			args: args{
				name: "",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "failed_email",
			args: args{
				name:  "Izhari Ishak Aksa",
				email: "email",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "failed_empty_email",
			args: args{
				name:  "Izhari Ishak Aksa",
				email: "",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewCustomer(tt.args.name, tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCustomer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr == false {
				assert.NotNil(t, got)
				assert.Nil(t, err)
				assert.Equal(t, tt.args.name, got.Name)
				assert.Equal(t, tt.args.email, got.Email)
			}
		})
	}
}

func TestCustomer_TopUp(t *testing.T) {
	type fields struct {
		Id      uuid.UUID
		Name    string
		Balance float64
	}
	type args struct {
		amount float64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "success",
			fields: struct {
				Id      uuid.UUID
				Name    string
				Balance float64
			}{Id: uuid.New(), Name: "customer name", Balance: 200},
			args:    struct{ amount float64 }{amount: 100},
			wantErr: false,
		},
		{
			name: "failed",
			fields: struct {
				Id      uuid.UUID
				Name    string
				Balance float64
			}{Id: uuid.New(), Name: "customer name", Balance: 200},
			args:    struct{ amount float64 }{amount: -100},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Customer{
				Id:      tt.fields.Id,
				Name:    tt.fields.Name,
				Balance: tt.fields.Balance,
			}
			if err := c.TopUp(tt.args.amount); (err != nil) != tt.wantErr {
				t.Errorf("TopUp() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr == false {
				assert.Equal(t, tt.fields.Balance+tt.args.amount, c.Balance)
			}
		})
	}
}
