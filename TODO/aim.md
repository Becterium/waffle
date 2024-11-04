### user

1.~~设置jwt,将用户信息序列化进token中~~  kratos的jwt已实现，github.com\go-kratos\kratos\v2@v2.2.0\middleware\auth\jwt\jwt.go有代码

```go
func Server(keyFunc jwt.Keyfunc, opts ...Option) middleware.Middleware {
 	...
    ctx = NewContext(ctx, tokenInfo.Claims)
    ...
}

// NewContext put auth info into context
func NewContext(ctx context.Context, info jwt.Claims) context.Context {
	return context.WithValue(ctx, authKey{}, info)
}

所以只要在context中找authKey就行
```



2.设置middleware，设置token认证范围，用Selector.Server包装go-kratos/jwt就行

