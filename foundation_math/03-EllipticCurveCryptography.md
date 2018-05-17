03 椭圆曲线加密
================

## 域参数 Domain parameters

我们需要在有限域的一个循环子群上使用椭圆曲线算法, 因此我们需要以下参数:

* 确定有限域大小的质数 **`p`**
* 椭圆曲线的两个系数 **`a`**, **`b`**
* 子群的生产点 **`G`**
* 子群的阶 **`n`**
* 子群的余因子 **`h`** (h = 有限域的阶 N / 子群的阶 n)

故域参数是一个六元组 `sextuple(p, a, b, G, n, h)`  

## 随机曲线



[ref](https://eng.paxos.com/blockchain-101-elliptic-curve-cryptography)  
[计算](https://en.wikipedia.org/wiki/Elliptic_curve_point_multiplication)

1. 有限域上的椭圆曲线计算  
2. secp256k1 曲线  
3. 加密应用  