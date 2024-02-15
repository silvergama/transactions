package postgres

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/silvergama/transations/transaction"
	"github.com/stretchr/testify/assert"
)

func TestRepository_Create(t *testing.T) {
	tests := []struct {
		name        string
		mockedQuery func(mock sqlmock.Sqlmock)
		wantID      int
		wantErr     error
	}{
		{
			name: "should return transaction id",
			mockedQuery: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("INSERT INTO transactions").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
			},
			wantID:  1,
			wantErr: nil,
		},
		{
			name: "should return an error",
			mockedQuery: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("INSERT INTO transactions").WillReturnError(errors.New("mock error"))
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

			repo := NewTransaction(mockDB)
			tt.mockedQuery(mock)

			gotID, err := repo.Create(context.Background(), &transaction.Transaction{AccountID: 1, OperationTypeID: 1, Amount: 100.0})

			assert.Equal(t, tt.wantID, gotID)
			assert.Equal(t, tt.wantErr, err)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}
