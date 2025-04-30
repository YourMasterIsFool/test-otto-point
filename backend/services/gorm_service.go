package services

import (
	usecase "backend/applications/usecase"
	postgresqlgorm "backend/infrastructure/postgresql_gorm"
	postgres_gorm_repository_impl "backend/infrastructure/postgresql_gorm/repository"
)

type GormService struct {
	BrandUsecase       usecase.BrandUsecase
	VoucherUsecase     usecase.VoucherUsecase
	TransactionUsecase usecase.TransactionUsecase
	AuthUsecase        usecase.AuthUsecase
}

func NewGormService() *GormService {

	// defining db
	dsn := "postgres://admin:password@localhost:5432/app_db?sslmode=disable"

	db, err := postgresqlgorm.NewDatabasePostgress(dsn)

	if err != nil {
		panic("Failed to connect database")
	}

	// auth

	authRepoImpl := postgres_gorm_repository_impl.NewAuthRepositoryImpl(db)
	authUsecase := usecase.NewAuthUsecase(authRepoImpl)

	//brand
	brandRepoImpl := postgres_gorm_repository_impl.NewBrandRepositoryImpl(db)
	brandUsecase := usecase.NewBrandUsecase(brandRepoImpl)

	//voucher
	voucherRepoImpl := postgres_gorm_repository_impl.NewVoucherRepositoryImpl(db)
	voucherUsecase := usecase.NewVoucherUsecase(voucherRepoImpl, brandUsecase)

	// transaction detail
	transactionDetailRepoImpl := postgres_gorm_repository_impl.NewTransactionDetailRepository(db)
	transactionDetailUsecase := usecase.NewTransactionDetailUsecase(transactionDetailRepoImpl)
	//transaction
	transactionRepoImpl := postgres_gorm_repository_impl.NewTransactionRepositoryImpl(db)
	transactionUsecase := usecase.NewTransactionUsecase(transactionRepoImpl, voucherUsecase, brandUsecase, transactionDetailUsecase)

	return &GormService{
		BrandUsecase:       brandUsecase,
		VoucherUsecase:     voucherUsecase,
		TransactionUsecase: transactionUsecase,
		AuthUsecase:        authUsecase,
	}
}
