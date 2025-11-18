### Feature 
- [x] wake up a remote pc (that remote pc using linux) 
- [x] check if remote pc is online or offline
- [x] turn off remote pc
- [x] ssh to remote pc 
- [x] disconnect ssh from remote pc
- [x] run a specific program inside remote pc 
- [x] check if the program is running or not
- [x] stop that specific program on remote pc
- [x] create --help command to show all available commands
- [ ] TUI that support multiple client

### CLI Commands
1. Check all available the command
```bash
wolremote --help
```
2. Wake-On-LAN
```bash
wolremote on --remote "config_1.toml"
```

3. Check Remote PC Status
```bash
wolremote status --remote "config_1.toml"
```

4. Turn Off Remote Comp
```bash
wolremote off --remote "config_1.toml"
```

5. Connect via SSH
```bash
wolremote connect --remote "config_1.toml"
```

6. Disconnect from SSH
```bash
wolremote disconnect --remote "config_1.toml"
```

7. Run Program
```bash
wolremote run --remote "config_1.toml"
```

8. Check Program
```bash
wolremote check --remote "config_1.toml"
```

9. Stop Program
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
Commands: 
[Arrow Up] Move up [Arrow Down] Move Down
[P] Power ON  [O] Power OFF 
[R] Start / Restart Program [S] Stop Program 
[Q] Quit
```

### Notes:
command to check in remote pc
```bash
ps aux | grep wolidar-linux-amd64
```
