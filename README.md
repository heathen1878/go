# Go

I'm learning Go primarily to aid my understanding of Terratest and I think it'll be worthwhile from a career point-of-view.

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

# Update Path - this example is updating the PowerShell profile
notepad $PROFILE
# Update Path
$env:Path += ';C:\Program Files\Go\bin\'


# or using Chocolatey
choco install go
```
