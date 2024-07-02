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

import "strconv"

type Spider struct {
	c *Client
}

// This component is optional and therefore the API will only work if it is installed
func (s Spider) Status(scanid string) (map[string]interface{}, error) {
	m := map[string]string{
		"scanId": scanid,
	}
	return s.c.Request("spider/view/status/", m)
}

// This component is optional and therefore the API will only work if it is installed
func (s Spider) Results(scanid string) (map[string]interface{}, error) {
	m := map[string]string{
		"scanId": scanid,
	}
	return s.c.Request("spider/view/results/", m)
}

// This component is optional and therefore the API will only work if it is installed
func (s Spider) FullResults(scanid string) (map[string]interface{}, error) {
	m := map[string]string{
		"scanId": scanid,
	}
	return s.c.Request("spider/view/fullResults/", m)
}

// This component is optional and therefore the API will only work if it is installed
func (s Spider) Scans() (map[string]interface{}, error) {
	return s.c.Request("spider/view/scans/", nil)
}

// Gets the regexes of URLs excluded from the spider scans.
//
// This component is optional and therefore the API will only work if it is installed
func (s Spider) ExcludedFromScan() (map[string]interface{}, error) {
	return s.c.Request("spider/view/excludedFromScan/", nil)
}

// Returns a list of unique URLs from the history table based on HTTP messages added by the Spider.
//
// This component is optional and therefore the API will only work if it is installed
func (s Spider) AllUrls() (map[string]interface{}, error) {
	return s.c.Request("spider/view/allUrls/", nil)
}

// Returns a list of the names of the nodes added to the Sites tree by the specified scan.
//
// This component is optional and therefore the API will only work if it is installed
func (s Spider) AddedNodes(scanid string) (map[string]interface{}, error) {
	m := map[string]string{
		"scanId": scanid,
	}
	return s.c.Request("spider/view/addedNodes/", m)
}

// Gets all the domains that are always in scope. For each domain the following are shown: the index, the value (domain), if enabled, and if specified as a regex.
//
// This component is optional and therefore the API will only work if it is installed
func (s Spider) DomainsAlwaysInScope() (map[string]interface{}, error) {
	return s.c.Request("spider/view/domainsAlwaysInScope/", nil)
}

// Use view domainsAlwaysInScope instead.
//
// This component is optional and therefore the API will only work if it is installed
func (s Spider) OptionDomainsAlwaysInScope() (map[string]interface{}, error) {
	return s.c.Request("spider/view/optionDomainsAlwaysInScope/", nil)
}

// Use view domainsAlwaysInScope instead.
//
// This component is optional and therefore the API will only work if it is installed
func (s Spider) OptionDomainsAlwaysInScopeEnabled() (map[string]interface{}, error) {
	return s.c.Request("spider/view/optionDomainsAlwaysInScopeEnabled/", nil)
}

// This component is optional and therefore the API will only work if it is installed
func (s Spider) OptionHandleParameters() (map[string]interface{}, error) {
	return s.c.Request("spider/view/optionHandleParameters/", nil)
}

// Gets the maximum number of child nodes (per node) that can be crawled, 0 means no limit.
//
// This component is optional and therefore the API will only work if it is installed
func (s Spider) OptionMaxChildren() (map[string]interface{}, error) {
	return s.c.Request("spider/view/optionMaxChildren/", nil)
}

// Gets the maximum depth the spider can crawl, 0 if unlimited.
//
// This component is optional and therefore the API will only work if it is installed
func (s Spider) OptionMaxDepth() (map[string]interface{}, error) {
	return s.c.Request("spider/view/optionMaxDepth/", nil)
}

// This component is optional and therefore the API will only work if it is installed
func (s Spider) OptionMaxDuration() (map[string]interface{}, error) {
	return s.c.Request("spider/view/optionMaxDuration/", nil)
}

// Gets the maximum size, in bytes, that a response might have to be parsed.
//
// This component is optional and therefore the API will only work if it is installed
func (s Spider) OptionMaxParseSizeBytes() (map[string]interface{}, error) {
	return s.c.Request("spider/view/optionMaxParseSizeBytes/", nil)
}

