package main

type IDataService interface {
	GetStatus() (int, error)
	Insert(string) error
	Exists(string) (bool, error)
	Delete(string) error
	SetCredentials(string, string)
}
