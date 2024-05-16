package domain

type AuthManager struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	Scope        string `json:"scope"`
	UserId       int    `json:"user_id"`
	RefreshToken string `json:"refresh_token"`
}

func (m AuthManager) GetAccessToken() string {
	return m.AccessToken
}
func (m AuthManager) SetAccessToken(accessToken string) {
	m.AccessToken = accessToken
}
