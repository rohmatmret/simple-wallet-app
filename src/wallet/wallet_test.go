package wallet

import (
	"github.com/scoop-wallet/src/db"
	"gorm.io/gorm"
	"testing"
)

func Test_repository_ValidateBalance(t *testing.T) {
	type fields struct {
		db gorm.DB
	}
	type args struct {
		amount int64
		userID string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Validate data",
			fields: fields{db: *db.Connect()},
			args: args{amount: 12000,userID: "21"},
			want: true,
		},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &repository{
				db: tt.fields.db,
			}
			if got := d.ValidateBalance(tt.args.amount, tt.args.userID); got != tt.want {
				t.Errorf("ValidateBalance() = %v, want %v", got, tt.want)
			}
		})
	}
}
