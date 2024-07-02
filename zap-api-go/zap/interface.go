// Package zap defines the interface a ZAP client should implement
package zap

// Interface defines the interface a ZAP client should implement
type Interface interface {
	Alert() *Alert
	Postman() *Postman
	Authentication() *Authentication
	Retest() *Retest
	Network() *Network
	Pnh() *Pnh
	Graphql() *Graphql
	Soap() *Soap
	Spider() *Spider
	Autoupdate() *Autoupdate
	Core() *Core
	Break() *Break
	Exim() *Exim
	Context() *Context
	Pscan() *Pscan
	ForcedUser() *ForcedUser
	Script() *Script
	Authorization() *Authorization
	Automation() *Automation
	Search() *Search
	Reports() *Reports
	Selenium() *Selenium
	HttpSessions() *HttpSessions
	Websocket() *Websocket
	Openapi() *Openapi
	AjaxSpider() *AjaxSpider
	AccessControl() *AccessControl
	Reveal() *Reveal
	Params() *Params
	Revisit() *Revisit
	AlertFilter() *AlertFilter
	Replacer() *Replacer
	Users() *Users
	Acsrf() *Acsrf
	Ascan() *Ascan
	Wappalyzer() *Wappalyzer
	Stats() *Stats
	RuleConfig() *RuleConfig
	SessionManagement() *SessionManagement
	Custompayloads() *Custompayloads
}

// Alert() returns a Alert client
func (c *Client) Alert() *Alert {
	return &Alert{c}
}

// Postman() returns a Postman client
func (c *Client) Postman() *Postman {
	return &Postman{c}
}

// Authentication() returns a Authentication client
func (c *Client) Authentication() *Authentication {
	return &Authentication{c}
}

// Retest() returns a Retest client
func (c *Client) Retest() *Retest {
	return &Retest{c}
}

// Network() returns a Network client
func (c *Client) Network() *Network {
	return &Network{c}
}

// Pnh() returns a Pnh client
func (c *Client) Pnh() *Pnh {
	return &Pnh{c}
}

// Graphql() returns a Graphql client
func (c *Client) Graphql() *Graphql {
	return &Graphql{c}
}

// Soap() returns a Soap client
func (c *Client) Soap() *Soap {
	return &Soap{c}
}

// Spider() returns a Spider client
func (c *Client) Spider() *Spider {
	return &Spider{c}
}

// Autoupdate() returns a Autoupdate client
func (c *Client) Autoupdate() *Autoupdate {
	return &Autoupdate{c}
}

// Core() returns a Core client
func (c *Client) Core() *Core {
	return &Core{c}
}

// Break() returns a Break client
func (c *Client) Break() *Break {
	return &Break{c}
}

// Exim() returns a Exim client
func (c *Client) Exim() *Exim {
	return &Exim{c}
}

// Context() returns a Context client
func (c *Client) Context() *Context {
	return &Context{c}
}

// Pscan() returns a Pscan client
func (c *Client) Pscan() *Pscan {
	return &Pscan{c}
}

// ForcedUser() returns a ForcedUser client
func (c *Client) ForcedUser() *ForcedUser {
	return &ForcedUser{c}
}

// Script() returns a Script client
func (c *Client) Script() *Script {
	return &Script{c}
}

// Authorization() returns a Authorization client
func (c *Client) Authorization() *Authorization {
	return &Authorization{c}
}

// Automation() returns a Automation client
func (c *Client) Automation() *Automation {
	return &Automation{c}
}

// Search() returns a Search client
func (c *Client) Search() *Search {
	return &Search{c}
}

// Reports() returns a Reports client
func (c *Client) Reports() *Reports {
	return &Reports{c}
}

// Selenium() returns a Selenium client
func (c *Client) Selenium() *Selenium {
	return &Selenium{c}
}

// HttpSessions() returns a HttpSessions client
func (c *Client) HttpSessions() *HttpSessions {
	return &HttpSessions{c}
}

// Websocket() returns a Websocket client
func (c *Client) Websocket() *Websocket {
	return &Websocket{c}
}

// Openapi() returns a Openapi client
func (c *Client) Openapi() *Openapi {
	return &Openapi{c}
}

// AjaxSpider() returns a AjaxSpider client
func (c *Client) AjaxSpider() *AjaxSpider {
	return &AjaxSpider{c}
}

// AccessControl() returns a AccessControl client
func (c *Client) AccessControl() *AccessControl {
	return &AccessControl{c}
}

// Reveal() returns a Reveal client
func (c *Client) Reveal() *Reveal {
	return &Reveal{c}
}

// Params() returns a Params client
func (c *Client) Params() *Params {
	return &Params{c}
}

// Revisit() returns a Revisit client
func (c *Client) Revisit() *Revisit {
	return &Revisit{c}
}

// AlertFilter() returns a AlertFilter client
func (c *Client) AlertFilter() *AlertFilter {
	return &AlertFilter{c}
}

// Replacer() returns a Replacer client
func (c *Client) Replacer() *Replacer {
	return &Replacer{c}
}

// Users() returns a Users client
func (c *Client) Users() *Users {
	return &Users{c}
}

// Acsrf() returns a Acsrf client
func (c *Client) Acsrf() *Acsrf {
	return &Acsrf{c}
}

// Ascan() returns a Ascan client
func (c *Client) Ascan() *Ascan {
	return &Ascan{c}
}

// Wappalyzer() returns a Wappalyzer client
func (c *Client) Wappalyzer() *Wappalyzer {
	return &Wappalyzer{c}
}

// Stats() returns a Stats client
func (c *Client) Stats() *Stats {
	return &Stats{c}
}

// RuleConfig() returns a RuleConfig client
func (c *Client) RuleConfig() *RuleConfig {
	return &RuleConfig{c}
}

// SessionManagement() returns a SessionManagement client
func (c *Client) SessionManagement() *SessionManagement {
	return &SessionManagement{c}
}

// Custompayloads() returns a Custompayloads client
func (c *Client) Custompayloads() *Custompayloads {
	return &Custompayloads{c}
}
