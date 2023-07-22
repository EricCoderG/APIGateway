# api-gateway 测试文档

## 测试概述

本文档是对api-gateway项目相关测试的记录和报告，包括单元测试，集成测试和性能测试。本次测试使用了go语言自带的测试框架，go test测试工具，以及Apache Bench测试工具。

所有测试环境均为windows系统下的wsl2虚拟机，cpu为AMD Ryzen 7 5800H 8核16线程。

所有的微服务都进行了预热，以避免初次测试时的性能误差。



## 单元测试

### Test_convertCase_func_normal

测试用例说明：

- 测试传入字符串 "abcd" 的大小写转换是否正常。
- 预期输出字符串为 "ABCD"。

测试代码：

```go
func Test_convertCase_func_normal(t *testing.T) {
    testString := "abcd"
    expectedResp := "ABCD"
    csi := new(convHandler.ConversionServiceImpl)
    req := &conversion.ConversionRequest{
        InputString: testString,
    }
    resp, err := csi.ConvertCase(context.Background(), req)
    if err != nil {
        t.Error(err)
    }
    if resp.ConvertedString != expectedResp {
        t.Errorf("expected resp is %s but resp in fact is %s", expectedResp, resp.ConvertedString)
    }
}
```

### Test_convertCase_func_Mixed

测试用例说明：

- 测试传入字符串 "AbCd" 的大小写转换是否正常。
- 预期输出字符串为 "ABCD"。

测试代码：

```go
func Test_convertCase_func_Mixed(t *testing.T) {
    testString := "AbCd"
    expectedResp := "ABCD"
    csi := new(convHandler.ConversionServiceImpl)
    req := &conversion.ConversionRequest{
        InputString: testString,
    }
    resp, err := csi.ConvertCase(context.Background(), req)
    if err != nil {
        t.Error(err)
    }
    if resp.ConvertedString != expectedResp {
        t.Errorf("expected resp is %s but resp in fact is %s", expectedResp, resp.ConvertedString)
    }
}
```

### Test_convertCase_func_zero

测试用例说明：

- 测试传入空字符串的大小写转换是否正常。
- 预期输出字符串为空字符串。

测试代码：

```go
func Test_convertCase_func_zero(t *testing.T) {
    testString := ""
    expectedResp := ""
    csi := new(convHandler.ConversionServiceImpl)
    req := &conversion.ConversionRequest{
        InputString: testString,
    }
    resp, err := csi.ConvertCase(context.Background(), req)
    if err != nil {
        t.Error(err)
    }
    if resp.ConvertedString != expectedResp {
        t.Errorf("expected resp is %s but resp in fact is %s", expectedResp, resp.ConvertedString)
    }
}
```

### Test_calculateLength_func_normal

测试用例说明：

- 测试传入字符串 "abcd" 的长度计算是否正常。
- 预期输出字符串长度为 4。

测试代码：

```go
func Test_calculateLength_func_normal(t *testing.T) {
    testString := "abcd"
    expectedResp := 4
    lsi := new(lenHandler.LengthServiceImpl)
    req := &length.LengthRequest{
        InputString: testString,
    }
    resp, err := lsi.CalculateLength(context.Background(), req)
    if err != nil {
        t.Error(err)
    }
    if resp.Length != int32(expectedResp) {
        t.Errorf("expected resp is %d but resp in fact is %d", expectedResp, resp.Length)
    }
}
```

### Test_calculateLength_func_zero

测试用例说明：

- 测试传入空字符串的长度计算是否正常。
- 预期输出字符串长度为 0。

测试代码：

```go
func Test_calculateLength_func_zero(t *testing.T) {
    testString := ""
    expectedResp := 0
    lsi := new(lenHandler.LengthServiceImpl)
    req := &length.LengthRequest{
        InputString: testString,
    }
    resp, err := lsi.CalculateLength(context.Background(), req)
    if err != nil {
        t.Error(err)
    }
    if resp.Length != int32(expectedResp) {
        t.Errorf("expected resp is %d but resp in fact is %d", expectedResp, resp.Length)
    }
}
```