// This component is optional and therefore the API will only work if it is installed
func (s Spider) OptionMaxScansInUI() (map[string]interface{}, error) {
	return s.c.Request("spider/view/optionMaxScansInUI/", nil)
}

// This component is optional and therefore the API will only work if it is installed
func (s Spider) OptionRequestWaitTime() (map[string]interface{}, error) {
	return s.c.Request("spider/view/optionRequestWaitTime/", nil)
}

// This component is optional and therefore the API will only work if it is installed
func (s Spider) OptionSkipURLString() (map[string]interface{}, error) {
	return s.c.Request("spider/view/optionSkipURLString/", nil)
}

// This component is optional and therefore the API will only work if it is installed
func (s Spider) OptionThreadCount() (map[string]interface{}, error) {
	return s.c.Request("spider/view/optionThreadCount/", nil)
}

// This component is optional and therefore the API will only work if it is installed
func (s Spider) OptionUserAgent() (map[string]interface{}, error) {
	return s.c.Request("spider/view/optionUserAgent/", nil)
}

// Gets whether or not a spider process should accept cookies while spidering.
//
// This component is optional and therefore the API will only work if it is installed
func (s Spider) OptionAcceptCookies() (map[string]interface{}, error) {
	return s.c.Request("spider/view/optionAcceptCookies/", nil)
}

// This component is optional and therefore the API will only work if it is installed
func (s Spider) OptionHandleODataParametersVisited() (map[string]interface{}, error) {
	return s.c.Request("spider/view/optionHandleODataParametersVisited/", nil)
}

// This component is optional and therefore the API will only work if it is installed
func (s Spider) OptionParseComments() (map[string]interface{}, error) {
	return s.c.Request("spider/view/optionParseComments/", nil)
}

// This component is optional and therefore the API will only work if it is installed
func (s Spider) OptionParseDsStore() (map[string]interface{}, error) {
	return s.c.Request("spider/view/optionParseDsStore/", nil)
}

// This component is optional and therefore the API will only work if it is installed
func (s Spider) OptionParseGit() (map[string]interface{}, error) {
	return s.c.Request("spider/view/optionParseGit/", nil)
}

// This component is optional and therefore the API will only work if it is installed
func (s Spider) OptionParseRobotsTxt() (map[string]interface{}, error) {
	return s.c.Request("spider/view/optionParseRobotsTxt/", nil)
}

// This component is optional and therefore the API will only work if it is installed
func (s Spider) OptionParseSVNEntries() (map[string]interface{}, error) {
	return s.c.Request("spider/view/optionParseSVNEntries/", nil)
}

// This component is optional and therefore the API will only work if it is installed
func (s Spider) OptionParseSitemapXml() (map[string]interface{}, error) {
	return s.c.Request("spider/view/optionParseSitemapXml/", nil)
}

// This component is optional and therefore the API will only work if it is installed
func (s Spider) OptionPostForm() (map[string]interface{}, error) {
	return s.c.Request("spider/view/optionPostForm/", nil)
}

// This component is optional and therefore the API will only work if it is installed
func (s Spider) OptionProcessForm() (map[string]interface{}, error) {
	return s.c.Request("spider/view/optionProcessForm/", nil)
}

// Gets whether or not the 'Referer' header should be sent while spidering.
//
// This component is optional and therefore the API will only work if it is installed
func (s Spider) OptionSendRefererHeader() (map[string]interface{}, error) {
	return s.c.Request("spider/view/optionSendRefererHeader/", nil)
}

// This component is optional and therefore the API will only work if it is installed
func (s Spider) OptionShowAdvancedDialog() (map[string]interface{}, error) {
	return s.c.Request("spider/view/optionShowAdvancedDialog/", nil)
}

// Runs the spider against the given URL (or context). Optionally, the 'maxChildren' parameter can be set to limit the number of children scanned, the 'recurse' parameter can be used to prevent the spider from seeding recursively, the parameter 'contextName' can be used to constrain the scan to a Context and the parameter 'subtreeOnly' allows to restrict the spider under a site's subtree (using the specified 'url').
//
// This component is optional and therefore the API will only work if it is installed
func (s Spider) Scan(url string, maxchildren string, recurse string, contextname string, subtreeonly string) (map[string]interface{}, error) {
	m := map[string]string{
		"url":         url,
		"maxChildren": maxchildren,
		"recurse":     recurse,
		"contextName": contextname,
		"subtreeOnly": subtreeonly,
	}
	return s.c.Request("spider/action/scan/", m)
}

