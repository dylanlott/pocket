## client Node EditStake

EditStake <fromAddr> <amount> <relayChainIDs> <serviceURI>

### Synopsis

Stakes a new <amount> for the Node actor with address <fromAddr> for the specified <relayChainIDs> and <serviceURI>.

```
client Node EditStake <fromAddr> <amount> <relayChainIDs> <serviceURI> [flags]
```

### Options

```
  -h, --help         help for EditStake
      --pwd string   passphrase used by the cmd, non empty usage bypass interactive prompt
```

### Options inherited from parent commands

```
      --path_to_private_key_file string   Path to private key to use when signing (default "./pk.json")
      --remote_cli_url string             takes a remote endpoint in the form of <protocol>://<host> (uses RPC Port) (default "http://localhost:50832")
```

### SEE ALSO

* [client Node](client_Node.md)	 - Node actor specific commands

###### Auto generated by spf13/cobra on 9-Nov-2022