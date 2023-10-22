/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    head := &ListNode{}
    currHead := head

    for list1 != nil && list2 != nil {
        if list1.Val < list2.Val {
            currHead.Next = &ListNode{Val: list1.Val}
            list1 = list1.Next
        } else {
            currHead.Next = &ListNode{Val: list2.Val}
            list2 = list2.Next
        }
        currHead = currHead.Next
    }


    if list1 == nil {
        currHead.Next = list2
    } else {
        currHead.Next = list1
    }

    return head.Next
}   