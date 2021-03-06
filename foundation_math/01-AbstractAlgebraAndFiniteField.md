@(Coding)[学习笔记|文章]

[TOC]

01 抽象代数基础与有限域
========================

## 抽象代数基础

抽象代数作为数学的一门学科, 主要研究对象是代数结构, 比如群, 环, 域, 模, 向量空间, 格与域代数.  

"抽象代数"一词出现于20世纪初, 作为与其他代数领域相区别之学科.  

代数结构与其相关之同态, 构成数学范畴. 范畴论是用来分析与比较不同代数结构的强大形式工具.

### 集合 (Set)

> 1. 确定性. 给定一个集合, 对任意一个元素, 这个元素要么属于这个集合, 要么不属于这个集合, 不存在其他情况.
> 2. 互异性. 集合中任何两个元素都是不相同的.
> 3. 无序性. 集合中的元素是无序的.

常见的集合有: 整数集 `Z`, 有理数集 `Q`, 实数集 `R` 等.

### 半群 (Semigroup)

在一个集合 S 中定义了某种运算, 记做 `+`, 那么在这个集合上, 如果这种运算满足以下性质, 那么它和集合 S 共同组成一个**半群**, 记做 `(S, +)`

> 1. 封闭性. 运算的结果始终在集合 S 内
> 2. 结合律. 满足 `(a + b) + c = a + (b + c)`

例: 如果集合 S 是全体实数(记做"R"), 而是运算是实数加法, 那么它们共同形成了一个半群, 记做 `(R, +)`

### 幺半群 (Monoid)

如果一个半群 `(S, +)` 中存在一个元素 **e**, 使得 S 中所有的元素 a 都满足:

> `a + e = e + a = a`

则该半群为**幺半群**, 元素**e**被称为**单位元**或者**幺元**.

例: `(R, +)` 中, 实数 0 符合这一要求, 所以 `(R, +)` 是幺半群, 0 是它的单位元.

### 群 (Group)

如果一个幺半群 `(S, +)` 中的每一个元素 a 都有唯一一个元素 b 与之对应且满足

> `a + b = b + a = e , 其中 e 为单位元`

则该幺半群就是一个**群**. 群中元素 a 和元素 b 互为**逆元**, 记作 `a = -b` 或者 `b = -a`.

逆元存在, 也就定义了群上的减法. a 减去 b, 其实就是 a 加上 b 的逆元.

> `a - b = a + (-b)`

例: `(R, +)` 中, 每一个正数都和一个负数一一对应, 他们的和为 0. 0 取负是它自身. 所以 `(R, +)` 是一个群.

### 交换群 (Commutative/Abelian Group)

如果一个群 `(S, +)` 符合交换律, 即对于集合中任意元素 a 和 b, 满足:

> `a + b = b + a`

那么这个群被称为**交换群**, 又叫**阿贝尔群**.

### 环 (Loop)

在一个交换群 `(S, +)` 上, 再定义另外一种运算, 记做 `*`, 得到 `(S, +,*)`. 如果 `(S, +,*)` 满足以下性质

> 1. `(S, *)` 是一个幺半群
> 2. 两种运算满足分配率, 即 `a(b + c) = ab + ac`

那么 `(S, +,*)` 形成一个**环**. 此时群 `(S, +)` 的单位元被称为环 `(S, +,*)` 的**零元**.

### 除环 (Division Ring)

> 如果幺半群 `(S, *)` 里除了**零元**以外的所有元素都有**逆元**, 那么 `(S, +,*)` 被称为**除环**.

为了避免和 `(S, +)` 里的逆元混淆, `(S, +)` 里的逆元称为加法逆元, `(S, *)` 里的则是乘法逆元.

`(S, +,*)` 中乘法逆元存在, 也就定义了除法, 元素 a 除以元素 b 实际上就是 a 乘以 b 的乘法逆元. 也就是

> `a ÷ b = ab^-1`

### 交换环 (Commutative Ring)

> 如果环 `(S, +,*)` 中, `(S, *)` 满足交换律, 那么 `(S, +,*)` 被称为交换环.

例: 实数乘法满足交换律, 所以 `(R, +,*)` 是一个交换环.

### 域 (Field)

> 如果一个环 `(S, +,*)`, 既是除环又是交换环, 那么它是一个域

例: `(R, +,*)` 既是除环又是交换环, 所以它是一个域, 称为**实数域**

## 有限域 (Finite Field)

如果一个域只包含有限个元素, 则其为**有限域**或**伽罗瓦域**

有限域中元素的个数称为该有限域的**阶(order)**

有限域的阶一定是某个素数 p 的 k 次幂 (k 是正整数), p 也被称为有限域的**特征值(characteristic)**

### 模 p 有限域 GF(p)

