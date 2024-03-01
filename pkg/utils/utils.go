package utils

import (
	"sync"

	"github.com/toastsandwich/CRUD-echo/pkg/model"
)

var (
	Users = map[int]*model.User{}
	Lock  = sync.Mutex{}
	Seq   = 1
)