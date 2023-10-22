/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    if list1 == nil {
        return list2
    }

    if list2 == nil {
        return list1
    }


    var head *ListNode
    if list1.Val < list2.Val {
        head = list1
        list1 = list1.Next
    } else {
        head = list2
        list2 = list2.Next
    }

    currHead := head
    for list1 != nil && list2 != nil {
        if list1.Val < list2.Val {
            currHead.Next = list1
            list1 = list1.Next
        } else {
            currHead.Next = list2
            list2 = list2.Next
        }
        currHead = currHead.Next
    }


    if list1 == nil {
        currHead.Next = list2
    } 

    if list2 == nil {
        currHead.Next = list1
    } 

    return head
}   