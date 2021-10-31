# UniqueWeb秋季招新

## 第一期任务

### 1.并查集

#### 数据结构：

* `int fa[n]`每个节点的父节点
* `int arc[n][n]`邻接矩阵

#### 基本函数

* 初始化：`fa[n]=n`
* 查询：判断两个节点是不是在一个连通图内，可以判断最小生成树有没有形成
* 合并：群体连接节点最小权重的节点

![image-20211031141703024](C:\Users\26254\AppData\Roaming\Typora\typora-user-images\image-20211031141703024.png)
