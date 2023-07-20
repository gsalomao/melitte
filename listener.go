// Copyright 2023 Gustavo Salomao
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package akira

import (
	"net"
	"sync"
)

// OnConnectionFunc is the function which the Listener must call when a new connection has been opened.
type OnConnectionFunc func(Listener, net.Conn)

// Listener is an interface which all network listeners must implement. A network listener is responsible for listen
// for network connections and notify any incoming connection.
type Listener interface {
	// Listen starts the listener. When the listener starts listening, it starts to accept any incoming connection,
	// and calls f with the new connection. If the listener fails to start listening, it returns the error.
	// This function does not block the caller and returns immediately after the listener is ready to accept incoming
	// connections.
	Listen(f OnConnectionFunc) error

	// Close closes the listener. Once the listener is closed, it does not accept any incoming connection. This
	// function blocks and returns only after the listener has closed.
	Close() error
}

type listeners struct {
	mutex    sync.RWMutex
	internal []Listener
}

func newListeners() *listeners {
	l := listeners{internal: make([]Listener, 0)}
	return &l
}

func (l *listeners) add(lsn Listener) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	l.internal = append(l.internal, lsn)
}

func (l *listeners) listenAll(f OnConnectionFunc) error {
	l.mutex.RLock()
	defer l.mutex.RUnlock()

	for _, lsn := range l.internal {
		err := lsn.Listen(f)
		if err != nil {
			return err
		}
	}

	return nil
}

func (l *listeners) closeAll() {
	l.mutex.RLock()
	defer l.mutex.RUnlock()

	for _, lsn := range l.internal {
		_ = lsn.Close()
	}
}
