package repository

const (
	UserTableName = "user_"

	UserIdColumn       = "id"
	UserUsernameColumn = "username"
	UserEmailColumn    = "email"
	UserPasswordColumn = "password"

	StatsTableName = "stats_"

	StatsIdColumn             = "id"
	StatsIdUserColumn         = "id_user"
	StatsUtilizeCounterColumn = "utilize_counter"
	StatsReportCounterColumn  = "report_counter"
	StatsRatingColumn         = "rating"

	DataTableName = "personal_data_"

	DataIdColumn         = "id"
	DataIdUserColumn     = "id_user"
	DataGivenNameColumn  = "given_name"
	DataSurnameColumn    = "surname"
	DataPatronymicColumn = "patronymic"
	DataCityColumn       = "city"
	DataBirthDateColumn  = "birthdate"

	SessionTableName = "session_"

	SessionIdColumn             = "id_session"
	SessionIdUserColumn         = "id_user"
	SessionExpirationTimeColumn = "expiration_time"

	RoleTableName = "role_"

	RoleIdColumn    = "id"
	RoleTitleColumn = "title"

	UserRoleTableName = "user_role_"

	UserRoleIdColumn     = "id"
	UserRoleIdUserColumn = "id_user"
	UserRoleIdRoleColumn = "id_role"
)
