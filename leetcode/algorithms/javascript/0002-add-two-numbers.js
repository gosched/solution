/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */
/**
 * @param {ListNode} l1
 * @param {ListNode} l2
 * @return {ListNode}
 */
var addTwoNumbers = function (l1, l2) {
    let dummy = new ListNode(0);
    let mark = dummy;
    let carry = 0;
    while (l1 || l2) {
        if (l1) {
            carry += l1.val;
            l1 = l1.next;
        }
        if (l2) {
            carry += l2.val;
            l2 = l2.next;
        }
        mark.next = new ListNode(carry % 10);
        mark = mark.next;
        carry = Math.floor(carry / 10);
    }
    if (carry) {
        mark.next = new ListNode(1);
    }
    return dummy.next;
};