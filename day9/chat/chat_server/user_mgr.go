package main

var (
	mgr *UserMgr
)

func initUserMgr() {
	mgr = NewUserMgr(pool)
}
