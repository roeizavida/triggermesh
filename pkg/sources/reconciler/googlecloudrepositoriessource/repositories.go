/*
Copyright 2021 TriggerMesh Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package googlecloudrepositoriessource

import (
	"context"

	gsourcerepo "google.golang.org/api/sourcerepo/v1"

	corev1 "k8s.io/api/core/v1"
	"knative.dev/pkg/controller"
	"knative.dev/pkg/reconciler"

	"github.com/triggermesh/triggermesh/pkg/apis/sources/v1alpha1"
	"github.com/triggermesh/triggermesh/pkg/sources/reconciler/common/event"
	"github.com/triggermesh/triggermesh/pkg/sources/reconciler/common/skip"
)

// Ensures that the Repo has the topic associated.
// Required permissions:
// - source.repos.updateRepoConfig
func ensureTopicAssociated(ctx context.Context, cli *gsourcerepo.Service, topicResName *v1alpha1.GCloudResourceName) error {
	if skip.Skip(ctx) {
		return nil
	}

	src := v1alpha1.SourceFromContext(ctx).(*v1alpha1.GoogleCloudRepositoriesSource)
	status := &src.Status

	repoName := src.Spec.Repository.String()

	updateRepoRequest := &gsourcerepo.UpdateRepoRequest{
		Repo: &gsourcerepo.Repo{
			PubsubConfigs: map[string]gsourcerepo.PubsubConfig{
				topicResName.String(): {
					Topic:         topicResName.String(),
					MessageFormat: "JSON",
				},
			},
		},
		UpdateMask: "pubsubConfigs",
	}

	patchRepo := cli.Projects.Repos.Patch(repoName, updateRepoRequest)

	_, err := patchRepo.Do()
	switch {
	case isDenied(err):
		status.MarkNotSubscribed(v1alpha1.GCloudReasonAPIError,
			"Access denied to Cloud Source Repositories API: "+toErrMsg(err))
		return controller.NewPermanentError(failCreatingRepositories(repoName, err))
	case isNotFound(err):
		status.MarkNotSubscribed(v1alpha1.GCloudReasonAPIError,
			"Repo does not exists: "+toErrMsg(err))
		return controller.NewPermanentError(failCreatingRepositories(repoName, err))
	case err != nil:
		return reconciler.NewEvent(corev1.EventTypeWarning, ReasonFailedSubscribe,
			"Cannot create repo notification %q: %s", repoName, toErrMsg(err))
	}
	event.Normal(ctx, ReasonSubscribed, "Created Repo notification %q", repoName)

	return err
}

// ensureNoTopicAssociated looks at status.Repositories and if non-empty will delete it
// Required permissions:
// - source.repos.updateRepoConfig
func (r *Reconciler) ensureNoTopicAssociated(ctx context.Context, cli *gsourcerepo.Service) error {
	if skip.Skip(ctx) {
		return nil
	}

	src := v1alpha1.SourceFromContext(ctx).(*v1alpha1.GoogleCloudRepositoriesSource)
	status := &src.Status

	repoName := src.Spec.Repository.String()

	updateRepoRequest := &gsourcerepo.UpdateRepoRequest{
		Repo: &gsourcerepo.Repo{
			PubsubConfigs: map[string]gsourcerepo.PubsubConfig{},
		},
		UpdateMask: "pubsubConfigs",
	}

	patchRepo := cli.Projects.Repos.Patch(repoName, updateRepoRequest)

	_, err := patchRepo.Do()
	switch {
	case isDenied(err):
		status.MarkNotSubscribed(v1alpha1.GCloudReasonAPIError,
			"Access denied to Cloud Source Repositories API: "+toErrMsg(err))
		return controller.NewPermanentError(failCreatingRepositories(repoName, err))
	case isNotFound(err):
		status.MarkNotSubscribed(v1alpha1.GCloudReasonAPIError,
			"Repo does not exist: "+toErrMsg(err))
		return controller.NewPermanentError(failCreatingRepositories(repoName, err))
	case err != nil:
		return reconciler.NewEvent(corev1.EventTypeWarning, ReasonFailedSubscribe,
			"Cannot delete Repo notification %q: %s", repoName, toErrMsg(err))
	}
	event.Normal(ctx, ReasonSubscribed, "Deleted Repo notification %q", repoName)

	return err
}

// failCreatingRepositories returns a reconciler event which indicates
// that a Repo could not be retrieved or created from the
// Google Cloud API.
func failCreatingRepositories(repoName string, origErr error) reconciler.Event {
	return reconciler.NewEvent(corev1.EventTypeWarning, ReasonFailedSubscribe,
		"Error creating Repo Notification %q: %s", repoName, toErrMsg(origErr))
}