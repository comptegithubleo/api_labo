#!/usr/bin/env bash

# get the smallest vlan id that's not yet registered in interfaces
vlan_id=$(grep -oE 'vlan[0-9]{1,4}' /etc/network/interfaces \
  | grep -oE '[0-9]{1,4}' \
  | sort -un \
  | awk 'BEGIN{min=10;found=0} {if($1==min) min++; else if($1>min){print min; found=1; exit}} END{if(!found) print min}')

echo $vlan_id

wg_ip=$(printf "10.%d.%d.0" $((vlan_id / 255)) $((vlan_id % 255)))

echo $wg_ip