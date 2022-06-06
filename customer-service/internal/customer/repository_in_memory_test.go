package customer

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInMemoryRepository_SaveCustomer(t *testing.T) {
	type fields struct {
		customers map[uuid.UUID]*Customer
	}
	type args struct {
		customer *Customer
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:   testing.CoverMode(),
			fields: struct{ customers map[uuid.UUID]*Customer }{customers: map[uuid.UUID]*Customer{}},
			args: struct{ customer *Customer }{customer: &Customer{
				Id:      uuid.New(),
				Name:    "Izhari Ishak Aksa",
				Email:   "izharishaksa@gmail.com",
				Balance: 0,
			}},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(t, err, i...)
			},
		},
		{
			name: testing.CoverMode(),
			fields: struct{ customers map[uuid.UUID]*Customer }{customers: map[uuid.UUID]*Customer{
				uuid.New(): {
					Id:    uuid.New(),
					Name:  "Izhari Ishak Aksa",
					Email: "izharishaksa@gmail.com",
				},
			}},
			args: struct{ customer *Customer }{customer: &Customer{
				Id:      uuid.New(),
				Name:    "Izhari Ishak Aksa",
				Email:   "izharishaksa@gmail.com",
				Balance: 0,
			}},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.Error(t, err, i...)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := inMemoryRepository{
				customers: tt.fields.customers,
			}
			tt.wantErr(t, repo.SaveCustomer(tt.args.customer), fmt.Sprintf("SaveCustomer(%v)", tt.args.customer))
		})
	}
}
