package constants

import "os"

var (
	IsGitpod           = os.Getenv("GITPOD_WORKSPACE_ID") != ""
	IsCodespace        = os.Getenv("CODESPACE_NAME") != "" && os.Getenv("CLOUDENV_ENVIRONMENT_ID") != ""
	IsGoogleCloud      = os.Getenv("DEVSHELL_GCLOUD_CONFIG") != "" || os.Getenv("BASHRC_GOOGLE_PATH") != ""
	IsCloudEnvironment = IsGitpod || IsCodespace || IsGoogleCloud
	TempFileName       = "temp_file.txt"
)
