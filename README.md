# Go

## Overview

Developed originally by Google...
Docker and Kubernetes are written in Go...
Go can be compiled into a binary...

## Getting started

### Installing on Linux

```Bash
# Download
wget https://go.dev/dl/go1.19.2.linux-amd64.tar.gz -O ~/go.tar.gz

# Install
sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf ~/go.tar.gz
```

```Bash
# Add Go to .bashrc
vim ~/.bashrc

# Add this to .bashrc
export PATH=$PATH:/usr/local/go/bin

# press escape w q to write and save

# Reload bashrc
source ~/.bashrc
```

### Installing on Windows

```PowerShell
# Download
Invoke-Webrequest -uri 'https://go.dev/dl/go1.19.2.windows-amd64.msi' -OutFile '~/go.msi'

# Install
.\go.msi /passive
```