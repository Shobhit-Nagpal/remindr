# Remindr

Remindr is a recurring reminder application written in Go. It notifies users on Linux using the `libnotify` library and the `dunst` notification daemon. The application consists of a CLI tool and a server that runs as a systemd service.

## Features
- Set recurring reminders via CLI.
- Notifications delivered via `dunst` on Linux.
- Easy installation and setup with a single command.

## Dependencies
- **libnotify**: For sending desktop notifications.
- **dunst**: A lightweight notification daemon for Linux.

Install the dependencies on Ubuntu/Debian:
```bash
sudo apt-get install libnotify-bin dunst
```

---

## Installation

### **Option 1: Install Using the CLI (Recommended)**

1. **Install Go** (if not already installed):
   ```bash
   sudo apt-get install golang
   ```

2. **Install Remindr**:
   ```bash
   go install github.com/yourusername/remindr/cli@latest
   ```

3. **Run the Setup Command**:
   ```bash
   remindr setup
   ```

   This will:
   - Install the server binary to `/usr/local/bin/remindr-server`.
   - Create a systemd service file at `/etc/systemd/system/remindr.service`.
   - Enable and start the service.

4. **Verify the Service**:
   ```bash
   systemctl status remindr.service
   ```

---

### **Option 2: Manual Installation**

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/yourusername/remindr.git
   cd remindr
   ```

2. **Build the CLI and Server**:
   ```bash
   go build -o remindr ./cli
   go build -o remindr-server ./server
   ```

3. **Install the Binaries**:
   ```bash
   sudo mv remindr /usr/local/bin/
   sudo mv remindr-server /usr/local/bin/
   ```

4. **Create the Systemd Service**:
   Create a service file at `/etc/systemd/system/remindr.service`:
   ```bash
   sudo bash -c "cat > /etc/systemd/system/remindr.service" <<EOF
   [Unit]
   Description=Remindr Recurring Reminder Application
   After=network.target

   [Service]
   ExecStart=/usr/local/bin/remindr-server
   Restart=always
   User=$USER
   Environment=GO_ENV=production

   [Install]
   WantedBy=default.target
   EOF
   ```

5. **Reload Systemd and Start the Service**:
   ```bash
   sudo systemctl daemon-reload
   sudo systemctl enable remindr.service
   sudo systemctl start remindr.service
   ```

6. **Verify the Service**:
   ```bash
   systemctl status remindr.service
   ```

---

## Usage

### **Set a Reminder**
Use the CLI to set a recurring reminder. For example:
```bash
remindr create "Meeting in 10 minutes" --interval 600
```

### **List Reminders**
View all active reminders:
```bash
remindr list
```

### **Stop a Reminder**
Stop a reminder by its ID:
```bash
remindr stop 1
```

### **Run a Reminder**
Run a reminder by its ID:
```bash
remindr run 1
```

### **Kill a Reminder**
Kill a reminder by its ID:
```bash
remindr kill 1
```

---

## Troubleshooting

### **Notifications Not Showing**
Ensure `dunst` is running:
```bash
dunst &
```

### **Service Fails to Start**
Check the logs for errors:
```bash
journalctl -u remindr.service
```

---

## Contributing
Contributions are welcome! Please open an issue or submit a pull request.

---

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---
