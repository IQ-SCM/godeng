# godeng

Godeng is used to generate test logs. For example, to generate a large amount of test data in ELK. you can define the rule and name of the fields, and godeng will automatically generate the data for you. 

Godeng depends on gofakeit, thanks to [gofakeit](https://github.com/mingrammer/flog) 🥰

## 👨‍💻 installation

You can install godeng using the following

### Docker



## Usage

There are useful options.

```
Flags:
      --config string      config file (default "./dodeng.json")
      --count int          count (default 100)
      --file string        output file, only used when output is file (default "./godeng.out")
      --format string      output format (default "json")
  -h, --help               help for godeng
      --loop               loop output forever until killed. if loop is set, then count is ignored
      --output string      output (default "stdout")
      --sleep int          fix creation time interval for each log (second)
      --tablename string   tablename, only used when output is sql (default "godeng")
      --url string         http request url,only used when output is http/https
```

## Support Formats

- json
- sql

if you need another format, please raise an issue.


## Support output

- stdout
- file
- sql
- http request

if you need another output way, please raise an issue.


## Support type

- string
- int
- float
- ipv4
- mac
- ipv6
- bool
- enum
- url
- datetime
- timestamp
- email
- sequence
- ua