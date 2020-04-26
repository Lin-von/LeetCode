package main

type userNode struct {
	tw []int
	fl map[int]bool
}

type Twitter struct {
	user map[int]userNode
	time int
	tweets map[int]int
}


/** Initialize your data structure here. */
func Constructor() Twitter {
	return Twitter{
		map[int]userNode{},
		0,
		map[int]int{},
	}
}


/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int)  {
	this.tweets[tweetId] = this.time
	this.time ++
	tmp, ok := this.user[userId]
	if !ok {
		tmp = userNode{
			[]int{},
			map[int]bool{},
		}
	}
	tmp.tw = append([]int{tweetId}, tmp.tw...)
	if len(tmp.tw) > 10 {
		tmp.tw = tmp.tw[:len(tmp.tw) - 1]
	}
	this.user[userId] = tmp
}


/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
	ret := make([]int, len(this.user[userId].tw))
	copy(ret, this.user[userId].tw)
	for fuser, _ := range this.user[userId].fl {
		if fuser == userId {
			continue
		}
		tmp := []int{}
		cur := this.user[fuser].tw
		p1 := 0
		p2 := 0
		for p1 < len(ret) && p2 < len(cur) {
			if this.tweets[ret[p1]] > this.tweets[cur[p2]] {
				tmp = append(tmp, ret[p1])
				p1 ++
			} else {
				tmp = append(tmp, cur[p2])
				p2 ++
			}
			if len(tmp) == 10 {
				break
			}
		}
		for p1 < len(ret) && len(tmp) < 10 {
			tmp = append(tmp, ret[p1])
			p1 ++
		}
		for p2 < len(cur) && len(tmp) < 10 {
			tmp = append(tmp, cur[p2])
			p2 ++
		}
		ret = tmp
	}
	return ret
}


/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int)  {
	tmp, ok := this.user[followerId]
	if !ok {
		tmp = userNode{
			[]int{},
			map[int]bool{},
		}
	}
	tmp.fl[followeeId] = true
	this.user[followerId] = tmp
}


/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int)  {
	tmp, ok := this.user[followerId]
	if !ok {
		tmp = userNode{
			[]int{},
			map[int]bool{},
		}
	}
	delete(tmp.fl, followeeId)
	this.user[followerId] = tmp
}


/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
