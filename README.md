# Remindr

Remindr is a recurring reminder application written in Go. It notifies users on Linux using the `libnotify` library and the `dunst` notification daemon. The application consists of a CLI tool and a server that runs as a systemd user service.

## Features
- Set recurring reminders via CLI
- Notifications delivered via `dunst` on Linux
- Easy installation and setup with a single command
- Runs as a user-level systemd service

## Dependencies
- **libnotify**: For sending desktop notifications
- **dunst**: A lightweight notification daemon for Linux

Install the dependencies on Ubuntu/Debian:
```bash
sudo apt-get install libnotify-bin dunst
```

## Installation

### Install Using Go

1. **Install Go** (if not already installed):
   ```bash
   sudo apt-get install golang
   ```

2. **Install Remindr**:
   ```bash
   go install github.com/Shobhit-Nagpal/remindr/cmd/remindr@latest
   ```

3. **Setup the Service**:
   ```bash
   remindr setup /path/to/working/directory
   ```
   This will:
   - Create a systemd user service file at `~/.config/systemd/user/remindr.service`
   - Start the service using `go run`

4. **Verify the Service**:
   ```bash
   systemctl --user status remindr.service
   ```

### Manual Installation

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/Shobhit-Nagpal/remindr.git
   cd remindr
   ```

2. **Create the Systemd User Service**:
   Create a service file at `~/.config/systemd/user/remindr.service`:
   ```ini
   [Unit]
   Description=Remindr Service
   After=default.target

   [Service]
   Environment=PATH=/usr/local/go/bin:/usr/local/bin:/usr/bin:/bin
   ExecStart=go run /path/to/remindr/server/main.go
   WorkingDirectory=/path/to/remindr/server
   Restart=on-failure

   [Install]
   WantedBy=default.target
   ```

3. **Enable and Start the Service**:
   ```bash
   systemctl --user daemon-reload
   systemctl --user enable remindr.service
   systemctl --user start remindr.service
   ```

## Usage

### Set a Reminder
Use the CLI to set a recurring reminder:
```bash
remindr create "Meeting in 10 minutes" --interval 600
```

### List Reminders
View all active reminders:
```bash
remindr list
```

### Stop a Reminder
Stop a reminder by its ID:
```bash
remindr stop <id>
```

### Run a Reminder
Run a reminder by its ID:
```bash
remindr run <id>
```

### Remove Service
To remove the Remindr service:
```bash
remindr destroy
```

## Troubleshooting

### Notifications Not Showing
Ensure `dunst` is running:
```bash
dunst &
```

### Service Fails to Start
Check the logs:
```bash
journalctl --user -u remindr.service
```

### Service Status
Check service status:
```bash
systemctl --user status remindr.service
```

## Contributing
Contributions are welcome! Please open an issue or submit a pull request.

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
