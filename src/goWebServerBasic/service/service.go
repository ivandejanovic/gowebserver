package service

import ()

func Method(key string) (string, error) {
	value := "Your key is: " + key
	return value, nil
}
