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

### 随机曲线

上一篇文章的结尾我们介绍了椭圆曲线上的对数问题, 以及一些特定的曲线容易受到攻击, 那么我们有没有办法确定一个椭圆曲线是"安全"的呢?  

为了解决这类问题, 我们经常需要引入一个新的域参数: 用来生成椭圆曲线系数 `a`, `b` 和/或基点 `G` 的种子随机数(seed) **`S`**.  

![img](http://andrea.corbellini.name/images/random-parameters-generation.png)  

通过对 `S` 哈希得到 `a` 和 `b` 非常简单  

![img](http://andrea.corbellini.name/images/seed-inversion.png)

而试图通过 `a` 和 `b` 反向推出 `S` 却非常困难  

种子数的选取应当遵循 (nothing up my sleeve)[https://en.wikipedia.org/wiki/Nothing_up_my_sleeve_number] 原则.  

## 椭圆曲线加密

> per aspera ad astra

有有限域上的椭圆曲线 `E(Fq)`, 域参数 `(p, a, b, G, n, h)`  

1. **私钥 private key** 是从 `{1, ..., n - 1}`, 其中 `n` 是子群的阶, 里随机选取的整数 `d`.
2. **公钥 public key** 是点 `H = dG`, 其中 `G` 是子群的基点.

### ECDH 加密

[ref](http://andrea.corbellini.name/2015/05/30/elliptic-curve-cryptography-ecdh-and-ecdsa/)  
[计算](https://en.wikipedia.org/wiki/Elliptic_curve_point_multiplication)

1. 有限域上的椭圆曲线计算  
2. secp256k1 曲线  
3. 加密应用  