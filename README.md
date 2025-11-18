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

### Notes:
command to check in remote pc
```bash
ps aux | grep wolidar-linux-amd64
```

### TUI Reference
```
+---------------------------------------------------------------+
|                 REMOTE PC CONTROLLER (v2.0)                   |
+---------------------------------------------------------------+
| ID | Host          | Status    | Program      | Last Update   |
|----+---------------+-----------+--------------+---------------|
| 01 | 192.168.1.100 | ONLINE    | RUNNING      | 10s ago       |
| 02 | 192.168.1.101 | OFFLINE   | STOPPED      | 1m ago        |
| 03 | 192.168.1.102 | ONLINE    | STOPPED      | 5s ago        |
+---------------------------------------------------------------+
Commands: 
[↑] Move up
[↓] Move Down
[P] Power ON
[O] Power OFF 
[X] Start Program 
[S] Stop Program
[R] Refresh table
[Q] Quit
```

### Step by Step for building UI
- [ ] create table
- [ ] list all config(s) (indexing)
- [ ] pass config value to new table row (ID = index, IP = get from toml file)
- [ ] navigate row using arrow keyboard ` [Arrow Up | Arrow Down] `
- [ ] quit application ` [Q] `
- [ ] execute wake-on-lan based on row (will update STATUS and LAST UPATE column for that row) ` [P] `
- [ ] execute turn off based on row (will update STATUS, PROGRAM and LAST UPATE column, this also automatically call stop program command for that row) ` [O] `
- [ ] execute run program based on row (will update PROGRAM and LAST UPATE column for that row) ` [X] `
- [ ] execute stop program based on row (will update PROGRAM and LAST UPATE column for that row) ` [S] `
- [ ] execute check pc status and program globally (will update STATUS, PROGRAM and LAST UPATE column globally) ` [R] `
- [ ] check pc status and program globally when opening the app (will update STATUS, PROGRAM and LAST UPATE column globally)
- [ ] add hotkeys information

### Library
- bubbletea (core tui library) [https://github.com/charmbracelet/bubbletea]
- lipgloss (styling) [https://github.com/charmbracelet/lipgloss]
- bubbles (reference component) [https://github.com/charmbracelet/bubbles]
