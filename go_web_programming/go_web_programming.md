# go web programming

## 基本概念

* 垂直扩展：指单机增加cpu或内存实现的扩展，因为go天然支持并发，所以增加cpu可以提高性能
* 水平扩展：指增加机器实现的扩展，因为go编译好的二进制文件不需要再提供三方依赖，所以实现水平扩展部署也比较简单
* HTTP协议交互是纯文本格式的，这种设计是的定位问题变得比较简单，因为不用做数据解析
* 一开始web application只用来传输**HTML文件**，后来人们发现如果能传输动态生产的内容就好了，最早传输动态生成内容的技术叫**CGI**(*Common Gateway Interface*)
    > client --> web server --环境变量-> cgi script
    > 
    > 然后cgi把内容通过stdout传给webserver
* HTTP请求方法的安全性和幂等性
    > GET,HEAD,OPTIONS,TRACE is safe
    > 
    > POST,PUT,DELETE not safe
    >
    > PUT,DELETE is idempotent, POST isn't idempotent
* HTML中的表单元素支持GET和POST两种HTTP方法