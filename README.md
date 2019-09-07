# Bloom Filter

## 简介

Bloom Filter 由 Burton H. Bloom 在 1970 年的论文 [1] 中被首次提出，用于解决 membership 问题，即检测某数据 s 是已知集合 S 的成员。

### 问题界定

给定一个数据集合 S，检测新的数据 d 是否是集合 S 的成员。这里需要考虑的计算因素 (computational factors) 包括：

* 数据集合 S 所占用的存储空间（内存、外存）
* 成员检测所需时间
* False Positive Rate (FPR)，即认定属于集合 S，实际上不是的比例

其中 FPR 的计算公式为：

![fpr formular](./statics/imgs/fpr.jpg)

* n_a：将被判定为成员的数据总数
* n_t：数据总数
* n：集合 S 中的数据总数

### 解决方案

#### 1. 传统解法

假设数据集合 S 中有 n 条数据，那么我们需要新建一个大小为 h (h > n) 的哈希表，表中每个元素的大小不定。假设 S 中某条数据 s 的大小为 b bits，那么将其存入哈希表中将占用 b+1 bits空间，多出的 1 bits用于判定该位置是否为空。

**初始化哈希表过程**：将每条数据 s 输入哈希函数，生成 [0,  h-1] 中的某个数，即为 s 将被放入的位置，若该位置为空，则直接放入；否则继续生成一个新的位置，直到找到空位置为止 (解哈希冲突)。伪码如下：

```js
function add(s) {
	times = 1
	pos = hash(s, seed, times)
	while S[pos] != nil {
		times += 1
		pos = hash(s, seed, times)
	}
	S[pos] = s
}
```

**成员检测过程**：将新数据 s 输入哈希函数，生成对应的位置，若该位置为空，则判定 s 不属于 S；若位置不空且数据与 s 相等，则判定 s 属于 S，若位置不空且数据与 s 不等，则继续生成新位置，直到可以判定为止。伪码如下:

```js
function check(s) {
	times = 1
	pos = hash(s, seed, times)
	while S[pos] != nil && S[pos] != s {
		times += 1
		pos = hash(s, seed, times)
	}
	return S[pos] == s
}
```

## 参考

##### Papers

1. [Space/Time Trade-offs in Hash Coding with Allowable Errors](https://people.cs.umass.edu/~emery/classes/cmpsci691st/readings/Misc/p422-bloom.pdf)[Less Hashing, Same Performance: Building a Better Bloom Filter](https://www.eecs.harvard.edu/~michaelm/postscripts/rsa2008.pdf)
2. [Less Hashing, Same Performance: Building a Better Bloom Filter](https://www.eecs.harvard.edu/~michaelm/postscripts/rsa2008.pdf)

##### Blogs

* [Bloom Filters: Is element x in set S?](https://www.abhishek-tiwari.com/bloom-filters-is-element-x-in-set-s/)

##### Projects

* [github.com/willf/bloom](https://github.com/willf/bloom)
* [github.com/steakknife/bloomfilter](https://github.com/steakknife/bloomfilter)
* [github.com/zentures/bloom](https://github.com/zentures/bloom)