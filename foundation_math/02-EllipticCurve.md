02 椭圆曲线
============

关于椭圆曲线比较全面的定义可以在 [MathWorld](http://mathworld.wolfram.com/EllipticCurve.html) 上查看.  

对于我们而言, 比较简单的定义是: **满足以下方程描述的点所形成的集合**  

> `y^2 = x^3 + ax + b 且 4a^3 + 27b^2 ≠ 0`

当 `4a^3 + 27b^2 = 0` 时, 曲线是[奇异曲线](https://en.wikipedia.org/wiki/Singularity_(mathematics)), 不在考量范围  

![img](https://upload.wikimedia.org/wikipedia/commons/thumb/d/db/EllipticCurveCatalog.svg/533px-EllipticCurveCatalog.svg.png)

可以看出, 椭圆曲线是 x 轴对称的, 同时我们需要一个[无穷远点](https://en.wikipedia.org/wiki/Point_at_infinity), 记做 **0** 或 **P∞**      

## 椭圆曲线上的群定律

由以下条件, 我们能够定义出椭圆曲线上的一个群:

> 1. 群的元素都是曲线上的点
> 2. 单位元是无穷远点 0
> 3. 曲线上点 P 的逆元是 P 点关于 x 轴对称的点

另外定义加法:

> 任意取椭圆曲线上两点 P, Q (若 P, Q 两点重合, 则作 P 点切线)  
> 过 P, Q 作直线与椭圆曲线相交于另一点 R'  
> 过 R' 做 y 轴平行线与椭圆曲线交于 R 点  

定义 `P + Q = R`

![img](http://andrea.corbellini.name/images/point-addition.png)

可知此加法也满足交换律结合律:  

![img](https://eng.paxos.com/hs-fs/hubfs/_02_Paxos_Engineering/Blockchain101-graphs-08.png?t=1524958561104&width=1280&name=Blockchain101-graphs-08.png)  
![img](https://eng.paxos.com/hs-fs/hubfs/_02_Paxos_Engineering/01-Blockchain101-graphs-09.png?t=1524958561104&width=1280&name=01-Blockchain101-graphs-09.png)  

`(A + B) + C = A + (B + C)`

则此群是一个阿贝尔群  

## 有限域上的椭圆曲线

`{ (x, y) ∈ GF(p) | y^2 = x^3 + ax + b (mod p), 4a^3 + 27b^2 ≠ 0 (mod p) } ∪ {0}`  

其中 0 是无穷远点; a, b 是 GF(p) 内的两个整数.  

![img](http://andrea.corbellini.name/images/elliptic-curves-mod-p.png)

曲线 `y^2=x^3-7x+10 (mod p), p=19,97,127,487` 的图像, 注意到对于每个 x 的值, 有对应的两个点, 且两点关于 `y = p/2` 对称  

### 点加法  

根据在实数域 `R` 上的加法, 如果椭圆曲线上的三点 P, Q, R 在同一直线 `ax + by + c = 0` 上, 可知 `P + Q + R = 0`  

在 `GF(p)` 上有类似的定义, 只是增加了一个条件, 即如果在`GF(p)`上的椭圆曲线上有过同一直线 `ax + by + c = 0 (mod p)` 的三点 P, Q, R, 则 `P + Q + R = 0`  

![img](http://andrea.corbellini.name/images/point-addition-mod-p.png)  

图中在椭圆曲线 `y^2 = x^3 - x + 3 (mod 127)` 上的点 `P(16, 20)`, `Q(41, 120)` 构成直线 `y = 4x + 83 (mod 127)`. 注意直线在平面上重复的样式    

同样的, 有限域上的椭圆曲线也是一个阿贝尔群, 满足以下性质:

> 1. `Q + 0 = 0 + Q = Q`
> 2. Q(xQ, yQ) 是曲线上的一个非 0 点, -Q = (xQ, -yQ mod p)
> 3. P + (-P) = 0

### 点的代数和

曲线上有点 `P = (xP, yP)`, `Q = (xQ, yQ)` 可以计算出 `P + Q = -R`:  

```plain
    xR = (m^2 - xP - xQ) mod p
    yR = [yP + m(xR - xP)] mod p
       = [yQ + m(xR - xQ)] mod p
       
P ≠ Q:
    m = (yP - yQ)(xP - xQ)^-1 mod p
P = Q:
    m = (3xP^2 + a)(2yP)^-1 mod p
```

### 标量乘法

`nP = P + P + ... + P, 其中 n 是自然数`  

  

### 有限域椭圆曲线群的阶  

有限域椭圆曲线群的阶就是曲线上所有点的个数, 我们可以通过计算 x = 0 到 x = p -1 对应的所有点来得到这个阶  

然而这种计算方式的复杂度是 `O(p)`, 而比特币使用的 secp256k1 曲线的 p 是 `2^256 - 2^32 - 977`  

另外一种更迅速的计算方式是使用 [Schoof 算法](https://en.wikipedia.org/wiki/Schoof%27s_algorithm)  

[实现](https://libraries.docs.miracl.com/miracl-user-manual/example-progs)  
[MIRACL](https://github.com/miracl/MIRACL)  

### 循环子群  





## 参考

[椭圆曲线介绍](http://andrea.corbellini.name/2015/05/17/elliptic-curve-cryptography-a-gentle-introduction/)  
[ref 偏理论](https://www.cnblogs.com/Kalafinaian/p/7392505.html)  
[ref 过程](http://www.freebuf.com/articles/database/155912.html)


连续空间上的椭圆曲线, 基本运算