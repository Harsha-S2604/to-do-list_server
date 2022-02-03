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

func CheckUserExists(email string, todoDB *sql.DB) (bool, string, int) {
	var userId int
	userExistsQuery := `SELECT user_id from users WHERE email=$1;`
	row := todoDB.QueryRow(userExistsQuery, email)
	err := row.Scan(&userId)
	switch err {
    case sql.ErrNoRows:
        return true, "user does not exist", 0
    case nil:
        return true, "user exist", userId
    default:
        return false, "Sorry something went wrong. Please try again later.", 0
    }
}