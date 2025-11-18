### Feature 
- [ ] wake up a remote pc (that remote pc using linux) 
- [ ] ssh to that remote pc 
- [ ] run a specific program inside that remote pc 
- [ ] stop that specific program on that remote pc 
- [ ] turn off that remote pc
- [ ] support multiple client

### CLI Commands
to check all the command
```bash
wolremote --help
```
Wake-On-LAN
```bash
wolremote on --remote "config_1.toml"
```

Check Remote PC Status
```bash
wolremote status --remote "config_1.toml"
```

Turn Off Remote Comp
```bash
wolremote off --remote "config_1.toml"
```

Connect via SSH
```bash
wolremote connect --remote "config_1.toml"
```

disconnect from SSH
```bash
wolremote disconnect --remote "config_1.toml"
```

Run Program
```bash
wolremote run --remote "config_1.toml"
```

Check Program
```bash
wolremote check --remote "config_1.toml"
```

Stop Program
```bash
wolremote stop --remote "config_1.toml"
```

### TUI Reference
```
+---------------------------------------------------------------+
|                 REMOTE PC CONTROLLER (v1.0)                   |
+---------------------------------------------------------------+
| ID | Host          | Status    | Program      | Last Update   |
|----+---------------+-----------+--------------+---------------|
| 01 | 192.168.1.100 | RUNNING   | ok           | 10s ago       |
| 02 | 192.168.1.101 | OFFLINE   | -            | -             |
| 03 | 192.168.1.102 | RUNNING   | exit code 1  | 5s ago        |
+---------------------------------------------------------------+
Commands: [P] Power ON  [O] Power OFF  [R] Restart Program
          [S] Stop Program  [L] Logs  [Q] Quit
          [Arrow Up] Move up [Arrow Down] Move Down
```

### Notes:
command to check in remote pc
```bash
ps aux | grep wolidar-linux-amd64
```
