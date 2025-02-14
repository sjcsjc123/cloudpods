// Copyright 2019 Yunion
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package notify

import (
	"yunion.io/x/onecloud/cmd/climc/shell"
	modules "yunion.io/x/onecloud/pkg/mcclient/modules/notify"
	"yunion.io/x/onecloud/pkg/mcclient/options/notify"
)

func init() {
	cmd := shell.NewResourceCmd(&modules.NotifyTopic).WithKeyword("notify-topic")
	cmd.List(new(notify.TopicListOptions))

	cmd1 := shell.NewResourceCmd(&modules.NotifySubscriber).WithKeyword("notify-subscriber")
	cmd1.List(new(notify.SubscriberListOptions))
	cmd1.Create(new(notify.SubscriberCreateOptions))
	cmd1.Show(new(notify.SubscriberOptions))
	cmd1.Delete(new(notify.SubscriberOptions))
	cmd1.Perform("change", new(notify.SubscriberChangeOptions))
	cmd1.Perform("enable", new(notify.SubscriberOptions))
	cmd1.Perform("disable", new(notify.SubscriberOptions))
}
