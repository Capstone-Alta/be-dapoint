package request

import "dapoint-api/service/auth/spec"

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (req *AuthRequest) ToSpec() *spec.InputLogin {
	return &spec.InputLogin{
		Username: req.Username,
		Password: req.Password,
	}
}
