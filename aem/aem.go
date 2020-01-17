// © Copyright 2017 XumaK, LLC. All rights reserved.
// Do not distribute.

package aem

import (
	"net/http"
	"net/url"
)

const (
	libraryVersion  = "1"
	userAgent       = "go-aem/" + libraryVersion
	defaultBasePort = "4502"
	defaultBaseURL  = "http://localhost" + defaultBasePort + "/"
)

// Client managed communication with Adobe CQ/AEM.
type Client struct {
	// HTTP client used to communicate with Sling.
	client *http.Client

	// Username is the user to authenticate for Adobe AEM's API.
	//
	// The default is 'admin'
	Username string
	// Password is the credential password to autneticate for Adobe AEM's API.
	//
	// The deafult is 'admin'
	Password string

	// Base URL for Adobe AEM API requests. Defaults to the localhost:4502.
	// Can be set to alternative domain and port (:4503 for publishers).
	BaseURL *url.URL

	// UserAgent used when communication with Adobe AEM API.
	// Default installations allow only for {"Firefox", "Shiretoko", "Chrome",
	// "Opera", "curl", "Wget" }
	// Display acceptable UserAgents:
	// http://localhost:4502/system/console/configMgr/com.day.cq.wcm.foundation.impl.HTTPAuthHandler
	//
	// The default is 'curl'
	UserAgent string

	// CSRFUserAgent used for CSRF token of modification requests.
	// Default installations allow only for {"Apache-HttpClient/*",
	// "Jakarta Commons-HttpClient/*", "Jakarta-Commons-VFS", "curl/*", "Wget/*",
	// "WebDAVFS/*", "Microsoft-WebDAV-MiniRedir/*", "Ruby", "Adobe-Campaign/*",
	// "Forms-Mobile *"}
	// Display acceptable Safe User Agents:
	// http://localhost:4502/system/console/configMgr/com.adobe.granite.csrf.impl.CSRFFilter
	//
	// The default is 'curl/*'
	CSRFUserAgent string

	// common prevents allocating one struct for each service on the heap.
	common service

	Bundles     *BundlesService
	Datastore   *DatastoreService
	Nodes       *NodesService
	Packages    *PackagesService
	Replication *ReplicationService
	Users       *UsersService
	Workflows   *WorkflowsService
}

type service struct {
	client *Client
}

// UploadOptions specifies  parameters to methods that support uploads.
type UploadOptions struct {
	Name string `url:"name,omitempty"`
}

// NewClient returns a new Adobe AEM client. If a nil httpClient is
// provided, http.DefaultClient will be used. To use API methods which require
// authentication, provide an http.Client that will perform the authentication
// for you (such as that provided by the golang.org/x/oauth2 library).
func NewClient(httpClient *http.Client, user, pass string) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	baseURL, _ := url.Parse(defaultBaseURL)
	//uploadURL, _ := url.Parse(uploadBaseURL)

	c := &Client{
		client:    httpClient,
		BaseURL:   baseURL,
		UserAgent: userAgent,
		Username:  user,
		Password:  pass,
	}

	c.common.client = c
	c.Bundles = (*BundlesService)(&c.common)
	c.Datastore = (*DatastoreService)(&c.common)
	c.Nodes = (*NodesService)(&c.common)
	c.Packages = (*PackagesService)(&c.common)
	c.Users = (*UsersService)(&c.common)
	c.Workflows = (*WorkflowsService)(&c.common)

	if c.UserAgent == "" {
		c.UserAgent = "curl"
	}

	if c.Username == "" {
		c.Username = "admin"
	}

	if c.Password == "" {
		c.Password = "admin"
	}

	return c
}

func (c *Client) Do() error {
	// new client request
	// set user agent
	/*
		AEM_SCHEME=http
		AEM_HOST=localhost
		AEM_PORT=4502
		AEM_LOGIN=admin:admin
		REFERER=${AEM_SCHEME}://${AEM_HOST}:${AEM_PORT}
		SERVICE_URL=/libs/granite/security/post/authorizables
		SERVICE_TOKEN=/libs/granite/csrf/token.json

		AEM_TOKEN="$(curl -s -H User-Agent:curl -H Referer:${REFERER} -u ${AEM_LOGIN} ${REFERER}${SERVICE_TOKEN}  | sed -e 's/[{"token":}]/''/g')"

		echo "CSRF Token: ${AEM_TOKEN}"

		HEADERS=" -u ${AEM_LOGIN} -H User-Agent:curl -H CSRF-Token:${AEM_TOKEN} -H Referer:${REFERER}"

		echo "Headers: $HEADERS"

		curl ${HEADERS} -FcreateGroup= -FauthorizableId=testGroup1 ${REFERER}${SERVICE_URL}
	*/
	return _, nil
}
