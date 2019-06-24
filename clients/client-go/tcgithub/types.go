// This source code file is AUTO-GENERATED by github.com/taskcluster/jsonschema2go

package tcgithub

import (
	tcclient "github.com/taskcluster/taskcluster/clients/client-go"
)

type (
	// See https://taskcluster-staging.net/schemas/github/v1/build-list.json#/properties/builds/items
	Build struct {

		// The initial creation time of the build. This is when it became pending.
		//
		// See https://taskcluster-staging.net/schemas/github/v1/build-list.json#/properties/builds/items/properties/created
		Created tcclient.Time `json:"created"`

		// The GitHub webhook deliveryId. Extracted from the header 'X-GitHub-Delivery'
		//
		// One of:
		//   * GithubGUID
		//   * UnknownGithubGUID
		//
		// See https://taskcluster-staging.net/schemas/github/v1/build-list.json#/properties/builds/items/properties/eventId
		EventID string `json:"eventId"`

		// Type of Github event that triggered the build (i.e. push, pull_request.opened).
		//
		// See https://taskcluster-staging.net/schemas/github/v1/build-list.json#/properties/builds/items/properties/eventType
		EventType string `json:"eventType"`

		// Github organization associated with the build.
		//
		// Syntax:     ^([a-zA-Z0-9-_%]*)$
		// Min length: 1
		// Max length: 100
		//
		// See https://taskcluster-staging.net/schemas/github/v1/build-list.json#/properties/builds/items/properties/organization
		Organization string `json:"organization"`

		// Github repository associated with the build.
		//
		// Syntax:     ^([a-zA-Z0-9-_%]*)$
		// Min length: 1
		// Max length: 100
		//
		// See https://taskcluster-staging.net/schemas/github/v1/build-list.json#/properties/builds/items/properties/repository
		Repository string `json:"repository"`

		// Github revision associated with the build.
		//
		// Min length: 40
		// Max length: 40
		//
		// See https://taskcluster-staging.net/schemas/github/v1/build-list.json#/properties/builds/items/properties/sha
		Sha string `json:"sha"`

		// Github status associated with the build.
		//
		// Possible values:
		//   * "pending"
		//   * "success"
		//   * "error"
		//   * "failure"
		//
		// See https://taskcluster-staging.net/schemas/github/v1/build-list.json#/properties/builds/items/properties/state
		State string `json:"state"`

		// Taskcluster task-group associated with the build.
		//
		// Syntax:     ^[A-Za-z0-9_-]{8}[Q-T][A-Za-z0-9_-][CGKOSWaeimquy26-][A-Za-z0-9_-]{10}[AQgw]$
		//
		// See https://taskcluster-staging.net/schemas/github/v1/build-list.json#/properties/builds/items/properties/taskGroupId
		TaskGroupID string `json:"taskGroupId"`

		// The last updated of the build. If it is done, this is when it finished.
		//
		// See https://taskcluster-staging.net/schemas/github/v1/build-list.json#/properties/builds/items/properties/updated
		Updated tcclient.Time `json:"updated"`
	}

	// A paginated list of builds
	//
	// See https://taskcluster-staging.net/schemas/github/v1/build-list.json#
	BuildsResponse struct {

		// A simple list of builds.
		//
		// See https://taskcluster-staging.net/schemas/github/v1/build-list.json#/properties/builds
		Builds []Build `json:"builds"`

		// Passed back from Azure to allow us to page through long result sets.
		//
		// See https://taskcluster-staging.net/schemas/github/v1/build-list.json#/properties/continuationToken
		ContinuationToken string `json:"continuationToken,omitempty"`
	}

	// Write a new comment on a GitHub Issue or Pull Request.
	// Full specification on [GitHub docs](https://developer.github.com/v3/issues/comments/#create-a-comment)
	//
	// See https://taskcluster-staging.net/schemas/github/v1/create-comment.json#
	CreateCommentRequest struct {

		// The contents of the comment.
		//
		// See https://taskcluster-staging.net/schemas/github/v1/create-comment.json#/properties/body
		Body string `json:"body"`
	}

	// Create a commit status on GitHub.
	// Full specification on [GitHub docs](https://developer.github.com/v3/repos/statuses/#create-a-status)
	//
	// See https://taskcluster-staging.net/schemas/github/v1/create-status.json#
	CreateStatusRequest struct {

		// A string label to differentiate this status from the status of other systems.
		//
		// See https://taskcluster-staging.net/schemas/github/v1/create-status.json#/properties/context
		Context string `json:"context,omitempty"`

		// A short description of the status.
		//
		// See https://taskcluster-staging.net/schemas/github/v1/create-status.json#/properties/description
		Description string `json:"description,omitempty"`

		// The state of the status.
		//
		// Possible values:
		//   * "pending"
		//   * "success"
		//   * "error"
		//   * "failure"
		//
		// See https://taskcluster-staging.net/schemas/github/v1/create-status.json#/properties/state
		State string `json:"state"`

		// The target URL to associate with this status. This URL will be linked from the GitHub UI to allow users to easily see the 'source' of the Status.
		//
		// See https://taskcluster-staging.net/schemas/github/v1/create-status.json#/properties/target_url
		Target_URL string `json:"target_url,omitempty"`
	}

	// The GitHub webhook deliveryId. Extracted from the header 'X-GitHub-Delivery'
	//
	// Syntax:     ^[a-zA-Z0-9]{8}-[a-zA-Z0-9]{4}-[a-zA-Z0-9]{4}-[a-zA-Z0-9]{4}-[a-zA-Z0-9]{12}$
	//
	// See https://taskcluster-staging.net/schemas/github/v1/build-list.json#/properties/builds/items/properties/eventId/oneOf[0]
	GithubGUID string

	// Any Taskcluster-specific Github repository information.
	//
	// See https://taskcluster-staging.net/schemas/github/v1/repository.json#
	RepositoryResponse struct {

		// True if integration is installed, False otherwise.
		//
		// See https://taskcluster-staging.net/schemas/github/v1/repository.json#/properties/installed
		Installed bool `json:"installed"`
	}

	// The GitHub webhook deliveryId. Extracted from the header 'X-GitHub-Delivery'
	//
	// Possible values:
	//   * "Unknown"
	//
	// See https://taskcluster-staging.net/schemas/github/v1/build-list.json#/properties/builds/items/properties/eventId/oneOf[1]
	UnknownGithubGUID string
)
