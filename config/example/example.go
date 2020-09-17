package example

const ExampleFile = `---
general:
  tmp_dir: /tmp
tasks:
  - name: demo
    loops:
      - name: A1
        case: case1
      - name: A2
        case: case2
    pipes:
    - name: print-path
      command: >
        echo "PATH: {{ .path }}/{{ .date }}/{{ .task_name }}/{{ .pipe_name }}/{{ .name_test }}"
      variables:
        name_test: "{{ .date }}"
    - name: create-register
      command: >
        echo -e "\n$(date +%Y%m%d%H%M)\n"
      register: register_test
    - name: print-register
      command: 'echo -e "R: {{ .register_test }}"'
    - name: try-re-print-name-test-var
      command: echo -e "{{ .name_test }}"
    - name: re-print-register
      command: echo -e "{{ .register_test }}"
    - name: print-loop-variable
      command: echo -e "{{ .name }}"
    - name: curl
      command: curl --silent ifconfig.me/all.json
      format: JSON
`