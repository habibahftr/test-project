package server_config

import (
	config2 "test/config"
)

func NewServerAttribute() serverAttribute {
	return serverAttribute{}
}

func (s *serverAttribute) Init() (err error) {
	s.DBConnection = config2.ConnectDB()
	defer func() {
		if err != nil {
			err = s.DBConnection.Close()
			if err != nil {
				return
			}
		}

	}()

	s.InitDao()

	return err
}
