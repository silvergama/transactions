package postgres

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/silvergama/transations/account"
	"github.com/stretchr/testify/assert"
)

func TestRepository_Create(t *testing.T) {
	type arg struct {
		account *account.Account
	}
	tests := []struct {
		name        string
		arg         arg
		mockedQuery func(mock sqlmock.Sqlmock)
		wantID      int
		wantErr     error
	}{
		{
			name: "should return account id",
			arg: arg{
				account: &account.Account{
					DocumentNumber: "123456",
				},
			},
			mockedQuery: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("INSERT INTO account").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
			},
			wantID:  1,
			wantErr: nil,
		},
		{
			name: "should return an error",
			arg: arg{
				account: &account.Account{
					AccoundID:      1,
					DocumentNumber: "987765",
				},
			},
			mockedQuery: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("INSERT INTO account").WillReturnError(errors.New("mock error"))
			},
			wantID:  0,
			wantErr: errors.New("mock error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("error initializing sqlmock: %v", err)
			}
			defer mockDB.Close()

			repo := NewAccount(mockDB)
			tt.mockedQuery(mock)

			gotID, err := repo.Create(context.Background(), tt.arg.account)

			assert.Equal(t, tt.wantID, gotID)
			assert.Equal(t, tt.wantErr, err)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestRepository_GetByID(t *testing.T) {
	type arg struct {
		accountID int
	}
	tests := []struct {
		name        string
		arg         arg
		mockedQuery func(mock sqlmock.Sqlmock)
		wantAccount *account.Account
		wantErr     error
	}{
		{
			name: "Should return account",
			arg: arg{
				accountID: 1,
			},
			mockedQuery: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "document_number"}).
					AddRow(1, "123456")
				mock.ExpectQuery("SELECT id, document_number FROM accounts").WithArgs(1).WillReturnRows(rows)
			},
			wantAccount: &account.Account{
				AccoundID:      1,
				DocumentNumber: "123456",
			},
		},
		{
			name: "Should return an error",
			arg: arg{
				accountID: 2,
			},
			mockedQuery: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT id, document_number FROM accounts").WithArgs(2).WillReturnError(errors.New("mock error"))
			},
			wantAccount: nil,
			wantErr:     errors.New("mock error"),
		},
		{
			name: "Should return ErrNoRows",
			arg: arg{
				accountID: 2,
			},
			mockedQuery: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT id, document_number FROM accounts").WithArgs(2).WillReturnError(sql.ErrNoRows)
			},
			wantAccount: nil,
			wantErr:     sql.ErrNoRows,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("error initializing sqlmock: %v", err)
			}
			defer mockDB.Close()

			repo := NewAccount(mockDB)
			tt.mockedQuery(mock)

			gotAccount, err := repo.GetByID(context.Background(), tt.arg.accountID)

			assert.Equal(t, tt.wantAccount, gotAccount)
			assert.Equal(t, tt.wantErr, err)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}
