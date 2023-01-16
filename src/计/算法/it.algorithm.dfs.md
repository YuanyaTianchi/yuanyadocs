

+++

title = "algorithm.dfs"
description = "algorithm.dfs"
tags = ["it", "Algorithm"]

+++

# algorithm.dfs

## DFS

- DFS：深度优先遍历，也叫回溯算法。不像动态规划存在重叠子问题可以优化，回溯算法就是纯暴力穷举，复杂度一般都很高。某种程度上来看，动态规划的暴力求解阶段就是回溯算法



- 过程：**解决一个回溯问题，实际上就是一个决策树的遍历过程**
  - 路径：记录已经做过的选择
  - 选择列表：当前节点可以做的选择。
  - 结束条件：到达决策树底层，无法再做选择的条件。
- 代码实现步骤：**其核心是 for 循环里面的递归，在递归调用之前「做选择」，在递归调用之后「撤销选择」**
  - 参数：将「选择列表」和「路径」作为递归函数的参数
  - 结束条件：也表现为递归的结束条件，到达决策树底层，将路径添加到结果集
  - 循环递归选择：**核心就是 for 循环中，在递归调用之前「做选择」，在递归调用之后「撤销选择」**
    - 进入循环时，如果「当前选择」已经在「路径」中了（可以用单独写一个方法遍历判断），表示已经选择过了，直接continue跳过即可。比如【全排列】【N皇后】中
    - 或者也可以，「做选择」时，先将该选择从选择列表移除，然后将选择添加到路径；「撤销选择」时，先从路径中移除选择，再将该选择再加入选择列表
    - 当然还有有更好的方法，通过交换元素达到目的

```python
result = []
def backtrack(路径, 选择列表):
    if 满足结束条件，到达决策树底层，无法再做选择的条件:
        result.add(路径)
        return

    for 选择 in 选择列表:
        if 路径 包含 选择:
            continue

        路径.add(选择)
        backtrack(路径, 选择列表)
        路径.remove(选择)
```





### 全排列

- 列出`n` 个不重复的数的全排列
- 比方说给三个数 `[1,2,3]`，你肯定不会无规律地乱穷举，一般是这样：先固定第一位为 1，然后第二位可以是 2，那么第三位只能是 3；然后可以把第二位变成 3，第三位就只能是 2 了；然后就只能变化第一位，变成 2，然后再穷举后两位……，就形成了一个决策树，为啥说这是决策树呢，因为你在每个节点上其实都在做决策，**定义的** **`backtrack`** **函数其实就像一个指针，在这棵树上游走，同时要正确维护每个节点的属性，每当走到树的底层，其「路径」就是一个全排列**。
  - 路径：记录在 track 中
  - 选择列表：nums 中不存在于 track 的那些元素
  - 结束条件：nums 中的元素全都在 track 中出现



### 全排列