### Test_reverseString_func_normal

测试用例说明：

- 测试传入字符串 "abcd" 的翻转是否正常。
- 预期输出字符串为 "dcba"。

测试代码：

```go
func Test_reverseString_func_normal(t *testing.T) {
    testString := "abcd"
    expectedResp := "dcba"
    csi := new(handler.ReverseServiceImpl)
    req := &reverse.ReverseRequest{
        InputString: testString,
    }
    resp, err := csi.ReverseString(context.Background(), req)
    if err != nil {
        t.Error(err)
    }
    if resp.ReversedString != expectedResp {
        t.Errorf("expected resp is %s butresp in fact is %s", expectedResp, resp.ReversedString)
    }
}
```

### Test_reverseString_func_zero

测试用例说明：

- 测试传入空字符串的翻转是否正常。
- 预期输出字符串为空字符串。

测试代码：

```go
func Test_reverseString_func_zero(t *testing.T) {
    testString := ""
    expectedResp := ""
    csi := new(handler.ReverseServiceImpl)
    req := &reverse.ReverseRequest{
        InputString: testString,
    }
    resp, err := csi.ReverseString(context.Background(), req)
    if err != nil {
        t.Error(err)
    }
    if resp.ReversedString != expectedResp {
        t.Errorf("expected resp is %s but resp in fact is %s", expectedResp, resp.ReversedString)
    }
}
```

### Test_reverseString_func_Mixed

测试用例说明：

- 测试传入字符串 "abcd123ABC" 的翻转是否正常。
- 预期输出字符串为 "CBA321dcba"。

测试代码：

```go
func Test_reverseString_func_Mixed(t *testing.T) {
    testString := "abcd123ABC"
    expectedResp := "CBA321dcba"
    csi := new(handler.ReverseServiceImpl)
    req := &reverse.ReverseRequest{
        InputString: testString,
    }
    resp, err := csi.ReverseString(context.Background(), req)
    if err != nil {
        t.Error(err)
    }
    if resp.ReversedString != expectedResp {
        t.Errorf("expected resp is %s but resp in fact is %s", expectedResp, resp.ReversedString)
    }
}
```

### Test_FindSubstring_func_normal

测试用例说明：

- 测试在主字符串 "abcd" 中查找子字符串 "bc" 是否正常。
- 预期输出字符串为 "[1]"，表示子字符串在主字符串中的位置为1。

测试代码：

```go
func Test_FindSubstring_func_normal(t *testing.T) {
    mainStr := "abcd"
    subStr := "bc"
    expectedResp := "[1]"
    ssi := new(handler.SubstringServiceImpl)
    req := &substring.SubstringRequest{
        MainString: mainStr,
        SubString: subStr,
    }
    resp, err := ssi.FindSubstring(context.Background(), req)
    if err != nil {
        t.Error(err)
    }
    respstr := fmt.Sprintf("%v", resp.Positions)
    if respstr != expectedResp {
        t.Errorf("expected resp is %s but resp in fact is %s", expectedResp, respstr)
    }
}
```

### Test_FindSubstring_func_zero

测试用例说明：

- 测试在主字符串 "abcd" 中查找子字符串 "ef" 是否正常。
- 预期输出字符串为空字符串。

测试代码：

```go
func Test_FindSubstring_func_zero(t *testing.T) {
    mainStr := "abcd"
    subStr := "ef"
    expectedResp := "[]"
    ssi := new(handler.SubstringServiceImpl)
    req := &substring.SubstringRequest{
        MainString: mainStr,
        SubString: subStr,
    }
    resp, err := ssi.FindSubstring(context.Background(), req)
    if err != nil {
        t.Error(err)
    }
    respstr := fmt.Sprintf("%v", resp.Positions)
    if respstr != expectedResp {
        t.Errorf("expected resp is %s but resp in fact is %s", expectedResp, respstr)
    }
}
```

