// Code generated by schema-generate. DO NOT EDIT.

package schema

// ContainerOverrides
type ContainerOverrides struct {
	Limits   *ResourceOverride `json:"limits,omitempty"`
	Requests *ResourceOverride `json:"requests,omitempty"`
}

// DatacenterOpsConfiguration Configuration for Sourcegraph Datacenter Ops. Defines all config parameters that can be used to change the operational parameters of Sourcegraph Datacenter (e.g., for scaling).
type DatacenterOpsConfiguration struct {
	AlertmanagerConfig  string                          `json:"alertmanagerConfig,omitempty"`
	AlertmanagerURL     string                          `json:"alertmanagerURL,omitempty"`
	AuthProxyIP         string                          `json:"authProxyIP,omitempty"`
	AuthProxyPassword   string                          `json:"authProxyPassword,omitempty"`
	DeploymentOverrides map[string]*DeploymentOverrides `json:"deploymentOverrides,omitempty"`
	GitoliteIP          string                          `json:"gitoliteIP,omitempty"`
	GitserverCount      int                             `json:"gitserverCount,omitempty"`
	GitserverSSH        map[string]string               `json:"gitserverSSH,omitempty"`
	HttpNodePort        int                             `json:"httpNodePort,omitempty"`
	HttpsNodePort       int                             `json:"httpsNodePort,omitempty"`
	LangGo              bool                            `json:"langGo,omitempty"`
	LangJava            bool                            `json:"langJava,omitempty"`
	LangJavaScript      bool                            `json:"langJavaScript,omitempty"`
	LangPHP             bool                            `json:"langPHP,omitempty"`
	LangPython          bool                            `json:"langPython,omitempty"`
	LangSwift           bool                            `json:"langSwift,omitempty"`
	LangTypeScript      bool                            `json:"langTypeScript,omitempty"`
	NodeSSDPath         string                          `json:"nodeSSDPath,omitempty"`
	PhabricatorIP       string                          `json:"phabricatorIP,omitempty"`
	StorageClass        string                          `json:"storageClass,omitempty"`
	UseRBAC             bool                            `json:"useRBAC,omitempty"`
}

// DeploymentOverrides
type DeploymentOverrides struct {
	Containers map[string]*ContainerOverrides `json:"containers,omitempty"`
	Replicas   *int                           `json:"replicas,omitempty"`
}

// Github
type Github struct {
	Url         string   `json:"url,omitempty"`
	Token       string   `json:"token,omitempty"`
	Certificate string   `json:"certificate,omitempty"`
	Repos       []string `json:"repos,omitempty"`
}

// Phabricator
type Phabricator struct {
	Url   string  `json:"url,omitempty"`
	Token string  `json:"token,omitempty"`
	Repos []Repos `json:"repos,omitempty"`
}

// Repos
type Repos struct {
	Callsign string `json:"callsign,omitempty"`
	Path     string `json:"path,omitempty"`
}

// ReposList
type ReposList struct {
	Path string `json:"path,omitempty"`
	Type string `json:"type,omitempty"`
	Url  string `json:"url,omitempty"`
}

// ResourceOverride
type ResourceOverride struct {
	Cpu    string `json:"cpu,omitempty"`
	Memory string `json:"memory,omitempty"`
}

// SearchSavedQueries
type SearchSavedQueries struct {
	Description string `json:"description"`
	Key         string `json:"key"`
	Query       string `json:"query"`
	ScopeQuery  string `json:"scopeQuery,omitempty"`
}

// SearchScope
type SearchScope struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// Settings Configuration settings for users and organizations on Sourcegraph Server.
type Settings struct {
	SearchSavedQueries []SearchSavedQueries `json:"search.savedQueries,omitempty"`
	SearchScopes       []SearchScope        `json:"search.scopes,omitempty"`
}

