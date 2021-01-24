#!/bin/bash
_path=$(dirname "${BASH_SOURCE[0]}")
/usr/bin/go build -buildmode=plugin -o $_path/../nmap.so $_path/nmap.go
echo "Loaded"
