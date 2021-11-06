import unittest
import logging
import time
from bitcoinrpc.authproxy import AuthServiceProxy


class RpcWrapper:
    def __init__(self, host='127.0.0.1', port=8432,
                 rpc_user='', rpc_password=''):
        self.rpc_connection = AuthServiceProxy('http://{}:{}@{}:{}'.format(
            rpc_user, rpc_password, host, port))

    def command(self, command, *args):
        return self.rpc_connection.command(args)

    def get_rpc(self):
        return self.rpc_connection


class TestElements(unittest.TestCase):
    def setUp(self):
        logging.basicConfig()
        logging.getLogger("BitcoinRPC").setLevel(logging.DEBUG)

        self.btcConn = RpcWrapper(
            host='testing-bitcoin', port=18443, rpc_user='bitcoinrpc', rpc_password='password')
        self.elmConn = RpcWrapper(
            host='testing-elements', port=18447, rpc_user='elementsrpc', rpc_password='password')
        # init command
        btc_rpc = self.btcConn.get_rpc()
        try:
            btc_rpc.settxfee(0.00001)
        except Exception as err:
            print('Exception({})'.format(err))
            btc_rpc.createwallet('wallet')

    def test_bitcoin_elements(self):
        btc_rpc = self.btcConn.get_rpc()
        elm_rpc = self.elmConn.get_rpc()

        past_btc_chaininfo = btc_rpc.getblockchaininfo()
        past_elm_chaininfo = elm_rpc.getblockchaininfo()

        time.sleep(10)
        btc_chaininfo = btc_rpc.getblockchaininfo()
        elm_chaininfo = elm_rpc.getblockchaininfo()

        self.assertGreaterEqual(btc_chaininfo['blocks'], past_btc_chaininfo['blocks'] + 5)
        self.assertGreaterEqual(elm_chaininfo['blocks'], past_elm_chaininfo['blocks'] + 5)


if __name__ == "__main__":
    unittest.main()
