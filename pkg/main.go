package pkg

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
	"fmt"
	"io"
	"log"
	"net"
)

// Node
type Node struct {
	Host string
	Port int
}

// Tunnel
type Tunnel struct {
	Protocol string
	Listener net.Listener
}

// Creates a new node
func NewNode(host string, port int) *Node {
	return &Node{
		Host: host,
		Port: port,
	}
}

// Listen announces on the local network address.
func (node *Node) Listen(protocol string) (Tunnel, error) {
	listener, err := net.Listen(protocol, fmt.Sprintf("%s:%d", node.Host, node.Port))

	if err != nil {
		return Tunnel{}, err
	}

	return Tunnel{
		Protocol: protocol,
		Listener: listener,
	}, nil
}

// Dial connects to the address on the named network.
func (node *Node) Dial(protocol string) (net.Conn, error) {
	conn, err := net.Dial(protocol, fmt.Sprintf("%s:%d", node.Host, node.Port))

	if err != nil {
		return nil, err
	}

	return conn, nil
}

// Reverse data transfer
func Handshake(local net.Conn, remote net.Conn) {
	defer local.Close()
	defer remote.Close()

	done := make(chan bool)

	// Start remote -> local data transfer
	go func() {
		_, err := io.Copy(local, remote)
		if err != nil {
			log.Println(err)
		}
		done <- true
	}()

	// Start local -> remote data transfer
	go func() {
		_, err := io.Copy(remote, local)
		if err != nil {
			log.Println(err)
		}
		done <- true
	}()

	// Wait for data transfer to finish
	<-done
}

// Creates a new reverse tunnel
func CreateTunnel(protocol string, local *Node, remote *Node) error {
	// Listen on local
	tunnel, err := local.Listen(protocol)

	if err != nil {
		return err
	}

	defer tunnel.Listener.Close()

	for {
		// Accept connections from local
		local, err := tunnel.Listener.Accept()
		if err != nil {
			return err
		}

		// Dial to remote
		remote, err := remote.Dial(protocol)
		if err != nil {
			return err
		}

		// Start data transfer
		go Handshake(local, remote)
	}
}
