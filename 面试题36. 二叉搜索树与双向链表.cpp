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
    void scanTree(Node* root) {
        if (root->left != NULL) {
            Node* tmp = root->left;
            while (tmp->right != NULL) {
                tmp = tmp->right;
            }

            scanTree(root->left);
            tmp->right = root;
            root->left = tmp;
        }

        if (root->right != NULL) {
            Node* tmp = root->right;
            while (tmp->left != NULL) {
                tmp = tmp->left;
            }

            scanTree(root->right);
            tmp->left = root;
            root->right = tmp;
        }
    }

    Node* treeToDoublyList(Node* root) {
        if (root == NULL) {
            return root;
        }
        Node* l = root;
        Node* r = root;
        while (l->left != NULL) {
            l = l->left;
        }
        while (r->right != NULL) {
            r = r->right;
        }
        scanTree(root);
        l->left = r;
        r->right = l;
        return l;
    }
};