package helper

import (
	"context"
	"log"

	"gitea.qwertysystem.net/BETS/ts-utils/entity"
	"gitea.qwertysystem.net/BETS/ts-utils/module/fixedasset/sql_builder"
	"gitea.qwertysystem.net/BETS/ts-utils/service"
)

type AssetsAliasCodeHelper interface {

	// The function builds and executes a SQL query to bulk-update the 'alias_code'
	// field for all assets
	//
	// Parameters:
	//   - ctx: The context for managing the request lifecycle, including cancellations and timeouts.
	//   - fixedAssetService: The database service for fixed asset db
	//   - companyCode: client id string
	//
	// Returns:
	//   - An *entity.HttpError if the database operation fails, otherwise it returns nil
	//     to indicate success.
	UpdateAssetsAliasCode(
		ctx context.Context,
		fixedAssetService service.PostgreSqlService,
		companyCode string,
	) *entity.HttpError
}

type AssetsAliasCodeHelperImpl struct{}

func MakeAssetsAliasCodeHelperImpl() *AssetsAliasCodeHelperImpl {
	return &AssetsAliasCodeHelperImpl{}
}

func (helper *AssetsAliasCodeHelperImpl) UpdateAssetsAliasCode(
	ctx context.Context,
	fixedAssetService service.PostgreSqlService,
	companyCode string,
) *entity.HttpError {
	querystring, args, err := sql_builder.UpdateAssetAliasCodeBuilder()

	_, err = fixedAssetService.UpdateMany(ctx, querystring, args...)
	if err != nil {
		log.Println(err.Error())
		return entity.ToHttpError(err)
	}

	return nil
}
