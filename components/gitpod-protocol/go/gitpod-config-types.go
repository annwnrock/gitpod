// Copyright (c) 2020 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

// Code generated by schema-generate. DO NOT EDIT.

package protocol

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

// Env Environment variables to set.
type Env struct {
}

// Github Configures Gitpod's GitHub app
type Github struct {

	// Set to true to enable workspace prebuilds, false to disable them. Defaults to true.
	Prebuilds interface{} `yaml:"prebuilds,omitempty"`
}

// GitpodConfig
type GitpodConfig struct {

	// Path to where the repository should be checked out.
	CheckoutLocation string `yaml:"checkoutLocation,omitempty"`

	// Git config values should be provided in pairs. E.g. `core.autocrlf: input`. See https://git-scm.com/docs/git-config#_values.
	GitConfig map[string]string `yaml:"gitConfig,omitempty"`

	// Configures Gitpod's GitHub app
	Github *Github `yaml:"github,omitempty"`

	// Controls what ide should be used for a workspace.
	Ide interface{} `yaml:"ide,omitempty"`

	// The Docker image to run your workspace in.
	Image interface{} `yaml:"image,omitempty"`

	// List of exposed ports.
	Ports []*PortsItems `yaml:"ports,omitempty"`

	// Whether the workspace is started in privileged mode.
	Privileged bool `yaml:"privileged,omitempty"`

	// List of tasks to run on start. Each task will open a terminal in the IDE.
	Tasks []*TasksItems `yaml:"tasks,omitempty"`

	// Configure VS Code integration
	Vscode *Vscode `yaml:"vscode,omitempty"`

	// Path to where the IDE's workspace should be opened.
	WorkspaceLocation string `yaml:"workspaceLocation,omitempty"`
}

// Image_object The Docker image to run your workspace in.
type Image_object struct {

	// Relative path to the context path (optional). Should only be set if you need to copy files into the image.
	Context string `yaml:"context,omitempty"`

	// Relative path to a docker file.
	File string `yaml:"file"`
}

// PortsItems
type PortsItems struct {

	// Port name (deprecated).
	Name string `yaml:"name,omitempty"`

	// What to do when a service on this port was detected. 'notify' (default) will show a notification asking the user what to do. 'open-browser' will open a new browser tab. 'open-preview' will open in the preview on the right of the IDE. 'ignore' will do nothing.
	OnOpen string `yaml:"onOpen,omitempty"`

	// The port number (e.g. 1337) or range (e.g. 3000-3999) to expose.
	Port interface{} `yaml:"port"`

	// The protocol to be used. (deprecated)
	Protocol string `yaml:"protocol,omitempty"`

	// Whether the port visibility should be private or public. 'public' (default) will allow everyone with the port URL to access the port. 'private' will only allow users with workspace access to access the port.
	Visibility string `yaml:"visibility,omitempty"`
}

// Prebuilds_object Set to true to enable workspace prebuilds, false to disable them. Defaults to true.
type Prebuilds_object struct {

	// Add a Review in Gitpod badge to pull requests. Defaults to true.
	AddBadge bool `yaml:"addBadge,omitempty"`

	// Add a label to a PR when it's prebuilt. Set to true to use the default label (prebuilt-in-gitpod) or set to a string to use a different label name. This is a beta feature and may be unreliable. Defaults to false.
	AddLabel interface{} `yaml:"addLabel,omitempty"`

	// Enable prebuilds for all branches. Defaults to false.
	Branches bool `yaml:"branches,omitempty"`

	// Enable prebuilds for the default branch (typically master). Defaults to true.
	Master bool `yaml:"master,omitempty"`

	// Enable prebuilds for pull-requests from the original repo. Defaults to true.
	PullRequests bool `yaml:"pullRequests,omitempty"`

	// Enable prebuilds for pull-requests from any repo (e.g. from forks). Defaults to false.
	PullRequestsFromForks bool `yaml:"pullRequestsFromForks,omitempty"`
}

