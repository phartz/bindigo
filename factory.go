package main

func GetDataService(params map[string]string) IDataService {
	var service IDataService = nil

	switch params["dataService"] {
	case "testservice":
		service = TestService{}
	}

	service.SetCredentials(params["id"], params["credentials"])

	return service
}
