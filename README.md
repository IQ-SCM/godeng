# godeng

![](https://img.shields.io/github/license/chenjiayao/godeng)
![](https://github.com/chenjiayao/godeng/actions/workflows/gorelease.yml/badge.svg)


Godeng is used to generate test logs. For example, to generate a large amount of test data in ELK. you can define the rule and name of the fields, and godeng will automatically **generate** the data for you. 

**Godeng depends on gofakeit, thanks to [gofakeit](https://github.com/mingrammer/flog) 🥰**

## 👨‍💻 installation

You can install godeng using the following

### 📦 Using .tar.gz archive

Download gzip file from [Github Releases](https://github.com/chenjiayao/godeng/releases) according to your OS. Then, copy the unzipped executable to under system path.

### 🐳 Docker

```bash
docker run --rm chenjiayao/godeng
```

## 🧑‍💻 Usage

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
      --url string         http request url,only used when output is http/https and output is json
```

## Support Formats

- json
- sql


## Support output

- stdout
- file
- http

## Examples

```
Send a post request to http://some.website every 3 seconds, and never exit
godeng --config=godeng.json --url=http://some.website --sleep=3 --loop

Generate 100 sql statements to insert into the faker table
godeng --config=godeng.json --format=sql --coun=100 --tablemame=faker 

Output 100 json data to output.json file
godeng --config=godeng.json --format=json --output=file --file=output.json
```

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
- ua (user-agent)


## Issue

if you need more ouput/format/support type or find a bug, please raise a [issue](https://github.com/chenjiayao/godeng/issues).