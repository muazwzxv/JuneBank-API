package middleware

// Converts from []byte to a json object according to the User struct.
// func toJson(val []byte)  {
//     user := User{}
//     err := json.Unmarshal(val, &user)
//     if err != nil {
//         panic(err)
//     }
//     return user
// }

// func veryfyCache(ctx *fiber.Ctx) error {
// 	id := ctx.Params("id")
// 	val, err := caching.GetRedisInstance().Cache.Get(caching.Ctx, id).Bytes()
// 	if err != nil {
// 		return ctx.Next()
// 	}
// 	data := toJs
// }
