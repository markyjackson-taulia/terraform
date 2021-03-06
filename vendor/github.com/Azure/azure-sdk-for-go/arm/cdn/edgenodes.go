package cdn

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator 1.0.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// EdgeNodesClient is the use these APIs to manage Azure CDN resources through
// the Azure Resource Manager. You must make sure that requests made to these
// resources are secure.
type EdgeNodesClient struct {
	ManagementClient
}

// NewEdgeNodesClient creates an instance of the EdgeNodesClient client.
func NewEdgeNodesClient(subscriptionID string) EdgeNodesClient {
	return NewEdgeNodesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewEdgeNodesClientWithBaseURI creates an instance of the EdgeNodesClient
// client.
func NewEdgeNodesClientWithBaseURI(baseURI string, subscriptionID string) EdgeNodesClient {
	return EdgeNodesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List lists all the edge nodes of a CDN service.
func (client EdgeNodesClient) List() (result EdgenodeResult, err error) {
	req, err := client.ListPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "cdn.EdgeNodesClient", "List", nil, "Failure preparing request")
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "cdn.EdgeNodesClient", "List", resp, "Failure sending request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.EdgeNodesClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client EdgeNodesClient) ListPreparer() (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Cdn/edgenodes"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client EdgeNodesClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client EdgeNodesClient) ListResponder(resp *http.Response) (result EdgenodeResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
