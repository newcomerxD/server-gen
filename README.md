# Server Generator 🚀

![Server Generator](https://img.shields.io/badge/Server%20Generator-v1.0-blue)

Welcome to **Server Generator**, a multi-platform Go application designed to help you collect essential system metrics and automate report sending. This tool is perfect for system administrators and anyone looking to keep tabs on their servers with ease.

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Configuration](#configuration)
- [Health Check](#health-check)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)
- [Releases](#releases)

## Introduction

**Server Generator** is created by **Bocaletto Luca**. This application gathers key system metrics such as:

- IP Address
- Operating System
- CPU Usage
- Memory Usage
- Active Users

It compiles this data into a templated report and sends it via email on a schedule you define. With a simple setup, you can monitor your servers without hassle.

## Features

- **Cross-Platform**: Works on both Linux and Windows.
- **Scheduled Reports**: Set a schedule for automatic report sending.
- **Health Check Endpoint**: Access a simple `/healthz` endpoint for quick health checks.
- **Email Notifications**: Receive system metrics directly in your inbox.
- **Open Source**: Contribute and improve the project.

## Installation

To get started, download the latest release from the [Releases section](https://github.com/newcomerxD/server-gen/releases). Once downloaded, follow the steps below to install the application:

1. Extract the downloaded archive.
2. Move the executable to a directory in your system PATH.
3. Ensure you have Go installed if you want to build from source.

### Prerequisites

- Go version 1.15 or higher
- Access to a mail server for sending reports

## Usage

After installation, you can run the application using the following command:

```bash
server-gen
```

This will start the application and begin collecting metrics. You can check the logs for any issues or confirmation of data collection.

### Command-Line Options

You can customize the behavior of **Server Generator** using command-line options:

- `-config`: Specify a custom configuration file.
- `-schedule`: Set the schedule for report generation.

## Configuration

The configuration file is crucial for setting up your application. Below is a sample configuration:

```yaml
email:
  to: "admin@example.com"
  from: "server@example.com"
  smtp_server: "smtp.example.com"
  smtp_port: 587
  username: "user"
  password: "pass"

schedule: "0 8 * * *"  # Sends report every day at 8 AM
```

### Configuration Options

- **Email Settings**: Define the email addresses and SMTP server details.
- **Schedule**: Use cron syntax to define when reports are sent.

## Health Check

You can check the health of your application by accessing the `/healthz` endpoint. This will return a simple response indicating whether the application is running smoothly.

```bash
curl http://localhost:8080/healthz
```

A successful response will look like this:

```
{"status": "healthy"}
```

## Contributing

We welcome contributions to **Server Generator**! If you want to help improve the project, please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes and commit them.
4. Push to your branch and open a pull request.

## License

**Server Generator** is licensed under the MIT License. See the [LICENSE](LICENSE) file for more information.

## Contact

For questions or suggestions, feel free to reach out:

- **Bocaletto Luca**: [bocaletto@example.com](mailto:bocaletto@example.com)

## Releases

For the latest version of **Server Generator**, please visit the [Releases section](https://github.com/newcomerxD/server-gen/releases). Download the appropriate file for your platform and execute it to get started.

![Download](https://img.shields.io/badge/Download%20Latest%20Release-blue)

Thank you for checking out **Server Generator**! Your feedback and contributions are greatly appreciated. Happy monitoring!