// Runs the spider from the perspective of a User, obtained using the given Context ID and User ID. See 'scan' action for more details.
//
// This component is optional and therefore the API will only work if it is installed
func (s Spider) ScanAsUser(contextid string, userid string, url string, maxchildren string, recurse string, subtreeonly string) (map[string]interface{}, error) {
	m := map[string]string{
		"contextId":   contextid,
		"userId":      userid,
		"url":         url,
		"maxChildren": maxchildren,
		"recurse":     recurse,
		"subtreeOnly": subtreeonly,
	}
	return s.c.Request("spider/action/scanAsUser/", m)
}

// This component is optional and therefore the API will only work if it is installed
func (s Spider) Pause(scanid string) (map[string]interface{}, error) {
	m := map[string]string{
		"scanId": scanid,
	}
	return s.c.Request("spider/action/pause/", m)
}

// This component is optional and therefore the API will only work if it is installed
func (s Spider) Resume(scanid string) (map[string]interface{}, error) {
	m := map[string]string{
		"scanId": scanid,
	}
	return s.c.Request("spider/action/resume/", m)
}

// This component is optional and therefore the API will only work if it is installed
func (s Spider) Stop(scanid string) (map[string]interface{}, error) {
	m := map[string]string{
		"scanId": scanid,
	}
	return s.c.Request("spider/action/stop/", m)
}

// This component is optional and therefore the API will only work if it is installed
func (s Spider) RemoveScan(scanid string) (map[string]interface{}, error) {
	m := map[string]string{
		"scanId": scanid,
	}
	return s.c.Request("spider/action/removeScan/", m)
}

// This component is optional and therefore the API will only work if it is installed
func (s Spider) PauseAllScans() (map[string]interface{}, error) {
	return s.c.Request("spider/action/pauseAllScans/", nil)
}

// This component is optional and therefore the API will only work if it is installed
func (s Spider) ResumeAllScans() (map[string]interface{}, error) {
	return s.c.Request("spider/action/resumeAllScans/", nil)
}

// This component is optional and therefore the API will only work if it is installed
func (s Spider) StopAllScans() (map[string]interface{}, error) {
	return s.c.Request("spider/action/stopAllScans/", nil)
}

// This component is optional and therefore the API will only work if it is installed
func (s Spider) RemoveAllScans() (map[string]interface{}, error) {
	return s.c.Request("spider/action/removeAllScans/", nil)
}

// Clears the regexes of URLs excluded from the spider scans.
//
// This component is optional and therefore the API will only work if it is installed
func (s Spider) ClearExcludedFromScan() (map[string]interface{}, error) {
	return s.c.Request("spider/action/clearExcludedFromScan/", nil)
}

// Adds a regex of URLs that should be excluded from the spider scans.
//
// This component is optional and therefore the API will only work if it is installed
func (s Spider) ExcludeFromScan(regex string) (map[string]interface{}, error) {
	m := map[string]string{
		"regex": regex,
	}
	return s.c.Request("spider/action/excludeFromScan/", m)
}

// Adds a new domain that's always in scope, using the specified value. Optionally sets if the new entry is enabled (default, true) and whether or not the new value is specified as a regex (default, false).
//
// This component is optional and therefore the API will only work if it is installed
func (s Spider) AddDomainAlwaysInScope(value string, isregex string, isenabled string) (map[string]interface{}, error) {
	m := map[string]string{
		"value":     value,
		"isRegex":   isregex,
		"isEnabled": isenabled,
	}
	return s.c.Request("spider/action/addDomainAlwaysInScope/", m)
}

// Modifies a domain that's always in scope. Allows to modify the value, if enabled or if a regex. The domain is selected with its index, which can be obtained with the view domainsAlwaysInScope.
//
// This component is optional and therefore the API will only work if it is installed
func (s Spider) ModifyDomainAlwaysInScope(idx string, value string, isregex string, isenabled string) (map[string]interface{}, error) {
	m := map[string]string{
		"idx":       idx,
		"value":     value,
		"isRegex":   isregex,
		"isEnabled": isenabled,
	}
	return s.c.Request("spider/action/modifyDomainAlwaysInScope/", m)
}

