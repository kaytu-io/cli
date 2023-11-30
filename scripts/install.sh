#!/bin/bash

sudo -v
curl -s -H "Accept: application/vnd.github.v3+json" https://api.github.com/repos/kaytu-io/cli/releases/latest | jq -r .assets[].browser_download_url | grep linux | grep amd | xargs wget -O ./kaytu -qi -
chmod +x ./kaytu
mv ./kaytu /usr/bin/