### Test_FindSubstring_func_MultiAnswer

测试用例说明：

- 测试在主字符串 "123456123" 中查找子字符串 "123" 是否正常。
- 预期输出字符串为 "[0 6]"，表示子字符串在主字符串中的位置分别为0和6。

测试代码：

```go
func Test_FindSubstring_func_MultiAnswer(t *testing.T) {
    mainStr := "123456123"
    subStr := "123"
    expectedResp := "[0 6]"
    ssi := new(handler.SubstringServiceImpl)
    req := &substring.SubstringRequest{
        MainString: mainStr,
        SubString: subStr,
    }
    resp, err := ssi.FindSubstring(context.Background(), req)
    if err != nil {
        t.Error(err)
    }
    respstr := fmt.Sprintf("%v", resp.Positions)
    if respstr != expectedResp {
        t.Errorf("expected resp is %s but resp in fact is %s", expectedResp, respstr)
    }
}
```

### 测试结果

每个测试用例都有一个预期的响应字符串。如果服务的响应与预期的响应匹配，那么测试就会通过。否则，测试就会失败，并记录错误信息。经检验，所有测试均通过。

### 注意事项

单元测试用例不依靠于任何已启动的服务，用户可以clone本项目之后直接运行以验证测试的正确性



## 集成测试

本部分以大小写转换服务为例，描述了对字符串转换服务的集成测试。这个服务接受一个输入字符串，并将其转换为大写形式。我们的测试目标是确保服务正确地处理了各种输入情况，并返回了预期的输出。关于其余集成测试的代码可以在本项目的github代码仓库中找到，在此就不一一列出了。

下面列出了关于字符串转换服务的测试用例:

### Test_conversion_normal

```go
func Test_conversion_normal(t *testing.T) {
	url := "http://127.0.0.1:8088/convert"
	jsonStr := `{"inputString":"abcccc"}`
	expectedResp := `{"convertedString":"ABCCCC"}`
	resp, err := http.Post(url, "application/json", bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		t.Fatal(err)
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	err = resp.Body.Close()

	if err != nil {
		return
	}

	respString := string(respBytes)

	if respString != expectedResp {
		t.Error("expected resp is " + expectedResp + " but real resp is " + respString + "!")
	} else {
		t.Log("Normal Test Pass!")
	}
}
```

测试目标：验证服务是否能正确地将小写字母转换为大写。

测试步骤：

1. 发送一个包含小写字母的 JSON 请求到服务端。
2. 读取并关闭响应体。
3. 检查响应体是否与预期的大写字符串匹配。

###  Test_conversion_ZERO

```go
func Test_conversion_ZERO(t *testing.T) {
	url := "http://127.0.0.1:8088/convert"
	jsonStr := `{"inputString":""}`
	expectedResp := `{"convertedString":""}`
	resp, err := http.Post(url, "application/json", bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		t.Fatal(err)
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	err = resp.Body.Close()

	if err != nil {
		return
	}

	respString := string(respBytes)

	if respString != expectedResp {
		t.Error("expected resp is " + expectedResp + " but real resp is " + respString + "!")
	} else {
		t.Log("ZERO Test Pass!")
	}
}
```

测试目标：验证服务是否能正确处理空字符串。

测试步骤：

1. 发送一个包含空字符串的 JSON 请求到服务端。
2. 读取并关闭响应体。
3. 检查响应体是否与预期的空字符串匹配。

###  Test_conversion_Mixed

```go
func Test_conversion_Mixed(t *testing.T) {
	url := "http://127.0.0.1:8088/convert"
	jsonStr := `{"inputString":"aBcCcC"}`
	expectedResp := `{"convertedString":"ABCCCC"}`
	resp, err := http.Post(url, "application/json", bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		t.Fatal(err)
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	err = resp.Body.Close()

	if err != nil {
		return
	}

	respString := string(respBytes)

	if respString != expectedResp {
		t.Error("expected resp is " + expectedResp + " but real resp is " + respString + "!")
	} else {
		t.Log("Mixed Test Pass!")
	}
}
```

