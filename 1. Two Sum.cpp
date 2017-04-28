/*
最容易想到的是O(nlogn)的解法，遍历一遍数组做二分查找。
下面代码是个比较聪明的解法，复杂度只有O(n)，遍历过程中查询map中是否存在对应值，不存在则加入mao继续遍历
*/

class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        unordered_map<int ,int > m;
        for(int i=0;i<nums.size();i++)
        {
            int cur=nums[i];
            if(m.count(target-cur)>0) return vector<int > {m[target-cur],i};
            m[cur]=i;
            
        }
    }
};