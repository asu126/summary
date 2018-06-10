/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */
class Solution {
public:
    TreeNode* lowestCommonAncestor(TreeNode* root, TreeNode* p, TreeNode* q) {
        if(root==NULL || p == NULL || q == NULL){
            return NULL;
        }
        if(p->val > q->val){
            TreeNode* temp = p;
            p = q;
            q = temp;
        }
        if(root->val >= p->val && root->val <= q->val){
        	return root;
        }
        if(root->val < p->val && root->val < q->val){
        	return lowestCommonAncestor(root->right, p, q);
        }
        if(root->val > p->val && root->val > q->val){
        	return lowestCommonAncestor(root->left, p, q);
        }
    }
};
