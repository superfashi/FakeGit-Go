package fakegit

import "errors"

const (
	ARGUMENT_ERROR_INVALID  = errors.New("Command excuted with inappropriate argument.")
	ARGUMENT_ERROR_USERNAME = errors.New("Username needed.")
	GITCONF_FILE_NOT_FOUND  = errors.New("No git config file found, make sure you are under a git repository folder.")
	GITHUB_USER_ERROR       = errors.New("No email found for specific user.")
)