`GF(p)` 是定义在整数集合 `{0, 1, ..., p-1}` 上的域. `GF(p)` 上的加法和乘法分别是模加法和模乘法.

### 模加法和模乘法

模加法和模乘法和普通的整数加法乘法类似, 不同的是要对运算结果对素数 p 取模.

例: `GF(7)` 的加法和乘法

```plain
1 + 2 = 3  mod 7 = 3
5 + 6 = 11 mod 7 = 4
1 * 2 = 2  mod 7 = 2
4 * 5 = 20 mod 7 = 6
```

### 模减法

a 减去 b, 就是 a 加上 b 的加法逆元.

```plain
1 - 2 = 1 + (-2) = 1 + 5 =  6 mod 7 = 6
                   1 - 2 = -1 mod 7 = 6
```

### 模除法

a 除以 b，需要找到 b 的乘法逆元 (在这里又被称为数论倒数). 即满足以下式子的整数 x:

> `bx = 1 mod 7`

也就是解以下方程:

> `bx = 1 + 7k, 其中 k ∈ Z+`

求解需要用到[扩展欧几里得算法(Extended Euclidean Algorithm)](https://zh.wikipedia.org/zh-cn/%E6%89%A9%E5%B1%95%E6%AC%A7%E5%87%A0%E9%87%8C%E5%BE%97%E7%AE%97%E6%B3%95)

> ` 3 / 4 = 3 * (4^-1) = 3 * 2 = 6 mod 7 = 6`

### 有限域 GF(2^m)

`GF(2^m)` 在 [Reed-Solomon 编码](https://en.wikipedia.org/wiki/Reed%E2%80%93Solomon_error_correction)与椭圆曲线加密中都有应用

`GF(2^m)` 的元素是二进制多项式, 即多项式的系数不是 1 就是 0, 域中一共有 2^m 个多项式, 每个多项式的最高次不超过 m-1.

例: `GF(2^3)` 包含 8 个元素, `{0, 1, x, x+1, x^2, x^2+1, x^2+x, x^2+x+1}`, `x+1` 实际上是 `0x^2+1x+1`, 则其可以表示为 `011`

### 加减法

当模为 2 时

> `1 + 1 = 2 mod 2 = 0`
> `1 + 0 = 1 mod 2 = 1`
> `0 + 0 = 0 mod 2 = 0`

而模减法, 则有

> `-1 = 0 - 1 = 0 + 1(-1 的加法逆元) = 1 mod 2 = 1`

可知, 对于 `GF(2^m)` 上的模加减法, 均等价为异或运算

例:

> `(x^2 + x + 1) + (x + 1) = x^2 + 2x + 2` 而 `2 mod 2 = 0`, 结果为 `x^2`, 等价于
> `111 xor 011 = 100`, 100 就是 x^2 的 bit 字符串表示

### 乘法

二进制多项式乘法可以通过位移和异或计算得出

例: 对于 GF(2^3)

```plain
(x^2 + x + 1)*(x^2 + 1) = x^4 + x^3 + 2x^2 + x + 1 = x^4 + x^3 + x + 1
   111
x  101
------
   111
  0000
 11100
------
 11011
```

在 `GF(2^m)` 上, 当结果的度(degree) 超过 m-1 时, 需要将结果对一个不可约多项式进行取模运算.

这个运算可以通过位移和异或实现.

例如 `x^3 + x + 1` 是一个不可约多项式 (另一个是 `x^3 + x^2 + 1`), 记作 1011, 因为 11011 的度是 4 而 1011 的度是 3

则计算过程以将 1011 左移 1 位开始

```plain
    11011
xor 10110
---------
    01101 度仍然大于 m - 1 = 2, 但此时 1011 不需要再左移了
xor 01011
---------
    00110
```

得 `111*101 = 11011 mod 1011 = 110 即 x^2 + x`

### 除法

除法依然是等效于乘以乘法逆元, 可以通过[扩展欧几里得算法](https://en.wikipedia.org/wiki/Extended_Euclidean_algorithm)计算

此处不再赘述

## 参考

[Wiki 上的抽象代数](https://zh.wikipedia.org/zh-cn/%E6%8A%BD%E8%B1%A1%E4%BB%A3%E6%95%B0)  
[Wiki 上的有限域计算](https://en.wikipedia.org/wiki/Finite_field_arithmetic)  
[谈谈有限域那些事儿](https://blog.csdn.net/qmickecs/article/details/77281602)  
[有限域上的代数运算](https://www.doc.ic.ac.uk/~mrh/330tutor/ch04s04.html)  
[资料](https://engineering.purdue.edu/kak/compsec/NewLectures/Lecture7.pdf)  
[资料](https://sites.math.washington.edu/~morrow/336_12/papers/juan.pdf)  
