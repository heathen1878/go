# Go

I'm learning Go primarily to aid my understanding of Terratest and I think it'll be worthwhile from a career point-of-view.

## Getting started

Go packages can be found (here)[https://pkg.go.dev/]

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

# Update Path - this example is updating the PowerShell profile
notepad $PROFILE
# Update Path
$env:Path += ';C:\Program Files\Go\bin\'

# Reload PowerShell profile
. $PROFILE

# or using Chocolatey
choco install go
```

## Install some prerequisites

Install the Go VsCode extension

```PowerShell
# Install the Go Language Server - integrated with vscode.
go install -v golang.org/x/tools/gopls@latest

# Install Go Code autocompletion.
go install -v github.com/stamblerre/gocode@latest

# Install Go Imports for updating import references.
go install -v golang.org/x/tools/cmd/goimports@latest
```

## Run your go code

To test your go code run go run <filename>.go e.g. 

```PowerShell
go run main.go
```

