package main

//  Copyright 2022- Emin Muhammadi and contributors
//
//  Licensed under the The GNU Affero General Public License,
//  Version 3.0 (the "License"); you may not use this file except
//  in compliance with the License. You may obtain a copy
//  of the License at
//
//     https://www.gnu.org/licenses/agpl-3.0.en.html
//
//  Unless required by applicable law or agreed to in writing,
//  software distributed under the License is distributed on an
//  "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the
//  specific language governing permissions and limitations
//  under the License.

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	cmd "github.com/eminmuhammadi/xtunnel/cmd"
	cli "github.com/urfave/cli/v2"
)

var (
	VERSION    = "0.0.1-development"
	BUILD_TIME = ""
	GO_VERSION = ""
)

var commands = []*cli.Command{
	cmd.Start(),
}

// Metadata for current release
const METADATA string = `{
    "name": "xtunnel",
	"description": "TCP/UDP reverse tunneling tool",
    "version": "%s",
    "author": {
        "name": "Emin Muhammadi",
        "email": "muemin.info@gmail.com"
    },
    "copyright": "(c) %d Emin Muhammadi and contributors",
    "license": "The GNU Affero General Public License"
}`

// Entry point for the application
func main() {
	// parse METADATA
	var md struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Version     string `json:"version"`
		Author      struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		} `json:"author"`
		Copyright string `json:"copyright"`
		License   string `json:"license"`
	}
	json.Unmarshal([]byte(METADATA), &md)

	app := &cli.App{
		Name:    md.Name,
		Usage:   md.Description,
		Version: fmt.Sprintf(md.Version, VERSION),
		Authors: []*cli.Author{{
			Name:  md.Author.Name,
			Email: md.Author.Email,
		}},
		Copyright: fmt.Sprintf(md.Copyright, time.Now().Year()),
		ExtraInfo: func() map[string]string {
			return map[string]string{
				"LICENSE":    md.License,
				"BUILD_TIME": BUILD_TIME,
				"GO_VERSION": GO_VERSION,
			}
		},
		Commands: commands,
		Suggest:  true,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
