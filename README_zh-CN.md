<h1 align="center">
  <br>xlsx2mysql<br>
</h1>

<h4 align="center">一个能快速帮助你把Excel数据转换成MySQL的SQL语句的工具</h4>

#### [中文文档](https://github.com/ToWeLong/xlsx2mysql/blob/master/README_zh-CN.md)

## 起源

为了实现把Excel数据转换成MySQL的SQL语句,我写了个工具去实现它。但是我为什么要写它呢？我找了一圈的工具但是仍然没有找到符合我需求的工具。所以这个工具就诞生了。

## 安装

### 从源码安装

xlsx2mysql 要求 Go >= 1.16. 你可以从源码编译:

```sh
$ go get -u -v github.com/ToWeLong/xlsx2mysql
```

### 安装二进制文件

可用的预下载文件: [release](https://github.com/ToWeLong/xlsx2mysql/releases)

## 使用

### 生成SQL文件

```
$ xlsx2mysql generate example.xlsx
```


## 另外

### 帮助

```
An tool of helping your fastly generate SQL from Excel.

Usage:
  xlsx2mysql [command]

Available Commands:
  completion  generate the autocompletion script for the specified shell
  generate    generate SQL from Excel
  help        Help about any command

Flags:
  -h, --help   help for xlsx2mysql

Use "xlsx2mysql [command] --help" for more information about a command.
```

## 开源协议

MIT