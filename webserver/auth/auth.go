package auth

import "strconv"

type AuthResult struct {
	UserId int64
}

// TODO: 修改为真实的鉴权逻辑
func Auth(token string) (info *AuthResult, err error) {
	uid, err := strconv.ParseInt(token, 10, 64)
	if err != nil {
		return nil, err
	}
	return &AuthResult{
		UserId: uid,
	}, nil
}
