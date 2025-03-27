package repository

const (
	UserTableName      = "user_"
	UserColumnId       = "id"
	UserColumnUsername = "username"
	UserColumnEmail    = "email"
	UserColumnPassword = "password"

	StatsTableName            = "stats_"
	StatsColumnId             = "id"
	StatsColumnIdUser         = "id_user"
	StatsColumnUtilizeCounter = "utilize_counter"
	StatsColumnReportCounter  = "report_counter"
	StatsColumnRating         = "rating"

	UserDataTableName        = "personal_data_"
	UserDataColumnId         = "id"
	UserDataColumnIdUser     = "id_user"
	UserDataColumnGivenName  = "given_name"
	UserDataColumnSurname    = "surname"
	UserDataColumnPatronymic = "patronymic"
	UserDataColumnCity       = "city"
	UserDataColumnBirthDate  = "birthdate"

	SessionTableName            = "session_"
	SessionColumnId             = "id_session"
	SessionColumnIdUser         = "id_user"
	SessionColumnExpirationTime = "expiration_time"

	RoleTableName   = "role_"
	RoleColumnId    = "id"
	RoleColumnTitle = "title"

	UserRoleTableName    = "user_role_"
	UserRoleColumnId     = "id"
	UserRoleColumnIdUser = "id_user"
	UserRoleColumnIdRole = "id_role"

	EventTableName    = "event_"
	EventColumnId     = "id"
	EventColumnStatus = "status_"

	EventStatusTableName   = "event_status_"
	EventStatusColumnId    = "id"
	EventStatusColumnTitle = "title"

	EventDataTableName          = "event_data"
	EventDataColumnEventId      = "event_id"
	EventDataColumnLatitude     = "latitude"
	EventDataColumnLongtitude   = "longtitude"
	EventDataColumnTitle        = "title"
	EventDataColumnDescription  = "description_"
	EventDataColumnCallerId     = "caller_id"
	EventDataColumnTimeCalled   = "time_called"
	EventDataColumnUtilizatorId = "utilizator_id"
	EventDataColumnTimeUtilized = "time_utilized"
)
