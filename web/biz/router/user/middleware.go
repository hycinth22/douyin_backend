// Code generated by hertz generator.

package user

import (
	"github.com/cloudwego/hertz/pkg/app"
)

func rootMw() []app.HandlerFunc {
	// your status...
	return nil
}

func _douyinMw() []app.HandlerFunc {
	// your status...
	return nil
}

func _userMw() []app.HandlerFunc {
	// your status...
	return nil
}

func _user0Mw() []app.HandlerFunc {
	// your status...
	return []app.HandlerFunc{
		//jwt.JwtMiddleware.MiddlewareFunc(),
	}
}

func _loginMw() []app.HandlerFunc {
	// your status...
	return nil
}

func _login0Mw() []app.HandlerFunc {
	// your status...
	return []app.HandlerFunc{
		//jwt.JwtMiddleware.LoginHandler,
	}
}

func _registerMw() []app.HandlerFunc {
	// your status...
	return nil
}

func _register0Mw() []app.HandlerFunc {
	// your status...
	return nil
}
