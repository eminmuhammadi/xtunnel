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

func Forward() *cli.Command {
	return &cli.Command{
		Name:    "forward",
		Aliases: []string{"f"},
		Usage:   "Forwards remote connections to local port",
		Flags: []cli.Flag{
			// -r, --remote flag
			&cli.StringFlag{
				Name:     "remote",
				Aliases:  []string{"r"},
				Usage:    "Remote node",
				Required: true,
			},
			// -l, --local flag
			&cli.StringFlag{
				Name:     "local",
				Aliases:  []string{"l"},
				Usage:    "Local node",
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

			local := ctx.String("local")
			log.Printf("Using %s as local node\n", local)

			remote := ctx.String("remote")
			log.Printf("Dialling %s, and using it as remote node\n", remote)

			protocol := ctx.String("protocol")
			log.Printf("Using %s as protocol\n", protocol)

			localHost, localPort := strings.Split(local, ":")[0], strings.Split(local, ":")[1]
			log.Printf("Local info: %s, port: %s\n", localHost, localPort)

			remoteHost, remotePort := strings.Split(remote, ":")[0], strings.Split(remote, ":")[1]
			log.Printf("Remote info: %s, port: %s\n", remoteHost, remotePort)

			localPortN, err := strconv.Atoi(localPort)
			if err != nil {
				return err
			}

			remotePortN, err := strconv.Atoi(remotePort)
			if err != nil {
				return err
			}

			localNode := pkg.NewNode(localHost, localPortN)
			remoteNode := pkg.NewNode(remoteHost, remotePortN)

			log.Println("Connection established")
			if err := pkg.CreateTunnel(protocol, localNode, remoteNode); err != nil {
				return err
			}

			return nil
		},
	}
}
