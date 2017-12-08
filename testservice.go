package main

type TestService struct {
	DataServiceBase
}

func (t TestService) GetStatus() (output int, err error) {
	return 1, nil
}

func (t TestService) SetCredentials(id string, credentials string) {
	t.SetCredentialsImpl(id, credentials)
}

func (t TestService) Insert(value string) (err error) {
	return nil
}

func (t TestService) Exists(value string) (exists bool, err error) {
	return true, nil
}
func (t TestService) Delete(value string) (err error) {
	return nil
}
