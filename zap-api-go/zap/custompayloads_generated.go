// Zed Attack Proxy (ZAP) and its related class files.
//
// ZAP is an HTTP/HTTPS proxy for assessing web application security.
//
// Copyright 2017 the ZAP development team
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// *** This file was automatically generated. ***
//

package zap

type Custompayloads struct {
	c *Client
}

// Lists all available categories.
//
// This component is optional and therefore the API will only work if it is installed
func (c Custompayloads) CustomPayloadsCategories() (map[string]interface{}, error) {
	return c.c.Request("custompayloads/view/customPayloadsCategories/", nil)
}

// Lists all the payloads currently loaded (category, payload, enabled state). Optionally filtered by category.
//
// This component is optional and therefore the API will only work if it is installed
func (c Custompayloads) CustomPayloads(category string) (map[string]interface{}, error) {
	m := map[string]string{
		"category": category,
	}
	return c.c.Request("custompayloads/view/customPayloads/", m)
}

// Disables payloads for a given category.
//
// This component is optional and therefore the API will only work if it is installed
func (c Custompayloads) DisableCustomPayloads(category string) (map[string]interface{}, error) {
	m := map[string]string{
		"category": category,
	}
	return c.c.Request("custompayloads/action/disableCustomPayloads/", m)
}

// Enables payloads for a given category.
//
// This component is optional and therefore the API will only work if it is installed
func (c Custompayloads) EnableCustomPayloads(category string) (map[string]interface{}, error) {
	m := map[string]string{
		"category": category,
	}
	return c.c.Request("custompayloads/action/enableCustomPayloads/", m)
}

// Removes a payload.
//
// This component is optional and therefore the API will only work if it is installed
func (c Custompayloads) RemoveCustomPayload(category string, payload string) (map[string]interface{}, error) {
	m := map[string]string{
		"category": category,
		"payload":  payload,
	}
	return c.c.Request("custompayloads/action/removeCustomPayload/", m)
}

// Adds a new payload.
//
// This component is optional and therefore the API will only work if it is installed
func (c Custompayloads) AddCustomPayload(category string, payload string) (map[string]interface{}, error) {
	m := map[string]string{
		"category": category,
		"payload":  payload,
	}
	return c.c.Request("custompayloads/action/addCustomPayload/", m)
}

// Enables a given payload.
//
// This component is optional and therefore the API will only work if it is installed
func (c Custompayloads) EnableCustomPayload(category string, payload string) (map[string]interface{}, error) {
	m := map[string]string{
		"category": category,
		"payload":  payload,
	}
	return c.c.Request("custompayloads/action/enableCustomPayload/", m)
}

// Disables a given payload.
//
// This component is optional and therefore the API will only work if it is installed
func (c Custompayloads) DisableCustomPayload(category string, payload string) (map[string]interface{}, error) {
	m := map[string]string{
		"category": category,
		"payload":  payload,
	}
	return c.c.Request("custompayloads/action/disableCustomPayload/", m)
}
