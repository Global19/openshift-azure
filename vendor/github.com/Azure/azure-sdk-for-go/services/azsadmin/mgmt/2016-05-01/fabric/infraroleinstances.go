package fabric

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
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// InfraRoleInstancesClient is the fabric Admin Client
type InfraRoleInstancesClient struct {
	BaseClient
}

// NewInfraRoleInstancesClient creates an instance of the InfraRoleInstancesClient client.
func NewInfraRoleInstancesClient(subscriptionID string) InfraRoleInstancesClient {
	return NewInfraRoleInstancesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewInfraRoleInstancesClientWithBaseURI creates an instance of the InfraRoleInstancesClient client.
func NewInfraRoleInstancesClientWithBaseURI(baseURI string, subscriptionID string) InfraRoleInstancesClient {
	return InfraRoleInstancesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get get an infra role instance.
//
// location is location of the resource. infraRoleInstance is name of an infra role instance.
func (client InfraRoleInstancesClient) Get(ctx context.Context, location string, infraRoleInstance string) (result InfraRoleInstance, err error) {
	req, err := client.GetPreparer(ctx, location, infraRoleInstance)
	if err != nil {
		err = autorest.NewErrorWithError(err, "fabric.InfraRoleInstancesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "fabric.InfraRoleInstancesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "fabric.InfraRoleInstancesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client InfraRoleInstancesClient) GetPreparer(ctx context.Context, location string, infraRoleInstance string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"infraRoleInstance": autorest.Encode("path", infraRoleInstance),
		"location":          autorest.Encode("path", location),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/System.{location}/providers/Microsoft.Fabric.Admin/fabricLocations/{location}/infraRoleInstances/{infraRoleInstance}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client InfraRoleInstancesClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client InfraRoleInstancesClient) GetResponder(resp *http.Response) (result InfraRoleInstance, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List get a list of all infra role instances at a location.
//
// location is location of the resource. filter is oData filter parameter.
func (client InfraRoleInstancesClient) List(ctx context.Context, location string, filter string) (result InfraRoleInstanceListPage, err error) {
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, location, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "fabric.InfraRoleInstancesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.iril.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "fabric.InfraRoleInstancesClient", "List", resp, "Failure sending request")
		return
	}

	result.iril, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "fabric.InfraRoleInstancesClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client InfraRoleInstancesClient) ListPreparer(ctx context.Context, location string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/System.{location}/providers/Microsoft.Fabric.Admin/fabricLocations/{location}/infraRoleInstances", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client InfraRoleInstancesClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client InfraRoleInstancesClient) ListResponder(resp *http.Response) (result InfraRoleInstanceList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client InfraRoleInstancesClient) listNextResults(lastResults InfraRoleInstanceList) (result InfraRoleInstanceList, err error) {
	req, err := lastResults.infraRoleInstanceListPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "fabric.InfraRoleInstancesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "fabric.InfraRoleInstancesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "fabric.InfraRoleInstancesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client InfraRoleInstancesClient) ListComplete(ctx context.Context, location string, filter string) (result InfraRoleInstanceListIterator, err error) {
	result.page, err = client.List(ctx, location, filter)
	return
}

// PowerOff power off an infra role instance.
//
// location is location of the resource. infraRoleInstance is name of an infra role instance.
func (client InfraRoleInstancesClient) PowerOff(ctx context.Context, location string, infraRoleInstance string) (result InfraRoleInstancesPowerOffFuture, err error) {
	req, err := client.PowerOffPreparer(ctx, location, infraRoleInstance)
	if err != nil {
		err = autorest.NewErrorWithError(err, "fabric.InfraRoleInstancesClient", "PowerOff", nil, "Failure preparing request")
		return
	}

	result, err = client.PowerOffSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "fabric.InfraRoleInstancesClient", "PowerOff", result.Response(), "Failure sending request")
		return
	}

	return
}

// PowerOffPreparer prepares the PowerOff request.
func (client InfraRoleInstancesClient) PowerOffPreparer(ctx context.Context, location string, infraRoleInstance string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"infraRoleInstance": autorest.Encode("path", infraRoleInstance),
		"location":          autorest.Encode("path", location),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/System.{location}/providers/Microsoft.Fabric.Admin/fabricLocations/{location}/infraRoleInstances/{infraRoleInstance}/PowerOff", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PowerOffSender sends the PowerOff request. The method will close the
// http.Response Body if it receives an error.
func (client InfraRoleInstancesClient) PowerOffSender(req *http.Request) (future InfraRoleInstancesPowerOffFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusInternalServerError))
	return
}

// PowerOffResponder handles the response to the PowerOff request. The method always
// closes the http.Response Body.
func (client InfraRoleInstancesClient) PowerOffResponder(resp *http.Response) (result OperationStatus, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusInternalServerError),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// PowerOn power on an infra role instance.
//
// location is location of the resource. infraRoleInstance is name of an infra role instance.
func (client InfraRoleInstancesClient) PowerOn(ctx context.Context, location string, infraRoleInstance string) (result InfraRoleInstancesPowerOnFuture, err error) {
	req, err := client.PowerOnPreparer(ctx, location, infraRoleInstance)
	if err != nil {
		err = autorest.NewErrorWithError(err, "fabric.InfraRoleInstancesClient", "PowerOn", nil, "Failure preparing request")
		return
	}

	result, err = client.PowerOnSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "fabric.InfraRoleInstancesClient", "PowerOn", result.Response(), "Failure sending request")
		return
	}

	return
}

