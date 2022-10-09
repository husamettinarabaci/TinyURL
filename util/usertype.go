package util

const (
	FREE    = "FREE"
	PREMIUM = "PREMIUM"
)

func IsSupportedUserType(userType string) bool {
	switch userType {
	case FREE, PREMIUM:
		return true
	}
	return false
}
