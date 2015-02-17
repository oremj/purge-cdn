# purge-cdn [![Build Status](https://travis-ci.org/oremj/purge-cdn.svg?branch=master)](https://travis-ci.org/oremj/purge-cdn)
A tool to purge cache objects from multiple CDNs

## Installing
```bash
go get github.com/oremj/purge-cdn
```

## Usage

### Purging edgecast
```bash
EDGECAST_TOKEN="your-token"
EDGECAST_ACCOUNT_ID="account-id"

purge-cdn --url https://url-to-purge
```
