#!/usr/bin/env bash

geth --bootnodes "enode://a9cc1b1888054e050b6b8c167744540354b0d8bf0361ff8a92364989a5912aefb7c0f1d1122f261c731b2338d85c0ffaa3a9b4d543447a583101e530989c7309@127.0.0.1:30301" \
     --identity "node01" \
     --rpc --rpcport "8000" --rpccorsdomain "*" \
     --ipcdisable \
     --datadir ./ \
     --ethash.dagdir ../temp/dag/ \
     --ethash.cachedir ../temp/cache/ \
     --port "8545" \
     --rpcapi "admin,db,eth,miner,net,personal,txpool,web3" \
     --networkid 714 \
     --nat "any"
