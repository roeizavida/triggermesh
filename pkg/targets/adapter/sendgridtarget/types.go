/*
Copyright (c) 2021 TriggerMesh Inc.

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

package sendgridtarget

import "net/url"

// EmailMessage targeting Sendgrid
type EmailMessage struct {
	Message   string     `json:"message,omitempty"`
	MediaURLs []*url.URL `json:"media_urls,omitempty"`
	FromName  string     `json:"fromname,omitempty"`
	FromEmail string     `json:"fromemail,omitempty"`
	ToName    string     `json:"toname,omitempty"`
	ToEmail   string     `json:"toemail,omitempty"`
	Subject   string     `json:"subject,omitempty"`
}
