/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
// #include <iostream>
// struct ListNode {
//     int val;
//     ListNode *next;
//     ListNode(int x) : val(x), next(NULL) {}
// };

class Solution {
public:
    ListNode *getIntersectionNode(ListNode *headA, ListNode *headB) {
        int lengthA = getListLength(headA);
        int lengthB = getListLength(headB);
       // std::cout << lengthA << " " <<lengthB << std::endl;
        int n = lengthA -lengthB;
        ListNode *longPtr = headA;
        ListNode *shortPtr = headB;
        
        if(lengthA < lengthB) {
            n = lengthB -lengthA;
            longPtr = headB;
            shortPtr = headA;
        }
        
        for(int i = 0; i<n; i++){
            if(longPtr!=NULL){
              longPtr = longPtr->next;  
            }
        }
        
        while(longPtr!=NULL && shortPtr != NULL){
            if(longPtr == shortPtr){
                return shortPtr;
            }
            longPtr = longPtr->next;
            shortPtr = shortPtr->next;
        }
        return NULL;
    }
    
    int getListLength(ListNode *head){
        int result =0;
        ListNode *nextP = head;
        
        while(nextP!=NULL){
            result++;
            nextP = nextP->next;
        }
        return result;
    }
};