// Removes a domain that's always in scope, with the given index. The index can be obtained with the view domainsAlwaysInScope.
//
// This component is optional and therefore the API will only work if it is installed
func (s Spider) RemoveDomainAlwaysInScope(idx string) (map[string]interface{}, error) {
	m := map[string]string{
		"idx": idx,
	}
	return s.c.Request("spider/action/removeDomainAlwaysInScope/", m)
}

// Enables all domains that are always in scope.
//
// This component is optional and therefore the API will only work if it is installed
func (s Spider) EnableAllDomainsAlwaysInScope() (map[string]interface{}, error) {
	return s.c.Request("spider/action/enableAllDomainsAlwaysInScope/", nil)
}

// Disables all domains that are always in scope.
//
// This component is optional and therefore the API will only work if it is installed
func (s Spider) DisableAllDomainsAlwaysInScope() (map[string]interface{}, error) {
	return s.c.Request("spider/action/disableAllDomainsAlwaysInScope/", nil)
}

// This component is optional and therefore the API will only work if it is installed
func (s Spider) SetOptionHandleParameters(str string) (map[string]interface{}, error) {
	m := map[string]string{
		"String": str,
	}
	return s.c.Request("spider/action/setOptionHandleParameters/", m)
}

// This component is optional and therefore the API will only work if it is installed
func (s Spider) SetOptionSkipURLString(str string) (map[string]interface{}, error) {
	m := map[string]string{
		"String": str,
	}
	return s.c.Request("spider/action/setOptionSkipURLString/", m)
}

// This component is optional and therefore the API will only work if it is installed
func (s Spider) SetOptionUserAgent(str string) (map[string]interface{}, error) {
	m := map[string]string{
		"String": str,
	}
	return s.c.Request("spider/action/setOptionUserAgent/", m)
}

// Sets whether or not a spider process should accept cookies while spidering.
//
// This component is optional and therefore the API will only work if it is installed
func (s Spider) SetOptionAcceptCookies(boolean bool) (map[string]interface{}, error) {
	m := map[string]string{
		"Boolean": strconv.FormatBool(boolean),
	}
	return s.c.Request("spider/action/setOptionAcceptCookies/", m)
}

// This component is optional and therefore the API will only work if it is installed
func (s Spider) SetOptionHandleODataParametersVisited(boolean bool) (map[string]interface{}, error) {
	m := map[string]string{
		"Boolean": strconv.FormatBool(boolean),
	}
	return s.c.Request("spider/action/setOptionHandleODataParametersVisited/", m)
}

// Sets the maximum number of child nodes (per node) that can be crawled, 0 means no limit.
//
// This component is optional and therefore the API will only work if it is installed
func (s Spider) SetOptionMaxChildren(i int) (map[string]interface{}, error) {
	m := map[string]string{
		"Integer": strconv.Itoa(i),
	}
	return s.c.Request("spider/action/setOptionMaxChildren/", m)
}

// Sets the maximum depth the spider can crawl, 0 for unlimited depth.
//
// This component is optional and therefore the API will only work if it is installed
func (s Spider) SetOptionMaxDepth(i int) (map[string]interface{}, error) {
	m := map[string]string{
		"Integer": strconv.Itoa(i),
	}
	return s.c.Request("spider/action/setOptionMaxDepth/", m)
}

// This component is optional and therefore the API will only work if it is installed
func (s Spider) SetOptionMaxDuration(i int) (map[string]interface{}, error) {
	m := map[string]string{
		"Integer": strconv.Itoa(i),
	}
	return s.c.Request("spider/action/setOptionMaxDuration/", m)
}

// Sets the maximum size, in bytes, that a response might have to be parsed. This allows the spider to skip big responses/files.
//
// This component is optional and therefore the API will only work if it is installed
func (s Spider) SetOptionMaxParseSizeBytes(i int) (map[string]interface{}, error) {
	m := map[string]string{
		"Integer": strconv.Itoa(i),
	}
	return s.c.Request("spider/action/setOptionMaxParseSizeBytes/", m)
}

