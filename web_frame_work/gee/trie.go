package gee

import "strings"

type node struct {
	pattern  string  // 待匹配路由，例如 /p/:lang
	part     string  // 路由中的一部分，例如 :lang
	children []*node // 子节点，例如 [doc, tutorial, intro]
	isWild   bool    // 是否精确匹配，part 含有 : 或 * 时为true
}

//插入节点
// 1. /p/a/doc  /p/a/cd
//以上两个节点组成的 前缀树
// 		  p
//		/
//     a
//    /  \
//  doc  cd
// 前缀树的每个节点都是node.
// 以p节点为例 pattern 为空 (因为这个pattern 是用来作为key的,value的值为函数.末尾节点才用,如 p/a/doc 只有doc节点pattern为p/a/doc ,路过的p和a全为空)
// 所以插入URL: p/a/cd 的方式就是在p节点搜索其子节点,看看是否有和 URL 二号位(/a/) 相同的
// 有 p的子节点a
// 迭代搜索 a的 子节点是否有和 URL 三号位(/cd/)
//没有 插入节点,为这个节点插入pattern
/* pattern 整体的url   parts 将pattern 以 / 为分界切分,分成的数组, height遍历到前缀树的第n层 OR URL的第n个位置*/
func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}
	// 遍历到URL的第height个位置
	part := parts[height]
	// 匹配子节点
	child := n.matchChild(part)
	// 没有插入子节点
	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	// 进行递归
	child.insert(pattern, parts, height+1)

}

// 第一个匹配成功的节点，用于插入
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {

		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

//查询节点
// 注意 查询  返回的实质上是一个key,用来从map中找出执行的function
// 当然

func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)

	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}

	return nil
}

// 查找匹配成功的子节点
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}
