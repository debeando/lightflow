package example

const ExampleFile = `---
tasks:
  - name: random
    subtask:
      - name: computer
        variables:
          url: https://random-data-api.com/api/computer/random_computer
    pipes:
      - name: check
        execute: >
          curl \
            --silent \
            -X GET "{{ .url }}"
        format: JSON
        unset:
          - error
        print:
          - error
          - type
        error: >
          {{ if .error }}
            {{ len .error }} > 0
          {{ end }}
        retry:
          attempts: 3
          wait: 1
          expression: '"{{ .type }}" != "server"'
  - name: ip
    subtask:
      - name: basic
    pipes:
      - name: curl
        execute: curl --silent ifconfig.me/all.json
        format: JSON
      - name: print
        execute: echo -n "{{ .forwarded }}"
        print:
          - forwarded
      - name: ip
        execute: curl --silent ifconfig.me/all.json | jq '.ip_addr' -r
        register: ip
        print:
          - ip
          - stdout
  - name: print
    subtask:
    - name: path
      case: "DATE ({{ .year }}-{{ .month }}-{{ .day }})"
    pipes:
    - name: print-path
      execute: >
        echo "PATH: {{ .path }}/{{ .date }}/{{ .task_name }}/{{ .pipe_name }}/{{ .name_test }}"
      variables:
        name_test: "{{ .date }}.json"
      print:
        - name_test
        - stdout
`
