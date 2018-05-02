ffcalc
======

ffcalc 包是一个学习性质的有限域计算器的 go 实现

## 安装与更新

```bash
$ go get -u github.com/master-g/big/src/foundation_math/ffcalc
```

## 使用

```bash
$ ffcalc prime degree [exp...]
```

其中`prime` 为有限域的素数, `degree` 为素数的正整数次幂, `exp` 为计算表达式 

例子

```bash
$ ffcalc 19 1 "11+6"
17
$ ffcalc 31 1 "4^-4*11"
13
```

## 参考

[谈谈有限域那些事儿](https://blog.csdn.net/qmickecs/article/details/77281602)  
[WikiPedia](https://en.wikipedia.org/wiki/Finite_field#GF(p2)_for_an_odd_prime_p)  
[Blockchain 101](https://eng.paxos.com/blockchain-101-foundational-math)  

## 许可证

ffcalc 遵循 MIT 许可证