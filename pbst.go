package main

type Node struct {
	Key         string
	Value       int
	Left        *Node
	Right       *Node
	LeftWeight  int
	RightWeight int
}

type Tree struct {
	Root *Node
}

type InvalidKeyError struct {
	TheKey string
}

func (theError InvalidKeyError) Error() string {
	return "Invalid Key: " + theError.TheKey
}

func (tree *Tree) put(newKey string, newValue int) {
	if tree.Root != nil {
		tree.Root.put(newKey, newValue)
	} else {
		tree.Root = &Node{Key: newKey, Value: newValue, Left: nil, Right: nil, LeftWeight: 0, RightWeight: 0}
	}
}

func (node *Node) put(newKey string, newValue int) {
	var normSide string
	var sideToPut string
	if newKey > node.Key {
		normSide = "Right"
	} else if newKey == node.Key {
		normSide = "Middle"
	} else {
		normSide = "Left"
	}
	if normSide == "Right" {
		if node.RightWeight > node.LeftWeight*2 {
			sideToPut = "Left"
		} else {
			sideToPut = "Right"
		}
	} else if normSide == "Left" {
		if node.LeftWeight > node.RightWeight*2 {
			sideToPut = "Right"
		} else {
			sideToPut = "Left"
		}
	} else {
		sideToPut = "Middle"
	}
	if sideToPut == "Right" {
		// TODO: This will incorrectly count weights, because if a put is done for an existing key, it won't increase the number of nodes
		node.RightWeight += 1
		if node.Right == nil {
			node.Right = &Node{Key: newKey, Value: newValue, Left: nil, Right: nil, LeftWeight: 0, RightWeight: 0}
		} else {
			node.Right.put(newKey, newValue)
		}
	} else if sideToPut == "Left" {
		node.LeftWeight += 1
		if node.Left == nil {
			node.Left = &Node{Key: newKey, Value: newValue, Left: nil, Right: nil, LeftWeight: 0, RightWeight: 0}
		} else {
			node.Left.put(newKey, newValue)
		}
	} else {
		node.Key, node.Value = newKey, newValue
	}
}

func (node *Node) get(theKey string) (int, error) {
	if node != nil {
		if node.Key == theKey {
			return node.Value, nil
		} else if theKey > node.Key {
			result, err := node.Right.get(theKey)
			if err != nil {
				return node.Left.get(theKey)
			} else {
				return result, err
			}
		} else {
			result, err := node.Left.get(theKey)
			if err != nil {
				return node.Right.get(theKey)
			} else {
				return result, err
			}
		}
	} else {
		return -1, InvalidKeyError{TheKey: theKey}
	}
}

func (tree *Tree) get(theKey string) (int, error) {
	if tree.Root != nil {
		return tree.Root.get(theKey)
	} else {
		return -1, InvalidKeyError{TheKey: theKey}
	}
}

func (node *Node) delete(theKey string, parent *Node) error {
	if theKey < node.Key {
		if node.Left == nil {
			return InvalidKeyError{TheKey: theKey}
		} else {
			return node.Left.delete(theKey, node)
		}
	} else if theKey > node.Key {
		if node.Right == nil {
			return InvalidKeyError{TheKey: theKey}
		} else {
			return node.Right.delete(theKey, node)
		}
	} else {
		if node.Left != nil && node.Right != nil {
			node.Key, node.Value = node.Right.minKeyVal()
			node.Right.delete(node.Key, node)
		} else if parent.Left == node {
			if node.Left != nil {
				parent.Left = node.Left
			} else {
				parent.Left = node.Right
			}
		} else if parent.Right == node {
			if node.Left != nil {
				parent.Right = node.Left
			} else {
				parent.Right = node.Right
			}
		}
	}
	return nil
}

func (node *Node) minKeyVal() (string, int) {
	if node.Left == nil && node.Right == nil {
		return node.Key, node.Value
	} else {
		minLeft, leftVal := node.Left.minKeyVal()
		minRight, rightVal := node.Right.minKeyVal()
		if minLeft < minRight {
			return minLeft, leftVal
		} else {
			return minRight, rightVal
		}

	}
}

func (tree *Tree) delete(theKey string) error {
	if tree.Root != nil {
		return tree.Root.delete(theKey, tree.Root)
	} else {
		return InvalidKeyError{TheKey: theKey}
	}
}

func main() {

}
