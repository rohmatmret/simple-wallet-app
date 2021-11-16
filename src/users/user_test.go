package users

import (
	"github.com/google/uuid"
	"github.com/scoop-wallet/src/db"
	"gorm.io/gorm"
	"testing"
)

func Test_repository_Save(t *testing.T) {
	id := uuid.New().String() // create uuid for primary userid
	newUser := NewUsers(id,"Haidir ali")
	type fields struct {
		db gorm.DB
	}
	type args struct {
		user Users
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Create new users",
			fields: fields{db: *db.Connect()},
			args: args{*newUser},
			wantErr: false,
		},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := repository{
				db: tt.fields.db,
			}
			if err := d.Save(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
