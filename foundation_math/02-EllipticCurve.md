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

有限域上的椭圆曲线可以看做是该曲线在有限域内的点的集合

`{ (x, y) ∈ GF(p) | y^2 = x^3 + ax + b (mod p), 4a^3 + 27b^2 ≠ 0 (mod p) } ∪ {0}`

其中 0 是无穷远点; a, b 是 GF(p) 内的两个整数.

也可以记作 `E(Fq)`, 其中 `q = p^n, n ≥ 1, n ∈ N`

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

### 有限域椭圆曲线群的阶

有限域椭圆曲线群的阶就是曲线上所有点的个数, 这个值是衡量一个椭圆曲线的密码学强度的一个重要指标

椭圆曲线 `E(Fq)` 的阶记做 `#E(Fq)`

我们可以通过计算 x = 0 到 x = p -1 对应的所有点来得到这个阶, 然而这种计算方式的复杂度是 `O(p)`

而比特币使用的 secp256k1 曲线的 p 是 `2^256 - 2^32 - 977`, 这样的计算方式显然效率太低

另一种更迅速的计算方式是使用 [Schoof 算法](https://en.wikipedia.org/wiki/Schoof%27s_algorithm)和其改进算法 [SEA 算法](https://en.wikipedia.org/wiki/Schoof%E2%80%93Elkies%E2%80%93Atkin_algorithm)


[Hasse 理论](https://en.wikipedia.org/wiki/Hasse%27s_theorem_on_elliptic_curves):

> 有有限域 Fq 上的椭圆曲线 E/Fq , 其阶 #E(Fq) 满足
> `|q + 1 - #E(Fq)| ≤ 2√q`

这样能够将曲线的阶限定在一个相对有限的范围内, 随后通过余数定理和其他定理就能够计算出 `#E(Fq)`

Schoof 算法时间复杂度满足[多项式时间](https://en.wikipedia.org/wiki/Time_complexity#Polynomial_time)

### 标量乘法与循环子群

`nP = P + P + ... + P, 其中 n 是自然数`

我们可以通过 double and add 算法来计算

例如当 `n = 151` 时, 其二进制表示为 `10010111b`, 可以表示为:

`151 = 1*2^7 + 0*2^6 + 0*2^5 + 1*2^4 + 0*2^3 + 1*2^2 + 1*2^1 + 1*2^0`

则有:

`151*P = 2^7*P + 2^4*P + 2^2*P + 2^1*P + 2^0*P`

```go
func DoubleAndAdd(n uint, x int) int {
	result := 0
	addend := x
	for i := 0; i < bits.Len(n); i++ {
		b := n >> uint(i) & 0x1
		if b == 1 {
			result += addend
		}
		addend *= 2
	}

	return result
}
```

`E(Fq)` 上点的标量乘法又一个非常有趣的性质. 有曲线 `y^2 = x^3 + 2x + 3 (mod 97)` 和曲线上一点 `P = (3, 6)`

有:

![img](http://andrea.corbellini.name/images/cyclic-subgroup.png)

* `0P = 0`
* `1P = (3, 6)`
* `2P = (80, 10)`
* `3P = (80, 87)`
* `4P = (3, 91)`
* `5P = 0`
* `6P = (3, 6)`
* `7P = (80, 10)`
* `8P = (80, 87)`
* `9P = (3, 91)`
* `10P = 0`
* ...

可以看到, `P` 的标量乘法结果只有5个点, 而且是循环的:

* `5k*P = 0`
* `(5k + 1)*P = P`
* `(5k + 2)*P = 2P`
* `(5k + 3)*P = 3P`
* `(5k + 4)*P = 4P`

这个结论对 `E(Fq)` 上的所有点都适用, 我们任取一点 `P`:

> `nP + mP = (P + 共 n 个 + P) + (P + 共 m 个 + P) = (n + m)P`

即将两个 `P` 的标量积相加, 得到 P 的另一个标量积

这也就[说明](https://en.wikipedia.org/wiki/Subgroup#Basic_properties_of_subgroups)了点 `P` 的标量积构成了 `E(Fq)` 的一个循环子群

点 `P` 则被称为子群的生成点 generator 或基点 base point

子群的概念是椭圆曲线加密 ECC 和其他密码学系统的关键基础

### 子群的阶

椭圆曲线 `E(Fq)` 上一点 `P` 生成子群 `S`

`S` 的阶也可以称为基点 `P` 的阶

有:

* 一定存在一个最小的正整数 `n`, 使得 `nP = 0`, n 是基点 `P` 的阶
* 根据[拉格朗日定理](https://en.wikipedia.org/wiki/Lagrange%27s_theorem_(group_theory)) `n` 必是 `#E(Fq)` 的约数

根据上面两个性质, 我们可以计算出点 `P` 的阶:

1. 利用 schoof 或其他算法求得椭圆曲线的阶 `N = #E(Fq)`
2. 找出 `N` 的所有约数
3. 对每一个 `{n| n ∈ N}` 计算 `nP`
4. 使得 `nP = 0` 的**最小**的 n 就是点 `P` 的阶

例:

`y^2 = x^3 - x + 3 (mod 37)` 的阶 `N = 42`. 子群的阶 `n` 有可能是 `1, 2, 3, 6, 7, 14, 21 或 42`.

如果基点为 `P = (2, 3)` 则有 `1P ≠ 0, 2P ≠ 0, ..., 7P = 0`, 即 `P` 的阶是 7.

### 确定基点

对于 ECC 算法而言, 我们希望子群的阶越大越好. 一般流程是这样:

1. 先确定椭圆曲线 `E(Fq)`, 计算它的阶 `N = #E(Fq)`
2. 取较大的 `N` 的余数 `n` 作为子群的阶
3. 根据 `n` 确定基点 `P`

令:

> `h = N/n`

根据拉格朗日定理, `n` 是 `N` 的约数, 则 `h` 必为正整数, 我们称 `h` 为子群的**余因子** (cofactor of the subgroup)

对于 `E(Fq)` 上的任意一点 `P`, 有 `NP = 0` 即 `n(hP) = 0`

假定 `n` 是一个素数 (原因后续介绍), 上述等式的意义就是: 对于点 `G = hP`, 构成了一个阶为 `n` 的子群 (`G = hP = 0` 除外, 此时子群的阶为 1).

由此, 我们可以勾勒出下面的算法:

1. 计算 `E(Fq)` 的阶 `N`
2. 选取子群的阶 `n`, `n` 是素数, 且是 `N` 的约数
3. 计算子群的余因子 `h = N/n`
4. 在 `E(Fq)` 上随机选取点 `P`
5. 计算 `G = hP`
6. 如果 `G = 0` 则退回第4步, 否则 `G` 点的阶就是 `n` 

### 离散对数 Discrete logarithm



## 参考

[椭圆曲线介绍](http://andrea.corbellini.name/2015/05/17/elliptic-curve-cryptography-a-gentle-introduction/)
[ref 偏理论](https://www.cnblogs.com/Kalafinaian/p/7392505.html)
[ref 过程](http://www.freebuf.com/articles/database/155912.html)
