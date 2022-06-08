package inventory

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewProductSuccess(t *testing.T) {
	newProduct, err := NewProduct("test", 3.0, 5)
	assert.Nil(t, err)
	assert.NotNil(t, newProduct)
	assert.Equal(t, "test", newProduct.Title)
	assert.Equal(t, 3.0, newProduct.AveragePrice)
	assert.Equal(t, 5, newProduct.Stock)
}

func TestNewProductErrTitle(t *testing.T) {
	newProduct, err := NewProduct("", 3.0, 5)
	assert.Nil(t, newProduct)
	assert.NotNil(t, err)
}

func TestNewProductErrPrice(t *testing.T) {
	newProduct, err := NewProduct("title", 0, 5)
	assert.Nil(t, newProduct)
	assert.NotNil(t, err)
}

func TestNewProductErrQuantity(t *testing.T) {
	newProduct, err := NewProduct("title", 1, 0)
	assert.Nil(t, newProduct)
	assert.NotNil(t, err)
}

func TestProduct_UpdateSalePrice(t *testing.T) {
	type fields struct {
		Id           uuid.UUID
		Title        string
		SalePrice    float64
		AveragePrice float64
		Stock        int
	}
	type args struct {
		newSalePrice float64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "success_test",
			fields: fields{
				Id:           uuid.New(),
				Title:        "title",
				SalePrice:    3,
				AveragePrice: 3,
				Stock:        3,
			},
			args: args{
				newSalePrice: 3.0,
			},
			wantErr: false,
		},
		{
			name: "fail_test",
			fields: fields{
				Id:           uuid.New(),
				Title:        "title",
				SalePrice:    3,
				AveragePrice: 3,
				Stock:        3,
			},
			args: args{
				newSalePrice: 2.99999,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Product{
				Id:           tt.fields.Id,
				Title:        tt.fields.Title,
				SalePrice:    tt.fields.SalePrice,
				AveragePrice: tt.fields.AveragePrice,
				Stock:        tt.fields.Stock,
			}
			if err := p.UpdateSalePrice(tt.args.newSalePrice); (err != nil) != tt.wantErr {
				t.Errorf("UpdateSalePrice() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, tt.args.newSalePrice, p.SalePrice)
			}
		})
	}
}
