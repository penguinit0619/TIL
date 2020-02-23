# Go Module

Go module은 go package manager이다.

다른 언어에서는 이런 역할을 Node에서는 NPM이 Python에서는 PyPI가 했는데 동일한 레벨로 보면 이해가 쉬울듯 하다.

Golang 1.11, 1.12버전에 베타로 들어갔다가 1.13 버전이후에는 공식적으로 포함이 되었다.

사실 내가 처음 Go를 접했던 시기가 1.13버전 이후라서 처음에 프로젝트를 빌드하는 과정에서부터 Go Module을 사용했기에 그렇기 큰 혼란은없었는데 이전에 사용했다면 조금 당황스러웠을 것 같긴하다.



>  Go Module이 나오기 이전에는 GOPATH에 설정된 bin, src, pkg에서 패키지들을 관리하였다.



단순히 생각해 봤을 때 위와 같이 하면 어떤 문제점이 있을지에 대해서 생각해보면 일단 버전관리가 힘들고, 또 개인이 툴을 만들어서 운영 할 경우 관리 포인트가 늘어나는 점도 힘들어보인다.

Go Blog에서 보니 Gopher들이 하나로 통합된 Package Manager에 대한 요구가 많았던 걸로 보인다. [링크](https://blog.golang.org/versioning-proposal)



#### 사용법

사용법은 간단하다. 만약에 기존 Go를 사용 해본 적이 있다면 go mod라는 명령어에 대해서만 이해를 하고 있으면된다.

프로젝트의 Root Directory에 아래와 같이 명령어를 수행하면 go.mod라는 파일이 생성된다.

```bash
go mod init github.com/penguinit0619/go-sample
$ ls
go.mod  hello.go  hello_test.go
$ cat go.mod 
module github.com/penguinit0619/go-sample

go 1.13
```

뒤에 붙은 `github.com/penguinit0619/go-sample`은 임의로 생성한 네임스페이스이다. go.mod의 하위 디렉토리들은 모두 네임스페이스를 기반으로 인식된다. 예를들어서 하위에 work라는 디렉토리를 만들었다면 import되는 path는 `github.com/penguinit0619/go-sample/work`라고 볼 수 있겠다.



위 처럼 설정하고 아래 코드를 돌려보자



> hello.go

```go
package hello

import "rsc.io/quote"

func Hello() string {
    return quote.Hello()
}
```



> hello_test.go

```go
package hello

import "testing"

func TestHello(t *testing.T) {
        want := "Hello, world."
        if got := Hello(); got != want {
                t.Errorf("Hello() = %q, want %q", got, want)
        }
}
```

> 결과

```bash
$ go test
go: finding rsc.io/quote v1.5.2
go: downloading rsc.io/quote v1.5.2
go: extracting rsc.io/quote v1.5.2
go: downloading rsc.io/sampler v1.3.0
go: extracting rsc.io/sampler v1.3.0
go: downloading golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
go: extracting golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
go: finding rsc.io/sampler v1.3.0
go: finding golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
--- FAIL: TestHello (0.00s)
    hello_test.go:8: Hello() = "안녕, 세상.", want "Hello, world."
FAIL
exit status 1
FAIL	github.com/penguinit0619/go-sample	0.004s
```



Go module에 대해서 조금 더 심화해서 파보자

우선 결과를 보면 test 수행과 함께 rsc.io/quote가 import 되는 것과 부차적으로 수행되는 dependency들을 볼 수 있다. 그리고 수행이 되면서 go.sum이라는 파일과 go.mod에 변화가 생긴 것 을 확인 할 수 있다.



```bash
$ ls
go.mod  go.sum  hello.go  hello_test.go

$ cat go.mod 
module github.com/penguinit0619/go-sample

go 1.13

require rsc.io/quote v1.5.2

$ cat go.sum 
golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c h1:qgOY6WgZOaTkIIMiVjBQcw93ERBE4m30iBm00nkL0i8=
golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c/go.mod h1:NqM8EUOU14njkJ3fqMW+pc6Ldnwhi/IjpwHt7yyuwOQ=
rsc.io/quote v1.5.2 h1:w5fcysjrx7yqtD/aO+QwRjYZOKnaM9Uh2b40tElTs3Y=
rsc.io/quote v1.5.2/go.mod h1:LzX7hefJvL54yjefDEDHNONDjII0t9xZLPXsUe+TKr0=
rsc.io/sampler v1.3.0 h1:7uVkIFmeBqHfdjD+gZwtXXI+RODJ2Wc4O7MPEh/QiW4=
rsc.io/sampler v1.3.0/go.mod h1:T1hPZKmBbMNahiBKFy5HrXp6adAjACjK9JXDnKaTXpA=

$ go list -m all
github.com/penguinit0619/go-sample
golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
rsc.io/quote v1.5.2
rsc.io/sampler v1.3.0
```

go.mod 파일에서는 rsc.io/quote v1.5.2 만 불러온 것으로 보이지만 그 하위에 간접적으로 참조하는 패키지들이 있는데 이는 go.sum에 명시되어 있다. 이곳에서 참조된 패키지들의 자세한 버전과 해쉬 값을 볼 수 있다. 



#### Upgrade Dependencies

rsc.io/quote는 뭔가 문제가 있어보인다. 현재 기대하는 값은 분명 hello world인데 다른 값이 찍히는걸 보니 버전이 상승하면서 값이 바뀐것으로 보인다. 

```bash
$ go get golang.org/x/text
```

우선 golang.org/x/text를 최신버전으로 받아보자. go get 명령어를 통해서 패키지를 수동으로 가져 올 수 있는데 이때 따로 버전을 명시하지 않으면 최신버전의 패키지를 가지고 온다.

```bash
$ go list -m -versions golang.org/x/text
golang.org/x/text v0.1.0 v0.2.0 v0.3.0 v0.3.1 v0.3.2
```



위 처럼되면 go.mod와 go.sum은 어떻게 값이 변할까?

```bash
$ cat go.mod 
module github.com/penguinit0619/go-sample

go 1.13

require (
	golang.org/x/text v0.3.2 // indirect
	rsc.io/quote v1.5.2
)

$ cat go.sum 
golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c h1:qgOY6WgZOaTkIIMiVjBQcw93ERBE4m30iBm00nkL0i8=
golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c/go.mod h1:NqM8EUOU14njkJ3fqMW+pc6Ldnwhi/IjpwHt7yyuwOQ=
golang.org/x/text v0.3.2 h1:tW2bmiBqwgJj/UpqtC8EpXEZVYOwU0yG4iWbprSVAcs=
golang.org/x/text v0.3.2/go.mod h1:bEr9sfX3Q8Zfm5fL9x+3itogRgK3+ptLWKqgva+5dAk=
golang.org/x/tools v0.0.0-20180917221912-90fa682c2a6e/go.mod h1:n7NCudcB/nEzxVGmLbDWY5pfWTLqBcC2KZ6jyYvM4mQ=
rsc.io/quote v1.5.2 h1:w5fcysjrx7yqtD/aO+QwRjYZOKnaM9Uh2b40tElTs3Y=
rsc.io/quote v1.5.2/go.mod h1:LzX7hefJvL54yjefDEDHNONDjII0t9xZLPXsUe+TKr0=
rsc.io/sampler v1.3.0 h1:7uVkIFmeBqHfdjD+gZwtXXI+RODJ2Wc4O7MPEh/QiW4=
rsc.io/sampler v1.3.0/go.mod h1:T1hPZKmBbMNahiBKFy5HrXp6adAjACjK9JXDnKaTXpA=
```



우선  go.mod에서는 golang.org/x/text v0.3.2 // indirect 으로 명시가 되게 되고 go.sum에는 이전버전인 v.0.0.0과 최신버전인 v.0.3.2가 모두 존재하는 것으로 보인다. go.mod에서 v0.3.2로 명시가 되어있기 때문에 go build시에 v0.3.2를 가져올 것이다.



```bash
go get rsc.io/sampler@v1.3.1
```

특정 버전으로 package를 받고 싶다면 @{version}을 통해서 가능하다.