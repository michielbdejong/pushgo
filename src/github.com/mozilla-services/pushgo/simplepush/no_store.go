/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package simplepush

import (
	"fmt"
	"strings"
	"time"
)

type NoStoreConfig struct {
	MaxChannels int `toml:"max_channels" env:"max_channels"`
}

type NoStore struct {
	logger      *SimpleLogger
	maxChannels int
}

func (*NoStore) KeyToIDs(key string) (suaid, schid string, ok bool) {
	items := strings.SplitN(key, ".", 2)
	if len(items) < 2 {
		return "", "", false
	}
	return items[0], items[1], true
}

func (*NoStore) IDsToKey(suaid, schid string) (string, bool) {
	if len(suaid) == 0 || len(schid) == 0 {
		return "", false
	}
	return fmt.Sprintf("%s.%s", suaid, schid), true
}

func (*NoStore) ConfigStruct() interface{} {
	return &NoStoreConfig{
		MaxChannels: 200,
	}
}

func (n *NoStore) Init(app *Application, config interface{}) error {
	n.logger = app.Logger()
	n.maxChannels = config.(*NoStoreConfig).MaxChannels
	return nil
}

func (n *NoStore) MaxChannels() int                                     { return n.maxChannels }
func (*NoStore) Close() error                                           { return nil }
func (*NoStore) Status() (bool, error)                                  { return true, nil }
func (*NoStore) Exists(string) bool                                     { return false }
func (*NoStore) Register(string, string, int64) error                   { return nil }
func (*NoStore) Update(string, int64) error                             { return nil }
func (*NoStore) Unregister(string, string) error                        { return nil }
func (*NoStore) Drop(string, string) error                              { return nil }
func (*NoStore) FetchAll(string, time.Time) ([]Update, []string, error) { return nil, nil, nil }
func (*NoStore) DropAll(string) error                                   { return nil }
func (*NoStore) FetchPing(string) ([]byte, error)                       { return nil, nil }
func (*NoStore) PutPing(string, []byte) error                           { return nil }
func (*NoStore) DropPing(string) error                                  { return nil }
