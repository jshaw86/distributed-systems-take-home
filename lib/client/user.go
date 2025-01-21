package client

import (
    "math/rand"
)

type UserInfo struct {
    ID string
    UserAgent string
    IPAddr string
    BadActor int 
}

type UserIds struct {
    userIDs []UserInfo
}

const (
    maxUsers = 20 
)

func (u *UserIds) CreateUser() UserInfo {
    user := UserInfo{RandomUserID(), RandomUserAgent(), RandomIP(), RandomBadActor()}
    u.userIDs = append(u.userIDs, user)
    return user 
}

func (u *UserIds) CreateUserOrUseExisting() UserInfo {
    if (len(u.userIDs) == 0 || DecisionBasedOnProbability(20)) && len(u.userIDs) < maxUsers{
        return u.CreateUser()
    }

    return u.userIDs[rand.Intn(len(u.userIDs))]
}
