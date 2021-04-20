# LightFlow

A flexible, light, easy to use, automation framework for typical data manipulation. With this tool you can define many bash command with special functions, for example you can complete with template variables, loops, chunks, autoincrement, retry and more. Maybe this tool not have sense, but please check with your eyes.

## Components:

![Flow](https://raw.githubusercontent.com/debeando/lightflow/master/assets/flow.png)

- **Tasks:** Collections of subtask and pipes.
- **Subtask:** Collections of pipes with own variables.
- **Pipes:** Collections of commands to execute in the bash with many abilities;
  - **MySQL:** Allow execute SELECT and save result into csv file.
  - **Interval:** Is auto increment or decrement values, for the moment only apply for date, you can define range.
  - **Chunks:** Split pipe in many chunks.
  - **Register:** You can save stdout into variable, and if JSON you can convert all first level into variables.
  - **Retry:** Specific pipe when satisfy a condition expression to continue or not the pipe.
  - **Skip:** Specific pipe when satisfy a condition expression to continue or not the subtask.
  - **Error:** You can define custom expression to evaluate pipe variables are error or not.
  - **Slack:** Send message to slack when satisfy a condition expression.
  - **Unset:** List of variables to unset every pipe loop.
  - **Print:** List of variables to print.
- **Variables:** You can use environment variables or define own variables to use in the template, subtask and pipes and between them.
- **Template:** You can build a command with many variables, defined in the subtask, pipes or register.

## Install tool:

For the moment, this tool only run in any Linux and MAC OS X distribution with 64 bits. Paste that at a Terminal prompt:

```bash
bash < <(curl -fsSL https://raw.githubusercontent.com/debeando/lightflow/master/scripts/install.sh)
```

## Configuration:

Your's tasks are definied in YAML format in one configuration file, please see this [example files](https://github.com/debeando/lightflow/tree/master/tests/flow).

To see all configuration options, please see comments on this [source code](https://github.com/debeando/lightflow/blob/master/config/structure.go).

## How to use it:

See usage with:

```
lightflow --help
```
