package constant

const (
	DB_PROVIDER_NAME       = "DBProvider"
	PASSWORD_PROVIDER_NAME = "PasswordProvider"
	TEMPORAL_SERVICE_NAME  = "TemporalService"

	API_GATEWAY_SERVICE_NAME           = "api-gateway-service"
	ANALYTICS_REPORTING_SERVICE_NAME   = "analytics-reporting-service"
	ENTERPRISE_MANAGEMENT_SERVICE_NAME = "enterprise-management-service"
	FILE_MANAGER_SERVICE_NAME          = "file-manager-service"
	HARDWARE_TAGGING_SERVICE_NAME      = "hardware-tagging-service"
	IAM_SERVICE_NAME                   = "iam-service"
	IMPORT_SERVICE_NAME                = "import-service"
	LANGUAGE_SERVICE_NAME              = "language-service"
	NOTIFICATION_SERVICE_NAME          = "notification-service"
	PLANNING_STRATEGY_SERVICE_NAME     = "planning-strategy-service"
	RESOURCE_ALLOCATION_SERVICE_NAME   = "resource-allocation-service"
)

type AnalyticsReportingActivityType string

const (
	AnalyticsReportingCreateChangeLogActivity     AnalyticsReportingActivityType = "CreateChangeLogActivity"
	AnalyticsReportingCreateSystemLogActivity     AnalyticsReportingActivityType = "CreateSystemLogActivity"
	AnalyticsReportingCreateUserActionLogActivity AnalyticsReportingActivityType = "CreateUserActionLogActivity"
	AnalyticsReportingCreateSessionLogActivity    AnalyticsReportingActivityType = "CreateSessionLogActivity"
	AnalyticsReportingDeleteChangeLogActivity     AnalyticsReportingActivityType = "DeleteChangeLogActivity"
	AnalyticsReportingDeleteSystemLogActivity     AnalyticsReportingActivityType = "DeleteSystemLogActivity"
	AnalyticsReportingDeleteUserActionLogActivity AnalyticsReportingActivityType = "DeleteUserActionLogActivity"
	AnalyticsReportingDeleteSessionLogActivity    AnalyticsReportingActivityType = "DeleteSessionLogActivity"
)
