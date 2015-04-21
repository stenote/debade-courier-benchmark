# debade-courier-benchmark

## 简介

用来发送消息给 **debade-courier**, 目的是用来比对 [debade-courier-go](https://github.com/stenote/debade-courier-go) 和 [debade-courier](https://github.com/iamfat/debade-courier) 的性能

## Usage

```bash

  -c=10: 并发数
  -d="": 发送的数据(json字符串)
  -n=100: 每秒发送数
  -p="tcp://0.0.0.0:3333": 0MQ 连接路径

```

## Example

```bash

$ debade-courier-benchmark  \
	-c=10 \
	-n=100 \
	-p="tcp://192.168.0.53:3333" \
	-d='{"queue":"lims","data":["hello world"]}'

```

## License
Benchmark For Debade-Courier
The MIT License (MIT)

Copyright (c) 2015 Rui Ma

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
