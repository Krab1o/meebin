package repository

// TODO: refactor column names
const (
	UserTableName      = "user_"
	UserIdColumn       = "id"
	UserUsernameColumn = "username"
	UserEmailColumn    = "email"
	UserPasswordColumn = "password"

	StatsTableName            = "stats_"
	StatsIdColumn             = "id"
	StatsIdUserColumn         = "id_user"
	StatsUtilizeCounterColumn = "utilize_counter"
	StatsReportCounterColumn  = "report_counter"
	StatsRatingColumn         = "rating"

	UserDataTableName        = "personal_data_"
	UserDataIdColumn         = "id"
	UserDataIdUserColumn     = "id_user"
	UserDataGivenNameColumn  = "given_name"
	UserDataSurnameColumn    = "surname"
	UserDataPatronymicColumn = "patronymic"
	UserDataCityColumn       = "city"
	UserDataBirthDateColumn  = "birthdate"

	SessionTableName            = "session_"
	SessionIdColumn             = "id_session"
	SessionIdUserColumn         = "id_user"
	SessionExpirationTimeColumn = "expiration_time"

	RoleTableName   = "role_"
	RoleIdColumn    = "id"
	RoleTitleColumn = "title"

	UserRoleTableName    = "user_role_"
	UserRoleIdColumn     = "id"
	UserRoleIdUserColumn = "id_user"
	UserRoleIdRoleColumn = "id_role"

	EventTableName              = "event_"
	EventIdColumn               = "id"
	EventDataCallerIdColumn     = "caller_id"
	EventDataUtilizatorIdColumn = "utilizator_id"
	EventStatusColumn           = "status"

	EventStatusTableName   = "event_status_"
	EventStatusIdColumn    = "id"
	EventStatusTitleColumn = "title"

	EventDataTableName         = "event_data"
	EventDataEventIdColumn     = "event_id"
	EventDataLatitudeColumn    = "latitude"
	EventDataLongtitudeColumn  = "longtitude"
	EventDataTitleColumn       = "title"
	EventDataDescriptionColumn = "description"
	EventDataTimeCalledColumn  = "time_called"
	EventDataTimeCleanedColumn = "time_cleaned"
)