// This component is optional and therefore the API will only work if it is installed
func (s Spider) SetOptionMaxScansInUI(i int) (map[string]interface{}, error) {
	m := map[string]string{
		"Integer": strconv.Itoa(i),
	}
	return s.c.Request("spider/action/setOptionMaxScansInUI/", m)
}

// This component is optional and therefore the API will only work if it is installed
func (s Spider) SetOptionParseComments(boolean bool) (map[string]interface{}, error) {
	m := map[string]string{
		"Boolean": strconv.FormatBool(boolean),
	}
	return s.c.Request("spider/action/setOptionParseComments/", m)
}

// This component is optional and therefore the API will only work if it is installed
func (s Spider) SetOptionParseDsStore(boolean bool) (map[string]interface{}, error) {
	m := map[string]string{
		"Boolean": strconv.FormatBool(boolean),
	}
	return s.c.Request("spider/action/setOptionParseDsStore/", m)
}

// This component is optional and therefore the API will only work if it is installed
func (s Spider) SetOptionParseGit(boolean bool) (map[string]interface{}, error) {
	m := map[string]string{
		"Boolean": strconv.FormatBool(boolean),
	}
	return s.c.Request("spider/action/setOptionParseGit/", m)
}

// This component is optional and therefore the API will only work if it is installed
func (s Spider) SetOptionParseRobotsTxt(boolean bool) (map[string]interface{}, error) {
	m := map[string]string{
		"Boolean": strconv.FormatBool(boolean),
	}
	return s.c.Request("spider/action/setOptionParseRobotsTxt/", m)
}

// This component is optional and therefore the API will only work if it is installed
func (s Spider) SetOptionParseSVNEntries(boolean bool) (map[string]interface{}, error) {
	m := map[string]string{
		"Boolean": strconv.FormatBool(boolean),
	}
	return s.c.Request("spider/action/setOptionParseSVNEntries/", m)
}

// This component is optional and therefore the API will only work if it is installed
func (s Spider) SetOptionParseSitemapXml(boolean bool) (map[string]interface{}, error) {
	m := map[string]string{
		"Boolean": strconv.FormatBool(boolean),
	}
	return s.c.Request("spider/action/setOptionParseSitemapXml/", m)
}

// This component is optional and therefore the API will only work if it is installed
func (s Spider) SetOptionPostForm(boolean bool) (map[string]interface{}, error) {
	m := map[string]string{
		"Boolean": strconv.FormatBool(boolean),
	}
	return s.c.Request("spider/action/setOptionPostForm/", m)
}

// This component is optional and therefore the API will only work if it is installed
func (s Spider) SetOptionProcessForm(boolean bool) (map[string]interface{}, error) {
	m := map[string]string{
		"Boolean": strconv.FormatBool(boolean),
	}
	return s.c.Request("spider/action/setOptionProcessForm/", m)
}

// This component is optional and therefore the API will only work if it is installed
func (s Spider) SetOptionRequestWaitTime(i int) (map[string]interface{}, error) {
	m := map[string]string{
		"Integer": strconv.Itoa(i),
	}
	return s.c.Request("spider/action/setOptionRequestWaitTime/", m)
}

// Sets whether or not the 'Referer' header should be sent while spidering.
//
// This component is optional and therefore the API will only work if it is installed
func (s Spider) SetOptionSendRefererHeader(boolean bool) (map[string]interface{}, error) {
	m := map[string]string{
		"Boolean": strconv.FormatBool(boolean),
	}
	return s.c.Request("spider/action/setOptionSendRefererHeader/", m)
}

// This component is optional and therefore the API will only work if it is installed
func (s Spider) SetOptionShowAdvancedDialog(boolean bool) (map[string]interface{}, error) {
	m := map[string]string{
		"Boolean": strconv.FormatBool(boolean),
	}
	return s.c.Request("spider/action/setOptionShowAdvancedDialog/", m)
}

// This component is optional and therefore the API will only work if it is installed
func (s Spider) SetOptionThreadCount(i int) (map[string]interface{}, error) {
	m := map[string]string{
		"Integer": strconv.Itoa(i),
	}
	return s.c.Request("spider/action/setOptionThreadCount/", m)
}
