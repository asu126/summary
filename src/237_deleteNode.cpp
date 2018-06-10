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
    void deleteNode(ListNode* node) {
        ListNode* nex = node->next;
        while(nex!=NULL){
            node->val = nex->val;
            if(nex->next == NULL){
                node->next = NULL;
                return;
            }
            node = nex;
            nex = nex->next;
        }
    }
};
