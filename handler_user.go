package main

import "fmt"

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("login command excepts an username")
	}
	username := cmd.args[0]
	err := s.config.SetUser(username)
	if err != nil {
		return err
	}
	fmt.Printf("username has been set to: %s", username)
	return nil
}
