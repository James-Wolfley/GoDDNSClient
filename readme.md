# GoDDNSClient

A lightweight command-line dynamic DNS client for updating Cloudflare DNS records when your public IP changes.

This project was built for self-hosted services running on residential internet connections where the external IP address may change over time. The client is designed to be simple, scriptable, and easy to run on a schedule through cron on Linux or Task Scheduler on Windows. It only updates records when the IP changes, unless forced manually.

## Features

- Updates Cloudflare DNS records from the command line
- Stores the last known IP in a local config file
- Avoids unnecessary updates when the IP has not changed
- Supports multiple domains in a single configuration
- Can be scheduled to run automatically on Linux or Windows
- Includes command-line options for configuration and maintenance

## Why I Built It

I wanted a simple dynamic DNS updater for self-hosted environments without relying on a heavier third-party client. This tool was made to fit into a home lab workflow and to work cleanly with scheduled execution.

## Current Support

- Cloudflare DNS

## Planned Improvements

- Additional DNS provider support
- Better logging and output formatting
- Improved validation and setup flow

## Build

Clone the repository and build with Go:

```bash
go build -o GoDDNSClient
```

## Quick Start

1. Create a Cloudflare API token with permission to edit DNS records.
2. Find the Zone ID for the domain you want to update.
3. Configure the application using `config.json` or command-line arguments.
4. Run the client manually or schedule it to run periodically.

## Configuration

Example:

```json
{
  "email": "Account Email Address",
  "token": "Account API Access Token",
  "current-ip": "0.0.0.0",
  "domains": [
    {
      "name": "www.yoursitehere.com",
      "zone": "The Site Zone"
    }
  ]
}
```

## Usage

Run:

```bash
./GoDDNSClient
```

## Example Commands

Set account info:

```bash
./GoDDNSClient -email you@example.com -token your_token_here
```

Add a site:

```bash
./GoDDNSClient -add-site -site-name www.example.com -site-zone your_zone_id
```

Force update:

```bash
./GoDDNSClient -force
```

## Automation

### Linux cron

```cron
*/5 * * * * /path/to/GoDDNSClient
```

### Windows Task Scheduler

Create a scheduled task to run the executable at your desired interval.

## Notes

Useful for home labs, self-hosted services, and environments where DNS must follow a changing public IP.
