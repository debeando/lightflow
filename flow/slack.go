package flow

import (
	"fmt"

	"github.com/debeando/lightflow/common/log"
	"github.com/debeando/lightflow/flow/evaluate"
	"github.com/debeando/lightflow/flow/slack"
)

// Slack send custom message.
func (f *Flow) slack() {
	expression := f.Render(f.GetSlackExpression())

	if evaluate.Expression(expression) {
		title := f.Render(f.GetSlackTitle())
		message := f.Render(f.GetSlackMessage())

		slack.Token = f.Config.General.Slack.Token
		slack.Send(
			f.GetSlackChannel(),
			title,
			message,
			f.GetSlackColor(),
		)

		log.Info(
			fmt.Sprintf(
				"%s/%s/%s Message sent to Slack",
				f.TaskName(),
				f.SubTaskName(),
				f.PipeName(),
			),
			nil,
		)
	}
}