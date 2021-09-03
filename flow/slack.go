package flow

import (
	"fmt"

	"github.com/debeando/lightflow/common/log"
	"github.com/debeando/lightflow/flow/evaluate"
	"github.com/debeando/lightflow/plugins/slack"
)

// Slack send custom message.
func (f *Flow) slack() {
	expression := f.Render(f.GetSlackExpression())

	if evaluate.Expression(expression) {
		title := f.Render(f.GetSlackTitle())
		message := f.Render(f.GetSlackMessage())

		slack.Token = f.Config.General.Slack.Token
		err := slack.Send(
			f.GetSlackChannel(),
			title,
			message,
			f.GetSlackColor(),
		)
		if err != nil {
			log.Error(err.Error(), nil)
		}

		log.Info(
			fmt.Sprintf(
				"%s/%s Send message to slack.",
				f.TaskName(),
				f.PipeName(),
			),
			nil,
		)
	}
}
