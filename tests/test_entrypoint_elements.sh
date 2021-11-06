#!/bin/bash -u

# while :; do sleep 10; done
export WORKDIR_ROOT=workspace
export WORK_DIR=tests
export WORKDIR_PATH=/${WORKDIR_ROOT}/${WORK_DIR}

cd /${WORKDIR_ROOT}
if [ ! -d ${WORK_DIR} ]; then
  mkdir ${WORK_DIR}
fi

cd ${WORKDIR_PATH}
rm -rf elementsd_datadir

mkdir elementsd_datadir
chmod 777 elementsd_datadir
# cp /root/.elements/elements.conf elementsd_datadir/
cp ./elements.conf elementsd_datadir/

# boot daemon
bitcoin-cli -rpcconnect=testing-bitcoin -rpcport=18443 -rpcuser=bitcoinrpc -rpcpassword=password ping > /dev/null 2>&1
while [ $? -ne 0 ]
do
  bitcoin-cli -rpcconnect=testing-bitcoin -rpcport=18443 -rpcuser=bitcoinrpc -rpcpassword=password ping > /dev/null 2>&1
done
echo "start bitcoin node"
bitcoin-cli -rpcconnect=testing-bitcoin -rpcport=18443 -rpcuser=bitcoinrpc -rpcpassword=password createwallet wallet

elementsd -chain=liquidregtest -datadir=${WORKDIR_PATH}/elementsd_datadir
