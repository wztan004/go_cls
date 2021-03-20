# Move to the last node
## Method 1 (Preferred due to simplicity)
for currentNode.next != nil {
	currentNode = currentNode.next
}

## Method 2
for i := 0; i < p.size - 1; i++ {
	currentNode = currentNode.next
}

# Adding and removing
for i := 0; i < index; i++ {
	prevNode = currentNode
	currentNode = currentNode.next
}


// seems to be an issue with get() index and addAtPos index. It seems to be 1-based index instead of 0-based inedx
// I've edited official solution in get() and addAtPos
// If you're using this for development, thoroughly test these first! E.g. negative int as argument for certain methods.

// If you want to have additional features in the code (contains, swap, prepend, indexOf, isEmpty, Clear, Sort)
// https://github.com/emirpasic/gods/blob/master/lists/singlylinkedlist/singlylinkedlist.go#L25