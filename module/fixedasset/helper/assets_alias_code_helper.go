package helper

import (
	"context"
	"log"

	utildb "gitea.qwertysystem.net/BETS/ts-utils/db"
	"gitea.qwertysystem.net/BETS/ts-utils/entity"
	"gitea.qwertysystem.net/BETS/ts-utils/module/fixedasset/sql_builder"
	utilprovider "gitea.qwertysystem.net/BETS/ts-utils/provider"
	"gitea.qwertysystem.net/BETS/ts-utils/sql_query"
)

type AssetsAliasCodeHelper interface {
	UpdateAssetsAliasCode(
		ctx context.Context,
		serviceFactory utilprovider.ServiceProvider,
		companyCode string,
		filter map[string]sql_query.SQLCondition,
	) *entity.HttpError
}

type AssetsAliasCodeHelperImpl struct{}

func MakeAssetsAliasCodeHelperImpl() *AssetsAliasCodeHelperImpl {
	return &AssetsAliasCodeHelperImpl{}
}

func (helper *AssetsAliasCodeHelperImpl) UpdateAssetsAliasCode(
	ctx context.Context,
	serviceFactory utilprovider.ServiceProvider,
	companyCode string,
) *entity.HttpError {
	adminDb := utildb.FixedAssetDBName.CompanyDBName(companyCode)

	fixedAssetService := serviceFactory.MakePostgreSqlService(adminDb)

	querystring, args, err := sql_builder.UpdateAssetAliasCodeBuilder()

	_, err = fixedAssetService.UpdateMany(ctx, querystring, args...)
	if err != nil {
		log.Println(err.Error())
		return entity.ToHttpError(err)
	}

	return nil
}
