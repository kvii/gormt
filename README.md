# gormt

gorm bug report.

## steps

1. clone https://github.com/kvii/gormt.git
2. cd gormt
3. git clone https://github.com/kvii/gorm.git gorm
4. cd gormt
5. go run .
6. Remove go.work file and try again.

You can see some log like this.

``` shell
2022/11/05 11:25:41 /Users/kvii/workspace/qs/gormt/gormt/main.go:14
[0.006ms] [rows:0] SELECT 1

2022/11/05 11:28:10 /Users/kvii/go/pkg/mod/gorm.io/gorm@v1.24.1/callbacks.go:134
[0.003ms] [rows:0] SELECT 1
```