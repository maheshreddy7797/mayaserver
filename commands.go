package main

import (
	"os"

	"github.com/maheshreddy7797/mayaserver/cmd"
	"github.com/mitchellh/cli"
)

// Commands returns the mapping of CLI commands for Maya server. The meta
// parameter lets you set meta options for all commands.
func Commands(metaPtr *cmd.Meta) map[string]cli.CommandFactory {
	if metaPtr == nil {
		metaPtr = new(cmd.Meta)
	}

	meta := *metaPtr
	if meta.Ui == nil {
		meta.Ui = &cli.BasicUi{
			Reader:      os.Stdin,
			Writer:      os.Stdout,
			ErrorWriter: os.Stderr,
		}
	}

	return map[string]cli.CommandFactory{
		"up": func() (cli.Command, error) {
			return &cmd.UpCommand{
				Revision:          GitCommit,
				Version:           Version,
				VersionPrerelease: VersionPrerelease,
				Ui:                meta.Ui,
				ShutdownCh:        make(chan struct{}),
			}, nil
		},
		"version": func() (cli.Command, error) {
			ver := Version
			rel := VersionPrerelease
			if GitDescribe != "" {
				ver = GitDescribe
				// Trim off a leading 'v', we append it anyways.
				if ver[0] == 'v' {
					ver = ver[1:]
				}
			}
			if GitDescribe == "" && rel == "" && VersionPrerelease != "" {
				rel = "dev"
			}

			return &cmd.VersionCommand{
				Revision:          GitCommit,
				Version:           ver,
				VersionPrerelease: rel,
				Ui:                meta.Ui,
			}, nil
		},
	}
}
