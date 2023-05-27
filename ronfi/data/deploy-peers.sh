#!/bin/bash

date "+%Y%m%d-%H%M%S"

peers=('abb' 'c12c' 'c2d' 'm5a' 'm5b' 'm5c' 'm5e' 'm0z' 'm1z' 'm2z' 'm3z' 'm4z' 'm5z' 'm6z' 'h1z' 'hza')
#peers=('m5f' 'c12h' 'c12i' 'c12j')

for i in "${!peers[@]}"; do
    peer=${peers[$i]}
    scp ronfi/data/obs_routers.json bsc@$peer:bin/.
    ssh bsc@$peer "cd ~/bin; ./reloadLoops.sh > /dev/null"
done
