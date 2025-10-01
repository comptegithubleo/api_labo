#!/usr/bin/env bash

grep -oE 'vlan[0-9]{1,4}' /etc/network/interfaces \
  | grep -oE '[0-9]{1,4}' \
  | sort -un