// PowerOnPreparer prepares the PowerOn request.
func (client InfraRoleInstancesClient) PowerOnPreparer(ctx context.Context, location string, infraRoleInstance string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"infraRoleInstance": autorest.Encode("path", infraRoleInstance),
		"location":          autorest.Encode("path", location),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/System.{location}/providers/Microsoft.Fabric.Admin/fabricLocations/{location}/infraRoleInstances/{infraRoleInstance}/PowerOn", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PowerOnSender sends the PowerOn request. The method will close the
// http.Response Body if it receives an error.
func (client InfraRoleInstancesClient) PowerOnSender(req *http.Request) (future InfraRoleInstancesPowerOnFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusInternalServerError))
	return
}

// PowerOnResponder handles the response to the PowerOn request. The method always
// closes the http.Response Body.
func (client InfraRoleInstancesClient) PowerOnResponder(resp *http.Response) (result OperationStatus, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusInternalServerError),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Reboot reboot an infra role instance.
//
// location is location of the resource. infraRoleInstance is name of an infra role instance.
func (client InfraRoleInstancesClient) Reboot(ctx context.Context, location string, infraRoleInstance string) (result InfraRoleInstancesRebootFuture, err error) {
	req, err := client.RebootPreparer(ctx, location, infraRoleInstance)
	if err != nil {
		err = autorest.NewErrorWithError(err, "fabric.InfraRoleInstancesClient", "Reboot", nil, "Failure preparing request")
		return
	}

	result, err = client.RebootSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "fabric.InfraRoleInstancesClient", "Reboot", result.Response(), "Failure sending request")
		return
	}

	return
}

// RebootPreparer prepares the Reboot request.
func (client InfraRoleInstancesClient) RebootPreparer(ctx context.Context, location string, infraRoleInstance string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"infraRoleInstance": autorest.Encode("path", infraRoleInstance),
		"location":          autorest.Encode("path", location),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/System.{location}/providers/Microsoft.Fabric.Admin/fabricLocations/{location}/infraRoleInstances/{infraRoleInstance}/Reboot", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RebootSender sends the Reboot request. The method will close the
// http.Response Body if it receives an error.
func (client InfraRoleInstancesClient) RebootSender(req *http.Request) (future InfraRoleInstancesRebootFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusInternalServerError))
	return
}

// RebootResponder handles the response to the Reboot request. The method always
// closes the http.Response Body.
func (client InfraRoleInstancesClient) RebootResponder(resp *http.Response) (result OperationStatus, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusInternalServerError),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Shutdown shut down an infra role instance.
//
// location is location of the resource. infraRoleInstance is name of an infra role instance.
func (client InfraRoleInstancesClient) Shutdown(ctx context.Context, location string, infraRoleInstance string) (result InfraRoleInstancesShutdownFuture, err error) {
	req, err := client.ShutdownPreparer(ctx, location, infraRoleInstance)
	if err != nil {
		err = autorest.NewErrorWithError(err, "fabric.InfraRoleInstancesClient", "Shutdown", nil, "Failure preparing request")
		return
	}

	result, err = client.ShutdownSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "fabric.InfraRoleInstancesClient", "Shutdown", result.Response(), "Failure sending request")
		return
	}

	return
}

// ShutdownPreparer prepares the Shutdown request.
func (client InfraRoleInstancesClient) ShutdownPreparer(ctx context.Context, location string, infraRoleInstance string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"infraRoleInstance": autorest.Encode("path", infraRoleInstance),
		"location":          autorest.Encode("path", location),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/System.{location}/providers/Microsoft.Fabric.Admin/fabricLocations/{location}/infraRoleInstances/{infraRoleInstance}/Shutdown", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ShutdownSender sends the Shutdown request. The method will close the
// http.Response Body if it receives an error.
func (client InfraRoleInstancesClient) ShutdownSender(req *http.Request) (future InfraRoleInstancesShutdownFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusInternalServerError))
	return
}

// ShutdownResponder handles the response to the Shutdown request. The method always
// closes the http.Response Body.
func (client InfraRoleInstancesClient) ShutdownResponder(resp *http.Response) (result OperationStatus, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusInternalServerError),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}