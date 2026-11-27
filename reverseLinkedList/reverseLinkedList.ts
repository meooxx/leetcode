import { ListNode } from '../types.ts'

/**
 * Definition for singly-linked list.
 * class ListNode {
 *     val: number
 *     next: ListNode | null
 *     constructor(val?: number, next?: ListNode | null) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.next = (next===undefined ? null : next)
 *     }
 * }
 */

function reverseList(head: ListNode | null): ListNode | null {
    if(head == null) return head
    let pre = head
    let cur = pre.next
    while (cur != null) {
        const next = cur.next
        cur.next = head
        pre.next = next
        head = cur
        cur = next
    }
    return head
};