// TasksItems
type TasksItems struct {

	// A shell command to run before `init` and the main `command`. This command is executed on every start and is expected to terminate. If it fails, the following commands will not be executed.
	Before string `yaml:"before,omitempty" json:"before,omitempty"`

	// The main shell command to run after `before` and `init`. This command is executed last on every start and doesn't have to terminate.
	Command string `yaml:"command,omitempty" json:"command,omitempty"`

	// Environment variables to set.
	Env *Env `yaml:"env,omitempty" json:"env,omitempty"`

	// A shell command to run between `before` and the main `command`. This command is executed only on after initializing a workspace with a fresh clone, but not on restarts and snapshots. This command is expected to terminate. If it fails, the `command` property will not be executed.
	Init string `yaml:"init,omitempty" json:"init,omitempty"`

	// Name of the task. Shown on the tab of the opened terminal.
	Name string `yaml:"name,omitempty" json:"name,omitempty"`

	// The panel/area where to open the terminal. Default is 'bottom' panel.
	OpenIn string `yaml:"openIn,omitempty" json:"openIn,omitempty"`

	// The opening mode. Default is 'tab-after'.
	OpenMode string `yaml:"openMode,omitempty" json:"openMode,omitempty"`

	// A shell command to run after `before`. This command is executed only on during workspace prebuilds. This command is expected to terminate. If it fails, the workspace build fails.
	Prebuild string `yaml:"prebuild,omitempty" json:"prebuild,omitempty"`
}

// Vscode Configure VS Code integration
type Vscode struct {

	// List of extensions which should be installed for users of this workspace. The identifier of an extension is always '${publisher}.${name}'. For example: 'vscode.csharp'.
	Extensions []string `yaml:"extensions,omitempty"`
}

func (strct *Github) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "prebuilds" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"prebuilds\": ")
	if tmp, err := json.Marshal(strct.Prebuilds); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Github) UnmarshalJSON(b []byte) error {
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "prebuilds":
			if err := json.Unmarshal([]byte(v), &strct.Prebuilds); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	return nil
}

func (strct *GitpodConfig) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "checkoutLocation" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"checkoutLocation\": ")
	if tmp, err := json.Marshal(strct.CheckoutLocation); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "gitConfig" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"gitConfig\": ")
	if tmp, err := json.Marshal(strct.GitConfig); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "github" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"github\": ")
	if tmp, err := json.Marshal(strct.Github); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "ide" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"ide\": ")
	if tmp, err := json.Marshal(strct.Ide); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "image" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"image\": ")
	if tmp, err := json.Marshal(strct.Image); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "ports" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"ports\": ")
	if tmp, err := json.Marshal(strct.Ports); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "privileged" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"privileged\": ")
	if tmp, err := json.Marshal(strct.Privileged); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "tasks" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"tasks\": ")
	if tmp, err := json.Marshal(strct.Tasks); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "vscode" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"vscode\": ")
	if tmp, err := json.Marshal(strct.Vscode); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "workspaceLocation" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"workspaceLocation\": ")
	if tmp, err := json.Marshal(strct.WorkspaceLocation); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *GitpodConfig) UnmarshalJSON(b []byte) error {
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "checkoutLocation":
			if err := json.Unmarshal([]byte(v), &strct.CheckoutLocation); err != nil {
				return err
			}
		case "gitConfig":
			if err := json.Unmarshal([]byte(v), &strct.GitConfig); err != nil {
				return err
			}
		case "github":
			if err := json.Unmarshal([]byte(v), &strct.Github); err != nil {
				return err
			}
		case "ide":
			if err := json.Unmarshal([]byte(v), &strct.Ide); err != nil {
				return err
			}
		case "image":
			if err := json.Unmarshal([]byte(v), &strct.Image); err != nil {
				return err
			}
		case "ports":
			if err := json.Unmarshal([]byte(v), &strct.Ports); err != nil {
				return err
			}
		case "privileged":
			if err := json.Unmarshal([]byte(v), &strct.Privileged); err != nil {
				return err
			}
		case "tasks":
			if err := json.Unmarshal([]byte(v), &strct.Tasks); err != nil {
				return err
			}
		case "vscode":
			if err := json.Unmarshal([]byte(v), &strct.Vscode); err != nil {
				return err
			}
		case "workspaceLocation":
			if err := json.Unmarshal([]byte(v), &strct.WorkspaceLocation); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	return nil
}

