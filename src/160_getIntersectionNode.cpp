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
    ListNode *getIntersectionNode(ListNode *headA, ListNode *headB) {
        int lengthA = getListLength(headA);
        int lengthB = getListLength(headB);
        int n = lengthA -lengthB;
        ListNode *longPtr = headA;
        ListNode *shortPtr = headB;
        
        if(lengthA < lengthB) {
            n = lengthB -lengthA;
            ListNode *longPtr = headB;
            ListNode *shortPtr = headA;
        }
        
        for(int i = 0; i<n; i++){
            if(longPtr!=nullptr){
              longPtr = longPtr->next;  
            }
        }
        
        while(longPtr!=nullptr && shortPtr != nullptr){
            if(longPtr == shortPtr){
                return shortPtr;
            }
            longPtr = longPtr->next;
            shortPtr = shortPtr->next;
        }
        return nullptr;
    }
    
    int getListLength(ListNode *head){
        int result =0;
        ListNode *nextP = head;
        
        while(nextP!=nullptr){
            result++;
            nextP = nextP->next;
        }
        return result;
    }
};
