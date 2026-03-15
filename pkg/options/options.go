package options

import (
	"fmt"
	"os"
)

type Options struct {
	Token     string
	MachineID string
}

func FromEnv(skipMachine bool) (*Options, error) {
	opts := &Options{}

	token, err := requiredEnv("SPRITE_TOKEN")
	if err != nil {
		return nil, err
	}
	opts.Token = token

	if !skipMachine {
		machineID, err := requiredEnv("MACHINE_ID")
		if err != nil {
			return nil, err
		}
		opts.MachineID = machineID
	}

	return opts, nil
}

func requiredEnv(name string) (string, error) {
	val := os.Getenv(name)
	if val == "" {
		return "", fmt.Errorf("required option %s not found in environment", name)
	}
	return val, nil
}
