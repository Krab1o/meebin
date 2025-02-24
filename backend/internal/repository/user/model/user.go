package repoModel

import "time"

type User struct {
	Id 			int64
	Email 		string
	Password	string
	PersonalData struct {
		FirstName 	string
		SurName		string
		LastName	string
		City		string
		Birthdate	time.Time
	}
	Stats struct {
		UtilizeCount	int64
		ReportCount		int64
		Rating			float64
	}
}