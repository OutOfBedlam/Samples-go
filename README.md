
## 준비사항


VSCode의 필요한 플러그인을 설치한다: Go, Git Graph

Linux에서는 zip 명령으로 배포파일을 생성하므로 zip 명령이 존재해야한다.

```
$ sudo apt install zip
```


아래 파일들의 내용을 확인해 볼것.

- Makefile

    scripts/build.sh, package.sh을 구동하여 컴파일하거나 배포를 위한 패키지를 .zip으로 생성한다.

- scripts 디렉터리
    - build.sh내의 MODNAME 변수에 주의
    - build.sh내에서 가장 최근의 git tag를 가져오게 되어 있으므로 v0.0.0을 initial commit에 tagging하였음에 주의

- .vscode/settings.json

    vscode에 golang에 관련된 설정 (옵션)

go.mod 파일이 이미 존재하는데 최초에 프로젝트를 생성할 때는
아래와 같이 생성한다.

```sh
$ go mod init github.com/OutOfBedlam/Samples-go
go: creating new go.mod: module github.com/OutOfBedlam/Samples-go
go: to add module requirements and sums:
        go mod tidy
```

## make 및 스크립트 사용법

make를 실행하여 build하면 ./tmp 아래에 바이너리가 생성된다.

```sh
$ make

$ ./tmp/hello
Hello World?

$ ./tmp/count 
Count sum: 55
```

make package를 하면 각 배포 팩키지가 생성된다.


```sh
$ make package

$ ls -l packages
total 2008
-rw-r--r-- 1 eirny eirny 1027386 Oct  7 11:04 count-v0.0.0-linux-amd64.zip
-rw-r--r-- 1 eirny eirny 1027284 Oct  7 11:04 hello-v0.0.0-linux-amd64.zip
```