测试目标：验证服务是否能正确处理混合大、小写字母的字符串。

测试步骤：

1. 发送一个包含混合大、小写字母的 JSON 请求到服务端。
2. 读取并关闭响应体。
3. 检查响应体是否与预期的全部大写字符串匹配。

### 测试结果

每个测试用例都有一个预期的响应字符串。如果服务的响应与预期的响应匹配，那么测试就会通过。否则，测试就会失败，并记录错误信息。经检验，所有测试均通过。

### 注意事项

在进行测试之前，确保服务已经启动并且可以接受请求。每个测试用例都会在完成后关闭响应体，以释放资源。如果在读取响应体或关闭响应体时发生错误，测试将会失败。

## 性能测试

### 关于性能测试

接下来以conversion的测试代码为例解释一下我们的性能测试代码，用户可以在我们项目的源代码test目录下找到：

用于测试 conversion 接口的性能的代码包含两个函数：BenchmarkConversion 和 BenchmarkConversionParallel。

BenchmarkConversion 函数模拟了一个循环调用 conversion 接口的场景，它接受一个参数 b，类型为 *testing.B。在函数内部，它使用 http.Post 方法发送请求，并读取响应内容。该函数会被运行 b.N 次，b.N 是基准测试框架自动生成的一个参数，表示测试的迭代次数。

BenchmarkConversionParallel 函数与 BenchmarkConversion 函数类似，但是它使用了 b.RunParallel 方法来并行运行测试。该函数还调用了 runtime.GOMAXPROCS 方法，将并发执行的最大 CPU 数量设置为 8。在函数内部，它使用 http.Post 方法发送请求，并读取响应内容。该函数会被运行 b.N 次，而且每次会使用多个 goroutine 并行执行。

这两个基准测试函数都使用了 http.Post 方法发送请求，并读取响应内容。它们的测试场景都是发送相同的请求数据，并模拟多次循环调用的场景，用于测试 conversion 接口的性能。其中，BenchmarkConversionParallel 函数还使用了并发执行的方式来进行测试，以测试接口的并发性能。



本项目的性能测试使用了 ApacheBench 工具进行 HTTP 压力测试。ApacheBench 是一个非常流行的 HTTP 压力测试工具，可以用于测试 Web 服务器的性能和稳定性。它可以模拟大量并发请求，并计算每个请求的响应时间、吞吐量等指标，从而评估 Web 服务器的性能和质量。

ab命令的参数说明：

- `-n 1000` 表示执行 1000 次请求。
- `-c 10` 表示并发请求数量为 10。
- `-T 'application/x-www-form-urlencoded'` 表示提交的数据类型为 `application/x-www-form-urlencoded`。
- `-p test/data/conversion.data` 表示使用 `test/data/conversion.data` 文件中的数据作为 POST 请求的 body。

### length-service

测试数据：inputString=12345678909876543210abc

测试命令： `go test -bench=. test/length_test.go`

测试结果：

> goos: linux
> goarch: amd64
> cpu: AMD Ryzen 7 5800H with Radeon Graphics         
> BenchmarkLength-16                  4341            250803 ns/op
> BenchmarkLengthParallel-16         28168             41971 ns/op
> testing: BenchmarkLengthParallel-16 left GOMAXPROCS set to 8
> PASS
> ok      command-line-arguments  2.742s

测试命令：`ab -n 1000 -c 10 -T 'application/x-www-form-urlencoded' -p test/data/length.data "http://127.0.0.1:8088/length"`

测试结果：

