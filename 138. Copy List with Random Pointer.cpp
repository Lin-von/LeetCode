/*
// Definition for a Node.
class Node {
public:
    int val;
    Node* next;
    Node* random;
    
    Node(int _val) {
        val = _val;
        next = NULL;
        random = NULL;
    }
};
*/

class Solution {
public:
    Node* copyRandomList(Node* head) {
        if (head == NULL)
            return head;

        Node* cur = head;
        while (cur)
        {
            Node* node = new Node(cur->val);
            Node* next = cur->next;
            node->next = next;
            cur->next = node;
            cur = next;
        }

        cur = head;
        while (cur)
        {
            cur->next->random = cur->random ? cur->random->next : NULL;
            cur = cur->next->next;
        }

        cur = head;
        Node* tmp = head->next;
        Node* ret = tmp;
        while (cur)
        {
            cur->next = cur->next->next;
            cur = cur->next;

            tmp->next = cur ? cur->next : NULL;
            tmp = tmp->next;
        }

        return ret;
    }
};