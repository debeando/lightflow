# LightFlow

A flexible, light, easy to use, automation framework for typical data manipulation. With this tool you can define many bash command with special functions, for example you can complete with template variables, loops, chunks, autoincrement, retry and more. Maybe this tool not have sense, but please check with your eyes.

## Components:

![Flow](https://raw.githubusercontent.com/debeando/lightflow/master/assets/flow.png)

- **Tasks:** Collections of subtask and pipes.
- **Subtask:** Collections of pipes with own variables.
- **Pipes:** Collections of commands to execute in the bash with many abilities;
	- **AutoIncrement:** For the moment only apply for date, you can define range.
	- **Chunks:** Split pipe in many chunks.
	- **Register:** You can save stdout into variable, and if JSON you can convert all first level into variables.
	- **Retry:** Specific pipe when satisfy a condition expression to continue or not the pipe.
  - **Skip:** Specific pipe when satisfy a condition expression to continue or not the subtask.
- **Variables:** You can use environment variables or define own variables to use in the template, subtask and pipes and between them.
- **Template:** You can build a command with many variables, defined in the subtask, pipes or register.

## Configuration:

```yaml
---
general:
  tmp_dir: /tmp
tasks:
  - name: T1
    subtask:
      - name: T1ST1
        variables:
          V1: T1ST1.A
          V2: T1ST1.A
      - name: T1ST2
        variables:
          V1: T1ST2.A
          V2: T1ST2.B
    pipes:
      # Print variables:
      - name: T1P0
        execute: echo -e "T1P0 {{ .date }} {{ .year }}{{ .month }}{{ .day }}"
        ignore: true
        variables:
          V3: T1P1.C
      - name: T1P1
        execute: echo -e "T1P1 {{ .V1 }}-{{ .V2 }}-{{ .V3 }}"
        variables:
          V3: T1P1.C

      # Retry command if exit code is different from definition:
      - name: T1P2
        execute: exit 1
        retry:
          attempts: 3
          wait: 1
          expression: "{{ .exit_code }} != 1"
      - name: T1P4
        execute: exit 1
        retry:
          attempts: 3
          wait: 1
          expression: "{{ .exit_code }} != 0"

      # Curl to identify external IP address and
      # save JSON output from stdout to variables and
      # retry if fail request:
      - name: curl
        execute: curl --silent ifconfig.me/all.json
        format: JSON
        retry:
          attempts: 1
          wait: 1
      - name: print
        execute: echo -n "{{ .ip_addr }}"

      # Chunk example:
      - name: T2P4
        execute: echo -e "LIMIT {{ .offset }},{{ .limit }};"
        chunk:
          total: 9
          limit: 2
```
