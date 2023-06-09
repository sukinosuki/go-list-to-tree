# go-list-to-tree

A common generic function to convert a list to a tree struct list.

一个将列表转为树形结构的通用泛型方法

This library is base on generic, so go1.8 or later is needed.

该方法基于泛型, 需要go1.8及之后的版本.

## Install

```
go get -u github.com/sukinosuki/go-list-to-tree
```

## Usage. 用法

You need to let your struct implement three function declared in the `Tree` interface.

你需要让结构体实现 `Tree` 接口的三个方法

```golang
type Tree[T any] interface {
	GetId() uint
	
	GetPid() uint
	
	SetChildren(arr []T)
}
```

- GetId() uint
- GetPid() uint
- SetChildren(arr []T)

## Example. 例子

```go
package main

import (
	"fmt"
	"github.com/sukinosuki/go-list-to-tree"
)

// declare your own struct like that.
// 定义你自己的类型, 比如这样.
type Node struct {
	ID       uint
	PID      uint
	Children []*Node // Children item need to be declared as point type.(children元素需要定义为指针类型)
	Name     string
}

// please implement GetId, GetPid, SetChildren. functions.
// 需要实现 GetId, GetPid, SetChildren这三个方法
func (n *Node) GetId() uint {
	return n.ID
}

func (n *Node) GetPid() uint {
	return n.PID
}

func (n *Node) SetChildren(arr []*Node) {
	n.Children = arr
}

func main() {
	// The list you need convert to tree struct.
	// Point type (*Node) item is needed.
	// 指针类型的列表元素是必须的
	list := []*Node{
		{
			ID:   1,
			PID:  0,
			Name: "A",
		},
		{
			ID:   2,
			PID:  0,
			Name: "B", 
			Children: make([]*Node, 0), // If you do not want the empty children is nil as a result 
		},
		{
			ID:   3,
			PID:  1,
			Name: "A-1",
		},
		{
			ID:   4,
			PID:  1,
			Name: "A-2",
		},
		{
			ID:   5,
			PID:  3,
			Name: "A-1-1",
		},
		{
			ID:   6,
			PID:  3,
			Name: "A-1-2",
		},
		{
			ID:   7,
			PID:  100, // not a node's id is 100, so this node will be ignored in result tree.(
			            // 没有id=100的记录, 所以这条记录会被忽略
			Name: "A-1-2",
		},
	}

	// Call `ListToTree` function, it will return a tree struct as result. 
	// 调用ListToTree方法, 返回得到树形结构
	result := tree.ListToTree(list)

	// you should check if the result are correct.
	// 你需要自己检查结果是否正确
	fmt.Println("result ", result)
}

```

#### result

```json
[
  {
    "ID":1,
    "PID":0,
    "Children":[
      {
        "ID":3,
        "PID":1,
        "Children":[
          {
            "ID":5,
            "PID":3,
            "Children":null,
            "Name":"A-1-1"
          },
          {
            "ID":6,
            "PID":3,
            "Children":null,
            "Name":"A-1-2"
          }
        ],
        "Name":"A-1"
      },
      {
        "ID":4,
        "PID":1,
        "Children":null,
        "Name":"A-2"
      }
    ],
    "Name":"A"
  },
  {
    "ID":2,
    "PID":0,
    "Children":[],
    "Name":"B"
  }
]
```

## Notice

I'm not a perfect developer. please do not use the library at your main project. :(

不保证代码的可用性, 请不要用在重要的项目上. :( 