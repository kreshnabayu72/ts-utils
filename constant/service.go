package constant

const (
	DB_PROVIDER_NAME       = "DBProvider"
	PASSWORD_PROVIDER_NAME = "PasswordProvider"
	TEMPORAL_SERVICE_NAME  = "TemporalService"

	API_GATEWAY_SERVICE_NAME           = "api-gateway-service"
	ANALYTICS_REPORTING_SERVICE_NAME   = "analytics-reporting-service"
	APPROVAL_SERVICE_NAME              = "approval-service"
	ASSET_ADDON_SERVICE_NAME           = "asset-addon-service"
	ASSET_TRANSACTION_SERVICE_NAME     = "asset-transaction-service"
	ENTERPRISE_MANAGEMENT_SERVICE_NAME = "enterprise-management-service"
	FILE_MANAGER_SERVICE_NAME          = "file-manager-service"
	HARDWARE_TAGGING_SERVICE_NAME      = "hardware-tagging-service"
	IAM_SERVICE_NAME                   = "iam-service"
	IMPORT_SERVICE_NAME                = "import-service"
	LANGUAGE_TRANSLATION_SERVICE_NAME  = "language-translation-service"
	MASTER_ASSET_REGISTRY_SERVICE_NAME = "master-asset-registry-service"
	NOTIFICATION_SERVICE_NAME          = "notification-service"
	PLANNING_STRATEGY_SERVICE_NAME     = "planning-strategy-service"
	RESOURCE_ALLOCATION_SERVICE_NAME   = "resource-allocation-service"
)

// Not used anywhere but stored as knowledge
const (
	IAM_SERVICE_PORT                   = 5432
	HARDWARE_TAGGING_SERVICE_PORT      = 5433
	ENTERPRISE_MANAGEMENT_SERVICE_PORT = 5434
	MASTER_ASSET_REGISTRY_SERVICE_PORT = 5435
	ANALYTICS_REPORTING_SERVICE_PORT   = 5436
	APPROVAL_SERVICE_PORT              = 5437
	FILE_MANAGER_SERVICE_PORT          = 5438
	IMPORT_SERVICE_PORT                = 5439
	LANGUAGE_TRANSLATION_SERVICE_PORT  = 5440
	NOTIFICATION_SERVICE_PORT          = 5441
	PLANNING_STRATEGY_SERVICE_PORT     = 5442
	RESOURCE_ALLOCATION_SERVICE_PORT   = 5443
	ASSET_ADDON_SERVICE_PORT           = 5444
	ASSET_TRANSACTION_SERVICE_PORT     = 5445
)

const (
	GetLoggedInUserActivity = "GetLoggedInUserActivity"

	AnalyticsReportingCreateChangeLogActivity              = "CreateChangeLogActivity"
	AnalyticsReportingDeleteChangeLogActivity              = "DeleteChangeLogActivity"
	AnalyticsReportingCreateSessionLogActivity             = "CreateSessionLogActivity"
	AnalyticsReportingDeleteSessionLogActivity             = "DeleteSessionLogActivity"
	AnalyticsReportingCreateHardwareChangeLogActivity      = "CreateHardwareChangeLogActivity"
	AnalyticsReportingDeleteHardwareChangeLogActivity      = "DeleteHardwareChangeLogActivity"
	AnalyticsReportingCreateIOTReaderUsageLogActivity      = "CreateIOTReaderUsageLogActivity"
	AnalyticsReportingDeleteIOTReaderUsageLogActivity      = "DeleteIOTReaderUsageLogActivity"
	AnalyticsReportingCreateHandheldReaderUsageLogActivity = "CreateHandheldReaderUsageLogActivity"
	AnalyticsReportingDeleteHandheldReaderUsageLogActivity = "DeleteHandheldReaderUsageLogActivity"
	AnalyticsReportingCreateReaderTaggingLogActivity       = "CreateReaderTaggingLogActivity"
	AnalyticsReportingDeleteReaderTaggingLogActivity       = "DeleteReaderTaggingLogActivity"
	AnalyticsReportingCreateReaderUserTagLogActivity       = "CreateReaderUserTagLogActivity"
	AnalyticsReportingDeleteReaderUserTagLogActivity       = "DeleteReaderUserTagLogActivity"
	AnalyticsReportingCreateTagEventLogActivity            = "CreateTagEventLogActivity"
	AnalyticsReportingDeleteTagEventLogActivity            = "DeleteTagEventLogActivity"
	AnalyticsReportingCreateTagEventLogAuditActivity       = "CreateTagEventLogAuditActivity"
	AnalyticsReportingDeleteTagEventLogAuditActivity       = "DeleteTagEventLogAuditActivity"

	AssetTransactionGetUsersBorrowedAssetCountActivity = "GetUsersBorrowedAssetCountActivity"
	AssetTransactionGetUsersAssignedAssetCountActivity = "GetUsersAssignedAssetCountActivity"

	ApprovalGetUsersPendingApprovalCountActivity = "GetUsersPendingApprovalCountActivity"
)
