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

> 任意取椭圆曲线上两点 P, Q (若 P, Q 两点重合, 则做 P 点切线)  
> 过 P, Q 作直线与椭圆曲线相交于另一点 R'  
> 过 R' 做 y 轴平行线与椭圆曲线交于 R 点  

定义 `P + Q = R`

![img](http://andrea.corbellini.name/images/point-addition.png)

可知此加法也满足交换律结合律:  

![img](https://eng.paxos.com/hs-fs/hubfs/_02_Paxos_Engineering/Blockchain101-graphs-08.png?t=1524958561104&width=1280&name=Blockchain101-graphs-08.png)  
![img](https://eng.paxos.com/hs-fs/hubfs/_02_Paxos_Engineering/01-Blockchain101-graphs-09.png?t=1524958561104&width=1280&name=01-Blockchain101-graphs-09.png)  

`(A + B) + C = A + (B + C)`

则此群是一个阿贝尔群  

## secp256k1

比特币使用的曲线名称是 `secp256k1`, 其方程式为 `y^2=x^3+7`  

![img](https://eng.paxos.com/hs-fs/hubfs/_02_Paxos_Engineering/Blockchain101-graphs-06.png?t=1524958561104&width=1280&name=Blockchain101-graphs-06.png)



## 参考

[椭圆曲线介绍](http://andrea.corbellini.name/2015/05/17/elliptic-curve-cryptography-a-gentle-introduction/)  
[ref 偏理论](https://www.cnblogs.com/Kalafinaian/p/7392505.html)  
[ref 过程](http://www.freebuf.com/articles/database/155912.html)


连续空间上的椭圆曲线, 基本运算