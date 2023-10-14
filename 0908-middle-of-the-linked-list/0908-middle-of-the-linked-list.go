/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    size := 0

    if head.Next == nil {
        return head
    }

    curr := head
    size++
    for curr.Next != nil {
        curr = curr.Next
        size++
    }

    midd := size / 2
    for i := 0; i < midd; i++ {
        head = head.Next
    }

    return head
}