// SiteConfiguration Configuration for a Sourcegraph Server site.
type SiteConfiguration struct {
	AdminUsernames                 string              `json:"adminUsernames,omitempty"`
	AppID                          string              `json:"appID,omitempty"`
	AppURL                         string              `json:"appURL,omitempty"`
	AuthUserOrgMap                 map[string][]string `json:"auth.userOrgMap,omitempty"`
	AutoRepoAdd                    bool                `json:"autoRepoAdd,omitempty"`
	CorsOrigin                     string              `json:"corsOrigin,omitempty"`
	DisablePublicRepoRedirects     bool                `json:"disablePublicRepoRedirects,omitempty"`
	DisableTelemetry               bool                `json:"disableTelemetry,omitempty"`
	ExecuteGradleOriginalRootPaths string              `json:"executeGradleOriginalRootPaths,omitempty"`
	GitOriginMap                   string              `json:"gitOriginMap,omitempty"`
	Github                         []Github            `json:"github,omitempty"`
	GithubClientID                 string              `json:"githubClientID,omitempty"`
	GithubClientSecret             string              `json:"githubClientSecret,omitempty"`
	GithubEnterpriseAccessToken    string              `json:"githubEnterpriseAccessToken,omitempty"`
	GithubEnterpriseCert           string              `json:"githubEnterpriseCert,omitempty"`
	GithubEnterpriseURL            string              `json:"githubEnterpriseURL,omitempty"`
	GithubPersonalAccessToken      string              `json:"githubPersonalAccessToken,omitempty"`
	GitoliteHosts                  string              `json:"gitoliteHosts,omitempty"`
	HtmlBodyBottom                 string              `json:"htmlBodyBottom,omitempty"`
	HtmlBodyTop                    string              `json:"htmlBodyTop,omitempty"`
	HtmlHeadBottom                 string              `json:"htmlHeadBottom,omitempty"`
	HtmlHeadTop                    string              `json:"htmlHeadTop,omitempty"`
	InactiveRepos                  string              `json:"inactiveRepos,omitempty"`
	LicenseKey                     string              `json:"licenseKey,omitempty"`
	LightstepAccessToken           string              `json:"lightstepAccessToken,omitempty"`
	LightstepProject               string              `json:"lightstepProject,omitempty"`
	MandrillKey                    string              `json:"mandrillKey,omitempty"`
	MaxReposToSearch               int                 `json:"maxReposToSearch,omitempty"`
	NoGoGetDomains                 string              `json:"noGoGetDomains,omitempty"`
	OidcClientID                   string              `json:"oidcClientID,omitempty"`
	OidcClientSecret               string              `json:"oidcClientSecret,omitempty"`
	OidcEmailDomain                string              `json:"oidcEmailDomain,omitempty"`
	OidcOverrideToken              string              `json:"oidcOverrideToken,omitempty"`
	OidcProvider                   string              `json:"oidcProvider,omitempty"`
	Phabricator                    []Phabricator       `json:"phabricator,omitempty"`
	PhabricatorURL                 string              `json:"phabricatorURL,omitempty"`
	PrivateArtifactRepoID          string              `json:"privateArtifactRepoID,omitempty"`
	PrivateArtifactRepoPassword    string              `json:"privateArtifactRepoPassword,omitempty"`
	PrivateArtifactRepoURL         string              `json:"privateArtifactRepoURL,omitempty"`
	PrivateArtifactRepoUsername    string              `json:"privateArtifactRepoUsername,omitempty"`
	RepoListUpdateInterval         int                 `json:"repoListUpdateInterval,omitempty"`
	ReposList                      []ReposList         `json:"repos.list,omitempty"`
	SamlIDProviderMetadataURL      string              `json:"samlIDProviderMetadataURL,omitempty"`
	SamlSPCert                     string              `json:"samlSPCert,omitempty"`
	SamlSPKey                      string              `json:"samlSPKey,omitempty"`
	SearchScopes                   []SearchScope       `json:"searchScopes,omitempty"`
	Settings                       *Settings           `json:"settings,omitempty"`
	SsoUserHeader                  string              `json:"ssoUserHeader,omitempty"`
	TlsCert                        string              `json:"tlsCert,omitempty"`
	TlsKey                         string              `json:"tlsKey,omitempty"`
}
