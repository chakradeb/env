# Env

![Build-Test](https://github.com/chakradeb/env/workflows/Build-Test/badge.svg)
[![codecov](https://codecov.io/gh/chakradeb/env/branch/master/graph/badge.svg)](https://codecov.io/gh/chakradeb/env)

This package helps you to parse environment variables using struct

```go
import "github.com/chakradeb/env"
```

## Examples

Lets start with declaring a couple of environment variables

```shell script
export PORT=5000
export HOST=localhost
export DEBUG=false
```

Now we will write a `struct` to consume these environment variables

```go
type Config struct {
    Port int `env:"PORT"`
    Host string `env:"HOST"`
    Debug bool `env:"DEBUG"`
}
```

Here the tag `env` helps the parser to find the variables from environment variables.
Parser will search for the same name in environment variable as provided by `env` tag.

Now, The main function will look something like this

```go
func main() {
    conf := &Config{}

    errs := env.Parse(conf)

    fmt.Println("Port: ", conf.Port)
    fmt.Println("Host: ", conf.Host)
    fmt.Println("Debug: ", conf.Debug)
}
```

Result:

```shell script
Port: 5000
Host: localhost
Debug: false
```

## Supported Struct Tags

For now `Env` supports only two kind of tags, `env` and `default`.

Tag `env` is mandatory to match with the environment variable.
Tag `default` is not mandatory but it is useful to set a default value to a struct field.

For example, consider these environment variables:

```shell script
export HOST=localhost
```

Now we will declare some default value to our struct field using the `default` tag

```go
type Config struct {
    Port int `env:"PORT" default:"8000"`
    Host string `env:"HOST" default:"github.chakradeb.com"`
    Debug bool `env:"DEBUG" default:"true"`
}
```

Now we will see the output of our `main` function

```go
func main() {
    conf := &Config{}

    errs := env.Parse(conf)

    fmt.Println("Port: ", conf.Port)
    fmt.Println("Host: ", conf.Host)
    fmt.Println("Debug: ", conf.Debug)
}
```

Result:

```shell script
Port: 8000
Host: localhost
Debug: true
```

As you can see `Port` and `Debug` has the default value what we've provided using the `default` tag.

But `Host` didn't get the default value provided in the `default` tag,
as it is declared in the environment variable already.

So `default` value will only be effective if there is no environment variable present with the name provided with the `env` tag

## Supported Struct Fields

For now `Env` supports only these struct field types:

- `string`
- `int`
- `int8`
- `int16`
- `int32`
- `int64`
- `float32`
- `float64`
- `bool`

More support will be added soon
