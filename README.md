# golab

A GitLab CLI tool bringing base on GitLab, for your XP Dev and Scrum Dev.

## Usage

```shell
$ git clone https://github.com/esperyong/golab.git 
$ cd golab
$ go install
$ golab -h
```

配置两个环境变量：
```shell
GITLAB_TOKEN=xxx
GITLAB_BASE_URL=https://xxx.yourgitlabxxx.com/api/v4
```
也可以在执行golab命令的地方创建一个.env文件,
将上面的环境变量写在文件中即可.
golab命令会自动加载文件中的设置到系统的环境变量中.

## 例子

1. 打印某一个milestone的统计信息
需要提供group milestone的名称和所在的组ID
```shell
$ golab milestone stat -m Sprint-7 -g 5
```