> Benchmarking 127.0.0.1 (be patient)
> Completed 100 requests
> Completed 200 requests
> Completed 300 requests
> Completed 400 requests
> Completed 500 requests
> Completed 600 requests
> Completed 700 requests
> Completed 800 requests
> Completed 900 requests
> Completed 1000 requests
> Finished 1000 requests
>
>
> Server Software:        hertz
> Server Hostname:        127.0.0.1
> Server Port:            8088
>
> Document Path:          /length
> Document Length:        13 bytes
>
> Concurrency Level:      10
> Time taken for tests:   0.053 seconds
> Complete requests:      1000
> Failed requests:        0
> Total transferred:      170000 bytes
> Total body sent:        194000
> HTML transferred:       13000 bytes
> Requests per second:    18866.86 [#/sec] (mean)
> Time per request:       0.530 [ms] (mean)
> Time per request:       0.053 [ms] (mean, across all concurrent requests)
> Transfer rate:          3132.19 [Kbytes/sec] received
>                         3574.38 kb/s sent
>                         6706.58 kb/s total
>
> Connection Times (ms)
>               min  mean[+/-sd] median   max
> Connect:        0    0   0.1      0       0
> Processing:     0    0   0.1      0       2
> Waiting:        0    0   0.1      0       1
> Total:          0    1   0.1      0       2
> ERROR: The median and mean for the total time are more than twice the standard
>        deviation apart. These results are NOT reliable.
>
> Percentage of the requests served within a certain time (ms)
>   50%      0
>   66%      1
>   75%      1
>   80%      1
>   90%      1
>   95%      1
>   98%      1
>   99%      1
>  100%      2 (longest request)

### conversion-service

测试数据：inputString=abcccc

测试命令:  `go test -bench=. test/conversion_test.go`

测试结果：

> goos: linux
> goarch: amd64
> cpu: AMD Ryzen 7 5800H with Radeon Graphics         
> BenchmarkConversion-16                      1039           1299434 ns/op
> BenchmarkConversionParallel-16               100          19997059 ns/op
> testing: BenchmarkConversionParallel-16 left GOMAXPROCS set to 8
> PASS
> ok      command-line-arguments  3.476s

测试命令：`ab -n 1000 -c 10 -T 'application/x-www-form-urlencoded' -p test/data/conversion.data "http://127.0.0.1:8088/convert"`

> Benchmarking 127.0.0.1 (be patient)
> Completed 100 requests
> Completed 200 requests
> Completed 300 requests
> Completed 400 requests
> Completed 500 requests
> Completed 600 requests
> Completed 700 requests
> Completed 800 requests
> Completed 900 requests
> Completed 1000 requests
> Finished 1000 requests
>
>
> Server Software:        hertz
> Server Hostname:        127.0.0.1
> Server Port:            8088
>
> Document Path:          /convert
> Document Length:        32 bytes
>
> Concurrency Level:      10
> Time taken for tests:   0.279 seconds
> Complete requests:      1000
> Failed requests:        0
> Total transferred:      189000 bytes
> Total body sent:        179000
> HTML transferred:       32000 bytes
> Requests per second:    3584.92 [#/sec] (mean)
> Time per request:       2.789 [ms] (mean)
> Time per request:       0.279 [ms] (mean, across all concurrent requests)
> Transfer rate:          661.67 [Kbytes/sec] received
>                         626.66 kb/s sent
>                         1288.33 kb/s total
>
> Connection Times (ms)
>               min  mean[+/-sd] median   max
> Connect:        0    0   0.1      0       1
> Processing:     1    3   1.0      2      11
> Waiting:        1    3   1.0      2      11
> Total:          1    3   1.0      2      11
>
> Percentage of the requests served within a certain time (ms)
>   50%      2
>   66%      3
>   75%      3
>   80%      3
>   90%      4
>   95%      5
>   98%      6
>   99%      7
>  100%     11 (longest request)

### substring-service

测试数据： mainString=123456123&subString=123

测试命令 ：`go test -bench=. test/substring_test.go`

测试结果：

> goos: linux
> goarch: amd64
> cpu: AMD Ryzen 7 5800H with Radeon Graphics         
> BenchmarkSubstring-16                        732           1530601 ns/op
> BenchmarkSubstringParallel-16               2893            626882 ns/op
> testing: BenchmarkSubstringParallel-16 left GOMAXPROCS set to 8
> PASS
> ok      command-line-arguments  5.009s

测试命令： `ab -n 1000 -c 10 -T 'application/x-www-form-urlencoded' -p test/data/substring.data "http://127.0.0.1:8088/substring"`

> Benchmarking 127.0.0.1 (be patient)
>
> Completed 100 requests
> Completed 200 requests
> Completed 300 requests
> Completed 400 requests
> Completed 500 requests
> Completed 600 requests
> Completed 700 requests
> Completed 800 requests
> Completed 900 requests
> Completed 1000 requests
> Finished 1000 requests
>
>
> Server Software:        hertz
> Server Hostname:        127.0.0.1
> Server Port:            8088
>
> Document Path:          /substring
> Document Length:        16 bytes
>
> Concurrency Level:      10
> Time taken for tests:   0.733 seconds
> Complete requests:      1000
> Failed requests:        0
> Total transferred:      173000 bytes
> Total body sent:        196000
> HTML transferred:       16000 bytes
> Requests per second:    1364.17 [#/sec] (mean)
> Time per request:       7.330 [ms] (mean)
> Time per request:       0.733 [ms] (mean, across all concurrent requests)
> Transfer rate:          230.47 [Kbytes/sec] received
>                         261.11 kb/s sent
>                         491.58 kb/s total
>
> Connection Times (ms)
>               min  mean[+/-sd] median   max
> Connect:        0    0   0.1      0       1
> Processing:     2    7   6.3      6      66
> Waiting:        2    7   6.3      6      66
> Total:          2    7   6.3      7      66
>
> Percentage of the requests served within a certain time (ms)
>   50%      7
>   66%      8
>   75%      8
>   80%      9
>   90%     11
>   95%     13
>   98%     25
>   99%     55
>  100%     66 (longest request)

### reverse-service

测试数据：inputString=123456

测试命令：`go test -bench=. test/reverse_test.go`

测试结果：

> goos: linux
> goarch: amd64
> cpu: AMD Ryzen 7 5800H with Radeon Graphics         
> BenchmarkReverse-16                  519           2010966 ns/op
> BenchmarkReverseParallel-16         3042            449251 ns/op
> testing: BenchmarkReverseParallel-16 left GOMAXPROCS set to 8
> PASS
> ok      command-line-arguments  2.705s

测试命令： `ab -n 1000 -c 10 -T 'application/x-www-form-urlencoded' -p test/data/reverse.data "http://127.0.0.1:8088/reverse"`

测试结果：

> Benchmarking 127.0.0.1 (be patient)
> Completed 100 requests
> Completed 200 requests
> Completed 300 requests
> Completed 400 requests
> Completed 500 requests
> Completed 600 requests
> Completed 700 requests
> Completed 800 requests
> Completed 900 requests
> Completed 1000 requests
> Finished 1000 requests
>
>
> Server Software:        hertz
> Server Hostname:        127.0.0.1
> Server Port:            8088
>
> Document Path:          /reverse
> Document Length:        27 bytes
>
> Concurrency Level:      10
> Time taken for tests:   0.898 seconds
> Complete requests:      1000
> Failed requests:        0
> Total transferred:      184000 bytes
> Total body sent:        177000
> HTML transferred:       27000 bytes
> Requests per second:    1113.27 [#/sec] (mean)
> Time per request:       8.983 [ms] (mean)
> Time per request:       0.898 [ms] (mean, across all concurrent requests)
> Transfer rate:          200.04 [Kbytes/sec] received
>                         192.43 kb/s sent
>                         392.47 kb/s total
>
> Connection Times (ms)
>               min  mean[+/-sd] median   max
> Connect:        0    0   0.2      0       3
> Processing:     2    9  12.7      5     103
> Waiting:        2    9  12.7      5     103
> Total:          2    9  12.7      5     103
>
> Percentage of the requests served within a certain time (ms)
>   50%      5
>   66%      7
>   75%      9
>   80%     11
>   90%     16
>   95%     25
>   98%     64
>   99%    100
>  100%    103 (longest request)

