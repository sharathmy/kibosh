// kibosh
//
// Copyright (c) 2017-Present Pivotal Software, Inc. All Rights Reserved.
//
// This program and the accompanying materials are made available under the terms of the under the Apache License,
// Version 2.0 (the "License”); you may not use this file except in compliance with the License. You may
// obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the
// License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing permissions and
// limitations under the License.

package config

import (
	"github.com/kelseyhightower/envconfig"

	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

type ClusterCredentials struct {
	CADataRaw string `envconfig:"CA_DATA" json:"caDataRaw"`
	CAData    []byte
	Server    string `envconfig:"SERVER" json:"server"`
	Token     string `envconfig:"TOKEN" json:"token"`
}

type RegistryConfig struct {
	Server string `envconfig:"REG_SERVER"`
	User   string `envconfig:"REG_USER"`
	Pass   string `envconfig:"REG_PASS"`
	Email  string `envconfig:"REG_EMAIL"`
}

type CFClientConfig struct {
	BrokerURL         string `envconfig:"CF_BROKER_URL"`
	BrokerName        string `envconfig:"CF_BROKER_NAME"`
	ApiAddress        string `envconfig:"CF_API_ADDRESS"`
	Username          string `envconfig:"CF_USERNAME"`
	Password          string `envconfig:"CF_PASSWORD"`
	SkipSslValidation bool   `envconfig:"CF_SKIP_SSL_VALIDATION"`
}

type TillerTLSConfig struct {
	TLSKeyFile    string `envconfig:"TILLER_TLS_KEY_FILE"`
	TLSCertFile   string `envconfig:"TILLER_CERT_FILE"`
	TLSCaCertFile string `envconfig:"TILLER_TLS_CA_CERT_FILE"`
}

type Config struct {
	AdminUsername string `envconfig:"SECURITY_USER_NAME" required:"true"`
	AdminPassword string `envconfig:"SECURITY_USER_PASSWORD" required:"true"`

	Port         int    `envconfig:"PORT" default:"8080"`
	HelmChartDir string `envconfig:"HELM_CHART_DIR" default:"charts"`
	OperatorDir  string `envconfig:"OPERATOR_DIR" default:"operators"`
	StateDir     string `envconfig:"STATE_DIR" default:"state"`

	ClusterCredentials *ClusterCredentials
	RegistryConfig     *RegistryConfig
	CFClientConfig     *CFClientConfig
	TillerTLSConfig    *TillerTLSConfig
}

func (r RegistryConfig) HasRegistryConfig() bool {
	return r.Server != ""
}

func (c CFClientConfig) HasCFClientConfig() bool {
	return c.ApiAddress != ""
}

func (r RegistryConfig) GetDockerConfigJson() ([]byte, error) {
	if r.Server == "" || r.Email == "" || r.Pass == "" || r.User == "" {
		return nil, errors.New("environment didn't have a proper registry Config")
	}

	dockerConfig := map[string]interface{}{
		"auths": map[string]interface{}{
			r.Server: map[string]interface{}{
				"username": r.User,
				"password": r.Pass,
				"email":    r.Email,
			},
		},
	}
	dockerConfigJson, err := json.Marshal(dockerConfig)
	if err != nil {
		return nil, err
	}

	return dockerConfigJson, nil
}

func Parse() (*Config, error) {
	c := &Config{}
	err := envconfig.Process("", c)
	if err != nil {
		return nil, err
	}

	err = c.ClusterCredentials.parseCAData()
	if err != nil {
		return nil, err
	}

	err = c.TillerTLSConfig.validateTillerParameters()
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c *ClusterCredentials) parseCAData() error {
	c.CADataRaw = strings.TrimSpace(c.CADataRaw)
	if strings.Index(c.CADataRaw, "-----BEGIN CERTIFICATE-----") == 0 {
		c.CAData = []byte(c.CADataRaw)
	} else {
		data, err := base64.StdEncoding.DecodeString(c.CADataRaw)
		if err != nil {
			return err
		} else {
			c.CAData = data
		}
	}

	return nil
}

func (t *TillerTLSConfig) validateTillerParameters() error {
	files := []string{
		t.TLSCertFile, t.TLSKeyFile, t.TLSCaCertFile,
	}
	for _, file := range files {
		if file != "" {
			exists, err := fileExists(file)
			if !exists {
				return errors.New(fmt.Sprintf("Tiller file [%s] does not exist", file))
			}
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func fileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	} else {
		if os.IsNotExist(err) {
			return false, nil
		} else {
			return false, err
		}
	}
}
