/*
Copyright 2018 Caicloud Authors
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

package tensorflow

type Job interface {
	Action() Event
}

type Event struct {
	Action ActionType
	Number int
}

type ActionType string

const (
	// ActionShouldAddPS is the action to add PS pods.
	ActionShouldAddPS ActionType = "ShouldAddPS"
	// ActionShouldAddPSService is the action to add PS services.
	ActionShouldAddPSService ActionType = "ShouldAddPSService"
	// ActionShouldAddWorker is the action to add worker pods.
	ActionShouldAddWorker ActionType = "ShouldAddWorker"
	// ActionShouldAddWorkerService is the action to add worker services.
	ActionShouldAddWorkerService ActionType = "ShouldAddWorkerService"
	// ActionShouldDelete is the action to delete something.
	ActionShouldDelete ActionType = "ShouldDelete"
	// ActionNothing is the action to do nothing.
	ActionNothing ActionType = "Nothing"
)