题目：[46. 全排列](https://leetcode-cn.com/problems/permutations/)

```go
var res = make([][]int, 0)

func permute(nums []int) [][]int {
	res = make([][]int, 0)
	path := make([]int, 0)
	tr(path, nums)
	return res
}

func tr(path []int, nums []int) {
	if len(nums) == 0 {
		pathRes := make([]int, len(path))
		copy(pathRes, path)
		res = append(res, pathRes)
		return
	}

	for i, _ := range nums {
		num := nums[i]
		path = append(path, num)

		// 需要 copy 后操作
		numsNew := make([]int, len(nums))
		copy(numsNew, nums)
		numsNew = append(numsNew[:i], numsNew[i+1:]...)

		tr(path, numsNew)
		path = path[:len(path)-1]
	}
}
```

### 全排列 II

题目：[47. 全排列 II](https://leetcode-cn.com/problems/permutations-ii/)

思路1：使用全排列去重，简单粗暴，效率低

思路2

- 当选择一个数字时，记录该数字，如果（决策树）同层、同前缀（）再次遇到该数字，则表示重复，跳过即可，即剪枝技巧
- 剪枝思路：同前缀的当前层记录该层使用过的数字，即一个递归的循环种要维持一个数组，即一个[][]int 或 []map[]

```go
var res [][]int

func permuteUnique(nums []int) [][]int {
	res = make([][]int, 0)
	path := make([]int, 0)
	visited := make([]map[int]bool, len(nums))
	tr(nums, path, visited, 0)
	return res
}

func tr(nums []int, path []int, visited []map[int]bool, deep int) {
	if len(nums) == 0 {
		pathCopy := make([]int, len(path))
		copy(pathCopy, path)
		res = append(res, pathCopy)
		return
	}

	// 每次进来都会重置所有元素为 false，因为即使同层，但跟之前已经不是相同前缀了
	visited[deep] = make(map[int]bool)

	for i, _ := range nums {
		num := nums[i]

		if visited[deep][num] {
			continue
		}
		visited[deep][num] = true

		numsCopy := make([]int, len(nums))
		copy(numsCopy, nums)
		numsCopy = append(numsCopy[:i], numsCopy[i+1:]...)

		path = append(path, nums[i])
		tr(numsCopy, path, visited, deep+1)
		path = path[:len(path)-1]
	}
}
```



### N皇后

-  N×N 的棋盘防止N个皇后，使得它们不能互相攻击，问有多少种放法。皇后可以攻击 同一行、同一列、左上、左下、右上、右下 8个方向任意距离的单位
- 本质上跟全排列问题差不多，决策树的每一层表示棋盘上的每一行；每个节点可以做出的选择是，在该行的任意一列放置一个皇后
  - 路径：board 中小于 row 的那些行都已经成功放置了皇后
  - 选择列表：第 row 行的所有列都是放置皇后的选择
  - 结束条件：row 超过 board 的最后一行
- 最坏时间复杂度仍然是 O(N^(N+1))，而且无法优化

```c
vector<vector<string>> res;

/* 输入棋盘边长 n，返回所有合法的放置 */
vector<vector<string>> solveNQueens(int n) {
    // '.' 表示空，'Q' 表示皇后，初始化空棋盘。
    vector<string> board(n, string(n, '.'));
    backtrack(board, 0);
    return res;
}

void backtrack(vector<string>& board, int row) {
    // 触发结束条件
    if (row == board.size()) {
        res.push_back(board);
        return;
    }

    int n = board[row].size();
    for (int col = 0; col < n; col++) {
        // 排除不合法选择
        if (!isValid(board, row, col)) 
            continue;
        // 做选择
        board[row][col] = 'Q';
        // 进入下一行决策
        backtrack(board, row + 1);
        // 撤销选择
        board[row][col] = '.';
    }
}

/* 是否可以在 board[row][col] 放置皇后？ */
bool isValid(vector<string>& board, int row, int col) {
    int n = board.size();
    // 检查列是否有皇后互相冲突
    for (int i = 0; i < n; i++) {
        if (board[i][col] == 'Q')
            return false;
    }
    // 检查右上方是否有皇后互相冲突
    for (int i = row - 1, j = col + 1; 
            i >= 0 && j < n; i--, j++) {
        if (board[i][j] == 'Q')
            return false;
    }
    // 检查左上方是否有皇后互相冲突
    for (int i = row - 1, j = col - 1;
            i >= 0 && j >= 0; i--, j--) {
        if (board[i][j] == 'Q')
            return false;
    }
    return true;
}
```

- 有的时候，我们并不想得到所有合法的答案，只想要一个答案。比如解数独的算法，找所有解法复杂度太高，只要找到一种解法就可以。只要稍微修改一下回溯算法的代码

```c
// 函数找到一个答案后就返回 true
bool backtrack(vector<string>& board, int row) {
    // 触发结束条件
    if (row == board.size()) {
        res.push_back(board);
        return true;
    }
    ...
    for (int col = 0; col < n; col++) {
        ...
        board[row][col] = 'Q';

        if (backtrack(board, row + 1))
            return true;

        board[row][col] = '.';
    }

    return false;
}
```



## BFS

- 常见场景：**问题的本质就是让你在一幅「图」中找到从起点** **`start`** **到终点** **`target`** **的最近距离，这个例子听起来很枯燥，但是 BFS 算法问题其实都是在干这个事儿**
  - 广义的描述可以有各种变体，比如走迷宫，有的格子是围墙不能走，从起点到终点的最短距离是多少？如果这个迷宫带「传送门」可以瞬间传送呢？
  - 比如说两个单词，要求你通过某些替换，把其中一个变成另一个，每次只能替换一个字符，最少要替换几次？
  - 比如说连连看游戏，两个方块消除的条件不仅仅是图案相同，还得保证两个方块之间的最短连线不能多于两个拐点。你玩连连看，点击两个坐标，游戏是如何判断它俩的最短连线有几个拐点的？
- **BFS 找到的路径一定是最短的，但代价就是空间复杂度比 DFS 大很多**
  - BFS 可以找到最短距离，但是空间复杂度高。处理二叉树问题的例子，假设给你的这个二叉树是满二叉树，节点数为 `N`，对于 DFS 算法来说，空间复杂度无非就是递归堆栈，最坏情况下顶多就是树的高度，也就是 `O(logN)`，
  - DFS 不能找最短路径吗？其实也是可以的，但是时间复杂度相对高很多。你想啊，DFS 实际上是靠递归的堆栈记录走过的路径，你要找到最短路径，肯定得把二叉树中所有树杈都探索完才能对比出最短的路径有多长




- 过程：就是在一个图中从一个start节点开始**扩散**，直到找到target节点，通过一个**队列**存储下一步要扩散的节点，通过一个**集合**存储所有已经走过的点，以避免走回头路
  - **要扩散的节点**：一般用一个队列存储，也可以是链表，或者其它的方便增删元素的结构
  - **已扩散的节点**：

- 代码实现步骤：
  - 声明队列q、集合visited、扩散步数step，将start节点加入 q 和 visited，因为直到start就要扩散了，就直接加入到已扩散的节点无妨，当然改变一下位置在循环中再加入visited也无妨
  - 在q.len=0之前一直循环：
    - 遍历q：
      - 移除元素，表示已扩散
      - 判断是否到是target节点，到达就返回step
      - 遍历元素的相邻节点
        - 如果不在已扩散的节点中，则将该相邻节点加入

```c++
// 计算从起点 start 到终点 target 的最近距离
int BFS(Node start, Node target) {
    Queue<Node> q; // 核心数据结构
    Set<Node> visited; // 避免走回头路，即避免再次走到走过的节点

    q.offer(start); // 将起点加入队列
    visited.add(start);
    int step = 0; // 记录扩散的步数

    while (q not empty) {
        int sz = q.size();
        /* 将当前队列中的所有节点向四周扩散 */
        for (int i = 0; i < sz; i++) {
            Node cur = q.poll();
            /* 划重点：这里判断是否到达终点 */
            if (cur is target)
                return step;
            /* 将 cur 的相邻节点加入队列。cur.adj() 泛指 cur 相邻的节点 */
            for (Node x : cur.adj())
                if (x not in visited) {
                    q.offer(x);
                    visited.add(x);
                }
        }
        /* 划重点：更新步数在这里 */
        step++;
    }
}
```



##### 二叉树的最小深度

https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/

- **显然起点就是** **`root`** **根节点，终点就是最靠近根节点的那个「叶子节点」嘛**，叶子节点就是两个子节点都是 `null` 的节点

```go
int minDepth(TreeNode root) {
    if (root == null) return 0;
    Queue<TreeNode> q = new LinkedList<>();
    q.offer(root);
    // root 本身就是一层，depth 初始化为 1
    int depth = 1;

    while (!q.isEmpty()) {
        int sz = q.size();
        /* 将当前队列中的所有节点向四周扩散 */
        for (int i = 0; i < sz; i++) {
            TreeNode cur = q.poll();
            /* 判断是否到达终点 */
            if (cur.left == null && cur.right == null) 
                return depth;
            /* 将 cur 的相邻节点加入队列 */
            if (cur.left != null)
                q.offer(cur.left);
            if (cur.right != null) 
                q.offer(cur.right);
        }
        /* 这里增加步数 */
        depth++;
    }
    return depth;
}


```



##### 打开转盘锁

https://leetcode-cn.com/problems/open-the-lock/

```go
int openLock(String[] deadends, String target) {
    // 记录需要跳过的死亡密码
    Set<String> deads = new HashSet<>();
    for (String s : deadends) deads.add(s);
    
    // 记录已经穷举过的密码，防止走回头路
    Set<String> visited = new HashSet<>();
    Queue<String> q = new LinkedList<>();
    
    // 从起点开始启动广度优先搜索
    int step = 0;
    q.offer("0000");
    visited.add("0000");

    while (!q.isEmpty()) {
        int sz = q.size();
        /* 将当前队列中的所有节点向周围扩散 */
        for (int i = 0; i < sz; i++) {
            String cur = q.poll();

            /* 判断是否到达终点 */
            if (deads.contains(cur))
                continue;
            if (cur.equals(target))
                return step;

            /* 将一个节点的未遍历相邻节点加入队列 */
            for (int j = 0; j < 4; j++) {
                String up = plusOne(cur, j);
                if (!visited.contains(up)) {
                    q.offer(up);
                    visited.add(up);
                }
                String down = minusOne(cur, j);
                if (!visited.contains(down)) {
                    q.offer(down);
                    visited.add(down);
                }
            }
        }
        /* 在这里增加步数 */
        step++;
    }
    // 如果穷举完都没找到目标密码，那就是找不到了
    return -1;
}
```



##### 双向 BFS 优化

- **无论传统 BFS 还是双向 BFS，无论做不做优化，从 Big O 衡量标准来看，时间复杂度都是一样的**，只能说双向 BFS 是一种 trick，算法运行的速度会相对快一点
- **传统的 BFS 框架就是从起点开始向四周扩散，遇到终点时停止；而双向 BFS 则是从起点和终点同时开始扩散，当两边有交集的时候停止**。
- 双向 BFS 还是遵循 BFS 算法框架的，只是**不再使用队列，而是使用 HashSet 方便快速判断两个集合是否有交集**。另外的一个技巧点就是 **while 循环的最后交换** **`q1`** **和** **`q2`** **的内容**，所以只要默认扩散 `q1` 就相当于轮流扩散 `q1` 和 `q2`。
- **不过，双向 BFS 也有局限，因为你必须知道终点在哪里**。二叉树最小高度的问题，你一开始根本就不知道终点在哪里，也就无法使用双向 BFS；但是第二个密码锁的问题，是可以使用双向 BFS 算法来提高效率的

```go
int openLock(String[] deadends, String target) {
    Set<String> deads = new HashSet<>();
    for (String s : deadends) deads.add(s);
    // 用集合不用队列，可以快速判断元素是否存在
    Set<String> q1 = new HashSet<>();
    Set<String> q2 = new HashSet<>();
    Set<String> visited = new HashSet<>();

    int step = 0;
    q1.add("0000");
    q2.add(target);

    while (!q1.isEmpty() && !q2.isEmpty()) {
        // 哈希集合在遍历的过程中不能修改，用 temp 存储扩散结果
        Set<String> temp = new HashSet<>();

        /* 将 q1 中的所有节点向周围扩散 */
        for (String cur : q1) {
            /* 判断是否到达终点 */
            if (deads.contains(cur))
                continue;
            if (q2.contains(cur))
                return step;
            visited.add(cur);

            /* 将一个节点的未遍历相邻节点加入集合 */
            for (int j = 0; j < 4; j++) {
                String up = plusOne(cur, j);
                if (!visited.contains(up))
                    temp.add(up);
                String down = minusOne(cur, j);
                if (!visited.contains(down))
                    temp.add(down);
            }
        }
        /* 在这里增加步数 */
        step++;
        // temp 相当于 q1
        // 这里交换 q1 q2，下一轮 while 就是扩散 q2
        q1 = q2;
        q2 = temp;
    }
    return -1;
}
```

- 双向 BFS 还有一个优化，就是在 while 循环开始时做一个判断：因为按照 BFS 的逻辑，队列（集合）中的元素越多，扩散之后新的队列（集合）中的元素就越多；在双向 BFS 算法中，如果我们每次都选择一个较小的集合进行扩散，那么占用的空间增长速度就会慢一些，效率就会高一些

```go
// ...
while (!q1.isEmpty() && !q2.isEmpty()) {
    if (q1.size() > q2.size()) {
        // 交换 q1 和 q2
        temp = q1;
        q1 = q2;
        q2 = temp;
    }
    // ...
```



## 雪花算法

- ID生成规则部分硬性要求
  - 全局唯一
    趋势递增
    单调递增
    信息安全
    含时间戳
- ID号生成系统的可用性要求
- 高可用：发一个获取分布式ID的请求，服务器就要保证99.999%的情况下给我创建一个唯一分布式ID
  低延迟：发一个获取分布式ID的请求，服务器就要快，极速
  高QPS：假如并发一口气10万个创建分布式ID请求同时杀过来，服务器要顶的住且一下子成功创建10万个分布式id





uuid完全可以保证唯一性，但是

为什么无序的UUID会导致入库性能变差呢?
1无序，无法预测他的生成顺序，不能生成递增有序的数字。
首先分布式id-般都会作为 主键，但是安装mysq|官方推荐主键要尽量越短越好，UUID每一一个 都很长，所以不是很推荐。
2主键，ID作为主键时在特定的环境会存在一些问题。
比如做DB主键的场景下，UUID就非常不适用MySQL官方有明确的建议主键要尽量越短越好36个字符长度的UUID不符合要求

3索引，B+树索引的分裂
既然分布式id是主键，然后主键是包含索引的，然后mysq|的索引是通过b+树来实现的，每一次新的UUID数据的插入， 为了查询的优化，都会对索引
底层的b+树进行修改，因为UUID数据 是无序的，所以每一次UUID数据的插入 都会对主键地城的b+树进行很大的修改，这一点很不好。 插入完全无
序，不但会导致一.些中间节点产生分裂，也会白白创造出很多不饱和的节点，这样大大降低了数据库插入的性能



- 数据库自增

  - 单机

  - 集群

    - 在分布式里面，数据库的自增ID机制的主要原理是:数据库自增ID和mysq|数据库的replace into实现的。
      这里的replace into跟insert功能类似，
      不同点在于: replace into首先尝试插入数据列表中，如果发现表中已经有此行数据(根据主键或唯一索引 判断)则先删除，再插入。
      否则直接插入新数据。
      REPLACE INTO的含义是插入一条记录，如果表中唯一索引的值遇到冲突，则替换老数据。

    - 可以保证唯一和递增，但仅适用于低并发

    - 那数据库自增ID机制适合作分布式ID吗?答案是不太适合
      1:系统水平扩展比较困难，比如定义好了步长和机器台数之后，如果要添加机器该怎么做?假设现在只有一台机器发号是1,2,3,4,5 (步长是1)，这
      个时候需要扩容机器一台。 可以这样做:把第二台机器的初始值设置得比第一- 台超过很多，貌似还好，现在想象-下如果我们线 上有100台机器，这
      个时候要扩容该怎么做?简直是噩梦。所以系统水平扩展方案复杂难以实现。
      2:数据库压力还是很大，每次获取ID都得读写一次 数据库，非常影响性能，不符合分布式ID里面的延迟低和要高QPS的规则(在高并发下，如果都
      去数据库里面获取id,那是非常影响性能的)

    - 基于Redis生成全局id策略

      - 单机：因为Redis是单线的天生保证原子性（redis6之后支持多线程了，这里说的是5及以下），可以使用原子操作，INCR和INCRBY来实现

      - 集群：注意:在Redis集群情况下，同样和MySQL--样需要设置不同的增长步长，同时key-定要设置有效期
        可以使用Redis集群来获取更高的吞吐量。
        假如一个集群中有5台Redis。可以初始化每台Redis的值分别是1,2,3,4,5,然后步长都是5。
        各个Redis生成的ID为: 
        A: 1,6,11,16,21
        B: 2,7,12,17,22
        C: 3,8,13,18,23
        D: 4,9,14,19,24
        E: 5,10,15,20,25

        维护麻烦，为了一个全局唯一id，要维护一个redis集群



而twitter的snowflake解决了这种需求，最初Twitter把存储系统从MySQL迁移到Cassandra(由Facebook开发一套 开源分布式NoSQL数据库系统)因
为Cassandra没有顺序ID生成机制，所以开发了这样一套全局唯一ID生成服务。
Twitter的分布式雪花算法SnowFlake，经测试snowflake每秒能够产生26万个自增可排序的ID
1、twitter的SnowFlake 生成ID能够按照时间有序生成
2、SnowFlake 算法生成id的结果是一个64bit大小的整数， 为一个Long型(转 换成字符串后长度最多19)。
3、分布式系统内不会产生ID碰撞(由datacenter和workerld作区分)并且效率较高。
分布式系统中，有一些需要使用全局唯一ID的场景， 生成ID的基本要求
1.在分布式的环境下必须全局且唯一. 。
2.一般都需要单调递增,因为一般唯一ID都会存到数据库,而Innodb的特性就是将内容存储在主键索引树上的叶子节点，而且是从左往右，递增的，所以考
虑到数据库性能，一般生成的id 也最好是单调递增。为了防止ID冲突可以使用36位的UUID，但是UUID有一-些缺点，首先他相对比较长，另外UUID
般是无序的

3.可能还会需要无规则，因为如果使用唯一ID作为订单号这种，为 了不然别人知道天的订单量是多少 ,就需要这个规则。
