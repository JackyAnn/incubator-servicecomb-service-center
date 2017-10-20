//Copyright 2017 Huawei Technologies Co., Ltd
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
package plain

import (
	"github.com/ServiceComb/service-center/server/infra/security"
)

type DefaultCipher struct {
}

func init() {
	security.CipherPlugins["default"] = New

}
func New() security.Cipher {

	return &DefaultCipher{}
}

func (c *DefaultCipher) Encrypt(src string) (string, error) {
	return src, nil
}

func (c *DefaultCipher) Decrypt(src string) (string, error) {
	return src, nil
}