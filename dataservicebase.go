package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	cfenv "github.com/cloudfoundry-community/go-cfenv"
)

type DataServiceBase struct {
	credentials map[string]interface{}
}

func (d *DataServiceBase) SetCredentialsImpl(id string, credentialsParam string) error {
	if len(credentialsParam) > 0 {
		err := json.Unmarshal([]byte(credentialsParam), &d.credentials)
		if err != nil {
			return err
		}
		return nil
	}

	// In any case of an exception recover to not stop the app
	defer fmt.Errorf("Could not parse VCAP_SERVICES.")

	// read the env variables
	appEnv, err := cfenv.Current()
	if err != nil {
		return err
	}

	// convert string to i
	i, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	var service cfenv.Service

	// get first service and than the index
	for _, v := range appEnv.Services {
		service = v[i]
		break
	}

	// set cedentials
	d.credentials = service.Credentials

	return nil
}
