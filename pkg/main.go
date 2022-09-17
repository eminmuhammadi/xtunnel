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
	"net"
)

// Node
type Node struct {
	Host string
	Port int
}

type RequestHandler func(conn net.Conn) ([]byte, error)
type ResponseHandler func(protocol string, target Node, request []byte) ([]byte, error)

// Tunnel
type Tunnel struct {
	Protocol string
	Listener net.Listener
	Target   *Node
}

// Creates a new node
func NewNode(host string, port int) *Node {
	return &Node{
		Host: host,
		Port: port,
	}
}

// Listen announces on the local network address.
func (node *Node) Forward(protocol string, target *Node) (Tunnel, error) {
	listener, err := net.Listen(protocol, fmt.Sprintf("%s:%d", node.Host, node.Port))

	if err != nil {
		return Tunnel{}, err
	}

	return Tunnel{
		Protocol: protocol,
		Listener: listener,
		Target:   target,
	}, nil
}

// Accept waits for and returns the next connection to the listener.
func (tunnel *Tunnel) Expose(requestHandler RequestHandler, responseHandler ResponseHandler) error {
	for {
		conn, err := tunnel.Listener.Accept()

		if err != nil {
			return err
		}

		go func() {
			defer conn.Close()

			for {
				request, err := requestHandler(conn)

				if err != nil {
					break
				}

				response, err := responseHandler(tunnel.Protocol, *tunnel.Target, request)

				if err != nil {
					break
				}

				conn.Write(response)
			}
		}()
	}
}

// Dial connects to the address on the named network.
func (node *Node) Dial(protocol string) (net.Conn, error) {
	conn, err := net.Dial(protocol, fmt.Sprintf("%s:%d", node.Host, node.Port))

	if err != nil {
		return nil, err
	}

	return conn, nil
}

// Request
func Request(conn net.Conn) ([]byte, error) {
	request := make([]byte, 1024)
	_, err := conn.Read(request)

	if err != nil {
		return nil, err
	}

	return request, nil
}

// Response
func Response(protocol string, target Node, request []byte) ([]byte, error) {
	conn, err := target.Dial(protocol)

	if err != nil {
		return nil, err
	}

	defer conn.Close()

	conn.Write(request)

	response := make([]byte, 1024)
	_, err = conn.Read(response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

// Create a tunnel
func CreateTunnel(protocol string, master *Node, target *Node) error {
	listener, err := master.Forward(protocol, target)
	if err != nil {
		return err
	}

	if err := listener.Expose(Request, Response); err != nil {
		return err
	}

	return nil
}
