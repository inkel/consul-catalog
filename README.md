# consul-catalog
Queries [Consul](https://consul.io/) catalog and prints out registered nodes and services.

## Usage
Install the latest version:

```
go install -u github.com/inkel/consul-catalog
```

By default it will print the nodes and services registered. You can control whether to include nodes or services with the following flags:

```
$ ./consul-catalog -h
Usage of ./consul-catalog:
  -nodes
        Include node names in output (default true)
  -services
        Include services in output (default true)
```

This is useful for TAB completion on the CLI. For instance, in Bash, to add TAB completion to your `ssh` command, you could do something like this:

```bash
_consul_catalog() {
    COMPRELY=( $( compgen -W "$(consul-catalog -services=false)" -- "$(_get_cword)" ) )
}
complete -F _consul_catalog ssh
```

Note we are using the `-services=false` flag to exclude services from being printed.

## License
See [LICENSE](LICENSE).
