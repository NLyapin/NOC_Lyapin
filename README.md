# Network Configuration Management System

The **Network Configuration Management System** is a Go-based tool designed to automate and manage the configuration of network devices. It integrates with version control (Git) to provide tracking and rollback capabilities for configuration changes, and supports interaction with devices over standard network protocols like SSH, SNMP, and Netconf.

## Features

- **Configuration Management**: Add, modify, and delete network configurations.
- **Version Control**: Track configuration changes and rollback to previous versions using Git.
- **Device Interaction**: Communicate with network devices over SSH, SNMP, and Netconf.
- **Logging**: Log all configuration operations for auditing and troubleshooting.

## Project Structure

The project follows Go best practices with a modular structure:
```

network-config-system/
├── cmd/
│ └── app/
│ └── main.go # Main entry point of the application
├── pkg/
│ └── config/
│ └── config.go # Configuration management logic
├── internal/
│ └── devices/
│ └── devices.go # Network device interaction logic
├── go.mod # Go module file
├── README.md # Project description and instructions
├── .gitignore # Git ignore file
```

### Directories:
- `cmd/`: Contains the main application entry point.
- `pkg/`: Contains reusable modules, such as configuration management.
- `internal/`: Contains modules specific to this application, such as device interaction.
  
## Prerequisites

- **Go**: The project requires [Go](https://golang.org/dl/) (version 1.16 or later) to be installed.
- **Git**: Ensure you have [Git](https://git-scm.com/) installed for version control functionality.

## Installation

1. Clone the repository:
```bash
git clone https://github.com/NLyapin/NOC_Lyapin.git
```
2. Navigate to the project directory:


```bash
cd network-config-system
```
 
3. Initialize Go modules (if required):


```bash
go mod tidy
```

## Usage 

You can start the application by running the following command:


```bash
go run cmd/app/main.go
```

The system will:

- Load network configurations from a specified file.

- Interact with a network device to apply configurations.

- Log any configuration changes or device interactions.

### Example Workflow 
 
1. **Load a Configuration File** :
The system will load a network configuration (e.g., `example-config.yaml`) using the `pkg/config` package. The configuration contains device details such as IP address, username, and password.
 
2. **Device Interaction** :
The system connects to a network device using the `internal/devices` package, simulating a connection via SSH or other protocols.
 
3. **Version Control** :
Any changes to configurations are logged in Git for tracking and rollback purposes.

## License 
This project is licensed under the MIT License. See the [LICENSE]()  file for more details.

## Contact 

If you have any questions or suggestions, feel free to reach out:
 
- **Email** : [NLyapinwork@bk.ru]()
 
- **GitHub** : [https://github.com/NLyapin](https://github.com/Nlyapin)