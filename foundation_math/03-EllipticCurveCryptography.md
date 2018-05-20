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

有限域上的椭圆曲线 `E(Fq)`, 域参数 `(p, a, b, G, n, h)`  

1. **私钥 private key** 是从 `{1, ..., n - 1}`, 其中 `n` 是子群的阶, 里随机选取的整数 `d`.
2. **公钥 public key** 是点 `H = dG`, 其中 `G` 是子群的基点.

可以看到, 当我们知道私钥的时候, 是能够通过椭圆曲线计算出公钥的.  

如果我们知道了 `d` 和 `G` (还有其他的域参数), 那么求出点 `H` 是非常容易的一件事情.  

但如果我们知道了 `H` 和 `G`, 想要找出私钥 `d`, 则会变得困难, 因为这涉及到求解离散对数问题.  

### ECDH 加密

ECDH 是 [Diffie-Hellman 算法](https://en.wikipedia.org/wiki/Diffie%E2%80%93Hellman_key_exchange)在椭圆曲线上的一个变种, 它实际上是一个密钥协商协议.  

该算法定义了密钥应当如何产生和交换, 如何使用这些密钥进行加密则取决于我们.  

ECDH 算法解决了以下问题: 如何在两个组织[Alice 和 Bob](https://en.wikipedia.org/wiki/Alice_and_Bob) 之间安全的交换信息, 同时保证即使有第三方 [Man in the middle](https://en.wikipedia.org/wiki/Man-in-the-middle_attack) 能够拦截这些信息, 但无法解读  

工作过程如下:  

1. Alice 和 Bob 在同一条椭圆曲线和域参数上分别生成各自的私钥 `dA`, `dB` 和公钥 `HA = dA*G`, `HB = dB * G`
2. Alice 和 Bob 在不安全的信道上交换各自的公钥, 中间人能够拦截到 `HA` 和 `HB`, 但他无法解出 `dA` 和 `dB`
3. Alice 通过 `S = dA*HB` 计算出共享密钥 `S`, Bob 通过 `S = dB*HA` 计算出共享密钥 `S`

注意 Alice 和 Bob 计算出来的 `S` 是一样的, 因为:

> ` S = dA*HB = dA(dB*G) = dB(dA*G) = dB*HA`

而中间人知道 `HA` 和 `HB`, 却无法计算出共享秘钥 `S`, 这被称为 Diffie-Hellman 问题, 描述如下:  

> 给定三个点 `P`, `aP` 和 `bP`, 求 `abP`

或等价的:

> 给定三个整数 `k`, `k^x`, `k^y`, 求 `k^xy`

后者在原始的 Diffie-Hellman 算法当中得到应用, 基于模运算的原理  

![img](http://andrea.corbellini.name/images/ecdh.png)  

原始的 DH 算法例子:  

1. Alice 和 Bob 有以下共识: `p=23`, `g=5`
2. Alice 选择了一个私钥 `a=6`, 并计算出公钥 `A = g^a mod p = 5^6 % 23 = 8`
3. Bob 选择了一个私钥 `b=15`, 并计算出公钥 `B = g^b mod p = 5^15 % 23 = 19`
4. Alice 和 Bob 交换公钥, 而中间人拦截到 `A=8`, `B=19`
5. Alice 和 Bob 计算出共享密钥 S = B^a mod p = 2 = A^b mod p = 2`
6. 中间人很难计算出 `S = 2`

现在 Alice 和 Bob 拥有了共享秘钥, 现在他们能够使用对称加密, 比如 AES 和 3DES 来进行安全的通信了.  

### 单向陷门函数 One way trapdoor function

在研究密码学的过程中我们经常会接触到几个名词, `单向函数`, `陷门函数`, `门函数` 和`单向陷门函数`  

其实一共就两个概念:  

1. 单向函数 One way function, 给定 `x`, 计算 `f(x)` 很容易, 而给定 `y`, 在 `f` 的范围内, 很难找到 `x` 使 `f(x) = y`
2. 单向陷门函数, 跟单向函数类似, 但是当知道了一些额外的信息(比如域参数)之后, 可以简单的进行反向计算

### ECDH 例子

我们选取 bitcoin 使用的椭圆曲线 `secp256k1` 来说明. `secp256k1` 的域参数如下:

* `p  =  0xFFFFFFFF FFFFFFFF FFFFFFFF FFFFFFFF FFFFFFFF FFFFFFFF FFFFFFFE FFFFFC2F`
* `a  = 0`
* `b  = 7`
* `xG = 0x79BE667E F9DCBBAC 55A06295 CE870B07 029BFCDB 2DCE28D9 59F2815B 16F81798`
* `yG = 0x483ADA77 26A3C465 5DA4FBFC 0E1108A8 FD17B448 A6855419 9C47D08F FB10D4B8`
* `n  = 0xFFFFFFFF FFFFFFFF FFFFFFFF FFFFFFFE BAAEDCE6 AF48A03B BFD25E8C D0364141`
* `h  = 1`

TODO: make a ECDH example here  

[ref](http://andrea.corbellini.name/2015/05/30/elliptic-curve-cryptography-ecdh-and-ecdsa/)  
[计算](https://en.wikipedia.org/wiki/Elliptic_curve_point_multiplication)

1. 有限域上的椭圆曲线计算  
2. secp256k1 曲线  
3. 加密应用  