/*
// Definition for a Node->
class Node {
public:
    int val;
    Node* left;
    Node* right;

    Node() {}

    Node(int _val) {
        val = _val;
        left = NULL;
        right = NULL;
    }

    Node(int _val, Node* _left, Node* _right) {
        val = _val;
        left = _left;
        right = _right;
    }
};
*/
class Solution {
public:
    Node* pre;
    void scanTree(Node* root) {
        if (!root) {
            return;
        }

        scanTree(root->left);
        if (pre) {
            pre->right = root;
            root->left = pre;
        }
        pre = root;
        scanTree(root->right);
    }

    Node* treeToDoublyList(Node* root) {
        if (root == NULL) {
            return root;
        }
        Node* l = root;
        while (l->left != NULL) {
            l = l->left;
        }
        pre = NULL;

        scanTree(root);
        l->left = pre;
        pre->right = l;
        return l;
    }
};