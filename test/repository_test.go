package test

import (
	"database/sql"
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/SyberiaEmperor/avito_task/models"
	"github.com/SyberiaEmperor/avito_task/pkg/repository"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Suite struct {
	suite.Suite
	DB *gorm.DB
	mock sqlmock.Sqlmock

	repo repository.Repository
	account models.Account
}

func (s *Suite) SetupSuite() {
	var (
		db *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(),err)

	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})

	gdb, err := gorm.Open(dialector,&gorm.Config{})
	require.NoError(s.T(),err)

	s.repo = *repository.NewRepository(gdb)
}

func (s *Suite) Test_RepositoryGetAccountInfo_NoAccount() {

	id := 1
	
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "accounts" WHERE "accounts"."id" = $1 ORDER BY "accounts"."id" LIMIT 1`)).WithArgs(id).
		WillReturnError(fmt.Errorf("record not found"))

	_, err := s.repo.GetAccountInfo(id)
	
	require.Error(s.T(),err)
}

func (s *Suite) Test_RepositoryGetAccountInfo_AccountExists() {

	s.repo.Deposit(models.AccountRequest{ID:1,Income: 100.0})

	id := 1
	balance := 100.0
	
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "accounts" WHERE "accounts"."id" = $1 ORDER BY "accounts"."id" LIMIT 1`)).WithArgs(id).
		WillReturnRows(sqlmock.NewRows([]string{"id","balance"}).AddRow(id,balance))

	res, err := s.repo.GetAccountInfo(id)
	
	require.NoError(s.T(),err)
	require.Equal(s.T(),balance,res)
}

func TestRepository(t *testing.T) {
	suite.Run(t,new(Suite))
}