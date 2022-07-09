package voucher

import (
	"dapoint-api/entities"
	"dapoint-api/util"
)

func RepositoryFactory(dbCon *util.DatabaseConnection) entities.VoucherRepository {
	var voucherRepo entities.VoucherRepository

	if dbCon.Driver == util.Postgres {
		// existing tetep jalan
		voucherRepo = NewPostgresRepository(dbCon.Postgres)
	}

	return voucherRepo
}
