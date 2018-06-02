/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
class Solution {
public:
    bool hasCycle(ListNode *head) {
        if(head == nullptr){
            return false;
        }
        ListNode *pNext = head->next;
        if(pNext == nullptr){
            return false;
        }
        ListNode *pNext2 = pNext->next;
        
        while(pNext != nullptr && pNext2 != nullptr){
            if(pNext == pNext2){
                return true;
            }
            
            pNext = pNext->next;
            
            pNext2 = pNext2->next;
            if(pNext2 != nullptr) {
                pNext2 = pNext2->next;
            }
        }
        return false;
    }
};
