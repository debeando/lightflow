package core

import (
	"fmt"

	"github.com/debeando/lightflow/common/log"
	"github.com/debeando/lightflow/plugins/evaluate"
	"github.com/debeando/lightflow/plugins/slack"
)

// Slack send custom message.
func (core *Core) slack() {
	expression := core.Render(core.GetSlackExpression())

	if evaluate.Expression(expression) {
		title := core.Render(core.GetSlackTitle())
		message := core.Render(core.GetSlackMessage())

		slack.Token = core.Config.General.Slack.Token
		err := slack.Send(
			core.GetSlackChannel(),
			title,
			message,
			core.GetSlackColor(),
		)
		if err != nil {
			log.Error(err.Error(), nil)
		}

		log.Info(
			fmt.Sprintf(
				"%s/%s Send message to slack.",
				core.TaskName(),
				core.PipeName(),
			),
			nil,
		)
	}
}
