// Package cache currently contains the main caching variable. In a real
// production environment, this would hold interfaces and methods related to
// interacting with a memory store cache such as Redis.
package cache

import "github.com/nkarpenko/koho-transaction/common/model"

// Cache var is the main cache variable used for storing transaction data. In
// a real production scenario, we would use some memory store caching mechanism
// such as Redis/Memcache or nosql/sql solution. Please review root directory
// README.md file for more details.
var Cache = &map[int][]model.Result{}
