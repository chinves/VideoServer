package defs


//requests
type UserCredential struct {
	Username string `json:"user_name"`
	Pwd string `json:"pwd"`
}

//response
type SignedUp struct {
	Success bool `json:"user_name"`
	SessionId string `json:"session_id"`
}
