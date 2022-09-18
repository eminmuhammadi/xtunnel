package cmd

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
	"log"
	"strconv"
	"strings"

	pkg "github.com/eminmuhammadi/xtunnel/pkg"
	cli "github.com/urfave/cli/v2"
)

func Start() *cli.Command {
	return &cli.Command{
		Name:    "start",
		Aliases: []string{"s"},
		Usage:   "Forwards remote connections to local port",
		Flags: []cli.Flag{
			// -m, --master flag
			&cli.StringFlag{
				Name:     "master",
				Aliases:  []string{"m"},
				Usage:    "Master node",
				Required: true,
			},
			// -t, --target flag
			&cli.StringFlag{
				Name:     "target",
				Aliases:  []string{"t"},
				Usage:    "Target node",
				Required: true,
			},
			// -p, --protocol flag
			&cli.StringFlag{
				Name:     "protocol",
				Aliases:  []string{"p"},
				Usage:    "Protocols for tunneling. E.g \"tcp\", \"tcp4\", \"tcp6\", \"unix\" or \"unixpacket\"",
				Required: true,
			},
		},
		Action: func(ctx *cli.Context) error {
			log.Println("Starting tunnel...")

			master := ctx.String("master")
			log.Printf("Using %s as master node\n", master)

			target := ctx.String("target")
			log.Printf("Dialling %s, and using it as target node\n", target)

			protocol := ctx.String("protocol")
			log.Printf("Using %s as protocol\n", protocol)

			masterHost, masterPort := strings.Split(master, ":")[0], strings.Split(master, ":")[1]
			log.Printf("Master info: %s, port: %s\n", masterHost, masterPort)

			targetHost, targetPort := strings.Split(target, ":")[0], strings.Split(target, ":")[1]
			log.Printf("Target info: %s, port: %s\n", targetHost, targetPort)

			masterPortN, err := strconv.Atoi(masterPort)
			if err != nil {
				return err
			}

			targetPortN, err := strconv.Atoi(targetPort)
			if err != nil {
				return err
			}

			masterNode := pkg.NewNode(masterHost, masterPortN)
			targetNode := pkg.NewNode(targetHost, targetPortN)

			log.Println("Connection established")
			if err := pkg.CreateTunnel(protocol, masterNode, targetNode); err != nil {
				return err
			}

			return nil
		},
	}
}
