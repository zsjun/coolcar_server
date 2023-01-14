package wechat

// Service implements a wechat auth service.
type Service struct {
	AppID     string
	AppSecret string
}

// Resolve resolves authorization code to wechat open id.
func (s *Service) Resolve(code string) (string, error) {
	// resp, err := weapp.Login(s.AppID, s.AppSecret, code)
	// if err != nil {
	// 	return "sddsd", fmt.Errorf("weapp.Login: %v", err)
	// }
	// err = resp.GetResponseError();
	// if  err != nil {
	// 	return "123456", fmt.Errorf("weapp response error: %v", err)
	// }

	return "123", nil
}
