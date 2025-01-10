/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode() {}
 *     ListNode(int val) { this.val = val; }
 *     ListNode(int val, ListNode next) { this.val = val; this.next = next; }
 * }
 */
class Solution {
    public ListNode deleteDuplicates(ListNode head) {
        if(head == null) return head;

        ListNode currNode = head;
        ListNode nextNode = head.next;

        while(nextNode != null) {
            if(currNode.val != nextNode.val) {
                currNode.next = nextNode;
                currNode = nextNode;
            } else {
                currNode.next = null;
            }

            nextNode = nextNode.next;
        }

        return head;
    }
}