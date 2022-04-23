package web

import "github.com/gin-contrib/sessions"

func getSession(session sessions.Session) int64 {
	value := session.Get(Key)
	if value == nil {
		return 0
	}

	return value.(int64)
}

func setSession(session sessions.Session, id int64) error {
	cookie := sessions.Options{
		MaxAge:   3600 * 24,
		Path:     "/",
		HttpOnly: true,
	}
	session.Options(cookie)
	session.Set(Key, id)

	return session.Save()
}

func deleteSession(session sessions.Session) error {
	session.Clear()
	session.Options(sessions.Options{MaxAge: -1})

	return session.Save()
}
