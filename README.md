<!-- File: README.md -->
# server-gen

**Author:** Bocaletto Luca (@bocaletto-luca)  
**Description:** Server Generator is Collects system information (IP, OS, CPU, memory, users) and sends templated email reports on a configurable schedule.

## Installation

    git clone https://github.com/bocaletto-luca/server-gen.git
    cd server-gen
    go build ./cmd/server-gen

## Configuration

#### Edit 
    configs/config.yaml:

schedule: "@every 30m"
smtp:
  host: "smtp.example.com"
  port: 587
  username: "alert@example.com"
  password: "supersecret"
  from: "alert@example.com"
  to:
    - "admin@example.com"
modules:
  - ip
  - os
  - cpu
  - mem
  - users

## Usage
    ./server-gen --config configs/config.yaml

    Press Ctrl+C to stop.

# File: .gitignore
    /server-gen
    /configs/*.lock
    *.exe
    *.log
