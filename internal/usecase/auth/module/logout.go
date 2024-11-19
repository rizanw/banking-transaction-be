package module

func (u *usecase) Logout(token string) error {
	return u.sessionMgr.BlacklistToken(token)
}
