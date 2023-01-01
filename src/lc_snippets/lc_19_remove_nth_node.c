#include <stdlib.h>
struct ListNode {
   int val;
   struct ListNode *next;
};

struct ListNode* removeNthFromEnd(struct ListNode* head, int n){
    typedef struct ListNode node;
    if (n <= 0 || head == NULL) return head;
	node *prev = &(node){.next=head};
    node *first = head; 

    for (int i = n-1; i > 0; --i) {
        first = first->next;
    }

    while (first->next) {
        prev = prev->next;
        first = first->next;
    }

    if (prev->next == head) {
        return head->next;
    }
    prev->next = prev->next->next;
    return head;
}