func (strct *Image_object) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "context" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"context\": ")
	if tmp, err := json.Marshal(strct.Context); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "File" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "file" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"file\": ")
	if tmp, err := json.Marshal(strct.File); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Image_object) UnmarshalJSON(b []byte) error {
	fileReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "context":
			if err := json.Unmarshal([]byte(v), &strct.Context); err != nil {
				return err
			}
		case "file":
			if err := json.Unmarshal([]byte(v), &strct.File); err != nil {
				return err
			}
			fileReceived = true
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	// check if file (a required property) was received
	if !fileReceived {
		return errors.New("\"file\" is required but was not present")
	}
	return nil
}

func (strct *PortsItems) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "name" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"name\": ")
	if tmp, err := json.Marshal(strct.Name); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "onOpen" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"onOpen\": ")
	if tmp, err := json.Marshal(strct.OnOpen); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Port" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "port" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"port\": ")
	if tmp, err := json.Marshal(strct.Port); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "protocol" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"protocol\": ")
	if tmp, err := json.Marshal(strct.Protocol); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "visibility" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"visibility\": ")
	if tmp, err := json.Marshal(strct.Visibility); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *PortsItems) UnmarshalJSON(b []byte) error {
	portReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "name":
			if err := json.Unmarshal([]byte(v), &strct.Name); err != nil {
				return err
			}
		case "onOpen":
			if err := json.Unmarshal([]byte(v), &strct.OnOpen); err != nil {
				return err
			}
		case "port":
			if err := json.Unmarshal([]byte(v), &strct.Port); err != nil {
				return err
			}
			portReceived = true
		case "protocol":
			if err := json.Unmarshal([]byte(v), &strct.Protocol); err != nil {
				return err
			}
		case "visibility":
			if err := json.Unmarshal([]byte(v), &strct.Visibility); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	// check if port (a required property) was received
	if !portReceived {
		return errors.New("\"port\" is required but was not present")
	}
	return nil
}

func (strct *TasksItems) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "before" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"before\": ")
	if tmp, err := json.Marshal(strct.Before); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "command" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"command\": ")
	if tmp, err := json.Marshal(strct.Command); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "env" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"env\": ")
	if tmp, err := json.Marshal(strct.Env); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "init" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"init\": ")
	if tmp, err := json.Marshal(strct.Init); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "name" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"name\": ")
	if tmp, err := json.Marshal(strct.Name); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "openIn" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"openIn\": ")
	if tmp, err := json.Marshal(strct.OpenIn); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "openMode" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"openMode\": ")
	if tmp, err := json.Marshal(strct.OpenMode); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "prebuild" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"prebuild\": ")
	if tmp, err := json.Marshal(strct.Prebuild); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *TasksItems) UnmarshalJSON(b []byte) error {
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "before":
			if err := json.Unmarshal([]byte(v), &strct.Before); err != nil {
				return err
			}
		case "command":
			if err := json.Unmarshal([]byte(v), &strct.Command); err != nil {
				return err
			}
		case "env":
			if err := json.Unmarshal([]byte(v), &strct.Env); err != nil {
				return err
			}
		case "init":
			if err := json.Unmarshal([]byte(v), &strct.Init); err != nil {
				return err
			}
		case "name":
			if err := json.Unmarshal([]byte(v), &strct.Name); err != nil {
				return err
			}
		case "openIn":
			if err := json.Unmarshal([]byte(v), &strct.OpenIn); err != nil {
				return err
			}
		case "openMode":
			if err := json.Unmarshal([]byte(v), &strct.OpenMode); err != nil {
				return err
			}
		case "prebuild":
			if err := json.Unmarshal([]byte(v), &strct.Prebuild); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	return nil
}

func (strct *Vscode) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "extensions" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"extensions\": ")
	if tmp, err := json.Marshal(strct.Extensions); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Vscode) UnmarshalJSON(b []byte) error {
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "extensions":
			if err := json.Unmarshal([]byte(v), &strct.Extensions); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	return nil
}
