package utility

import (
	"database/sql"
	"regexp"
)

func ValidateEmail(email string) (bool, string) {
	if email == "" {
		return false, "email is required"
	}
	var re = regexp.MustCompile(`^[a-zA-Z0-9+_.-]+@[a-zA-Z0-9.-]+$`)
	if !re.MatchString(email) {
		return false, "not a valid email"
	}

	return true, ""
}

func CheckUserExists(email string, todoDB *sql.DB) (bool, string) {
	var emailFromDB string
	userExistsQuery := `SELECT email from users WHERE email=$1;`
	row := todoDB.QueryRow(userExistsQuery, email)
	err := row.Scan(&emailFromDB)
	switch err {
    case sql.ErrNoRows:
        return true, "user does not exist"
    case nil:
        return true, "user exist"
    default:
        return false, "Sorry something went wrong. Please try again later."
    }
}