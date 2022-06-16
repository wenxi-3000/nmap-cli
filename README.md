# nmap-cli
### 说明
读取host:ip格式文件，进行处理，交给nmap探测，保存结果
### 使用方法
git clone https://github.com/ShadowFl0w/nmap-cli.git<br>
cd nmap-cli<br>
go build<br>
./nmap-cli -f ./hostip.txt -o nmap.txt

### 输入的文件
指定输入文件，host:ip类型，
```
172.16.42.151:8081
172.16.42.1:8082
172.16.2.151:8081
172.16.42.121:8082
172.1.42.151:8080
172.16.42.132:8080
```

### 输出nmap的结果
```
PORT     STATE  SERVICE    VERSION
8080/tcp closed http-proxy
8081/tcp open   http       SimpleHTTPServer 0.6 (Python 2.7.16)
8082/tcp open   http       SimpleHTTPServer 0.6 (Python 2.7.16)
```
