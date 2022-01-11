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

set -e

python3 --version

pip3 install python-bitcoinrpc

# boot daemon
elements-cli -rpcconnect=testing-elements -rpcport=18447 -rpcuser=elementsrpc -rpcpassword=password ping > /dev/null 2>&1
while [ $? -ne 0 ]
do
  elements-cli -rpcconnect=testing-elements -rpcport=18447 -rpcuser=elementsrpc -rpcpassword=password ping > /dev/null 2>&1
done

python3 test.py -v
if [ $? -gt 0 ]; then
  cd ..
  exit 1
fi
