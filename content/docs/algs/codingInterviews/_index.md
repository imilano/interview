---
weight: 5
---

这里是一些《剑指 Offer》的刷题题解。

## 链表

24. 反转链表

	TODO recheck
	注意这里递归方法的反转思路。

52. 两个链表的第一个公共节点
   
   TODO recheck
    
	注意条件

23. 链表中环的入口节点
   
   TODO recheck

   首先需要注意如何进行环检测，如何判断环终止。其次需要注意的是，如何找到环中的入口节点。

35. 复杂链表的复制

	链表的复制和图的复制都可以考虑 map。

1. 删除链表中的重复节点

	这里一个非常需要注意的点就是，如果头结点也要被删除的话怎么办？如果头结点的也要被删除的话，那么判断条件就会变得比较复杂，这里的一个解决办法就是，增加一个虚拟的头结点，然后让这个虚拟的头结点指向真正的头结点，遍历的时候就从这个虚拟的头结点开始遍历，这样就能够应付真实头结点也要被删除的情况。
	TODO rechck 增加虚拟头结点的技巧

1. 从尾到头打印链表
   简单题

1. 合并两个排序的链表
   简单题

22. 链表中倒数最后 k 个节点
	简单题

76. 删除链表汇中重复的节点
	TODO recheck

18. 删除链表的节点
   简单，注意虚拟头结点的使用

## 树
55. 二叉树的深度
    简单题

77. 按之字形顺序打印二叉树
    简单题

54. 二叉搜索树的第 k 个节点
    简单题

7. 重建二叉树
   TODO recheck

26. 树的子结构
    简单题

27. 二叉树的镜像
    简单题

32. 从上往下打印二叉树
    简单题

33. 二叉搜索树的后序遍历序列
    注意 corner case

82. 二叉树和为某一值的路径
	TODO 注意叶子节点的判断方式

34. 二叉树中和为某一值的路径 II
    TODO 注意叶子节点的判断方式

36. 二叉搜索树与双向链表
	TODO recheck
	这个题需要注意，尤其是怎么确定 head 节点，以及 pre 节点的选择，还有就是怎么将 pre 节点和当前节点做连接。题目本身其实并不算难，就是中序遍历，但是问题就是一些边界条件的考量，在这里尤其需要注意。

79. 判断是不是平衡二叉树
    TODO recheck
	这里虽然是个简单题，但是还是要注意一些，要快速做出来。

8. 二叉树的下一个节点
   TODO recheck
    这里需要好好注意一下问题是怎么分析的。

28. 对称的二叉树
    简单题
    
78. 把二叉树打印成多行
    简单题，其实就是一个二叉树的层次遍历。